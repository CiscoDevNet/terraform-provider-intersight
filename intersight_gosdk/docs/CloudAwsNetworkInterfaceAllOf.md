# CloudAwsNetworkInterfaceAllOf

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

### NewCloudAwsNetworkInterfaceAllOf

`func NewCloudAwsNetworkInterfaceAllOf(classId string, objectType string, ) *CloudAwsNetworkInterfaceAllOf`

NewCloudAwsNetworkInterfaceAllOf instantiates a new CloudAwsNetworkInterfaceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAwsNetworkInterfaceAllOfWithDefaults

`func NewCloudAwsNetworkInterfaceAllOfWithDefaults() *CloudAwsNetworkInterfaceAllOf`

NewCloudAwsNetworkInterfaceAllOfWithDefaults instantiates a new CloudAwsNetworkInterfaceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudAwsNetworkInterfaceAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudAwsNetworkInterfaceAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudAwsNetworkInterfaceAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudAwsNetworkInterfaceAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudAwsNetworkInterfaceAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudAwsNetworkInterfaceAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetInstanceAttachment

`func (o *CloudAwsNetworkInterfaceAllOf) GetInstanceAttachment() CloudNetworkInstanceAttachment`

GetInstanceAttachment returns the InstanceAttachment field if non-nil, zero value otherwise.

### GetInstanceAttachmentOk

`func (o *CloudAwsNetworkInterfaceAllOf) GetInstanceAttachmentOk() (*CloudNetworkInstanceAttachment, bool)`

GetInstanceAttachmentOk returns a tuple with the InstanceAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceAttachment

`func (o *CloudAwsNetworkInterfaceAllOf) SetInstanceAttachment(v CloudNetworkInstanceAttachment)`

SetInstanceAttachment sets InstanceAttachment field to given value.

### HasInstanceAttachment

`func (o *CloudAwsNetworkInterfaceAllOf) HasInstanceAttachment() bool`

HasInstanceAttachment returns a boolean if a field has been set.

### SetInstanceAttachmentNil

`func (o *CloudAwsNetworkInterfaceAllOf) SetInstanceAttachmentNil(b bool)`

 SetInstanceAttachmentNil sets the value for InstanceAttachment to be an explicit nil

### UnsetInstanceAttachment
`func (o *CloudAwsNetworkInterfaceAllOf) UnsetInstanceAttachment()`

UnsetInstanceAttachment ensures that no value is present for InstanceAttachment, not even an explicit nil
### GetNicCreateTime

`func (o *CloudAwsNetworkInterfaceAllOf) GetNicCreateTime() time.Time`

GetNicCreateTime returns the NicCreateTime field if non-nil, zero value otherwise.

### GetNicCreateTimeOk

`func (o *CloudAwsNetworkInterfaceAllOf) GetNicCreateTimeOk() (*time.Time, bool)`

GetNicCreateTimeOk returns a tuple with the NicCreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNicCreateTime

`func (o *CloudAwsNetworkInterfaceAllOf) SetNicCreateTime(v time.Time)`

SetNicCreateTime sets NicCreateTime field to given value.

### HasNicCreateTime

`func (o *CloudAwsNetworkInterfaceAllOf) HasNicCreateTime() bool`

HasNicCreateTime returns a boolean if a field has been set.

### GetPrivateDnsName

`func (o *CloudAwsNetworkInterfaceAllOf) GetPrivateDnsName() string`

GetPrivateDnsName returns the PrivateDnsName field if non-nil, zero value otherwise.

### GetPrivateDnsNameOk

`func (o *CloudAwsNetworkInterfaceAllOf) GetPrivateDnsNameOk() (*string, bool)`

GetPrivateDnsNameOk returns a tuple with the PrivateDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateDnsName

`func (o *CloudAwsNetworkInterfaceAllOf) SetPrivateDnsName(v string)`

SetPrivateDnsName sets PrivateDnsName field to given value.

### HasPrivateDnsName

`func (o *CloudAwsNetworkInterfaceAllOf) HasPrivateDnsName() bool`

HasPrivateDnsName returns a boolean if a field has been set.

### GetPrivateIpAddress

`func (o *CloudAwsNetworkInterfaceAllOf) GetPrivateIpAddress() []string`

GetPrivateIpAddress returns the PrivateIpAddress field if non-nil, zero value otherwise.

### GetPrivateIpAddressOk

`func (o *CloudAwsNetworkInterfaceAllOf) GetPrivateIpAddressOk() (*[]string, bool)`

