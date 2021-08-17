# CloudAwsSubnetAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cloud.AwsSubnet"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cloud.AwsSubnet"]
**AutoAssignPrivateIpV6** | Pointer to **bool** | If true, Ipv4 address is assigned automatically. | [optional] [readonly] 
**AutoAssignPublicIpV4** | Pointer to **bool** | If true, Ipv4 address is assigned automatically. | [optional] [readonly] 
**AvailabilityZoneName** | Pointer to **string** | The Availability Zone name of the subnet. | [optional] [readonly] 
**Ipv4Cidr** | Pointer to **string** | The IPv4 CIDR block assigned to the subnet.. | [optional] [readonly] 
**Ipv6Cidr** | Pointer to **string** | The IPv6 CIDR block assigned to the subnet.. | [optional] [readonly] 
**IsDefault** | Pointer to **bool** | If true, indicates that this is default subnet. | [optional] [readonly] 
**State** | Pointer to **string** | The state of the subnet (pending | available). | [optional] [readonly] 
**SubnetTags** | Pointer to [**[]CloudCloudTag**](CloudCloudTag.md) |  | [optional] 
**AwsVpc** | Pointer to [**CloudAwsVpcRelationship**](cloud.AwsVpc.Relationship.md) |  | [optional] 

## Methods

### NewCloudAwsSubnetAllOf

`func NewCloudAwsSubnetAllOf(classId string, objectType string, ) *CloudAwsSubnetAllOf`

NewCloudAwsSubnetAllOf instantiates a new CloudAwsSubnetAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAwsSubnetAllOfWithDefaults

`func NewCloudAwsSubnetAllOfWithDefaults() *CloudAwsSubnetAllOf`

NewCloudAwsSubnetAllOfWithDefaults instantiates a new CloudAwsSubnetAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudAwsSubnetAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudAwsSubnetAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudAwsSubnetAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudAwsSubnetAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudAwsSubnetAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudAwsSubnetAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAutoAssignPrivateIpV6

`func (o *CloudAwsSubnetAllOf) GetAutoAssignPrivateIpV6() bool`

GetAutoAssignPrivateIpV6 returns the AutoAssignPrivateIpV6 field if non-nil, zero value otherwise.

### GetAutoAssignPrivateIpV6Ok

`func (o *CloudAwsSubnetAllOf) GetAutoAssignPrivateIpV6Ok() (*bool, bool)`

GetAutoAssignPrivateIpV6Ok returns a tuple with the AutoAssignPrivateIpV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAssignPrivateIpV6

`func (o *CloudAwsSubnetAllOf) SetAutoAssignPrivateIpV6(v bool)`

SetAutoAssignPrivateIpV6 sets AutoAssignPrivateIpV6 field to given value.

### HasAutoAssignPrivateIpV6

`func (o *CloudAwsSubnetAllOf) HasAutoAssignPrivateIpV6() bool`

HasAutoAssignPrivateIpV6 returns a boolean if a field has been set.

### GetAutoAssignPublicIpV4

`func (o *CloudAwsSubnetAllOf) GetAutoAssignPublicIpV4() bool`

GetAutoAssignPublicIpV4 returns the AutoAssignPublicIpV4 field if non-nil, zero value otherwise.

### GetAutoAssignPublicIpV4Ok

`func (o *CloudAwsSubnetAllOf) GetAutoAssignPublicIpV4Ok() (*bool, bool)`

GetAutoAssignPublicIpV4Ok returns a tuple with the AutoAssignPublicIpV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAssignPublicIpV4

`func (o *CloudAwsSubnetAllOf) SetAutoAssignPublicIpV4(v bool)`

SetAutoAssignPublicIpV4 sets AutoAssignPublicIpV4 field to given value.

### HasAutoAssignPublicIpV4

`func (o *CloudAwsSubnetAllOf) HasAutoAssignPublicIpV4() bool`

HasAutoAssignPublicIpV4 returns a boolean if a field has been set.

### GetAvailabilityZoneName

`func (o *CloudAwsSubnetAllOf) GetAvailabilityZoneName() string`

GetAvailabilityZoneName returns the AvailabilityZoneName field if non-nil, zero value otherwise.

### GetAvailabilityZoneNameOk

`func (o *CloudAwsSubnetAllOf) GetAvailabilityZoneNameOk() (*string, bool)`

