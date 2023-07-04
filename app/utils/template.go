package utils

import (
	"fmt"
	hTemplate "html/template"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
)

func DebugPrintLoadTemplate(tmpl *template.Template) {
	var buf strings.Builder
	for _, tmpl := range tmpl.Templates() {
		buf.WriteString("\t- ")
		buf.WriteString(tmpl.Name())
		buf.WriteString("\n")
	}
	DebugPrint("Loaded HTML Templates (%d): \n%s\n", len(tmpl.Templates()), buf.String())
}

func DebugPrint(format string, values ...any) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	fmt.Fprintf(gin.DefaultWriter, "[GIN-debug] "+format, values...)
}

func Filewalk() []string {
	files := []string{}

	filepath.Walk("resources/views", func(path string, info os.FileInfo, err error) error {

		extension := strings.ToLower(filepath.Ext(path))

		if extension == ".html" {
			files = append(files, path)
		}

		return nil
	})

	return files
}

// Custom template function to replace newline characters with <br> tags
func Nl2br(s string) hTemplate.HTML {
	return hTemplate.HTML(strings.ReplaceAll(s, "\n", "<br>"))
}

func In(value any, options []string) bool {
	strValue := fmt.Sprintf("%v", value)
	for _, option := range options {
		if strValue == fmt.Sprintf("%v", option) {
			return true
		}
	}
	return false
}

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func ExplodeFunc(input string, separator string) []string {
	return strings.Split(input, separator)
}
