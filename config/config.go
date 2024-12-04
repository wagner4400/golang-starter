package config

import (
	"context"
	"fmt"
	"lawise-go/pkg/aws/credential"
	"lawise-go/pkg/server"
	"log"
	"runtime"
	"strings"
	"sync"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
	"lawise-go/pkg/database"
)

var once sync.Once

const (
	filePathFormat = "%s/config/env/.%s"
)

type Config struct {
	LawiseDb      database.PGConfig
	HttpServer    server.Config
	AWSCredential credential.AWSConfig
}

func GetConfig() *Config {
	var config Config
	once.Do(func() {
		LoadEnvs()

		if err := envconfig.Process(context.Background(), &config); err != nil {
			log.Fatal(err)
		}
	})
	return &config
}

func LoadEnvs() {
	fileName := getFileName()
	_ = godotenv.Load(fileName)
}

func getFileName() string {
	defaultFile := ".env"
	_, file, _, _ := runtime.Caller(1)
	basePath := strings.TrimSuffix(file, "/config/config.go")
	filePath := fmt.Sprintf(filePathFormat, basePath, defaultFile)
	return filePath
}
