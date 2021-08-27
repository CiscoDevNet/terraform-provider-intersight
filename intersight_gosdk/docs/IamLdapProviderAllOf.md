# IamLdapProviderAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.LdapProvider"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.LdapProvider"]
**Port** | Pointer to **int64** | LDAP Server Port for connection establishment. | [optional] [default to 389]
**Server** | Pointer to **string** | LDAP Server Address, can be IP address or hostname. | [optional] 
**LdapPolicy** | Pointer to [**IamLdapPolicyRelationship**](IamLdapPolicyRelationship.md) |  | [optional] 

## Methods

### NewIamLdapProviderAllOf

`func NewIamLdapProviderAllOf(classId string, objectType string, ) *IamLdapProviderAllOf`

NewIamLdapProviderAllOf instantiates a new IamLdapProviderAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamLdapProviderAllOfWithDefaults

`func NewIamLdapProviderAllOfWithDefaults() *IamLdapProviderAllOf`

NewIamLdapProviderAllOfWithDefaults instantiates a new IamLdapProviderAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamLdapProviderAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamLdapProviderAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamLdapProviderAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamLdapProviderAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamLdapProviderAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamLdapProviderAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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


