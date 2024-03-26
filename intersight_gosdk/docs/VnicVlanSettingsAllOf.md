# VnicVlanSettingsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.VlanSettings"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.VlanSettings"]
**AllowedVlans** | Pointer to **string** | Allowed VLAN IDs of the virtual interface. | [optional] 
**DefaultVlan** | Pointer to **int64** | Native VLAN ID of the virtual interface or the corresponding vethernet on the peer Fabric Interconnect to which the virtual interface is connected. Setting the ID to 0 will not associate any native VLAN to the traffic on the virtual interface. | [optional] [default to 0]
**Mode** | Pointer to **string** | Option to determine if the port can carry single VLAN (Access) or multiple VLANs (Trunk) traffic. * &#x60;ACCESS&#x60; - An access port carries traffic only for a single VLAN on the interface. * &#x60;TRUNK&#x60; - A trunk port can have two or more VLANs configured on the interface. It can carry traffic for several VLANs simultaneously. | [optional] [default to "ACCESS"]
**QinqEnabled** | Pointer to **bool** | Enable QinQ (802.1Q-in-802.1Q) Tunneling on the vNIC. | [optional] [default to false]
**QinqVlan** | Pointer to **int64** | When activating VIC QinQ (802.1Q-in-802.1Q) Tunneling, a particular VLAN ID is set. In Access VLAN mode, this QinQ VLAN ID is established as the default VLAN. | [optional] [default to 2]

## Methods

### NewVnicVlanSettingsAllOf

`func NewVnicVlanSettingsAllOf(classId string, objectType string, ) *VnicVlanSettingsAllOf`

NewVnicVlanSettingsAllOf instantiates a new VnicVlanSettingsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicVlanSettingsAllOfWithDefaults

`func NewVnicVlanSettingsAllOfWithDefaults() *VnicVlanSettingsAllOf`

NewVnicVlanSettingsAllOfWithDefaults instantiates a new VnicVlanSettingsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicVlanSettingsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicVlanSettingsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicVlanSettingsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicVlanSettingsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicVlanSettingsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicVlanSettingsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAllowedVlans

`func (o *VnicVlanSettingsAllOf) GetAllowedVlans() string`

GetAllowedVlans returns the AllowedVlans field if non-nil, zero value otherwise.

### GetAllowedVlansOk

`func (o *VnicVlanSettingsAllOf) GetAllowedVlansOk() (*string, bool)`

GetAllowedVlansOk returns a tuple with the AllowedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedVlans

`func (o *VnicVlanSettingsAllOf) SetAllowedVlans(v string)`

SetAllowedVlans sets AllowedVlans field to given value.

### HasAllowedVlans

`func (o *VnicVlanSettingsAllOf) HasAllowedVlans() bool`

HasAllowedVlans returns a boolean if a field has been set.

### GetDefaultVlan

`func (o *VnicVlanSettingsAllOf) GetDefaultVlan() int64`

GetDefaultVlan returns the DefaultVlan field if non-nil, zero value otherwise.

### GetDefaultVlanOk

`func (o *VnicVlanSettingsAllOf) GetDefaultVlanOk() (*int64, bool)`

GetDefaultVlanOk returns a tuple with the DefaultVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVlan

`func (o *VnicVlanSettingsAllOf) SetDefaultVlan(v int64)`

SetDefaultVlan sets DefaultVlan field to given value.

### HasDefaultVlan

`func (o *VnicVlanSettingsAllOf) HasDefaultVlan() bool`

HasDefaultVlan returns a boolean if a field has been set.

### GetMode

`func (o *VnicVlanSettingsAllOf) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *VnicVlanSettingsAllOf) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *VnicVlanSettingsAllOf) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *VnicVlanSettingsAllOf) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetQinqEnabled

`func (o *VnicVlanSettingsAllOf) GetQinqEnabled() bool`

GetQinqEnabled returns the QinqEnabled field if non-nil, zero value otherwise.

### GetQinqEnabledOk

`func (o *VnicVlanSettingsAllOf) GetQinqEnabledOk() (*bool, bool)`

GetQinqEnabledOk returns a tuple with the QinqEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQinqEnabled

`func (o *VnicVlanSettingsAllOf) SetQinqEnabled(v bool)`

SetQinqEnabled sets QinqEnabled field to given value.

### HasQinqEnabled

`func (o *VnicVlanSettingsAllOf) HasQinqEnabled() bool`

HasQinqEnabled returns a boolean if a field has been set.

### GetQinqVlan

`func (o *VnicVlanSettingsAllOf) GetQinqVlan() int64`

GetQinqVlan returns the QinqVlan field if non-nil, zero value otherwise.

### GetQinqVlanOk

`func (o *VnicVlanSettingsAllOf) GetQinqVlanOk() (*int64, bool)`

GetQinqVlanOk returns a tuple with the QinqVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQinqVlan

`func (o *VnicVlanSettingsAllOf) SetQinqVlan(v int64)`

SetQinqVlan sets QinqVlan field to given value.

### HasQinqVlan

`func (o *VnicVlanSettingsAllOf) HasQinqVlan() bool`

HasQinqVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


