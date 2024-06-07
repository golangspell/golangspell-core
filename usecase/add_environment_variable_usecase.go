package usecase

import (
	"fmt"

	"github.com/golangspell/golangspell-core/appcontext"
	"github.com/golangspell/golangspell-core/config"
	"github.com/golangspell/golangspell-core/domain"
	toolconfig "github.com/golangspell/golangspell/config"
)

// AddEnvironmentVariable
type AddEnvironmentVariable struct {
}

// Execute the usecase AddEnvironmentVariable
func (u *AddEnvironmentVariable) Execute(currentPath string, attributeName string, attributeType string, attributeValue string) error {
	contextFilePath := fmt.Sprintf("%s%sconfig%senvironment.go", currentPath, toolconfig.PlatformSeparator, toolconfig.PlatformSeparator)
	return new(domain.CodeFile).
		ParseFromPath(contextFilePath).
		AddAttributeToStruct("config", attributeName, attributeType, attributeValue).
		Save()
}

func buildAddEnvironmentVariable() appcontext.Component {
	addEnvironmentVariable := &AddEnvironmentVariable{}

	return addEnvironmentVariable
}

func GetAddEnvironmentVariable() *AddEnvironmentVariable {
	return appcontext.Current.Get(appcontext.AddEnvironmentVariable).(*AddEnvironmentVariable)
}

func init() {
	if config.Values.TestRun {
		return
	}
	appcontext.Current.Add(appcontext.AddEnvironmentVariable, buildAddEnvironmentVariable)
}
