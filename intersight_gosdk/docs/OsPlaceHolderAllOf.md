# OsPlaceHolderAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "os.PlaceHolder"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "os.PlaceHolder"]
**IsValueSet** | Pointer to **bool** | Flag to indicate if value is set. Value will be used to check if any edit. | [optional] [default to true]
**Type** | Pointer to [**WorkflowPrimitiveDataType**](WorkflowPrimitiveDataType.md) |  | [optional] 
**Value** | Pointer to **interface{}** | Value for placeholder provided by user. | [optional] 

## Methods

### NewOsPlaceHolderAllOf

`func NewOsPlaceHolderAllOf(classId string, objectType string, ) *OsPlaceHolderAllOf`

NewOsPlaceHolderAllOf instantiates a new OsPlaceHolderAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsPlaceHolderAllOfWithDefaults

`func NewOsPlaceHolderAllOfWithDefaults() *OsPlaceHolderAllOf`

NewOsPlaceHolderAllOfWithDefaults instantiates a new OsPlaceHolderAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OsPlaceHolderAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OsPlaceHolderAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OsPlaceHolderAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OsPlaceHolderAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OsPlaceHolderAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OsPlaceHolderAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsValueSet

`func (o *OsPlaceHolderAllOf) GetIsValueSet() bool`

GetIsValueSet returns the IsValueSet field if non-nil, zero value otherwise.

### GetIsValueSetOk

`func (o *OsPlaceHolderAllOf) GetIsValueSetOk() (*bool, bool)`

GetIsValueSetOk returns a tuple with the IsValueSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValueSet

`func (o *OsPlaceHolderAllOf) SetIsValueSet(v bool)`

SetIsValueSet sets IsValueSet field to given value.

### HasIsValueSet

`func (o *OsPlaceHolderAllOf) HasIsValueSet() bool`

HasIsValueSet returns a boolean if a field has been set.

### GetType

`func (o *OsPlaceHolderAllOf) GetType() WorkflowPrimitiveDataType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OsPlaceHolderAllOf) GetTypeOk() (*WorkflowPrimitiveDataType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OsPlaceHolderAllOf) SetType(v WorkflowPrimitiveDataType)`

SetType sets Type field to given value.

### HasType

`func (o *OsPlaceHolderAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *OsPlaceHolderAllOf) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *OsPlaceHolderAllOf) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *OsPlaceHolderAllOf) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *OsPlaceHolderAllOf) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *OsPlaceHolderAllOf) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *OsPlaceHolderAllOf) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


