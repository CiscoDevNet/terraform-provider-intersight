# StorageExternalParityGroupAllOf

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

### NewStorageExternalParityGroupAllOf

`func NewStorageExternalParityGroupAllOf(classId string, objectType string, ) *StorageExternalParityGroupAllOf`

NewStorageExternalParityGroupAllOf instantiates a new StorageExternalParityGroupAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageExternalParityGroupAllOfWithDefaults

`func NewStorageExternalParityGroupAllOfWithDefaults() *StorageExternalParityGroupAllOf`

NewStorageExternalParityGroupAllOfWithDefaults instantiates a new StorageExternalParityGroupAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageExternalParityGroupAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageExternalParityGroupAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageExternalParityGroupAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageExternalParityGroupAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageExternalParityGroupAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageExternalParityGroupAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCacheMode

`func (o *StorageExternalParityGroupAllOf) GetCacheMode() string`

GetCacheMode returns the CacheMode field if non-nil, zero value otherwise.

### GetCacheModeOk

`func (o *StorageExternalParityGroupAllOf) GetCacheModeOk() (*string, bool)`

GetCacheModeOk returns a tuple with the CacheMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheMode

`func (o *StorageExternalParityGroupAllOf) SetCacheMode(v string)`

SetCacheMode sets CacheMode field to given value.

### HasCacheMode

`func (o *StorageExternalParityGroupAllOf) HasCacheMode() bool`

HasCacheMode returns a boolean if a field has been set.

### GetExternalLuns

`func (o *StorageExternalParityGroupAllOf) GetExternalLuns() []StorageExternalLun`

GetExternalLuns returns the ExternalLuns field if non-nil, zero value otherwise.

### GetExternalLunsOk

`func (o *StorageExternalParityGroupAllOf) GetExternalLunsOk() (*[]StorageExternalLun, bool)`

GetExternalLunsOk returns a tuple with the ExternalLuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLuns

`func (o *StorageExternalParityGroupAllOf) SetExternalLuns(v []StorageExternalLun)`

SetExternalLuns sets ExternalLuns field to given value.

### HasExternalLuns

`func (o *StorageExternalParityGroupAllOf) HasExternalLuns() bool`

HasExternalLuns returns a boolean if a field has been set.

### SetExternalLunsNil

`func (o *StorageExternalParityGroupAllOf) SetExternalLunsNil(b bool)`

 SetExternalLunsNil sets the value for ExternalLuns to be an explicit nil

### UnsetExternalLuns
`func (o *StorageExternalParityGroupAllOf) UnsetExternalLuns()`

UnsetExternalLuns ensures that no value is present for ExternalLuns, not even an explicit nil
### GetExternalParityGroupId

`func (o *StorageExternalParityGroupAllOf) GetExternalParityGroupId() string`

GetExternalParityGroupId returns the ExternalParityGroupId field if non-nil, zero value otherwise.

### GetExternalParityGroupIdOk

`func (o *StorageExternalParityGroupAllOf) GetExternalParityGroupIdOk() (*string, bool)`

GetExternalParityGroupIdOk returns a tuple with the ExternalParityGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalParityGroupId

`func (o *StorageExternalParityGroupAllOf) SetExternalParityGroupId(v string)`

SetExternalParityGroupId sets ExternalParityGroupId field to given value.

### HasExternalParityGroupId

`func (o *StorageExternalParityGroupAllOf) HasExternalParityGroupId() bool`

HasExternalParityGroupId returns a boolean if a field has been set.

### GetExternalParityGroupStatus

`func (o *StorageExternalParityGroupAllOf) GetExternalParityGroupStatus() string`

GetExternalParityGroupStatus returns the ExternalParityGroupStatus field if non-nil, zero value otherwise.

### GetExternalParityGroupStatusOk

`func (o *StorageExternalParityGroupAllOf) GetExternalParityGroupStatusOk() (*string, bool)`

GetExternalParityGroupStatusOk returns a tuple with the ExternalParityGroupStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalParityGroupStatus

`func (o *StorageExternalParityGroupAllOf) SetExternalParityGroupStatus(v string)`

