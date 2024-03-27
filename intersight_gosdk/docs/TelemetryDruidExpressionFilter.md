# TelemetryDruidExpressionFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Expression** | **string** | Expression string to evaluate into true or false. See the Druid expression system for more details. | 

## Methods

### NewTelemetryDruidExpressionFilter

`func NewTelemetryDruidExpressionFilter(type_ string, expression string, ) *TelemetryDruidExpressionFilter`

NewTelemetryDruidExpressionFilter instantiates a new TelemetryDruidExpressionFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidExpressionFilterWithDefaults

`func NewTelemetryDruidExpressionFilterWithDefaults() *TelemetryDruidExpressionFilter`

NewTelemetryDruidExpressionFilterWithDefaults instantiates a new TelemetryDruidExpressionFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidExpressionFilter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidExpressionFilter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidExpressionFilter) SetType(v string)`

SetType sets Type field to given value.


### GetExpression

`func (o *TelemetryDruidExpressionFilter) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *TelemetryDruidExpressionFilter) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *TelemetryDruidExpressionFilter) SetExpression(v string)`

SetExpression sets Expression field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


