# TelemetryNestedFieldVirtualColumn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The virtual-column type. | 
**ColumnName** | Pointer to **string** | Name of the COMPLEX&lt;json&gt; input column. | [optional] 
**OutputName** | Pointer to **string** | Name of the virtual column. | [optional] 
**ExpectedType** | Pointer to **string** | Native druid output type of the column. Druid will coerce output to this type if it does not match the underlying type. | [optional] [default to "STRING"]
**PathParts** | Pointer to [**[]TelemetryNestedFieldVirtualColumnAllOfPathParts**](TelemetryNestedFieldVirtualColumnAllOfPathParts.md) | A list of path parts that represent the path to the desired value in the json column. Each path part can be a string or an integer. If a path part is an integer, it is treated as an index into an array. | [optional] 
**ProcessFromRaw** | Pointer to **bool** | If set to true, the virtual column will process the \&quot;raw\&quot; JSON data to extract values rather than using an optimized \&quot;literal\&quot; value selector. This option allows extracting non-literal values from JSON the cost of much slower performance. | [optional] [default to false]
**Path** | Pointer to **string** | A JSONPath or jq syntax string representation of the path to the desired value in the json column. This field is optional and can be used instead of pathParts. | [optional] 
**UseJqSyntax** | Pointer to **bool** | If set to true, the virtual column will use jq syntax instead of JSONPath. | [optional] [default to false]

## Methods

### NewTelemetryNestedFieldVirtualColumn

`func NewTelemetryNestedFieldVirtualColumn(type_ string, ) *TelemetryNestedFieldVirtualColumn`

NewTelemetryNestedFieldVirtualColumn instantiates a new TelemetryNestedFieldVirtualColumn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryNestedFieldVirtualColumnWithDefaults

`func NewTelemetryNestedFieldVirtualColumnWithDefaults() *TelemetryNestedFieldVirtualColumn`

NewTelemetryNestedFieldVirtualColumnWithDefaults instantiates a new TelemetryNestedFieldVirtualColumn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryNestedFieldVirtualColumn) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryNestedFieldVirtualColumn) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryNestedFieldVirtualColumn) SetType(v string)`

SetType sets Type field to given value.


### GetColumnName

`func (o *TelemetryNestedFieldVirtualColumn) GetColumnName() string`

GetColumnName returns the ColumnName field if non-nil, zero value otherwise.

### GetColumnNameOk

`func (o *TelemetryNestedFieldVirtualColumn) GetColumnNameOk() (*string, bool)`

GetColumnNameOk returns a tuple with the ColumnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnName

`func (o *TelemetryNestedFieldVirtualColumn) SetColumnName(v string)`

SetColumnName sets ColumnName field to given value.

### HasColumnName

`func (o *TelemetryNestedFieldVirtualColumn) HasColumnName() bool`

HasColumnName returns a boolean if a field has been set.

### GetOutputName

`func (o *TelemetryNestedFieldVirtualColumn) GetOutputName() string`

GetOutputName returns the OutputName field if non-nil, zero value otherwise.

### GetOutputNameOk

`func (o *TelemetryNestedFieldVirtualColumn) GetOutputNameOk() (*string, bool)`

GetOutputNameOk returns a tuple with the OutputName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputName

`func (o *TelemetryNestedFieldVirtualColumn) SetOutputName(v string)`

SetOutputName sets OutputName field to given value.

### HasOutputName

`func (o *TelemetryNestedFieldVirtualColumn) HasOutputName() bool`

HasOutputName returns a boolean if a field has been set.

### GetExpectedType

`func (o *TelemetryNestedFieldVirtualColumn) GetExpectedType() string`

GetExpectedType returns the ExpectedType field if non-nil, zero value otherwise.

### GetExpectedTypeOk

`func (o *TelemetryNestedFieldVirtualColumn) GetExpectedTypeOk() (*string, bool)`

GetExpectedTypeOk returns a tuple with the ExpectedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedType

`func (o *TelemetryNestedFieldVirtualColumn) SetExpectedType(v string)`

SetExpectedType sets ExpectedType field to given value.

### HasExpectedType

`func (o *TelemetryNestedFieldVirtualColumn) HasExpectedType() bool`

HasExpectedType returns a boolean if a field has been set.

### GetPathParts

`func (o *TelemetryNestedFieldVirtualColumn) GetPathParts() []TelemetryNestedFieldVirtualColumnAllOfPathParts`

GetPathParts returns the PathParts field if non-nil, zero value otherwise.

### GetPathPartsOk

`func (o *TelemetryNestedFieldVirtualColumn) GetPathPartsOk() (*[]TelemetryNestedFieldVirtualColumnAllOfPathParts, bool)`

GetPathPartsOk returns a tuple with the PathParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathParts

`func (o *TelemetryNestedFieldVirtualColumn) SetPathParts(v []TelemetryNestedFieldVirtualColumnAllOfPathParts)`

SetPathParts sets PathParts field to given value.

### HasPathParts

`func (o *TelemetryNestedFieldVirtualColumn) HasPathParts() bool`

HasPathParts returns a boolean if a field has been set.

### GetProcessFromRaw

`func (o *TelemetryNestedFieldVirtualColumn) GetProcessFromRaw() bool`

GetProcessFromRaw returns the ProcessFromRaw field if non-nil, zero value otherwise.

### GetProcessFromRawOk

`func (o *TelemetryNestedFieldVirtualColumn) GetProcessFromRawOk() (*bool, bool)`

GetProcessFromRawOk returns a tuple with the ProcessFromRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessFromRaw

`func (o *TelemetryNestedFieldVirtualColumn) SetProcessFromRaw(v bool)`

SetProcessFromRaw sets ProcessFromRaw field to given value.

### HasProcessFromRaw

`func (o *TelemetryNestedFieldVirtualColumn) HasProcessFromRaw() bool`

HasProcessFromRaw returns a boolean if a field has been set.

### GetPath

`func (o *TelemetryNestedFieldVirtualColumn) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *TelemetryNestedFieldVirtualColumn) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *TelemetryNestedFieldVirtualColumn) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *TelemetryNestedFieldVirtualColumn) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetUseJqSyntax

`func (o *TelemetryNestedFieldVirtualColumn) GetUseJqSyntax() bool`

GetUseJqSyntax returns the UseJqSyntax field if non-nil, zero value otherwise.

### GetUseJqSyntaxOk

`func (o *TelemetryNestedFieldVirtualColumn) GetUseJqSyntaxOk() (*bool, bool)`

GetUseJqSyntaxOk returns a tuple with the UseJqSyntax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseJqSyntax

`func (o *TelemetryNestedFieldVirtualColumn) SetUseJqSyntax(v bool)`

SetUseJqSyntax sets UseJqSyntax field to given value.

### HasUseJqSyntax

`func (o *TelemetryNestedFieldVirtualColumn) HasUseJqSyntax() bool`

HasUseJqSyntax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


