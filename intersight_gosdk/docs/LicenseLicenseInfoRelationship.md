# LicenseLicenseInfoRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountMoid** | Pointer to **string** | The Account ID for this managed object. | [optional] [readonly] 
**ClassId** | **string** | The concrete type of this complex type. Its value must be the same as the &#39;objectType&#39; property. The OpenAPI document references this property as a discriminator value. | [readonly] 
**CreateTime** | Pointer to [**time.Time**](time.Time.md) | The time when this managed object was created. | [optional] [readonly] 
**DomainGroupMoid** | Pointer to **string** | The DomainGroup ID for this managed object. | [optional] [readonly] 
**ModTime** | Pointer to [**time.Time**](time.Time.md) | The time when this managed object was last modified. | [optional] [readonly] 
**Moid** | Pointer to **string** | The unique identifier of this Managed Object instance. | [optional] 
**ObjectType** | **string** | The fully-qualified type of this managed object, i.e. the class name. This property is optional. The ObjectType is implied from the URL path. If specified, the value of objectType must match the class name specified in the URL path. | [readonly] 
**Owners** | Pointer to **[]string** |  | [optional] 
**SharedScope** | Pointer to **string** | Intersight provides pre-built workflows, tasks and policies to end users through global catalogs. Objects that are made available through global catalogs are said to have a &#39;shared&#39; ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs. | [optional] [readonly] 
**Tags** | Pointer to [**[]MoTag**](mo.Tag.md) |  | [optional] 
**VersionContext** | Pointer to [**MoVersionContext**](mo.VersionContext.md) |  | [optional] 
**Ancestors** | Pointer to [**[]MoBaseMoRelationship**](mo.BaseMo.Relationship.md) | An array of relationships to moBaseMo resources. | [optional] [readonly] 
**Parent** | Pointer to [**MoBaseMoRelationship**](mo.BaseMo.Relationship.md) |  | [optional] 
**PermissionResources** | Pointer to [**[]MoBaseMoRelationship**](mo.BaseMo.Relationship.md) | An array of relationships to moBaseMo resources. | [optional] [readonly] 
**DisplayNames** | Pointer to [**map[string][]string**](array.md) | A set of display names for the MO resource. These names are calculated based on other properties of the MO and potentially properties of Ancestor MOs. Displaynames are intended as a way to provide a normalized user appropriate name for an MO, especially for MOs which do not have a &#39;Name&#39; property, which is the case for much of the inventory discovered from managed targets. There are a limited number of keys, currently &#39;short&#39; and &#39;hierarchical&#39;. The value is an array and clients should use the first element of the array. | [optional] [readonly] 
**ActiveAdmin** | Pointer to **bool** | The license administrative state. Set this property to &#39;true&#39; to activate the license entitlements. | [optional] [readonly] 
**DaysLeft** | Pointer to **int64** | The number of days left for licenseState to stay in TrialPeriod or OutOfCompliance state. | [optional] [readonly] 
**EndTime** | Pointer to [**time.Time**](time.Time.md) | The date and time when the trial period expires. The value of the &#39;endTime&#39; property is set when the account enters the TrialPeriod or OutOfCompliance state. | [optional] [readonly] 
**EnforceMode** | Pointer to **string** | The entitlement mode reported by Cisco Smart Software Manager. | [optional] [readonly] 
**ErrorDesc** | Pointer to **string** | The detailed error message when there is any error related to this licensing entitlement. | [optional] [readonly] 
**EvaluationPeriod** | Pointer to **int64** | The default Trial or Grace period customer is entitled to. | [optional] 
**ExtraEvaluation** | Pointer to **int64** | The number of days the trial Trial or Grace period is extended. The trial or grace period can be extended once. | [optional] 
**LicenseCount** | Pointer to **int64** | The total number of devices claimed in the Intersight account. | [optional] [readonly] 
**LicenseState** | Pointer to **string** | The license state defined by Intersight. The value may be one of NotLicensed, TrialPeriod, OutOfCompliance, Compliance, GraceExpired, or TrialExpired. * &#x60;NotLicensed&#x60; - The license token is neither activated nor registered. * &#x60;GraceExpired&#x60; - The license grace period has expired. * &#x60;TrialPeriod&#x60; - The 90 days of trial period. * &#x60;OutOfCompliance&#x60; - The license is out of compliance. * &#x60;Compliance&#x60; - The license is in compliance. * &#x60;TrialExpired&#x60; - The trial period of 90 days has expired. | [optional] [readonly] [default to "NotLicensed"]
**LicenseType** | Pointer to **string** | The name of the Intersight license entitlement. For example, this property may be set to &#39;Essential&#39;. * &#x60;Base&#x60; - Base as a License type. It is default license type. * &#x60;Essential&#x60; - Essential as a License type. * &#x60;Standard&#x60; - Standard as a License type. * &#x60;Advantage&#x60; - Advantage as a License type. * &#x60;Premier&#x60; - Premier as a License type. | [optional] [readonly] [default to "Base"]
**StartTime** | Pointer to [**time.Time**](time.Time.md) | The date and time when the licenseState entered the TrialPeriod or OutOfCompliance state. | [optional] [readonly] 
**TrialAdmin** | Pointer to **bool** | The administrative state of the trial license. When the LicenseState is set to &#39;NotLicensed&#39;, &#39;trialAdmin&#39; can be set to true to start the trial period, i.e. licenseState is set to be TrialPeriod. | [optional] [readonly] 
**AccountLicenseData** | Pointer to [**LicenseAccountLicenseDataRelationship**](license.AccountLicenseData.Relationship.md) |  | [optional] 

