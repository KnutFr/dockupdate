package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) == 1 {
		usage()
	} else {
		cli := initLocalDocker()
		getContainer(cli)
		hub := initRegistryDocker()
		fmt.Println("Hub version ", hub)
	}
	os.Exit(0)
}

func usage() {

}
