package intersight

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	endsOfWorkflow = map[string]bool{"RUNNING": false, "WAITING": false, "COMPLETED": true, "TIME_OUT": true, "FAILED": true}
)

// SuppressDiffAdditionProps Suppress Difference functions for additional properties
//old is from tfstate file
// new is from tf config file
func SuppressDiffAdditionProps(k, old, new string, d *schema.ResourceData) bool {
	if old == "null" && new == "" {
		return true
	}
	if new == "" {
		new = "{}"
	}
	var oldJson = make(map[string]interface{})
	var newJson = make(map[string]interface{})
	err := json.Unmarshal([]byte(old), &oldJson)
	err1 := json.Unmarshal([]byte(new), &newJson)
	if err != nil || err1 != nil {
		log.Printf("Error while json parsing ERR: %+v ERR1: %+v", err, err1)
		return false
	}
	different := true
	for k, v := range newJson {
		different = different && recursiveValueCheck(oldJson, k, v)
	}
	return different
}

func parseFilterMap(m map[string]interface{}) ([]string, []string) {
	var keyArray []string
	var valArray []string
	for k, v := range m {
		if k == "ClassId" || k == "ObjectType" {
			continue
		}
		switch reflect.TypeOf(v).Kind() {
		case reflect.Map:
			temp_k, temp_v := parseFilterMap(v.(map[string]interface{}))
			for i := range temp_k {
				keyArray = append(keyArray, k+"."+temp_k[i])
				valArray = append(valArray, temp_v[i])
			}
		case reflect.String:
			keyArray = append(keyArray, k)
			valArray = append(valArray, "'"+v.(string)+"'")
		case reflect.Bool:
			keyArray = append(keyArray, k)
			valArray = append(valArray, strconv.FormatBool(v.(bool)))
		case reflect.Int:
			keyArray = append(keyArray, k)
			valArray = append(valArray, strconv.FormatInt(v.(int64), 10))
		case reflect.Float64:
			keyArray = append(keyArray, k)
			valArray = append(valArray, fmt.Sprintf("%f", v.(float64)))
		}
	}
	return keyArray, valArray
}

func getRequestParams(in []byte) string {
	var o string
	var s map[string]interface{}
	err := json.Unmarshal(in, &s)
	if err != nil {
		return ""
	}
	keyArr, valArr := parseFilterMap(s)
	for i := range keyArr {
		o += fmt.Sprintf("(%s eq %s) and ", keyArr[i], valArr[i])
	}
	o = strings.TrimSuffix(o, " and ")
	return o
}

func recursiveValueCheck(oldM map[string]interface{}, k string, v interface{}) bool {
	if k == "Password" {
		return true
	}
	x := reflect.TypeOf(v).String()
	b := true
	simpleDatatypes := "string int int32 int64 float bool float64"
	complexDatatypes := "map[string]interface {}"
	if strings.Contains(simpleDatatypes, x) {
		if oldM[k] != v {
			b = false
		}
	} else if strings.Contains(complexDatatypes, x) {
		if oldM[k] == nil || v == nil {
			b = false
		} else {
			oldMap := oldM[k].(map[string]interface{})
			newMap := v.(map[string]interface{})
			for k1, v1 := range newMap {
				b = b && recursiveValueCheck(oldMap, k1, v1)
			}
		}
	} else {
		log.Printf("Type did not match: %s", x)
	}
	return b
}

func checkWorkflowStatus(conn *Config, w models.WorkflowWorkflowInfoRelationship) (string, error) {
	var workflowInfo models.WorkflowWorkflowInfo
	var responseErr error
	for { // infinite loop
		moid := w.MoMoRef.GetMoid()
		workflowInfo, _, responseErr = conn.ApiClient.WorkflowApi.GetWorkflowWorkflowInfoByMoid(conn.ctx, moid).Execute()
		if responseErr != nil {
			return "", responseErr
		}
		if endsOfWorkflow[workflowInfo.GetStatus()] {
			break
		} else {
			time.Sleep(5 * time.Second)
			log.Printf("Workflow %s is %s ... %f percent \n", workflowInfo.GetName(), workflowInfo.GetStatus(), workflowInfo.GetProgress())
		}
	}
	warning := fmt.Sprintf("Workflow moid: %s name: %s is in status %s", workflowInfo.GetMoid(), workflowInfo.GetName(), workflowInfo.GetStatus())
	return warning, nil
}

// CustomizeTagDiff Customize tag diff to ignore cisco.meta tags
func CustomizeTagDiff(ctx context.Context, diff *schema.ResourceDiff, i interface{}) error {
	if tags, ok := diff.GetOk("tags"); ok {
		tagsArray := tags.([]interface{})
		log.Println("tags", tagsArray) // user config
		old, _new := diff.GetChange("tags")
		oldArray := old.([]interface{})
		newArray := _new.([]interface{})
		var newTags, subNew []map[string]interface{}
		var existingKeys = make(map[string]bool)
		for _, item := range newArray {
			x := item.(map[string]interface{})
			existingKeys[x["key"].(string)] = true
		}
		for _, item := range oldArray {
			x := item.(map[string]interface{})
			var s = x["key"].(string)
			_, ok := existingKeys[s]
			if strings.HasPrefix(s, "cisco.meta") && !ok {
				newTags = append(newTags, x)
			}
		}
		for _, item := range tagsArray {
			x := item.(map[string]interface{})
			subNew = append(subNew, x)
		}
		newTags = append(newTags, subNew...)
		err := diff.SetNew("tags", newTags)
		return err
	}
	return nil
}
