package main

import (
	"fmt"

	"github.com/danilovalente/golangspell-core/cmd"
	_ "github.com/danilovalente/golangspell-core/config"
	_ "github.com/danilovalente/golangspell-core/gateway/template"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Printf("An error occurred while executing the command. Message: %s\n", err.Error())
	}
}
