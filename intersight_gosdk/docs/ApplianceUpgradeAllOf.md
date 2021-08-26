# ApplianceUpgradeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.Upgrade"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.Upgrade"]
**Active** | Pointer to **bool** | Indicates if the software upgrade is active or not. | [optional] [readonly] 
**AutoCreated** | Pointer to **bool** | Indicates that the request was automatically created by the system. | [optional] [readonly] 
**CompletedPhases** | Pointer to [**[]OnpremUpgradePhase**](OnpremUpgradePhase.md) |  | [optional] 
**CurrentPhase** | Pointer to [**NullableOnpremUpgradePhase**](OnpremUpgradePhase.md) |  | [optional] 
**Description** | Pointer to **string** | Description of the software upgrade. | [optional] [readonly] 
**ElapsedTime** | Pointer to **int64** | Elapsed time in seconds during the software upgrade. | [optional] [readonly] 
**EndTime** | Pointer to **time.Time** | End date of the software upgrade. | [optional] [readonly] 
**ErrorCode** | Pointer to **int64** | Error code for Intersight Appliance&#39;s software upgrade. In case of failure - this code will help decide if software upgrade can be retried. | [optional] [readonly] 
**Fingerprint** | Pointer to **string** | Software upgrade manifest&#39;s fingerprint. | [optional] [readonly] 
**IsRollingBack** | Pointer to **bool** | Track if software upgrade is upgrading or rolling back. | [optional] [readonly] [default to false]
**IsUserTriggered** | Pointer to **bool** | Indicates if the upgrade is triggered by user or due to schedule. | [optional] [readonly] [default to false]
**Messages** | Pointer to **[]string** |  | [optional] 
**RollbackNeeded** | Pointer to **bool** | Track if rollback is needed. | [optional] [default to false]
**RollbackPhases** | Pointer to [**[]OnpremUpgradePhase**](OnpremUpgradePhase.md) |  | [optional] 
**RollbackStatus** | Pointer to **string** | Status of the Intersight Appliance&#39;s software rollback status. | [optional] [readonly] 
**Services** | Pointer to **[]string** |  | [optional] 
**StartTime** | Pointer to **time.Time** | Start date of the software upgrade. UI can modify startTime to re-schedule an upgrade. | [optional] 
**Status** | Pointer to **string** | Status of the Intersight Appliance&#39;s software upgrade. | [optional] [readonly] 
**TotalPhases** | Pointer to **int64** | TotalPhase represents the total number of the upgradePhases for one upgrade. | [optional] [readonly] 
**UiPackages** | Pointer to **[]string** |  | [optional] 
**Version** | Pointer to **string** | Software upgrade manifest&#39;s version. | [optional] [readonly] 
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**ImageBundle** | Pointer to [**ApplianceImageBundleRelationship**](ApplianceImageBundleRelationship.md) |  | [optional] 

## Methods

### NewApplianceUpgradeAllOf

`func NewApplianceUpgradeAllOf(classId string, objectType string, ) *ApplianceUpgradeAllOf`

NewApplianceUpgradeAllOf instantiates a new ApplianceUpgradeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceUpgradeAllOfWithDefaults

`func NewApplianceUpgradeAllOfWithDefaults() *ApplianceUpgradeAllOf`

NewApplianceUpgradeAllOfWithDefaults instantiates a new ApplianceUpgradeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceUpgradeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceUpgradeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceUpgradeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceUpgradeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceUpgradeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceUpgradeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetActive

`func (o *ApplianceUpgradeAllOf) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ApplianceUpgradeAllOf) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ApplianceUpgradeAllOf) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ApplianceUpgradeAllOf) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAutoCreated

`func (o *ApplianceUpgradeAllOf) GetAutoCreated() bool`

GetAutoCreated returns the AutoCreated field if non-nil, zero value otherwise.

### GetAutoCreatedOk

`func (o *ApplianceUpgradeAllOf) GetAutoCreatedOk() (*bool, bool)`

GetAutoCreatedOk returns a tuple with the AutoCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreated

`func (o *ApplianceUpgradeAllOf) SetAutoCreated(v bool)`

SetAutoCreated sets AutoCreated field to given value.

### HasAutoCreated

`func (o *ApplianceUpgradeAllOf) HasAutoCreated() bool`

HasAutoCreated returns a boolean if a field has been set.

### GetCompletedPhases

`func (o *ApplianceUpgradeAllOf) GetCompletedPhases() []OnpremUpgradePhase`

GetCompletedPhases returns the CompletedPhases field if non-nil, zero value otherwise.

### GetCompletedPhasesOk

`func (o *ApplianceUpgradeAllOf) GetCompletedPhasesOk() (*[]OnpremUpgradePhase, bool)`

GetCompletedPhasesOk returns a tuple with the CompletedPhases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedPhases

`func (o *ApplianceUpgradeAllOf) SetCompletedPhases(v []OnpremUpgradePhase)`

SetCompletedPhases sets CompletedPhases field to given value.

### HasCompletedPhases

`func (o *ApplianceUpgradeAllOf) HasCompletedPhases() bool`

HasCompletedPhases returns a boolean if a field has been set.

### SetCompletedPhasesNil

`func (o *ApplianceUpgradeAllOf) SetCompletedPhasesNil(b bool)`

 SetCompletedPhasesNil sets the value for CompletedPhases to be an explicit nil

### UnsetCompletedPhases
`func (o *ApplianceUpgradeAllOf) UnsetCompletedPhases()`

UnsetCompletedPhases ensures that no value is present for CompletedPhases, not even an explicit nil
### GetCurrentPhase

`func (o *ApplianceUpgradeAllOf) GetCurrentPhase() OnpremUpgradePhase`

GetCurrentPhase returns the CurrentPhase field if non-nil, zero value otherwise.

### GetCurrentPhaseOk

`func (o *ApplianceUpgradeAllOf) GetCurrentPhaseOk() (*OnpremUpgradePhase, bool)`

GetCurrentPhaseOk returns a tuple with the CurrentPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPhase

`func (o *ApplianceUpgradeAllOf) SetCurrentPhase(v OnpremUpgradePhase)`

SetCurrentPhase sets CurrentPhase field to given value.

### HasCurrentPhase

`func (o *ApplianceUpgradeAllOf) HasCurrentPhase() bool`

HasCurrentPhase returns a boolean if a field has been set.

### SetCurrentPhaseNil

`func (o *ApplianceUpgradeAllOf) SetCurrentPhaseNil(b bool)`

 SetCurrentPhaseNil sets the value for CurrentPhase to be an explicit nil

### UnsetCurrentPhase
`func (o *ApplianceUpgradeAllOf) UnsetCurrentPhase()`

UnsetCurrentPhase ensures that no value is present for CurrentPhase, not even an explicit nil
### GetDescription

`func (o *ApplianceUpgradeAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplianceUpgradeAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplianceUpgradeAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplianceUpgradeAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetElapsedTime

`func (o *ApplianceUpgradeAllOf) GetElapsedTime() int64`

GetElapsedTime returns the ElapsedTime field if non-nil, zero value otherwise.

### GetElapsedTimeOk

`func (o *ApplianceUpgradeAllOf) GetElapsedTimeOk() (*int64, bool)`

GetElapsedTimeOk returns a tuple with the ElapsedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsedTime

`func (o *ApplianceUpgradeAllOf) SetElapsedTime(v int64)`

SetElapsedTime sets ElapsedTime field to given value.

### HasElapsedTime

`func (o *ApplianceUpgradeAllOf) HasElapsedTime() bool`

HasElapsedTime returns a boolean if a field has been set.

### GetEndTime

`func (o *ApplianceUpgradeAllOf) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ApplianceUpgradeAllOf) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ApplianceUpgradeAllOf) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ApplianceUpgradeAllOf) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetErrorCode

`func (o *ApplianceUpgradeAllOf) GetErrorCode() int64`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ApplianceUpgradeAllOf) GetErrorCodeOk() (*int64, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ApplianceUpgradeAllOf) SetErrorCode(v int64)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ApplianceUpgradeAllOf) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetFingerprint

`func (o *ApplianceUpgradeAllOf) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *ApplianceUpgradeAllOf) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *ApplianceUpgradeAllOf) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *ApplianceUpgradeAllOf) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetIsRollingBack

`func (o *ApplianceUpgradeAllOf) GetIsRollingBack() bool`

