# ApplianceImageBundleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnsiblePackages** | Pointer to [**[]OnpremImagePackage**](onprem.ImagePackage.md) |  | [optional] 
**AutoUpgrade** | Pointer to **bool** | Indicates that the software upgrade was automatically initiated by the Intersight Appliance. | [optional] [readonly] 
**DcPackages** | Pointer to [**[]OnpremImagePackage**](onprem.ImagePackage.md) |  | [optional] 
**DebugPackages** | Pointer to [**[]OnpremImagePackage**](onprem.ImagePackage.md) |  | [optional] 
**Description** | Pointer to **string** | Short description of the software upgrade bundle. | [optional] [readonly] 
**EndpointPackages** | Pointer to [**[]OnpremImagePackage**](onprem.ImagePackage.md) |  | [optional] 
**Fingerprint** | Pointer to **string** | Fingerprint of the software manifest from which this bundle is created. Fingerprint is calculated using the SHA256 algorithm. | [optional] [readonly] 
**HasError** | Pointer to **bool** | Indicates that the ImageBundle has errors. The upgrade service sets this field when it encounters errors during the manifest processing. | [optional] [readonly] 
**InfraPackages** | Pointer to [**[]OnpremImagePackage**](onprem.ImagePackage.md) |  | [optional] 
**InitPackages** | Pointer to [**[]OnpremImagePackage**](onprem.ImagePackage.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the software upgrade bundle. | [optional] [readonly] 
**Notes** | Pointer to **string** | Detailed description of the software upgrade bundle. | [optional] [readonly] 
**Priority** | Pointer to **string** | Software upgrade manifest&#39;s upgrade priority. The upgrade service supports two priorities, Normal and Critical. Normal priority is used for regular software upgrades, and the upgrade service uses the Upgrade Policy to compute upgrade start time. Critical priority is used for the critical software security patches, and the upgrade service ignores the Upgrade Policy when it computes the upgrade start time. * &#x60;Normal&#x60; - Normal upgrade priority is used for all the software upgrades except for the critical security updates. The upgrade service of Intersight Appliance uses the Software Upgrade Policy settings to start the upgrade process. * &#x60;Critical&#x60; - Critical upgrade priority is used for critical updates such as security patches. The upgrade service of the Intersight Appliance starts the upgrade as specified by the upgrade properties in the software manifest file. The upgrade service will not use the settings specified in the Software Upgrade Policy. | [optional] [readonly] [default to "Normal"]
**ReleaseTime** | Pointer to [**time.Time**](time.Time.md) | Software upgrade manifest&#39;s release date and time. | [optional] [readonly] 
**ServicePackages** | Pointer to [**[]OnpremImagePackage**](onprem.ImagePackage.md) |  | [optional] 
**StatusMessage** | Pointer to **string** | Status message set during the manifest processing. | [optional] [readonly] 
**SystemPackages** | Pointer to [**[]OnpremImagePackage**](onprem.ImagePackage.md) |  | [optional] 
**UiPackages** | Pointer to [**[]OnpremImagePackage**](onprem.ImagePackage.md) |  | [optional] 
**UpgradeEndTime** | Pointer to [**time.Time**](time.Time.md) | End date of the software upgrade process. | [optional] [readonly] 
**UpgradeGracePeriod** | Pointer to **int64** | Grace period in seconds before the automatic upgrade is initiated. The upgrade service uses the grace period to compute the upgrade start time when it receives an upgrade notfication from the Intersight. If there is an Upgrade Policy configured for the Intersight Appliance, then the upgrade service uses the policy to compute the upgrade start time. However, the upgrade start time cannot not exceed the limit enforced by the grace period. | [optional] [readonly] 
**UpgradeImpactDuration** | Pointer to **int64** | Duration (in minutes) for which services will be disrupted. | [optional] [readonly] 
**UpgradeImpactEnum** | Pointer to **string** | UpgradeImpactEnum is used to indicate the kind of impact the upgrade has on currently running services on the appliance. * &#x60;None&#x60; - The upgrade has no effect on the system. * &#x60;Disruptive&#x60; - The services will not be functional during the upgrade. * &#x60;Disruptive-reboot&#x60; - The upgrade needs a reboot. | [optional] [readonly] [default to "None"]
**UpgradeStartTime** | Pointer to [**time.Time**](time.Time.md) | Start date of the software upgrade process. | [optional] [readonly] 
**Version** | Pointer to **string** | Software upgrade manifest&#39;s version. | [optional] [readonly] 

## Methods

### NewApplianceImageBundleAllOf

`func NewApplianceImageBundleAllOf() *ApplianceImageBundleAllOf`

NewApplianceImageBundleAllOf instantiates a new ApplianceImageBundleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceImageBundleAllOfWithDefaults

`func NewApplianceImageBundleAllOfWithDefaults() *ApplianceImageBundleAllOf`

NewApplianceImageBundleAllOfWithDefaults instantiates a new ApplianceImageBundleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnsiblePackages

`func (o *ApplianceImageBundleAllOf) GetAnsiblePackages() []OnpremImagePackage`

GetAnsiblePackages returns the AnsiblePackages field if non-nil, zero value otherwise.

### GetAnsiblePackagesOk

`func (o *ApplianceImageBundleAllOf) GetAnsiblePackagesOk() (*[]OnpremImagePackage, bool)`

GetAnsiblePackagesOk returns a tuple with the AnsiblePackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsiblePackages

`func (o *ApplianceImageBundleAllOf) SetAnsiblePackages(v []OnpremImagePackage)`

SetAnsiblePackages sets AnsiblePackages field to given value.

### HasAnsiblePackages

`func (o *ApplianceImageBundleAllOf) HasAnsiblePackages() bool`

HasAnsiblePackages returns a boolean if a field has been set.

### GetAutoUpgrade

`func (o *ApplianceImageBundleAllOf) GetAutoUpgrade() bool`

GetAutoUpgrade returns the AutoUpgrade field if non-nil, zero value otherwise.

### GetAutoUpgradeOk

`func (o *ApplianceImageBundleAllOf) GetAutoUpgradeOk() (*bool, bool)`

GetAutoUpgradeOk returns a tuple with the AutoUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpgrade

`func (o *ApplianceImageBundleAllOf) SetAutoUpgrade(v bool)`

SetAutoUpgrade sets AutoUpgrade field to given value.

### HasAutoUpgrade

`func (o *ApplianceImageBundleAllOf) HasAutoUpgrade() bool`

HasAutoUpgrade returns a boolean if a field has been set.

### GetDcPackages

`func (o *ApplianceImageBundleAllOf) GetDcPackages() []OnpremImagePackage`

GetDcPackages returns the DcPackages field if non-nil, zero value otherwise.

### GetDcPackagesOk

`func (o *ApplianceImageBundleAllOf) GetDcPackagesOk() (*[]OnpremImagePackage, bool)`

GetDcPackagesOk returns a tuple with the DcPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcPackages

`func (o *ApplianceImageBundleAllOf) SetDcPackages(v []OnpremImagePackage)`

SetDcPackages sets DcPackages field to given value.

### HasDcPackages

`func (o *ApplianceImageBundleAllOf) HasDcPackages() bool`

HasDcPackages returns a boolean if a field has been set.

### GetDebugPackages

`func (o *ApplianceImageBundleAllOf) GetDebugPackages() []OnpremImagePackage`

GetDebugPackages returns the DebugPackages field if non-nil, zero value otherwise.

### GetDebugPackagesOk

`func (o *ApplianceImageBundleAllOf) GetDebugPackagesOk() (*[]OnpremImagePackage, bool)`

GetDebugPackagesOk returns a tuple with the DebugPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugPackages

`func (o *ApplianceImageBundleAllOf) SetDebugPackages(v []OnpremImagePackage)`

SetDebugPackages sets DebugPackages field to given value.

### HasDebugPackages

`func (o *ApplianceImageBundleAllOf) HasDebugPackages() bool`

HasDebugPackages returns a boolean if a field has been set.

### GetDescription

`func (o *ApplianceImageBundleAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplianceImageBundleAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplianceImageBundleAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplianceImageBundleAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndpointPackages

`func (o *ApplianceImageBundleAllOf) GetEndpointPackages() []OnpremImagePackage`

GetEndpointPackages returns the EndpointPackages field if non-nil, zero value otherwise.

### GetEndpointPackagesOk

`func (o *ApplianceImageBundleAllOf) GetEndpointPackagesOk() (*[]OnpremImagePackage, bool)`

GetEndpointPackagesOk returns a tuple with the EndpointPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointPackages

`func (o *ApplianceImageBundleAllOf) SetEndpointPackages(v []OnpremImagePackage)`

SetEndpointPackages sets EndpointPackages field to given value.

### HasEndpointPackages

`func (o *ApplianceImageBundleAllOf) HasEndpointPackages() bool`

HasEndpointPackages returns a boolean if a field has been set.

### GetFingerprint

`func (o *ApplianceImageBundleAllOf) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *ApplianceImageBundleAllOf) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *ApplianceImageBundleAllOf) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *ApplianceImageBundleAllOf) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetHasError

`func (o *ApplianceImageBundleAllOf) GetHasError() bool`

GetHasError returns the HasError field if non-nil, zero value otherwise.

### GetHasErrorOk

`func (o *ApplianceImageBundleAllOf) GetHasErrorOk() (*bool, bool)`

GetHasErrorOk returns a tuple with the HasError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasError

`func (o *ApplianceImageBundleAllOf) SetHasError(v bool)`

SetHasError sets HasError field to given value.

### HasHasError

`func (o *ApplianceImageBundleAllOf) HasHasError() bool`

HasHasError returns a boolean if a field has been set.

### GetInfraPackages

`func (o *ApplianceImageBundleAllOf) GetInfraPackages() []OnpremImagePackage`

GetInfraPackages returns the InfraPackages field if non-nil, zero value otherwise.

### GetInfraPackagesOk

`func (o *ApplianceImageBundleAllOf) GetInfraPackagesOk() (*[]OnpremImagePackage, bool)`

GetInfraPackagesOk returns a tuple with the InfraPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraPackages

`func (o *ApplianceImageBundleAllOf) SetInfraPackages(v []OnpremImagePackage)`

SetInfraPackages sets InfraPackages field to given value.

### HasInfraPackages

`func (o *ApplianceImageBundleAllOf) HasInfraPackages() bool`

