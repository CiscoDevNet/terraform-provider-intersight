# NetworkVfc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "network.Vfc"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "network.Vfc"]
**BoundInterfaceDn** | Pointer to **string** | Port or PortChannel interface configured for Vfcs virtual etherent interface. | [optional] [readonly] 
**Burst** | Pointer to **int64** | Burst defined on QoS policy. | [optional] [readonly] 
**Description** | Pointer to **string** | Description for the vHBA. | [optional] [readonly] 
**OperReason** | Pointer to **string** | Operational Reason of the Vfcs Vethernet on the Fabric Interconnect. When the operational state is down, Operreason indicates the reason why the vHBA is not operational. Some common reasons are admindown, error disabled. | [optional] [readonly] 
**OperState** | Pointer to **string** | The operational state of the Vfcs Vethernet peer of a vNIC in Intersight Managed Mode. This state is updated by events from Fabric Interconnect or by periodic updates from Fabric Interconnect. When a Fabric Interconnect is not connected to Intersight or if the Fabric Interconnect is powered down, this property will not be updated. * &#x60;unknown&#x60; - The operational state of the Vethernet is not known. * &#x60;adminDown&#x60; - The operational state of the Vethernet is admin down. * &#x60;up&#x60; - The operational state of the Vethernet is Up. * &#x60;down&#x60; - The operational state of the Vethernet is Down. * &#x60;noLicense&#x60; - The operational state of the Vethernet is no license. * &#x60;linkUp&#x60; - The operational state of the Vethernet is link up. * &#x60;hardwareFailure&#x60; - The operational state of the Vethernet is hardware failure. * &#x60;softwareFailure&#x60; - The operational state of the Vethernet is software failure. * &#x60;errorDisabled&#x60; - The operational state of the Vethernet is error disabled. * &#x60;linkDown&#x60; - The operational state of the Vethernet is link down. * &#x60;sfpNotPresent&#x60; - The operational state of the Vethernet is SFP not present. * &#x60;udldAggrDown&#x60; - The operational state of the Vethernet is UDLD aggregate down. | [optional] [readonly] [default to "unknown"]
**PinnedInterfaceDn** | Pointer to **string** | Uplink port or portchannel pinned to a Vfc. | [optional] [readonly] 
**Ratelimit** | Pointer to **int64** | Rate limit defined on QoS policy. | [optional] [readonly] 
**VfcId** | Pointer to **int64** | Vfc Identifier on a Fabric Interconnect. | [optional] [readonly] 
**AdapterHostFcInterface** | Pointer to [**NullableAdapterHostFcInterfaceRelationship**](AdapterHostFcInterfaceRelationship.md) |  | [optional] 
**BoundInterface** | Pointer to [**NullableInventoryInterfaceRelationship**](InventoryInterfaceRelationship.md) |  | [optional] 
**NetworkElement** | Pointer to [**NullableNetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**PinnedInterface** | Pointer to [**NullableInventoryInterfaceRelationship**](InventoryInterfaceRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNetworkVfc

`func NewNetworkVfc(classId string, objectType string, ) *NetworkVfc`

NewNetworkVfc instantiates a new NetworkVfc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkVfcWithDefaults

`func NewNetworkVfcWithDefaults() *NetworkVfc`

NewNetworkVfcWithDefaults instantiates a new NetworkVfc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NetworkVfc) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NetworkVfc) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NetworkVfc) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NetworkVfc) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NetworkVfc) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NetworkVfc) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBoundInterfaceDn

`func (o *NetworkVfc) GetBoundInterfaceDn() string`

GetBoundInterfaceDn returns the BoundInterfaceDn field if non-nil, zero value otherwise.

### GetBoundInterfaceDnOk

`func (o *NetworkVfc) GetBoundInterfaceDnOk() (*string, bool)`

GetBoundInterfaceDnOk returns a tuple with the BoundInterfaceDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundInterfaceDn

`func (o *NetworkVfc) SetBoundInterfaceDn(v string)`

