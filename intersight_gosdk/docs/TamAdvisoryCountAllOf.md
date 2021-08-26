# TamAdvisoryCountAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "tam.AdvisoryCount"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "tam.AdvisoryCount"]
**AdvisoryCount** | Pointer to **int64** | Total number of advisories affecting the account. | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewTamAdvisoryCountAllOf

`func NewTamAdvisoryCountAllOf(classId string, objectType string, ) *TamAdvisoryCountAllOf`

NewTamAdvisoryCountAllOf instantiates a new TamAdvisoryCountAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTamAdvisoryCountAllOfWithDefaults

`func NewTamAdvisoryCountAllOfWithDefaults() *TamAdvisoryCountAllOf`

NewTamAdvisoryCountAllOfWithDefaults instantiates a new TamAdvisoryCountAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TamAdvisoryCountAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TamAdvisoryCountAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TamAdvisoryCountAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TamAdvisoryCountAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TamAdvisoryCountAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TamAdvisoryCountAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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


