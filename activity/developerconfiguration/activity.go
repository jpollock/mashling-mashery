package developerconfiguration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/go-redis/redis"
	"github.com/jpollock/mashling-mashery/models"
)

// log is the default package logger
var log = logger.GetLogger("activity-mashery-developer-configuration")

const (
	ivActivityEnabled = "activityEnabled"
	ivURI             = "uri"
	ivPathParams      = "pathParams"
	ivQueryParams     = "queryParams"
	ivContent         = "content"
	ivRedisAddress    = "redisAddress"

	ovError     = "error"
	ovErrorData = "errorData"
)

// CacheActivity is a Cache Activity implementation
type DeveloperConfigurationActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new CacheActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &DeveloperConfigurationActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *DeveloperConfigurationActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *DeveloperConfigurationActivity) Eval(context activity.Context) (done bool, err error) {
	activityEnabled := false
	if context.GetInput(ivActivityEnabled) != nil {
		activityEnabled = context.GetInput(ivActivityEnabled).(bool)
	}
	if activityEnabled == false {
		log.Info("Not enabled")
		context.SetOutput(ovError, false)
		context.SetOutput(ovErrorData, nil)
		return true, nil
	}

	apiConfigValue, ok := data.GetGlobalScope().GetAttr("apiConfiguration")
	log.Info(apiConfigValue)
	d := apiConfigValue.Value
	apiConfiguration, ok := d.(models.ApiConfiguration)

	if ok == false {
		log.Info(ok)
	}

	var developerConfiguration models.DeveloperConfiguration
	found := false
	if context.GetInput(ivQueryParams) != nil {
		queryParams := context.GetInput(ivQueryParams).(map[string]string)
		developerConfiguration, found = GetDeveloperConfiguration(context, queryParams, apiConfiguration)
		if (developerConfiguration != models.DeveloperConfiguration{}) {
			found = true
			dt, ok := data.ToTypeEnum("object")
			if ok {
				data.GetGlobalScope().AddAttr("developerConfiguration", dt, developerConfiguration)
			}

		}
	}
	log.Info(found)
	if found == false {
		log.Info("Not found")
		context.SetOutput(ovError, true)
		errorData := activity.NewError("test", "403", nil)
		log.Info(errorData.Code())
		//errorData.errorStr = "test"
		//Error{errorStr: errorText, errorData: errorData, errorCode: code}
		dt, ok := data.ToTypeEnum("object")
		if ok {
			log.Info("Adding error to global")
			data.GetGlobalScope().AddAttr("error", dt, errorData)
		}

	} else {
		context.SetOutput(ovError, false)
		context.SetOutput(ovErrorData, nil)

	}

	return true, nil

}

func GetDeveloperConfiguration(context activity.Context, queryParams map[string]string, apiConfiguration models.ApiConfiguration) (developerConfiguration models.DeveloperConfiguration, found bool) {
	var s models.DeveloperConfiguration
	log.Info(apiConfiguration.Endpoints[0].ApiKeyValueLocationKey)
	for key, value := range queryParams {

		if key == apiConfiguration.Endpoints[0].ApiKeyValueLocationKey {
			redisAddress := context.GetInput(ivRedisAddress).(string)
			client := redis.NewClient(&redis.Options{
				Addr:     redisAddress,
				Password: "", // no password set
				DB:       0,  // use default DB
			})

			var cache_key bytes.Buffer
			cache_key.WriteString(apiConfiguration.ID)
			cache_key.WriteString("_")
			cache_key.WriteString(value)

			val2, err := client.Get(cache_key.String()).Result()
			if err == redis.Nil {
				fmt.Println("key2", val2)
			} else if err != nil {
				panic(err)
			} else {
				found = true
				json.Unmarshal([]byte(val2), &s)

			}
		}
	}

	return s, found
}
