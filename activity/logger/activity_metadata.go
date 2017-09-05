package test

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var jsonMetadata = `{
  "name": "mashery-log",
  "type": "flogo:activity",
  "ref": "github.com/jpollock/mashling-mashery/activity/test",
  "version": "0.0.1",
  "title": "Log Message",
  "description": "Simple Log Activity",
  "homepage": "https://github.com/jpollock/mashling-mashery/tree/master/activity/test",
  "inputs":[
    {
      "name": "fluentdHost",
      "type": "string",
      "required": true
    },
    {
      "name": "fluentdPort",
      "type": "string",
      "required": true
    },
    {
      "name": "fluentdTag",
      "type": "string",
      "required": true
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
