# IamIpAccessManagement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.IpAccessManagement"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.IpAccessManagement"]
**Enable** | Pointer to **bool** | Flag stores the state of IP address based access management. Access management is enabled when it&#39;s true. | [optional] 
**LastRecoveryTime** | Pointer to **time.Time** | The access to account gets locked out if wrong IP addresses are configured. Account Administrators have privilege to unblock the account. It stores the time when the account was last recovered from lock out. | [optional] [readonly] 
**Holder** | Pointer to [**IamSecurityHolderRelationship**](IamSecurityHolderRelationship.md) |  | [optional] 
**IpAddresses** | Pointer to [**[]IamIpAddressRelationship**](IamIpAddressRelationship.md) | An array of relationships to iamIpAddress resources. | [optional] [readonly] 

## Methods

### NewIamIpAccessManagement

`func NewIamIpAccessManagement(classId string, objectType string, ) *IamIpAccessManagement`

NewIamIpAccessManagement instantiates a new IamIpAccessManagement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamIpAccessManagementWithDefaults

`func NewIamIpAccessManagementWithDefaults() *IamIpAccessManagement`

NewIamIpAccessManagementWithDefaults instantiates a new IamIpAccessManagement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamIpAccessManagement) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamIpAccessManagement) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamIpAccessManagement) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamIpAccessManagement) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamIpAccessManagement) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamIpAccessManagement) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnable

`func (o *IamIpAccessManagement) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *IamIpAccessManagement) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *IamIpAccessManagement) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *IamIpAccessManagement) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetLastRecoveryTime

`func (o *IamIpAccessManagement) GetLastRecoveryTime() time.Time`

GetLastRecoveryTime returns the LastRecoveryTime field if non-nil, zero value otherwise.

### GetLastRecoveryTimeOk

`func (o *IamIpAccessManagement) GetLastRecoveryTimeOk() (*time.Time, bool)`

GetLastRecoveryTimeOk returns a tuple with the LastRecoveryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRecoveryTime

`func (o *IamIpAccessManagement) SetLastRecoveryTime(v time.Time)`

SetLastRecoveryTime sets LastRecoveryTime field to given value.

### HasLastRecoveryTime

`func (o *IamIpAccessManagement) HasLastRecoveryTime() bool`

HasLastRecoveryTime returns a boolean if a field has been set.

### GetHolder

`func (o *IamIpAccessManagement) GetHolder() IamSecurityHolderRelationship`

GetHolder returns the Holder field if non-nil, zero value otherwise.

### GetHolderOk

`func (o *IamIpAccessManagement) GetHolderOk() (*IamSecurityHolderRelationship, bool)`

GetHolderOk returns a tuple with the Holder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolder

`func (o *IamIpAccessManagement) SetHolder(v IamSecurityHolderRelationship)`

SetHolder sets Holder field to given value.

### HasHolder

`func (o *IamIpAccessManagement) HasHolder() bool`

HasHolder returns a boolean if a field has been set.

### GetIpAddresses

`func (o *IamIpAccessManagement) GetIpAddresses() []IamIpAddressRelationship`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *IamIpAccessManagement) GetIpAddressesOk() (*[]IamIpAddressRelationship, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *IamIpAccessManagement) SetIpAddresses(v []IamIpAddressRelationship)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *IamIpAccessManagement) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### SetIpAddressesNil

`func (o *IamIpAccessManagement) SetIpAddressesNil(b bool)`

 SetIpAddressesNil sets the value for IpAddresses to be an explicit nil

### UnsetIpAddresses
`func (o *IamIpAccessManagement) UnsetIpAddresses()`

UnsetIpAddresses ensures that no value is present for IpAddresses, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


