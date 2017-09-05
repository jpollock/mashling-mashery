package eventlogsetup

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/jpollock/mashling-mashery/models"
	"time"
)

// log is the default package logger
var log = logger.GetLogger("activity-eventlog-setup")

// CacheActivity is a Cache Activity implementation
type EventLogActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new CacheActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &EventLogActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *EventLogActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *EventLogActivity) Eval(context activity.Context) (done bool, err error) {
	eventLog := new(models.EventLog)
	eventLog.ServerName = "-"
	eventLog.SrcIpd = "-"
	eventLog.Ident = "-"
	eventLog.RecordType = "-"
	eventLog.HttpMethod = "-"
	eventLog.Status = "-"
	eventLog.Referrer = "-"
	eventLog.UserAgent = "-"
	eventLog.RequestId = "-"
	eventLog.ServiceDevKey = "-"
	eventLog.ServiceKey = "-"
	eventLog.ReferrerDomain = "-"
	eventLog.ProxyWorker = "-"
	eventLog.ApiMethod = "-"
	eventLog.CacheHit = 0
	eventLog.ProxyErrorCode = "-"
	eventLog.ReferenceGuid = "-"
	eventLog.ExecTimeStart = time.Now().UTC()

	dt, ok := data.ToTypeEnum("object")
	if ok {
		data.GetGlobalScope().AddAttr("eventLog", dt, eventLog)
	}

	return true, nil
}
