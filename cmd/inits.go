package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(InitsCmd)
}

//Exported var InitsCmd is the command to run by cobra
var InitsCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a new password store",
	Long:  `Initialize a new password store and use a asc file, or two for encryption. This command must be ran first`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Must pass an asc file to this command. To generate file run gpg2 --export -a 'key uid' > <path to pubkey.asc> and use the uid of the key you wish to encrypt")
			os.Exit(2)
		}

		user, err := user.Current()
		if err != nil {
			panic(err)
		}
		homeDir := user.HomeDir

		defaultDir := fmt.Sprintf("%s/.passwordcache", homeDir)
		cachedFile := fmt.Sprintf("%s/.pernia", homeDir)
		cachedIndicator := fmt.Sprintf("%s", args[0])
		os.Mkdir(defaultDir, 0755)
		os.Mkdir(cachedFile, 0755)
		cachedFilePath := fmt.Sprintf("%s/.pernia/cachedFile", homeDir)
		os.Create(cachedFilePath)

		write1 := []byte(cachedIndicator)
		err = ioutil.WriteFile(cachedFilePath, write1, 0755)
		if err != nil {
			panic(err)
		}
		os.Mkdir(defaultDir, 0755)

		userFeelGoodMsg := fmt.Sprintf("Initalized new dir at: %s", defaultDir)
		if _, err := os.Stat(defaultDir); os.IsNotExist(err) {
			log.Fatal(err)
			os.Exit(2)
		}

		fmt.Println(userFeelGoodMsg)
	},
}
