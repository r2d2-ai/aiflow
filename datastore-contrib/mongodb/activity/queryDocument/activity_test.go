package queryDocument

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/r2d2-ai/core/activity"
	"github.com/r2d2-ai/core/data/mapper"
	"github.com/r2d2-ai/core/data/resolve"
	"github.com/r2d2-ai/core/support"
	"github.com/r2d2-ai/core/support/log"
	"github.com/r2d2-ai/core/support/test"
	_ "github.com/r2d2-ai/datastore-contrib/mongodb/connection"
	"github.com/stretchr/testify/assert"
)

var activityMetadata *activity.Metadata

var settingsjson = `{
	"settings": {
		"connection": {
			"id": "e1e890d0-de91-11e9-aef0-13201957902e",
			"name": "mongocon",
			"ref": "github.com/r2d2-ai/datastore-contrib/mongodb/connection",
			"settings": {
				"name": "mongocon",
				"description": "",
				"connectionURI": "mongodb://admin:admin@10.102.169.188:27017",
				"credType": "None",
				"ssl": false
			}
		},
		"operation": "Find One Document",
        "databaseName": "sample",
        "collectionName": "test",
        "timeout": 0
	}
}`
var settingsjson1 = `{
	"settings": {
		"connection": {
			"id": "e1e890d0-de91-11e9-aef0-13201957902e",
			"name": "mongocon",
			"ref": "github.com/r2d2-ai/datastore-contrib/mongodb/connection",
			"settings": {
				"name": "mongocon",
				"description": "",
				"connectionURI": "mongodb://admin:admin@10.102.169.188:27017",
				"credType": "None",
				"ssl": false
			}
		},
		"operation": "Find Many Documents",
        "databaseName": "sample",
        "collectionName": "test",
        "timeout": 0
	}
}`

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.ToMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}
func Test_FindOne(t *testing.T) {
	log.RootLogger().Info("****TEST : Executing Find One start****")
	m := make(map[string]interface{})
	err1 := json.Unmarshal([]byte(settingsjson), &m)
	assert.Nil(t, err1)
	mf := mapper.NewFactory(resolve.GetBasicResolver())

	support.RegisterAlias("connection", "connection", "github.com/r2d2-ai/datastore-contrib/mongodb/connection")
	fmt.Println("=======Settings========", m["settings"])
	iCtx := test.NewActivityInitContext(m["settings"], mf)
	act, err := New(iCtx)
	assert.Nil(t, err)
	tc := test.NewActivityContext(act.Metadata())
	tc.SetInput("criteria", map[string]string{"location": "abcd"})
	_, err = act.Eval(tc)
	// Getting outputs
	testOutput := tc.GetOutput("response")
	jsonOutput, _ := json.Marshal(testOutput)
	log.RootLogger().Infof("jsonOutput is : %s", string(jsonOutput))
	log.RootLogger().Info("****TEST : Executing Find One ends****")
	assert.Nil(t, err)
}
func Test_FindAll(t *testing.T) {
	log.RootLogger().Info("****TEST : Executing Find All start****")
	m := make(map[string]interface{})
	err1 := json.Unmarshal([]byte(settingsjson1), &m)
	assert.Nil(t, err1)
	mf := mapper.NewFactory(resolve.GetBasicResolver())
	support.RegisterAlias("connection", "connection", "github.com/r2d2-ai/datastore-contrib/mongodb/connection")
	iCtx := test.NewActivityInitContext(m["settings"], mf)
	act, err := New(iCtx)
	assert.Nil(t, err)
	tc := test.NewActivityContext(act.Metadata())
	tc.SetInput("criteria", map[string]string{"location": "abcde"})

	_, err = act.Eval(tc)
	// Getting outputs
	testOutput := tc.GetOutput("response")
	jsonOutput, _ := json.Marshal(testOutput)
	log.RootLogger().Infof("jsonOutput is : %s", string(jsonOutput))
	log.RootLogger().Info("****TEST : Executing Find All ends****")
	assert.Nil(t, err)
}
func Test_FindMany(t *testing.T) {
	log.RootLogger().Info("****TEST : Executing Find Many start****")
	m := make(map[string]interface{})
	err1 := json.Unmarshal([]byte(settingsjson1), &m)
	assert.Nil(t, err1)
	mf := mapper.NewFactory(resolve.GetBasicResolver())
	support.RegisterAlias("connection", "connection", "github.com/r2d2-ai/datastore-contrib/mongodb/connection")
	iCtx := test.NewActivityInitContext(m["settings"], mf)
	act, err := New(iCtx)
	assert.Nil(t, err)

	tc := test.NewActivityContext(act.Metadata())
	tc.SetInput("criteria", `[{"location" : "Hyderabad" },{"location" : "Chennai" }]`)
	_, err = act.Eval(tc)
	// Getting outputs
	testOutput := tc.GetOutput("response")
	jsonOutput, _ := json.Marshal(testOutput)
	log.RootLogger().Infof("jsonOutput is : %s", string(jsonOutput))
	log.RootLogger().Info("****TEST : Executing Find Many ends****")
	assert.Nil(t, err)
}
