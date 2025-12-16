# WorkloadBlueprint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workload.Blueprint"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workload.Blueprint"]
**BlueprintDependencies** | Pointer to [**[]BlueprintBlueprintDefinitionDependency**](BlueprintBlueprintDefinitionDependency.md) |  | [optional] 
**DefaultVersion** | Pointer to **bool** | The flag to indicate that this is the default version of the blueprint. | [optional] [default to false]
**Description** | Pointer to **string** | The description for this blueprint which provides information on what are the pre-requisites to deploy the blueprint and what features are supported on the blueprint. | [optional] 
**ExternalMeta** | Pointer to **bool** | When set to false the blueprint is created for use within internal services. | [optional] [default to false]
**GeneratedObjectDefinition** | Pointer to [**[]BlueprintGeneratedObjectMetadata**](BlueprintGeneratedObjectMetadata.md) |  | [optional] 
**InputDefinition** | Pointer to [**[]WorkflowBaseDataType**](WorkflowBaseDataType.md) |  | [optional] 
**InputOperationMetadata** | Pointer to [**[]BlueprintInputOperationType**](BlueprintInputOperationType.md) |  | [optional] 
**Label** | Pointer to **string** | A user friendly short name to identify the blueprint. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_). | [optional] 
**Name** | Pointer to **string** | The name for this blueprint. You can have multiple versions of the blueprint with the same name. Name can only contain letters (a-z, A-Z), numbers (0-9), or an underscore (_). | [optional] 
**PlatformType** | Pointer to **string** | The Intersight platforms supported by this blueprint. * &#x60;None&#x60; - Default value is none, platform is not mentioned. * &#x60;UnifiedEdge&#x60; - The blueprint is created for Unified Edge platform. | [optional] [default to "None"]
**ResourceDefinition** | Pointer to [**NullableBlueprintResourceConstraints**](BlueprintResourceConstraints.md) |  | [optional] 
**ServiceItems** | Pointer to [**[]BlueprintServiceItemDefinition**](BlueprintServiceItemDefinition.md) |  | [optional] 
**ValidationInformation** | Pointer to [**NullableWorkflowValidationInformation**](WorkflowValidationInformation.md) |  | [optional] 
**Version** | Pointer to **int64** | The version of the blueprint to support multiple versions. | [optional] [default to 1]
**AssociatedRoles** | Pointer to [**[]IamRoleRelationship**](IamRoleRelationship.md) | An array of relationships to iamRole resources. | [optional] 
**Catalog** | Pointer to [**NullableWorkflowCatalogRelationship**](WorkflowCatalogRelationship.md) |  | [optional] 

## Methods

### NewWorkloadBlueprint

`func NewWorkloadBlueprint(classId string, objectType string, ) *WorkloadBlueprint`

NewWorkloadBlueprint instantiates a new WorkloadBlueprint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadBlueprintWithDefaults

`func NewWorkloadBlueprintWithDefaults() *WorkloadBlueprint`

NewWorkloadBlueprintWithDefaults instantiates a new WorkloadBlueprint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkloadBlueprint) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkloadBlueprint) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkloadBlueprint) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkloadBlueprint) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkloadBlueprint) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkloadBlueprint) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBlueprintDependencies

`func (o *WorkloadBlueprint) GetBlueprintDependencies() []BlueprintBlueprintDefinitionDependency`

GetBlueprintDependencies returns the BlueprintDependencies field if non-nil, zero value otherwise.

### GetBlueprintDependenciesOk

`func (o *WorkloadBlueprint) GetBlueprintDependenciesOk() (*[]BlueprintBlueprintDefinitionDependency, bool)`

GetBlueprintDependenciesOk returns a tuple with the BlueprintDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprintDependencies

`func (o *WorkloadBlueprint) SetBlueprintDependencies(v []BlueprintBlueprintDefinitionDependency)`

SetBlueprintDependencies sets BlueprintDependencies field to given value.

### HasBlueprintDependencies

