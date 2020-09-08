# IamAccountExperience

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Features** | Pointer to [**[]IamFeatureDefinition**](iam.FeatureDefinition.md) |  | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 

## Methods

### NewIamAccountExperience

`func NewIamAccountExperience() *IamAccountExperience`

NewIamAccountExperience instantiates a new IamAccountExperience object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamAccountExperienceWithDefaults

`func NewIamAccountExperienceWithDefaults() *IamAccountExperience`

NewIamAccountExperienceWithDefaults instantiates a new IamAccountExperience object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatures

`func (o *IamAccountExperience) GetFeatures() []IamFeatureDefinition`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *IamAccountExperience) GetFeaturesOk() (*[]IamFeatureDefinition, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *IamAccountExperience) SetFeatures(v []IamFeatureDefinition)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *IamAccountExperience) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetAccount

`func (o *IamAccountExperience) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamAccountExperience) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamAccountExperience) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamAccountExperience) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


