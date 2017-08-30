package reply

import (
	"bytes"
	"encoding/json"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"io/ioutil"
	"net/http"
	"reflect"
)

// log is the default package logger
var log = logger.GetLogger("activity-mashery-reply")

const (
	ivData = "data"
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
	log.Info("REPLY::1")
	code := 200
	data := context.GetInput(ivData)
	log.Info("REPLY::2")
	log.Info(data)
	log.Info(reflect.TypeOf(data))
	//	log.Info("Code :'%d', Data: '%+v'", code, data)

	log.Info("REPLY::3")
	var result interface{}
	resp, ok := data.(*http.Response)
	log.Info("REPLY::4")
	log.Info(ok)
	if ok {
		defer resp.Body.Close()

		log.Info("response Status:", resp.Status)
		respBody, _ := ioutil.ReadAll(resp.Body)

		d := json.NewDecoder(bytes.NewReader(respBody))
		d.UseNumber()
		err = d.Decode(&result)

		json.Unmarshal(respBody, &result)

		log.Info("response Body:", result)
	} else {

	}
	replyHandler := context.FlowDetails().ReplyHandler()
	if replyHandler != nil {
		replyHandler.Reply(code, result, nil)
	}

	return true, nil
}
