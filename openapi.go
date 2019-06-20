package openapi_parser

type OpenAPI struct {
	OpenAPI      string                `yaml:"openapi"`
	Info         Info                  `yaml:"info"`
	Servers      []Server              `yaml:"servers"`
	Paths        Paths                 `yaml:"paths"`
	Components   Components            `yaml:"components"`
	Security     SecurityRequirement   `yaml:"security"`
	Tags         []Tag                 `yaml:"tags"`
	ExternalDocs ExternalDocumentation `yaml:"externalDocs"`
}
