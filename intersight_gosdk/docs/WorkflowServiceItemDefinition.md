# WorkflowServiceItemDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.ServiceItemDefinition"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.ServiceItemDefinition"]
**AllowMultipleServiceItemInstances** | Pointer to **bool** | Service item definition can declare that only one instance can be allowed within the customer account. | [optional] [default to true]
**AttributeDefinition** | Pointer to [**[]WorkflowBaseDataType**](WorkflowBaseDataType.md) |  | [optional] 
**CreateUser** | Pointer to **string** | The user identifier who created or cloned the service item definition. | [optional] [readonly] 
**CvdId** | Pointer to **string** | The Cisco Validated Design (CVD) Identifier that this service item provides. | [optional] 
**DeleteInstanceOnDecommission** | Pointer to **bool** | The flag to indicate that service item instance will be deleted after the completion of decommission action. | [optional] [default to false]
**Description** | Pointer to **string** | The description for this service item which provides information on what are the pre-requisites to deploy the service item and what features are supported on the service item. | [optional] 
**Label** | Pointer to **string** | A user friendly short name to identify the service item. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_). | [optional] 
**LicenseEntitlement** | Pointer to **string** | License entitlement required to run this service item. * &#x60;Base&#x60; - Base as a License type. It is default license type. * &#x60;Essential&#x60; - Essential as a License type. * &#x60;Standard&#x60; - Standard as a License type. * &#x60;Advantage&#x60; - Advantage as a License type. * &#x60;Premier&#x60; - Premier as a License type. * &#x60;IWO-Essential&#x60; - IWO-Essential as a License type. * &#x60;IWO-Advantage&#x60; - IWO-Advantage as a License type. * &#x60;IWO-Premier&#x60; - IWO-Premier as a License type. * &#x60;IKS-Advantage&#x60; - IKS-Advantage as a License type. * &#x60;INC-Premier-1GFixed&#x60; - Premier 1G Fixed license tier for Intersight Nexus Cloud. * &#x60;INC-Premier-10GFixed&#x60; - Premier 10G Fixed license tier for Intersight Nexus Cloud. * &#x60;INC-Premier-100GFixed&#x60; - Premier 100G Fixed license tier for Intersight Nexus Cloud. * &#x60;INC-Premier-Mod4Slot&#x60; - Premier Modular 4 slot license tier for Intersight Nexus Cloud. * &#x60;INC-Premier-Mod8Slot&#x60; - Premier Modular 8 slot license tier for Intersight Nexus Cloud. * &#x60;INC-Premier-D2OpsFixed&#x60; - Premier D2Ops fixed license tier for Intersight Nexus Cloud. * &#x60;INC-Premier-D2OpsMod&#x60; - Premier D2Ops modular license tier for Intersight Nexus Cloud. * &#x60;INC-Premier-CentralizedMod8Slot&#x60; - Premier modular license tier of switch type CentralizedMod8Slot for Intersight Nexus Cloud. * &#x60;INC-Premier-DistributedMod8Slot&#x60; - Premier modular license tier of switch type DistributedMod8Slot for Intersight Nexus Cloud. * &#x60;IntersightTrial&#x60; - Virtual dummy license type to indicate trial. Used for UI display of trial mode Intersight tiers. * &#x60;IWOTrial&#x60; - Virtual dummy license type to indicate trial. Used for UI display of trial mode IKS tiers. * &#x60;IKSTrial&#x60; - Virtual dummy license type to indicate trial. Used for UI display of trial mode IWO tiers. * &#x60;INCTrial&#x60; - Virtual dummy license type to indicate trial. Used for UI display of trial mode Nexus tiers. | [optional] [readonly] [default to "Base"]
**ModUser** | Pointer to **string** | The user identifier who last updated the service item definition. | [optional] [readonly] 
**Name** | Pointer to **string** | The name for this service item definition. You can have multiple versions of the service item with the same name. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:) or an underscore (_). | [optional] 
**PublishStatus** | Pointer to **string** | The publish status of service item (Draft, Published, Archived). * &#x60;Draft&#x60; - The enum specifies the option as Draft which means the meta definition is being designed and tested. * &#x60;Published&#x60; - The enum specifies the option as Published which means the meta definition is ready for consumption. * &#x60;Archived&#x60; - The enum specifies the option as Archived which means the meta definition is archived and can no longer be consumed. | [optional] [default to "Draft"]
**Status** | Pointer to **string** | State of service item considering the state of underlying service item actions definitions. * &#x60;Okay&#x60; - Deployment and other post-deployment actions are in valid state. * &#x60;Critical&#x60; - Deployment action is not in valid state. * &#x60;Warning&#x60; - Deployment action is in valid state, and one or more post-deployment actions are not in valid state. | [optional] [readonly] [default to "Okay"]
**SupportStatus** | Pointer to **string** | The service item can be marked as deprecated, supported or beta, the support status indicates that. When a new service item is introduced, it can be marked beta to indicate this is experimental and later moved to Supported status. When Service item is deprecated, it cannot be instantiated and used for a Catalog Item design. * &#x60;Supported&#x60; - The definition is a supported version and there will be no changes to the mandatory inputs or outputs. * &#x60;Beta&#x60; - The definition is a Beta version and this version can under go changes until the version is marked supported. * &#x60;Deprecated&#x60; - The version of definition is deprecated and typically there will be a higher version of the same definition that has been added. | [optional] [default to "Supported"]
**UserIdOrEmail** | Pointer to **string** | This attribute is deprecated and replaced by createUser and modUser. | [optional] [readonly] 
**ValidationInformation** | Pointer to [**NullableWorkflowValidationInformation**](WorkflowValidationInformation.md) |  | [optional] 
**Version** | Pointer to **int64** | The version of the service item to support multiple versions. | [optional] [default to 1]
**ActionDefinitions** | Pointer to [**[]WorkflowServiceItemActionDefinitionRelationship**](WorkflowServiceItemActionDefinitionRelationship.md) | An array of relationships to workflowServiceItemActionDefinition resources. | [optional] 
**Catalog** | Pointer to [**NullableWorkflowCatalogRelationship**](WorkflowCatalogRelationship.md) |  | [optional] 

