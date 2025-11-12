# IamBasicAuthCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.BasicAuthCredentials"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.BasicAuthCredentials"]
**Password** | Pointer to **string** | Contains the password for the basic auth flow. | [optional] 
**Username** | Pointer to **string** | Contains the username for the basic auth flow. | [optional] 

## Methods

### NewIamBasicAuthCredentials

`func NewIamBasicAuthCredentials(classId string, objectType string, ) *IamBasicAuthCredentials`

NewIamBasicAuthCredentials instantiates a new IamBasicAuthCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamBasicAuthCredentialsWithDefaults

`func NewIamBasicAuthCredentialsWithDefaults() *IamBasicAuthCredentials`

NewIamBasicAuthCredentialsWithDefaults instantiates a new IamBasicAuthCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamBasicAuthCredentials) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamBasicAuthCredentials) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamBasicAuthCredentials) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamBasicAuthCredentials) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamBasicAuthCredentials) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamBasicAuthCredentials) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPassword

`func (o *IamBasicAuthCredentials) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *IamBasicAuthCredentials) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *IamBasicAuthCredentials) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *IamBasicAuthCredentials) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUsername

`func (o *IamBasicAuthCredentials) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *IamBasicAuthCredentials) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *IamBasicAuthCredentials) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *IamBasicAuthCredentials) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


