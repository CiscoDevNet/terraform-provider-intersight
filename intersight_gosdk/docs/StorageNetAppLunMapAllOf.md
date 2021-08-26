# StorageNetAppLunMapAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppLunMap"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppLunMap"]
**Uuid** | Pointer to **string** | UUID of the LUN. | [optional] [readonly] 
**Host** | Pointer to [**[]StorageNetAppInitiatorGroupRelationship**](StorageNetAppInitiatorGroupRelationship.md) | An array of relationships to storageNetAppInitiatorGroup resources. | [optional] [readonly] 
**Tenant** | Pointer to [**StorageNetAppStorageVmRelationship**](StorageNetAppStorageVmRelationship.md) |  | [optional] 
**Volume** | Pointer to [**StorageNetAppLunRelationship**](StorageNetAppLunRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppLunMapAllOf

`func NewStorageNetAppLunMapAllOf(classId string, objectType string, ) *StorageNetAppLunMapAllOf`

NewStorageNetAppLunMapAllOf instantiates a new StorageNetAppLunMapAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppLunMapAllOfWithDefaults

`func NewStorageNetAppLunMapAllOfWithDefaults() *StorageNetAppLunMapAllOf`

NewStorageNetAppLunMapAllOfWithDefaults instantiates a new StorageNetAppLunMapAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppLunMapAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppLunMapAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppLunMapAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppLunMapAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppLunMapAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppLunMapAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetUuid

`func (o *StorageNetAppLunMapAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StorageNetAppLunMapAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StorageNetAppLunMapAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StorageNetAppLunMapAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetHost

`func (o *StorageNetAppLunMapAllOf) GetHost() []StorageNetAppInitiatorGroupRelationship`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *StorageNetAppLunMapAllOf) GetHostOk() (*[]StorageNetAppInitiatorGroupRelationship, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *StorageNetAppLunMapAllOf) SetHost(v []StorageNetAppInitiatorGroupRelationship)`

SetHost sets Host field to given value.

### HasHost

`func (o *StorageNetAppLunMapAllOf) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *StorageNetAppLunMapAllOf) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *StorageNetAppLunMapAllOf) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetTenant

`func (o *StorageNetAppLunMapAllOf) GetTenant() StorageNetAppStorageVmRelationship`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *StorageNetAppLunMapAllOf) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *StorageNetAppLunMapAllOf) SetTenant(v StorageNetAppStorageVmRelationship)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *StorageNetAppLunMapAllOf) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetVolume

`func (o *StorageNetAppLunMapAllOf) GetVolume() StorageNetAppLunRelationship`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *StorageNetAppLunMapAllOf) GetVolumeOk() (*StorageNetAppLunRelationship, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *StorageNetAppLunMapAllOf) SetVolume(v StorageNetAppLunRelationship)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *StorageNetAppLunMapAllOf) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


