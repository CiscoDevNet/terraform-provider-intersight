# IaasWorkflowSteps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iaas.WorkflowSteps"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iaas.WorkflowSteps"]
**CompletedTime** | Pointer to **string** | Completed time of the workflow step. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the workflow step. | [optional] [readonly] 
**Status** | Pointer to **string** | Status of the workflow step. | [optional] [readonly] 
**StatusMessage** | Pointer to **string** | Status message of the workflow step. | [optional] [readonly] 

## Methods

### NewIaasWorkflowSteps

`func NewIaasWorkflowSteps(classId string, objectType string, ) *IaasWorkflowSteps`

NewIaasWorkflowSteps instantiates a new IaasWorkflowSteps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIaasWorkflowStepsWithDefaults

`func NewIaasWorkflowStepsWithDefaults() *IaasWorkflowSteps`

NewIaasWorkflowStepsWithDefaults instantiates a new IaasWorkflowSteps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IaasWorkflowSteps) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IaasWorkflowSteps) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IaasWorkflowSteps) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IaasWorkflowSteps) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IaasWorkflowSteps) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IaasWorkflowSteps) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCompletedTime

`func (o *IaasWorkflowSteps) GetCompletedTime() string`

GetCompletedTime returns the CompletedTime field if non-nil, zero value otherwise.

### GetCompletedTimeOk

`func (o *IaasWorkflowSteps) GetCompletedTimeOk() (*string, bool)`

GetCompletedTimeOk returns a tuple with the CompletedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedTime

`func (o *IaasWorkflowSteps) SetCompletedTime(v string)`

SetCompletedTime sets CompletedTime field to given value.

### HasCompletedTime

`func (o *IaasWorkflowSteps) HasCompletedTime() bool`

HasCompletedTime returns a boolean if a field has been set.

### GetName

`func (o *IaasWorkflowSteps) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IaasWorkflowSteps) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IaasWorkflowSteps) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IaasWorkflowSteps) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *IaasWorkflowSteps) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IaasWorkflowSteps) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IaasWorkflowSteps) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IaasWorkflowSteps) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *IaasWorkflowSteps) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *IaasWorkflowSteps) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *IaasWorkflowSteps) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *IaasWorkflowSteps) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


