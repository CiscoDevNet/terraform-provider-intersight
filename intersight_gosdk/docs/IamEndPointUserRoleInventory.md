# IamEndPointUserRoleInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.EndPointUserRoleInventory"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.EndPointUserRoleInventory"]
**ChangePassword** | Pointer to **bool** | Denotes whether password has changed. | [optional] [readonly] 
**Enabled** | Pointer to **bool** | Enables the user account on the endpoint. | [optional] [readonly] 
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] [default to false]
**Password** | Pointer to **string** | The password must have a minimum of 8 and a maximum of 127 characters. For servers with IPMI user role enabled, the maximum length is limited to 20 characters. When strong password is enabled, must satisfy below requirements: A. The password must not contain the User&#39;s Name. B. The password must contain characters from three of the following four categories. 1) English uppercase characters (A through Z). 2) English lowercase characters (a through z). 3) Base 10 digits (0 through 9). 4) Non-alphabetic characters (! , @, #, $, %, ^, &amp;, *, -, _, +, &#x3D;). | [optional] 
**EndPointRole** | Pointer to [**[]IamEndPointRoleRelationship**](IamEndPointRoleRelationship.md) | An array of relationships to iamEndPointRole resources. | [optional] [readonly] 
**EndPointUser** | Pointer to [**NullableIamEndPointUserInventoryRelationship**](IamEndPointUserInventoryRelationship.md) |  | [optional] 
**EndPointUserPolicy** | Pointer to [**NullableIamEndPointUserPolicyInventoryRelationship**](IamEndPointUserPolicyInventoryRelationship.md) |  | [optional] 

## Methods

### NewIamEndPointUserRoleInventory

`func NewIamEndPointUserRoleInventory(classId string, objectType string, ) *IamEndPointUserRoleInventory`

NewIamEndPointUserRoleInventory instantiates a new IamEndPointUserRoleInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamEndPointUserRoleInventoryWithDefaults

`func NewIamEndPointUserRoleInventoryWithDefaults() *IamEndPointUserRoleInventory`

NewIamEndPointUserRoleInventoryWithDefaults instantiates a new IamEndPointUserRoleInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamEndPointUserRoleInventory) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamEndPointUserRoleInventory) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamEndPointUserRoleInventory) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamEndPointUserRoleInventory) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamEndPointUserRoleInventory) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamEndPointUserRoleInventory) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetChangePassword

`func (o *IamEndPointUserRoleInventory) GetChangePassword() bool`

GetChangePassword returns the ChangePassword field if non-nil, zero value otherwise.

### GetChangePasswordOk

`func (o *IamEndPointUserRoleInventory) GetChangePasswordOk() (*bool, bool)`

GetChangePasswordOk returns a tuple with the ChangePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangePassword

`func (o *IamEndPointUserRoleInventory) SetChangePassword(v bool)`

SetChangePassword sets ChangePassword field to given value.

### HasChangePassword

`func (o *IamEndPointUserRoleInventory) HasChangePassword() bool`

HasChangePassword returns a boolean if a field has been set.

### GetEnabled

`func (o *IamEndPointUserRoleInventory) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IamEndPointUserRoleInventory) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IamEndPointUserRoleInventory) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IamEndPointUserRoleInventory) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIsPasswordSet

`func (o *IamEndPointUserRoleInventory) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *IamEndPointUserRoleInventory) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *IamEndPointUserRoleInventory) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *IamEndPointUserRoleInventory) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetPassword

`func (o *IamEndPointUserRoleInventory) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *IamEndPointUserRoleInventory) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *IamEndPointUserRoleInventory) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *IamEndPointUserRoleInventory) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetEndPointRole

`func (o *IamEndPointUserRoleInventory) GetEndPointRole() []IamEndPointRoleRelationship`

GetEndPointRole returns the EndPointRole field if non-nil, zero value otherwise.

### GetEndPointRoleOk

`func (o *IamEndPointUserRoleInventory) GetEndPointRoleOk() (*[]IamEndPointRoleRelationship, bool)`

GetEndPointRoleOk returns a tuple with the EndPointRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPointRole

`func (o *IamEndPointUserRoleInventory) SetEndPointRole(v []IamEndPointRoleRelationship)`

SetEndPointRole sets EndPointRole field to given value.

### HasEndPointRole

`func (o *IamEndPointUserRoleInventory) HasEndPointRole() bool`

HasEndPointRole returns a boolean if a field has been set.

### SetEndPointRoleNil

`func (o *IamEndPointUserRoleInventory) SetEndPointRoleNil(b bool)`

 SetEndPointRoleNil sets the value for EndPointRole to be an explicit nil

### UnsetEndPointRole
`func (o *IamEndPointUserRoleInventory) UnsetEndPointRole()`

UnsetEndPointRole ensures that no value is present for EndPointRole, not even an explicit nil
### GetEndPointUser

`func (o *IamEndPointUserRoleInventory) GetEndPointUser() IamEndPointUserInventoryRelationship`

GetEndPointUser returns the EndPointUser field if non-nil, zero value otherwise.

### GetEndPointUserOk

`func (o *IamEndPointUserRoleInventory) GetEndPointUserOk() (*IamEndPointUserInventoryRelationship, bool)`

GetEndPointUserOk returns a tuple with the EndPointUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPointUser

`func (o *IamEndPointUserRoleInventory) SetEndPointUser(v IamEndPointUserInventoryRelationship)`

SetEndPointUser sets EndPointUser field to given value.

### HasEndPointUser

`func (o *IamEndPointUserRoleInventory) HasEndPointUser() bool`

HasEndPointUser returns a boolean if a field has been set.

### SetEndPointUserNil

`func (o *IamEndPointUserRoleInventory) SetEndPointUserNil(b bool)`

 SetEndPointUserNil sets the value for EndPointUser to be an explicit nil

### UnsetEndPointUser
`func (o *IamEndPointUserRoleInventory) UnsetEndPointUser()`

UnsetEndPointUser ensures that no value is present for EndPointUser, not even an explicit nil
### GetEndPointUserPolicy

`func (o *IamEndPointUserRoleInventory) GetEndPointUserPolicy() IamEndPointUserPolicyInventoryRelationship`

GetEndPointUserPolicy returns the EndPointUserPolicy field if non-nil, zero value otherwise.

### GetEndPointUserPolicyOk

`func (o *IamEndPointUserRoleInventory) GetEndPointUserPolicyOk() (*IamEndPointUserPolicyInventoryRelationship, bool)`

GetEndPointUserPolicyOk returns a tuple with the EndPointUserPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPointUserPolicy

`func (o *IamEndPointUserRoleInventory) SetEndPointUserPolicy(v IamEndPointUserPolicyInventoryRelationship)`

SetEndPointUserPolicy sets EndPointUserPolicy field to given value.

### HasEndPointUserPolicy

`func (o *IamEndPointUserRoleInventory) HasEndPointUserPolicy() bool`

HasEndPointUserPolicy returns a boolean if a field has been set.

### SetEndPointUserPolicyNil

`func (o *IamEndPointUserRoleInventory) SetEndPointUserPolicyNil(b bool)`

 SetEndPointUserPolicyNil sets the value for EndPointUserPolicy to be an explicit nil

### UnsetEndPointUserPolicy
`func (o *IamEndPointUserRoleInventory) UnsetEndPointUserPolicy()`

UnsetEndPointUserPolicy ensures that no value is present for EndPointUserPolicy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


