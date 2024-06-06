package usecase

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/golangspell/golangspell-core/appcontext"
	"github.com/golangspell/golangspell-core/domain"
	toolconfig "github.com/golangspell/golangspell/config"
	tooldomain "github.com/golangspell/golangspell/domain"
	"github.com/iancoleman/strcase"
)

func renameAddusecaseTemplateFileNames(currentPath string, componentName string) error {
	sourcePath := fmt.Sprintf("%s%susecase%snewusecase.go", currentPath, toolconfig.PlatformSeparator, toolconfig.PlatformSeparator)
	directory := filepath.Dir(sourcePath)
	destinationPath := fmt.Sprintf("%s%s%s_usecase.go", directory, toolconfig.PlatformSeparator, strcase.ToSnake(componentName))

	err := os.Rename(sourcePath, destinationPath)
	if err != nil {
		return err
	}
	return nil
}

// RenderaddusecaseTemplate renders the templates defined to the addusecase command with the proper variables
func RenderaddusecaseTemplate(args []string) error {
	spell := appcontext.Current.Get(appcontext.Spell).(tooldomain.Spell)
	renderer := domain.GetRenderer()
	currentPath, err := os.Getwd()
	if err != nil {
		fmt.Printf("An error occurred while trying to add the new usecase %s. Error: %s\n", args[0], err.Error())
		return err
	}
	moduleName := toolconfig.GetModuleName(currentPath)
	safeNewUsecaseName := strings.ReplaceAll(strings.ReplaceAll(args[0], "-", ""), " ", "")
	camelSafeNewUsecaseName := strcase.ToCamel(safeNewUsecaseName)
	globalVariables := map[string]interface{}{
		"SafeNewUsecaseName":           camelSafeNewUsecaseName,
		"SafeNewUsecaseNameLowerCamel": strcase.ToLowerCamel(safeNewUsecaseName),
		"ModuleName":                   moduleName,
	}

	err = renderer.RenderTemplate(spell, "addusecase", globalVariables, nil)
	if err != nil {
		fmt.Printf("An error occurred while trying to render the template. Error: %s\n", err.Error())
		return err
	}

	err = GetAddComponentConstantToContext().Execute(currentPath, camelSafeNewUsecaseName)

	if err != nil {
		fmt.Printf("An error occurred while trying to save the new constant to the context file. Error: %s\n", err.Error())
		return err
	}

	err = GetAddPackageImportToMain().Execute(moduleName, currentPath, fmt.Sprintf("%s/usecase", moduleName))

	if err != nil {
		fmt.Printf("An error occurred while trying to save the new import to the main file. Error: %s\n", err.Error())
		return err
	}

	err = renameAddusecaseTemplateFileNames(currentPath, camelSafeNewUsecaseName)
	if err != nil {
		fmt.Printf("An error occurred while trying to rename the rendered template files. Error: %s\n", err.Error())
		return err
	}

	return nil
}
