/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.11-2025051220
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// checks if the RecoveryAbstractBackupConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecoveryAbstractBackupConfig{}

// RecoveryAbstractBackupConfig Base Backup config which contains all the required inputs to do backup on a local or remote server.
type RecoveryAbstractBackupConfig struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The file name for the backup image. This name is added as a prefix in the name for the backup image. A unique file name for the backup image is created along with a timestamp. For example: prefix-1572431305418.
	FileNamePrefix *string `json:"FileNamePrefix,omitempty"`
	// Indicates whether the value of the 'password' property has been set.
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
	// Specifies whether the backup will be stored locally or remotely. * `Network Share` - The backup is stored remotely on a separate server. * `Local Storage` - The backup is stored locally on the endpoint.
	LocationType *string `json:"LocationType,omitempty"`
	// Password of Backup server.
	Password *string `json:"Password,omitempty"`
	// The file system path where the backup images must be stored. Include the IP address/hostname of the network share location and the complete file system path. For example: 172.29.109.234/var/backups/.
	Path *string `json:"Path,omitempty"`
	// Protocol for transferring the backup image to the network share location. * `SCP` - Secure Copy Protocol (SCP) to access the file server. * `SFTP` - SSH File Transfer Protocol (SFTP) to access file server. * `FTP` - File Transfer Protocol (FTP) to access file server.
	Protocol *string `json:"Protocol,omitempty"`
	// Number of backup copies maintained on the local or remote server. When the created backup files exceed this number, the initial backup files are overwritten in a sequential manner.
	RetentionCount *int64 `json:"RetentionCount,omitempty"`
	// Username for the backup server.
	UserName             *string `json:"UserName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecoveryAbstractBackupConfig RecoveryAbstractBackupConfig

// NewRecoveryAbstractBackupConfig instantiates a new RecoveryAbstractBackupConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecoveryAbstractBackupConfig(classId string, objectType string) *RecoveryAbstractBackupConfig {
	this := RecoveryAbstractBackupConfig{}
	this.ClassId = classId
	this.ObjectType = objectType
	var locationType string = "Network Share"
	this.LocationType = &locationType
	var protocol string = "SCP"
	this.Protocol = &protocol
	var retentionCount int64 = 10
	this.RetentionCount = &retentionCount
	return &this
}

// NewRecoveryAbstractBackupConfigWithDefaults instantiates a new RecoveryAbstractBackupConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecoveryAbstractBackupConfigWithDefaults() *RecoveryAbstractBackupConfig {
	this := RecoveryAbstractBackupConfig{}
	var locationType string = "Network Share"
	this.LocationType = &locationType
	var protocol string = "SCP"
	this.Protocol = &protocol
	var retentionCount int64 = 10
	this.RetentionCount = &retentionCount
	return &this
}

// GetClassId returns the ClassId field value
func (o *RecoveryAbstractBackupConfig) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *RecoveryAbstractBackupConfig) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *RecoveryAbstractBackupConfig) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *RecoveryAbstractBackupConfig) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *RecoveryAbstractBackupConfig) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *RecoveryAbstractBackupConfig) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFileNamePrefix returns the FileNamePrefix field value if set, zero value otherwise.
func (o *RecoveryAbstractBackupConfig) GetFileNamePrefix() string {
	if o == nil || IsNil(o.FileNamePrefix) {
		var ret string
		return ret
	}
	return *o.FileNamePrefix
}

// GetFileNamePrefixOk returns a tuple with the FileNamePrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryAbstractBackupConfig) GetFileNamePrefixOk() (*string, bool) {
	if o == nil || IsNil(o.FileNamePrefix) {
		return nil, false
	}
	return o.FileNamePrefix, true
}

// HasFileNamePrefix returns a boolean if a field has been set.
func (o *RecoveryAbstractBackupConfig) HasFileNamePrefix() bool {
	if o != nil && !IsNil(o.FileNamePrefix) {
		return true
	}

	return false
}

// SetFileNamePrefix gets a reference to the given string and assigns it to the FileNamePrefix field.
func (o *RecoveryAbstractBackupConfig) SetFileNamePrefix(v string) {
	o.FileNamePrefix = &v
}

// GetIsPasswordSet returns the IsPasswordSet field value if set, zero value otherwise.
func (o *RecoveryAbstractBackupConfig) GetIsPasswordSet() bool {
	if o == nil || IsNil(o.IsPasswordSet) {
		var ret bool
		return ret
	}
	return *o.IsPasswordSet
}

// GetIsPasswordSetOk returns a tuple with the IsPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryAbstractBackupConfig) GetIsPasswordSetOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPasswordSet) {
		return nil, false
	}
	return o.IsPasswordSet, true
}

// HasIsPasswordSet returns a boolean if a field has been set.
func (o *RecoveryAbstractBackupConfig) HasIsPasswordSet() bool {
	if o != nil && !IsNil(o.IsPasswordSet) {
		return true
	}

	return false
}

// SetIsPasswordSet gets a reference to the given bool and assigns it to the IsPasswordSet field.
func (o *RecoveryAbstractBackupConfig) SetIsPasswordSet(v bool) {
	o.IsPasswordSet = &v
}

// GetLocationType returns the LocationType field value if set, zero value otherwise.
func (o *RecoveryAbstractBackupConfig) GetLocationType() string {
	if o == nil || IsNil(o.LocationType) {
		var ret string
		return ret
	}
	return *o.LocationType
}

// GetLocationTypeOk returns a tuple with the LocationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryAbstractBackupConfig) GetLocationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.LocationType) {
		return nil, false
	}
	return o.LocationType, true
}

// HasLocationType returns a boolean if a field has been set.
func (o *RecoveryAbstractBackupConfig) HasLocationType() bool {
	if o != nil && !IsNil(o.LocationType) {
		return true
	}

	return false
}

// SetLocationType gets a reference to the given string and assigns it to the LocationType field.
func (o *RecoveryAbstractBackupConfig) SetLocationType(v string) {
	o.LocationType = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *RecoveryAbstractBackupConfig) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryAbstractBackupConfig) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *RecoveryAbstractBackupConfig) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *RecoveryAbstractBackupConfig) SetPassword(v string) {
	o.Password = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *RecoveryAbstractBackupConfig) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryAbstractBackupConfig) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *RecoveryAbstractBackupConfig) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *RecoveryAbstractBackupConfig) SetPath(v string) {
	o.Path = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *RecoveryAbstractBackupConfig) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryAbstractBackupConfig) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *RecoveryAbstractBackupConfig) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *RecoveryAbstractBackupConfig) SetProtocol(v string) {
	o.Protocol = &v
}

// GetRetentionCount returns the RetentionCount field value if set, zero value otherwise.
func (o *RecoveryAbstractBackupConfig) GetRetentionCount() int64 {
	if o == nil || IsNil(o.RetentionCount) {
		var ret int64
		return ret
	}
	return *o.RetentionCount
}

// GetRetentionCountOk returns a tuple with the RetentionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryAbstractBackupConfig) GetRetentionCountOk() (*int64, bool) {
	if o == nil || IsNil(o.RetentionCount) {
		return nil, false
	}
	return o.RetentionCount, true
}

// HasRetentionCount returns a boolean if a field has been set.
func (o *RecoveryAbstractBackupConfig) HasRetentionCount() bool {
	if o != nil && !IsNil(o.RetentionCount) {
		return true
	}

	return false
}

// SetRetentionCount gets a reference to the given int64 and assigns it to the RetentionCount field.
func (o *RecoveryAbstractBackupConfig) SetRetentionCount(v int64) {
	o.RetentionCount = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *RecoveryAbstractBackupConfig) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecoveryAbstractBackupConfig) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *RecoveryAbstractBackupConfig) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *RecoveryAbstractBackupConfig) SetUserName(v string) {
	o.UserName = &v
}

func (o RecoveryAbstractBackupConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecoveryAbstractBackupConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return map[string]interface{}{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return map[string]interface{}{}, errPolicyAbstractPolicy
	}
	toSerialize["ClassId"] = o.ClassId
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.FileNamePrefix) {
		toSerialize["FileNamePrefix"] = o.FileNamePrefix
	}
	if !IsNil(o.IsPasswordSet) {
		toSerialize["IsPasswordSet"] = o.IsPasswordSet
	}
	if !IsNil(o.LocationType) {
		toSerialize["LocationType"] = o.LocationType
	}
	if !IsNil(o.Password) {
		toSerialize["Password"] = o.Password
	}
	if !IsNil(o.Path) {
		toSerialize["Path"] = o.Path
	}
	if !IsNil(o.Protocol) {
		toSerialize["Protocol"] = o.Protocol
	}
	if !IsNil(o.RetentionCount) {
		toSerialize["RetentionCount"] = o.RetentionCount
	}
	if !IsNil(o.UserName) {
		toSerialize["UserName"] = o.UserName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RecoveryAbstractBackupConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{}
	var defaultValueApplied bool
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			if _, ok := defaultValueFuncMap[requiredProperty]; ok {
				allProperties[requiredProperty] = defaultValueFuncMap[requiredProperty]()
				defaultValueApplied = true
			}
		}
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	if defaultValueApplied {
		data, err = json.Marshal(allProperties)
		if err != nil {
			return err
		}
	}
	type RecoveryAbstractBackupConfigWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// The file name for the backup image. This name is added as a prefix in the name for the backup image. A unique file name for the backup image is created along with a timestamp. For example: prefix-1572431305418.
		FileNamePrefix *string `json:"FileNamePrefix,omitempty"`
		// Indicates whether the value of the 'password' property has been set.
		IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
		// Specifies whether the backup will be stored locally or remotely. * `Network Share` - The backup is stored remotely on a separate server. * `Local Storage` - The backup is stored locally on the endpoint.
		LocationType *string `json:"LocationType,omitempty"`
		// Password of Backup server.
		Password *string `json:"Password,omitempty"`
		// The file system path where the backup images must be stored. Include the IP address/hostname of the network share location and the complete file system path. For example: 172.29.109.234/var/backups/.
		Path *string `json:"Path,omitempty"`
		// Protocol for transferring the backup image to the network share location. * `SCP` - Secure Copy Protocol (SCP) to access the file server. * `SFTP` - SSH File Transfer Protocol (SFTP) to access file server. * `FTP` - File Transfer Protocol (FTP) to access file server.
		Protocol *string `json:"Protocol,omitempty"`
		// Number of backup copies maintained on the local or remote server. When the created backup files exceed this number, the initial backup files are overwritten in a sequential manner.
		RetentionCount *int64 `json:"RetentionCount,omitempty"`
		// Username for the backup server.
		UserName *string `json:"UserName,omitempty"`
	}

	varRecoveryAbstractBackupConfigWithoutEmbeddedStruct := RecoveryAbstractBackupConfigWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varRecoveryAbstractBackupConfigWithoutEmbeddedStruct)
	if err == nil {
		varRecoveryAbstractBackupConfig := _RecoveryAbstractBackupConfig{}
		varRecoveryAbstractBackupConfig.ClassId = varRecoveryAbstractBackupConfigWithoutEmbeddedStruct.ClassId
		varRecoveryAbstractBackupConfig.ObjectType = varRecoveryAbstractBackupConfigWithoutEmbeddedStruct.ObjectType
		varRecoveryAbstractBackupConfig.FileNamePrefix = varRecoveryAbstractBackupConfigWithoutEmbeddedStruct.FileNamePrefix
		varRecoveryAbstractBackupConfig.IsPasswordSet = varRecoveryAbstractBackupConfigWithoutEmbeddedStruct.IsPasswordSet
		varRecoveryAbstractBackupConfig.LocationType = varRecoveryAbstractBackupConfigWithoutEmbeddedStruct.LocationType
		varRecoveryAbstractBackupConfig.Password = varRecoveryAbstractBackupConfigWithoutEmbeddedStruct.Password
		varRecoveryAbstractBackupConfig.Path = varRecoveryAbstractBackupConfigWithoutEmbeddedStruct.Path
		varRecoveryAbstractBackupConfig.Protocol = varRecoveryAbstractBackupConfigWithoutEmbeddedStruct.Protocol
		varRecoveryAbstractBackupConfig.RetentionCount = varRecoveryAbstractBackupConfigWithoutEmbeddedStruct.RetentionCount
		varRecoveryAbstractBackupConfig.UserName = varRecoveryAbstractBackupConfigWithoutEmbeddedStruct.UserName
		*o = RecoveryAbstractBackupConfig(varRecoveryAbstractBackupConfig)
	} else {
		return err
	}

	varRecoveryAbstractBackupConfig := _RecoveryAbstractBackupConfig{}

	err = json.Unmarshal(data, &varRecoveryAbstractBackupConfig)
	if err == nil {
		o.PolicyAbstractPolicy = varRecoveryAbstractBackupConfig.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "FileNamePrefix")
		delete(additionalProperties, "IsPasswordSet")
		delete(additionalProperties, "LocationType")
		delete(additionalProperties, "Password")
		delete(additionalProperties, "Path")
		delete(additionalProperties, "Protocol")
		delete(additionalProperties, "RetentionCount")
		delete(additionalProperties, "UserName")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicy := reflect.ValueOf(o.PolicyAbstractPolicy)
		for i := 0; i < reflectPolicyAbstractPolicy.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicy.Type().Field(i)

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

type NullableRecoveryAbstractBackupConfig struct {
	value *RecoveryAbstractBackupConfig
	isSet bool
}

func (v NullableRecoveryAbstractBackupConfig) Get() *RecoveryAbstractBackupConfig {
	return v.value
}

func (v *NullableRecoveryAbstractBackupConfig) Set(val *RecoveryAbstractBackupConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableRecoveryAbstractBackupConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableRecoveryAbstractBackupConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecoveryAbstractBackupConfig(val *RecoveryAbstractBackupConfig) *NullableRecoveryAbstractBackupConfig {
	return &NullableRecoveryAbstractBackupConfig{value: val, isSet: true}
}

func (v NullableRecoveryAbstractBackupConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecoveryAbstractBackupConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
