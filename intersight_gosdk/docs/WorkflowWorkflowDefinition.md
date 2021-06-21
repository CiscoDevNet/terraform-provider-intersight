# WorkflowWorkflowDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.WorkflowDefinition"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.WorkflowDefinition"]
**DefaultVersion** | Pointer to **bool** | When true this will be the workflow version that is used when a specific workflow definition version is not specified. The default version is used when user executes a workflow without specifying a version or when workflow is included in another workflow without a specific version. The very first workflow definition created with a name will be set as the default version, after that user can explicitly set any version of the workflow definition as the default version. | [optional] 
**Description** | Pointer to **string** | The description for this workflow. | [optional] 
**InputDefinition** | Pointer to [**[]WorkflowBaseDataType**](WorkflowBaseDataType.md) |  | [optional] 
**InputParameterSet** | Pointer to [**[]WorkflowParameterSet**](WorkflowParameterSet.md) |  | [optional] 
**Label** | Pointer to **string** | A user friendly short name to identify the workflow. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), forward slash (/), or an underscore (_). | [optional] 
**LicenseEntitlement** | Pointer to **string** | License entitlement required to run this workflow. It is calculated based on the highest license requirement of all its tasks. * &#x60;Base&#x60; - Base as a License type. It is default license type. * &#x60;Essential&#x60; - Essential as a License type. * &#x60;Standard&#x60; - Standard as a License type. * &#x60;Advantage&#x60; - Advantage as a License type. * &#x60;Premier&#x60; - Premier as a License type. * &#x60;IWO-Essential&#x60; - IWO-Essential as a License type. * &#x60;IWO-Advantage&#x60; - IWO-Advantage as a License type. * &#x60;IWO-Premier&#x60; - IWO-Premier as a License type. | [optional] [readonly] [default to "Base"]
**MaxTaskCount** | Pointer to **int64** | The maximum number of tasks that can be executed on this workflow. | [optional] [readonly] 
**MaxWorkerTaskCount** | Pointer to **int64** | The maximum number of external (worker) tasks that can be executed on this workflow. | [optional] [readonly] 
**Name** | Pointer to **string** | The name for this workflow. You can have multiple versions of the workflow with the same name. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.) or an underscore (_). | [optional] 
**OutputDefinition** | Pointer to [**[]WorkflowBaseDataType**](WorkflowBaseDataType.md) |  | [optional] 
**OutputParameters** | Pointer to **interface{}** | The output mappings for the workflow. The outputs for workflows will generally be task output variables that we want to export out at the end of the workflow. The format to specify the mapping is &#39;${Source.output.JsonPath}&#39;, where &#39;Source&#39; is the name of the task within the workflow. Any task output can be mapped to a workflow output as long as the types are compatible. It&#39;s followed by a JSON path expression to extract JSON fragment from source&#39;s output. | [optional] 
**Properties** | Pointer to [**NullableWorkflowWorkflowProperties**](workflow.WorkflowProperties.md) |  | [optional] 
**Tasks** | Pointer to [**[]WorkflowWorkflowTask**](WorkflowWorkflowTask.md) |  | [optional] 
**UiInputFilters** | Pointer to [**[]WorkflowUiInputFilter**](WorkflowUiInputFilter.md) |  | [optional] 
**UiRenderingData** | Pointer to **interface{}** | This will hold the data needed for workflow to be rendered in the user interface. | [optional] 
**ValidationInformation** | Pointer to [**NullableWorkflowValidationInformation**](workflow.ValidationInformation.md) |  | [optional] 
**Version** | Pointer to **int64** | The version of the workflow to support multiple versions. | [optional] [default to 1]
**Catalog** | Pointer to [**WorkflowCatalogRelationship**](workflow.Catalog.Relationship.md) |  | [optional] 
**ClonedFrom** | Pointer to [**WorkflowWorkflowDefinitionRelationship**](workflow.WorkflowDefinition.Relationship.md) |  | [optional] 
**WorkflowMetadata** | Pointer to [**WorkflowWorkflowMetadataRelationship**](workflow.WorkflowMetadata.Relationship.md) |  | [optional] 

