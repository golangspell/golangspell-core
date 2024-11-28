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

func ensureDir(dirPath string) error {
	// Check if the directory exists
	_, err := os.Stat(dirPath)
	if err == nil {
		// Directory exists, return nil
		return nil
	}

	if os.IsNotExist(err) {
		// Directory does not exist, create it
		err = os.MkdirAll(dirPath, 0755)
		if err != nil {
			return err
		}
	} else {
		// Handle other errors
		return err
	}

	return nil
}

func renameAddgatewayTemplateFileNames(currentPath string, safeNewGatewayPackageName string, componentName string) error {
	sourcePath := fmt.Sprintf("%s%sgateway%snewgateway.go", currentPath, toolconfig.PlatformSeparator, toolconfig.PlatformSeparator)
	directory := filepath.Dir(sourcePath)
	destinationPath := fmt.Sprintf("%s%s%s%s%s.go", directory, toolconfig.PlatformSeparator, safeNewGatewayPackageName, toolconfig.PlatformSeparator, strcase.ToSnake(componentName))
	destinationDirectory := filepath.Dir(destinationPath)
	err := ensureDir(destinationDirectory)
	if err != nil {
		return err
	}
	err = os.Rename(sourcePath, destinationPath)
	if err != nil {
		return err
	}
	return nil
}

// RenderaddgatewayTemplate renders the templates defined to the addgateway command with the proper variables
func RenderaddgatewayTemplate(args []string) error {
	spell := appcontext.Current.Get(appcontext.Spell).(tooldomain.Spell)
	renderer := domain.GetRenderer()
	currentPath, err := os.Getwd()
	if err != nil {
		fmt.Printf("An error occurred while trying to add the new gateway %s. Error: %s\n", args[1], err.Error())
		return err
	}
	moduleName := toolconfig.GetModuleName(currentPath)
	safeNewGatewayName := strings.ReplaceAll(strings.ReplaceAll(args[1], "-", ""), " ", "")
	camelSafeNewGatewayName := strcase.ToCamel(safeNewGatewayName)
	safeNewGatewayPackageName := strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(args[0], "_", ""), " ", ""), "-", ""))
	var safeNewGatewayPackageNameLastName string
	if strings.Contains(safeNewGatewayPackageName, "/") {
		splittedSafeNewGatewayPackageName := strings.Split(safeNewGatewayPackageName, "/")
		safeNewGatewayPackageNameLastName = splittedSafeNewGatewayPackageName[len(splittedSafeNewGatewayPackageName)-1]
	} else {
		safeNewGatewayPackageNameLastName = safeNewGatewayPackageName
	}
	globalVariables := map[string]interface{}{
		"SafeNewGatewayPackageName":    safeNewGatewayPackageNameLastName,
		"SafeNewGatewayName":           camelSafeNewGatewayName,
		"SafeNewGatewayNameLowerCamel": strcase.ToLowerCamel(safeNewGatewayName),
		"ModuleName":                   moduleName,
	}

	err = renderer.RenderTemplate(spell, "addgateway", globalVariables, nil)
	if err != nil {
		fmt.Printf("An error occurred while trying to render the template. Error: %s\n", err.Error())
		return err
	}

	err = GetAddComponentConstantToContext().Execute(currentPath, camelSafeNewGatewayName)

	if err != nil {
		fmt.Printf("An error occurred while trying to save the new constant to the context file. Error: %s\n", err.Error())
		return err
	}

	err = GetAddPackageImportToMain().Execute(moduleName, currentPath, fmt.Sprintf("%s/gateway/%s", moduleName, safeNewGatewayPackageName))

	if err != nil {
		fmt.Printf("An error occurred while trying to save the new import to the main file. Error: %s\n", err.Error())
		return err
	}

	err = renameAddgatewayTemplateFileNames(currentPath, safeNewGatewayPackageName, camelSafeNewGatewayName)
	if err != nil {
		fmt.Printf("An error occurred while trying to rename the rendered template files. Error: %s\n", err.Error())
		return err
	}

	return nil

}
