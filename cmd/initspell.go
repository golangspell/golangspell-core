package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/danilovalente/golangspell-core/usecase"
	"github.com/spf13/cobra"
)

func init() {
	RunCommandFunctions["initspell"] = runInitSpellCommand
}

func runInitSpellCommand(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		fmt.Println(`The command initspell requires exactly two parameters: ModuleName and NewSpellName
Args:
ModuleName: New Module's name (required). Example: github.com/your-git-account/my-new-spell
NewSpellName: New Spell's name (required). Example: my-new-spell
		
Syntax: 
golangspell initspell [ModuleName] [NewSpellName]`)
		return
	}

	err := usecase.RenderInitSpellTemplate(args)
	if err != nil {
		fmt.Printf("An error occurred while trying to create the spell. Error: %s\n", err.Error())
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
	fmt.Println("Spell created!")
	fmt.Println("---------------------------------------------------------")
	fmt.Println("To use your spell, first build it with the command:")
	fmt.Println("go build")
	fmt.Println("---------------------------------------------------------")
	fmt.Println("Then commit and push your spell to a Git repository")
	fmt.Println("Tag your Git repository: Ex:")
	fmt.Println("git tag -a v0.0.1 -m \"My First Spell version\"")
	fmt.Println("git push origin v0.0.1")
	fmt.Println("---------------------------------------------------------")
	fmt.Println("How to add your new Spell to Golangspell:")
	fmt.Printf("golangspell addspell https://%s %s", args[0], args[1])
	fmt.Println("Remark: If your Git repository is private, before adding you need to define the env variable GOPRIVATE=[Your module path]")
	fmt.Println("---------------------------------------------------------")
	fmt.Println("How to run your new Spell (you have an example command ready):")
	fmt.Printf("golangspell %s-hello [name]\n", args[1])
}
