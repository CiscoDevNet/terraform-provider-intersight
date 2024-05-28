# StorageNetAppNonDataIpInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppNonDataIpInterface"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppNonDataIpInterface"]
**Array** | Pointer to [**NullableStorageNetAppClusterRelationship**](StorageNetAppClusterRelationship.md) |  | [optional] 
**ArrayController** | Pointer to [**NullableStorageNetAppNodeRelationship**](StorageNetAppNodeRelationship.md) |  | [optional] 
**Events** | Pointer to [**[]StorageNetAppNonDataIpInterfaceEventRelationship**](StorageNetAppNonDataIpInterfaceEventRelationship.md) | An array of relationships to storageNetAppNonDataIpInterfaceEvent resources. | [optional] [readonly] 
**NetAppEthernetPort** | Pointer to [**NullableStorageNetAppEthernetPortRelationship**](StorageNetAppEthernetPortRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppNonDataIpInterface

`func NewStorageNetAppNonDataIpInterface(classId string, objectType string, ) *StorageNetAppNonDataIpInterface`

NewStorageNetAppNonDataIpInterface instantiates a new StorageNetAppNonDataIpInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppNonDataIpInterfaceWithDefaults

`func NewStorageNetAppNonDataIpInterfaceWithDefaults() *StorageNetAppNonDataIpInterface`

NewStorageNetAppNonDataIpInterfaceWithDefaults instantiates a new StorageNetAppNonDataIpInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppNonDataIpInterface) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppNonDataIpInterface) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppNonDataIpInterface) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppNonDataIpInterface) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppNonDataIpInterface) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppNonDataIpInterface) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetArray

`func (o *StorageNetAppNonDataIpInterface) GetArray() StorageNetAppClusterRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageNetAppNonDataIpInterface) GetArrayOk() (*StorageNetAppClusterRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageNetAppNonDataIpInterface) SetArray(v StorageNetAppClusterRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageNetAppNonDataIpInterface) HasArray() bool`

HasArray returns a boolean if a field has been set.

### SetArrayNil

`func (o *StorageNetAppNonDataIpInterface) SetArrayNil(b bool)`

 SetArrayNil sets the value for Array to be an explicit nil

### UnsetArray
`func (o *StorageNetAppNonDataIpInterface) UnsetArray()`

UnsetArray ensures that no value is present for Array, not even an explicit nil
### GetArrayController

`func (o *StorageNetAppNonDataIpInterface) GetArrayController() StorageNetAppNodeRelationship`

GetArrayController returns the ArrayController field if non-nil, zero value otherwise.

### GetArrayControllerOk

`func (o *StorageNetAppNonDataIpInterface) GetArrayControllerOk() (*StorageNetAppNodeRelationship, bool)`

GetArrayControllerOk returns a tuple with the ArrayController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayController

`func (o *StorageNetAppNonDataIpInterface) SetArrayController(v StorageNetAppNodeRelationship)`

SetArrayController sets ArrayController field to given value.

### HasArrayController

`func (o *StorageNetAppNonDataIpInterface) HasArrayController() bool`

HasArrayController returns a boolean if a field has been set.

### SetArrayControllerNil

`func (o *StorageNetAppNonDataIpInterface) SetArrayControllerNil(b bool)`

 SetArrayControllerNil sets the value for ArrayController to be an explicit nil

### UnsetArrayController
`func (o *StorageNetAppNonDataIpInterface) UnsetArrayController()`

UnsetArrayController ensures that no value is present for ArrayController, not even an explicit nil
### GetEvents

`func (o *StorageNetAppNonDataIpInterface) GetEvents() []StorageNetAppNonDataIpInterfaceEventRelationship`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *StorageNetAppNonDataIpInterface) GetEventsOk() (*[]StorageNetAppNonDataIpInterfaceEventRelationship, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *StorageNetAppNonDataIpInterface) SetEvents(v []StorageNetAppNonDataIpInterfaceEventRelationship)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *StorageNetAppNonDataIpInterface) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *StorageNetAppNonDataIpInterface) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *StorageNetAppNonDataIpInterface) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil
### GetNetAppEthernetPort

`func (o *StorageNetAppNonDataIpInterface) GetNetAppEthernetPort() StorageNetAppEthernetPortRelationship`

GetNetAppEthernetPort returns the NetAppEthernetPort field if non-nil, zero value otherwise.

### GetNetAppEthernetPortOk

`func (o *StorageNetAppNonDataIpInterface) GetNetAppEthernetPortOk() (*StorageNetAppEthernetPortRelationship, bool)`

GetNetAppEthernetPortOk returns a tuple with the NetAppEthernetPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAppEthernetPort

`func (o *StorageNetAppNonDataIpInterface) SetNetAppEthernetPort(v StorageNetAppEthernetPortRelationship)`

SetNetAppEthernetPort sets NetAppEthernetPort field to given value.

### HasNetAppEthernetPort

`func (o *StorageNetAppNonDataIpInterface) HasNetAppEthernetPort() bool`

HasNetAppEthernetPort returns a boolean if a field has been set.

### SetNetAppEthernetPortNil

`func (o *StorageNetAppNonDataIpInterface) SetNetAppEthernetPortNil(b bool)`

 SetNetAppEthernetPortNil sets the value for NetAppEthernetPort to be an explicit nil

### UnsetNetAppEthernetPort
`func (o *StorageNetAppNonDataIpInterface) UnsetNetAppEthernetPort()`

UnsetNetAppEthernetPort ensures that no value is present for NetAppEthernetPort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


