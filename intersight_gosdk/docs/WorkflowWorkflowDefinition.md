# WorkflowWorkflowDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultVersion** | Pointer to **bool** | When true this will be the workflow version that is used when a specific workflow definition version is not specified. The default version is used when user executes a workflow without specifying a version or when workflow is included in another workflow without a specific version. The very first workflow definition created with a name will be set as the default version, after that user can explicitly set any version of the workflow definition as the default version. | [optional] 
**Description** | Pointer to **string** | The description for this workflow. | [optional] 
**InputDefinition** | Pointer to [**[]WorkflowBaseDataType**](workflow.BaseDataType.md) |  | [optional] 
**Label** | Pointer to **string** | A user friendly short name to identify the workflow. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_). | [optional] 
**LicenseEntitlement** | Pointer to **string** | License entitlement required to run this workflow. It is calculated based on the highest license requirement of all its tasks. * &#x60;Base&#x60; - Base as a License type. It is default license type. * &#x60;Essential&#x60; - Essential as a License type. * &#x60;Standard&#x60; - Standard as a License type. * &#x60;Advantage&#x60; - Advantage as a License type. * &#x60;Premier&#x60; - Premier as a License type. | [optional] [readonly] [default to "Base"]
**MaxTaskCount** | Pointer to **int64** | The maximum number of tasks that can be executed on this workflow. | [optional] [readonly] 
**Name** | Pointer to **string** | The name for this workflow. You can have multiple versions of the workflow with the same name. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.) or an underscore (_). | [optional] 
**OutputDefinition** | Pointer to [**[]WorkflowBaseDataType**](workflow.BaseDataType.md) |  | [optional] 
**OutputParameters** | Pointer to **interface{}** | The output mappings for the workflow. The outputs for worflows will generally be task output variables that we want to export out at the end of the workflow. The format to specify the mapping is &#39;${Source.output.JsonPath}&#39;. &#39;Source&#39; is the name of the task within the workflow. You can map any task output to a workflow output as long as the types are compatible. Following this is JSON path expression to extract JSON fragment from source&#39;s output. | [optional] 
**Properties** | Pointer to [**WorkflowWorkflowProperties**](workflow.WorkflowProperties.md) |  | [optional] 
**Tasks** | Pointer to [**[]WorkflowWorkflowTask**](workflow.WorkflowTask.md) |  | [optional] 
**UiRenderingData** | Pointer to **interface{}** | This will hold the data needed for workflow to be rendered in the user interface. | [optional] 
**ValidationInformation** | Pointer to [**WorkflowValidationInformation**](workflow.ValidationInformation.md) |  | [optional] 
**Version** | Pointer to **int64** | The version of the workflow to support multiple versions. | [optional] 
**Catalog** | Pointer to [**WorkflowCatalogRelationship**](workflow.Catalog.Relationship.md) |  | [optional] 

## Methods

### NewWorkflowWorkflowDefinition

`func NewWorkflowWorkflowDefinition() *WorkflowWorkflowDefinition`

NewWorkflowWorkflowDefinition instantiates a new WorkflowWorkflowDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowWorkflowDefinitionWithDefaults

`func NewWorkflowWorkflowDefinitionWithDefaults() *WorkflowWorkflowDefinition`

NewWorkflowWorkflowDefinitionWithDefaults instantiates a new WorkflowWorkflowDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


