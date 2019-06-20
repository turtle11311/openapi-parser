package openapi_parser

type Operation struct {
	Tags         []string              `yaml:"tags"`
	Summary      string                `yaml:"summary"`
	Description  string                `yaml:"description"`
	ExternalDocs ExternalDocumentation `yaml:"externalDocs"`
	OperationID  string                `yaml:"operationId"`
	Parameters   []Parameter           `yaml:"parameters"`
	RequestBody  RequestBody           `yaml:"requestBody"`
	Responses    Responses             `yaml:"responses"`
	Callbacks    map[string]Callback   `yaml:"callbacks"`
	Deprecated   bool                  `yaml:"deprecated"`
	Security     []SecurityRequirement `yaml:"security"`
	Servers      []Server              `yaml:"servers"`
}
