package eventlogsetup

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var jsonMetadata = `{
  "name": "mashery-event-log-setup",
  "type": "flogo:activity",
  "ref": "github.com/jpollock/mashling-mashery/activity/eventlogsetup",
  "version": "0.0.1",
  "title": "Log Message",
  "description": "Simple Log Activity",
  "homepage": "https://github.com/jpollock/mashling-mashery/tree/master/activity/eventlogsetup",
  "inputs":[
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
