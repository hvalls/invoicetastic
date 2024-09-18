package file

import (
	"invoicetastic/util"
	"io"
	"net/http"
	"os"
)

// Read content from filepath or URL
func ReadContent(location string) (string, error) {
	if util.IsURL(location) {
		resp, err := http.Get(location)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		return string(body), nil
	} else {
		fileContent, err := os.ReadFile(location)
		if err != nil {
			return "", err
		}
		return string(fileContent), nil
	}
}

func WriteContent(filename, header string, content any) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	text, err := util.MarshalYAML(content)
	if err != nil {
		return err
	}

	fileFirstLine := "# " + header + "\n"
	if header == "" {
		fileFirstLine = ""
	}

	_, err = file.Write([]byte(fileFirstLine + text))
	if err != nil {
		return err
	}
	return file.Close()
}
