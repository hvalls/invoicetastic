package latextemplate

import (
	"errors"
	"fmt"
	"invoicetastic/util"
	"io"
	"net/http"
	"os"
	"os/exec"
	"text/template"
)

const (
	pdfLatex = "pdflatex"
)

var ErrLaTexNotInstalled = errors.New("TexLive is not installed. Check https://www.latex-project.org/get/")

type LatexTemplate struct {
	t *template.Template
}

func New(tplLocation string) (*LatexTemplate, error) {
	if util.IsURL(tplLocation) {
		return newFromURL(tplLocation)
	}
	return newFromFile(tplLocation)
}

func newFromFile(filePath string) (*LatexTemplate, error) {
	templateFileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return newFromContent(string(templateFileContent))
}

func newFromURL(url string) (*LatexTemplate, error) {
	resp, err := http.Get(url)
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

func (t *LatexTemplate) RenderPDF(fileName string, data any) (string, error) {
	if err := checkLaTexInstalled(); err != nil {
		return "", err
	}
	texFile, err := os.Create(fileName + ".pdf")
	if err != nil {
		return "", err
	}
	defer texFile.Close()
	err = t.t.Execute(texFile, data)
	if err != nil {
		return "", err
	}
	cmd := exec.Command(pdfLatex, texFile.Name())
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("%s\n", output)
		return "", err
	}
	err = os.Remove(fileName + ".aux")
	if err != nil {
		return "", err
	}
	err = os.Remove(fileName + ".log")
	if err != nil {
		return "", err
	}
	return texFile.Name(), err
}

func checkLaTexInstalled() error {
	cmd := exec.Command("command", "-v", pdfLatex)
	err := cmd.Run()
	if err != nil {
		return ErrLaTexNotInstalled
	}
	return nil
}