## Methods

### NewWorkflowWorkflowDefinition

`func NewWorkflowWorkflowDefinition(classId string, objectType string, ) *WorkflowWorkflowDefinition`

NewWorkflowWorkflowDefinition instantiates a new WorkflowWorkflowDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowWorkflowDefinitionWithDefaults

`func NewWorkflowWorkflowDefinitionWithDefaults() *WorkflowWorkflowDefinition`

NewWorkflowWorkflowDefinitionWithDefaults instantiates a new WorkflowWorkflowDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowWorkflowDefinition) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowWorkflowDefinition) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowWorkflowDefinition) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowWorkflowDefinition) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowWorkflowDefinition) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowWorkflowDefinition) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDefaultVersion

`func (o *WorkflowWorkflowDefinition) GetDefaultVersion() bool`

GetDefaultVersion returns the DefaultVersion field if non-nil, zero value otherwise.

### GetDefaultVersionOk

`func (o *WorkflowWorkflowDefinition) GetDefaultVersionOk() (*bool, bool)`

GetDefaultVersionOk returns a tuple with the DefaultVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVersion

`func (o *WorkflowWorkflowDefinition) SetDefaultVersion(v bool)`

SetDefaultVersion sets DefaultVersion field to given value.

### HasDefaultVersion

`func (o *WorkflowWorkflowDefinition) HasDefaultVersion() bool`

HasDefaultVersion returns a boolean if a field has been set.

### GetDescription

`func (o *WorkflowWorkflowDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowWorkflowDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowWorkflowDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowWorkflowDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInputDefinition

`func (o *WorkflowWorkflowDefinition) GetInputDefinition() []WorkflowBaseDataType`

GetInputDefinition returns the InputDefinition field if non-nil, zero value otherwise.

### GetInputDefinitionOk

`func (o *WorkflowWorkflowDefinition) GetInputDefinitionOk() (*[]WorkflowBaseDataType, bool)`

GetInputDefinitionOk returns a tuple with the InputDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputDefinition

`func (o *WorkflowWorkflowDefinition) SetInputDefinition(v []WorkflowBaseDataType)`

SetInputDefinition sets InputDefinition field to given value.

### HasInputDefinition

`func (o *WorkflowWorkflowDefinition) HasInputDefinition() bool`

HasInputDefinition returns a boolean if a field has been set.

### SetInputDefinitionNil

`func (o *WorkflowWorkflowDefinition) SetInputDefinitionNil(b bool)`

 SetInputDefinitionNil sets the value for InputDefinition to be an explicit nil

### UnsetInputDefinition
`func (o *WorkflowWorkflowDefinition) UnsetInputDefinition()`

UnsetInputDefinition ensures that no value is present for InputDefinition, not even an explicit nil
### GetInputParameterSet

`func (o *WorkflowWorkflowDefinition) GetInputParameterSet() []WorkflowParameterSet`

GetInputParameterSet returns the InputParameterSet field if non-nil, zero value otherwise.

### GetInputParameterSetOk

`func (o *WorkflowWorkflowDefinition) GetInputParameterSetOk() (*[]WorkflowParameterSet, bool)`

GetInputParameterSetOk returns a tuple with the InputParameterSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputParameterSet

`func (o *WorkflowWorkflowDefinition) SetInputParameterSet(v []WorkflowParameterSet)`

SetInputParameterSet sets InputParameterSet field to given value.

### HasInputParameterSet

`func (o *WorkflowWorkflowDefinition) HasInputParameterSet() bool`

HasInputParameterSet returns a boolean if a field has been set.

### SetInputParameterSetNil

`func (o *WorkflowWorkflowDefinition) SetInputParameterSetNil(b bool)`

 SetInputParameterSetNil sets the value for InputParameterSet to be an explicit nil

### UnsetInputParameterSet
`func (o *WorkflowWorkflowDefinition) UnsetInputParameterSet()`

UnsetInputParameterSet ensures that no value is present for InputParameterSet, not even an explicit nil
### GetLabel

`func (o *WorkflowWorkflowDefinition) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WorkflowWorkflowDefinition) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WorkflowWorkflowDefinition) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WorkflowWorkflowDefinition) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLicenseEntitlement

