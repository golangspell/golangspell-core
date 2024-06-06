package domain

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/ast/astutil"
)

// CodeFile provides the features for parsing and editing
type CodeFile struct {
	path       string
	code       *ast.File
	fset       *token.FileSet
	testOutput *bytes.Buffer
}

// CodeFileLoaded is true if the code file is not nil and the file was parsed
func CodeFileLoaded(codeFile *CodeFile) bool {
	return codeFile != nil && codeFile.code != nil
}

// ParseFromPath provided as a parameter
func (codeFile *CodeFile) ParseFromPath(path string) *CodeFile {
	if strings.TrimSpace(path) == "" {
		fmt.Println("Path parameter is required for parsing from path.")
		return nil
	}

	codeFile.path = path
	err := codeFile.Parse()
	if err != nil {
		fmt.Printf("Code file could not be parsed. Error: %s\n", err.Error())
		return nil
	}
	return codeFile
}

// Parse the code file contained in the configured path
func (codeFile *CodeFile) Parse() error {
	codeFile.fset = token.NewFileSet() // positions are relative to fset
	var err error
	codeFile.code, err = parser.ParseFile(codeFile.fset, codeFile.path, nil, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("an error occurred while trying to parse the file %s. Error: %s", codeFile.path, err.Error())
	}

	return nil
}

// AddImport to the loaded code file
func (codeFile *CodeFile) AddImport(importPath string, alias string) *CodeFile {
	if !CodeFileLoaded(codeFile) {
		fmt.Printf("Code file not loaded. Impossible to add import %s\n", importPath)
		return codeFile
	}
	importPath = strings.TrimSpace(importPath)
	if importPath == "" {
		fmt.Printf("Import path not provided to be added to the code file %s\n", codeFile.path)
		return codeFile
	}
	alias = strings.TrimSpace(alias)
	var added bool
	if alias == "" {
		added = astutil.AddImport(codeFile.fset, codeFile.code, importPath)
	} else {
		added = astutil.AddNamedImport(codeFile.fset, codeFile.code, alias, importPath)
	}
	if !added {
		fmt.Printf("An error occurred while trying to add the import %s to the code file %s.\n", importPath, codeFile.path)
		return codeFile
	}
	err := codeFile.Save()
	if err != nil {
		fmt.Printf("An error occurred while trying to save the code file %s. Message: %s\n", codeFile.path, err.Error())
		return codeFile
	}
	fmt.Printf("Added import to package %s to code file %s\n", importPath, codeFile.path)

	return codeFile
}

func (codeFile *CodeFile) GetTokenIndexByKind(kind token.Token) int {
	for i, decl := range codeFile.code.Decls {
		switch decl := decl.(type) {
		case *ast.GenDecl:
			if decl.Tok == kind {
				return i
			}
		}
	}
	return -1
}

// GetConst declaration
func (codeFile *CodeFile) GetConst() *ast.GenDecl {
	if codeFile == nil || codeFile.code == nil || codeFile.code.Decls == nil {
		return nil
	}
	for _, decl := range codeFile.code.Decls {
		switch decl := decl.(type) {
		case *ast.GenDecl:
			if decl.Tok == token.CONST {
				return decl
			}
		}
	}

	return nil
}

// AddConstant to the loaded code file
func (codeFile *CodeFile) AddConstant(constantName string, constantKind token.Token, constantValue string) *CodeFile {
	if !CodeFileLoaded(codeFile) {
		fmt.Printf("Code file not loaded. Impossible to add constant %s\n", constantName)
		return codeFile
	}
	constantName = strings.TrimSpace(constantName)
	if constantName == "" {
		fmt.Printf("Constant name not provided to be added to the code file %s\n", codeFile.path)
		return codeFile
	}
	//TODO: Check properly constant value
	constantValue = strings.TrimSpace(string(constantValue))
	if constantValue == "" {
		fmt.Printf("Constant value not provided to be added to the code file %s\n", codeFile.path)
		return codeFile
	}

	// Create the new constant declaration
	newConst := &ast.GenDecl{
		Tok: token.CONST,
		Specs: []ast.Spec{
			&ast.ValueSpec{
				Names: []*ast.Ident{
					ast.NewIdent(constantName),
				},
				Values: []ast.Expr{
					&ast.BasicLit{
						Kind:  constantKind,
						Value: fmt.Sprintf(`"%s"`, constantValue),
					},
				},
			},
		},
	}
	insertIndex := codeFile.GetTokenIndexByKind(token.IMPORT)
	if insertIndex < 0 {
		insertIndex = codeFile.GetTokenIndexByKind(token.PACKAGE)
	}
	insertIndex++
	// Add the new constant to the file
	codeFile.code.Decls = append(codeFile.code.Decls[:insertIndex], append([]ast.Decl{newConst}, codeFile.code.Decls[insertIndex:]...)...)

	err := codeFile.Save()
	if err != nil {
		fmt.Printf("An error occurred while trying to save the code file %s. Message: %s\n", codeFile.path, err.Error())
		return codeFile
	}
	fmt.Printf("Added constant %s to code file %s\n", constantName, codeFile.path)
	return codeFile
}

