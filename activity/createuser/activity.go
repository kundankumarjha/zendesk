package createuser

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// activityAdd is the default logger for the createuser Activity
var activityLog = logger.GetLogger("activity-kundankumarjha-createuser")

const (
	companyUrl = "companyUrl"
	oAuthToken = "oAuthToken"
	name  = "name"
	email = "email"
	role = "role"
	phonenumber = "phoneNumber"
	alias = "alias"
	externalID = "externalID"
	userId = "userId"
)

type User struct {
	companyUrl string
	OAuthToken string
	Name        string
	Email string
	Role  string
	Phone  string
	Alias  string
	External_id  string
}

type AddActivity struct {
	metadata *activity.Metadata
}

func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &AddActivity{metadata: metadata}
}

func (a *AddActivity) Metadata() *activity.Metadata {
	return a.metadata
}

func (a *AddActivity) Eval(context activity.Context) (done bool, err error) {

	user := &User{}
	user.companyUrl = context.GetInput(companyUrl).(string)
	user.OAuthToken = context.GetInput(oAuthToken).(string)
	user.Name = context.GetInput(name).(string)
	user.Email = context.GetInput(email).(string)
	user.Role = context.GetInput(role).(string)
	user.Phone = context.GetInput(phonenumber).(string)
	user.Alias = context.GetInput(alias).(string)
	user.External_id = context.GetInput(externalID).(string)

	//activityLog.Info(fmt.Sprintf("Num1: %d, Num2: %d", companyUrl, num2))

	context.SetOutput(userId, "")
	context.SetOutput(name, user.Name)
	context.SetOutput(externalID, user.External_id)
	return true, nil
}