`func (o *WorkflowWorkflowDefinition) GetLicenseEntitlement() string`

GetLicenseEntitlement returns the LicenseEntitlement field if non-nil, zero value otherwise.

### GetLicenseEntitlementOk

`func (o *WorkflowWorkflowDefinition) GetLicenseEntitlementOk() (*string, bool)`

GetLicenseEntitlementOk returns a tuple with the LicenseEntitlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseEntitlement

`func (o *WorkflowWorkflowDefinition) SetLicenseEntitlement(v string)`

SetLicenseEntitlement sets LicenseEntitlement field to given value.

### HasLicenseEntitlement

`func (o *WorkflowWorkflowDefinition) HasLicenseEntitlement() bool`

HasLicenseEntitlement returns a boolean if a field has been set.

### GetMaxTaskCount

`func (o *WorkflowWorkflowDefinition) GetMaxTaskCount() int64`

GetMaxTaskCount returns the MaxTaskCount field if non-nil, zero value otherwise.

### GetMaxTaskCountOk

`func (o *WorkflowWorkflowDefinition) GetMaxTaskCountOk() (*int64, bool)`

GetMaxTaskCountOk returns a tuple with the MaxTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTaskCount

`func (o *WorkflowWorkflowDefinition) SetMaxTaskCount(v int64)`

SetMaxTaskCount sets MaxTaskCount field to given value.

### HasMaxTaskCount

`func (o *WorkflowWorkflowDefinition) HasMaxTaskCount() bool`

HasMaxTaskCount returns a boolean if a field has been set.

### GetMaxWorkerTaskCount

`func (o *WorkflowWorkflowDefinition) GetMaxWorkerTaskCount() int64`

GetMaxWorkerTaskCount returns the MaxWorkerTaskCount field if non-nil, zero value otherwise.

### GetMaxWorkerTaskCountOk

`func (o *WorkflowWorkflowDefinition) GetMaxWorkerTaskCountOk() (*int64, bool)`

GetMaxWorkerTaskCountOk returns a tuple with the MaxWorkerTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWorkerTaskCount

`func (o *WorkflowWorkflowDefinition) SetMaxWorkerTaskCount(v int64)`

SetMaxWorkerTaskCount sets MaxWorkerTaskCount field to given value.

### HasMaxWorkerTaskCount

`func (o *WorkflowWorkflowDefinition) HasMaxWorkerTaskCount() bool`

HasMaxWorkerTaskCount returns a boolean if a field has been set.

### GetName

