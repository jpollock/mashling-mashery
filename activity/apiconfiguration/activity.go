package apiconfiguration

import (
	"encoding/json"
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"io/ioutil"
	"os"
)

type ApiConfiguration struct {
	ID              string     `json:"id"`
	Name            string     `json:"name"`
	QpsLimitOverall int        `json:"qpsLimitOverall"`
	Endpoints       []Endpoint `json:"endpoints"`
}

type Endpoint struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Cache Cache  `json:"cache"`
}

type Cache struct {
	CacheTtlOverride int `json:"cacheTtlOverride"`
}

// log is the default package logger
var log = logger.GetLogger("activity-mashery-api-configuration")

const (
	ivFilePath         = "filePath"
	ovApiConfiguration = "apiConfiguration"
)

// CacheActivity is a Cache Activity implementation
type ApiConfigurationActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new CacheActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &ApiConfigurationActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *ApiConfigurationActivity) Metadata() *activity.Metadata {
	log.Info(a.metadata)
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *ApiConfigurationActivity) Eval(context activity.Context) (done bool, err error) {
	filePath := context.GetInput(ivFilePath).(string)
	apiConfiguration := getApiConfiguration(filePath)

	b, err := json.Marshal(apiConfiguration)
	if err != nil {
		fmt.Println("error:", err)
	}
	context.SetOutput(ovApiConfiguration, string(b))
	return true, nil
}

func getApiConfiguration(filePath string) ApiConfiguration {
	raw, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c ApiConfiguration
	json.Unmarshal(raw, &c)
	return c
}
