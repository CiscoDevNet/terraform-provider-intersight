# AdapterHostFcInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "adapter.HostFcInterface"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "adapter.HostFcInterface"]
**AdminState** | Pointer to **string** | Admin Configured State of Host Fibre Channel Interface. | [optional] [readonly] 
**EpDn** | Pointer to **string** | The Endpoint Config Dn of the Host Fibre Channel Interface. | [optional] [readonly] 
**HostFcInterfaceId** | Pointer to **int64** | Identifier of Host Fibre Channel Interface. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of Host Fibre Channel Interface. | [optional] [readonly] 
**OperReason** | Pointer to **[]string** |  | [optional] 
**OperState** | Pointer to **string** | Operational State of Host Fibre Channel Interface. | [optional] [readonly] 
**Operability** | Pointer to **string** | Operability status of Host Fibre Channel Interface. | [optional] [readonly] 
**OriginalWwnn** | Pointer to **string** | The uniquely distinguishable factory default  World Wide Node Name of the Host. | [optional] [readonly] 
**OriginalWwpn** | Pointer to **string** | The uniquely distinguishable factory default World Wide Port Name of the Host Fibre Channel Interface. | [optional] [readonly] 
**PeerDn** | Pointer to **string** | PeerPort Dn of Host Fibre Channel Interface. | [optional] [readonly] 
**Wwnn** | Pointer to **string** | The uniquely distinguishable user configured World Wide Node Name of the Host. | [optional] [readonly] 
**Wwpn** | Pointer to **string** | The uniquely distinguishable user configured World Wide Port Name of the Host Fibre Channel Interface. | [optional] [readonly] 
**AdapterUnit** | Pointer to [**AdapterUnitRelationship**](AdapterUnitRelationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewAdapterHostFcInterface

`func NewAdapterHostFcInterface(classId string, objectType string, ) *AdapterHostFcInterface`

NewAdapterHostFcInterface instantiates a new AdapterHostFcInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdapterHostFcInterfaceWithDefaults

`func NewAdapterHostFcInterfaceWithDefaults() *AdapterHostFcInterface`

NewAdapterHostFcInterfaceWithDefaults instantiates a new AdapterHostFcInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AdapterHostFcInterface) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AdapterHostFcInterface) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AdapterHostFcInterface) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AdapterHostFcInterface) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AdapterHostFcInterface) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AdapterHostFcInterface) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminState

`func (o *AdapterHostFcInterface) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *AdapterHostFcInterface) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *AdapterHostFcInterface) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *AdapterHostFcInterface) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetEpDn

`func (o *AdapterHostFcInterface) GetEpDn() string`

GetEpDn returns the EpDn field if non-nil, zero value otherwise.

### GetEpDnOk

`func (o *AdapterHostFcInterface) GetEpDnOk() (*string, bool)`

GetEpDnOk returns a tuple with the EpDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpDn

`func (o *AdapterHostFcInterface) SetEpDn(v string)`

SetEpDn sets EpDn field to given value.

### HasEpDn

`func (o *AdapterHostFcInterface) HasEpDn() bool`

HasEpDn returns a boolean if a field has been set.

### GetHostFcInterfaceId

`func (o *AdapterHostFcInterface) GetHostFcInterfaceId() int64`

GetHostFcInterfaceId returns the HostFcInterfaceId field if non-nil, zero value otherwise.

### GetHostFcInterfaceIdOk

`func (o *AdapterHostFcInterface) GetHostFcInterfaceIdOk() (*int64, bool)`

GetHostFcInterfaceIdOk returns a tuple with the HostFcInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostFcInterfaceId

`func (o *AdapterHostFcInterface) SetHostFcInterfaceId(v int64)`

SetHostFcInterfaceId sets HostFcInterfaceId field to given value.

### HasHostFcInterfaceId

`func (o *AdapterHostFcInterface) HasHostFcInterfaceId() bool`

HasHostFcInterfaceId returns a boolean if a field has been set.

### GetName

`func (o *AdapterHostFcInterface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdapterHostFcInterface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdapterHostFcInterface) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdapterHostFcInterface) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperReason

`func (o *AdapterHostFcInterface) GetOperReason() []string`

GetOperReason returns the OperReason field if non-nil, zero value otherwise.

### GetOperReasonOk

`func (o *AdapterHostFcInterface) GetOperReasonOk() (*[]string, bool)`

GetOperReasonOk returns a tuple with the OperReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperReason

`func (o *AdapterHostFcInterface) SetOperReason(v []string)`

SetOperReason sets OperReason field to given value.

### HasOperReason

`func (o *AdapterHostFcInterface) HasOperReason() bool`

HasOperReason returns a boolean if a field has been set.

### SetOperReasonNil

`func (o *AdapterHostFcInterface) SetOperReasonNil(b bool)`

 SetOperReasonNil sets the value for OperReason to be an explicit nil

### UnsetOperReason
`func (o *AdapterHostFcInterface) UnsetOperReason()`

UnsetOperReason ensures that no value is present for OperReason, not even an explicit nil
### GetOperState

