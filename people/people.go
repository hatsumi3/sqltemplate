package people

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gocarina/gocsv"
)

type Person struct {
	Name   string `csv:"name"`
	Age    string `csv:"age"`
	Height string `csv:"height"`
	Width  string `csv:"width"`
}

var templateText string = `INSERT INTO XXX(Name, Age, Height, Width) VALUES ({{ .Name }}, {{ .Age }}, {{ .Height }}, {{ .Width }});`

func ReadCsvFile(fileName string) []*Person {
	file, err := os.Open(filepath.Join("in", fileName))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	people := []*Person{}
	if err := gocsv.UnmarshalFile(file, &people); err != nil {
		panic(err)
	}
	return people
}

func CreateTemplateString(person *Person) string {
	tmpl, err := template.New("test").Parse(templateText)
	if err != nil {
		panic(err)
	}
	writer := new(strings.Builder)
	if err := tmpl.Execute(writer, person); err != nil {
		panic(err)
	}
	return writer.String()
}

func WriteTemplateFile(fileName string, people []*Person) error {
	f, err := os.Create(filepath.Join("out", fileName))
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	for _, person := range people {
		_, err := fmt.Fprint(f, CreateTemplateString(person)+"\n")
		if err != nil {
			return err
		}
	}
	return nil
}
