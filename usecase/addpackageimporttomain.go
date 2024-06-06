package usecase

import (
	"fmt"

	"github.com/golangspell/golangspell-core/appcontext"
	"github.com/golangspell/golangspell-core/config"
	"github.com/golangspell/golangspell-core/domain"
	toolconfig "github.com/golangspell/golangspell/config"
	"golang.org/x/tools/go/ast/astutil"
)

// AddPackageImportToMain adds a new import to the main file
type AddPackageImportToMain struct {
}

// Execute the usecase AddPackageImportToMain
func (u *AddPackageImportToMain) Execute(moduleName string, currentPath string, importPath string) bool {
	mainFilePath := fmt.Sprintf("%s%smain.go", currentPath, toolconfig.PlatformSeparator)
	mainFile := new(domain.CodeFile).ParseFromPath(mainFilePath)
	return astutil.AddNamedImport(mainFile.Fset(), mainFile.Code(), "_", importPath)
}

func buildAddPackageImportToMain() appcontext.Component {
	addComponentConstantToContext := &AddPackageImportToMain{}

	return addComponentConstantToContext
}

func GetAddPackageImportToMain() *AddPackageImportToMain {
	return appcontext.Current.Get(appcontext.AddPackageImportToMain).(*AddPackageImportToMain)
}

func init() {
	if config.Values.TestRun {
		return
	}
	appcontext.Current.Add(appcontext.AddPackageImportToMain, buildAddPackageImportToMain)
}
