# ApplianceDeviceUpgradePolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.DeviceUpgradePolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.DeviceUpgradePolicy"]
**AutoUpgrade** | Pointer to **bool** | Indicates if the upgrade service is set to automatically start the software upgrade or not. If autoUpgrade is true, then the value of the schedule field is ignored. | [optional] 
**BlackoutDatesEnabled** | Pointer to **bool** | If enabled, allows the user to define a blackout period during which the appliance will not be upgraded. | [optional] 
**BlackoutEndDate** | Pointer to **time.Time** | End date of the black out period. | [optional] 
**BlackoutStartDate** | Pointer to **time.Time** | Start date of the black out period. The appliance will not be upgraded during this period. | [optional] 
**EnableMetaDataSync** | Pointer to **bool** | Indicates if the updated metadata files should be synced immediately or at the next upgrade. | [optional] [default to true]
**Schedule** | Pointer to [**NullableOnpremSchedule**](OnpremSchedule.md) |  | [optional] 
**SerialId** | Pointer to **string** | SerialId of the Intersight Appliance. SerialId is generated when the Intersight Appliance is setup. It is a unique UUID string, and serialId will not change for the life time of the Intersight Appliance. | [optional] [readonly] 
**SoftwareDownloadType** | Pointer to **string** | UpgradeType is used to indicate the kink of software upload to upgrade. * &#x60;connected&#x60; - Indicates if the upgrade service is set to upload software to latest version automatically. * &#x60;manual&#x60; - Indicates if the upgrade service is set to upload software to user picked verison manually . | [optional] [default to "connected"]
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewApplianceDeviceUpgradePolicyAllOf

`func NewApplianceDeviceUpgradePolicyAllOf(classId string, objectType string, ) *ApplianceDeviceUpgradePolicyAllOf`

NewApplianceDeviceUpgradePolicyAllOf instantiates a new ApplianceDeviceUpgradePolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceDeviceUpgradePolicyAllOfWithDefaults

`func NewApplianceDeviceUpgradePolicyAllOfWithDefaults() *ApplianceDeviceUpgradePolicyAllOf`

NewApplianceDeviceUpgradePolicyAllOfWithDefaults instantiates a new ApplianceDeviceUpgradePolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceDeviceUpgradePolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceDeviceUpgradePolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceDeviceUpgradePolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceDeviceUpgradePolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceDeviceUpgradePolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceDeviceUpgradePolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAutoUpgrade

`func (o *ApplianceDeviceUpgradePolicyAllOf) GetAutoUpgrade() bool`

GetAutoUpgrade returns the AutoUpgrade field if non-nil, zero value otherwise.

### GetAutoUpgradeOk

`func (o *ApplianceDeviceUpgradePolicyAllOf) GetAutoUpgradeOk() (*bool, bool)`

GetAutoUpgradeOk returns a tuple with the AutoUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpgrade

`func (o *ApplianceDeviceUpgradePolicyAllOf) SetAutoUpgrade(v bool)`

SetAutoUpgrade sets AutoUpgrade field to given value.

### HasAutoUpgrade

`func (o *ApplianceDeviceUpgradePolicyAllOf) HasAutoUpgrade() bool`

HasAutoUpgrade returns a boolean if a field has been set.

### GetBlackoutDatesEnabled

`func (o *ApplianceDeviceUpgradePolicyAllOf) GetBlackoutDatesEnabled() bool`

GetBlackoutDatesEnabled returns the BlackoutDatesEnabled field if non-nil, zero value otherwise.

### GetBlackoutDatesEnabledOk

`func (o *ApplianceDeviceUpgradePolicyAllOf) GetBlackoutDatesEnabledOk() (*bool, bool)`

GetBlackoutDatesEnabledOk returns a tuple with the BlackoutDatesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutDatesEnabled

`func (o *ApplianceDeviceUpgradePolicyAllOf) SetBlackoutDatesEnabled(v bool)`

SetBlackoutDatesEnabled sets BlackoutDatesEnabled field to given value.

### HasBlackoutDatesEnabled

`func (o *ApplianceDeviceUpgradePolicyAllOf) HasBlackoutDatesEnabled() bool`

HasBlackoutDatesEnabled returns a boolean if a field has been set.

### GetBlackoutEndDate

`func (o *ApplianceDeviceUpgradePolicyAllOf) GetBlackoutEndDate() time.Time`

GetBlackoutEndDate returns the BlackoutEndDate field if non-nil, zero value otherwise.

### GetBlackoutEndDateOk

`func (o *ApplianceDeviceUpgradePolicyAllOf) GetBlackoutEndDateOk() (*time.Time, bool)`

