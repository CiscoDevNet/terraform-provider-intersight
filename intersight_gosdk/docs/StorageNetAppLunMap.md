# StorageNetAppLunMap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppLunMap"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppLunMap"]
**IgroupUuid** | Pointer to **string** | UUID of the initiator group. | [optional] [readonly] 
**LunUuid** | Pointer to **string** | Universally unique identifier of the LUN. | [optional] [readonly] 
**Host** | Pointer to [**[]StorageNetAppInitiatorGroupRelationship**](StorageNetAppInitiatorGroupRelationship.md) | An array of relationships to storageNetAppInitiatorGroup resources. | [optional] [readonly] 
**Tenant** | Pointer to [**NullableStorageNetAppStorageVmRelationship**](StorageNetAppStorageVmRelationship.md) |  | [optional] 
**Volume** | Pointer to [**NullableStorageNetAppLunRelationship**](StorageNetAppLunRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppLunMap

`func NewStorageNetAppLunMap(classId string, objectType string, ) *StorageNetAppLunMap`

NewStorageNetAppLunMap instantiates a new StorageNetAppLunMap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppLunMapWithDefaults

`func NewStorageNetAppLunMapWithDefaults() *StorageNetAppLunMap`

NewStorageNetAppLunMapWithDefaults instantiates a new StorageNetAppLunMap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppLunMap) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppLunMap) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppLunMap) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppLunMap) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppLunMap) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppLunMap) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIgroupUuid

`func (o *StorageNetAppLunMap) GetIgroupUuid() string`

GetIgroupUuid returns the IgroupUuid field if non-nil, zero value otherwise.

### GetIgroupUuidOk

`func (o *StorageNetAppLunMap) GetIgroupUuidOk() (*string, bool)`

GetIgroupUuidOk returns a tuple with the IgroupUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgroupUuid

`func (o *StorageNetAppLunMap) SetIgroupUuid(v string)`

SetIgroupUuid sets IgroupUuid field to given value.

### HasIgroupUuid

`func (o *StorageNetAppLunMap) HasIgroupUuid() bool`

HasIgroupUuid returns a boolean if a field has been set.

### GetLunUuid

`func (o *StorageNetAppLunMap) GetLunUuid() string`

GetLunUuid returns the LunUuid field if non-nil, zero value otherwise.

### GetLunUuidOk

`func (o *StorageNetAppLunMap) GetLunUuidOk() (*string, bool)`

GetLunUuidOk returns a tuple with the LunUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLunUuid

`func (o *StorageNetAppLunMap) SetLunUuid(v string)`

SetLunUuid sets LunUuid field to given value.

### HasLunUuid

`func (o *StorageNetAppLunMap) HasLunUuid() bool`

HasLunUuid returns a boolean if a field has been set.

### GetHost

`func (o *StorageNetAppLunMap) GetHost() []StorageNetAppInitiatorGroupRelationship`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *StorageNetAppLunMap) GetHostOk() (*[]StorageNetAppInitiatorGroupRelationship, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *StorageNetAppLunMap) SetHost(v []StorageNetAppInitiatorGroupRelationship)`

SetHost sets Host field to given value.

### HasHost

`func (o *StorageNetAppLunMap) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *StorageNetAppLunMap) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *StorageNetAppLunMap) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetTenant

`func (o *StorageNetAppLunMap) GetTenant() StorageNetAppStorageVmRelationship`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *StorageNetAppLunMap) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *StorageNetAppLunMap) SetTenant(v StorageNetAppStorageVmRelationship)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *StorageNetAppLunMap) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *StorageNetAppLunMap) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *StorageNetAppLunMap) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetVolume

`func (o *StorageNetAppLunMap) GetVolume() StorageNetAppLunRelationship`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *StorageNetAppLunMap) GetVolumeOk() (*StorageNetAppLunRelationship, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *StorageNetAppLunMap) SetVolume(v StorageNetAppLunRelationship)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *StorageNetAppLunMap) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### SetVolumeNil

`func (o *StorageNetAppLunMap) SetVolumeNil(b bool)`

 SetVolumeNil sets the value for Volume to be an explicit nil

### UnsetVolume
`func (o *StorageNetAppLunMap) UnsetVolume()`

UnsetVolume ensures that no value is present for Volume, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


