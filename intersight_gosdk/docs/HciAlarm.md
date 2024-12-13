# HciAlarm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.Alarm"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.Alarm"]
**AlarmExtId** | Pointer to **string** | The unique identifier for the alarm on the endpoint. | [optional] [readonly] 
**AlertType** | Pointer to **string** | The code for the reported alarm. | [optional] [readonly] 
**ClusterExtId** | Pointer to **string** | The unique identifer for the cluster associated with the source entity on the endpoint. | [optional] [readonly] 
**IsResolved** | Pointer to **bool** | The status of the alarm. If an alarm is resolved, this value is set as true on the endpoint. | [optional] [readonly] 
**Message** | Pointer to **string** | The description from the endpoint explaining the cause of the alarm. | [optional] [readonly] 
**OriginAcknowledgeTime** | Pointer to **time.Time** | The time the alarm was acknowledged on the endpoint. | [optional] [readonly] 
**OriginCreationTime** | Pointer to **time.Time** | The time the alarm was created on the endpoint. | [optional] [readonly] 
**OriginIsAcknowledged** | Pointer to **bool** | The acknowledgement status of the alert on the endpoint. | [optional] [readonly] 
**Parameters** | Pointer to [**[]HciAlarmParameter**](HciAlarmParameter.md) |  | [optional] 
**ResolvedTime** | Pointer to **time.Time** | The time the alarm was resolved on the endpoint. | [optional] [readonly] 
**Severity** | Pointer to **string** | The severity of the alarm. Valid values are Critical, Warning, Info. * &#x60;None&#x60; - The Enum value None represents that there is no severity. * &#x60;Info&#x60; - The Enum value Info represents the Informational level of severity. * &#x60;Critical&#x60; - The Enum value Critical represents the Critical level of severity. * &#x60;Warning&#x60; - The Enum value Warning represents the Warning level of severity. * &#x60;Cleared&#x60; - The Enum value Cleared represents that the alarm severity has been cleared. | [optional] [readonly] [default to "None"]
**SourceEntityExtId** | Pointer to **string** | The unique identifer for the entity on the endpoint for which the alarm was raised. | [optional] [readonly] 
**SourceEntityName** | Pointer to **string** | The name of the entity on the endpoint for which the alarm was raised. | [optional] [readonly] 
**SourceEntityType** | Pointer to **string** | The object type for the entity corresponding to the objects inventoried. | [optional] [readonly] 
**Title** | Pointer to **string** | The title of the reported alarm. | [optional] [readonly] 
**Cluster** | Pointer to [**NullableHciClusterRelationship**](HciClusterRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewHciAlarm

`func NewHciAlarm(classId string, objectType string, ) *HciAlarm`

NewHciAlarm instantiates a new HciAlarm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciAlarmWithDefaults

`func NewHciAlarmWithDefaults() *HciAlarm`

NewHciAlarmWithDefaults instantiates a new HciAlarm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciAlarm) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciAlarm) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciAlarm) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciAlarm) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciAlarm) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciAlarm) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAlarmExtId

`func (o *HciAlarm) GetAlarmExtId() string`

GetAlarmExtId returns the AlarmExtId field if non-nil, zero value otherwise.

### GetAlarmExtIdOk

`func (o *HciAlarm) GetAlarmExtIdOk() (*string, bool)`

GetAlarmExtIdOk returns a tuple with the AlarmExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmExtId

`func (o *HciAlarm) SetAlarmExtId(v string)`

SetAlarmExtId sets AlarmExtId field to given value.

### HasAlarmExtId

`func (o *HciAlarm) HasAlarmExtId() bool`

HasAlarmExtId returns a boolean if a field has been set.

### GetAlertType

`func (o *HciAlarm) GetAlertType() string`

GetAlertType returns the AlertType field if non-nil, zero value otherwise.

### GetAlertTypeOk

`func (o *HciAlarm) GetAlertTypeOk() (*string, bool)`

GetAlertTypeOk returns a tuple with the AlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertType

`func (o *HciAlarm) SetAlertType(v string)`

SetAlertType sets AlertType field to given value.

### HasAlertType

`func (o *HciAlarm) HasAlertType() bool`

HasAlertType returns a boolean if a field has been set.

### GetClusterExtId

`func (o *HciAlarm) GetClusterExtId() string`

GetClusterExtId returns the ClusterExtId field if non-nil, zero value otherwise.

### GetClusterExtIdOk

`func (o *HciAlarm) GetClusterExtIdOk() (*string, bool)`

GetClusterExtIdOk returns a tuple with the ClusterExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterExtId

`func (o *HciAlarm) SetClusterExtId(v string)`

SetClusterExtId sets ClusterExtId field to given value.

### HasClusterExtId

`func (o *HciAlarm) HasClusterExtId() bool`

HasClusterExtId returns a boolean if a field has been set.

### GetIsResolved

`func (o *HciAlarm) GetIsResolved() bool`

GetIsResolved returns the IsResolved field if non-nil, zero value otherwise.

### GetIsResolvedOk

`func (o *HciAlarm) GetIsResolvedOk() (*bool, bool)`

GetIsResolvedOk returns a tuple with the IsResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsResolved

`func (o *HciAlarm) SetIsResolved(v bool)`

SetIsResolved sets IsResolved field to given value.

### HasIsResolved

`func (o *HciAlarm) HasIsResolved() bool`

HasIsResolved returns a boolean if a field has been set.

### GetMessage

