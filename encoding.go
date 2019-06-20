package openapi_parser

type Encoding struct {
	ContentType   string            `yaml:"contentType"`
	Headers       map[string]Header `yaml:"headers"`
	Style         string            `yaml:"style"`
	Explode       bool              `yaml:"explode"`
	AllowReserved bool              `yaml:"allowReserved"`
}
