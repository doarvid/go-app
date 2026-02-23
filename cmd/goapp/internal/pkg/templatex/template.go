package templatex

import (
	"strings"

	"github.com/doarvid/go-app/core/templatex"

	"github.com/doarvid/go-app/cmd/goapp/internal/config"
)

// ParseTemplate template
func ParseTemplate(name string, data map[string]any, tplT []byte) ([]byte, error) {
	for _, v := range config.C.RegisterTplVal {
		split := strings.Split(v, "=")
		if len(split) == 2 {
			data[split[0]] = split[1]
		}
	}
	return templatex.ParseTemplateWithName(name, data, tplT, templatex.WithFuncMaps(registerFuncMap))
}
