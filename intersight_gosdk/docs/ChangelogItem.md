# ChangelogItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "changelog.Item"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "changelog.Item"]
**DateVersion** | Pointer to **time.Time** | The date version for the API contract changelog item in the format rfc3339 with no fraction seconds set.  Note that there can be more than one item per DateVersion. Example: 2023-12-19T00:00:00Z . | [optional] 
**Entity** | Pointer to **string** | The operationId of the endpoint for which changelog item is being generated. | [optional] 
**SemanticVersion** | Pointer to **string** | The semantic version for the API contract changelog item. Note that there can be more than one item per SemanticVersion. | [optional] 
**Value** | Pointer to **string** | The value of the API contract changelog item. | [optional] 
**Catalog** | Pointer to [**NullableWorkflowCatalogRelationship**](WorkflowCatalogRelationship.md) |  | [optional] 

## Methods

### NewChangelogItem

`func NewChangelogItem(classId string, objectType string, ) *ChangelogItem`

NewChangelogItem instantiates a new ChangelogItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangelogItemWithDefaults

`func NewChangelogItemWithDefaults() *ChangelogItem`

NewChangelogItemWithDefaults instantiates a new ChangelogItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ChangelogItem) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ChangelogItem) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ChangelogItem) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ChangelogItem) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ChangelogItem) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ChangelogItem) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDateVersion

`func (o *ChangelogItem) GetDateVersion() time.Time`

GetDateVersion returns the DateVersion field if non-nil, zero value otherwise.

### GetDateVersionOk

`func (o *ChangelogItem) GetDateVersionOk() (*time.Time, bool)`

GetDateVersionOk returns a tuple with the DateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateVersion

`func (o *ChangelogItem) SetDateVersion(v time.Time)`

SetDateVersion sets DateVersion field to given value.

### HasDateVersion

`func (o *ChangelogItem) HasDateVersion() bool`

HasDateVersion returns a boolean if a field has been set.

### GetEntity

`func (o *ChangelogItem) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *ChangelogItem) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *ChangelogItem) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *ChangelogItem) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetSemanticVersion

`func (o *ChangelogItem) GetSemanticVersion() string`

GetSemanticVersion returns the SemanticVersion field if non-nil, zero value otherwise.

### GetSemanticVersionOk

`func (o *ChangelogItem) GetSemanticVersionOk() (*string, bool)`

GetSemanticVersionOk returns a tuple with the SemanticVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSemanticVersion

`func (o *ChangelogItem) SetSemanticVersion(v string)`

SetSemanticVersion sets SemanticVersion field to given value.

### HasSemanticVersion

`func (o *ChangelogItem) HasSemanticVersion() bool`

HasSemanticVersion returns a boolean if a field has been set.

### GetValue

`func (o *ChangelogItem) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ChangelogItem) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ChangelogItem) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ChangelogItem) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCatalog

`func (o *ChangelogItem) GetCatalog() WorkflowCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *ChangelogItem) GetCatalogOk() (*WorkflowCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *ChangelogItem) SetCatalog(v WorkflowCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *ChangelogItem) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### SetCatalogNil

`func (o *ChangelogItem) SetCatalogNil(b bool)`

 SetCatalogNil sets the value for Catalog to be an explicit nil

### UnsetCatalog
`func (o *ChangelogItem) UnsetCatalog()`

UnsetCatalog ensures that no value is present for Catalog, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


