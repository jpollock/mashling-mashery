package apiconfiguration

import (
	"encoding/json"
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/jpollock/mashling-mashery/models"
)

// log is the default package logger
var log = logger.GetLogger("activity-mashery-api-configuration")

const (
	ivServiceJSON      = "serviceJSON"
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
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *ApiConfigurationActivity) Eval(context activity.Context) (done bool, err error) {
	//	serviceJSON := context.GetInput(ivServiceJSON).(string)
	if context.GetInput(ivServiceJSON) != nil {
		if serviceJSON := context.GetInput(ivServiceJSON).(string); serviceJSON != "" {
			apiConfiguration := getApiConfiguration(serviceJSON)

			dt, ok := data.ToTypeEnum("object")
			if ok {
				data.GetGlobalScope().AddAttr("apiConfiguration", dt, apiConfiguration)
			}

			b, err := json.Marshal(apiConfiguration)
			if err != nil {
				fmt.Println("error:", err)
			}
			context.SetOutput(ovApiConfiguration, string(b))

		}

	}

	eventLogValue, ok := data.GetGlobalScope().GetAttr("eventLog")
	d := eventLogValue.Value
	eventLog, ok := d.(models.EventLog)
	if ok == false {
		log.Info(ok)
	} else {
		log.Info(eventLog.RequestId)
	}

	return true, nil
}

func getApiConfiguration(serviceJSON string) models.ApiConfiguration {
	var c models.ApiConfiguration
	json.Unmarshal([]byte(serviceJSON), &c)
	return c
}
