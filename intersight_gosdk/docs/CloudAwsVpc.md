# CloudAwsVpc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cloud.AwsVpc"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cloud.AwsVpc"]
**DnsHostName** | Pointer to **bool** | If true, instances in the vpc get public DNS hostnames. | [optional] [readonly] 
**DnsResolution** | Pointer to **bool** | Indicates whether the Dns resolution is supported. | [optional] [readonly] 
**Ipv4Cidr** | Pointer to **[]string** |  | [optional] 
**Ipv6Cidr** | Pointer to **[]string** |  | [optional] 
**IsDefault** | Pointer to **bool** | If true, indicates that this is default VPC. | [optional] [readonly] 
**State** | Pointer to **string** | The state of the VPC (pending | available). | [optional] [readonly] 
**Tenancy** | Pointer to **string** | The allowed tenancy of instances launched into the VPC. | [optional] [readonly] 
**VpcTags** | Pointer to [**[]CloudCloudTag**](CloudCloudTag.md) |  | [optional] 
**AwsBillingUnit** | Pointer to [**CloudAwsBillingUnitRelationship**](cloud.AwsBillingUnit.Relationship.md) |  | [optional] 

## Methods

### NewCloudAwsVpc

`func NewCloudAwsVpc(classId string, objectType string, ) *CloudAwsVpc`

NewCloudAwsVpc instantiates a new CloudAwsVpc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAwsVpcWithDefaults

`func NewCloudAwsVpcWithDefaults() *CloudAwsVpc`

NewCloudAwsVpcWithDefaults instantiates a new CloudAwsVpc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudAwsVpc) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudAwsVpc) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudAwsVpc) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudAwsVpc) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudAwsVpc) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudAwsVpc) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDnsHostName

`func (o *CloudAwsVpc) GetDnsHostName() bool`

GetDnsHostName returns the DnsHostName field if non-nil, zero value otherwise.

### GetDnsHostNameOk

`func (o *CloudAwsVpc) GetDnsHostNameOk() (*bool, bool)`

GetDnsHostNameOk returns a tuple with the DnsHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsHostName

`func (o *CloudAwsVpc) SetDnsHostName(v bool)`

SetDnsHostName sets DnsHostName field to given value.

### HasDnsHostName

`func (o *CloudAwsVpc) HasDnsHostName() bool`

HasDnsHostName returns a boolean if a field has been set.

### GetDnsResolution

`func (o *CloudAwsVpc) GetDnsResolution() bool`

GetDnsResolution returns the DnsResolution field if non-nil, zero value otherwise.

### GetDnsResolutionOk

`func (o *CloudAwsVpc) GetDnsResolutionOk() (*bool, bool)`

GetDnsResolutionOk returns a tuple with the DnsResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsResolution

`func (o *CloudAwsVpc) SetDnsResolution(v bool)`

SetDnsResolution sets DnsResolution field to given value.

### HasDnsResolution

`func (o *CloudAwsVpc) HasDnsResolution() bool`

HasDnsResolution returns a boolean if a field has been set.

### GetIpv4Cidr

`func (o *CloudAwsVpc) GetIpv4Cidr() []string`

GetIpv4Cidr returns the Ipv4Cidr field if non-nil, zero value otherwise.

### GetIpv4CidrOk

`func (o *CloudAwsVpc) GetIpv4CidrOk() (*[]string, bool)`

GetIpv4CidrOk returns a tuple with the Ipv4Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Cidr

`func (o *CloudAwsVpc) SetIpv4Cidr(v []string)`

SetIpv4Cidr sets Ipv4Cidr field to given value.

### HasIpv4Cidr

`func (o *CloudAwsVpc) HasIpv4Cidr() bool`

HasIpv4Cidr returns a boolean if a field has been set.

### SetIpv4CidrNil

`func (o *CloudAwsVpc) SetIpv4CidrNil(b bool)`

 SetIpv4CidrNil sets the value for Ipv4Cidr to be an explicit nil

### UnsetIpv4Cidr
`func (o *CloudAwsVpc) UnsetIpv4Cidr()`