## Methods

### NewWorkflowServiceItemDefinition

`func NewWorkflowServiceItemDefinition(classId string, objectType string, ) *WorkflowServiceItemDefinition`

NewWorkflowServiceItemDefinition instantiates a new WorkflowServiceItemDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowServiceItemDefinitionWithDefaults

`func NewWorkflowServiceItemDefinitionWithDefaults() *WorkflowServiceItemDefinition`

NewWorkflowServiceItemDefinitionWithDefaults instantiates a new WorkflowServiceItemDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowServiceItemDefinition) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowServiceItemDefinition) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowServiceItemDefinition) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowServiceItemDefinition) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowServiceItemDefinition) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowServiceItemDefinition) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAllowMultipleServiceItemInstances

`func (o *WorkflowServiceItemDefinition) GetAllowMultipleServiceItemInstances() bool`

GetAllowMultipleServiceItemInstances returns the AllowMultipleServiceItemInstances field if non-nil, zero value otherwise.

### GetAllowMultipleServiceItemInstancesOk

`func (o *WorkflowServiceItemDefinition) GetAllowMultipleServiceItemInstancesOk() (*bool, bool)`

GetAllowMultipleServiceItemInstancesOk returns a tuple with the AllowMultipleServiceItemInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultipleServiceItemInstances

`func (o *WorkflowServiceItemDefinition) SetAllowMultipleServiceItemInstances(v bool)`

SetAllowMultipleServiceItemInstances sets AllowMultipleServiceItemInstances field to given value.

### HasAllowMultipleServiceItemInstances

`func (o *WorkflowServiceItemDefinition) HasAllowMultipleServiceItemInstances() bool`

HasAllowMultipleServiceItemInstances returns a boolean if a field has been set.

### GetAttributeDefinition

`func (o *WorkflowServiceItemDefinition) GetAttributeDefinition() []WorkflowBaseDataType`

GetAttributeDefinition returns the AttributeDefinition field if non-nil, zero value otherwise.

### GetAttributeDefinitionOk

`func (o *WorkflowServiceItemDefinition) GetAttributeDefinitionOk() (*[]WorkflowBaseDataType, bool)`

GetAttributeDefinitionOk returns a tuple with the AttributeDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeDefinition

`func (o *WorkflowServiceItemDefinition) SetAttributeDefinition(v []WorkflowBaseDataType)`

SetAttributeDefinition sets AttributeDefinition field to given value.

### HasAttributeDefinition

`func (o *WorkflowServiceItemDefinition) HasAttributeDefinition() bool`

HasAttributeDefinition returns a boolean if a field has been set.

### SetAttributeDefinitionNil

`func (o *WorkflowServiceItemDefinition) SetAttributeDefinitionNil(b bool)`

 SetAttributeDefinitionNil sets the value for AttributeDefinition to be an explicit nil

### UnsetAttributeDefinition
`func (o *WorkflowServiceItemDefinition) UnsetAttributeDefinition()`

UnsetAttributeDefinition ensures that no value is present for AttributeDefinition, not even an explicit nil
### GetCreateUser

