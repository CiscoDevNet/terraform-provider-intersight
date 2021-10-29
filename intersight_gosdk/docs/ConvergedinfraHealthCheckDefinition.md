# ConvergedinfraHealthCheckDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "convergedinfra.HealthCheckDefinition"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "convergedinfra.HealthCheckDefinition"]
**Category** | Pointer to **string** | Category that the health check belongs to. | [optional] 
**CommonCauses** | Pointer to **string** | Static information detailing the common causes for the health check failure. | [optional] 
**Description** | Pointer to **string** | Description of the health check definition. | [optional] 
**ExecutionMode** | Pointer to **string** | Execution mode of the health check on converged infrastructure pod. * &#x60;OnDemand&#x60; - Execute the health check on-demand. * &#x60;Periodic&#x60; - Execute the health check on a periodic basis. | [optional] [default to "OnDemand"]
**Label** | Pointer to **string** | Label for the health check definition that is displayed on UI. | [optional] 
**Name** | Pointer to **string** | Name of the health check definition. | [optional] 
**SuggestedResolution** | Pointer to **string** | Static information detailing the possible remediation actions that can be taken to remedy the health check failure. | [optional] 

## Methods

### NewConvergedinfraHealthCheckDefinition

`func NewConvergedinfraHealthCheckDefinition(classId string, objectType string, ) *ConvergedinfraHealthCheckDefinition`

NewConvergedinfraHealthCheckDefinition instantiates a new ConvergedinfraHealthCheckDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvergedinfraHealthCheckDefinitionWithDefaults

`func NewConvergedinfraHealthCheckDefinitionWithDefaults() *ConvergedinfraHealthCheckDefinition`

NewConvergedinfraHealthCheckDefinitionWithDefaults instantiates a new ConvergedinfraHealthCheckDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConvergedinfraHealthCheckDefinition) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConvergedinfraHealthCheckDefinition) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConvergedinfraHealthCheckDefinition) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConvergedinfraHealthCheckDefinition) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConvergedinfraHealthCheckDefinition) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConvergedinfraHealthCheckDefinition) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCategory

`func (o *ConvergedinfraHealthCheckDefinition) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ConvergedinfraHealthCheckDefinition) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ConvergedinfraHealthCheckDefinition) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ConvergedinfraHealthCheckDefinition) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCommonCauses

`func (o *ConvergedinfraHealthCheckDefinition) GetCommonCauses() string`

GetCommonCauses returns the CommonCauses field if non-nil, zero value otherwise.

### GetCommonCausesOk

`func (o *ConvergedinfraHealthCheckDefinition) GetCommonCausesOk() (*string, bool)`

GetCommonCausesOk returns a tuple with the CommonCauses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonCauses

`func (o *ConvergedinfraHealthCheckDefinition) SetCommonCauses(v string)`

SetCommonCauses sets CommonCauses field to given value.

### HasCommonCauses

`func (o *ConvergedinfraHealthCheckDefinition) HasCommonCauses() bool`

HasCommonCauses returns a boolean if a field has been set.

### GetDescription

`func (o *ConvergedinfraHealthCheckDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConvergedinfraHealthCheckDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConvergedinfraHealthCheckDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConvergedinfraHealthCheckDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExecutionMode

`func (o *ConvergedinfraHealthCheckDefinition) GetExecutionMode() string`

GetExecutionMode returns the ExecutionMode field if non-nil, zero value otherwise.

### GetExecutionModeOk

`func (o *ConvergedinfraHealthCheckDefinition) GetExecutionModeOk() (*string, bool)`

GetExecutionModeOk returns a tuple with the ExecutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionMode

`func (o *ConvergedinfraHealthCheckDefinition) SetExecutionMode(v string)`

SetExecutionMode sets ExecutionMode field to given value.

### HasExecutionMode

`func (o *ConvergedinfraHealthCheckDefinition) HasExecutionMode() bool`

HasExecutionMode returns a boolean if a field has been set.

### GetLabel

`func (o *ConvergedinfraHealthCheckDefinition) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ConvergedinfraHealthCheckDefinition) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ConvergedinfraHealthCheckDefinition) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ConvergedinfraHealthCheckDefinition) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *ConvergedinfraHealthCheckDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConvergedinfraHealthCheckDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConvergedinfraHealthCheckDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConvergedinfraHealthCheckDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSuggestedResolution

`func (o *ConvergedinfraHealthCheckDefinition) GetSuggestedResolution() string`

GetSuggestedResolution returns the SuggestedResolution field if non-nil, zero value otherwise.

### GetSuggestedResolutionOk

`func (o *ConvergedinfraHealthCheckDefinition) GetSuggestedResolutionOk() (*string, bool)`

GetSuggestedResolutionOk returns a tuple with the SuggestedResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedResolution

`func (o *ConvergedinfraHealthCheckDefinition) SetSuggestedResolution(v string)`

SetSuggestedResolution sets SuggestedResolution field to given value.

### HasSuggestedResolution

`func (o *ConvergedinfraHealthCheckDefinition) HasSuggestedResolution() bool`

HasSuggestedResolution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


