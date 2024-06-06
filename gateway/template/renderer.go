package template

import (
	"github.com/golangspell/golangspell-core/appcontext"
	"github.com/golangspell/golangspell-core/config"
	"github.com/golangspell/golangspell/gateway/template"
)

// GetRenderer lazy loads a Renderer
func GetRenderer() appcontext.Component {
	return &template.Renderer{}
}

func init() {
	if config.Values.TestRun {
		return
	}

	appcontext.Current.Add(appcontext.Renderer, GetRenderer)
}
