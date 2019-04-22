package handle

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Handle create services.cfg
func Handle(info ...string) {
	dir := info[0]
	if dir[len(dir)-1] != '/' {
		dir += "/"
	}

	buf, err := ioutil.ReadFile(info[2])
	if err != nil {
		fmt.Println("Error: need", info[2])
		os.Exit(0)
	}

	err = ioutil.WriteFile(info[1], stream(buf, "_", dir), 0666)
	if err != nil {
		fmt.Println("Error: write error", err)
		os.Exit(0)
	}
}
