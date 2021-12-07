# ConnectorSshConfigAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "connector.SshConfig"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "connector.SshConfig"]
**JumpHost** | Pointer to **string** | A jump host for establishing a connection to a server. Plugin will first establish a connection to this server, then create a tunneled connection to the target host. | [optional] 
**Passphrase** | Pointer to **string** | Optional passphrase if provided while creating the private key. | [optional] 
**Password** | Pointer to **string** | Password to use in the connection credentials (If empty the private key will be used). | [optional] 
**Pkey** | Pointer to **string** | The private key to use in the connection credentials (Optional if password is given). | [optional] 
**PkeyString** | Pointer to **string** | The private key to use in the connection credentials in string format. | [optional] 
**Target** | Pointer to **string** | The remote server to connect to. | [optional] 
**User** | Pointer to **string** | Username for the remote connection. | [optional] 

## Methods

### NewConnectorSshConfigAllOf

`func NewConnectorSshConfigAllOf(classId string, objectType string, ) *ConnectorSshConfigAllOf`

NewConnectorSshConfigAllOf instantiates a new ConnectorSshConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorSshConfigAllOfWithDefaults

`func NewConnectorSshConfigAllOfWithDefaults() *ConnectorSshConfigAllOf`

NewConnectorSshConfigAllOfWithDefaults instantiates a new ConnectorSshConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConnectorSshConfigAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConnectorSshConfigAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConnectorSshConfigAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConnectorSshConfigAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConnectorSshConfigAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConnectorSshConfigAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetJumpHost

`func (o *ConnectorSshConfigAllOf) GetJumpHost() string`

GetJumpHost returns the JumpHost field if non-nil, zero value otherwise.

### GetJumpHostOk

`func (o *ConnectorSshConfigAllOf) GetJumpHostOk() (*string, bool)`

GetJumpHostOk returns a tuple with the JumpHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJumpHost

`func (o *ConnectorSshConfigAllOf) SetJumpHost(v string)`

SetJumpHost sets JumpHost field to given value.

### HasJumpHost

`func (o *ConnectorSshConfigAllOf) HasJumpHost() bool`

HasJumpHost returns a boolean if a field has been set.

### GetPassphrase

`func (o *ConnectorSshConfigAllOf) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *ConnectorSshConfigAllOf) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *ConnectorSshConfigAllOf) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *ConnectorSshConfigAllOf) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### GetPassword

`func (o *ConnectorSshConfigAllOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ConnectorSshConfigAllOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ConnectorSshConfigAllOf) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ConnectorSshConfigAllOf) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPkey

`func (o *ConnectorSshConfigAllOf) GetPkey() string`

GetPkey returns the Pkey field if non-nil, zero value otherwise.

### GetPkeyOk

`func (o *ConnectorSshConfigAllOf) GetPkeyOk() (*string, bool)`

GetPkeyOk returns a tuple with the Pkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkey

`func (o *ConnectorSshConfigAllOf) SetPkey(v string)`

SetPkey sets Pkey field to given value.

### HasPkey

`func (o *ConnectorSshConfigAllOf) HasPkey() bool`

HasPkey returns a boolean if a field has been set.

### GetPkeyString

`func (o *ConnectorSshConfigAllOf) GetPkeyString() string`

GetPkeyString returns the PkeyString field if non-nil, zero value otherwise.

### GetPkeyStringOk

`func (o *ConnectorSshConfigAllOf) GetPkeyStringOk() (*string, bool)`

GetPkeyStringOk returns a tuple with the PkeyString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkeyString

`func (o *ConnectorSshConfigAllOf) SetPkeyString(v string)`

SetPkeyString sets PkeyString field to given value.

### HasPkeyString

`func (o *ConnectorSshConfigAllOf) HasPkeyString() bool`

HasPkeyString returns a boolean if a field has been set.

### GetTarget

`func (o *ConnectorSshConfigAllOf) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ConnectorSshConfigAllOf) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ConnectorSshConfigAllOf) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *ConnectorSshConfigAllOf) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetUser

`func (o *ConnectorSshConfigAllOf) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ConnectorSshConfigAllOf) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ConnectorSshConfigAllOf) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *ConnectorSshConfigAllOf) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


