# WorkflowTaskMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.TaskMetadata"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.TaskMetadata"]
**Description** | Pointer to **string** | The task metadata description to describe what this task will do when executed. | [optional] 
**Label** | Pointer to **string** | A user friendly short name to identify the task metadata. | [optional] 
**Name** | Pointer to **string** | The name of the task metadata. The name should follow this convention &lt;Verb or Action&gt;&lt;Category&gt;&lt;Vendor&gt;&lt;Product&gt;&lt;Noun or object&gt; Verb or Action is a required portion of the name and this must be part of the pre-approved verb list. Category is an optional field and this will refer to the broad category of the task referring to the type of resource or endpoint. If there is no specific category then use \&quot;Generic\&quot; if required. Vendor is an optional field and this will refer to the specific vendor this task applies to. If the task is generic and not tied to a vendor, then do not specify anything. Product is an optional field, this will contain the vendor product and model when desired. Noun or object is a required field and  this will contain the noun or object on which the action is being performed. Examples SendEmail  - This is a task in Generic category for sending email. NewStorageVolume - This is a vendor agnostic task under Storage device category for creating a new volume. | [optional] 
**AssociatedRoles** | Pointer to [**[]IamRoleRelationship**](IamRoleRelationship.md) | An array of relationships to iamRole resources. | [optional] 
**Catalog** | Pointer to [**WorkflowCatalogRelationship**](WorkflowCatalogRelationship.md) |  | [optional] 

## Methods

### NewWorkflowTaskMetadata

`func NewWorkflowTaskMetadata(classId string, objectType string, ) *WorkflowTaskMetadata`

NewWorkflowTaskMetadata instantiates a new WorkflowTaskMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskMetadataWithDefaults

`func NewWorkflowTaskMetadataWithDefaults() *WorkflowTaskMetadata`

NewWorkflowTaskMetadataWithDefaults instantiates a new WorkflowTaskMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowTaskMetadata) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowTaskMetadata) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowTaskMetadata) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowTaskMetadata) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowTaskMetadata) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowTaskMetadata) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *WorkflowTaskMetadata) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowTaskMetadata) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowTaskMetadata) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowTaskMetadata) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabel

`func (o *WorkflowTaskMetadata) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WorkflowTaskMetadata) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WorkflowTaskMetadata) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WorkflowTaskMetadata) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *WorkflowTaskMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowTaskMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowTaskMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowTaskMetadata) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAssociatedRoles

`func (o *WorkflowTaskMetadata) GetAssociatedRoles() []IamRoleRelationship`

GetAssociatedRoles returns the AssociatedRoles field if non-nil, zero value otherwise.

### GetAssociatedRolesOk

`func (o *WorkflowTaskMetadata) GetAssociatedRolesOk() (*[]IamRoleRelationship, bool)`

GetAssociatedRolesOk returns a tuple with the AssociatedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedRoles

`func (o *WorkflowTaskMetadata) SetAssociatedRoles(v []IamRoleRelationship)`

SetAssociatedRoles sets AssociatedRoles field to given value.

### HasAssociatedRoles

`func (o *WorkflowTaskMetadata) HasAssociatedRoles() bool`

HasAssociatedRoles returns a boolean if a field has been set.

### SetAssociatedRolesNil

`func (o *WorkflowTaskMetadata) SetAssociatedRolesNil(b bool)`

 SetAssociatedRolesNil sets the value for AssociatedRoles to be an explicit nil

### UnsetAssociatedRoles
`func (o *WorkflowTaskMetadata) UnsetAssociatedRoles()`

UnsetAssociatedRoles ensures that no value is present for AssociatedRoles, not even an explicit nil
### GetCatalog

`func (o *WorkflowTaskMetadata) GetCatalog() WorkflowCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *WorkflowTaskMetadata) GetCatalogOk() (*WorkflowCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *WorkflowTaskMetadata) SetCatalog(v WorkflowCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *WorkflowTaskMetadata) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


