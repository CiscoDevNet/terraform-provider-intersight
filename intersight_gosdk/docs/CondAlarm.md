# CondAlarm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cond.Alarm"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cond.Alarm"]
**Acknowledge** | Pointer to **string** | Alarm acknowledgment state. Default value is None. * &#x60;None&#x60; - The Enum value None represents that the alarm is not acknowledged and is included as part of health status and overall alarm count. * &#x60;Acknowledge&#x60; - The Enum value Acknowledge represents that the alarm is acknowledged by user. The alarm will be ignored from the health status and overall alarm count. | [optional] [default to "None"]
**AcknowledgeBy** | Pointer to **string** | User who acknowledged the alarm. | [optional] [readonly] 
**AcknowledgeTime** | Pointer to **time.Time** | Time at which the alarm was acknowledged by the user. | [optional] [readonly] 
**AffectedMoDisplayName** | Pointer to **string** | Display name of the affected object on which the alarm is raised. The name uniquely identifies an alarm target such that it can be distinguished from similar objects managed by Intersight. | [optional] [readonly] 
**AffectedMoId** | Pointer to **string** | MoId of the affected object from the managed system&#39;s point of view. | [optional] [readonly] 
**AffectedMoType** | Pointer to **string** | Managed system affected object type. For example Adaptor, FI, CIMC. | [optional] [readonly] 
**AffectedObject** | Pointer to **string** | A unique key for an alarm instance, consists of a combination of a unique system name and msAffectedObject. | [optional] [readonly] 
**AncestorMoId** | Pointer to **string** | Parent MoId of the fault from managed system. For example, Blade moid for adaptor fault. | [optional] [readonly] 
**AncestorMoType** | Pointer to **string** | Parent MO type of the fault from managed system. For example, Blade for adaptor fault. | [optional] [readonly] 
**Code** | Pointer to **string** | A unique alarm code. For alarms mapped from UCS faults, this will be the same as the UCS fault code. | [optional] [readonly] 
**CreationTime** | Pointer to **time.Time** | The time the alarm was created. | [optional] [readonly] 
**Description** | Pointer to **string** | A longer description of the alarm than the name. The description contains details of the component reporting the issue. | [optional] [readonly] 
**Flapping** | Pointer to **string** | Alarm flapping state. This will be set to Flapping or Cooldown if both (A) this type of alarm is being monitored for flapping conditions, and (B) the alarm has recently transitioned to an active state (Critical, Warning or Info) followed by a Cleared state or vice versa. LastTransitionTime is a better field to use to know whether a particular alarm recently changed state. * &#x60;NotFlapping&#x60; - The enum value None says that no recent flaps have occurred. * &#x60;Flapping&#x60; - The enum value Flapping says that the alarm has become active recently, after being active and then cleared previously. * &#x60;Cooldown&#x60; - The enum value Cooldown says that the alarm is cleared, but was recently active. * &#x60;Unknown&#x60; - The enum value Unknown indicates that you might not have the latest version of the property meta. | [optional] [readonly] [default to "NotFlapping"]
**FlappingCount** | Pointer to **int64** | Alarm flapping counter. This will be incremented every time the state of the alarm transitions to an active state (Critical, Warning or Info) followed by a Cleared state or vice versa. If no more transitions occur within the system-defined flap interval (usually less than 5 minutes), the counter will be reset to zero. This represents the amount of times the alarm has flapped between an active and a cleared state since the last time the Flapping state was cleared. | [optional] [readonly] 
**FlappingStartTime** | Pointer to **time.Time** | Alarm flapping start time. Only when the flapping state is Flapping or Cooldown, this will be set to the time the alarm began flapping. If the flapping state is NotFlapping, this timestamp may be set to zero or any other time and should be ignored. | [optional] [readonly] 
**LastTransitionTime** | Pointer to **time.Time** | The time the alarm last had a change in severity. | [optional] [readonly] 
**MsAffectedObject** | Pointer to **string** | A unique key for the alarm from the managed system&#39;s point of view. For example, in the case of UCS, this is the fault&#39;s dn. | [optional] [readonly] 
**Name** | Pointer to **string** | Uniquely identifies the type of alarm. For alarms originating from Intersight, this will be a descriptive name. For alarms that are mapped from faults, the name will be derived from fault properties. For example, alarms mapped from UCS faults will use a prefix of UCS and appended with the fault code. | [optional] [readonly] 
**OrigSeverity** | Pointer to **string** | The original severity when the alarm was first created. * &#x60;None&#x60; - The Enum value None represents that there is no severity. * &#x60;Info&#x60; - The Enum value Info represents the Informational level of severity. * &#x60;Critical&#x60; - The Enum value Critical represents the Critical level of severity. * &#x60;Warning&#x60; - The Enum value Warning represents the Warning level of severity. * &#x60;Cleared&#x60; - The Enum value Cleared represents that the alarm severity has been cleared. | [optional] [readonly] [default to "None"]
**Severity** | Pointer to **string** | The severity of the alarm. Valid values are Critical, Warning, Info, and Cleared. * &#x60;None&#x60; - The Enum value None represents that there is no severity. * &#x60;Info&#x60; - The Enum value Info represents the Informational level of severity. * &#x60;Critical&#x60; - The Enum value Critical represents the Critical level of severity. * &#x60;Warning&#x60; - The Enum value Warning represents the Warning level of severity. * &#x60;Cleared&#x60; - The Enum value Cleared represents that the alarm severity has been cleared. | [optional] [readonly] [default to "None"]
**Suppressed** | Pointer to **bool** | Indicates whether the alarm is marked for suppression or not. | [optional] 
**AffectedMo** | Pointer to [**NullableMoBaseMoRelationship**](MoBaseMoRelationship.md) |  | [optional] 
**AlarmSummaryAggregators** | Pointer to [**[]MoBaseMoRelationship**](MoBaseMoRelationship.md) | An array of relationships to moBaseMo resources. | [optional] [readonly] 
**Definition** | Pointer to [**NullableCondAlarmDefinitionRelationship**](CondAlarmDefinitionRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewCondAlarm

`func NewCondAlarm(classId string, objectType string, ) *CondAlarm`

NewCondAlarm instantiates a new CondAlarm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCondAlarmWithDefaults

`func NewCondAlarmWithDefaults() *CondAlarm`

NewCondAlarmWithDefaults instantiates a new CondAlarm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CondAlarm) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CondAlarm) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CondAlarm) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CondAlarm) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CondAlarm) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CondAlarm) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAcknowledge

