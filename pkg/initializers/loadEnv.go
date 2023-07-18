package initializers

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func LoadEnvVar() {

	readFile, err := os.Open("./config/config.yml")

	if err != nil {
		log.Fatal("Unable to read config file", err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		env := strings.Split(fileScanner.Text(), ":")
		// setting env var, if not set
		if _, isExists := os.LookupEnv(env[0]); !isExists {
			err := os.Setenv(env[0], strings.Trim(env[1], " "))
			if err != nil {
				log.Fatal("Unable to set the env var", err)
			}
		}
	}
}