## Methods

### NewLicenseLicenseInfoRelationship

`func NewLicenseLicenseInfoRelationship(classId string, objectType string, ) *LicenseLicenseInfoRelationship`

NewLicenseLicenseInfoRelationship instantiates a new LicenseLicenseInfoRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseLicenseInfoRelationshipWithDefaults

`func NewLicenseLicenseInfoRelationshipWithDefaults() *LicenseLicenseInfoRelationship`

NewLicenseLicenseInfoRelationshipWithDefaults instantiates a new LicenseLicenseInfoRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *LicenseLicenseInfoRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *LicenseLicenseInfoRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *LicenseLicenseInfoRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *LicenseLicenseInfoRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *LicenseLicenseInfoRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *LicenseLicenseInfoRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *LicenseLicenseInfoRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *LicenseLicenseInfoRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *LicenseLicenseInfoRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *LicenseLicenseInfoRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *LicenseLicenseInfoRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *LicenseLicenseInfoRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *LicenseLicenseInfoRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *LicenseLicenseInfoRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *LicenseLicenseInfoRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *LicenseLicenseInfoRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *LicenseLicenseInfoRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *LicenseLicenseInfoRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *LicenseLicenseInfoRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *LicenseLicenseInfoRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *LicenseLicenseInfoRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *LicenseLicenseInfoRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *LicenseLicenseInfoRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *LicenseLicenseInfoRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *LicenseLicenseInfoRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *LicenseLicenseInfoRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *LicenseLicenseInfoRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *LicenseLicenseInfoRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *LicenseLicenseInfoRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *LicenseLicenseInfoRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *LicenseLicenseInfoRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *LicenseLicenseInfoRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *LicenseLicenseInfoRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *LicenseLicenseInfoRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *LicenseLicenseInfoRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LicenseLicenseInfoRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LicenseLicenseInfoRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *LicenseLicenseInfoRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *LicenseLicenseInfoRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *LicenseLicenseInfoRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *LicenseLicenseInfoRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *LicenseLicenseInfoRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *LicenseLicenseInfoRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *LicenseLicenseInfoRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *LicenseLicenseInfoRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *LicenseLicenseInfoRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *LicenseLicenseInfoRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *LicenseLicenseInfoRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *LicenseLicenseInfoRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *LicenseLicenseInfoRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *LicenseLicenseInfoRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *LicenseLicenseInfoRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *LicenseLicenseInfoRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *LicenseLicenseInfoRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *LicenseLicenseInfoRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *LicenseLicenseInfoRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *LicenseLicenseInfoRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *LicenseLicenseInfoRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *LicenseLicenseInfoRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *LicenseLicenseInfoRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *LicenseLicenseInfoRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *LicenseLicenseInfoRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *LicenseLicenseInfoRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *LicenseLicenseInfoRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetActiveAdmin

