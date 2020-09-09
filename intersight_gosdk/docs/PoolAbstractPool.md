# PoolAbstractPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assigned** | Pointer to **int64** | Number of IDs that are currently assigned. | [optional] [readonly] 
**AssignmentOrder** | Pointer to **string** | Assignment order decides the order in which the next identifier is allocated. * &#x60;sequential&#x60; - Identifiers are assigned in a sequential order. * &#x60;default&#x60; - Assignment order is decided by the system. | [optional] [default to "sequential"]
**Size** | Pointer to **int64** | Total number of identifiers in this pool. | [optional] [readonly] 

## Methods

### NewPoolAbstractPool

`func NewPoolAbstractPool() *PoolAbstractPool`

NewPoolAbstractPool instantiates a new PoolAbstractPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolAbstractPoolWithDefaults

`func NewPoolAbstractPoolWithDefaults() *PoolAbstractPool`

NewPoolAbstractPoolWithDefaults instantiates a new PoolAbstractPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssigned

`func (o *PoolAbstractPool) GetAssigned() int64`

GetAssigned returns the Assigned field if non-nil, zero value otherwise.

### GetAssignedOk

`func (o *PoolAbstractPool) GetAssignedOk() (*int64, bool)`

GetAssignedOk returns a tuple with the Assigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigned

`func (o *PoolAbstractPool) SetAssigned(v int64)`

SetAssigned sets Assigned field to given value.

### HasAssigned

`func (o *PoolAbstractPool) HasAssigned() bool`

HasAssigned returns a boolean if a field has been set.

### GetAssignmentOrder

`func (o *PoolAbstractPool) GetAssignmentOrder() string`

GetAssignmentOrder returns the AssignmentOrder field if non-nil, zero value otherwise.

### GetAssignmentOrderOk

`func (o *PoolAbstractPool) GetAssignmentOrderOk() (*string, bool)`

GetAssignmentOrderOk returns a tuple with the AssignmentOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentOrder

`func (o *PoolAbstractPool) SetAssignmentOrder(v string)`

SetAssignmentOrder sets AssignmentOrder field to given value.

### HasAssignmentOrder

`func (o *PoolAbstractPool) HasAssignmentOrder() bool`

HasAssignmentOrder returns a boolean if a field has been set.

### GetSize

`func (o *PoolAbstractPool) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PoolAbstractPool) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PoolAbstractPool) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *PoolAbstractPool) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


