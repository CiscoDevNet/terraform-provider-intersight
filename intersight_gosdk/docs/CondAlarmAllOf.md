# CondAlarmAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acknowledge** | Pointer to **string** | Alarm acknowledgment state. Default value is None. * &#x60;None&#x60; - The Enum value None represents that the alarm is not acknowledged and is included as part of health status and overall alarm count. * &#x60;Acknowledge&#x60; - The Enum value Acknowledge represents that the alarm is acknowledged by user. The alarm will be ignored from the health status and overall alarm count. | [optional] [default to "None"]
**AcknowledgeBy** | Pointer to **string** | User who acknowledged the alarm. | [optional] [readonly] 
**AcknowledgeTime** | Pointer to [**time.Time**](time.Time.md) | Time at which the alarm was acknowledged by the user. | [optional] [readonly] 
**AffectedMoId** | Pointer to **string** | MoId of the affected object from the managed system&#39;s point of view. | [optional] [readonly] 
**AffectedMoType** | Pointer to **string** | Managed system affected object type. For example Adaptor, FI, CIMC. | [optional] [readonly] 
**AffectedObject** | Pointer to **string** | A unique key for an alarm instance, consists of a combination of a unique system name and msAffectedObject. | [optional] [readonly] 
**AncestorMoId** | Pointer to **string** | Parent MoId of the fault from managed system. For example, Blade moid for adaptor fault. | [optional] [readonly] 
**AncestorMoType** | Pointer to **string** | Parent MO type of the fault from managed system. For example, Blade for adaptor fault. | [optional] [readonly] 
**Code** | Pointer to **string** | A unique alarm code. For alarms mapped from UCS faults, this will be the same as the UCS fault code. | [optional] [readonly] 
**CreationTime** | Pointer to [**time.Time**](time.Time.md) | The time the alarm was created. | [optional] [readonly] 
**Description** | Pointer to **string** | A longer description of the alarm than the name. The description contains details of the component reporting the issue. | [optional] [readonly] 
**LastTransitionTime** | Pointer to [**time.Time**](time.Time.md) | The time the alarm last had a change in severity. | [optional] [readonly] 
**MsAffectedObject** | Pointer to **string** | A unique key for the alarm from the managed system&#39;s point of view. For example, in the case of UCS, this is the fault&#39;s dn. | [optional] [readonly] 
**Name** | Pointer to **string** | Uniquely identifies the type of alarm. For alarms originating from Intersight, this will be a descriptive name. For alarms that are mapped from faults, the name will be derived from fault properties. For example, alarms mapped from UCS faults will use a prefix of UCS and appended with the fault code. | [optional] [readonly] 
**OrigSeverity** | Pointer to **string** | The original severity when the alarm was first created. * &#x60;None&#x60; - The Enum value None represents that there is no severity. * &#x60;Info&#x60; - The Enum value Info represents the Informational level of severity. * &#x60;Critical&#x60; - The Enum value Critical represents the Critical level of severity. * &#x60;Warning&#x60; - The Enum value Warning represents the Warning level of severity. * &#x60;Cleared&#x60; - The Enum value Cleared represents that the alarm severity has been cleared. | [optional] [readonly] [default to "None"]
**Severity** | Pointer to **string** | The severity of the alarm. Valid values are Critical, Warning, Info, and Cleared. * &#x60;None&#x60; - The Enum value None represents that there is no severity. * &#x60;Info&#x60; - The Enum value Info represents the Informational level of severity. * &#x60;Critical&#x60; - The Enum value Critical represents the Critical level of severity. * &#x60;Warning&#x60; - The Enum value Warning represents the Warning level of severity. * &#x60;Cleared&#x60; - The Enum value Cleared represents that the alarm severity has been cleared. | [optional] [readonly] [default to "None"]
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewCondAlarmAllOf

`func NewCondAlarmAllOf() *CondAlarmAllOf`

NewCondAlarmAllOf instantiates a new CondAlarmAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCondAlarmAllOfWithDefaults

`func NewCondAlarmAllOfWithDefaults() *CondAlarmAllOf`

NewCondAlarmAllOfWithDefaults instantiates a new CondAlarmAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcknowledge

