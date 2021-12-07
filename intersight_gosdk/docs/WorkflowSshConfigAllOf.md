# WorkflowSshConfigAllOf

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

### NewWorkflowSshConfigAllOf

`func NewWorkflowSshConfigAllOf(classId string, objectType string, ) *WorkflowSshConfigAllOf`

NewWorkflowSshConfigAllOf instantiates a new WorkflowSshConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSshConfigAllOfWithDefaults

`func NewWorkflowSshConfigAllOfWithDefaults() *WorkflowSshConfigAllOf`

NewWorkflowSshConfigAllOfWithDefaults instantiates a new WorkflowSshConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowSshConfigAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowSshConfigAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowSshConfigAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowSshConfigAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowSshConfigAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowSshConfigAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsPassphraseSet

`func (o *WorkflowSshConfigAllOf) GetIsPassphraseSet() bool`

GetIsPassphraseSet returns the IsPassphraseSet field if non-nil, zero value otherwise.

### GetIsPassphraseSetOk

`func (o *WorkflowSshConfigAllOf) GetIsPassphraseSetOk() (*bool, bool)`

GetIsPassphraseSetOk returns a tuple with the IsPassphraseSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPassphraseSet

`func (o *WorkflowSshConfigAllOf) SetIsPassphraseSet(v bool)`

SetIsPassphraseSet sets IsPassphraseSet field to given value.

### HasIsPassphraseSet

`func (o *WorkflowSshConfigAllOf) HasIsPassphraseSet() bool`

HasIsPassphraseSet returns a boolean if a field has been set.

### GetIsPasswordSet

`func (o *WorkflowSshConfigAllOf) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *WorkflowSshConfigAllOf) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *WorkflowSshConfigAllOf) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *WorkflowSshConfigAllOf) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetIsPrivateKeySet

`func (o *WorkflowSshConfigAllOf) GetIsPrivateKeySet() bool`

GetIsPrivateKeySet returns the IsPrivateKeySet field if non-nil, zero value otherwise.

### GetIsPrivateKeySetOk

`func (o *WorkflowSshConfigAllOf) GetIsPrivateKeySetOk() (*bool, bool)`

GetIsPrivateKeySetOk returns a tuple with the IsPrivateKeySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivateKeySet

`func (o *WorkflowSshConfigAllOf) SetIsPrivateKeySet(v bool)`

SetIsPrivateKeySet sets IsPrivateKeySet field to given value.

### HasIsPrivateKeySet

`func (o *WorkflowSshConfigAllOf) HasIsPrivateKeySet() bool`

HasIsPrivateKeySet returns a boolean if a field has been set.

### GetPassphrase

`func (o *WorkflowSshConfigAllOf) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *WorkflowSshConfigAllOf) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *WorkflowSshConfigAllOf) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *WorkflowSshConfigAllOf) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### GetPassword

`func (o *WorkflowSshConfigAllOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *WorkflowSshConfigAllOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *WorkflowSshConfigAllOf) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *WorkflowSshConfigAllOf) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPrivateKey

`func (o *WorkflowSshConfigAllOf) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *WorkflowSshConfigAllOf) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *WorkflowSshConfigAllOf) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *WorkflowSshConfigAllOf) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetTarget

`func (o *WorkflowSshConfigAllOf) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *WorkflowSshConfigAllOf) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *WorkflowSshConfigAllOf) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *WorkflowSshConfigAllOf) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetUser

`func (o *WorkflowSshConfigAllOf) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *WorkflowSshConfigAllOf) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *WorkflowSshConfigAllOf) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *WorkflowSshConfigAllOf) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


