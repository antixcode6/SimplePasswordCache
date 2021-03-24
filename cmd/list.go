package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"syscall"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Print the initialized password cache",
	Long:  `Print the initialized password cache directory structure using tree`,
	Run: func(cmd *cobra.Command, args []string) {
		//";", "tree -C -l --noreport $HOME/.passwordcache", "| tail -n +2")
		binary, lookErr := exec.LookPath("tree")
		if lookErr != nil {
			log.Fatal(lookErr)
		}

		user, err := user.Current()
		if err != nil {
			panic(err)
		}
		homeDir := user.HomeDir

		cachedir := fmt.Sprintf("%s/.passwordcache", homeDir)
		args = []string{"-C", "-l", "--noreport", cachedir}

		env := os.Environ()

		execErr := syscall.Exec(binary, args, env)
		if execErr != nil {
			log.Fatal(execErr)
		}
	},
}
