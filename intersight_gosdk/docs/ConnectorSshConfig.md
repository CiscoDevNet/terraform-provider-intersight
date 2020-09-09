# ConnectorSshConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** | Password to use in the connection credentials (If empty the private key will be used). | [optional] 
**Pkey** | Pointer to **string** | The private key to use in the connection credentials (Optional if password is given). | [optional] 
**Target** | Pointer to **string** | The remote server to connect to. | [optional] 
**User** | Pointer to **string** | Username for the remote connection. | [optional] 

## Methods

### NewConnectorSshConfig

`func NewConnectorSshConfig() *ConnectorSshConfig`

NewConnectorSshConfig instantiates a new ConnectorSshConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorSshConfigWithDefaults

`func NewConnectorSshConfigWithDefaults() *ConnectorSshConfig`

NewConnectorSshConfigWithDefaults instantiates a new ConnectorSshConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *ConnectorSshConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ConnectorSshConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ConnectorSshConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ConnectorSshConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPkey

`func (o *ConnectorSshConfig) GetPkey() string`

GetPkey returns the Pkey field if non-nil, zero value otherwise.

### GetPkeyOk

`func (o *ConnectorSshConfig) GetPkeyOk() (*string, bool)`

GetPkeyOk returns a tuple with the Pkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkey

`func (o *ConnectorSshConfig) SetPkey(v string)`

SetPkey sets Pkey field to given value.

### HasPkey

`func (o *ConnectorSshConfig) HasPkey() bool`

HasPkey returns a boolean if a field has been set.

### GetTarget

`func (o *ConnectorSshConfig) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ConnectorSshConfig) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ConnectorSshConfig) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *ConnectorSshConfig) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetUser

`func (o *ConnectorSshConfig) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ConnectorSshConfig) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ConnectorSshConfig) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *ConnectorSshConfig) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


