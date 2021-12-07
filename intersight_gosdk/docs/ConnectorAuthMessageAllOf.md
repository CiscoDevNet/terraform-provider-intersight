# ConnectorAuthMessageAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**RemoteUserLocale** | Pointer to **string** | The platform locale to assign user. A locale defines one or more organizations (domains) the user is allowed access, and access is limited to the organizations specified in the locale. | [optional] 
**RemoteUserName** | Pointer to **string** | The user name passed to the platform for use in platform audit logs. | [optional] 
**RemoteUserRoles** | Pointer to **string** | The list of roles to pass to the platform to validate the action against. | [optional] 
**RemoteUserSessionId** | Pointer to **string** | The session Id passed to the platform for use in platforms auditing. | [optional] 

## Methods

### NewConnectorAuthMessageAllOf

`func NewConnectorAuthMessageAllOf(classId string, objectType string, ) *ConnectorAuthMessageAllOf`

NewConnectorAuthMessageAllOf instantiates a new ConnectorAuthMessageAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorAuthMessageAllOfWithDefaults

`func NewConnectorAuthMessageAllOfWithDefaults() *ConnectorAuthMessageAllOf`

NewConnectorAuthMessageAllOfWithDefaults instantiates a new ConnectorAuthMessageAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConnectorAuthMessageAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConnectorAuthMessageAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConnectorAuthMessageAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConnectorAuthMessageAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConnectorAuthMessageAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConnectorAuthMessageAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetRemoteUserLocale

`func (o *ConnectorAuthMessageAllOf) GetRemoteUserLocale() string`

GetRemoteUserLocale returns the RemoteUserLocale field if non-nil, zero value otherwise.

### GetRemoteUserLocaleOk

`func (o *ConnectorAuthMessageAllOf) GetRemoteUserLocaleOk() (*string, bool)`

GetRemoteUserLocaleOk returns a tuple with the RemoteUserLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteUserLocale

`func (o *ConnectorAuthMessageAllOf) SetRemoteUserLocale(v string)`

SetRemoteUserLocale sets RemoteUserLocale field to given value.

### HasRemoteUserLocale

`func (o *ConnectorAuthMessageAllOf) HasRemoteUserLocale() bool`

HasRemoteUserLocale returns a boolean if a field has been set.

### GetRemoteUserName

`func (o *ConnectorAuthMessageAllOf) GetRemoteUserName() string`

GetRemoteUserName returns the RemoteUserName field if non-nil, zero value otherwise.

### GetRemoteUserNameOk

`func (o *ConnectorAuthMessageAllOf) GetRemoteUserNameOk() (*string, bool)`

GetRemoteUserNameOk returns a tuple with the RemoteUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteUserName

`func (o *ConnectorAuthMessageAllOf) SetRemoteUserName(v string)`

SetRemoteUserName sets RemoteUserName field to given value.

### HasRemoteUserName

`func (o *ConnectorAuthMessageAllOf) HasRemoteUserName() bool`

HasRemoteUserName returns a boolean if a field has been set.

### GetRemoteUserRoles

`func (o *ConnectorAuthMessageAllOf) GetRemoteUserRoles() string`

GetRemoteUserRoles returns the RemoteUserRoles field if non-nil, zero value otherwise.

### GetRemoteUserRolesOk

`func (o *ConnectorAuthMessageAllOf) GetRemoteUserRolesOk() (*string, bool)`

GetRemoteUserRolesOk returns a tuple with the RemoteUserRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteUserRoles

`func (o *ConnectorAuthMessageAllOf) SetRemoteUserRoles(v string)`

SetRemoteUserRoles sets RemoteUserRoles field to given value.

### HasRemoteUserRoles

`func (o *ConnectorAuthMessageAllOf) HasRemoteUserRoles() bool`

HasRemoteUserRoles returns a boolean if a field has been set.

### GetRemoteUserSessionId

`func (o *ConnectorAuthMessageAllOf) GetRemoteUserSessionId() string`

GetRemoteUserSessionId returns the RemoteUserSessionId field if non-nil, zero value otherwise.

### GetRemoteUserSessionIdOk

`func (o *ConnectorAuthMessageAllOf) GetRemoteUserSessionIdOk() (*string, bool)`

GetRemoteUserSessionIdOk returns a tuple with the RemoteUserSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteUserSessionId

`func (o *ConnectorAuthMessageAllOf) SetRemoteUserSessionId(v string)`

SetRemoteUserSessionId sets RemoteUserSessionId field to given value.

### HasRemoteUserSessionId

`func (o *ConnectorAuthMessageAllOf) HasRemoteUserSessionId() bool`

HasRemoteUserSessionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


