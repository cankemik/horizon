package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"horizon/structs"
	"io/ioutil"
	"os"
	"strings"
)

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func isDirectory(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}

	return fileInfo.IsDir()
}

func copyFile(sourcePath string, targetPath string) {
	input, err := ioutil.ReadFile(sourcePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	moduleConfiguration := GetConfiguration()
	output := bytes.ReplaceAll(input, []byte("Boilerplate"), []byte(moduleConfiguration.Name))
	output = bytes.ReplaceAll(output, []byte("boilerplate"), []byte(strings.ToLower(moduleConfiguration.Name)))
	err = ioutil.WriteFile(targetPath, output, 0644)
	if err != nil {
		fmt.Println("Error creating", targetPath)
		fmt.Println(err)
		return
	}
}

func moveFile(path string) bool {
	path = "./" + path
	moduleConfiguration := GetConfiguration()
	correctedPath := strings.ReplaceAll(path, "Boilerplate", moduleConfiguration.Name)
	if fileExists(correctedPath) || !strings.Contains(path, "Boilerplate") {
		return false
	}

	if isDirectory(path) {
		err := os.Mkdir(correctedPath, 0755)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		copyFile(path, correctedPath)
	}
	return true
}

func CreateModuleDirectory(name string) {
	err := os.Mkdir("Horizon.Modules."+name, 0755)
	if err != nil {
		fmt.Println(err)
	}
}

func SaveModuleConfigurationFile(moduleConfiguration structs.ModuleConfiguration) {
	configurationFile := structs.ModuleConfigurationFile{
		Id:             CreateUUID(),
		Name:           moduleConfiguration.Name,
		Type:           moduleConfiguration.Type,
		ExeProjectPath: "Horizon.Modules." + moduleConfiguration.Name + "/Infrastructure/Horizon.Modules." + moduleConfiguration.Name + ".API.csproj",
	}

	configurationFileBytes, err := json.Marshal(configurationFile)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./Horizon.Modules."+moduleConfiguration.Name+"/horizon.module.json", configurationFileBytes, 0644)
	if err != nil {
		panic(err)
	}
}
