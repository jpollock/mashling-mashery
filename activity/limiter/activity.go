package limiter

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("activity-mashery-limiter")

const (
	ivCount = "count"
	ivLimit = "limit"

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
	log.Info("HERE::1")
	count, _ := getIntValue(context, ivCount, 0)
	limit, _ := getIntValue(context, ivLimit, 0)
	log.Info("HERE::2")
	if count > limit {
		log.Info("HERE::3")
		context.SetOutput(ovLimited, true)
	} else {
		log.Info("HERE::4")
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
