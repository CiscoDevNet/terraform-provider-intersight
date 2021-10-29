# ConvergedinfraHealthCheckDefinitionAllOf

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

### NewConvergedinfraHealthCheckDefinitionAllOf

`func NewConvergedinfraHealthCheckDefinitionAllOf(classId string, objectType string, ) *ConvergedinfraHealthCheckDefinitionAllOf`

NewConvergedinfraHealthCheckDefinitionAllOf instantiates a new ConvergedinfraHealthCheckDefinitionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvergedinfraHealthCheckDefinitionAllOfWithDefaults

`func NewConvergedinfraHealthCheckDefinitionAllOfWithDefaults() *ConvergedinfraHealthCheckDefinitionAllOf`

NewConvergedinfraHealthCheckDefinitionAllOfWithDefaults instantiates a new ConvergedinfraHealthCheckDefinitionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ConvergedinfraHealthCheckDefinitionAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConvergedinfraHealthCheckDefinitionAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConvergedinfraHealthCheckDefinitionAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ConvergedinfraHealthCheckDefinitionAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConvergedinfraHealthCheckDefinitionAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConvergedinfraHealthCheckDefinitionAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCategory

`func (o *ConvergedinfraHealthCheckDefinitionAllOf) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ConvergedinfraHealthCheckDefinitionAllOf) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ConvergedinfraHealthCheckDefinitionAllOf) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ConvergedinfraHealthCheckDefinitionAllOf) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCommonCauses

`func (o *ConvergedinfraHealthCheckDefinitionAllOf) GetCommonCauses() string`

GetCommonCauses returns the CommonCauses field if non-nil, zero value otherwise.

### GetCommonCausesOk

`func (o *ConvergedinfraHealthCheckDefinitionAllOf) GetCommonCausesOk() (*string, bool)`

GetCommonCausesOk returns a tuple with the CommonCauses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonCauses

`func (o *ConvergedinfraHealthCheckDefinitionAllOf) SetCommonCauses(v string)`

SetCommonCauses sets CommonCauses field to given value.

### HasCommonCauses

`func (o *ConvergedinfraHealthCheckDefinitionAllOf) HasCommonCauses() bool`

HasCommonCauses returns a boolean if a field has been set.

### GetDescription

`func (o *ConvergedinfraHealthCheckDefinitionAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConvergedinfraHealthCheckDefinitionAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConvergedinfraHealthCheckDefinitionAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConvergedinfraHealthCheckDefinitionAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExecutionMode

`func (o *ConvergedinfraHealthCheckDefinitionAllOf) GetExecutionMode() string`

GetExecutionMode returns the ExecutionMode field if non-nil, zero value otherwise.

### GetExecutionModeOk

`func (o *ConvergedinfraHealthCheckDefinitionAllOf) GetExecutionModeOk() (*string, bool)`

GetExecutionModeOk returns a tuple with the ExecutionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionMode

`func (o *ConvergedinfraHealthCheckDefinitionAllOf) SetExecutionMode(v string)`

SetExecutionMode sets ExecutionMode field to given value.

### HasExecutionMode

`func (o *ConvergedinfraHealthCheckDefinitionAllOf) HasExecutionMode() bool`

HasExecutionMode returns a boolean if a field has been set.

### GetLabel

`func (o *ConvergedinfraHealthCheckDefinitionAllOf) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ConvergedinfraHealthCheckDefinitionAllOf) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ConvergedinfraHealthCheckDefinitionAllOf) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ConvergedinfraHealthCheckDefinitionAllOf) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *ConvergedinfraHealthCheckDefinitionAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConvergedinfraHealthCheckDefinitionAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConvergedinfraHealthCheckDefinitionAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConvergedinfraHealthCheckDefinitionAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSuggestedResolution

`func (o *ConvergedinfraHealthCheckDefinitionAllOf) GetSuggestedResolution() string`

GetSuggestedResolution returns the SuggestedResolution field if non-nil, zero value otherwise.

### GetSuggestedResolutionOk

`func (o *ConvergedinfraHealthCheckDefinitionAllOf) GetSuggestedResolutionOk() (*string, bool)`

GetSuggestedResolutionOk returns a tuple with the SuggestedResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedResolution

`func (o *ConvergedinfraHealthCheckDefinitionAllOf) SetSuggestedResolution(v string)`

SetSuggestedResolution sets SuggestedResolution field to given value.

### HasSuggestedResolution

`func (o *ConvergedinfraHealthCheckDefinitionAllOf) HasSuggestedResolution() bool`

HasSuggestedResolution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


