# StorageNetAppNonDataIpInterfaceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppNonDataIpInterface"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppNonDataIpInterface"]
**Array** | Pointer to [**StorageNetAppClusterRelationship**](StorageNetAppClusterRelationship.md) |  | [optional] 
**ArrayController** | Pointer to [**StorageNetAppNodeRelationship**](StorageNetAppNodeRelationship.md) |  | [optional] 
**Events** | Pointer to [**[]StorageNetAppNonDataIpInterfaceEventRelationship**](StorageNetAppNonDataIpInterfaceEventRelationship.md) | An array of relationships to storageNetAppNonDataIpInterfaceEvent resources. | [optional] [readonly] 
**NetAppEthernetPort** | Pointer to [**StorageNetAppEthernetPortRelationship**](StorageNetAppEthernetPortRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppNonDataIpInterfaceAllOf

`func NewStorageNetAppNonDataIpInterfaceAllOf(classId string, objectType string, ) *StorageNetAppNonDataIpInterfaceAllOf`

NewStorageNetAppNonDataIpInterfaceAllOf instantiates a new StorageNetAppNonDataIpInterfaceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppNonDataIpInterfaceAllOfWithDefaults

`func NewStorageNetAppNonDataIpInterfaceAllOfWithDefaults() *StorageNetAppNonDataIpInterfaceAllOf`

NewStorageNetAppNonDataIpInterfaceAllOfWithDefaults instantiates a new StorageNetAppNonDataIpInterfaceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppNonDataIpInterfaceAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppNonDataIpInterfaceAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppNonDataIpInterfaceAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppNonDataIpInterfaceAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppNonDataIpInterfaceAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppNonDataIpInterfaceAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetArray

`func (o *StorageNetAppNonDataIpInterfaceAllOf) GetArray() StorageNetAppClusterRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageNetAppNonDataIpInterfaceAllOf) GetArrayOk() (*StorageNetAppClusterRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageNetAppNonDataIpInterfaceAllOf) SetArray(v StorageNetAppClusterRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageNetAppNonDataIpInterfaceAllOf) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetArrayController

`func (o *StorageNetAppNonDataIpInterfaceAllOf) GetArrayController() StorageNetAppNodeRelationship`

GetArrayController returns the ArrayController field if non-nil, zero value otherwise.

### GetArrayControllerOk

`func (o *StorageNetAppNonDataIpInterfaceAllOf) GetArrayControllerOk() (*StorageNetAppNodeRelationship, bool)`

GetArrayControllerOk returns a tuple with the ArrayController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayController

`func (o *StorageNetAppNonDataIpInterfaceAllOf) SetArrayController(v StorageNetAppNodeRelationship)`

SetArrayController sets ArrayController field to given value.

### HasArrayController

`func (o *StorageNetAppNonDataIpInterfaceAllOf) HasArrayController() bool`

HasArrayController returns a boolean if a field has been set.

### GetEvents

`func (o *StorageNetAppNonDataIpInterfaceAllOf) GetEvents() []StorageNetAppNonDataIpInterfaceEventRelationship`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *StorageNetAppNonDataIpInterfaceAllOf) GetEventsOk() (*[]StorageNetAppNonDataIpInterfaceEventRelationship, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *StorageNetAppNonDataIpInterfaceAllOf) SetEvents(v []StorageNetAppNonDataIpInterfaceEventRelationship)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *StorageNetAppNonDataIpInterfaceAllOf) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *StorageNetAppNonDataIpInterfaceAllOf) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *StorageNetAppNonDataIpInterfaceAllOf) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil
### GetNetAppEthernetPort

`func (o *StorageNetAppNonDataIpInterfaceAllOf) GetNetAppEthernetPort() StorageNetAppEthernetPortRelationship`

GetNetAppEthernetPort returns the NetAppEthernetPort field if non-nil, zero value otherwise.

### GetNetAppEthernetPortOk

`func (o *StorageNetAppNonDataIpInterfaceAllOf) GetNetAppEthernetPortOk() (*StorageNetAppEthernetPortRelationship, bool)`

GetNetAppEthernetPortOk returns a tuple with the NetAppEthernetPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAppEthernetPort

`func (o *StorageNetAppNonDataIpInterfaceAllOf) SetNetAppEthernetPort(v StorageNetAppEthernetPortRelationship)`

SetNetAppEthernetPort sets NetAppEthernetPort field to given value.

### HasNetAppEthernetPort

`func (o *StorageNetAppNonDataIpInterfaceAllOf) HasNetAppEthernetPort() bool`

HasNetAppEthernetPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


