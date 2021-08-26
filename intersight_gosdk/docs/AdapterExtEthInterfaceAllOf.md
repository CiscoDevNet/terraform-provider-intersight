# AdapterExtEthInterfaceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "adapter.ExtEthInterface"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "adapter.ExtEthInterface"]
**AdminState** | Pointer to **string** | Admin configured state of an External Ethernet Interface. | [optional] [readonly] 
**EpDn** | Pointer to **string** | Endpoint Config DN of an External Ethernet Interface. | [optional] [readonly] 
**ExtEthInterfaceId** | Pointer to **string** | Unique Identifier for an External Ethernet Interface within the adapter object. | [optional] [readonly] 
**InterfaceType** | Pointer to **string** | Type of an External Ethernet Interface. | [optional] [readonly] 
**MacAddress** | Pointer to **string** | MAC address of an External Ethernet Interface. | [optional] [readonly] 
**PeerAggrPortId** | Pointer to **int64** | Peer Aggregate Port Id attached to an External Ethernet Interface. | [optional] [readonly] 
**PeerDn** | Pointer to **string** | DN of peer end-point attached to an External Ethernet Interface. | [optional] [readonly] 
**PeerPortId** | Pointer to **int64** | Peer Port Id attached to an External Ethernet Interface. | [optional] [readonly] 
**PeerSlotId** | Pointer to **int64** | Peer Slot Id attached to an External Ethernet Interface. | [optional] [readonly] 
**SwitchId** | Pointer to **string** | SwitchId attached to an External Ethernet Interface. | [optional] [readonly] 
**AdapterUnit** | Pointer to [**AdapterUnitRelationship**](AdapterUnitRelationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewAdapterExtEthInterfaceAllOf

`func NewAdapterExtEthInterfaceAllOf(classId string, objectType string, ) *AdapterExtEthInterfaceAllOf`

NewAdapterExtEthInterfaceAllOf instantiates a new AdapterExtEthInterfaceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdapterExtEthInterfaceAllOfWithDefaults

`func NewAdapterExtEthInterfaceAllOfWithDefaults() *AdapterExtEthInterfaceAllOf`

NewAdapterExtEthInterfaceAllOfWithDefaults instantiates a new AdapterExtEthInterfaceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AdapterExtEthInterfaceAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AdapterExtEthInterfaceAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AdapterExtEthInterfaceAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AdapterExtEthInterfaceAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AdapterExtEthInterfaceAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AdapterExtEthInterfaceAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminState

`func (o *AdapterExtEthInterfaceAllOf) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *AdapterExtEthInterfaceAllOf) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *AdapterExtEthInterfaceAllOf) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *AdapterExtEthInterfaceAllOf) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetEpDn

`func (o *AdapterExtEthInterfaceAllOf) GetEpDn() string`

GetEpDn returns the EpDn field if non-nil, zero value otherwise.

### GetEpDnOk

`func (o *AdapterExtEthInterfaceAllOf) GetEpDnOk() (*string, bool)`

GetEpDnOk returns a tuple with the EpDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpDn

`func (o *AdapterExtEthInterfaceAllOf) SetEpDn(v string)`

SetEpDn sets EpDn field to given value.

### HasEpDn

`func (o *AdapterExtEthInterfaceAllOf) HasEpDn() bool`

HasEpDn returns a boolean if a field has been set.

### GetExtEthInterfaceId

`func (o *AdapterExtEthInterfaceAllOf) GetExtEthInterfaceId() string`

GetExtEthInterfaceId returns the ExtEthInterfaceId field if non-nil, zero value otherwise.

### GetExtEthInterfaceIdOk

`func (o *AdapterExtEthInterfaceAllOf) GetExtEthInterfaceIdOk() (*string, bool)`

GetExtEthInterfaceIdOk returns a tuple with the ExtEthInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtEthInterfaceId

`func (o *AdapterExtEthInterfaceAllOf) SetExtEthInterfaceId(v string)`

SetExtEthInterfaceId sets ExtEthInterfaceId field to given value.

### HasExtEthInterfaceId

`func (o *AdapterExtEthInterfaceAllOf) HasExtEthInterfaceId() bool`

HasExtEthInterfaceId returns a boolean if a field has been set.

### GetInterfaceType

`func (o *AdapterExtEthInterfaceAllOf) GetInterfaceType() string`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *AdapterExtEthInterfaceAllOf) GetInterfaceTypeOk() (*string, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *AdapterExtEthInterfaceAllOf) SetInterfaceType(v string)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *AdapterExtEthInterfaceAllOf) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### GetMacAddress

`func (o *AdapterExtEthInterfaceAllOf) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *AdapterExtEthInterfaceAllOf) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *AdapterExtEthInterfaceAllOf) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *AdapterExtEthInterfaceAllOf) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetPeerAggrPortId

