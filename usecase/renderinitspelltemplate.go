package usecase

import (
	"github.com/golangspell/golangspell-core/appcontext"
	"github.com/golangspell/golangspell-core/domain"
	tooldomain "github.com/golangspell/golangspell/domain"
)

//RenderInitSpellTemplate renders the templates defined to the initspell command with the proper variables
func RenderInitSpellTemplate(args []string) error {
	spell := appcontext.Current.Get(appcontext.Spell).(tooldomain.Spell)
	renderer := domain.GetRenderer()
	globalVariables := map[string]interface{}{
		"ModuleName":   args[0],
		"NewSpellName": args[1],
	}

	return renderer.RenderTemplate(spell, "initspell", globalVariables, nil)
}
