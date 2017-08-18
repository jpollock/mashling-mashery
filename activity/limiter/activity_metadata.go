package limiter

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var jsonMetadata = `{
  "name": "mashery-limiter",
  "type": "flogo:activity",
  "ref": "github.com/jpollock/mashling-mashery/activity/limiter",
  "version": "0.0.1",
  "title": "Limiter",
  "description": "Simple Limiter",
  "homepage": "https://github.com/jpollock/mashling-mashery/tree/master/activity/limiter",
  "inputs":[
    {
      "name": "count",
      "type": "integer",
      "value": "10"
    },
    {
      "name": "limit",
      "type": "integer",
      "value": "9"
    }
  ],
  "outputs": [
    {
      "name": "limited",
      "type": "boolean"
    }
  ]
}
`

// init create & register activity
func init() {
	md := activity.NewMetadata(jsonMetadata)
	activity.Register(NewActivity(md))
}
