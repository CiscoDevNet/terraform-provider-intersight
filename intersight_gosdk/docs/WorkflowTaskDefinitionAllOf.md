# WorkflowTaskDefinitionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.TaskDefinition"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.TaskDefinition"]
**DefaultVersion** | Pointer to **bool** | When true this will be the task version that is used when a specific task definition version is not specified. The very first task definition created with a name will be set as the default version, after that user can explicitly set any version of the task definition as the default version. | [optional] 
**Description** | Pointer to **string** | A user friendly description about task on what operations are done as part of the task execution and any other specific information about task input and output. | [optional] 
**InternalProperties** | Pointer to [**NullableWorkflowInternalProperties**](WorkflowInternalProperties.md) |  | [optional] 
**Label** | Pointer to **string** | A user friendly short name to identify the task definition. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), single quote (&#39;), forward slash (/), or an underscore (_) and must be at least 2 characters. | [optional] 
**LicenseEntitlement** | Pointer to **string** | License entitlement required to run this task. It is determined by license requirement of features. * &#x60;Base&#x60; - Base as a License type. It is default license type. * &#x60;Essential&#x60; - Essential as a License type. * &#x60;Standard&#x60; - Standard as a License type. * &#x60;Advantage&#x60; - Advantage as a License type. * &#x60;Premier&#x60; - Premier as a License type. * &#x60;IWO-Essential&#x60; - IWO-Essential as a License type. * &#x60;IWO-Advantage&#x60; - IWO-Advantage as a License type. * &#x60;IWO-Premier&#x60; - IWO-Premier as a License type. * &#x60;IKS-Advantage&#x60; - IKS-Advantage as a License type. | [optional] [readonly] [default to "Base"]
**Name** | Pointer to **string** | The name of the task definition. The name should follow this convention &lt;Verb or Action&gt;&lt;Category&gt;&lt;Vendor&gt;&lt;Product&gt;&lt;Noun or object&gt; Verb or Action is a required portion of the name and this must be part of the pre-approved verb list. Category is an optional field and this will refer to the broad category of the task referring to the type of resource or endpoint. If there is no specific category then use \&quot;Generic\&quot; if required. Vendor is an optional field and this will refer to the specific vendor this task applies to. If the task is generic and not tied to a vendor, then do not specify anything. Product is an optional field, this will contain the vendor product and model when desired. Noun or object is a required field and  this will contain the noun or object on which the action is being performed. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), or an underscore (_). Examples SendEmail  - This is a task in Generic category for sending email. NewStorageVolume - This is a vendor agnostic task under Storage device category for creating a new volume. | [optional] 
**Properties** | Pointer to [**NullableWorkflowProperties**](WorkflowProperties.md) |  | [optional] 
**RollbackTasks** | Pointer to [**[]WorkflowRollbackTask**](WorkflowRollbackTask.md) |  | [optional] 
**SecurePropAccess** | Pointer to **bool** | If set to true, the task requires access to secure properties and uses an encryption token associated with a workflow moid to encrypt or decrypt the secure properties. | [optional] 
**Version** | Pointer to **int64** | The version of the task definition so we can support multiple versions of a task definition. | [optional] [default to 1]
**Catalog** | Pointer to [**WorkflowCatalogRelationship**](WorkflowCatalogRelationship.md) |  | [optional] 
**ClonedFrom** | Pointer to [**WorkflowTaskDefinitionRelationship**](WorkflowTaskDefinitionRelationship.md) |  | [optional] 
**ImplementedTasks** | Pointer to [**[]WorkflowTaskDefinitionRelationship**](WorkflowTaskDefinitionRelationship.md) | An array of relationships to workflowTaskDefinition resources. | [optional] 
**InterfaceTask** | Pointer to [**WorkflowTaskDefinitionRelationship**](WorkflowTaskDefinitionRelationship.md) |  | [optional] 
**TaskMetadata** | Pointer to [**WorkflowTaskMetadataRelationship**](WorkflowTaskMetadataRelationship.md) |  | [optional] 

## Methods

### NewWorkflowTaskDefinitionAllOf

`func NewWorkflowTaskDefinitionAllOf(classId string, objectType string, ) *WorkflowTaskDefinitionAllOf`

NewWorkflowTaskDefinitionAllOf instantiates a new WorkflowTaskDefinitionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskDefinitionAllOfWithDefaults

`func NewWorkflowTaskDefinitionAllOfWithDefaults() *WorkflowTaskDefinitionAllOf`

NewWorkflowTaskDefinitionAllOfWithDefaults instantiates a new WorkflowTaskDefinitionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowTaskDefinitionAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowTaskDefinitionAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowTaskDefinitionAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowTaskDefinitionAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowTaskDefinitionAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowTaskDefinitionAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDefaultVersion

`func (o *WorkflowTaskDefinitionAllOf) GetDefaultVersion() bool`

GetDefaultVersion returns the DefaultVersion field if non-nil, zero value otherwise.

### GetDefaultVersionOk

