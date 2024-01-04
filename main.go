package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
)

func main() {
	v := uint32(math.MaxUint32)
	u := int(0)
	if uint32(int(v)) == v {
		u = int(v)
	}
	log.Printf("%v", u)
	var analysisPath string
	if len(os.Args) == 2 {
		var err error
		analysisPath, err = filepath.Abs(os.Args[1])
		if err != nil {
			panic(fmt.Errorf("unable to expand path '%v' to absolute: %w", os.Args[1], err))
		}
	} else if len(os.Args) == 1 {
		var err error
		analysisPath, err = os.Getwd()
		if err != nil {
			panic(fmt.Errorf("unable to get working directory: %w", err))
		}
	}
	log.Printf("module path: %v", analysisPath)
	assemblyLines, err := AnalyzeModuleAssembly(analysisPath)
	if err != nil {
		panic(fmt.Errorf("failed to analyze module assembly: %w", err))
	}
	err = AnalyzeModule(analysisPath, assemblyLines, Govanish)
	if err != nil {
		panic(fmt.Errorf("failed to analyze module AST: %w", err))
	}
}
