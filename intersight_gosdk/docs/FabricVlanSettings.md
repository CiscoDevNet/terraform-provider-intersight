# FabricVlanSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedVlans** | Pointer to **string** | Allowed VLAN IDs of the virtual interface. | [optional] 
**NativeVlan** | Pointer to **int64** | Native VLAN ID of the virtual interface or the corresponding vethernet on the peer Fabric Interconnect to which the virtual interface is connected. Setting the ID to 0 will not associate any native VLAN to the traffic on the virtual interface. | [optional] 

## Methods

### NewFabricVlanSettings

`func NewFabricVlanSettings() *FabricVlanSettings`

NewFabricVlanSettings instantiates a new FabricVlanSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricVlanSettingsWithDefaults

`func NewFabricVlanSettingsWithDefaults() *FabricVlanSettings`

NewFabricVlanSettingsWithDefaults instantiates a new FabricVlanSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


