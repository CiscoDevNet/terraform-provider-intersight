# OsPlaceHolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "os.PlaceHolder"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "os.PlaceHolder"]
**IsValueSet** | Pointer to **bool** | Flag to indicate if value is set. Value will be used to check if any edit. | [optional] [default to true]
**Type** | Pointer to [**WorkflowPrimitiveDataType**](WorkflowPrimitiveDataType.md) |  | [optional] 
**Value** | Pointer to **interface{}** | Value for placeholder provided by user. | [optional] 

## Methods

### NewOsPlaceHolder

`func NewOsPlaceHolder(classId string, objectType string, ) *OsPlaceHolder`

NewOsPlaceHolder instantiates a new OsPlaceHolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsPlaceHolderWithDefaults

`func NewOsPlaceHolderWithDefaults() *OsPlaceHolder`

NewOsPlaceHolderWithDefaults instantiates a new OsPlaceHolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OsPlaceHolder) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OsPlaceHolder) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OsPlaceHolder) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OsPlaceHolder) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OsPlaceHolder) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OsPlaceHolder) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsValueSet

`func (o *OsPlaceHolder) GetIsValueSet() bool`

GetIsValueSet returns the IsValueSet field if non-nil, zero value otherwise.

### GetIsValueSetOk

`func (o *OsPlaceHolder) GetIsValueSetOk() (*bool, bool)`

GetIsValueSetOk returns a tuple with the IsValueSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValueSet

`func (o *OsPlaceHolder) SetIsValueSet(v bool)`

SetIsValueSet sets IsValueSet field to given value.

### HasIsValueSet

`func (o *OsPlaceHolder) HasIsValueSet() bool`

HasIsValueSet returns a boolean if a field has been set.

### GetType

`func (o *OsPlaceHolder) GetType() WorkflowPrimitiveDataType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OsPlaceHolder) GetTypeOk() (*WorkflowPrimitiveDataType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OsPlaceHolder) SetType(v WorkflowPrimitiveDataType)`

SetType sets Type field to given value.

### HasType

`func (o *OsPlaceHolder) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *OsPlaceHolder) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *OsPlaceHolder) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *OsPlaceHolder) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *OsPlaceHolder) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *OsPlaceHolder) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *OsPlaceHolder) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