`func (o *WorkloadBlueprint) HasBlueprintDependencies() bool`

HasBlueprintDependencies returns a boolean if a field has been set.

### SetBlueprintDependenciesNil

`func (o *WorkloadBlueprint) SetBlueprintDependenciesNil(b bool)`

 SetBlueprintDependenciesNil sets the value for BlueprintDependencies to be an explicit nil

### UnsetBlueprintDependencies
`func (o *WorkloadBlueprint) UnsetBlueprintDependencies()`

UnsetBlueprintDependencies ensures that no value is present for BlueprintDependencies, not even an explicit nil
### GetDefaultVersion

`func (o *WorkloadBlueprint) GetDefaultVersion() bool`

GetDefaultVersion returns the DefaultVersion field if non-nil, zero value otherwise.

### GetDefaultVersionOk

`func (o *WorkloadBlueprint) GetDefaultVersionOk() (*bool, bool)`

GetDefaultVersionOk returns a tuple with the DefaultVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVersion

`func (o *WorkloadBlueprint) SetDefaultVersion(v bool)`

SetDefaultVersion sets DefaultVersion field to given value.

### HasDefaultVersion

`func (o *WorkloadBlueprint) HasDefaultVersion() bool`

HasDefaultVersion returns a boolean if a field has been set.

### GetDescription

`func (o *WorkloadBlueprint) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkloadBlueprint) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkloadBlueprint) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkloadBlueprint) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalMeta

`func (o *WorkloadBlueprint) GetExternalMeta() bool`

GetExternalMeta returns the ExternalMeta field if non-nil, zero value otherwise.

### GetExternalMetaOk

`func (o *WorkloadBlueprint) GetExternalMetaOk() (*bool, bool)`

GetExternalMetaOk returns a tuple with the ExternalMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalMeta

`func (o *WorkloadBlueprint) SetExternalMeta(v bool)`

SetExternalMeta sets ExternalMeta field to given value.

### HasExternalMeta

`func (o *WorkloadBlueprint) HasExternalMeta() bool`

HasExternalMeta returns a boolean if a field has been set.

### GetGeneratedObjectDefinition

`func (o *WorkloadBlueprint) GetGeneratedObjectDefinition() []BlueprintGeneratedObjectMetadata`

GetGeneratedObjectDefinition returns the GeneratedObjectDefinition field if non-nil, zero value otherwise.

### GetGeneratedObjectDefinitionOk

`func (o *WorkloadBlueprint) GetGeneratedObjectDefinitionOk() (*[]BlueprintGeneratedObjectMetadata, bool)`

GetGeneratedObjectDefinitionOk returns a tuple with the GeneratedObjectDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedObjectDefinition

`func (o *WorkloadBlueprint) SetGeneratedObjectDefinition(v []BlueprintGeneratedObjectMetadata)`

SetGeneratedObjectDefinition sets GeneratedObjectDefinition field to given value.

### HasGeneratedObjectDefinition

`func (o *WorkloadBlueprint) HasGeneratedObjectDefinition() bool`

HasGeneratedObjectDefinition returns a boolean if a field has been set.

### SetGeneratedObjectDefinitionNil

`func (o *WorkloadBlueprint) SetGeneratedObjectDefinitionNil(b bool)`

 SetGeneratedObjectDefinitionNil sets the value for GeneratedObjectDefinition to be an explicit nil

### UnsetGeneratedObjectDefinition
`func (o *WorkloadBlueprint) UnsetGeneratedObjectDefinition()`

UnsetGeneratedObjectDefinition ensures that no value is present for GeneratedObjectDefinition, not even an explicit nil
### GetInputDefinition

`func (o *WorkloadBlueprint) GetInputDefinition() []WorkflowBaseDataType`

GetInputDefinition returns the InputDefinition field if non-nil, zero value otherwise.

### GetInputDefinitionOk

`func (o *WorkloadBlueprint) GetInputDefinitionOk() (*[]WorkflowBaseDataType, bool)`

GetInputDefinitionOk returns a tuple with the InputDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputDefinition

`func (o *WorkloadBlueprint) SetInputDefinition(v []WorkflowBaseDataType)`

SetInputDefinition sets InputDefinition field to given value.

### HasInputDefinition

`func (o *WorkloadBlueprint) HasInputDefinition() bool`

HasInputDefinition returns a boolean if a field has been set.

### SetInputDefinitionNil

`func (o *WorkloadBlueprint) SetInputDefinitionNil(b bool)`

 SetInputDefinitionNil sets the value for InputDefinition to be an explicit nil

### UnsetInputDefinition
`func (o *WorkloadBlueprint) UnsetInputDefinition()`

UnsetInputDefinition ensures that no value is present for InputDefinition, not even an explicit nil
### GetInputOperationMetadata

`func (o *WorkloadBlueprint) GetInputOperationMetadata() []BlueprintInputOperationType`

GetInputOperationMetadata returns the InputOperationMetadata field if non-nil, zero value otherwise.

### GetInputOperationMetadataOk

`func (o *WorkloadBlueprint) GetInputOperationMetadataOk() (*[]BlueprintInputOperationType, bool)`

GetInputOperationMetadataOk returns a tuple with the InputOperationMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputOperationMetadata

`func (o *WorkloadBlueprint) SetInputOperationMetadata(v []BlueprintInputOperationType)`

SetInputOperationMetadata sets InputOperationMetadata field to given value.

### HasInputOperationMetadata

`func (o *WorkloadBlueprint) HasInputOperationMetadata() bool`

HasInputOperationMetadata returns a boolean if a field has been set.

### SetInputOperationMetadataNil

`func (o *WorkloadBlueprint) SetInputOperationMetadataNil(b bool)`

 SetInputOperationMetadataNil sets the value for InputOperationMetadata to be an explicit nil

### UnsetInputOperationMetadata
`func (o *WorkloadBlueprint) UnsetInputOperationMetadata()`

UnsetInputOperationMetadata ensures that no value is present for InputOperationMetadata, not even an explicit nil
### GetLabel

`func (o *WorkloadBlueprint) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WorkloadBlueprint) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WorkloadBlueprint) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WorkloadBlueprint) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *WorkloadBlueprint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkloadBlueprint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkloadBlueprint) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkloadBlueprint) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlatformType

