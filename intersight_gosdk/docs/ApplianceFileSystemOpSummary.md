# ApplianceFileSystemOpSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.FileSystemOpSummary"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.FileSystemOpSummary"]
**AlarmSummary** | Pointer to [**NullableCondAlarmSummary**](CondAlarmSummary.md) |  | [optional] 
**NodeOpStatus** | Pointer to [**NullableApplianceNodeOpStatusRelationship**](ApplianceNodeOpStatusRelationship.md) |  | [optional] 

## Methods

### NewApplianceFileSystemOpSummary

`func NewApplianceFileSystemOpSummary(classId string, objectType string, ) *ApplianceFileSystemOpSummary`

NewApplianceFileSystemOpSummary instantiates a new ApplianceFileSystemOpSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceFileSystemOpSummaryWithDefaults

`func NewApplianceFileSystemOpSummaryWithDefaults() *ApplianceFileSystemOpSummary`

NewApplianceFileSystemOpSummaryWithDefaults instantiates a new ApplianceFileSystemOpSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceFileSystemOpSummary) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceFileSystemOpSummary) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceFileSystemOpSummary) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceFileSystemOpSummary) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceFileSystemOpSummary) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceFileSystemOpSummary) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAlarmSummary

`func (o *ApplianceFileSystemOpSummary) GetAlarmSummary() CondAlarmSummary`

GetAlarmSummary returns the AlarmSummary field if non-nil, zero value otherwise.

### GetAlarmSummaryOk

`func (o *ApplianceFileSystemOpSummary) GetAlarmSummaryOk() (*CondAlarmSummary, bool)`

GetAlarmSummaryOk returns a tuple with the AlarmSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmSummary

`func (o *ApplianceFileSystemOpSummary) SetAlarmSummary(v CondAlarmSummary)`

SetAlarmSummary sets AlarmSummary field to given value.

### HasAlarmSummary

`func (o *ApplianceFileSystemOpSummary) HasAlarmSummary() bool`

HasAlarmSummary returns a boolean if a field has been set.

### SetAlarmSummaryNil

`func (o *ApplianceFileSystemOpSummary) SetAlarmSummaryNil(b bool)`

 SetAlarmSummaryNil sets the value for AlarmSummary to be an explicit nil

### UnsetAlarmSummary
`func (o *ApplianceFileSystemOpSummary) UnsetAlarmSummary()`

UnsetAlarmSummary ensures that no value is present for AlarmSummary, not even an explicit nil
### GetNodeOpStatus

`func (o *ApplianceFileSystemOpSummary) GetNodeOpStatus() ApplianceNodeOpStatusRelationship`

GetNodeOpStatus returns the NodeOpStatus field if non-nil, zero value otherwise.

### GetNodeOpStatusOk

`func (o *ApplianceFileSystemOpSummary) GetNodeOpStatusOk() (*ApplianceNodeOpStatusRelationship, bool)`

GetNodeOpStatusOk returns a tuple with the NodeOpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeOpStatus

`func (o *ApplianceFileSystemOpSummary) SetNodeOpStatus(v ApplianceNodeOpStatusRelationship)`

SetNodeOpStatus sets NodeOpStatus field to given value.

### HasNodeOpStatus

`func (o *ApplianceFileSystemOpSummary) HasNodeOpStatus() bool`

HasNodeOpStatus returns a boolean if a field has been set.

### SetNodeOpStatusNil

`func (o *ApplianceFileSystemOpSummary) SetNodeOpStatusNil(b bool)`

 SetNodeOpStatusNil sets the value for NodeOpStatus to be an explicit nil

### UnsetNodeOpStatus
`func (o *ApplianceFileSystemOpSummary) UnsetNodeOpStatus()`

UnsetNodeOpStatus ensures that no value is present for NodeOpStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


