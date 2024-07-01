package cmd

import (
	"fmt"

	"github.com/golangspell/golangspell-core/usecase"
	"github.com/spf13/cobra"
)

func init() {
	RunCommandFunctions["addgateway"] = runaddgateway
}

func runaddgateway(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		fmt.Println(`The command addgateway requires 2 arguments
Args:
gatewayPackage: The name of the package that will contain the new gateway. Example: gateway-package
gatewayName: The name of the new gateway. Example: MyNewGateway

Syntax: 
golangspell addgateway [gateway-package] [GatewayName]

Examples:
# Adds a new gateway with the name "MyNewGateway" to the package "my-db" of the current Golangspell application
golangspell addgateway my-db MyNewGateway`)
		return
	}

	//Here your template, hosted on the folder "templates" is rendered
	err := usecase.RenderaddgatewayTemplate(args)
	if err != nil {
		fmt.Printf("An error occurred while trying to render the template. Error: %s\n", err.Error())
		return
	}
	//TODO: Create your additional logic here
	fmt.Println("addgateway executed!")
}
