package test

import (
	"bytes"
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/jpollock/mashling-mashery/models"
	"strconv"
	"time"
)

// activityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-tibco-log-2")

const (
	ivMessage   = "message"
	ivFlowInfo  = "flowInfo"
	ivAddToFlow = "addToFlow"

	ovMessage = "message"
)

func init() {
	activityLog.SetLogLevel(logger.InfoLevel)
}

// LogActivity is an Activity that is used to log a message to the console
// inputs : {message, flowInfo}
// outputs: none
type LogActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new AppActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &LogActivity{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *LogActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *LogActivity) Eval(context activity.Context) (done bool, err error) {
	//mv := context.GetInput(ivMessage)
	apiConfigValue, ok := data.GetGlobalScope().GetAttr("apiConfiguration")
	activityLog.Info(ok)
	d := apiConfigValue.Value
	apiConfiguration, ok := d.(models.ApiConfiguration)
	activityLog.Info(apiConfiguration)

	var developerConfiguration models.DeveloperConfiguration
	developerConfigValue, ok := data.GetGlobalScope().GetAttr("developerConfiguration")
	activityLog.Info(ok)
	if ok {
		d3 := developerConfigValue.Value
		developerConfiguration, ok = d3.(models.DeveloperConfiguration)
		activityLog.Info(developerConfiguration)

	} else {
		developerConfiguration2 := new(models.DeveloperConfiguration)
		developerConfiguration = *developerConfiguration2 // I have no idea what i'm doing in golang!
		developerConfiguration.ApiKey = "unknown"
	}
	activityLog.Info("here")
	eventLogValue, ok := data.GetGlobalScope().GetAttr("eventLog")
	activityLog.Info(eventLogValue.Value)
	//d2 := eventLogValue.Value
	eventLog, ok := eventLogValue.Value.(models.EventLog)
	activityLog.Info(ok)
	activityLog.Info(eventLog)
	eventLog.ExecTimeEnd = time.Now()
	activityLog.Info(eventLog)

	//	message, _ := context.GetInput(ivMessage).(models.DeveloperConfiguration)
	//	activityLog.Info(message)

	flowInfo, _ := toBool(context.GetInput(ivFlowInfo))

	//	msg := message.ApiKey

	msg := eventLogToString(eventLog, apiConfiguration, developerConfiguration)

	//- - - - [12/Jun/2012:21:53:03 +0000] "GET - HTTP/1.1" 0 200 "-" "-" 0_u2cbu87r6f2q3m66j6yc2uce_ ygnj8v68nqb76akfzetwb799 "-" "-" "-" 0 - 0 0 0 0 -

	if flowInfo {

		//msg = fmt.Sprintf("'%s' - FlowInstanceID  HEY [%s], Flow [%s], Task [%s]", msg,
		//	context.FlowDetails().ID(), context.FlowDetails().Name(), context.TaskName())
	}

	activityLog.Info(msg)

	context.SetOutput(ovMessage, msg)

	return true, nil
}

