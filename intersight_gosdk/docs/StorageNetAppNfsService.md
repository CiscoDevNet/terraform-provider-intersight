# StorageNetAppNfsService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppNfsService"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppNfsService"]
**NfsV3Enabled** | Pointer to **bool** | Specifies whether NFSv3 protocol is enabled. | [optional] [readonly] 
**NfsV41Enabled** | Pointer to **bool** | Specifies whether NFSv4.1 protocol is enabled. | [optional] [readonly] 
**NfsV4Enabled** | Pointer to **bool** | Specifies whether NFSv4.0 protocol is enabled. | [optional] [readonly] 
**SvmUuid** | Pointer to **string** | Unique identifier for the NetApp Storage Virtual Machine. | [optional] [readonly] 
**Tenant** | Pointer to [**NullableStorageNetAppStorageVmRelationship**](StorageNetAppStorageVmRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppNfsService

`func NewStorageNetAppNfsService(classId string, objectType string, ) *StorageNetAppNfsService`

NewStorageNetAppNfsService instantiates a new StorageNetAppNfsService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppNfsServiceWithDefaults

`func NewStorageNetAppNfsServiceWithDefaults() *StorageNetAppNfsService`

NewStorageNetAppNfsServiceWithDefaults instantiates a new StorageNetAppNfsService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppNfsService) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppNfsService) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppNfsService) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppNfsService) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppNfsService) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppNfsService) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetNfsV3Enabled

`func (o *StorageNetAppNfsService) GetNfsV3Enabled() bool`

GetNfsV3Enabled returns the NfsV3Enabled field if non-nil, zero value otherwise.

### GetNfsV3EnabledOk

`func (o *StorageNetAppNfsService) GetNfsV3EnabledOk() (*bool, bool)`

GetNfsV3EnabledOk returns a tuple with the NfsV3Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsV3Enabled

`func (o *StorageNetAppNfsService) SetNfsV3Enabled(v bool)`

SetNfsV3Enabled sets NfsV3Enabled field to given value.

### HasNfsV3Enabled

`func (o *StorageNetAppNfsService) HasNfsV3Enabled() bool`

HasNfsV3Enabled returns a boolean if a field has been set.

### GetNfsV41Enabled

`func (o *StorageNetAppNfsService) GetNfsV41Enabled() bool`

GetNfsV41Enabled returns the NfsV41Enabled field if non-nil, zero value otherwise.

### GetNfsV41EnabledOk

`func (o *StorageNetAppNfsService) GetNfsV41EnabledOk() (*bool, bool)`

GetNfsV41EnabledOk returns a tuple with the NfsV41Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsV41Enabled

`func (o *StorageNetAppNfsService) SetNfsV41Enabled(v bool)`

SetNfsV41Enabled sets NfsV41Enabled field to given value.

### HasNfsV41Enabled

`func (o *StorageNetAppNfsService) HasNfsV41Enabled() bool`

HasNfsV41Enabled returns a boolean if a field has been set.

### GetNfsV4Enabled

`func (o *StorageNetAppNfsService) GetNfsV4Enabled() bool`

GetNfsV4Enabled returns the NfsV4Enabled field if non-nil, zero value otherwise.

### GetNfsV4EnabledOk

`func (o *StorageNetAppNfsService) GetNfsV4EnabledOk() (*bool, bool)`

GetNfsV4EnabledOk returns a tuple with the NfsV4Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsV4Enabled

`func (o *StorageNetAppNfsService) SetNfsV4Enabled(v bool)`

SetNfsV4Enabled sets NfsV4Enabled field to given value.

### HasNfsV4Enabled

`func (o *StorageNetAppNfsService) HasNfsV4Enabled() bool`

HasNfsV4Enabled returns a boolean if a field has been set.

### GetSvmUuid

`func (o *StorageNetAppNfsService) GetSvmUuid() string`

GetSvmUuid returns the SvmUuid field if non-nil, zero value otherwise.

### GetSvmUuidOk

`func (o *StorageNetAppNfsService) GetSvmUuidOk() (*string, bool)`

GetSvmUuidOk returns a tuple with the SvmUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvmUuid

`func (o *StorageNetAppNfsService) SetSvmUuid(v string)`

SetSvmUuid sets SvmUuid field to given value.

### HasSvmUuid

`func (o *StorageNetAppNfsService) HasSvmUuid() bool`

HasSvmUuid returns a boolean if a field has been set.

### GetTenant

`func (o *StorageNetAppNfsService) GetTenant() StorageNetAppStorageVmRelationship`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *StorageNetAppNfsService) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *StorageNetAppNfsService) SetTenant(v StorageNetAppStorageVmRelationship)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *StorageNetAppNfsService) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *StorageNetAppNfsService) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *StorageNetAppNfsService) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


