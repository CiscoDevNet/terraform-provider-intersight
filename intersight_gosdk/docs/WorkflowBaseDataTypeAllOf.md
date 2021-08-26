# WorkflowBaseDataTypeAllOf

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

### NewWorkflowBaseDataTypeAllOf

`func NewWorkflowBaseDataTypeAllOf(classId string, objectType string, ) *WorkflowBaseDataTypeAllOf`

NewWorkflowBaseDataTypeAllOf instantiates a new WorkflowBaseDataTypeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowBaseDataTypeAllOfWithDefaults

`func NewWorkflowBaseDataTypeAllOfWithDefaults() *WorkflowBaseDataTypeAllOf`

NewWorkflowBaseDataTypeAllOfWithDefaults instantiates a new WorkflowBaseDataTypeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowBaseDataTypeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowBaseDataTypeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowBaseDataTypeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowBaseDataTypeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowBaseDataTypeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowBaseDataTypeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDefault

`func (o *WorkflowBaseDataTypeAllOf) GetDefault() WorkflowDefaultValue`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *WorkflowBaseDataTypeAllOf) GetDefaultOk() (*WorkflowDefaultValue, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *WorkflowBaseDataTypeAllOf) SetDefault(v WorkflowDefaultValue)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *WorkflowBaseDataTypeAllOf) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### SetDefaultNil

`func (o *WorkflowBaseDataTypeAllOf) SetDefaultNil(b bool)`

 SetDefaultNil sets the value for Default to be an explicit nil

### UnsetDefault
`func (o *WorkflowBaseDataTypeAllOf) UnsetDefault()`

UnsetDefault ensures that no value is present for Default, not even an explicit nil
### GetDescription

`func (o *WorkflowBaseDataTypeAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowBaseDataTypeAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowBaseDataTypeAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowBaseDataTypeAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayMeta

`func (o *WorkflowBaseDataTypeAllOf) GetDisplayMeta() WorkflowDisplayMeta`

GetDisplayMeta returns the DisplayMeta field if non-nil, zero value otherwise.

### GetDisplayMetaOk

`func (o *WorkflowBaseDataTypeAllOf) GetDisplayMetaOk() (*WorkflowDisplayMeta, bool)`

GetDisplayMetaOk returns a tuple with the DisplayMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayMeta

`func (o *WorkflowBaseDataTypeAllOf) SetDisplayMeta(v WorkflowDisplayMeta)`

SetDisplayMeta sets DisplayMeta field to given value.

### HasDisplayMeta

`func (o *WorkflowBaseDataTypeAllOf) HasDisplayMeta() bool`

HasDisplayMeta returns a boolean if a field has been set.

### SetDisplayMetaNil

`func (o *WorkflowBaseDataTypeAllOf) SetDisplayMetaNil(b bool)`

 SetDisplayMetaNil sets the value for DisplayMeta to be an explicit nil

### UnsetDisplayMeta
`func (o *WorkflowBaseDataTypeAllOf) UnsetDisplayMeta()`

UnsetDisplayMeta ensures that no value is present for DisplayMeta, not even an explicit nil
### GetInputParameters

`func (o *WorkflowBaseDataTypeAllOf) GetInputParameters() interface{}`

GetInputParameters returns the InputParameters field if non-nil, zero value otherwise.

### GetInputParametersOk

`func (o *WorkflowBaseDataTypeAllOf) GetInputParametersOk() (*interface{}, bool)`

GetInputParametersOk returns a tuple with the InputParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputParameters

`func (o *WorkflowBaseDataTypeAllOf) SetInputParameters(v interface{})`

SetInputParameters sets InputParameters field to given value.

### HasInputParameters

`func (o *WorkflowBaseDataTypeAllOf) HasInputParameters() bool`

HasInputParameters returns a boolean if a field has been set.

### SetInputParametersNil

`func (o *WorkflowBaseDataTypeAllOf) SetInputParametersNil(b bool)`

 SetInputParametersNil sets the value for InputParameters to be an explicit nil

### UnsetInputParameters
`func (o *WorkflowBaseDataTypeAllOf) UnsetInputParameters()`

UnsetInputParameters ensures that no value is present for InputParameters, not even an explicit nil
### GetLabel

`func (o *WorkflowBaseDataTypeAllOf) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WorkflowBaseDataTypeAllOf) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WorkflowBaseDataTypeAllOf) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WorkflowBaseDataTypeAllOf) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *WorkflowBaseDataTypeAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowBaseDataTypeAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowBaseDataTypeAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowBaseDataTypeAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRequired

`func (o *WorkflowBaseDataTypeAllOf) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *WorkflowBaseDataTypeAllOf) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *WorkflowBaseDataTypeAllOf) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *WorkflowBaseDataTypeAllOf) HasRequired() bool`

HasRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