SetBoundInterfaceDn sets BoundInterfaceDn field to given value.

### HasBoundInterfaceDn

`func (o *NetworkVfc) HasBoundInterfaceDn() bool`

HasBoundInterfaceDn returns a boolean if a field has been set.

### GetBurst

`func (o *NetworkVfc) GetBurst() int64`

GetBurst returns the Burst field if non-nil, zero value otherwise.

### GetBurstOk

`func (o *NetworkVfc) GetBurstOk() (*int64, bool)`

GetBurstOk returns a tuple with the Burst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurst

`func (o *NetworkVfc) SetBurst(v int64)`

SetBurst sets Burst field to given value.

### HasBurst

`func (o *NetworkVfc) HasBurst() bool`

HasBurst returns a boolean if a field has been set.

### GetDescription

`func (o *NetworkVfc) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkVfc) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkVfc) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkVfc) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOperReason

`func (o *NetworkVfc) GetOperReason() string`

GetOperReason returns the OperReason field if non-nil, zero value otherwise.

### GetOperReasonOk

`func (o *NetworkVfc) GetOperReasonOk() (*string, bool)`

GetOperReasonOk returns a tuple with the OperReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperReason

`func (o *NetworkVfc) SetOperReason(v string)`

SetOperReason sets OperReason field to given value.

### HasOperReason

`func (o *NetworkVfc) HasOperReason() bool`

HasOperReason returns a boolean if a field has been set.

### GetOperState

`func (o *NetworkVfc) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *NetworkVfc) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *NetworkVfc) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *NetworkVfc) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetPinnedInterfaceDn

`func (o *NetworkVfc) GetPinnedInterfaceDn() string`

GetPinnedInterfaceDn returns the PinnedInterfaceDn field if non-nil, zero value otherwise.

### GetPinnedInterfaceDnOk

`func (o *NetworkVfc) GetPinnedInterfaceDnOk() (*string, bool)`

GetPinnedInterfaceDnOk returns a tuple with the PinnedInterfaceDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinnedInterfaceDn

`func (o *NetworkVfc) SetPinnedInterfaceDn(v string)`

SetPinnedInterfaceDn sets PinnedInterfaceDn field to given value.

### HasPinnedInterfaceDn

`func (o *NetworkVfc) HasPinnedInterfaceDn() bool`

HasPinnedInterfaceDn returns a boolean if a field has been set.

### GetRatelimit

`func (o *NetworkVfc) GetRatelimit() int64`

GetRatelimit returns the Ratelimit field if non-nil, zero value otherwise.

### GetRatelimitOk

`func (o *NetworkVfc) GetRatelimitOk() (*int64, bool)`

GetRatelimitOk returns a tuple with the Ratelimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatelimit

`func (o *NetworkVfc) SetRatelimit(v int64)`

SetRatelimit sets Ratelimit field to given value.

### HasRatelimit

`func (o *NetworkVfc) HasRatelimit() bool`

HasRatelimit returns a boolean if a field has been set.

### GetVfcId

`func (o *NetworkVfc) GetVfcId() int64`

GetVfcId returns the VfcId field if non-nil, zero value otherwise.

### GetVfcIdOk

`func (o *NetworkVfc) GetVfcIdOk() (*int64, bool)`

GetVfcIdOk returns a tuple with the VfcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVfcId

`func (o *NetworkVfc) SetVfcId(v int64)`

SetVfcId sets VfcId field to given value.

### HasVfcId

`func (o *NetworkVfc) HasVfcId() bool`

HasVfcId returns a boolean if a field has been set.

### GetAdapterHostFcInterface

`func (o *NetworkVfc) GetAdapterHostFcInterface() AdapterHostFcInterfaceRelationship`

GetAdapterHostFcInterface returns the AdapterHostFcInterface field if non-nil, zero value otherwise.

### GetAdapterHostFcInterfaceOk

`func (o *NetworkVfc) GetAdapterHostFcInterfaceOk() (*AdapterHostFcInterfaceRelationship, bool)`

