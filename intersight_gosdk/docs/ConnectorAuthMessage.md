# ConnectorAuthMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteUserLocale** | Pointer to **string** | The platform locale to assign user. A locale defines one or more organizations (domains) the user is allowed access, and access is limited to the organizations specified in the locale. | [optional] 
**RemoteUserName** | Pointer to **string** | The user name passed to the platform for use in platform audit logs. | [optional] 
**RemoteUserRoles** | Pointer to **string** | The list of roles to pass to the platform to validate the action against. | [optional] 
**RemoteUserSessionId** | Pointer to **string** | The session Id passed to the platform for use in platforms auditing. | [optional] 

## Methods

### NewConnectorAuthMessage

`func NewConnectorAuthMessage() *ConnectorAuthMessage`

NewConnectorAuthMessage instantiates a new ConnectorAuthMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorAuthMessageWithDefaults

`func NewConnectorAuthMessageWithDefaults() *ConnectorAuthMessage`

NewConnectorAuthMessageWithDefaults instantiates a new ConnectorAuthMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteUserLocale

`func (o *ConnectorAuthMessage) GetRemoteUserLocale() string`

GetRemoteUserLocale returns the RemoteUserLocale field if non-nil, zero value otherwise.

### GetRemoteUserLocaleOk

`func (o *ConnectorAuthMessage) GetRemoteUserLocaleOk() (*string, bool)`

GetRemoteUserLocaleOk returns a tuple with the RemoteUserLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteUserLocale

`func (o *ConnectorAuthMessage) SetRemoteUserLocale(v string)`

SetRemoteUserLocale sets RemoteUserLocale field to given value.

### HasRemoteUserLocale

`func (o *ConnectorAuthMessage) HasRemoteUserLocale() bool`

HasRemoteUserLocale returns a boolean if a field has been set.

### GetRemoteUserName

`func (o *ConnectorAuthMessage) GetRemoteUserName() string`

GetRemoteUserName returns the RemoteUserName field if non-nil, zero value otherwise.

### GetRemoteUserNameOk

`func (o *ConnectorAuthMessage) GetRemoteUserNameOk() (*string, bool)`

GetRemoteUserNameOk returns a tuple with the RemoteUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteUserName

`func (o *ConnectorAuthMessage) SetRemoteUserName(v string)`

SetRemoteUserName sets RemoteUserName field to given value.

### HasRemoteUserName

`func (o *ConnectorAuthMessage) HasRemoteUserName() bool`

HasRemoteUserName returns a boolean if a field has been set.

### GetRemoteUserRoles

`func (o *ConnectorAuthMessage) GetRemoteUserRoles() string`

GetRemoteUserRoles returns the RemoteUserRoles field if non-nil, zero value otherwise.

### GetRemoteUserRolesOk

`func (o *ConnectorAuthMessage) GetRemoteUserRolesOk() (*string, bool)`

GetRemoteUserRolesOk returns a tuple with the RemoteUserRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteUserRoles

`func (o *ConnectorAuthMessage) SetRemoteUserRoles(v string)`

SetRemoteUserRoles sets RemoteUserRoles field to given value.

### HasRemoteUserRoles

`func (o *ConnectorAuthMessage) HasRemoteUserRoles() bool`

HasRemoteUserRoles returns a boolean if a field has been set.

### GetRemoteUserSessionId

`func (o *ConnectorAuthMessage) GetRemoteUserSessionId() string`

GetRemoteUserSessionId returns the RemoteUserSessionId field if non-nil, zero value otherwise.

### GetRemoteUserSessionIdOk

`func (o *ConnectorAuthMessage) GetRemoteUserSessionIdOk() (*string, bool)`

GetRemoteUserSessionIdOk returns a tuple with the RemoteUserSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteUserSessionId

`func (o *ConnectorAuthMessage) SetRemoteUserSessionId(v string)`

SetRemoteUserSessionId sets RemoteUserSessionId field to given value.

### HasRemoteUserSessionId

`func (o *ConnectorAuthMessage) HasRemoteUserSessionId() bool`

HasRemoteUserSessionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