SetExternalParityGroupStatus sets ExternalParityGroupStatus field to given value.

### HasExternalParityGroupStatus

`func (o *StorageExternalParityGroupAllOf) HasExternalParityGroupStatus() bool`

HasExternalParityGroupStatus returns a boolean if a field has been set.

### GetIsDataDirectMapping

`func (o *StorageExternalParityGroupAllOf) GetIsDataDirectMapping() bool`

GetIsDataDirectMapping returns the IsDataDirectMapping field if non-nil, zero value otherwise.

### GetIsDataDirectMappingOk

`func (o *StorageExternalParityGroupAllOf) GetIsDataDirectMappingOk() (*bool, bool)`

GetIsDataDirectMappingOk returns a tuple with the IsDataDirectMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDataDirectMapping

`func (o *StorageExternalParityGroupAllOf) SetIsDataDirectMapping(v bool)`

SetIsDataDirectMapping sets IsDataDirectMapping field to given value.

### HasIsDataDirectMapping

`func (o *StorageExternalParityGroupAllOf) HasIsDataDirectMapping() bool`

HasIsDataDirectMapping returns a boolean if a field has been set.

### GetIsInflowControlEnabled

`func (o *StorageExternalParityGroupAllOf) GetIsInflowControlEnabled() bool`

GetIsInflowControlEnabled returns the IsInflowControlEnabled field if non-nil, zero value otherwise.

### GetIsInflowControlEnabledOk

`func (o *StorageExternalParityGroupAllOf) GetIsInflowControlEnabledOk() (*bool, bool)`

GetIsInflowControlEnabledOk returns a tuple with the IsInflowControlEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInflowControlEnabled

`func (o *StorageExternalParityGroupAllOf) SetIsInflowControlEnabled(v bool)`

SetIsInflowControlEnabled sets IsInflowControlEnabled field to given value.

### HasIsInflowControlEnabled

`func (o *StorageExternalParityGroupAllOf) HasIsInflowControlEnabled() bool`

HasIsInflowControlEnabled returns a boolean if a field has been set.

### GetLoadBalanceMode

`func (o *StorageExternalParityGroupAllOf) GetLoadBalanceMode() string`

GetLoadBalanceMode returns the LoadBalanceMode field if non-nil, zero value otherwise.

### GetLoadBalanceModeOk

`func (o *StorageExternalParityGroupAllOf) GetLoadBalanceModeOk() (*string, bool)`

GetLoadBalanceModeOk returns a tuple with the LoadBalanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalanceMode

`func (o *StorageExternalParityGroupAllOf) SetLoadBalanceMode(v string)`

SetLoadBalanceMode sets LoadBalanceMode field to given value.

### HasLoadBalanceMode

`func (o *StorageExternalParityGroupAllOf) HasLoadBalanceMode() bool`

HasLoadBalanceMode returns a boolean if a field has been set.

### GetMpBladeId

`func (o *StorageExternalParityGroupAllOf) GetMpBladeId() int64`

GetMpBladeId returns the MpBladeId field if non-nil, zero value otherwise.

### GetMpBladeIdOk

`func (o *StorageExternalParityGroupAllOf) GetMpBladeIdOk() (*int64, bool)`

GetMpBladeIdOk returns a tuple with the MpBladeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpBladeId

`func (o *StorageExternalParityGroupAllOf) SetMpBladeId(v int64)`

SetMpBladeId sets MpBladeId field to given value.

### HasMpBladeId

`func (o *StorageExternalParityGroupAllOf) HasMpBladeId() bool`

HasMpBladeId returns a boolean if a field has been set.

### GetPathMode

`func (o *StorageExternalParityGroupAllOf) GetPathMode() string`

GetPathMode returns the PathMode field if non-nil, zero value otherwise.

### GetPathModeOk

`func (o *StorageExternalParityGroupAllOf) GetPathModeOk() (*string, bool)`

GetPathModeOk returns a tuple with the PathMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathMode

`func (o *StorageExternalParityGroupAllOf) SetPathMode(v string)`

SetPathMode sets PathMode field to given value.

### HasPathMode

`func (o *StorageExternalParityGroupAllOf) HasPathMode() bool`

HasPathMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


