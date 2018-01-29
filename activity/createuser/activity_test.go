package createuser

import (
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput(companyUrl, "https://kundantest.zendesk.com")
	tc.SetInput(oAuthToken, "aXQua3VuZGFubWNhQGdtYWlsLmNvbS90b2tlbjpNZG1iR1ZxTUQweXlqdkVZd3VsRzJMSWlYVGdTWVlrdW9wcnFpSmNW")
	tc.SetInput(name, "Test")
	tc.SetInput(email, "abc111@gmail.com")

	act.Eval(tc)

	name := tc.GetOutput(name).(string)

	if name != "Test" {
		t.Error("Name did not match")
	}
}

