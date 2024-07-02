package latextemplate

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"text/template"
)

const DefaultTemplateURL = "https://raw.githubusercontent.com/hvalls/invoicetastic/main/_templates/english-usd.tex"

type LatexTemplate struct {
	t *template.Template
}

func New(filePath string) (*LatexTemplate, error) {
	if filePath == "" {
		return newDefault()
	}
	return newFromFile(filePath)
}

func newFromFile(filePath string) (*LatexTemplate, error) {
	templateFileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return newFromContent(string(templateFileContent))
}

func newDefault() (*LatexTemplate, error) {
	resp, err := http.Get(DefaultTemplateURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return newFromContent(string(body))
}

func newFromContent(content string) (*LatexTemplate, error) {
	tmpl, err := template.New("latex").Parse(string(content))
	if err != nil {
		return nil, err
	}
	return &LatexTemplate{tmpl}, nil
}

func (t *LatexTemplate) Render(fileName string, data any) error {
	texFile, err := os.Create(fileName + ".pdf")
	if err != nil {
		panic(err)
	}
	defer texFile.Close()
	err = t.t.Execute(texFile, data)
	if err != nil {
		return err
	}
	cmd := exec.Command("pdflatex", texFile.Name())
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("%s\n", output)
		return err
	}
	err = os.Remove(fileName + ".aux")
	if err != nil {
		panic(err)
	}
	err = os.Remove(fileName + ".log")
	if err != nil {
		panic(err)
	}
	return err
}
