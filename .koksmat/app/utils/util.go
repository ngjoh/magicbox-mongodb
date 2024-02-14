package utils

import (
	"os"
	"path"
)

func WorkDir(kitchenname string) string {
	if os.Getenv("WORKDIR") != "" {
		return os.Getenv("WORKDIR")
	}
	return path.Join(os.Getenv("KITCHENROOT"), kitchenname, ".koksmat", "workdir")
}
