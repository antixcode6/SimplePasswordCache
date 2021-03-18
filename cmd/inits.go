package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(InitsCmd)
}

var InitsCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a new password store",
	Long:  `Initialize a new password store and use a gpg-id for encryption. This command must be ran first`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Must pass gpg key to this command")
			os.Exit(2)
		}

		user, err := user.Current()
		if err != nil {
			panic(err)
		}
		homeDir := user.HomeDir

		defaultDir := fmt.Sprintf("%s/.passwordcache", homeDir)
		now := time.Now()
		sec := now.Unix()
		cachedFile := fmt.Sprintf("%s/.%dpernia", homeDir, sec)
		cachedIndicator := fmt.Sprintf("%s:%s", defaultDir, args[0])
		os.Mkdir(defaultDir, 0755)
		os.Mkdir(cachedFile, 0755)
		cachedFilePath := fmt.Sprintf("%s/.%dpernia/cachedFile", homeDir, sec)
		os.Create(cachedFilePath)

		write1 := []byte(cachedIndicator)
		err = ioutil.WriteFile(cachedFilePath, write1, 0755)
		if err != nil {
			panic(err)
		}

		userFeelGoodMsg := fmt.Sprintf("Initalized new dir at: %s", defaultDir)
		if _, err := os.Stat(defaultDir); os.IsNotExist(err) {
			log.Fatal(err)
			os.Exit(2)
		}

		fmt.Println(userFeelGoodMsg)
	},
}
