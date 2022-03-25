# WorkflowSolutionDefinitionAllOf

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
**UpgradedMoid** | Pointer to **string** | Stores the upgraded Moid for help during future lookups. | [optional] [readonly] 
**Version** | Pointer to **int64** | The version of the solution to support multiple versions. | [optional] [default to 1]
**ActionDefinitions** | Pointer to [**[]WorkflowSolutionActionDefinitionRelationship**](WorkflowSolutionActionDefinitionRelationship.md) | An array of relationships to workflowSolutionActionDefinition resources. | [optional] 
**Catalog** | Pointer to [**WorkflowCatalogRelationship**](WorkflowCatalogRelationship.md) |  | [optional] 

## Methods

### NewWorkflowSolutionDefinitionAllOf

`func NewWorkflowSolutionDefinitionAllOf(classId string, objectType string, ) *WorkflowSolutionDefinitionAllOf`

NewWorkflowSolutionDefinitionAllOf instantiates a new WorkflowSolutionDefinitionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSolutionDefinitionAllOfWithDefaults

`func NewWorkflowSolutionDefinitionAllOfWithDefaults() *WorkflowSolutionDefinitionAllOf`

NewWorkflowSolutionDefinitionAllOfWithDefaults instantiates a new WorkflowSolutionDefinitionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowSolutionDefinitionAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowSolutionDefinitionAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowSolutionDefinitionAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowSolutionDefinitionAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowSolutionDefinitionAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowSolutionDefinitionAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAllowMultipleSolutionInstances

`func (o *WorkflowSolutionDefinitionAllOf) GetAllowMultipleSolutionInstances() bool`

GetAllowMultipleSolutionInstances returns the AllowMultipleSolutionInstances field if non-nil, zero value otherwise.

### GetAllowMultipleSolutionInstancesOk

`func (o *WorkflowSolutionDefinitionAllOf) GetAllowMultipleSolutionInstancesOk() (*bool, bool)`

GetAllowMultipleSolutionInstancesOk returns a tuple with the AllowMultipleSolutionInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultipleSolutionInstances

`func (o *WorkflowSolutionDefinitionAllOf) SetAllowMultipleSolutionInstances(v bool)`

SetAllowMultipleSolutionInstances sets AllowMultipleSolutionInstances field to given value.

### HasAllowMultipleSolutionInstances

`func (o *WorkflowSolutionDefinitionAllOf) HasAllowMultipleSolutionInstances() bool`

HasAllowMultipleSolutionInstances returns a boolean if a field has been set.

### GetCvdId

`func (o *WorkflowSolutionDefinitionAllOf) GetCvdId() string`

GetCvdId returns the CvdId field if non-nil, zero value otherwise.

### GetCvdIdOk

`func (o *WorkflowSolutionDefinitionAllOf) GetCvdIdOk() (*string, bool)`

GetCvdIdOk returns a tuple with the CvdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvdId

`func (o *WorkflowSolutionDefinitionAllOf) SetCvdId(v string)`

SetCvdId sets CvdId field to given value.

### HasCvdId

`func (o *WorkflowSolutionDefinitionAllOf) HasCvdId() bool`

HasCvdId returns a boolean if a field has been set.

### GetDeleteInstanceOnDecommission

`func (o *WorkflowSolutionDefinitionAllOf) GetDeleteInstanceOnDecommission() bool`

GetDeleteInstanceOnDecommission returns the DeleteInstanceOnDecommission field if non-nil, zero value otherwise.

### GetDeleteInstanceOnDecommissionOk

`func (o *WorkflowSolutionDefinitionAllOf) GetDeleteInstanceOnDecommissionOk() (*bool, bool)`

GetDeleteInstanceOnDecommissionOk returns a tuple with the DeleteInstanceOnDecommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteInstanceOnDecommission

`func (o *WorkflowSolutionDefinitionAllOf) SetDeleteInstanceOnDecommission(v bool)`

SetDeleteInstanceOnDecommission sets DeleteInstanceOnDecommission field to given value.

### HasDeleteInstanceOnDecommission

`func (o *WorkflowSolutionDefinitionAllOf) HasDeleteInstanceOnDecommission() bool`

HasDeleteInstanceOnDecommission returns a boolean if a field has been set.

### GetDescription

`func (o *WorkflowSolutionDefinitionAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowSolutionDefinitionAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowSolutionDefinitionAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowSolutionDefinitionAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabel

`func (o *WorkflowSolutionDefinitionAllOf) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WorkflowSolutionDefinitionAllOf) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WorkflowSolutionDefinitionAllOf) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WorkflowSolutionDefinitionAllOf) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLicenseEntitlement

`func (o *WorkflowSolutionDefinitionAllOf) GetLicenseEntitlement() string`

GetLicenseEntitlement returns the LicenseEntitlement field if non-nil, zero value otherwise.

