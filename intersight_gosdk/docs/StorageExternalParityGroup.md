# StorageExternalParityGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.ExternalParityGroup"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.ExternalParityGroup"]
**CacheMode** | Pointer to **string** | Cache mode. Either E (enabled) or D (disabled) is displayed. | [optional] [readonly] 
**ExternalLuns** | Pointer to [**[]StorageExternalLun**](StorageExternalLun.md) |  | [optional] 
**ExternalParityGroupId** | Pointer to **string** | External parity group number. | [optional] [readonly] 
**ExternalParityGroupStatus** | Pointer to **string** | Status of the external parity group. | [optional] [readonly] 
**IsDataDirectMapping** | Pointer to **bool** | Whether the data direct mapping attribute is enabled. | [optional] [readonly] 
**IsInflowControlEnabled** | Pointer to **bool** | Inflow cache control. Either true (enabled) or false (disabled) is displayed. | [optional] [readonly] 
**LoadBalanceMode** | Pointer to **string** | The load balancing method for I-O operations for the external storage system. | [optional] [readonly] 
**MpBladeId** | Pointer to **int64** | The blade number of the MP blade assigned to the external parity group. | [optional] [readonly] 
**PathMode** | Pointer to **string** | Path mode of the external storage system. | [optional] [readonly] 

## Methods

### NewStorageExternalParityGroup

`func NewStorageExternalParityGroup(classId string, objectType string, ) *StorageExternalParityGroup`

NewStorageExternalParityGroup instantiates a new StorageExternalParityGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageExternalParityGroupWithDefaults

`func NewStorageExternalParityGroupWithDefaults() *StorageExternalParityGroup`

NewStorageExternalParityGroupWithDefaults instantiates a new StorageExternalParityGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageExternalParityGroup) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageExternalParityGroup) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageExternalParityGroup) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageExternalParityGroup) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageExternalParityGroup) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageExternalParityGroup) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCacheMode

`func (o *StorageExternalParityGroup) GetCacheMode() string`

GetCacheMode returns the CacheMode field if non-nil, zero value otherwise.

### GetCacheModeOk

`func (o *StorageExternalParityGroup) GetCacheModeOk() (*string, bool)`

GetCacheModeOk returns a tuple with the CacheMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheMode

`func (o *StorageExternalParityGroup) SetCacheMode(v string)`

SetCacheMode sets CacheMode field to given value.

### HasCacheMode

`func (o *StorageExternalParityGroup) HasCacheMode() bool`

HasCacheMode returns a boolean if a field has been set.

### GetExternalLuns

`func (o *StorageExternalParityGroup) GetExternalLuns() []StorageExternalLun`

GetExternalLuns returns the ExternalLuns field if non-nil, zero value otherwise.

### GetExternalLunsOk

`func (o *StorageExternalParityGroup) GetExternalLunsOk() (*[]StorageExternalLun, bool)`

GetExternalLunsOk returns a tuple with the ExternalLuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLuns

`func (o *StorageExternalParityGroup) SetExternalLuns(v []StorageExternalLun)`

SetExternalLuns sets ExternalLuns field to given value.

### HasExternalLuns

`func (o *StorageExternalParityGroup) HasExternalLuns() bool`

HasExternalLuns returns a boolean if a field has been set.

### SetExternalLunsNil

`func (o *StorageExternalParityGroup) SetExternalLunsNil(b bool)`

 SetExternalLunsNil sets the value for ExternalLuns to be an explicit nil

### UnsetExternalLuns
`func (o *StorageExternalParityGroup) UnsetExternalLuns()`

UnsetExternalLuns ensures that no value is present for ExternalLuns, not even an explicit nil
### GetExternalParityGroupId

`func (o *StorageExternalParityGroup) GetExternalParityGroupId() string`

GetExternalParityGroupId returns the ExternalParityGroupId field if non-nil, zero value otherwise.

### GetExternalParityGroupIdOk

`func (o *StorageExternalParityGroup) GetExternalParityGroupIdOk() (*string, bool)`

GetExternalParityGroupIdOk returns a tuple with the ExternalParityGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalParityGroupId

`func (o *StorageExternalParityGroup) SetExternalParityGroupId(v string)`

SetExternalParityGroupId sets ExternalParityGroupId field to given value.

### HasExternalParityGroupId

`func (o *StorageExternalParityGroup) HasExternalParityGroupId() bool`

HasExternalParityGroupId returns a boolean if a field has been set.

### GetExternalParityGroupStatus

