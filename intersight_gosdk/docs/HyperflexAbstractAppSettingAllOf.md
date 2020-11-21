# HyperflexAbstractAppSettingAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Name** | Pointer to **string** | The application setting identifier. | [optional] 
**Value** | Pointer to **string** | The application setting value. | [optional] 

## Methods

### NewHyperflexAbstractAppSettingAllOf

`func NewHyperflexAbstractAppSettingAllOf(classId string, objectType string, ) *HyperflexAbstractAppSettingAllOf`

NewHyperflexAbstractAppSettingAllOf instantiates a new HyperflexAbstractAppSettingAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexAbstractAppSettingAllOfWithDefaults

`func NewHyperflexAbstractAppSettingAllOfWithDefaults() *HyperflexAbstractAppSettingAllOf`

NewHyperflexAbstractAppSettingAllOfWithDefaults instantiates a new HyperflexAbstractAppSettingAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexAbstractAppSettingAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexAbstractAppSettingAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexAbstractAppSettingAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexAbstractAppSettingAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexAbstractAppSettingAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexAbstractAppSettingAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *HyperflexAbstractAppSettingAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HyperflexAbstractAppSettingAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HyperflexAbstractAppSettingAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HyperflexAbstractAppSettingAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *HyperflexAbstractAppSettingAllOf) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *HyperflexAbstractAppSettingAllOf) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *HyperflexAbstractAppSettingAllOf) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *HyperflexAbstractAppSettingAllOf) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


