# IamIpAccessManagementAllOf

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

### NewIamIpAccessManagementAllOf

`func NewIamIpAccessManagementAllOf(classId string, objectType string, ) *IamIpAccessManagementAllOf`

NewIamIpAccessManagementAllOf instantiates a new IamIpAccessManagementAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamIpAccessManagementAllOfWithDefaults

`func NewIamIpAccessManagementAllOfWithDefaults() *IamIpAccessManagementAllOf`

NewIamIpAccessManagementAllOfWithDefaults instantiates a new IamIpAccessManagementAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamIpAccessManagementAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamIpAccessManagementAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamIpAccessManagementAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamIpAccessManagementAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamIpAccessManagementAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamIpAccessManagementAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetEnable

`func (o *IamIpAccessManagementAllOf) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *IamIpAccessManagementAllOf) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *IamIpAccessManagementAllOf) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *IamIpAccessManagementAllOf) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetLastRecoveryTime

`func (o *IamIpAccessManagementAllOf) GetLastRecoveryTime() time.Time`

GetLastRecoveryTime returns the LastRecoveryTime field if non-nil, zero value otherwise.

### GetLastRecoveryTimeOk

`func (o *IamIpAccessManagementAllOf) GetLastRecoveryTimeOk() (*time.Time, bool)`

GetLastRecoveryTimeOk returns a tuple with the LastRecoveryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRecoveryTime

`func (o *IamIpAccessManagementAllOf) SetLastRecoveryTime(v time.Time)`

SetLastRecoveryTime sets LastRecoveryTime field to given value.

### HasLastRecoveryTime

`func (o *IamIpAccessManagementAllOf) HasLastRecoveryTime() bool`

HasLastRecoveryTime returns a boolean if a field has been set.

### GetHolder

`func (o *IamIpAccessManagementAllOf) GetHolder() IamSecurityHolderRelationship`

GetHolder returns the Holder field if non-nil, zero value otherwise.

### GetHolderOk

`func (o *IamIpAccessManagementAllOf) GetHolderOk() (*IamSecurityHolderRelationship, bool)`

GetHolderOk returns a tuple with the Holder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolder

`func (o *IamIpAccessManagementAllOf) SetHolder(v IamSecurityHolderRelationship)`

SetHolder sets Holder field to given value.

### HasHolder

`func (o *IamIpAccessManagementAllOf) HasHolder() bool`

HasHolder returns a boolean if a field has been set.

### GetIpAddresses

`func (o *IamIpAccessManagementAllOf) GetIpAddresses() []IamIpAddressRelationship`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *IamIpAccessManagementAllOf) GetIpAddressesOk() (*[]IamIpAddressRelationship, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *IamIpAccessManagementAllOf) SetIpAddresses(v []IamIpAddressRelationship)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *IamIpAccessManagementAllOf) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### SetIpAddressesNil

`func (o *IamIpAccessManagementAllOf) SetIpAddressesNil(b bool)`

 SetIpAddressesNil sets the value for IpAddresses to be an explicit nil

### UnsetIpAddresses
`func (o *IamIpAccessManagementAllOf) UnsetIpAddresses()`

UnsetIpAddresses ensures that no value is present for IpAddresses, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


