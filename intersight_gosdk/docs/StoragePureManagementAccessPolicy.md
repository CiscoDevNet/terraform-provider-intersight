# StoragePureManagementAccessPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.PureManagementAccessPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.PureManagementAccessPolicy"]
**AggregationStrategy** | Pointer to **string** | The strategy used to combine the effects of multiple rules in this policy (e.g., &#39;deny-override&#39;). | [optional] [readonly] 
**Enabled** | Pointer to **bool** | Returns a value of true if the management access policy is currently active and enforced. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the management access policy. | [optional] [readonly] 
**PolicyType** | Pointer to **string** | The type of policy, typically &#39;management-access&#39;. | [optional] [readonly] 
**Array** | Pointer to [**NullableStoragePureArrayRelationship**](StoragePureArrayRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStoragePureManagementAccessPolicy

`func NewStoragePureManagementAccessPolicy(classId string, objectType string, ) *StoragePureManagementAccessPolicy`

NewStoragePureManagementAccessPolicy instantiates a new StoragePureManagementAccessPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePureManagementAccessPolicyWithDefaults

`func NewStoragePureManagementAccessPolicyWithDefaults() *StoragePureManagementAccessPolicy`

NewStoragePureManagementAccessPolicyWithDefaults instantiates a new StoragePureManagementAccessPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StoragePureManagementAccessPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StoragePureManagementAccessPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StoragePureManagementAccessPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StoragePureManagementAccessPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StoragePureManagementAccessPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StoragePureManagementAccessPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAggregationStrategy

`func (o *StoragePureManagementAccessPolicy) GetAggregationStrategy() string`

GetAggregationStrategy returns the AggregationStrategy field if non-nil, zero value otherwise.

### GetAggregationStrategyOk

`func (o *StoragePureManagementAccessPolicy) GetAggregationStrategyOk() (*string, bool)`

GetAggregationStrategyOk returns a tuple with the AggregationStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationStrategy

`func (o *StoragePureManagementAccessPolicy) SetAggregationStrategy(v string)`

SetAggregationStrategy sets AggregationStrategy field to given value.

### HasAggregationStrategy

`func (o *StoragePureManagementAccessPolicy) HasAggregationStrategy() bool`

HasAggregationStrategy returns a boolean if a field has been set.

### GetEnabled

`func (o *StoragePureManagementAccessPolicy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StoragePureManagementAccessPolicy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StoragePureManagementAccessPolicy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *StoragePureManagementAccessPolicy) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *StoragePureManagementAccessPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StoragePureManagementAccessPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StoragePureManagementAccessPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StoragePureManagementAccessPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolicyType

`func (o *StoragePureManagementAccessPolicy) GetPolicyType() string`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *StoragePureManagementAccessPolicy) GetPolicyTypeOk() (*string, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *StoragePureManagementAccessPolicy) SetPolicyType(v string)`

SetPolicyType sets PolicyType field to given value.

### HasPolicyType

`func (o *StoragePureManagementAccessPolicy) HasPolicyType() bool`

HasPolicyType returns a boolean if a field has been set.

### GetArray

`func (o *StoragePureManagementAccessPolicy) GetArray() StoragePureArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StoragePureManagementAccessPolicy) GetArrayOk() (*StoragePureArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StoragePureManagementAccessPolicy) SetArray(v StoragePureArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StoragePureManagementAccessPolicy) HasArray() bool`

HasArray returns a boolean if a field has been set.

### SetArrayNil

`func (o *StoragePureManagementAccessPolicy) SetArrayNil(b bool)`

 SetArrayNil sets the value for Array to be an explicit nil

### UnsetArray
`func (o *StoragePureManagementAccessPolicy) UnsetArray()`

UnsetArray ensures that no value is present for Array, not even an explicit nil
### GetRegisteredDevice

`func (o *StoragePureManagementAccessPolicy) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StoragePureManagementAccessPolicy) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StoragePureManagementAccessPolicy) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StoragePureManagementAccessPolicy) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *StoragePureManagementAccessPolicy) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *StoragePureManagementAccessPolicy) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


