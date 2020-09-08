# IamLdapProviderAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **int64** | LDAP Server Port for connection establishment. | [optional] 
**Server** | Pointer to **string** | LDAP Server Address, can be IP address or hostname. | [optional] 
**LdapPolicy** | Pointer to [**IamLdapPolicyRelationship**](iam.LdapPolicy.Relationship.md) |  | [optional] 

## Methods

### NewIamLdapProviderAllOf

`func NewIamLdapProviderAllOf() *IamLdapProviderAllOf`

NewIamLdapProviderAllOf instantiates a new IamLdapProviderAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamLdapProviderAllOfWithDefaults

`func NewIamLdapProviderAllOfWithDefaults() *IamLdapProviderAllOf`

NewIamLdapProviderAllOfWithDefaults instantiates a new IamLdapProviderAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *IamLdapProviderAllOf) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *IamLdapProviderAllOf) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *IamLdapProviderAllOf) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *IamLdapProviderAllOf) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetServer

`func (o *IamLdapProviderAllOf) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *IamLdapProviderAllOf) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *IamLdapProviderAllOf) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *IamLdapProviderAllOf) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetLdapPolicy

`func (o *IamLdapProviderAllOf) GetLdapPolicy() IamLdapPolicyRelationship`

GetLdapPolicy returns the LdapPolicy field if non-nil, zero value otherwise.

### GetLdapPolicyOk

`func (o *IamLdapProviderAllOf) GetLdapPolicyOk() (*IamLdapPolicyRelationship, bool)`

GetLdapPolicyOk returns a tuple with the LdapPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapPolicy

`func (o *IamLdapProviderAllOf) SetLdapPolicy(v IamLdapPolicyRelationship)`

SetLdapPolicy sets LdapPolicy field to given value.

### HasLdapPolicy

`func (o *IamLdapProviderAllOf) HasLdapPolicy() bool`

HasLdapPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


