# CloudAwsNetworkInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cloud.AwsNetworkInterface"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cloud.AwsNetworkInterface"]
**InstanceAttachment** | Pointer to [**NullableCloudNetworkInstanceAttachment**](cloud.NetworkInstanceAttachment.md) |  | [optional] 
**NicCreateTime** | Pointer to **time.Time** | Time when this network interface is created. | [optional] [readonly] 
**PrivateDnsName** | Pointer to **string** | The private DNS hostname name assigned to the network interface. | [optional] [readonly] 
**PrivateIpAddress** | Pointer to **[]string** |  | [optional] 
**PublicDnsName** | Pointer to **string** | The public DNS hostname name assigned to the network interface. | [optional] [readonly] 
**PublicIpAddress** | Pointer to **[]string** |  | [optional] 
**SecurityGroups** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to **string** | The status of the network interface. If the network interface is not attached to an instance, the status is available; if a network interface is attached to an instance the status is in-use. | [optional] [readonly] 
**AwsSubnet** | Pointer to [**CloudAwsSubnetRelationship**](cloud.AwsSubnet.Relationship.md) |  | [optional] 

## Methods

### NewCloudAwsNetworkInterface

`func NewCloudAwsNetworkInterface(classId string, objectType string, ) *CloudAwsNetworkInterface`

NewCloudAwsNetworkInterface instantiates a new CloudAwsNetworkInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAwsNetworkInterfaceWithDefaults

`func NewCloudAwsNetworkInterfaceWithDefaults() *CloudAwsNetworkInterface`

NewCloudAwsNetworkInterfaceWithDefaults instantiates a new CloudAwsNetworkInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudAwsNetworkInterface) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudAwsNetworkInterface) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudAwsNetworkInterface) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudAwsNetworkInterface) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudAwsNetworkInterface) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudAwsNetworkInterface) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetInstanceAttachment

`func (o *CloudAwsNetworkInterface) GetInstanceAttachment() CloudNetworkInstanceAttachment`

GetInstanceAttachment returns the InstanceAttachment field if non-nil, zero value otherwise.

### GetInstanceAttachmentOk

`func (o *CloudAwsNetworkInterface) GetInstanceAttachmentOk() (*CloudNetworkInstanceAttachment, bool)`

GetInstanceAttachmentOk returns a tuple with the InstanceAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceAttachment

`func (o *CloudAwsNetworkInterface) SetInstanceAttachment(v CloudNetworkInstanceAttachment)`

SetInstanceAttachment sets InstanceAttachment field to given value.

### HasInstanceAttachment

`func (o *CloudAwsNetworkInterface) HasInstanceAttachment() bool`

HasInstanceAttachment returns a boolean if a field has been set.

### SetInstanceAttachmentNil

`func (o *CloudAwsNetworkInterface) SetInstanceAttachmentNil(b bool)`

 SetInstanceAttachmentNil sets the value for InstanceAttachment to be an explicit nil

### UnsetInstanceAttachment
`func (o *CloudAwsNetworkInterface) UnsetInstanceAttachment()`

UnsetInstanceAttachment ensures that no value is present for InstanceAttachment, not even an explicit nil
### GetNicCreateTime

`func (o *CloudAwsNetworkInterface) GetNicCreateTime() time.Time`

GetNicCreateTime returns the NicCreateTime field if non-nil, zero value otherwise.

### GetNicCreateTimeOk

`func (o *CloudAwsNetworkInterface) GetNicCreateTimeOk() (*time.Time, bool)`

GetNicCreateTimeOk returns a tuple with the NicCreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicCreateTime

`func (o *CloudAwsNetworkInterface) SetNicCreateTime(v time.Time)`

SetNicCreateTime sets NicCreateTime field to given value.

### HasNicCreateTime

`func (o *CloudAwsNetworkInterface) HasNicCreateTime() bool`

HasNicCreateTime returns a boolean if a field has been set.

### GetPrivateDnsName

`func (o *CloudAwsNetworkInterface) GetPrivateDnsName() string`

GetPrivateDnsName returns the PrivateDnsName field if non-nil, zero value otherwise.

### GetPrivateDnsNameOk

`func (o *CloudAwsNetworkInterface) GetPrivateDnsNameOk() (*string, bool)`

GetPrivateDnsNameOk returns a tuple with the PrivateDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateDnsName

`func (o *CloudAwsNetworkInterface) SetPrivateDnsName(v string)`

SetPrivateDnsName sets PrivateDnsName field to given value.

### HasPrivateDnsName

`func (o *CloudAwsNetworkInterface) HasPrivateDnsName() bool`

