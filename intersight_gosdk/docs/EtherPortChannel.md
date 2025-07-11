# EtherPortChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "ether.PortChannel"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "ether.PortChannel"]
**AccessVlan** | Pointer to **string** | Access VLANs for this port-channel, on this FI. | [optional] [readonly] 
**AdminState** | Pointer to **string** | Administratively configured state (enabled/disabled) for this port-channel. | [optional] [readonly] 
**AllowedVlans** | Pointer to **string** | Allowed VLANs on this port-channel, on this FI. | [optional] [readonly] 
**BandWidth** | Pointer to **string** | Bandwidth of this port-channel. | [optional] [readonly] 
**Description** | Pointer to **string** | Description of this port-channel. | [optional] [readonly] 
**IpAddress** | Pointer to **string** | IP address of this port-channel. | [optional] [readonly] 
**IpAddressMask** | Pointer to **int64** | IP address mask of this port-channel. | [optional] [readonly] 
**Ipv6SubnetCidr** | Pointer to **string** | IPv6 subnet in CIDR notation of this port-channel. Ex. 2000::/8. | [optional] [readonly] 
**MacAddress** | Pointer to **string** | MAC address of this port-channel. | [optional] [readonly] 
**Mode** | Pointer to **string** | Operating mode of this port-channel. | [optional] [readonly] 
**Mtu** | Pointer to **int64** | Maximum transmission unit of this port-channel. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the port channel. | [optional] [readonly] 
**NativeVlan** | Pointer to **string** | Native VLAN for this port-channel, on this FI. | [optional] [readonly] 
**OperSpeed** | Pointer to **string** | Operational speed of this port-channel. | [optional] [readonly] 
**OperState** | Pointer to **string** | Operational state of this port-channel. | [optional] [readonly] 
**OperStateQual** | Pointer to **string** | Reason for this port-channel&#39;s Operational state. | [optional] [readonly] 
**PortChannelId** | Pointer to **int64** | Unique identifier for this port-channel on the FI. | [optional] [readonly] 
**Role** | Pointer to **string** | This port-channel&#39;s configured role (uplink, server, etc.). | [optional] [readonly] 
**Status** | Pointer to **string** | Detailed status of this port-channel. | [optional] [readonly] 
**SwitchId** | Pointer to **string** | Switch Identifier that is local to a cluster. | [optional] [readonly] 
**UserLabel** | Pointer to **string** | The user defined label assigned to the port channel. | [optional] [readonly] 
**EquipmentSwitchCard** | Pointer to [**NullableEquipmentSwitchCardRelationship**](EquipmentSwitchCardRelationship.md) |  | [optional] 
**NetworkElement** | Pointer to [**NullableNetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewEtherPortChannel

`func NewEtherPortChannel(classId string, objectType string, ) *EtherPortChannel`

NewEtherPortChannel instantiates a new EtherPortChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEtherPortChannelWithDefaults

`func NewEtherPortChannelWithDefaults() *EtherPortChannel`

NewEtherPortChannelWithDefaults instantiates a new EtherPortChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EtherPortChannel) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EtherPortChannel) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EtherPortChannel) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EtherPortChannel) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EtherPortChannel) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EtherPortChannel) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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

### GetBandWidth

`func (o *EtherPortChannel) GetBandWidth() string`

GetBandWidth returns the BandWidth field if non-nil, zero value otherwise.

### GetBandWidthOk

`func (o *EtherPortChannel) GetBandWidthOk() (*string, bool)`

GetBandWidthOk returns a tuple with the BandWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandWidth

`func (o *EtherPortChannel) SetBandWidth(v string)`

SetBandWidth sets BandWidth field to given value.

### HasBandWidth

`func (o *EtherPortChannel) HasBandWidth() bool`

HasBandWidth returns a boolean if a field has been set.

### GetDescription

`func (o *EtherPortChannel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EtherPortChannel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EtherPortChannel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EtherPortChannel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpAddress

`func (o *EtherPortChannel) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *EtherPortChannel) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *EtherPortChannel) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *EtherPortChannel) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetIpAddressMask

