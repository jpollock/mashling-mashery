package reply

import (
	"bytes"
	"encoding/json"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/jpollock/mashling-mashery/models"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
)

// log is the default package logger
var log = logger.GetLogger("activity-mashery-reply")

const (
	ivData    = "data"
	ovContent = "content"
)

// ReplyActivity is an Activity that is used to reply via the trigger
// inputs : {method,uri,params}
// outputs: {result}
type ReplyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new ReplyActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &ReplyActivity{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *ReplyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - Invokes a REST Operation
func (a *ReplyActivity) Eval(context activity.Context) (done bool, err error) {
	var result interface{}
	eventLogValue, ok := data.GetGlobalScope().GetAttr("eventLog")
	t_eventLog := eventLogValue.Value
	eventLog, ok := t_eventLog.(*models.EventLog)

	responseData := context.GetInput(ivData)

	if reflect.TypeOf(responseData).Name() == "string" {

		d := json.NewDecoder(bytes.NewReader([]byte(responseData.(string))))
		d.UseNumber()
		err = d.Decode(&result)

		replyHandler := context.FlowDetails().ReplyHandler()
		if replyHandler != nil {
			replyHandler.Reply(200, result, nil)
		}
	} else {

		var result interface{}
		resp, ok := responseData.(*http.Response)
		if ok {
			defer resp.Body.Close()

			eventLog.Status = strconv.Itoa(resp.StatusCode)
			eventLog.Bytes = resp.ContentLength
			respBody, _ := ioutil.ReadAll(resp.Body)
			bytesBody := bytes.NewReader(respBody)
			d := json.NewDecoder(bytesBody)
			d.UseNumber()
			err = d.Decode(&result)

			//json.Unmarshal(respBody, &result)
			//content := string(result[:])
			context.SetOutput(ovContent, result)
			replyHandler := context.FlowDetails().ReplyHandler()
			if replyHandler != nil {
				replyHandler.Reply(resp.StatusCode, result, nil)
			}

		} else {
			errorData := activity.NewError("Server Error", "500", nil)
			//errorData.errorStr = "test"
			//Error{errorStr: errorText, errorData: errorData, errorCode: code}
			dt, ok := data.ToTypeEnum("object")
			if ok {
				data.GetGlobalScope().AddAttr("error", dt, errorData)
			}

		}
	}
	dt_eventLog, ok := data.ToTypeEnum("object")
	if ok {
		data.GetGlobalScope().AddAttr("eventLog", dt_eventLog, eventLog)
	}

	return true, nil
}
