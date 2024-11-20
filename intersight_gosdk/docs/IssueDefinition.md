# IssueDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "cond.AlarmDefinition"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "cond.AlarmDefinition"]
**Condition** | Pointer to [**NullableMoBaseComplexType**](MoBaseComplexType.md) | Condition defines the set of criteria under which an issue exists. | [optional] 
**Description** | Pointer to **string** | A description of the issue which is common to all instances of the issue. | [optional] 
**Name** | Pointer to **string** | An informational display name. | [optional] 
**ProbableCause** | Pointer to **string** | An explanation of the likely causes of the detected issue. | [optional] 
**Remediation** | Pointer to **string** | An explanation of the steps to perform to remediate the detected issue. | [optional] 
**SystemClassifications** | Pointer to **[]string** |  | [optional] 

## Methods

### NewIssueDefinition

`func NewIssueDefinition(classId string, objectType string, ) *IssueDefinition`

NewIssueDefinition instantiates a new IssueDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueDefinitionWithDefaults

`func NewIssueDefinitionWithDefaults() *IssueDefinition`

NewIssueDefinitionWithDefaults instantiates a new IssueDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IssueDefinition) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IssueDefinition) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IssueDefinition) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IssueDefinition) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IssueDefinition) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IssueDefinition) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCondition

`func (o *IssueDefinition) GetCondition() MoBaseComplexType`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *IssueDefinition) GetConditionOk() (*MoBaseComplexType, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *IssueDefinition) SetCondition(v MoBaseComplexType)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *IssueDefinition) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *IssueDefinition) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *IssueDefinition) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetDescription

`func (o *IssueDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IssueDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IssueDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IssueDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *IssueDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IssueDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IssueDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IssueDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProbableCause

`func (o *IssueDefinition) GetProbableCause() string`

GetProbableCause returns the ProbableCause field if non-nil, zero value otherwise.

### GetProbableCauseOk

`func (o *IssueDefinition) GetProbableCauseOk() (*string, bool)`

GetProbableCauseOk returns a tuple with the ProbableCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbableCause

`func (o *IssueDefinition) SetProbableCause(v string)`

SetProbableCause sets ProbableCause field to given value.

### HasProbableCause

`func (o *IssueDefinition) HasProbableCause() bool`

HasProbableCause returns a boolean if a field has been set.

### GetRemediation

`func (o *IssueDefinition) GetRemediation() string`

GetRemediation returns the Remediation field if non-nil, zero value otherwise.

### GetRemediationOk

`func (o *IssueDefinition) GetRemediationOk() (*string, bool)`

GetRemediationOk returns a tuple with the Remediation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediation

`func (o *IssueDefinition) SetRemediation(v string)`

SetRemediation sets Remediation field to given value.

### HasRemediation

`func (o *IssueDefinition) HasRemediation() bool`

HasRemediation returns a boolean if a field has been set.

### GetSystemClassifications

`func (o *IssueDefinition) GetSystemClassifications() []string`

GetSystemClassifications returns the SystemClassifications field if non-nil, zero value otherwise.

### GetSystemClassificationsOk

`func (o *IssueDefinition) GetSystemClassificationsOk() (*[]string, bool)`

GetSystemClassificationsOk returns a tuple with the SystemClassifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemClassifications

`func (o *IssueDefinition) SetSystemClassifications(v []string)`

SetSystemClassifications sets SystemClassifications field to given value.

### HasSystemClassifications

`func (o *IssueDefinition) HasSystemClassifications() bool`

HasSystemClassifications returns a boolean if a field has been set.

### SetSystemClassificationsNil

`func (o *IssueDefinition) SetSystemClassificationsNil(b bool)`

 SetSystemClassificationsNil sets the value for SystemClassifications to be an explicit nil

### UnsetSystemClassifications
`func (o *IssueDefinition) UnsetSystemClassifications()`

UnsetSystemClassifications ensures that no value is present for SystemClassifications, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


