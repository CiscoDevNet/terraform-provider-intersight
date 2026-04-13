# FabricVlanExportInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.VlanExportInterface"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.VlanExportInterface"]
**SwitchAip** | Pointer to **string** | Switch IP for the export interface. | [optional] 
**SwitchAsubnetmask** | Pointer to **string** | Subnet mask for the export interface network. | [optional] 
**SwitchBip** | Pointer to **string** | Switch IP for the export interface. | [optional] 
**SwitchBsubnetmask** | Pointer to **string** | Subnet mask for the export interface network. | [optional] 
**VlanId** | Pointer to **int64** | VLAN ID for the export interface. | [optional] 

## Methods

### NewFabricVlanExportInterface

`func NewFabricVlanExportInterface(classId string, objectType string, ) *FabricVlanExportInterface`

NewFabricVlanExportInterface instantiates a new FabricVlanExportInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricVlanExportInterfaceWithDefaults

`func NewFabricVlanExportInterfaceWithDefaults() *FabricVlanExportInterface`

NewFabricVlanExportInterfaceWithDefaults instantiates a new FabricVlanExportInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricVlanExportInterface) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricVlanExportInterface) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricVlanExportInterface) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricVlanExportInterface) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricVlanExportInterface) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricVlanExportInterface) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSwitchAip

`func (o *FabricVlanExportInterface) GetSwitchAip() string`

GetSwitchAip returns the SwitchAip field if non-nil, zero value otherwise.

### GetSwitchAipOk

`func (o *FabricVlanExportInterface) GetSwitchAipOk() (*string, bool)`

GetSwitchAipOk returns a tuple with the SwitchAip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchAip

`func (o *FabricVlanExportInterface) SetSwitchAip(v string)`

SetSwitchAip sets SwitchAip field to given value.

### HasSwitchAip

`func (o *FabricVlanExportInterface) HasSwitchAip() bool`

HasSwitchAip returns a boolean if a field has been set.

### GetSwitchAsubnetmask

`func (o *FabricVlanExportInterface) GetSwitchAsubnetmask() string`

GetSwitchAsubnetmask returns the SwitchAsubnetmask field if non-nil, zero value otherwise.

### GetSwitchAsubnetmaskOk

`func (o *FabricVlanExportInterface) GetSwitchAsubnetmaskOk() (*string, bool)`

GetSwitchAsubnetmaskOk returns a tuple with the SwitchAsubnetmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchAsubnetmask

`func (o *FabricVlanExportInterface) SetSwitchAsubnetmask(v string)`

SetSwitchAsubnetmask sets SwitchAsubnetmask field to given value.

### HasSwitchAsubnetmask

`func (o *FabricVlanExportInterface) HasSwitchAsubnetmask() bool`

HasSwitchAsubnetmask returns a boolean if a field has been set.

### GetSwitchBip

`func (o *FabricVlanExportInterface) GetSwitchBip() string`

GetSwitchBip returns the SwitchBip field if non-nil, zero value otherwise.

### GetSwitchBipOk

`func (o *FabricVlanExportInterface) GetSwitchBipOk() (*string, bool)`

GetSwitchBipOk returns a tuple with the SwitchBip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchBip

`func (o *FabricVlanExportInterface) SetSwitchBip(v string)`

SetSwitchBip sets SwitchBip field to given value.

### HasSwitchBip

`func (o *FabricVlanExportInterface) HasSwitchBip() bool`

HasSwitchBip returns a boolean if a field has been set.

### GetSwitchBsubnetmask

`func (o *FabricVlanExportInterface) GetSwitchBsubnetmask() string`

GetSwitchBsubnetmask returns the SwitchBsubnetmask field if non-nil, zero value otherwise.

### GetSwitchBsubnetmaskOk

`func (o *FabricVlanExportInterface) GetSwitchBsubnetmaskOk() (*string, bool)`

GetSwitchBsubnetmaskOk returns a tuple with the SwitchBsubnetmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchBsubnetmask

`func (o *FabricVlanExportInterface) SetSwitchBsubnetmask(v string)`

SetSwitchBsubnetmask sets SwitchBsubnetmask field to given value.

### HasSwitchBsubnetmask

`func (o *FabricVlanExportInterface) HasSwitchBsubnetmask() bool`

HasSwitchBsubnetmask returns a boolean if a field has been set.

### GetVlanId

`func (o *FabricVlanExportInterface) GetVlanId() int64`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *FabricVlanExportInterface) GetVlanIdOk() (*int64, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *FabricVlanExportInterface) SetVlanId(v int64)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *FabricVlanExportInterface) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


