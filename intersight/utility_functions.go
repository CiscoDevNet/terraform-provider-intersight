package intersight

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"sort"
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
// old is from tfstate file
// new is from tf config file
func SuppressDiffAdditionProps(k, old, new string, d *schema.ResourceData) bool {
	if old == "" && new == "" {
		return true
	}
	if new == "" {
		new = "{}"
	}
	var oldJson = make(map[string]interface{})
	var newJson = make(map[string]interface{})
	if old != "" {
		err := json.Unmarshal([]byte(old), &oldJson)
		if err != nil {
			log.Printf("Error while json parsing ERR: %+v", err)
			return false
		}
	}
	err1 := json.Unmarshal([]byte(new), &newJson)
	if err1 != nil {
		log.Printf("Error while json parsing ERR1: %+v", err1)
		return false
	}
	different := true
	for k, v := range newJson {
		different = recursiveValueCheck(oldJson, k, v)
		if different == false {
			return different
		}
	}
	return different
}

func parseFilterMap(m map[string]interface{}) ([]string, []string) {
	var keyArray []string
	var valArray []string
	for k, v := range m {
		if (k == "ClassId") || (k == "ObjectType" && v == "") {
			continue
		}
		switch reflect.TypeOf(v).Kind() {
		case reflect.Map:
			temp_k, temp_v := parseFilterMap(v.(map[string]interface{}))
			for i := range temp_k {
				keyArray = append(keyArray, k+"."+temp_k[i])
				valArray = append(valArray, temp_v[i])
			}
		case reflect.Slice:
			valueArray := make([]string, 0)
			for _, val := range v.([]interface{}) {
				switch value := val.(type) {
				case int64:
					valueArray = append(valueArray, strconv.FormatInt(value, 10))
				case string:
					valueArray = append(valueArray, "'"+value+"'")
				case float64:
					valueArray = append(valueArray, fmt.Sprintf("%f", value))
				case bool:
					valueArray = append(valueArray, strconv.FormatBool(value))
				}
			}
			keyArray = append(keyArray, k)
			valArray = append(valArray, "("+strings.Join(valueArray, ",")+")")
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
func sortArray(arr []interface{}) []interface{} {
	sort.SliceStable(arr, func(i, j int) bool {
		return fmt.Sprintf("%v", arr[i]) < fmt.Sprintf("%v", arr[j])
	})
	return arr
}

func recursiveValueCheck(oldM map[string]interface{}, k string, v interface{}) bool {
	log.Printf("Checking key: %s, Value: %+v", k, v)
	if k == "Password" || k == "Passphrase" {
		return true
	}
	y := reflect.TypeOf(v)
	if y == nil {
		if oldM[k] != nil {
			return false
		}
		return true
	}
	x := y.String()
	b := true
	simpleDatatypes := "string int int32 int64 float bool float64"
	complexDatatypes := "map[string]interface {}"
	if strings.Contains(simpleDatatypes, x) {
		if oldM[k] != v {
			return false
		}
	} else if strings.Contains(complexDatatypes, x) {
		if (oldM[k] != nil && v == nil) || (oldM[k] == nil && v != nil) {
			return false
		} else {
			oldMap := oldM[k].(map[string]interface{})
			newMap := v.(map[string]interface{})
			return compareMapRecursively(oldMap, newMap)
		}
	} else if x == "[]string" {
		oldStringArray, oldOk := oldM[k].([]string)
		newStringArray, newOk := v.([]string)
		if !oldOk || !newOk || len(oldStringArray) != len(newStringArray) {
			return false
		}
		sort.Strings(oldStringArray)
		sort.Strings(newStringArray)
		for i := range oldStringArray {
			if oldStringArray[i] != newStringArray[i] {
				return false
			}
		}
		return true
	} else if x == "[]interface {}" {

		oldArray, _ := oldM[k].([]interface{})
		newArray, _ := v.([]interface{})

		if len(oldArray) != len(newArray) {
			return false
		}
		sortedOldArray := sortArray(oldArray)
		sortedNewArray := sortArray(newArray)

		for i := range sortedNewArray {
			oldItem := sortedOldArray[i]
			newItem := sortedNewArray[i]

			oldMap, oldIsMap := oldItem.(map[string]interface{})
			newMap, newIsMap := newItem.(map[string]interface{})

			if oldIsMap && newIsMap {
				if !compareMapRecursively(oldMap, newMap) {
					return false
				}
			}
		}
		return true
	} else {
		log.Printf("Type did not match: %s", x)
	}
	return b
}

func compareMapRecursively(oldMap, newMap map[string]interface{}) bool {
	for k1, v1 := range newMap {
		b := recursiveValueCheck(oldMap, k1, v1)
		if b == false {
			return false
		}
	}
	return true
}

func checkWorkflowStatus(conn *Config, w models.WorkflowWorkflowInfoRelationship) (string, error) {
	var workflowInfo models.WorkflowWorkflowInfo
	for { // infinite loop
		moid := w.MoMoRef.GetMoid()
		if len(moid) == 0 {
			// nothing to wait for workflowInfo Moid does not exist
			return fmt.Sprintf("workflow moid is empty %v", w), nil
		}
		// accepting pointer to workflowInfo and assigning to workflowInfo after de-referencing
		workflowInfoRef, _, responseErr := conn.ApiClient.WorkflowApi.GetWorkflowWorkflowInfoByMoid(conn.ctx, moid).Execute()
		if responseErr != nil {
			return "", responseErr
		}
		workflowInfo = *workflowInfoRef
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

func sortTags(tags []map[string]interface{}) {
	sort.SliceStable(tags, func(i, j int) bool {
		return tags[i]["key"].(string) < tags[j]["key"].(string)
	})
}

// CustomizeTagDiff Customize tag diff to ignore cisco.meta tags
func CustomizeTagDiff(tags interface{}, diff *schema.ResourceDiff) error {
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
		sortTags(newTags)
		err := diff.SetNew("tags", newTags)
		return err
	}
	return nil
}

func isNonEmptySliceOfMaps(v interface{}) bool {
	if vSlice, ok := v.([]interface{}); ok && len(vSlice) > 0 {
		if _, ok := vSlice[0].(map[string]interface{}); ok {
			return ok
		}
	}
	return false
}

func ReorderPolicyBucket(old, new []interface{}) []interface{} {
	if len(old) == 0 {
		return new
	}

	oldMap := make(map[string]interface{})
	newMap := make(map[string]interface{})
	var oldMoidArray []string

	for _, item := range old {
		m := item.(map[string]interface{})
		moid := m["moid"].(string)
		oldMap[moid] = m
		oldMoidArray = append(oldMoidArray, moid)
	}

	for _, item := range new {
		m := item.(map[string]interface{})
		moid := m["moid"].(string)
		newMap[moid] = m
	}

	var newModArray []interface{}

	// Arrange newArray according to oldArray
	for _, moid := range oldMoidArray {
		newPolicy, ok := newMap[moid]
		if ok {
			newModArray = append(newModArray, newPolicy)
			delete(newMap, moid)
		}
	}

	// Append left out data to newArray
	for _, item := range new {
		m := item.(map[string]interface{})
		moid := m["moid"].(string)
		newPolicy, ok := newMap[moid]
		if ok {
			newModArray = append(newModArray, newPolicy)
		}
	}

	return newModArray
}

// CustomizePolicyBucketDiff Customizes the diff for policy_bucket to handle ordering
func CustomizePolicyBucketDiff(policyBuckets interface{}, diff *schema.ResourceDiff) error {
	policyBucketArray := policyBuckets.([]interface{})
	log.Println("policy_bucket", policyBucketArray)
	old, _new := diff.GetChange("policy_bucket")
	oldArray := old.([]interface{})
	newArray := _new.([]interface{})
	log.Printf("old Array %#v", oldArray)
	log.Printf("new Array %#v", newArray)

	newModArray := ReorderPolicyBucket(oldArray, newArray)
	log.Printf("new mod Array %#v", newModArray)

	err := diff.SetNew("policy_bucket", newModArray)
	return err
}

func CombinedCustomizeDiff(ctx context.Context, diff *schema.ResourceDiff, i interface{}) error {
	// Handle tag diff logic
	var err error
	if tags, ok := diff.GetOk("tags"); ok {
		err = CustomizeTagDiff(tags, diff)
	}
	// Handle policy bucket diff logic
	if policy_bucket, ok := diff.GetOk("policy_bucket"); ok {
		err = CustomizePolicyBucketDiff(policy_bucket, diff)
	}
	return err
}

// parseVlanRange converts a VLAN range string (e.g., "1-5") to a slice of integers
func parseVlanRange(rangeStr string) ([]int, error) {
	var vlans []int

	if strings.Contains(rangeStr, "-") {
		parts := strings.Split(rangeStr, "-")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid range format: %s", rangeStr)
		}

		start, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			return nil, fmt.Errorf("invalid start range: %s", parts[0])
		}

		end, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			return nil, fmt.Errorf("invalid end range: %s", parts[1])
		}

		if start < 1 {
			return nil, fmt.Errorf("Invalid Vlan ID: %d Valid VLAN range is 1-4093.", start)
		}

		if start > end {
			return nil, fmt.Errorf("Invalid Vlan ID: start %d is greater than end %d", start, end)
		}

		if end > 4093 {
			return nil, fmt.Errorf("Invalid Vlan ID: %d Valid VLAN range is 1-4093.", end)
		}

		for i := start; i <= end; i++ {
			vlans = append(vlans, i)
		}
	} else {
		vlan, err := strconv.Atoi(strings.TrimSpace(rangeStr))
		if err != nil {
			return nil, fmt.Errorf("invalid VLAN: %s", rangeStr)
		}
		vlans = []int{vlan}
	}

	return vlans, nil
}

// parseVlanString converts a VLAN string (comma and/or hyphen separated) to a sorted slice of integers
func parseVlanString(vlanStr string) ([]int, error) {
	if vlanStr == "" {
		return []int{}, nil
	}

	var allVlans []int
	parts := strings.Split(vlanStr, ",")

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		vlans, err := parseVlanRange(part)
		if err != nil {
			return nil, err
		}

		allVlans = append(allVlans, vlans...)
	}

	// Remove duplicates and sort
	vlanMap := make(map[int]bool)
	for _, vlan := range allVlans {
		vlanMap[vlan] = true
	}

	var uniqueVlans []int
	for vlan := range vlanMap {
		uniqueVlans = append(uniqueVlans, vlan)
	}

	sort.Ints(uniqueVlans)
	return uniqueVlans, nil
}

// compareVlanRanges compares two VLAN range strings and returns true if they represent the same set of VLANs
func compareVlanRanges(oldVlans, newVlans string) bool {
	oldParsed, err := parseVlanString(oldVlans)
	if err != nil {
		log.Printf("Error parsing old VLAN string '%s': %v", oldVlans, err)
		return false
	}

	newParsed, err := parseVlanString(newVlans)
	if err != nil {
		log.Printf("Error parsing new VLAN string '%s': %v", newVlans, err)
		return false
	}

	if len(oldParsed) != len(newParsed) {
		return false
	}

	for i, vlan := range oldParsed {
		if vlan != newParsed[i] {
			return false
		}
	}

	return true
}

// SuppressDiffVlanRanges is a Terraform suppress diff function for VLAN ranges
// It handles comparison between comma-separated and hyphen-separated VLAN values
func SuppressDiffVlanRanges(k, old, new string, d *schema.ResourceData) bool {
	if old == "" && new == "" {
		return true
	}

	// Direct string comparison first
	if old == new {
		return true
	}

	// Compare VLAN ranges
	return compareVlanRanges(old, new)
}
