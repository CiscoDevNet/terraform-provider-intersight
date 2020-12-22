# BootPxeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "boot.Pxe"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "boot.Pxe"]
**InterfaceName** | Pointer to **string** | The name of the underlying virtual ethernet interface used by the PXE boot device. | [optional] 
**InterfaceSource** | Pointer to **string** | Lists the supported Interface Source for PXE device. Supported values are \&quot;name\&quot; and \&quot;mac\&quot;. * &#x60;name&#x60; - Use interface name to select virtual ethernet interface. * &#x60;mac&#x60; - Use MAC address to select virtual ethernet interface. * &#x60;port&#x60; - Use port to select virtual ethernet interface. | [optional] [default to "name"]
**IpType** | Pointer to **string** | The IP Address family type to use during the PXE Boot process. * &#x60;None&#x60; - Default value if IpType is not specified. * &#x60;IPv4&#x60; - The IPv4 address family type. * &#x60;IPv6&#x60; - The IPv6 address family type. | [optional] [default to "None"]
**MacAddress** | Pointer to **string** | The MAC Address of the underlying virtual ethernet interface used by the PXE boot device. | [optional] 
**Port** | Pointer to **int64** | The Port ID of the adapter on which the underlying virtual ethernet interface is present. If no port is specified, the default value is -1. Supported values are -1 to 255. | [optional] [default to -1]
**Slot** | Pointer to **string** | The slot ID of the adapter on which the underlying virtual ethernet interface is present. Supported values are ( 1 - 255, \&quot;MLOM\&quot;, \&quot;L\&quot;, \&quot;L1\&quot;, \&quot;L2\&quot;, \&quot;OCP\&quot;). | [optional] 

## Methods

### NewBootPxeAllOf

`func NewBootPxeAllOf(classId string, objectType string, ) *BootPxeAllOf`

NewBootPxeAllOf instantiates a new BootPxeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootPxeAllOfWithDefaults

`func NewBootPxeAllOfWithDefaults() *BootPxeAllOf`

NewBootPxeAllOfWithDefaults instantiates a new BootPxeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BootPxeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BootPxeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BootPxeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BootPxeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BootPxeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BootPxeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetInterfaceName

`func (o *BootPxeAllOf) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *BootPxeAllOf) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *BootPxeAllOf) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *BootPxeAllOf) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### GetInterfaceSource

`func (o *BootPxeAllOf) GetInterfaceSource() string`

GetInterfaceSource returns the InterfaceSource field if non-nil, zero value otherwise.

### GetInterfaceSourceOk

`func (o *BootPxeAllOf) GetInterfaceSourceOk() (*string, bool)`

GetInterfaceSourceOk returns a tuple with the InterfaceSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceSource

`func (o *BootPxeAllOf) SetInterfaceSource(v string)`

SetInterfaceSource sets InterfaceSource field to given value.

### HasInterfaceSource

`func (o *BootPxeAllOf) HasInterfaceSource() bool`

HasInterfaceSource returns a boolean if a field has been set.

### GetIpType

`func (o *BootPxeAllOf) GetIpType() string`

GetIpType returns the IpType field if non-nil, zero value otherwise.

### GetIpTypeOk

`func (o *BootPxeAllOf) GetIpTypeOk() (*string, bool)`

GetIpTypeOk returns a tuple with the IpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpType

`func (o *BootPxeAllOf) SetIpType(v string)`

SetIpType sets IpType field to given value.

### HasIpType

`func (o *BootPxeAllOf) HasIpType() bool`

HasIpType returns a boolean if a field has been set.

### GetMacAddress

`func (o *BootPxeAllOf) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *BootPxeAllOf) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *BootPxeAllOf) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *BootPxeAllOf) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetPort

`func (o *BootPxeAllOf) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *BootPxeAllOf) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *BootPxeAllOf) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *BootPxeAllOf) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSlot

`func (o *BootPxeAllOf) GetSlot() string`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *BootPxeAllOf) GetSlotOk() (*string, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *BootPxeAllOf) SetSlot(v string)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *BootPxeAllOf) HasSlot() bool`

HasSlot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


