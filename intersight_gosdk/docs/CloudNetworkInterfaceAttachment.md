# CloudNetworkInterfaceAttachment

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

### NewCloudNetworkInterfaceAttachment

`func NewCloudNetworkInterfaceAttachment(classId string, objectType string, ) *CloudNetworkInterfaceAttachment`

NewCloudNetworkInterfaceAttachment instantiates a new CloudNetworkInterfaceAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudNetworkInterfaceAttachmentWithDefaults

`func NewCloudNetworkInterfaceAttachmentWithDefaults() *CloudNetworkInterfaceAttachment`

NewCloudNetworkInterfaceAttachmentWithDefaults instantiates a new CloudNetworkInterfaceAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudNetworkInterfaceAttachment) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudNetworkInterfaceAttachment) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudNetworkInterfaceAttachment) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudNetworkInterfaceAttachment) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudNetworkInterfaceAttachment) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudNetworkInterfaceAttachment) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccessConfig

`func (o *CloudNetworkInterfaceAttachment) GetAccessConfig() CloudNetworkAccessConfig`

GetAccessConfig returns the AccessConfig field if non-nil, zero value otherwise.

### GetAccessConfigOk

`func (o *CloudNetworkInterfaceAttachment) GetAccessConfigOk() (*CloudNetworkAccessConfig, bool)`

GetAccessConfigOk returns a tuple with the AccessConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessConfig

`func (o *CloudNetworkInterfaceAttachment) SetAccessConfig(v CloudNetworkAccessConfig)`

SetAccessConfig sets AccessConfig field to given value.

### HasAccessConfig

`func (o *CloudNetworkInterfaceAttachment) HasAccessConfig() bool`

HasAccessConfig returns a boolean if a field has been set.

### SetAccessConfigNil

`func (o *CloudNetworkInterfaceAttachment) SetAccessConfigNil(b bool)`

 SetAccessConfigNil sets the value for AccessConfig to be an explicit nil

### UnsetAccessConfig
`func (o *CloudNetworkInterfaceAttachment) UnsetAccessConfig()`

UnsetAccessConfig ensures that no value is present for AccessConfig, not even an explicit nil
### GetIdentity

`func (o *CloudNetworkInterfaceAttachment) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *CloudNetworkInterfaceAttachment) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *CloudNetworkInterfaceAttachment) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *CloudNetworkInterfaceAttachment) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetIpForwardingEnabled

`func (o *CloudNetworkInterfaceAttachment) GetIpForwardingEnabled() bool`

GetIpForwardingEnabled returns the IpForwardingEnabled field if non-nil, zero value otherwise.

### GetIpForwardingEnabledOk

`func (o *CloudNetworkInterfaceAttachment) GetIpForwardingEnabledOk() (*bool, bool)`

GetIpForwardingEnabledOk returns a tuple with the IpForwardingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpForwardingEnabled

`func (o *CloudNetworkInterfaceAttachment) SetIpForwardingEnabled(v bool)`

SetIpForwardingEnabled sets IpForwardingEnabled field to given value.

### HasIpForwardingEnabled

`func (o *CloudNetworkInterfaceAttachment) HasIpForwardingEnabled() bool`

HasIpForwardingEnabled returns a boolean if a field has been set.

### GetMacAddress

`func (o *CloudNetworkInterfaceAttachment) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *CloudNetworkInterfaceAttachment) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *CloudNetworkInterfaceAttachment) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *CloudNetworkInterfaceAttachment) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetNetworkId

`func (o *CloudNetworkInterfaceAttachment) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *CloudNetworkInterfaceAttachment) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *CloudNetworkInterfaceAttachment) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *CloudNetworkInterfaceAttachment) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetNetworkName

`func (o *CloudNetworkInterfaceAttachment) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *CloudNetworkInterfaceAttachment) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *CloudNetworkInterfaceAttachment) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *CloudNetworkInterfaceAttachment) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

### GetNicIndex

`func (o *CloudNetworkInterfaceAttachment) GetNicIndex() int64`

GetNicIndex returns the NicIndex field if non-nil, zero value otherwise.

### GetNicIndexOk

`func (o *CloudNetworkInterfaceAttachment) GetNicIndexOk() (*int64, bool)`

GetNicIndexOk returns a tuple with the NicIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicIndex

`func (o *CloudNetworkInterfaceAttachment) SetNicIndex(v int64)`

SetNicIndex sets NicIndex field to given value.

### HasNicIndex

`func (o *CloudNetworkInterfaceAttachment) HasNicIndex() bool`

