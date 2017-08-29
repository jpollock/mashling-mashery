package authenticate

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var jsonMetadata = `{
  "name": "mashery-authenticate",
  "type": "flogo:activity",
  "ref": "github.com/jpollock/mashling-mashery/activity/authenticate",
  "version": "0.0.1",
  "title": "Increment Counter",
  "description": "Simple Global Counter Activity",
  "homepage": "https://github.com/jpollock/mashling-mashery/tree/master/activity/authenticate",
  "inputs":[
    {
      "name": "activityEnabled",
      "type": "boolean",
      "value": false
    }
  ],
  "outputs": [
    {
      "name": "authenticated",
      "type": "boolean"
    }    
  ]
}`

// init create & register activity
func init() {
	md := activity.NewMetadata(jsonMetadata)
	activity.Register(NewActivity(md))
}
