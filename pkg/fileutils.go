package pkg

import (
	"fmt"
	"path/filepath"
	"strings"
)

func GenerateNewFilename(path, operation, format string) string {
	ext := filepath.Ext(path)
	base := strings.TrimSuffix(path, ext)
	return fmt.Sprintf("%s-%s.%s", base, operation, format)
}