GetPrivateIpAddressOk returns a tuple with the PrivateIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIpAddress

`func (o *CloudAwsNetworkInterfaceAllOf) SetPrivateIpAddress(v []string)`

SetPrivateIpAddress sets PrivateIpAddress field to given value.

### HasPrivateIpAddress

`func (o *CloudAwsNetworkInterfaceAllOf) HasPrivateIpAddress() bool`

HasPrivateIpAddress returns a boolean if a field has been set.

### SetPrivateIpAddressNil

`func (o *CloudAwsNetworkInterfaceAllOf) SetPrivateIpAddressNil(b bool)`

 SetPrivateIpAddressNil sets the value for PrivateIpAddress to be an explicit nil

### UnsetPrivateIpAddress
`func (o *CloudAwsNetworkInterfaceAllOf) UnsetPrivateIpAddress()`

UnsetPrivateIpAddress ensures that no value is present for PrivateIpAddress, not even an explicit nil
### GetPublicDnsName

`func (o *CloudAwsNetworkInterfaceAllOf) GetPublicDnsName() string`

GetPublicDnsName returns the PublicDnsName field if non-nil, zero value otherwise.

### GetPublicDnsNameOk

`func (o *CloudAwsNetworkInterfaceAllOf) GetPublicDnsNameOk() (*string, bool)`

GetPublicDnsNameOk returns a tuple with the PublicDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicDnsName

`func (o *CloudAwsNetworkInterfaceAllOf) SetPublicDnsName(v string)`

SetPublicDnsName sets PublicDnsName field to given value.

### HasPublicDnsName

`func (o *CloudAwsNetworkInterfaceAllOf) HasPublicDnsName() bool`

HasPublicDnsName returns a boolean if a field has been set.

### GetPublicIpAddress

`func (o *CloudAwsNetworkInterfaceAllOf) GetPublicIpAddress() []string`

GetPublicIpAddress returns the PublicIpAddress field if non-nil, zero value otherwise.

### GetPublicIpAddressOk

`func (o *CloudAwsNetworkInterfaceAllOf) GetPublicIpAddressOk() (*[]string, bool)`

GetPublicIpAddressOk returns a tuple with the PublicIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpAddress

`func (o *CloudAwsNetworkInterfaceAllOf) SetPublicIpAddress(v []string)`

SetPublicIpAddress sets PublicIpAddress field to given value.

### HasPublicIpAddress

`func (o *CloudAwsNetworkInterfaceAllOf) HasPublicIpAddress() bool`

HasPublicIpAddress returns a boolean if a field has been set.

### SetPublicIpAddressNil

`func (o *CloudAwsNetworkInterfaceAllOf) SetPublicIpAddressNil(b bool)`

 SetPublicIpAddressNil sets the value for PublicIpAddress to be an explicit nil

### UnsetPublicIpAddress
`func (o *CloudAwsNetworkInterfaceAllOf) UnsetPublicIpAddress()`

UnsetPublicIpAddress ensures that no value is present for PublicIpAddress, not even an explicit nil
### GetSecurityGroups

`func (o *CloudAwsNetworkInterfaceAllOf) GetSecurityGroups() []string`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *CloudAwsNetworkInterfaceAllOf) GetSecurityGroupsOk() (*[]string, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *CloudAwsNetworkInterfaceAllOf) SetSecurityGroups(v []string)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *CloudAwsNetworkInterfaceAllOf) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *CloudAwsNetworkInterfaceAllOf) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *CloudAwsNetworkInterfaceAllOf) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil
### GetStatus

`func (o *CloudAwsNetworkInterfaceAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudAwsNetworkInterfaceAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudAwsNetworkInterfaceAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudAwsNetworkInterfaceAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAwsSubnet

`func (o *CloudAwsNetworkInterfaceAllOf) GetAwsSubnet() CloudAwsSubnetRelationship`

GetAwsSubnet returns the AwsSubnet field if non-nil, zero value otherwise.

### GetAwsSubnetOk

`func (o *CloudAwsNetworkInterfaceAllOf) GetAwsSubnetOk() (*CloudAwsSubnetRelationship, bool)`

GetAwsSubnetOk returns a tuple with the AwsSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSubnet

`func (o *CloudAwsNetworkInterfaceAllOf) SetAwsSubnet(v CloudAwsSubnetRelationship)`

SetAwsSubnet sets AwsSubnet field to given value.

### HasAwsSubnet

`func (o *CloudAwsNetworkInterfaceAllOf) HasAwsSubnet() bool`

HasAwsSubnet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


