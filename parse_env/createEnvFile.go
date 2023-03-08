package parse_env

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func CreateEnvFile(key interface{}, value interface{}, targetFolder string) {
	log.Println("Creating file for ", key.(string))
	path := filepath.Join(".", ".envs", targetFolder, "."+key.(string)+".env")

	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0777)
	CheckError(err)

	defer file.Close()

	for k, v := range value.(map[string]interface{}) {
		n, err := file.WriteString(strings.ToUpper(k) + "=" + v.(string) + "\n")
		CheckError(err)

		log.Println("Wrote ", n, " bytes to file")
	}

}
