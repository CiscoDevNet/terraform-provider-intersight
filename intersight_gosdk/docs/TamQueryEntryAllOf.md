# TamQueryEntryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "tam.QueryEntry"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "tam.QueryEntry"]
**Name** | Pointer to **string** | Name is used to unique identify and result of the given query which can be used by subsequent queries as input data source. | [optional] 
**Priority** | Pointer to **int64** | An integer value depicting the priority of the query among the queries that are part of the same QueryEntry collection. | [optional] 
**Query** | Pointer to **string** | A SparkSQL query to be used on a given data source. | [optional] 

## Methods

### NewTamQueryEntryAllOf

`func NewTamQueryEntryAllOf(classId string, objectType string, ) *TamQueryEntryAllOf`

NewTamQueryEntryAllOf instantiates a new TamQueryEntryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTamQueryEntryAllOfWithDefaults

`func NewTamQueryEntryAllOfWithDefaults() *TamQueryEntryAllOf`

NewTamQueryEntryAllOfWithDefaults instantiates a new TamQueryEntryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *TamQueryEntryAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *TamQueryEntryAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *TamQueryEntryAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *TamQueryEntryAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *TamQueryEntryAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *TamQueryEntryAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *TamQueryEntryAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TamQueryEntryAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TamQueryEntryAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TamQueryEntryAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriority

`func (o *TamQueryEntryAllOf) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TamQueryEntryAllOf) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TamQueryEntryAllOf) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *TamQueryEntryAllOf) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetQuery

`func (o *TamQueryEntryAllOf) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *TamQueryEntryAllOf) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *TamQueryEntryAllOf) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *TamQueryEntryAllOf) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


