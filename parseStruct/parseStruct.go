package parseStruct

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
)

func Parse(filename string) (ssInfo []StructInfo) {
	f, err := parseFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return parseStruct(f)
}

func parseFile(filename string) (f *ast.File, err error) {
	log.Printf("Running on %s\n", filename)

	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("os.ReadFile(%s) error: %v", filename, err)
	}

	fset := token.NewFileSet()
	f, err = parser.ParseFile(fset, filename, fileContent, parser.ParseComments|parser.AllErrors)
	if err != nil {
		return nil, fmt.Errorf("parser.ParseFile(%s) error: %v", filename, err)
	}

	return f, nil

}

func parseStruct(f *ast.File) (ssInfo []StructInfo) {
	for _, decl := range f.Decls {
		if s, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range s.Specs {
				if ts, ok := spec.(*ast.TypeSpec); ok {
					if st, ok := ts.Type.(*ast.StructType); ok {

						sName, isMatch := extractStructName(ts.Name.Name)
						log.Printf("sName = %s, isMatch = %v \n", sName, isMatch)

						if isMatch {
							sInfo := StructInfo{
								Name:   sName,
								Fields: make([]StructField, len(st.Fields.List)),
							}
							for j, field := range st.Fields.List {
								sInfo.Fields[j] = StructField{
									Name:    field.Names[0].Name,
									Tag:     field.Tag.Value,
									TagMap:  extractTag(field.Tag.Value),
									Comment: field.Comment.Text(),
								}
								switch field.Type.(type) {
								case *ast.Ident:
									sInfo.Fields[j].Type = field.Type.(*ast.Ident).Name
								case *ast.SelectorExpr:
									sInfo.Fields[j].Type = field.Type.(*ast.SelectorExpr).X.(*ast.Ident).Name +
										"." +
										field.Type.(*ast.SelectorExpr).Sel.Name
								}
							}
							ssInfo = append(ssInfo, sInfo)
						}
					}
				}
			}
		}
	}
	return ssInfo
}
