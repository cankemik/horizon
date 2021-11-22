package utils

import "horizon/structs"

var configuration structs.ModuleConfiguration

func SetConfiguration(moduleConfiguration structs.ModuleConfiguration) {
	configuration = moduleConfiguration
}

func GetConfiguration() structs.ModuleConfiguration {
	return configuration
}
