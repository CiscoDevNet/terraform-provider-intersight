# NiatelemetryNxosVtp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.NxosVtp"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.NxosVtp"]
**OperMode** | Pointer to **string** | Returns the status of operational mode of vtp. | [optional] 
**PruningMode** | Pointer to **string** | Returns the status pruning mode of vtp. | [optional] 
**RunningVersion** | Pointer to **string** | Returns the running version of vtp. | [optional] 
**TrapEnabled** | Pointer to **string** | Returns the status of trap in vtp. | [optional] 
**V2Mode** | Pointer to **string** | Returns the status of v2 mode of vtp. | [optional] 
**Version** | Pointer to **int64** | Returns version of vtp running. | [optional] 

## Methods

### NewNiatelemetryNxosVtp

`func NewNiatelemetryNxosVtp(classId string, objectType string, ) *NiatelemetryNxosVtp`

NewNiatelemetryNxosVtp instantiates a new NiatelemetryNxosVtp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryNxosVtpWithDefaults

`func NewNiatelemetryNxosVtpWithDefaults() *NiatelemetryNxosVtp`

NewNiatelemetryNxosVtpWithDefaults instantiates a new NiatelemetryNxosVtp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryNxosVtp) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryNxosVtp) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryNxosVtp) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryNxosVtp) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryNxosVtp) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryNxosVtp) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOperMode

`func (o *NiatelemetryNxosVtp) GetOperMode() string`

GetOperMode returns the OperMode field if non-nil, zero value otherwise.

### GetOperModeOk

`func (o *NiatelemetryNxosVtp) GetOperModeOk() (*string, bool)`

GetOperModeOk returns a tuple with the OperMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperMode

`func (o *NiatelemetryNxosVtp) SetOperMode(v string)`

SetOperMode sets OperMode field to given value.

### HasOperMode

`func (o *NiatelemetryNxosVtp) HasOperMode() bool`

HasOperMode returns a boolean if a field has been set.

### GetPruningMode

`func (o *NiatelemetryNxosVtp) GetPruningMode() string`

GetPruningMode returns the PruningMode field if non-nil, zero value otherwise.

### GetPruningModeOk

`func (o *NiatelemetryNxosVtp) GetPruningModeOk() (*string, bool)`

GetPruningModeOk returns a tuple with the PruningMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPruningMode

`func (o *NiatelemetryNxosVtp) SetPruningMode(v string)`

SetPruningMode sets PruningMode field to given value.

### HasPruningMode

`func (o *NiatelemetryNxosVtp) HasPruningMode() bool`

HasPruningMode returns a boolean if a field has been set.

### GetRunningVersion

`func (o *NiatelemetryNxosVtp) GetRunningVersion() string`

GetRunningVersion returns the RunningVersion field if non-nil, zero value otherwise.

### GetRunningVersionOk

`func (o *NiatelemetryNxosVtp) GetRunningVersionOk() (*string, bool)`

GetRunningVersionOk returns a tuple with the RunningVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningVersion

`func (o *NiatelemetryNxosVtp) SetRunningVersion(v string)`

SetRunningVersion sets RunningVersion field to given value.

### HasRunningVersion

`func (o *NiatelemetryNxosVtp) HasRunningVersion() bool`

HasRunningVersion returns a boolean if a field has been set.

### GetTrapEnabled

`func (o *NiatelemetryNxosVtp) GetTrapEnabled() string`

GetTrapEnabled returns the TrapEnabled field if non-nil, zero value otherwise.

### GetTrapEnabledOk

`func (o *NiatelemetryNxosVtp) GetTrapEnabledOk() (*string, bool)`

GetTrapEnabledOk returns a tuple with the TrapEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrapEnabled

`func (o *NiatelemetryNxosVtp) SetTrapEnabled(v string)`

SetTrapEnabled sets TrapEnabled field to given value.

### HasTrapEnabled

`func (o *NiatelemetryNxosVtp) HasTrapEnabled() bool`

HasTrapEnabled returns a boolean if a field has been set.

### GetV2Mode

`func (o *NiatelemetryNxosVtp) GetV2Mode() string`

GetV2Mode returns the V2Mode field if non-nil, zero value otherwise.

### GetV2ModeOk

`func (o *NiatelemetryNxosVtp) GetV2ModeOk() (*string, bool)`

GetV2ModeOk returns a tuple with the V2Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV2Mode

`func (o *NiatelemetryNxosVtp) SetV2Mode(v string)`

SetV2Mode sets V2Mode field to given value.

### HasV2Mode

`func (o *NiatelemetryNxosVtp) HasV2Mode() bool`

HasV2Mode returns a boolean if a field has been set.

### GetVersion

`func (o *NiatelemetryNxosVtp) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NiatelemetryNxosVtp) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NiatelemetryNxosVtp) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NiatelemetryNxosVtp) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


