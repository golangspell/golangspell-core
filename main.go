package main

import (
	"github.com/danilovalente/golangspell-core/cmd"
	_ "github.com/danilovalente/golangspell-core/config"
	_ "github.com/danilovalente/golangspell-core/gateway/template"
)

func main() {
	cmd.Execute()
}
