# NiatelemetrySwitchInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.SwitchInterface"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.SwitchInterface"]
**AdminStatus** | Pointer to **string** | Administrative status of the interface. | [optional] [readonly] 
**Discovered** | Pointer to **bool** | Indicates whether the interface was discovered. | [optional] [readonly] 
**DiscoveredConfigMode** | Pointer to **string** | Discovered configuration mode of the interface. | [optional] [readonly] 
**DisplayName** | Pointer to **string** | Display name of the interface. | [optional] [readonly] 
**ExternalFabric** | Pointer to **bool** | Indicates whether the interface belongs to an external fabric. | [optional] [readonly] 
**FabricName** | Pointer to **string** | Fabric name for the interface. | [optional] [readonly] 
**InterfaceName** | Pointer to **string** | The name of the interface. | [optional] [readonly] 
**InterfaceType** | Pointer to **string** | The type of the interface. | [optional] [readonly] 
**Ip** | Pointer to **string** | IP address configured on the interface. | [optional] [readonly] 
**OperMode** | Pointer to **string** | Operational mode of the interface. | [optional] [readonly] 
**OperationalStatus** | Pointer to **string** | Operational status of the interface. | [optional] [readonly] 
**Platform** | Pointer to **string** | Platform of the interface. | [optional] [readonly] 
**PortChannelId** | Pointer to **int64** | The portchannel ID of the associated interface. | [optional] [readonly] 
**Speed** | Pointer to **string** | The speed of the interface. | [optional] [readonly] 
**SwitchId** | Pointer to **string** | Identifier of the switch that owns the interface. | [optional] [readonly] 
**SwitchName** | Pointer to **string** | Name of the switch that owns the interface. | [optional] [readonly] 
**VlanId** | Pointer to **int64** | VLAN identifier associated with the interface. | [optional] [readonly] 
**VpcId** | Pointer to **int64** | VPC identifier for the interface. | [optional] [readonly] 
**Switch** | Pointer to [**NullableNiatelemetrySwitchRelationship**](NiatelemetrySwitchRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetrySwitchInterface

`func NewNiatelemetrySwitchInterface(classId string, objectType string, ) *NiatelemetrySwitchInterface`

NewNiatelemetrySwitchInterface instantiates a new NiatelemetrySwitchInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetrySwitchInterfaceWithDefaults

`func NewNiatelemetrySwitchInterfaceWithDefaults() *NiatelemetrySwitchInterface`

NewNiatelemetrySwitchInterfaceWithDefaults instantiates a new NiatelemetrySwitchInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetrySwitchInterface) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetrySwitchInterface) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetrySwitchInterface) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetrySwitchInterface) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetrySwitchInterface) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetrySwitchInterface) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminStatus

`func (o *NiatelemetrySwitchInterface) GetAdminStatus() string`

GetAdminStatus returns the AdminStatus field if non-nil, zero value otherwise.

### GetAdminStatusOk

`func (o *NiatelemetrySwitchInterface) GetAdminStatusOk() (*string, bool)`

GetAdminStatusOk returns a tuple with the AdminStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminStatus

`func (o *NiatelemetrySwitchInterface) SetAdminStatus(v string)`

SetAdminStatus sets AdminStatus field to given value.

### HasAdminStatus

`func (o *NiatelemetrySwitchInterface) HasAdminStatus() bool`

HasAdminStatus returns a boolean if a field has been set.

### GetDiscovered

`func (o *NiatelemetrySwitchInterface) GetDiscovered() bool`

GetDiscovered returns the Discovered field if non-nil, zero value otherwise.

### GetDiscoveredOk

`func (o *NiatelemetrySwitchInterface) GetDiscoveredOk() (*bool, bool)`

GetDiscoveredOk returns a tuple with the Discovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscovered

`func (o *NiatelemetrySwitchInterface) SetDiscovered(v bool)`

SetDiscovered sets Discovered field to given value.

### HasDiscovered

`func (o *NiatelemetrySwitchInterface) HasDiscovered() bool`

HasDiscovered returns a boolean if a field has been set.

