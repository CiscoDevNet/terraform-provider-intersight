# EtherPortChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessVlan** | Pointer to **string** | Access VLANs for this port-channel, on this FI. | [optional] 
**AdminState** | Pointer to **string** | Administratively configured state (enabled/disabled) for this port-channel. | [optional] 
**AllowedVlans** | Pointer to **string** | Allowed VLANs on this port-channel, on this FI. | [optional] 
**Mode** | Pointer to **string** | Operating mode of this port-channel. | [optional] 
**NativeVlan** | Pointer to **string** | Native VLAN for this port-channel, on this FI. | [optional] 
**OperSpeed** | Pointer to **string** | Operational speed of this port-channel. | [optional] 
**OperState** | Pointer to **string** | Operational state of this port-channel. | [optional] 
**OperStateQual** | Pointer to **string** | Reason for this port-channel&#39;s Operational state. | [optional] 
**PortChannelId** | Pointer to **int64** | Unique identifier for this port-channel on the FI. | [optional] 
**Role** | Pointer to **string** | This port-channel&#39;s configured role (uplink, server, etc.). | [optional] 
**SwitchId** | Pointer to **string** | Switch Identifier that is local to a cluster. | [optional] 
**EquipmentSwitchCard** | Pointer to [**EquipmentSwitchCardRelationship**](equipment.SwitchCard.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewEtherPortChannel

`func NewEtherPortChannel() *EtherPortChannel`

NewEtherPortChannel instantiates a new EtherPortChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEtherPortChannelWithDefaults

`func NewEtherPortChannelWithDefaults() *EtherPortChannel`

NewEtherPortChannelWithDefaults instantiates a new EtherPortChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessVlan

`func (o *EtherPortChannel) GetAccessVlan() string`

GetAccessVlan returns the AccessVlan field if non-nil, zero value otherwise.

### GetAccessVlanOk

`func (o *EtherPortChannel) GetAccessVlanOk() (*string, bool)`

GetAccessVlanOk returns a tuple with the AccessVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessVlan

`func (o *EtherPortChannel) SetAccessVlan(v string)`

SetAccessVlan sets AccessVlan field to given value.

### HasAccessVlan

`func (o *EtherPortChannel) HasAccessVlan() bool`

HasAccessVlan returns a boolean if a field has been set.

### GetAdminState

`func (o *EtherPortChannel) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *EtherPortChannel) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *EtherPortChannel) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *EtherPortChannel) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetAllowedVlans

`func (o *EtherPortChannel) GetAllowedVlans() string`

GetAllowedVlans returns the AllowedVlans field if non-nil, zero value otherwise.

### GetAllowedVlansOk

`func (o *EtherPortChannel) GetAllowedVlansOk() (*string, bool)`

GetAllowedVlansOk returns a tuple with the AllowedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedVlans

`func (o *EtherPortChannel) SetAllowedVlans(v string)`

SetAllowedVlans sets AllowedVlans field to given value.

### HasAllowedVlans

`func (o *EtherPortChannel) HasAllowedVlans() bool`

HasAllowedVlans returns a boolean if a field has been set.

### GetMode

`func (o *EtherPortChannel) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *EtherPortChannel) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *EtherPortChannel) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *EtherPortChannel) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetNativeVlan

`func (o *EtherPortChannel) GetNativeVlan() string`

GetNativeVlan returns the NativeVlan field if non-nil, zero value otherwise.

### GetNativeVlanOk

`func (o *EtherPortChannel) GetNativeVlanOk() (*string, bool)`

GetNativeVlanOk returns a tuple with the NativeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeVlan

`func (o *EtherPortChannel) SetNativeVlan(v string)`

SetNativeVlan sets NativeVlan field to given value.

### HasNativeVlan

`func (o *EtherPortChannel) HasNativeVlan() bool`

HasNativeVlan returns a boolean if a field has been set.

### GetOperSpeed

`func (o *EtherPortChannel) GetOperSpeed() string`

GetOperSpeed returns the OperSpeed field if non-nil, zero value otherwise.

### GetOperSpeedOk

`func (o *EtherPortChannel) GetOperSpeedOk() (*string, bool)`

GetOperSpeedOk returns a tuple with the OperSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperSpeed

`func (o *EtherPortChannel) SetOperSpeed(v string)`

SetOperSpeed sets OperSpeed field to given value.

### HasOperSpeed

`func (o *EtherPortChannel) HasOperSpeed() bool`

HasOperSpeed returns a boolean if a field has been set.

### GetOperState

`func (o *EtherPortChannel) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *EtherPortChannel) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *EtherPortChannel) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *EtherPortChannel) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetOperStateQual

`func (o *EtherPortChannel) GetOperStateQual() string`

GetOperStateQual returns the OperStateQual field if non-nil, zero value otherwise.

### GetOperStateQualOk

`func (o *EtherPortChannel) GetOperStateQualOk() (*string, bool)`

GetOperStateQualOk returns a tuple with the OperStateQual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperStateQual

`func (o *EtherPortChannel) SetOperStateQual(v string)`

SetOperStateQual sets OperStateQual field to given value.

### HasOperStateQual

`func (o *EtherPortChannel) HasOperStateQual() bool`

HasOperStateQual returns a boolean if a field has been set.

### GetPortChannelId

`func (o *EtherPortChannel) GetPortChannelId() int64`

GetPortChannelId returns the PortChannelId field if non-nil, zero value otherwise.

### GetPortChannelIdOk

`func (o *EtherPortChannel) GetPortChannelIdOk() (*int64, bool)`

GetPortChannelIdOk returns a tuple with the PortChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortChannelId

`func (o *EtherPortChannel) SetPortChannelId(v int64)`

SetPortChannelId sets PortChannelId field to given value.

### HasPortChannelId

`func (o *EtherPortChannel) HasPortChannelId() bool`

HasPortChannelId returns a boolean if a field has been set.

### GetRole

`func (o *EtherPortChannel) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *EtherPortChannel) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *EtherPortChannel) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *EtherPortChannel) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSwitchId

`func (o *EtherPortChannel) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *EtherPortChannel) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *EtherPortChannel) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *EtherPortChannel) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetEquipmentSwitchCard

`func (o *EtherPortChannel) GetEquipmentSwitchCard() EquipmentSwitchCardRelationship`

GetEquipmentSwitchCard returns the EquipmentSwitchCard field if non-nil, zero value otherwise.

### GetEquipmentSwitchCardOk

`func (o *EtherPortChannel) GetEquipmentSwitchCardOk() (*EquipmentSwitchCardRelationship, bool)`

GetEquipmentSwitchCardOk returns a tuple with the EquipmentSwitchCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentSwitchCard

`func (o *EtherPortChannel) SetEquipmentSwitchCard(v EquipmentSwitchCardRelationship)`

SetEquipmentSwitchCard sets EquipmentSwitchCard field to given value.

### HasEquipmentSwitchCard

`func (o *EtherPortChannel) HasEquipmentSwitchCard() bool`

HasEquipmentSwitchCard returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *EtherPortChannel) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EtherPortChannel) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EtherPortChannel) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EtherPortChannel) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


