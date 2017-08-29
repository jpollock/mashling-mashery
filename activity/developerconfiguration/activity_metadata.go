package developerconfiguration

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var jsonMetadata = `{
  "name": "mashery-developer-configuration",
  "type": "flogo:activity",
  "ref": "github.com/jpollock/mashling-mashery/activity/developerconfiguration",
  "version": "0.0.1",
  "title": "Log Message",
  "description": "Simple Log Activity",
  "homepage": "https://github.com/jpollock/mashling-mashery/tree/master/activity/developerconfiguration",
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
      "name": "pathParams",
      "type": "params"
    },
    {
      "name": "queryParams",
      "type": "params"
    },
    {
      "name": "content",
      "type": "any"
    }
  ],
  "outputs": [
  ]
}
`

// init create & register activity
func init() {
	md := activity.NewMetadata(jsonMetadata)
	activity.Register(NewActivity(md))
}
