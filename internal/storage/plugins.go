package storage

import (
	"database/sql"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/datagenx/license-generator/internal/generate"
)

type Storage interface {
	Save(data string) error
	ReadAll() ([]string, error)
}

func Plugins(sl generate.Slic, data string) error {

	// Initialize file
	if filePath, isExists := os.LookupEnv("FILE_PATH"); isExists && filePath != "" {
		log.Println("file plugin is enabled")

		fp := FileStorage{filePath}
		return fp.Save(sl, data)

	}

	// Initialize the mongo
	if mongoConnStr, isExists := os.LookupEnv("MONGODB_CONN_STRING"); isExists && mongoConnStr != "" {
		log.Println("mongodb plugin is enabled")
		getMongoDBClient(mongoConnStr)
	}

	// Initialize the postgres
	if postgresConnStr, isExists := os.LookupEnv("POSTGRES_CONN_STRING"); isExists && postgresConnStr != "" {
		log.Println("postgres plugin is enabled")
		getPostgresClient(postgresConnStr)
	}

	return nil

}

func fileInitialize(filePath string) (*os.File, error) {

	fh, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0640)
	defer fh.Close()
	return fh, err
}

func getMongoDBClient(mongoConnStr string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(mongoConnStr)
	return mongo.NewClient(clientOptions)
}

func getPostgresClient(postgresConnStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", postgresConnStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return db, nil
}