HasInfraPackages returns a boolean if a field has been set.

### GetInitPackages

`func (o *ApplianceImageBundleAllOf) GetInitPackages() []OnpremImagePackage`

GetInitPackages returns the InitPackages field if non-nil, zero value otherwise.

### GetInitPackagesOk

`func (o *ApplianceImageBundleAllOf) GetInitPackagesOk() (*[]OnpremImagePackage, bool)`

GetInitPackagesOk returns a tuple with the InitPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitPackages

`func (o *ApplianceImageBundleAllOf) SetInitPackages(v []OnpremImagePackage)`

SetInitPackages sets InitPackages field to given value.

### HasInitPackages

`func (o *ApplianceImageBundleAllOf) HasInitPackages() bool`

HasInitPackages returns a boolean if a field has been set.

### GetName

`func (o *ApplianceImageBundleAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplianceImageBundleAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplianceImageBundleAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplianceImageBundleAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotes

`func (o *ApplianceImageBundleAllOf) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ApplianceImageBundleAllOf) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ApplianceImageBundleAllOf) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ApplianceImageBundleAllOf) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetPriority

`func (o *ApplianceImageBundleAllOf) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ApplianceImageBundleAllOf) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ApplianceImageBundleAllOf) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ApplianceImageBundleAllOf) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetReleaseTime

`func (o *ApplianceImageBundleAllOf) GetReleaseTime() time.Time`

GetReleaseTime returns the ReleaseTime field if non-nil, zero value otherwise.

### GetReleaseTimeOk

`func (o *ApplianceImageBundleAllOf) GetReleaseTimeOk() (*time.Time, bool)`

GetReleaseTimeOk returns a tuple with the ReleaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseTime

`func (o *ApplianceImageBundleAllOf) SetReleaseTime(v time.Time)`

SetReleaseTime sets ReleaseTime field to given value.

### HasReleaseTime

`func (o *ApplianceImageBundleAllOf) HasReleaseTime() bool`

HasReleaseTime returns a boolean if a field has been set.

### GetServicePackages

`func (o *ApplianceImageBundleAllOf) GetServicePackages() []OnpremImagePackage`

GetServicePackages returns the ServicePackages field if non-nil, zero value otherwise.

### GetServicePackagesOk

`func (o *ApplianceImageBundleAllOf) GetServicePackagesOk() (*[]OnpremImagePackage, bool)`

GetServicePackagesOk returns a tuple with the ServicePackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePackages

`func (o *ApplianceImageBundleAllOf) SetServicePackages(v []OnpremImagePackage)`

SetServicePackages sets ServicePackages field to given value.

### HasServicePackages

`func (o *ApplianceImageBundleAllOf) HasServicePackages() bool`

HasServicePackages returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ApplianceImageBundleAllOf) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ApplianceImageBundleAllOf) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ApplianceImageBundleAllOf) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ApplianceImageBundleAllOf) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetSystemPackages

`func (o *ApplianceImageBundleAllOf) GetSystemPackages() []OnpremImagePackage`

GetSystemPackages returns the SystemPackages field if non-nil, zero value otherwise.

### GetSystemPackagesOk

`func (o *ApplianceImageBundleAllOf) GetSystemPackagesOk() (*[]OnpremImagePackage, bool)`

GetSystemPackagesOk returns a tuple with the SystemPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPackages

`func (o *ApplianceImageBundleAllOf) SetSystemPackages(v []OnpremImagePackage)`

SetSystemPackages sets SystemPackages field to given value.

### HasSystemPackages

`func (o *ApplianceImageBundleAllOf) HasSystemPackages() bool`

HasSystemPackages returns a boolean if a field has been set.

### GetUiPackages

`func (o *ApplianceImageBundleAllOf) GetUiPackages() []OnpremImagePackage`

GetUiPackages returns the UiPackages field if non-nil, zero value otherwise.

### GetUiPackagesOk

`func (o *ApplianceImageBundleAllOf) GetUiPackagesOk() (*[]OnpremImagePackage, bool)`

GetUiPackagesOk returns a tuple with the UiPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiPackages

`func (o *ApplianceImageBundleAllOf) SetUiPackages(v []OnpremImagePackage)`

SetUiPackages sets UiPackages field to given value.

### HasUiPackages

`func (o *ApplianceImageBundleAllOf) HasUiPackages() bool`

HasUiPackages returns a boolean if a field has been set.

### GetUpgradeEndTime

`func (o *ApplianceImageBundleAllOf) GetUpgradeEndTime() time.Time`

GetUpgradeEndTime returns the UpgradeEndTime field if non-nil, zero value otherwise.

### GetUpgradeEndTimeOk

`func (o *ApplianceImageBundleAllOf) GetUpgradeEndTimeOk() (*time.Time, bool)`

GetUpgradeEndTimeOk returns a tuple with the UpgradeEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeEndTime

`func (o *ApplianceImageBundleAllOf) SetUpgradeEndTime(v time.Time)`

SetUpgradeEndTime sets UpgradeEndTime field to given value.

### HasUpgradeEndTime

`func (o *ApplianceImageBundleAllOf) HasUpgradeEndTime() bool`

HasUpgradeEndTime returns a boolean if a field has been set.

### GetUpgradeGracePeriod

`func (o *ApplianceImageBundleAllOf) GetUpgradeGracePeriod() int64`

GetUpgradeGracePeriod returns the UpgradeGracePeriod field if non-nil, zero value otherwise.

### GetUpgradeGracePeriodOk

`func (o *ApplianceImageBundleAllOf) GetUpgradeGracePeriodOk() (*int64, bool)`

GetUpgradeGracePeriodOk returns a tuple with the UpgradeGracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeGracePeriod

`func (o *ApplianceImageBundleAllOf) SetUpgradeGracePeriod(v int64)`

SetUpgradeGracePeriod sets UpgradeGracePeriod field to given value.

### HasUpgradeGracePeriod

`func (o *ApplianceImageBundleAllOf) HasUpgradeGracePeriod() bool`

HasUpgradeGracePeriod returns a boolean if a field has been set.

### GetUpgradeImpactDuration

`func (o *ApplianceImageBundleAllOf) GetUpgradeImpactDuration() int64`

GetUpgradeImpactDuration returns the UpgradeImpactDuration field if non-nil, zero value otherwise.

### GetUpgradeImpactDurationOk

`func (o *ApplianceImageBundleAllOf) GetUpgradeImpactDurationOk() (*int64, bool)`

GetUpgradeImpactDurationOk returns a tuple with the UpgradeImpactDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeImpactDuration

`func (o *ApplianceImageBundleAllOf) SetUpgradeImpactDuration(v int64)`

SetUpgradeImpactDuration sets UpgradeImpactDuration field to given value.

### HasUpgradeImpactDuration

`func (o *ApplianceImageBundleAllOf) HasUpgradeImpactDuration() bool`

HasUpgradeImpactDuration returns a boolean if a field has been set.

### GetUpgradeImpactEnum

`func (o *ApplianceImageBundleAllOf) GetUpgradeImpactEnum() string`

GetUpgradeImpactEnum returns the UpgradeImpactEnum field if non-nil, zero value otherwise.

### GetUpgradeImpactEnumOk

`func (o *ApplianceImageBundleAllOf) GetUpgradeImpactEnumOk() (*string, bool)`

GetUpgradeImpactEnumOk returns a tuple with the UpgradeImpactEnum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeImpactEnum

`func (o *ApplianceImageBundleAllOf) SetUpgradeImpactEnum(v string)`

SetUpgradeImpactEnum sets UpgradeImpactEnum field to given value.

### HasUpgradeImpactEnum

`func (o *ApplianceImageBundleAllOf) HasUpgradeImpactEnum() bool`

HasUpgradeImpactEnum returns a boolean if a field has been set.

### GetUpgradeStartTime

`func (o *ApplianceImageBundleAllOf) GetUpgradeStartTime() time.Time`

GetUpgradeStartTime returns the UpgradeStartTime field if non-nil, zero value otherwise.

### GetUpgradeStartTimeOk

`func (o *ApplianceImageBundleAllOf) GetUpgradeStartTimeOk() (*time.Time, bool)`

GetUpgradeStartTimeOk returns a tuple with the UpgradeStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeStartTime

`func (o *ApplianceImageBundleAllOf) SetUpgradeStartTime(v time.Time)`

SetUpgradeStartTime sets UpgradeStartTime field to given value.

### HasUpgradeStartTime

`func (o *ApplianceImageBundleAllOf) HasUpgradeStartTime() bool`

HasUpgradeStartTime returns a boolean if a field has been set.

### GetVersion

`func (o *ApplianceImageBundleAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApplianceImageBundleAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApplianceImageBundleAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ApplianceImageBundleAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