### GetDiscoveredConfigMode

`func (o *NiatelemetrySwitchInterface) GetDiscoveredConfigMode() string`

GetDiscoveredConfigMode returns the DiscoveredConfigMode field if non-nil, zero value otherwise.

### GetDiscoveredConfigModeOk

`func (o *NiatelemetrySwitchInterface) GetDiscoveredConfigModeOk() (*string, bool)`

GetDiscoveredConfigModeOk returns a tuple with the DiscoveredConfigMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredConfigMode

`func (o *NiatelemetrySwitchInterface) SetDiscoveredConfigMode(v string)`

SetDiscoveredConfigMode sets DiscoveredConfigMode field to given value.

### HasDiscoveredConfigMode

`func (o *NiatelemetrySwitchInterface) HasDiscoveredConfigMode() bool`

HasDiscoveredConfigMode returns a boolean if a field has been set.

### GetDisplayName

`func (o *NiatelemetrySwitchInterface) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NiatelemetrySwitchInterface) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NiatelemetrySwitchInterface) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *NiatelemetrySwitchInterface) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExternalFabric

`func (o *NiatelemetrySwitchInterface) GetExternalFabric() bool`

GetExternalFabric returns the ExternalFabric field if non-nil, zero value otherwise.

### GetExternalFabricOk

`func (o *NiatelemetrySwitchInterface) GetExternalFabricOk() (*bool, bool)`

GetExternalFabricOk returns a tuple with the ExternalFabric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalFabric

`func (o *NiatelemetrySwitchInterface) SetExternalFabric(v bool)`

SetExternalFabric sets ExternalFabric field to given value.

### HasExternalFabric

`func (o *NiatelemetrySwitchInterface) HasExternalFabric() bool`

HasExternalFabric returns a boolean if a field has been set.

### GetFabricName

`func (o *NiatelemetrySwitchInterface) GetFabricName() string`

GetFabricName returns the FabricName field if non-nil, zero value otherwise.

### GetFabricNameOk

`func (o *NiatelemetrySwitchInterface) GetFabricNameOk() (*string, bool)`

GetFabricNameOk returns a tuple with the FabricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricName

`func (o *NiatelemetrySwitchInterface) SetFabricName(v string)`

SetFabricName sets FabricName field to given value.

### HasFabricName

`func (o *NiatelemetrySwitchInterface) HasFabricName() bool`

HasFabricName returns a boolean if a field has been set.

### GetInterfaceName

`func (o *NiatelemetrySwitchInterface) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *NiatelemetrySwitchInterface) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *NiatelemetrySwitchInterface) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *NiatelemetrySwitchInterface) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### GetInterfaceType

`func (o *NiatelemetrySwitchInterface) GetInterfaceType() string`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *NiatelemetrySwitchInterface) GetInterfaceTypeOk() (*string, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *NiatelemetrySwitchInterface) SetInterfaceType(v string)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *NiatelemetrySwitchInterface) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### GetIp

`func (o *NiatelemetrySwitchInterface) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *NiatelemetrySwitchInterface) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *NiatelemetrySwitchInterface) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *NiatelemetrySwitchInterface) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetOperMode

`func (o *NiatelemetrySwitchInterface) GetOperMode() string`

GetOperMode returns the OperMode field if non-nil, zero value otherwise.

### GetOperModeOk

`func (o *NiatelemetrySwitchInterface) GetOperModeOk() (*string, bool)`

GetOperModeOk returns a tuple with the OperMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperMode

`func (o *NiatelemetrySwitchInterface) SetOperMode(v string)`

SetOperMode sets OperMode field to given value.

### HasOperMode

`func (o *NiatelemetrySwitchInterface) HasOperMode() bool`

HasOperMode returns a boolean if a field has been set.

### GetOperationalStatus

`func (o *NiatelemetrySwitchInterface) GetOperationalStatus() string`

GetOperationalStatus returns the OperationalStatus field if non-nil, zero value otherwise.

### GetOperationalStatusOk

`func (o *NiatelemetrySwitchInterface) GetOperationalStatusOk() (*string, bool)`

GetOperationalStatusOk returns a tuple with the OperationalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalStatus

`func (o *NiatelemetrySwitchInterface) SetOperationalStatus(v string)`

SetOperationalStatus sets OperationalStatus field to given value.

### HasOperationalStatus

`func (o *NiatelemetrySwitchInterface) HasOperationalStatus() bool`

HasOperationalStatus returns a boolean if a field has been set.

### GetPlatform

`func (o *NiatelemetrySwitchInterface) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *NiatelemetrySwitchInterface) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *NiatelemetrySwitchInterface) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *NiatelemetrySwitchInterface) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetPortChannelId

`func (o *NiatelemetrySwitchInterface) GetPortChannelId() int64`

GetPortChannelId returns the PortChannelId field if non-nil, zero value otherwise.

### GetPortChannelIdOk

`func (o *NiatelemetrySwitchInterface) GetPortChannelIdOk() (*int64, bool)`

GetPortChannelIdOk returns a tuple with the PortChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortChannelId

`func (o *NiatelemetrySwitchInterface) SetPortChannelId(v int64)`

SetPortChannelId sets PortChannelId field to given value.

### HasPortChannelId

`func (o *NiatelemetrySwitchInterface) HasPortChannelId() bool`

HasPortChannelId returns a boolean if a field has been set.

### GetSpeed

`func (o *NiatelemetrySwitchInterface) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *NiatelemetrySwitchInterface) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *NiatelemetrySwitchInterface) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *NiatelemetrySwitchInterface) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetSwitchId

`func (o *NiatelemetrySwitchInterface) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *NiatelemetrySwitchInterface) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *NiatelemetrySwitchInterface) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *NiatelemetrySwitchInterface) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetSwitchName

`func (o *NiatelemetrySwitchInterface) GetSwitchName() string`

GetSwitchName returns the SwitchName field if non-nil, zero value otherwise.

### GetSwitchNameOk

`func (o *NiatelemetrySwitchInterface) GetSwitchNameOk() (*string, bool)`

GetSwitchNameOk returns a tuple with the SwitchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchName

`func (o *NiatelemetrySwitchInterface) SetSwitchName(v string)`

SetSwitchName sets SwitchName field to given value.

### HasSwitchName

`func (o *NiatelemetrySwitchInterface) HasSwitchName() bool`

HasSwitchName returns a boolean if a field has been set.

### GetVlanId

`func (o *NiatelemetrySwitchInterface) GetVlanId() int64`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *NiatelemetrySwitchInterface) GetVlanIdOk() (*int64, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *NiatelemetrySwitchInterface) SetVlanId(v int64)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *NiatelemetrySwitchInterface) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetVpcId

`func (o *NiatelemetrySwitchInterface) GetVpcId() int64`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *NiatelemetrySwitchInterface) GetVpcIdOk() (*int64, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *NiatelemetrySwitchInterface) SetVpcId(v int64)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *NiatelemetrySwitchInterface) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### GetSwitch

`func (o *NiatelemetrySwitchInterface) GetSwitch() NiatelemetrySwitchRelationship`

GetSwitch returns the Switch field if non-nil, zero value otherwise.

### GetSwitchOk

`func (o *NiatelemetrySwitchInterface) GetSwitchOk() (*NiatelemetrySwitchRelationship, bool)`

GetSwitchOk returns a tuple with the Switch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitch

`func (o *NiatelemetrySwitchInterface) SetSwitch(v NiatelemetrySwitchRelationship)`

SetSwitch sets Switch field to given value.

### HasSwitch

`func (o *NiatelemetrySwitchInterface) HasSwitch() bool`

HasSwitch returns a boolean if a field has been set.

### SetSwitchNil

`func (o *NiatelemetrySwitchInterface) SetSwitchNil(b bool)`

 SetSwitchNil sets the value for Switch to be an explicit nil

### UnsetSwitch
`func (o *NiatelemetrySwitchInterface) UnsetSwitch()`

UnsetSwitch ensures that no value is present for Switch, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