// AddStatementToFunction adds a new statement to the function with the given "functionName", under the statement found by the "matchStatement" parameter
// If nil is provided as the parameter "matchStatement", the newStatement is added as the last line of the function
func (codeFile *CodeFile) AddStatementToFunction(functionName string, newStatement string, matchStatement func(statementCode string) bool) *CodeFile {
	if !CodeFileLoaded(codeFile) {
		fmt.Printf("Code file not loaded. Impossible to add the statement %s\n", newStatement)
		return codeFile
	}
	newStatement = strings.TrimSpace(newStatement)
	if newStatement == "" {
		fmt.Printf("newStatement not provided to be added to the code file %s\n", codeFile.path)
		return codeFile
	}

	// Find the function to be modified
	var targetFunc *ast.FuncDecl
	ast.Inspect(codeFile.code, func(node ast.Node) bool {
		if fn, ok := node.(*ast.FuncDecl); ok {
			if fn.Name.Name == functionName {
				targetFunc = fn
				return false
			}
		}
		return true
	})

	if targetFunc == nil {
		fmt.Printf("Function '%s' was not found at the code file %s\n", functionName, codeFile.path)
		return codeFile
	}

	// The new statement will be added under the statement found here using the matchStatement function
	var targetStmt *ast.Stmt
	for i, stmt := range targetFunc.Body.List {
		var buf strings.Builder
		if err := printer.Fprint(&buf, codeFile.fset, stmt); err != nil {
			fmt.Printf("Impossible to print the code line number %d of the function '%s' of the code file %s\n", i, functionName, codeFile.path)
		} else {
			if matchStatement(buf.String()) {
				targetStmt = &stmt
				break
			}
		}
	}

	// Create the new statement
	newStmt := &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun: &ast.Ident{
				Name: "fmt",
			},
			Args: []ast.Expr{
				&ast.BasicLit{
					Kind:  token.STRING,
					Value: newStatement,
				},
			},
		},
	}

	// Add the new statement to the function
	if targetStmt != nil {
		var newBody []ast.Stmt
		for _, stmt := range targetFunc.Body.List {
			newBody = append(newBody, stmt)
			if stmt == *targetStmt {
				newBody = append(newBody, newStmt)
			}
		}
		targetFunc.Body.List = newBody
	} else {
		targetFunc.Body.List = append(targetFunc.Body.List, newStmt)
	}

	err := codeFile.Save()
	if err != nil {
		fmt.Printf("An error occurred while trying to save the code file %s. Message: %s\n", codeFile.path, err.Error())
		return codeFile
	}
	fmt.Printf("Added statement '%s' to code file %s\n", newStatement, codeFile.path)

	return codeFile
}

func (codeFile *CodeFile) saveToFile() error {
	if destFileInfo, err := os.Stat(codeFile.path); err == nil && !destFileInfo.IsDir() {
		err := GetRenderer().BackupExistingCode(codeFile.path)
		if err != nil {
			return err
		}
	}
	directory := filepath.Dir(codeFile.path)
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		err = os.MkdirAll(directory, 0755)
		if err != nil {
			return err
		}
	}
	file, err := os.Create(codeFile.path)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := printer.Fprint(file, codeFile.fset, codeFile.code); err != nil {
		return err
	}

	return nil
}

func (codeFile *CodeFile) saveToTestBuffer() error {
	if err := printer.Fprint(codeFile.testOutput, codeFile.fset, codeFile.code); err != nil {
		return err
	}
	return nil
}

func (codeFile *CodeFile) Save() error {
	if strings.TrimSpace(codeFile.path) == "" {
		return errors.New("codeFile.path attribute is required for saving the code file")
	}
	if codeFile.testOutput != nil {
		return codeFile.saveToTestBuffer()
	} else {
		return codeFile.saveToFile()
	}
}

func (codeFile *CodeFile) Code() *ast.File {
	return codeFile.code
}

func (codeFile *CodeFile) Fset() *token.FileSet {
	return codeFile.fset
}
