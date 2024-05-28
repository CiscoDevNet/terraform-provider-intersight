# StorageNetAppIscsiService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppIscsiService"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppIscsiService"]
**SvmUuid** | Pointer to **string** | Unique identifier for the NetApp Storage Virtual Machine. | [optional] [readonly] 
**TargetAlias** | Pointer to **string** | The iSCSI target alias of the iSCSI service. | [optional] [readonly] 
**TargetName** | Pointer to **string** | The iSCSI target name of the iSCSI service. | [optional] [readonly] 
**Tenant** | Pointer to [**NullableStorageNetAppStorageVmRelationship**](StorageNetAppStorageVmRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppIscsiService

`func NewStorageNetAppIscsiService(classId string, objectType string, ) *StorageNetAppIscsiService`

NewStorageNetAppIscsiService instantiates a new StorageNetAppIscsiService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppIscsiServiceWithDefaults

`func NewStorageNetAppIscsiServiceWithDefaults() *StorageNetAppIscsiService`

NewStorageNetAppIscsiServiceWithDefaults instantiates a new StorageNetAppIscsiService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppIscsiService) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppIscsiService) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppIscsiService) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppIscsiService) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppIscsiService) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppIscsiService) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSvmUuid

`func (o *StorageNetAppIscsiService) GetSvmUuid() string`

GetSvmUuid returns the SvmUuid field if non-nil, zero value otherwise.

### GetSvmUuidOk

`func (o *StorageNetAppIscsiService) GetSvmUuidOk() (*string, bool)`

GetSvmUuidOk returns a tuple with the SvmUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvmUuid

`func (o *StorageNetAppIscsiService) SetSvmUuid(v string)`

SetSvmUuid sets SvmUuid field to given value.

### HasSvmUuid

`func (o *StorageNetAppIscsiService) HasSvmUuid() bool`

HasSvmUuid returns a boolean if a field has been set.

### GetTargetAlias

`func (o *StorageNetAppIscsiService) GetTargetAlias() string`

GetTargetAlias returns the TargetAlias field if non-nil, zero value otherwise.

### GetTargetAliasOk

`func (o *StorageNetAppIscsiService) GetTargetAliasOk() (*string, bool)`

GetTargetAliasOk returns a tuple with the TargetAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAlias

`func (o *StorageNetAppIscsiService) SetTargetAlias(v string)`

SetTargetAlias sets TargetAlias field to given value.

### HasTargetAlias

`func (o *StorageNetAppIscsiService) HasTargetAlias() bool`

HasTargetAlias returns a boolean if a field has been set.

### GetTargetName

`func (o *StorageNetAppIscsiService) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *StorageNetAppIscsiService) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *StorageNetAppIscsiService) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *StorageNetAppIscsiService) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetTenant

`func (o *StorageNetAppIscsiService) GetTenant() StorageNetAppStorageVmRelationship`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *StorageNetAppIscsiService) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *StorageNetAppIscsiService) SetTenant(v StorageNetAppStorageVmRelationship)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *StorageNetAppIscsiService) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *StorageNetAppIscsiService) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *StorageNetAppIscsiService) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


