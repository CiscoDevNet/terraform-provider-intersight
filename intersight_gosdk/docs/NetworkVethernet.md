# NetworkVethernet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "network.Vethernet"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "network.Vethernet"]
**BoundInterfaceDn** | Pointer to **string** | Port or portchannel interface configured for vitual ethernet interface. | [optional] [readonly] 
**Burst** | Pointer to **int64** | Burst defined on QoS policy. | [optional] [readonly] 
**Description** | Pointer to **string** | Description for the vNIC. | [optional] [readonly] 
**OperReason** | Pointer to **string** | Operational Reason of the virtual etherent interface on the Fabric Interconnect. When the operational state is down, Operreason indicates the reason why the vNIC is not operational. Some common reasons are admindown, error disabled. | [optional] [readonly] 
**OperState** | Pointer to **string** | The operational state of the Active Vethernet peer of a vNIC in Intersight Managed Mode. This state is updated by events from Fabric Interconnect or by periodic updates from Fabric Interconnect. When a Fabric Interconnect is not connected to Intersight or if the Fabric Interconnect is powered down, this property will not be updated. * &#x60;unknown&#x60; - The operational state of the Vethernet is not known. * &#x60;adminDown&#x60; - The operational state of the Vethernet is admin down. * &#x60;up&#x60; - The operational state of the Vethernet is Up. * &#x60;down&#x60; - The operational state of the Vethernet is Down. * &#x60;noLicense&#x60; - The operational state of the Vethernet is no license. * &#x60;linkUp&#x60; - The operational state of the Vethernet is link up. * &#x60;hardwareFailure&#x60; - The operational state of the Vethernet is hardware failure. * &#x60;softwareFailure&#x60; - The operational state of the Vethernet is software failure. * &#x60;errorDisabled&#x60; - The operational state of the Vethernet is error disabled. * &#x60;linkDown&#x60; - The operational state of the Vethernet is link down. * &#x60;sfpNotPresent&#x60; - The operational state of the Vethernet is SFP not present. * &#x60;udldAggrDown&#x60; - The operational state of the Vethernet is UDLD aggregate down. | [optional] [readonly] [default to "unknown"]
**PinnedInterfaceDn** | Pointer to **string** | Uplink port or portchannel interface pinned to a vitual ethernet interface. | [optional] [readonly] 
**Ratelimit** | Pointer to **int64** | Rate limit defined on QoS policy. | [optional] [readonly] 
**VethId** | Pointer to **int64** | Vethernet Identifier on a Fabric Interconnect. | [optional] [readonly] 
**AdapterHostEthInterface** | Pointer to [**NullableAdapterHostEthInterfaceRelationship**](AdapterHostEthInterfaceRelationship.md) |  | [optional] 
**BoundInterface** | Pointer to [**NullableInventoryInterfaceRelationship**](InventoryInterfaceRelationship.md) |  | [optional] 
**NetworkElement** | Pointer to [**NullableNetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**PinnedInterface** | Pointer to [**NullableInventoryInterfaceRelationship**](InventoryInterfaceRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNetworkVethernet

`func NewNetworkVethernet(classId string, objectType string, ) *NetworkVethernet`

NewNetworkVethernet instantiates a new NetworkVethernet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkVethernetWithDefaults

`func NewNetworkVethernetWithDefaults() *NetworkVethernet`

NewNetworkVethernetWithDefaults instantiates a new NetworkVethernet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NetworkVethernet) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NetworkVethernet) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NetworkVethernet) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NetworkVethernet) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NetworkVethernet) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NetworkVethernet) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBoundInterfaceDn

`func (o *NetworkVethernet) GetBoundInterfaceDn() string`

GetBoundInterfaceDn returns the BoundInterfaceDn field if non-nil, zero value otherwise.

### GetBoundInterfaceDnOk

`func (o *NetworkVethernet) GetBoundInterfaceDnOk() (*string, bool)`

GetBoundInterfaceDnOk returns a tuple with the BoundInterfaceDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundInterfaceDn

`func (o *NetworkVethernet) SetBoundInterfaceDn(v string)`