`func (o *HciAlarm) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *HciAlarm) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *HciAlarm) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *HciAlarm) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOriginAcknowledgeTime

`func (o *HciAlarm) GetOriginAcknowledgeTime() time.Time`

GetOriginAcknowledgeTime returns the OriginAcknowledgeTime field if non-nil, zero value otherwise.

### GetOriginAcknowledgeTimeOk

`func (o *HciAlarm) GetOriginAcknowledgeTimeOk() (*time.Time, bool)`

GetOriginAcknowledgeTimeOk returns a tuple with the OriginAcknowledgeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginAcknowledgeTime

`func (o *HciAlarm) SetOriginAcknowledgeTime(v time.Time)`

SetOriginAcknowledgeTime sets OriginAcknowledgeTime field to given value.

### HasOriginAcknowledgeTime

`func (o *HciAlarm) HasOriginAcknowledgeTime() bool`

HasOriginAcknowledgeTime returns a boolean if a field has been set.

### GetOriginCreationTime

`func (o *HciAlarm) GetOriginCreationTime() time.Time`

GetOriginCreationTime returns the OriginCreationTime field if non-nil, zero value otherwise.

### GetOriginCreationTimeOk

`func (o *HciAlarm) GetOriginCreationTimeOk() (*time.Time, bool)`

GetOriginCreationTimeOk returns a tuple with the OriginCreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginCreationTime

`func (o *HciAlarm) SetOriginCreationTime(v time.Time)`

SetOriginCreationTime sets OriginCreationTime field to given value.

### HasOriginCreationTime

`func (o *HciAlarm) HasOriginCreationTime() bool`

HasOriginCreationTime returns a boolean if a field has been set.

### GetOriginIsAcknowledged

`func (o *HciAlarm) GetOriginIsAcknowledged() bool`

GetOriginIsAcknowledged returns the OriginIsAcknowledged field if non-nil, zero value otherwise.

### GetOriginIsAcknowledgedOk

`func (o *HciAlarm) GetOriginIsAcknowledgedOk() (*bool, bool)`

GetOriginIsAcknowledgedOk returns a tuple with the OriginIsAcknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginIsAcknowledged

`func (o *HciAlarm) SetOriginIsAcknowledged(v bool)`

SetOriginIsAcknowledged sets OriginIsAcknowledged field to given value.

### HasOriginIsAcknowledged

`func (o *HciAlarm) HasOriginIsAcknowledged() bool`

HasOriginIsAcknowledged returns a boolean if a field has been set.

### GetParameters

`func (o *HciAlarm) GetParameters() []HciAlarmParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *HciAlarm) GetParametersOk() (*[]HciAlarmParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *HciAlarm) SetParameters(v []HciAlarmParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *HciAlarm) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *HciAlarm) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *HciAlarm) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil
### GetResolvedTime

`func (o *HciAlarm) GetResolvedTime() time.Time`

GetResolvedTime returns the ResolvedTime field if non-nil, zero value otherwise.

### GetResolvedTimeOk

`func (o *HciAlarm) GetResolvedTimeOk() (*time.Time, bool)`

GetResolvedTimeOk returns a tuple with the ResolvedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedTime

`func (o *HciAlarm) SetResolvedTime(v time.Time)`

SetResolvedTime sets ResolvedTime field to given value.

### HasResolvedTime

`func (o *HciAlarm) HasResolvedTime() bool`

HasResolvedTime returns a boolean if a field has been set.

### GetSeverity

`func (o *HciAlarm) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *HciAlarm) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *HciAlarm) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *HciAlarm) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSourceEntityExtId

`func (o *HciAlarm) GetSourceEntityExtId() string`

GetSourceEntityExtId returns the SourceEntityExtId field if non-nil, zero value otherwise.

### GetSourceEntityExtIdOk

`func (o *HciAlarm) GetSourceEntityExtIdOk() (*string, bool)`

GetSourceEntityExtIdOk returns a tuple with the SourceEntityExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEntityExtId

`func (o *HciAlarm) SetSourceEntityExtId(v string)`

SetSourceEntityExtId sets SourceEntityExtId field to given value.

### HasSourceEntityExtId

`func (o *HciAlarm) HasSourceEntityExtId() bool`

HasSourceEntityExtId returns a boolean if a field has been set.

### GetSourceEntityName

`func (o *HciAlarm) GetSourceEntityName() string`

GetSourceEntityName returns the SourceEntityName field if non-nil, zero value otherwise.

### GetSourceEntityNameOk

`func (o *HciAlarm) GetSourceEntityNameOk() (*string, bool)`

GetSourceEntityNameOk returns a tuple with the SourceEntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEntityName

`func (o *HciAlarm) SetSourceEntityName(v string)`

SetSourceEntityName sets SourceEntityName field to given value.

### HasSourceEntityName

`func (o *HciAlarm) HasSourceEntityName() bool`

HasSourceEntityName returns a boolean if a field has been set.

### GetSourceEntityType

`func (o *HciAlarm) GetSourceEntityType() string`

GetSourceEntityType returns the SourceEntityType field if non-nil, zero value otherwise.

### GetSourceEntityTypeOk

`func (o *HciAlarm) GetSourceEntityTypeOk() (*string, bool)`

GetSourceEntityTypeOk returns a tuple with the SourceEntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEntityType

`func (o *HciAlarm) SetSourceEntityType(v string)`

SetSourceEntityType sets SourceEntityType field to given value.

### HasSourceEntityType

`func (o *HciAlarm) HasSourceEntityType() bool`

HasSourceEntityType returns a boolean if a field has been set.

### GetTitle

`func (o *HciAlarm) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *HciAlarm) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *HciAlarm) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *HciAlarm) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetCluster

`func (o *HciAlarm) GetCluster() HciClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HciAlarm) GetClusterOk() (*HciClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HciAlarm) SetCluster(v HciClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HciAlarm) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *HciAlarm) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *HciAlarm) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetRegisteredDevice

`func (o *HciAlarm) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *HciAlarm) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *HciAlarm) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *HciAlarm) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *HciAlarm) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *HciAlarm) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


