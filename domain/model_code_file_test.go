package domain

import (
	"bytes"
	"go/ast"
	"go/token"
	"strings"
	"testing"
)

func TestCodeFile_AddConstant(t *testing.T) {
	type fields struct {
		path string
		code *ast.File
	}
	type args struct {
		constantName  string
		constantKind  token.Token
		constantValue string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CodeFile
	}{
		{name: "codeFile loaded and constant definition provided", fields: fields{path: "../appcontext/context_test.go", code: new(CodeFile).ParseFromPath("../appcontext/context_test.go").code}, args: args{constantName: "TestRepository", constantKind: token.STRING, constantValue: "\"TestRepository\""}, want: new(CodeFile).ParseFromPath("../appcontext/context_test.go")},
		{name: "codeFile loaded and constant definition provided", fields: fields{path: "../appcontext/context.go", code: new(CodeFile).ParseFromPath("../appcontext/context.go").code}, args: args{constantName: "TestRepository", constantKind: token.STRING, constantValue: "\"TestRepository\""}, want: new(CodeFile).ParseFromPath("../appcontext/context.go")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			codeFile := &CodeFile{
				path: tt.fields.path,
				code: tt.fields.code,
			}
			codeFile.testOutput = new(bytes.Buffer)
			err := codeFile.Parse()
			if err != nil {
				t.Errorf("An error occurred while trying to parse the file %s. Message: %s", tt.fields.path, err.Error())
			}
			got := codeFile.AddConstant(tt.args.constantName, tt.args.constantKind, tt.args.constantValue)
			resultFileContents := got.testOutput.String()
			newConstValue := "\"TestRepository\""
			if !strings.Contains(resultFileContents, newConstValue) {
				t.Errorf("New constant not added: %s", newConstValue)
			}
		})
	}
}

func TestCodeFile_AddImport(t *testing.T) {
	type fields struct {
		path        string
		code        *ast.File
		importValue string
	}
	type args struct {
		importPath string
		alias      string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CodeFile
	}{
		{name: "codeFile loaded and constant definition provided", fields: fields{path: "../appcontext/context_test.go", code: new(CodeFile).ParseFromPath("../appcontext/context_test.go").code, importValue: "\"path/filepath\""}, args: args{importPath: "path/filepath", alias: ""}, want: new(CodeFile).ParseFromPath("../appcontext/context_test.go")},
		{name: "codeFile loaded and constant definition provided", fields: fields{path: "../appcontext/context.go", code: new(CodeFile).ParseFromPath("../appcontext/context.go").code, importValue: "_ \"github.com/golangspell/golangspell-core/config\""}, args: args{importPath: "github.com/golangspell/golangspell-core/config", alias: "_"}, want: new(CodeFile).ParseFromPath("../appcontext/context.go")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			codeFile := &CodeFile{
				path: tt.fields.path,
				code: tt.fields.code,
			}
			codeFile.testOutput = new(bytes.Buffer)
			err := codeFile.Parse()
			if err != nil {
				t.Errorf("An error occurred while trying to parse the file %s. Message: %s", tt.fields.path, err.Error())
			}
			var got *CodeFile = codeFile.AddImport(tt.args.importPath, tt.args.alias)
			resultFileContents := got.testOutput.String()
			if !strings.Contains(resultFileContents, tt.fields.importValue) {
				t.Errorf("New import not added: %s", tt.fields.importValue)
			}
		})
	}
}

func TestCodeFile_AddCallToFunction(t *testing.T) {
	type fields struct {
		path       string
		code       *ast.File
		testOutput *bytes.Buffer
	}
	type args struct {
		functionName          string
		newStatementCall      string
		newStatementArguments string
		matchStatement        func(statementCode string) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CodeFile
	}{
		{
			name:   "Add new fmt.Printf to code",
			fields: fields{path: "../appcontext/context.go", code: new(CodeFile).ParseFromPath("../appcontext/context.go").code, testOutput: new(bytes.Buffer)},
			args: args{functionName: "Add", newStatementCall: `fmt.Printf`, newStatementArguments: `"My pretty new statement %s", componentName`, matchStatement: func(statementCode string) bool {
				return strings.Contains(statementCode, "defer applicationContext.componentMutex.Unlock()")
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			codeFile := &CodeFile{
				path:       tt.fields.path,
				code:       tt.fields.code,
				testOutput: tt.fields.testOutput,
			}
			err := codeFile.Parse()
			if err != nil {
				t.Errorf("An error occurred while trying to parse the file %s. Message: %s", tt.fields.path, err.Error())
			}
			got := codeFile.AddCallToFunction(tt.args.functionName, tt.args.newStatementCall, tt.args.newStatementArguments, tt.args.matchStatement)
			resultFileContents := got.testOutput.String()
			if !strings.Contains(resultFileContents, tt.args.newStatementCall) {
				t.Errorf("New statement not added: %s", tt.args.newStatementCall)
			}
		})
	}
}

func TestCodeFile_AddAttributeToStruct(t *testing.T) {
	type fields struct {
		path       string
		code       *ast.File
		fset       *token.FileSet
		testOutput *bytes.Buffer
	}
	type args struct {
		structTypeName string
		attributeName  string
		attributeType  string
		attributeValue string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CodeFile
	}{
		{
			name:   "Add DBConnectionString attribute",
			fields: fields{path: "../domain/model_code_file.go", code: new(CodeFile).ParseFromPath("../domain/model_code_file.go").code, testOutput: new(bytes.Buffer)},
			args:   args{structTypeName: "CodeFile", attributeName: "DBConnectionString", attributeType: "string", attributeValue: "`env:\"DB_CONNECTION_STRING\" envDefault:\"\"`"},
		},
		{
			name:   "Add DBConnectionCertificateFileName attribute",
			fields: fields{path: "../domain/model_code_file.go", code: new(CodeFile).ParseFromPath("../domain/model_code_file.go").code, testOutput: new(bytes.Buffer)},
			args:   args{structTypeName: "CodeFile", attributeName: "DBConnectionCertificateFileName", attributeType: "string", attributeValue: "`env:\"DB_CONNECTION_CERTIFICATE_FILE_NAME\" envDefault:\"\"`"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			codeFile := &CodeFile{
				path:       tt.fields.path,
				code:       tt.fields.code,
				fset:       tt.fields.fset,
				testOutput: tt.fields.testOutput,
			}
			err := codeFile.Parse()
			if err != nil {
				t.Errorf("An error occurred while trying to parse the file %s. Message: %s", tt.fields.path, err.Error())
			}
			got := codeFile.AddAttributeToStruct(tt.args.structTypeName, tt.args.attributeName, tt.args.attributeType, tt.args.attributeValue)
			resultFileContents := got.testOutput.String()
			if !strings.Contains(resultFileContents, tt.args.attributeName) ||
				!strings.Contains(resultFileContents, tt.args.attributeType) ||
				!strings.Contains(resultFileContents, tt.args.attributeValue) {
				t.Errorf("New attribute not added: %s", tt.args.attributeName)
			}
		})
	}
}
