# ResourceChassisServersQualifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "resource.ChassisServersQualifier"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "resource.ChassisServersQualifier"]
**SlotIds** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewResourceChassisServersQualifier

`func NewResourceChassisServersQualifier(classId string, objectType string, ) *ResourceChassisServersQualifier`

NewResourceChassisServersQualifier instantiates a new ResourceChassisServersQualifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceChassisServersQualifierWithDefaults

`func NewResourceChassisServersQualifierWithDefaults() *ResourceChassisServersQualifier`

NewResourceChassisServersQualifierWithDefaults instantiates a new ResourceChassisServersQualifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ResourceChassisServersQualifier) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ResourceChassisServersQualifier) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ResourceChassisServersQualifier) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ResourceChassisServersQualifier) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ResourceChassisServersQualifier) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ResourceChassisServersQualifier) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSlotIds

`func (o *ResourceChassisServersQualifier) GetSlotIds() []int64`

GetSlotIds returns the SlotIds field if non-nil, zero value otherwise.

### GetSlotIdsOk

`func (o *ResourceChassisServersQualifier) GetSlotIdsOk() (*[]int64, bool)`

GetSlotIdsOk returns a tuple with the SlotIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotIds

`func (o *ResourceChassisServersQualifier) SetSlotIds(v []int64)`

SetSlotIds sets SlotIds field to given value.

### HasSlotIds

`func (o *ResourceChassisServersQualifier) HasSlotIds() bool`

HasSlotIds returns a boolean if a field has been set.

### SetSlotIdsNil

`func (o *ResourceChassisServersQualifier) SetSlotIdsNil(b bool)`

 SetSlotIdsNil sets the value for SlotIds to be an explicit nil

### UnsetSlotIds
`func (o *ResourceChassisServersQualifier) UnsetSlotIds()`

UnsetSlotIds ensures that no value is present for SlotIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


