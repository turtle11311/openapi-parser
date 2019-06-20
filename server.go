package openapi_parser

type Server struct {
	Url         string                    `yaml:"url"`
	Description string                    `yaml:"description"`
	Variables   map[string]ServerVariable `yaml:"variables"`
}
