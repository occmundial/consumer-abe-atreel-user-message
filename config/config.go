package config

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/occmundial/consumer-abe-atreel-user-message/libs/logger"
	"github.com/occmundial/consumer-abe-atreel-user-message/models"

	"github.com/sethvargo/go-envconfig"
)

// https://medium.com/@felipedutratine/manage-config-in-golang-to-get-variables-from-file-and-env-variables-33d876887152

var (
	_, b, _, _    = runtime.Caller(0)
	basepath      = filepath.Dir(b)
	configuration = models.Configuration{}
	envKeys       = []string{"GO_ENV"}
)

// NewConfiguration :
func NewConfiguration() *models.Configuration {
	loadConfig()
	return &configuration
}

// loadConfig :
func loadConfig() {
	validateEnvironment()
	readConfigFile()
	readConfigEnv()
}

func readConfigFile() {
	logger.LogSimpleDebug(fmt.Sprintf("Environment: %s", os.Getenv("GO_ENV")))
	file, err := os.Open(basepath + "/config." + os.Getenv("GO_ENV") + ".json")

	if err != nil {
		logger.Fatal("config", "readConfigFile", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		logger.Fatal("config", "readConfigFile", err)
	}
}

func readConfigEnv() {
	ctx := context.Background()
	err := envconfig.Process(ctx, &configuration)
	if err != nil {
		logger.Fatal("config", "readConfigEnv", err)
	}
}

func validateEnvironment() {
	if !isValidEnvironment(envKeys...) {
		logger.FatalS("config", "validateEnvironment", fmt.Sprintf("Environment variables must be set: %s", envKeys))
	}
}

func isValidEnvironment(keys ...string) bool {
	for _, key := range keys {
		if len(strings.TrimSpace(os.Getenv(key))) == 0 {
			return false
		}
	}
	return true
}