`func (o *StorageExternalParityGroup) GetExternalParityGroupStatus() string`

GetExternalParityGroupStatus returns the ExternalParityGroupStatus field if non-nil, zero value otherwise.

### GetExternalParityGroupStatusOk

`func (o *StorageExternalParityGroup) GetExternalParityGroupStatusOk() (*string, bool)`

GetExternalParityGroupStatusOk returns a tuple with the ExternalParityGroupStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalParityGroupStatus

`func (o *StorageExternalParityGroup) SetExternalParityGroupStatus(v string)`

SetExternalParityGroupStatus sets ExternalParityGroupStatus field to given value.

### HasExternalParityGroupStatus

`func (o *StorageExternalParityGroup) HasExternalParityGroupStatus() bool`

HasExternalParityGroupStatus returns a boolean if a field has been set.

### GetIsDataDirectMapping

`func (o *StorageExternalParityGroup) GetIsDataDirectMapping() bool`

GetIsDataDirectMapping returns the IsDataDirectMapping field if non-nil, zero value otherwise.

### GetIsDataDirectMappingOk

`func (o *StorageExternalParityGroup) GetIsDataDirectMappingOk() (*bool, bool)`

GetIsDataDirectMappingOk returns a tuple with the IsDataDirectMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDataDirectMapping

`func (o *StorageExternalParityGroup) SetIsDataDirectMapping(v bool)`

SetIsDataDirectMapping sets IsDataDirectMapping field to given value.

### HasIsDataDirectMapping

`func (o *StorageExternalParityGroup) HasIsDataDirectMapping() bool`

HasIsDataDirectMapping returns a boolean if a field has been set.

### GetIsInflowControlEnabled

`func (o *StorageExternalParityGroup) GetIsInflowControlEnabled() bool`

GetIsInflowControlEnabled returns the IsInflowControlEnabled field if non-nil, zero value otherwise.

### GetIsInflowControlEnabledOk

`func (o *StorageExternalParityGroup) GetIsInflowControlEnabledOk() (*bool, bool)`

GetIsInflowControlEnabledOk returns a tuple with the IsInflowControlEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInflowControlEnabled

`func (o *StorageExternalParityGroup) SetIsInflowControlEnabled(v bool)`

SetIsInflowControlEnabled sets IsInflowControlEnabled field to given value.

### HasIsInflowControlEnabled

`func (o *StorageExternalParityGroup) HasIsInflowControlEnabled() bool`

HasIsInflowControlEnabled returns a boolean if a field has been set.

### GetLoadBalanceMode

`func (o *StorageExternalParityGroup) GetLoadBalanceMode() string`

GetLoadBalanceMode returns the LoadBalanceMode field if non-nil, zero value otherwise.

### GetLoadBalanceModeOk

`func (o *StorageExternalParityGroup) GetLoadBalanceModeOk() (*string, bool)`

GetLoadBalanceModeOk returns a tuple with the LoadBalanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalanceMode

`func (o *StorageExternalParityGroup) SetLoadBalanceMode(v string)`

SetLoadBalanceMode sets LoadBalanceMode field to given value.

### HasLoadBalanceMode

`func (o *StorageExternalParityGroup) HasLoadBalanceMode() bool`

HasLoadBalanceMode returns a boolean if a field has been set.

### GetMpBladeId

`func (o *StorageExternalParityGroup) GetMpBladeId() int64`

GetMpBladeId returns the MpBladeId field if non-nil, zero value otherwise.

### GetMpBladeIdOk

`func (o *StorageExternalParityGroup) GetMpBladeIdOk() (*int64, bool)`

GetMpBladeIdOk returns a tuple with the MpBladeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpBladeId

`func (o *StorageExternalParityGroup) SetMpBladeId(v int64)`

SetMpBladeId sets MpBladeId field to given value.

### HasMpBladeId

`func (o *StorageExternalParityGroup) HasMpBladeId() bool`

HasMpBladeId returns a boolean if a field has been set.

### GetPathMode

`func (o *StorageExternalParityGroup) GetPathMode() string`

GetPathMode returns the PathMode field if non-nil, zero value otherwise.

### GetPathModeOk

`func (o *StorageExternalParityGroup) GetPathModeOk() (*string, bool)`

GetPathModeOk returns a tuple with the PathMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathMode

`func (o *StorageExternalParityGroup) SetPathMode(v string)`

SetPathMode sets PathMode field to given value.

### HasPathMode

`func (o *StorageExternalParityGroup) HasPathMode() bool`

HasPathMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


