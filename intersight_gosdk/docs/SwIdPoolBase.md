# SwIdPoolBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**GapAvailableIds** | Pointer to **[]int64** |  | [optional] 
**NextAvailableId** | Pointer to **int64** | Shows the next available Chassis ID to be allocated. | [optional] [readonly] 

## Methods

### NewSwIdPoolBase

`func NewSwIdPoolBase(classId string, objectType string, ) *SwIdPoolBase`

NewSwIdPoolBase instantiates a new SwIdPoolBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwIdPoolBaseWithDefaults

`func NewSwIdPoolBaseWithDefaults() *SwIdPoolBase`

NewSwIdPoolBaseWithDefaults instantiates a new SwIdPoolBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SwIdPoolBase) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SwIdPoolBase) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SwIdPoolBase) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SwIdPoolBase) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SwIdPoolBase) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SwIdPoolBase) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetGapAvailableIds

`func (o *SwIdPoolBase) GetGapAvailableIds() []int64`

GetGapAvailableIds returns the GapAvailableIds field if non-nil, zero value otherwise.

### GetGapAvailableIdsOk

`func (o *SwIdPoolBase) GetGapAvailableIdsOk() (*[]int64, bool)`

GetGapAvailableIdsOk returns a tuple with the GapAvailableIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGapAvailableIds

`func (o *SwIdPoolBase) SetGapAvailableIds(v []int64)`

SetGapAvailableIds sets GapAvailableIds field to given value.

### HasGapAvailableIds

`func (o *SwIdPoolBase) HasGapAvailableIds() bool`

HasGapAvailableIds returns a boolean if a field has been set.

### SetGapAvailableIdsNil

`func (o *SwIdPoolBase) SetGapAvailableIdsNil(b bool)`

 SetGapAvailableIdsNil sets the value for GapAvailableIds to be an explicit nil

### UnsetGapAvailableIds
`func (o *SwIdPoolBase) UnsetGapAvailableIds()`

UnsetGapAvailableIds ensures that no value is present for GapAvailableIds, not even an explicit nil
### GetNextAvailableId

`func (o *SwIdPoolBase) GetNextAvailableId() int64`

GetNextAvailableId returns the NextAvailableId field if non-nil, zero value otherwise.

### GetNextAvailableIdOk

`func (o *SwIdPoolBase) GetNextAvailableIdOk() (*int64, bool)`

GetNextAvailableIdOk returns a tuple with the NextAvailableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAvailableId

`func (o *SwIdPoolBase) SetNextAvailableId(v int64)`

SetNextAvailableId sets NextAvailableId field to given value.

### HasNextAvailableId

`func (o *SwIdPoolBase) HasNextAvailableId() bool`

HasNextAvailableId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


