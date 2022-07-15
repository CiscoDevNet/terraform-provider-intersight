# NetworkInterfaceListAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "network.InterfaceList"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "network.InterfaceList"]
**AdminState** | Pointer to **string** | Admin state of the interface list. | [optional] [readonly] 
**AllowedVlans** | Pointer to **string** | Allowed VLANs of the interface list. | [optional] [readonly] 
**Description** | Pointer to **string** | Description of the interface list. | [optional] [readonly] 
**IpAddress** | Pointer to **string** | IP address of the interface list. | [optional] [readonly] 
**IpSubnet** | Pointer to **int64** | IP subnet of the interface list. | [optional] [readonly] 
**Mac** | Pointer to **string** | MAC address of the interface list. | [optional] [readonly] 
**Mtu** | Pointer to **int64** | Maximum transmission unit of the interface list. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the interface list. | [optional] [readonly] 
**OperState** | Pointer to **string** | Operational state of the interface list. | [optional] [readonly] 
**PortSubType** | Pointer to **string** | Interface types supported in Network device like Subinterfaces, Breakout Interfaces. | [optional] [readonly] 
**PortType** | Pointer to **string** | Port type of interface list. | [optional] [readonly] 
**SlotId** | Pointer to **string** | Slot id of the interface list. | [optional] [readonly] 
**Speed** | Pointer to **string** | Port speed of the interface list. | [optional] [readonly] 
**SpeedGroup** | Pointer to **string** | Speed Group of the interface list. | [optional] [readonly] 
**Vlan** | Pointer to **string** | VLAN of the interface list. | [optional] [readonly] 
**NetworkElement** | Pointer to [**NetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNetworkInterfaceListAllOf

`func NewNetworkInterfaceListAllOf(classId string, objectType string, ) *NetworkInterfaceListAllOf`

NewNetworkInterfaceListAllOf instantiates a new NetworkInterfaceListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkInterfaceListAllOfWithDefaults

`func NewNetworkInterfaceListAllOfWithDefaults() *NetworkInterfaceListAllOf`

NewNetworkInterfaceListAllOfWithDefaults instantiates a new NetworkInterfaceListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NetworkInterfaceListAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NetworkInterfaceListAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NetworkInterfaceListAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NetworkInterfaceListAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NetworkInterfaceListAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NetworkInterfaceListAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminState

`func (o *NetworkInterfaceListAllOf) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *NetworkInterfaceListAllOf) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *NetworkInterfaceListAllOf) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *NetworkInterfaceListAllOf) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetAllowedVlans

`func (o *NetworkInterfaceListAllOf) GetAllowedVlans() string`

GetAllowedVlans returns the AllowedVlans field if non-nil, zero value otherwise.

### GetAllowedVlansOk

`func (o *NetworkInterfaceListAllOf) GetAllowedVlansOk() (*string, bool)`

GetAllowedVlansOk returns a tuple with the AllowedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedVlans

`func (o *NetworkInterfaceListAllOf) SetAllowedVlans(v string)`

SetAllowedVlans sets AllowedVlans field to given value.

### HasAllowedVlans

`func (o *NetworkInterfaceListAllOf) HasAllowedVlans() bool`

HasAllowedVlans returns a boolean if a field has been set.

### GetDescription

`func (o *NetworkInterfaceListAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkInterfaceListAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkInterfaceListAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkInterfaceListAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpAddress

`func (o *NetworkInterfaceListAllOf) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *NetworkInterfaceListAllOf) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *NetworkInterfaceListAllOf) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *NetworkInterfaceListAllOf) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetIpSubnet

`func (o *NetworkInterfaceListAllOf) GetIpSubnet() int64`

GetIpSubnet returns the IpSubnet field if non-nil, zero value otherwise.

### GetIpSubnetOk

`func (o *NetworkInterfaceListAllOf) GetIpSubnetOk() (*int64, bool)`

GetIpSubnetOk returns a tuple with the IpSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSubnet

`func (o *NetworkInterfaceListAllOf) SetIpSubnet(v int64)`

SetIpSubnet sets IpSubnet field to given value.

### HasIpSubnet

`func (o *NetworkInterfaceListAllOf) HasIpSubnet() bool`

HasIpSubnet returns a boolean if a field has been set.

### GetMac

`func (o *NetworkInterfaceListAllOf) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *NetworkInterfaceListAllOf) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *NetworkInterfaceListAllOf) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *NetworkInterfaceListAllOf) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMtu

