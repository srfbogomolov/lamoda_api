package api_test

import (
	jsonparse "encoding/json"

	"go.uber.org/zap"
)

var testLogger *zap.SugaredLogger

func init() {
	testLogger = createTestLogger()
}

func createTestLogger() *zap.SugaredLogger {
	cfgJSON := []byte(`{
		"level": "info",
		"encoding": "json"
	  }`)
	var cfg zap.Config
	if err := jsonparse.Unmarshal(cfgJSON, &cfg); err != nil {
		panic(err)
	}
	return zap.Must(cfg.Build()).Sugar()
}
