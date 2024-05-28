# StorageNetAppSvmSnapshotPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppSvmSnapshotPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppSvmSnapshotPolicy"]
**SvmName** | Pointer to **string** | The storage virtual machine name for the policy. | [optional] [readonly] 
**Tenant** | Pointer to [**NullableStorageNetAppStorageVmRelationship**](StorageNetAppStorageVmRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppSvmSnapshotPolicy

`func NewStorageNetAppSvmSnapshotPolicy(classId string, objectType string, ) *StorageNetAppSvmSnapshotPolicy`

NewStorageNetAppSvmSnapshotPolicy instantiates a new StorageNetAppSvmSnapshotPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppSvmSnapshotPolicyWithDefaults

`func NewStorageNetAppSvmSnapshotPolicyWithDefaults() *StorageNetAppSvmSnapshotPolicy`

NewStorageNetAppSvmSnapshotPolicyWithDefaults instantiates a new StorageNetAppSvmSnapshotPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppSvmSnapshotPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppSvmSnapshotPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppSvmSnapshotPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppSvmSnapshotPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppSvmSnapshotPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppSvmSnapshotPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSvmName

`func (o *StorageNetAppSvmSnapshotPolicy) GetSvmName() string`

GetSvmName returns the SvmName field if non-nil, zero value otherwise.

### GetSvmNameOk

`func (o *StorageNetAppSvmSnapshotPolicy) GetSvmNameOk() (*string, bool)`

GetSvmNameOk returns a tuple with the SvmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvmName

`func (o *StorageNetAppSvmSnapshotPolicy) SetSvmName(v string)`

SetSvmName sets SvmName field to given value.

### HasSvmName

`func (o *StorageNetAppSvmSnapshotPolicy) HasSvmName() bool`

HasSvmName returns a boolean if a field has been set.

### GetTenant

`func (o *StorageNetAppSvmSnapshotPolicy) GetTenant() StorageNetAppStorageVmRelationship`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *StorageNetAppSvmSnapshotPolicy) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *StorageNetAppSvmSnapshotPolicy) SetTenant(v StorageNetAppStorageVmRelationship)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *StorageNetAppSvmSnapshotPolicy) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *StorageNetAppSvmSnapshotPolicy) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *StorageNetAppSvmSnapshotPolicy) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


