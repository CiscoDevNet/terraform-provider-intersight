# IamLdapProviderInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.LdapProviderInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.LdapProviderInventory"]
**Port** | Pointer to **int64** | LDAP Server Port for connection establishment. | [optional] [readonly] [default to 389]
**Server** | Pointer to **string** | LDAP Server Address, such as an IP address or hostname. | [optional] [readonly] 
**Vendor** | Pointer to **string** | LDAP server vendor type used for authentication. * &#x60;OpenLDAP&#x60; - Open source LDAP server for remote authentication. * &#x60;MSAD&#x60; - Microsoft active directory for remote authentication. | [optional] [readonly] [default to "OpenLDAP"]
**LdapPolicy** | Pointer to [**NullableIamLdapPolicyInventoryRelationship**](IamLdapPolicyInventoryRelationship.md) |  | [optional] 

## Methods

### NewIamLdapProviderInventory

`func NewIamLdapProviderInventory(classId string, objectType string, ) *IamLdapProviderInventory`

NewIamLdapProviderInventory instantiates a new IamLdapProviderInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamLdapProviderInventoryWithDefaults

`func NewIamLdapProviderInventoryWithDefaults() *IamLdapProviderInventory`

NewIamLdapProviderInventoryWithDefaults instantiates a new IamLdapProviderInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamLdapProviderInventory) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamLdapProviderInventory) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamLdapProviderInventory) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamLdapProviderInventory) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamLdapProviderInventory) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamLdapProviderInventory) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetPort

`func (o *IamLdapProviderInventory) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *IamLdapProviderInventory) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *IamLdapProviderInventory) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *IamLdapProviderInventory) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetServer

`func (o *IamLdapProviderInventory) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *IamLdapProviderInventory) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *IamLdapProviderInventory) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *IamLdapProviderInventory) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetVendor

`func (o *IamLdapProviderInventory) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *IamLdapProviderInventory) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *IamLdapProviderInventory) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *IamLdapProviderInventory) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetLdapPolicy

`func (o *IamLdapProviderInventory) GetLdapPolicy() IamLdapPolicyInventoryRelationship`

GetLdapPolicy returns the LdapPolicy field if non-nil, zero value otherwise.

### GetLdapPolicyOk

`func (o *IamLdapProviderInventory) GetLdapPolicyOk() (*IamLdapPolicyInventoryRelationship, bool)`

GetLdapPolicyOk returns a tuple with the LdapPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapPolicy

`func (o *IamLdapProviderInventory) SetLdapPolicy(v IamLdapPolicyInventoryRelationship)`

SetLdapPolicy sets LdapPolicy field to given value.

### HasLdapPolicy

`func (o *IamLdapProviderInventory) HasLdapPolicy() bool`

HasLdapPolicy returns a boolean if a field has been set.

### SetLdapPolicyNil

`func (o *IamLdapProviderInventory) SetLdapPolicyNil(b bool)`

 SetLdapPolicyNil sets the value for LdapPolicy to be an explicit nil

### UnsetLdapPolicy
`func (o *IamLdapProviderInventory) UnsetLdapPolicy()`

UnsetLdapPolicy ensures that no value is present for LdapPolicy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


