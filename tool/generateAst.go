// package tool

// import (
// 	"fmt"
// 	"os"
// )

//	func generateAst(args []string) {
//		if len(args) != 1 {
//			fmt.Fprint(os.Stderr, "Usage: generate_ast <output directory>")
//			os.Exit(64)
//		}
//		outputDir := args[0]
//		defineAst(outputDir, "")
//	}
//
//	func defineAst(outputDir string, baseName string, types []string) {
//		path := outputDir + baseName + ".go"
//		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644)
//		f.WriteString()
//	}
package tool
