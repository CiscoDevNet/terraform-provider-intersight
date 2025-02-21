# IamAbstractClientApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "webhook.Endpoint"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "webhook.Endpoint"]
**AppCredentials** | Pointer to [**NullableIamAppCredentials**](IamAppCredentials.md) |  | [optional] 
**AuthType** | Pointer to **string** | Type of authentication used by the clientApp. * &#x60;basic&#x60; - The client uses basic username/password authentication. The password is expected to be a JWT token. * &#x60;none&#x60; - No authentication method specified by the client. * &#x60;bearer-token&#x60; - The client uses a long-lived bearer token to authenticate. * &#x60;auth-code&#x60; - The client uses OAuth Authorization Grant Flow without PKCE for authentication. * &#x60;client-credentials&#x60; - The client uses OAuth Client Credentials Flow for authentication. | [optional] [default to "basic"]
**CredentialsAction** | Pointer to **string** | An action to be performed on the credentials. * &#x60;none&#x60; - No action to be performed. * &#x60;regenerateCredentials&#x60; - Allows for revocation and regeneration of a token. The old token associated with the client application. will not be usable and a new token will be generated. | [optional] [default to "none"]
**AppRegistration** | Pointer to [**NullableIamAppRegistrationRelationship**](IamAppRegistrationRelationship.md) |  | [optional] 

## Methods

### NewIamAbstractClientApplication

`func NewIamAbstractClientApplication(classId string, objectType string, ) *IamAbstractClientApplication`

NewIamAbstractClientApplication instantiates a new IamAbstractClientApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamAbstractClientApplicationWithDefaults

`func NewIamAbstractClientApplicationWithDefaults() *IamAbstractClientApplication`

NewIamAbstractClientApplicationWithDefaults instantiates a new IamAbstractClientApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamAbstractClientApplication) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamAbstractClientApplication) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamAbstractClientApplication) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamAbstractClientApplication) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamAbstractClientApplication) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamAbstractClientApplication) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAppCredentials

`func (o *IamAbstractClientApplication) GetAppCredentials() IamAppCredentials`

GetAppCredentials returns the AppCredentials field if non-nil, zero value otherwise.

### GetAppCredentialsOk

`func (o *IamAbstractClientApplication) GetAppCredentialsOk() (*IamAppCredentials, bool)`

GetAppCredentialsOk returns a tuple with the AppCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppCredentials

`func (o *IamAbstractClientApplication) SetAppCredentials(v IamAppCredentials)`

SetAppCredentials sets AppCredentials field to given value.

### HasAppCredentials

`func (o *IamAbstractClientApplication) HasAppCredentials() bool`

HasAppCredentials returns a boolean if a field has been set.

### SetAppCredentialsNil

`func (o *IamAbstractClientApplication) SetAppCredentialsNil(b bool)`

 SetAppCredentialsNil sets the value for AppCredentials to be an explicit nil

### UnsetAppCredentials
`func (o *IamAbstractClientApplication) UnsetAppCredentials()`

UnsetAppCredentials ensures that no value is present for AppCredentials, not even an explicit nil
### GetAuthType

`func (o *IamAbstractClientApplication) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *IamAbstractClientApplication) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *IamAbstractClientApplication) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *IamAbstractClientApplication) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetCredentialsAction

`func (o *IamAbstractClientApplication) GetCredentialsAction() string`

GetCredentialsAction returns the CredentialsAction field if non-nil, zero value otherwise.

### GetCredentialsActionOk

`func (o *IamAbstractClientApplication) GetCredentialsActionOk() (*string, bool)`

GetCredentialsActionOk returns a tuple with the CredentialsAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsAction

`func (o *IamAbstractClientApplication) SetCredentialsAction(v string)`

SetCredentialsAction sets CredentialsAction field to given value.

### HasCredentialsAction

`func (o *IamAbstractClientApplication) HasCredentialsAction() bool`

HasCredentialsAction returns a boolean if a field has been set.

### GetAppRegistration

`func (o *IamAbstractClientApplication) GetAppRegistration() IamAppRegistrationRelationship`

GetAppRegistration returns the AppRegistration field if non-nil, zero value otherwise.

### GetAppRegistrationOk

`func (o *IamAbstractClientApplication) GetAppRegistrationOk() (*IamAppRegistrationRelationship, bool)`

GetAppRegistrationOk returns a tuple with the AppRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRegistration

`func (o *IamAbstractClientApplication) SetAppRegistration(v IamAppRegistrationRelationship)`

SetAppRegistration sets AppRegistration field to given value.

### HasAppRegistration

`func (o *IamAbstractClientApplication) HasAppRegistration() bool`

HasAppRegistration returns a boolean if a field has been set.

### SetAppRegistrationNil

`func (o *IamAbstractClientApplication) SetAppRegistrationNil(b bool)`

 SetAppRegistrationNil sets the value for AppRegistration to be an explicit nil

### UnsetAppRegistration
`func (o *IamAbstractClientApplication) UnsetAppRegistration()`

UnsetAppRegistration ensures that no value is present for AppRegistration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