GetBlackoutEndDateOk returns a tuple with the BlackoutEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutEndDate

`func (o *ApplianceDeviceUpgradePolicyAllOf) SetBlackoutEndDate(v time.Time)`

SetBlackoutEndDate sets BlackoutEndDate field to given value.

### HasBlackoutEndDate

`func (o *ApplianceDeviceUpgradePolicyAllOf) HasBlackoutEndDate() bool`

HasBlackoutEndDate returns a boolean if a field has been set.

### GetBlackoutStartDate

`func (o *ApplianceDeviceUpgradePolicyAllOf) GetBlackoutStartDate() time.Time`

GetBlackoutStartDate returns the BlackoutStartDate field if non-nil, zero value otherwise.

### GetBlackoutStartDateOk

`func (o *ApplianceDeviceUpgradePolicyAllOf) GetBlackoutStartDateOk() (*time.Time, bool)`

GetBlackoutStartDateOk returns a tuple with the BlackoutStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutStartDate

`func (o *ApplianceDeviceUpgradePolicyAllOf) SetBlackoutStartDate(v time.Time)`

SetBlackoutStartDate sets BlackoutStartDate field to given value.

### HasBlackoutStartDate

`func (o *ApplianceDeviceUpgradePolicyAllOf) HasBlackoutStartDate() bool`

HasBlackoutStartDate returns a boolean if a field has been set.

### GetEnableMetaDataSync

`func (o *ApplianceDeviceUpgradePolicyAllOf) GetEnableMetaDataSync() bool`

GetEnableMetaDataSync returns the EnableMetaDataSync field if non-nil, zero value otherwise.

### GetEnableMetaDataSyncOk

`func (o *ApplianceDeviceUpgradePolicyAllOf) GetEnableMetaDataSyncOk() (*bool, bool)`

GetEnableMetaDataSyncOk returns a tuple with the EnableMetaDataSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMetaDataSync

`func (o *ApplianceDeviceUpgradePolicyAllOf) SetEnableMetaDataSync(v bool)`

SetEnableMetaDataSync sets EnableMetaDataSync field to given value.

### HasEnableMetaDataSync

`func (o *ApplianceDeviceUpgradePolicyAllOf) HasEnableMetaDataSync() bool`

HasEnableMetaDataSync returns a boolean if a field has been set.

### GetSchedule

`func (o *ApplianceDeviceUpgradePolicyAllOf) GetSchedule() OnpremSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ApplianceDeviceUpgradePolicyAllOf) GetScheduleOk() (*OnpremSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ApplianceDeviceUpgradePolicyAllOf) SetSchedule(v OnpremSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ApplianceDeviceUpgradePolicyAllOf) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### SetScheduleNil

`func (o *ApplianceDeviceUpgradePolicyAllOf) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *ApplianceDeviceUpgradePolicyAllOf) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
### GetSerialId

`func (o *ApplianceDeviceUpgradePolicyAllOf) GetSerialId() string`

GetSerialId returns the SerialId field if non-nil, zero value otherwise.

### GetSerialIdOk

`func (o *ApplianceDeviceUpgradePolicyAllOf) GetSerialIdOk() (*string, bool)`

GetSerialIdOk returns a tuple with the SerialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialId

`func (o *ApplianceDeviceUpgradePolicyAllOf) SetSerialId(v string)`

SetSerialId sets SerialId field to given value.

### HasSerialId

`func (o *ApplianceDeviceUpgradePolicyAllOf) HasSerialId() bool`

HasSerialId returns a boolean if a field has been set.

### GetSoftwareDownloadType

`func (o *ApplianceDeviceUpgradePolicyAllOf) GetSoftwareDownloadType() string`

GetSoftwareDownloadType returns the SoftwareDownloadType field if non-nil, zero value otherwise.

### GetSoftwareDownloadTypeOk

`func (o *ApplianceDeviceUpgradePolicyAllOf) GetSoftwareDownloadTypeOk() (*string, bool)`

GetSoftwareDownloadTypeOk returns a tuple with the SoftwareDownloadType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareDownloadType

`func (o *ApplianceDeviceUpgradePolicyAllOf) SetSoftwareDownloadType(v string)`

SetSoftwareDownloadType sets SoftwareDownloadType field to given value.

### HasSoftwareDownloadType

`func (o *ApplianceDeviceUpgradePolicyAllOf) HasSoftwareDownloadType() bool`

HasSoftwareDownloadType returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *ApplianceDeviceUpgradePolicyAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *ApplianceDeviceUpgradePolicyAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *ApplianceDeviceUpgradePolicyAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *ApplianceDeviceUpgradePolicyAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


