# ConnectorSshConfigAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** | Password to use in the connection credentials (If empty the private key will be used). | [optional] 
**Pkey** | Pointer to **string** | The private key to use in the connection credentials (Optional if password is given). | [optional] 
**Target** | Pointer to **string** | The remote server to connect to. | [optional] 
**User** | Pointer to **string** | Username for the remote connection. | [optional] 

## Methods

### NewConnectorSshConfigAllOf

`func NewConnectorSshConfigAllOf() *ConnectorSshConfigAllOf`

NewConnectorSshConfigAllOf instantiates a new ConnectorSshConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorSshConfigAllOfWithDefaults

`func NewConnectorSshConfigAllOfWithDefaults() *ConnectorSshConfigAllOf`

NewConnectorSshConfigAllOfWithDefaults instantiates a new ConnectorSshConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


