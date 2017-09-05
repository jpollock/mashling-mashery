package reply

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var jsonMetadata = `{
  "name": "mashery-reply",
  "type": "flogo:activity",
  "ref": "github.com/jpollock/mashling-mashery/activity/reply",
  "version": "0.0.1",
  "title": "Reply To Trigger",
  "description": "Simple Reply Activity",
  "homepage": "https://github.com/jpollock/mashling-mashery/tree/master/activity/reply",
  "inputs":[
    {
      "name": "data",
      "type": "any"
    }
  ],
  "outputs": [
    {
      "name": "content",
      "type": "string"
    }
  ]
}
`

// init create & register activity
func init() {
	md := activity.NewMetadata(jsonMetadata)
	activity.Register(NewActivity(md))
}
