# HciSbdf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.Sbdf"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.Sbdf"]
**Bus** | Pointer to **int32** | PCI Bus number information. | [optional] [readonly] 
**Device** | Pointer to **int32** | PCI Device number information. | [optional] [readonly] 
**Func** | Pointer to **int32** | PCI Function number information. | [optional] [readonly] 
**Segment** | Pointer to **int32** | PCI&#39;s Segment information. | [optional] [readonly] 

## Methods

### NewHciSbdf

`func NewHciSbdf(classId string, objectType string, ) *HciSbdf`

NewHciSbdf instantiates a new HciSbdf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciSbdfWithDefaults

`func NewHciSbdfWithDefaults() *HciSbdf`

NewHciSbdfWithDefaults instantiates a new HciSbdf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciSbdf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciSbdf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciSbdf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciSbdf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciSbdf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciSbdf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBus

`func (o *HciSbdf) GetBus() int32`

GetBus returns the Bus field if non-nil, zero value otherwise.

### GetBusOk

`func (o *HciSbdf) GetBusOk() (*int32, bool)`

GetBusOk returns a tuple with the Bus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBus

`func (o *HciSbdf) SetBus(v int32)`

SetBus sets Bus field to given value.

### HasBus

`func (o *HciSbdf) HasBus() bool`

HasBus returns a boolean if a field has been set.

### GetDevice

`func (o *HciSbdf) GetDevice() int32`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *HciSbdf) GetDeviceOk() (*int32, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *HciSbdf) SetDevice(v int32)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *HciSbdf) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetFunc

`func (o *HciSbdf) GetFunc() int32`

GetFunc returns the Func field if non-nil, zero value otherwise.

### GetFuncOk

`func (o *HciSbdf) GetFuncOk() (*int32, bool)`

GetFuncOk returns a tuple with the Func field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunc

`func (o *HciSbdf) SetFunc(v int32)`

SetFunc sets Func field to given value.

### HasFunc

`func (o *HciSbdf) HasFunc() bool`

HasFunc returns a boolean if a field has been set.

### GetSegment

`func (o *HciSbdf) GetSegment() int32`

GetSegment returns the Segment field if non-nil, zero value otherwise.

### GetSegmentOk

`func (o *HciSbdf) GetSegmentOk() (*int32, bool)`

GetSegmentOk returns a tuple with the Segment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment

`func (o *HciSbdf) SetSegment(v int32)`

SetSegment sets Segment field to given value.

### HasSegment

`func (o *HciSbdf) HasSegment() bool`

HasSegment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


