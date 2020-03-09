package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/danilovalente/golangspell-core/usecase"
	"github.com/spf13/cobra"
)

func init() {
	RunCommandFunctions["init"] = runInitCommand
}

//Module name to initialize with 'Go Modules'
var Module string

//AppName used to define the application's directory and the default value to the config variable with the same name
var AppName string

func runInitCommand(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		fmt.Println(`The command init requires exactly two parameters: ModuleName and AppName
Args:
ModuleName: New Module's name (required). Example: github.com/your-git-account/my-new-application
AppName: New App's name (required). Example: my-new-application
		
Syntax: 
golangspell init [ModuleName] [AppName]`)
		return
	}

	err := usecase.RenderInitTemplate(args)
	if err != nil {
		fmt.Printf("An error occurred while trying to create the application. Error: %s\n", err.Error())
		return
	}
	execCmd := exec.Command("go", "mod", "init", args[0])
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr
	err = execCmd.Run()
	if err != nil {
		fmt.Printf("An error occurred while trying to init the mudule. Error: %s\n", err.Error())
		return
	}
	fmt.Println("Application created!")
	fmt.Println("---------------------------------------------------------")
	fmt.Println("To run your application, first build it with the command:")
	fmt.Println("go build")
	fmt.Println("---------------------------------------------------------")
	fmt.Println("Then execute the application with the command (Unix based systems):")
	fmt.Printf("./%s\n", args[1])
	fmt.Println("Then execute the application with the command (Windows based systems):")
	fmt.Printf("%s\n", args[1])
	fmt.Println("---------------------------------------------------------")
	fmt.Println("Available endpoints:")
	fmt.Printf("http://localhost:8080/%s/v1/info\n", args[1])
	fmt.Printf("http://localhost:8080/%s/v1/health\n", args[1])
	fmt.Println("---------------------------------------------------------")
	fmt.Println("Find more details on README.md file")
}
