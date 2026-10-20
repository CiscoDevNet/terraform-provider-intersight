# FabricVsanInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.VsanInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.VsanInventory"]
**AdminState** | Pointer to **string** | Admin state of the VSAN. It can be active or inactive. * &#x60;&#x60; - Default value to represent the administrative state of isolated vsan. * &#x60;Up&#x60; - VSAN administrative state is up. * &#x60;Down&#x60; - VSAN administrative state is down. | [optional] [readonly] [default to ""]
**InteropMode** | Pointer to **string** | Interoperability mode of the VSAN. It enables products of multiple vendors to communicate with each other. * &#x60;&#x60; - Default value to represent the interoperability mode of isolated vsan when it is not configured. * &#x60;Default&#x60; - Default mode for a VSAN that is communicating between a SAN composed entirely of MDS 9000 switches. * &#x60;1&#x60; - Allows integration with Brocade and McData switches that have been configured for their own interoperability modes. Brocade and McData switches must be running in interop mode to work with this VSAN mode. * &#x60;2&#x60; - Allows seamless integration with specific Brocade switches running in their own native mode of operation. * &#x60;3&#x60; - Allows seamless integration with Brocade switches that contains more than 16 ports. * &#x60;4&#x60; - Allows seamless integration between MDS VSANs and McData switches running in McData Fabric 1.0 interop mode. | [optional] [readonly] [default to ""]
**LoadBalancing** | Pointer to **string** | These attributes indicate the use of the source-destination ID or the originator exchange OX ID for load balancing path selection. * &#x60;&#x60; - Default value to represent the load balancing scheme of isolated vsan. * &#x60;src-id/dst-id&#x60; - Directs the switch to use the source and destination ID for its path selection process. * &#x60;src-id/dst-id/oxid&#x60; - Directs the switch to use the source ID, the destination ID, and the OX ID for its path selection process. | [optional] [readonly] [default to ""]
**Name** | Pointer to **string** | User-specified name of the VSAN. | [optional] [readonly] 
**OperState** | Pointer to **string** | Operational state of the VSAN. * &#x60;&#x60; - Default value to represent the operational state of isolated vsan. * &#x60;Up&#x60; - VSAN operational state is up. * &#x60;Down&#x60; - VSAN operational state is down. | [optional] [readonly] [default to ""]
**SmartZoning** | Pointer to **string** | Smart zoning status on the VSAN. It can be enabled or disabled. * &#x60;&#x60; - Default value to represent the smart zoning status of isolated vsan. * &#x60;Enabled&#x60; - VSAN smart zoning state is enabled. * &#x60;Disabled&#x60; - VSAN smart zoning state is disabled. | [optional] [readonly] [default to ""]
**VsanId** | Pointer to **int64** | Identifier for the VSAN. It is an integer from 1 to 4094. | [optional] [readonly] 
**NetworkElement** | Pointer to [**NullableNetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewFabricVsanInventory

`func NewFabricVsanInventory(classId string, objectType string, ) *FabricVsanInventory`

NewFabricVsanInventory instantiates a new FabricVsanInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricVsanInventoryWithDefaults

`func NewFabricVsanInventoryWithDefaults() *FabricVsanInventory`

NewFabricVsanInventoryWithDefaults instantiates a new FabricVsanInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricVsanInventory) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricVsanInventory) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricVsanInventory) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricVsanInventory) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricVsanInventory) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricVsanInventory) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminState

`func (o *FabricVsanInventory) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *FabricVsanInventory) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *FabricVsanInventory) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *FabricVsanInventory) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetInteropMode

`func (o *FabricVsanInventory) GetInteropMode() string`

GetInteropMode returns the InteropMode field if non-nil, zero value otherwise.

### GetInteropModeOk

`func (o *FabricVsanInventory) GetInteropModeOk() (*string, bool)`

GetInteropModeOk returns a tuple with the InteropMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteropMode

`func (o *FabricVsanInventory) SetInteropMode(v string)`

SetInteropMode sets InteropMode field to given value.

### HasInteropMode

`func (o *FabricVsanInventory) HasInteropMode() bool`

HasInteropMode returns a boolean if a field has been set.

### GetLoadBalancing

`func (o *FabricVsanInventory) GetLoadBalancing() string`

GetLoadBalancing returns the LoadBalancing field if non-nil, zero value otherwise.

### GetLoadBalancingOk

`func (o *FabricVsanInventory) GetLoadBalancingOk() (*string, bool)`

GetLoadBalancingOk returns a tuple with the LoadBalancing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancing

`func (o *FabricVsanInventory) SetLoadBalancing(v string)`

SetLoadBalancing sets LoadBalancing field to given value.

### HasLoadBalancing

`func (o *FabricVsanInventory) HasLoadBalancing() bool`

HasLoadBalancing returns a boolean if a field has been set.

### GetName

`func (o *FabricVsanInventory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FabricVsanInventory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FabricVsanInventory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FabricVsanInventory) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperState

`func (o *FabricVsanInventory) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *FabricVsanInventory) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *FabricVsanInventory) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *FabricVsanInventory) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetSmartZoning

`func (o *FabricVsanInventory) GetSmartZoning() string`

GetSmartZoning returns the SmartZoning field if non-nil, zero value otherwise.

### GetSmartZoningOk

`func (o *FabricVsanInventory) GetSmartZoningOk() (*string, bool)`

GetSmartZoningOk returns a tuple with the SmartZoning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartZoning

`func (o *FabricVsanInventory) SetSmartZoning(v string)`

SetSmartZoning sets SmartZoning field to given value.

### HasSmartZoning

`func (o *FabricVsanInventory) HasSmartZoning() bool`

HasSmartZoning returns a boolean if a field has been set.

### GetVsanId

`func (o *FabricVsanInventory) GetVsanId() int64`

GetVsanId returns the VsanId field if non-nil, zero value otherwise.

### GetVsanIdOk

`func (o *FabricVsanInventory) GetVsanIdOk() (*int64, bool)`

GetVsanIdOk returns a tuple with the VsanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsanId

`func (o *FabricVsanInventory) SetVsanId(v int64)`

SetVsanId sets VsanId field to given value.

### HasVsanId

`func (o *FabricVsanInventory) HasVsanId() bool`

HasVsanId returns a boolean if a field has been set.

### GetNetworkElement

`func (o *FabricVsanInventory) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *FabricVsanInventory) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *FabricVsanInventory) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *FabricVsanInventory) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### SetNetworkElementNil

`func (o *FabricVsanInventory) SetNetworkElementNil(b bool)`

 SetNetworkElementNil sets the value for NetworkElement to be an explicit nil

### UnsetNetworkElement
`func (o *FabricVsanInventory) UnsetNetworkElement()`

UnsetNetworkElement ensures that no value is present for NetworkElement, not even an explicit nil
### GetRegisteredDevice

`func (o *FabricVsanInventory) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *FabricVsanInventory) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *FabricVsanInventory) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *FabricVsanInventory) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *FabricVsanInventory) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *FabricVsanInventory) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


