# NiatelemetryCompensatingControlDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.CompensatingControlDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.CompensatingControlDetails"]
**AvailableControlsCount** | Pointer to [**NullableNiatelemetryCompensatingControlCounts**](NiatelemetryCompensatingControlCounts.md) |  | [optional] 
**MonitoringControlsCount** | Pointer to [**NullableNiatelemetryCompensatingControlCounts**](NiatelemetryCompensatingControlCounts.md) |  | [optional] 
**ProtectedControlsCount** | Pointer to [**NullableNiatelemetryCompensatingControlCounts**](NiatelemetryCompensatingControlCounts.md) |  | [optional] 

## Methods

### NewNiatelemetryCompensatingControlDetails

`func NewNiatelemetryCompensatingControlDetails(classId string, objectType string, ) *NiatelemetryCompensatingControlDetails`

NewNiatelemetryCompensatingControlDetails instantiates a new NiatelemetryCompensatingControlDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryCompensatingControlDetailsWithDefaults

`func NewNiatelemetryCompensatingControlDetailsWithDefaults() *NiatelemetryCompensatingControlDetails`

NewNiatelemetryCompensatingControlDetailsWithDefaults instantiates a new NiatelemetryCompensatingControlDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryCompensatingControlDetails) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryCompensatingControlDetails) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryCompensatingControlDetails) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryCompensatingControlDetails) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryCompensatingControlDetails) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryCompensatingControlDetails) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAvailableControlsCount

`func (o *NiatelemetryCompensatingControlDetails) GetAvailableControlsCount() NiatelemetryCompensatingControlCounts`

GetAvailableControlsCount returns the AvailableControlsCount field if non-nil, zero value otherwise.

### GetAvailableControlsCountOk

`func (o *NiatelemetryCompensatingControlDetails) GetAvailableControlsCountOk() (*NiatelemetryCompensatingControlCounts, bool)`

GetAvailableControlsCountOk returns a tuple with the AvailableControlsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableControlsCount

`func (o *NiatelemetryCompensatingControlDetails) SetAvailableControlsCount(v NiatelemetryCompensatingControlCounts)`

SetAvailableControlsCount sets AvailableControlsCount field to given value.

### HasAvailableControlsCount

`func (o *NiatelemetryCompensatingControlDetails) HasAvailableControlsCount() bool`

HasAvailableControlsCount returns a boolean if a field has been set.

### SetAvailableControlsCountNil

`func (o *NiatelemetryCompensatingControlDetails) SetAvailableControlsCountNil(b bool)`

 SetAvailableControlsCountNil sets the value for AvailableControlsCount to be an explicit nil

### UnsetAvailableControlsCount
`func (o *NiatelemetryCompensatingControlDetails) UnsetAvailableControlsCount()`

UnsetAvailableControlsCount ensures that no value is present for AvailableControlsCount, not even an explicit nil
### GetMonitoringControlsCount

`func (o *NiatelemetryCompensatingControlDetails) GetMonitoringControlsCount() NiatelemetryCompensatingControlCounts`

GetMonitoringControlsCount returns the MonitoringControlsCount field if non-nil, zero value otherwise.

### GetMonitoringControlsCountOk

`func (o *NiatelemetryCompensatingControlDetails) GetMonitoringControlsCountOk() (*NiatelemetryCompensatingControlCounts, bool)`

GetMonitoringControlsCountOk returns a tuple with the MonitoringControlsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringControlsCount

`func (o *NiatelemetryCompensatingControlDetails) SetMonitoringControlsCount(v NiatelemetryCompensatingControlCounts)`

SetMonitoringControlsCount sets MonitoringControlsCount field to given value.

### HasMonitoringControlsCount

`func (o *NiatelemetryCompensatingControlDetails) HasMonitoringControlsCount() bool`

HasMonitoringControlsCount returns a boolean if a field has been set.

### SetMonitoringControlsCountNil

`func (o *NiatelemetryCompensatingControlDetails) SetMonitoringControlsCountNil(b bool)`

 SetMonitoringControlsCountNil sets the value for MonitoringControlsCount to be an explicit nil

### UnsetMonitoringControlsCount
`func (o *NiatelemetryCompensatingControlDetails) UnsetMonitoringControlsCount()`

UnsetMonitoringControlsCount ensures that no value is present for MonitoringControlsCount, not even an explicit nil
### GetProtectedControlsCount

`func (o *NiatelemetryCompensatingControlDetails) GetProtectedControlsCount() NiatelemetryCompensatingControlCounts`

GetProtectedControlsCount returns the ProtectedControlsCount field if non-nil, zero value otherwise.

### GetProtectedControlsCountOk

`func (o *NiatelemetryCompensatingControlDetails) GetProtectedControlsCountOk() (*NiatelemetryCompensatingControlCounts, bool)`

GetProtectedControlsCountOk returns a tuple with the ProtectedControlsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedControlsCount

`func (o *NiatelemetryCompensatingControlDetails) SetProtectedControlsCount(v NiatelemetryCompensatingControlCounts)`

SetProtectedControlsCount sets ProtectedControlsCount field to given value.

### HasProtectedControlsCount

`func (o *NiatelemetryCompensatingControlDetails) HasProtectedControlsCount() bool`

HasProtectedControlsCount returns a boolean if a field has been set.

### SetProtectedControlsCountNil

`func (o *NiatelemetryCompensatingControlDetails) SetProtectedControlsCountNil(b bool)`

 SetProtectedControlsCountNil sets the value for ProtectedControlsCount to be an explicit nil

### UnsetProtectedControlsCount
`func (o *NiatelemetryCompensatingControlDetails) UnsetProtectedControlsCount()`

UnsetProtectedControlsCount ensures that no value is present for ProtectedControlsCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


