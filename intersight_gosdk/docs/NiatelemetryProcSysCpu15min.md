# NiatelemetryProcSysCpu15min

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.ProcSysCpu15min"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.ProcSysCpu15min"]
**KernalAvg** | Pointer to **string** | Returns the kernal average value. | [optional] 
**UserAvg** | Pointer to **string** | Returns the user average value. | [optional] 

## Methods

### NewNiatelemetryProcSysCpu15min

`func NewNiatelemetryProcSysCpu15min(classId string, objectType string, ) *NiatelemetryProcSysCpu15min`

NewNiatelemetryProcSysCpu15min instantiates a new NiatelemetryProcSysCpu15min object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryProcSysCpu15minWithDefaults

`func NewNiatelemetryProcSysCpu15minWithDefaults() *NiatelemetryProcSysCpu15min`

NewNiatelemetryProcSysCpu15minWithDefaults instantiates a new NiatelemetryProcSysCpu15min object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryProcSysCpu15min) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryProcSysCpu15min) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryProcSysCpu15min) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryProcSysCpu15min) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryProcSysCpu15min) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryProcSysCpu15min) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetKernalAvg

`func (o *NiatelemetryProcSysCpu15min) GetKernalAvg() string`

GetKernalAvg returns the KernalAvg field if non-nil, zero value otherwise.

### GetKernalAvgOk

`func (o *NiatelemetryProcSysCpu15min) GetKernalAvgOk() (*string, bool)`

GetKernalAvgOk returns a tuple with the KernalAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernalAvg

`func (o *NiatelemetryProcSysCpu15min) SetKernalAvg(v string)`

SetKernalAvg sets KernalAvg field to given value.

### HasKernalAvg

`func (o *NiatelemetryProcSysCpu15min) HasKernalAvg() bool`

HasKernalAvg returns a boolean if a field has been set.

### GetUserAvg

`func (o *NiatelemetryProcSysCpu15min) GetUserAvg() string`

GetUserAvg returns the UserAvg field if non-nil, zero value otherwise.

### GetUserAvgOk

`func (o *NiatelemetryProcSysCpu15min) GetUserAvgOk() (*string, bool)`

GetUserAvgOk returns a tuple with the UserAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAvg

`func (o *NiatelemetryProcSysCpu15min) SetUserAvg(v string)`

SetUserAvg sets UserAvg field to given value.

### HasUserAvg

`func (o *NiatelemetryProcSysCpu15min) HasUserAvg() bool`

HasUserAvg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


