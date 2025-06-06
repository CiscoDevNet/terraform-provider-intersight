# IamLdapProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.LdapProvider"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.LdapProvider"]
**Port** | Pointer to **int64** | LDAP Server Port for connection establishment. | [optional] [default to 389]
**Server** | Pointer to **string** | LDAP Server Address, can be IP address or hostname. | [optional] 
**Vendor** | Pointer to **string** | LDAP server vendor type used for authentication. * &#x60;OpenLDAP&#x60; - Open source LDAP server for remote authentication. * &#x60;MSAD&#x60; - Microsoft active directory for remote authentication. | [optional] [default to "OpenLDAP"]
**LdapPolicy** | Pointer to [**NullableIamLdapPolicyRelationship**](IamLdapPolicyRelationship.md) |  | [optional] 

## Methods

### NewIamLdapProvider

`func NewIamLdapProvider(classId string, objectType string, ) *IamLdapProvider`

NewIamLdapProvider instantiates a new IamLdapProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamLdapProviderWithDefaults

`func NewIamLdapProviderWithDefaults() *IamLdapProvider`

NewIamLdapProviderWithDefaults instantiates a new IamLdapProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamLdapProvider) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamLdapProvider) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamLdapProvider) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamLdapProvider) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamLdapProvider) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamLdapProvider) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPort

`func (o *IamLdapProvider) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *IamLdapProvider) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *IamLdapProvider) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *IamLdapProvider) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetServer

`func (o *IamLdapProvider) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *IamLdapProvider) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *IamLdapProvider) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *IamLdapProvider) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetVendor

`func (o *IamLdapProvider) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *IamLdapProvider) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *IamLdapProvider) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *IamLdapProvider) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetLdapPolicy

`func (o *IamLdapProvider) GetLdapPolicy() IamLdapPolicyRelationship`

GetLdapPolicy returns the LdapPolicy field if non-nil, zero value otherwise.

### GetLdapPolicyOk

`func (o *IamLdapProvider) GetLdapPolicyOk() (*IamLdapPolicyRelationship, bool)`

GetLdapPolicyOk returns a tuple with the LdapPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapPolicy

`func (o *IamLdapProvider) SetLdapPolicy(v IamLdapPolicyRelationship)`

SetLdapPolicy sets LdapPolicy field to given value.

### HasLdapPolicy

`func (o *IamLdapProvider) HasLdapPolicy() bool`

HasLdapPolicy returns a boolean if a field has been set.

### SetLdapPolicyNil

`func (o *IamLdapProvider) SetLdapPolicyNil(b bool)`

 SetLdapPolicyNil sets the value for LdapPolicy to be an explicit nil

### UnsetLdapPolicy
`func (o *IamLdapProvider) UnsetLdapPolicy()`

UnsetLdapPolicy ensures that no value is present for LdapPolicy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