`func (o *LicenseLicenseInfoRelationship) GetActiveAdmin() bool`

GetActiveAdmin returns the ActiveAdmin field if non-nil, zero value otherwise.

### GetActiveAdminOk

`func (o *LicenseLicenseInfoRelationship) GetActiveAdminOk() (*bool, bool)`

GetActiveAdminOk returns a tuple with the ActiveAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveAdmin

`func (o *LicenseLicenseInfoRelationship) SetActiveAdmin(v bool)`

SetActiveAdmin sets ActiveAdmin field to given value.

### HasActiveAdmin

`func (o *LicenseLicenseInfoRelationship) HasActiveAdmin() bool`

HasActiveAdmin returns a boolean if a field has been set.

### GetDaysLeft

`func (o *LicenseLicenseInfoRelationship) GetDaysLeft() int64`

GetDaysLeft returns the DaysLeft field if non-nil, zero value otherwise.

### GetDaysLeftOk

`func (o *LicenseLicenseInfoRelationship) GetDaysLeftOk() (*int64, bool)`

GetDaysLeftOk returns a tuple with the DaysLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysLeft

`func (o *LicenseLicenseInfoRelationship) SetDaysLeft(v int64)`

SetDaysLeft sets DaysLeft field to given value.

### HasDaysLeft

`func (o *LicenseLicenseInfoRelationship) HasDaysLeft() bool`

HasDaysLeft returns a boolean if a field has been set.

### GetEndTime

`func (o *LicenseLicenseInfoRelationship) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *LicenseLicenseInfoRelationship) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *LicenseLicenseInfoRelationship) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *LicenseLicenseInfoRelationship) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetEnforceMode

`func (o *LicenseLicenseInfoRelationship) GetEnforceMode() string`

GetEnforceMode returns the EnforceMode field if non-nil, zero value otherwise.

### GetEnforceModeOk

`func (o *LicenseLicenseInfoRelationship) GetEnforceModeOk() (*string, bool)`

GetEnforceModeOk returns a tuple with the EnforceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceMode

`func (o *LicenseLicenseInfoRelationship) SetEnforceMode(v string)`

SetEnforceMode sets EnforceMode field to given value.

### HasEnforceMode

`func (o *LicenseLicenseInfoRelationship) HasEnforceMode() bool`

HasEnforceMode returns a boolean if a field has been set.

### GetErrorDesc

`func (o *LicenseLicenseInfoRelationship) GetErrorDesc() string`

GetErrorDesc returns the ErrorDesc field if non-nil, zero value otherwise.

### GetErrorDescOk

`func (o *LicenseLicenseInfoRelationship) GetErrorDescOk() (*string, bool)`

GetErrorDescOk returns a tuple with the ErrorDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDesc

`func (o *LicenseLicenseInfoRelationship) SetErrorDesc(v string)`

SetErrorDesc sets ErrorDesc field to given value.

### HasErrorDesc

`func (o *LicenseLicenseInfoRelationship) HasErrorDesc() bool`

HasErrorDesc returns a boolean if a field has been set.

### GetEvaluationPeriod

`func (o *LicenseLicenseInfoRelationship) GetEvaluationPeriod() int64`

GetEvaluationPeriod returns the EvaluationPeriod field if non-nil, zero value otherwise.

### GetEvaluationPeriodOk

`func (o *LicenseLicenseInfoRelationship) GetEvaluationPeriodOk() (*int64, bool)`

GetEvaluationPeriodOk returns a tuple with the EvaluationPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationPeriod

`func (o *LicenseLicenseInfoRelationship) SetEvaluationPeriod(v int64)`

SetEvaluationPeriod sets EvaluationPeriod field to given value.

### HasEvaluationPeriod

`func (o *LicenseLicenseInfoRelationship) HasEvaluationPeriod() bool`

HasEvaluationPeriod returns a boolean if a field has been set.

### GetExtraEvaluation

`func (o *LicenseLicenseInfoRelationship) GetExtraEvaluation() int64`

GetExtraEvaluation returns the ExtraEvaluation field if non-nil, zero value otherwise.

### GetExtraEvaluationOk

`func (o *LicenseLicenseInfoRelationship) GetExtraEvaluationOk() (*int64, bool)`

GetExtraEvaluationOk returns a tuple with the ExtraEvaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraEvaluation

`func (o *LicenseLicenseInfoRelationship) SetExtraEvaluation(v int64)`

SetExtraEvaluation sets ExtraEvaluation field to given value.

### HasExtraEvaluation

`func (o *LicenseLicenseInfoRelationship) HasExtraEvaluation() bool`

HasExtraEvaluation returns a boolean if a field has been set.

### GetLicenseCount

`func (o *LicenseLicenseInfoRelationship) GetLicenseCount() int64`

GetLicenseCount returns the LicenseCount field if non-nil, zero value otherwise.

### GetLicenseCountOk

`func (o *LicenseLicenseInfoRelationship) GetLicenseCountOk() (*int64, bool)`

GetLicenseCountOk returns a tuple with the LicenseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseCount

`func (o *LicenseLicenseInfoRelationship) SetLicenseCount(v int64)`

SetLicenseCount sets LicenseCount field to given value.

### HasLicenseCount

`func (o *LicenseLicenseInfoRelationship) HasLicenseCount() bool`

HasLicenseCount returns a boolean if a field has been set.

### GetLicenseState

`func (o *LicenseLicenseInfoRelationship) GetLicenseState() string`

GetLicenseState returns the LicenseState field if non-nil, zero value otherwise.

### GetLicenseStateOk

`func (o *LicenseLicenseInfoRelationship) GetLicenseStateOk() (*string, bool)`

GetLicenseStateOk returns a tuple with the LicenseState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseState

`func (o *LicenseLicenseInfoRelationship) SetLicenseState(v string)`

SetLicenseState sets LicenseState field to given value.

### HasLicenseState

`func (o *LicenseLicenseInfoRelationship) HasLicenseState() bool`

HasLicenseState returns a boolean if a field has been set.

### GetLicenseType

`func (o *LicenseLicenseInfoRelationship) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *LicenseLicenseInfoRelationship) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *LicenseLicenseInfoRelationship) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *LicenseLicenseInfoRelationship) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### GetStartTime

