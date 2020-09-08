# WorkflowSshConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** | Password to use in the SSH connection credentials (If empty then private key will be used). | [optional] 
**Target** | Pointer to **string** | The remote server to connect to. IPv4 address represented in dot decimal notation. | [optional] 
**User** | Pointer to **string** | Username for the remote SSH connection. | [optional] 

## Methods

### NewWorkflowSshConfig

`func NewWorkflowSshConfig() *WorkflowSshConfig`

NewWorkflowSshConfig instantiates a new WorkflowSshConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSshConfigWithDefaults

`func NewWorkflowSshConfigWithDefaults() *WorkflowSshConfig`

NewWorkflowSshConfigWithDefaults instantiates a new WorkflowSshConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


