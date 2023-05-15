# StorageNetAppFcPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppFcPort"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppFcPort"]
**NodeName** | Pointer to **string** | The node name for the port. | [optional] [readonly] 
**PortStatus** | Pointer to **string** | Status of storage array port. | [optional] [readonly] 
**SpeedConfigured** | Pointer to **string** | The configured speed of the FC port in gigabits per second. | [optional] [readonly] 
**SpeedMaximum** | Pointer to **string** | The maximum speed supported by the FC port in gigabits per second. | [optional] [readonly] 
**State** | Pointer to **string** | State of the port available in storage array. * &#x60;Unknown&#x60; - Default unknown port state. * &#x60;StartUp&#x60; - The port in the storage array is booting up. * &#x60;LinkNotConnected&#x60; - The port has finished initialization, but a link with the fabric is not established. * &#x60;Online&#x60; - The port is initialized and a link with the fabric has been established. * &#x60;LinkDisconnected&#x60; - The link on this port is currently not established. * &#x60;OfflineUser&#x60; - The port is administratively disabled. * &#x60;OfflineSystem&#x60; - The port is set to offline by the system. This happens when the port encounters too many errors. * &#x60;NodeOffline&#x60; - The state information for the port cannot be retrieved. The node is offline or inaccessible. | [optional] [readonly] [default to "Unknown"]
**Uuid** | Pointer to **string** | Universally unique identifier of the FC port. | [optional] [readonly] 
**ArrayController** | Pointer to [**StorageNetAppNodeRelationship**](StorageNetAppNodeRelationship.md) |  | [optional] 
**Events** | Pointer to [**[]StorageNetAppFcPortEventRelationship**](StorageNetAppFcPortEventRelationship.md) | An array of relationships to storageNetAppFcPortEvent resources. | [optional] [readonly] 

## Methods

### NewStorageNetAppFcPort

`func NewStorageNetAppFcPort(classId string, objectType string, ) *StorageNetAppFcPort`

NewStorageNetAppFcPort instantiates a new StorageNetAppFcPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppFcPortWithDefaults

`func NewStorageNetAppFcPortWithDefaults() *StorageNetAppFcPort`

NewStorageNetAppFcPortWithDefaults instantiates a new StorageNetAppFcPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppFcPort) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppFcPort) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppFcPort) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppFcPort) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppFcPort) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppFcPort) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetNodeName

`func (o *StorageNetAppFcPort) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *StorageNetAppFcPort) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *StorageNetAppFcPort) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.

### HasNodeName

`func (o *StorageNetAppFcPort) HasNodeName() bool`

HasNodeName returns a boolean if a field has been set.

### GetPortStatus

`func (o *StorageNetAppFcPort) GetPortStatus() string`

GetPortStatus returns the PortStatus field if non-nil, zero value otherwise.

### GetPortStatusOk

`func (o *StorageNetAppFcPort) GetPortStatusOk() (*string, bool)`

GetPortStatusOk returns a tuple with the PortStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortStatus

`func (o *StorageNetAppFcPort) SetPortStatus(v string)`

SetPortStatus sets PortStatus field to given value.

### HasPortStatus

`func (o *StorageNetAppFcPort) HasPortStatus() bool`

HasPortStatus returns a boolean if a field has been set.

### GetSpeedConfigured

`func (o *StorageNetAppFcPort) GetSpeedConfigured() string`

GetSpeedConfigured returns the SpeedConfigured field if non-nil, zero value otherwise.

### GetSpeedConfiguredOk

`func (o *StorageNetAppFcPort) GetSpeedConfiguredOk() (*string, bool)`

GetSpeedConfiguredOk returns a tuple with the SpeedConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedConfigured

`func (o *StorageNetAppFcPort) SetSpeedConfigured(v string)`

SetSpeedConfigured sets SpeedConfigured field to given value.

### HasSpeedConfigured

`func (o *StorageNetAppFcPort) HasSpeedConfigured() bool`

HasSpeedConfigured returns a boolean if a field has been set.

### GetSpeedMaximum

`func (o *StorageNetAppFcPort) GetSpeedMaximum() string`

GetSpeedMaximum returns the SpeedMaximum field if non-nil, zero value otherwise.

### GetSpeedMaximumOk

`func (o *StorageNetAppFcPort) GetSpeedMaximumOk() (*string, bool)`

GetSpeedMaximumOk returns a tuple with the SpeedMaximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedMaximum

`func (o *StorageNetAppFcPort) SetSpeedMaximum(v string)`

SetSpeedMaximum sets SpeedMaximum field to given value.

### HasSpeedMaximum

`func (o *StorageNetAppFcPort) HasSpeedMaximum() bool`

HasSpeedMaximum returns a boolean if a field has been set.

### GetState

`func (o *StorageNetAppFcPort) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StorageNetAppFcPort) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StorageNetAppFcPort) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StorageNetAppFcPort) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUuid

`func (o *StorageNetAppFcPort) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StorageNetAppFcPort) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StorageNetAppFcPort) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StorageNetAppFcPort) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetArrayController

`func (o *StorageNetAppFcPort) GetArrayController() StorageNetAppNodeRelationship`

GetArrayController returns the ArrayController field if non-nil, zero value otherwise.

### GetArrayControllerOk

`func (o *StorageNetAppFcPort) GetArrayControllerOk() (*StorageNetAppNodeRelationship, bool)`

GetArrayControllerOk returns a tuple with the ArrayController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayController

`func (o *StorageNetAppFcPort) SetArrayController(v StorageNetAppNodeRelationship)`

SetArrayController sets ArrayController field to given value.

### HasArrayController

`func (o *StorageNetAppFcPort) HasArrayController() bool`

HasArrayController returns a boolean if a field has been set.

### GetEvents

`func (o *StorageNetAppFcPort) GetEvents() []StorageNetAppFcPortEventRelationship`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *StorageNetAppFcPort) GetEventsOk() (*[]StorageNetAppFcPortEventRelationship, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *StorageNetAppFcPort) SetEvents(v []StorageNetAppFcPortEventRelationship)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *StorageNetAppFcPort) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *StorageNetAppFcPort) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *StorageNetAppFcPort) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


