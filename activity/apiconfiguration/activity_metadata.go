package apiconfiguration

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var jsonMetadata = `{
  "name": "mashery-log",
  "type": "flogo:activity",
  "ref": "github.com/jpollock/mashling-mashery/activity/apiconfiguration",
  "version": "0.0.1",
  "title": "Log Message",
  "description": "Simple Log Activity",
  "homepage": "https://github.com/jpollock/mashling-mashery/tree/master/activity/apiconfiguration",
  "inputs":[
    {
      "name": "filePath",
      "type": "string"
    }
  ],
  "outputs": [
    {
      "name": "apiConfiguration",
      "type": "object"
    }
  ]
}
`

// init create & register activity
func init() {
	md := activity.NewMetadata(jsonMetadata)
	activity.Register(NewActivity(md))
}
