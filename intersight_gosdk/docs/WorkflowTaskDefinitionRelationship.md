# WorkflowTaskDefinitionRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountMoid** | Pointer to **string** | The Account ID for this managed object. | [optional] [readonly] 
**ClassId** | **string** | The concrete type of this complex type. Its value must be the same as the &#39;objectType&#39; property. The OpenAPI document references this property as a discriminator value. | [readonly] 
**CreateTime** | Pointer to [**time.Time**](time.Time.md) | The time when this managed object was created. | [optional] [readonly] 
**DomainGroupMoid** | Pointer to **string** | The DomainGroup ID for this managed object. | [optional] [readonly] 
**ModTime** | Pointer to [**time.Time**](time.Time.md) | The time when this managed object was last modified. | [optional] [readonly] 
**Moid** | Pointer to **string** | The unique identifier of this Managed Object instance. | [optional] 
**ObjectType** | **string** | The fully-qualified type of this managed object, i.e. the class name. This property is optional. The ObjectType is implied from the URL path. If specified, the value of objectType must match the class name specified in the URL path. | [readonly] 
**Owners** | Pointer to **[]string** |  | [optional] 
**SharedScope** | Pointer to **string** | Intersight provides pre-built workflows, tasks and policies to end users through global catalogs. Objects that are made available through global catalogs are said to have a &#39;shared&#39; ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. | [optional] [readonly] 
**Tags** | Pointer to [**[]MoTag**](mo.Tag.md) |  | [optional] 
**VersionContext** | Pointer to [**MoVersionContext**](mo.VersionContext.md) |  | [optional] 
**Ancestors** | Pointer to [**[]MoBaseMoRelationship**](mo.BaseMo.Relationship.md) | An array of relationships to moBaseMo resources. | [optional] [readonly] 
**Parent** | Pointer to [**MoBaseMoRelationship**](mo.BaseMo.Relationship.md) |  | [optional] 
**PermissionResources** | Pointer to [**[]MoBaseMoRelationship**](mo.BaseMo.Relationship.md) | An array of relationships to moBaseMo resources. | [optional] [readonly] 
**DisplayNames** | Pointer to [**map[string][]string**](array.md) | A set of display names for the MO resource. These names are calculated based on other properties of the MO and potentially properties of Ancestor MOs. Displaynames are intended as a way to provide a normalized user appropriate name for an MO, especially for MOs which do not have a &#39;Name&#39; property, which is the case for much of the inventory discovered from managed targets. There are a limited number of keys, currently &#39;short&#39; and &#39;hierarchical&#39;. The value is an array and clients should use the first element of the array. | [optional] [readonly] 
**DefaultVersion** | Pointer to **bool** | When true this will be the task version that is used when a specific task definition version is not specified. The very first task definition created with a name will be set as the default version, after that user can explicitly set any version of the task definition as the default version. | [optional] 
**Description** | Pointer to **string** | The task definition description to describe what this task will do when executed. | [optional] 
**InternalProperties** | Pointer to [**WorkflowInternalProperties**](workflow.InternalProperties.md) |  | [optional] 
**Label** | Pointer to **string** | A user friendly short name to identify the task definition. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), single quote (&#39;), or an underscore (_). | [optional] 
**LicenseEntitlement** | Pointer to **string** | License entitlement required to run this task. It is determined by license requirement of features. * &#x60;Base&#x60; - Base as a License type. It is default license type. * &#x60;Essential&#x60; - Essential as a License type. * &#x60;Standard&#x60; - Standard as a License type. * &#x60;Advantage&#x60; - Advantage as a License type. * &#x60;Premier&#x60; - Premier as a License type. | [optional] [readonly] [default to "Base"]
**Name** | Pointer to **string** | The name of the task definition. The name should follow this convention &lt;Verb or Action&gt;&lt;Category&gt;&lt;Vendor&gt;&lt;Product&gt;&lt;Noun or object&gt; Verb or Action is a required portion of the name and this must be part of the pre-approved verb list. Category is an optional field and this will refer to the broad category of the task referring to the type of resource or endpoint. If there is no specific category then use \&quot;Generic\&quot; if required. Vendor is an optional field and this will refer to the specific vendor this task applies to. If the task is generic and not tied to a vendor, then do not specify anything. Product is an optional field, this will contain the vendor product and model when desired. Noun or object is a required field and  this will contain the noun or object on which the action is being performed. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), or an underscore (_). Examples SendEmail  - This is a task in Generic category for sending email. NewStorageVolume - This is a vendor agnostic task under Storage device category for creating a new volume. | [optional] 
**Properties** | Pointer to [**WorkflowProperties**](workflow.Properties.md) |  | [optional] 
**SecurePropAccess** | Pointer to **bool** | If set to true, the task requires access to secure properties and uses an encyption token associated with a workflow moid to encrypt or decrypt the secure properties. | [optional] 
**Version** | Pointer to **int64** | The version of the task definition so we can support multiple versions of a task definition. | [optional] 
**Catalog** | Pointer to [**WorkflowCatalogRelationship**](workflow.Catalog.Relationship.md) |  | [optional] 
**ImplementedTasks** | Pointer to [**[]WorkflowTaskDefinitionRelationship**](workflow.TaskDefinition.Relationship.md) | An array of relationships to workflowTaskDefinition resources. | [optional] 
**InterfaceTask** | Pointer to [**WorkflowTaskDefinitionRelationship**](workflow.TaskDefinition.Relationship.md) |  | [optional] 

