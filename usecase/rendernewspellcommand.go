package usecase

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/golangspell/golangspell-core/appcontext"
	"github.com/golangspell/golangspell-core/domain"
	toolconfig "github.com/golangspell/golangspell/config"
	tooldomain "github.com/golangspell/golangspell/domain"
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
	var validArgs string
	if len(args) > 1 {
		commandArgs := make([]string, len(args)-1)
		for i := 1; i < len(args); i++ {
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

	variables := make(map[string]interface{})
	variables["CommandName"] = args[0]
	variables["ValidArgs"] = validArgs
	renderer := domain.GetRenderer()
	spell := appcontext.Current.Get(appcontext.Spell).(tooldomain.Spell)
	renderedTemplateString, err := renderer.RenderString(spell, "addspellcommand", "buildconfig_add.got", variables)
	if err != nil {
		fmt.Printf("An error occurred while trying to create the new spell command. Error: %s\n", err.Error())
		return ""
	}

	code = r.ReplaceAllString(code, renderedTemplateString)
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
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("An error occurred while trying to create the new spell command. Error: %s\n", err.Error())
		return err
	}
	code := addCommandConfig(string(content), args)
	err = os.WriteFile(filePath, []byte(code), 0644)
	if err != nil {
		fmt.Printf("An error occurred while trying to create the new spell command. Error: %s\n", err.Error())
		return err
	}

	return nil
}

func createTemplateDirectory(currentPath string, newSpellCommandName string) error {
	templateDirectory := fmt.Sprintf("%s%stemplates%s%s", currentPath, toolconfig.PlatformSeparator, toolconfig.PlatformSeparator, newSpellCommandName)
	if _, err := os.Stat(templateDirectory); os.IsNotExist(err) {
		err = os.MkdirAll(templateDirectory, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

// RenderNewSpellCommandTemplate renders the templates defined to the addspellcommand command with the proper variables
func RenderNewSpellCommandTemplate(args []string) error {
	spell := appcontext.Current.Get(appcontext.Spell).(tooldomain.Spell)
	renderer := domain.GetRenderer()
	newSpellCommandName := args[0]
	safeNewSpellCommandName := strings.ReplaceAll(strings.ReplaceAll(newSpellCommandName, "-", ""), " ", "")
	currentPath, err := os.Getwd()
	if err != nil {
		fmt.Printf("An error occurred while trying to create the new spell command. Error: %s\n", err.Error())
		return err
	}
	moduleName := toolconfig.GetModuleName(currentPath)
	globalVariables := map[string]interface{}{
		"NewSpellCommandName":     newSpellCommandName,
		"SafeNewSpellCommandName": safeNewSpellCommandName,
		"ModuleName":              moduleName,
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

	err = createTemplateDirectory(currentPath, newSpellCommandName)
	if err != nil {
		fmt.Printf("An error occurred while trying to create the new spell command. Error: %s\n", err.Error())
		return err
	}

	return nil
}
