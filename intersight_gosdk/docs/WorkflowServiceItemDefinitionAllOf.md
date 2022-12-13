# WorkflowServiceItemDefinitionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.ServiceItemDefinition"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.ServiceItemDefinition"]
**AllowMultipleServiceItemInstances** | Pointer to **bool** | Service item definition can declare that only one instance can be allowed within the customer account. | [optional] [default to true]
**CvdId** | Pointer to **string** | The Cisco Validated Design (CVD) Identifier that this service item provides. | [optional] 
**DeleteInstanceOnDecommission** | Pointer to **bool** | The flag to indicate that service item instance will be deleted after the completion of decommission action. | [optional] [default to false]
**Description** | Pointer to **string** | The description for this service item which provides information on what are the pre-requisites to deploy the service item and what features are supported on the service item. | [optional] 
**Label** | Pointer to **string** | A user friendly short name to identify the service item. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_). | [optional] 
**LicenseEntitlement** | Pointer to **string** | License entitlement required to run this service item. * &#x60;Base&#x60; - Base as a License type. It is default license type. * &#x60;Essential&#x60; - Essential as a License type. * &#x60;Standard&#x60; - Standard as a License type. * &#x60;Advantage&#x60; - Advantage as a License type. * &#x60;Premier&#x60; - Premier as a License type. * &#x60;IWO-Essential&#x60; - IWO-Essential as a License type. * &#x60;IWO-Advantage&#x60; - IWO-Advantage as a License type. * &#x60;IWO-Premier&#x60; - IWO-Premier as a License type. * &#x60;IKS-Advantage&#x60; - IKS-Advantage as a License type. * &#x60;INC-Premier-1GFixed&#x60; - Premier 1G Fixed license tier for Intersight Nexus Cloud. * &#x60;INC-Premier-10GFixed&#x60; - Premier 10G Fixed license tier for Intersight Nexus Cloud. * &#x60;INC-Premier-100GFixed&#x60; - Premier 100G Fixed license tier for Intersight Nexus Cloud. * &#x60;INC-Premier-Mod4Slot&#x60; - Premier Modular 4 slot license tier for Intersight Nexus Cloud. * &#x60;INC-Premier-Mod8Slot&#x60; - Premier Modular 8 slot license tier for Intersight Nexus Cloud. * &#x60;INC-Premier-D2OpsFixed&#x60; - Premier D2Ops fixed license tier for Intersight Nexus Cloud. * &#x60;INC-Premier-D2OpsMod&#x60; - Premier D2Ops modular license tier for Intersight Nexus Cloud. * &#x60;IntersightTrial&#x60; - Virtual dummy license type to indicate trial. Used for UI display of trial mode Intersight tiers. * &#x60;IWOTrial&#x60; - Virtual dummy license type to indicate trial. Used for UI display of trial mode IKS tiers. * &#x60;IKSTrial&#x60; - Virtual dummy license type to indicate trial. Used for UI display of trial mode IWO tiers. * &#x60;INCTrial&#x60; - Virtual dummy license type to indicate trial. Used for UI display of trial mode Nexus tiers. | [optional] [readonly] [default to "Base"]
**Name** | Pointer to **string** | The name for this service item definition. You can have multiple versions of the service item with the same name. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:) or an underscore (_). | [optional] 
**OutputDefinition** | Pointer to [**[]WorkflowBaseDataType**](WorkflowBaseDataType.md) |  | [optional] 
**Version** | Pointer to **int64** | The version of the service item to support multiple versions. | [optional] [default to 1]
**ActionDefinitions** | Pointer to [**[]WorkflowServiceItemActionDefinitionRelationship**](WorkflowServiceItemActionDefinitionRelationship.md) | An array of relationships to workflowServiceItemActionDefinition resources. | [optional] 
**Catalog** | Pointer to [**WorkflowCatalogRelationship**](WorkflowCatalogRelationship.md) |  | [optional] 

## Methods

### NewWorkflowServiceItemDefinitionAllOf

`func NewWorkflowServiceItemDefinitionAllOf(classId string, objectType string, ) *WorkflowServiceItemDefinitionAllOf`

