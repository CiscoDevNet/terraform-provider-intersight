# IssueDefinitionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "cond.AlarmDefinition"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "cond.AlarmDefinition"]
**Condition** | Pointer to [**NullableIssueCondition**](IssueCondition.md) |  | [optional] 
**Description** | Pointer to **string** | A description of the issue which is common to all instances of the issue. | [optional] 
**Name** | Pointer to **string** | An informational display name. | [optional] 
**ProbableCause** | Pointer to **string** | An explanation of the likely causes of the detected issue. | [optional] 
**Remediation** | Pointer to **string** | An explanation of the steps to perform to remediate the detected issue. | [optional] 
**SystemClassifications** | Pointer to **[]string** |  | [optional] 

## Methods

### NewIssueDefinitionAllOf

`func NewIssueDefinitionAllOf(classId string, objectType string, ) *IssueDefinitionAllOf`

NewIssueDefinitionAllOf instantiates a new IssueDefinitionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueDefinitionAllOfWithDefaults

`func NewIssueDefinitionAllOfWithDefaults() *IssueDefinitionAllOf`

NewIssueDefinitionAllOfWithDefaults instantiates a new IssueDefinitionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IssueDefinitionAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IssueDefinitionAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IssueDefinitionAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IssueDefinitionAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IssueDefinitionAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IssueDefinitionAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCondition

`func (o *IssueDefinitionAllOf) GetCondition() IssueCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *IssueDefinitionAllOf) GetConditionOk() (*IssueCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *IssueDefinitionAllOf) SetCondition(v IssueCondition)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *IssueDefinitionAllOf) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *IssueDefinitionAllOf) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *IssueDefinitionAllOf) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetDescription

`func (o *IssueDefinitionAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IssueDefinitionAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IssueDefinitionAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IssueDefinitionAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *IssueDefinitionAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IssueDefinitionAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IssueDefinitionAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IssueDefinitionAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProbableCause

`func (o *IssueDefinitionAllOf) GetProbableCause() string`

GetProbableCause returns the ProbableCause field if non-nil, zero value otherwise.

### GetProbableCauseOk

`func (o *IssueDefinitionAllOf) GetProbableCauseOk() (*string, bool)`

GetProbableCauseOk returns a tuple with the ProbableCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbableCause

`func (o *IssueDefinitionAllOf) SetProbableCause(v string)`

SetProbableCause sets ProbableCause field to given value.

### HasProbableCause

`func (o *IssueDefinitionAllOf) HasProbableCause() bool`

HasProbableCause returns a boolean if a field has been set.

### GetRemediation

`func (o *IssueDefinitionAllOf) GetRemediation() string`

GetRemediation returns the Remediation field if non-nil, zero value otherwise.

### GetRemediationOk

`func (o *IssueDefinitionAllOf) GetRemediationOk() (*string, bool)`

GetRemediationOk returns a tuple with the Remediation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediation

`func (o *IssueDefinitionAllOf) SetRemediation(v string)`

SetRemediation sets Remediation field to given value.

### HasRemediation

`func (o *IssueDefinitionAllOf) HasRemediation() bool`

HasRemediation returns a boolean if a field has been set.

### GetSystemClassifications

`func (o *IssueDefinitionAllOf) GetSystemClassifications() []string`

GetSystemClassifications returns the SystemClassifications field if non-nil, zero value otherwise.

### GetSystemClassificationsOk

`func (o *IssueDefinitionAllOf) GetSystemClassificationsOk() (*[]string, bool)`

GetSystemClassificationsOk returns a tuple with the SystemClassifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemClassifications

`func (o *IssueDefinitionAllOf) SetSystemClassifications(v []string)`

SetSystemClassifications sets SystemClassifications field to given value.

### HasSystemClassifications

`func (o *IssueDefinitionAllOf) HasSystemClassifications() bool`

HasSystemClassifications returns a boolean if a field has been set.

### SetSystemClassificationsNil

`func (o *IssueDefinitionAllOf) SetSystemClassificationsNil(b bool)`

 SetSystemClassificationsNil sets the value for SystemClassifications to be an explicit nil

### UnsetSystemClassifications
`func (o *IssueDefinitionAllOf) UnsetSystemClassifications()`

UnsetSystemClassifications ensures that no value is present for SystemClassifications, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


