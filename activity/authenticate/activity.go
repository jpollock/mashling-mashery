package authenticate

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// log is the default package logger
var log = logger.GetLogger("activity-mashery-authenticate")

const (
	ivActivityEnabled = "activityEnabled"
	ovAuthenticated   = "authenticated"
)

// CacheActivity is a Cache Activity implementation
type AuthenticateActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new CacheActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &AuthenticateActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *AuthenticateActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *AuthenticateActivity) Eval(context activity.Context) (done bool, err error) {

	activityEnabled := false

	if context.GetInput(ivActivityEnabled) != nil {
		activityEnabled = context.GetInput(ivActivityEnabled).(bool)
	}

	context.SetOutput(ovAuthenticated, true)
	if activityEnabled == false {
		return true, nil
	}

	return true, nil

}
