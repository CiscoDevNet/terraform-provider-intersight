# VirtualizationVmwarePhysicalNetworkInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VmwarePhysicalNetworkInterface"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VmwarePhysicalNetworkInterface"]
**Driver** | Pointer to **string** | Driver type associated with physical network interface. | [optional] 
**LinkSpeed** | Pointer to **int32** | Link speed of the physical network interface. | [optional] 
**MacAddress** | Pointer to **string** | Standard MAC address assigned to physical network interface. | [optional] 
**Pci** | Pointer to **string** | PCI info for physical network interface. | [optional] 
**SwitchName** | Pointer to **string** | Switch associated with the physical network interface. | [optional] 
**Host** | Pointer to [**VirtualizationVmwareHostRelationship**](virtualization.VmwareHost.Relationship.md) |  | [optional] 

## Methods

### NewVirtualizationVmwarePhysicalNetworkInterface

`func NewVirtualizationVmwarePhysicalNetworkInterface(classId string, objectType string, ) *VirtualizationVmwarePhysicalNetworkInterface`

NewVirtualizationVmwarePhysicalNetworkInterface instantiates a new VirtualizationVmwarePhysicalNetworkInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwarePhysicalNetworkInterfaceWithDefaults

`func NewVirtualizationVmwarePhysicalNetworkInterfaceWithDefaults() *VirtualizationVmwarePhysicalNetworkInterface`

NewVirtualizationVmwarePhysicalNetworkInterfaceWithDefaults instantiates a new VirtualizationVmwarePhysicalNetworkInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmwarePhysicalNetworkInterface) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwarePhysicalNetworkInterface) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwarePhysicalNetworkInterface) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmwarePhysicalNetworkInterface) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwarePhysicalNetworkInterface) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwarePhysicalNetworkInterface) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDriver

`func (o *VirtualizationVmwarePhysicalNetworkInterface) GetDriver() string`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *VirtualizationVmwarePhysicalNetworkInterface) GetDriverOk() (*string, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *VirtualizationVmwarePhysicalNetworkInterface) SetDriver(v string)`

SetDriver sets Driver field to given value.

### HasDriver

`func (o *VirtualizationVmwarePhysicalNetworkInterface) HasDriver() bool`

HasDriver returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *VirtualizationVmwarePhysicalNetworkInterface) GetLinkSpeed() int32`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *VirtualizationVmwarePhysicalNetworkInterface) GetLinkSpeedOk() (*int32, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *VirtualizationVmwarePhysicalNetworkInterface) SetLinkSpeed(v int32)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *VirtualizationVmwarePhysicalNetworkInterface) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetMacAddress

`func (o *VirtualizationVmwarePhysicalNetworkInterface) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *VirtualizationVmwarePhysicalNetworkInterface) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *VirtualizationVmwarePhysicalNetworkInterface) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *VirtualizationVmwarePhysicalNetworkInterface) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetPci

`func (o *VirtualizationVmwarePhysicalNetworkInterface) GetPci() string`

GetPci returns the Pci field if non-nil, zero value otherwise.

### GetPciOk

`func (o *VirtualizationVmwarePhysicalNetworkInterface) GetPciOk() (*string, bool)`

GetPciOk returns a tuple with the Pci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPci

`func (o *VirtualizationVmwarePhysicalNetworkInterface) SetPci(v string)`

SetPci sets Pci field to given value.

### HasPci

`func (o *VirtualizationVmwarePhysicalNetworkInterface) HasPci() bool`

HasPci returns a boolean if a field has been set.

### GetSwitchName

`func (o *VirtualizationVmwarePhysicalNetworkInterface) GetSwitchName() string`

GetSwitchName returns the SwitchName field if non-nil, zero value otherwise.

### GetSwitchNameOk

`func (o *VirtualizationVmwarePhysicalNetworkInterface) GetSwitchNameOk() (*string, bool)`

GetSwitchNameOk returns a tuple with the SwitchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchName

`func (o *VirtualizationVmwarePhysicalNetworkInterface) SetSwitchName(v string)`

SetSwitchName sets SwitchName field to given value.

### HasSwitchName

`func (o *VirtualizationVmwarePhysicalNetworkInterface) HasSwitchName() bool`

HasSwitchName returns a boolean if a field has been set.

### GetHost

`func (o *VirtualizationVmwarePhysicalNetworkInterface) GetHost() VirtualizationVmwareHostRelationship`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *VirtualizationVmwarePhysicalNetworkInterface) GetHostOk() (*VirtualizationVmwareHostRelationship, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *VirtualizationVmwarePhysicalNetworkInterface) SetHost(v VirtualizationVmwareHostRelationship)`

SetHost sets Host field to given value.

### HasHost

`func (o *VirtualizationVmwarePhysicalNetworkInterface) HasHost() bool`

HasHost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


