# WorkflowServiceItemActionWorkflowDefinitionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.ServiceItemActionWorkflowDefinition"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.ServiceItemActionWorkflowDefinition"]
**CatalogMoid** | Pointer to **string** | Specify the catalog moid that this workflow belongs. When catalog moid is not specified then the catalog of the service item is used first and if no workflow can be found in that catalog, then the shared catalog from Intersight is used. | [optional] 
**Description** | Pointer to **string** | The description of this workflow instance. | [optional] 
**InputParameters** | Pointer to **interface{}** | Capture the mapping of ActionDefinition inputDefinition to workflow definition. | [optional] 
**Label** | Pointer to **string** | A user defined label identifier of the workflow used for UI display. | [optional] 
**Name** | Pointer to **string** | The name of the workflow, this name must be unique across all the workflow definition used within the action definitions. | [optional] 
**Version** | Pointer to **int64** | The workflow definition version to use as subworkflow. When no version is specified then the default version of the workflow at the time of creating or updating this workflow is used. | [optional] 
**WorkflowDefinitionName** | Pointer to **string** | The qualified name of workflow that should be executed. | [optional] 

## Methods

### NewWorkflowServiceItemActionWorkflowDefinitionAllOf

`func NewWorkflowServiceItemActionWorkflowDefinitionAllOf(classId string, objectType string, ) *WorkflowServiceItemActionWorkflowDefinitionAllOf`

NewWorkflowServiceItemActionWorkflowDefinitionAllOf instantiates a new WorkflowServiceItemActionWorkflowDefinitionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowServiceItemActionWorkflowDefinitionAllOfWithDefaults

`func NewWorkflowServiceItemActionWorkflowDefinitionAllOfWithDefaults() *WorkflowServiceItemActionWorkflowDefinitionAllOf`

NewWorkflowServiceItemActionWorkflowDefinitionAllOfWithDefaults instantiates a new WorkflowServiceItemActionWorkflowDefinitionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCatalogMoid

`func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) GetCatalogMoid() string`

GetCatalogMoid returns the CatalogMoid field if non-nil, zero value otherwise.

### GetCatalogMoidOk

`func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) GetCatalogMoidOk() (*string, bool)`

GetCatalogMoidOk returns a tuple with the CatalogMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogMoid

`func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) SetCatalogMoid(v string)`

SetCatalogMoid sets CatalogMoid field to given value.

### HasCatalogMoid

`func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) HasCatalogMoid() bool`

HasCatalogMoid returns a boolean if a field has been set.

### GetDescription

`func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInputParameters

`func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) GetInputParameters() interface{}`

GetInputParameters returns the InputParameters field if non-nil, zero value otherwise.

### GetInputParametersOk

`func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) GetInputParametersOk() (*interface{}, bool)`

GetInputParametersOk returns a tuple with the InputParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputParameters

`func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) SetInputParameters(v interface{})`

SetInputParameters sets InputParameters field to given value.

### HasInputParameters

`func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) HasInputParameters() bool`

HasInputParameters returns a boolean if a field has been set.

### SetInputParametersNil

`func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) SetInputParametersNil(b bool)`

 SetInputParametersNil sets the value for InputParameters to be an explicit nil

### UnsetInputParameters
`func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) UnsetInputParameters()`

UnsetInputParameters ensures that no value is present for InputParameters, not even an explicit nil
### GetLabel

`func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWorkflowDefinitionName

`func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) GetWorkflowDefinitionName() string`

GetWorkflowDefinitionName returns the WorkflowDefinitionName field if non-nil, zero value otherwise.

### GetWorkflowDefinitionNameOk

`func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) GetWorkflowDefinitionNameOk() (*string, bool)`

GetWorkflowDefinitionNameOk returns a tuple with the WorkflowDefinitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowDefinitionName

`func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) SetWorkflowDefinitionName(v string)`

SetWorkflowDefinitionName sets WorkflowDefinitionName field to given value.

### HasWorkflowDefinitionName

`func (o *WorkflowServiceItemActionWorkflowDefinitionAllOf) HasWorkflowDefinitionName() bool`

HasWorkflowDefinitionName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


