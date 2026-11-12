# NetworkIpAddressInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "network.IpAddressInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "network.IpAddressInfo"]
**Address** | Pointer to **string** | The IPv4 or IPv6 address string. Empty placeholder addresses are not stored; a network.IpAddressInfo record is only created when a non-empty address is present. | [optional] [readonly] 
**AddressFamily** | Pointer to **string** | The address family of this IP address (IPv4 or IPv6). * &#x60;&#x60; - The address family is not available or not applicable. * &#x60;IPv4&#x60; - The address is an IPv4 address. * &#x60;IPv6&#x60; - The address is an IPv6 address. | [optional] [readonly] [default to ""]
**AddressOrigin** | Pointer to **string** | The origin of this IP address assignment. Normalized from platform-specific values to DHCP, SLAAC, Static, LinkLocal, or NotApplicable. * &#x60;&#x60; - The address origin is not available or not applicable. * &#x60;DHCP&#x60; - The IP address was assigned via DHCP or DHCPv6. * &#x60;SLAAC&#x60; - The IPv6 address was assigned via Stateless Address Autoconfiguration (SLAAC). * &#x60;Static&#x60; - The IP address is statically configured. * &#x60;LinkLocal&#x60; - The IPv6 address is a link-local address auto-derived from the interface MAC. | [optional] [readonly] [default to ""]
**AddressState** | Pointer to **string** | The operational state of this IP address (Preferred, Deprecated, Failed). Set to NotApplicable for IPv4 addresses where per-address state is not applicable. * &#x60;&#x60; - The address state is not available or not applicable (e.g., for IPv4 addresses). * &#x60;Preferred&#x60; - The IP address is preferred and fully operational. * &#x60;Deprecated&#x60; - The IP address is deprecated (preferred lifetime has expired). * &#x60;Failed&#x60; - The IP address has failed (e.g., Duplicate Address Detection failure for IPv6). | [optional] [readonly] [default to ""]
**InterfaceType** | Pointer to **string** | The management interface type this address is associated with. OutOfBand for the OOB management interface; InBand for the inband management interface. * &#x60;&#x60; - The interface type is not available or not applicable. * &#x60;OutOfBand&#x60; - The IP address is associated with the out-of-band management interface. * &#x60;InBand&#x60; - The IP address is associated with the in-band management interface. | [optional] [readonly] [default to ""]
**IpAddressError** | Pointer to **string** | Any error or diagnostic message associated with this IP address. Empty when no error is reported. | [optional] [readonly] 
**PrefixLength** | Pointer to **string** | The prefix length of this IP address in bits. Used for IPv6 addresses. Set to zero for IPv4 addresses where SubnetMask is used instead. | [optional] [readonly] 
**SubnetMask** | Pointer to **string** | The subnet mask of this IP address in dotted-decimal notation. Used for IPv4 addresses. Empty for IPv6 addresses where PrefixLength is used instead. | [optional] [readonly] 
**NetworkElement** | Pointer to [**NullableNetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNetworkIpAddressInfo

`func NewNetworkIpAddressInfo(classId string, objectType string, ) *NetworkIpAddressInfo`

NewNetworkIpAddressInfo instantiates a new NetworkIpAddressInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkIpAddressInfoWithDefaults

`func NewNetworkIpAddressInfoWithDefaults() *NetworkIpAddressInfo`

NewNetworkIpAddressInfoWithDefaults instantiates a new NetworkIpAddressInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NetworkIpAddressInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NetworkIpAddressInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NetworkIpAddressInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NetworkIpAddressInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NetworkIpAddressInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NetworkIpAddressInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAddress

`func (o *NetworkIpAddressInfo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NetworkIpAddressInfo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NetworkIpAddressInfo) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *NetworkIpAddressInfo) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAddressFamily

`func (o *NetworkIpAddressInfo) GetAddressFamily() string`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *NetworkIpAddressInfo) GetAddressFamilyOk() (*string, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *NetworkIpAddressInfo) SetAddressFamily(v string)`

SetAddressFamily sets AddressFamily field to given value.

### HasAddressFamily

`func (o *NetworkIpAddressInfo) HasAddressFamily() bool`

HasAddressFamily returns a boolean if a field has been set.

### GetAddressOrigin

`func (o *NetworkIpAddressInfo) GetAddressOrigin() string`

GetAddressOrigin returns the AddressOrigin field if non-nil, zero value otherwise.

### GetAddressOriginOk

`func (o *NetworkIpAddressInfo) GetAddressOriginOk() (*string, bool)`

GetAddressOriginOk returns a tuple with the AddressOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressOrigin

`func (o *NetworkIpAddressInfo) SetAddressOrigin(v string)`

SetAddressOrigin sets AddressOrigin field to given value.

### HasAddressOrigin

`func (o *NetworkIpAddressInfo) HasAddressOrigin() bool`

HasAddressOrigin returns a boolean if a field has been set.

### GetAddressState

`func (o *NetworkIpAddressInfo) GetAddressState() string`

GetAddressState returns the AddressState field if non-nil, zero value otherwise.

### GetAddressStateOk

`func (o *NetworkIpAddressInfo) GetAddressStateOk() (*string, bool)`

GetAddressStateOk returns a tuple with the AddressState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressState

`func (o *NetworkIpAddressInfo) SetAddressState(v string)`

SetAddressState sets AddressState field to given value.

### HasAddressState

`func (o *NetworkIpAddressInfo) HasAddressState() bool`

HasAddressState returns a boolean if a field has been set.

### GetInterfaceType

`func (o *NetworkIpAddressInfo) GetInterfaceType() string`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *NetworkIpAddressInfo) GetInterfaceTypeOk() (*string, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *NetworkIpAddressInfo) SetInterfaceType(v string)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *NetworkIpAddressInfo) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### GetIpAddressError