SetBoundInterfaceDn sets BoundInterfaceDn field to given value.

### HasBoundInterfaceDn

`func (o *NetworkVethernet) HasBoundInterfaceDn() bool`

HasBoundInterfaceDn returns a boolean if a field has been set.

### GetBurst

`func (o *NetworkVethernet) GetBurst() int64`

GetBurst returns the Burst field if non-nil, zero value otherwise.

### GetBurstOk

`func (o *NetworkVethernet) GetBurstOk() (*int64, bool)`

GetBurstOk returns a tuple with the Burst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurst

`func (o *NetworkVethernet) SetBurst(v int64)`

SetBurst sets Burst field to given value.

### HasBurst

`func (o *NetworkVethernet) HasBurst() bool`

HasBurst returns a boolean if a field has been set.

### GetDescription

`func (o *NetworkVethernet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkVethernet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkVethernet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkVethernet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOperReason

`func (o *NetworkVethernet) GetOperReason() string`

GetOperReason returns the OperReason field if non-nil, zero value otherwise.

### GetOperReasonOk

`func (o *NetworkVethernet) GetOperReasonOk() (*string, bool)`

GetOperReasonOk returns a tuple with the OperReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperReason

`func (o *NetworkVethernet) SetOperReason(v string)`

SetOperReason sets OperReason field to given value.

### HasOperReason

`func (o *NetworkVethernet) HasOperReason() bool`

HasOperReason returns a boolean if a field has been set.

### GetOperState

`func (o *NetworkVethernet) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *NetworkVethernet) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *NetworkVethernet) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *NetworkVethernet) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetPinnedInterfaceDn

`func (o *NetworkVethernet) GetPinnedInterfaceDn() string`

GetPinnedInterfaceDn returns the PinnedInterfaceDn field if non-nil, zero value otherwise.

### GetPinnedInterfaceDnOk

`func (o *NetworkVethernet) GetPinnedInterfaceDnOk() (*string, bool)`

GetPinnedInterfaceDnOk returns a tuple with the PinnedInterfaceDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinnedInterfaceDn

`func (o *NetworkVethernet) SetPinnedInterfaceDn(v string)`

SetPinnedInterfaceDn sets PinnedInterfaceDn field to given value.

### HasPinnedInterfaceDn

`func (o *NetworkVethernet) HasPinnedInterfaceDn() bool`

HasPinnedInterfaceDn returns a boolean if a field has been set.

### GetRatelimit

`func (o *NetworkVethernet) GetRatelimit() int64`

GetRatelimit returns the Ratelimit field if non-nil, zero value otherwise.

### GetRatelimitOk

`func (o *NetworkVethernet) GetRatelimitOk() (*int64, bool)`

GetRatelimitOk returns a tuple with the Ratelimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatelimit

`func (o *NetworkVethernet) SetRatelimit(v int64)`

SetRatelimit sets Ratelimit field to given value.

### HasRatelimit

`func (o *NetworkVethernet) HasRatelimit() bool`

HasRatelimit returns a boolean if a field has been set.

### GetVethId

`func (o *NetworkVethernet) GetVethId() int64`

GetVethId returns the VethId field if non-nil, zero value otherwise.

### GetVethIdOk

`func (o *NetworkVethernet) GetVethIdOk() (*int64, bool)`

GetVethIdOk returns a tuple with the VethId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVethId

`func (o *NetworkVethernet) SetVethId(v int64)`

SetVethId sets VethId field to given value.

### HasVethId

`func (o *NetworkVethernet) HasVethId() bool`

HasVethId returns a boolean if a field has been set.

### GetAdapterHostEthInterface

`func (o *NetworkVethernet) GetAdapterHostEthInterface() AdapterHostEthInterfaceRelationship`

GetAdapterHostEthInterface returns the AdapterHostEthInterface field if non-nil, zero value otherwise.

### GetAdapterHostEthInterfaceOk

`func (o *NetworkVethernet) GetAdapterHostEthInterfaceOk() (*AdapterHostEthInterfaceRelationship, bool)`

