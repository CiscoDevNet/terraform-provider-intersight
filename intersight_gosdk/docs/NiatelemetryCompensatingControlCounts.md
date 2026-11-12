# NiatelemetryCompensatingControlCounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.CompensatingControlCounts"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.CompensatingControlCounts"]
**Critical** | Pointer to **int64** | Count of critical severity compensating controls. | [optional] [readonly] 
**Major** | Pointer to **int64** | Count of major severity compensating controls. | [optional] [readonly] 
**Minor** | Pointer to **int64** | Count of minor severity compensating controls. | [optional] [readonly] 
**Warning** | Pointer to **int64** | Count of warning severity compensating controls. | [optional] [readonly] 

## Methods

### NewNiatelemetryCompensatingControlCounts

`func NewNiatelemetryCompensatingControlCounts(classId string, objectType string, ) *NiatelemetryCompensatingControlCounts`

NewNiatelemetryCompensatingControlCounts instantiates a new NiatelemetryCompensatingControlCounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryCompensatingControlCountsWithDefaults

`func NewNiatelemetryCompensatingControlCountsWithDefaults() *NiatelemetryCompensatingControlCounts`

NewNiatelemetryCompensatingControlCountsWithDefaults instantiates a new NiatelemetryCompensatingControlCounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryCompensatingControlCounts) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryCompensatingControlCounts) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryCompensatingControlCounts) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryCompensatingControlCounts) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryCompensatingControlCounts) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryCompensatingControlCounts) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCritical

`func (o *NiatelemetryCompensatingControlCounts) GetCritical() int64`

GetCritical returns the Critical field if non-nil, zero value otherwise.

### GetCriticalOk

`func (o *NiatelemetryCompensatingControlCounts) GetCriticalOk() (*int64, bool)`

GetCriticalOk returns a tuple with the Critical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCritical

`func (o *NiatelemetryCompensatingControlCounts) SetCritical(v int64)`

SetCritical sets Critical field to given value.

### HasCritical

`func (o *NiatelemetryCompensatingControlCounts) HasCritical() bool`

HasCritical returns a boolean if a field has been set.

### GetMajor

`func (o *NiatelemetryCompensatingControlCounts) GetMajor() int64`

GetMajor returns the Major field if non-nil, zero value otherwise.

### GetMajorOk

`func (o *NiatelemetryCompensatingControlCounts) GetMajorOk() (*int64, bool)`

GetMajorOk returns a tuple with the Major field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajor

`func (o *NiatelemetryCompensatingControlCounts) SetMajor(v int64)`

SetMajor sets Major field to given value.

### HasMajor

`func (o *NiatelemetryCompensatingControlCounts) HasMajor() bool`

HasMajor returns a boolean if a field has been set.

### GetMinor

`func (o *NiatelemetryCompensatingControlCounts) GetMinor() int64`

GetMinor returns the Minor field if non-nil, zero value otherwise.

### GetMinorOk

`func (o *NiatelemetryCompensatingControlCounts) GetMinorOk() (*int64, bool)`

GetMinorOk returns a tuple with the Minor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinor

`func (o *NiatelemetryCompensatingControlCounts) SetMinor(v int64)`

SetMinor sets Minor field to given value.

### HasMinor

`func (o *NiatelemetryCompensatingControlCounts) HasMinor() bool`

HasMinor returns a boolean if a field has been set.

### GetWarning

`func (o *NiatelemetryCompensatingControlCounts) GetWarning() int64`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *NiatelemetryCompensatingControlCounts) GetWarningOk() (*int64, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *NiatelemetryCompensatingControlCounts) SetWarning(v int64)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *NiatelemetryCompensatingControlCounts) HasWarning() bool`

HasWarning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