GetAdapterHostFcInterfaceOk returns a tuple with the AdapterHostFcInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterHostFcInterface

`func (o *NetworkVfc) SetAdapterHostFcInterface(v AdapterHostFcInterfaceRelationship)`

SetAdapterHostFcInterface sets AdapterHostFcInterface field to given value.

### HasAdapterHostFcInterface

`func (o *NetworkVfc) HasAdapterHostFcInterface() bool`

HasAdapterHostFcInterface returns a boolean if a field has been set.

### SetAdapterHostFcInterfaceNil

`func (o *NetworkVfc) SetAdapterHostFcInterfaceNil(b bool)`

 SetAdapterHostFcInterfaceNil sets the value for AdapterHostFcInterface to be an explicit nil

### UnsetAdapterHostFcInterface
`func (o *NetworkVfc) UnsetAdapterHostFcInterface()`

UnsetAdapterHostFcInterface ensures that no value is present for AdapterHostFcInterface, not even an explicit nil
### GetBoundInterface

`func (o *NetworkVfc) GetBoundInterface() InventoryInterfaceRelationship`

GetBoundInterface returns the BoundInterface field if non-nil, zero value otherwise.

### GetBoundInterfaceOk

`func (o *NetworkVfc) GetBoundInterfaceOk() (*InventoryInterfaceRelationship, bool)`

GetBoundInterfaceOk returns a tuple with the BoundInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundInterface

`func (o *NetworkVfc) SetBoundInterface(v InventoryInterfaceRelationship)`

SetBoundInterface sets BoundInterface field to given value.

### HasBoundInterface

`func (o *NetworkVfc) HasBoundInterface() bool`

HasBoundInterface returns a boolean if a field has been set.

### SetBoundInterfaceNil

`func (o *NetworkVfc) SetBoundInterfaceNil(b bool)`

 SetBoundInterfaceNil sets the value for BoundInterface to be an explicit nil

### UnsetBoundInterface
`func (o *NetworkVfc) UnsetBoundInterface()`

UnsetBoundInterface ensures that no value is present for BoundInterface, not even an explicit nil
### GetNetworkElement

`func (o *NetworkVfc) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *NetworkVfc) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *NetworkVfc) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *NetworkVfc) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### SetNetworkElementNil

`func (o *NetworkVfc) SetNetworkElementNil(b bool)`

 SetNetworkElementNil sets the value for NetworkElement to be an explicit nil

### UnsetNetworkElement
`func (o *NetworkVfc) UnsetNetworkElement()`

UnsetNetworkElement ensures that no value is present for NetworkElement, not even an explicit nil
### GetPinnedInterface

`func (o *NetworkVfc) GetPinnedInterface() InventoryInterfaceRelationship`

GetPinnedInterface returns the PinnedInterface field if non-nil, zero value otherwise.

### GetPinnedInterfaceOk

`func (o *NetworkVfc) GetPinnedInterfaceOk() (*InventoryInterfaceRelationship, bool)`

GetPinnedInterfaceOk returns a tuple with the PinnedInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinnedInterface

`func (o *NetworkVfc) SetPinnedInterface(v InventoryInterfaceRelationship)`

SetPinnedInterface sets PinnedInterface field to given value.

### HasPinnedInterface

`func (o *NetworkVfc) HasPinnedInterface() bool`

HasPinnedInterface returns a boolean if a field has been set.

### SetPinnedInterfaceNil

`func (o *NetworkVfc) SetPinnedInterfaceNil(b bool)`

 SetPinnedInterfaceNil sets the value for PinnedInterface to be an explicit nil

### UnsetPinnedInterface
`func (o *NetworkVfc) UnsetPinnedInterface()`

UnsetPinnedInterface ensures that no value is present for PinnedInterface, not even an explicit nil
### GetRegisteredDevice

`func (o *NetworkVfc) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NetworkVfc) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NetworkVfc) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NetworkVfc) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *NetworkVfc) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *NetworkVfc) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


