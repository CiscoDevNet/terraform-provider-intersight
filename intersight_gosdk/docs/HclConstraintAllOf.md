# HclConstraintAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hcl.Constraint"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hcl.Constraint"]
**ConstraintName** | Pointer to **string** | Name or key of the applicable compatibility constraint. | [optional] 
**ConstraintValue** | Pointer to **string** | Value of the applicable compatibility constraint. Could either be a string value or a regex. | [optional] 

## Methods

### NewHclConstraintAllOf

`func NewHclConstraintAllOf(classId string, objectType string, ) *HclConstraintAllOf`

NewHclConstraintAllOf instantiates a new HclConstraintAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHclConstraintAllOfWithDefaults

`func NewHclConstraintAllOfWithDefaults() *HclConstraintAllOf`

NewHclConstraintAllOfWithDefaults instantiates a new HclConstraintAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HclConstraintAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HclConstraintAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HclConstraintAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HclConstraintAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HclConstraintAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HclConstraintAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConstraintName

`func (o *HclConstraintAllOf) GetConstraintName() string`

GetConstraintName returns the ConstraintName field if non-nil, zero value otherwise.

### GetConstraintNameOk

`func (o *HclConstraintAllOf) GetConstraintNameOk() (*string, bool)`

GetConstraintNameOk returns a tuple with the ConstraintName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraintName

`func (o *HclConstraintAllOf) SetConstraintName(v string)`

SetConstraintName sets ConstraintName field to given value.

### HasConstraintName

`func (o *HclConstraintAllOf) HasConstraintName() bool`

HasConstraintName returns a boolean if a field has been set.

### GetConstraintValue

`func (o *HclConstraintAllOf) GetConstraintValue() string`

GetConstraintValue returns the ConstraintValue field if non-nil, zero value otherwise.

### GetConstraintValueOk

`func (o *HclConstraintAllOf) GetConstraintValueOk() (*string, bool)`

GetConstraintValueOk returns a tuple with the ConstraintValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraintValue

`func (o *HclConstraintAllOf) SetConstraintValue(v string)`

SetConstraintValue sets ConstraintValue field to given value.

### HasConstraintValue

`func (o *HclConstraintAllOf) HasConstraintValue() bool`

HasConstraintValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


