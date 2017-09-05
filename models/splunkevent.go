package models

type SplunkEvent struct {
	SourceType string    `json:"sourcetype"`
	Event      *EventLog `json:"event"`
}