GetAdapterHostEthInterfaceOk returns a tuple with the AdapterHostEthInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapterHostEthInterface

`func (o *NetworkVethernet) SetAdapterHostEthInterface(v AdapterHostEthInterfaceRelationship)`

SetAdapterHostEthInterface sets AdapterHostEthInterface field to given value.

### HasAdapterHostEthInterface

`func (o *NetworkVethernet) HasAdapterHostEthInterface() bool`

HasAdapterHostEthInterface returns a boolean if a field has been set.

### SetAdapterHostEthInterfaceNil

`func (o *NetworkVethernet) SetAdapterHostEthInterfaceNil(b bool)`

 SetAdapterHostEthInterfaceNil sets the value for AdapterHostEthInterface to be an explicit nil

### UnsetAdapterHostEthInterface
`func (o *NetworkVethernet) UnsetAdapterHostEthInterface()`

UnsetAdapterHostEthInterface ensures that no value is present for AdapterHostEthInterface, not even an explicit nil
### GetBoundInterface

`func (o *NetworkVethernet) GetBoundInterface() InventoryInterfaceRelationship`

GetBoundInterface returns the BoundInterface field if non-nil, zero value otherwise.

### GetBoundInterfaceOk

`func (o *NetworkVethernet) GetBoundInterfaceOk() (*InventoryInterfaceRelationship, bool)`

GetBoundInterfaceOk returns a tuple with the BoundInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundInterface

`func (o *NetworkVethernet) SetBoundInterface(v InventoryInterfaceRelationship)`

SetBoundInterface sets BoundInterface field to given value.

### HasBoundInterface

`func (o *NetworkVethernet) HasBoundInterface() bool`

HasBoundInterface returns a boolean if a field has been set.

### SetBoundInterfaceNil

`func (o *NetworkVethernet) SetBoundInterfaceNil(b bool)`

 SetBoundInterfaceNil sets the value for BoundInterface to be an explicit nil

### UnsetBoundInterface
`func (o *NetworkVethernet) UnsetBoundInterface()`

UnsetBoundInterface ensures that no value is present for BoundInterface, not even an explicit nil
### GetNetworkElement

`func (o *NetworkVethernet) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *NetworkVethernet) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *NetworkVethernet) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *NetworkVethernet) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### SetNetworkElementNil

`func (o *NetworkVethernet) SetNetworkElementNil(b bool)`

 SetNetworkElementNil sets the value for NetworkElement to be an explicit nil

### UnsetNetworkElement
`func (o *NetworkVethernet) UnsetNetworkElement()`

UnsetNetworkElement ensures that no value is present for NetworkElement, not even an explicit nil
### GetPinnedInterface

`func (o *NetworkVethernet) GetPinnedInterface() InventoryInterfaceRelationship`

GetPinnedInterface returns the PinnedInterface field if non-nil, zero value otherwise.

### GetPinnedInterfaceOk

`func (o *NetworkVethernet) GetPinnedInterfaceOk() (*InventoryInterfaceRelationship, bool)`

GetPinnedInterfaceOk returns a tuple with the PinnedInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinnedInterface

`func (o *NetworkVethernet) SetPinnedInterface(v InventoryInterfaceRelationship)`

SetPinnedInterface sets PinnedInterface field to given value.

### HasPinnedInterface

`func (o *NetworkVethernet) HasPinnedInterface() bool`

HasPinnedInterface returns a boolean if a field has been set.

### SetPinnedInterfaceNil

`func (o *NetworkVethernet) SetPinnedInterfaceNil(b bool)`

 SetPinnedInterfaceNil sets the value for PinnedInterface to be an explicit nil

### UnsetPinnedInterface
`func (o *NetworkVethernet) UnsetPinnedInterface()`

UnsetPinnedInterface ensures that no value is present for PinnedInterface, not even an explicit nil
### GetRegisteredDevice

`func (o *NetworkVethernet) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NetworkVethernet) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NetworkVethernet) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NetworkVethernet) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *NetworkVethernet) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *NetworkVethernet) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