### GetLicenseEntitlementOk

`func (o *WorkflowSolutionDefinitionAllOf) GetLicenseEntitlementOk() (*string, bool)`

GetLicenseEntitlementOk returns a tuple with the LicenseEntitlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseEntitlement

`func (o *WorkflowSolutionDefinitionAllOf) SetLicenseEntitlement(v string)`

SetLicenseEntitlement sets LicenseEntitlement field to given value.

### HasLicenseEntitlement

`func (o *WorkflowSolutionDefinitionAllOf) HasLicenseEntitlement() bool`

HasLicenseEntitlement returns a boolean if a field has been set.

### GetName

`func (o *WorkflowSolutionDefinitionAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowSolutionDefinitionAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowSolutionDefinitionAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowSolutionDefinitionAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutputDefinition

`func (o *WorkflowSolutionDefinitionAllOf) GetOutputDefinition() []WorkflowBaseDataType`

GetOutputDefinition returns the OutputDefinition field if non-nil, zero value otherwise.

### GetOutputDefinitionOk

`func (o *WorkflowSolutionDefinitionAllOf) GetOutputDefinitionOk() (*[]WorkflowBaseDataType, bool)`

GetOutputDefinitionOk returns a tuple with the OutputDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputDefinition

`func (o *WorkflowSolutionDefinitionAllOf) SetOutputDefinition(v []WorkflowBaseDataType)`

SetOutputDefinition sets OutputDefinition field to given value.

### HasOutputDefinition

`func (o *WorkflowSolutionDefinitionAllOf) HasOutputDefinition() bool`

HasOutputDefinition returns a boolean if a field has been set.

### SetOutputDefinitionNil

`func (o *WorkflowSolutionDefinitionAllOf) SetOutputDefinitionNil(b bool)`

 SetOutputDefinitionNil sets the value for OutputDefinition to be an explicit nil

### UnsetOutputDefinition
`func (o *WorkflowSolutionDefinitionAllOf) UnsetOutputDefinition()`

UnsetOutputDefinition ensures that no value is present for OutputDefinition, not even an explicit nil
### GetUpgradedMoid

`func (o *WorkflowSolutionDefinitionAllOf) GetUpgradedMoid() string`

GetUpgradedMoid returns the UpgradedMoid field if non-nil, zero value otherwise.

### GetUpgradedMoidOk

`func (o *WorkflowSolutionDefinitionAllOf) GetUpgradedMoidOk() (*string, bool)`

GetUpgradedMoidOk returns a tuple with the UpgradedMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradedMoid

`func (o *WorkflowSolutionDefinitionAllOf) SetUpgradedMoid(v string)`

SetUpgradedMoid sets UpgradedMoid field to given value.

### HasUpgradedMoid

`func (o *WorkflowSolutionDefinitionAllOf) HasUpgradedMoid() bool`

HasUpgradedMoid returns a boolean if a field has been set.

### GetVersion

`func (o *WorkflowSolutionDefinitionAllOf) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkflowSolutionDefinitionAllOf) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkflowSolutionDefinitionAllOf) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WorkflowSolutionDefinitionAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetActionDefinitions

`func (o *WorkflowSolutionDefinitionAllOf) GetActionDefinitions() []WorkflowSolutionActionDefinitionRelationship`

GetActionDefinitions returns the ActionDefinitions field if non-nil, zero value otherwise.

### GetActionDefinitionsOk

`func (o *WorkflowSolutionDefinitionAllOf) GetActionDefinitionsOk() (*[]WorkflowSolutionActionDefinitionRelationship, bool)`

GetActionDefinitionsOk returns a tuple with the ActionDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionDefinitions

`func (o *WorkflowSolutionDefinitionAllOf) SetActionDefinitions(v []WorkflowSolutionActionDefinitionRelationship)`

SetActionDefinitions sets ActionDefinitions field to given value.

### HasActionDefinitions

`func (o *WorkflowSolutionDefinitionAllOf) HasActionDefinitions() bool`

HasActionDefinitions returns a boolean if a field has been set.

### SetActionDefinitionsNil

`func (o *WorkflowSolutionDefinitionAllOf) SetActionDefinitionsNil(b bool)`

 SetActionDefinitionsNil sets the value for ActionDefinitions to be an explicit nil

### UnsetActionDefinitions
`func (o *WorkflowSolutionDefinitionAllOf) UnsetActionDefinitions()`

UnsetActionDefinitions ensures that no value is present for ActionDefinitions, not even an explicit nil
### GetCatalog

`func (o *WorkflowSolutionDefinitionAllOf) GetCatalog() WorkflowCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *WorkflowSolutionDefinitionAllOf) GetCatalogOk() (*WorkflowCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *WorkflowSolutionDefinitionAllOf) SetCatalog(v WorkflowCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *WorkflowSolutionDefinitionAllOf) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