`func (o *WorkflowTaskDefinitionAllOf) GetDefaultVersionOk() (*bool, bool)`

GetDefaultVersionOk returns a tuple with the DefaultVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVersion

`func (o *WorkflowTaskDefinitionAllOf) SetDefaultVersion(v bool)`

SetDefaultVersion sets DefaultVersion field to given value.

### HasDefaultVersion

`func (o *WorkflowTaskDefinitionAllOf) HasDefaultVersion() bool`

HasDefaultVersion returns a boolean if a field has been set.

### GetDescription

`func (o *WorkflowTaskDefinitionAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowTaskDefinitionAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowTaskDefinitionAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowTaskDefinitionAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInternalProperties

`func (o *WorkflowTaskDefinitionAllOf) GetInternalProperties() WorkflowInternalProperties`

GetInternalProperties returns the InternalProperties field if non-nil, zero value otherwise.

### GetInternalPropertiesOk

`func (o *WorkflowTaskDefinitionAllOf) GetInternalPropertiesOk() (*WorkflowInternalProperties, bool)`

GetInternalPropertiesOk returns a tuple with the InternalProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalProperties

`func (o *WorkflowTaskDefinitionAllOf) SetInternalProperties(v WorkflowInternalProperties)`

SetInternalProperties sets InternalProperties field to given value.

### HasInternalProperties

`func (o *WorkflowTaskDefinitionAllOf) HasInternalProperties() bool`

HasInternalProperties returns a boolean if a field has been set.

### SetInternalPropertiesNil

`func (o *WorkflowTaskDefinitionAllOf) SetInternalPropertiesNil(b bool)`

 SetInternalPropertiesNil sets the value for InternalProperties to be an explicit nil

### UnsetInternalProperties
`func (o *WorkflowTaskDefinitionAllOf) UnsetInternalProperties()`

UnsetInternalProperties ensures that no value is present for InternalProperties, not even an explicit nil
### GetLabel

`func (o *WorkflowTaskDefinitionAllOf) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WorkflowTaskDefinitionAllOf) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WorkflowTaskDefinitionAllOf) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WorkflowTaskDefinitionAllOf) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLicenseEntitlement

`func (o *WorkflowTaskDefinitionAllOf) GetLicenseEntitlement() string`

GetLicenseEntitlement returns the LicenseEntitlement field if non-nil, zero value otherwise.

### GetLicenseEntitlementOk

`func (o *WorkflowTaskDefinitionAllOf) GetLicenseEntitlementOk() (*string, bool)`

GetLicenseEntitlementOk returns a tuple with the LicenseEntitlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseEntitlement

`func (o *WorkflowTaskDefinitionAllOf) SetLicenseEntitlement(v string)`

SetLicenseEntitlement sets LicenseEntitlement field to given value.

### HasLicenseEntitlement

`func (o *WorkflowTaskDefinitionAllOf) HasLicenseEntitlement() bool`

HasLicenseEntitlement returns a boolean if a field has been set.

### GetName

`func (o *WorkflowTaskDefinitionAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowTaskDefinitionAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowTaskDefinitionAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowTaskDefinitionAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProperties

`func (o *WorkflowTaskDefinitionAllOf) GetProperties() WorkflowProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WorkflowTaskDefinitionAllOf) GetPropertiesOk() (*WorkflowProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WorkflowTaskDefinitionAllOf) SetProperties(v WorkflowProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *WorkflowTaskDefinitionAllOf) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *WorkflowTaskDefinitionAllOf) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *WorkflowTaskDefinitionAllOf) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetRollbackTasks

`func (o *WorkflowTaskDefinitionAllOf) GetRollbackTasks() []WorkflowRollbackTask`

GetRollbackTasks returns the RollbackTasks field if non-nil, zero value otherwise.

### GetRollbackTasksOk

`func (o *WorkflowTaskDefinitionAllOf) GetRollbackTasksOk() (*[]WorkflowRollbackTask, bool)`

GetRollbackTasksOk returns a tuple with the RollbackTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollbackTasks

`func (o *WorkflowTaskDefinitionAllOf) SetRollbackTasks(v []WorkflowRollbackTask)`

SetRollbackTasks sets RollbackTasks field to given value.

### HasRollbackTasks

`func (o *WorkflowTaskDefinitionAllOf) HasRollbackTasks() bool`

HasRollbackTasks returns a boolean if a field has been set.

### SetRollbackTasksNil

`func (o *WorkflowTaskDefinitionAllOf) SetRollbackTasksNil(b bool)`

 SetRollbackTasksNil sets the value for RollbackTasks to be an explicit nil

### UnsetRollbackTasks
`func (o *WorkflowTaskDefinitionAllOf) UnsetRollbackTasks()`

UnsetRollbackTasks ensures that no value is present for RollbackTasks, not even an explicit nil
### GetSecurePropAccess

`func (o *WorkflowTaskDefinitionAllOf) GetSecurePropAccess() bool`

GetSecurePropAccess returns the SecurePropAccess field if non-nil, zero value otherwise.

### GetSecurePropAccessOk

