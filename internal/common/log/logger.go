package log

import (
	"encoding/json"

	"go.uber.org/zap"
)

func InitLogger(LogLevel string, LogOutPutPath string, LogErrorOutputPath string) (*zap.Logger, error) {

	rawJSON := []byte(`{
		"level": "` + LogLevel + `",
		"encoding": "json",
		"outputPaths": ["` + LogOutPutPath + `"],
		"errorOutputPaths": ["` + LogErrorOutputPath + `"],
		"encoderConfig": {
		  "nameKey": "name",
		  "messageKey": "message",
		  "levelKey": "level",
		  "levelEncoder": "lowercase",
		  "timeKey": "datetime",
		  "timeEncoder": "ISO8601"
		}
	  }`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		return nil, err
	}

	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}
