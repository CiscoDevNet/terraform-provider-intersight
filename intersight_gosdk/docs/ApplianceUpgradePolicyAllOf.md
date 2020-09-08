# ApplianceUpgradePolicyAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoUpgrade** | Pointer to **bool** | Indicates if the upgrade service is set to automatically start the software upgrade or not. If autoUpgrade is true, then the value of the schedule field is ignored. | [optional] 
**BlackoutDatesEnabled** | Pointer to **bool** | If enabled, allows the user to define a blackout period during which the appliance will not be upgraded. | [optional] 
**BlackoutEndDate** | Pointer to [**time.Time**](time.Time.md) | End date of the black out period. | [optional] 
**BlackoutStartDate** | Pointer to [**time.Time**](time.Time.md) | Start date of the black out period. The appliance will not be upgraded during this period. | [optional] 
**Schedule** | Pointer to [**OnpremSchedule**](onprem.Schedule.md) |  | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 

## Methods

### NewApplianceUpgradePolicyAllOf

`func NewApplianceUpgradePolicyAllOf() *ApplianceUpgradePolicyAllOf`

NewApplianceUpgradePolicyAllOf instantiates a new ApplianceUpgradePolicyAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceUpgradePolicyAllOfWithDefaults

`func NewApplianceUpgradePolicyAllOfWithDefaults() *ApplianceUpgradePolicyAllOf`

NewApplianceUpgradePolicyAllOfWithDefaults instantiates a new ApplianceUpgradePolicyAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


