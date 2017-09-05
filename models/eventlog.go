package models

import (
	"encoding/json"
	"fmt"
	"time"
)

type EventLog struct {
	ServerName                    string    `json:"server_name"`
	SrcIpd                        string    `json:"src_ip"`
	Ident                         string    `json:"ident"`
	RecordType                    string    `json:"record_type"`
	LogTimestamp                  time.Time `json:"log_timestamp"`
	HttpMethod                    string    `json:"http_method"`
	Uri                           string    `json:"uri"`
	HttpVersion                   string    `json:"http_version"`
	Bytes                         int64     `json:"bytes"`
	Status                        string    `json:"status"`
	Referrer                      string    `json:"referrer"`
	UserAgent                     string    `json:"useragent"`
	RequestId                     string    `json:"request_id"`
	RequestTime                   float64   `json:"request_time"`
	ServiceDevKey                 string    `json:"service_dev_key"`
	ServiceKey                    string    `json:"service_key"`
	ReferrerDomain                string    `json:"referrer_domain"`
	ProxyWorker                   string    `json:"proxy_worker"`
	ApiMethod                     string    `json:"api_method"`
	CacheHit                      int       `json:"cache_hit"`
	ProxyErrorCode                string    `json:"proxy_error_code"`
	ExecTimeStart                 time.Time `json:"exec_time_start"`
	ExecTimeEnd                   time.Time `json:"exec_time_end"`
	ExecTime                      float64   `json:"exec_time"`
	RemoteTotalTimeStart          time.Time `json:"remote_total_time_start"`
	RemoteTotalTimeEnd            time.Time `json:"remote_total_time_end"`
	RemoteTotalTime               float64   `json:"remote_total_time"`
	ConnectTimeStart              time.Time `json:"connect_time_start"`
	ConnectTimeEnd                time.Time `json:"connect_time_end"`
	ConnectTime                   float64   `json:"connect_time"`
	PreTransferTimeStart          time.Time `json:"pre_transfer_time_start"`
	PreTransferTimeEnd            time.Time `json:"pre_transfer_time_end"`
	PreTransferTime               float64   `json:"pre_transfer_time"`
	ReferenceGuid                 string    `json:"reference_guid"`
	OauthAccessToken              string    `json:"oauth_access_token"`
	CapacityCenterId              int       `json:"capacity_center_id"`
	CapacityCenterName            string    `json:"capacity_center_name"`
	SslEnabled                    bool      `json:"ssl_enabled"`
	QuotaId                       string    `json:"quota_id"`
	QuotaObjectType               string    `json:"quota_object_type"`
	QuotaValue                    int       `json:"quota_value"`
	ThrottleId                    int       `json:"throttle_id"`
	ThrottleObjectType            string    `json:"throttle_object_type"`
	ThrottleValue                 int       `json:"throttle_value"`
	ServiceAggregateThrottleValue int       `json:"service_aggregate_throttle_value"`
	AreaAggregateThrottleValue    int       `json:"area_aggregate_throttle_value"`
	PlanId                        int       `json:"plan_id"`
	ServiceDefinitionEndpointId   int       `json:"service_definition_endpoint_id"`
	ServiceDefinitionMethodId     int       `json:"service_definition_method_id"`
	ResourceFilterId              int       `json:"resource_filter_id"`
	StaleCache                    bool      `json:"stale_cache"`
	DataSource                    int       `json:"data_source"`
	BytesIn                       int       `json:"bytes_in"`
	ClientTransferTime            float64   `json:"client_transfer_time"`
}

func (d *EventLog) MarshalJSON() ([]byte, error) {
	type Alias EventLog
	fmt.Println("test")
	return json.Marshal(&struct {
		*Alias
		LogTimestamp string `json:"log_timestamp"`
	}{
		Alias:        (*Alias)(d),
		LogTimestamp: d.LogTimestamp.Format("2006-01-02T15:04:05.12340Z"),
	})
}

func (r EventLog) SetExecTimeStart() {
	r.ExecTimeStart = time.Now()
	return
}

func (r EventLog) SetExecTimeEnd() {
	r.ExecTimeEnd = time.Now()
	return
}
