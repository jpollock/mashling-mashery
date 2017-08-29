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
	"unsafe"
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

	ovDeveloperConfiguration = "developerConfiguration"
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
		return true, nil
	}

	apiConfigValue, ok := data.GetGlobalScope().GetAttr("apiConfiguration")
	d := apiConfigValue.Value
	apiConfiguration, ok := d.(models.ApiConfiguration)

	if ok == false {
		log.Info(ok)
	}
	log.Info(apiConfiguration)

	if context.GetInput(ivQueryParams) != nil {
		queryParams := context.GetInput(ivQueryParams).(map[string]string)
		developerConfiguration, found := GetDeveloperConfiguration(context, queryParams, apiConfiguration)
		if unsafe.Sizeof(developerConfiguration) > 0 {
			dt, ok := data.ToTypeEnum("object")
			if ok {
				data.GetGlobalScope().AddAttr("developerConfiguration", dt, developerConfiguration)
			}

		}

		return found, nil

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
