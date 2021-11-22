package structs

type ModuleType struct {
	Id          string
	Name        string
	Repository  string
	Description string
}

type ModuleConfigurationFile struct {
	Name           string
	Id             string
	Type           ModuleType
	ExeProjectPath string
}

type ModuleConfiguration struct {
	Name       string
	Type       ModuleType
	Repository string
}
