package usecase

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/danilovalente/golangspell-core/appcontext"
	"github.com/danilovalente/golangspell/domain"
)

//RenderFile renders a file
func RenderFile(sourcePath string, info os.FileInfo) error {
	fileName := filepath.Base(sourcePath)
	fmt.Printf("Rendering template: %s\n", fileName)

	file, err := os.Open(sourcePath)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	return nil
}

//RenderPath renders an object (file os directory) in the templates directory
func RenderPath(sourcePath string, info os.FileInfo, err error) error {
	if err != nil {
		log.Printf("An error occurred while trying to analyze path %s\n", err)
		return nil
	}
	if !info.IsDir() {
		RenderFile(sourcePath, info)
	} else {
		fmt.Printf("Templates Directory: %s\n", sourcePath)
	}
	return nil
}

//RenderTemplate renders all templates in the template directory providing the respective variables
//commandName: specifies the name of the command for which the template will be rendered
//variables: specifies the list of variables (value) which should be provided for rendering each file (key)
func RenderTemplate(commandName string, variables map[string]interface{}) error {
	spell := appcontext.Current.Get(appcontext.Spell).(domain.Spell)
	spellInstallation := domain.GolangLibrary{Name: spell.Name, URL: spell.URL}
	rootTemplatePath := fmt.Sprintf("%s/templates/%s", spellInstallation.SrcPath(), commandName)
	err := filepath.Walk(rootTemplatePath, RenderPath)
	return err
}