## Methods

### NewWorkflowTaskDefinitionRelationship

`func NewWorkflowTaskDefinitionRelationship(classId string, objectType string, ) *WorkflowTaskDefinitionRelationship`

NewWorkflowTaskDefinitionRelationship instantiates a new WorkflowTaskDefinitionRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskDefinitionRelationshipWithDefaults

`func NewWorkflowTaskDefinitionRelationshipWithDefaults() *WorkflowTaskDefinitionRelationship`

NewWorkflowTaskDefinitionRelationshipWithDefaults instantiates a new WorkflowTaskDefinitionRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *WorkflowTaskDefinitionRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *WorkflowTaskDefinitionRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *WorkflowTaskDefinitionRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *WorkflowTaskDefinitionRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *WorkflowTaskDefinitionRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowTaskDefinitionRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowTaskDefinitionRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *WorkflowTaskDefinitionRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *WorkflowTaskDefinitionRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *WorkflowTaskDefinitionRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *WorkflowTaskDefinitionRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *WorkflowTaskDefinitionRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *WorkflowTaskDefinitionRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *WorkflowTaskDefinitionRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *WorkflowTaskDefinitionRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *WorkflowTaskDefinitionRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *WorkflowTaskDefinitionRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *WorkflowTaskDefinitionRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *WorkflowTaskDefinitionRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *WorkflowTaskDefinitionRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *WorkflowTaskDefinitionRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *WorkflowTaskDefinitionRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *WorkflowTaskDefinitionRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *WorkflowTaskDefinitionRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowTaskDefinitionRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowTaskDefinitionRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *WorkflowTaskDefinitionRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *WorkflowTaskDefinitionRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *WorkflowTaskDefinitionRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *WorkflowTaskDefinitionRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *WorkflowTaskDefinitionRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *WorkflowTaskDefinitionRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *WorkflowTaskDefinitionRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *WorkflowTaskDefinitionRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *WorkflowTaskDefinitionRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WorkflowTaskDefinitionRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WorkflowTaskDefinitionRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WorkflowTaskDefinitionRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *WorkflowTaskDefinitionRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *WorkflowTaskDefinitionRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *WorkflowTaskDefinitionRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *WorkflowTaskDefinitionRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *WorkflowTaskDefinitionRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *WorkflowTaskDefinitionRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *WorkflowTaskDefinitionRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *WorkflowTaskDefinitionRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *WorkflowTaskDefinitionRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *WorkflowTaskDefinitionRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *WorkflowTaskDefinitionRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *WorkflowTaskDefinitionRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *WorkflowTaskDefinitionRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *WorkflowTaskDefinitionRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *WorkflowTaskDefinitionRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *WorkflowTaskDefinitionRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *WorkflowTaskDefinitionRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *WorkflowTaskDefinitionRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *WorkflowTaskDefinitionRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *WorkflowTaskDefinitionRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *WorkflowTaskDefinitionRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *WorkflowTaskDefinitionRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *WorkflowTaskDefinitionRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *WorkflowTaskDefinitionRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *WorkflowTaskDefinitionRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *WorkflowTaskDefinitionRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDefaultVersion

