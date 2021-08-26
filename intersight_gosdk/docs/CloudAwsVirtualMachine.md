# CloudAwsVirtualMachine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cloud.AwsVirtualMachine"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cloud.AwsVirtualMachine"]
**KeyPair** | Pointer to [**CloudAwsKeyPairRelationship**](CloudAwsKeyPairRelationship.md) |  | [optional] 
**Location** | Pointer to [**CloudAwsVpcRelationship**](CloudAwsVpcRelationship.md) |  | [optional] 
**SecurityGroups** | Pointer to [**[]CloudAwsSecurityGroupRelationship**](CloudAwsSecurityGroupRelationship.md) | An array of relationships to cloudAwsSecurityGroup resources. | [optional] [readonly] 

## Methods

### NewCloudAwsVirtualMachine

`func NewCloudAwsVirtualMachine(classId string, objectType string, ) *CloudAwsVirtualMachine`

NewCloudAwsVirtualMachine instantiates a new CloudAwsVirtualMachine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAwsVirtualMachineWithDefaults

`func NewCloudAwsVirtualMachineWithDefaults() *CloudAwsVirtualMachine`

NewCloudAwsVirtualMachineWithDefaults instantiates a new CloudAwsVirtualMachine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudAwsVirtualMachine) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudAwsVirtualMachine) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudAwsVirtualMachine) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudAwsVirtualMachine) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudAwsVirtualMachine) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudAwsVirtualMachine) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetKeyPair

`func (o *CloudAwsVirtualMachine) GetKeyPair() CloudAwsKeyPairRelationship`

GetKeyPair returns the KeyPair field if non-nil, zero value otherwise.

### GetKeyPairOk

`func (o *CloudAwsVirtualMachine) GetKeyPairOk() (*CloudAwsKeyPairRelationship, bool)`

GetKeyPairOk returns a tuple with the KeyPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPair

`func (o *CloudAwsVirtualMachine) SetKeyPair(v CloudAwsKeyPairRelationship)`

SetKeyPair sets KeyPair field to given value.

### HasKeyPair

`func (o *CloudAwsVirtualMachine) HasKeyPair() bool`

HasKeyPair returns a boolean if a field has been set.

### GetLocation

`func (o *CloudAwsVirtualMachine) GetLocation() CloudAwsVpcRelationship`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CloudAwsVirtualMachine) GetLocationOk() (*CloudAwsVpcRelationship, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CloudAwsVirtualMachine) SetLocation(v CloudAwsVpcRelationship)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *CloudAwsVirtualMachine) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *CloudAwsVirtualMachine) GetSecurityGroups() []CloudAwsSecurityGroupRelationship`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *CloudAwsVirtualMachine) GetSecurityGroupsOk() (*[]CloudAwsSecurityGroupRelationship, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *CloudAwsVirtualMachine) SetSecurityGroups(v []CloudAwsSecurityGroupRelationship)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *CloudAwsVirtualMachine) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *CloudAwsVirtualMachine) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *CloudAwsVirtualMachine) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