`func (o *WorkloadBlueprint) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *WorkloadBlueprint) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *WorkloadBlueprint) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *WorkloadBlueprint) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### GetResourceDefinition

`func (o *WorkloadBlueprint) GetResourceDefinition() BlueprintResourceConstraints`

GetResourceDefinition returns the ResourceDefinition field if non-nil, zero value otherwise.

### GetResourceDefinitionOk

`func (o *WorkloadBlueprint) GetResourceDefinitionOk() (*BlueprintResourceConstraints, bool)`

GetResourceDefinitionOk returns a tuple with the ResourceDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDefinition

`func (o *WorkloadBlueprint) SetResourceDefinition(v BlueprintResourceConstraints)`

SetResourceDefinition sets ResourceDefinition field to given value.

### HasResourceDefinition

`func (o *WorkloadBlueprint) HasResourceDefinition() bool`

HasResourceDefinition returns a boolean if a field has been set.

### SetResourceDefinitionNil

`func (o *WorkloadBlueprint) SetResourceDefinitionNil(b bool)`

 SetResourceDefinitionNil sets the value for ResourceDefinition to be an explicit nil

### UnsetResourceDefinition
`func (o *WorkloadBlueprint) UnsetResourceDefinition()`

UnsetResourceDefinition ensures that no value is present for ResourceDefinition, not even an explicit nil
### GetServiceItems

`func (o *WorkloadBlueprint) GetServiceItems() []BlueprintServiceItemDefinition`

GetServiceItems returns the ServiceItems field if non-nil, zero value otherwise.

### GetServiceItemsOk

`func (o *WorkloadBlueprint) GetServiceItemsOk() (*[]BlueprintServiceItemDefinition, bool)`

GetServiceItemsOk returns a tuple with the ServiceItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItems

`func (o *WorkloadBlueprint) SetServiceItems(v []BlueprintServiceItemDefinition)`

SetServiceItems sets ServiceItems field to given value.

### HasServiceItems

`func (o *WorkloadBlueprint) HasServiceItems() bool`

HasServiceItems returns a boolean if a field has been set.

### SetServiceItemsNil

`func (o *WorkloadBlueprint) SetServiceItemsNil(b bool)`

 SetServiceItemsNil sets the value for ServiceItems to be an explicit nil

### UnsetServiceItems
`func (o *WorkloadBlueprint) UnsetServiceItems()`

UnsetServiceItems ensures that no value is present for ServiceItems, not even an explicit nil
### GetValidationInformation

`func (o *WorkloadBlueprint) GetValidationInformation() WorkflowValidationInformation`

GetValidationInformation returns the ValidationInformation field if non-nil, zero value otherwise.

### GetValidationInformationOk

`func (o *WorkloadBlueprint) GetValidationInformationOk() (*WorkflowValidationInformation, bool)`

GetValidationInformationOk returns a tuple with the ValidationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationInformation

`func (o *WorkloadBlueprint) SetValidationInformation(v WorkflowValidationInformation)`

SetValidationInformation sets ValidationInformation field to given value.

### HasValidationInformation

`func (o *WorkloadBlueprint) HasValidationInformation() bool`

HasValidationInformation returns a boolean if a field has been set.

### SetValidationInformationNil

`func (o *WorkloadBlueprint) SetValidationInformationNil(b bool)`

 SetValidationInformationNil sets the value for ValidationInformation to be an explicit nil

### UnsetValidationInformation
`func (o *WorkloadBlueprint) UnsetValidationInformation()`

UnsetValidationInformation ensures that no value is present for ValidationInformation, not even an explicit nil
### GetVersion

`func (o *WorkloadBlueprint) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkloadBlueprint) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkloadBlueprint) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WorkloadBlueprint) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAssociatedRoles

