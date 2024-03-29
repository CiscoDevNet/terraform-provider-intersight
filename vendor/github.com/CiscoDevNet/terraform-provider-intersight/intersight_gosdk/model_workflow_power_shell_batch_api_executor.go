/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-15711
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// WorkflowPowerShellBatchApiExecutor Intersight allows generic API tasks to be created by taking the API request body and a response parser specification in the form of content.Grammar object. Batch API associates the list of API requests to be executed as part of single task execution. Each API request takes the request body and a response parser specification.
type WorkflowPowerShellBatchApiExecutor struct {
	WorkflowBatchExecutor
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                    `json:"ObjectType"`
	ErrorResponseHandler *WorkflowErrorResponseHandlerRelationship `json:"ErrorResponseHandler,omitempty"`
	TaskDefinition       *WorkflowTaskDefinitionRelationship       `json:"TaskDefinition,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowPowerShellBatchApiExecutor WorkflowPowerShellBatchApiExecutor

// NewWorkflowPowerShellBatchApiExecutor instantiates a new WorkflowPowerShellBatchApiExecutor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowPowerShellBatchApiExecutor(classId string, objectType string) *WorkflowPowerShellBatchApiExecutor {
	this := WorkflowPowerShellBatchApiExecutor{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowPowerShellBatchApiExecutorWithDefaults instantiates a new WorkflowPowerShellBatchApiExecutor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowPowerShellBatchApiExecutorWithDefaults() *WorkflowPowerShellBatchApiExecutor {
	this := WorkflowPowerShellBatchApiExecutor{}
	var classId string = "workflow.PowerShellBatchApiExecutor"
	this.ClassId = classId
	var objectType string = "workflow.PowerShellBatchApiExecutor"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowPowerShellBatchApiExecutor) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowPowerShellBatchApiExecutor) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowPowerShellBatchApiExecutor) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowPowerShellBatchApiExecutor) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowPowerShellBatchApiExecutor) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowPowerShellBatchApiExecutor) SetObjectType(v string) {
	o.ObjectType = v
}

// GetErrorResponseHandler returns the ErrorResponseHandler field value if set, zero value otherwise.
func (o *WorkflowPowerShellBatchApiExecutor) GetErrorResponseHandler() WorkflowErrorResponseHandlerRelationship {
	if o == nil || o.ErrorResponseHandler == nil {
		var ret WorkflowErrorResponseHandlerRelationship
		return ret
	}
	return *o.ErrorResponseHandler
}

// GetErrorResponseHandlerOk returns a tuple with the ErrorResponseHandler field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPowerShellBatchApiExecutor) GetErrorResponseHandlerOk() (*WorkflowErrorResponseHandlerRelationship, bool) {
	if o == nil || o.ErrorResponseHandler == nil {
		return nil, false
	}
	return o.ErrorResponseHandler, true
}

// HasErrorResponseHandler returns a boolean if a field has been set.
func (o *WorkflowPowerShellBatchApiExecutor) HasErrorResponseHandler() bool {
	if o != nil && o.ErrorResponseHandler != nil {
		return true
	}

	return false
}

// SetErrorResponseHandler gets a reference to the given WorkflowErrorResponseHandlerRelationship and assigns it to the ErrorResponseHandler field.
func (o *WorkflowPowerShellBatchApiExecutor) SetErrorResponseHandler(v WorkflowErrorResponseHandlerRelationship) {
	o.ErrorResponseHandler = &v
}

// GetTaskDefinition returns the TaskDefinition field value if set, zero value otherwise.
func (o *WorkflowPowerShellBatchApiExecutor) GetTaskDefinition() WorkflowTaskDefinitionRelationship {
	if o == nil || o.TaskDefinition == nil {
		var ret WorkflowTaskDefinitionRelationship
		return ret
	}
	return *o.TaskDefinition
}

// GetTaskDefinitionOk returns a tuple with the TaskDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowPowerShellBatchApiExecutor) GetTaskDefinitionOk() (*WorkflowTaskDefinitionRelationship, bool) {
	if o == nil || o.TaskDefinition == nil {
		return nil, false
	}
	return o.TaskDefinition, true
}

// HasTaskDefinition returns a boolean if a field has been set.
func (o *WorkflowPowerShellBatchApiExecutor) HasTaskDefinition() bool {
	if o != nil && o.TaskDefinition != nil {
		return true
	}

	return false
}

// SetTaskDefinition gets a reference to the given WorkflowTaskDefinitionRelationship and assigns it to the TaskDefinition field.
func (o *WorkflowPowerShellBatchApiExecutor) SetTaskDefinition(v WorkflowTaskDefinitionRelationship) {
	o.TaskDefinition = &v
}

func (o WorkflowPowerShellBatchApiExecutor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedWorkflowBatchExecutor, errWorkflowBatchExecutor := json.Marshal(o.WorkflowBatchExecutor)
	if errWorkflowBatchExecutor != nil {
		return []byte{}, errWorkflowBatchExecutor
	}
	errWorkflowBatchExecutor = json.Unmarshal([]byte(serializedWorkflowBatchExecutor), &toSerialize)
	if errWorkflowBatchExecutor != nil {
		return []byte{}, errWorkflowBatchExecutor
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ErrorResponseHandler != nil {
		toSerialize["ErrorResponseHandler"] = o.ErrorResponseHandler
	}
	if o.TaskDefinition != nil {
		toSerialize["TaskDefinition"] = o.TaskDefinition
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowPowerShellBatchApiExecutor) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowPowerShellBatchApiExecutorWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType           string                                    `json:"ObjectType"`
		ErrorResponseHandler *WorkflowErrorResponseHandlerRelationship `json:"ErrorResponseHandler,omitempty"`
		TaskDefinition       *WorkflowTaskDefinitionRelationship       `json:"TaskDefinition,omitempty"`
	}

	varWorkflowPowerShellBatchApiExecutorWithoutEmbeddedStruct := WorkflowPowerShellBatchApiExecutorWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowPowerShellBatchApiExecutorWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowPowerShellBatchApiExecutor := _WorkflowPowerShellBatchApiExecutor{}
		varWorkflowPowerShellBatchApiExecutor.ClassId = varWorkflowPowerShellBatchApiExecutorWithoutEmbeddedStruct.ClassId
		varWorkflowPowerShellBatchApiExecutor.ObjectType = varWorkflowPowerShellBatchApiExecutorWithoutEmbeddedStruct.ObjectType
		varWorkflowPowerShellBatchApiExecutor.ErrorResponseHandler = varWorkflowPowerShellBatchApiExecutorWithoutEmbeddedStruct.ErrorResponseHandler
		varWorkflowPowerShellBatchApiExecutor.TaskDefinition = varWorkflowPowerShellBatchApiExecutorWithoutEmbeddedStruct.TaskDefinition
		*o = WorkflowPowerShellBatchApiExecutor(varWorkflowPowerShellBatchApiExecutor)
	} else {
		return err
	}

	varWorkflowPowerShellBatchApiExecutor := _WorkflowPowerShellBatchApiExecutor{}

	err = json.Unmarshal(bytes, &varWorkflowPowerShellBatchApiExecutor)
	if err == nil {
		o.WorkflowBatchExecutor = varWorkflowPowerShellBatchApiExecutor.WorkflowBatchExecutor
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ErrorResponseHandler")
		delete(additionalProperties, "TaskDefinition")

		// remove fields from embedded structs
		reflectWorkflowBatchExecutor := reflect.ValueOf(o.WorkflowBatchExecutor)
		for i := 0; i < reflectWorkflowBatchExecutor.Type().NumField(); i++ {
			t := reflectWorkflowBatchExecutor.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowPowerShellBatchApiExecutor struct {
	value *WorkflowPowerShellBatchApiExecutor
	isSet bool
}

func (v NullableWorkflowPowerShellBatchApiExecutor) Get() *WorkflowPowerShellBatchApiExecutor {
	return v.value
}

func (v *NullableWorkflowPowerShellBatchApiExecutor) Set(val *WorkflowPowerShellBatchApiExecutor) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowPowerShellBatchApiExecutor) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowPowerShellBatchApiExecutor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowPowerShellBatchApiExecutor(val *WorkflowPowerShellBatchApiExecutor) *NullableWorkflowPowerShellBatchApiExecutor {
	return &NullableWorkflowPowerShellBatchApiExecutor{value: val, isSet: true}
}

func (v NullableWorkflowPowerShellBatchApiExecutor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowPowerShellBatchApiExecutor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
