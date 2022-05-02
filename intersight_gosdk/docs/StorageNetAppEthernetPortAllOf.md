# StorageNetAppEthernetPortAllOf

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
**Speed** | Pointer to **int64** | Operational speed of port measured. | [optional] [readonly] 
**State** | Pointer to **string** | State of the port available in storage array. * &#x60;down&#x60; - An inactive port is listed as Down. * &#x60;up&#x60; - An active port is listed as Up. * &#x60;present&#x60; - An active port is listed as present. | [optional] [readonly] [default to "down"]
**Type** | Pointer to **string** | Type of the port available in storage array. * &#x60;LAG&#x60; - Storage port of type lag. * &#x60;physical&#x60; - LIFs can be configured directly on physical ports. * &#x60;VLAN&#x60; - A logical port that receives and sends VLAN-tagged (IEEE 802.1Q standard) traffic. VLAN port characteristics include the VLAN ID for the port. | [optional] [readonly] [default to "LAG"]
**Uuid** | Pointer to **string** | Universally unique identifier of the physical port. | [optional] [readonly] 
**ArrayController** | Pointer to [**StorageNetAppNodeRelationship**](StorageNetAppNodeRelationship.md) |  | [optional] 
**Events** | Pointer to [**[]StorageNetAppEthernetPortEventRelationship**](StorageNetAppEthernetPortEventRelationship.md) | An array of relationships to storageNetAppEthernetPortEvent resources. | [optional] [readonly] 

## Methods

### NewStorageNetAppEthernetPortAllOf

`func NewStorageNetAppEthernetPortAllOf(classId string, objectType string, ) *StorageNetAppEthernetPortAllOf`

NewStorageNetAppEthernetPortAllOf instantiates a new StorageNetAppEthernetPortAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppEthernetPortAllOfWithDefaults

`func NewStorageNetAppEthernetPortAllOfWithDefaults() *StorageNetAppEthernetPortAllOf`

NewStorageNetAppEthernetPortAllOfWithDefaults instantiates a new StorageNetAppEthernetPortAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppEthernetPortAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppEthernetPortAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppEthernetPortAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppEthernetPortAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppEthernetPortAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppEthernetPortAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBroadcastDomainName

`func (o *StorageNetAppEthernetPortAllOf) GetBroadcastDomainName() string`

GetBroadcastDomainName returns the BroadcastDomainName field if non-nil, zero value otherwise.

### GetBroadcastDomainNameOk

`func (o *StorageNetAppEthernetPortAllOf) GetBroadcastDomainNameOk() (*string, bool)`

GetBroadcastDomainNameOk returns a tuple with the BroadcastDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcastDomainName

`func (o *StorageNetAppEthernetPortAllOf) SetBroadcastDomainName(v string)`

SetBroadcastDomainName sets BroadcastDomainName field to given value.

### HasBroadcastDomainName

`func (o *StorageNetAppEthernetPortAllOf) HasBroadcastDomainName() bool`

HasBroadcastDomainName returns a boolean if a field has been set.

### GetEnabled

`func (o *StorageNetAppEthernetPortAllOf) GetEnabled() string`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StorageNetAppEthernetPortAllOf) GetEnabledOk() (*string, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StorageNetAppEthernetPortAllOf) SetEnabled(v string)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *StorageNetAppEthernetPortAllOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMacAddress

