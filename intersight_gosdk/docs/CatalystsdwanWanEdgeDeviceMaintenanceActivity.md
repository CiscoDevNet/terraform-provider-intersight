# CatalystsdwanWanEdgeDeviceMaintenanceActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "catalystsdwan.WanEdgeDeviceMaintenanceActivity"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "catalystsdwan.WanEdgeDeviceMaintenanceActivity"]
**ActivityType** | Pointer to **string** | Type of the maintenance action (e.g., Reboot, upgrade). | [optional] [readonly] 
**ChassisNumber** | Pointer to **string** | The chassis number of the WAN Edge Device under maintenance. | [optional] [readonly] 
**DeviceModel** | Pointer to **string** | The model of the WAN Edge Device under maintenance. | [optional] [readonly] 
**DurationInMins** | Pointer to **string** | Duration of the maintenance activity in minutes. | [optional] [readonly] 
**HostName** | Pointer to **string** | The host name of the WAN Edge Device under maintenance. | [optional] [readonly] 
**Location** | Pointer to **string** | The location where the maintenance activity is scheduled. | [optional] [readonly] 
**SiteName** | Pointer to **string** | The site name of the WAN Edge Device scheduled for maintenance. | [optional] [readonly] 
**StartTimeInUtc** | Pointer to **string** | The scheduled start time for the maintenance of the WAN Edge Device. | [optional] [readonly] 
**SystemIp** | Pointer to **string** | System IP of the WAN Edge Device under maintenance. | [optional] [readonly] 
**WanEdgeDevice** | Pointer to [**NullableCatalystsdwanWanEdgeDeviceRelationship**](CatalystsdwanWanEdgeDeviceRelationship.md) |  | [optional] 

## Methods

### NewCatalystsdwanWanEdgeDeviceMaintenanceActivity

`func NewCatalystsdwanWanEdgeDeviceMaintenanceActivity(classId string, objectType string, ) *CatalystsdwanWanEdgeDeviceMaintenanceActivity`

NewCatalystsdwanWanEdgeDeviceMaintenanceActivity instantiates a new CatalystsdwanWanEdgeDeviceMaintenanceActivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalystsdwanWanEdgeDeviceMaintenanceActivityWithDefaults

`func NewCatalystsdwanWanEdgeDeviceMaintenanceActivityWithDefaults() *CatalystsdwanWanEdgeDeviceMaintenanceActivity`

NewCatalystsdwanWanEdgeDeviceMaintenanceActivityWithDefaults instantiates a new CatalystsdwanWanEdgeDeviceMaintenanceActivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetActivityType

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) GetActivityType() string`

GetActivityType returns the ActivityType field if non-nil, zero value otherwise.

### GetActivityTypeOk

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) GetActivityTypeOk() (*string, bool)`

GetActivityTypeOk returns a tuple with the ActivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityType

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) SetActivityType(v string)`

SetActivityType sets ActivityType field to given value.

### HasActivityType

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) HasActivityType() bool`

HasActivityType returns a boolean if a field has been set.

### GetChassisNumber

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) GetChassisNumber() string`

GetChassisNumber returns the ChassisNumber field if non-nil, zero value otherwise.

### GetChassisNumberOk

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) GetChassisNumberOk() (*string, bool)`

GetChassisNumberOk returns a tuple with the ChassisNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisNumber

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) SetChassisNumber(v string)`

SetChassisNumber sets ChassisNumber field to given value.

### HasChassisNumber

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) HasChassisNumber() bool`

HasChassisNumber returns a boolean if a field has been set.

### GetDeviceModel

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) GetDeviceModelOk() (*string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModel

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) SetDeviceModel(v string)`

SetDeviceModel sets DeviceModel field to given value.

### HasDeviceModel

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### GetDurationInMins

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) GetDurationInMins() string`

GetDurationInMins returns the DurationInMins field if non-nil, zero value otherwise.

### GetDurationInMinsOk

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) GetDurationInMinsOk() (*string, bool)`

GetDurationInMinsOk returns a tuple with the DurationInMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInMins

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) SetDurationInMins(v string)`

SetDurationInMins sets DurationInMins field to given value.

### HasDurationInMins

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) HasDurationInMins() bool`

HasDurationInMins returns a boolean if a field has been set.

### GetHostName

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetLocation

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetSiteName

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetStartTimeInUtc

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) GetStartTimeInUtc() string`

GetStartTimeInUtc returns the StartTimeInUtc field if non-nil, zero value otherwise.

### GetStartTimeInUtcOk

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) GetStartTimeInUtcOk() (*string, bool)`

GetStartTimeInUtcOk returns a tuple with the StartTimeInUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeInUtc

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) SetStartTimeInUtc(v string)`

SetStartTimeInUtc sets StartTimeInUtc field to given value.

### HasStartTimeInUtc

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) HasStartTimeInUtc() bool`

HasStartTimeInUtc returns a boolean if a field has been set.

### GetSystemIp

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) GetSystemIp() string`

GetSystemIp returns the SystemIp field if non-nil, zero value otherwise.

### GetSystemIpOk

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) GetSystemIpOk() (*string, bool)`

GetSystemIpOk returns a tuple with the SystemIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIp

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) SetSystemIp(v string)`

SetSystemIp sets SystemIp field to given value.

### HasSystemIp

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) HasSystemIp() bool`

HasSystemIp returns a boolean if a field has been set.

### GetWanEdgeDevice

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) GetWanEdgeDevice() CatalystsdwanWanEdgeDeviceRelationship`

GetWanEdgeDevice returns the WanEdgeDevice field if non-nil, zero value otherwise.

### GetWanEdgeDeviceOk

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) GetWanEdgeDeviceOk() (*CatalystsdwanWanEdgeDeviceRelationship, bool)`

GetWanEdgeDeviceOk returns a tuple with the WanEdgeDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanEdgeDevice

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) SetWanEdgeDevice(v CatalystsdwanWanEdgeDeviceRelationship)`

SetWanEdgeDevice sets WanEdgeDevice field to given value.

### HasWanEdgeDevice

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) HasWanEdgeDevice() bool`

HasWanEdgeDevice returns a boolean if a field has been set.

### SetWanEdgeDeviceNil

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) SetWanEdgeDeviceNil(b bool)`

 SetWanEdgeDeviceNil sets the value for WanEdgeDevice to be an explicit nil

### UnsetWanEdgeDevice
`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivity) UnsetWanEdgeDevice()`

UnsetWanEdgeDevice ensures that no value is present for WanEdgeDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


