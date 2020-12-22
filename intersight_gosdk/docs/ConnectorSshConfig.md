# ConnectorSshConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "connector.SshConfig"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "connector.SshConfig"]
**JumpHost** | Pointer to **string** | A jump host for establishing a connection to a server. Plugin will first establish a connection to this server, then create a tunneled connection to the target host. | [optional] 
**Password** | Pointer to **string** | Password to use in the connection credentials (If empty the private key will be used). | [optional] 
**Pkey** | Pointer to **string** | The private key to use in the connection credentials (Optional if password is given). | [optional] 
**Target** | Pointer to **string** | The remote server to connect to. | [optional] 
**User** | Pointer to **string** | Username for the remote connection. | [optional] 

## Methods

### NewConnectorSshConfig

`func NewConnectorSshConfig(classId string, objectType string, ) *ConnectorSshConfig`

NewConnectorSshConfig instantiates a new ConnectorSshConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorSshConfigWithDefaults

`func NewConnectorSshConfigWithDefaults() *ConnectorSshConfig`

NewConnectorSshConfigWithDefaults instantiates a new ConnectorSshConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConnectorSshConfig) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConnectorSshConfig) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConnectorSshConfig) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConnectorSshConfig) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConnectorSshConfig) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConnectorSshConfig) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetJumpHost

`func (o *ConnectorSshConfig) GetJumpHost() string`

GetJumpHost returns the JumpHost field if non-nil, zero value otherwise.

### GetJumpHostOk

`func (o *ConnectorSshConfig) GetJumpHostOk() (*string, bool)`

GetJumpHostOk returns a tuple with the JumpHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJumpHost

`func (o *ConnectorSshConfig) SetJumpHost(v string)`

SetJumpHost sets JumpHost field to given value.

### HasJumpHost

`func (o *ConnectorSshConfig) HasJumpHost() bool`

HasJumpHost returns a boolean if a field has been set.

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


