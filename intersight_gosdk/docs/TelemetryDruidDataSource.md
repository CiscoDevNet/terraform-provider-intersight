# TelemetryDruidDataSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of data source. | 
**Name** | **string** | The name of a data source. | 
**DataSources** | **[]string** | A list of data sources. | 
**Query** | [**TelemetryDruidGroupByRequest**](TelemetryDruidGroupByRequest.md) |  | 
**Lookup** | **string** | the name of the lookup object. | 
**Left** | **string** | Left-hand datasource. Must be of type table, join, lookup, query, or inline. Placing another join as the left datasource allows you to join arbitrarily many datasources. | 
**Right** | **string** | Right-hand datasource. Must be of type lookup, query, or inline. | 
**RightPrefix** | **string** | String prefix that will be applied to all columns from the right-hand datasource, to prevent them from colliding with columns from the left-hand datasource. Can be any string, so long as it is nonempty and is not be a prefix of the string __time. Any columns from the left-hand side that start with your rightPrefix will be shadowed. It is up to you to provide a prefix that will not shadow any important columns from the left side. | 
**Condition** | **string** | Expression that must be an equality where one side is an expression of the left-hand side, and the other side is a simple column reference to the right-hand side. The right-hand reference must be a simple column reference. | 
**JoinType** | **string** |  | 
**ColumnNames** | **[]string** | the column names. | 
**Rows** | **[][]string** | an array of rows. | 

## Methods

### NewTelemetryDruidDataSource

`func NewTelemetryDruidDataSource(type_ string, name string, dataSources []string, query TelemetryDruidGroupByRequest, lookup string, left string, right string, rightPrefix string, condition string, joinType string, columnNames []string, rows [][]string, ) *TelemetryDruidDataSource`

NewTelemetryDruidDataSource instantiates a new TelemetryDruidDataSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidDataSourceWithDefaults

`func NewTelemetryDruidDataSourceWithDefaults() *TelemetryDruidDataSource`

NewTelemetryDruidDataSourceWithDefaults instantiates a new TelemetryDruidDataSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidDataSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidDataSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidDataSource) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *TelemetryDruidDataSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TelemetryDruidDataSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TelemetryDruidDataSource) SetName(v string)`

SetName sets Name field to given value.


### GetDataSources

`func (o *TelemetryDruidDataSource) GetDataSources() []string`

GetDataSources returns the DataSources field if non-nil, zero value otherwise.

### GetDataSourcesOk

`func (o *TelemetryDruidDataSource) GetDataSourcesOk() (*[]string, bool)`

GetDataSourcesOk returns a tuple with the DataSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSources

`func (o *TelemetryDruidDataSource) SetDataSources(v []string)`

SetDataSources sets DataSources field to given value.


### GetQuery

`func (o *TelemetryDruidDataSource) GetQuery() TelemetryDruidGroupByRequest`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *TelemetryDruidDataSource) GetQueryOk() (*TelemetryDruidGroupByRequest, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *TelemetryDruidDataSource) SetQuery(v TelemetryDruidGroupByRequest)`

SetQuery sets Query field to given value.


### GetLookup

`func (o *TelemetryDruidDataSource) GetLookup() string`

GetLookup returns the Lookup field if non-nil, zero value otherwise.

### GetLookupOk

`func (o *TelemetryDruidDataSource) GetLookupOk() (*string, bool)`

GetLookupOk returns a tuple with the Lookup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookup

`func (o *TelemetryDruidDataSource) SetLookup(v string)`

SetLookup sets Lookup field to given value.


### GetLeft

`func (o *TelemetryDruidDataSource) GetLeft() string`

GetLeft returns the Left field if non-nil, zero value otherwise.

### GetLeftOk

`func (o *TelemetryDruidDataSource) GetLeftOk() (*string, bool)`

GetLeftOk returns a tuple with the Left field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeft

`func (o *TelemetryDruidDataSource) SetLeft(v string)`

SetLeft sets Left field to given value.


### GetRight

`func (o *TelemetryDruidDataSource) GetRight() string`

GetRight returns the Right field if non-nil, zero value otherwise.

### GetRightOk

`func (o *TelemetryDruidDataSource) GetRightOk() (*string, bool)`

GetRightOk returns a tuple with the Right field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRight

`func (o *TelemetryDruidDataSource) SetRight(v string)`

SetRight sets Right field to given value.


### GetRightPrefix

`func (o *TelemetryDruidDataSource) GetRightPrefix() string`

GetRightPrefix returns the RightPrefix field if non-nil, zero value otherwise.

### GetRightPrefixOk

`func (o *TelemetryDruidDataSource) GetRightPrefixOk() (*string, bool)`

GetRightPrefixOk returns a tuple with the RightPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRightPrefix

`func (o *TelemetryDruidDataSource) SetRightPrefix(v string)`

SetRightPrefix sets RightPrefix field to given value.


### GetCondition

`func (o *TelemetryDruidDataSource) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *TelemetryDruidDataSource) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *TelemetryDruidDataSource) SetCondition(v string)`

SetCondition sets Condition field to given value.


### GetJoinType

`func (o *TelemetryDruidDataSource) GetJoinType() string`

GetJoinType returns the JoinType field if non-nil, zero value otherwise.

### GetJoinTypeOk

`func (o *TelemetryDruidDataSource) GetJoinTypeOk() (*string, bool)`

GetJoinTypeOk returns a tuple with the JoinType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinType

`func (o *TelemetryDruidDataSource) SetJoinType(v string)`

SetJoinType sets JoinType field to given value.


### GetColumnNames

`func (o *TelemetryDruidDataSource) GetColumnNames() []string`

GetColumnNames returns the ColumnNames field if non-nil, zero value otherwise.

### GetColumnNamesOk

`func (o *TelemetryDruidDataSource) GetColumnNamesOk() (*[]string, bool)`

GetColumnNamesOk returns a tuple with the ColumnNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnNames

`func (o *TelemetryDruidDataSource) SetColumnNames(v []string)`

SetColumnNames sets ColumnNames field to given value.


### GetRows

`func (o *TelemetryDruidDataSource) GetRows() [][]string`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *TelemetryDruidDataSource) GetRowsOk() (*[][]string, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *TelemetryDruidDataSource) SetRows(v [][]string)`

SetRows sets Rows field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


