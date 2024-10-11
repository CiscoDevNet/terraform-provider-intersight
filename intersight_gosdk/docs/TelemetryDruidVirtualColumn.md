# TelemetryDruidVirtualColumn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The virtual-column type. | 
**Name** | Pointer to **string** | name of the virtual column. | [optional] 
**Expression** | Pointer to **string** | The Druid expression. | [optional] 
**OutputType** | Pointer to **string** | Type of the resulting column. | [optional] [default to "FLOAT"]
**ColumnName** | Pointer to **string** | Name of the COMPLEX&lt;json&gt; input column. | [optional] 
**OutputName** | Pointer to **string** | Name of the virtual column. | [optional] 
**ExpectedType** | Pointer to **string** | Native druid output type of the column. Druid will coerce output to this type if it does not match the underlying type. | [optional] [default to "STRING"]
**PathParts** | Pointer to [**[]TelemetryNestedFieldVirtualColumnAllOfPathParts**](TelemetryNestedFieldVirtualColumnAllOfPathParts.md) | A list of path parts that represent the path to the desired value in the json column. Each path part can be a string or an integer. If a path part is an integer, it is treated as an index into an array. | [optional] 
**ProcessFromRaw** | Pointer to **bool** | If set to true, the virtual column will process the \&quot;raw\&quot; JSON data to extract values rather than using an optimized \&quot;literal\&quot; value selector. This option allows extracting non-literal values from JSON the cost of much slower performance. | [optional] [default to false]
**Path** | Pointer to **string** | A JSONPath or jq syntax string representation of the path to the desired value in the json column. This field is optional and can be used instead of pathParts. | [optional] 
**UseJqSyntax** | Pointer to **bool** | If set to true, the virtual column will use jq syntax instead of JSONPath. | [optional] [default to false]
**Delegate** | Pointer to **string** | Name of the multi-value string input column to filter. | [optional] 
**Values** | Pointer to **[]string** | Set of string value to allow or deny. | [optional] 
**IsAllowList** | Pointer to **bool** | If set to true, the virtual column will allow only the values in the list. If set to false, the virtual column will provide all values except those specified. | [optional] [default to true]

## Methods

### NewTelemetryDruidVirtualColumn

`func NewTelemetryDruidVirtualColumn(type_ string, ) *TelemetryDruidVirtualColumn`

NewTelemetryDruidVirtualColumn instantiates a new TelemetryDruidVirtualColumn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidVirtualColumnWithDefaults

`func NewTelemetryDruidVirtualColumnWithDefaults() *TelemetryDruidVirtualColumn`

NewTelemetryDruidVirtualColumnWithDefaults instantiates a new TelemetryDruidVirtualColumn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidVirtualColumn) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidVirtualColumn) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidVirtualColumn) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *TelemetryDruidVirtualColumn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TelemetryDruidVirtualColumn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TelemetryDruidVirtualColumn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TelemetryDruidVirtualColumn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExpression

