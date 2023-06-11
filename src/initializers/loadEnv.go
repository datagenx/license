package initializers

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func LoadEnvVar() {

	readFile, err := os.Open("./config")

	if err != nil {
		log.Fatal("Unable to read config file", err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		env := strings.Split(fileScanner.Text(), "=")
		// setting env var
		err := os.Setenv(env[0], env[1])
		if err != nil {
			log.Fatal("Unable to set the env var", err)
		}
	}
	readFile.Close()
}