GetIsRollingBack returns the IsRollingBack field if non-nil, zero value otherwise.

### GetIsRollingBackOk

`func (o *ApplianceUpgradeAllOf) GetIsRollingBackOk() (*bool, bool)`

GetIsRollingBackOk returns a tuple with the IsRollingBack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRollingBack

`func (o *ApplianceUpgradeAllOf) SetIsRollingBack(v bool)`

SetIsRollingBack sets IsRollingBack field to given value.

### HasIsRollingBack

`func (o *ApplianceUpgradeAllOf) HasIsRollingBack() bool`

HasIsRollingBack returns a boolean if a field has been set.

### GetIsUserTriggered

`func (o *ApplianceUpgradeAllOf) GetIsUserTriggered() bool`

GetIsUserTriggered returns the IsUserTriggered field if non-nil, zero value otherwise.

### GetIsUserTriggeredOk

`func (o *ApplianceUpgradeAllOf) GetIsUserTriggeredOk() (*bool, bool)`

GetIsUserTriggeredOk returns a tuple with the IsUserTriggered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUserTriggered

`func (o *ApplianceUpgradeAllOf) SetIsUserTriggered(v bool)`

SetIsUserTriggered sets IsUserTriggered field to given value.

### HasIsUserTriggered

`func (o *ApplianceUpgradeAllOf) HasIsUserTriggered() bool`

HasIsUserTriggered returns a boolean if a field has been set.

### GetMessages

`func (o *ApplianceUpgradeAllOf) GetMessages() []string`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ApplianceUpgradeAllOf) GetMessagesOk() (*[]string, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ApplianceUpgradeAllOf) SetMessages(v []string)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *ApplianceUpgradeAllOf) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### SetMessagesNil

`func (o *ApplianceUpgradeAllOf) SetMessagesNil(b bool)`

 SetMessagesNil sets the value for Messages to be an explicit nil

### UnsetMessages
`func (o *ApplianceUpgradeAllOf) UnsetMessages()`

UnsetMessages ensures that no value is present for Messages, not even an explicit nil
### GetRollbackNeeded

`func (o *ApplianceUpgradeAllOf) GetRollbackNeeded() bool`

GetRollbackNeeded returns the RollbackNeeded field if non-nil, zero value otherwise.

### GetRollbackNeededOk

`func (o *ApplianceUpgradeAllOf) GetRollbackNeededOk() (*bool, bool)`

GetRollbackNeededOk returns a tuple with the RollbackNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollbackNeeded

`func (o *ApplianceUpgradeAllOf) SetRollbackNeeded(v bool)`

SetRollbackNeeded sets RollbackNeeded field to given value.

### HasRollbackNeeded

`func (o *ApplianceUpgradeAllOf) HasRollbackNeeded() bool`

HasRollbackNeeded returns a boolean if a field has been set.

### GetRollbackPhases

`func (o *ApplianceUpgradeAllOf) GetRollbackPhases() []OnpremUpgradePhase`

GetRollbackPhases returns the RollbackPhases field if non-nil, zero value otherwise.

### GetRollbackPhasesOk

`func (o *ApplianceUpgradeAllOf) GetRollbackPhasesOk() (*[]OnpremUpgradePhase, bool)`

GetRollbackPhasesOk returns a tuple with the RollbackPhases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollbackPhases

`func (o *ApplianceUpgradeAllOf) SetRollbackPhases(v []OnpremUpgradePhase)`

SetRollbackPhases sets RollbackPhases field to given value.

### HasRollbackPhases

`func (o *ApplianceUpgradeAllOf) HasRollbackPhases() bool`

HasRollbackPhases returns a boolean if a field has been set.

### SetRollbackPhasesNil

`func (o *ApplianceUpgradeAllOf) SetRollbackPhasesNil(b bool)`

 SetRollbackPhasesNil sets the value for RollbackPhases to be an explicit nil

### UnsetRollbackPhases
`func (o *ApplianceUpgradeAllOf) UnsetRollbackPhases()`

UnsetRollbackPhases ensures that no value is present for RollbackPhases, not even an explicit nil
### GetRollbackStatus

`func (o *ApplianceUpgradeAllOf) GetRollbackStatus() string`

GetRollbackStatus returns the RollbackStatus field if non-nil, zero value otherwise.

