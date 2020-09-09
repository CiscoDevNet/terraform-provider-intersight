# WorkflowWorkflowDefinitionAllOf

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

### NewWorkflowWorkflowDefinitionAllOf

`func NewWorkflowWorkflowDefinitionAllOf() *WorkflowWorkflowDefinitionAllOf`

NewWorkflowWorkflowDefinitionAllOf instantiates a new WorkflowWorkflowDefinitionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowWorkflowDefinitionAllOfWithDefaults

`func NewWorkflowWorkflowDefinitionAllOfWithDefaults() *WorkflowWorkflowDefinitionAllOf`

NewWorkflowWorkflowDefinitionAllOfWithDefaults instantiates a new WorkflowWorkflowDefinitionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultVersion

`func (o *WorkflowWorkflowDefinitionAllOf) GetDefaultVersion() bool`

GetDefaultVersion returns the DefaultVersion field if non-nil, zero value otherwise.

### GetDefaultVersionOk

`func (o *WorkflowWorkflowDefinitionAllOf) GetDefaultVersionOk() (*bool, bool)`

GetDefaultVersionOk returns a tuple with the DefaultVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVersion

`func (o *WorkflowWorkflowDefinitionAllOf) SetDefaultVersion(v bool)`

SetDefaultVersion sets DefaultVersion field to given value.

### HasDefaultVersion

`func (o *WorkflowWorkflowDefinitionAllOf) HasDefaultVersion() bool`

HasDefaultVersion returns a boolean if a field has been set.

### GetDescription

`func (o *WorkflowWorkflowDefinitionAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowWorkflowDefinitionAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowWorkflowDefinitionAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowWorkflowDefinitionAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInputDefinition

`func (o *WorkflowWorkflowDefinitionAllOf) GetInputDefinition() []WorkflowBaseDataType`

GetInputDefinition returns the InputDefinition field if non-nil, zero value otherwise.

### GetInputDefinitionOk

`func (o *WorkflowWorkflowDefinitionAllOf) GetInputDefinitionOk() (*[]WorkflowBaseDataType, bool)`

GetInputDefinitionOk returns a tuple with the InputDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputDefinition

`func (o *WorkflowWorkflowDefinitionAllOf) SetInputDefinition(v []WorkflowBaseDataType)`

SetInputDefinition sets InputDefinition field to given value.

### HasInputDefinition

`func (o *WorkflowWorkflowDefinitionAllOf) HasInputDefinition() bool`

HasInputDefinition returns a boolean if a field has been set.

### GetLabel

`func (o *WorkflowWorkflowDefinitionAllOf) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WorkflowWorkflowDefinitionAllOf) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WorkflowWorkflowDefinitionAllOf) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WorkflowWorkflowDefinitionAllOf) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLicenseEntitlement

`func (o *WorkflowWorkflowDefinitionAllOf) GetLicenseEntitlement() string`

GetLicenseEntitlement returns the LicenseEntitlement field if non-nil, zero value otherwise.

### GetLicenseEntitlementOk

`func (o *WorkflowWorkflowDefinitionAllOf) GetLicenseEntitlementOk() (*string, bool)`

GetLicenseEntitlementOk returns a tuple with the LicenseEntitlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseEntitlement

`func (o *WorkflowWorkflowDefinitionAllOf) SetLicenseEntitlement(v string)`

SetLicenseEntitlement sets LicenseEntitlement field to given value.

### HasLicenseEntitlement

`func (o *WorkflowWorkflowDefinitionAllOf) HasLicenseEntitlement() bool`

HasLicenseEntitlement returns a boolean if a field has been set.

### GetMaxTaskCount

`func (o *WorkflowWorkflowDefinitionAllOf) GetMaxTaskCount() int64`

GetMaxTaskCount returns the MaxTaskCount field if non-nil, zero value otherwise.

### GetMaxTaskCountOk

`func (o *WorkflowWorkflowDefinitionAllOf) GetMaxTaskCountOk() (*int64, bool)`

GetMaxTaskCountOk returns a tuple with the MaxTaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTaskCount

`func (o *WorkflowWorkflowDefinitionAllOf) SetMaxTaskCount(v int64)`

SetMaxTaskCount sets MaxTaskCount field to given value.

### HasMaxTaskCount

`func (o *WorkflowWorkflowDefinitionAllOf) HasMaxTaskCount() bool`

HasMaxTaskCount returns a boolean if a field has been set.

### GetName

`func (o *WorkflowWorkflowDefinitionAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowWorkflowDefinitionAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowWorkflowDefinitionAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowWorkflowDefinitionAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutputDefinition

`func (o *WorkflowWorkflowDefinitionAllOf) GetOutputDefinition() []WorkflowBaseDataType`

GetOutputDefinition returns the OutputDefinition field if non-nil, zero value otherwise.

### GetOutputDefinitionOk

`func (o *WorkflowWorkflowDefinitionAllOf) GetOutputDefinitionOk() (*[]WorkflowBaseDataType, bool)`

GetOutputDefinitionOk returns a tuple with the OutputDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputDefinition

`func (o *WorkflowWorkflowDefinitionAllOf) SetOutputDefinition(v []WorkflowBaseDataType)`

SetOutputDefinition sets OutputDefinition field to given value.

### HasOutputDefinition

`func (o *WorkflowWorkflowDefinitionAllOf) HasOutputDefinition() bool`

HasOutputDefinition returns a boolean if a field has been set.

### GetOutputParameters

`func (o *WorkflowWorkflowDefinitionAllOf) GetOutputParameters() interface{}`

GetOutputParameters returns the OutputParameters field if non-nil, zero value otherwise.

### GetOutputParametersOk

`func (o *WorkflowWorkflowDefinitionAllOf) GetOutputParametersOk() (*interface{}, bool)`

GetOutputParametersOk returns a tuple with the OutputParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputParameters

`func (o *WorkflowWorkflowDefinitionAllOf) SetOutputParameters(v interface{})`

SetOutputParameters sets OutputParameters field to given value.

### HasOutputParameters

`func (o *WorkflowWorkflowDefinitionAllOf) HasOutputParameters() bool`

HasOutputParameters returns a boolean if a field has been set.

### SetOutputParametersNil

`func (o *WorkflowWorkflowDefinitionAllOf) SetOutputParametersNil(b bool)`

 SetOutputParametersNil sets the value for OutputParameters to be an explicit nil

### UnsetOutputParameters
`func (o *WorkflowWorkflowDefinitionAllOf) UnsetOutputParameters()`

UnsetOutputParameters ensures that no value is present for OutputParameters, not even an explicit nil
### GetProperties

`func (o *WorkflowWorkflowDefinitionAllOf) GetProperties() WorkflowWorkflowProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WorkflowWorkflowDefinitionAllOf) GetPropertiesOk() (*WorkflowWorkflowProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WorkflowWorkflowDefinitionAllOf) SetProperties(v WorkflowWorkflowProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *WorkflowWorkflowDefinitionAllOf) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetTasks

`func (o *WorkflowWorkflowDefinitionAllOf) GetTasks() []WorkflowWorkflowTask`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *WorkflowWorkflowDefinitionAllOf) GetTasksOk() (*[]WorkflowWorkflowTask, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *WorkflowWorkflowDefinitionAllOf) SetTasks(v []WorkflowWorkflowTask)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *WorkflowWorkflowDefinitionAllOf) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetUiRenderingData

`func (o *WorkflowWorkflowDefinitionAllOf) GetUiRenderingData() interface{}`

GetUiRenderingData returns the UiRenderingData field if non-nil, zero value otherwise.

### GetUiRenderingDataOk

`func (o *WorkflowWorkflowDefinitionAllOf) GetUiRenderingDataOk() (*interface{}, bool)`

GetUiRenderingDataOk returns a tuple with the UiRenderingData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiRenderingData

`func (o *WorkflowWorkflowDefinitionAllOf) SetUiRenderingData(v interface{})`

SetUiRenderingData sets UiRenderingData field to given value.

### HasUiRenderingData

`func (o *WorkflowWorkflowDefinitionAllOf) HasUiRenderingData() bool`

HasUiRenderingData returns a boolean if a field has been set.

### SetUiRenderingDataNil

`func (o *WorkflowWorkflowDefinitionAllOf) SetUiRenderingDataNil(b bool)`

 SetUiRenderingDataNil sets the value for UiRenderingData to be an explicit nil

### UnsetUiRenderingData
`func (o *WorkflowWorkflowDefinitionAllOf) UnsetUiRenderingData()`

UnsetUiRenderingData ensures that no value is present for UiRenderingData, not even an explicit nil
### GetValidationInformation

`func (o *WorkflowWorkflowDefinitionAllOf) GetValidationInformation() WorkflowValidationInformation`

GetValidationInformation returns the ValidationInformation field if non-nil, zero value otherwise.

### GetValidationInformationOk

`func (o *WorkflowWorkflowDefinitionAllOf) GetValidationInformationOk() (*WorkflowValidationInformation, bool)`

GetValidationInformationOk returns a tuple with the ValidationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationInformation

`func (o *WorkflowWorkflowDefinitionAllOf) SetValidationInformation(v WorkflowValidationInformation)`

SetValidationInformation sets ValidationInformation field to given value.

### HasValidationInformation

`func (o *WorkflowWorkflowDefinitionAllOf) HasValidationInformation() bool`

HasValidationInformation returns a boolean if a field has been set.

### GetVersion

`func (o *WorkflowWorkflowDefinitionAllOf) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkflowWorkflowDefinitionAllOf) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkflowWorkflowDefinitionAllOf) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WorkflowWorkflowDefinitionAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCatalog

`func (o *WorkflowWorkflowDefinitionAllOf) GetCatalog() WorkflowCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *WorkflowWorkflowDefinitionAllOf) GetCatalogOk() (*WorkflowCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *WorkflowWorkflowDefinitionAllOf) SetCatalog(v WorkflowCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *WorkflowWorkflowDefinitionAllOf) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