`func (o *StorageNetAppEthernetPortAllOf) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *StorageNetAppEthernetPortAllOf) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *StorageNetAppEthernetPortAllOf) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *StorageNetAppEthernetPortAllOf) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMtu

`func (o *StorageNetAppEthernetPortAllOf) GetMtu() int64`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *StorageNetAppEthernetPortAllOf) GetMtuOk() (*int64, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *StorageNetAppEthernetPortAllOf) SetMtu(v int64)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *StorageNetAppEthernetPortAllOf) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetName

`func (o *StorageNetAppEthernetPortAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageNetAppEthernetPortAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageNetAppEthernetPortAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageNetAppEthernetPortAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetAppEthernetPortLag

`func (o *StorageNetAppEthernetPortAllOf) GetNetAppEthernetPortLag() StorageNetAppEthernetPortLag`

GetNetAppEthernetPortLag returns the NetAppEthernetPortLag field if non-nil, zero value otherwise.

### GetNetAppEthernetPortLagOk

`func (o *StorageNetAppEthernetPortAllOf) GetNetAppEthernetPortLagOk() (*StorageNetAppEthernetPortLag, bool)`

GetNetAppEthernetPortLagOk returns a tuple with the NetAppEthernetPortLag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAppEthernetPortLag

`func (o *StorageNetAppEthernetPortAllOf) SetNetAppEthernetPortLag(v StorageNetAppEthernetPortLag)`

SetNetAppEthernetPortLag sets NetAppEthernetPortLag field to given value.

### HasNetAppEthernetPortLag

`func (o *StorageNetAppEthernetPortAllOf) HasNetAppEthernetPortLag() bool`

HasNetAppEthernetPortLag returns a boolean if a field has been set.

### SetNetAppEthernetPortLagNil

`func (o *StorageNetAppEthernetPortAllOf) SetNetAppEthernetPortLagNil(b bool)`

 SetNetAppEthernetPortLagNil sets the value for NetAppEthernetPortLag to be an explicit nil

### UnsetNetAppEthernetPortLag
`func (o *StorageNetAppEthernetPortAllOf) UnsetNetAppEthernetPortLag()`

UnsetNetAppEthernetPortLag ensures that no value is present for NetAppEthernetPortLag, not even an explicit nil
### GetNetAppEthernetPortVlan

`func (o *StorageNetAppEthernetPortAllOf) GetNetAppEthernetPortVlan() StorageNetAppEthernetPortVlan`

GetNetAppEthernetPortVlan returns the NetAppEthernetPortVlan field if non-nil, zero value otherwise.

### GetNetAppEthernetPortVlanOk

`func (o *StorageNetAppEthernetPortAllOf) GetNetAppEthernetPortVlanOk() (*StorageNetAppEthernetPortVlan, bool)`

GetNetAppEthernetPortVlanOk returns a tuple with the NetAppEthernetPortVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAppEthernetPortVlan

`func (o *StorageNetAppEthernetPortAllOf) SetNetAppEthernetPortVlan(v StorageNetAppEthernetPortVlan)`

SetNetAppEthernetPortVlan sets NetAppEthernetPortVlan field to given value.

### HasNetAppEthernetPortVlan

`func (o *StorageNetAppEthernetPortAllOf) HasNetAppEthernetPortVlan() bool`

HasNetAppEthernetPortVlan returns a boolean if a field has been set.

### SetNetAppEthernetPortVlanNil

`func (o *StorageNetAppEthernetPortAllOf) SetNetAppEthernetPortVlanNil(b bool)`

 SetNetAppEthernetPortVlanNil sets the value for NetAppEthernetPortVlan to be an explicit nil

### UnsetNetAppEthernetPortVlan
`func (o *StorageNetAppEthernetPortAllOf) UnsetNetAppEthernetPortVlan()`

UnsetNetAppEthernetPortVlan ensures that no value is present for NetAppEthernetPortVlan, not even an explicit nil
### GetSpeed

`func (o *StorageNetAppEthernetPortAllOf) GetSpeed() int64`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *StorageNetAppEthernetPortAllOf) GetSpeedOk() (*int64, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *StorageNetAppEthernetPortAllOf) SetSpeed(v int64)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *StorageNetAppEthernetPortAllOf) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetState

`func (o *StorageNetAppEthernetPortAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StorageNetAppEthernetPortAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StorageNetAppEthernetPortAllOf) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StorageNetAppEthernetPortAllOf) HasState() bool`

HasState returns a boolean if a field has been set.

### GetType

`func (o *StorageNetAppEthernetPortAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageNetAppEthernetPortAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageNetAppEthernetPortAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StorageNetAppEthernetPortAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *StorageNetAppEthernetPortAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StorageNetAppEthernetPortAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StorageNetAppEthernetPortAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StorageNetAppEthernetPortAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetArrayController

`func (o *StorageNetAppEthernetPortAllOf) GetArrayController() StorageNetAppNodeRelationship`

GetArrayController returns the ArrayController field if non-nil, zero value otherwise.

### GetArrayControllerOk

`func (o *StorageNetAppEthernetPortAllOf) GetArrayControllerOk() (*StorageNetAppNodeRelationship, bool)`

GetArrayControllerOk returns a tuple with the ArrayController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrayController

`func (o *StorageNetAppEthernetPortAllOf) SetArrayController(v StorageNetAppNodeRelationship)`

SetArrayController sets ArrayController field to given value.

### HasArrayController

`func (o *StorageNetAppEthernetPortAllOf) HasArrayController() bool`

HasArrayController returns a boolean if a field has been set.

### GetEvents

`func (o *StorageNetAppEthernetPortAllOf) GetEvents() []StorageNetAppEthernetPortEventRelationship`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *StorageNetAppEthernetPortAllOf) GetEventsOk() (*[]StorageNetAppEthernetPortEventRelationship, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *StorageNetAppEthernetPortAllOf) SetEvents(v []StorageNetAppEthernetPortEventRelationship)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *StorageNetAppEthernetPortAllOf) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *StorageNetAppEthernetPortAllOf) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *StorageNetAppEthernetPortAllOf) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