### GetRollbackStatusOk

`func (o *ApplianceUpgradeAllOf) GetRollbackStatusOk() (*string, bool)`

GetRollbackStatusOk returns a tuple with the RollbackStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollbackStatus

`func (o *ApplianceUpgradeAllOf) SetRollbackStatus(v string)`

SetRollbackStatus sets RollbackStatus field to given value.

### HasRollbackStatus

`func (o *ApplianceUpgradeAllOf) HasRollbackStatus() bool`

HasRollbackStatus returns a boolean if a field has been set.

### GetServices

`func (o *ApplianceUpgradeAllOf) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ApplianceUpgradeAllOf) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ApplianceUpgradeAllOf) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *ApplianceUpgradeAllOf) HasServices() bool`

HasServices returns a boolean if a field has been set.

### SetServicesNil

`func (o *ApplianceUpgradeAllOf) SetServicesNil(b bool)`

 SetServicesNil sets the value for Services to be an explicit nil

### UnsetServices
`func (o *ApplianceUpgradeAllOf) UnsetServices()`

UnsetServices ensures that no value is present for Services, not even an explicit nil
### GetStartTime

`func (o *ApplianceUpgradeAllOf) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ApplianceUpgradeAllOf) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ApplianceUpgradeAllOf) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ApplianceUpgradeAllOf) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *ApplianceUpgradeAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApplianceUpgradeAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApplianceUpgradeAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApplianceUpgradeAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotalPhases

`func (o *ApplianceUpgradeAllOf) GetTotalPhases() int64`

GetTotalPhases returns the TotalPhases field if non-nil, zero value otherwise.

### GetTotalPhasesOk

`func (o *ApplianceUpgradeAllOf) GetTotalPhasesOk() (*int64, bool)`

GetTotalPhasesOk returns a tuple with the TotalPhases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPhases

`func (o *ApplianceUpgradeAllOf) SetTotalPhases(v int64)`

SetTotalPhases sets TotalPhases field to given value.

### HasTotalPhases

`func (o *ApplianceUpgradeAllOf) HasTotalPhases() bool`

HasTotalPhases returns a boolean if a field has been set.

### GetUiPackages

`func (o *ApplianceUpgradeAllOf) GetUiPackages() []string`

GetUiPackages returns the UiPackages field if non-nil, zero value otherwise.

### GetUiPackagesOk

`func (o *ApplianceUpgradeAllOf) GetUiPackagesOk() (*[]string, bool)`

GetUiPackagesOk returns a tuple with the UiPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiPackages

`func (o *ApplianceUpgradeAllOf) SetUiPackages(v []string)`

SetUiPackages sets UiPackages field to given value.

### HasUiPackages

`func (o *ApplianceUpgradeAllOf) HasUiPackages() bool`

HasUiPackages returns a boolean if a field has been set.

### SetUiPackagesNil

`func (o *ApplianceUpgradeAllOf) SetUiPackagesNil(b bool)`

 SetUiPackagesNil sets the value for UiPackages to be an explicit nil

### UnsetUiPackages
`func (o *ApplianceUpgradeAllOf) UnsetUiPackages()`

UnsetUiPackages ensures that no value is present for UiPackages, not even an explicit nil
### GetVersion

`func (o *ApplianceUpgradeAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApplianceUpgradeAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApplianceUpgradeAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ApplianceUpgradeAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAccount

`func (o *ApplianceUpgradeAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ApplianceUpgradeAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ApplianceUpgradeAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ApplianceUpgradeAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetImageBundle

`func (o *ApplianceUpgradeAllOf) GetImageBundle() ApplianceImageBundleRelationship`

GetImageBundle returns the ImageBundle field if non-nil, zero value otherwise.

### GetImageBundleOk

`func (o *ApplianceUpgradeAllOf) GetImageBundleOk() (*ApplianceImageBundleRelationship, bool)`

GetImageBundleOk returns a tuple with the ImageBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBundle

`func (o *ApplianceUpgradeAllOf) SetImageBundle(v ApplianceImageBundleRelationship)`

SetImageBundle sets ImageBundle field to given value.

### HasImageBundle

`func (o *ApplianceUpgradeAllOf) HasImageBundle() bool`

HasImageBundle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


