# StorageNetAppDataIpInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppDataIpInterface"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppDataIpInterface"]
**ArrayController** | Pointer to [**NullableStorageNetAppNodeRelationship**](StorageNetAppNodeRelationship.md) |  | [optional] 
**Events** | Pointer to [**[]StorageNetAppDataIpInterfaceEventRelationship**](StorageNetAppDataIpInterfaceEventRelationship.md) | An array of relationships to storageNetAppDataIpInterfaceEvent resources. | [optional] [readonly] 
**NetAppEthernetPort** | Pointer to [**NullableStorageNetAppEthernetPortRelationship**](StorageNetAppEthernetPortRelationship.md) |  | [optional] 
**Tenant** | Pointer to [**NullableStorageNetAppStorageVmRelationship**](StorageNetAppStorageVmRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppDataIpInterface

`func NewStorageNetAppDataIpInterface(classId string, objectType string, ) *StorageNetAppDataIpInterface`

NewStorageNetAppDataIpInterface instantiates a new StorageNetAppDataIpInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppDataIpInterfaceWithDefaults

`func NewStorageNetAppDataIpInterfaceWithDefaults() *StorageNetAppDataIpInterface`

NewStorageNetAppDataIpInterfaceWithDefaults instantiates a new StorageNetAppDataIpInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppDataIpInterface) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppDataIpInterface) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppDataIpInterface) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppDataIpInterface) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppDataIpInterface) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppDataIpInterface) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetArrayController

`func (o *StorageNetAppDataIpInterface) GetArrayController() StorageNetAppNodeRelationship`

GetArrayController returns the ArrayController field if non-nil, zero value otherwise.

### GetArrayControllerOk

`func (o *StorageNetAppDataIpInterface) GetArrayControllerOk() (*StorageNetAppNodeRelationship, bool)`

GetArrayControllerOk returns a tuple with the ArrayController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayController

`func (o *StorageNetAppDataIpInterface) SetArrayController(v StorageNetAppNodeRelationship)`

SetArrayController sets ArrayController field to given value.

### HasArrayController

`func (o *StorageNetAppDataIpInterface) HasArrayController() bool`

HasArrayController returns a boolean if a field has been set.

### SetArrayControllerNil

`func (o *StorageNetAppDataIpInterface) SetArrayControllerNil(b bool)`

 SetArrayControllerNil sets the value for ArrayController to be an explicit nil

### UnsetArrayController
`func (o *StorageNetAppDataIpInterface) UnsetArrayController()`

UnsetArrayController ensures that no value is present for ArrayController, not even an explicit nil
### GetEvents

`func (o *StorageNetAppDataIpInterface) GetEvents() []StorageNetAppDataIpInterfaceEventRelationship`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *StorageNetAppDataIpInterface) GetEventsOk() (*[]StorageNetAppDataIpInterfaceEventRelationship, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *StorageNetAppDataIpInterface) SetEvents(v []StorageNetAppDataIpInterfaceEventRelationship)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *StorageNetAppDataIpInterface) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *StorageNetAppDataIpInterface) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *StorageNetAppDataIpInterface) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil
### GetNetAppEthernetPort

`func (o *StorageNetAppDataIpInterface) GetNetAppEthernetPort() StorageNetAppEthernetPortRelationship`

GetNetAppEthernetPort returns the NetAppEthernetPort field if non-nil, zero value otherwise.

### GetNetAppEthernetPortOk

`func (o *StorageNetAppDataIpInterface) GetNetAppEthernetPortOk() (*StorageNetAppEthernetPortRelationship, bool)`

GetNetAppEthernetPortOk returns a tuple with the NetAppEthernetPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAppEthernetPort

`func (o *StorageNetAppDataIpInterface) SetNetAppEthernetPort(v StorageNetAppEthernetPortRelationship)`

SetNetAppEthernetPort sets NetAppEthernetPort field to given value.

### HasNetAppEthernetPort

`func (o *StorageNetAppDataIpInterface) HasNetAppEthernetPort() bool`

HasNetAppEthernetPort returns a boolean if a field has been set.

### SetNetAppEthernetPortNil

`func (o *StorageNetAppDataIpInterface) SetNetAppEthernetPortNil(b bool)`

 SetNetAppEthernetPortNil sets the value for NetAppEthernetPort to be an explicit nil

### UnsetNetAppEthernetPort
`func (o *StorageNetAppDataIpInterface) UnsetNetAppEthernetPort()`

UnsetNetAppEthernetPort ensures that no value is present for NetAppEthernetPort, not even an explicit nil
### GetTenant

`func (o *StorageNetAppDataIpInterface) GetTenant() StorageNetAppStorageVmRelationship`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *StorageNetAppDataIpInterface) GetTenantOk() (*StorageNetAppStorageVmRelationship, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *StorageNetAppDataIpInterface) SetTenant(v StorageNetAppStorageVmRelationship)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *StorageNetAppDataIpInterface) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *StorageNetAppDataIpInterface) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *StorageNetAppDataIpInterface) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


