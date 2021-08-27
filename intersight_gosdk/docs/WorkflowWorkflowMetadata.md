# WorkflowWorkflowMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.WorkflowMetadata"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.WorkflowMetadata"]
**Description** | Pointer to **string** | The description for this workflow. | [optional] 
**Label** | Pointer to **string** | A user friendly short name to identify the workflow metadata. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_). | [optional] 
**Name** | Pointer to **string** | The name for this workflow metadata. You can have multiple versions of the workflow with the same name. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.) or an underscore (_). | [optional] 
**AssociatedRoles** | Pointer to [**[]IamRoleRelationship**](IamRoleRelationship.md) | An array of relationships to iamRole resources. | [optional] 
**Catalog** | Pointer to [**WorkflowCatalogRelationship**](WorkflowCatalogRelationship.md) |  | [optional] 

## Methods

### NewWorkflowWorkflowMetadata

`func NewWorkflowWorkflowMetadata(classId string, objectType string, ) *WorkflowWorkflowMetadata`

NewWorkflowWorkflowMetadata instantiates a new WorkflowWorkflowMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowWorkflowMetadataWithDefaults

`func NewWorkflowWorkflowMetadataWithDefaults() *WorkflowWorkflowMetadata`

NewWorkflowWorkflowMetadataWithDefaults instantiates a new WorkflowWorkflowMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowWorkflowMetadata) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowWorkflowMetadata) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowWorkflowMetadata) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowWorkflowMetadata) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowWorkflowMetadata) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowWorkflowMetadata) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *WorkflowWorkflowMetadata) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowWorkflowMetadata) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowWorkflowMetadata) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowWorkflowMetadata) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabel

`func (o *WorkflowWorkflowMetadata) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WorkflowWorkflowMetadata) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WorkflowWorkflowMetadata) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WorkflowWorkflowMetadata) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *WorkflowWorkflowMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowWorkflowMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowWorkflowMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowWorkflowMetadata) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAssociatedRoles

`func (o *WorkflowWorkflowMetadata) GetAssociatedRoles() []IamRoleRelationship`

GetAssociatedRoles returns the AssociatedRoles field if non-nil, zero value otherwise.

### GetAssociatedRolesOk

`func (o *WorkflowWorkflowMetadata) GetAssociatedRolesOk() (*[]IamRoleRelationship, bool)`

GetAssociatedRolesOk returns a tuple with the AssociatedRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedRoles

`func (o *WorkflowWorkflowMetadata) SetAssociatedRoles(v []IamRoleRelationship)`

SetAssociatedRoles sets AssociatedRoles field to given value.

### HasAssociatedRoles

`func (o *WorkflowWorkflowMetadata) HasAssociatedRoles() bool`

HasAssociatedRoles returns a boolean if a field has been set.

### SetAssociatedRolesNil

`func (o *WorkflowWorkflowMetadata) SetAssociatedRolesNil(b bool)`

 SetAssociatedRolesNil sets the value for AssociatedRoles to be an explicit nil

### UnsetAssociatedRoles
`func (o *WorkflowWorkflowMetadata) UnsetAssociatedRoles()`

UnsetAssociatedRoles ensures that no value is present for AssociatedRoles, not even an explicit nil
### GetCatalog

`func (o *WorkflowWorkflowMetadata) GetCatalog() WorkflowCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *WorkflowWorkflowMetadata) GetCatalogOk() (*WorkflowCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *WorkflowWorkflowMetadata) SetCatalog(v WorkflowCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *WorkflowWorkflowMetadata) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


