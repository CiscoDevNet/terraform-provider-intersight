# NetworkDiscoveredNeighborAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "network.DiscoveredNeighbor"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "network.DiscoveredNeighbor"]
**DeviceId** | Pointer to **string** | Device identifier of the switch. | [optional] [readonly] 
**ManagementAddress** | Pointer to **string** | The management address of the switch. | [optional] [readonly] 
**NeighborDeviceCapability** | Pointer to **[]string** |  | [optional] 
**NeighborDeviceId** | Pointer to **string** | Device identifier of the neighboring device. | [optional] [readonly] 
**NeighborInterface** | Pointer to **string** | Neighboring device interface. | [optional] [readonly] 
**NeighborPortId** | Pointer to **string** | PortID of the neighbor device configured. | [optional] [readonly] 
**SwitchPortId** | Pointer to **string** | Name of the local interface. | [optional] [readonly] 
**CdpNeighbor** | Pointer to [**NetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**LldpNeighbor** | Pointer to [**NetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 

## Methods

### NewNetworkDiscoveredNeighborAllOf

`func NewNetworkDiscoveredNeighborAllOf(classId string, objectType string, ) *NetworkDiscoveredNeighborAllOf`

NewNetworkDiscoveredNeighborAllOf instantiates a new NetworkDiscoveredNeighborAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDiscoveredNeighborAllOfWithDefaults

`func NewNetworkDiscoveredNeighborAllOfWithDefaults() *NetworkDiscoveredNeighborAllOf`

NewNetworkDiscoveredNeighborAllOfWithDefaults instantiates a new NetworkDiscoveredNeighborAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NetworkDiscoveredNeighborAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NetworkDiscoveredNeighborAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NetworkDiscoveredNeighborAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NetworkDiscoveredNeighborAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NetworkDiscoveredNeighborAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NetworkDiscoveredNeighborAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDeviceId

`func (o *NetworkDiscoveredNeighborAllOf) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *NetworkDiscoveredNeighborAllOf) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *NetworkDiscoveredNeighborAllOf) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *NetworkDiscoveredNeighborAllOf) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetManagementAddress

`func (o *NetworkDiscoveredNeighborAllOf) GetManagementAddress() string`

GetManagementAddress returns the ManagementAddress field if non-nil, zero value otherwise.

### GetManagementAddressOk

`func (o *NetworkDiscoveredNeighborAllOf) GetManagementAddressOk() (*string, bool)`

GetManagementAddressOk returns a tuple with the ManagementAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementAddress

`func (o *NetworkDiscoveredNeighborAllOf) SetManagementAddress(v string)`

SetManagementAddress sets ManagementAddress field to given value.

### HasManagementAddress

`func (o *NetworkDiscoveredNeighborAllOf) HasManagementAddress() bool`

HasManagementAddress returns a boolean if a field has been set.

### GetNeighborDeviceCapability

`func (o *NetworkDiscoveredNeighborAllOf) GetNeighborDeviceCapability() []string`

GetNeighborDeviceCapability returns the NeighborDeviceCapability field if non-nil, zero value otherwise.

### GetNeighborDeviceCapabilityOk

`func (o *NetworkDiscoveredNeighborAllOf) GetNeighborDeviceCapabilityOk() (*[]string, bool)`

GetNeighborDeviceCapabilityOk returns a tuple with the NeighborDeviceCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighborDeviceCapability

`func (o *NetworkDiscoveredNeighborAllOf) SetNeighborDeviceCapability(v []string)`

SetNeighborDeviceCapability sets NeighborDeviceCapability field to given value.

### HasNeighborDeviceCapability

`func (o *NetworkDiscoveredNeighborAllOf) HasNeighborDeviceCapability() bool`

HasNeighborDeviceCapability returns a boolean if a field has been set.

### SetNeighborDeviceCapabilityNil

`func (o *NetworkDiscoveredNeighborAllOf) SetNeighborDeviceCapabilityNil(b bool)`

 SetNeighborDeviceCapabilityNil sets the value for NeighborDeviceCapability to be an explicit nil

### UnsetNeighborDeviceCapability
`func (o *NetworkDiscoveredNeighborAllOf) UnsetNeighborDeviceCapability()`

UnsetNeighborDeviceCapability ensures that no value is present for NeighborDeviceCapability, not even an explicit nil
### GetNeighborDeviceId

`func (o *NetworkDiscoveredNeighborAllOf) GetNeighborDeviceId() string`

GetNeighborDeviceId returns the NeighborDeviceId field if non-nil, zero value otherwise.

### GetNeighborDeviceIdOk

`func (o *NetworkDiscoveredNeighborAllOf) GetNeighborDeviceIdOk() (*string, bool)`

GetNeighborDeviceIdOk returns a tuple with the NeighborDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighborDeviceId

`func (o *NetworkDiscoveredNeighborAllOf) SetNeighborDeviceId(v string)`

SetNeighborDeviceId sets NeighborDeviceId field to given value.

### HasNeighborDeviceId

`func (o *NetworkDiscoveredNeighborAllOf) HasNeighborDeviceId() bool`

HasNeighborDeviceId returns a boolean if a field has been set.

### GetNeighborInterface

`func (o *NetworkDiscoveredNeighborAllOf) GetNeighborInterface() string`

GetNeighborInterface returns the NeighborInterface field if non-nil, zero value otherwise.

### GetNeighborInterfaceOk

`func (o *NetworkDiscoveredNeighborAllOf) GetNeighborInterfaceOk() (*string, bool)`

GetNeighborInterfaceOk returns a tuple with the NeighborInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighborInterface

`func (o *NetworkDiscoveredNeighborAllOf) SetNeighborInterface(v string)`

SetNeighborInterface sets NeighborInterface field to given value.

### HasNeighborInterface

`func (o *NetworkDiscoveredNeighborAllOf) HasNeighborInterface() bool`

HasNeighborInterface returns a boolean if a field has been set.

### GetNeighborPortId

`func (o *NetworkDiscoveredNeighborAllOf) GetNeighborPortId() string`

GetNeighborPortId returns the NeighborPortId field if non-nil, zero value otherwise.

### GetNeighborPortIdOk

`func (o *NetworkDiscoveredNeighborAllOf) GetNeighborPortIdOk() (*string, bool)`

GetNeighborPortIdOk returns a tuple with the NeighborPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighborPortId

`func (o *NetworkDiscoveredNeighborAllOf) SetNeighborPortId(v string)`

SetNeighborPortId sets NeighborPortId field to given value.

### HasNeighborPortId

`func (o *NetworkDiscoveredNeighborAllOf) HasNeighborPortId() bool`

HasNeighborPortId returns a boolean if a field has been set.

### GetSwitchPortId

`func (o *NetworkDiscoveredNeighborAllOf) GetSwitchPortId() string`

GetSwitchPortId returns the SwitchPortId field if non-nil, zero value otherwise.

### GetSwitchPortIdOk

`func (o *NetworkDiscoveredNeighborAllOf) GetSwitchPortIdOk() (*string, bool)`

GetSwitchPortIdOk returns a tuple with the SwitchPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchPortId

`func (o *NetworkDiscoveredNeighborAllOf) SetSwitchPortId(v string)`

SetSwitchPortId sets SwitchPortId field to given value.

### HasSwitchPortId

`func (o *NetworkDiscoveredNeighborAllOf) HasSwitchPortId() bool`

HasSwitchPortId returns a boolean if a field has been set.

### GetCdpNeighbor

`func (o *NetworkDiscoveredNeighborAllOf) GetCdpNeighbor() NetworkElementRelationship`

GetCdpNeighbor returns the CdpNeighbor field if non-nil, zero value otherwise.

### GetCdpNeighborOk

`func (o *NetworkDiscoveredNeighborAllOf) GetCdpNeighborOk() (*NetworkElementRelationship, bool)`

GetCdpNeighborOk returns a tuple with the CdpNeighbor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdpNeighbor

`func (o *NetworkDiscoveredNeighborAllOf) SetCdpNeighbor(v NetworkElementRelationship)`

SetCdpNeighbor sets CdpNeighbor field to given value.

### HasCdpNeighbor

`func (o *NetworkDiscoveredNeighborAllOf) HasCdpNeighbor() bool`

HasCdpNeighbor returns a boolean if a field has been set.

### GetLldpNeighbor

`func (o *NetworkDiscoveredNeighborAllOf) GetLldpNeighbor() NetworkElementRelationship`

GetLldpNeighbor returns the LldpNeighbor field if non-nil, zero value otherwise.

### GetLldpNeighborOk

`func (o *NetworkDiscoveredNeighborAllOf) GetLldpNeighborOk() (*NetworkElementRelationship, bool)`

GetLldpNeighborOk returns a tuple with the LldpNeighbor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldpNeighbor

`func (o *NetworkDiscoveredNeighborAllOf) SetLldpNeighbor(v NetworkElementRelationship)`

SetLldpNeighbor sets LldpNeighbor field to given value.

### HasLldpNeighbor

`func (o *NetworkDiscoveredNeighborAllOf) HasLldpNeighbor() bool`

HasLldpNeighbor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


