package main

import (
	"GenFromStruct/GenFromFile/generator"
	"GenFromStruct/GenFromFile/parser"
	"GenFromStruct/GenFromFile/processor"
	"fmt"
	"os"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("pwd: %s \n", pwd)

	filename := "./sampleData/mypack/mypack_ToG.go"

	sfi := parser.Parse(filename)

	nSfi := processor.ProcessTag(sfi)

	generator.GenerateModelFile(nSfi)
}
