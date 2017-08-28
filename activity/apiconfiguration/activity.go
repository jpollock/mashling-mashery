package apiconfiguration

import (
	"encoding/json"
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/jpollock/mashling-mashery/models"
	"io/ioutil"
	"os"
)

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

	dt, ok := data.ToTypeEnum("object")
	if ok {
		data.GetGlobalScope().AddAttr("apiConfiguration", dt, apiConfiguration)
	}

	b, err := json.Marshal(apiConfiguration)
	if err != nil {
		fmt.Println("error:", err)
	}
	context.SetOutput(ovApiConfiguration, string(b))

	eventLogValue, ok := data.GetGlobalScope().GetAttr("eventLog")
	log.Info(eventLogValue)
	d := eventLogValue.Value
	eventLog, ok := d.(models.EventLog)
	log.Info(eventLog)
	if ok == false {
		log.Info(ok)
	} else {
		log.Info(eventLog.RequestId)
	}

	return true, nil
}

func getApiConfiguration(filePath string) models.ApiConfiguration {
	raw, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c models.ApiConfiguration
	json.Unmarshal(raw, &c)
	return c
}
