# FabricVlanSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.VlanSettings"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.VlanSettings"]
**AllowedVlans** | Pointer to **string** | Allowed VLAN IDs of the virtual interface. A list of comma separated VLAN ids and/or VLAN id ranges. | [optional] 
**NativeVlan** | Pointer to **int64** | Native VLAN ID of the virtual interface or the corresponding Vethernet on the peer Fabric Interconnect to which the virtual interface is connected. Native VLAN ID maps all incoming untagged traffic i.e. packets without a VLAN tag to the native VLAN for switching purposes. If the native VLAN is not a part of the allowed VLANs, it will automatically be added to the list of allowed VLANs. A native VLAN ID of 0 will indicate to the system to use the system default native VLAN ID and will also prevent native VLAN from being added to the allowed VLAN list. | [optional] [default to 0]
**QinqEnabled** | Pointer to **bool** | Enable QinQ (802.1Q-in-802.1Q) Tunneling on the vNIC. | [optional] [default to false]
**QinqVlan** | Pointer to **int64** | Select the VLAN ID for VIC QinQ (802.1Q-in-802.1Q) Tunneling. | [optional] [default to 2]

## Methods

### NewFabricVlanSettings

`func NewFabricVlanSettings(classId string, objectType string, ) *FabricVlanSettings`

NewFabricVlanSettings instantiates a new FabricVlanSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricVlanSettingsWithDefaults

`func NewFabricVlanSettingsWithDefaults() *FabricVlanSettings`

NewFabricVlanSettingsWithDefaults instantiates a new FabricVlanSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricVlanSettings) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricVlanSettings) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricVlanSettings) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricVlanSettings) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricVlanSettings) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricVlanSettings) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAllowedVlans

`func (o *FabricVlanSettings) GetAllowedVlans() string`

GetAllowedVlans returns the AllowedVlans field if non-nil, zero value otherwise.

### GetAllowedVlansOk

`func (o *FabricVlanSettings) GetAllowedVlansOk() (*string, bool)`

GetAllowedVlansOk returns a tuple with the AllowedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedVlans

`func (o *FabricVlanSettings) SetAllowedVlans(v string)`

SetAllowedVlans sets AllowedVlans field to given value.

### HasAllowedVlans

`func (o *FabricVlanSettings) HasAllowedVlans() bool`

HasAllowedVlans returns a boolean if a field has been set.

### GetNativeVlan

`func (o *FabricVlanSettings) GetNativeVlan() int64`

GetNativeVlan returns the NativeVlan field if non-nil, zero value otherwise.

### GetNativeVlanOk

`func (o *FabricVlanSettings) GetNativeVlanOk() (*int64, bool)`

GetNativeVlanOk returns a tuple with the NativeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeVlan

`func (o *FabricVlanSettings) SetNativeVlan(v int64)`

SetNativeVlan sets NativeVlan field to given value.

### HasNativeVlan

`func (o *FabricVlanSettings) HasNativeVlan() bool`

HasNativeVlan returns a boolean if a field has been set.

### GetQinqEnabled

`func (o *FabricVlanSettings) GetQinqEnabled() bool`

GetQinqEnabled returns the QinqEnabled field if non-nil, zero value otherwise.

### GetQinqEnabledOk

`func (o *FabricVlanSettings) GetQinqEnabledOk() (*bool, bool)`

GetQinqEnabledOk returns a tuple with the QinqEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQinqEnabled

`func (o *FabricVlanSettings) SetQinqEnabled(v bool)`

SetQinqEnabled sets QinqEnabled field to given value.

### HasQinqEnabled

`func (o *FabricVlanSettings) HasQinqEnabled() bool`

HasQinqEnabled returns a boolean if a field has been set.

### GetQinqVlan

`func (o *FabricVlanSettings) GetQinqVlan() int64`

GetQinqVlan returns the QinqVlan field if non-nil, zero value otherwise.

### GetQinqVlanOk

`func (o *FabricVlanSettings) GetQinqVlanOk() (*int64, bool)`

GetQinqVlanOk returns a tuple with the QinqVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQinqVlan

`func (o *FabricVlanSettings) SetQinqVlan(v int64)`

SetQinqVlan sets QinqVlan field to given value.

### HasQinqVlan

`func (o *FabricVlanSettings) HasQinqVlan() bool`

HasQinqVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