`func (o *TelemetryDruidVirtualColumn) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *TelemetryDruidVirtualColumn) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *TelemetryDruidVirtualColumn) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *TelemetryDruidVirtualColumn) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetOutputType

`func (o *TelemetryDruidVirtualColumn) GetOutputType() string`

GetOutputType returns the OutputType field if non-nil, zero value otherwise.

### GetOutputTypeOk

`func (o *TelemetryDruidVirtualColumn) GetOutputTypeOk() (*string, bool)`

GetOutputTypeOk returns a tuple with the OutputType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputType

`func (o *TelemetryDruidVirtualColumn) SetOutputType(v string)`

SetOutputType sets OutputType field to given value.

### HasOutputType

`func (o *TelemetryDruidVirtualColumn) HasOutputType() bool`

HasOutputType returns a boolean if a field has been set.

### GetColumnName

`func (o *TelemetryDruidVirtualColumn) GetColumnName() string`

GetColumnName returns the ColumnName field if non-nil, zero value otherwise.

### GetColumnNameOk

`func (o *TelemetryDruidVirtualColumn) GetColumnNameOk() (*string, bool)`

GetColumnNameOk returns a tuple with the ColumnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnName

`func (o *TelemetryDruidVirtualColumn) SetColumnName(v string)`

SetColumnName sets ColumnName field to given value.

### HasColumnName

`func (o *TelemetryDruidVirtualColumn) HasColumnName() bool`

HasColumnName returns a boolean if a field has been set.

### GetOutputName

`func (o *TelemetryDruidVirtualColumn) GetOutputName() string`

GetOutputName returns the OutputName field if non-nil, zero value otherwise.

### GetOutputNameOk

`func (o *TelemetryDruidVirtualColumn) GetOutputNameOk() (*string, bool)`

GetOutputNameOk returns a tuple with the OutputName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputName

`func (o *TelemetryDruidVirtualColumn) SetOutputName(v string)`

SetOutputName sets OutputName field to given value.

### HasOutputName

`func (o *TelemetryDruidVirtualColumn) HasOutputName() bool`

HasOutputName returns a boolean if a field has been set.

### GetExpectedType

`func (o *TelemetryDruidVirtualColumn) GetExpectedType() string`

GetExpectedType returns the ExpectedType field if non-nil, zero value otherwise.

### GetExpectedTypeOk

`func (o *TelemetryDruidVirtualColumn) GetExpectedTypeOk() (*string, bool)`

GetExpectedTypeOk returns a tuple with the ExpectedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedType

`func (o *TelemetryDruidVirtualColumn) SetExpectedType(v string)`

SetExpectedType sets ExpectedType field to given value.

### HasExpectedType

`func (o *TelemetryDruidVirtualColumn) HasExpectedType() bool`

HasExpectedType returns a boolean if a field has been set.

### GetPathParts

`func (o *TelemetryDruidVirtualColumn) GetPathParts() []TelemetryNestedFieldVirtualColumnAllOfPathParts`

GetPathParts returns the PathParts field if non-nil, zero value otherwise.

### GetPathPartsOk

`func (o *TelemetryDruidVirtualColumn) GetPathPartsOk() (*[]TelemetryNestedFieldVirtualColumnAllOfPathParts, bool)`

GetPathPartsOk returns a tuple with the PathParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathParts

`func (o *TelemetryDruidVirtualColumn) SetPathParts(v []TelemetryNestedFieldVirtualColumnAllOfPathParts)`

SetPathParts sets PathParts field to given value.

### HasPathParts

`func (o *TelemetryDruidVirtualColumn) HasPathParts() bool`

HasPathParts returns a boolean if a field has been set.

### GetProcessFromRaw

`func (o *TelemetryDruidVirtualColumn) GetProcessFromRaw() bool`

GetProcessFromRaw returns the ProcessFromRaw field if non-nil, zero value otherwise.

### GetProcessFromRawOk

`func (o *TelemetryDruidVirtualColumn) GetProcessFromRawOk() (*bool, bool)`

GetProcessFromRawOk returns a tuple with the ProcessFromRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessFromRaw

`func (o *TelemetryDruidVirtualColumn) SetProcessFromRaw(v bool)`

SetProcessFromRaw sets ProcessFromRaw field to given value.

### HasProcessFromRaw

`func (o *TelemetryDruidVirtualColumn) HasProcessFromRaw() bool`

HasProcessFromRaw returns a boolean if a field has been set.

### GetPath

`func (o *TelemetryDruidVirtualColumn) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *TelemetryDruidVirtualColumn) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *TelemetryDruidVirtualColumn) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *TelemetryDruidVirtualColumn) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetUseJqSyntax

`func (o *TelemetryDruidVirtualColumn) GetUseJqSyntax() bool`

GetUseJqSyntax returns the UseJqSyntax field if non-nil, zero value otherwise.

### GetUseJqSyntaxOk

`func (o *TelemetryDruidVirtualColumn) GetUseJqSyntaxOk() (*bool, bool)`

GetUseJqSyntaxOk returns a tuple with the UseJqSyntax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseJqSyntax

`func (o *TelemetryDruidVirtualColumn) SetUseJqSyntax(v bool)`

SetUseJqSyntax sets UseJqSyntax field to given value.

### HasUseJqSyntax

`func (o *TelemetryDruidVirtualColumn) HasUseJqSyntax() bool`

HasUseJqSyntax returns a boolean if a field has been set.

### GetDelegate

`func (o *TelemetryDruidVirtualColumn) GetDelegate() string`

GetDelegate returns the Delegate field if non-nil, zero value otherwise.

### GetDelegateOk

`func (o *TelemetryDruidVirtualColumn) GetDelegateOk() (*string, bool)`

GetDelegateOk returns a tuple with the Delegate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegate

`func (o *TelemetryDruidVirtualColumn) SetDelegate(v string)`

SetDelegate sets Delegate field to given value.

### HasDelegate

`func (o *TelemetryDruidVirtualColumn) HasDelegate() bool`

HasDelegate returns a boolean if a field has been set.

### GetValues

`func (o *TelemetryDruidVirtualColumn) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *TelemetryDruidVirtualColumn) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *TelemetryDruidVirtualColumn) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *TelemetryDruidVirtualColumn) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetIsAllowList

`func (o *TelemetryDruidVirtualColumn) GetIsAllowList() bool`

GetIsAllowList returns the IsAllowList field if non-nil, zero value otherwise.

### GetIsAllowListOk

`func (o *TelemetryDruidVirtualColumn) GetIsAllowListOk() (*bool, bool)`

GetIsAllowListOk returns a tuple with the IsAllowList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllowList

`func (o *TelemetryDruidVirtualColumn) SetIsAllowList(v bool)`

SetIsAllowList sets IsAllowList field to given value.

### HasIsAllowList

`func (o *TelemetryDruidVirtualColumn) HasIsAllowList() bool`

HasIsAllowList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


