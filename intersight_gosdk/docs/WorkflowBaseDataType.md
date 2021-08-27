# WorkflowBaseDataType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Default** | Pointer to [**NullableWorkflowDefaultValue**](WorkflowDefaultValue.md) |  | [optional] 
**Description** | Pointer to **string** | Provide a detailed description of the data type. | [optional] 
**DisplayMeta** | Pointer to [**NullableWorkflowDisplayMeta**](WorkflowDisplayMeta.md) |  | [optional] 
**InputParameters** | Pointer to **interface{}** | JSON formatted mapping from other property of the definition to the current property. Input parameter mapping is supported only for custom data type property in workflow definition and custom data type definition. The format to specify mapping ina workflow definition when source property is of scalar types is &#39;${workflow.input.property}&#39;. The format to specify mapping when the source property is of object reference and mapping needs to be made to the property of the object is &#39;${workflow.input.property.subproperty}&#39;. The format to specify mapping in a custom data type definition is &#39;${datatype.type.property}&#39;. When the current property is of non-scalar type like composite custom data type, then mapping can be provided to the individual property of the custom data type like &#39;cdt_property:${workflow.input.property}&#39;. | [optional] 
**Label** | Pointer to **string** | Descriptive label for the data type. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), space ( ) or an underscore (_). The first and last character in label must be an alphanumeric character. | [optional] 
**Name** | Pointer to **string** | Descriptive name for the data type. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_). The first and last character in name must be an alphanumeric character. | [optional] 
**Required** | Pointer to **bool** | Specifies whether this parameter is required. The field is applicable for task and workflow. | [optional] 

## Methods

### NewWorkflowBaseDataType

`func NewWorkflowBaseDataType(classId string, objectType string, ) *WorkflowBaseDataType`

NewWorkflowBaseDataType instantiates a new WorkflowBaseDataType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowBaseDataTypeWithDefaults

`func NewWorkflowBaseDataTypeWithDefaults() *WorkflowBaseDataType`

NewWorkflowBaseDataTypeWithDefaults instantiates a new WorkflowBaseDataType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowBaseDataType) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowBaseDataType) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowBaseDataType) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowBaseDataType) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowBaseDataType) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowBaseDataType) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDefault

`func (o *WorkflowBaseDataType) GetDefault() WorkflowDefaultValue`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *WorkflowBaseDataType) GetDefaultOk() (*WorkflowDefaultValue, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *WorkflowBaseDataType) SetDefault(v WorkflowDefaultValue)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *WorkflowBaseDataType) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefaultNil

`func (o *WorkflowBaseDataType) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *WorkflowBaseDataType) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil
### GetDescription

`func (o *WorkflowBaseDataType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowBaseDataType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowBaseDataType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowBaseDataType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayMeta

`func (o *WorkflowBaseDataType) GetDisplayMeta() WorkflowDisplayMeta`

GetDisplayMeta returns the DisplayMeta field if non-nil, zero value otherwise.

### GetDisplayMetaOk

`func (o *WorkflowBaseDataType) GetDisplayMetaOk() (*WorkflowDisplayMeta, bool)`

GetDisplayMetaOk returns a tuple with the DisplayMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayMeta

`func (o *WorkflowBaseDataType) SetDisplayMeta(v WorkflowDisplayMeta)`

SetDisplayMeta sets DisplayMeta field to given value.

### HasDisplayMeta

`func (o *WorkflowBaseDataType) HasDisplayMeta() bool`

HasDisplayMeta returns a boolean if a field has been set.

### SetDisplayMetaNil

`func (o *WorkflowBaseDataType) SetDisplayMetaNil(b bool)`

 SetDisplayMetaNil sets the value for DisplayMeta to be an explicit nil

### UnsetDisplayMeta
`func (o *WorkflowBaseDataType) UnsetDisplayMeta()`

UnsetDisplayMeta ensures that no value is present for DisplayMeta, not even an explicit nil
### GetInputParameters

`func (o *WorkflowBaseDataType) GetInputParameters() interface{}`

GetInputParameters returns the InputParameters field if non-nil, zero value otherwise.

### GetInputParametersOk

`func (o *WorkflowBaseDataType) GetInputParametersOk() (*interface{}, bool)`

GetInputParametersOk returns a tuple with the InputParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputParameters

`func (o *WorkflowBaseDataType) SetInputParameters(v interface{})`

SetInputParameters sets InputParameters field to given value.

### HasInputParameters

`func (o *WorkflowBaseDataType) HasInputParameters() bool`

HasInputParameters returns a boolean if a field has been set.

### SetInputParametersNil

`func (o *WorkflowBaseDataType) SetInputParametersNil(b bool)`

 SetInputParametersNil sets the value for InputParameters to be an explicit nil

### UnsetInputParameters
`func (o *WorkflowBaseDataType) UnsetInputParameters()`

UnsetInputParameters ensures that no value is present for InputParameters, not even an explicit nil
### GetLabel

`func (o *WorkflowBaseDataType) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WorkflowBaseDataType) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WorkflowBaseDataType) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WorkflowBaseDataType) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *WorkflowBaseDataType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowBaseDataType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowBaseDataType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowBaseDataType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRequired

`func (o *WorkflowBaseDataType) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *WorkflowBaseDataType) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *WorkflowBaseDataType) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *WorkflowBaseDataType) HasRequired() bool`

HasRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