`func (o *AdapterHostFcInterface) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *AdapterHostFcInterface) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *AdapterHostFcInterface) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *AdapterHostFcInterface) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetOperability

`func (o *AdapterHostFcInterface) GetOperability() string`

GetOperability returns the Operability field if non-nil, zero value otherwise.

### GetOperabilityOk

`func (o *AdapterHostFcInterface) GetOperabilityOk() (*string, bool)`

GetOperabilityOk returns a tuple with the Operability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperability

`func (o *AdapterHostFcInterface) SetOperability(v string)`

SetOperability sets Operability field to given value.

### HasOperability

`func (o *AdapterHostFcInterface) HasOperability() bool`

HasOperability returns a boolean if a field has been set.

### GetOriginalWwnn

`func (o *AdapterHostFcInterface) GetOriginalWwnn() string`

GetOriginalWwnn returns the OriginalWwnn field if non-nil, zero value otherwise.

### GetOriginalWwnnOk

`func (o *AdapterHostFcInterface) GetOriginalWwnnOk() (*string, bool)`

GetOriginalWwnnOk returns a tuple with the OriginalWwnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalWwnn

`func (o *AdapterHostFcInterface) SetOriginalWwnn(v string)`

SetOriginalWwnn sets OriginalWwnn field to given value.

### HasOriginalWwnn

`func (o *AdapterHostFcInterface) HasOriginalWwnn() bool`

HasOriginalWwnn returns a boolean if a field has been set.

### GetOriginalWwpn

`func (o *AdapterHostFcInterface) GetOriginalWwpn() string`

GetOriginalWwpn returns the OriginalWwpn field if non-nil, zero value otherwise.

### GetOriginalWwpnOk

`func (o *AdapterHostFcInterface) GetOriginalWwpnOk() (*string, bool)`

GetOriginalWwpnOk returns a tuple with the OriginalWwpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalWwpn

`func (o *AdapterHostFcInterface) SetOriginalWwpn(v string)`

SetOriginalWwpn sets OriginalWwpn field to given value.

### HasOriginalWwpn

`func (o *AdapterHostFcInterface) HasOriginalWwpn() bool`

HasOriginalWwpn returns a boolean if a field has been set.

### GetPeerDn

`func (o *AdapterHostFcInterface) GetPeerDn() string`

GetPeerDn returns the PeerDn field if non-nil, zero value otherwise.

### GetPeerDnOk

`func (o *AdapterHostFcInterface) GetPeerDnOk() (*string, bool)`

GetPeerDnOk returns a tuple with the PeerDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerDn

`func (o *AdapterHostFcInterface) SetPeerDn(v string)`

SetPeerDn sets PeerDn field to given value.

### HasPeerDn

`func (o *AdapterHostFcInterface) HasPeerDn() bool`

HasPeerDn returns a boolean if a field has been set.

### GetWwnn

`func (o *AdapterHostFcInterface) GetWwnn() string`

GetWwnn returns the Wwnn field if non-nil, zero value otherwise.

### GetWwnnOk

`func (o *AdapterHostFcInterface) GetWwnnOk() (*string, bool)`

GetWwnnOk returns a tuple with the Wwnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwnn

`func (o *AdapterHostFcInterface) SetWwnn(v string)`

SetWwnn sets Wwnn field to given value.

### HasWwnn

`func (o *AdapterHostFcInterface) HasWwnn() bool`

HasWwnn returns a boolean if a field has been set.

### GetWwpn

`func (o *AdapterHostFcInterface) GetWwpn() string`

GetWwpn returns the Wwpn field if non-nil, zero value otherwise.

### GetWwpnOk

`func (o *AdapterHostFcInterface) GetWwpnOk() (*string, bool)`

GetWwpnOk returns a tuple with the Wwpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwpn

`func (o *AdapterHostFcInterface) SetWwpn(v string)`

SetWwpn sets Wwpn field to given value.

### HasWwpn

`func (o *AdapterHostFcInterface) HasWwpn() bool`

HasWwpn returns a boolean if a field has been set.

### GetAdapterUnit

`func (o *AdapterHostFcInterface) GetAdapterUnit() AdapterUnitRelationship`

GetAdapterUnit returns the AdapterUnit field if non-nil, zero value otherwise.

### GetAdapterUnitOk

`func (o *AdapterHostFcInterface) GetAdapterUnitOk() (*AdapterUnitRelationship, bool)`

GetAdapterUnitOk returns a tuple with the AdapterUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterUnit

`func (o *AdapterHostFcInterface) SetAdapterUnit(v AdapterUnitRelationship)`

SetAdapterUnit sets AdapterUnit field to given value.

### HasAdapterUnit

`func (o *AdapterHostFcInterface) HasAdapterUnit() bool`

HasAdapterUnit returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *AdapterHostFcInterface) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *AdapterHostFcInterface) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *AdapterHostFcInterface) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *AdapterHostFcInterface) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *AdapterHostFcInterface) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *AdapterHostFcInterface) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *AdapterHostFcInterface) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *AdapterHostFcInterface) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


