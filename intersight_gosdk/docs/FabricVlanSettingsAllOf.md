# FabricVlanSettingsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedVlans** | Pointer to **string** | Allowed VLAN IDs of the virtual interface. | [optional] 
**NativeVlan** | Pointer to **int64** | Native VLAN ID of the virtual interface or the corresponding vethernet on the peer Fabric Interconnect to which the virtual interface is connected. Setting the ID to 0 will not associate any native VLAN to the traffic on the virtual interface. | [optional] 

## Methods

### NewFabricVlanSettingsAllOf

`func NewFabricVlanSettingsAllOf() *FabricVlanSettingsAllOf`

NewFabricVlanSettingsAllOf instantiates a new FabricVlanSettingsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricVlanSettingsAllOfWithDefaults

`func NewFabricVlanSettingsAllOfWithDefaults() *FabricVlanSettingsAllOf`

NewFabricVlanSettingsAllOfWithDefaults instantiates a new FabricVlanSettingsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedVlans

`func (o *FabricVlanSettingsAllOf) GetAllowedVlans() string`

GetAllowedVlans returns the AllowedVlans field if non-nil, zero value otherwise.

### GetAllowedVlansOk

`func (o *FabricVlanSettingsAllOf) GetAllowedVlansOk() (*string, bool)`

GetAllowedVlansOk returns a tuple with the AllowedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedVlans

`func (o *FabricVlanSettingsAllOf) SetAllowedVlans(v string)`

SetAllowedVlans sets AllowedVlans field to given value.

### HasAllowedVlans

`func (o *FabricVlanSettingsAllOf) HasAllowedVlans() bool`

HasAllowedVlans returns a boolean if a field has been set.

### GetNativeVlan

`func (o *FabricVlanSettingsAllOf) GetNativeVlan() int64`

GetNativeVlan returns the NativeVlan field if non-nil, zero value otherwise.

### GetNativeVlanOk

`func (o *FabricVlanSettingsAllOf) GetNativeVlanOk() (*int64, bool)`

GetNativeVlanOk returns a tuple with the NativeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeVlan

`func (o *FabricVlanSettingsAllOf) SetNativeVlan(v int64)`

SetNativeVlan sets NativeVlan field to given value.

### HasNativeVlan

`func (o *FabricVlanSettingsAllOf) HasNativeVlan() bool`

HasNativeVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


