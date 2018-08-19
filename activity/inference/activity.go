package inference

import (

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)


// log is the default package logger
var log = logger.GetLogger("activity-tibco-inference")


// InferenceActivity is an Activity that is used to invoke a a ML Model using flogo-ml framework
type InferenceActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new InferenceActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &InferenceActivity{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *InferenceActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - Runs an ML model
func (a *InferenceActivity) Eval(context activity.Context) (done bool, err error) {


	return true, nil
}
