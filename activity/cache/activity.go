package cache

import (
	"sync"

	"bytes"
	"encoding/json"
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/go-redis/redis"
	"net/url"
	"strings"
	"time"
)

type ApiConfiguration struct {
	ID              string
	Name            string
	QpsLimitOverall int
	Endpoints       []Endpoint
}

type Endpoint struct {
	ID    string
	Name  string
	Cache Cache
}

type Cache struct {
	CacheTtlOverride int
}

// log is the default package logger
var log = logger.GetLogger("activity-mashery-cache")

const (
	ivRedisAddress = "redisAddress"
	ivURI          = "uri"
	ivPathParams   = "pathParams"
	ivQueryParams  = "queryParams"
	ivParams       = "params"

	ivContent          = "content"
	ivApiConfiguration = "apiConfiguration"

	ovValue        = "value"
	ovFoundContent = "foundContent"
)

// CacheActivity is a Cache Activity implementation
type CacheActivity struct {
	sync.Mutex
	metadata *activity.Metadata
	counters map[string]int
}

// NewActivity creates a new CacheActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &CacheActivity{metadata: metadata, counters: make(map[string]int)}
}

// Metadata implements activity.Activity.Metadata
func (a *CacheActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *CacheActivity) Eval(context activity.Context) (done bool, err error) {

	apiConfiguration := context.GetInput(ivApiConfiguration).(string)
	var c ApiConfiguration
	json.Unmarshal([]byte(apiConfiguration), &c)
	log.Info(c.Name)

	redisAddress := context.GetInput(ivRedisAddress).(string)
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	log.Info("CACHE::1")

	uri := context.GetInput(ivURI).(string)

	containsParam := strings.Index(uri, "/:") > -1

	if containsParam {

		val := context.GetInput(ivPathParams)

		if val == nil {
			val = context.GetInput(ivParams)

			if val == nil {
				err := activity.NewError("Path Params not specified, required for URI: "+uri, "", nil)
				return false, err
			}
		}

		pathParams := val.(map[string]string)
		uri = BuildURI(uri, pathParams)
	}

	if context.GetInput(ivQueryParams) != nil {
		queryParams := context.GetInput(ivQueryParams).(map[string]string)

		qp := url.Values{}

		for key, value := range queryParams {
			qp.Set(key, value)
		}

		uri = uri + "?" + qp.Encode()
	}

	fmt.Printf("URI= %s\n", uri)
	content := context.GetInput(ivContent).(string)
	log.Info("CACHE::2")
	log.Info(len(content))
	if content != "" {
		log.Info("CACHE::ee")

		//fmt.Print(*time.Second) /
		err := client.Set(uri, content, time.Duration(c.Endpoints[0].Cache.CacheTtlOverride)*time.Second).Err()
		if err != nil {
			log.Info("CACHE::weeoe")
			panic(err)
		}
		context.SetOutput(ovFoundContent, false)

	} else {
		log.Info("CACHE::4")
		val2, err := client.Get(uri).Result()
		if err == redis.Nil {
			fmt.Println("key2 does not exists")
			context.SetOutput(ovFoundContent, false)
		} else if err != nil {
			panic(err)
		} else {
			fmt.Println("key2", val2)
			content = val2
			var result interface{}

			d := json.NewDecoder(bytes.NewReader([]byte(content)))
			d.UseNumber()
			err = d.Decode(&result)

			//json.Unmarshal(respBody, &result)

			log.Debug("response Body:", result)

			context.SetOutput(ovValue, result)

			context.SetOutput(ovFoundContent, true)

		}
	}
	log.Info("CACHE::6")
	//context.SetOutput(ovValue, content)

	return true, nil
}

// BuildURI is a temporary crude URI builder
func BuildURI(uri string, values map[string]string) string {

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

	buffer.WriteString(uri[0:i])

	for i < len(uri) {
		if uri[i] == ':' {
			j := i + 1
			for j < len(uri) && uri[j] != '/' {
				j++
			}

			if i+1 == j {

				buffer.WriteByte(uri[i])
				i++
			} else {

				param := uri[i+1 : j]
				value := values[param]
				buffer.WriteString(value)
				if j < len(uri) {
					buffer.WriteString("/")
				}
				i = j + 1
			}

		} else {
			buffer.WriteByte(uri[i])
			i++
		}
	}

	return buffer.String()
}
