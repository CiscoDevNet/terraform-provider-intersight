# StorageSpanDrives

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.SpanDrives"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.SpanDrives"]
**Slots** | Pointer to **string** | Collection of local disks that are part of this span group. Allowed value is a comma or hyphen separated number range. The minimum number of disks needed in a span group varies based on RAID level. RAID0 requires at least one disk, RAID1 and RAID10 requires at least 2 and in multiples of 2, RAID5 RAID50 RAID6 and RAID60 require at least 3 disks in a span group. | [optional] 

## Methods

### NewStorageSpanDrives

`func NewStorageSpanDrives(classId string, objectType string, ) *StorageSpanDrives`

NewStorageSpanDrives instantiates a new StorageSpanDrives object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageSpanDrivesWithDefaults

`func NewStorageSpanDrivesWithDefaults() *StorageSpanDrives`

NewStorageSpanDrivesWithDefaults instantiates a new StorageSpanDrives object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageSpanDrives) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageSpanDrives) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageSpanDrives) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageSpanDrives) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageSpanDrives) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageSpanDrives) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSlots

`func (o *StorageSpanDrives) GetSlots() string`

GetSlots returns the Slots field if non-nil, zero value otherwise.

### GetSlotsOk

`func (o *StorageSpanDrives) GetSlotsOk() (*string, bool)`

GetSlotsOk returns a tuple with the Slots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlots

`func (o *StorageSpanDrives) SetSlots(v string)`

SetSlots sets Slots field to given value.

### HasSlots

`func (o *StorageSpanDrives) HasSlots() bool`

HasSlots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