NewWorkflowServiceItemDefinitionAllOf instantiates a new WorkflowServiceItemDefinitionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowServiceItemDefinitionAllOfWithDefaults

`func NewWorkflowServiceItemDefinitionAllOfWithDefaults() *WorkflowServiceItemDefinitionAllOf`

NewWorkflowServiceItemDefinitionAllOfWithDefaults instantiates a new WorkflowServiceItemDefinitionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowServiceItemDefinitionAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowServiceItemDefinitionAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowServiceItemDefinitionAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowServiceItemDefinitionAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowServiceItemDefinitionAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowServiceItemDefinitionAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAllowMultipleServiceItemInstances

`func (o *WorkflowServiceItemDefinitionAllOf) GetAllowMultipleServiceItemInstances() bool`

GetAllowMultipleServiceItemInstances returns the AllowMultipleServiceItemInstances field if non-nil, zero value otherwise.

### GetAllowMultipleServiceItemInstancesOk

`func (o *WorkflowServiceItemDefinitionAllOf) GetAllowMultipleServiceItemInstancesOk() (*bool, bool)`

GetAllowMultipleServiceItemInstancesOk returns a tuple with the AllowMultipleServiceItemInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultipleServiceItemInstances

`func (o *WorkflowServiceItemDefinitionAllOf) SetAllowMultipleServiceItemInstances(v bool)`

SetAllowMultipleServiceItemInstances sets AllowMultipleServiceItemInstances field to given value.

### HasAllowMultipleServiceItemInstances

`func (o *WorkflowServiceItemDefinitionAllOf) HasAllowMultipleServiceItemInstances() bool`

HasAllowMultipleServiceItemInstances returns a boolean if a field has been set.

### GetCvdId

`func (o *WorkflowServiceItemDefinitionAllOf) GetCvdId() string`

GetCvdId returns the CvdId field if non-nil, zero value otherwise.

### GetCvdIdOk

`func (o *WorkflowServiceItemDefinitionAllOf) GetCvdIdOk() (*string, bool)`

GetCvdIdOk returns a tuple with the CvdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvdId

`func (o *WorkflowServiceItemDefinitionAllOf) SetCvdId(v string)`

SetCvdId sets CvdId field to given value.

### HasCvdId

`func (o *WorkflowServiceItemDefinitionAllOf) HasCvdId() bool`

HasCvdId returns a boolean if a field has been set.

### GetDeleteInstanceOnDecommission

`func (o *WorkflowServiceItemDefinitionAllOf) GetDeleteInstanceOnDecommission() bool`

GetDeleteInstanceOnDecommission returns the DeleteInstanceOnDecommission field if non-nil, zero value otherwise.

### GetDeleteInstanceOnDecommissionOk

`func (o *WorkflowServiceItemDefinitionAllOf) GetDeleteInstanceOnDecommissionOk() (*bool, bool)`

GetDeleteInstanceOnDecommissionOk returns a tuple with the DeleteInstanceOnDecommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteInstanceOnDecommission

`func (o *WorkflowServiceItemDefinitionAllOf) SetDeleteInstanceOnDecommission(v bool)`

SetDeleteInstanceOnDecommission sets DeleteInstanceOnDecommission field to given value.

### HasDeleteInstanceOnDecommission

`func (o *WorkflowServiceItemDefinitionAllOf) HasDeleteInstanceOnDecommission() bool`

HasDeleteInstanceOnDecommission returns a boolean if a field has been set.

### GetDescription

`func (o *WorkflowServiceItemDefinitionAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowServiceItemDefinitionAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowServiceItemDefinitionAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowServiceItemDefinitionAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabel

`func (o *WorkflowServiceItemDefinitionAllOf) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WorkflowServiceItemDefinitionAllOf) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WorkflowServiceItemDefinitionAllOf) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WorkflowServiceItemDefinitionAllOf) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLicenseEntitlement

`func (o *WorkflowServiceItemDefinitionAllOf) GetLicenseEntitlement() string`

GetLicenseEntitlement returns the LicenseEntitlement field if non-nil, zero value otherwise.

