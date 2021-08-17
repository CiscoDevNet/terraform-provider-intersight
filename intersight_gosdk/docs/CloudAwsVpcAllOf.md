# CloudAwsVpcAllOf

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

### NewCloudAwsVpcAllOf

`func NewCloudAwsVpcAllOf(classId string, objectType string, ) *CloudAwsVpcAllOf`

NewCloudAwsVpcAllOf instantiates a new CloudAwsVpcAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAwsVpcAllOfWithDefaults

`func NewCloudAwsVpcAllOfWithDefaults() *CloudAwsVpcAllOf`

NewCloudAwsVpcAllOfWithDefaults instantiates a new CloudAwsVpcAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudAwsVpcAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudAwsVpcAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudAwsVpcAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudAwsVpcAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudAwsVpcAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudAwsVpcAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDnsHostName

`func (o *CloudAwsVpcAllOf) GetDnsHostName() bool`

GetDnsHostName returns the DnsHostName field if non-nil, zero value otherwise.

### GetDnsHostNameOk

`func (o *CloudAwsVpcAllOf) GetDnsHostNameOk() (*bool, bool)`

GetDnsHostNameOk returns a tuple with the DnsHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsHostName

`func (o *CloudAwsVpcAllOf) SetDnsHostName(v bool)`

SetDnsHostName sets DnsHostName field to given value.

### HasDnsHostName

`func (o *CloudAwsVpcAllOf) HasDnsHostName() bool`

HasDnsHostName returns a boolean if a field has been set.

### GetDnsResolution

`func (o *CloudAwsVpcAllOf) GetDnsResolution() bool`

GetDnsResolution returns the DnsResolution field if non-nil, zero value otherwise.

### GetDnsResolutionOk

`func (o *CloudAwsVpcAllOf) GetDnsResolutionOk() (*bool, bool)`

GetDnsResolutionOk returns a tuple with the DnsResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsResolution

`func (o *CloudAwsVpcAllOf) SetDnsResolution(v bool)`

SetDnsResolution sets DnsResolution field to given value.

### HasDnsResolution

`func (o *CloudAwsVpcAllOf) HasDnsResolution() bool`

HasDnsResolution returns a boolean if a field has been set.

### GetIpv4Cidr

`func (o *CloudAwsVpcAllOf) GetIpv4Cidr() []string`

GetIpv4Cidr returns the Ipv4Cidr field if non-nil, zero value otherwise.

### GetIpv4CidrOk

`func (o *CloudAwsVpcAllOf) GetIpv4CidrOk() (*[]string, bool)`

GetIpv4CidrOk returns a tuple with the Ipv4Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Cidr

`func (o *CloudAwsVpcAllOf) SetIpv4Cidr(v []string)`

SetIpv4Cidr sets Ipv4Cidr field to given value.

### HasIpv4Cidr

`func (o *CloudAwsVpcAllOf) HasIpv4Cidr() bool`

HasIpv4Cidr returns a boolean if a field has been set.

### SetIpv4CidrNil

`func (o *CloudAwsVpcAllOf) SetIpv4CidrNil(b bool)`

 SetIpv4CidrNil sets the value for Ipv4Cidr to be an explicit nil

### UnsetIpv4Cidr
`func (o *CloudAwsVpcAllOf) UnsetIpv4Cidr()`

UnsetIpv4Cidr ensures that no value is present for Ipv4Cidr, not even an explicit nil
### GetIpv6Cidr

`func (o *CloudAwsVpcAllOf) GetIpv6Cidr() []string`

GetIpv6Cidr returns the Ipv6Cidr field if non-nil, zero value otherwise.

### GetIpv6CidrOk

`func (o *CloudAwsVpcAllOf) GetIpv6CidrOk() (*[]string, bool)`

GetIpv6CidrOk returns a tuple with the Ipv6Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Cidr

`func (o *CloudAwsVpcAllOf) SetIpv6Cidr(v []string)`

SetIpv6Cidr sets Ipv6Cidr field to given value.

### HasIpv6Cidr

`func (o *CloudAwsVpcAllOf) HasIpv6Cidr() bool`

HasIpv6Cidr returns a boolean if a field has been set.

### SetIpv6CidrNil

`func (o *CloudAwsVpcAllOf) SetIpv6CidrNil(b bool)`

 SetIpv6CidrNil sets the value for Ipv6Cidr to be an explicit nil

### UnsetIpv6Cidr
`func (o *CloudAwsVpcAllOf) UnsetIpv6Cidr()`

UnsetIpv6Cidr ensures that no value is present for Ipv6Cidr, not even an explicit nil
### GetIsDefault

`func (o *CloudAwsVpcAllOf) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *CloudAwsVpcAllOf) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *CloudAwsVpcAllOf) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *CloudAwsVpcAllOf) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetState

`func (o *CloudAwsVpcAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CloudAwsVpcAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CloudAwsVpcAllOf) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CloudAwsVpcAllOf) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTenancy

`func (o *CloudAwsVpcAllOf) GetTenancy() string`

GetTenancy returns the Tenancy field if non-nil, zero value otherwise.

### GetTenancyOk

`func (o *CloudAwsVpcAllOf) GetTenancyOk() (*string, bool)`

GetTenancyOk returns a tuple with the Tenancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenancy

`func (o *CloudAwsVpcAllOf) SetTenancy(v string)`

SetTenancy sets Tenancy field to given value.

### HasTenancy

`func (o *CloudAwsVpcAllOf) HasTenancy() bool`

HasTenancy returns a boolean if a field has been set.

### GetVpcTags

`func (o *CloudAwsVpcAllOf) GetVpcTags() []CloudCloudTag`

GetVpcTags returns the VpcTags field if non-nil, zero value otherwise.

### GetVpcTagsOk

`func (o *CloudAwsVpcAllOf) GetVpcTagsOk() (*[]CloudCloudTag, bool)`

GetVpcTagsOk returns a tuple with the VpcTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcTags

`func (o *CloudAwsVpcAllOf) SetVpcTags(v []CloudCloudTag)`

SetVpcTags sets VpcTags field to given value.

### HasVpcTags

`func (o *CloudAwsVpcAllOf) HasVpcTags() bool`

HasVpcTags returns a boolean if a field has been set.

### SetVpcTagsNil

`func (o *CloudAwsVpcAllOf) SetVpcTagsNil(b bool)`

 SetVpcTagsNil sets the value for VpcTags to be an explicit nil

### UnsetVpcTags
`func (o *CloudAwsVpcAllOf) UnsetVpcTags()`

UnsetVpcTags ensures that no value is present for VpcTags, not even an explicit nil
### GetAwsBillingUnit

`func (o *CloudAwsVpcAllOf) GetAwsBillingUnit() CloudAwsBillingUnitRelationship`

GetAwsBillingUnit returns the AwsBillingUnit field if non-nil, zero value otherwise.

### GetAwsBillingUnitOk

`func (o *CloudAwsVpcAllOf) GetAwsBillingUnitOk() (*CloudAwsBillingUnitRelationship, bool)`

GetAwsBillingUnitOk returns a tuple with the AwsBillingUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsBillingUnit

`func (o *CloudAwsVpcAllOf) SetAwsBillingUnit(v CloudAwsBillingUnitRelationship)`

SetAwsBillingUnit sets AwsBillingUnit field to given value.

### HasAwsBillingUnit

`func (o *CloudAwsVpcAllOf) HasAwsBillingUnit() bool`

HasAwsBillingUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


