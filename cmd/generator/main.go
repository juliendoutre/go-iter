package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

var (
	out              string
	pkg              string
	elementTypes     string
	accumulatorTypes string

	config = map[string]func(elements []string, accumulators []string) string{
		"iterator.go": func(elements []string, accumulators []string) string {
			return fmt.Sprintf("Element=%s", strings.Join(elements, ","))
		},
		"iterators.go": func(elements []string, accumulators []string) string {
			return fmt.Sprintf("Element=%s", strings.Join(elements, ","))
		},
		"option.go": func(elements []string, accumulators []string) string {
			return fmt.Sprintf("Element=%s", strings.Join(append(elements, "uint"), ","))
		},
		"folding.go": func(elements []string, accumulators []string) string {
			types := append([]string{"uint", "Empty"}, elements...)
			for _, element := range elements {
				types = append(types, fmt.Sprintf("OptionFor%s", strings.Title(element)))
			}

			return fmt.Sprintf(
				"Element=%s Accumulator=%s",
				strings.Join(elements, ","),
				strings.Join(removeDuplicates(append(accumulators, types...)), ","),
			)
		},
		"vector.go": func(elements []string, accumulators []string) string {
			return fmt.Sprintf("Element=%s", strings.Join(elements, ","))
		},
	}
)

func main() {
	flag.StringVar(&out, "out", ".", "path where to write the generated files")
	flag.StringVar(&pkg, "pkg", "iter", "Name of the package to be used by generated files")
	flag.StringVar(&elementTypes, "items", "int", "comma separated types for the items of iterators")
	flag.StringVar(&accumulatorTypes, "acc", "int", "comma separated types for the accumulators over iterators")
	flag.Parse()

	in := path.Join(os.Getenv("GOPATH"), "src", "github.com", "juliendoutre", "go-iter", "pkg", "templates")

	elements := removeDuplicates(strings.Split(elementTypes, ","))
	accumulators := removeDuplicates(strings.Split(accumulatorTypes, ","))

	for file, generateExpression := range config {
		if err := genny(path.Join(in, file), path.Join(out, file), pkg, generateExpression(elements, accumulators)); err != nil {
			log.Fatal(err)
		}
	}

	for _, file := range []string{"types.go", "range.go"} {
		if err := copy(path.Join(in, file), path.Join(out, file), pkg); err != nil {
			log.Fatal(err)
		}
	}
}

func removeDuplicates(data []string) []string {
	cache := map[string]struct{}{}
	for _, entry := range data {
		cache[entry] = struct{}{}
	}

	output := []string{}
	for key := range cache {
		output = append(output, key)
	}

	return output
}

func genny(in, out, pkg, types string) error {
	cmd := exec.Command("genny", "-in", in, "-out", out, "-pkg", pkg, "gen", types)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func copy(in, out, pkg string) error {
	r, err := os.Open(in)
	if err != nil {
		return err
	}
	defer r.Close()

	w, err := os.Create(out)
	if err != nil {
		return err
	}
	defer w.Close()

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := fmt.Sprintln(scanner.Text())
		if strings.HasPrefix(line, "package") {
			if _, err := w.Write([]byte(fmt.Sprintf("package %s\n", pkg))); err != nil {
				return err
			}
		} else {
			if _, err := w.Write([]byte(line)); err != nil {
				return err
			}
		}
	}

	return scanner.Err()
}
