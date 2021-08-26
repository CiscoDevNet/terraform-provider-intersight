# CommAbstractHttpProxyPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "comm.HttpProxyPolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "comm.HttpProxyPolicy"]
**Hostname** | Pointer to **string** | HTTP Proxy server FQDN or IP. | [optional] 
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] [default to false]
**Password** | Pointer to **string** | The password for the HTTP Proxy. | [optional] 
**Port** | Pointer to **int64** | The HTTP Proxy port number. The port number of the HTTP proxy must be between 1 and 65535, inclusive. | [optional] 
**Username** | Pointer to **string** | The username for the HTTP Proxy. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewCommAbstractHttpProxyPolicy

`func NewCommAbstractHttpProxyPolicy(classId string, objectType string, ) *CommAbstractHttpProxyPolicy`

NewCommAbstractHttpProxyPolicy instantiates a new CommAbstractHttpProxyPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommAbstractHttpProxyPolicyWithDefaults

`func NewCommAbstractHttpProxyPolicyWithDefaults() *CommAbstractHttpProxyPolicy`

NewCommAbstractHttpProxyPolicyWithDefaults instantiates a new CommAbstractHttpProxyPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CommAbstractHttpProxyPolicy) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CommAbstractHttpProxyPolicy) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CommAbstractHttpProxyPolicy) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CommAbstractHttpProxyPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CommAbstractHttpProxyPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CommAbstractHttpProxyPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetHostname

`func (o *CommAbstractHttpProxyPolicy) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *CommAbstractHttpProxyPolicy) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *CommAbstractHttpProxyPolicy) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *CommAbstractHttpProxyPolicy) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIsPasswordSet

`func (o *CommAbstractHttpProxyPolicy) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *CommAbstractHttpProxyPolicy) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *CommAbstractHttpProxyPolicy) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *CommAbstractHttpProxyPolicy) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetPassword

`func (o *CommAbstractHttpProxyPolicy) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CommAbstractHttpProxyPolicy) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CommAbstractHttpProxyPolicy) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CommAbstractHttpProxyPolicy) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPort

`func (o *CommAbstractHttpProxyPolicy) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CommAbstractHttpProxyPolicy) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CommAbstractHttpProxyPolicy) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *CommAbstractHttpProxyPolicy) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUsername

`func (o *CommAbstractHttpProxyPolicy) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CommAbstractHttpProxyPolicy) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CommAbstractHttpProxyPolicy) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CommAbstractHttpProxyPolicy) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetOrganization

`func (o *CommAbstractHttpProxyPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *CommAbstractHttpProxyPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *CommAbstractHttpProxyPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *CommAbstractHttpProxyPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