`func (o *CondAlarm) GetAcknowledge() string`

GetAcknowledge returns the Acknowledge field if non-nil, zero value otherwise.

### GetAcknowledgeOk

`func (o *CondAlarm) GetAcknowledgeOk() (*string, bool)`

GetAcknowledgeOk returns a tuple with the Acknowledge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledge

`func (o *CondAlarm) SetAcknowledge(v string)`

SetAcknowledge sets Acknowledge field to given value.

### HasAcknowledge

`func (o *CondAlarm) HasAcknowledge() bool`

HasAcknowledge returns a boolean if a field has been set.

### GetAcknowledgeBy

`func (o *CondAlarm) GetAcknowledgeBy() string`

GetAcknowledgeBy returns the AcknowledgeBy field if non-nil, zero value otherwise.

### GetAcknowledgeByOk

`func (o *CondAlarm) GetAcknowledgeByOk() (*string, bool)`

GetAcknowledgeByOk returns a tuple with the AcknowledgeBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgeBy

`func (o *CondAlarm) SetAcknowledgeBy(v string)`

SetAcknowledgeBy sets AcknowledgeBy field to given value.

### HasAcknowledgeBy

`func (o *CondAlarm) HasAcknowledgeBy() bool`

HasAcknowledgeBy returns a boolean if a field has been set.

### GetAcknowledgeTime

`func (o *CondAlarm) GetAcknowledgeTime() time.Time`

GetAcknowledgeTime returns the AcknowledgeTime field if non-nil, zero value otherwise.

### GetAcknowledgeTimeOk

`func (o *CondAlarm) GetAcknowledgeTimeOk() (*time.Time, bool)`

GetAcknowledgeTimeOk returns a tuple with the AcknowledgeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgeTime

`func (o *CondAlarm) SetAcknowledgeTime(v time.Time)`

SetAcknowledgeTime sets AcknowledgeTime field to given value.

### HasAcknowledgeTime

`func (o *CondAlarm) HasAcknowledgeTime() bool`

HasAcknowledgeTime returns a boolean if a field has been set.

### GetAffectedMoDisplayName

`func (o *CondAlarm) GetAffectedMoDisplayName() string`

GetAffectedMoDisplayName returns the AffectedMoDisplayName field if non-nil, zero value otherwise.

### GetAffectedMoDisplayNameOk

`func (o *CondAlarm) GetAffectedMoDisplayNameOk() (*string, bool)`

GetAffectedMoDisplayNameOk returns a tuple with the AffectedMoDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedMoDisplayName

`func (o *CondAlarm) SetAffectedMoDisplayName(v string)`

SetAffectedMoDisplayName sets AffectedMoDisplayName field to given value.

### HasAffectedMoDisplayName

`func (o *CondAlarm) HasAffectedMoDisplayName() bool`

HasAffectedMoDisplayName returns a boolean if a field has been set.

### GetAffectedMoId

`func (o *CondAlarm) GetAffectedMoId() string`

GetAffectedMoId returns the AffectedMoId field if non-nil, zero value otherwise.

### GetAffectedMoIdOk

`func (o *CondAlarm) GetAffectedMoIdOk() (*string, bool)`

GetAffectedMoIdOk returns a tuple with the AffectedMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedMoId

`func (o *CondAlarm) SetAffectedMoId(v string)`

SetAffectedMoId sets AffectedMoId field to given value.

### HasAffectedMoId

`func (o *CondAlarm) HasAffectedMoId() bool`

HasAffectedMoId returns a boolean if a field has been set.

### GetAffectedMoType

`func (o *CondAlarm) GetAffectedMoType() string`

GetAffectedMoType returns the AffectedMoType field if non-nil, zero value otherwise.

### GetAffectedMoTypeOk

`func (o *CondAlarm) GetAffectedMoTypeOk() (*string, bool)`

GetAffectedMoTypeOk returns a tuple with the AffectedMoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedMoType

`func (o *CondAlarm) SetAffectedMoType(v string)`

SetAffectedMoType sets AffectedMoType field to given value.

### HasAffectedMoType

`func (o *CondAlarm) HasAffectedMoType() bool`

HasAffectedMoType returns a boolean if a field has been set.

### GetAffectedObject

`func (o *CondAlarm) GetAffectedObject() string`

GetAffectedObject returns the AffectedObject field if non-nil, zero value otherwise.

### GetAffectedObjectOk

`func (o *CondAlarm) GetAffectedObjectOk() (*string, bool)`

GetAffectedObjectOk returns a tuple with the AffectedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedObject

`func (o *CondAlarm) SetAffectedObject(v string)`

SetAffectedObject sets AffectedObject field to given value.

### HasAffectedObject

`func (o *CondAlarm) HasAffectedObject() bool`

HasAffectedObject returns a boolean if a field has been set.

### GetAncestorMoId

`func (o *CondAlarm) GetAncestorMoId() string`

GetAncestorMoId returns the AncestorMoId field if non-nil, zero value otherwise.

### GetAncestorMoIdOk

`func (o *CondAlarm) GetAncestorMoIdOk() (*string, bool)`

GetAncestorMoIdOk returns a tuple with the AncestorMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestorMoId

`func (o *CondAlarm) SetAncestorMoId(v string)`

SetAncestorMoId sets AncestorMoId field to given value.

### HasAncestorMoId

`func (o *CondAlarm) HasAncestorMoId() bool`

HasAncestorMoId returns a boolean if a field has been set.

### GetAncestorMoType

`func (o *CondAlarm) GetAncestorMoType() string`

GetAncestorMoType returns the AncestorMoType field if non-nil, zero value otherwise.

### GetAncestorMoTypeOk

`func (o *CondAlarm) GetAncestorMoTypeOk() (*string, bool)`

GetAncestorMoTypeOk returns a tuple with the AncestorMoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestorMoType

`func (o *CondAlarm) SetAncestorMoType(v string)`

