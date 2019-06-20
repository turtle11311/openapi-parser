package openapi_parser

type MediaType struct {
	Schema   Schema              `yaml:"schema"`
	Example  interface{}         `yaml:"example"`
	Examples map[string]Example  `yaml:"examples"`
	Encoding map[string]Encoding `yaml:"encoding"`
}
