# BootHttp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "boot.Http"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "boot.Http"]
**InterfaceName** | Pointer to **string** | For integrated NICs, specify the underlying network interface card name (such as 1 or 2). For virtual NICs, specify the vNIC name (customized name like eth0). | [optional] 
**InterfaceSource** | Pointer to **string** | Lists the supported Interface Source for HTTP device. Supported values are \&quot;name\&quot; and \&quot;mac\&quot;. * &#x60;name&#x60; - Use interface name to select virtual ethernet interface. * &#x60;mac&#x60; - Use MAC address to select virtual ethernet interface. * &#x60;port&#x60; - Use port to select virtual ethernet interface. | [optional] [default to "name"]
**IpConfigType** | Pointer to **string** | The IP config type to use during the HTTP boot process. For DHCP configuration, the IP address, DNS server, netmask and gateway details are obtained from DHCP server. For static configuration, please provide the IP address, DNS server, netmask, and gateway details. * &#x60;DHCP&#x60; - The type of the IP config is DHCP. * &#x60;Static&#x60; - The type of the IP config is Static. | [optional] [default to "DHCP"]
**IpType** | Pointer to **string** | The IP address family type to use during the HTTP boot process. * &#x60;IPv4&#x60; - The type of the IP address is IPv4. * &#x60;IPv6&#x60; - The type of the IP address is IPv6. | [optional] [default to "IPv4"]
**MacAddress** | Pointer to **string** | The MAC Address of the underlying virtual ethernet interface used by the HTTP boot device. | [optional] 
**Port** | Pointer to **int64** | The Port ID of the adapter on which the underlying virtual ethernet interface is present. If no port is specified, the default value is -1. Supported values are 0 to 255. | [optional] [default to -1]
**Protocol** | Pointer to **string** | Protocol to be used for HTTP boot. HTTPS require root certificate for authentication. * &#x60;HTTPS&#x60; - Secure HTTP protocol, certificate required for authentication. * &#x60;HTTP&#x60; - HTTP protocol without security certificate requirement. | [optional] [default to "HTTPS"]
**Slot** | Pointer to **string** | The slot ID of the adapter on which the underlying virtual ethernet interface is present. Supported values are ( 1 - 255, \&quot;MLOM\&quot;, \&quot;L\&quot;, \&quot;L1\&quot;, \&quot;L2\&quot;, \&quot;OCP\&quot;). | [optional] 
**StaticIpV4Settings** | Pointer to [**NullableBootStaticIpV4Settings**](BootStaticIpV4Settings.md) |  | [optional] 
**StaticIpV6Settings** | Pointer to [**NullableBootStaticIpV6Settings**](BootStaticIpV6Settings.md) |  | [optional] 
**Uri** | Pointer to **string** | Boot resource location in URI format. | [optional] 

## Methods

### NewBootHttp

`func NewBootHttp(classId string, objectType string, ) *BootHttp`

NewBootHttp instantiates a new BootHttp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootHttpWithDefaults

`func NewBootHttpWithDefaults() *BootHttp`

NewBootHttpWithDefaults instantiates a new BootHttp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BootHttp) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BootHttp) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BootHttp) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BootHttp) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BootHttp) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BootHttp) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetInterfaceName

`func (o *BootHttp) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *BootHttp) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *BootHttp) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *BootHttp) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### GetInterfaceSource

`func (o *BootHttp) GetInterfaceSource() string`

GetInterfaceSource returns the InterfaceSource field if non-nil, zero value otherwise.

### GetInterfaceSourceOk

`func (o *BootHttp) GetInterfaceSourceOk() (*string, bool)`

GetInterfaceSourceOk returns a tuple with the InterfaceSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceSource

`func (o *BootHttp) SetInterfaceSource(v string)`

SetInterfaceSource sets InterfaceSource field to given value.

### HasInterfaceSource

`func (o *BootHttp) HasInterfaceSource() bool`

HasInterfaceSource returns a boolean if a field has been set.

### GetIpConfigType

`func (o *BootHttp) GetIpConfigType() string`

GetIpConfigType returns the IpConfigType field if non-nil, zero value otherwise.

### GetIpConfigTypeOk

`func (o *BootHttp) GetIpConfigTypeOk() (*string, bool)`

GetIpConfigTypeOk returns a tuple with the IpConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpConfigType

`func (o *BootHttp) SetIpConfigType(v string)`

SetIpConfigType sets IpConfigType field to given value.

### HasIpConfigType

`func (o *BootHttp) HasIpConfigType() bool`

HasIpConfigType returns a boolean if a field has been set.

### GetIpType

`func (o *BootHttp) GetIpType() string`

GetIpType returns the IpType field if non-nil, zero value otherwise.

### GetIpTypeOk

`func (o *BootHttp) GetIpTypeOk() (*string, bool)`

GetIpTypeOk returns a tuple with the IpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpType

`func (o *BootHttp) SetIpType(v string)`

SetIpType sets IpType field to given value.

### HasIpType

`func (o *BootHttp) HasIpType() bool`

HasIpType returns a boolean if a field has been set.

### GetMacAddress

`func (o *BootHttp) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *BootHttp) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *BootHttp) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *BootHttp) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetPort

`func (o *BootHttp) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *BootHttp) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *BootHttp) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *BootHttp) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *BootHttp) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *BootHttp) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *BootHttp) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *BootHttp) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSlot

`func (o *BootHttp) GetSlot() string`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *BootHttp) GetSlotOk() (*string, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *BootHttp) SetSlot(v string)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *BootHttp) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### GetStaticIpV4Settings

`func (o *BootHttp) GetStaticIpV4Settings() BootStaticIpV4Settings`

GetStaticIpV4Settings returns the StaticIpV4Settings field if non-nil, zero value otherwise.

### GetStaticIpV4SettingsOk

`func (o *BootHttp) GetStaticIpV4SettingsOk() (*BootStaticIpV4Settings, bool)`

GetStaticIpV4SettingsOk returns a tuple with the StaticIpV4Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticIpV4Settings

`func (o *BootHttp) SetStaticIpV4Settings(v BootStaticIpV4Settings)`

SetStaticIpV4Settings sets StaticIpV4Settings field to given value.

### HasStaticIpV4Settings

`func (o *BootHttp) HasStaticIpV4Settings() bool`

HasStaticIpV4Settings returns a boolean if a field has been set.

### SetStaticIpV4SettingsNil

`func (o *BootHttp) SetStaticIpV4SettingsNil(b bool)`

 SetStaticIpV4SettingsNil sets the value for StaticIpV4Settings to be an explicit nil

### UnsetStaticIpV4Settings
`func (o *BootHttp) UnsetStaticIpV4Settings()`

UnsetStaticIpV4Settings ensures that no value is present for StaticIpV4Settings, not even an explicit nil
### GetStaticIpV6Settings

`func (o *BootHttp) GetStaticIpV6Settings() BootStaticIpV6Settings`

GetStaticIpV6Settings returns the StaticIpV6Settings field if non-nil, zero value otherwise.

### GetStaticIpV6SettingsOk

`func (o *BootHttp) GetStaticIpV6SettingsOk() (*BootStaticIpV6Settings, bool)`

GetStaticIpV6SettingsOk returns a tuple with the StaticIpV6Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticIpV6Settings

`func (o *BootHttp) SetStaticIpV6Settings(v BootStaticIpV6Settings)`

SetStaticIpV6Settings sets StaticIpV6Settings field to given value.

### HasStaticIpV6Settings

`func (o *BootHttp) HasStaticIpV6Settings() bool`

HasStaticIpV6Settings returns a boolean if a field has been set.

### SetStaticIpV6SettingsNil

`func (o *BootHttp) SetStaticIpV6SettingsNil(b bool)`

 SetStaticIpV6SettingsNil sets the value for StaticIpV6Settings to be an explicit nil

### UnsetStaticIpV6Settings
`func (o *BootHttp) UnsetStaticIpV6Settings()`

UnsetStaticIpV6Settings ensures that no value is present for StaticIpV6Settings, not even an explicit nil
### GetUri

`func (o *BootHttp) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *BootHttp) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *BootHttp) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *BootHttp) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


