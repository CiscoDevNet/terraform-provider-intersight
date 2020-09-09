# AdapterHostIscsiInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminState** | Pointer to **string** | Admin Configured State of Host ISCSI Interface. | [optional] [readonly] 
**EpDn** | Pointer to **string** | The Endpoint Config Dn of the Host ISCSI Interface. | [optional] [readonly] 
**HostIscsiInterfaceId** | Pointer to **int64** | Identifier of the Host ISCSI Interface. | [optional] [readonly] 
**HostVisible** | Pointer to **string** | The visibility of the Host to the endpoint. | [optional] [readonly] 
**MacAddress** | Pointer to **string** | MAC address of Host ISCSI Interface. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the Host ISCSI Interface. | [optional] [readonly] 
**OperState** | Pointer to **string** | Operational State of Host ISCSI Interface. | [optional] [readonly] 
**Operability** | Pointer to **string** | Operability status of Host ISCSI Interface. | [optional] [readonly] 
**PeerDn** | Pointer to **string** | PeerPort Dn of Host ISCSI Interface. | [optional] [readonly] 
**AdapterUnit** | Pointer to [**AdapterUnitRelationship**](adapter.Unit.Relationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](inventory.DeviceInfo.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewAdapterHostIscsiInterface

`func NewAdapterHostIscsiInterface() *AdapterHostIscsiInterface`

NewAdapterHostIscsiInterface instantiates a new AdapterHostIscsiInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdapterHostIscsiInterfaceWithDefaults

`func NewAdapterHostIscsiInterfaceWithDefaults() *AdapterHostIscsiInterface`

NewAdapterHostIscsiInterfaceWithDefaults instantiates a new AdapterHostIscsiInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminState

`func (o *AdapterHostIscsiInterface) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *AdapterHostIscsiInterface) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *AdapterHostIscsiInterface) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *AdapterHostIscsiInterface) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetEpDn

`func (o *AdapterHostIscsiInterface) GetEpDn() string`

GetEpDn returns the EpDn field if non-nil, zero value otherwise.

### GetEpDnOk

`func (o *AdapterHostIscsiInterface) GetEpDnOk() (*string, bool)`

GetEpDnOk returns a tuple with the EpDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpDn

`func (o *AdapterHostIscsiInterface) SetEpDn(v string)`

SetEpDn sets EpDn field to given value.

### HasEpDn

`func (o *AdapterHostIscsiInterface) HasEpDn() bool`

HasEpDn returns a boolean if a field has been set.

### GetHostIscsiInterfaceId

`func (o *AdapterHostIscsiInterface) GetHostIscsiInterfaceId() int64`

GetHostIscsiInterfaceId returns the HostIscsiInterfaceId field if non-nil, zero value otherwise.

### GetHostIscsiInterfaceIdOk

`func (o *AdapterHostIscsiInterface) GetHostIscsiInterfaceIdOk() (*int64, bool)`

GetHostIscsiInterfaceIdOk returns a tuple with the HostIscsiInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostIscsiInterfaceId

`func (o *AdapterHostIscsiInterface) SetHostIscsiInterfaceId(v int64)`

SetHostIscsiInterfaceId sets HostIscsiInterfaceId field to given value.

### HasHostIscsiInterfaceId

`func (o *AdapterHostIscsiInterface) HasHostIscsiInterfaceId() bool`

HasHostIscsiInterfaceId returns a boolean if a field has been set.

### GetHostVisible

`func (o *AdapterHostIscsiInterface) GetHostVisible() string`

GetHostVisible returns the HostVisible field if non-nil, zero value otherwise.

### GetHostVisibleOk

`func (o *AdapterHostIscsiInterface) GetHostVisibleOk() (*string, bool)`

GetHostVisibleOk returns a tuple with the HostVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostVisible

`func (o *AdapterHostIscsiInterface) SetHostVisible(v string)`

SetHostVisible sets HostVisible field to given value.

### HasHostVisible

`func (o *AdapterHostIscsiInterface) HasHostVisible() bool`

HasHostVisible returns a boolean if a field has been set.

### GetMacAddress

`func (o *AdapterHostIscsiInterface) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *AdapterHostIscsiInterface) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *AdapterHostIscsiInterface) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *AdapterHostIscsiInterface) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetName

`func (o *AdapterHostIscsiInterface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdapterHostIscsiInterface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdapterHostIscsiInterface) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdapterHostIscsiInterface) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperState

`func (o *AdapterHostIscsiInterface) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *AdapterHostIscsiInterface) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *AdapterHostIscsiInterface) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *AdapterHostIscsiInterface) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetOperability

`func (o *AdapterHostIscsiInterface) GetOperability() string`

GetOperability returns the Operability field if non-nil, zero value otherwise.

### GetOperabilityOk

`func (o *AdapterHostIscsiInterface) GetOperabilityOk() (*string, bool)`

GetOperabilityOk returns a tuple with the Operability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperability

`func (o *AdapterHostIscsiInterface) SetOperability(v string)`

SetOperability sets Operability field to given value.

### HasOperability

`func (o *AdapterHostIscsiInterface) HasOperability() bool`

HasOperability returns a boolean if a field has been set.

### GetPeerDn

`func (o *AdapterHostIscsiInterface) GetPeerDn() string`

GetPeerDn returns the PeerDn field if non-nil, zero value otherwise.

### GetPeerDnOk

`func (o *AdapterHostIscsiInterface) GetPeerDnOk() (*string, bool)`

GetPeerDnOk returns a tuple with the PeerDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerDn

`func (o *AdapterHostIscsiInterface) SetPeerDn(v string)`

SetPeerDn sets PeerDn field to given value.

### HasPeerDn

`func (o *AdapterHostIscsiInterface) HasPeerDn() bool`

HasPeerDn returns a boolean if a field has been set.

### GetAdapterUnit

`func (o *AdapterHostIscsiInterface) GetAdapterUnit() AdapterUnitRelationship`

GetAdapterUnit returns the AdapterUnit field if non-nil, zero value otherwise.

### GetAdapterUnitOk

`func (o *AdapterHostIscsiInterface) GetAdapterUnitOk() (*AdapterUnitRelationship, bool)`

GetAdapterUnitOk returns a tuple with the AdapterUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterUnit

`func (o *AdapterHostIscsiInterface) SetAdapterUnit(v AdapterUnitRelationship)`

SetAdapterUnit sets AdapterUnit field to given value.

### HasAdapterUnit

`func (o *AdapterHostIscsiInterface) HasAdapterUnit() bool`

HasAdapterUnit returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *AdapterHostIscsiInterface) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *AdapterHostIscsiInterface) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *AdapterHostIscsiInterface) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *AdapterHostIscsiInterface) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *AdapterHostIscsiInterface) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *AdapterHostIscsiInterface) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *AdapterHostIscsiInterface) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *AdapterHostIscsiInterface) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


