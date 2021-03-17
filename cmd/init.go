package init

import (
	"fmt"
	"os"
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
	os.MkdirAll(defaultDir, 0644)
	userFeelGoodMsg := fmt.Sprintf("Initalized new dir at: %s", defaultDir)
	fmt.Println(userFeelGoodMsg)
}