### GetLicenseEntitlementOk

`func (o *WorkflowServiceItemDefinitionAllOf) GetLicenseEntitlementOk() (*string, bool)`

GetLicenseEntitlementOk returns a tuple with the LicenseEntitlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseEntitlement

`func (o *WorkflowServiceItemDefinitionAllOf) SetLicenseEntitlement(v string)`

SetLicenseEntitlement sets LicenseEntitlement field to given value.

### HasLicenseEntitlement

`func (o *WorkflowServiceItemDefinitionAllOf) HasLicenseEntitlement() bool`

HasLicenseEntitlement returns a boolean if a field has been set.

### GetName

`func (o *WorkflowServiceItemDefinitionAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowServiceItemDefinitionAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowServiceItemDefinitionAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowServiceItemDefinitionAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutputDefinition

`func (o *WorkflowServiceItemDefinitionAllOf) GetOutputDefinition() []WorkflowBaseDataType`

GetOutputDefinition returns the OutputDefinition field if non-nil, zero value otherwise.

### GetOutputDefinitionOk

`func (o *WorkflowServiceItemDefinitionAllOf) GetOutputDefinitionOk() (*[]WorkflowBaseDataType, bool)`

GetOutputDefinitionOk returns a tuple with the OutputDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputDefinition

`func (o *WorkflowServiceItemDefinitionAllOf) SetOutputDefinition(v []WorkflowBaseDataType)`

SetOutputDefinition sets OutputDefinition field to given value.

### HasOutputDefinition

`func (o *WorkflowServiceItemDefinitionAllOf) HasOutputDefinition() bool`

HasOutputDefinition returns a boolean if a field has been set.

### SetOutputDefinitionNil

`func (o *WorkflowServiceItemDefinitionAllOf) SetOutputDefinitionNil(b bool)`

 SetOutputDefinitionNil sets the value for OutputDefinition to be an explicit nil

### UnsetOutputDefinition
`func (o *WorkflowServiceItemDefinitionAllOf) UnsetOutputDefinition()`

UnsetOutputDefinition ensures that no value is present for OutputDefinition, not even an explicit nil
### GetVersion

`func (o *WorkflowServiceItemDefinitionAllOf) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkflowServiceItemDefinitionAllOf) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkflowServiceItemDefinitionAllOf) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WorkflowServiceItemDefinitionAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetActionDefinitions

`func (o *WorkflowServiceItemDefinitionAllOf) GetActionDefinitions() []WorkflowServiceItemActionDefinitionRelationship`

GetActionDefinitions returns the ActionDefinitions field if non-nil, zero value otherwise.

### GetActionDefinitionsOk

`func (o *WorkflowServiceItemDefinitionAllOf) GetActionDefinitionsOk() (*[]WorkflowServiceItemActionDefinitionRelationship, bool)`

GetActionDefinitionsOk returns a tuple with the ActionDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionDefinitions

`func (o *WorkflowServiceItemDefinitionAllOf) SetActionDefinitions(v []WorkflowServiceItemActionDefinitionRelationship)`

SetActionDefinitions sets ActionDefinitions field to given value.

### HasActionDefinitions

`func (o *WorkflowServiceItemDefinitionAllOf) HasActionDefinitions() bool`

HasActionDefinitions returns a boolean if a field has been set.

### SetActionDefinitionsNil

`func (o *WorkflowServiceItemDefinitionAllOf) SetActionDefinitionsNil(b bool)`

 SetActionDefinitionsNil sets the value for ActionDefinitions to be an explicit nil

### UnsetActionDefinitions
`func (o *WorkflowServiceItemDefinitionAllOf) UnsetActionDefinitions()`

UnsetActionDefinitions ensures that no value is present for ActionDefinitions, not even an explicit nil
### GetCatalog

`func (o *WorkflowServiceItemDefinitionAllOf) GetCatalog() WorkflowCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *WorkflowServiceItemDefinitionAllOf) GetCatalogOk() (*WorkflowCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *WorkflowServiceItemDefinitionAllOf) SetCatalog(v WorkflowCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *WorkflowServiceItemDefinitionAllOf) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


