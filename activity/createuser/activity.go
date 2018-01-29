package createuser

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"fmt"
	"encoding/json"
	"strings"
	"net/http"
	"io/ioutil"
)

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
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Role  string `json:"role,omitempty"`
	Phone  string `json:"phone,omitempty"`
	Alias  string `json:"alias,omitempty"`
	External_id  string `json:"external_id,omitempty"`
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
	companyUrl := context.GetInput(companyUrl).(string)
	OAuthToken := context.GetInput(oAuthToken).(string)
	user.Name = context.GetInput(name).(string)
	user.Email = context.GetInput(email).(string)
	user.Role = context.GetInput(role).(string)
	user.Phone = context.GetInput(phonenumber).(string)
	user.Alias = context.GetInput(alias).(string)
	user.External_id = context.GetInput(externalID).(string)

	b, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return;
	}
	fmt.Println(string(b))

	data := "{\"user\":" + string(b) + "}"
	fmt.Println(data)
	fmt.Println(companyUrl)

	request, _ := http.NewRequest("POST", companyUrl + "/api/v2/users.json",  strings.NewReader(data))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Basic " + OAuthToken)

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
		var dat map[string]interface{}
		err := json.Unmarshal(data, &dat)
		if err != nil {
			fmt.Printf("Error: %s", err)
		}

		//var dat1 map[string]interface{}
		//dat1 = dat["user"]["id"]
		dat_user := dat["user"].(map[string]interface{})
		fmt.Println(dat_user["id"])
		context.SetOutput(userId, dat_user["id"])
	}

	//activityLog.Info(fmt.Sprintf("Num1: %d, Num2: %d", companyUrl, num2))

	context.SetOutput(name, user.Name)
	context.SetOutput(externalID, user.External_id)
	return true, nil
}
