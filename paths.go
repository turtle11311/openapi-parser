package openapi_parser

type Paths map[string]PathItem

type PathItem struct {
	Ref         string      `yaml:"$ref"`
	Summary     string      `yaml:"summary"`
	Description string      `yaml:"description"`
	GET         *Operation  `yaml:"get"`
	PUT         *Operation  `yaml:"put"`
	POST        *Operation  `yaml:"post"`
	DELETE      *Operation  `yaml:"delete"`
	OPTIONS     *Operation  `yaml:"options"`
	HEAD        *Operation  `yaml:"head"`
	PATCH       *Operation  `yaml:"patch"`
	TRACE       *Operation  `yaml:"trace"`
	Servers     []Server    `yaml:"servers"`
	Parameters  []Parameter `yaml:"parameters"`
}