`func (o *AdapterExtEthInterfaceAllOf) GetPeerAggrPortId() int64`

GetPeerAggrPortId returns the PeerAggrPortId field if non-nil, zero value otherwise.

### GetPeerAggrPortIdOk

`func (o *AdapterExtEthInterfaceAllOf) GetPeerAggrPortIdOk() (*int64, bool)`

GetPeerAggrPortIdOk returns a tuple with the PeerAggrPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerAggrPortId

`func (o *AdapterExtEthInterfaceAllOf) SetPeerAggrPortId(v int64)`

SetPeerAggrPortId sets PeerAggrPortId field to given value.

### HasPeerAggrPortId

`func (o *AdapterExtEthInterfaceAllOf) HasPeerAggrPortId() bool`

HasPeerAggrPortId returns a boolean if a field has been set.

### GetPeerDn

`func (o *AdapterExtEthInterfaceAllOf) GetPeerDn() string`

GetPeerDn returns the PeerDn field if non-nil, zero value otherwise.

### GetPeerDnOk

`func (o *AdapterExtEthInterfaceAllOf) GetPeerDnOk() (*string, bool)`

GetPeerDnOk returns a tuple with the PeerDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerDn

`func (o *AdapterExtEthInterfaceAllOf) SetPeerDn(v string)`

SetPeerDn sets PeerDn field to given value.

### HasPeerDn

`func (o *AdapterExtEthInterfaceAllOf) HasPeerDn() bool`

HasPeerDn returns a boolean if a field has been set.

### GetPeerPortId

`func (o *AdapterExtEthInterfaceAllOf) GetPeerPortId() int64`

GetPeerPortId returns the PeerPortId field if non-nil, zero value otherwise.

### GetPeerPortIdOk

`func (o *AdapterExtEthInterfaceAllOf) GetPeerPortIdOk() (*int64, bool)`

GetPeerPortIdOk returns a tuple with the PeerPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerPortId

`func (o *AdapterExtEthInterfaceAllOf) SetPeerPortId(v int64)`

SetPeerPortId sets PeerPortId field to given value.

### HasPeerPortId

`func (o *AdapterExtEthInterfaceAllOf) HasPeerPortId() bool`

HasPeerPortId returns a boolean if a field has been set.

### GetPeerSlotId

`func (o *AdapterExtEthInterfaceAllOf) GetPeerSlotId() int64`

GetPeerSlotId returns the PeerSlotId field if non-nil, zero value otherwise.

### GetPeerSlotIdOk

`func (o *AdapterExtEthInterfaceAllOf) GetPeerSlotIdOk() (*int64, bool)`

GetPeerSlotIdOk returns a tuple with the PeerSlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerSlotId

`func (o *AdapterExtEthInterfaceAllOf) SetPeerSlotId(v int64)`

SetPeerSlotId sets PeerSlotId field to given value.

### HasPeerSlotId

`func (o *AdapterExtEthInterfaceAllOf) HasPeerSlotId() bool`

HasPeerSlotId returns a boolean if a field has been set.

### GetSwitchId

`func (o *AdapterExtEthInterfaceAllOf) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *AdapterExtEthInterfaceAllOf) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *AdapterExtEthInterfaceAllOf) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *AdapterExtEthInterfaceAllOf) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetAdapterUnit

`func (o *AdapterExtEthInterfaceAllOf) GetAdapterUnit() AdapterUnitRelationship`

GetAdapterUnit returns the AdapterUnit field if non-nil, zero value otherwise.

### GetAdapterUnitOk

`func (o *AdapterExtEthInterfaceAllOf) GetAdapterUnitOk() (*AdapterUnitRelationship, bool)`

GetAdapterUnitOk returns a tuple with the AdapterUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterUnit

`func (o *AdapterExtEthInterfaceAllOf) SetAdapterUnit(v AdapterUnitRelationship)`

SetAdapterUnit sets AdapterUnit field to given value.

### HasAdapterUnit

`func (o *AdapterExtEthInterfaceAllOf) HasAdapterUnit() bool`

HasAdapterUnit returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *AdapterExtEthInterfaceAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *AdapterExtEthInterfaceAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *AdapterExtEthInterfaceAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *AdapterExtEthInterfaceAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *AdapterExtEthInterfaceAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *AdapterExtEthInterfaceAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *AdapterExtEthInterfaceAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *AdapterExtEthInterfaceAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


