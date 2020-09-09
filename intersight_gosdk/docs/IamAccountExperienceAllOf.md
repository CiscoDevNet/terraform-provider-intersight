# IamAccountExperienceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Features** | Pointer to [**[]IamFeatureDefinition**](iam.FeatureDefinition.md) |  | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 

## Methods

### NewIamAccountExperienceAllOf

`func NewIamAccountExperienceAllOf() *IamAccountExperienceAllOf`

NewIamAccountExperienceAllOf instantiates a new IamAccountExperienceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamAccountExperienceAllOfWithDefaults

`func NewIamAccountExperienceAllOfWithDefaults() *IamAccountExperienceAllOf`

NewIamAccountExperienceAllOfWithDefaults instantiates a new IamAccountExperienceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatures

`func (o *IamAccountExperienceAllOf) GetFeatures() []IamFeatureDefinition`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *IamAccountExperienceAllOf) GetFeaturesOk() (*[]IamFeatureDefinition, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *IamAccountExperienceAllOf) SetFeatures(v []IamFeatureDefinition)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *IamAccountExperienceAllOf) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetAccount

`func (o *IamAccountExperienceAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamAccountExperienceAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamAccountExperienceAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamAccountExperienceAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


