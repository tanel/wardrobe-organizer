package template

import (
	"fmt"
	"html/template"
	"io"
	"path/filepath"
	"time"

	"github.com/juju/errors"
	"github.com/tanel/webapp/configuration"
)

// Render renders a template with given data
func Render(w io.Writer, templateName string, data interface{}) error {
	list, err := filepath.Glob(configuration.SharedInstance.TemplatePath)
	if err != nil {
		return errors.Annotate(err, "globbing templates failed")
	}

	if len(list) == 0 {
		return fmt.Errorf("no templates found: %s", configuration.SharedInstance.TemplatePath)
	}

	funcMap := template.FuncMap{
		"formatDate": formatDate,
	}

	t, err := template.New("").Funcs(funcMap).ParseFiles(list...)
	if err != nil {
		return errors.Annotate(err, "parsing templates failed")
	}

	if err := t.ExecuteTemplate(w, templateName, data); err != nil {
		return errors.Annotate(err, "executing template failed")
	}

	return nil
}

func formatDate(value time.Time) string {
	return value.Format("Mon Jan 2 15:04")
}
