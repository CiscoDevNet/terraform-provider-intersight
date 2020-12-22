# VnicScsiQueueSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.ScsiQueueSettings"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.ScsiQueueSettings"]
**Count** | Pointer to **int64** | The number of SCSI I/O queue resources the system should allocate. | [optional] [default to 1]
**RingSize** | Pointer to **int64** | The number of descriptors in each SCSI I/O queue. | [optional] [default to 512]

## Methods

### NewVnicScsiQueueSettings

`func NewVnicScsiQueueSettings(classId string, objectType string, ) *VnicScsiQueueSettings`

NewVnicScsiQueueSettings instantiates a new VnicScsiQueueSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicScsiQueueSettingsWithDefaults

`func NewVnicScsiQueueSettingsWithDefaults() *VnicScsiQueueSettings`

NewVnicScsiQueueSettingsWithDefaults instantiates a new VnicScsiQueueSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicScsiQueueSettings) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicScsiQueueSettings) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicScsiQueueSettings) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicScsiQueueSettings) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicScsiQueueSettings) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicScsiQueueSettings) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCount

`func (o *VnicScsiQueueSettings) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *VnicScsiQueueSettings) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *VnicScsiQueueSettings) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *VnicScsiQueueSettings) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetRingSize

`func (o *VnicScsiQueueSettings) GetRingSize() int64`

GetRingSize returns the RingSize field if non-nil, zero value otherwise.

### GetRingSizeOk

`func (o *VnicScsiQueueSettings) GetRingSizeOk() (*int64, bool)`

GetRingSizeOk returns a tuple with the RingSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRingSize

`func (o *VnicScsiQueueSettings) SetRingSize(v int64)`

SetRingSize sets RingSize field to given value.

### HasRingSize

`func (o *VnicScsiQueueSettings) HasRingSize() bool`

HasRingSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


