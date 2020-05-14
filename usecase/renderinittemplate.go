package usecase

import (
	"github.com/golangspell/golangspell-core/appcontext"
	"github.com/golangspell/golangspell-core/domain"
	tooldomain "github.com/golangspell/golangspell/domain"
)

//RenderInitTemplate renders the templates defined to the init command with the proper variables
func RenderInitTemplate(args []string) error {
	spell := appcontext.Current.Get(appcontext.Spell).(tooldomain.Spell)
	renderer := domain.GetRenderer()
	globalVariables := map[string]interface{}{
		"ModuleName": args[0],
		"AppName":    args[1],
	}

	return renderer.RenderTemplate(spell, "init", globalVariables, nil)
}
