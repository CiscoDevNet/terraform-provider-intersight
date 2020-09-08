# IamIpAccessManagement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** | Flag stores the state of IP address based access management. Access management is enabled when it&#39;s true. | [optional] 
**LastRecoveryTime** | Pointer to [**time.Time**](time.Time.md) | The access to account gets locked out if wrong IP addresses are configured. Account Administrators have privilege to unblock the account. It stores the time when the account was last recovered from lock out. | [optional] [readonly] 
**Holder** | Pointer to [**IamSecurityHolderRelationship**](iam.SecurityHolder.Relationship.md) |  | [optional] 
**IpAddresses** | Pointer to [**[]IamIpAddressRelationship**](iam.IpAddress.Relationship.md) | An array of relationships to iamIpAddress resources. | [optional] [readonly] 

## Methods

### NewIamIpAccessManagement

`func NewIamIpAccessManagement() *IamIpAccessManagement`

NewIamIpAccessManagement instantiates a new IamIpAccessManagement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamIpAccessManagementWithDefaults

`func NewIamIpAccessManagementWithDefaults() *IamIpAccessManagement`

NewIamIpAccessManagementWithDefaults instantiates a new IamIpAccessManagement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


