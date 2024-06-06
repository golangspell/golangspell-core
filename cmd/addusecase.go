package cmd

import (
	"fmt"

	"github.com/golangspell/golangspell-core/usecase"
	"github.com/spf13/cobra"
)

func init() {
	RunCommandFunctions["addusecase"] = runaddusecase
}

func runaddusecase(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println(`The command addusecase requires 1 argument
Args:
usecaseName: The name of the new usecase. Example: MyNewUsecase

Syntax: 
golangspell addusecase [UsecaseName]

Examples:
# Adds a new usecase to the current Golangspell with the name "MyNewUsecase"
golangspell addusecase MyNewUsecase`)
		return
	}

	//Here your template, hosted on the folder "templates" is rendered
	err := usecase.RenderaddusecaseTemplate(args)
	if err != nil {
		fmt.Printf("An error occurred while trying to render the template. Error: %s\n", err.Error())
		return
	}
	//TODO: Create your additional logic here
	fmt.Println("addusecase executed!")
}
