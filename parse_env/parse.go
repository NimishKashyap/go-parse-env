package parse_env

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal("Error: ", err)
	}
}

func convert(dirPath string, file os.DirEntry) {
	log.Println(file.Name())

	yamlPath := filepath.Join(".", dirPath, file.Name())
	data := make(map[interface{}]interface{})

	yamlFile, err := os.ReadFile(yamlPath)
	CheckError(err)

	CheckError(yaml.Unmarshal(yamlFile, &data))
	environmentName := strings.Split(file.Name(), ".")[0]

	targetFolder := filepath.Join(".", ".envs", environmentName)
	log.Println("Target folder: ", targetFolder)

	CheckError(os.MkdirAll(targetFolder, 0777))

	for key, value := range data {
		CreateEnvFile(key, value, environmentName)
	}

}

func Parse(dirPath string) {
	files, err := os.ReadDir(dirPath)

	if err != nil {
		log.Fatal("Cannot read dir: Is the directory path correct? ", err)
	}

	for _, file := range files {
		convert(dirPath, file)
	}
}
