# TamAdvisoryCountAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvisoryCount** | Pointer to **int64** | Total number of advisories affecting the account. | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 

## Methods

### NewTamAdvisoryCountAllOf

`func NewTamAdvisoryCountAllOf() *TamAdvisoryCountAllOf`

NewTamAdvisoryCountAllOf instantiates a new TamAdvisoryCountAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTamAdvisoryCountAllOfWithDefaults

`func NewTamAdvisoryCountAllOfWithDefaults() *TamAdvisoryCountAllOf`

NewTamAdvisoryCountAllOfWithDefaults instantiates a new TamAdvisoryCountAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvisoryCount

`func (o *TamAdvisoryCountAllOf) GetAdvisoryCount() int64`

GetAdvisoryCount returns the AdvisoryCount field if non-nil, zero value otherwise.

### GetAdvisoryCountOk

`func (o *TamAdvisoryCountAllOf) GetAdvisoryCountOk() (*int64, bool)`

GetAdvisoryCountOk returns a tuple with the AdvisoryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvisoryCount

`func (o *TamAdvisoryCountAllOf) SetAdvisoryCount(v int64)`

SetAdvisoryCount sets AdvisoryCount field to given value.

### HasAdvisoryCount

`func (o *TamAdvisoryCountAllOf) HasAdvisoryCount() bool`

HasAdvisoryCount returns a boolean if a field has been set.

### GetAccount

`func (o *TamAdvisoryCountAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *TamAdvisoryCountAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *TamAdvisoryCountAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *TamAdvisoryCountAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


