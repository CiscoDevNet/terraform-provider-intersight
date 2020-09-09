# VnicScsiQueueSettingsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** | The number of SCSI I/O queue resources the system should allocate. | [optional] 
**RingSize** | Pointer to **int64** | The number of descriptors in each SCSI I/O queue. | [optional] 

## Methods

### NewVnicScsiQueueSettingsAllOf

`func NewVnicScsiQueueSettingsAllOf() *VnicScsiQueueSettingsAllOf`

NewVnicScsiQueueSettingsAllOf instantiates a new VnicScsiQueueSettingsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicScsiQueueSettingsAllOfWithDefaults

`func NewVnicScsiQueueSettingsAllOfWithDefaults() *VnicScsiQueueSettingsAllOf`

NewVnicScsiQueueSettingsAllOfWithDefaults instantiates a new VnicScsiQueueSettingsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *VnicScsiQueueSettingsAllOf) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *VnicScsiQueueSettingsAllOf) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *VnicScsiQueueSettingsAllOf) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *VnicScsiQueueSettingsAllOf) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetRingSize

`func (o *VnicScsiQueueSettingsAllOf) GetRingSize() int64`

GetRingSize returns the RingSize field if non-nil, zero value otherwise.

### GetRingSizeOk

`func (o *VnicScsiQueueSettingsAllOf) GetRingSizeOk() (*int64, bool)`

GetRingSizeOk returns a tuple with the RingSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRingSize

`func (o *VnicScsiQueueSettingsAllOf) SetRingSize(v int64)`

SetRingSize sets RingSize field to given value.

### HasRingSize

`func (o *VnicScsiQueueSettingsAllOf) HasRingSize() bool`

HasRingSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


