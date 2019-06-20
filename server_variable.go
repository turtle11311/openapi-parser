package openapi_parser

type ServerVariable struct {
	Enum        []string `yaml:"enum"`
	Default     string   `yaml:"default"`
	Description string   `yaml:"description"`
}
