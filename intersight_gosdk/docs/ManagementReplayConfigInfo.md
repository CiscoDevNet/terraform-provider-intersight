# ManagementReplayConfigInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "management.ReplayConfigInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "management.ReplayConfigInfo"]
**ReplayConfigEndTime** | Pointer to **time.Time** | Time at which replay config ended. | [optional] [readonly] 
**ReplayConfigStartTime** | Pointer to **time.Time** | Time at which replay config started. | [optional] [readonly] 
**ReplayConfigStatus** | Pointer to **string** | Status of replay config of the Fabric Interconnect. * &#x60;Completed&#x60; - Configuration replay is complete. * &#x60;InProgress&#x60; - Configuration replay is in progress. * &#x60;TimedOut&#x60; - Configuration replay timed out. * &#x60;Fail&#x60; - Configuration replay failed. | [optional] [readonly] [default to "Completed"]

## Methods

### NewManagementReplayConfigInfo

`func NewManagementReplayConfigInfo(classId string, objectType string, ) *ManagementReplayConfigInfo`

NewManagementReplayConfigInfo instantiates a new ManagementReplayConfigInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementReplayConfigInfoWithDefaults

`func NewManagementReplayConfigInfoWithDefaults() *ManagementReplayConfigInfo`

NewManagementReplayConfigInfoWithDefaults instantiates a new ManagementReplayConfigInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ManagementReplayConfigInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ManagementReplayConfigInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ManagementReplayConfigInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ManagementReplayConfigInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ManagementReplayConfigInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ManagementReplayConfigInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetReplayConfigEndTime

`func (o *ManagementReplayConfigInfo) GetReplayConfigEndTime() time.Time`

GetReplayConfigEndTime returns the ReplayConfigEndTime field if non-nil, zero value otherwise.

### GetReplayConfigEndTimeOk

`func (o *ManagementReplayConfigInfo) GetReplayConfigEndTimeOk() (*time.Time, bool)`

GetReplayConfigEndTimeOk returns a tuple with the ReplayConfigEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayConfigEndTime

`func (o *ManagementReplayConfigInfo) SetReplayConfigEndTime(v time.Time)`

SetReplayConfigEndTime sets ReplayConfigEndTime field to given value.

### HasReplayConfigEndTime

`func (o *ManagementReplayConfigInfo) HasReplayConfigEndTime() bool`

HasReplayConfigEndTime returns a boolean if a field has been set.

### GetReplayConfigStartTime

`func (o *ManagementReplayConfigInfo) GetReplayConfigStartTime() time.Time`

GetReplayConfigStartTime returns the ReplayConfigStartTime field if non-nil, zero value otherwise.

### GetReplayConfigStartTimeOk

`func (o *ManagementReplayConfigInfo) GetReplayConfigStartTimeOk() (*time.Time, bool)`

GetReplayConfigStartTimeOk returns a tuple with the ReplayConfigStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayConfigStartTime

`func (o *ManagementReplayConfigInfo) SetReplayConfigStartTime(v time.Time)`

SetReplayConfigStartTime sets ReplayConfigStartTime field to given value.

### HasReplayConfigStartTime

`func (o *ManagementReplayConfigInfo) HasReplayConfigStartTime() bool`

HasReplayConfigStartTime returns a boolean if a field has been set.

### GetReplayConfigStatus

`func (o *ManagementReplayConfigInfo) GetReplayConfigStatus() string`

GetReplayConfigStatus returns the ReplayConfigStatus field if non-nil, zero value otherwise.

### GetReplayConfigStatusOk

`func (o *ManagementReplayConfigInfo) GetReplayConfigStatusOk() (*string, bool)`

GetReplayConfigStatusOk returns a tuple with the ReplayConfigStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayConfigStatus

`func (o *ManagementReplayConfigInfo) SetReplayConfigStatus(v string)`

SetReplayConfigStatus sets ReplayConfigStatus field to given value.

### HasReplayConfigStatus

`func (o *ManagementReplayConfigInfo) HasReplayConfigStatus() bool`

HasReplayConfigStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


