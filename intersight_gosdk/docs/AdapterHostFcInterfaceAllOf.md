# AdapterHostFcInterfaceAllOf

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
**PinGroupName** | Pointer to **string** | Name given for San PinGroup. | [optional] 
**VfcAction** | Pointer to **string** | Denotes the action to be performed on the vfc corresponding to the vHBA. * &#x60;None&#x60; - Default value for vif operation. * &#x60;ResetConnectivity&#x60; - Resets connectivity on both active and passive vif. * &#x60;ResetConnectivityActive&#x60; - Resets connectivity on the active vif. * &#x60;ResetConnectivityPassive&#x60; - Resets connectivity on the passive vif. * &#x60;Enable&#x60; - Enables the vif on both the FIs. * &#x60;Disable&#x60; - Disables the vif on both the FIs. * &#x60;EnableActive&#x60; - Enables the corresponding active vif. * &#x60;EnablePassive&#x60; - Enables the corresponding standby vif. * &#x60;DisableActive&#x60; - Disables the corresponding active vif. * &#x60;DisablePassive&#x60; - Disables the corresponding standby vif. | [optional] [default to "None"]
**VifId** | Pointer to **int64** | Identifier of the virtual fibre channel (Vfc) interface on the networking component (e.g., Fabric Interconnect) for the corresponding Host Fibre Channel Interface. | [optional] [readonly] 
**Wwnn** | Pointer to **string** | The uniquely distinguishable user configured World Wide Node Name of the Host. | [optional] [readonly] 
**Wwpn** | Pointer to **string** | The uniquely distinguishable user configured World Wide Port Name of the Host Fibre Channel Interface. | [optional] [readonly] 
**AdapterUnit** | Pointer to [**AdapterUnitRelationship**](AdapterUnitRelationship.md) |  | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**PinnedInterface** | Pointer to [**InventoryInterfaceRelationship**](InventoryInterfaceRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**Vfc** | Pointer to [**NetworkVfcRelationship**](NetworkVfcRelationship.md) |  | [optional] 

## Methods

### NewAdapterHostFcInterfaceAllOf

`func NewAdapterHostFcInterfaceAllOf(classId string, objectType string, ) *AdapterHostFcInterfaceAllOf`

NewAdapterHostFcInterfaceAllOf instantiates a new AdapterHostFcInterfaceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdapterHostFcInterfaceAllOfWithDefaults

`func NewAdapterHostFcInterfaceAllOfWithDefaults() *AdapterHostFcInterfaceAllOf`

NewAdapterHostFcInterfaceAllOfWithDefaults instantiates a new AdapterHostFcInterfaceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *AdapterHostFcInterfaceAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *AdapterHostFcInterfaceAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *AdapterHostFcInterfaceAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *AdapterHostFcInterfaceAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AdapterHostFcInterfaceAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AdapterHostFcInterfaceAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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

### GetOperReason

`func (o *AdapterHostFcInterfaceAllOf) GetOperReason() []string`

GetOperReason returns the OperReason field if non-nil, zero value otherwise.

### GetOperReasonOk

`func (o *AdapterHostFcInterfaceAllOf) GetOperReasonOk() (*[]string, bool)`

GetOperReasonOk returns a tuple with the OperReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperReason

`func (o *AdapterHostFcInterfaceAllOf) SetOperReason(v []string)`

SetOperReason sets OperReason field to given value.

### HasOperReason

`func (o *AdapterHostFcInterfaceAllOf) HasOperReason() bool`

HasOperReason returns a boolean if a field has been set.

### SetOperReasonNil

`func (o *AdapterHostFcInterfaceAllOf) SetOperReasonNil(b bool)`

 SetOperReasonNil sets the value for OperReason to be an explicit nil

### UnsetOperReason
`func (o *AdapterHostFcInterfaceAllOf) UnsetOperReason()`

UnsetOperReason ensures that no value is present for OperReason, not even an explicit nil
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

### GetPinGroupName

`func (o *AdapterHostFcInterfaceAllOf) GetPinGroupName() string`

GetPinGroupName returns the PinGroupName field if non-nil, zero value otherwise.

### GetPinGroupNameOk

`func (o *AdapterHostFcInterfaceAllOf) GetPinGroupNameOk() (*string, bool)`

GetPinGroupNameOk returns a tuple with the PinGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinGroupName

`func (o *AdapterHostFcInterfaceAllOf) SetPinGroupName(v string)`

SetPinGroupName sets PinGroupName field to given value.

### HasPinGroupName

`func (o *AdapterHostFcInterfaceAllOf) HasPinGroupName() bool`

HasPinGroupName returns a boolean if a field has been set.

### GetVfcAction

`func (o *AdapterHostFcInterfaceAllOf) GetVfcAction() string`

GetVfcAction returns the VfcAction field if non-nil, zero value otherwise.

### GetVfcActionOk

`func (o *AdapterHostFcInterfaceAllOf) GetVfcActionOk() (*string, bool)`

GetVfcActionOk returns a tuple with the VfcAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVfcAction

`func (o *AdapterHostFcInterfaceAllOf) SetVfcAction(v string)`

SetVfcAction sets VfcAction field to given value.

### HasVfcAction

`func (o *AdapterHostFcInterfaceAllOf) HasVfcAction() bool`

HasVfcAction returns a boolean if a field has been set.

### GetVifId

`func (o *AdapterHostFcInterfaceAllOf) GetVifId() int64`

GetVifId returns the VifId field if non-nil, zero value otherwise.

### GetVifIdOk

`func (o *AdapterHostFcInterfaceAllOf) GetVifIdOk() (*int64, bool)`

GetVifIdOk returns a tuple with the VifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVifId

`func (o *AdapterHostFcInterfaceAllOf) SetVifId(v int64)`

SetVifId sets VifId field to given value.

### HasVifId

`func (o *AdapterHostFcInterfaceAllOf) HasVifId() bool`

HasVifId returns a boolean if a field has been set.

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

### GetPinnedInterface

`func (o *AdapterHostFcInterfaceAllOf) GetPinnedInterface() InventoryInterfaceRelationship`

GetPinnedInterface returns the PinnedInterface field if non-nil, zero value otherwise.

### GetPinnedInterfaceOk

`func (o *AdapterHostFcInterfaceAllOf) GetPinnedInterfaceOk() (*InventoryInterfaceRelationship, bool)`

GetPinnedInterfaceOk returns a tuple with the PinnedInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinnedInterface

`func (o *AdapterHostFcInterfaceAllOf) SetPinnedInterface(v InventoryInterfaceRelationship)`

SetPinnedInterface sets PinnedInterface field to given value.

### HasPinnedInterface

`func (o *AdapterHostFcInterfaceAllOf) HasPinnedInterface() bool`

HasPinnedInterface returns a boolean if a field has been set.

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

### GetVfc

`func (o *AdapterHostFcInterfaceAllOf) GetVfc() NetworkVfcRelationship`

GetVfc returns the Vfc field if non-nil, zero value otherwise.

### GetVfcOk

`func (o *AdapterHostFcInterfaceAllOf) GetVfcOk() (*NetworkVfcRelationship, bool)`

GetVfcOk returns a tuple with the Vfc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVfc

`func (o *AdapterHostFcInterfaceAllOf) SetVfc(v NetworkVfcRelationship)`

SetVfc sets Vfc field to given value.

### HasVfc

`func (o *AdapterHostFcInterfaceAllOf) HasVfc() bool`

HasVfc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


