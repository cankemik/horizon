package module

import (
	"errors"
	"fmt"
	"horizon/structs"
	"horizon/utils"
	"os"
	"regexp"
	"strings"

	"github.com/manifoldco/promptui"
)

var moduleConfiguration structs.ModuleConfiguration

func Create() {

	initializeConfiguration()

	fmt.Println("\U0001F9D9 Crafting a module named " + moduleConfiguration.Name)

	initializeWorkingDirectory()

	prepareTemplate(moduleConfiguration)

	fmt.Println("\U0001F9D9 Hold on. Casting some spells..")

	utils.InitiazeGit()

	setRepository()

	utils.SyncGit()

	fmt.Println("\U0001F9D9 " + moduleConfiguration.Name + " Module crafted ! ")

}

func initializeConfiguration() {

	setName()

	setType()

	utils.SetConfiguration(moduleConfiguration)
}

func initializeWorkingDirectory() {
	utils.CreateModuleDirectory(moduleConfiguration.Name)

	utils.SaveModuleConfigurationFile(moduleConfiguration)
}

func setName() {
	if len(os.Args) <= 2 {
		panic("Please define a module name. Ex: horizon start <moduleName>")
	}
	moduleConfiguration.Name = os.Args[2]
}

func setRepository() {
	var gitRepositoryRegex = regexp.MustCompile(`((git|ssh|http(s)?)|(git@[\w\.]+))(:(//)?)([\w\.@\:/\-~]+)(\.git)(/)?`)
	validate := func(input string) error {
		if !gitRepositoryRegex.MatchString(input) {
			return errors.New("Invalid repository URL.")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Git Remote Repository URL",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	moduleConfiguration.Repository = result
}

func setType() {
	moduleTypes := []structs.ModuleType{
		{Name: "REST API", Repository: "https://t4csharedfiles.blob.core.windows.net/horizon/APIBoilerplate.zip?sv=2020-04-08&st=2021-11-21T20%3A01%3A39Z&se=2030-11-22T20%3A01%3A00Z&sr=b&sp=r&sig=ve7PJnCyspitB3KELBzxosBub7r5%2FTFOjqXXR4mO%2FfU%3D", Description: "Creates a .net worker service REST API boilerplate with sample DI, controller and swagger documentation."},
		{Name: "Console Application", Repository: "https://t4csharedfiles.blob.core.windows.net/horizon/APIBoilerplate.zip?sv=2020-04-08&st=2021-11-21T20%3A01%3A39Z&se=2030-11-22T20%3A01%3A00Z&sr=b&sp=r&sig=ve7PJnCyspitB3KELBzxosBub7r5%2FTFOjqXXR4mO%2FfU%3D", Description: "Creates a .net worker service console application boilerplate with sample DI and scheduled job"},
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }} ?",
		Active:   "\U0001f9e9 {{ .Name | cyan }}",
		Inactive: "  {{ .Name | cyan }}",
		Selected: "\U0001f9e9 {{ .Name | red | cyan }}",
		Details: `
--------- Available Module Types ----------
{{ "Module:" | faint }}	{{ .Name }}
{{ "Description:" | faint }}	{{ .Description }}`,
	}

	searcher := func(input string, index int) bool {
		moduleType := moduleTypes[index]
		name := strings.Replace(strings.ToLower(moduleType.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Module Selection",
		Items:     moduleTypes,
		Templates: templates,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	moduleConfiguration.Type = moduleTypes[i]

	fmt.Printf("Selection:  %s\n", moduleTypes[i].Name)
}

func prepareTemplate(moduleConfiguration structs.ModuleConfiguration) {
	utils.CloneTemplate(moduleConfiguration.Type.Repository)
}
