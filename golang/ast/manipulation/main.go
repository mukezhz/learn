package main

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"strings"
)

func main() {
	src := `package infrastructure

			import "go.uber.org/fx"

			var Modules = fx.Options(
				fx.Provide(NewRouter),
				fx.Provide(NewDatabase),
				// ... other providers ...
				fx.Provide(NewGmailService),
			)
			`
	// Parse the source code
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	// Traverse the AST and find the fx.Options call
	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.CallExpr:
			if sel, ok := x.Fun.(*ast.SelectorExpr); ok {
				if sel.Sel.Name == "Options" {
					for _, arg := range x.Args {
						call := arg.(*ast.CallExpr)
						log.Printf("CALL:: %#v\n\n", call)
					}
					x.Args = append(x.Args, &ast.CallExpr{
						Fun: &ast.SelectorExpr{
							X:   ast.NewIdent("fx"),
							Sel: ast.NewIdent("Provide"),
						},
						Args: []ast.Expr{
							ast.NewIdent("NewS3Client"),
						},
					})
				}
			}
		}
		return true
	})

	// Add the source code in buffer
	var buf bytes.Buffer
	if err := format.Node(&buf, fset, node); err != nil {
		log.Fatal(err)
	}
	log.Println("FORMATTED CODE", buf.String())
	formattedCode := buf.String()
	providerToInsert := "fx.Provide(NewS3Client),"
	formattedCode = strings.Replace(formattedCode, providerToInsert, "\n\t"+providerToInsert, 1)

	// Output the modified source
	log.Println("After modification:", formattedCode)
}