`func (o *NetworkIpAddressInfo) GetIpAddressError() string`

GetIpAddressError returns the IpAddressError field if non-nil, zero value otherwise.

### GetIpAddressErrorOk

`func (o *NetworkIpAddressInfo) GetIpAddressErrorOk() (*string, bool)`

GetIpAddressErrorOk returns a tuple with the IpAddressError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddressError

`func (o *NetworkIpAddressInfo) SetIpAddressError(v string)`

SetIpAddressError sets IpAddressError field to given value.

### HasIpAddressError

`func (o *NetworkIpAddressInfo) HasIpAddressError() bool`

HasIpAddressError returns a boolean if a field has been set.

### GetPrefixLength

`func (o *NetworkIpAddressInfo) GetPrefixLength() string`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *NetworkIpAddressInfo) GetPrefixLengthOk() (*string, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *NetworkIpAddressInfo) SetPrefixLength(v string)`

SetPrefixLength sets PrefixLength field to given value.

### HasPrefixLength

`func (o *NetworkIpAddressInfo) HasPrefixLength() bool`

HasPrefixLength returns a boolean if a field has been set.

### GetSubnetMask

`func (o *NetworkIpAddressInfo) GetSubnetMask() string`

GetSubnetMask returns the SubnetMask field if non-nil, zero value otherwise.

### GetSubnetMaskOk

`func (o *NetworkIpAddressInfo) GetSubnetMaskOk() (*string, bool)`

GetSubnetMaskOk returns a tuple with the SubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetMask

`func (o *NetworkIpAddressInfo) SetSubnetMask(v string)`

SetSubnetMask sets SubnetMask field to given value.

### HasSubnetMask

`func (o *NetworkIpAddressInfo) HasSubnetMask() bool`

HasSubnetMask returns a boolean if a field has been set.

### GetNetworkElement

`func (o *NetworkIpAddressInfo) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *NetworkIpAddressInfo) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *NetworkIpAddressInfo) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *NetworkIpAddressInfo) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### SetNetworkElementNil

`func (o *NetworkIpAddressInfo) SetNetworkElementNil(b bool)`

 SetNetworkElementNil sets the value for NetworkElement to be an explicit nil

### UnsetNetworkElement
`func (o *NetworkIpAddressInfo) UnsetNetworkElement()`

UnsetNetworkElement ensures that no value is present for NetworkElement, not even an explicit nil
### GetRegisteredDevice

`func (o *NetworkIpAddressInfo) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NetworkIpAddressInfo) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NetworkIpAddressInfo) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NetworkIpAddressInfo) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *NetworkIpAddressInfo) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *NetworkIpAddressInfo) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


