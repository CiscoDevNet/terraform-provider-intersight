# TelemetryDruidListFilteredVirtualColumn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The virtual-column type. | 
**Name** | Pointer to **string** | name of the virtual column. | [optional] 
**Delegate** | Pointer to **string** | Name of the multi-value string input column to filter. | [optional] 
**Values** | Pointer to **[]string** | Set of string value to allow or deny. | [optional] 
**IsAllowList** | Pointer to **bool** | If set to true, the virtual column will allow only the values in the list. If set to false, the virtual column will provide all values except those specified. | [optional] [default to true]

## Methods

### NewTelemetryDruidListFilteredVirtualColumn

`func NewTelemetryDruidListFilteredVirtualColumn(type_ string, ) *TelemetryDruidListFilteredVirtualColumn`

NewTelemetryDruidListFilteredVirtualColumn instantiates a new TelemetryDruidListFilteredVirtualColumn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetryDruidListFilteredVirtualColumnWithDefaults

`func NewTelemetryDruidListFilteredVirtualColumnWithDefaults() *TelemetryDruidListFilteredVirtualColumn`

NewTelemetryDruidListFilteredVirtualColumnWithDefaults instantiates a new TelemetryDruidListFilteredVirtualColumn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TelemetryDruidListFilteredVirtualColumn) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TelemetryDruidListFilteredVirtualColumn) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TelemetryDruidListFilteredVirtualColumn) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *TelemetryDruidListFilteredVirtualColumn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TelemetryDruidListFilteredVirtualColumn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TelemetryDruidListFilteredVirtualColumn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TelemetryDruidListFilteredVirtualColumn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDelegate

`func (o *TelemetryDruidListFilteredVirtualColumn) GetDelegate() string`

GetDelegate returns the Delegate field if non-nil, zero value otherwise.

### GetDelegateOk

`func (o *TelemetryDruidListFilteredVirtualColumn) GetDelegateOk() (*string, bool)`

GetDelegateOk returns a tuple with the Delegate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegate

`func (o *TelemetryDruidListFilteredVirtualColumn) SetDelegate(v string)`

SetDelegate sets Delegate field to given value.

### HasDelegate

`func (o *TelemetryDruidListFilteredVirtualColumn) HasDelegate() bool`

HasDelegate returns a boolean if a field has been set.

### GetValues

`func (o *TelemetryDruidListFilteredVirtualColumn) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *TelemetryDruidListFilteredVirtualColumn) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *TelemetryDruidListFilteredVirtualColumn) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *TelemetryDruidListFilteredVirtualColumn) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetIsAllowList

`func (o *TelemetryDruidListFilteredVirtualColumn) GetIsAllowList() bool`

GetIsAllowList returns the IsAllowList field if non-nil, zero value otherwise.

### GetIsAllowListOk

`func (o *TelemetryDruidListFilteredVirtualColumn) GetIsAllowListOk() (*bool, bool)`

GetIsAllowListOk returns a tuple with the IsAllowList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllowList

`func (o *TelemetryDruidListFilteredVirtualColumn) SetIsAllowList(v bool)`

SetIsAllowList sets IsAllowList field to given value.

### HasIsAllowList

`func (o *TelemetryDruidListFilteredVirtualColumn) HasIsAllowList() bool`

HasIsAllowList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


