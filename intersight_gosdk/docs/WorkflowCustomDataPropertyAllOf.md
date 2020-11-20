# WorkflowCustomDataPropertyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.CustomDataProperty"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.CustomDataProperty"]
**CatalogMoid** | Pointer to **string** | Specify the catalog moid that this custom data type belongs. | [optional] 
**CustomDataTypeId** | Pointer to **string** | The resolved custom data type definition managed object. | [optional] [readonly] 
**CustomDataTypeName** | Pointer to **string** | Name of the custom data type for this input. | [optional] 

## Methods

### NewWorkflowCustomDataPropertyAllOf

`func NewWorkflowCustomDataPropertyAllOf(classId string, objectType string, ) *WorkflowCustomDataPropertyAllOf`

NewWorkflowCustomDataPropertyAllOf instantiates a new WorkflowCustomDataPropertyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowCustomDataPropertyAllOfWithDefaults

`func NewWorkflowCustomDataPropertyAllOfWithDefaults() *WorkflowCustomDataPropertyAllOf`

NewWorkflowCustomDataPropertyAllOfWithDefaults instantiates a new WorkflowCustomDataPropertyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowCustomDataPropertyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowCustomDataPropertyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowCustomDataPropertyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowCustomDataPropertyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowCustomDataPropertyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowCustomDataPropertyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCatalogMoid

`func (o *WorkflowCustomDataPropertyAllOf) GetCatalogMoid() string`

GetCatalogMoid returns the CatalogMoid field if non-nil, zero value otherwise.

### GetCatalogMoidOk

`func (o *WorkflowCustomDataPropertyAllOf) GetCatalogMoidOk() (*string, bool)`

GetCatalogMoidOk returns a tuple with the CatalogMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogMoid

`func (o *WorkflowCustomDataPropertyAllOf) SetCatalogMoid(v string)`

SetCatalogMoid sets CatalogMoid field to given value.

### HasCatalogMoid

`func (o *WorkflowCustomDataPropertyAllOf) HasCatalogMoid() bool`

HasCatalogMoid returns a boolean if a field has been set.

### GetCustomDataTypeId

`func (o *WorkflowCustomDataPropertyAllOf) GetCustomDataTypeId() string`

GetCustomDataTypeId returns the CustomDataTypeId field if non-nil, zero value otherwise.

### GetCustomDataTypeIdOk

`func (o *WorkflowCustomDataPropertyAllOf) GetCustomDataTypeIdOk() (*string, bool)`

GetCustomDataTypeIdOk returns a tuple with the CustomDataTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDataTypeId

`func (o *WorkflowCustomDataPropertyAllOf) SetCustomDataTypeId(v string)`

SetCustomDataTypeId sets CustomDataTypeId field to given value.

### HasCustomDataTypeId

`func (o *WorkflowCustomDataPropertyAllOf) HasCustomDataTypeId() bool`

HasCustomDataTypeId returns a boolean if a field has been set.

### GetCustomDataTypeName

`func (o *WorkflowCustomDataPropertyAllOf) GetCustomDataTypeName() string`

GetCustomDataTypeName returns the CustomDataTypeName field if non-nil, zero value otherwise.

### GetCustomDataTypeNameOk

`func (o *WorkflowCustomDataPropertyAllOf) GetCustomDataTypeNameOk() (*string, bool)`

GetCustomDataTypeNameOk returns a tuple with the CustomDataTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDataTypeName

`func (o *WorkflowCustomDataPropertyAllOf) SetCustomDataTypeName(v string)`

SetCustomDataTypeName sets CustomDataTypeName field to given value.

### HasCustomDataTypeName

`func (o *WorkflowCustomDataPropertyAllOf) HasCustomDataTypeName() bool`

HasCustomDataTypeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