HasNicIndex returns a boolean if a field has been set.

### GetPrivateAddress

`func (o *CloudNetworkInterfaceAttachment) GetPrivateAddress() []CloudNetworkAddress`

GetPrivateAddress returns the PrivateAddress field if non-nil, zero value otherwise.

### GetPrivateAddressOk

`func (o *CloudNetworkInterfaceAttachment) GetPrivateAddressOk() (*[]CloudNetworkAddress, bool)`

GetPrivateAddressOk returns a tuple with the PrivateAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateAddress

`func (o *CloudNetworkInterfaceAttachment) SetPrivateAddress(v []CloudNetworkAddress)`

SetPrivateAddress sets PrivateAddress field to given value.

### HasPrivateAddress

`func (o *CloudNetworkInterfaceAttachment) HasPrivateAddress() bool`

HasPrivateAddress returns a boolean if a field has been set.

### SetPrivateAddressNil

`func (o *CloudNetworkInterfaceAttachment) SetPrivateAddressNil(b bool)`

 SetPrivateAddressNil sets the value for PrivateAddress to be an explicit nil

### UnsetPrivateAddress
`func (o *CloudNetworkInterfaceAttachment) UnsetPrivateAddress()`

UnsetPrivateAddress ensures that no value is present for PrivateAddress, not even an explicit nil
### GetPublicAddress

`func (o *CloudNetworkInterfaceAttachment) GetPublicAddress() []CloudNetworkAddress`

GetPublicAddress returns the PublicAddress field if non-nil, zero value otherwise.

### GetPublicAddressOk

`func (o *CloudNetworkInterfaceAttachment) GetPublicAddressOk() (*[]CloudNetworkAddress, bool)`

GetPublicAddressOk returns a tuple with the PublicAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicAddress

`func (o *CloudNetworkInterfaceAttachment) SetPublicAddress(v []CloudNetworkAddress)`

SetPublicAddress sets PublicAddress field to given value.

### HasPublicAddress

`func (o *CloudNetworkInterfaceAttachment) HasPublicAddress() bool`

HasPublicAddress returns a boolean if a field has been set.

### SetPublicAddressNil

`func (o *CloudNetworkInterfaceAttachment) SetPublicAddressNil(b bool)`

 SetPublicAddressNil sets the value for PublicAddress to be an explicit nil

### UnsetPublicAddress
`func (o *CloudNetworkInterfaceAttachment) UnsetPublicAddress()`

UnsetPublicAddress ensures that no value is present for PublicAddress, not even an explicit nil
### GetSecurityGroups

`func (o *CloudNetworkInterfaceAttachment) GetSecurityGroups() []string`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *CloudNetworkInterfaceAttachment) GetSecurityGroupsOk() (*[]string, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *CloudNetworkInterfaceAttachment) SetSecurityGroups(v []string)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *CloudNetworkInterfaceAttachment) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *CloudNetworkInterfaceAttachment) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *CloudNetworkInterfaceAttachment) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil
### GetSubNetworkId

`func (o *CloudNetworkInterfaceAttachment) GetSubNetworkId() string`

GetSubNetworkId returns the SubNetworkId field if non-nil, zero value otherwise.

### GetSubNetworkIdOk

`func (o *CloudNetworkInterfaceAttachment) GetSubNetworkIdOk() (*string, bool)`

GetSubNetworkIdOk returns a tuple with the SubNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubNetworkId

`func (o *CloudNetworkInterfaceAttachment) SetSubNetworkId(v string)`

SetSubNetworkId sets SubNetworkId field to given value.

### HasSubNetworkId

`func (o *CloudNetworkInterfaceAttachment) HasSubNetworkId() bool`

HasSubNetworkId returns a boolean if a field has been set.

### GetSubNetworkName

`func (o *CloudNetworkInterfaceAttachment) GetSubNetworkName() string`

GetSubNetworkName returns the SubNetworkName field if non-nil, zero value otherwise.

### GetSubNetworkNameOk

`func (o *CloudNetworkInterfaceAttachment) GetSubNetworkNameOk() (*string, bool)`

GetSubNetworkNameOk returns a tuple with the SubNetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubNetworkName

`func (o *CloudNetworkInterfaceAttachment) SetSubNetworkName(v string)`

SetSubNetworkName sets SubNetworkName field to given value.

### HasSubNetworkName

`func (o *CloudNetworkInterfaceAttachment) HasSubNetworkName() bool`

HasSubNetworkName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