HasPrivateDnsName returns a boolean if a field has been set.

### GetPrivateIpAddress

`func (o *CloudAwsNetworkInterface) GetPrivateIpAddress() []string`

GetPrivateIpAddress returns the PrivateIpAddress field if non-nil, zero value otherwise.

### GetPrivateIpAddressOk

`func (o *CloudAwsNetworkInterface) GetPrivateIpAddressOk() (*[]string, bool)`

GetPrivateIpAddressOk returns a tuple with the PrivateIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIpAddress

`func (o *CloudAwsNetworkInterface) SetPrivateIpAddress(v []string)`

SetPrivateIpAddress sets PrivateIpAddress field to given value.

### HasPrivateIpAddress

`func (o *CloudAwsNetworkInterface) HasPrivateIpAddress() bool`

HasPrivateIpAddress returns a boolean if a field has been set.

### SetPrivateIpAddressNil

`func (o *CloudAwsNetworkInterface) SetPrivateIpAddressNil(b bool)`

 SetPrivateIpAddressNil sets the value for PrivateIpAddress to be an explicit nil

### UnsetPrivateIpAddress
`func (o *CloudAwsNetworkInterface) UnsetPrivateIpAddress()`

UnsetPrivateIpAddress ensures that no value is present for PrivateIpAddress, not even an explicit nil
### GetPublicDnsName

`func (o *CloudAwsNetworkInterface) GetPublicDnsName() string`

GetPublicDnsName returns the PublicDnsName field if non-nil, zero value otherwise.

### GetPublicDnsNameOk

`func (o *CloudAwsNetworkInterface) GetPublicDnsNameOk() (*string, bool)`

GetPublicDnsNameOk returns a tuple with the PublicDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicDnsName

`func (o *CloudAwsNetworkInterface) SetPublicDnsName(v string)`

SetPublicDnsName sets PublicDnsName field to given value.

### HasPublicDnsName

`func (o *CloudAwsNetworkInterface) HasPublicDnsName() bool`

HasPublicDnsName returns a boolean if a field has been set.

### GetPublicIpAddress

`func (o *CloudAwsNetworkInterface) GetPublicIpAddress() []string`

GetPublicIpAddress returns the PublicIpAddress field if non-nil, zero value otherwise.

### GetPublicIpAddressOk

`func (o *CloudAwsNetworkInterface) GetPublicIpAddressOk() (*[]string, bool)`

GetPublicIpAddressOk returns a tuple with the PublicIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpAddress

`func (o *CloudAwsNetworkInterface) SetPublicIpAddress(v []string)`

SetPublicIpAddress sets PublicIpAddress field to given value.

### HasPublicIpAddress

`func (o *CloudAwsNetworkInterface) HasPublicIpAddress() bool`

HasPublicIpAddress returns a boolean if a field has been set.

### SetPublicIpAddressNil

`func (o *CloudAwsNetworkInterface) SetPublicIpAddressNil(b bool)`

 SetPublicIpAddressNil sets the value for PublicIpAddress to be an explicit nil

### UnsetPublicIpAddress
`func (o *CloudAwsNetworkInterface) UnsetPublicIpAddress()`

UnsetPublicIpAddress ensures that no value is present for PublicIpAddress, not even an explicit nil
### GetSecurityGroups

`func (o *CloudAwsNetworkInterface) GetSecurityGroups() []string`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *CloudAwsNetworkInterface) GetSecurityGroupsOk() (*[]string, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *CloudAwsNetworkInterface) SetSecurityGroups(v []string)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *CloudAwsNetworkInterface) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *CloudAwsNetworkInterface) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *CloudAwsNetworkInterface) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil
### GetStatus

`func (o *CloudAwsNetworkInterface) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudAwsNetworkInterface) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudAwsNetworkInterface) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudAwsNetworkInterface) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAwsSubnet

`func (o *CloudAwsNetworkInterface) GetAwsSubnet() CloudAwsSubnetRelationship`

GetAwsSubnet returns the AwsSubnet field if non-nil, zero value otherwise.

### GetAwsSubnetOk

`func (o *CloudAwsNetworkInterface) GetAwsSubnetOk() (*CloudAwsSubnetRelationship, bool)`

GetAwsSubnetOk returns a tuple with the AwsSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSubnet

`func (o *CloudAwsNetworkInterface) SetAwsSubnet(v CloudAwsSubnetRelationship)`

SetAwsSubnet sets AwsSubnet field to given value.

### HasAwsSubnet

`func (o *CloudAwsNetworkInterface) HasAwsSubnet() bool`

HasAwsSubnet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