`func (o *LicenseLicenseInfoRelationship) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *LicenseLicenseInfoRelationship) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *LicenseLicenseInfoRelationship) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *LicenseLicenseInfoRelationship) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTrialAdmin

`func (o *LicenseLicenseInfoRelationship) GetTrialAdmin() bool`

GetTrialAdmin returns the TrialAdmin field if non-nil, zero value otherwise.

### GetTrialAdminOk

`func (o *LicenseLicenseInfoRelationship) GetTrialAdminOk() (*bool, bool)`

GetTrialAdminOk returns a tuple with the TrialAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialAdmin

`func (o *LicenseLicenseInfoRelationship) SetTrialAdmin(v bool)`

SetTrialAdmin sets TrialAdmin field to given value.

### HasTrialAdmin

`func (o *LicenseLicenseInfoRelationship) HasTrialAdmin() bool`

HasTrialAdmin returns a boolean if a field has been set.

### GetAccountLicenseData

`func (o *LicenseLicenseInfoRelationship) GetAccountLicenseData() LicenseAccountLicenseDataRelationship`

GetAccountLicenseData returns the AccountLicenseData field if non-nil, zero value otherwise.

### GetAccountLicenseDataOk

`func (o *LicenseLicenseInfoRelationship) GetAccountLicenseDataOk() (*LicenseAccountLicenseDataRelationship, bool)`

GetAccountLicenseDataOk returns a tuple with the AccountLicenseData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLicenseData

`func (o *LicenseLicenseInfoRelationship) SetAccountLicenseData(v LicenseAccountLicenseDataRelationship)`

SetAccountLicenseData sets AccountLicenseData field to given value.

### HasAccountLicenseData

`func (o *LicenseLicenseInfoRelationship) HasAccountLicenseData() bool`

HasAccountLicenseData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


