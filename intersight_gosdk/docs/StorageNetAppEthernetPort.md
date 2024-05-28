# StorageNetAppEthernetPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppEthernetPort"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppEthernetPort"]
**BroadcastDomainName** | Pointer to **string** | Name of the broadcast domain, scoped to its IPspace. | [optional] [readonly] 
**Enabled** | Pointer to **string** | Status of port to determine if its enabled or not. | [optional] [readonly] 
**MacAddress** | Pointer to **string** | MAC address of the port available in storage array. | [optional] [readonly] 
**Mtu** | Pointer to **int64** | Maximum transmission unit of the physical port available in storage array. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the port available in storage array. | [optional] [readonly] 
**NetAppEthernetPortLag** | Pointer to [**NullableStorageNetAppEthernetPortLag**](StorageNetAppEthernetPortLag.md) |  | [optional] 
**NetAppEthernetPortVlan** | Pointer to [**NullableStorageNetAppEthernetPortVlan**](StorageNetAppEthernetPortVlan.md) |  | [optional] 
**NodeName** | Pointer to **string** | The node name for the port. | [optional] [readonly] 
**PortState** | Pointer to **string** | State of the port available in storage array. * &#x60;Down&#x60; - An inactive port is listed as Down. * &#x60;Up&#x60; - An active port is listed as Up. * &#x60;Degraded&#x60; - An active port that is Up but unhealthy. | [optional] [readonly] [default to "Down"]
**Speed** | Pointer to **int64** | Operational speed of port measured. | [optional] [readonly] 
**State** | Pointer to **string** | State of the port available in storage array. * &#x60;down&#x60; - An inactive port is listed as Down. * &#x60;up&#x60; - An active port is listed as Up. * &#x60;present&#x60; - An active port is listed as present. | [optional] [readonly] [default to "down"]
**Type** | Pointer to **string** | Type of the port available in storage array. * &#x60;LAG&#x60; - Storage port of type lag. * &#x60;physical&#x60; - LIFs can be configured directly on physical ports. * &#x60;VLAN&#x60; - A logical port that receives and sends VLAN-tagged (IEEE 802.1Q standard) traffic. VLAN port characteristics include the VLAN ID for the port. | [optional] [readonly] [default to "LAG"]
**Uuid** | Pointer to **string** | Universally unique identifier of the physical port. | [optional] [readonly] 
**ArrayController** | Pointer to [**NullableStorageNetAppNodeRelationship**](StorageNetAppNodeRelationship.md) |  | [optional] 
**Events** | Pointer to [**[]StorageNetAppEthernetPortEventRelationship**](StorageNetAppEthernetPortEventRelationship.md) | An array of relationships to storageNetAppEthernetPortEvent resources. | [optional] [readonly] 

## Methods

### NewStorageNetAppEthernetPort

`func NewStorageNetAppEthernetPort(classId string, objectType string, ) *StorageNetAppEthernetPort`

NewStorageNetAppEthernetPort instantiates a new StorageNetAppEthernetPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppEthernetPortWithDefaults

`func NewStorageNetAppEthernetPortWithDefaults() *StorageNetAppEthernetPort`

NewStorageNetAppEthernetPortWithDefaults instantiates a new StorageNetAppEthernetPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppEthernetPort) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppEthernetPort) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppEthernetPort) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppEthernetPort) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppEthernetPort) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppEthernetPort) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBroadcastDomainName

`func (o *StorageNetAppEthernetPort) GetBroadcastDomainName() string`

GetBroadcastDomainName returns the BroadcastDomainName field if non-nil, zero value otherwise.

### GetBroadcastDomainNameOk

`func (o *StorageNetAppEthernetPort) GetBroadcastDomainNameOk() (*string, bool)`

GetBroadcastDomainNameOk returns a tuple with the BroadcastDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcastDomainName

`func (o *StorageNetAppEthernetPort) SetBroadcastDomainName(v string)`

SetBroadcastDomainName sets BroadcastDomainName field to given value.

### HasBroadcastDomainName

`func (o *StorageNetAppEthernetPort) HasBroadcastDomainName() bool`

HasBroadcastDomainName returns a boolean if a field has been set.

### GetEnabled

`func (o *StorageNetAppEthernetPort) GetEnabled() string`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StorageNetAppEthernetPort) GetEnabledOk() (*string, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StorageNetAppEthernetPort) SetEnabled(v string)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *StorageNetAppEthernetPort) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMacAddress

`func (o *StorageNetAppEthernetPort) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *StorageNetAppEthernetPort) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *StorageNetAppEthernetPort) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *StorageNetAppEthernetPort) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMtu

`func (o *StorageNetAppEthernetPort) GetMtu() int64`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *StorageNetAppEthernetPort) GetMtuOk() (*int64, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *StorageNetAppEthernetPort) SetMtu(v int64)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *StorageNetAppEthernetPort) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetName

`func (o *StorageNetAppEthernetPort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageNetAppEthernetPort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageNetAppEthernetPort) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageNetAppEthernetPort) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetAppEthernetPortLag

`func (o *StorageNetAppEthernetPort) GetNetAppEthernetPortLag() StorageNetAppEthernetPortLag`

GetNetAppEthernetPortLag returns the NetAppEthernetPortLag field if non-nil, zero value otherwise.

### GetNetAppEthernetPortLagOk

`func (o *StorageNetAppEthernetPort) GetNetAppEthernetPortLagOk() (*StorageNetAppEthernetPortLag, bool)`

GetNetAppEthernetPortLagOk returns a tuple with the NetAppEthernetPortLag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAppEthernetPortLag

`func (o *StorageNetAppEthernetPort) SetNetAppEthernetPortLag(v StorageNetAppEthernetPortLag)`

