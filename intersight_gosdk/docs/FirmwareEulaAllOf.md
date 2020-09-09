# FirmwareEulaAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accepted** | Pointer to **bool** | EULA acceptance status for the account. | [optional] [readonly] 
**Content** | Pointer to **string** | EULA acceptance form content provided by cisco.com. | [optional] [readonly] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 

## Methods

### NewFirmwareEulaAllOf

`func NewFirmwareEulaAllOf() *FirmwareEulaAllOf`

NewFirmwareEulaAllOf instantiates a new FirmwareEulaAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareEulaAllOfWithDefaults

`func NewFirmwareEulaAllOfWithDefaults() *FirmwareEulaAllOf`

NewFirmwareEulaAllOfWithDefaults instantiates a new FirmwareEulaAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccepted

`func (o *FirmwareEulaAllOf) GetAccepted() bool`

GetAccepted returns the Accepted field if non-nil, zero value otherwise.

### GetAcceptedOk

`func (o *FirmwareEulaAllOf) GetAcceptedOk() (*bool, bool)`

GetAcceptedOk returns a tuple with the Accepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepted

`func (o *FirmwareEulaAllOf) SetAccepted(v bool)`

SetAccepted sets Accepted field to given value.

### HasAccepted

`func (o *FirmwareEulaAllOf) HasAccepted() bool`

HasAccepted returns a boolean if a field has been set.

### GetContent

`func (o *FirmwareEulaAllOf) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *FirmwareEulaAllOf) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *FirmwareEulaAllOf) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *FirmwareEulaAllOf) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetAccount

`func (o *FirmwareEulaAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *FirmwareEulaAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *FirmwareEulaAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *FirmwareEulaAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


