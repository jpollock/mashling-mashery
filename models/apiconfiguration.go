package models

type ApiConfiguration struct {
	ID              string     `json:"id"`
	Name            string     `json:"name"`
	QpsLimitOverall int        `json:"qpsLimitOverall"`
	Endpoints       []Endpoint `json:"endpoints"`
}

type Endpoint struct {
	ID                     string `json:"id"`
	Name                   string `json:"name"`
	Cache                  Cache  `json:"cache"`
	ApiKeyValueLocationKey string `json:"apiKeyValueLocationKey"`
	Method                 Method
}

type Cache struct {
	CacheTtlOverride int `json:"cacheTtlOverride"`
}

type Method struct {
	Name    string
	Verb    string
	Version string
}