`func (o *CondAlarmAllOf) GetAcknowledge() string`

GetAcknowledge returns the Acknowledge field if non-nil, zero value otherwise.

### GetAcknowledgeOk

`func (o *CondAlarmAllOf) GetAcknowledgeOk() (*string, bool)`

GetAcknowledgeOk returns a tuple with the Acknowledge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledge

`func (o *CondAlarmAllOf) SetAcknowledge(v string)`

SetAcknowledge sets Acknowledge field to given value.

### HasAcknowledge

`func (o *CondAlarmAllOf) HasAcknowledge() bool`

HasAcknowledge returns a boolean if a field has been set.

### GetAcknowledgeBy

`func (o *CondAlarmAllOf) GetAcknowledgeBy() string`

GetAcknowledgeBy returns the AcknowledgeBy field if non-nil, zero value otherwise.

### GetAcknowledgeByOk

`func (o *CondAlarmAllOf) GetAcknowledgeByOk() (*string, bool)`

GetAcknowledgeByOk returns a tuple with the AcknowledgeBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgeBy

`func (o *CondAlarmAllOf) SetAcknowledgeBy(v string)`

SetAcknowledgeBy sets AcknowledgeBy field to given value.

### HasAcknowledgeBy

`func (o *CondAlarmAllOf) HasAcknowledgeBy() bool`

HasAcknowledgeBy returns a boolean if a field has been set.

### GetAcknowledgeTime

`func (o *CondAlarmAllOf) GetAcknowledgeTime() time.Time`

GetAcknowledgeTime returns the AcknowledgeTime field if non-nil, zero value otherwise.

### GetAcknowledgeTimeOk

`func (o *CondAlarmAllOf) GetAcknowledgeTimeOk() (*time.Time, bool)`

GetAcknowledgeTimeOk returns a tuple with the AcknowledgeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgeTime

`func (o *CondAlarmAllOf) SetAcknowledgeTime(v time.Time)`

SetAcknowledgeTime sets AcknowledgeTime field to given value.

### HasAcknowledgeTime

`func (o *CondAlarmAllOf) HasAcknowledgeTime() bool`

HasAcknowledgeTime returns a boolean if a field has been set.

### GetAffectedMoId

`func (o *CondAlarmAllOf) GetAffectedMoId() string`

GetAffectedMoId returns the AffectedMoId field if non-nil, zero value otherwise.

### GetAffectedMoIdOk

`func (o *CondAlarmAllOf) GetAffectedMoIdOk() (*string, bool)`

GetAffectedMoIdOk returns a tuple with the AffectedMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedMoId

`func (o *CondAlarmAllOf) SetAffectedMoId(v string)`

SetAffectedMoId sets AffectedMoId field to given value.

### HasAffectedMoId

`func (o *CondAlarmAllOf) HasAffectedMoId() bool`

HasAffectedMoId returns a boolean if a field has been set.

### GetAffectedMoType

`func (o *CondAlarmAllOf) GetAffectedMoType() string`

GetAffectedMoType returns the AffectedMoType field if non-nil, zero value otherwise.

### GetAffectedMoTypeOk

`func (o *CondAlarmAllOf) GetAffectedMoTypeOk() (*string, bool)`

GetAffectedMoTypeOk returns a tuple with the AffectedMoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedMoType

`func (o *CondAlarmAllOf) SetAffectedMoType(v string)`

SetAffectedMoType sets AffectedMoType field to given value.

### HasAffectedMoType

`func (o *CondAlarmAllOf) HasAffectedMoType() bool`

HasAffectedMoType returns a boolean if a field has been set.

### GetAffectedObject

`func (o *CondAlarmAllOf) GetAffectedObject() string`

GetAffectedObject returns the AffectedObject field if non-nil, zero value otherwise.

### GetAffectedObjectOk

`func (o *CondAlarmAllOf) GetAffectedObjectOk() (*string, bool)`

GetAffectedObjectOk returns a tuple with the AffectedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedObject

`func (o *CondAlarmAllOf) SetAffectedObject(v string)`

SetAffectedObject sets AffectedObject field to given value.

### HasAffectedObject

`func (o *CondAlarmAllOf) HasAffectedObject() bool`

