# ApplianceDiagSettingAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] 
**Message** | Pointer to **string** | Status message of the password change operation. | [optional] 
**Password** | Pointer to **string** | Password of the Intersight Appliance&#39;s OS diagnostic user account. | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 

## Methods

### NewApplianceDiagSettingAllOf

`func NewApplianceDiagSettingAllOf() *ApplianceDiagSettingAllOf`

NewApplianceDiagSettingAllOf instantiates a new ApplianceDiagSettingAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceDiagSettingAllOfWithDefaults

`func NewApplianceDiagSettingAllOfWithDefaults() *ApplianceDiagSettingAllOf`

NewApplianceDiagSettingAllOfWithDefaults instantiates a new ApplianceDiagSettingAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsPasswordSet

`func (o *ApplianceDiagSettingAllOf) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *ApplianceDiagSettingAllOf) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *ApplianceDiagSettingAllOf) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *ApplianceDiagSettingAllOf) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetMessage

`func (o *ApplianceDiagSettingAllOf) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApplianceDiagSettingAllOf) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApplianceDiagSettingAllOf) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApplianceDiagSettingAllOf) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPassword

`func (o *ApplianceDiagSettingAllOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ApplianceDiagSettingAllOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ApplianceDiagSettingAllOf) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ApplianceDiagSettingAllOf) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetAccount

`func (o *ApplianceDiagSettingAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ApplianceDiagSettingAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ApplianceDiagSettingAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ApplianceDiagSettingAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


