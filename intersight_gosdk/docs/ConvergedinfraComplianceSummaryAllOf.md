# ConvergedinfraComplianceSummaryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "convergedinfra.ComplianceSummary"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "convergedinfra.ComplianceSummary"]
**Incomplete** | Pointer to **int64** | The count of elements where compliance information is not available. eg. For HCL of server, OS information is not available. | [optional] 
**NotEvaluated** | Pointer to **int64** | The count of elements where compliance has not been evaluated. e.g. For HCL of server, status has not been evaluated due to lack of information. | [optional] 
**NotListed** | Pointer to **int64** | The count of elements where compliance has failed for one or more reason. e.g. For HCL of server, some part of the HCL validation has failed. | [optional] 
**Validated** | Pointer to **int64** | The count of elements where compliance has passed validation for all components. e.g. For HCL of server, all of the components have passed validation. | [optional] 

## Methods

### NewConvergedinfraComplianceSummaryAllOf

`func NewConvergedinfraComplianceSummaryAllOf(classId string, objectType string, ) *ConvergedinfraComplianceSummaryAllOf`

NewConvergedinfraComplianceSummaryAllOf instantiates a new ConvergedinfraComplianceSummaryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvergedinfraComplianceSummaryAllOfWithDefaults

`func NewConvergedinfraComplianceSummaryAllOfWithDefaults() *ConvergedinfraComplianceSummaryAllOf`

NewConvergedinfraComplianceSummaryAllOfWithDefaults instantiates a new ConvergedinfraComplianceSummaryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConvergedinfraComplianceSummaryAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConvergedinfraComplianceSummaryAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConvergedinfraComplianceSummaryAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConvergedinfraComplianceSummaryAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConvergedinfraComplianceSummaryAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConvergedinfraComplianceSummaryAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIncomplete

`func (o *ConvergedinfraComplianceSummaryAllOf) GetIncomplete() int64`

GetIncomplete returns the Incomplete field if non-nil, zero value otherwise.

### GetIncompleteOk

`func (o *ConvergedinfraComplianceSummaryAllOf) GetIncompleteOk() (*int64, bool)`

GetIncompleteOk returns a tuple with the Incomplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomplete

`func (o *ConvergedinfraComplianceSummaryAllOf) SetIncomplete(v int64)`

SetIncomplete sets Incomplete field to given value.

### HasIncomplete

`func (o *ConvergedinfraComplianceSummaryAllOf) HasIncomplete() bool`

HasIncomplete returns a boolean if a field has been set.

### GetNotEvaluated

`func (o *ConvergedinfraComplianceSummaryAllOf) GetNotEvaluated() int64`

GetNotEvaluated returns the NotEvaluated field if non-nil, zero value otherwise.

### GetNotEvaluatedOk

`func (o *ConvergedinfraComplianceSummaryAllOf) GetNotEvaluatedOk() (*int64, bool)`

GetNotEvaluatedOk returns a tuple with the NotEvaluated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotEvaluated

`func (o *ConvergedinfraComplianceSummaryAllOf) SetNotEvaluated(v int64)`

SetNotEvaluated sets NotEvaluated field to given value.

### HasNotEvaluated

`func (o *ConvergedinfraComplianceSummaryAllOf) HasNotEvaluated() bool`

HasNotEvaluated returns a boolean if a field has been set.

### GetNotListed

`func (o *ConvergedinfraComplianceSummaryAllOf) GetNotListed() int64`

GetNotListed returns the NotListed field if non-nil, zero value otherwise.

### GetNotListedOk

`func (o *ConvergedinfraComplianceSummaryAllOf) GetNotListedOk() (*int64, bool)`

GetNotListedOk returns a tuple with the NotListed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotListed

`func (o *ConvergedinfraComplianceSummaryAllOf) SetNotListed(v int64)`

SetNotListed sets NotListed field to given value.

### HasNotListed

`func (o *ConvergedinfraComplianceSummaryAllOf) HasNotListed() bool`

HasNotListed returns a boolean if a field has been set.

### GetValidated

`func (o *ConvergedinfraComplianceSummaryAllOf) GetValidated() int64`

GetValidated returns the Validated field if non-nil, zero value otherwise.

### GetValidatedOk

`func (o *ConvergedinfraComplianceSummaryAllOf) GetValidatedOk() (*int64, bool)`

GetValidatedOk returns a tuple with the Validated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidated

`func (o *ConvergedinfraComplianceSummaryAllOf) SetValidated(v int64)`

SetValidated sets Validated field to given value.

### HasValidated

`func (o *ConvergedinfraComplianceSummaryAllOf) HasValidated() bool`

HasValidated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


