# VirtualizationVmwarePhysicalNetworkInterfaceAllOf

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

### NewVirtualizationVmwarePhysicalNetworkInterfaceAllOf

`func NewVirtualizationVmwarePhysicalNetworkInterfaceAllOf(classId string, objectType string, ) *VirtualizationVmwarePhysicalNetworkInterfaceAllOf`

NewVirtualizationVmwarePhysicalNetworkInterfaceAllOf instantiates a new VirtualizationVmwarePhysicalNetworkInterfaceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwarePhysicalNetworkInterfaceAllOfWithDefaults

`func NewVirtualizationVmwarePhysicalNetworkInterfaceAllOfWithDefaults() *VirtualizationVmwarePhysicalNetworkInterfaceAllOf`

NewVirtualizationVmwarePhysicalNetworkInterfaceAllOfWithDefaults instantiates a new VirtualizationVmwarePhysicalNetworkInterfaceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDriver

`func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) GetDriver() string`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) GetDriverOk() (*string, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) SetDriver(v string)`

SetDriver sets Driver field to given value.

### HasDriver

`func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) HasDriver() bool`

HasDriver returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) GetLinkSpeed() int32`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) GetLinkSpeedOk() (*int32, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) SetLinkSpeed(v int32)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetMacAddress

`func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetPci

`func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) GetPci() string`

GetPci returns the Pci field if non-nil, zero value otherwise.

### GetPciOk

`func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) GetPciOk() (*string, bool)`

GetPciOk returns a tuple with the Pci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPci

`func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) SetPci(v string)`

SetPci sets Pci field to given value.

### HasPci

`func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) HasPci() bool`

HasPci returns a boolean if a field has been set.

### GetSwitchName

`func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) GetSwitchName() string`

GetSwitchName returns the SwitchName field if non-nil, zero value otherwise.

### GetSwitchNameOk

`func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) GetSwitchNameOk() (*string, bool)`

GetSwitchNameOk returns a tuple with the SwitchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchName

`func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) SetSwitchName(v string)`

SetSwitchName sets SwitchName field to given value.

### HasSwitchName

`func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) HasSwitchName() bool`

HasSwitchName returns a boolean if a field has been set.

### GetHost

`func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) GetHost() VirtualizationVmwareHostRelationship`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) GetHostOk() (*VirtualizationVmwareHostRelationship, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) SetHost(v VirtualizationVmwareHostRelationship)`

SetHost sets Host field to given value.

### HasHost

`func (o *VirtualizationVmwarePhysicalNetworkInterfaceAllOf) HasHost() bool`

HasHost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