SetAncestorMoType sets AncestorMoType field to given value.

### HasAncestorMoType

`func (o *CondAlarm) HasAncestorMoType() bool`

HasAncestorMoType returns a boolean if a field has been set.

### GetCode

`func (o *CondAlarm) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CondAlarm) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CondAlarm) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CondAlarm) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCreationTime

`func (o *CondAlarm) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *CondAlarm) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *CondAlarm) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *CondAlarm) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetDescription

`func (o *CondAlarm) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CondAlarm) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CondAlarm) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CondAlarm) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFlapping

`func (o *CondAlarm) GetFlapping() string`

GetFlapping returns the Flapping field if non-nil, zero value otherwise.

### GetFlappingOk

`func (o *CondAlarm) GetFlappingOk() (*string, bool)`

GetFlappingOk returns a tuple with the Flapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlapping

`func (o *CondAlarm) SetFlapping(v string)`

SetFlapping sets Flapping field to given value.

### HasFlapping

`func (o *CondAlarm) HasFlapping() bool`

HasFlapping returns a boolean if a field has been set.

### GetFlappingCount

`func (o *CondAlarm) GetFlappingCount() int64`

GetFlappingCount returns the FlappingCount field if non-nil, zero value otherwise.

### GetFlappingCountOk

`func (o *CondAlarm) GetFlappingCountOk() (*int64, bool)`

GetFlappingCountOk returns a tuple with the FlappingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlappingCount

`func (o *CondAlarm) SetFlappingCount(v int64)`

SetFlappingCount sets FlappingCount field to given value.

### HasFlappingCount

`func (o *CondAlarm) HasFlappingCount() bool`

HasFlappingCount returns a boolean if a field has been set.

### GetFlappingStartTime

`func (o *CondAlarm) GetFlappingStartTime() time.Time`

GetFlappingStartTime returns the FlappingStartTime field if non-nil, zero value otherwise.

### GetFlappingStartTimeOk

`func (o *CondAlarm) GetFlappingStartTimeOk() (*time.Time, bool)`

GetFlappingStartTimeOk returns a tuple with the FlappingStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlappingStartTime

`func (o *CondAlarm) SetFlappingStartTime(v time.Time)`

SetFlappingStartTime sets FlappingStartTime field to given value.

### HasFlappingStartTime

`func (o *CondAlarm) HasFlappingStartTime() bool`

HasFlappingStartTime returns a boolean if a field has been set.

### GetLastTransitionTime

`func (o *CondAlarm) GetLastTransitionTime() time.Time`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *CondAlarm) GetLastTransitionTimeOk() (*time.Time, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *CondAlarm) SetLastTransitionTime(v time.Time)`

SetLastTransitionTime sets LastTransitionTime field to given value.

### HasLastTransitionTime

`func (o *CondAlarm) HasLastTransitionTime() bool`

HasLastTransitionTime returns a boolean if a field has been set.

### GetMsAffectedObject

`func (o *CondAlarm) GetMsAffectedObject() string`

GetMsAffectedObject returns the MsAffectedObject field if non-nil, zero value otherwise.

### GetMsAffectedObjectOk

`func (o *CondAlarm) GetMsAffectedObjectOk() (*string, bool)`

GetMsAffectedObjectOk returns a tuple with the MsAffectedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAffectedObject

`func (o *CondAlarm) SetMsAffectedObject(v string)`

SetMsAffectedObject sets MsAffectedObject field to given value.

### HasMsAffectedObject

`func (o *CondAlarm) HasMsAffectedObject() bool`

HasMsAffectedObject returns a boolean if a field has been set.

### GetName

`func (o *CondAlarm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CondAlarm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CondAlarm) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CondAlarm) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrigSeverity

`func (o *CondAlarm) GetOrigSeverity() string`

GetOrigSeverity returns the OrigSeverity field if non-nil, zero value otherwise.

### GetOrigSeverityOk

`func (o *CondAlarm) GetOrigSeverityOk() (*string, bool)`

GetOrigSeverityOk returns a tuple with the OrigSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigSeverity

`func (o *CondAlarm) SetOrigSeverity(v string)`

SetOrigSeverity sets OrigSeverity field to given value.

### HasOrigSeverity

`func (o *CondAlarm) HasOrigSeverity() bool`

HasOrigSeverity returns a boolean if a field has been set.

### GetSeverity

`func (o *CondAlarm) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *CondAlarm) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *CondAlarm) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *CondAlarm) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSuppressed

