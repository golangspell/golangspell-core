package appcontext

import (
	"testing"
)

func TestContext_Add(t *testing.T) {
	type fields struct {
		components map[string]Component
	}
	type args struct {
		componentName string
		component     Component
	}
	components := make(map[string]Component)
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "Add complete repository",
			fields: fields{components: components},
			args: args{
				componentName: Logger,
				component:     ApplicationContext{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			context := CreateApplicationContext()
			context.Add(tt.args.componentName, tt.args.component)
			if context.Count() == 0 {
				t.Error("Component not added")
			}
			shopDealerServiceRatingRepository :=
				context.Get(Logger)
			if shopDealerServiceRatingRepository == nil {
				t.Error("Component not found")
			}
			context.Delete(Logger)
			shopDealerServiceRatingRepository =
				context.Get(Logger)
			if shopDealerServiceRatingRepository != nil {
				t.Error("Component not deleted")
			}
		})
	}

}