`func (o *EtherPortChannel) GetIpAddressMask() int64`

GetIpAddressMask returns the IpAddressMask field if non-nil, zero value otherwise.

### GetIpAddressMaskOk

`func (o *EtherPortChannel) GetIpAddressMaskOk() (*int64, bool)`

GetIpAddressMaskOk returns a tuple with the IpAddressMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddressMask

`func (o *EtherPortChannel) SetIpAddressMask(v int64)`

SetIpAddressMask sets IpAddressMask field to given value.

### HasIpAddressMask

`func (o *EtherPortChannel) HasIpAddressMask() bool`

HasIpAddressMask returns a boolean if a field has been set.

### GetIpv6SubnetCidr

`func (o *EtherPortChannel) GetIpv6SubnetCidr() string`

GetIpv6SubnetCidr returns the Ipv6SubnetCidr field if non-nil, zero value otherwise.

### GetIpv6SubnetCidrOk

`func (o *EtherPortChannel) GetIpv6SubnetCidrOk() (*string, bool)`

GetIpv6SubnetCidrOk returns a tuple with the Ipv6SubnetCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6SubnetCidr

`func (o *EtherPortChannel) SetIpv6SubnetCidr(v string)`

SetIpv6SubnetCidr sets Ipv6SubnetCidr field to given value.

### HasIpv6SubnetCidr

`func (o *EtherPortChannel) HasIpv6SubnetCidr() bool`

HasIpv6SubnetCidr returns a boolean if a field has been set.

### GetMacAddress

`func (o *EtherPortChannel) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *EtherPortChannel) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *EtherPortChannel) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *EtherPortChannel) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

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

### GetMtu

`func (o *EtherPortChannel) GetMtu() int64`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *EtherPortChannel) GetMtuOk() (*int64, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *EtherPortChannel) SetMtu(v int64)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *EtherPortChannel) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetName

`func (o *EtherPortChannel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EtherPortChannel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EtherPortChannel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EtherPortChannel) HasName() bool`

HasName returns a boolean if a field has been set.

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

### GetStatus

`func (o *EtherPortChannel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EtherPortChannel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EtherPortChannel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EtherPortChannel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

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

### GetUserLabel

`func (o *EtherPortChannel) GetUserLabel() string`

GetUserLabel returns the UserLabel field if non-nil, zero value otherwise.

### GetUserLabelOk

`func (o *EtherPortChannel) GetUserLabelOk() (*string, bool)`

GetUserLabelOk returns a tuple with the UserLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLabel

`func (o *EtherPortChannel) SetUserLabel(v string)`

SetUserLabel sets UserLabel field to given value.

### HasUserLabel

`func (o *EtherPortChannel) HasUserLabel() bool`

HasUserLabel returns a boolean if a field has been set.

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

### SetEquipmentSwitchCardNil

`func (o *EtherPortChannel) SetEquipmentSwitchCardNil(b bool)`

 SetEquipmentSwitchCardNil sets the value for EquipmentSwitchCard to be an explicit nil

### UnsetEquipmentSwitchCard
`func (o *EtherPortChannel) UnsetEquipmentSwitchCard()`

UnsetEquipmentSwitchCard ensures that no value is present for EquipmentSwitchCard, not even an explicit nil
### GetNetworkElement

`func (o *EtherPortChannel) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *EtherPortChannel) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *EtherPortChannel) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *EtherPortChannel) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### SetNetworkElementNil

`func (o *EtherPortChannel) SetNetworkElementNil(b bool)`

 SetNetworkElementNil sets the value for NetworkElement to be an explicit nil

### UnsetNetworkElement
`func (o *EtherPortChannel) UnsetNetworkElement()`

UnsetNetworkElement ensures that no value is present for NetworkElement, not even an explicit nil
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

### SetRegisteredDeviceNil

`func (o *EtherPortChannel) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *EtherPortChannel) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