`func (o *NetworkInterfaceListAllOf) GetMtu() int64`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *NetworkInterfaceListAllOf) GetMtuOk() (*int64, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *NetworkInterfaceListAllOf) SetMtu(v int64)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *NetworkInterfaceListAllOf) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetName

`func (o *NetworkInterfaceListAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkInterfaceListAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkInterfaceListAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkInterfaceListAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperState

`func (o *NetworkInterfaceListAllOf) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *NetworkInterfaceListAllOf) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *NetworkInterfaceListAllOf) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *NetworkInterfaceListAllOf) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetPortSubType

`func (o *NetworkInterfaceListAllOf) GetPortSubType() string`

GetPortSubType returns the PortSubType field if non-nil, zero value otherwise.

### GetPortSubTypeOk

`func (o *NetworkInterfaceListAllOf) GetPortSubTypeOk() (*string, bool)`

GetPortSubTypeOk returns a tuple with the PortSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortSubType

`func (o *NetworkInterfaceListAllOf) SetPortSubType(v string)`

SetPortSubType sets PortSubType field to given value.

### HasPortSubType

`func (o *NetworkInterfaceListAllOf) HasPortSubType() bool`

HasPortSubType returns a boolean if a field has been set.

### GetPortType

`func (o *NetworkInterfaceListAllOf) GetPortType() string`

GetPortType returns the PortType field if non-nil, zero value otherwise.

### GetPortTypeOk

`func (o *NetworkInterfaceListAllOf) GetPortTypeOk() (*string, bool)`

GetPortTypeOk returns a tuple with the PortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortType

`func (o *NetworkInterfaceListAllOf) SetPortType(v string)`

SetPortType sets PortType field to given value.

### HasPortType

`func (o *NetworkInterfaceListAllOf) HasPortType() bool`

HasPortType returns a boolean if a field has been set.

### GetSlotId

`func (o *NetworkInterfaceListAllOf) GetSlotId() string`

GetSlotId returns the SlotId field if non-nil, zero value otherwise.

### GetSlotIdOk

`func (o *NetworkInterfaceListAllOf) GetSlotIdOk() (*string, bool)`

GetSlotIdOk returns a tuple with the SlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotId

`func (o *NetworkInterfaceListAllOf) SetSlotId(v string)`

SetSlotId sets SlotId field to given value.

### HasSlotId

`func (o *NetworkInterfaceListAllOf) HasSlotId() bool`

HasSlotId returns a boolean if a field has been set.

### GetSpeed

`func (o *NetworkInterfaceListAllOf) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *NetworkInterfaceListAllOf) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *NetworkInterfaceListAllOf) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *NetworkInterfaceListAllOf) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetSpeedGroup

`func (o *NetworkInterfaceListAllOf) GetSpeedGroup() string`

GetSpeedGroup returns the SpeedGroup field if non-nil, zero value otherwise.

### GetSpeedGroupOk

`func (o *NetworkInterfaceListAllOf) GetSpeedGroupOk() (*string, bool)`

GetSpeedGroupOk returns a tuple with the SpeedGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedGroup

`func (o *NetworkInterfaceListAllOf) SetSpeedGroup(v string)`

SetSpeedGroup sets SpeedGroup field to given value.

### HasSpeedGroup

`func (o *NetworkInterfaceListAllOf) HasSpeedGroup() bool`

HasSpeedGroup returns a boolean if a field has been set.

### GetVlan

`func (o *NetworkInterfaceListAllOf) GetVlan() string`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *NetworkInterfaceListAllOf) GetVlanOk() (*string, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *NetworkInterfaceListAllOf) SetVlan(v string)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *NetworkInterfaceListAllOf) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetNetworkElement

`func (o *NetworkInterfaceListAllOf) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *NetworkInterfaceListAllOf) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *NetworkInterfaceListAllOf) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *NetworkInterfaceListAllOf) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NetworkInterfaceListAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NetworkInterfaceListAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NetworkInterfaceListAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NetworkInterfaceListAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


