# IamTestIdpConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.TestIdpConfiguration"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.TestIdpConfiguration"]
**ErrorDetails** | Pointer to **string** | Error returned by the IdP when the configuration test fails. | [optional] [readonly] 
**IdpEntityId** | Pointer to **string** | Entity ID of the IdP whose configuration needs to be tested. | [optional] [readonly] 
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] [default to false]
**NoOfGroups** | Pointer to **int64** | Total number of groups of the user received from the IdP after successful authentication. | [optional] [readonly] 
**Password** | Pointer to **string** | The password of the test user for testing the IdP configuration settings. It can be any string that adheres to the following constraints. It can have character except spaces, tabs, line breaks. It cannot be more than 256 characters. | [optional] 
**Status** | Pointer to **string** | IdP configuration test result. * &#x60;Unverified&#x60; - The IdP configuration is yet to be verified. * &#x60;Verified&#x60; - Successfully authenticated the user via the configured IdP and verified the configuration parameters. * &#x60;Pending&#x60; - Pending verification of the IdP configuration. * &#x60;Failed&#x60; - Could not authenticate the user via the configured IdP and verify its configuration. | [optional] [readonly] [default to "Unverified"]
**TriggerTest** | Pointer to **bool** | Trigger property used to initiate an IdP configuration test. | [optional] 
**Username** | Pointer to **string** | Email or userId of the test user for testing the IdP configuration settings. | [optional] 
**Idp** | Pointer to [**NullableIamIdpRelationship**](IamIdpRelationship.md) |  | [optional] 

## Methods

### NewIamTestIdpConfiguration

`func NewIamTestIdpConfiguration(classId string, objectType string, ) *IamTestIdpConfiguration`

NewIamTestIdpConfiguration instantiates a new IamTestIdpConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamTestIdpConfigurationWithDefaults

`func NewIamTestIdpConfigurationWithDefaults() *IamTestIdpConfiguration`

NewIamTestIdpConfigurationWithDefaults instantiates a new IamTestIdpConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamTestIdpConfiguration) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamTestIdpConfiguration) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamTestIdpConfiguration) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamTestIdpConfiguration) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamTestIdpConfiguration) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamTestIdpConfiguration) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetErrorDetails

`func (o *IamTestIdpConfiguration) GetErrorDetails() string`

GetErrorDetails returns the ErrorDetails field if non-nil, zero value otherwise.

### GetErrorDetailsOk

`func (o *IamTestIdpConfiguration) GetErrorDetailsOk() (*string, bool)`

GetErrorDetailsOk returns a tuple with the ErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetails

`func (o *IamTestIdpConfiguration) SetErrorDetails(v string)`

SetErrorDetails sets ErrorDetails field to given value.

### HasErrorDetails

`func (o *IamTestIdpConfiguration) HasErrorDetails() bool`

HasErrorDetails returns a boolean if a field has been set.

### GetIdpEntityId

`func (o *IamTestIdpConfiguration) GetIdpEntityId() string`

GetIdpEntityId returns the IdpEntityId field if non-nil, zero value otherwise.

### GetIdpEntityIdOk

`func (o *IamTestIdpConfiguration) GetIdpEntityIdOk() (*string, bool)`

GetIdpEntityIdOk returns a tuple with the IdpEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpEntityId

`func (o *IamTestIdpConfiguration) SetIdpEntityId(v string)`

SetIdpEntityId sets IdpEntityId field to given value.

### HasIdpEntityId

`func (o *IamTestIdpConfiguration) HasIdpEntityId() bool`

HasIdpEntityId returns a boolean if a field has been set.

### GetIsPasswordSet

`func (o *IamTestIdpConfiguration) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *IamTestIdpConfiguration) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *IamTestIdpConfiguration) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *IamTestIdpConfiguration) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetNoOfGroups

`func (o *IamTestIdpConfiguration) GetNoOfGroups() int64`

GetNoOfGroups returns the NoOfGroups field if non-nil, zero value otherwise.

### GetNoOfGroupsOk

`func (o *IamTestIdpConfiguration) GetNoOfGroupsOk() (*int64, bool)`

GetNoOfGroupsOk returns a tuple with the NoOfGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOfGroups

`func (o *IamTestIdpConfiguration) SetNoOfGroups(v int64)`

SetNoOfGroups sets NoOfGroups field to given value.

### HasNoOfGroups

`func (o *IamTestIdpConfiguration) HasNoOfGroups() bool`

HasNoOfGroups returns a boolean if a field has been set.

### GetPassword

`func (o *IamTestIdpConfiguration) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *IamTestIdpConfiguration) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *IamTestIdpConfiguration) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *IamTestIdpConfiguration) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetStatus

`func (o *IamTestIdpConfiguration) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IamTestIdpConfiguration) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IamTestIdpConfiguration) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IamTestIdpConfiguration) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTriggerTest

`func (o *IamTestIdpConfiguration) GetTriggerTest() bool`

GetTriggerTest returns the TriggerTest field if non-nil, zero value otherwise.

### GetTriggerTestOk

`func (o *IamTestIdpConfiguration) GetTriggerTestOk() (*bool, bool)`

GetTriggerTestOk returns a tuple with the TriggerTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerTest

`func (o *IamTestIdpConfiguration) SetTriggerTest(v bool)`

SetTriggerTest sets TriggerTest field to given value.

### HasTriggerTest

`func (o *IamTestIdpConfiguration) HasTriggerTest() bool`

HasTriggerTest returns a boolean if a field has been set.

### GetUsername

`func (o *IamTestIdpConfiguration) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *IamTestIdpConfiguration) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *IamTestIdpConfiguration) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *IamTestIdpConfiguration) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetIdp

`func (o *IamTestIdpConfiguration) GetIdp() IamIdpRelationship`

GetIdp returns the Idp field if non-nil, zero value otherwise.

### GetIdpOk

`func (o *IamTestIdpConfiguration) GetIdpOk() (*IamIdpRelationship, bool)`

GetIdpOk returns a tuple with the Idp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdp

`func (o *IamTestIdpConfiguration) SetIdp(v IamIdpRelationship)`

SetIdp sets Idp field to given value.

### HasIdp

`func (o *IamTestIdpConfiguration) HasIdp() bool`

HasIdp returns a boolean if a field has been set.

### SetIdpNil

`func (o *IamTestIdpConfiguration) SetIdpNil(b bool)`

 SetIdpNil sets the value for Idp to be an explicit nil

### UnsetIdp
`func (o *IamTestIdpConfiguration) UnsetIdp()`

UnsetIdp ensures that no value is present for Idp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


