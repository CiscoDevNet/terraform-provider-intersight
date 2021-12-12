# WorkflowSolutionDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.SolutionDefinition"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.SolutionDefinition"]
**AllowMultipleSolutionInstances** | Pointer to **bool** | Solution definition can declare that only one instance can be allowed within the customer account. | [optional] [default to true]
**CvdId** | Pointer to **string** | The Cisco Validated Design (CVD) Identifier that this solution provides. | [optional] 
**DeleteInstanceOnDecommission** | Pointer to **bool** | The flag to indicate that solution instance will be deleted after the completion of decommission action. | [optional] [default to false]
**Description** | Pointer to **string** | The description for this solution which provides information on what are the pre-requisites to deploy the solution and what features are supported on the solution. | [optional] 
**Label** | Pointer to **string** | A user friendly short name to identify the solution. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_). | [optional] 
**LicenseEntitlement** | Pointer to **string** | License entitlement required to run this solution. * &#x60;Base&#x60; - Base as a License type. It is default license type. * &#x60;Essential&#x60; - Essential as a License type. * &#x60;Standard&#x60; - Standard as a License type. * &#x60;Advantage&#x60; - Advantage as a License type. * &#x60;Premier&#x60; - Premier as a License type. * &#x60;IWO-Essential&#x60; - IWO-Essential as a License type. * &#x60;IWO-Advantage&#x60; - IWO-Advantage as a License type. * &#x60;IWO-Premier&#x60; - IWO-Premier as a License type. * &#x60;IKS-Advantage&#x60; - IKS-Advantage as a License type. | [optional] [readonly] [default to "Base"]
**Name** | Pointer to **string** | The name for this solution definition. You can have multiple versions of the solution with the same name. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:) or an underscore (_). | [optional] 
**OutputDefinition** | Pointer to [**[]WorkflowBaseDataType**](WorkflowBaseDataType.md) |  | [optional] 
**Version** | Pointer to **int64** | The version of the solution to support multiple versions. | [optional] [default to 1]
**ActionDefinitions** | Pointer to [**[]WorkflowSolutionActionDefinitionRelationship**](WorkflowSolutionActionDefinitionRelationship.md) | An array of relationships to workflowSolutionActionDefinition resources. | [optional] 
**Catalog** | Pointer to [**WorkflowCatalogRelationship**](WorkflowCatalogRelationship.md) |  | [optional] 

## Methods

### NewWorkflowSolutionDefinition

`func NewWorkflowSolutionDefinition(classId string, objectType string, ) *WorkflowSolutionDefinition`

NewWorkflowSolutionDefinition instantiates a new WorkflowSolutionDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSolutionDefinitionWithDefaults

`func NewWorkflowSolutionDefinitionWithDefaults() *WorkflowSolutionDefinition`

NewWorkflowSolutionDefinitionWithDefaults instantiates a new WorkflowSolutionDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowSolutionDefinition) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowSolutionDefinition) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowSolutionDefinition) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowSolutionDefinition) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowSolutionDefinition) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowSolutionDefinition) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAllowMultipleSolutionInstances

`func (o *WorkflowSolutionDefinition) GetAllowMultipleSolutionInstances() bool`

GetAllowMultipleSolutionInstances returns the AllowMultipleSolutionInstances field if non-nil, zero value otherwise.

### GetAllowMultipleSolutionInstancesOk

`func (o *WorkflowSolutionDefinition) GetAllowMultipleSolutionInstancesOk() (*bool, bool)`

GetAllowMultipleSolutionInstancesOk returns a tuple with the AllowMultipleSolutionInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultipleSolutionInstances

`func (o *WorkflowSolutionDefinition) SetAllowMultipleSolutionInstances(v bool)`

SetAllowMultipleSolutionInstances sets AllowMultipleSolutionInstances field to given value.

### HasAllowMultipleSolutionInstances

`func (o *WorkflowSolutionDefinition) HasAllowMultipleSolutionInstances() bool`

HasAllowMultipleSolutionInstances returns a boolean if a field has been set.

### GetCvdId

`func (o *WorkflowSolutionDefinition) GetCvdId() string`

GetCvdId returns the CvdId field if non-nil, zero value otherwise.

### GetCvdIdOk

`func (o *WorkflowSolutionDefinition) GetCvdIdOk() (*string, bool)`

GetCvdIdOk returns a tuple with the CvdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvdId

`func (o *WorkflowSolutionDefinition) SetCvdId(v string)`

SetCvdId sets CvdId field to given value.

### HasCvdId

`func (o *WorkflowSolutionDefinition) HasCvdId() bool`

HasCvdId returns a boolean if a field has been set.

### GetDeleteInstanceOnDecommission

`func (o *WorkflowSolutionDefinition) GetDeleteInstanceOnDecommission() bool`

GetDeleteInstanceOnDecommission returns the DeleteInstanceOnDecommission field if non-nil, zero value otherwise.

