package entfga

// Annotations of the fga extension
type Annotations struct {
	ObjectType   string `yaml:"ObjectType,omitempty"`   // Object type for the fga relationship
	IncludeHooks bool   `yaml:"includeHooks,omitempty"` // Include hooks for the fga extension to add tuples to FGA
	IDField      string `yaml:"idField,omitempty"`      // ID field for the object type
}

// Name of the annotation
func (Annotations) Name() string {
	return "Authz"
}
