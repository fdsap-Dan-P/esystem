package file

import (
	"path/filepath"
	"strings"
)

func FileSpecs(fileName string) (string, string, string) {
	if fileName == "" {
		return "", "", ""
	}
	ext := strings.ToLower(filepath.Ext(fileName))
	name := filepath.Base(fileName)
	return fileName[:len(fileName)-len(name)], name[:len(name)-len(ext)], ext
}
