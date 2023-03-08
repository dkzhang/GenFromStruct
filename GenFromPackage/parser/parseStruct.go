package parser

import (
	"GenFromStruct/GenFromPackage/model"
	"fmt"
	"go/types"
	"log"
	"os"
	"regexp"
	"strings"

	"golang.org/x/tools/go/packages"
)

func Parse(sourceType string, filePath string) model.StructInfo {
	// 1. Split source type into package and type name
	sourceTypePackage, sourceTypeName := splitSourceType(sourceType)

	// 2. Inspect package and use type checker to infer imported types
	pkg := loadPackage(sourceTypePackage)

	// 3. Lookup the given source type name in the package declarations
	obj := pkg.Types.Scope().Lookup(sourceTypeName)
	if obj == nil {
		failErr(fmt.Errorf("%s not found in declared types of %s",
			sourceTypeName, pkg))
	}

	// 4. We check if it is a declared type
	if _, ok := obj.(*types.TypeName); !ok {
		failErr(fmt.Errorf("%v is not a named type", obj))
	}
	// 5. We expect the underlying type to be a struct
	structType, ok := obj.Type().Underlying().(*types.Struct)
	if !ok {
		failErr(fmt.Errorf("type %v is not a struct", obj))
	}

	si := model.StructInfo{
		PackagePath: sourceTypePackage,
		FilePath:    filePath,
		StructName:  extractStructName(sourceTypeName),
		Fields:      make([]model.StructField, structType.NumFields()),
	}

	// 6. Now we can iterate through fields and access tags
	for i := 0; i < structType.NumFields(); i++ {
		si.Fields[i] = model.StructField{
			Name:   structType.Field(i).Name(),
			Type:   structType.Field(i).Type(),
			TagMap: extractTag(structType.Tag(i)),
		}
	}
	return si
}

func loadPackage(path string) *packages.Package {
	cfg := &packages.Config{Mode: packages.NeedTypes | packages.NeedImports}
	pkgs, err := packages.Load(cfg, path)
	if err != nil {
		failErr(fmt.Errorf("loading packages for inspection: %v", err))
	}
	if packages.PrintErrors(pkgs) > 0 {
		os.Exit(1)
	}

	return pkgs[0]
}

func splitSourceType(sourceType string) (string, string) {
	idx := strings.LastIndexByte(sourceType, '.')
	if idx == -1 {
		failErr(fmt.Errorf(`expected qualified type as "pkg/path.MyType"`))
	}
	sourceTypePackage := sourceType[0:idx]
	sourceTypeName := sourceType[idx+1:]
	return sourceTypePackage, sourceTypeName
}

func failErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func extractTag(tag string) map[string]string {
	m := map[string]string{}
	if len(tag) > 2 {
		re := regexp.MustCompile(`(\w+):"([^"]+)"`)
		for _, s := range re.FindAllStringSubmatch(tag, -1) {
			m[strings.TrimSpace(s[1])] = strings.TrimSpace(s[2])
		}
	}
	return m
}

func extractStructName(structName string) string {
	re := regexp.MustCompile(`(\w+)_ToG`)
	matches := re.FindStringSubmatch(structName)
	if len(matches) > 1 {
		return matches[1]
	} else {
		log.Fatal("No match found for struct name")
		return ""
	}
}