UnsetIpv4Cidr ensures that no value is present for Ipv4Cidr, not even an explicit nil
### GetIpv6Cidr

`func (o *CloudAwsVpc) GetIpv6Cidr() []string`

GetIpv6Cidr returns the Ipv6Cidr field if non-nil, zero value otherwise.

### GetIpv6CidrOk

`func (o *CloudAwsVpc) GetIpv6CidrOk() (*[]string, bool)`

GetIpv6CidrOk returns a tuple with the Ipv6Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Cidr

`func (o *CloudAwsVpc) SetIpv6Cidr(v []string)`

SetIpv6Cidr sets Ipv6Cidr field to given value.

### HasIpv6Cidr

`func (o *CloudAwsVpc) HasIpv6Cidr() bool`

HasIpv6Cidr returns a boolean if a field has been set.

### SetIpv6CidrNil

`func (o *CloudAwsVpc) SetIpv6CidrNil(b bool)`

 SetIpv6CidrNil sets the value for Ipv6Cidr to be an explicit nil

### UnsetIpv6Cidr
`func (o *CloudAwsVpc) UnsetIpv6Cidr()`

UnsetIpv6Cidr ensures that no value is present for Ipv6Cidr, not even an explicit nil
### GetIsDefault

`func (o *CloudAwsVpc) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *CloudAwsVpc) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *CloudAwsVpc) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *CloudAwsVpc) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetState

`func (o *CloudAwsVpc) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CloudAwsVpc) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CloudAwsVpc) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CloudAwsVpc) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTenancy

`func (o *CloudAwsVpc) GetTenancy() string`

GetTenancy returns the Tenancy field if non-nil, zero value otherwise.

### GetTenancyOk

`func (o *CloudAwsVpc) GetTenancyOk() (*string, bool)`

GetTenancyOk returns a tuple with the Tenancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenancy

`func (o *CloudAwsVpc) SetTenancy(v string)`

SetTenancy sets Tenancy field to given value.

### HasTenancy

`func (o *CloudAwsVpc) HasTenancy() bool`

HasTenancy returns a boolean if a field has been set.

### GetVpcTags

`func (o *CloudAwsVpc) GetVpcTags() []CloudCloudTag`

GetVpcTags returns the VpcTags field if non-nil, zero value otherwise.

### GetVpcTagsOk

`func (o *CloudAwsVpc) GetVpcTagsOk() (*[]CloudCloudTag, bool)`

GetVpcTagsOk returns a tuple with the VpcTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcTags

`func (o *CloudAwsVpc) SetVpcTags(v []CloudCloudTag)`

SetVpcTags sets VpcTags field to given value.

### HasVpcTags

`func (o *CloudAwsVpc) HasVpcTags() bool`

HasVpcTags returns a boolean if a field has been set.

### SetVpcTagsNil

`func (o *CloudAwsVpc) SetVpcTagsNil(b bool)`

 SetVpcTagsNil sets the value for VpcTags to be an explicit nil

### UnsetVpcTags
`func (o *CloudAwsVpc) UnsetVpcTags()`

UnsetVpcTags ensures that no value is present for VpcTags, not even an explicit nil
### GetAwsBillingUnit

`func (o *CloudAwsVpc) GetAwsBillingUnit() CloudAwsBillingUnitRelationship`

GetAwsBillingUnit returns the AwsBillingUnit field if non-nil, zero value otherwise.

### GetAwsBillingUnitOk

`func (o *CloudAwsVpc) GetAwsBillingUnitOk() (*CloudAwsBillingUnitRelationship, bool)`

GetAwsBillingUnitOk returns a tuple with the AwsBillingUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsBillingUnit

`func (o *CloudAwsVpc) SetAwsBillingUnit(v CloudAwsBillingUnitRelationship)`

SetAwsBillingUnit sets AwsBillingUnit field to given value.

### HasAwsBillingUnit

`func (o *CloudAwsVpc) HasAwsBillingUnit() bool`

HasAwsBillingUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


