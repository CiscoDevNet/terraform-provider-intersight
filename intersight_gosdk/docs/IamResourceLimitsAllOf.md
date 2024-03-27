# IamResourceLimitsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.ResourceLimits"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.ResourceLimits"]
**AllowApiKeysWithoutExpiry** | Pointer to **bool** | Boolean value used to decide whether API keys that never expire are allowed for the account. This allows creation of API keys which are perpetual which can used for specific applications where rotation of API keys are not feasible. | [optional] [default to false]
**AllowAppRegistrationsWithoutExpiry** | Pointer to **bool** | Boolean value used to decide whether App Registration that never expire are allowed for the account. | [optional] [default to false]
**MaxApiKeyExpiry** | Pointer to **int64** | The maximum expiration period (in seconds) allowed for API keys. The default value is 180 days or 15552000 seconds. It is shown to user in days for readability. | [optional] [default to 15552000]
**MaxAppRegistrationExpiry** | Pointer to **int64** | The maximum expiration period (in seconds) allowed for App Registration. The default value is 180 days or 15552000 seconds. It is shown to user in days for readability. | [optional] [default to 15552000]
**PerAccountUserLimit** | Pointer to **int64** | The maximum number of users allowed in an account. The default value is 200. | [optional] [readonly] 
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewIamResourceLimitsAllOf

`func NewIamResourceLimitsAllOf(classId string, objectType string, ) *IamResourceLimitsAllOf`

NewIamResourceLimitsAllOf instantiates a new IamResourceLimitsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamResourceLimitsAllOfWithDefaults

`func NewIamResourceLimitsAllOfWithDefaults() *IamResourceLimitsAllOf`

NewIamResourceLimitsAllOfWithDefaults instantiates a new IamResourceLimitsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamResourceLimitsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamResourceLimitsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamResourceLimitsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamResourceLimitsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamResourceLimitsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamResourceLimitsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAllowApiKeysWithoutExpiry

`func (o *IamResourceLimitsAllOf) GetAllowApiKeysWithoutExpiry() bool`

GetAllowApiKeysWithoutExpiry returns the AllowApiKeysWithoutExpiry field if non-nil, zero value otherwise.

### GetAllowApiKeysWithoutExpiryOk

`func (o *IamResourceLimitsAllOf) GetAllowApiKeysWithoutExpiryOk() (*bool, bool)`

GetAllowApiKeysWithoutExpiryOk returns a tuple with the AllowApiKeysWithoutExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowApiKeysWithoutExpiry

`func (o *IamResourceLimitsAllOf) SetAllowApiKeysWithoutExpiry(v bool)`

SetAllowApiKeysWithoutExpiry sets AllowApiKeysWithoutExpiry field to given value.

### HasAllowApiKeysWithoutExpiry

`func (o *IamResourceLimitsAllOf) HasAllowApiKeysWithoutExpiry() bool`

HasAllowApiKeysWithoutExpiry returns a boolean if a field has been set.

### GetAllowAppRegistrationsWithoutExpiry

`func (o *IamResourceLimitsAllOf) GetAllowAppRegistrationsWithoutExpiry() bool`

GetAllowAppRegistrationsWithoutExpiry returns the AllowAppRegistrationsWithoutExpiry field if non-nil, zero value otherwise.

### GetAllowAppRegistrationsWithoutExpiryOk

`func (o *IamResourceLimitsAllOf) GetAllowAppRegistrationsWithoutExpiryOk() (*bool, bool)`

GetAllowAppRegistrationsWithoutExpiryOk returns a tuple with the AllowAppRegistrationsWithoutExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAppRegistrationsWithoutExpiry

`func (o *IamResourceLimitsAllOf) SetAllowAppRegistrationsWithoutExpiry(v bool)`

SetAllowAppRegistrationsWithoutExpiry sets AllowAppRegistrationsWithoutExpiry field to given value.

### HasAllowAppRegistrationsWithoutExpiry

`func (o *IamResourceLimitsAllOf) HasAllowAppRegistrationsWithoutExpiry() bool`

HasAllowAppRegistrationsWithoutExpiry returns a boolean if a field has been set.

### GetMaxApiKeyExpiry

`func (o *IamResourceLimitsAllOf) GetMaxApiKeyExpiry() int64`

GetMaxApiKeyExpiry returns the MaxApiKeyExpiry field if non-nil, zero value otherwise.

### GetMaxApiKeyExpiryOk

`func (o *IamResourceLimitsAllOf) GetMaxApiKeyExpiryOk() (*int64, bool)`

GetMaxApiKeyExpiryOk returns a tuple with the MaxApiKeyExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxApiKeyExpiry

`func (o *IamResourceLimitsAllOf) SetMaxApiKeyExpiry(v int64)`

SetMaxApiKeyExpiry sets MaxApiKeyExpiry field to given value.

### HasMaxApiKeyExpiry

`func (o *IamResourceLimitsAllOf) HasMaxApiKeyExpiry() bool`

HasMaxApiKeyExpiry returns a boolean if a field has been set.

### GetMaxAppRegistrationExpiry

`func (o *IamResourceLimitsAllOf) GetMaxAppRegistrationExpiry() int64`

GetMaxAppRegistrationExpiry returns the MaxAppRegistrationExpiry field if non-nil, zero value otherwise.

### GetMaxAppRegistrationExpiryOk

`func (o *IamResourceLimitsAllOf) GetMaxAppRegistrationExpiryOk() (*int64, bool)`

GetMaxAppRegistrationExpiryOk returns a tuple with the MaxAppRegistrationExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAppRegistrationExpiry

`func (o *IamResourceLimitsAllOf) SetMaxAppRegistrationExpiry(v int64)`

SetMaxAppRegistrationExpiry sets MaxAppRegistrationExpiry field to given value.

### HasMaxAppRegistrationExpiry

`func (o *IamResourceLimitsAllOf) HasMaxAppRegistrationExpiry() bool`

HasMaxAppRegistrationExpiry returns a boolean if a field has been set.

### GetPerAccountUserLimit

`func (o *IamResourceLimitsAllOf) GetPerAccountUserLimit() int64`

GetPerAccountUserLimit returns the PerAccountUserLimit field if non-nil, zero value otherwise.

### GetPerAccountUserLimitOk

`func (o *IamResourceLimitsAllOf) GetPerAccountUserLimitOk() (*int64, bool)`

GetPerAccountUserLimitOk returns a tuple with the PerAccountUserLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerAccountUserLimit

`func (o *IamResourceLimitsAllOf) SetPerAccountUserLimit(v int64)`

SetPerAccountUserLimit sets PerAccountUserLimit field to given value.

### HasPerAccountUserLimit

`func (o *IamResourceLimitsAllOf) HasPerAccountUserLimit() bool`

HasPerAccountUserLimit returns a boolean if a field has been set.

### GetAccount

`func (o *IamResourceLimitsAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamResourceLimitsAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamResourceLimitsAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamResourceLimitsAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


