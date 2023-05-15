# WorkflowCatalogItemDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.CatalogItemDefinition"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.CatalogItemDefinition"]
**Description** | Pointer to **string** | The description for the catalog item which provides information on what are the pre-requisites to deploy the resource. | [optional] 
**Label** | Pointer to **string** | A user friendly short name to identify the catalog item. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_). | [optional] 
**Name** | Pointer to **string** | The name for this catalog item definition. You can have multiple versions of the catalog item with the same name. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:) or an underscore (_). | [optional] 
**PublishStatus** | Pointer to **string** | Publish status of the catalog item. * &#x60;NotPublished&#x60; - A state of the service item or catalog item which is not yet published. * &#x60;Published&#x60; - A state denoting that the service item or catalog item is published. | [optional] [default to "NotPublished"]
**ServiceItems** | Pointer to [**[]WorkflowServiceItemType**](WorkflowServiceItemType.md) |  | [optional] 
**SupportStatus** | Pointer to **string** | The CatalogItem Support depicts the support status of catalog, the values will be any of Supported or Deprecated state. The user can create a Catalog Service Request if the support status is supported, if its Deprecated then it cannot be instantiated. * &#x60;Supported&#x60; - The definition is a supported version and there will be no changes to the mandatory inputs or outputs. * &#x60;Beta&#x60; - The definition is a Beta version and this version can under go changes until the version is marked supported. * &#x60;Deprecated&#x60; - The version of definition is deprecated and typically there will be a higher version of the same definition that has been added. | [optional] [default to "Supported"]
**UserIdOrEmail** | Pointer to **string** | The user identifier who created or updated the catalog item definition. | [optional] [readonly] 
**ValidationInformation** | Pointer to [**NullableWorkflowValidationInformation**](WorkflowValidationInformation.md) |  | [optional] 
**Version** | Pointer to **int64** | The version of the catalog item to support multiple versions. | [optional] [default to 1]
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 

## Methods

### NewWorkflowCatalogItemDefinition

`func NewWorkflowCatalogItemDefinition(classId string, objectType string, ) *WorkflowCatalogItemDefinition`

NewWorkflowCatalogItemDefinition instantiates a new WorkflowCatalogItemDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowCatalogItemDefinitionWithDefaults

`func NewWorkflowCatalogItemDefinitionWithDefaults() *WorkflowCatalogItemDefinition`

NewWorkflowCatalogItemDefinitionWithDefaults instantiates a new WorkflowCatalogItemDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowCatalogItemDefinition) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowCatalogItemDefinition) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowCatalogItemDefinition) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowCatalogItemDefinition) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowCatalogItemDefinition) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowCatalogItemDefinition) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDescription

`func (o *WorkflowCatalogItemDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowCatalogItemDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowCatalogItemDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowCatalogItemDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabel

`func (o *WorkflowCatalogItemDefinition) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WorkflowCatalogItemDefinition) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WorkflowCatalogItemDefinition) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WorkflowCatalogItemDefinition) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *WorkflowCatalogItemDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowCatalogItemDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowCatalogItemDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowCatalogItemDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublishStatus

`func (o *WorkflowCatalogItemDefinition) GetPublishStatus() string`

GetPublishStatus returns the PublishStatus field if non-nil, zero value otherwise.

### GetPublishStatusOk

`func (o *WorkflowCatalogItemDefinition) GetPublishStatusOk() (*string, bool)`

GetPublishStatusOk returns a tuple with the PublishStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishStatus

`func (o *WorkflowCatalogItemDefinition) SetPublishStatus(v string)`

SetPublishStatus sets PublishStatus field to given value.

### HasPublishStatus

`func (o *WorkflowCatalogItemDefinition) HasPublishStatus() bool`

HasPublishStatus returns a boolean if a field has been set.

### GetServiceItems

`func (o *WorkflowCatalogItemDefinition) GetServiceItems() []WorkflowServiceItemType`

GetServiceItems returns the ServiceItems field if non-nil, zero value otherwise.

### GetServiceItemsOk

