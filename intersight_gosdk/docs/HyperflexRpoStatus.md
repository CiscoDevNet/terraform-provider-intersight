# HyperflexRpoStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.RpoStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.RpoStatus"]
**Actual** | Pointer to **int64** | Actual end time for the snapshot. | [optional] [readonly] 
**Expected** | Pointer to **int64** | Expected end time for the snapshot. | [optional] [readonly] 
**RpoExceeded** | Pointer to **bool** | A flag to determine if snapshot delivery is delayed. | [optional] [readonly] 

## Methods

### NewHyperflexRpoStatus

`func NewHyperflexRpoStatus(classId string, objectType string, ) *HyperflexRpoStatus`

NewHyperflexRpoStatus instantiates a new HyperflexRpoStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexRpoStatusWithDefaults

`func NewHyperflexRpoStatusWithDefaults() *HyperflexRpoStatus`

NewHyperflexRpoStatusWithDefaults instantiates a new HyperflexRpoStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexRpoStatus) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexRpoStatus) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexRpoStatus) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexRpoStatus) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexRpoStatus) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexRpoStatus) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetActual

`func (o *HyperflexRpoStatus) GetActual() int64`

GetActual returns the Actual field if non-nil, zero value otherwise.

### GetActualOk

`func (o *HyperflexRpoStatus) GetActualOk() (*int64, bool)`

GetActualOk returns a tuple with the Actual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActual

`func (o *HyperflexRpoStatus) SetActual(v int64)`

SetActual sets Actual field to given value.

### HasActual

`func (o *HyperflexRpoStatus) HasActual() bool`

HasActual returns a boolean if a field has been set.

### GetExpected

`func (o *HyperflexRpoStatus) GetExpected() int64`

GetExpected returns the Expected field if non-nil, zero value otherwise.

### GetExpectedOk

`func (o *HyperflexRpoStatus) GetExpectedOk() (*int64, bool)`

GetExpectedOk returns a tuple with the Expected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpected

`func (o *HyperflexRpoStatus) SetExpected(v int64)`

SetExpected sets Expected field to given value.

### HasExpected

`func (o *HyperflexRpoStatus) HasExpected() bool`

HasExpected returns a boolean if a field has been set.

### GetRpoExceeded

`func (o *HyperflexRpoStatus) GetRpoExceeded() bool`

GetRpoExceeded returns the RpoExceeded field if non-nil, zero value otherwise.

### GetRpoExceededOk

`func (o *HyperflexRpoStatus) GetRpoExceededOk() (*bool, bool)`

GetRpoExceededOk returns a tuple with the RpoExceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpoExceeded

`func (o *HyperflexRpoStatus) SetRpoExceeded(v bool)`

SetRpoExceeded sets RpoExceeded field to given value.

### HasRpoExceeded

`func (o *HyperflexRpoStatus) HasRpoExceeded() bool`

HasRpoExceeded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


