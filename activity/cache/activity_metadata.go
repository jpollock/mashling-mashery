package cache

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var jsonMetadata = `{
  "name": "mashery-cache",
  "type": "flogo:activity",
  "ref": "github.com/jpollock/mashling-mashery/activity/cache",
  "version": "0.0.1",
  "title": "Increment Counter",
  "description": "Simple Global Counter Activity",
  "homepage": "https://github.com/jpollock/mashling-mashery/tree/master/activity/cache",
  "inputs":[
    {
      "name": "activityEnabled",
      "type": "boolean",
      "value": false
    },
    {
      "name": "redisAddress",
      "type": "string",
      "required": true
    },
    {
      "name": "uri",
      "type": "string",
      "required": true
    },
    {
      "name": "pathParams",
      "type": "params"
    },
    {
      "name": "queryParams",
      "type": "params"
    },
    {
      "name": "content",
      "type": "string",
      "required": false
    },
    {
      "name": "apiConfiguration",
      "type": "string"
    }
  ],
  "outputs": [
    {
      "name": "value",
      "type": "string"
    },
    {
      "name": "foundContent",
      "type": "boolean"
    }    
  ]
}`

// init create & register activity
func init() {
	md := activity.NewMetadata(jsonMetadata)
	activity.Register(NewActivity(md))
}