`func (o *CondAlarm) GetSuppressed() bool`

GetSuppressed returns the Suppressed field if non-nil, zero value otherwise.

### GetSuppressedOk

`func (o *CondAlarm) GetSuppressedOk() (*bool, bool)`

GetSuppressedOk returns a tuple with the Suppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressed

`func (o *CondAlarm) SetSuppressed(v bool)`

SetSuppressed sets Suppressed field to given value.

### HasSuppressed

`func (o *CondAlarm) HasSuppressed() bool`

HasSuppressed returns a boolean if a field has been set.

### GetAffectedMo

`func (o *CondAlarm) GetAffectedMo() MoBaseMoRelationship`

GetAffectedMo returns the AffectedMo field if non-nil, zero value otherwise.

### GetAffectedMoOk

`func (o *CondAlarm) GetAffectedMoOk() (*MoBaseMoRelationship, bool)`

GetAffectedMoOk returns a tuple with the AffectedMo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedMo

`func (o *CondAlarm) SetAffectedMo(v MoBaseMoRelationship)`

SetAffectedMo sets AffectedMo field to given value.

### HasAffectedMo

`func (o *CondAlarm) HasAffectedMo() bool`

HasAffectedMo returns a boolean if a field has been set.

### SetAffectedMoNil

`func (o *CondAlarm) SetAffectedMoNil(b bool)`

 SetAffectedMoNil sets the value for AffectedMo to be an explicit nil

### UnsetAffectedMo
`func (o *CondAlarm) UnsetAffectedMo()`

UnsetAffectedMo ensures that no value is present for AffectedMo, not even an explicit nil
### GetAlarmSummaryAggregators

`func (o *CondAlarm) GetAlarmSummaryAggregators() []MoBaseMoRelationship`

GetAlarmSummaryAggregators returns the AlarmSummaryAggregators field if non-nil, zero value otherwise.

### GetAlarmSummaryAggregatorsOk

`func (o *CondAlarm) GetAlarmSummaryAggregatorsOk() (*[]MoBaseMoRelationship, bool)`

GetAlarmSummaryAggregatorsOk returns a tuple with the AlarmSummaryAggregators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmSummaryAggregators

`func (o *CondAlarm) SetAlarmSummaryAggregators(v []MoBaseMoRelationship)`

SetAlarmSummaryAggregators sets AlarmSummaryAggregators field to given value.

### HasAlarmSummaryAggregators

`func (o *CondAlarm) HasAlarmSummaryAggregators() bool`

HasAlarmSummaryAggregators returns a boolean if a field has been set.

### SetAlarmSummaryAggregatorsNil

`func (o *CondAlarm) SetAlarmSummaryAggregatorsNil(b bool)`

 SetAlarmSummaryAggregatorsNil sets the value for AlarmSummaryAggregators to be an explicit nil

### UnsetAlarmSummaryAggregators
`func (o *CondAlarm) UnsetAlarmSummaryAggregators()`

UnsetAlarmSummaryAggregators ensures that no value is present for AlarmSummaryAggregators, not even an explicit nil
### GetDefinition

`func (o *CondAlarm) GetDefinition() CondAlarmDefinitionRelationship`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *CondAlarm) GetDefinitionOk() (*CondAlarmDefinitionRelationship, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *CondAlarm) SetDefinition(v CondAlarmDefinitionRelationship)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *CondAlarm) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### SetDefinitionNil

`func (o *CondAlarm) SetDefinitionNil(b bool)`

 SetDefinitionNil sets the value for Definition to be an explicit nil

### UnsetDefinition
`func (o *CondAlarm) UnsetDefinition()`

UnsetDefinition ensures that no value is present for Definition, not even an explicit nil
### GetRegisteredDevice

`func (o *CondAlarm) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *CondAlarm) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *CondAlarm) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *CondAlarm) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *CondAlarm) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *CondAlarm) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


