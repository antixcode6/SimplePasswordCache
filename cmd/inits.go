package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(InitsCmd)
}

var InitsCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a new password store",
	Long:  `Initialize a new password store and use a gpg-id, or two for encryption. This command must be ran first`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Must pass gpg key UID to this command: run gpg2 -k and use the uid of the key you wish to encrypt")
			os.Exit(2)
		}

		user, err := user.Current()
		if err != nil {
			panic(err)
		}
		homeDir := user.HomeDir

		defaultDir := fmt.Sprintf("%s/.passwordcache", homeDir)

		os.Mkdir(defaultDir, 0755)
		pgpcmduid := fmt.Sprintf("--export -a %s > /tmp/pubKey.asc", args[0])
		fmt.Println(pgpcmduid)

		cmdOutput, err := exec.Command("gpg2", pgpcmduid).Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", cmdOutput)

		userFeelGoodMsg := fmt.Sprintf("Initalized new dir at: %s", defaultDir)
		if _, err := os.Stat(defaultDir); os.IsNotExist(err) {
			log.Fatal(err)
			os.Exit(2)
		}

		fmt.Println(userFeelGoodMsg)
	},
}
