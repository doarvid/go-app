package addproto

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/doarvid/go-app/cmd/goapp/internal/config"
	"github.com/doarvid/go-app/cmd/goapp/internal/desc"
	"github.com/doarvid/go-app/cmd/goapp/internal/embeded"
	"github.com/doarvid/go-app/cmd/goapp/internal/pkg/filex"
	"github.com/doarvid/go-app/cmd/goapp/internal/pkg/stringx"
	"github.com/doarvid/go-app/cmd/goapp/internal/pkg/templatex"
)

func Run(args []string) error {
	baseDir := filepath.Join("desc", "proto")

	protoName := args[0]

	if strings.HasSuffix(protoName, ".proto") {
		protoName = strings.TrimSuffix(protoName, ".proto")
	}

	frameType, _ := desc.GetFrameType()
	if frameType == "" {
		frameType = "rpc"
	}

	var template []byte

	template, err := templatex.ParseTemplate(filepath.Join(frameType, "template.proto.tpl"), map[string]any{
		"Package": protoName,
		"Service": stringx.ToCamel(protoName),
	}, embeded.ReadTemplateFile(filepath.Join(frameType, "template.proto.tpl")))
	if err != nil {
		return err
	}

	if config.C.Add.Output == "file" {
		if filex.FileExists(filepath.Join(baseDir, protoName+".proto")) {
			return fmt.Errorf("%s already exists", protoName)
		}

		_ = os.MkdirAll(filepath.Dir(filepath.Join(baseDir, protoName)), 0o755)

		err = os.WriteFile(filepath.Join(baseDir, protoName+".proto"), template, 0o644)
		if err != nil {
			return err
		}
		return nil
	}
	fmt.Println(string(template))
	return nil
}
