package generator

import (
	"GenFromStruct/GenFromFile/model"
	"bytes"
	"fmt"
	"go/format"
	"log"
	"os"
	"path/filepath"
)

func GenerateModelFile(sfi model.StructFileInfo) error {
	// create a new file
	dir := filepath.Dir(sfi.OriginFileName)
	modelFileName := filepath.Join(dir, sfi.FileName+".go")

	var buf bytes.Buffer

	// write package name
	buf.WriteString(fmt.Sprintf("package %s \n", sfi.PackageName))

	// write import
	for _, s := range sfi.Imports {
		fmt.Fprintf(&buf, "import %s \n", s)
	}

	// write struct
	for _, sInfo := range sfi.Structs {
		// start of struct
		fmt.Fprintf(&buf, "type %s struct { \n", sInfo.ClassName)

		// write fields
		for _, f := range sInfo.Fields {
			fmt.Fprintf(&buf, "%s %s", f.Name, f.Type)

			// write tag
			fmt.Fprintf(&buf, "`db:\"%s\" json:\"%s\"` ", f.TagMap["db"], f.TagMap["json"])

			// write comment
			if len(f.Comment) > 0 {
				fmt.Fprintf(&buf, "// %s ", f.Comment)
			}

			// end of line
			fmt.Fprintf(&buf, "\n")
		}

		// end of struct
		fmt.Fprintf(&buf, "} \n")
	}

	// write schema
	for _, sInfo := range sfi.Structs {
		// start of schema
		fmt.Fprintf(&buf, "var Schema%s = ` \n", sInfo.ClassName)

		// create table
		fmt.Fprintf(&buf, "\t\tCREATE TABLE %s ( \n", sInfo.TableName)

		// write fields
		for _, f := range sInfo.Fields {
			fmt.Fprintf(&buf, "\t\t\t%s %s", f.TagMap["db"], f.TagMap["pgType"])

			if pgA, ok := f.TagMap["pgA"]; ok {
				fmt.Fprintf(&buf, " %s", pgA)
			}
			// end of line
			fmt.Fprintf(&buf, ", \n")
		}

		// end of create table
		fmt.Fprintf(&buf, "); \n")

		// end of schema
		fmt.Fprintf(&buf, "` \n")
	}

	// format the output
	src, err := format.Source(buf.Bytes())
	if err != nil {
		log.Printf("formatting source: %s", err)
		src = buf.Bytes()
	}
	//src := buf.Bytes()

	// write to file
	err = os.WriteFile(modelFileName, src, 0644)
	if err != nil {
		log.Fatalf("writing output file: %s", err)
	}

	return nil
}
