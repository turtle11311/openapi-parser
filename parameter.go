package openapi_parser

type Parameter struct {
	Name            string               `yaml:"name"`
	In              string               `yaml:"in"`
	Description     string               `yaml:"description"`
	Required        bool                 `yaml:"required"`
	Deprecated      bool                 `yaml:"deprecated"`
	AllowEmptyValue bool                 `yaml:"allowEmptyValue"`
	Style           string               `yaml:"style"`
	Explode         bool                 `yaml:"explode"`
	AllowReserved   bool                 `yaml:"allowReserved"`
	Schema          Schema               `yaml:"schema"`
	Example         interface{}          `yaml:"example"`
	Examples        map[string]Example   `yaml:"examples"`
	Content         map[string]MediaType `yaml:"content"`
}
