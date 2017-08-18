package error

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var jsonMetadata = `{
  "name": "mashery-error",
  "type": "flogo:activity",
  "ref": "github.com/jpollock/mashling-mashery/activity/error",
  "version": "0.0.1",
  "title": "Log Message",
  "description": "Simple Log Activity",
  "homepage": "https://github.com/jpollock/mashling-mashery/tree/master/activity/error",
  "inputs":[
    {
      "name": "code",
      "type": "string",
      "value": ""
    },
    {
      "name": "status",
      "type": "string",
      "value": ""
    },
    {
      "name": "message",
      "type": "string",
      "value": ""
    }
  ],
  "outputs": [
    {
      "name": "result",
      "type": "any"
    }
  ]
}
`

// init create & register activity
func init() {
	md := activity.NewMetadata(jsonMetadata)
	activity.Register(NewActivity(md))
}
