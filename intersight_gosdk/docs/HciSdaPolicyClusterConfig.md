# HciSdaPolicyClusterConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.SdaPolicyClusterConfig"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.SdaPolicyClusterConfig"]
**ClusterExtId** | Pointer to **string** | Unique ID of the cluster. | [optional] [readonly] 
**ConfigurableParameters** | Pointer to [**[]HciNccConfigurableParameter**](HciNccConfigurableParameter.md) |  | [optional] 
**IsEnabled** | Pointer to **bool** | Indicates whether the NCC policy is enabled or not on the cluster. | [optional] [readonly] 
**LastModifiedTime** | Pointer to **string** | Time in ISO 8601 format when the NCC policy was last modified. | [optional] [readonly] 
**ScheduleIntervalSeconds** | Pointer to **int32** | Interval in seconds for periodically executing the NCC policy. | [optional] [readonly] 

## Methods

### NewHciSdaPolicyClusterConfig

`func NewHciSdaPolicyClusterConfig(classId string, objectType string, ) *HciSdaPolicyClusterConfig`

NewHciSdaPolicyClusterConfig instantiates a new HciSdaPolicyClusterConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciSdaPolicyClusterConfigWithDefaults

`func NewHciSdaPolicyClusterConfigWithDefaults() *HciSdaPolicyClusterConfig`

NewHciSdaPolicyClusterConfigWithDefaults instantiates a new HciSdaPolicyClusterConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciSdaPolicyClusterConfig) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciSdaPolicyClusterConfig) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciSdaPolicyClusterConfig) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciSdaPolicyClusterConfig) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciSdaPolicyClusterConfig) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciSdaPolicyClusterConfig) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClusterExtId

`func (o *HciSdaPolicyClusterConfig) GetClusterExtId() string`

GetClusterExtId returns the ClusterExtId field if non-nil, zero value otherwise.

### GetClusterExtIdOk

`func (o *HciSdaPolicyClusterConfig) GetClusterExtIdOk() (*string, bool)`

GetClusterExtIdOk returns a tuple with the ClusterExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterExtId

`func (o *HciSdaPolicyClusterConfig) SetClusterExtId(v string)`

SetClusterExtId sets ClusterExtId field to given value.

### HasClusterExtId

`func (o *HciSdaPolicyClusterConfig) HasClusterExtId() bool`

HasClusterExtId returns a boolean if a field has been set.

### GetConfigurableParameters

`func (o *HciSdaPolicyClusterConfig) GetConfigurableParameters() []HciNccConfigurableParameter`

GetConfigurableParameters returns the ConfigurableParameters field if non-nil, zero value otherwise.

### GetConfigurableParametersOk

`func (o *HciSdaPolicyClusterConfig) GetConfigurableParametersOk() (*[]HciNccConfigurableParameter, bool)`

GetConfigurableParametersOk returns a tuple with the ConfigurableParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurableParameters

`func (o *HciSdaPolicyClusterConfig) SetConfigurableParameters(v []HciNccConfigurableParameter)`

SetConfigurableParameters sets ConfigurableParameters field to given value.

### HasConfigurableParameters

`func (o *HciSdaPolicyClusterConfig) HasConfigurableParameters() bool`

HasConfigurableParameters returns a boolean if a field has been set.

### SetConfigurableParametersNil

`func (o *HciSdaPolicyClusterConfig) SetConfigurableParametersNil(b bool)`

 SetConfigurableParametersNil sets the value for ConfigurableParameters to be an explicit nil

### UnsetConfigurableParameters
`func (o *HciSdaPolicyClusterConfig) UnsetConfigurableParameters()`

UnsetConfigurableParameters ensures that no value is present for ConfigurableParameters, not even an explicit nil
### GetIsEnabled

`func (o *HciSdaPolicyClusterConfig) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *HciSdaPolicyClusterConfig) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *HciSdaPolicyClusterConfig) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *HciSdaPolicyClusterConfig) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *HciSdaPolicyClusterConfig) GetLastModifiedTime() string`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *HciSdaPolicyClusterConfig) GetLastModifiedTimeOk() (*string, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *HciSdaPolicyClusterConfig) SetLastModifiedTime(v string)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *HciSdaPolicyClusterConfig) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetScheduleIntervalSeconds

`func (o *HciSdaPolicyClusterConfig) GetScheduleIntervalSeconds() int32`

GetScheduleIntervalSeconds returns the ScheduleIntervalSeconds field if non-nil, zero value otherwise.

### GetScheduleIntervalSecondsOk

`func (o *HciSdaPolicyClusterConfig) GetScheduleIntervalSecondsOk() (*int32, bool)`

GetScheduleIntervalSecondsOk returns a tuple with the ScheduleIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleIntervalSeconds

`func (o *HciSdaPolicyClusterConfig) SetScheduleIntervalSeconds(v int32)`

SetScheduleIntervalSeconds sets ScheduleIntervalSeconds field to given value.

### HasScheduleIntervalSeconds

`func (o *HciSdaPolicyClusterConfig) HasScheduleIntervalSeconds() bool`

HasScheduleIntervalSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


