# ManagementInterfaceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "management.Interface"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "management.Interface"]
**Gateway** | Pointer to **string** | Default gateway for the interface. | [optional] [readonly] 
**HostName** | Pointer to **string** | Hostname configured for the interface. | [optional] 
**IpAddress** | Pointer to **string** | IP address of the interface. | [optional] [readonly] 
**Ipv4Address** | Pointer to **string** | IPv4 address of the interface. | [optional] [readonly] 
**Ipv4Gateway** | Pointer to **string** | IPv4 default gateway for the interface. | [optional] [readonly] 
**Ipv4Mask** | Pointer to **string** | IPv4 Netmask for the interface. | [optional] [readonly] 
**Ipv6Address** | Pointer to **string** | IPv6 address of the interface. | [optional] 
**Ipv6Gateway** | Pointer to **string** | IPv6 default gateway for the interface. | [optional] 
**Ipv6Prefix** | Pointer to **int64** | IPv6 prefix for the interface. | [optional] 
**MacAddress** | Pointer to **string** | MAC address configured for the interface. | [optional] [readonly] 
**Mask** | Pointer to **string** | Netmask for the interface. | [optional] [readonly] 
**SwitchId** | Pointer to **string** | Switch Id connected to the interface. | [optional] 
**UemConnStatus** | Pointer to **string** | The event channel connection status for the interface. | [optional] 
**VirtualHostName** | Pointer to **string** | Virtual hostname configured for the interface in case of clustered environment. | [optional] 
**VlanId** | Pointer to **int64** | VlanId configured for the interface. | [optional] 
**InventoryDeviceInfo** | Pointer to [**InventoryDeviceInfoRelationship**](InventoryDeviceInfoRelationship.md) |  | [optional] 
**ManagementController** | Pointer to [**ManagementControllerRelationship**](ManagementControllerRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewManagementInterfaceAllOf

`func NewManagementInterfaceAllOf(classId string, objectType string, ) *ManagementInterfaceAllOf`

NewManagementInterfaceAllOf instantiates a new ManagementInterfaceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementInterfaceAllOfWithDefaults

`func NewManagementInterfaceAllOfWithDefaults() *ManagementInterfaceAllOf`

NewManagementInterfaceAllOfWithDefaults instantiates a new ManagementInterfaceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ManagementInterfaceAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ManagementInterfaceAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ManagementInterfaceAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ManagementInterfaceAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ManagementInterfaceAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ManagementInterfaceAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetGateway

`func (o *ManagementInterfaceAllOf) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *ManagementInterfaceAllOf) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *ManagementInterfaceAllOf) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *ManagementInterfaceAllOf) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetHostName

`func (o *ManagementInterfaceAllOf) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *ManagementInterfaceAllOf) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *ManagementInterfaceAllOf) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *ManagementInterfaceAllOf) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetIpAddress

`func (o *ManagementInterfaceAllOf) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ManagementInterfaceAllOf) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ManagementInterfaceAllOf) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *ManagementInterfaceAllOf) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetIpv4Address

`func (o *ManagementInterfaceAllOf) GetIpv4Address() string`

GetIpv4Address returns the Ipv4Address field if non-nil, zero value otherwise.

### GetIpv4AddressOk

`func (o *ManagementInterfaceAllOf) GetIpv4AddressOk() (*string, bool)`

GetIpv4AddressOk returns a tuple with the Ipv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Address

`func (o *ManagementInterfaceAllOf) SetIpv4Address(v string)`

SetIpv4Address sets Ipv4Address field to given value.

### HasIpv4Address

`func (o *ManagementInterfaceAllOf) HasIpv4Address() bool`

HasIpv4Address returns a boolean if a field has been set.

### GetIpv4Gateway

`func (o *ManagementInterfaceAllOf) GetIpv4Gateway() string`

GetIpv4Gateway returns the Ipv4Gateway field if non-nil, zero value otherwise.

### GetIpv4GatewayOk

`func (o *ManagementInterfaceAllOf) GetIpv4GatewayOk() (*string, bool)`

GetIpv4GatewayOk returns a tuple with the Ipv4Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Gateway

`func (o *ManagementInterfaceAllOf) SetIpv4Gateway(v string)`

SetIpv4Gateway sets Ipv4Gateway field to given value.

### HasIpv4Gateway

`func (o *ManagementInterfaceAllOf) HasIpv4Gateway() bool`

HasIpv4Gateway returns a boolean if a field has been set.

### GetIpv4Mask

`func (o *ManagementInterfaceAllOf) GetIpv4Mask() string`

GetIpv4Mask returns the Ipv4Mask field if non-nil, zero value otherwise.

### GetIpv4MaskOk

`func (o *ManagementInterfaceAllOf) GetIpv4MaskOk() (*string, bool)`

GetIpv4MaskOk returns a tuple with the Ipv4Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Mask

`func (o *ManagementInterfaceAllOf) SetIpv4Mask(v string)`

SetIpv4Mask sets Ipv4Mask field to given value.

### HasIpv4Mask

`func (o *ManagementInterfaceAllOf) HasIpv4Mask() bool`

HasIpv4Mask returns a boolean if a field has been set.

### GetIpv6Address

`func (o *ManagementInterfaceAllOf) GetIpv6Address() string`

GetIpv6Address returns the Ipv6Address field if non-nil, zero value otherwise.

### GetIpv6AddressOk

`func (o *ManagementInterfaceAllOf) GetIpv6AddressOk() (*string, bool)`

GetIpv6AddressOk returns a tuple with the Ipv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Address

`func (o *ManagementInterfaceAllOf) SetIpv6Address(v string)`

SetIpv6Address sets Ipv6Address field to given value.

### HasIpv6Address

`func (o *ManagementInterfaceAllOf) HasIpv6Address() bool`

HasIpv6Address returns a boolean if a field has been set.

### GetIpv6Gateway

`func (o *ManagementInterfaceAllOf) GetIpv6Gateway() string`

GetIpv6Gateway returns the Ipv6Gateway field if non-nil, zero value otherwise.

### GetIpv6GatewayOk

`func (o *ManagementInterfaceAllOf) GetIpv6GatewayOk() (*string, bool)`

GetIpv6GatewayOk returns a tuple with the Ipv6Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Gateway

`func (o *ManagementInterfaceAllOf) SetIpv6Gateway(v string)`

SetIpv6Gateway sets Ipv6Gateway field to given value.

### HasIpv6Gateway

`func (o *ManagementInterfaceAllOf) HasIpv6Gateway() bool`

HasIpv6Gateway returns a boolean if a field has been set.

### GetIpv6Prefix

`func (o *ManagementInterfaceAllOf) GetIpv6Prefix() int64`

GetIpv6Prefix returns the Ipv6Prefix field if non-nil, zero value otherwise.

### GetIpv6PrefixOk

`func (o *ManagementInterfaceAllOf) GetIpv6PrefixOk() (*int64, bool)`

GetIpv6PrefixOk returns a tuple with the Ipv6Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Prefix

`func (o *ManagementInterfaceAllOf) SetIpv6Prefix(v int64)`

SetIpv6Prefix sets Ipv6Prefix field to given value.

### HasIpv6Prefix

`func (o *ManagementInterfaceAllOf) HasIpv6Prefix() bool`

HasIpv6Prefix returns a boolean if a field has been set.

### GetMacAddress

`func (o *ManagementInterfaceAllOf) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *ManagementInterfaceAllOf) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *ManagementInterfaceAllOf) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *ManagementInterfaceAllOf) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMask

