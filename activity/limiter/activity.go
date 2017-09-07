package limiter

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/go-redis/redis"
	"github.com/jpollock/mashling-mashery/models"
	"strconv"
)

var log = logger.GetLogger("activity-mashery-limiter")

const (
	ivActivityEnabled = "activityEnabled"
	ivCount           = "count"
	ivLimit           = "limit"
	ivRedisAddress    = "redisAddress"

	ovLimited = "limited"
)

// LimiterActivity is an Activity that is used to check a count against a limit
// and return 403 if count is above limit
// inputs : {count, limit}
// outputs: result
type LimiterActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new AppActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &LimiterActivity{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *LimiterActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *LimiterActivity) Eval(context activity.Context) (done bool, err error) {
	// Get cache
	//redisAddress := context.GetInput(ivRedisAddress).(string)
	//cacheClient := getCache(redisAddress)
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
	limit := apiConfiguration.QpsLimitOverall
	count, _ := getIntValue(context, ivCount, 0)
	//limit, _ := getIntValue(context, ivLimit, 0)
	if count > limit {
		context.SetOutput(ovLimited, true)

		errorData := activity.NewError("Service Over Limit", "403", nil)
		//errorData.errorStr = "test"
		//Error{errorStr: errorText, errorData: errorData, errorCode: code}
		dt, ok := data.ToTypeEnum("object")
		if ok {
			log.Info("Adding error to global")
			data.GetGlobalScope().AddAttr("error", dt, errorData)
		}

	} else {
		context.SetOutput(ovLimited, false)
	}

	return true, nil
}

func getIntValue(context activity.Context, attrName string, defValue interface{}) (int, bool) {

	val := context.GetInput(attrName)
	found := true

	if val == nil {
		found = false

		if defValue == nil {
			return 0, false
		}
		val = defValue
	}

	return val.(int), found
}

func getCurrentCount(key string, client redis.Client) (int, bool) {
	val, err := client.Get(key).Result()
	if err == redis.Nil {
		return 0, false
	} else if err != nil {
		panic(err)
	} else {
		intVal, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		} else {
			return intVal, true
		}

	}
}

func getCache(redisAddress string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client

}