`func (o *WorkflowServiceItemDefinition) GetCreateUser() string`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *WorkflowServiceItemDefinition) GetCreateUserOk() (*string, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *WorkflowServiceItemDefinition) SetCreateUser(v string)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *WorkflowServiceItemDefinition) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetCvdId

`func (o *WorkflowServiceItemDefinition) GetCvdId() string`

GetCvdId returns the CvdId field if non-nil, zero value otherwise.

### GetCvdIdOk

`func (o *WorkflowServiceItemDefinition) GetCvdIdOk() (*string, bool)`

GetCvdIdOk returns a tuple with the CvdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvdId

`func (o *WorkflowServiceItemDefinition) SetCvdId(v string)`

SetCvdId sets CvdId field to given value.

### HasCvdId

`func (o *WorkflowServiceItemDefinition) HasCvdId() bool`

HasCvdId returns a boolean if a field has been set.

### GetDeleteInstanceOnDecommission

`func (o *WorkflowServiceItemDefinition) GetDeleteInstanceOnDecommission() bool`

GetDeleteInstanceOnDecommission returns the DeleteInstanceOnDecommission field if non-nil, zero value otherwise.

### GetDeleteInstanceOnDecommissionOk

`func (o *WorkflowServiceItemDefinition) GetDeleteInstanceOnDecommissionOk() (*bool, bool)`

GetDeleteInstanceOnDecommissionOk returns a tuple with the DeleteInstanceOnDecommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteInstanceOnDecommission

`func (o *WorkflowServiceItemDefinition) SetDeleteInstanceOnDecommission(v bool)`

SetDeleteInstanceOnDecommission sets DeleteInstanceOnDecommission field to given value.

### HasDeleteInstanceOnDecommission

`func (o *WorkflowServiceItemDefinition) HasDeleteInstanceOnDecommission() bool`

HasDeleteInstanceOnDecommission returns a boolean if a field has been set.

### GetDescription

`func (o *WorkflowServiceItemDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowServiceItemDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowServiceItemDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowServiceItemDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabel

`func (o *WorkflowServiceItemDefinition) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WorkflowServiceItemDefinition) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WorkflowServiceItemDefinition) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WorkflowServiceItemDefinition) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLicenseEntitlement

`func (o *WorkflowServiceItemDefinition) GetLicenseEntitlement() string`

GetLicenseEntitlement returns the LicenseEntitlement field if non-nil, zero value otherwise.

### GetLicenseEntitlementOk

`func (o *WorkflowServiceItemDefinition) GetLicenseEntitlementOk() (*string, bool)`

GetLicenseEntitlementOk returns a tuple with the LicenseEntitlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseEntitlement

`func (o *WorkflowServiceItemDefinition) SetLicenseEntitlement(v string)`

SetLicenseEntitlement sets LicenseEntitlement field to given value.

### HasLicenseEntitlement

`func (o *WorkflowServiceItemDefinition) HasLicenseEntitlement() bool`

HasLicenseEntitlement returns a boolean if a field has been set.

### GetModUser

`func (o *WorkflowServiceItemDefinition) GetModUser() string`

GetModUser returns the ModUser field if non-nil, zero value otherwise.

### GetModUserOk

`func (o *WorkflowServiceItemDefinition) GetModUserOk() (*string, bool)`

GetModUserOk returns a tuple with the ModUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModUser

`func (o *WorkflowServiceItemDefinition) SetModUser(v string)`

SetModUser sets ModUser field to given value.

### HasModUser

`func (o *WorkflowServiceItemDefinition) HasModUser() bool`

HasModUser returns a boolean if a field has been set.

### GetName

`func (o *WorkflowServiceItemDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowServiceItemDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowServiceItemDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowServiceItemDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublishStatus

`func (o *WorkflowServiceItemDefinition) GetPublishStatus() string`

GetPublishStatus returns the PublishStatus field if non-nil, zero value otherwise.

### GetPublishStatusOk

`func (o *WorkflowServiceItemDefinition) GetPublishStatusOk() (*string, bool)`

GetPublishStatusOk returns a tuple with the PublishStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishStatus

`func (o *WorkflowServiceItemDefinition) SetPublishStatus(v string)`

SetPublishStatus sets PublishStatus field to given value.

### HasPublishStatus

`func (o *WorkflowServiceItemDefinition) HasPublishStatus() bool`

HasPublishStatus returns a boolean if a field has been set.

### GetStatus

`func (o *WorkflowServiceItemDefinition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowServiceItemDefinition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowServiceItemDefinition) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowServiceItemDefinition) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSupportStatus

`func (o *WorkflowServiceItemDefinition) GetSupportStatus() string`

GetSupportStatus returns the SupportStatus field if non-nil, zero value otherwise.

