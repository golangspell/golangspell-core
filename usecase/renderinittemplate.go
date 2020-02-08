package usecase

import "github.com/danilovalente/golangspell-core/domain"

//RenderInitTemplate renders the templates defined to the init command with the proper variables
func RenderInitTemplate(args []string) error {
	renderer := domain.GetRenderer()
	globalVariables := map[string]interface{}{
		"ModuleName": args[0],
		"AppName":    args[1],
	}

	return renderer.RenderTemplate("init", globalVariables, nil)
}
