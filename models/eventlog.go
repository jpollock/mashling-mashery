package models

import (
	"time"
)

type EventLog struct {
	ServerName           string
	SrcIpd               string
	Ident                string
	RecordType           string
	LogTimestamp         string
	HttpMethodVersion    string
	Bytes                uint64
	Status               string
	Referrer             string
	UserAgent            string
	RequestId            string
	ServiceDevKey        string
	ServiceKey           string
	ReferrerDomain       string
	ProxyWorker          string
	ApiMethod            string
	CacheHit             int
	ProxyErrorCode       string
	ExecTimeStart        time.Time
	ExecTimeEnd          time.Time
	ExecTime             float64
	RemoteTotalTimeStart time.Time
	RemoteTotalTimeEnd   time.Time
	RemoteTotalTime      float64
	ConnectTimeStart     time.Time
	ConnectTimeEnd       time.Time
	ConnectTime          float64
	PreTransferTimeStart time.Time
	PreTransferTimeEnd   time.Time
	PreTransferTime      float64
	ReferenceGuid        string
}

func (r EventLog) SetExecTimeStart() {
	r.ExecTimeStart = time.Now()
	return
}

func (r EventLog) SetExecTimeEnd() {
	r.ExecTimeEnd = time.Now()
	return
}