SetNetAppEthernetPortLag sets NetAppEthernetPortLag field to given value.

### HasNetAppEthernetPortLag

`func (o *StorageNetAppEthernetPort) HasNetAppEthernetPortLag() bool`

HasNetAppEthernetPortLag returns a boolean if a field has been set.

### SetNetAppEthernetPortLagNil

`func (o *StorageNetAppEthernetPort) SetNetAppEthernetPortLagNil(b bool)`

 SetNetAppEthernetPortLagNil sets the value for NetAppEthernetPortLag to be an explicit nil

### UnsetNetAppEthernetPortLag
`func (o *StorageNetAppEthernetPort) UnsetNetAppEthernetPortLag()`

UnsetNetAppEthernetPortLag ensures that no value is present for NetAppEthernetPortLag, not even an explicit nil
### GetNetAppEthernetPortVlan

`func (o *StorageNetAppEthernetPort) GetNetAppEthernetPortVlan() StorageNetAppEthernetPortVlan`

GetNetAppEthernetPortVlan returns the NetAppEthernetPortVlan field if non-nil, zero value otherwise.

### GetNetAppEthernetPortVlanOk

`func (o *StorageNetAppEthernetPort) GetNetAppEthernetPortVlanOk() (*StorageNetAppEthernetPortVlan, bool)`

GetNetAppEthernetPortVlanOk returns a tuple with the NetAppEthernetPortVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAppEthernetPortVlan

`func (o *StorageNetAppEthernetPort) SetNetAppEthernetPortVlan(v StorageNetAppEthernetPortVlan)`

SetNetAppEthernetPortVlan sets NetAppEthernetPortVlan field to given value.

### HasNetAppEthernetPortVlan

`func (o *StorageNetAppEthernetPort) HasNetAppEthernetPortVlan() bool`

HasNetAppEthernetPortVlan returns a boolean if a field has been set.

### SetNetAppEthernetPortVlanNil

`func (o *StorageNetAppEthernetPort) SetNetAppEthernetPortVlanNil(b bool)`

 SetNetAppEthernetPortVlanNil sets the value for NetAppEthernetPortVlan to be an explicit nil

### UnsetNetAppEthernetPortVlan
`func (o *StorageNetAppEthernetPort) UnsetNetAppEthernetPortVlan()`

UnsetNetAppEthernetPortVlan ensures that no value is present for NetAppEthernetPortVlan, not even an explicit nil
### GetNodeName

`func (o *StorageNetAppEthernetPort) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *StorageNetAppEthernetPort) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *StorageNetAppEthernetPort) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.

### HasNodeName

`func (o *StorageNetAppEthernetPort) HasNodeName() bool`

HasNodeName returns a boolean if a field has been set.

### GetPortState

`func (o *StorageNetAppEthernetPort) GetPortState() string`

GetPortState returns the PortState field if non-nil, zero value otherwise.

### GetPortStateOk

`func (o *StorageNetAppEthernetPort) GetPortStateOk() (*string, bool)`

GetPortStateOk returns a tuple with the PortState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortState

`func (o *StorageNetAppEthernetPort) SetPortState(v string)`

SetPortState sets PortState field to given value.

### HasPortState

`func (o *StorageNetAppEthernetPort) HasPortState() bool`

HasPortState returns a boolean if a field has been set.

### GetSpeed

`func (o *StorageNetAppEthernetPort) GetSpeed() int64`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *StorageNetAppEthernetPort) GetSpeedOk() (*int64, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *StorageNetAppEthernetPort) SetSpeed(v int64)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *StorageNetAppEthernetPort) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetState

`func (o *StorageNetAppEthernetPort) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StorageNetAppEthernetPort) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StorageNetAppEthernetPort) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StorageNetAppEthernetPort) HasState() bool`

HasState returns a boolean if a field has been set.

### GetType

`func (o *StorageNetAppEthernetPort) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageNetAppEthernetPort) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageNetAppEthernetPort) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StorageNetAppEthernetPort) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *StorageNetAppEthernetPort) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StorageNetAppEthernetPort) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StorageNetAppEthernetPort) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StorageNetAppEthernetPort) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetArrayController

`func (o *StorageNetAppEthernetPort) GetArrayController() StorageNetAppNodeRelationship`

GetArrayController returns the ArrayController field if non-nil, zero value otherwise.

### GetArrayControllerOk

`func (o *StorageNetAppEthernetPort) GetArrayControllerOk() (*StorageNetAppNodeRelationship, bool)`

GetArrayControllerOk returns a tuple with the ArrayController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayController

`func (o *StorageNetAppEthernetPort) SetArrayController(v StorageNetAppNodeRelationship)`

SetArrayController sets ArrayController field to given value.

### HasArrayController

`func (o *StorageNetAppEthernetPort) HasArrayController() bool`

HasArrayController returns a boolean if a field has been set.

### SetArrayControllerNil

`func (o *StorageNetAppEthernetPort) SetArrayControllerNil(b bool)`

 SetArrayControllerNil sets the value for ArrayController to be an explicit nil

### UnsetArrayController
`func (o *StorageNetAppEthernetPort) UnsetArrayController()`

UnsetArrayController ensures that no value is present for ArrayController, not even an explicit nil
### GetEvents

`func (o *StorageNetAppEthernetPort) GetEvents() []StorageNetAppEthernetPortEventRelationship`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *StorageNetAppEthernetPort) GetEventsOk() (*[]StorageNetAppEthernetPortEventRelationship, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *StorageNetAppEthernetPort) SetEvents(v []StorageNetAppEthernetPortEventRelationship)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *StorageNetAppEthernetPort) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *StorageNetAppEthernetPort) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *StorageNetAppEthernetPort) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