`func (o *WorkloadBlueprint) GetAssociatedRoles() []IamRoleRelationship`

GetAssociatedRoles returns the AssociatedRoles field if non-nil, zero value otherwise.

### GetAssociatedRolesOk

`func (o *WorkloadBlueprint) GetAssociatedRolesOk() (*[]IamRoleRelationship, bool)`

GetAssociatedRolesOk returns a tuple with the AssociatedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedRoles

`func (o *WorkloadBlueprint) SetAssociatedRoles(v []IamRoleRelationship)`

SetAssociatedRoles sets AssociatedRoles field to given value.

### HasAssociatedRoles

`func (o *WorkloadBlueprint) HasAssociatedRoles() bool`

HasAssociatedRoles returns a boolean if a field has been set.

### SetAssociatedRolesNil

`func (o *WorkloadBlueprint) SetAssociatedRolesNil(b bool)`

 SetAssociatedRolesNil sets the value for AssociatedRoles to be an explicit nil

### UnsetAssociatedRoles
`func (o *WorkloadBlueprint) UnsetAssociatedRoles()`

UnsetAssociatedRoles ensures that no value is present for AssociatedRoles, not even an explicit nil
### GetCatalog

`func (o *WorkloadBlueprint) GetCatalog() WorkflowCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *WorkloadBlueprint) GetCatalogOk() (*WorkflowCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *WorkloadBlueprint) SetCatalog(v WorkflowCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *WorkloadBlueprint) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### SetCatalogNil

`func (o *WorkloadBlueprint) SetCatalogNil(b bool)`

 SetCatalogNil sets the value for Catalog to be an explicit nil

### UnsetCatalog
`func (o *WorkloadBlueprint) UnsetCatalog()`

UnsetCatalog ensures that no value is present for Catalog, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


