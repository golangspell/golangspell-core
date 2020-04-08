package usecase

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/danilovalente/golangspell-core/appcontext"
	"github.com/danilovalente/golangspell-core/domain"
	toolconfig "github.com/danilovalente/golangspell/config"
	tooldomain "github.com/danilovalente/golangspell/domain"
)

func renameTemplateFileNames(currentPath string, newSpellCommandName string) error {
	sourcePath := fmt.Sprintf("%s%scmd%snewspellcommand.go", currentPath, toolconfig.PlatformSeparator, toolconfig.PlatformSeparator)
	directory := filepath.Dir(sourcePath)
	destinationPath := fmt.Sprintf("%s%s%s.go", directory, toolconfig.PlatformSeparator, newSpellCommandName)

	err := os.Rename(sourcePath, destinationPath)
	if err != nil {
		return err
	}

	sourcePath = fmt.Sprintf("%s%susecase%snewspellusecase.go", currentPath, toolconfig.PlatformSeparator, toolconfig.PlatformSeparator)
	directory = filepath.Dir(sourcePath)
	destinationPath = fmt.Sprintf("%s%srender%stemplate.go", directory, toolconfig.PlatformSeparator, newSpellCommandName)

	return os.Rename(sourcePath, destinationPath)
}

func addCommandConfig(code string, args []string) string {
	commandName := args[0]
	var validArgs string
	if len(args) > 1 {
		commandArgs := make([]string, len(args)-1)
		for i := 1; i < len(args)-1; i++ {
			commandArgs[i-1] = args[i]
		}

		argsContent, _ := json.Marshal(commandArgs)
		argstext := string(argsContent)
		argstext = strings.ReplaceAll(strings.ReplaceAll(argstext, "[", "{"), "]", "}")
		validArgs = fmt.Sprintf("ValidArgs: []string%s,", argstext)
	} else {
		validArgs = ""
	}

	r, _ := regexp.Compile("},\n.*}\n.*}")

	code = r.ReplaceAllString(code, fmt.Sprintf(`	"%s": &domain.Command{
				Name:             "%s",
				ShortDescription: "The %s [TODO: PUT HERE THE NEW COMMAND FEATURE]",
				LongDescription: The %s [TODO: PUT HERE THE NEW COMMAND FEATURE EXTENDED DESCRIPTION]
Args:
[TODO: PUT HERE THE NEW COMMAND ARGS DESCRIPTION]

Syntax: 
golangspell [TODO: PUT HERE THE NEW COMMAND SYNTAX]

Examples:
[TODO: PUT HERE THE NEW COMMAND EXAMPLES IF NEEDED],
		%s
	},
},
}
}`, commandName, commandName, commandName, commandName, validArgs))
	return code
}

func addCommandToBuildConfigCommand(currentPath string, args []string) error {
	filePath := fmt.Sprintf("%s%scmd%sbuildconfig.go", currentPath, toolconfig.PlatformSeparator, toolconfig.PlatformSeparator)
	renderer := domain.GetRenderer()
	err := renderer.BackupExistingCode(filePath)
	if err != nil {
		fmt.Printf("An error occurred while trying to create the new spell command. Error: %s\n", err.Error())
		return err
	}
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("An error occurred while trying to create the new spell command. Error: %s\n", err.Error())
		return err
	}
	code := addCommandConfig(string(content), args)
	err = ioutil.WriteFile(filePath, []byte(code), 0644)
	if err != nil {
		fmt.Printf("An error occurred while trying to create the new spell command. Error: %s\n", err.Error())
		return err
	}

	return nil
}

func getModuleName(currentPath string) string {
	filePath := fmt.Sprintf("%s%sgo.mod", currentPath, toolconfig.PlatformSeparator)
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("An error occurred while trying to create the new spell command. Error: %s\n", err.Error())
		return ""
	}
	contentText := string(content)
	re := regexp.MustCompile("module (.*?)\n")
	match := re.FindStringSubmatch(contentText)
	if len(match) >= 2 {
		return strings.Trim(match[1], " ")
	}
	return ""
}

//RenderNewSpellCommandTemplate renders the templates defined to the addspellcommand command with the proper variables
func RenderNewSpellCommandTemplate(args []string) error {
	spell := appcontext.Current.Get(appcontext.Spell).(tooldomain.Spell)
	renderer := domain.GetRenderer()
	newSpellCommandName := args[0]
	currentPath, err := os.Getwd()
	if err != nil {
		fmt.Printf("An error occurred while trying to create the new spell command. Error: %s\n", err.Error())
		return err
	}
	moduleName := getModuleName(currentPath)
	globalVariables := map[string]interface{}{
		"NewSpellCommandName": strings.ReplaceAll(strings.ReplaceAll(newSpellCommandName, "-", ""), " ", ""),
		"ModuleName":          moduleName,
	}

	err = renderer.RenderTemplate(spell, "addspellcommand", globalVariables, nil)
	if err != nil {
		fmt.Printf("An error occurred while trying to create the new spell command. Error: %s\n", err.Error())
		return err
	}

	err = renameTemplateFileNames(currentPath, newSpellCommandName)
	if err != nil {
		fmt.Printf("An error occurred while trying to create the new spell command. Error: %s\n", err.Error())
		return err
	}

	err = addCommandToBuildConfigCommand(currentPath, args)
	if err != nil {
		fmt.Printf("An error occurred while trying to create the new spell command. Error: %s\n", err.Error())
		return err
	}

	return nil
}
