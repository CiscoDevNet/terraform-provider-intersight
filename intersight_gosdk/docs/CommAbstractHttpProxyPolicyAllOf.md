# CommAbstractHttpProxyPolicyAllOf

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

### NewCommAbstractHttpProxyPolicyAllOf

`func NewCommAbstractHttpProxyPolicyAllOf(classId string, objectType string, ) *CommAbstractHttpProxyPolicyAllOf`

NewCommAbstractHttpProxyPolicyAllOf instantiates a new CommAbstractHttpProxyPolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommAbstractHttpProxyPolicyAllOfWithDefaults

`func NewCommAbstractHttpProxyPolicyAllOfWithDefaults() *CommAbstractHttpProxyPolicyAllOf`

NewCommAbstractHttpProxyPolicyAllOfWithDefaults instantiates a new CommAbstractHttpProxyPolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CommAbstractHttpProxyPolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CommAbstractHttpProxyPolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CommAbstractHttpProxyPolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CommAbstractHttpProxyPolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CommAbstractHttpProxyPolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CommAbstractHttpProxyPolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetHostname

`func (o *CommAbstractHttpProxyPolicyAllOf) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *CommAbstractHttpProxyPolicyAllOf) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *CommAbstractHttpProxyPolicyAllOf) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *CommAbstractHttpProxyPolicyAllOf) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIsPasswordSet

`func (o *CommAbstractHttpProxyPolicyAllOf) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *CommAbstractHttpProxyPolicyAllOf) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *CommAbstractHttpProxyPolicyAllOf) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *CommAbstractHttpProxyPolicyAllOf) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetPassword

`func (o *CommAbstractHttpProxyPolicyAllOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CommAbstractHttpProxyPolicyAllOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CommAbstractHttpProxyPolicyAllOf) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CommAbstractHttpProxyPolicyAllOf) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPort

`func (o *CommAbstractHttpProxyPolicyAllOf) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CommAbstractHttpProxyPolicyAllOf) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CommAbstractHttpProxyPolicyAllOf) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *CommAbstractHttpProxyPolicyAllOf) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUsername

`func (o *CommAbstractHttpProxyPolicyAllOf) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CommAbstractHttpProxyPolicyAllOf) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CommAbstractHttpProxyPolicyAllOf) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CommAbstractHttpProxyPolicyAllOf) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetOrganization

`func (o *CommAbstractHttpProxyPolicyAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *CommAbstractHttpProxyPolicyAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *CommAbstractHttpProxyPolicyAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *CommAbstractHttpProxyPolicyAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


