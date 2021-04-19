# EquipmentSlotAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "storage.DiskSlot"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "storage.DiskSlot"]
**SlotId** | Pointer to **string** | This field identifies a specific slot on the server. | [optional] [readonly] 

## Methods

### NewEquipmentSlotAllOf

`func NewEquipmentSlotAllOf(classId string, objectType string, ) *EquipmentSlotAllOf`

NewEquipmentSlotAllOf instantiates a new EquipmentSlotAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentSlotAllOfWithDefaults

`func NewEquipmentSlotAllOfWithDefaults() *EquipmentSlotAllOf`

NewEquipmentSlotAllOfWithDefaults instantiates a new EquipmentSlotAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *EquipmentSlotAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *EquipmentSlotAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *EquipmentSlotAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *EquipmentSlotAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *EquipmentSlotAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *EquipmentSlotAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSlotId

`func (o *EquipmentSlotAllOf) GetSlotId() string`

GetSlotId returns the SlotId field if non-nil, zero value otherwise.

### GetSlotIdOk

`func (o *EquipmentSlotAllOf) GetSlotIdOk() (*string, bool)`

GetSlotIdOk returns a tuple with the SlotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotId

`func (o *EquipmentSlotAllOf) SetSlotId(v string)`

SetSlotId sets SlotId field to given value.

### HasSlotId

`func (o *EquipmentSlotAllOf) HasSlotId() bool`

HasSlotId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


