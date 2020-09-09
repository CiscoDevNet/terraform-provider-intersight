# AdapterHostFcInterfaceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminState** | Pointer to **string** | Admin Configured State of Host Fibre Channel Interface. | [optional] [readonly] 
**EpDn** | Pointer to **string** | The Endpoint Config Dn of the Host Fibre Channel Interface. | [optional] [readonly] 
**HostFcInterfaceId** | Pointer to **int64** | Identifier of Host Fibre Channel Interface. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of Host Fibre Channel Interface. | [optional] [readonly] 
**OperState** | Pointer to **string** | Operational State of Host Fibre Channel Interface. | [optional] [readonly] 
**Operability** | Pointer to **string** | Operability status of Host Fibre Channel Interface. | [optional] [readonly] 
**OriginalWwnn** | Pointer to **string** | The uniquely distinguishable factory default  World Wide Node Name of the Host. | [optional] [readonly] 
**OriginalWwpn** | Pointer to **string** | The uniquely distinguishable factory default World Wide Port Name of the Host Fibre Channel Interface. | [optional] [readonly] 
**PeerDn** | Pointer to **string** | PeerPort Dn of Host Fibre Channel Interface. | [optional] [readonly] 
**Wwnn** | Pointer to **string** | The uniquely distinguishable user configured World Wide Node Name of the Host. | [optional] [readonly] 
**Wwpn** | Pointer to **string** | The uniquely distinguishable user configured World Wide Port Name of the Host Fibre Channel Interface. | [optional] [readonly] 
**AdapterUnit** | Pointer to [**AdapterUnitRelationship**](adapter.Unit.Relationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewAdapterHostFcInterfaceAllOf

`func NewAdapterHostFcInterfaceAllOf() *AdapterHostFcInterfaceAllOf`

NewAdapterHostFcInterfaceAllOf instantiates a new AdapterHostFcInterfaceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdapterHostFcInterfaceAllOfWithDefaults

`func NewAdapterHostFcInterfaceAllOfWithDefaults() *AdapterHostFcInterfaceAllOf`

NewAdapterHostFcInterfaceAllOfWithDefaults instantiates a new AdapterHostFcInterfaceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminState

`func (o *AdapterHostFcInterfaceAllOf) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *AdapterHostFcInterfaceAllOf) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *AdapterHostFcInterfaceAllOf) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *AdapterHostFcInterfaceAllOf) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetEpDn

`func (o *AdapterHostFcInterfaceAllOf) GetEpDn() string`

GetEpDn returns the EpDn field if non-nil, zero value otherwise.

### GetEpDnOk

`func (o *AdapterHostFcInterfaceAllOf) GetEpDnOk() (*string, bool)`

GetEpDnOk returns a tuple with the EpDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpDn

`func (o *AdapterHostFcInterfaceAllOf) SetEpDn(v string)`

SetEpDn sets EpDn field to given value.

### HasEpDn

`func (o *AdapterHostFcInterfaceAllOf) HasEpDn() bool`

HasEpDn returns a boolean if a field has been set.

### GetHostFcInterfaceId

`func (o *AdapterHostFcInterfaceAllOf) GetHostFcInterfaceId() int64`

GetHostFcInterfaceId returns the HostFcInterfaceId field if non-nil, zero value otherwise.

### GetHostFcInterfaceIdOk

`func (o *AdapterHostFcInterfaceAllOf) GetHostFcInterfaceIdOk() (*int64, bool)`

GetHostFcInterfaceIdOk returns a tuple with the HostFcInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostFcInterfaceId

`func (o *AdapterHostFcInterfaceAllOf) SetHostFcInterfaceId(v int64)`

SetHostFcInterfaceId sets HostFcInterfaceId field to given value.

### HasHostFcInterfaceId

`func (o *AdapterHostFcInterfaceAllOf) HasHostFcInterfaceId() bool`

HasHostFcInterfaceId returns a boolean if a field has been set.

### GetName

`func (o *AdapterHostFcInterfaceAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdapterHostFcInterfaceAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdapterHostFcInterfaceAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdapterHostFcInterfaceAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperState

`func (o *AdapterHostFcInterfaceAllOf) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *AdapterHostFcInterfaceAllOf) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *AdapterHostFcInterfaceAllOf) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *AdapterHostFcInterfaceAllOf) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetOperability

`func (o *AdapterHostFcInterfaceAllOf) GetOperability() string`

GetOperability returns the Operability field if non-nil, zero value otherwise.

### GetOperabilityOk

`func (o *AdapterHostFcInterfaceAllOf) GetOperabilityOk() (*string, bool)`

GetOperabilityOk returns a tuple with the Operability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperability

`func (o *AdapterHostFcInterfaceAllOf) SetOperability(v string)`

SetOperability sets Operability field to given value.

### HasOperability

`func (o *AdapterHostFcInterfaceAllOf) HasOperability() bool`

HasOperability returns a boolean if a field has been set.

### GetOriginalWwnn

`func (o *AdapterHostFcInterfaceAllOf) GetOriginalWwnn() string`

GetOriginalWwnn returns the OriginalWwnn field if non-nil, zero value otherwise.

### GetOriginalWwnnOk

`func (o *AdapterHostFcInterfaceAllOf) GetOriginalWwnnOk() (*string, bool)`

GetOriginalWwnnOk returns a tuple with the OriginalWwnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalWwnn

`func (o *AdapterHostFcInterfaceAllOf) SetOriginalWwnn(v string)`

SetOriginalWwnn sets OriginalWwnn field to given value.

### HasOriginalWwnn

`func (o *AdapterHostFcInterfaceAllOf) HasOriginalWwnn() bool`

HasOriginalWwnn returns a boolean if a field has been set.

### GetOriginalWwpn

`func (o *AdapterHostFcInterfaceAllOf) GetOriginalWwpn() string`

GetOriginalWwpn returns the OriginalWwpn field if non-nil, zero value otherwise.

### GetOriginalWwpnOk

`func (o *AdapterHostFcInterfaceAllOf) GetOriginalWwpnOk() (*string, bool)`

GetOriginalWwpnOk returns a tuple with the OriginalWwpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalWwpn

`func (o *AdapterHostFcInterfaceAllOf) SetOriginalWwpn(v string)`

SetOriginalWwpn sets OriginalWwpn field to given value.

### HasOriginalWwpn

`func (o *AdapterHostFcInterfaceAllOf) HasOriginalWwpn() bool`

HasOriginalWwpn returns a boolean if a field has been set.

### GetPeerDn

`func (o *AdapterHostFcInterfaceAllOf) GetPeerDn() string`

GetPeerDn returns the PeerDn field if non-nil, zero value otherwise.

### GetPeerDnOk

`func (o *AdapterHostFcInterfaceAllOf) GetPeerDnOk() (*string, bool)`

GetPeerDnOk returns a tuple with the PeerDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerDn

`func (o *AdapterHostFcInterfaceAllOf) SetPeerDn(v string)`

SetPeerDn sets PeerDn field to given value.

### HasPeerDn

`func (o *AdapterHostFcInterfaceAllOf) HasPeerDn() bool`

HasPeerDn returns a boolean if a field has been set.

### GetWwnn

`func (o *AdapterHostFcInterfaceAllOf) GetWwnn() string`

GetWwnn returns the Wwnn field if non-nil, zero value otherwise.

### GetWwnnOk

`func (o *AdapterHostFcInterfaceAllOf) GetWwnnOk() (*string, bool)`

GetWwnnOk returns a tuple with the Wwnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwnn

`func (o *AdapterHostFcInterfaceAllOf) SetWwnn(v string)`

SetWwnn sets Wwnn field to given value.

### HasWwnn

`func (o *AdapterHostFcInterfaceAllOf) HasWwnn() bool`

HasWwnn returns a boolean if a field has been set.

### GetWwpn

`func (o *AdapterHostFcInterfaceAllOf) GetWwpn() string`

GetWwpn returns the Wwpn field if non-nil, zero value otherwise.

### GetWwpnOk

`func (o *AdapterHostFcInterfaceAllOf) GetWwpnOk() (*string, bool)`

GetWwpnOk returns a tuple with the Wwpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwpn

`func (o *AdapterHostFcInterfaceAllOf) SetWwpn(v string)`

SetWwpn sets Wwpn field to given value.

### HasWwpn

`func (o *AdapterHostFcInterfaceAllOf) HasWwpn() bool`

HasWwpn returns a boolean if a field has been set.

### GetAdapterUnit

`func (o *AdapterHostFcInterfaceAllOf) GetAdapterUnit() AdapterUnitRelationship`

GetAdapterUnit returns the AdapterUnit field if non-nil, zero value otherwise.

### GetAdapterUnitOk

`func (o *AdapterHostFcInterfaceAllOf) GetAdapterUnitOk() (*AdapterUnitRelationship, bool)`

GetAdapterUnitOk returns a tuple with the AdapterUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterUnit

`func (o *AdapterHostFcInterfaceAllOf) SetAdapterUnit(v AdapterUnitRelationship)`

SetAdapterUnit sets AdapterUnit field to given value.

### HasAdapterUnit

`func (o *AdapterHostFcInterfaceAllOf) HasAdapterUnit() bool`

HasAdapterUnit returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *AdapterHostFcInterfaceAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *AdapterHostFcInterfaceAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *AdapterHostFcInterfaceAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *AdapterHostFcInterfaceAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *AdapterHostFcInterfaceAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *AdapterHostFcInterfaceAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *AdapterHostFcInterfaceAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *AdapterHostFcInterfaceAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


