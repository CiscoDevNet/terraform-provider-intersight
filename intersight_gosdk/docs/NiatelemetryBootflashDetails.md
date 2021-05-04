# NiatelemetryBootflashDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.BootflashDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.BootflashDetails"]
**FwRev** | Pointer to **string** | Return firmware revision in boot flash details. | [optional] 
**ModelType** | Pointer to **string** | Return model type in boot flash details. | [optional] 
**Serial** | Pointer to **string** | Return serial id in boot flash details. | [optional] 

## Methods

### NewNiatelemetryBootflashDetails

`func NewNiatelemetryBootflashDetails(classId string, objectType string, ) *NiatelemetryBootflashDetails`

NewNiatelemetryBootflashDetails instantiates a new NiatelemetryBootflashDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryBootflashDetailsWithDefaults

`func NewNiatelemetryBootflashDetailsWithDefaults() *NiatelemetryBootflashDetails`

NewNiatelemetryBootflashDetailsWithDefaults instantiates a new NiatelemetryBootflashDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryBootflashDetails) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryBootflashDetails) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryBootflashDetails) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryBootflashDetails) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryBootflashDetails) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryBootflashDetails) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFwRev

`func (o *NiatelemetryBootflashDetails) GetFwRev() string`

GetFwRev returns the FwRev field if non-nil, zero value otherwise.

### GetFwRevOk

`func (o *NiatelemetryBootflashDetails) GetFwRevOk() (*string, bool)`

GetFwRevOk returns a tuple with the FwRev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwRev

`func (o *NiatelemetryBootflashDetails) SetFwRev(v string)`

SetFwRev sets FwRev field to given value.

### HasFwRev

`func (o *NiatelemetryBootflashDetails) HasFwRev() bool`

HasFwRev returns a boolean if a field has been set.

### GetModelType

`func (o *NiatelemetryBootflashDetails) GetModelType() string`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *NiatelemetryBootflashDetails) GetModelTypeOk() (*string, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *NiatelemetryBootflashDetails) SetModelType(v string)`

SetModelType sets ModelType field to given value.

### HasModelType

`func (o *NiatelemetryBootflashDetails) HasModelType() bool`

HasModelType returns a boolean if a field has been set.

### GetSerial

`func (o *NiatelemetryBootflashDetails) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *NiatelemetryBootflashDetails) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *NiatelemetryBootflashDetails) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *NiatelemetryBootflashDetails) HasSerial() bool`

HasSerial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


