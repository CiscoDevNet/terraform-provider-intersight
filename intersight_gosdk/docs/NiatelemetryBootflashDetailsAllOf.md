# NiatelemetryBootflashDetailsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.BootflashDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.BootflashDetails"]
**FwRev** | Pointer to **string** | Return firmware revision in boot flash details. | [optional] 
**ModelType** | Pointer to **string** | Return model type in boot flash details. | [optional] 
**Serial** | Pointer to **string** | Return serial id in boot flash details. | [optional] 

## Methods

### NewNiatelemetryBootflashDetailsAllOf

`func NewNiatelemetryBootflashDetailsAllOf(classId string, objectType string, ) *NiatelemetryBootflashDetailsAllOf`

NewNiatelemetryBootflashDetailsAllOf instantiates a new NiatelemetryBootflashDetailsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryBootflashDetailsAllOfWithDefaults

`func NewNiatelemetryBootflashDetailsAllOfWithDefaults() *NiatelemetryBootflashDetailsAllOf`

NewNiatelemetryBootflashDetailsAllOfWithDefaults instantiates a new NiatelemetryBootflashDetailsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryBootflashDetailsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryBootflashDetailsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryBootflashDetailsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryBootflashDetailsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryBootflashDetailsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryBootflashDetailsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFwRev

`func (o *NiatelemetryBootflashDetailsAllOf) GetFwRev() string`

GetFwRev returns the FwRev field if non-nil, zero value otherwise.

### GetFwRevOk

`func (o *NiatelemetryBootflashDetailsAllOf) GetFwRevOk() (*string, bool)`

GetFwRevOk returns a tuple with the FwRev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwRev

`func (o *NiatelemetryBootflashDetailsAllOf) SetFwRev(v string)`

SetFwRev sets FwRev field to given value.

### HasFwRev

`func (o *NiatelemetryBootflashDetailsAllOf) HasFwRev() bool`

HasFwRev returns a boolean if a field has been set.

### GetModelType

`func (o *NiatelemetryBootflashDetailsAllOf) GetModelType() string`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *NiatelemetryBootflashDetailsAllOf) GetModelTypeOk() (*string, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *NiatelemetryBootflashDetailsAllOf) SetModelType(v string)`

SetModelType sets ModelType field to given value.

### HasModelType

`func (o *NiatelemetryBootflashDetailsAllOf) HasModelType() bool`

HasModelType returns a boolean if a field has been set.

### GetSerial

`func (o *NiatelemetryBootflashDetailsAllOf) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *NiatelemetryBootflashDetailsAllOf) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *NiatelemetryBootflashDetailsAllOf) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *NiatelemetryBootflashDetailsAllOf) HasSerial() bool`

HasSerial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


