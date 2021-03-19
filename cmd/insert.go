package cmd

import (
	"fmt"
	"io"
	"os"
	"os/user"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"
)

func init() {
	rootCmd.AddCommand(insertCmd)
}

var insertCmd = &cobra.Command{
	Use:   "insert [pass-name]",
	Short: "Add a password to the cache",
	Long:  `This command adds a password to the cache. You can specify -n or --no-echo disables keybord echo of password`,
	Run: func(cmd *cobra.Command, args []string) {
		//naming convention is weird, but this is the file created out of the arg user specifies
		//it is not trying to encrypt an existing file.
		fileToEnc := args[0]

		user, err := user.Current()
		if err != nil {
			panic(err)
		}
		homeDir := user.HomeDir
		cachedFilePath := fmt.Sprintf("%s/.pernia/.cachedFile.asc", homeDir)

		// Read in public key
		recipient, err := readEntity(cachedFilePath)
		if err != nil {
			panic(err)
		}

		//	f, err := os.Open(fileToEnc)
		//	if err != nil {
		//		fmt.Println(err)
		//		return
		//	}
		//	defer f.Close()

		destination := fmt.Sprintf("%s/.passwordcache/%s", homeDir, fileToEnc)
		dst, err := os.Create(destination + ".gpg")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer dst.Close()
		encrypt([]*openpgp.Entity{recipient}, nil, nil, dst)

	},
}

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