`func (o *WorkflowCatalogItemDefinition) GetServiceItemsOk() (*[]WorkflowServiceItemType, bool)`

GetServiceItemsOk returns a tuple with the ServiceItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItems

`func (o *WorkflowCatalogItemDefinition) SetServiceItems(v []WorkflowServiceItemType)`

SetServiceItems sets ServiceItems field to given value.

### HasServiceItems

`func (o *WorkflowCatalogItemDefinition) HasServiceItems() bool`

HasServiceItems returns a boolean if a field has been set.

### SetServiceItemsNil

`func (o *WorkflowCatalogItemDefinition) SetServiceItemsNil(b bool)`

 SetServiceItemsNil sets the value for ServiceItems to be an explicit nil

### UnsetServiceItems
`func (o *WorkflowCatalogItemDefinition) UnsetServiceItems()`

UnsetServiceItems ensures that no value is present for ServiceItems, not even an explicit nil
### GetSupportStatus

`func (o *WorkflowCatalogItemDefinition) GetSupportStatus() string`

GetSupportStatus returns the SupportStatus field if non-nil, zero value otherwise.

### GetSupportStatusOk

`func (o *WorkflowCatalogItemDefinition) GetSupportStatusOk() (*string, bool)`

GetSupportStatusOk returns a tuple with the SupportStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportStatus

`func (o *WorkflowCatalogItemDefinition) SetSupportStatus(v string)`

SetSupportStatus sets SupportStatus field to given value.

### HasSupportStatus

`func (o *WorkflowCatalogItemDefinition) HasSupportStatus() bool`

HasSupportStatus returns a boolean if a field has been set.

### GetUserIdOrEmail

`func (o *WorkflowCatalogItemDefinition) GetUserIdOrEmail() string`

GetUserIdOrEmail returns the UserIdOrEmail field if non-nil, zero value otherwise.

### GetUserIdOrEmailOk

`func (o *WorkflowCatalogItemDefinition) GetUserIdOrEmailOk() (*string, bool)`

GetUserIdOrEmailOk returns a tuple with the UserIdOrEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdOrEmail

`func (o *WorkflowCatalogItemDefinition) SetUserIdOrEmail(v string)`

SetUserIdOrEmail sets UserIdOrEmail field to given value.

### HasUserIdOrEmail

`func (o *WorkflowCatalogItemDefinition) HasUserIdOrEmail() bool`

HasUserIdOrEmail returns a boolean if a field has been set.

### GetValidationInformation

`func (o *WorkflowCatalogItemDefinition) GetValidationInformation() WorkflowValidationInformation`

GetValidationInformation returns the ValidationInformation field if non-nil, zero value otherwise.

### GetValidationInformationOk

`func (o *WorkflowCatalogItemDefinition) GetValidationInformationOk() (*WorkflowValidationInformation, bool)`

GetValidationInformationOk returns a tuple with the ValidationInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationInformation

`func (o *WorkflowCatalogItemDefinition) SetValidationInformation(v WorkflowValidationInformation)`

SetValidationInformation sets ValidationInformation field to given value.

### HasValidationInformation

`func (o *WorkflowCatalogItemDefinition) HasValidationInformation() bool`

HasValidationInformation returns a boolean if a field has been set.

### SetValidationInformationNil

`func (o *WorkflowCatalogItemDefinition) SetValidationInformationNil(b bool)`

 SetValidationInformationNil sets the value for ValidationInformation to be an explicit nil

### UnsetValidationInformation
`func (o *WorkflowCatalogItemDefinition) UnsetValidationInformation()`

UnsetValidationInformation ensures that no value is present for ValidationInformation, not even an explicit nil
### GetVersion

`func (o *WorkflowCatalogItemDefinition) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkflowCatalogItemDefinition) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkflowCatalogItemDefinition) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WorkflowCatalogItemDefinition) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetOrganization

`func (o *WorkflowCatalogItemDefinition) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *WorkflowCatalogItemDefinition) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *WorkflowCatalogItemDefinition) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *WorkflowCatalogItemDefinition) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


