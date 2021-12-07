# WorkflowSshConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.SshConfig"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.SshConfig"]
**IsPassphraseSet** | Pointer to **bool** | Indicates whether the value of the &#39;passphrase&#39; property has been set. | [optional] [readonly] [default to false]
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] [default to false]
**IsPrivateKeySet** | Pointer to **bool** | Indicates whether the value of the &#39;privateKey&#39; property has been set. | [optional] [readonly] [default to false]
**Passphrase** | Pointer to **string** | Optional passphrase if provided while creating the private key. | [optional] 
**Password** | Pointer to **string** | Password to use in the SSH connection credentials (If empty, private key will be used). | [optional] 
**PrivateKey** | Pointer to **string** | PEM encoded private key to be used in the SSH connection credentials (Optional if password is given). | [optional] 
**Target** | Pointer to **string** | The remote server to connect to. IPv4 address represented in dot decimal notation or hostname can be specified. | [optional] 
**User** | Pointer to **string** | Username for the remote SSH connection. | [optional] 

## Methods

### NewWorkflowSshConfig

`func NewWorkflowSshConfig(classId string, objectType string, ) *WorkflowSshConfig`

NewWorkflowSshConfig instantiates a new WorkflowSshConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSshConfigWithDefaults

`func NewWorkflowSshConfigWithDefaults() *WorkflowSshConfig`

NewWorkflowSshConfigWithDefaults instantiates a new WorkflowSshConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowSshConfig) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowSshConfig) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowSshConfig) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowSshConfig) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowSshConfig) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowSshConfig) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsPassphraseSet

`func (o *WorkflowSshConfig) GetIsPassphraseSet() bool`

GetIsPassphraseSet returns the IsPassphraseSet field if non-nil, zero value otherwise.

### GetIsPassphraseSetOk

`func (o *WorkflowSshConfig) GetIsPassphraseSetOk() (*bool, bool)`

GetIsPassphraseSetOk returns a tuple with the IsPassphraseSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPassphraseSet

`func (o *WorkflowSshConfig) SetIsPassphraseSet(v bool)`

SetIsPassphraseSet sets IsPassphraseSet field to given value.

### HasIsPassphraseSet

`func (o *WorkflowSshConfig) HasIsPassphraseSet() bool`

HasIsPassphraseSet returns a boolean if a field has been set.

### GetIsPasswordSet

`func (o *WorkflowSshConfig) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *WorkflowSshConfig) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *WorkflowSshConfig) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *WorkflowSshConfig) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetIsPrivateKeySet

`func (o *WorkflowSshConfig) GetIsPrivateKeySet() bool`

GetIsPrivateKeySet returns the IsPrivateKeySet field if non-nil, zero value otherwise.

### GetIsPrivateKeySetOk

`func (o *WorkflowSshConfig) GetIsPrivateKeySetOk() (*bool, bool)`

GetIsPrivateKeySetOk returns a tuple with the IsPrivateKeySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivateKeySet

`func (o *WorkflowSshConfig) SetIsPrivateKeySet(v bool)`

SetIsPrivateKeySet sets IsPrivateKeySet field to given value.

### HasIsPrivateKeySet

`func (o *WorkflowSshConfig) HasIsPrivateKeySet() bool`

HasIsPrivateKeySet returns a boolean if a field has been set.

### GetPassphrase

`func (o *WorkflowSshConfig) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *WorkflowSshConfig) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *WorkflowSshConfig) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *WorkflowSshConfig) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### GetPassword

`func (o *WorkflowSshConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *WorkflowSshConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *WorkflowSshConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *WorkflowSshConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPrivateKey

`func (o *WorkflowSshConfig) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *WorkflowSshConfig) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *WorkflowSshConfig) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *WorkflowSshConfig) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetTarget

`func (o *WorkflowSshConfig) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *WorkflowSshConfig) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *WorkflowSshConfig) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *WorkflowSshConfig) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetUser

`func (o *WorkflowSshConfig) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *WorkflowSshConfig) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *WorkflowSshConfig) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *WorkflowSshConfig) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