`func (o *WorkflowTaskDefinitionRelationship) GetDefaultVersion() bool`

GetDefaultVersion returns the DefaultVersion field if non-nil, zero value otherwise.

### GetDefaultVersionOk

`func (o *WorkflowTaskDefinitionRelationship) GetDefaultVersionOk() (*bool, bool)`

GetDefaultVersionOk returns a tuple with the DefaultVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVersion

`func (o *WorkflowTaskDefinitionRelationship) SetDefaultVersion(v bool)`

SetDefaultVersion sets DefaultVersion field to given value.

### HasDefaultVersion

`func (o *WorkflowTaskDefinitionRelationship) HasDefaultVersion() bool`

HasDefaultVersion returns a boolean if a field has been set.

### GetDescription

`func (o *WorkflowTaskDefinitionRelationship) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowTaskDefinitionRelationship) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowTaskDefinitionRelationship) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowTaskDefinitionRelationship) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInternalProperties

`func (o *WorkflowTaskDefinitionRelationship) GetInternalProperties() WorkflowInternalProperties`

GetInternalProperties returns the InternalProperties field if non-nil, zero value otherwise.

### GetInternalPropertiesOk

`func (o *WorkflowTaskDefinitionRelationship) GetInternalPropertiesOk() (*WorkflowInternalProperties, bool)`

GetInternalPropertiesOk returns a tuple with the InternalProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalProperties

`func (o *WorkflowTaskDefinitionRelationship) SetInternalProperties(v WorkflowInternalProperties)`

SetInternalProperties sets InternalProperties field to given value.

### HasInternalProperties

`func (o *WorkflowTaskDefinitionRelationship) HasInternalProperties() bool`

HasInternalProperties returns a boolean if a field has been set.

### GetLabel

`func (o *WorkflowTaskDefinitionRelationship) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WorkflowTaskDefinitionRelationship) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WorkflowTaskDefinitionRelationship) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WorkflowTaskDefinitionRelationship) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLicenseEntitlement

`func (o *WorkflowTaskDefinitionRelationship) GetLicenseEntitlement() string`

GetLicenseEntitlement returns the LicenseEntitlement field if non-nil, zero value otherwise.

### GetLicenseEntitlementOk

`func (o *WorkflowTaskDefinitionRelationship) GetLicenseEntitlementOk() (*string, bool)`

GetLicenseEntitlementOk returns a tuple with the LicenseEntitlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseEntitlement

`func (o *WorkflowTaskDefinitionRelationship) SetLicenseEntitlement(v string)`

SetLicenseEntitlement sets LicenseEntitlement field to given value.

### HasLicenseEntitlement

`func (o *WorkflowTaskDefinitionRelationship) HasLicenseEntitlement() bool`

HasLicenseEntitlement returns a boolean if a field has been set.

### GetName

