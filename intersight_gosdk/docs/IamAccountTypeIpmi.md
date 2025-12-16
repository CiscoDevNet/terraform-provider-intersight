# IamAccountTypeIpmi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.AccountTypeIpmi"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.AccountTypeIpmi"]
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] [default to false]
**Password** | Pointer to **string** | The password must have a minimum of 8 and a maximum of 127 characters. For servers with IPMI user role enabled, the maximum length is limited to 20 characters. When strong password is enabled, must satisfy below requirements: A. The password must not contain the User&#39;s Name. B. The password must contain characters from three of the following four categories. 1) English uppercase characters (A through Z). 2) English lowercase characters (a through z). 3) Base 10 digits (0 through 9). 4) Non-alphabetic characters (! , @, #, $, %, ^, &amp;, *, -, _, +, &#x3D;). | [optional] 
**UseDiffIpmiPassword** | Pointer to **bool** | Use Different IPMI Password. If false then same Local User Password will be used for IPMI. | [optional] [default to false]

## Methods

### NewIamAccountTypeIpmi

`func NewIamAccountTypeIpmi(classId string, objectType string, ) *IamAccountTypeIpmi`

NewIamAccountTypeIpmi instantiates a new IamAccountTypeIpmi object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamAccountTypeIpmiWithDefaults

`func NewIamAccountTypeIpmiWithDefaults() *IamAccountTypeIpmi`

NewIamAccountTypeIpmiWithDefaults instantiates a new IamAccountTypeIpmi object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamAccountTypeIpmi) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamAccountTypeIpmi) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamAccountTypeIpmi) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamAccountTypeIpmi) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamAccountTypeIpmi) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamAccountTypeIpmi) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsPasswordSet

`func (o *IamAccountTypeIpmi) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *IamAccountTypeIpmi) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *IamAccountTypeIpmi) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *IamAccountTypeIpmi) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetPassword

`func (o *IamAccountTypeIpmi) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *IamAccountTypeIpmi) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *IamAccountTypeIpmi) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *IamAccountTypeIpmi) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUseDiffIpmiPassword

`func (o *IamAccountTypeIpmi) GetUseDiffIpmiPassword() bool`

GetUseDiffIpmiPassword returns the UseDiffIpmiPassword field if non-nil, zero value otherwise.

### GetUseDiffIpmiPasswordOk

`func (o *IamAccountTypeIpmi) GetUseDiffIpmiPasswordOk() (*bool, bool)`

GetUseDiffIpmiPasswordOk returns a tuple with the UseDiffIpmiPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDiffIpmiPassword

`func (o *IamAccountTypeIpmi) SetUseDiffIpmiPassword(v bool)`

SetUseDiffIpmiPassword sets UseDiffIpmiPassword field to given value.

### HasUseDiffIpmiPassword

`func (o *IamAccountTypeIpmi) HasUseDiffIpmiPassword() bool`

HasUseDiffIpmiPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


