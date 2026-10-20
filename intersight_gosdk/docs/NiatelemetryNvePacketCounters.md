# NiatelemetryNvePacketCounters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.NvePacketCounters"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.NvePacketCounters"]
**McastInpkts** | Pointer to **int64** | Return mcast in packet count. | [optional] 
**McastOutbytes** | Pointer to **int64** | Return mcast outbytes count. | [optional] 
**UcastInpkts** | Pointer to **int64** | Return ucast in packet count. | [optional] 
**UcastOutpkts** | Pointer to **int64** | Return ucast out packet count. | [optional] 

## Methods

### NewNiatelemetryNvePacketCounters

`func NewNiatelemetryNvePacketCounters(classId string, objectType string, ) *NiatelemetryNvePacketCounters`

NewNiatelemetryNvePacketCounters instantiates a new NiatelemetryNvePacketCounters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryNvePacketCountersWithDefaults

`func NewNiatelemetryNvePacketCountersWithDefaults() *NiatelemetryNvePacketCounters`

NewNiatelemetryNvePacketCountersWithDefaults instantiates a new NiatelemetryNvePacketCounters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryNvePacketCounters) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryNvePacketCounters) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryNvePacketCounters) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryNvePacketCounters) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryNvePacketCounters) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryNvePacketCounters) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetMcastInpkts

`func (o *NiatelemetryNvePacketCounters) GetMcastInpkts() int64`

GetMcastInpkts returns the McastInpkts field if non-nil, zero value otherwise.

### GetMcastInpktsOk

`func (o *NiatelemetryNvePacketCounters) GetMcastInpktsOk() (*int64, bool)`

GetMcastInpktsOk returns a tuple with the McastInpkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcastInpkts

`func (o *NiatelemetryNvePacketCounters) SetMcastInpkts(v int64)`

SetMcastInpkts sets McastInpkts field to given value.

### HasMcastInpkts

`func (o *NiatelemetryNvePacketCounters) HasMcastInpkts() bool`

HasMcastInpkts returns a boolean if a field has been set.

### GetMcastOutbytes

`func (o *NiatelemetryNvePacketCounters) GetMcastOutbytes() int64`

GetMcastOutbytes returns the McastOutbytes field if non-nil, zero value otherwise.

### GetMcastOutbytesOk

`func (o *NiatelemetryNvePacketCounters) GetMcastOutbytesOk() (*int64, bool)`

GetMcastOutbytesOk returns a tuple with the McastOutbytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcastOutbytes

`func (o *NiatelemetryNvePacketCounters) SetMcastOutbytes(v int64)`

SetMcastOutbytes sets McastOutbytes field to given value.

### HasMcastOutbytes

`func (o *NiatelemetryNvePacketCounters) HasMcastOutbytes() bool`

HasMcastOutbytes returns a boolean if a field has been set.

### GetUcastInpkts

`func (o *NiatelemetryNvePacketCounters) GetUcastInpkts() int64`

GetUcastInpkts returns the UcastInpkts field if non-nil, zero value otherwise.

### GetUcastInpktsOk

`func (o *NiatelemetryNvePacketCounters) GetUcastInpktsOk() (*int64, bool)`

GetUcastInpktsOk returns a tuple with the UcastInpkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUcastInpkts

`func (o *NiatelemetryNvePacketCounters) SetUcastInpkts(v int64)`

SetUcastInpkts sets UcastInpkts field to given value.

### HasUcastInpkts

`func (o *NiatelemetryNvePacketCounters) HasUcastInpkts() bool`

HasUcastInpkts returns a boolean if a field has been set.

### GetUcastOutpkts

`func (o *NiatelemetryNvePacketCounters) GetUcastOutpkts() int64`

GetUcastOutpkts returns the UcastOutpkts field if non-nil, zero value otherwise.

### GetUcastOutpktsOk

`func (o *NiatelemetryNvePacketCounters) GetUcastOutpktsOk() (*int64, bool)`

GetUcastOutpktsOk returns a tuple with the UcastOutpkts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUcastOutpkts

`func (o *NiatelemetryNvePacketCounters) SetUcastOutpkts(v int64)`

SetUcastOutpkts sets UcastOutpkts field to given value.

### HasUcastOutpkts

`func (o *NiatelemetryNvePacketCounters) HasUcastOutpkts() bool`

HasUcastOutpkts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


