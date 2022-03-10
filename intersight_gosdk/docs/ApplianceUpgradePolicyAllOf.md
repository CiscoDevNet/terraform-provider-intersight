# ApplianceUpgradePolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.UpgradePolicy"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.UpgradePolicy"]
**AutoUpgrade** | Pointer to **bool** | Indicates if the upgrade service is set to automatically start the software upgrade or not. If autoUpgrade is true, then the value of the schedule field is ignored. | [optional] 
**BlackoutDatesEnabled** | Pointer to **bool** | If enabled, allows the user to define a blackout period during which the appliance will not be upgraded. | [optional] 
**BlackoutEndDate** | Pointer to **time.Time** | End date of the black out period. | [optional] 
**BlackoutStartDate** | Pointer to **time.Time** | Start date of the black out period. The appliance will not be upgraded during this period. | [optional] 
**EnableMetaDataSync** | Pointer to **bool** | Indicates if the updated metadata files should be synced immediately or at the next upgrade. | [optional] [default to true]
**IsSynced** | Pointer to **bool** | Flag to indicate software upgrade setting is synchronized with Intersight SaaS. | [optional] [readonly] 
**ManualInstallationStartTime** | Pointer to **time.Time** | Intersight Appliance manual upgrade start time. | [optional] 
**Schedule** | Pointer to [**NullableOnpremSchedule**](OnpremSchedule.md) |  | [optional] 
**SoftwareDownloadType** | Pointer to **string** | SoftwareDownloadType is used to indicate the kind of software download. * &#x60;unknown&#x60; - Indicates user setting of upgrade service to unknown. * &#x60;connected&#x60; - Indicates if the upgrade service is set to upload software to latest version automatically. * &#x60;manual&#x60; - Indicates if the upgrade service is set to upload software to user picked verison manually. | [optional] [default to "unknown"]
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 

## Methods

### NewApplianceUpgradePolicyAllOf

`func NewApplianceUpgradePolicyAllOf(classId string, objectType string, ) *ApplianceUpgradePolicyAllOf`

NewApplianceUpgradePolicyAllOf instantiates a new ApplianceUpgradePolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceUpgradePolicyAllOfWithDefaults

`func NewApplianceUpgradePolicyAllOfWithDefaults() *ApplianceUpgradePolicyAllOf`

NewApplianceUpgradePolicyAllOfWithDefaults instantiates a new ApplianceUpgradePolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceUpgradePolicyAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceUpgradePolicyAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceUpgradePolicyAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceUpgradePolicyAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceUpgradePolicyAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceUpgradePolicyAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAutoUpgrade

`func (o *ApplianceUpgradePolicyAllOf) GetAutoUpgrade() bool`

GetAutoUpgrade returns the AutoUpgrade field if non-nil, zero value otherwise.

### GetAutoUpgradeOk

`func (o *ApplianceUpgradePolicyAllOf) GetAutoUpgradeOk() (*bool, bool)`

GetAutoUpgradeOk returns a tuple with the AutoUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpgrade

`func (o *ApplianceUpgradePolicyAllOf) SetAutoUpgrade(v bool)`

SetAutoUpgrade sets AutoUpgrade field to given value.

### HasAutoUpgrade

`func (o *ApplianceUpgradePolicyAllOf) HasAutoUpgrade() bool`

HasAutoUpgrade returns a boolean if a field has been set.

### GetBlackoutDatesEnabled

`func (o *ApplianceUpgradePolicyAllOf) GetBlackoutDatesEnabled() bool`

GetBlackoutDatesEnabled returns the BlackoutDatesEnabled field if non-nil, zero value otherwise.

### GetBlackoutDatesEnabledOk

`func (o *ApplianceUpgradePolicyAllOf) GetBlackoutDatesEnabledOk() (*bool, bool)`

GetBlackoutDatesEnabledOk returns a tuple with the BlackoutDatesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutDatesEnabled

`func (o *ApplianceUpgradePolicyAllOf) SetBlackoutDatesEnabled(v bool)`

SetBlackoutDatesEnabled sets BlackoutDatesEnabled field to given value.

### HasBlackoutDatesEnabled

`func (o *ApplianceUpgradePolicyAllOf) HasBlackoutDatesEnabled() bool`

HasBlackoutDatesEnabled returns a boolean if a field has been set.

### GetBlackoutEndDate

`func (o *ApplianceUpgradePolicyAllOf) GetBlackoutEndDate() time.Time`

GetBlackoutEndDate returns the BlackoutEndDate field if non-nil, zero value otherwise.

### GetBlackoutEndDateOk

`func (o *ApplianceUpgradePolicyAllOf) GetBlackoutEndDateOk() (*time.Time, bool)`

GetBlackoutEndDateOk returns a tuple with the BlackoutEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutEndDate

`func (o *ApplianceUpgradePolicyAllOf) SetBlackoutEndDate(v time.Time)`

SetBlackoutEndDate sets BlackoutEndDate field to given value.

### HasBlackoutEndDate

