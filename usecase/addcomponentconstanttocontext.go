package usecase

import (
	"fmt"
	"go/token"

	"github.com/golangspell/golangspell-core/appcontext"
	"github.com/golangspell/golangspell-core/config"
	"github.com/golangspell/golangspell-core/domain"
	toolconfig "github.com/golangspell/golangspell/config"
)

// AddComponentConstantToContext adds a new component constant to the context file
type AddComponentConstantToContext struct {
}

// Execute the usecase AddComponentConstantToContext
func (u *AddComponentConstantToContext) Execute(currentPath string, componentName string) error {
	contextFilePath := fmt.Sprintf("%s%sappcontext%scontext.go", currentPath, toolconfig.PlatformSeparator, toolconfig.PlatformSeparator)
	return new(domain.CodeFile).
		ParseFromPath(contextFilePath).
		AddConstant(componentName, token.STRING, componentName).
		Save()
}

func buildAddComponentConstantToContext() appcontext.Component {
	addComponentConstantToContext := &AddComponentConstantToContext{}

	return addComponentConstantToContext
}

func GetAddComponentConstantToContext() *AddComponentConstantToContext {
	return appcontext.Current.Get(appcontext.AddComponentConstantToContext).(*AddComponentConstantToContext)
}

func init() {
	if config.Values.TestRun {
		return
	}
	appcontext.Current.Add(appcontext.AddComponentConstantToContext, buildAddComponentConstantToContext)
}
