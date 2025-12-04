# PolicyReportedPolicyChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "policy.ReportedPolicyChange"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "policy.ReportedPolicyChange"]
**ChangeId** | Pointer to **string** | The change evaluation identifier for which the change is reported. | [optional] [readonly] 
**ChangeStatus** | Pointer to **string** | The status of policy change evaluation which has been reported. * &#x60;Initiated&#x60; - The status when policy change evaluation is triggered for a policy. * &#x60;Reported&#x60; - The status when policy change evaluation is reported for a policy. | [optional] [readonly] [default to "Initiated"]
**PolicyType** | Pointer to **string** | The type of policy for which the change has been reported. | [optional] [readonly] 

## Methods

### NewPolicyReportedPolicyChange

`func NewPolicyReportedPolicyChange(classId string, objectType string, ) *PolicyReportedPolicyChange`

NewPolicyReportedPolicyChange instantiates a new PolicyReportedPolicyChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyReportedPolicyChangeWithDefaults

`func NewPolicyReportedPolicyChangeWithDefaults() *PolicyReportedPolicyChange`

NewPolicyReportedPolicyChangeWithDefaults instantiates a new PolicyReportedPolicyChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PolicyReportedPolicyChange) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PolicyReportedPolicyChange) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PolicyReportedPolicyChange) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PolicyReportedPolicyChange) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PolicyReportedPolicyChange) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PolicyReportedPolicyChange) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetChangeId

`func (o *PolicyReportedPolicyChange) GetChangeId() string`

GetChangeId returns the ChangeId field if non-nil, zero value otherwise.

### GetChangeIdOk

`func (o *PolicyReportedPolicyChange) GetChangeIdOk() (*string, bool)`

GetChangeIdOk returns a tuple with the ChangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeId

`func (o *PolicyReportedPolicyChange) SetChangeId(v string)`

SetChangeId sets ChangeId field to given value.

### HasChangeId

`func (o *PolicyReportedPolicyChange) HasChangeId() bool`

HasChangeId returns a boolean if a field has been set.

### GetChangeStatus

`func (o *PolicyReportedPolicyChange) GetChangeStatus() string`

GetChangeStatus returns the ChangeStatus field if non-nil, zero value otherwise.

### GetChangeStatusOk

`func (o *PolicyReportedPolicyChange) GetChangeStatusOk() (*string, bool)`

GetChangeStatusOk returns a tuple with the ChangeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeStatus

`func (o *PolicyReportedPolicyChange) SetChangeStatus(v string)`

SetChangeStatus sets ChangeStatus field to given value.

### HasChangeStatus

`func (o *PolicyReportedPolicyChange) HasChangeStatus() bool`

HasChangeStatus returns a boolean if a field has been set.

### GetPolicyType

`func (o *PolicyReportedPolicyChange) GetPolicyType() string`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *PolicyReportedPolicyChange) GetPolicyTypeOk() (*string, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *PolicyReportedPolicyChange) SetPolicyType(v string)`

SetPolicyType sets PolicyType field to given value.

### HasPolicyType

`func (o *PolicyReportedPolicyChange) HasPolicyType() bool`

HasPolicyType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