`func (o *ApplianceUpgradePolicyAllOf) HasBlackoutEndDate() bool`

HasBlackoutEndDate returns a boolean if a field has been set.

### GetBlackoutStartDate

`func (o *ApplianceUpgradePolicyAllOf) GetBlackoutStartDate() time.Time`

GetBlackoutStartDate returns the BlackoutStartDate field if non-nil, zero value otherwise.

### GetBlackoutStartDateOk

`func (o *ApplianceUpgradePolicyAllOf) GetBlackoutStartDateOk() (*time.Time, bool)`

GetBlackoutStartDateOk returns a tuple with the BlackoutStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutStartDate

`func (o *ApplianceUpgradePolicyAllOf) SetBlackoutStartDate(v time.Time)`

SetBlackoutStartDate sets BlackoutStartDate field to given value.

### HasBlackoutStartDate

`func (o *ApplianceUpgradePolicyAllOf) HasBlackoutStartDate() bool`

HasBlackoutStartDate returns a boolean if a field has been set.

### GetEnableMetaDataSync

`func (o *ApplianceUpgradePolicyAllOf) GetEnableMetaDataSync() bool`

GetEnableMetaDataSync returns the EnableMetaDataSync field if non-nil, zero value otherwise.

### GetEnableMetaDataSyncOk

`func (o *ApplianceUpgradePolicyAllOf) GetEnableMetaDataSyncOk() (*bool, bool)`

GetEnableMetaDataSyncOk returns a tuple with the EnableMetaDataSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMetaDataSync

`func (o *ApplianceUpgradePolicyAllOf) SetEnableMetaDataSync(v bool)`

SetEnableMetaDataSync sets EnableMetaDataSync field to given value.

### HasEnableMetaDataSync

`func (o *ApplianceUpgradePolicyAllOf) HasEnableMetaDataSync() bool`

HasEnableMetaDataSync returns a boolean if a field has been set.

### GetIsSynced

`func (o *ApplianceUpgradePolicyAllOf) GetIsSynced() bool`

GetIsSynced returns the IsSynced field if non-nil, zero value otherwise.

### GetIsSyncedOk

`func (o *ApplianceUpgradePolicyAllOf) GetIsSyncedOk() (*bool, bool)`

GetIsSyncedOk returns a tuple with the IsSynced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSynced

`func (o *ApplianceUpgradePolicyAllOf) SetIsSynced(v bool)`

SetIsSynced sets IsSynced field to given value.

### HasIsSynced

`func (o *ApplianceUpgradePolicyAllOf) HasIsSynced() bool`

HasIsSynced returns a boolean if a field has been set.

### GetManualInstallationStartTime

`func (o *ApplianceUpgradePolicyAllOf) GetManualInstallationStartTime() time.Time`

GetManualInstallationStartTime returns the ManualInstallationStartTime field if non-nil, zero value otherwise.

### GetManualInstallationStartTimeOk

`func (o *ApplianceUpgradePolicyAllOf) GetManualInstallationStartTimeOk() (*time.Time, bool)`

GetManualInstallationStartTimeOk returns a tuple with the ManualInstallationStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualInstallationStartTime

`func (o *ApplianceUpgradePolicyAllOf) SetManualInstallationStartTime(v time.Time)`

SetManualInstallationStartTime sets ManualInstallationStartTime field to given value.

### HasManualInstallationStartTime

`func (o *ApplianceUpgradePolicyAllOf) HasManualInstallationStartTime() bool`

HasManualInstallationStartTime returns a boolean if a field has been set.

### GetSchedule

`func (o *ApplianceUpgradePolicyAllOf) GetSchedule() OnpremSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ApplianceUpgradePolicyAllOf) GetScheduleOk() (*OnpremSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ApplianceUpgradePolicyAllOf) SetSchedule(v OnpremSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ApplianceUpgradePolicyAllOf) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### SetScheduleNil

`func (o *ApplianceUpgradePolicyAllOf) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *ApplianceUpgradePolicyAllOf) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
### GetSoftwareDownloadType

`func (o *ApplianceUpgradePolicyAllOf) GetSoftwareDownloadType() string`

GetSoftwareDownloadType returns the SoftwareDownloadType field if non-nil, zero value otherwise.

### GetSoftwareDownloadTypeOk

`func (o *ApplianceUpgradePolicyAllOf) GetSoftwareDownloadTypeOk() (*string, bool)`

GetSoftwareDownloadTypeOk returns a tuple with the SoftwareDownloadType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareDownloadType

`func (o *ApplianceUpgradePolicyAllOf) SetSoftwareDownloadType(v string)`

SetSoftwareDownloadType sets SoftwareDownloadType field to given value.

### HasSoftwareDownloadType

`func (o *ApplianceUpgradePolicyAllOf) HasSoftwareDownloadType() bool`

HasSoftwareDownloadType returns a boolean if a field has been set.

### GetAccount

`func (o *ApplianceUpgradePolicyAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ApplianceUpgradePolicyAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ApplianceUpgradePolicyAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ApplianceUpgradePolicyAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


