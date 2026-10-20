# CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "catalystsdwan.WanEdgeDeviceMaintenanceActivityInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "catalystsdwan.WanEdgeDeviceMaintenanceActivityInfo"]
**ActivityType** | Pointer to **string** | Type of the maintenance action (e.g., Reboot, Upgrade). | [optional] [readonly] 
**ChassisIdentifier** | Pointer to [**NullableCatalystsdwanChassisIdentifier**](CatalystsdwanChassisIdentifier.md) |  | [optional] 
**ChassisNumber** | Pointer to **string** | Catalyst SD-WAN Manager&#39;s model-serial identifier of the Secure Router (vEdge module) under maintenance. Distinct from the enclosing Unified Edge chassis, which is identified separately by chassisIdentifier. | [optional] [readonly] 
**DeviceModel** | Pointer to **string** | The model of the Secure Router under maintenance. | [optional] [readonly] 
**DurationInMins** | Pointer to **string** | Duration of the maintenance activity in minutes. | [optional] [readonly] 
**HostName** | Pointer to **string** | The host name of the Secure Router under maintenance. | [optional] [readonly] 
**Location** | Pointer to **string** | The location where the maintenance activity is scheduled. | [optional] [readonly] 
**SiteName** | Pointer to **string** | The site name of the Secure Router scheduled for maintenance. | [optional] [readonly] 
**StartTimeInUtc** | Pointer to **string** | The scheduled start time for the maintenance of the Secure Router. | [optional] [readonly] 
**SystemIp** | Pointer to **string** | System IP of the Secure Router under maintenance. | [optional] [readonly] 

## Methods

### NewCatalystsdwanWanEdgeDeviceMaintenanceActivityInfo

`func NewCatalystsdwanWanEdgeDeviceMaintenanceActivityInfo(classId string, objectType string, ) *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo`

NewCatalystsdwanWanEdgeDeviceMaintenanceActivityInfo instantiates a new CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalystsdwanWanEdgeDeviceMaintenanceActivityInfoWithDefaults

`func NewCatalystsdwanWanEdgeDeviceMaintenanceActivityInfoWithDefaults() *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo`

NewCatalystsdwanWanEdgeDeviceMaintenanceActivityInfoWithDefaults instantiates a new CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetActivityType

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) GetActivityType() string`

GetActivityType returns the ActivityType field if non-nil, zero value otherwise.

### GetActivityTypeOk

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) GetActivityTypeOk() (*string, bool)`

GetActivityTypeOk returns a tuple with the ActivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityType

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) SetActivityType(v string)`

SetActivityType sets ActivityType field to given value.

### HasActivityType

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) HasActivityType() bool`

HasActivityType returns a boolean if a field has been set.

### GetChassisIdentifier

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) GetChassisIdentifier() CatalystsdwanChassisIdentifier`

GetChassisIdentifier returns the ChassisIdentifier field if non-nil, zero value otherwise.

### GetChassisIdentifierOk

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) GetChassisIdentifierOk() (*CatalystsdwanChassisIdentifier, bool)`

GetChassisIdentifierOk returns a tuple with the ChassisIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisIdentifier

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) SetChassisIdentifier(v CatalystsdwanChassisIdentifier)`

SetChassisIdentifier sets ChassisIdentifier field to given value.

### HasChassisIdentifier

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) HasChassisIdentifier() bool`

HasChassisIdentifier returns a boolean if a field has been set.

### SetChassisIdentifierNil

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) SetChassisIdentifierNil(b bool)`

 SetChassisIdentifierNil sets the value for ChassisIdentifier to be an explicit nil

### UnsetChassisIdentifier
`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) UnsetChassisIdentifier()`

UnsetChassisIdentifier ensures that no value is present for ChassisIdentifier, not even an explicit nil
### GetChassisNumber

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) GetChassisNumber() string`

GetChassisNumber returns the ChassisNumber field if non-nil, zero value otherwise.

### GetChassisNumberOk

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) GetChassisNumberOk() (*string, bool)`

GetChassisNumberOk returns a tuple with the ChassisNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisNumber

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) SetChassisNumber(v string)`

SetChassisNumber sets ChassisNumber field to given value.

### HasChassisNumber

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) HasChassisNumber() bool`

HasChassisNumber returns a boolean if a field has been set.

### GetDeviceModel

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) GetDeviceModelOk() (*string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModel

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) SetDeviceModel(v string)`

SetDeviceModel sets DeviceModel field to given value.

### HasDeviceModel

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### GetDurationInMins

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) GetDurationInMins() string`

GetDurationInMins returns the DurationInMins field if non-nil, zero value otherwise.

### GetDurationInMinsOk

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) GetDurationInMinsOk() (*string, bool)`

GetDurationInMinsOk returns a tuple with the DurationInMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInMins

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) SetDurationInMins(v string)`

SetDurationInMins sets DurationInMins field to given value.

### HasDurationInMins

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) HasDurationInMins() bool`

HasDurationInMins returns a boolean if a field has been set.

### GetHostName

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetLocation

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetSiteName

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetStartTimeInUtc

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) GetStartTimeInUtc() string`

GetStartTimeInUtc returns the StartTimeInUtc field if non-nil, zero value otherwise.

### GetStartTimeInUtcOk

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) GetStartTimeInUtcOk() (*string, bool)`

GetStartTimeInUtcOk returns a tuple with the StartTimeInUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeInUtc

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) SetStartTimeInUtc(v string)`

SetStartTimeInUtc sets StartTimeInUtc field to given value.

### HasStartTimeInUtc

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) HasStartTimeInUtc() bool`

HasStartTimeInUtc returns a boolean if a field has been set.

### GetSystemIp

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) GetSystemIp() string`

GetSystemIp returns the SystemIp field if non-nil, zero value otherwise.

### GetSystemIpOk

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) GetSystemIpOk() (*string, bool)`

GetSystemIpOk returns a tuple with the SystemIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIp

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) SetSystemIp(v string)`

SetSystemIp sets SystemIp field to given value.

### HasSystemIp

`func (o *CatalystsdwanWanEdgeDeviceMaintenanceActivityInfo) HasSystemIp() bool`

HasSystemIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


