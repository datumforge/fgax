package entfga

import (
	"strings"
	"text/template"

	"entgo.io/ent/entc/gen"
	"github.com/stoewer/go-strcase"
)

// extractObjectType gets the key that is used for the object type
func extractObjectType(val any) string {
	objectType, ok := val.(string)
	if !ok {
		return ""
	}

	return objectType
}

// extractIDField gets the key that is used for the id field
func extractIDField(val any) string {
	idField, ok := val.(string)
	if !ok {
		return ""
	}

	if idField == "" {
		return "ID"
	}

	return idField
}

// hasCreateID checks if the input would have the ID to check permissions
func hasCreateID(val any) bool {
	idField, ok := val.(string)
	if !ok {
		return false
	}

	if idField == "" || idField == "ID" {
		return false
	}

	return true
}

// extractIncludeHooks gets the key that is used to determine if the hooks should be included
func extractIncludeHooks(val any) bool {
	includeHooks, ok := val.(bool)
	if !ok {
		return true
	}

	return includeHooks
}

// useSoftDeletes checks the config properties for the Soft Delete setting
func useSoftDeletes(config Config) bool {
	return config.SoftDeletes
}

// parseTemplate parses the template and sets values in the template
func parseTemplate(name, path string) *gen.Template {
	t := gen.NewTemplate(name)
	t.Funcs(template.FuncMap{
		"extractObjectType":   extractObjectType,
		"extractIDField":      extractIDField,
		"hasCreateID":         hasCreateID,
		"extractIncludeHooks": extractIncludeHooks,
		"useSoftDeletes":      useSoftDeletes,
		"ToUpperCamel":        strcase.UpperCamelCase,
		"ToLower":             strings.ToLower,
	})

	return gen.MustParse(t.ParseFS(_templates, path))
}
