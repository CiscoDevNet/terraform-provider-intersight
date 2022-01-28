# StorageNetAppSvmEventAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppSvmEvent"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppSvmEvent"]
**Tenant** | Pointer to [**StorageNetAppStorageVmRelationship**](StorageNetAppStorageVmRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppSvmEventAllOf

`func NewStorageNetAppSvmEventAllOf(classId string, objectType string, ) *StorageNetAppSvmEventAllOf`

NewStorageNetAppSvmEventAllOf instantiates a new StorageNetAppSvmEventAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppSvmEventAllOfWithDefaults

`func NewStorageNetAppSvmEventAllOfWithDefaults() *StorageNetAppSvmEventAllOf`

NewStorageNetAppSvmEventAllOfWithDefaults instantiates a new StorageNetAppSvmEventAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppSvmEventAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppSvmEventAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppSvmEventAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppSvmEventAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppSvmEventAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppSvmEventAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetTenant

`func (o *StorageNetAppSvmEventAllOf) GetTenant() StorageNetAppStorageVmRelationship`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *StorageNetAppSvmEventAllOf) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *StorageNetAppSvmEventAllOf) SetTenant(v StorageNetAppStorageVmRelationship)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *StorageNetAppSvmEventAllOf) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


