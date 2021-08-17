# CloudAwsSecurityGroupAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cloud.AwsSecurityGroup"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cloud.AwsSecurityGroup"]
**EgressRules** | Pointer to [**[]CloudSecurityGroupRule**](CloudSecurityGroupRule.md) |  | [optional] 
**IngressRules** | Pointer to [**[]CloudSecurityGroupRule**](CloudSecurityGroupRule.md) |  | [optional] 
**SecurityGroupTags** | Pointer to [**[]CloudCloudTag**](CloudCloudTag.md) |  | [optional] 
**Location** | Pointer to [**CloudAwsVpcRelationship**](cloud.AwsVpc.Relationship.md) |  | [optional] 

## Methods

### NewCloudAwsSecurityGroupAllOf

`func NewCloudAwsSecurityGroupAllOf(classId string, objectType string, ) *CloudAwsSecurityGroupAllOf`

NewCloudAwsSecurityGroupAllOf instantiates a new CloudAwsSecurityGroupAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAwsSecurityGroupAllOfWithDefaults

`func NewCloudAwsSecurityGroupAllOfWithDefaults() *CloudAwsSecurityGroupAllOf`

NewCloudAwsSecurityGroupAllOfWithDefaults instantiates a new CloudAwsSecurityGroupAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudAwsSecurityGroupAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudAwsSecurityGroupAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudAwsSecurityGroupAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudAwsSecurityGroupAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudAwsSecurityGroupAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudAwsSecurityGroupAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEgressRules

`func (o *CloudAwsSecurityGroupAllOf) GetEgressRules() []CloudSecurityGroupRule`

GetEgressRules returns the EgressRules field if non-nil, zero value otherwise.

### GetEgressRulesOk

`func (o *CloudAwsSecurityGroupAllOf) GetEgressRulesOk() (*[]CloudSecurityGroupRule, bool)`

GetEgressRulesOk returns a tuple with the EgressRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressRules

`func (o *CloudAwsSecurityGroupAllOf) SetEgressRules(v []CloudSecurityGroupRule)`

SetEgressRules sets EgressRules field to given value.

### HasEgressRules

`func (o *CloudAwsSecurityGroupAllOf) HasEgressRules() bool`

HasEgressRules returns a boolean if a field has been set.

### SetEgressRulesNil

`func (o *CloudAwsSecurityGroupAllOf) SetEgressRulesNil(b bool)`

 SetEgressRulesNil sets the value for EgressRules to be an explicit nil

### UnsetEgressRules
`func (o *CloudAwsSecurityGroupAllOf) UnsetEgressRules()`

UnsetEgressRules ensures that no value is present for EgressRules, not even an explicit nil
### GetIngressRules

`func (o *CloudAwsSecurityGroupAllOf) GetIngressRules() []CloudSecurityGroupRule`

GetIngressRules returns the IngressRules field if non-nil, zero value otherwise.

### GetIngressRulesOk

`func (o *CloudAwsSecurityGroupAllOf) GetIngressRulesOk() (*[]CloudSecurityGroupRule, bool)`

GetIngressRulesOk returns a tuple with the IngressRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressRules

`func (o *CloudAwsSecurityGroupAllOf) SetIngressRules(v []CloudSecurityGroupRule)`

SetIngressRules sets IngressRules field to given value.

### HasIngressRules

`func (o *CloudAwsSecurityGroupAllOf) HasIngressRules() bool`

HasIngressRules returns a boolean if a field has been set.

### SetIngressRulesNil

`func (o *CloudAwsSecurityGroupAllOf) SetIngressRulesNil(b bool)`

 SetIngressRulesNil sets the value for IngressRules to be an explicit nil

### UnsetIngressRules
`func (o *CloudAwsSecurityGroupAllOf) UnsetIngressRules()`

UnsetIngressRules ensures that no value is present for IngressRules, not even an explicit nil
### GetSecurityGroupTags

`func (o *CloudAwsSecurityGroupAllOf) GetSecurityGroupTags() []CloudCloudTag`

GetSecurityGroupTags returns the SecurityGroupTags field if non-nil, zero value otherwise.

### GetSecurityGroupTagsOk

`func (o *CloudAwsSecurityGroupAllOf) GetSecurityGroupTagsOk() (*[]CloudCloudTag, bool)`

GetSecurityGroupTagsOk returns a tuple with the SecurityGroupTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupTags

`func (o *CloudAwsSecurityGroupAllOf) SetSecurityGroupTags(v []CloudCloudTag)`

SetSecurityGroupTags sets SecurityGroupTags field to given value.

### HasSecurityGroupTags

`func (o *CloudAwsSecurityGroupAllOf) HasSecurityGroupTags() bool`

HasSecurityGroupTags returns a boolean if a field has been set.

### SetSecurityGroupTagsNil

`func (o *CloudAwsSecurityGroupAllOf) SetSecurityGroupTagsNil(b bool)`

 SetSecurityGroupTagsNil sets the value for SecurityGroupTags to be an explicit nil

### UnsetSecurityGroupTags
`func (o *CloudAwsSecurityGroupAllOf) UnsetSecurityGroupTags()`

UnsetSecurityGroupTags ensures that no value is present for SecurityGroupTags, not even an explicit nil
### GetLocation

`func (o *CloudAwsSecurityGroupAllOf) GetLocation() CloudAwsVpcRelationship`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CloudAwsSecurityGroupAllOf) GetLocationOk() (*CloudAwsVpcRelationship, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CloudAwsSecurityGroupAllOf) SetLocation(v CloudAwsVpcRelationship)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *CloudAwsSecurityGroupAllOf) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


