# ConvergedinfraBasePodSummaryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "convergedinfra.PodSummary"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "convergedinfra.PodSummary"]
**AlarmSummary** | Pointer to [**NullableConvergedinfraAlarmSummary**](ConvergedinfraAlarmSummary.md) |  | [optional] 
**ComplianceSummary** | Pointer to [**NullableConvergedinfraComplianceSummary**](ConvergedinfraComplianceSummary.md) |  | [optional] 
**NodeCount** | Pointer to **int64** | Number of nodes associated with the pod. | [optional] [readonly] 
**StorageAvailable** | Pointer to **int64** | The available storage capacity for this pod. | [optional] [readonly] 
**StorageCapacity** | Pointer to **int64** | The total storage capacity for this pod. | [optional] [readonly] 
**StorageUtilization** | Pointer to **float32** | The percentage storage utilization for this pod. | [optional] [readonly] 

## Methods

### NewConvergedinfraBasePodSummaryAllOf

`func NewConvergedinfraBasePodSummaryAllOf(classId string, objectType string, ) *ConvergedinfraBasePodSummaryAllOf`

NewConvergedinfraBasePodSummaryAllOf instantiates a new ConvergedinfraBasePodSummaryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvergedinfraBasePodSummaryAllOfWithDefaults

`func NewConvergedinfraBasePodSummaryAllOfWithDefaults() *ConvergedinfraBasePodSummaryAllOf`

NewConvergedinfraBasePodSummaryAllOfWithDefaults instantiates a new ConvergedinfraBasePodSummaryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConvergedinfraBasePodSummaryAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConvergedinfraBasePodSummaryAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConvergedinfraBasePodSummaryAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConvergedinfraBasePodSummaryAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConvergedinfraBasePodSummaryAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConvergedinfraBasePodSummaryAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAlarmSummary

`func (o *ConvergedinfraBasePodSummaryAllOf) GetAlarmSummary() ConvergedinfraAlarmSummary`

GetAlarmSummary returns the AlarmSummary field if non-nil, zero value otherwise.

### GetAlarmSummaryOk

`func (o *ConvergedinfraBasePodSummaryAllOf) GetAlarmSummaryOk() (*ConvergedinfraAlarmSummary, bool)`

GetAlarmSummaryOk returns a tuple with the AlarmSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmSummary

`func (o *ConvergedinfraBasePodSummaryAllOf) SetAlarmSummary(v ConvergedinfraAlarmSummary)`

SetAlarmSummary sets AlarmSummary field to given value.

### HasAlarmSummary

`func (o *ConvergedinfraBasePodSummaryAllOf) HasAlarmSummary() bool`

HasAlarmSummary returns a boolean if a field has been set.

### SetAlarmSummaryNil

`func (o *ConvergedinfraBasePodSummaryAllOf) SetAlarmSummaryNil(b bool)`

 SetAlarmSummaryNil sets the value for AlarmSummary to be an explicit nil

### UnsetAlarmSummary
`func (o *ConvergedinfraBasePodSummaryAllOf) UnsetAlarmSummary()`

UnsetAlarmSummary ensures that no value is present for AlarmSummary, not even an explicit nil
### GetComplianceSummary

`func (o *ConvergedinfraBasePodSummaryAllOf) GetComplianceSummary() ConvergedinfraComplianceSummary`

GetComplianceSummary returns the ComplianceSummary field if non-nil, zero value otherwise.

### GetComplianceSummaryOk

`func (o *ConvergedinfraBasePodSummaryAllOf) GetComplianceSummaryOk() (*ConvergedinfraComplianceSummary, bool)`

GetComplianceSummaryOk returns a tuple with the ComplianceSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceSummary

`func (o *ConvergedinfraBasePodSummaryAllOf) SetComplianceSummary(v ConvergedinfraComplianceSummary)`

SetComplianceSummary sets ComplianceSummary field to given value.

### HasComplianceSummary

`func (o *ConvergedinfraBasePodSummaryAllOf) HasComplianceSummary() bool`

HasComplianceSummary returns a boolean if a field has been set.

### SetComplianceSummaryNil

`func (o *ConvergedinfraBasePodSummaryAllOf) SetComplianceSummaryNil(b bool)`

 SetComplianceSummaryNil sets the value for ComplianceSummary to be an explicit nil

### UnsetComplianceSummary
`func (o *ConvergedinfraBasePodSummaryAllOf) UnsetComplianceSummary()`

UnsetComplianceSummary ensures that no value is present for ComplianceSummary, not even an explicit nil
### GetNodeCount

`func (o *ConvergedinfraBasePodSummaryAllOf) GetNodeCount() int64`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *ConvergedinfraBasePodSummaryAllOf) GetNodeCountOk() (*int64, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *ConvergedinfraBasePodSummaryAllOf) SetNodeCount(v int64)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *ConvergedinfraBasePodSummaryAllOf) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetStorageAvailable

`func (o *ConvergedinfraBasePodSummaryAllOf) GetStorageAvailable() int64`

GetStorageAvailable returns the StorageAvailable field if non-nil, zero value otherwise.

### GetStorageAvailableOk

`func (o *ConvergedinfraBasePodSummaryAllOf) GetStorageAvailableOk() (*int64, bool)`

GetStorageAvailableOk returns a tuple with the StorageAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAvailable

`func (o *ConvergedinfraBasePodSummaryAllOf) SetStorageAvailable(v int64)`

SetStorageAvailable sets StorageAvailable field to given value.

### HasStorageAvailable

`func (o *ConvergedinfraBasePodSummaryAllOf) HasStorageAvailable() bool`

HasStorageAvailable returns a boolean if a field has been set.

### GetStorageCapacity

`func (o *ConvergedinfraBasePodSummaryAllOf) GetStorageCapacity() int64`

GetStorageCapacity returns the StorageCapacity field if non-nil, zero value otherwise.

### GetStorageCapacityOk

`func (o *ConvergedinfraBasePodSummaryAllOf) GetStorageCapacityOk() (*int64, bool)`

GetStorageCapacityOk returns a tuple with the StorageCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageCapacity

`func (o *ConvergedinfraBasePodSummaryAllOf) SetStorageCapacity(v int64)`

SetStorageCapacity sets StorageCapacity field to given value.

### HasStorageCapacity

`func (o *ConvergedinfraBasePodSummaryAllOf) HasStorageCapacity() bool`

HasStorageCapacity returns a boolean if a field has been set.

### GetStorageUtilization

`func (o *ConvergedinfraBasePodSummaryAllOf) GetStorageUtilization() float32`

GetStorageUtilization returns the StorageUtilization field if non-nil, zero value otherwise.

### GetStorageUtilizationOk

`func (o *ConvergedinfraBasePodSummaryAllOf) GetStorageUtilizationOk() (*float32, bool)`

GetStorageUtilizationOk returns a tuple with the StorageUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUtilization

`func (o *ConvergedinfraBasePodSummaryAllOf) SetStorageUtilization(v float32)`

SetStorageUtilization sets StorageUtilization field to given value.

### HasStorageUtilization

`func (o *ConvergedinfraBasePodSummaryAllOf) HasStorageUtilization() bool`

HasStorageUtilization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


