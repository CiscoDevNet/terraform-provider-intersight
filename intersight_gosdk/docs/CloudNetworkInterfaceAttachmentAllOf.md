# CloudNetworkInterfaceAttachmentAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cloud.NetworkInterfaceAttachment"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cloud.NetworkInterfaceAttachment"]
**AccessConfig** | Pointer to [**NullableCloudNetworkAccessConfig**](cloud.NetworkAccessConfig.md) |  | [optional] 
**Identity** | Pointer to **string** | The internally generated identity of this network interface attachment. | [optional] [readonly] 
**IpForwardingEnabled** | Pointer to **bool** | Set to true, if IP forwarding is enabled on this interface. | [optional] [readonly] 
**MacAddress** | Pointer to **string** | The MAC (Media Access Control) address of the network interface. | [optional] [readonly] 
**NetworkId** | Pointer to **string** | The identity of the network to which this network interface attachment belongs. | [optional] [readonly] 
**NetworkName** | Pointer to **string** | User friendly name of the network to which this network interface attachment belongs. | [optional] [readonly] 
**NicIndex** | Pointer to **int64** | The device index of the network interface attachment in the virtual machine. | [optional] [readonly] 
**PrivateAddress** | Pointer to [**[]CloudNetworkAddress**](CloudNetworkAddress.md) |  | [optional] 
**PublicAddress** | Pointer to [**[]CloudNetworkAddress**](CloudNetworkAddress.md) |  | [optional] 
**SecurityGroups** | Pointer to **[]string** |  | [optional] 
**SubNetworkId** | Pointer to **string** | The identity of this network interface attachment&#39;s subnet. | [optional] [readonly] 
**SubNetworkName** | Pointer to **string** | User friendly name of this network interface attachment&#39;s subnet. | [optional] [readonly] 

## Methods

### NewCloudNetworkInterfaceAttachmentAllOf

`func NewCloudNetworkInterfaceAttachmentAllOf(classId string, objectType string, ) *CloudNetworkInterfaceAttachmentAllOf`

NewCloudNetworkInterfaceAttachmentAllOf instantiates a new CloudNetworkInterfaceAttachmentAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudNetworkInterfaceAttachmentAllOfWithDefaults

`func NewCloudNetworkInterfaceAttachmentAllOfWithDefaults() *CloudNetworkInterfaceAttachmentAllOf`

NewCloudNetworkInterfaceAttachmentAllOfWithDefaults instantiates a new CloudNetworkInterfaceAttachmentAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudNetworkInterfaceAttachmentAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudNetworkInterfaceAttachmentAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudNetworkInterfaceAttachmentAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudNetworkInterfaceAttachmentAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudNetworkInterfaceAttachmentAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudNetworkInterfaceAttachmentAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccessConfig

`func (o *CloudNetworkInterfaceAttachmentAllOf) GetAccessConfig() CloudNetworkAccessConfig`

GetAccessConfig returns the AccessConfig field if non-nil, zero value otherwise.

### GetAccessConfigOk

`func (o *CloudNetworkInterfaceAttachmentAllOf) GetAccessConfigOk() (*CloudNetworkAccessConfig, bool)`

GetAccessConfigOk returns a tuple with the AccessConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessConfig

`func (o *CloudNetworkInterfaceAttachmentAllOf) SetAccessConfig(v CloudNetworkAccessConfig)`

SetAccessConfig sets AccessConfig field to given value.

### HasAccessConfig

`func (o *CloudNetworkInterfaceAttachmentAllOf) HasAccessConfig() bool`

HasAccessConfig returns a boolean if a field has been set.

### SetAccessConfigNil

`func (o *CloudNetworkInterfaceAttachmentAllOf) SetAccessConfigNil(b bool)`

 SetAccessConfigNil sets the value for AccessConfig to be an explicit nil

### UnsetAccessConfig
`func (o *CloudNetworkInterfaceAttachmentAllOf) UnsetAccessConfig()`

UnsetAccessConfig ensures that no value is present for AccessConfig, not even an explicit nil
### GetIdentity

`func (o *CloudNetworkInterfaceAttachmentAllOf) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *CloudNetworkInterfaceAttachmentAllOf) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *CloudNetworkInterfaceAttachmentAllOf) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *CloudNetworkInterfaceAttachmentAllOf) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetIpForwardingEnabled

`func (o *CloudNetworkInterfaceAttachmentAllOf) GetIpForwardingEnabled() bool`

GetIpForwardingEnabled returns the IpForwardingEnabled field if non-nil, zero value otherwise.

### GetIpForwardingEnabledOk

`func (o *CloudNetworkInterfaceAttachmentAllOf) GetIpForwardingEnabledOk() (*bool, bool)`

GetIpForwardingEnabledOk returns a tuple with the IpForwardingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpForwardingEnabled

`func (o *CloudNetworkInterfaceAttachmentAllOf) SetIpForwardingEnabled(v bool)`

SetIpForwardingEnabled sets IpForwardingEnabled field to given value.

### HasIpForwardingEnabled

`func (o *CloudNetworkInterfaceAttachmentAllOf) HasIpForwardingEnabled() bool`

HasIpForwardingEnabled returns a boolean if a field has been set.

### GetMacAddress

`func (o *CloudNetworkInterfaceAttachmentAllOf) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *CloudNetworkInterfaceAttachmentAllOf) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *CloudNetworkInterfaceAttachmentAllOf) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *CloudNetworkInterfaceAttachmentAllOf) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetNetworkId

`func (o *CloudNetworkInterfaceAttachmentAllOf) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *CloudNetworkInterfaceAttachmentAllOf) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *CloudNetworkInterfaceAttachmentAllOf) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *CloudNetworkInterfaceAttachmentAllOf) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetNetworkName

`func (o *CloudNetworkInterfaceAttachmentAllOf) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *CloudNetworkInterfaceAttachmentAllOf) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *CloudNetworkInterfaceAttachmentAllOf) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *CloudNetworkInterfaceAttachmentAllOf) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

### GetNicIndex

`func (o *CloudNetworkInterfaceAttachmentAllOf) GetNicIndex() int64`

GetNicIndex returns the NicIndex field if non-nil, zero value otherwise.

### GetNicIndexOk

`func (o *CloudNetworkInterfaceAttachmentAllOf) GetNicIndexOk() (*int64, bool)`

GetNicIndexOk returns a tuple with the NicIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicIndex

`func (o *CloudNetworkInterfaceAttachmentAllOf) SetNicIndex(v int64)`

SetNicIndex sets NicIndex field to given value.

### HasNicIndex

`func (o *CloudNetworkInterfaceAttachmentAllOf) HasNicIndex() bool`

HasNicIndex returns a boolean if a field has been set.

### GetPrivateAddress

`func (o *CloudNetworkInterfaceAttachmentAllOf) GetPrivateAddress() []CloudNetworkAddress`

GetPrivateAddress returns the PrivateAddress field if non-nil, zero value otherwise.

### GetPrivateAddressOk

`func (o *CloudNetworkInterfaceAttachmentAllOf) GetPrivateAddressOk() (*[]CloudNetworkAddress, bool)`

GetPrivateAddressOk returns a tuple with the PrivateAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateAddress

`func (o *CloudNetworkInterfaceAttachmentAllOf) SetPrivateAddress(v []CloudNetworkAddress)`

SetPrivateAddress sets PrivateAddress field to given value.

### HasPrivateAddress

`func (o *CloudNetworkInterfaceAttachmentAllOf) HasPrivateAddress() bool`

HasPrivateAddress returns a boolean if a field has been set.

### SetPrivateAddressNil

`func (o *CloudNetworkInterfaceAttachmentAllOf) SetPrivateAddressNil(b bool)`

 SetPrivateAddressNil sets the value for PrivateAddress to be an explicit nil

### UnsetPrivateAddress
`func (o *CloudNetworkInterfaceAttachmentAllOf) UnsetPrivateAddress()`

UnsetPrivateAddress ensures that no value is present for PrivateAddress, not even an explicit nil
### GetPublicAddress

`func (o *CloudNetworkInterfaceAttachmentAllOf) GetPublicAddress() []CloudNetworkAddress`

GetPublicAddress returns the PublicAddress field if non-nil, zero value otherwise.

### GetPublicAddressOk

`func (o *CloudNetworkInterfaceAttachmentAllOf) GetPublicAddressOk() (*[]CloudNetworkAddress, bool)`

GetPublicAddressOk returns a tuple with the PublicAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicAddress

`func (o *CloudNetworkInterfaceAttachmentAllOf) SetPublicAddress(v []CloudNetworkAddress)`

SetPublicAddress sets PublicAddress field to given value.

### HasPublicAddress

`func (o *CloudNetworkInterfaceAttachmentAllOf) HasPublicAddress() bool`

HasPublicAddress returns a boolean if a field has been set.

### SetPublicAddressNil

`func (o *CloudNetworkInterfaceAttachmentAllOf) SetPublicAddressNil(b bool)`

 SetPublicAddressNil sets the value for PublicAddress to be an explicit nil

### UnsetPublicAddress
`func (o *CloudNetworkInterfaceAttachmentAllOf) UnsetPublicAddress()`

UnsetPublicAddress ensures that no value is present for PublicAddress, not even an explicit nil
### GetSecurityGroups

`func (o *CloudNetworkInterfaceAttachmentAllOf) GetSecurityGroups() []string`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *CloudNetworkInterfaceAttachmentAllOf) GetSecurityGroupsOk() (*[]string, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *CloudNetworkInterfaceAttachmentAllOf) SetSecurityGroups(v []string)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *CloudNetworkInterfaceAttachmentAllOf) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *CloudNetworkInterfaceAttachmentAllOf) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *CloudNetworkInterfaceAttachmentAllOf) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil
### GetSubNetworkId

`func (o *CloudNetworkInterfaceAttachmentAllOf) GetSubNetworkId() string`

GetSubNetworkId returns the SubNetworkId field if non-nil, zero value otherwise.

### GetSubNetworkIdOk

`func (o *CloudNetworkInterfaceAttachmentAllOf) GetSubNetworkIdOk() (*string, bool)`

GetSubNetworkIdOk returns a tuple with the SubNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubNetworkId

`func (o *CloudNetworkInterfaceAttachmentAllOf) SetSubNetworkId(v string)`

SetSubNetworkId sets SubNetworkId field to given value.

### HasSubNetworkId

`func (o *CloudNetworkInterfaceAttachmentAllOf) HasSubNetworkId() bool`

HasSubNetworkId returns a boolean if a field has been set.

### GetSubNetworkName

`func (o *CloudNetworkInterfaceAttachmentAllOf) GetSubNetworkName() string`

GetSubNetworkName returns the SubNetworkName field if non-nil, zero value otherwise.

### GetSubNetworkNameOk

`func (o *CloudNetworkInterfaceAttachmentAllOf) GetSubNetworkNameOk() (*string, bool)`

GetSubNetworkNameOk returns a tuple with the SubNetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubNetworkName

`func (o *CloudNetworkInterfaceAttachmentAllOf) SetSubNetworkName(v string)`

SetSubNetworkName sets SubNetworkName field to given value.

### HasSubNetworkName

`func (o *CloudNetworkInterfaceAttachmentAllOf) HasSubNetworkName() bool`

HasSubNetworkName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


