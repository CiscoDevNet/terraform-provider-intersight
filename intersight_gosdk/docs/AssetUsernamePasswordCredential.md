# AssetUsernamePasswordCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] 
**Password** | Pointer to **string** | The password used to authenticate with a managed target. | [optional] 
**Username** | Pointer to **string** | The username used to authenticate with a managed target. | [optional] 

## Methods

### NewAssetUsernamePasswordCredential

`func NewAssetUsernamePasswordCredential() *AssetUsernamePasswordCredential`

NewAssetUsernamePasswordCredential instantiates a new AssetUsernamePasswordCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetUsernamePasswordCredentialWithDefaults

`func NewAssetUsernamePasswordCredentialWithDefaults() *AssetUsernamePasswordCredential`

NewAssetUsernamePasswordCredentialWithDefaults instantiates a new AssetUsernamePasswordCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsPasswordSet

`func (o *AssetUsernamePasswordCredential) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *AssetUsernamePasswordCredential) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *AssetUsernamePasswordCredential) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *AssetUsernamePasswordCredential) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetPassword

`func (o *AssetUsernamePasswordCredential) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AssetUsernamePasswordCredential) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AssetUsernamePasswordCredential) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AssetUsernamePasswordCredential) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUsername

`func (o *AssetUsernamePasswordCredential) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AssetUsernamePasswordCredential) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AssetUsernamePasswordCredential) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AssetUsernamePasswordCredential) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


