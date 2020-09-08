# TamBaseDataSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is used to unique identify and refer a given data source in an alert definition. | [optional] 
**Type** | Pointer to **string** | Type of data source (for e.g. TextFsmTempalate based, Intersight API based etc.). * &#x60;nxos&#x60; - Collector type for this data collection is NXOS. * &#x60;intersightApi&#x60; - Collector type for this data collection is Intersight APIs. | [optional] [default to "nxos"]

## Methods

### NewTamBaseDataSource

`func NewTamBaseDataSource() *TamBaseDataSource`

NewTamBaseDataSource instantiates a new TamBaseDataSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTamBaseDataSourceWithDefaults

`func NewTamBaseDataSourceWithDefaults() *TamBaseDataSource`

NewTamBaseDataSourceWithDefaults instantiates a new TamBaseDataSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TamBaseDataSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TamBaseDataSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TamBaseDataSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TamBaseDataSource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *TamBaseDataSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TamBaseDataSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TamBaseDataSource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TamBaseDataSource) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


