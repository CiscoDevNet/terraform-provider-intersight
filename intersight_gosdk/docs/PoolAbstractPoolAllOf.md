# PoolAbstractPoolAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Assigned** | Pointer to **int64** | Number of IDs that are currently assigned. | [optional] [readonly] 
**AssignmentOrder** | Pointer to **string** | Assignment order decides the order in which the next identifier is allocated. * &#x60;sequential&#x60; - Identifiers are assigned in a sequential order. * &#x60;default&#x60; - Assignment order is decided by the system. | [optional] [default to "sequential"]
**Size** | Pointer to **int64** | Total number of identifiers in this pool. | [optional] [readonly] 

## Methods

### NewPoolAbstractPoolAllOf

`func NewPoolAbstractPoolAllOf(classId string, objectType string, ) *PoolAbstractPoolAllOf`

NewPoolAbstractPoolAllOf instantiates a new PoolAbstractPoolAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolAbstractPoolAllOfWithDefaults

`func NewPoolAbstractPoolAllOfWithDefaults() *PoolAbstractPoolAllOf`

NewPoolAbstractPoolAllOfWithDefaults instantiates a new PoolAbstractPoolAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *PoolAbstractPoolAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *PoolAbstractPoolAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *PoolAbstractPoolAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *PoolAbstractPoolAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PoolAbstractPoolAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PoolAbstractPoolAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAssigned

`func (o *PoolAbstractPoolAllOf) GetAssigned() int64`

GetAssigned returns the Assigned field if non-nil, zero value otherwise.

### GetAssignedOk

`func (o *PoolAbstractPoolAllOf) GetAssignedOk() (*int64, bool)`

GetAssignedOk returns a tuple with the Assigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigned

`func (o *PoolAbstractPoolAllOf) SetAssigned(v int64)`

SetAssigned sets Assigned field to given value.

### HasAssigned

`func (o *PoolAbstractPoolAllOf) HasAssigned() bool`

HasAssigned returns a boolean if a field has been set.

### GetAssignmentOrder

`func (o *PoolAbstractPoolAllOf) GetAssignmentOrder() string`

GetAssignmentOrder returns the AssignmentOrder field if non-nil, zero value otherwise.

### GetAssignmentOrderOk

`func (o *PoolAbstractPoolAllOf) GetAssignmentOrderOk() (*string, bool)`

GetAssignmentOrderOk returns a tuple with the AssignmentOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentOrder

`func (o *PoolAbstractPoolAllOf) SetAssignmentOrder(v string)`

SetAssignmentOrder sets AssignmentOrder field to given value.

### HasAssignmentOrder

`func (o *PoolAbstractPoolAllOf) HasAssignmentOrder() bool`

HasAssignmentOrder returns a boolean if a field has been set.

### GetSize

`func (o *PoolAbstractPoolAllOf) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PoolAbstractPoolAllOf) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PoolAbstractPoolAllOf) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *PoolAbstractPoolAllOf) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


