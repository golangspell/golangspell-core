package main

import (
	"fmt"

	"github.com/golangspell/golangspell-core/cmd"
	_ "github.com/golangspell/golangspell-core/config"
	_ "github.com/golangspell/golangspell-core/gateway/template"
	_ "github.com/golangspell/golangspell/gateway/filesystem"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Printf("An error occurred while executing the command. Message: %s\n", err.Error())
	}
}
