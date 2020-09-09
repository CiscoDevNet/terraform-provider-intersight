# HclConstraint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConstraintName** | Pointer to **string** | Name or key of the applicable compatibility constraint. | [optional] 
**ConstraintValue** | Pointer to **string** | Value of the applicable compatibility constraint. Could either be a string value or a regex. | [optional] 

## Methods

### NewHclConstraint

`func NewHclConstraint() *HclConstraint`

NewHclConstraint instantiates a new HclConstraint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHclConstraintWithDefaults

`func NewHclConstraintWithDefaults() *HclConstraint`

NewHclConstraintWithDefaults instantiates a new HclConstraint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConstraintName

`func (o *HclConstraint) GetConstraintName() string`

GetConstraintName returns the ConstraintName field if non-nil, zero value otherwise.

### GetConstraintNameOk

`func (o *HclConstraint) GetConstraintNameOk() (*string, bool)`

GetConstraintNameOk returns a tuple with the ConstraintName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraintName

`func (o *HclConstraint) SetConstraintName(v string)`

SetConstraintName sets ConstraintName field to given value.

### HasConstraintName

`func (o *HclConstraint) HasConstraintName() bool`

HasConstraintName returns a boolean if a field has been set.

### GetConstraintValue

`func (o *HclConstraint) GetConstraintValue() string`

GetConstraintValue returns the ConstraintValue field if non-nil, zero value otherwise.

### GetConstraintValueOk

`func (o *HclConstraint) GetConstraintValueOk() (*string, bool)`

GetConstraintValueOk returns a tuple with the ConstraintValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraintValue

`func (o *HclConstraint) SetConstraintValue(v string)`

SetConstraintValue sets ConstraintValue field to given value.

### HasConstraintValue

`func (o *HclConstraint) HasConstraintValue() bool`

HasConstraintValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


