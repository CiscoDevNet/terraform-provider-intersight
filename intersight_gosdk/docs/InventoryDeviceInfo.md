# InventoryDeviceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "inventory.DeviceInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "inventory.DeviceInfo"]
**DisableEvents** | Pointer to **bool** | Subscribe/Unsubscribe events for the device. | [optional] [readonly] 
**EventCounter** | Pointer to **int64** | Information regarding the number of events recorded for this device. | [optional] [readonly] 
**EventCounterEnabled** | Pointer to **bool** | Indicates whether the event counter enabled for the device. | [optional] [readonly] 
**Interval** | Pointer to **int64** | The time interval (in hours) between job runs. | [optional] [readonly] 
**JobInfo** | Pointer to [**[]InventoryJobInfo**](InventoryJobInfo.md) |  | [optional] 
**LastReInventoryTime** | Pointer to **time.Time** | Last Reinventory time of the device. | [optional] [readonly] 
**ReInventoryCount** | Pointer to **int64** | Number of full inventory within the given time. | [optional] [readonly] 
**ClusterMember** | Pointer to [**NullableAssetClusterMemberRelationship**](AssetClusterMemberRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewInventoryDeviceInfo

`func NewInventoryDeviceInfo(classId string, objectType string, ) *InventoryDeviceInfo`

NewInventoryDeviceInfo instantiates a new InventoryDeviceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryDeviceInfoWithDefaults

`func NewInventoryDeviceInfoWithDefaults() *InventoryDeviceInfo`

NewInventoryDeviceInfoWithDefaults instantiates a new InventoryDeviceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *InventoryDeviceInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *InventoryDeviceInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *InventoryDeviceInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *InventoryDeviceInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *InventoryDeviceInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *InventoryDeviceInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisableEvents

`func (o *InventoryDeviceInfo) GetDisableEvents() bool`

GetDisableEvents returns the DisableEvents field if non-nil, zero value otherwise.

### GetDisableEventsOk

`func (o *InventoryDeviceInfo) GetDisableEventsOk() (*bool, bool)`

GetDisableEventsOk returns a tuple with the DisableEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEvents

`func (o *InventoryDeviceInfo) SetDisableEvents(v bool)`

SetDisableEvents sets DisableEvents field to given value.

### HasDisableEvents

`func (o *InventoryDeviceInfo) HasDisableEvents() bool`

HasDisableEvents returns a boolean if a field has been set.

### GetEventCounter

`func (o *InventoryDeviceInfo) GetEventCounter() int64`

GetEventCounter returns the EventCounter field if non-nil, zero value otherwise.

### GetEventCounterOk

`func (o *InventoryDeviceInfo) GetEventCounterOk() (*int64, bool)`

GetEventCounterOk returns a tuple with the EventCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCounter

`func (o *InventoryDeviceInfo) SetEventCounter(v int64)`

SetEventCounter sets EventCounter field to given value.

### HasEventCounter

`func (o *InventoryDeviceInfo) HasEventCounter() bool`

HasEventCounter returns a boolean if a field has been set.

### GetEventCounterEnabled

`func (o *InventoryDeviceInfo) GetEventCounterEnabled() bool`

GetEventCounterEnabled returns the EventCounterEnabled field if non-nil, zero value otherwise.

### GetEventCounterEnabledOk

`func (o *InventoryDeviceInfo) GetEventCounterEnabledOk() (*bool, bool)`

GetEventCounterEnabledOk returns a tuple with the EventCounterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCounterEnabled

`func (o *InventoryDeviceInfo) SetEventCounterEnabled(v bool)`

SetEventCounterEnabled sets EventCounterEnabled field to given value.

### HasEventCounterEnabled

`func (o *InventoryDeviceInfo) HasEventCounterEnabled() bool`

HasEventCounterEnabled returns a boolean if a field has been set.

### GetInterval

`func (o *InventoryDeviceInfo) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *InventoryDeviceInfo) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *InventoryDeviceInfo) SetInterval(v int64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *InventoryDeviceInfo) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetJobInfo

`func (o *InventoryDeviceInfo) GetJobInfo() []InventoryJobInfo`

GetJobInfo returns the JobInfo field if non-nil, zero value otherwise.

### GetJobInfoOk

`func (o *InventoryDeviceInfo) GetJobInfoOk() (*[]InventoryJobInfo, bool)`

GetJobInfoOk returns a tuple with the JobInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobInfo

`func (o *InventoryDeviceInfo) SetJobInfo(v []InventoryJobInfo)`

SetJobInfo sets JobInfo field to given value.

### HasJobInfo

`func (o *InventoryDeviceInfo) HasJobInfo() bool`

HasJobInfo returns a boolean if a field has been set.

### SetJobInfoNil

`func (o *InventoryDeviceInfo) SetJobInfoNil(b bool)`

 SetJobInfoNil sets the value for JobInfo to be an explicit nil

### UnsetJobInfo
`func (o *InventoryDeviceInfo) UnsetJobInfo()`

UnsetJobInfo ensures that no value is present for JobInfo, not even an explicit nil
### GetLastReInventoryTime

`func (o *InventoryDeviceInfo) GetLastReInventoryTime() time.Time`

GetLastReInventoryTime returns the LastReInventoryTime field if non-nil, zero value otherwise.

### GetLastReInventoryTimeOk

`func (o *InventoryDeviceInfo) GetLastReInventoryTimeOk() (*time.Time, bool)`

GetLastReInventoryTimeOk returns a tuple with the LastReInventoryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReInventoryTime

`func (o *InventoryDeviceInfo) SetLastReInventoryTime(v time.Time)`

SetLastReInventoryTime sets LastReInventoryTime field to given value.

### HasLastReInventoryTime

`func (o *InventoryDeviceInfo) HasLastReInventoryTime() bool`

HasLastReInventoryTime returns a boolean if a field has been set.

### GetReInventoryCount

`func (o *InventoryDeviceInfo) GetReInventoryCount() int64`

GetReInventoryCount returns the ReInventoryCount field if non-nil, zero value otherwise.

### GetReInventoryCountOk

`func (o *InventoryDeviceInfo) GetReInventoryCountOk() (*int64, bool)`

GetReInventoryCountOk returns a tuple with the ReInventoryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReInventoryCount

`func (o *InventoryDeviceInfo) SetReInventoryCount(v int64)`

SetReInventoryCount sets ReInventoryCount field to given value.

### HasReInventoryCount

`func (o *InventoryDeviceInfo) HasReInventoryCount() bool`

HasReInventoryCount returns a boolean if a field has been set.

### GetClusterMember

`func (o *InventoryDeviceInfo) GetClusterMember() AssetClusterMemberRelationship`

GetClusterMember returns the ClusterMember field if non-nil, zero value otherwise.

### GetClusterMemberOk

`func (o *InventoryDeviceInfo) GetClusterMemberOk() (*AssetClusterMemberRelationship, bool)`

GetClusterMemberOk returns a tuple with the ClusterMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterMember

`func (o *InventoryDeviceInfo) SetClusterMember(v AssetClusterMemberRelationship)`

SetClusterMember sets ClusterMember field to given value.

### HasClusterMember

`func (o *InventoryDeviceInfo) HasClusterMember() bool`

HasClusterMember returns a boolean if a field has been set.

### SetClusterMemberNil

`func (o *InventoryDeviceInfo) SetClusterMemberNil(b bool)`

 SetClusterMemberNil sets the value for ClusterMember to be an explicit nil

### UnsetClusterMember
`func (o *InventoryDeviceInfo) UnsetClusterMember()`

UnsetClusterMember ensures that no value is present for ClusterMember, not even an explicit nil
### GetRegisteredDevice

`func (o *InventoryDeviceInfo) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *InventoryDeviceInfo) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *InventoryDeviceInfo) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *InventoryDeviceInfo) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *InventoryDeviceInfo) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *InventoryDeviceInfo) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