### GetDeleteInstanceOnDecommissionOk

`func (o *WorkflowSolutionDefinition) GetDeleteInstanceOnDecommissionOk() (*bool, bool)`

GetDeleteInstanceOnDecommissionOk returns a tuple with the DeleteInstanceOnDecommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteInstanceOnDecommission

`func (o *WorkflowSolutionDefinition) SetDeleteInstanceOnDecommission(v bool)`

SetDeleteInstanceOnDecommission sets DeleteInstanceOnDecommission field to given value.

### HasDeleteInstanceOnDecommission

`func (o *WorkflowSolutionDefinition) HasDeleteInstanceOnDecommission() bool`

HasDeleteInstanceOnDecommission returns a boolean if a field has been set.

### GetDescription

`func (o *WorkflowSolutionDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowSolutionDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowSolutionDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowSolutionDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabel

`func (o *WorkflowSolutionDefinition) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WorkflowSolutionDefinition) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WorkflowSolutionDefinition) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WorkflowSolutionDefinition) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLicenseEntitlement

`func (o *WorkflowSolutionDefinition) GetLicenseEntitlement() string`

GetLicenseEntitlement returns the LicenseEntitlement field if non-nil, zero value otherwise.

### GetLicenseEntitlementOk

`func (o *WorkflowSolutionDefinition) GetLicenseEntitlementOk() (*string, bool)`

GetLicenseEntitlementOk returns a tuple with the LicenseEntitlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseEntitlement

`func (o *WorkflowSolutionDefinition) SetLicenseEntitlement(v string)`

SetLicenseEntitlement sets LicenseEntitlement field to given value.

### HasLicenseEntitlement

`func (o *WorkflowSolutionDefinition) HasLicenseEntitlement() bool`

HasLicenseEntitlement returns a boolean if a field has been set.

### GetName

`func (o *WorkflowSolutionDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowSolutionDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowSolutionDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowSolutionDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutputDefinition

`func (o *WorkflowSolutionDefinition) GetOutputDefinition() []WorkflowBaseDataType`

GetOutputDefinition returns the OutputDefinition field if non-nil, zero value otherwise.

### GetOutputDefinitionOk

`func (o *WorkflowSolutionDefinition) GetOutputDefinitionOk() (*[]WorkflowBaseDataType, bool)`

GetOutputDefinitionOk returns a tuple with the OutputDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputDefinition

`func (o *WorkflowSolutionDefinition) SetOutputDefinition(v []WorkflowBaseDataType)`

SetOutputDefinition sets OutputDefinition field to given value.

### HasOutputDefinition

`func (o *WorkflowSolutionDefinition) HasOutputDefinition() bool`

HasOutputDefinition returns a boolean if a field has been set.

### SetOutputDefinitionNil

`func (o *WorkflowSolutionDefinition) SetOutputDefinitionNil(b bool)`

 SetOutputDefinitionNil sets the value for OutputDefinition to be an explicit nil

### UnsetOutputDefinition
`func (o *WorkflowSolutionDefinition) UnsetOutputDefinition()`

UnsetOutputDefinition ensures that no value is present for OutputDefinition, not even an explicit nil
### GetVersion

`func (o *WorkflowSolutionDefinition) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkflowSolutionDefinition) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkflowSolutionDefinition) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WorkflowSolutionDefinition) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetActionDefinitions

`func (o *WorkflowSolutionDefinition) GetActionDefinitions() []WorkflowSolutionActionDefinitionRelationship`

GetActionDefinitions returns the ActionDefinitions field if non-nil, zero value otherwise.

### GetActionDefinitionsOk

`func (o *WorkflowSolutionDefinition) GetActionDefinitionsOk() (*[]WorkflowSolutionActionDefinitionRelationship, bool)`

GetActionDefinitionsOk returns a tuple with the ActionDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionDefinitions

`func (o *WorkflowSolutionDefinition) SetActionDefinitions(v []WorkflowSolutionActionDefinitionRelationship)`

SetActionDefinitions sets ActionDefinitions field to given value.

### HasActionDefinitions

`func (o *WorkflowSolutionDefinition) HasActionDefinitions() bool`

HasActionDefinitions returns a boolean if a field has been set.

### SetActionDefinitionsNil

`func (o *WorkflowSolutionDefinition) SetActionDefinitionsNil(b bool)`

 SetActionDefinitionsNil sets the value for ActionDefinitions to be an explicit nil

### UnsetActionDefinitions
`func (o *WorkflowSolutionDefinition) UnsetActionDefinitions()`

UnsetActionDefinitions ensures that no value is present for ActionDefinitions, not even an explicit nil
### GetCatalog

`func (o *WorkflowSolutionDefinition) GetCatalog() WorkflowCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *WorkflowSolutionDefinition) GetCatalogOk() (*WorkflowCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *WorkflowSolutionDefinition) SetCatalog(v WorkflowCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *WorkflowSolutionDefinition) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


