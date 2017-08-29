package error

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("activity-mashery-error")

const (
	ivCode    = "code"
	ivStatus  = "status"
	ivMessage = "message"

	ovResult = "result"
)

// ErrorActivity is an Activity that responds with HTTP error
// inputs : {code, status, message}
// outputs: result
type ErrorActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new AppActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &ErrorActivity{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *ErrorActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *ErrorActivity) Eval(context activity.Context) (done bool, err error) {
	/*code, _ := context.GetInput(ivCode).(string)
	status, _ := context.GetInput(ivStatus).(string)
	message, _ := context.GetInput(ivMessage).(string)
	*/

	replyHandler := context.FlowDetails().ReplyHandler()

	//todo support replying with error

	if replyHandler != nil {
		replyHandler.Reply(403, "ERROR!!!!!!!!!", nil)
	}

	return true, nil
}
