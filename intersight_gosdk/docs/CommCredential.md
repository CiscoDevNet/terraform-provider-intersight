# CommCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] 
**Password** | Pointer to **string** | Password for the Managed Device. | [optional] 
**Username** | Pointer to **string** | Username for the Managed Device. Format and restrictions are not enforced here but usually follow the ManagedDevice requirements. | [optional] 

## Methods

### NewCommCredential

`func NewCommCredential() *CommCredential`

NewCommCredential instantiates a new CommCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommCredentialWithDefaults

`func NewCommCredentialWithDefaults() *CommCredential`

NewCommCredentialWithDefaults instantiates a new CommCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsPasswordSet

`func (o *CommCredential) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *CommCredential) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *CommCredential) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *CommCredential) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetPassword

`func (o *CommCredential) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CommCredential) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CommCredential) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CommCredential) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUsername

`func (o *CommCredential) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CommCredential) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CommCredential) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CommCredential) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


