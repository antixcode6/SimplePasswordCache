package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/user"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"
	"golang.org/x/crypto/ssh/terminal"
)

func init() {
	rootCmd.AddCommand(insertCmd)
}

var insertCmd = &cobra.Command{
	Use:   "insert [pass-name]",
	Short: "Add a password to the cache",
	Long:  `This command adds a password to the cache. You can specify -n or --no-echo disables keybord echo of password`,
	Run: func(cmd *cobra.Command, args []string) {
		user, err := user.Current()
		homeDir := user.HomeDir
		pubkeyPath := fmt.Sprintf("%s/.pernia/cachedFile", homeDir)

		dat, err := ioutil.ReadFile(pubkeyPath)
		pubkey_str := fmt.Sprintf("%s", dat)
		fmt.Println(pubkey_str)
		//naming convention is weird, but this is the file created out of the arg user specifies
		//it is not trying to encrypt an existing file.
		fileToEnc := args[0]

		if err != nil {
			panic(err)
		}
		//cachedFilePath := fmt.Sprintf("%s/.pernia/.cachedFile.asc", homeDir)

		// Read in public key
		recipient, err := readEntity(pubkey_str)
		if err != nil {
			panic(err)
		}

		destination := fmt.Sprintf("%s/.passwordcache/%s", homeDir, fileToEnc)
		os.Create(destination)
		//Create File with name of insert param
		f, err := os.Open(destination)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()

		//Get User Input
		fmt.Printf("Enter password for %s:", args[0])
		userpwd, err := terminal.ReadPassword(0)
		if err != nil {
			panic(err)
		}

		write1 := []byte(userpwd)
		err = ioutil.WriteFile(destination, write1, 0755)
		if err != nil {
			panic(err)
		}

		dst, err := os.Create(destination + ".gpg")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer dst.Close()
		os.Remove(destination)
		encrypt([]*openpgp.Entity{recipient}, nil, f, dst)

	},
}

//WHEN BACK WRITE CODE TO WRITE TO GPG FILE THEN ENCRYPT
func encrypt(recip []*openpgp.Entity, signer *openpgp.Entity, r io.Reader, w io.Writer) error {
	wc, err := openpgp.Encrypt(w, recip, signer, &openpgp.FileHints{IsBinary: true}, nil)
	if err != nil {
		return err
	}
	if _, err := io.Copy(wc, r); err != nil {
		return err
	}
	return wc.Close()
}

func readEntity(name string) (*openpgp.Entity, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	block, err := armor.Decode(f)
	if err != nil {
		return nil, err
	}
	return openpgp.ReadEntity(packet.NewReader(block.Body))
}
