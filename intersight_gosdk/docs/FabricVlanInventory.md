# FabricVlanInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.VlanInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.VlanInventory"]
**AdminState** | Pointer to **string** | Admin State of the VLAN attribute in the switch. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the VLAN Inventory for the switch. | [optional] [readonly] 
**Port** | Pointer to **string** | Ports associated to the VLAN. | [optional] [readonly] 
**Status** | Pointer to **string** | Status of the VLAN attribute in the switch. | [optional] [readonly] 
**VlanId** | Pointer to **string** | The Identifier for VLAN in the switch. | [optional] [readonly] 
**NetworkElement** | Pointer to [**NetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewFabricVlanInventory

`func NewFabricVlanInventory(classId string, objectType string, ) *FabricVlanInventory`

NewFabricVlanInventory instantiates a new FabricVlanInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricVlanInventoryWithDefaults

`func NewFabricVlanInventoryWithDefaults() *FabricVlanInventory`

NewFabricVlanInventoryWithDefaults instantiates a new FabricVlanInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricVlanInventory) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricVlanInventory) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricVlanInventory) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricVlanInventory) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricVlanInventory) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricVlanInventory) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminState

`func (o *FabricVlanInventory) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *FabricVlanInventory) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *FabricVlanInventory) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *FabricVlanInventory) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetName

`func (o *FabricVlanInventory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FabricVlanInventory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FabricVlanInventory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FabricVlanInventory) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *FabricVlanInventory) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *FabricVlanInventory) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *FabricVlanInventory) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *FabricVlanInventory) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetStatus

`func (o *FabricVlanInventory) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FabricVlanInventory) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FabricVlanInventory) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FabricVlanInventory) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVlanId

`func (o *FabricVlanInventory) GetVlanId() string`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *FabricVlanInventory) GetVlanIdOk() (*string, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *FabricVlanInventory) SetVlanId(v string)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *FabricVlanInventory) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetNetworkElement

`func (o *FabricVlanInventory) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *FabricVlanInventory) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *FabricVlanInventory) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *FabricVlanInventory) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *FabricVlanInventory) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *FabricVlanInventory) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *FabricVlanInventory) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *FabricVlanInventory) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