`func (o *WorkflowTaskDefinitionAllOf) GetSecurePropAccessOk() (*bool, bool)`

GetSecurePropAccessOk returns a tuple with the SecurePropAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurePropAccess

`func (o *WorkflowTaskDefinitionAllOf) SetSecurePropAccess(v bool)`

SetSecurePropAccess sets SecurePropAccess field to given value.

### HasSecurePropAccess

`func (o *WorkflowTaskDefinitionAllOf) HasSecurePropAccess() bool`

HasSecurePropAccess returns a boolean if a field has been set.

### GetVersion

`func (o *WorkflowTaskDefinitionAllOf) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkflowTaskDefinitionAllOf) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkflowTaskDefinitionAllOf) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WorkflowTaskDefinitionAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCatalog

`func (o *WorkflowTaskDefinitionAllOf) GetCatalog() WorkflowCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *WorkflowTaskDefinitionAllOf) GetCatalogOk() (*WorkflowCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *WorkflowTaskDefinitionAllOf) SetCatalog(v WorkflowCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *WorkflowTaskDefinitionAllOf) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### GetClonedFrom

`func (o *WorkflowTaskDefinitionAllOf) GetClonedFrom() WorkflowTaskDefinitionRelationship`

GetClonedFrom returns the ClonedFrom field if non-nil, zero value otherwise.

### GetClonedFromOk

`func (o *WorkflowTaskDefinitionAllOf) GetClonedFromOk() (*WorkflowTaskDefinitionRelationship, bool)`

GetClonedFromOk returns a tuple with the ClonedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClonedFrom

`func (o *WorkflowTaskDefinitionAllOf) SetClonedFrom(v WorkflowTaskDefinitionRelationship)`

SetClonedFrom sets ClonedFrom field to given value.

### HasClonedFrom

`func (o *WorkflowTaskDefinitionAllOf) HasClonedFrom() bool`

HasClonedFrom returns a boolean if a field has been set.

### GetImplementedTasks

`func (o *WorkflowTaskDefinitionAllOf) GetImplementedTasks() []WorkflowTaskDefinitionRelationship`

GetImplementedTasks returns the ImplementedTasks field if non-nil, zero value otherwise.

### GetImplementedTasksOk

`func (o *WorkflowTaskDefinitionAllOf) GetImplementedTasksOk() (*[]WorkflowTaskDefinitionRelationship, bool)`

GetImplementedTasksOk returns a tuple with the ImplementedTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementedTasks

`func (o *WorkflowTaskDefinitionAllOf) SetImplementedTasks(v []WorkflowTaskDefinitionRelationship)`

SetImplementedTasks sets ImplementedTasks field to given value.

### HasImplementedTasks

`func (o *WorkflowTaskDefinitionAllOf) HasImplementedTasks() bool`

HasImplementedTasks returns a boolean if a field has been set.

### SetImplementedTasksNil

`func (o *WorkflowTaskDefinitionAllOf) SetImplementedTasksNil(b bool)`

 SetImplementedTasksNil sets the value for ImplementedTasks to be an explicit nil

### UnsetImplementedTasks
`func (o *WorkflowTaskDefinitionAllOf) UnsetImplementedTasks()`

UnsetImplementedTasks ensures that no value is present for ImplementedTasks, not even an explicit nil
### GetInterfaceTask

`func (o *WorkflowTaskDefinitionAllOf) GetInterfaceTask() WorkflowTaskDefinitionRelationship`

GetInterfaceTask returns the InterfaceTask field if non-nil, zero value otherwise.

### GetInterfaceTaskOk

`func (o *WorkflowTaskDefinitionAllOf) GetInterfaceTaskOk() (*WorkflowTaskDefinitionRelationship, bool)`

GetInterfaceTaskOk returns a tuple with the InterfaceTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceTask

`func (o *WorkflowTaskDefinitionAllOf) SetInterfaceTask(v WorkflowTaskDefinitionRelationship)`

SetInterfaceTask sets InterfaceTask field to given value.

### HasInterfaceTask

`func (o *WorkflowTaskDefinitionAllOf) HasInterfaceTask() bool`

HasInterfaceTask returns a boolean if a field has been set.

### GetTaskMetadata

`func (o *WorkflowTaskDefinitionAllOf) GetTaskMetadata() WorkflowTaskMetadataRelationship`

GetTaskMetadata returns the TaskMetadata field if non-nil, zero value otherwise.

### GetTaskMetadataOk

`func (o *WorkflowTaskDefinitionAllOf) GetTaskMetadataOk() (*WorkflowTaskMetadataRelationship, bool)`

GetTaskMetadataOk returns a tuple with the TaskMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskMetadata

`func (o *WorkflowTaskDefinitionAllOf) SetTaskMetadata(v WorkflowTaskMetadataRelationship)`

SetTaskMetadata sets TaskMetadata field to given value.

### HasTaskMetadata

`func (o *WorkflowTaskDefinitionAllOf) HasTaskMetadata() bool`

HasTaskMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


