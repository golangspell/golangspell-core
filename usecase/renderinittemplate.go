package usecase

import (
	"github.com/danilovalente/golangspell-core/appcontext"
	"github.com/danilovalente/golangspell-core/domain"
	tooldomain "github.com/danilovalente/golangspell/domain"
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
