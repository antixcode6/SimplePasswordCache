package init

import (
	"fmt"
	"os/user"
)

//Exported Function Init
func Init(fileName string) {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	homeDir := user.HomeDir
	defaultDir := fmt.Sprintf("%s/.passwordcache", homeDir)
	fmt.Printf(defaultDir)
}
