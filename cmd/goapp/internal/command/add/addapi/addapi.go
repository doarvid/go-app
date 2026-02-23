package addapi

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/zeromicro/go-zero/tools/goctl/api/format"

	"github.com/doarvid/go-app/cmd/goapp/internal/config"
	"github.com/doarvid/go-app/cmd/goapp/internal/desc"
	"github.com/doarvid/go-app/cmd/goapp/internal/embeded"
	"github.com/doarvid/go-app/cmd/goapp/internal/pkg/filex"
	"github.com/doarvid/go-app/cmd/goapp/internal/pkg/templatex"
)

func Run(args []string) error {
	baseDir := filepath.Join("desc", "api")

	service := desc.GetApiServiceName(filepath.Join("desc", "api"))

	apiName := args[0]

	if strings.HasSuffix(apiName, ".api") {
		apiName = strings.TrimSuffix(apiName, ".api")
	}

	if service == "" {
		service = apiName
	}

	template, err := templatex.ParseTemplate(filepath.Join("api", "template.api.tpl"), map[string]any{
		"Service": service,
		"Group":   apiName,
	}, embeded.ReadTemplateFile(filepath.Join("api", "template.api.tpl")))
	if err != nil {
		return err
	}

	if config.C.Add.Output == "file" {
		if filex.FileExists(filepath.Join(baseDir, apiName+".api")) {
			return fmt.Errorf("%s already exists", apiName)
		}

		_ = os.MkdirAll(filepath.Dir(filepath.Join(baseDir, apiName)), 0o755)

		err = os.WriteFile(filepath.Join(baseDir, apiName+".api"), template, 0o644)
		if err != nil {
			return err
		}

		// format
		return format.ApiFormatByPath(filepath.Join(baseDir, apiName+".api"), false)
	}
	fmt.Println(string(template))
	return nil
}