`func (o *WorkflowWorkflowDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowWorkflowDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowWorkflowDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowWorkflowDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutputDefinition

`func (o *WorkflowWorkflowDefinition) GetOutputDefinition() []WorkflowBaseDataType`

GetOutputDefinition returns the OutputDefinition field if non-nil, zero value otherwise.

### GetOutputDefinitionOk

`func (o *WorkflowWorkflowDefinition) GetOutputDefinitionOk() (*[]WorkflowBaseDataType, bool)`

GetOutputDefinitionOk returns a tuple with the OutputDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputDefinition

`func (o *WorkflowWorkflowDefinition) SetOutputDefinition(v []WorkflowBaseDataType)`

SetOutputDefinition sets OutputDefinition field to given value.

### HasOutputDefinition

`func (o *WorkflowWorkflowDefinition) HasOutputDefinition() bool`

HasOutputDefinition returns a boolean if a field has been set.

### SetOutputDefinitionNil

`func (o *WorkflowWorkflowDefinition) SetOutputDefinitionNil(b bool)`

 SetOutputDefinitionNil sets the value for OutputDefinition to be an explicit nil

### UnsetOutputDefinition
`func (o *WorkflowWorkflowDefinition) UnsetOutputDefinition()`

UnsetOutputDefinition ensures that no value is present for OutputDefinition, not even an explicit nil
### GetOutputParameters

`func (o *WorkflowWorkflowDefinition) GetOutputParameters() interface{}`

GetOutputParameters returns the OutputParameters field if non-nil, zero value otherwise.

### GetOutputParametersOk

`func (o *WorkflowWorkflowDefinition) GetOutputParametersOk() (*interface{}, bool)`

GetOutputParametersOk returns a tuple with the OutputParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputParameters

`func (o *WorkflowWorkflowDefinition) SetOutputParameters(v interface{})`

SetOutputParameters sets OutputParameters field to given value.

### HasOutputParameters

`func (o *WorkflowWorkflowDefinition) HasOutputParameters() bool`

HasOutputParameters returns a boolean if a field has been set.

### SetOutputParametersNil

`func (o *WorkflowWorkflowDefinition) SetOutputParametersNil(b bool)`

 SetOutputParametersNil sets the value for OutputParameters to be an explicit nil

### UnsetOutputParameters
`func (o *WorkflowWorkflowDefinition) UnsetOutputParameters()`

UnsetOutputParameters ensures that no value is present for OutputParameters, not even an explicit nil
### GetProperties

`func (o *WorkflowWorkflowDefinition) GetProperties() WorkflowWorkflowProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WorkflowWorkflowDefinition) GetPropertiesOk() (*WorkflowWorkflowProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WorkflowWorkflowDefinition) SetProperties(v WorkflowWorkflowProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *WorkflowWorkflowDefinition) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *WorkflowWorkflowDefinition) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *WorkflowWorkflowDefinition) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetTasks

`func (o *WorkflowWorkflowDefinition) GetTasks() []WorkflowWorkflowTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *WorkflowWorkflowDefinition) GetTasksOk() (*[]WorkflowWorkflowTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *WorkflowWorkflowDefinition) SetTasks(v []WorkflowWorkflowTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *WorkflowWorkflowDefinition) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### SetTasksNil

`func (o *WorkflowWorkflowDefinition) SetTasksNil(b bool)`

 SetTasksNil sets the value for Tasks to be an explicit nil

### UnsetTasks
`func (o *WorkflowWorkflowDefinition) UnsetTasks()`

UnsetTasks ensures that no value is present for Tasks, not even an explicit nil
### GetUiInputFilters

`func (o *WorkflowWorkflowDefinition) GetUiInputFilters() []WorkflowUiInputFilter`

GetUiInputFilters returns the UiInputFilters field if non-nil, zero value otherwise.

### GetUiInputFiltersOk

`func (o *WorkflowWorkflowDefinition) GetUiInputFiltersOk() (*[]WorkflowUiInputFilter, bool)`

GetUiInputFiltersOk returns a tuple with the UiInputFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiInputFilters

`func (o *WorkflowWorkflowDefinition) SetUiInputFilters(v []WorkflowUiInputFilter)`

SetUiInputFilters sets UiInputFilters field to given value.

### HasUiInputFilters

`func (o *WorkflowWorkflowDefinition) HasUiInputFilters() bool`

HasUiInputFilters returns a boolean if a field has been set.

### SetUiInputFiltersNil

`func (o *WorkflowWorkflowDefinition) SetUiInputFiltersNil(b bool)`

 SetUiInputFiltersNil sets the value for UiInputFilters to be an explicit nil

### UnsetUiInputFilters
`func (o *WorkflowWorkflowDefinition) UnsetUiInputFilters()`

UnsetUiInputFilters ensures that no value is present for UiInputFilters, not even an explicit nil
### GetUiRenderingData

`func (o *WorkflowWorkflowDefinition) GetUiRenderingData() interface{}`

GetUiRenderingData returns the UiRenderingData field if non-nil, zero value otherwise.

### GetUiRenderingDataOk

`func (o *WorkflowWorkflowDefinition) GetUiRenderingDataOk() (*interface{}, bool)`

GetUiRenderingDataOk returns a tuple with the UiRenderingData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiRenderingData

`func (o *WorkflowWorkflowDefinition) SetUiRenderingData(v interface{})`

SetUiRenderingData sets UiRenderingData field to given value.

### HasUiRenderingData

`func (o *WorkflowWorkflowDefinition) HasUiRenderingData() bool`

HasUiRenderingData returns a boolean if a field has been set.

### SetUiRenderingDataNil

`func (o *WorkflowWorkflowDefinition) SetUiRenderingDataNil(b bool)`

 SetUiRenderingDataNil sets the value for UiRenderingData to be an explicit nil

### UnsetUiRenderingData
`func (o *WorkflowWorkflowDefinition) UnsetUiRenderingData()`

UnsetUiRenderingData ensures that no value is present for UiRenderingData, not even an explicit nil
### GetValidationInformation

`func (o *WorkflowWorkflowDefinition) GetValidationInformation() WorkflowValidationInformation`

GetValidationInformation returns the ValidationInformation field if non-nil, zero value otherwise.

### GetValidationInformationOk

`func (o *WorkflowWorkflowDefinition) GetValidationInformationOk() (*WorkflowValidationInformation, bool)`

GetValidationInformationOk returns a tuple with the ValidationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationInformation

`func (o *WorkflowWorkflowDefinition) SetValidationInformation(v WorkflowValidationInformation)`

SetValidationInformation sets ValidationInformation field to given value.

### HasValidationInformation

`func (o *WorkflowWorkflowDefinition) HasValidationInformation() bool`

HasValidationInformation returns a boolean if a field has been set.

### SetValidationInformationNil

`func (o *WorkflowWorkflowDefinition) SetValidationInformationNil(b bool)`

 SetValidationInformationNil sets the value for ValidationInformation to be an explicit nil

### UnsetValidationInformation
`func (o *WorkflowWorkflowDefinition) UnsetValidationInformation()`

UnsetValidationInformation ensures that no value is present for ValidationInformation, not even an explicit nil
### GetVersion

`func (o *WorkflowWorkflowDefinition) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkflowWorkflowDefinition) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkflowWorkflowDefinition) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WorkflowWorkflowDefinition) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCatalog

