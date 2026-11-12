# WorkflowSupportObjectReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.SupportObjectReference"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.SupportObjectReference"]
**CatalogMoid** | Pointer to **string** | The catalog MOID of the referenced object. If this field is not specified, the resolver uses the catalog of the versioned object first, followed by the default catalog for the referenced object type. | [optional] 
**Name** | Pointer to **string** | The qualified name of the referenced object. | [optional] 
**Object** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 
**Version** | Pointer to **int64** | The version of the referenced object. If not specified, the default version is used. | [optional] 

## Methods

### NewWorkflowSupportObjectReference

`func NewWorkflowSupportObjectReference(classId string, objectType string, ) *WorkflowSupportObjectReference`

NewWorkflowSupportObjectReference instantiates a new WorkflowSupportObjectReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSupportObjectReferenceWithDefaults

`func NewWorkflowSupportObjectReferenceWithDefaults() *WorkflowSupportObjectReference`

NewWorkflowSupportObjectReferenceWithDefaults instantiates a new WorkflowSupportObjectReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowSupportObjectReference) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowSupportObjectReference) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowSupportObjectReference) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowSupportObjectReference) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowSupportObjectReference) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowSupportObjectReference) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCatalogMoid

`func (o *WorkflowSupportObjectReference) GetCatalogMoid() string`

GetCatalogMoid returns the CatalogMoid field if non-nil, zero value otherwise.

### GetCatalogMoidOk

`func (o *WorkflowSupportObjectReference) GetCatalogMoidOk() (*string, bool)`

GetCatalogMoidOk returns a tuple with the CatalogMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogMoid

`func (o *WorkflowSupportObjectReference) SetCatalogMoid(v string)`

SetCatalogMoid sets CatalogMoid field to given value.

### HasCatalogMoid

`func (o *WorkflowSupportObjectReference) HasCatalogMoid() bool`

HasCatalogMoid returns a boolean if a field has been set.

### GetName

`func (o *WorkflowSupportObjectReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowSupportObjectReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowSupportObjectReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowSupportObjectReference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObject

`func (o *WorkflowSupportObjectReference) GetObject() MoMoRef`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *WorkflowSupportObjectReference) GetObjectOk() (*MoMoRef, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *WorkflowSupportObjectReference) SetObject(v MoMoRef)`

SetObject sets Object field to given value.

### HasObject

`func (o *WorkflowSupportObjectReference) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetVersion

`func (o *WorkflowSupportObjectReference) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkflowSupportObjectReference) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkflowSupportObjectReference) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WorkflowSupportObjectReference) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


