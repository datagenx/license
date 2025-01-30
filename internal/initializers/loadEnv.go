package initializers

import (
	"bufio"
	"crypto/ed25519"
	"encoding/base64"
	"log"
	"os"
	"strings"

	"github.com/alecthomas/kong"
)

type Environment struct {
	Env            string `env:"ENV" default:"PROD" help:"It should be PROD or DEV"`
	Port           int    `env:PORT default:9200 help:"Defaulted to 9200"`
	Key            string `env:KEY required:"true" help:"It should be 32 char"`
	PublicKey      string
	PrivateKey     ed25519.PrivateKey
	PostgresEnable bool   `env:"POSTGRES_ENABLE" required:"false" help:"The CircleCI API domain of the test instance."`
	PostgresHost   string `env:"POSTGRES_HOST" required:"false" help:"The CircleCI API domain of the test instance."`
	PostgresPort   int    `env:"POSTGRES_PORT" default:5432 required:"false" help:"The CircleCI API domain of the test instance."`
	PostgresUser   string `env:"POSTGRES_USER" required:"false" help:"The CircleCI API domain of the test instance."`
	PostgresPass   string `env:"POSTGRES_PASS" required:"false" help:"The CircleCI API domain of the test instance."`
	PostgresDB     string `env:"POSTGRES_DB" required:"false" help:"The CircleCI API domain of the test instance."`

	T string `short:"t"`
}

func Setup() *Environment {

	// Loading Env file if CONFIG_FILE is set
	loadEnvVarFile()

	st := &Environment{}
	_ = kong.Parse(st, kong.UsageOnError())
	// generating Public/Private Key
	st.setKeys()

	return st
}

func loadEnvVarFile() {
	if fh, isExists := os.LookupEnv("CONFIG_FILE"); isExists {
		log.Printf("Loading Env variables from config file: %s", fh)
		readFile, err := os.Open(fh)

		if err != nil {
			log.Fatal("Unable to read config file: ", err.Error())
		}
		fileScanner := bufio.NewScanner(readFile)

		fileScanner.Split(bufio.ScanLines)
		for fileScanner.Scan() {
			env := strings.Split(fileScanner.Text(), "=")
			err := os.Setenv(env[0], strings.Trim(env[1], " "))
			if err != nil {
				log.Fatal("Unable to set the env var", err)

			}
		}
	}

}

func (st *Environment) setKeys() {

	// Compute the full 64 byte <private key><public key> from the private key
	if len(st.Key) > 32 {
		log.Fatal("KEY should be 32 character only")
	}
	priv := ed25519.NewKeyFromSeed([]byte(st.Key))

	// Get the public key and base64 encode it
	pub := base64.StdEncoding.EncodeToString(priv.Public().(ed25519.PublicKey))

	// Generating Public/Private Key
	st.PrivateKey, st.PublicKey = priv, pub
}

func (st *Environment) Sign(data []byte) string {
	return base64.StdEncoding.EncodeToString(ed25519.Sign(st.PrivateKey, data))
}
