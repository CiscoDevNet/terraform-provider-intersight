# TamBaseAdvisoryDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "tam.SecurityAdvisoryDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "tam.SecurityAdvisoryDetails"]
**Description** | Pointer to **string** | Brief description of details specified for an advisory type. | [optional] 

## Methods

### NewTamBaseAdvisoryDetails

`func NewTamBaseAdvisoryDetails(classId string, objectType string, ) *TamBaseAdvisoryDetails`

NewTamBaseAdvisoryDetails instantiates a new TamBaseAdvisoryDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTamBaseAdvisoryDetailsWithDefaults

`func NewTamBaseAdvisoryDetailsWithDefaults() *TamBaseAdvisoryDetails`

NewTamBaseAdvisoryDetailsWithDefaults instantiates a new TamBaseAdvisoryDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TamBaseAdvisoryDetails) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TamBaseAdvisoryDetails) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TamBaseAdvisoryDetails) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TamBaseAdvisoryDetails) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TamBaseAdvisoryDetails) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TamBaseAdvisoryDetails) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *TamBaseAdvisoryDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TamBaseAdvisoryDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TamBaseAdvisoryDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TamBaseAdvisoryDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