HasAffectedObject returns a boolean if a field has been set.

### GetAncestorMoId

`func (o *CondAlarmAllOf) GetAncestorMoId() string`

GetAncestorMoId returns the AncestorMoId field if non-nil, zero value otherwise.

### GetAncestorMoIdOk

`func (o *CondAlarmAllOf) GetAncestorMoIdOk() (*string, bool)`

GetAncestorMoIdOk returns a tuple with the AncestorMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestorMoId

`func (o *CondAlarmAllOf) SetAncestorMoId(v string)`

SetAncestorMoId sets AncestorMoId field to given value.

### HasAncestorMoId

`func (o *CondAlarmAllOf) HasAncestorMoId() bool`

HasAncestorMoId returns a boolean if a field has been set.

### GetAncestorMoType

`func (o *CondAlarmAllOf) GetAncestorMoType() string`

GetAncestorMoType returns the AncestorMoType field if non-nil, zero value otherwise.

### GetAncestorMoTypeOk

`func (o *CondAlarmAllOf) GetAncestorMoTypeOk() (*string, bool)`

GetAncestorMoTypeOk returns a tuple with the AncestorMoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestorMoType

`func (o *CondAlarmAllOf) SetAncestorMoType(v string)`

SetAncestorMoType sets AncestorMoType field to given value.

### HasAncestorMoType

`func (o *CondAlarmAllOf) HasAncestorMoType() bool`

HasAncestorMoType returns a boolean if a field has been set.

### GetCode

`func (o *CondAlarmAllOf) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CondAlarmAllOf) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CondAlarmAllOf) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CondAlarmAllOf) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCreationTime

`func (o *CondAlarmAllOf) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *CondAlarmAllOf) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *CondAlarmAllOf) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *CondAlarmAllOf) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetDescription

`func (o *CondAlarmAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CondAlarmAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CondAlarmAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CondAlarmAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLastTransitionTime

`func (o *CondAlarmAllOf) GetLastTransitionTime() time.Time`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *CondAlarmAllOf) GetLastTransitionTimeOk() (*time.Time, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *CondAlarmAllOf) SetLastTransitionTime(v time.Time)`

SetLastTransitionTime sets LastTransitionTime field to given value.

### HasLastTransitionTime

`func (o *CondAlarmAllOf) HasLastTransitionTime() bool`

HasLastTransitionTime returns a boolean if a field has been set.

### GetMsAffectedObject

`func (o *CondAlarmAllOf) GetMsAffectedObject() string`

GetMsAffectedObject returns the MsAffectedObject field if non-nil, zero value otherwise.

### GetMsAffectedObjectOk

`func (o *CondAlarmAllOf) GetMsAffectedObjectOk() (*string, bool)`

GetMsAffectedObjectOk returns a tuple with the MsAffectedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAffectedObject

`func (o *CondAlarmAllOf) SetMsAffectedObject(v string)`

SetMsAffectedObject sets MsAffectedObject field to given value.

### HasMsAffectedObject

`func (o *CondAlarmAllOf) HasMsAffectedObject() bool`

HasMsAffectedObject returns a boolean if a field has been set.

### GetName

`func (o *CondAlarmAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CondAlarmAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CondAlarmAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CondAlarmAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrigSeverity

`func (o *CondAlarmAllOf) GetOrigSeverity() string`

GetOrigSeverity returns the OrigSeverity field if non-nil, zero value otherwise.

### GetOrigSeverityOk

`func (o *CondAlarmAllOf) GetOrigSeverityOk() (*string, bool)`

GetOrigSeverityOk returns a tuple with the OrigSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigSeverity

`func (o *CondAlarmAllOf) SetOrigSeverity(v string)`

SetOrigSeverity sets OrigSeverity field to given value.

### HasOrigSeverity

`func (o *CondAlarmAllOf) HasOrigSeverity() bool`

HasOrigSeverity returns a boolean if a field has been set.

### GetSeverity

`func (o *CondAlarmAllOf) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *CondAlarmAllOf) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *CondAlarmAllOf) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *CondAlarmAllOf) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *CondAlarmAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *CondAlarmAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *CondAlarmAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *CondAlarmAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


