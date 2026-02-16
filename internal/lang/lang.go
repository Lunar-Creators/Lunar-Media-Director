package lang

import (
	"embed"
	"strings"

	"github.com/BurntSushi/toml"

	"lunar-md/internal/config"
)

//go:embed *.toml
var localeFiles embed.FS

var currentData map[string]interface{}
var currentLang = config.Current.Language

func init() {
	LoadLanguage(currentLang)
}

// LoadLanguage switching language
func LoadLanguage(langCode string) {
	data, _ := localeFiles.ReadFile(langCode + ".toml")
	toml.Unmarshal(data, &currentData)
	currentLang = langCode
}

// Get searching lang key with tile. E.g "main.title"
func Get(path string) string {
	parts := strings.Split(path, ".")
	var v interface{} = currentData

	for _, part := range parts {
		if m, ok := v.(map[string]interface{}); ok {
			v = m[part]
		} else {
			return path // return path if key not found
		}
	}

	if s, ok := v.(string); ok {
		return s
	}
	return path
}