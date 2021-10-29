# CloudAwsVirtualMachineAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cloud.AwsVirtualMachine"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cloud.AwsVirtualMachine"]
**AwsBillingUnit** | Pointer to [**CloudAwsBillingUnitRelationship**](CloudAwsBillingUnitRelationship.md) |  | [optional] 
**KeyPair** | Pointer to [**CloudAwsKeyPairRelationship**](CloudAwsKeyPairRelationship.md) |  | [optional] 
**Location** | Pointer to [**CloudAwsVpcRelationship**](CloudAwsVpcRelationship.md) |  | [optional] 
**SecurityGroups** | Pointer to [**[]CloudAwsSecurityGroupRelationship**](CloudAwsSecurityGroupRelationship.md) | An array of relationships to cloudAwsSecurityGroup resources. | [optional] [readonly] 

## Methods

### NewCloudAwsVirtualMachineAllOf

`func NewCloudAwsVirtualMachineAllOf(classId string, objectType string, ) *CloudAwsVirtualMachineAllOf`

NewCloudAwsVirtualMachineAllOf instantiates a new CloudAwsVirtualMachineAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAwsVirtualMachineAllOfWithDefaults

`func NewCloudAwsVirtualMachineAllOfWithDefaults() *CloudAwsVirtualMachineAllOf`

NewCloudAwsVirtualMachineAllOfWithDefaults instantiates a new CloudAwsVirtualMachineAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudAwsVirtualMachineAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudAwsVirtualMachineAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudAwsVirtualMachineAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudAwsVirtualMachineAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudAwsVirtualMachineAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudAwsVirtualMachineAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAwsBillingUnit

`func (o *CloudAwsVirtualMachineAllOf) GetAwsBillingUnit() CloudAwsBillingUnitRelationship`

GetAwsBillingUnit returns the AwsBillingUnit field if non-nil, zero value otherwise.

### GetAwsBillingUnitOk

`func (o *CloudAwsVirtualMachineAllOf) GetAwsBillingUnitOk() (*CloudAwsBillingUnitRelationship, bool)`

GetAwsBillingUnitOk returns a tuple with the AwsBillingUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsBillingUnit

`func (o *CloudAwsVirtualMachineAllOf) SetAwsBillingUnit(v CloudAwsBillingUnitRelationship)`

SetAwsBillingUnit sets AwsBillingUnit field to given value.

### HasAwsBillingUnit

`func (o *CloudAwsVirtualMachineAllOf) HasAwsBillingUnit() bool`

HasAwsBillingUnit returns a boolean if a field has been set.

### GetKeyPair

`func (o *CloudAwsVirtualMachineAllOf) GetKeyPair() CloudAwsKeyPairRelationship`

GetKeyPair returns the KeyPair field if non-nil, zero value otherwise.

### GetKeyPairOk

`func (o *CloudAwsVirtualMachineAllOf) GetKeyPairOk() (*CloudAwsKeyPairRelationship, bool)`

GetKeyPairOk returns a tuple with the KeyPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPair

`func (o *CloudAwsVirtualMachineAllOf) SetKeyPair(v CloudAwsKeyPairRelationship)`

SetKeyPair sets KeyPair field to given value.

### HasKeyPair

`func (o *CloudAwsVirtualMachineAllOf) HasKeyPair() bool`

HasKeyPair returns a boolean if a field has been set.

### GetLocation

`func (o *CloudAwsVirtualMachineAllOf) GetLocation() CloudAwsVpcRelationship`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CloudAwsVirtualMachineAllOf) GetLocationOk() (*CloudAwsVpcRelationship, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CloudAwsVirtualMachineAllOf) SetLocation(v CloudAwsVpcRelationship)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *CloudAwsVirtualMachineAllOf) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *CloudAwsVirtualMachineAllOf) GetSecurityGroups() []CloudAwsSecurityGroupRelationship`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *CloudAwsVirtualMachineAllOf) GetSecurityGroupsOk() (*[]CloudAwsSecurityGroupRelationship, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *CloudAwsVirtualMachineAllOf) SetSecurityGroups(v []CloudAwsSecurityGroupRelationship)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *CloudAwsVirtualMachineAllOf) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *CloudAwsVirtualMachineAllOf) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *CloudAwsVirtualMachineAllOf) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


