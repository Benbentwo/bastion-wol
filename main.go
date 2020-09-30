package main

import (
	"github.com/Benbentwo/bastion-wol/app"
	"os"
)

func main() {
	if err := app.Run(nil); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
