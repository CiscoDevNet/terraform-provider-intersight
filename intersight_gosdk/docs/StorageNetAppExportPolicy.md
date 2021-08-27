# StorageNetAppExportPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppExportPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppExportPolicy"]
**ClusterUuid** | Pointer to **string** | Unique identity of the device. | [optional] [readonly] 
**NetAppExportPolicyRule** | Pointer to [**[]StorageNetAppExportPolicyRule**](StorageNetAppExportPolicyRule.md) |  | [optional] 
**PolicyId** | Pointer to **int64** | ID for the Export Policy. | [optional] [readonly] 
**Array** | Pointer to [**StorageNetAppClusterRelationship**](StorageNetAppClusterRelationship.md) |  | [optional] 
**Tenant** | Pointer to [**StorageNetAppStorageVmRelationship**](StorageNetAppStorageVmRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppExportPolicy

`func NewStorageNetAppExportPolicy(classId string, objectType string, ) *StorageNetAppExportPolicy`

NewStorageNetAppExportPolicy instantiates a new StorageNetAppExportPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppExportPolicyWithDefaults

`func NewStorageNetAppExportPolicyWithDefaults() *StorageNetAppExportPolicy`

NewStorageNetAppExportPolicyWithDefaults instantiates a new StorageNetAppExportPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppExportPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppExportPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppExportPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppExportPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppExportPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppExportPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClusterUuid

`func (o *StorageNetAppExportPolicy) GetClusterUuid() string`

GetClusterUuid returns the ClusterUuid field if non-nil, zero value otherwise.

### GetClusterUuidOk

`func (o *StorageNetAppExportPolicy) GetClusterUuidOk() (*string, bool)`

GetClusterUuidOk returns a tuple with the ClusterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterUuid

`func (o *StorageNetAppExportPolicy) SetClusterUuid(v string)`

SetClusterUuid sets ClusterUuid field to given value.

### HasClusterUuid

`func (o *StorageNetAppExportPolicy) HasClusterUuid() bool`

HasClusterUuid returns a boolean if a field has been set.

### GetNetAppExportPolicyRule

`func (o *StorageNetAppExportPolicy) GetNetAppExportPolicyRule() []StorageNetAppExportPolicyRule`

GetNetAppExportPolicyRule returns the NetAppExportPolicyRule field if non-nil, zero value otherwise.

### GetNetAppExportPolicyRuleOk

`func (o *StorageNetAppExportPolicy) GetNetAppExportPolicyRuleOk() (*[]StorageNetAppExportPolicyRule, bool)`

GetNetAppExportPolicyRuleOk returns a tuple with the NetAppExportPolicyRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAppExportPolicyRule

`func (o *StorageNetAppExportPolicy) SetNetAppExportPolicyRule(v []StorageNetAppExportPolicyRule)`

SetNetAppExportPolicyRule sets NetAppExportPolicyRule field to given value.

### HasNetAppExportPolicyRule

`func (o *StorageNetAppExportPolicy) HasNetAppExportPolicyRule() bool`

HasNetAppExportPolicyRule returns a boolean if a field has been set.

### SetNetAppExportPolicyRuleNil

`func (o *StorageNetAppExportPolicy) SetNetAppExportPolicyRuleNil(b bool)`

 SetNetAppExportPolicyRuleNil sets the value for NetAppExportPolicyRule to be an explicit nil

### UnsetNetAppExportPolicyRule
`func (o *StorageNetAppExportPolicy) UnsetNetAppExportPolicyRule()`

UnsetNetAppExportPolicyRule ensures that no value is present for NetAppExportPolicyRule, not even an explicit nil
### GetPolicyId

`func (o *StorageNetAppExportPolicy) GetPolicyId() int64`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *StorageNetAppExportPolicy) GetPolicyIdOk() (*int64, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *StorageNetAppExportPolicy) SetPolicyId(v int64)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *StorageNetAppExportPolicy) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetArray

`func (o *StorageNetAppExportPolicy) GetArray() StorageNetAppClusterRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageNetAppExportPolicy) GetArrayOk() (*StorageNetAppClusterRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageNetAppExportPolicy) SetArray(v StorageNetAppClusterRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageNetAppExportPolicy) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetTenant

`func (o *StorageNetAppExportPolicy) GetTenant() StorageNetAppStorageVmRelationship`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *StorageNetAppExportPolicy) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *StorageNetAppExportPolicy) SetTenant(v StorageNetAppStorageVmRelationship)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *StorageNetAppExportPolicy) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