### GetSupportStatusOk

`func (o *WorkflowServiceItemDefinition) GetSupportStatusOk() (*string, bool)`

GetSupportStatusOk returns a tuple with the SupportStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportStatus

`func (o *WorkflowServiceItemDefinition) SetSupportStatus(v string)`

SetSupportStatus sets SupportStatus field to given value.

### HasSupportStatus

`func (o *WorkflowServiceItemDefinition) HasSupportStatus() bool`

HasSupportStatus returns a boolean if a field has been set.

### GetUserIdOrEmail

`func (o *WorkflowServiceItemDefinition) GetUserIdOrEmail() string`

GetUserIdOrEmail returns the UserIdOrEmail field if non-nil, zero value otherwise.

### GetUserIdOrEmailOk

`func (o *WorkflowServiceItemDefinition) GetUserIdOrEmailOk() (*string, bool)`

GetUserIdOrEmailOk returns a tuple with the UserIdOrEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdOrEmail

`func (o *WorkflowServiceItemDefinition) SetUserIdOrEmail(v string)`

SetUserIdOrEmail sets UserIdOrEmail field to given value.

### HasUserIdOrEmail

`func (o *WorkflowServiceItemDefinition) HasUserIdOrEmail() bool`

HasUserIdOrEmail returns a boolean if a field has been set.

### GetValidationInformation

`func (o *WorkflowServiceItemDefinition) GetValidationInformation() WorkflowValidationInformation`

GetValidationInformation returns the ValidationInformation field if non-nil, zero value otherwise.

### GetValidationInformationOk

`func (o *WorkflowServiceItemDefinition) GetValidationInformationOk() (*WorkflowValidationInformation, bool)`

GetValidationInformationOk returns a tuple with the ValidationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationInformation

`func (o *WorkflowServiceItemDefinition) SetValidationInformation(v WorkflowValidationInformation)`

SetValidationInformation sets ValidationInformation field to given value.

### HasValidationInformation

`func (o *WorkflowServiceItemDefinition) HasValidationInformation() bool`

HasValidationInformation returns a boolean if a field has been set.

### SetValidationInformationNil

`func (o *WorkflowServiceItemDefinition) SetValidationInformationNil(b bool)`

 SetValidationInformationNil sets the value for ValidationInformation to be an explicit nil

### UnsetValidationInformation
`func (o *WorkflowServiceItemDefinition) UnsetValidationInformation()`

UnsetValidationInformation ensures that no value is present for ValidationInformation, not even an explicit nil
### GetVersion

`func (o *WorkflowServiceItemDefinition) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkflowServiceItemDefinition) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkflowServiceItemDefinition) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WorkflowServiceItemDefinition) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetActionDefinitions

`func (o *WorkflowServiceItemDefinition) GetActionDefinitions() []WorkflowServiceItemActionDefinitionRelationship`

GetActionDefinitions returns the ActionDefinitions field if non-nil, zero value otherwise.

### GetActionDefinitionsOk

`func (o *WorkflowServiceItemDefinition) GetActionDefinitionsOk() (*[]WorkflowServiceItemActionDefinitionRelationship, bool)`

GetActionDefinitionsOk returns a tuple with the ActionDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionDefinitions

`func (o *WorkflowServiceItemDefinition) SetActionDefinitions(v []WorkflowServiceItemActionDefinitionRelationship)`

SetActionDefinitions sets ActionDefinitions field to given value.

### HasActionDefinitions

`func (o *WorkflowServiceItemDefinition) HasActionDefinitions() bool`

HasActionDefinitions returns a boolean if a field has been set.

### SetActionDefinitionsNil

`func (o *WorkflowServiceItemDefinition) SetActionDefinitionsNil(b bool)`

 SetActionDefinitionsNil sets the value for ActionDefinitions to be an explicit nil

### UnsetActionDefinitions
`func (o *WorkflowServiceItemDefinition) UnsetActionDefinitions()`

UnsetActionDefinitions ensures that no value is present for ActionDefinitions, not even an explicit nil
### GetCatalog

`func (o *WorkflowServiceItemDefinition) GetCatalog() WorkflowCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *WorkflowServiceItemDefinition) GetCatalogOk() (*WorkflowCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *WorkflowServiceItemDefinition) SetCatalog(v WorkflowCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *WorkflowServiceItemDefinition) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### SetCatalogNil

`func (o *WorkflowServiceItemDefinition) SetCatalogNil(b bool)`

 SetCatalogNil sets the value for Catalog to be an explicit nil

### UnsetCatalog
`func (o *WorkflowServiceItemDefinition) UnsetCatalog()`

UnsetCatalog ensures that no value is present for Catalog, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


