# PciPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "pci.Port"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "pci.Port"]
**Name** | Pointer to **string** | The name of the PCIe Switch Port. Name will be of the format &#39;PCIe Port &lt;$id&gt;&#39;. | [optional] [readonly] 
**OperReason** | Pointer to **[]string** |  | [optional] 
**OperState** | Pointer to **string** | The operational status of the PCIe switch port. | [optional] [readonly] 
**PortId** | Pointer to **int64** | The identifier of the PCIe switch port. | [optional] [readonly] 
**Role** | Pointer to **string** | The current role of the PCIe switch port. * &#x60;Unconfigured&#x60; - Indicates that the PCIe switch port role is currently unconfigured. * &#x60;Downstream&#x60; - Indicates that the PCIe switch port role is currently downstream. A downstream port connects to a PCIe peripheral device such as a GPU or a network adapter. * &#x60;Upstream&#x60; - Indicates that the PCIe switch port role is currently upstream. An upstream port connects to a PCIe root complex such as a CPU. * &#x60;Unknown&#x60; - Indicates that the PCIe switch port role is currently unknown. | [optional] [readonly] [default to "Unconfigured"]
**Uri** | Pointer to **string** | The unique identifier of the PCIe switch port as reported by chassis expander module management controller. | [optional] 
**Width** | Pointer to **int64** | The link width of the PCIe switch port. | [optional] [readonly] 
**InventoryDeviceInfo** | Pointer to [**NullableInventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**PciEndpoint** | Pointer to [**NullablePciEndpointRelationship**](PciEndpointRelationship.md) |  | [optional] 
**PciSwitch** | Pointer to [**NullablePciSwitchRelationship**](PciSwitchRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewPciPort

`func NewPciPort(classId string, objectType string, ) *PciPort`

NewPciPort instantiates a new PciPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPciPortWithDefaults

`func NewPciPortWithDefaults() *PciPort`

NewPciPortWithDefaults instantiates a new PciPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PciPort) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PciPort) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PciPort) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PciPort) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PciPort) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PciPort) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *PciPort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PciPort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PciPort) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PciPort) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperReason

`func (o *PciPort) GetOperReason() []string`

GetOperReason returns the OperReason field if non-nil, zero value otherwise.

### GetOperReasonOk

`func (o *PciPort) GetOperReasonOk() (*[]string, bool)`

GetOperReasonOk returns a tuple with the OperReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperReason

`func (o *PciPort) SetOperReason(v []string)`

SetOperReason sets OperReason field to given value.

### HasOperReason

`func (o *PciPort) HasOperReason() bool`

HasOperReason returns a boolean if a field has been set.

### SetOperReasonNil

`func (o *PciPort) SetOperReasonNil(b bool)`

 SetOperReasonNil sets the value for OperReason to be an explicit nil

### UnsetOperReason
`func (o *PciPort) UnsetOperReason()`

UnsetOperReason ensures that no value is present for OperReason, not even an explicit nil
### GetOperState

`func (o *PciPort) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *PciPort) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *PciPort) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *PciPort) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetPortId

`func (o *PciPort) GetPortId() int64`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *PciPort) GetPortIdOk() (*int64, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *PciPort) SetPortId(v int64)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *PciPort) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetRole

`func (o *PciPort) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PciPort) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PciPort) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *PciPort) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetUri

`func (o *PciPort) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *PciPort) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *PciPort) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *PciPort) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetWidth

`func (o *PciPort) GetWidth() int64`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *PciPort) GetWidthOk() (*int64, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *PciPort) SetWidth(v int64)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *PciPort) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *PciPort) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *PciPort) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *PciPort) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *PciPort) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### SetInventoryDeviceInfoNil

`func (o *PciPort) SetInventoryDeviceInfoNil(b bool)`

 SetInventoryDeviceInfoNil sets the value for InventoryDeviceInfo to be an explicit nil

### UnsetInventoryDeviceInfo
`func (o *PciPort) UnsetInventoryDeviceInfo()`

UnsetInventoryDeviceInfo ensures that no value is present for InventoryDeviceInfo, not even an explicit nil
### GetPciEndpoint

`func (o *PciPort) GetPciEndpoint() PciEndpointRelationship`

GetPciEndpoint returns the PciEndpoint field if non-nil, zero value otherwise.

### GetPciEndpointOk

`func (o *PciPort) GetPciEndpointOk() (*PciEndpointRelationship, bool)`

GetPciEndpointOk returns a tuple with the PciEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciEndpoint

`func (o *PciPort) SetPciEndpoint(v PciEndpointRelationship)`

SetPciEndpoint sets PciEndpoint field to given value.

### HasPciEndpoint

`func (o *PciPort) HasPciEndpoint() bool`

HasPciEndpoint returns a boolean if a field has been set.

### SetPciEndpointNil

`func (o *PciPort) SetPciEndpointNil(b bool)`

 SetPciEndpointNil sets the value for PciEndpoint to be an explicit nil

### UnsetPciEndpoint
`func (o *PciPort) UnsetPciEndpoint()`

UnsetPciEndpoint ensures that no value is present for PciEndpoint, not even an explicit nil
### GetPciSwitch

`func (o *PciPort) GetPciSwitch() PciSwitchRelationship`

GetPciSwitch returns the PciSwitch field if non-nil, zero value otherwise.

### GetPciSwitchOk

`func (o *PciPort) GetPciSwitchOk() (*PciSwitchRelationship, bool)`

GetPciSwitchOk returns a tuple with the PciSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciSwitch

`func (o *PciPort) SetPciSwitch(v PciSwitchRelationship)`

SetPciSwitch sets PciSwitch field to given value.

### HasPciSwitch

`func (o *PciPort) HasPciSwitch() bool`

HasPciSwitch returns a boolean if a field has been set.

### SetPciSwitchNil

`func (o *PciPort) SetPciSwitchNil(b bool)`

 SetPciSwitchNil sets the value for PciSwitch to be an explicit nil

### UnsetPciSwitch
`func (o *PciPort) UnsetPciSwitch()`

UnsetPciSwitch ensures that no value is present for PciSwitch, not even an explicit nil
### GetRegisteredDevice

`func (o *PciPort) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *PciPort) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *PciPort) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *PciPort) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *PciPort) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *PciPort) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