`func (o *WorkflowWorkflowDefinition) GetCatalog() WorkflowCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *WorkflowWorkflowDefinition) GetCatalogOk() (*WorkflowCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *WorkflowWorkflowDefinition) SetCatalog(v WorkflowCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *WorkflowWorkflowDefinition) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### GetClonedFrom

`func (o *WorkflowWorkflowDefinition) GetClonedFrom() WorkflowWorkflowDefinitionRelationship`

GetClonedFrom returns the ClonedFrom field if non-nil, zero value otherwise.

### GetClonedFromOk

`func (o *WorkflowWorkflowDefinition) GetClonedFromOk() (*WorkflowWorkflowDefinitionRelationship, bool)`

GetClonedFromOk returns a tuple with the ClonedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClonedFrom

`func (o *WorkflowWorkflowDefinition) SetClonedFrom(v WorkflowWorkflowDefinitionRelationship)`

SetClonedFrom sets ClonedFrom field to given value.

### HasClonedFrom

`func (o *WorkflowWorkflowDefinition) HasClonedFrom() bool`

HasClonedFrom returns a boolean if a field has been set.

### GetWorkflowMetadata

`func (o *WorkflowWorkflowDefinition) GetWorkflowMetadata() WorkflowWorkflowMetadataRelationship`

GetWorkflowMetadata returns the WorkflowMetadata field if non-nil, zero value otherwise.

### GetWorkflowMetadataOk

`func (o *WorkflowWorkflowDefinition) GetWorkflowMetadataOk() (*WorkflowWorkflowMetadataRelationship, bool)`

GetWorkflowMetadataOk returns a tuple with the WorkflowMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowMetadata

`func (o *WorkflowWorkflowDefinition) SetWorkflowMetadata(v WorkflowWorkflowMetadataRelationship)`

SetWorkflowMetadata sets WorkflowMetadata field to given value.

### HasWorkflowMetadata

`func (o *WorkflowWorkflowDefinition) HasWorkflowMetadata() bool`

HasWorkflowMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