GetAvailabilityZoneNameOk returns a tuple with the AvailabilityZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneName

`func (o *CloudAwsSubnetAllOf) SetAvailabilityZoneName(v string)`

SetAvailabilityZoneName sets AvailabilityZoneName field to given value.

### HasAvailabilityZoneName

`func (o *CloudAwsSubnetAllOf) HasAvailabilityZoneName() bool`

HasAvailabilityZoneName returns a boolean if a field has been set.

### GetIpv4Cidr

`func (o *CloudAwsSubnetAllOf) GetIpv4Cidr() string`

GetIpv4Cidr returns the Ipv4Cidr field if non-nil, zero value otherwise.

### GetIpv4CidrOk

`func (o *CloudAwsSubnetAllOf) GetIpv4CidrOk() (*string, bool)`

GetIpv4CidrOk returns a tuple with the Ipv4Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Cidr

`func (o *CloudAwsSubnetAllOf) SetIpv4Cidr(v string)`

SetIpv4Cidr sets Ipv4Cidr field to given value.

### HasIpv4Cidr

`func (o *CloudAwsSubnetAllOf) HasIpv4Cidr() bool`

HasIpv4Cidr returns a boolean if a field has been set.

### GetIpv6Cidr

`func (o *CloudAwsSubnetAllOf) GetIpv6Cidr() string`

GetIpv6Cidr returns the Ipv6Cidr field if non-nil, zero value otherwise.

### GetIpv6CidrOk

`func (o *CloudAwsSubnetAllOf) GetIpv6CidrOk() (*string, bool)`

GetIpv6CidrOk returns a tuple with the Ipv6Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Cidr

`func (o *CloudAwsSubnetAllOf) SetIpv6Cidr(v string)`

SetIpv6Cidr sets Ipv6Cidr field to given value.

### HasIpv6Cidr

`func (o *CloudAwsSubnetAllOf) HasIpv6Cidr() bool`

HasIpv6Cidr returns a boolean if a field has been set.

### GetIsDefault

`func (o *CloudAwsSubnetAllOf) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *CloudAwsSubnetAllOf) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *CloudAwsSubnetAllOf) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *CloudAwsSubnetAllOf) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetState

`func (o *CloudAwsSubnetAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CloudAwsSubnetAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CloudAwsSubnetAllOf) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CloudAwsSubnetAllOf) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubnetTags

`func (o *CloudAwsSubnetAllOf) GetSubnetTags() []CloudCloudTag`

GetSubnetTags returns the SubnetTags field if non-nil, zero value otherwise.

### GetSubnetTagsOk

`func (o *CloudAwsSubnetAllOf) GetSubnetTagsOk() (*[]CloudCloudTag, bool)`

GetSubnetTagsOk returns a tuple with the SubnetTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetTags

`func (o *CloudAwsSubnetAllOf) SetSubnetTags(v []CloudCloudTag)`

SetSubnetTags sets SubnetTags field to given value.

### HasSubnetTags

`func (o *CloudAwsSubnetAllOf) HasSubnetTags() bool`

HasSubnetTags returns a boolean if a field has been set.

### SetSubnetTagsNil

`func (o *CloudAwsSubnetAllOf) SetSubnetTagsNil(b bool)`

 SetSubnetTagsNil sets the value for SubnetTags to be an explicit nil

### UnsetSubnetTags
`func (o *CloudAwsSubnetAllOf) UnsetSubnetTags()`

UnsetSubnetTags ensures that no value is present for SubnetTags, not even an explicit nil
### GetAwsVpc

`func (o *CloudAwsSubnetAllOf) GetAwsVpc() CloudAwsVpcRelationship`

GetAwsVpc returns the AwsVpc field if non-nil, zero value otherwise.

### GetAwsVpcOk

`func (o *CloudAwsSubnetAllOf) GetAwsVpcOk() (*CloudAwsVpcRelationship, bool)`

GetAwsVpcOk returns a tuple with the AwsVpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsVpc

`func (o *CloudAwsSubnetAllOf) SetAwsVpc(v CloudAwsVpcRelationship)`

SetAwsVpc sets AwsVpc field to given value.

### HasAwsVpc

`func (o *CloudAwsSubnetAllOf) HasAwsVpc() bool`

HasAwsVpc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


