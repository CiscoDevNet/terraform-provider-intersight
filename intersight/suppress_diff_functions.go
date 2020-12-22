package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
func getRequestParams(in []byte) string {
	var o string
	var s map[string]interface{}
	err := json.Unmarshal(in, &s)
	if err != nil {
		return ""
	}
	for k, v := range s {
		if k == "ClassId" || k == "ObjectType" {
			continue
		}
		log.Printf("Type: %+v", reflect.TypeOf(v).Kind())
		switch reflect.TypeOf(v).Kind() {
		case reflect.String:
			o += k + " eq '" + v.(string) + "'"
		case reflect.Bool:
			o += k + " eq " + strconv.FormatBool(v.(bool))
		case reflect.Int:
			o += k + " eq " + strconv.FormatInt(v.(int64), 10)
		case reflect.Float64:
			o += k + " eq " + fmt.Sprintf("%f", v.(float64))
		}
		o += " and "
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
