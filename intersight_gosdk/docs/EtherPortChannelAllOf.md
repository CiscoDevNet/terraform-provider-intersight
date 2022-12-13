# EtherPortChannelAllOf

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
**EquipmentSwitchCard** | Pointer to [**EquipmentSwitchCardRelationship**](EquipmentSwitchCardRelationship.md) |  | [optional] 
**NetworkElement** | Pointer to [**NetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewEtherPortChannelAllOf

`func NewEtherPortChannelAllOf(classId string, objectType string, ) *EtherPortChannelAllOf`

NewEtherPortChannelAllOf instantiates a new EtherPortChannelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEtherPortChannelAllOfWithDefaults

`func NewEtherPortChannelAllOfWithDefaults() *EtherPortChannelAllOf`

NewEtherPortChannelAllOfWithDefaults instantiates a new EtherPortChannelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EtherPortChannelAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EtherPortChannelAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EtherPortChannelAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EtherPortChannelAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EtherPortChannelAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EtherPortChannelAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccessVlan

`func (o *EtherPortChannelAllOf) GetAccessVlan() string`

GetAccessVlan returns the AccessVlan field if non-nil, zero value otherwise.

### GetAccessVlanOk

`func (o *EtherPortChannelAllOf) GetAccessVlanOk() (*string, bool)`

GetAccessVlanOk returns a tuple with the AccessVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessVlan

`func (o *EtherPortChannelAllOf) SetAccessVlan(v string)`

SetAccessVlan sets AccessVlan field to given value.

### HasAccessVlan

`func (o *EtherPortChannelAllOf) HasAccessVlan() bool`

HasAccessVlan returns a boolean if a field has been set.

### GetAdminState

`func (o *EtherPortChannelAllOf) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *EtherPortChannelAllOf) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *EtherPortChannelAllOf) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *EtherPortChannelAllOf) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetAllowedVlans

`func (o *EtherPortChannelAllOf) GetAllowedVlans() string`

GetAllowedVlans returns the AllowedVlans field if non-nil, zero value otherwise.

### GetAllowedVlansOk

`func (o *EtherPortChannelAllOf) GetAllowedVlansOk() (*string, bool)`

GetAllowedVlansOk returns a tuple with the AllowedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedVlans

`func (o *EtherPortChannelAllOf) SetAllowedVlans(v string)`

SetAllowedVlans sets AllowedVlans field to given value.

### HasAllowedVlans

`func (o *EtherPortChannelAllOf) HasAllowedVlans() bool`

HasAllowedVlans returns a boolean if a field has been set.

### GetBandWidth

`func (o *EtherPortChannelAllOf) GetBandWidth() string`

GetBandWidth returns the BandWidth field if non-nil, zero value otherwise.

### GetBandWidthOk

`func (o *EtherPortChannelAllOf) GetBandWidthOk() (*string, bool)`

GetBandWidthOk returns a tuple with the BandWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandWidth

`func (o *EtherPortChannelAllOf) SetBandWidth(v string)`

SetBandWidth sets BandWidth field to given value.

### HasBandWidth

`func (o *EtherPortChannelAllOf) HasBandWidth() bool`

HasBandWidth returns a boolean if a field has been set.

### GetDescription

`func (o *EtherPortChannelAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EtherPortChannelAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EtherPortChannelAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EtherPortChannelAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpAddress

`func (o *EtherPortChannelAllOf) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *EtherPortChannelAllOf) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *EtherPortChannelAllOf) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *EtherPortChannelAllOf) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetIpAddressMask

`func (o *EtherPortChannelAllOf) GetIpAddressMask() int64`

GetIpAddressMask returns the IpAddressMask field if non-nil, zero value otherwise.

### GetIpAddressMaskOk

`func (o *EtherPortChannelAllOf) GetIpAddressMaskOk() (*int64, bool)`

GetIpAddressMaskOk returns a tuple with the IpAddressMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddressMask

`func (o *EtherPortChannelAllOf) SetIpAddressMask(v int64)`

SetIpAddressMask sets IpAddressMask field to given value.

### HasIpAddressMask

`func (o *EtherPortChannelAllOf) HasIpAddressMask() bool`

HasIpAddressMask returns a boolean if a field has been set.

### GetIpv6SubnetCidr

`func (o *EtherPortChannelAllOf) GetIpv6SubnetCidr() string`

GetIpv6SubnetCidr returns the Ipv6SubnetCidr field if non-nil, zero value otherwise.

### GetIpv6SubnetCidrOk

`func (o *EtherPortChannelAllOf) GetIpv6SubnetCidrOk() (*string, bool)`

GetIpv6SubnetCidrOk returns a tuple with the Ipv6SubnetCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6SubnetCidr

`func (o *EtherPortChannelAllOf) SetIpv6SubnetCidr(v string)`

SetIpv6SubnetCidr sets Ipv6SubnetCidr field to given value.

### HasIpv6SubnetCidr

`func (o *EtherPortChannelAllOf) HasIpv6SubnetCidr() bool`

HasIpv6SubnetCidr returns a boolean if a field has been set.

### GetMacAddress

`func (o *EtherPortChannelAllOf) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *EtherPortChannelAllOf) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *EtherPortChannelAllOf) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *EtherPortChannelAllOf) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMode

`func (o *EtherPortChannelAllOf) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *EtherPortChannelAllOf) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *EtherPortChannelAllOf) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *EtherPortChannelAllOf) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetMtu

`func (o *EtherPortChannelAllOf) GetMtu() int64`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *EtherPortChannelAllOf) GetMtuOk() (*int64, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *EtherPortChannelAllOf) SetMtu(v int64)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *EtherPortChannelAllOf) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetName

`func (o *EtherPortChannelAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EtherPortChannelAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EtherPortChannelAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EtherPortChannelAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNativeVlan

`func (o *EtherPortChannelAllOf) GetNativeVlan() string`

GetNativeVlan returns the NativeVlan field if non-nil, zero value otherwise.

### GetNativeVlanOk

`func (o *EtherPortChannelAllOf) GetNativeVlanOk() (*string, bool)`

GetNativeVlanOk returns a tuple with the NativeVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeVlan

`func (o *EtherPortChannelAllOf) SetNativeVlan(v string)`

SetNativeVlan sets NativeVlan field to given value.

### HasNativeVlan

`func (o *EtherPortChannelAllOf) HasNativeVlan() bool`

HasNativeVlan returns a boolean if a field has been set.

### GetOperSpeed

`func (o *EtherPortChannelAllOf) GetOperSpeed() string`

GetOperSpeed returns the OperSpeed field if non-nil, zero value otherwise.

### GetOperSpeedOk

`func (o *EtherPortChannelAllOf) GetOperSpeedOk() (*string, bool)`

GetOperSpeedOk returns a tuple with the OperSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperSpeed

`func (o *EtherPortChannelAllOf) SetOperSpeed(v string)`

SetOperSpeed sets OperSpeed field to given value.

### HasOperSpeed

`func (o *EtherPortChannelAllOf) HasOperSpeed() bool`

HasOperSpeed returns a boolean if a field has been set.

### GetOperState

`func (o *EtherPortChannelAllOf) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *EtherPortChannelAllOf) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *EtherPortChannelAllOf) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *EtherPortChannelAllOf) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetOperStateQual

`func (o *EtherPortChannelAllOf) GetOperStateQual() string`

GetOperStateQual returns the OperStateQual field if non-nil, zero value otherwise.

### GetOperStateQualOk

`func (o *EtherPortChannelAllOf) GetOperStateQualOk() (*string, bool)`

GetOperStateQualOk returns a tuple with the OperStateQual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperStateQual

`func (o *EtherPortChannelAllOf) SetOperStateQual(v string)`

SetOperStateQual sets OperStateQual field to given value.

### HasOperStateQual

`func (o *EtherPortChannelAllOf) HasOperStateQual() bool`

HasOperStateQual returns a boolean if a field has been set.

### GetPortChannelId

`func (o *EtherPortChannelAllOf) GetPortChannelId() int64`

GetPortChannelId returns the PortChannelId field if non-nil, zero value otherwise.

### GetPortChannelIdOk

`func (o *EtherPortChannelAllOf) GetPortChannelIdOk() (*int64, bool)`

GetPortChannelIdOk returns a tuple with the PortChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortChannelId

`func (o *EtherPortChannelAllOf) SetPortChannelId(v int64)`

SetPortChannelId sets PortChannelId field to given value.

### HasPortChannelId

`func (o *EtherPortChannelAllOf) HasPortChannelId() bool`

HasPortChannelId returns a boolean if a field has been set.

### GetRole

`func (o *EtherPortChannelAllOf) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *EtherPortChannelAllOf) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *EtherPortChannelAllOf) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *EtherPortChannelAllOf) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetStatus

`func (o *EtherPortChannelAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EtherPortChannelAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EtherPortChannelAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EtherPortChannelAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSwitchId

`func (o *EtherPortChannelAllOf) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *EtherPortChannelAllOf) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *EtherPortChannelAllOf) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *EtherPortChannelAllOf) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetEquipmentSwitchCard

`func (o *EtherPortChannelAllOf) GetEquipmentSwitchCard() EquipmentSwitchCardRelationship`

GetEquipmentSwitchCard returns the EquipmentSwitchCard field if non-nil, zero value otherwise.

### GetEquipmentSwitchCardOk

`func (o *EtherPortChannelAllOf) GetEquipmentSwitchCardOk() (*EquipmentSwitchCardRelationship, bool)`

GetEquipmentSwitchCardOk returns a tuple with the EquipmentSwitchCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentSwitchCard

`func (o *EtherPortChannelAllOf) SetEquipmentSwitchCard(v EquipmentSwitchCardRelationship)`

SetEquipmentSwitchCard sets EquipmentSwitchCard field to given value.

### HasEquipmentSwitchCard

`func (o *EtherPortChannelAllOf) HasEquipmentSwitchCard() bool`

HasEquipmentSwitchCard returns a boolean if a field has been set.

### GetNetworkElement

`func (o *EtherPortChannelAllOf) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *EtherPortChannelAllOf) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *EtherPortChannelAllOf) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *EtherPortChannelAllOf) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *EtherPortChannelAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *EtherPortChannelAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *EtherPortChannelAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *EtherPortChannelAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


