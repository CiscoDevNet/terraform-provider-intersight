# ResourcepoolResourceEvaluationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resourcepool.ResourceEvaluationStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resourcepool.ResourceEvaluationStatus"]
**ErrMsg** | Pointer to **string** | The reason for the failure in ResourceEvaluation. | [optional] [readonly] 
**State** | Pointer to **string** | The state of the evaluation operation. * &#x60;ok&#x60; - The policy association or validation is successful. * &#x60;error&#x60; - The policy association or validation has failed. * &#x60;validating&#x60; - The policy association or validation is in progress. | [optional] [readonly] [default to "ok"]

## Methods

### NewResourcepoolResourceEvaluationStatus

`func NewResourcepoolResourceEvaluationStatus(classId string, objectType string, ) *ResourcepoolResourceEvaluationStatus`

NewResourcepoolResourceEvaluationStatus instantiates a new ResourcepoolResourceEvaluationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcepoolResourceEvaluationStatusWithDefaults

`func NewResourcepoolResourceEvaluationStatusWithDefaults() *ResourcepoolResourceEvaluationStatus`

NewResourcepoolResourceEvaluationStatusWithDefaults instantiates a new ResourcepoolResourceEvaluationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourcepoolResourceEvaluationStatus) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourcepoolResourceEvaluationStatus) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourcepoolResourceEvaluationStatus) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourcepoolResourceEvaluationStatus) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourcepoolResourceEvaluationStatus) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourcepoolResourceEvaluationStatus) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetErrMsg

`func (o *ResourcepoolResourceEvaluationStatus) GetErrMsg() string`

GetErrMsg returns the ErrMsg field if non-nil, zero value otherwise.

### GetErrMsgOk

`func (o *ResourcepoolResourceEvaluationStatus) GetErrMsgOk() (*string, bool)`

GetErrMsgOk returns a tuple with the ErrMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrMsg

`func (o *ResourcepoolResourceEvaluationStatus) SetErrMsg(v string)`

SetErrMsg sets ErrMsg field to given value.

### HasErrMsg

`func (o *ResourcepoolResourceEvaluationStatus) HasErrMsg() bool`

HasErrMsg returns a boolean if a field has been set.

### GetState

`func (o *ResourcepoolResourceEvaluationStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResourcepoolResourceEvaluationStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResourcepoolResourceEvaluationStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ResourcepoolResourceEvaluationStatus) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


