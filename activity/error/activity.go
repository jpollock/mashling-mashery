package error

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/jpollock/mashling-mashery/models"
	"strconv"
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

	errorDataValue, ok := data.GetGlobalScope().GetAttr("error")
	d := errorDataValue.Value
	errorData, ok := d.(*activity.Error)

	if ok == false {
		log.Info(ok)
	}

	replyHandler := context.FlowDetails().ReplyHandler()

	//todo support replying with error

	var eventLog models.EventLog
	eventLogValue, ok := data.GetGlobalScope().GetAttr("eventLog")
	t_eventLog := eventLogValue.Value
	eventLog, ok = t_eventLog.(models.EventLog)

	eventLog.Status = errorData.Code()

	dt_eventLog, ok := data.ToTypeEnum("object")
	if ok {
		data.GetGlobalScope().AddAttr("eventLog", dt_eventLog, eventLog)
	}

	if replyHandler != nil {

		if code, err := strconv.Atoi(errorData.Code()); err == nil {
			replyHandler.Reply(code, errorData.Error(), nil)
		}

	}

	return true, nil
}
