package logger

import (
	"bytes"

	//	"encoding/json"
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/jpollock/mashling-mashery/models"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// activityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-tibco-log-2")

const (
	ivHost = "fluentdHost"
	ivPort = "fluentdPort"
	ivTag  = "fluentdTag"

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

	/*fluentdHost := context.GetInput(ivHost).(string)
	fluentdPort := context.GetInput(ivPort).(string)
	fluentdTag := context.GetInput(ivTag).(string)*/

	//mv := context.GetInput(ivMessage)
	apiConfigValue, ok := data.GetGlobalScope().GetAttr("apiConfiguration")
	d := apiConfigValue.Value
	apiConfiguration, ok := d.(models.ApiConfiguration)

	var developerConfiguration models.DeveloperConfiguration
	developerConfigValue, ok := data.GetGlobalScope().GetAttr("developerConfiguration")
	if ok {
		d3 := developerConfigValue.Value
		developerConfiguration, ok = d3.(models.DeveloperConfiguration)

	} else {
		developerConfiguration2 := new(models.DeveloperConfiguration)
		developerConfiguration = *developerConfiguration2 // I have no idea what i'm doing in golang!
		developerConfiguration.ApiKey = "unknown"
	}
	eventLogValue, ok := data.GetGlobalScope().GetAttr("eventLog")
	d = eventLogValue.Value
	eventLog, ok := d.(*models.EventLog)

	eventLog.ExecTimeEnd = time.Now().UTC()
	eventLog.LogTimestamp = time.Now().UTC()

	msg := eventLogToString(eventLog, apiConfiguration, developerConfiguration)

	//- - - - [12/Jun/2012:21:53:03 +0000] "GET - HTTP/1.1" 0 200 "-" "-" 0_u2cbu87r6f2q3m66j6yc2uce_ ygnj8v68nqb76akfzetwb799 "-" "-" "-" 0 - 0 0 0 0 -

	context.SetOutput(ovMessage, msg)

	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}

	var jsonStr = []byte("{\"message\":\"" + msg + "\"}")
	netClient.Post("http://localhost:8888/local.ddd", "application/json", bytes.NewBuffer(jsonStr))
	//response, _ := netClient.PostForm("http://localhost:8888/local.ddd", url.Values{"message": []string("Test")})

	/*
	   	    try:
	           message_json = json.loads(message)
	           event = {}
	           event['sourcetype'] = "mysourcetype"
	           event['event'] = message_json
	           headers = {"Authorization": "Splunk 0EB426FE-51D5-4C6E-9AA4-5C15222A522F"}
	           #print json.dumps(event)
	           response = requests.post('http://127.0.0.1:8088/services/collector', headers=headers, data=json.dumps(event))
	           #print response.status
	       except Exception:
	           return
	*/
	/*splunkEvent := new(models.SplunkEvent)
	splunkEvent.SourceType = "mysourcetype"
	activityLog.Info(splunkEvent)
	splunkEvent.Event = eventLog

	data, err := json.Marshal(splunkEvent)
	if err != nil {
		activityLog.Error(err)
	}

	activityLog.Info(string(data))

	//netClient.Post("http://127.0.0.1:8088/services/collector", "application/json", bytes.NewBuffer(data))

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://127.0.0.1:8088/services/collector", bytes.NewReader(data))
	req.Header.Add("Authorization", "Splunk 0EB426FE-51D5-4C6E-9AA4-5C15222A522F")
	client.Do(req)*/
	//defer resp.Body.Close()
	return true, nil
}

func eventLogToString(eventLog *models.EventLog, apiConfiguration models.ApiConfiguration, developerConfiguration models.DeveloperConfiguration) string {
	activityLog.Info(eventLog)
	//byteSize := "0"
	byteSize := strconv.FormatInt(eventLog.Bytes, 10)

	execTime := "0"
	activityLog.Info("timestart")
	activityLog.Info(eventLog.ExecTimeStart)
	activityLog.Info("timeend")
	activityLog.Info(eventLog.ExecTimeEnd)

	diff := eventLog.ExecTimeEnd.Sub(eventLog.ExecTimeStart)
	execTime = strconv.FormatFloat(diff.Seconds(), 'f', -1, 64)
	//execTime = strconv.FormatFloat(eventLog.ExecTime, 'f', -1, 64)

	remoteTotalTime := "0"
	activityLog.Info("timestart")
	activityLog.Info(eventLog.RemoteTotalTimeStart)
	activityLog.Info("timeend")
	activityLog.Info(eventLog.RemoteTotalTimeEnd)
	diff = eventLog.RemoteTotalTimeEnd.Sub(eventLog.RemoteTotalTimeStart)
	remoteTotalTime = strconv.FormatFloat(diff.Seconds(), 'f', -1, 64)

	//remoteTotalTime = strconv.FormatFloat(eventLog.RemoteTotalTime, 'f', -1, 64)

	connectTime := "0"
	//connectTime = strconv.FormatFloat(eventLog.ConnectTime, 'f', -1, 64)

	preTransferTime := "0"
	//preTransferTime = strconv.FormatFloat(eventLog.PreTransferTime, 'f', -1, 64)

	t := time.Now().UTC()
	logTimestamp := t.Format("02/Jan/2006:15:04:05 +0000")

	httpMethodVersion := "-"
	var httpMethodVersionBuf bytes.Buffer
	httpMethodVersionBuf.WriteString(apiConfiguration.Endpoints[0].Method.Verb)
	httpMethodVersionBuf.WriteString(" ")
	httpMethodVersionBuf.WriteString(GetUri(eventLog.Uri))
	httpMethodVersionBuf.WriteString(" HTTP/1.1")
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

	return fmt.Sprintf("%v %v %v %v [%v] \\\"%v\\\" %v %v \\\"%v\\\" \\\"%v\\\" %v_%v_%v \\\"%v\\\" \\\"%v\\\" \\\"%v\\\" %v %v %v %v %v %v %v",
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

// BuildURI is a temporary crude URI builder
func GetUri(uri string) string {

	var buffer bytes.Buffer
	buffer.Grow(len(uri))

	addrStart := strings.Index(uri, "://")

	i := addrStart + 3

	for i < len(uri) {
		if uri[i] == '/' {
			break
		}
		i++
	}

	buffer.WriteString(uri[i:len(uri)])

	return buffer.String()
}
