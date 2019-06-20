package openapi_parser

type RequestBody struct {
	Description string               `yaml:"description"`
	Content     map[string]MediaType `yaml:"content"`
	Required    bool                 `yaml:"required"`
}
