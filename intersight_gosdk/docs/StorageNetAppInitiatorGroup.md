# StorageNetAppInitiatorGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppInitiatorGroup"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppInitiatorGroup"]
**Protocol** | Pointer to **string** | Initiator group protocol. * &#x60;FCP&#x60; - Fibre channel initiator type which contains WWN of an HBA on the host. * &#x60;iSCSI&#x60; - An iSCSI initiator type used by the host. * &#x60;mixed&#x60; - For systems using both FC and iSCSI connections to the same LUN, create two igroups, one for FC and one for iSCSI. Then map the LUN to both igroups. | [optional] [readonly] [default to "FCP"]
**Uuid** | Pointer to **string** | UUID of the LUN. | [optional] [readonly] 
**Tenant** | Pointer to [**StorageNetAppStorageVmRelationship**](StorageNetAppStorageVmRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppInitiatorGroup

`func NewStorageNetAppInitiatorGroup(classId string, objectType string, ) *StorageNetAppInitiatorGroup`

NewStorageNetAppInitiatorGroup instantiates a new StorageNetAppInitiatorGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppInitiatorGroupWithDefaults

`func NewStorageNetAppInitiatorGroupWithDefaults() *StorageNetAppInitiatorGroup`

NewStorageNetAppInitiatorGroupWithDefaults instantiates a new StorageNetAppInitiatorGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppInitiatorGroup) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppInitiatorGroup) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppInitiatorGroup) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppInitiatorGroup) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppInitiatorGroup) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppInitiatorGroup) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetProtocol

`func (o *StorageNetAppInitiatorGroup) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *StorageNetAppInitiatorGroup) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *StorageNetAppInitiatorGroup) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *StorageNetAppInitiatorGroup) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetUuid

`func (o *StorageNetAppInitiatorGroup) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StorageNetAppInitiatorGroup) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StorageNetAppInitiatorGroup) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StorageNetAppInitiatorGroup) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetTenant

`func (o *StorageNetAppInitiatorGroup) GetTenant() StorageNetAppStorageVmRelationship`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *StorageNetAppInitiatorGroup) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *StorageNetAppInitiatorGroup) SetTenant(v StorageNetAppStorageVmRelationship)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *StorageNetAppInitiatorGroup) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