func eventLogToString(eventLog models.EventLog, apiConfiguration models.ApiConfiguration, developerConfiguration models.DeveloperConfiguration) string {
	activityLog.Info(eventLog)
	byteSize := "0"
	//bytes = strconv.FormatUint(eventLog.Bytes, 10)

	execTime := "0"
	//execTime = strconv.FormatFloat(eventLog.ExecTime, 'f', -1, 64)

	remoteTotalTime := "0"
	//remoteTotalTime = strconv.FormatFloat(eventLog.RemoteTotalTime, 'f', -1, 64)

	connectTime := "0"
	//connectTime = strconv.FormatFloat(eventLog.ConnectTime, 'f', -1, 64)

	preTransferTime := "0"
	//preTransferTime = strconv.FormatFloat(eventLog.PreTransferTime, 'f', -1, 64)

	t := time.Now()
	logTimestamp := t.Format("02/Jan/2006:15:04:05 +0000")

	httpMethodVersion := "-"
	var httpMethodVersionBuf bytes.Buffer
	httpMethodVersionBuf.WriteString(apiConfiguration.Endpoints[0].Method.Verb)
	httpMethodVersionBuf.WriteString(" - HTTP/1.1")
	//httpMethodVersionBuf.WriteString(value)

	httpMethodVersion = httpMethodVersionBuf.String()

	serverName := eventLog.ServerName
	if serverName == "" {
		serverName = "-"
	}

	srcIpd := eventLog.SrcIpd
	if srcIpd == "" {
		srcIpd = "127.0.0.1"
	}

	ident := eventLog.Ident
	if ident == "" {
		ident = "-"
	}

	recordType := eventLog.RecordType
	if recordType == "" {
		recordType = "-"
	}

	proxyErrorCode := eventLog.ProxyErrorCode
	if proxyErrorCode == "" {
		proxyErrorCode = "-"
	}

	apiMethod := eventLog.ApiMethod
	if apiMethod == "" {
		apiMethod = "-"
	}

	proxyWorker := eventLog.ProxyWorker
	if proxyWorker == "" {
		proxyWorker = "-"
	}

	referrerDomain := eventLog.ReferrerDomain
	if referrerDomain == "" {
		referrerDomain = "-"
	}

	referenceGuid := eventLog.ReferenceGuid
	if referenceGuid == "" {
		referenceGuid = "-"
	}

	status := eventLog.Status
	if status == "" {
		status = "200"
	}

	referrer := eventLog.Referrer
	if referrer == "" {
		referrer = "-"
	}

	userAgent := eventLog.UserAgent
	if userAgent == "" {
		userAgent = "-"
	}

	requestId := eventLog.RequestId
	if requestId == "" {
		requestId = "0"
	}
	//- 158.151.240.64 - - [12/Jun/2012:21:53:03 +0000] "GET - HTTP/1.1" 11111 200 "-" "-" 0_u2cbu87r6f2q3m66j6yc2uce_ygnj8v68nqb76akfzetwb799 "-" "-" "GetCompanyDetailRequest" 0 - 5.555555 4.444444 0.333333 0.222222 -
	//- 158.151.240.64 - - [30/Aug/2017:10:55:35 +0000] "GET - HTTP/1.1" 0 200 "-" "-" 0_-_unknown "ty4zvpr9dbnssb496pq3yhhe" "-" "-" - 0 - 0 0 0 0%!(EXTRA string=-)
	//- - - - [30/Aug/2017:10:40:16 +0000] "GET - HTTP/1.1" 0 200 "-" "-" - "unknown" "ty4zvpr9dbnssb496pq3yhhe" "-" - - 0 - 0 0 0 0 -
	//- - - - [12/Jun/2012:21:53:03 +0000] "GET - HTTP/1.1" 0 200 "-" "-" 0_u2cbu87r6f2q3m66j6yc2uce_ ygnj8v68nqb76akfzetwb799 "-" "-" "-" 0 - 0 0 0 0 -

	return fmt.Sprintf("%v %v %v %v [%v] \"%v\" %v %v \"%v\" \"%v\" %v_%v_%v \"%v\" \"%v\" \"%v\" %v %v %v %v %v %v %v",
		serverName, srcIpd, ident, recordType, logTimestamp,
		httpMethodVersion, byteSize, status, referrer, userAgent,
		requestId, developerConfiguration.ApiKey, apiConfiguration.ID, referrerDomain, proxyWorker,
		apiMethod, eventLog.CacheHit, proxyErrorCode, execTime, remoteTotalTime,
		connectTime, preTransferTime, referenceGuid)

}
func toBool(val interface{}) (bool, error) {

	b, ok := val.(bool)
	if !ok {
		s, ok := val.(string)

		if !ok {
			return false, fmt.Errorf("unable to convert to boolean")
		}

		var err error
		b, err = strconv.ParseBool(s)

		if err != nil {
			return false, err
		}
	}

	return b, nil
}
