package main

import (
	"fmt"
	"github.com/Tsotsi/go-cli-template/pkg/commands"
)

func main() {
	if err := commands.Exec(); err != nil {
		fmt.Printf("err: %s\n", err)
	}
}
