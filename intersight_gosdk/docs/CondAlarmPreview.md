# CondAlarmPreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cond.AlarmPreview"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cond.AlarmPreview"]
**AffectedMo** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 
**AffectedName** | Pointer to **string** | The AffectedMoDisplayName field of the alarm, known in the UI as Affected Name. | [optional] [readonly] 
**Code** | Pointer to **string** | The alarm code. For alarms mapped from UCS faults, this will be the same as the UCS fault code. | [optional] [readonly] 
**Description** | Pointer to **string** | The description of the alarm. | [optional] [readonly] 
**LastTransitionTime** | Pointer to **time.Time** | The LastTransitionTime field of the alarm. | [optional] [readonly] 
**Severity** | Pointer to **string** | The severity of the alarm. * &#x60;None&#x60; - The Enum value None represents that there is no severity. * &#x60;Info&#x60; - The Enum value Info represents the Informational level of severity. * &#x60;Critical&#x60; - The Enum value Critical represents the Critical level of severity. * &#x60;Warning&#x60; - The Enum value Warning represents the Warning level of severity. * &#x60;Cleared&#x60; - The Enum value Cleared represents that the alarm severity has been cleared. | [optional] [readonly] [default to "None"]

## Methods

### NewCondAlarmPreview

`func NewCondAlarmPreview(classId string, objectType string, ) *CondAlarmPreview`

NewCondAlarmPreview instantiates a new CondAlarmPreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCondAlarmPreviewWithDefaults

`func NewCondAlarmPreviewWithDefaults() *CondAlarmPreview`

NewCondAlarmPreviewWithDefaults instantiates a new CondAlarmPreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CondAlarmPreview) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CondAlarmPreview) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CondAlarmPreview) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CondAlarmPreview) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CondAlarmPreview) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CondAlarmPreview) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAffectedMo

`func (o *CondAlarmPreview) GetAffectedMo() MoMoRef`

GetAffectedMo returns the AffectedMo field if non-nil, zero value otherwise.

### GetAffectedMoOk

`func (o *CondAlarmPreview) GetAffectedMoOk() (*MoMoRef, bool)`

GetAffectedMoOk returns a tuple with the AffectedMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedMo

`func (o *CondAlarmPreview) SetAffectedMo(v MoMoRef)`

SetAffectedMo sets AffectedMo field to given value.

### HasAffectedMo

`func (o *CondAlarmPreview) HasAffectedMo() bool`

HasAffectedMo returns a boolean if a field has been set.

### GetAffectedName

`func (o *CondAlarmPreview) GetAffectedName() string`

GetAffectedName returns the AffectedName field if non-nil, zero value otherwise.

### GetAffectedNameOk

`func (o *CondAlarmPreview) GetAffectedNameOk() (*string, bool)`

GetAffectedNameOk returns a tuple with the AffectedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedName

`func (o *CondAlarmPreview) SetAffectedName(v string)`

SetAffectedName sets AffectedName field to given value.

### HasAffectedName

`func (o *CondAlarmPreview) HasAffectedName() bool`

HasAffectedName returns a boolean if a field has been set.

### GetCode

`func (o *CondAlarmPreview) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CondAlarmPreview) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CondAlarmPreview) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CondAlarmPreview) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDescription

`func (o *CondAlarmPreview) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CondAlarmPreview) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CondAlarmPreview) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CondAlarmPreview) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLastTransitionTime

`func (o *CondAlarmPreview) GetLastTransitionTime() time.Time`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *CondAlarmPreview) GetLastTransitionTimeOk() (*time.Time, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *CondAlarmPreview) SetLastTransitionTime(v time.Time)`

SetLastTransitionTime sets LastTransitionTime field to given value.

### HasLastTransitionTime

`func (o *CondAlarmPreview) HasLastTransitionTime() bool`

HasLastTransitionTime returns a boolean if a field has been set.

### GetSeverity

`func (o *CondAlarmPreview) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *CondAlarmPreview) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *CondAlarmPreview) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *CondAlarmPreview) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


