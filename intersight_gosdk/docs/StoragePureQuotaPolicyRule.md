# StoragePureQuotaPolicyRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.PureQuotaPolicyRule"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.PureQuotaPolicyRule"]
**Destroyed** | Pointer to **bool** | A boolean value, if set to true, lists only destroyed objects that are in the eradication pending state. If set to false, lists only objects that are not destroyed. | [optional] [readonly] 
**Enforced** | Pointer to **bool** | Enforced of the policy rule. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the quota policy rule. | [optional] [readonly] 
**Notifications** | Pointer to **string** | Notifications of the policy rule. | [optional] [readonly] 
**PolicyName** | Pointer to **string** | Name of the NFS policy rule. | [optional] [readonly] 
**QuotaLimit** | Pointer to **int64** | Quota limit of the policy rule. | [optional] [readonly] 
**Array** | Pointer to [**NullableStoragePureArrayRelationship**](StoragePureArrayRelationship.md) |  | [optional] 
**Policy** | Pointer to [**NullableStoragePureDirectoryPolicyRelationship**](StoragePureDirectoryPolicyRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewStoragePureQuotaPolicyRule

`func NewStoragePureQuotaPolicyRule(classId string, objectType string, ) *StoragePureQuotaPolicyRule`

NewStoragePureQuotaPolicyRule instantiates a new StoragePureQuotaPolicyRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePureQuotaPolicyRuleWithDefaults

`func NewStoragePureQuotaPolicyRuleWithDefaults() *StoragePureQuotaPolicyRule`

NewStoragePureQuotaPolicyRuleWithDefaults instantiates a new StoragePureQuotaPolicyRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StoragePureQuotaPolicyRule) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StoragePureQuotaPolicyRule) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StoragePureQuotaPolicyRule) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StoragePureQuotaPolicyRule) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StoragePureQuotaPolicyRule) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StoragePureQuotaPolicyRule) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDestroyed

`func (o *StoragePureQuotaPolicyRule) GetDestroyed() bool`

GetDestroyed returns the Destroyed field if non-nil, zero value otherwise.

### GetDestroyedOk

`func (o *StoragePureQuotaPolicyRule) GetDestroyedOk() (*bool, bool)`

GetDestroyedOk returns a tuple with the Destroyed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyed

`func (o *StoragePureQuotaPolicyRule) SetDestroyed(v bool)`

SetDestroyed sets Destroyed field to given value.

### HasDestroyed

`func (o *StoragePureQuotaPolicyRule) HasDestroyed() bool`

HasDestroyed returns a boolean if a field has been set.

### GetEnforced

`func (o *StoragePureQuotaPolicyRule) GetEnforced() bool`

GetEnforced returns the Enforced field if non-nil, zero value otherwise.

### GetEnforcedOk

`func (o *StoragePureQuotaPolicyRule) GetEnforcedOk() (*bool, bool)`

GetEnforcedOk returns a tuple with the Enforced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforced

`func (o *StoragePureQuotaPolicyRule) SetEnforced(v bool)`

SetEnforced sets Enforced field to given value.

### HasEnforced

`func (o *StoragePureQuotaPolicyRule) HasEnforced() bool`

HasEnforced returns a boolean if a field has been set.

### GetName

`func (o *StoragePureQuotaPolicyRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StoragePureQuotaPolicyRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StoragePureQuotaPolicyRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StoragePureQuotaPolicyRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotifications

`func (o *StoragePureQuotaPolicyRule) GetNotifications() string`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *StoragePureQuotaPolicyRule) GetNotificationsOk() (*string, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *StoragePureQuotaPolicyRule) SetNotifications(v string)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *StoragePureQuotaPolicyRule) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetPolicyName

`func (o *StoragePureQuotaPolicyRule) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *StoragePureQuotaPolicyRule) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *StoragePureQuotaPolicyRule) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *StoragePureQuotaPolicyRule) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetQuotaLimit

`func (o *StoragePureQuotaPolicyRule) GetQuotaLimit() int64`

GetQuotaLimit returns the QuotaLimit field if non-nil, zero value otherwise.

### GetQuotaLimitOk

`func (o *StoragePureQuotaPolicyRule) GetQuotaLimitOk() (*int64, bool)`

GetQuotaLimitOk returns a tuple with the QuotaLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaLimit

`func (o *StoragePureQuotaPolicyRule) SetQuotaLimit(v int64)`

SetQuotaLimit sets QuotaLimit field to given value.

### HasQuotaLimit

`func (o *StoragePureQuotaPolicyRule) HasQuotaLimit() bool`

HasQuotaLimit returns a boolean if a field has been set.

### GetArray

`func (o *StoragePureQuotaPolicyRule) GetArray() StoragePureArrayRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StoragePureQuotaPolicyRule) GetArrayOk() (*StoragePureArrayRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StoragePureQuotaPolicyRule) SetArray(v StoragePureArrayRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StoragePureQuotaPolicyRule) HasArray() bool`

HasArray returns a boolean if a field has been set.

### SetArrayNil

`func (o *StoragePureQuotaPolicyRule) SetArrayNil(b bool)`

 SetArrayNil sets the value for Array to be an explicit nil

### UnsetArray
`func (o *StoragePureQuotaPolicyRule) UnsetArray()`

UnsetArray ensures that no value is present for Array, not even an explicit nil
### GetPolicy

`func (o *StoragePureQuotaPolicyRule) GetPolicy() StoragePureDirectoryPolicyRelationship`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *StoragePureQuotaPolicyRule) GetPolicyOk() (*StoragePureDirectoryPolicyRelationship, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *StoragePureQuotaPolicyRule) SetPolicy(v StoragePureDirectoryPolicyRelationship)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *StoragePureQuotaPolicyRule) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### SetPolicyNil

`func (o *StoragePureQuotaPolicyRule) SetPolicyNil(b bool)`

 SetPolicyNil sets the value for Policy to be an explicit nil

### UnsetPolicy
`func (o *StoragePureQuotaPolicyRule) UnsetPolicy()`

UnsetPolicy ensures that no value is present for Policy, not even an explicit nil
### GetRegisteredDevice

`func (o *StoragePureQuotaPolicyRule) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *StoragePureQuotaPolicyRule) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *StoragePureQuotaPolicyRule) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *StoragePureQuotaPolicyRule) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *StoragePureQuotaPolicyRule) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *StoragePureQuotaPolicyRule) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