`func (o *WorkflowTaskDefinitionRelationship) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowTaskDefinitionRelationship) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowTaskDefinitionRelationship) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowTaskDefinitionRelationship) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProperties

`func (o *WorkflowTaskDefinitionRelationship) GetProperties() WorkflowProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WorkflowTaskDefinitionRelationship) GetPropertiesOk() (*WorkflowProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WorkflowTaskDefinitionRelationship) SetProperties(v WorkflowProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *WorkflowTaskDefinitionRelationship) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetSecurePropAccess

`func (o *WorkflowTaskDefinitionRelationship) GetSecurePropAccess() bool`

GetSecurePropAccess returns the SecurePropAccess field if non-nil, zero value otherwise.

### GetSecurePropAccessOk

`func (o *WorkflowTaskDefinitionRelationship) GetSecurePropAccessOk() (*bool, bool)`

GetSecurePropAccessOk returns a tuple with the SecurePropAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurePropAccess

`func (o *WorkflowTaskDefinitionRelationship) SetSecurePropAccess(v bool)`

SetSecurePropAccess sets SecurePropAccess field to given value.

### HasSecurePropAccess

`func (o *WorkflowTaskDefinitionRelationship) HasSecurePropAccess() bool`

HasSecurePropAccess returns a boolean if a field has been set.

### GetVersion

`func (o *WorkflowTaskDefinitionRelationship) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkflowTaskDefinitionRelationship) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkflowTaskDefinitionRelationship) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WorkflowTaskDefinitionRelationship) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCatalog

`func (o *WorkflowTaskDefinitionRelationship) GetCatalog() WorkflowCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *WorkflowTaskDefinitionRelationship) GetCatalogOk() (*WorkflowCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *WorkflowTaskDefinitionRelationship) SetCatalog(v WorkflowCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *WorkflowTaskDefinitionRelationship) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### GetImplementedTasks

`func (o *WorkflowTaskDefinitionRelationship) GetImplementedTasks() []WorkflowTaskDefinitionRelationship`

GetImplementedTasks returns the ImplementedTasks field if non-nil, zero value otherwise.

### GetImplementedTasksOk

`func (o *WorkflowTaskDefinitionRelationship) GetImplementedTasksOk() (*[]WorkflowTaskDefinitionRelationship, bool)`

GetImplementedTasksOk returns a tuple with the ImplementedTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementedTasks

`func (o *WorkflowTaskDefinitionRelationship) SetImplementedTasks(v []WorkflowTaskDefinitionRelationship)`

SetImplementedTasks sets ImplementedTasks field to given value.

### HasImplementedTasks

`func (o *WorkflowTaskDefinitionRelationship) HasImplementedTasks() bool`

HasImplementedTasks returns a boolean if a field has been set.

### SetImplementedTasksNil

`func (o *WorkflowTaskDefinitionRelationship) SetImplementedTasksNil(b bool)`

 SetImplementedTasksNil sets the value for ImplementedTasks to be an explicit nil

### UnsetImplementedTasks
`func (o *WorkflowTaskDefinitionRelationship) UnsetImplementedTasks()`

UnsetImplementedTasks ensures that no value is present for ImplementedTasks, not even an explicit nil
### GetInterfaceTask

`func (o *WorkflowTaskDefinitionRelationship) GetInterfaceTask() WorkflowTaskDefinitionRelationship`

GetInterfaceTask returns the InterfaceTask field if non-nil, zero value otherwise.

### GetInterfaceTaskOk

`func (o *WorkflowTaskDefinitionRelationship) GetInterfaceTaskOk() (*WorkflowTaskDefinitionRelationship, bool)`

GetInterfaceTaskOk returns a tuple with the InterfaceTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceTask

`func (o *WorkflowTaskDefinitionRelationship) SetInterfaceTask(v WorkflowTaskDefinitionRelationship)`

SetInterfaceTask sets InterfaceTask field to given value.

### HasInterfaceTask

`func (o *WorkflowTaskDefinitionRelationship) HasInterfaceTask() bool`

HasInterfaceTask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


