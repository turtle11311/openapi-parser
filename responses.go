package openapi_parser

type Responses map[string]Response

type Response struct {
	Description string               `yaml:"description"`
	Headers     map[string]Header    `yaml:"headers"`
	Content     map[string]MediaType `yaml:"content"`
	Links       map[string]Link
}
