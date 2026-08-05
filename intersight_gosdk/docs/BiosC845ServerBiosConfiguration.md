# BiosC845ServerBiosConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "bios.C845ServerBiosConfiguration"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "bios.C845ServerBiosConfiguration"]
**BiosTokens** | Pointer to [**[]BiosC845BaseBiosToken**](BiosC845BaseBiosToken.md) |  | [optional] 

## Methods

### NewBiosC845ServerBiosConfiguration

`func NewBiosC845ServerBiosConfiguration(classId string, objectType string, ) *BiosC845ServerBiosConfiguration`

NewBiosC845ServerBiosConfiguration instantiates a new BiosC845ServerBiosConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiosC845ServerBiosConfigurationWithDefaults

`func NewBiosC845ServerBiosConfigurationWithDefaults() *BiosC845ServerBiosConfiguration`

NewBiosC845ServerBiosConfigurationWithDefaults instantiates a new BiosC845ServerBiosConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BiosC845ServerBiosConfiguration) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BiosC845ServerBiosConfiguration) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BiosC845ServerBiosConfiguration) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BiosC845ServerBiosConfiguration) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BiosC845ServerBiosConfiguration) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BiosC845ServerBiosConfiguration) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBiosTokens

`func (o *BiosC845ServerBiosConfiguration) GetBiosTokens() []BiosC845BaseBiosToken`

GetBiosTokens returns the BiosTokens field if non-nil, zero value otherwise.

### GetBiosTokensOk

`func (o *BiosC845ServerBiosConfiguration) GetBiosTokensOk() (*[]BiosC845BaseBiosToken, bool)`

GetBiosTokensOk returns a tuple with the BiosTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosTokens

`func (o *BiosC845ServerBiosConfiguration) SetBiosTokens(v []BiosC845BaseBiosToken)`

SetBiosTokens sets BiosTokens field to given value.

### HasBiosTokens

`func (o *BiosC845ServerBiosConfiguration) HasBiosTokens() bool`

HasBiosTokens returns a boolean if a field has been set.

### SetBiosTokensNil

`func (o *BiosC845ServerBiosConfiguration) SetBiosTokensNil(b bool)`

 SetBiosTokensNil sets the value for BiosTokens to be an explicit nil

### UnsetBiosTokens
`func (o *BiosC845ServerBiosConfiguration) UnsetBiosTokens()`

UnsetBiosTokens ensures that no value is present for BiosTokens, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


