package utils

import (
	"log"
	"os"
	"os/exec"
)

func InitiazeGit() {
	initiazeGitDirectory()

}

func SyncGit() {
	moduleConfiguration := GetConfiguration()
	cmd := exec.Command("sh", "-c", "git remote add horizon.modules."+moduleConfiguration.Name+" https://gitlab.lelyonline.com/pd/horizon/poc/horizon.modules."+moduleConfiguration.Name+".git && git push horizon.modules."+moduleConfiguration.Name+" master")
	cmd.Dir = getCurrentWorkingDirectory() + "/Horizon.Modules." + moduleConfiguration.Name

	err := cmd.Run()
	if err != nil {
		log.Fatalln("Failed to sync git working directory. => " + err.Error())
	}
}

func initiazeGitDirectory() {
	moduleConfiguration := GetConfiguration()
	cmd := exec.Command("bash", "-c", "git init")
	cmd.Dir = getCurrentWorkingDirectory() + "/Horizon.Modules." + moduleConfiguration.Name
	err := cmd.Run()
	if err != nil {
		log.Fatalln("Failed to initialize git working directory. => " + err.Error())
	}

	cmd = exec.Command("bash", "-c", "git add .")
	cmd.Dir = getCurrentWorkingDirectory() + "/Horizon.Modules." + moduleConfiguration.Name
	err = cmd.Run()
	if err != nil {
		log.Fatalln("Failed to initialize git working directory. => " + err.Error())
	}

	cmd = exec.Command("bash", "-c", "git commit --m \"Initial Commit\"")
	cmd.Dir = getCurrentWorkingDirectory() + "/Horizon.Modules." + moduleConfiguration.Name
	err = cmd.Run()
	if err != nil {
		log.Fatalln("Failed to initialize git working directory. => " + err.Error())
	}
}

func getCurrentWorkingDirectory() string {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	return path
}
