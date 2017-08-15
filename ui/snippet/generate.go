//+build ignore

package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/gu-io/gu/assets"
	"github.com/gu-io/gu/assets/packers"
	"github.com/influx6/moz/gen"
)

func main() {
	aspacker := assets.New(packers.RawPacker{})

	aspacker.Register(".js", packers.JSPacker{})
	aspacker.Register(".css", packers.CSSPacker{CleanCSS: true})
	aspacker.Register(".static.html", packers.StaticMarkupPacker{
		PackageName:     "snippet",
		DestinationFile: ".//snippet_static_bundle.go",
	})

	writer, statics, err := aspacker.Compile("./", false)
	if err != nil {
		panic(err)
	}

	pipeGen := gen.Block(
		gen.Package(
			gen.Name("snippet"),
			writer,
		),
	)

	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if err := writeToFile(pipeGen, fmt.Sprintf("%s_bundle.go", "snippet"), "./", currentDir); err != nil {
		panic(err)
	}

	for _, directives := range statics {
		for _, directive := range directives {
			if directive.Static == nil {
				continue
			}

			if err := writeToFile(directive.Writer, directive.Static.FileName, directive.Static.DirName, currentDir); err != nil {
				panic(err)
			}
		}
	}

	fmt.Println("Bundling completed for 'snippet'")
}

// writeToFile writes the giving content from the WriterTo instance to the file of
// the giving file.
func writeToFile(w io.WriterTo, fileName string, dirName string, currentDir string) error {
	coDir := filepath.Join(currentDir, dirName)

	if dirName != "" {
		if _, err := os.Stat(coDir); err != nil {
			if err := os.MkdirAll(coDir, 0700); err != nil && err != os.ErrExist {
				return err
			}

			fmt.Printf("- Created package directory: %q\n", coDir)
		}
	}

	coFile := filepath.Join(coDir, fileName)
	file, err := os.Create(coFile)
	if err != nil {
		return err
	}

	defer file.Close()

	if _, err := w.WriteTo(file); err != nil {
		return err
	}

	fmt.Printf("- Created directory file: %q\n", filepath.Join(dirName, fileName))
	return nil
}
