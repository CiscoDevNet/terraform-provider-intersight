# TelemetryDruidExpressionVirtualColumn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The virtual-column type. | 
**Name** | Pointer to **string** | Name of the virtual column. | [optional] 
**Expression** | Pointer to **string** | The Druid expression. | [optional] 
**OutputType** | Pointer to **string** | Type of the resulting column. | [optional] [default to "FLOAT"]

## Methods

### NewTelemetryDruidExpressionVirtualColumn

`func NewTelemetryDruidExpressionVirtualColumn(type_ string, ) *TelemetryDruidExpressionVirtualColumn`

NewTelemetryDruidExpressionVirtualColumn instantiates a new TelemetryDruidExpressionVirtualColumn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidExpressionVirtualColumnWithDefaults

`func NewTelemetryDruidExpressionVirtualColumnWithDefaults() *TelemetryDruidExpressionVirtualColumn`

NewTelemetryDruidExpressionVirtualColumnWithDefaults instantiates a new TelemetryDruidExpressionVirtualColumn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidExpressionVirtualColumn) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidExpressionVirtualColumn) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidExpressionVirtualColumn) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *TelemetryDruidExpressionVirtualColumn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TelemetryDruidExpressionVirtualColumn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TelemetryDruidExpressionVirtualColumn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TelemetryDruidExpressionVirtualColumn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExpression

`func (o *TelemetryDruidExpressionVirtualColumn) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *TelemetryDruidExpressionVirtualColumn) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *TelemetryDruidExpressionVirtualColumn) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *TelemetryDruidExpressionVirtualColumn) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetOutputType

`func (o *TelemetryDruidExpressionVirtualColumn) GetOutputType() string`

GetOutputType returns the OutputType field if non-nil, zero value otherwise.

### GetOutputTypeOk

`func (o *TelemetryDruidExpressionVirtualColumn) GetOutputTypeOk() (*string, bool)`

GetOutputTypeOk returns a tuple with the OutputType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputType

`func (o *TelemetryDruidExpressionVirtualColumn) SetOutputType(v string)`

SetOutputType sets OutputType field to given value.

### HasOutputType

`func (o *TelemetryDruidExpressionVirtualColumn) HasOutputType() bool`

HasOutputType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


