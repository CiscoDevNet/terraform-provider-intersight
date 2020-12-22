# TelemetryDruidInlineDataSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of data source. | 
**ColumnNames** | **[]string** | the column names. | 
**Rows** | **[][]string** | an array of rows. | 

## Methods

### NewTelemetryDruidInlineDataSource

`func NewTelemetryDruidInlineDataSource(type_ string, columnNames []string, rows [][]string, ) *TelemetryDruidInlineDataSource`

NewTelemetryDruidInlineDataSource instantiates a new TelemetryDruidInlineDataSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidInlineDataSourceWithDefaults

`func NewTelemetryDruidInlineDataSourceWithDefaults() *TelemetryDruidInlineDataSource`

NewTelemetryDruidInlineDataSourceWithDefaults instantiates a new TelemetryDruidInlineDataSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidInlineDataSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidInlineDataSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidInlineDataSource) SetType(v string)`

SetType sets Type field to given value.


### GetColumnNames

`func (o *TelemetryDruidInlineDataSource) GetColumnNames() []string`

GetColumnNames returns the ColumnNames field if non-nil, zero value otherwise.

### GetColumnNamesOk

`func (o *TelemetryDruidInlineDataSource) GetColumnNamesOk() (*[]string, bool)`

GetColumnNamesOk returns a tuple with the ColumnNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnNames

`func (o *TelemetryDruidInlineDataSource) SetColumnNames(v []string)`

SetColumnNames sets ColumnNames field to given value.


### GetRows

`func (o *TelemetryDruidInlineDataSource) GetRows() [][]string`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *TelemetryDruidInlineDataSource) GetRowsOk() (*[][]string, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *TelemetryDruidInlineDataSource) SetRows(v [][]string)`

SetRows sets Rows field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