`func (o *ManagementInterfaceAllOf) GetMask() string`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *ManagementInterfaceAllOf) GetMaskOk() (*string, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *ManagementInterfaceAllOf) SetMask(v string)`

SetMask sets Mask field to given value.

### HasMask

`func (o *ManagementInterfaceAllOf) HasMask() bool`

HasMask returns a boolean if a field has been set.

### GetSwitchId

`func (o *ManagementInterfaceAllOf) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *ManagementInterfaceAllOf) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *ManagementInterfaceAllOf) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *ManagementInterfaceAllOf) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetUemConnStatus

`func (o *ManagementInterfaceAllOf) GetUemConnStatus() string`

GetUemConnStatus returns the UemConnStatus field if non-nil, zero value otherwise.

### GetUemConnStatusOk

`func (o *ManagementInterfaceAllOf) GetUemConnStatusOk() (*string, bool)`

GetUemConnStatusOk returns a tuple with the UemConnStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUemConnStatus

`func (o *ManagementInterfaceAllOf) SetUemConnStatus(v string)`

SetUemConnStatus sets UemConnStatus field to given value.

### HasUemConnStatus

`func (o *ManagementInterfaceAllOf) HasUemConnStatus() bool`

HasUemConnStatus returns a boolean if a field has been set.

### GetVirtualHostName

`func (o *ManagementInterfaceAllOf) GetVirtualHostName() string`

GetVirtualHostName returns the VirtualHostName field if non-nil, zero value otherwise.

### GetVirtualHostNameOk

`func (o *ManagementInterfaceAllOf) GetVirtualHostNameOk() (*string, bool)`

GetVirtualHostNameOk returns a tuple with the VirtualHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualHostName

`func (o *ManagementInterfaceAllOf) SetVirtualHostName(v string)`

SetVirtualHostName sets VirtualHostName field to given value.

### HasVirtualHostName

`func (o *ManagementInterfaceAllOf) HasVirtualHostName() bool`

HasVirtualHostName returns a boolean if a field has been set.

### GetVlanId

`func (o *ManagementInterfaceAllOf) GetVlanId() int64`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *ManagementInterfaceAllOf) GetVlanIdOk() (*int64, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *ManagementInterfaceAllOf) SetVlanId(v int64)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *ManagementInterfaceAllOf) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetInventoryDeviceInfo

`func (o *ManagementInterfaceAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship`

GetInventoryDeviceInfo returns the InventoryDeviceInfo field if non-nil, zero value otherwise.

### GetInventoryDeviceInfoOk

`func (o *ManagementInterfaceAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool)`

GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryDeviceInfo

`func (o *ManagementInterfaceAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship)`

SetInventoryDeviceInfo sets InventoryDeviceInfo field to given value.

### HasInventoryDeviceInfo

`func (o *ManagementInterfaceAllOf) HasInventoryDeviceInfo() bool`

HasInventoryDeviceInfo returns a boolean if a field has been set.

### GetManagementController

`func (o *ManagementInterfaceAllOf) GetManagementController() ManagementControllerRelationship`

GetManagementController returns the ManagementController field if non-nil, zero value otherwise.

### GetManagementControllerOk

`func (o *ManagementInterfaceAllOf) GetManagementControllerOk() (*ManagementControllerRelationship, bool)`

GetManagementControllerOk returns a tuple with the ManagementController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementController

`func (o *ManagementInterfaceAllOf) SetManagementController(v ManagementControllerRelationship)`

SetManagementController sets ManagementController field to given value.

### HasManagementController

`func (o *ManagementInterfaceAllOf) HasManagementController() bool`

HasManagementController returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *ManagementInterfaceAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *ManagementInterfaceAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *ManagementInterfaceAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *ManagementInterfaceAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


