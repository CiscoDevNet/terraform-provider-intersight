# WorkflowCustomDataTypeProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.CustomDataTypeProperties"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.CustomDataTypeProperties"]
**Cloneable** | Pointer to **bool** | When set to false custom data type is not cloneable. It is set to true only if data type is not internal and it is not using any internal custom data type. | [optional] [readonly] [default to true]
**ExternalMeta** | Pointer to **bool** | When set to false the custom data type is owned by the system and used for internal services. Such custom data type cannot be directly used by external entities. | [optional] [readonly] [default to false]

## Methods

### NewWorkflowCustomDataTypeProperties

`func NewWorkflowCustomDataTypeProperties(classId string, objectType string, ) *WorkflowCustomDataTypeProperties`

NewWorkflowCustomDataTypeProperties instantiates a new WorkflowCustomDataTypeProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowCustomDataTypePropertiesWithDefaults

`func NewWorkflowCustomDataTypePropertiesWithDefaults() *WorkflowCustomDataTypeProperties`

NewWorkflowCustomDataTypePropertiesWithDefaults instantiates a new WorkflowCustomDataTypeProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowCustomDataTypeProperties) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowCustomDataTypeProperties) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowCustomDataTypeProperties) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowCustomDataTypeProperties) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowCustomDataTypeProperties) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowCustomDataTypeProperties) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCloneable

`func (o *WorkflowCustomDataTypeProperties) GetCloneable() bool`

GetCloneable returns the Cloneable field if non-nil, zero value otherwise.

### GetCloneableOk

`func (o *WorkflowCustomDataTypeProperties) GetCloneableOk() (*bool, bool)`

GetCloneableOk returns a tuple with the Cloneable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneable

`func (o *WorkflowCustomDataTypeProperties) SetCloneable(v bool)`

SetCloneable sets Cloneable field to given value.

### HasCloneable

`func (o *WorkflowCustomDataTypeProperties) HasCloneable() bool`

HasCloneable returns a boolean if a field has been set.

### GetExternalMeta

`func (o *WorkflowCustomDataTypeProperties) GetExternalMeta() bool`

GetExternalMeta returns the ExternalMeta field if non-nil, zero value otherwise.

### GetExternalMetaOk

`func (o *WorkflowCustomDataTypeProperties) GetExternalMetaOk() (*bool, bool)`

GetExternalMetaOk returns a tuple with the ExternalMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalMeta

`func (o *WorkflowCustomDataTypeProperties) SetExternalMeta(v bool)`

SetExternalMeta sets ExternalMeta field to given value.

### HasExternalMeta

`func (o *WorkflowCustomDataTypeProperties) HasExternalMeta() bool`

HasExternalMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


