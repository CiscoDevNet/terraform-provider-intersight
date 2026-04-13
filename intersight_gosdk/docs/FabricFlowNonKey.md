# FabricFlowNonKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.FlowNonKey"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.FlowNonKey"]
**ByteCounters** | Pointer to **bool** | Total number of bytes transferred for flow record. | [optional] 
**FirstSystemTime** | Pointer to **bool** | The timestamp when a flow was first observed or created in the system. | [optional] [default to false]
**LastSystemTime** | Pointer to **bool** | The last timestamp recorded for flow record by the system. | [optional] [default to false]
**PacketCounters** | Pointer to **bool** | Total number of packets transferred for flow record. | [optional] [default to false]

## Methods

### NewFabricFlowNonKey

`func NewFabricFlowNonKey(classId string, objectType string, ) *FabricFlowNonKey`

NewFabricFlowNonKey instantiates a new FabricFlowNonKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricFlowNonKeyWithDefaults

`func NewFabricFlowNonKeyWithDefaults() *FabricFlowNonKey`

NewFabricFlowNonKeyWithDefaults instantiates a new FabricFlowNonKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricFlowNonKey) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricFlowNonKey) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricFlowNonKey) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricFlowNonKey) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricFlowNonKey) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricFlowNonKey) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetByteCounters

`func (o *FabricFlowNonKey) GetByteCounters() bool`

GetByteCounters returns the ByteCounters field if non-nil, zero value otherwise.

### GetByteCountersOk

`func (o *FabricFlowNonKey) GetByteCountersOk() (*bool, bool)`

GetByteCountersOk returns a tuple with the ByteCounters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByteCounters

`func (o *FabricFlowNonKey) SetByteCounters(v bool)`

SetByteCounters sets ByteCounters field to given value.

### HasByteCounters

`func (o *FabricFlowNonKey) HasByteCounters() bool`

HasByteCounters returns a boolean if a field has been set.

### GetFirstSystemTime

`func (o *FabricFlowNonKey) GetFirstSystemTime() bool`

GetFirstSystemTime returns the FirstSystemTime field if non-nil, zero value otherwise.

### GetFirstSystemTimeOk

`func (o *FabricFlowNonKey) GetFirstSystemTimeOk() (*bool, bool)`

GetFirstSystemTimeOk returns a tuple with the FirstSystemTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSystemTime

`func (o *FabricFlowNonKey) SetFirstSystemTime(v bool)`

SetFirstSystemTime sets FirstSystemTime field to given value.

### HasFirstSystemTime

`func (o *FabricFlowNonKey) HasFirstSystemTime() bool`

HasFirstSystemTime returns a boolean if a field has been set.

### GetLastSystemTime

`func (o *FabricFlowNonKey) GetLastSystemTime() bool`

GetLastSystemTime returns the LastSystemTime field if non-nil, zero value otherwise.

### GetLastSystemTimeOk

`func (o *FabricFlowNonKey) GetLastSystemTimeOk() (*bool, bool)`

GetLastSystemTimeOk returns a tuple with the LastSystemTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSystemTime

`func (o *FabricFlowNonKey) SetLastSystemTime(v bool)`

SetLastSystemTime sets LastSystemTime field to given value.

### HasLastSystemTime

`func (o *FabricFlowNonKey) HasLastSystemTime() bool`

HasLastSystemTime returns a boolean if a field has been set.

### GetPacketCounters

`func (o *FabricFlowNonKey) GetPacketCounters() bool`

GetPacketCounters returns the PacketCounters field if non-nil, zero value otherwise.

### GetPacketCountersOk

`func (o *FabricFlowNonKey) GetPacketCountersOk() (*bool, bool)`

GetPacketCountersOk returns a tuple with the PacketCounters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketCounters

`func (o *FabricFlowNonKey) SetPacketCounters(v bool)`

SetPacketCounters sets PacketCounters field to given value.

### HasPacketCounters

`func (o *FabricFlowNonKey) HasPacketCounters() bool`

HasPacketCounters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


