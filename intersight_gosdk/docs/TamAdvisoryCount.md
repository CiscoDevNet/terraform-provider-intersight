# TamAdvisoryCount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvisoryCount** | Pointer to **int64** | Total number of advisories affecting the account. | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 

## Methods

### NewTamAdvisoryCount

`func NewTamAdvisoryCount() *TamAdvisoryCount`

NewTamAdvisoryCount instantiates a new TamAdvisoryCount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTamAdvisoryCountWithDefaults

`func NewTamAdvisoryCountWithDefaults() *TamAdvisoryCount`

NewTamAdvisoryCountWithDefaults instantiates a new TamAdvisoryCount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvisoryCount

`func (o *TamAdvisoryCount) GetAdvisoryCount() int64`

GetAdvisoryCount returns the AdvisoryCount field if non-nil, zero value otherwise.

### GetAdvisoryCountOk

`func (o *TamAdvisoryCount) GetAdvisoryCountOk() (*int64, bool)`

GetAdvisoryCountOk returns a tuple with the AdvisoryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvisoryCount

`func (o *TamAdvisoryCount) SetAdvisoryCount(v int64)`

SetAdvisoryCount sets AdvisoryCount field to given value.

### HasAdvisoryCount

`func (o *TamAdvisoryCount) HasAdvisoryCount() bool`

HasAdvisoryCount returns a boolean if a field has been set.

### GetAccount

`func (o *TamAdvisoryCount) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *TamAdvisoryCount) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *TamAdvisoryCount) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *TamAdvisoryCount) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


