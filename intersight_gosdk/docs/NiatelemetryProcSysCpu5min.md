# NiatelemetryProcSysCpu5min

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.ProcSysCpu5min"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.ProcSysCpu5min"]
**KernalAvg** | Pointer to **string** | Returns the kernal average value. | [optional] 
**UserAvg** | Pointer to **string** | Returns the user average value. | [optional] 

## Methods

### NewNiatelemetryProcSysCpu5min

`func NewNiatelemetryProcSysCpu5min(classId string, objectType string, ) *NiatelemetryProcSysCpu5min`

NewNiatelemetryProcSysCpu5min instantiates a new NiatelemetryProcSysCpu5min object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryProcSysCpu5minWithDefaults

`func NewNiatelemetryProcSysCpu5minWithDefaults() *NiatelemetryProcSysCpu5min`

NewNiatelemetryProcSysCpu5minWithDefaults instantiates a new NiatelemetryProcSysCpu5min object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryProcSysCpu5min) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryProcSysCpu5min) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryProcSysCpu5min) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryProcSysCpu5min) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryProcSysCpu5min) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryProcSysCpu5min) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetKernalAvg

`func (o *NiatelemetryProcSysCpu5min) GetKernalAvg() string`

GetKernalAvg returns the KernalAvg field if non-nil, zero value otherwise.

### GetKernalAvgOk

`func (o *NiatelemetryProcSysCpu5min) GetKernalAvgOk() (*string, bool)`

GetKernalAvgOk returns a tuple with the KernalAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernalAvg

`func (o *NiatelemetryProcSysCpu5min) SetKernalAvg(v string)`

SetKernalAvg sets KernalAvg field to given value.

### HasKernalAvg

`func (o *NiatelemetryProcSysCpu5min) HasKernalAvg() bool`

HasKernalAvg returns a boolean if a field has been set.

### GetUserAvg

`func (o *NiatelemetryProcSysCpu5min) GetUserAvg() string`

GetUserAvg returns the UserAvg field if non-nil, zero value otherwise.

### GetUserAvgOk

`func (o *NiatelemetryProcSysCpu5min) GetUserAvgOk() (*string, bool)`

GetUserAvgOk returns a tuple with the UserAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAvg

`func (o *NiatelemetryProcSysCpu5min) SetUserAvg(v string)`

SetUserAvg sets UserAvg field to given value.

### HasUserAvg

`func (o *NiatelemetryProcSysCpu5min) HasUserAvg() bool`

HasUserAvg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


