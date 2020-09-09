# ApplianceImageBundleRelationship

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

### NewApplianceImageBundleRelationship

`func NewApplianceImageBundleRelationship(classId string, objectType string, ) *ApplianceImageBundleRelationship`

NewApplianceImageBundleRelationship instantiates a new ApplianceImageBundleRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceImageBundleRelationshipWithDefaults

`func NewApplianceImageBundleRelationshipWithDefaults() *ApplianceImageBundleRelationship`

NewApplianceImageBundleRelationshipWithDefaults instantiates a new ApplianceImageBundleRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *ApplianceImageBundleRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *ApplianceImageBundleRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *ApplianceImageBundleRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *ApplianceImageBundleRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *ApplianceImageBundleRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceImageBundleRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceImageBundleRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *ApplianceImageBundleRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *ApplianceImageBundleRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *ApplianceImageBundleRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *ApplianceImageBundleRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *ApplianceImageBundleRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *ApplianceImageBundleRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *ApplianceImageBundleRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *ApplianceImageBundleRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *ApplianceImageBundleRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *ApplianceImageBundleRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *ApplianceImageBundleRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *ApplianceImageBundleRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *ApplianceImageBundleRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *ApplianceImageBundleRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *ApplianceImageBundleRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *ApplianceImageBundleRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *ApplianceImageBundleRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceImageBundleRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceImageBundleRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *ApplianceImageBundleRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *ApplianceImageBundleRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *ApplianceImageBundleRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *ApplianceImageBundleRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *ApplianceImageBundleRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *ApplianceImageBundleRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *ApplianceImageBundleRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *ApplianceImageBundleRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *ApplianceImageBundleRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ApplianceImageBundleRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ApplianceImageBundleRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ApplianceImageBundleRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *ApplianceImageBundleRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *ApplianceImageBundleRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *ApplianceImageBundleRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *ApplianceImageBundleRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *ApplianceImageBundleRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *ApplianceImageBundleRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *ApplianceImageBundleRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *ApplianceImageBundleRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *ApplianceImageBundleRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *ApplianceImageBundleRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *ApplianceImageBundleRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ApplianceImageBundleRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ApplianceImageBundleRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *ApplianceImageBundleRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *ApplianceImageBundleRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *ApplianceImageBundleRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *ApplianceImageBundleRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *ApplianceImageBundleRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *ApplianceImageBundleRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *ApplianceImageBundleRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *ApplianceImageBundleRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *ApplianceImageBundleRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *ApplianceImageBundleRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *ApplianceImageBundleRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *ApplianceImageBundleRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *ApplianceImageBundleRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetAnsiblePackages

`func (o *ApplianceImageBundleRelationship) GetAnsiblePackages() []OnpremImagePackage`

GetAnsiblePackages returns the AnsiblePackages field if non-nil, zero value otherwise.

### GetAnsiblePackagesOk

`func (o *ApplianceImageBundleRelationship) GetAnsiblePackagesOk() (*[]OnpremImagePackage, bool)`

GetAnsiblePackagesOk returns a tuple with the AnsiblePackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsiblePackages

`func (o *ApplianceImageBundleRelationship) SetAnsiblePackages(v []OnpremImagePackage)`

SetAnsiblePackages sets AnsiblePackages field to given value.

### HasAnsiblePackages

`func (o *ApplianceImageBundleRelationship) HasAnsiblePackages() bool`

HasAnsiblePackages returns a boolean if a field has been set.

### GetAutoUpgrade

`func (o *ApplianceImageBundleRelationship) GetAutoUpgrade() bool`

GetAutoUpgrade returns the AutoUpgrade field if non-nil, zero value otherwise.

### GetAutoUpgradeOk

`func (o *ApplianceImageBundleRelationship) GetAutoUpgradeOk() (*bool, bool)`

GetAutoUpgradeOk returns a tuple with the AutoUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpgrade

`func (o *ApplianceImageBundleRelationship) SetAutoUpgrade(v bool)`

SetAutoUpgrade sets AutoUpgrade field to given value.

### HasAutoUpgrade

`func (o *ApplianceImageBundleRelationship) HasAutoUpgrade() bool`

HasAutoUpgrade returns a boolean if a field has been set.

### GetDcPackages

`func (o *ApplianceImageBundleRelationship) GetDcPackages() []OnpremImagePackage`

GetDcPackages returns the DcPackages field if non-nil, zero value otherwise.

### GetDcPackagesOk

`func (o *ApplianceImageBundleRelationship) GetDcPackagesOk() (*[]OnpremImagePackage, bool)`

GetDcPackagesOk returns a tuple with the DcPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcPackages

`func (o *ApplianceImageBundleRelationship) SetDcPackages(v []OnpremImagePackage)`

SetDcPackages sets DcPackages field to given value.

### HasDcPackages

`func (o *ApplianceImageBundleRelationship) HasDcPackages() bool`

HasDcPackages returns a boolean if a field has been set.

### GetDebugPackages

`func (o *ApplianceImageBundleRelationship) GetDebugPackages() []OnpremImagePackage`

GetDebugPackages returns the DebugPackages field if non-nil, zero value otherwise.

### GetDebugPackagesOk

`func (o *ApplianceImageBundleRelationship) GetDebugPackagesOk() (*[]OnpremImagePackage, bool)`

GetDebugPackagesOk returns a tuple with the DebugPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugPackages

`func (o *ApplianceImageBundleRelationship) SetDebugPackages(v []OnpremImagePackage)`

SetDebugPackages sets DebugPackages field to given value.

### HasDebugPackages

`func (o *ApplianceImageBundleRelationship) HasDebugPackages() bool`

HasDebugPackages returns a boolean if a field has been set.

### GetDescription

`func (o *ApplianceImageBundleRelationship) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplianceImageBundleRelationship) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplianceImageBundleRelationship) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplianceImageBundleRelationship) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndpointPackages

`func (o *ApplianceImageBundleRelationship) GetEndpointPackages() []OnpremImagePackage`

GetEndpointPackages returns the EndpointPackages field if non-nil, zero value otherwise.

### GetEndpointPackagesOk

`func (o *ApplianceImageBundleRelationship) GetEndpointPackagesOk() (*[]OnpremImagePackage, bool)`

GetEndpointPackagesOk returns a tuple with the EndpointPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointPackages

`func (o *ApplianceImageBundleRelationship) SetEndpointPackages(v []OnpremImagePackage)`

SetEndpointPackages sets EndpointPackages field to given value.

### HasEndpointPackages

`func (o *ApplianceImageBundleRelationship) HasEndpointPackages() bool`

HasEndpointPackages returns a boolean if a field has been set.

### GetFingerprint

`func (o *ApplianceImageBundleRelationship) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *ApplianceImageBundleRelationship) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *ApplianceImageBundleRelationship) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *ApplianceImageBundleRelationship) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetHasError

`func (o *ApplianceImageBundleRelationship) GetHasError() bool`

GetHasError returns the HasError field if non-nil, zero value otherwise.

### GetHasErrorOk

`func (o *ApplianceImageBundleRelationship) GetHasErrorOk() (*bool, bool)`

GetHasErrorOk returns a tuple with the HasError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasError

`func (o *ApplianceImageBundleRelationship) SetHasError(v bool)`

SetHasError sets HasError field to given value.

### HasHasError

`func (o *ApplianceImageBundleRelationship) HasHasError() bool`

HasHasError returns a boolean if a field has been set.

### GetInfraPackages

`func (o *ApplianceImageBundleRelationship) GetInfraPackages() []OnpremImagePackage`

GetInfraPackages returns the InfraPackages field if non-nil, zero value otherwise.

### GetInfraPackagesOk

`func (o *ApplianceImageBundleRelationship) GetInfraPackagesOk() (*[]OnpremImagePackage, bool)`

GetInfraPackagesOk returns a tuple with the InfraPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraPackages

`func (o *ApplianceImageBundleRelationship) SetInfraPackages(v []OnpremImagePackage)`

SetInfraPackages sets InfraPackages field to given value.

### HasInfraPackages

`func (o *ApplianceImageBundleRelationship) HasInfraPackages() bool`

HasInfraPackages returns a boolean if a field has been set.

### GetInitPackages

`func (o *ApplianceImageBundleRelationship) GetInitPackages() []OnpremImagePackage`

GetInitPackages returns the InitPackages field if non-nil, zero value otherwise.

### GetInitPackagesOk

`func (o *ApplianceImageBundleRelationship) GetInitPackagesOk() (*[]OnpremImagePackage, bool)`

GetInitPackagesOk returns a tuple with the InitPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitPackages

`func (o *ApplianceImageBundleRelationship) SetInitPackages(v []OnpremImagePackage)`

SetInitPackages sets InitPackages field to given value.

### HasInitPackages

`func (o *ApplianceImageBundleRelationship) HasInitPackages() bool`

HasInitPackages returns a boolean if a field has been set.

### GetName

`func (o *ApplianceImageBundleRelationship) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplianceImageBundleRelationship) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplianceImageBundleRelationship) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplianceImageBundleRelationship) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotes

`func (o *ApplianceImageBundleRelationship) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ApplianceImageBundleRelationship) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ApplianceImageBundleRelationship) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ApplianceImageBundleRelationship) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetPriority

`func (o *ApplianceImageBundleRelationship) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ApplianceImageBundleRelationship) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ApplianceImageBundleRelationship) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ApplianceImageBundleRelationship) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetReleaseTime

`func (o *ApplianceImageBundleRelationship) GetReleaseTime() time.Time`

GetReleaseTime returns the ReleaseTime field if non-nil, zero value otherwise.

### GetReleaseTimeOk

`func (o *ApplianceImageBundleRelationship) GetReleaseTimeOk() (*time.Time, bool)`

GetReleaseTimeOk returns a tuple with the ReleaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseTime

`func (o *ApplianceImageBundleRelationship) SetReleaseTime(v time.Time)`

SetReleaseTime sets ReleaseTime field to given value.

### HasReleaseTime

`func (o *ApplianceImageBundleRelationship) HasReleaseTime() bool`

HasReleaseTime returns a boolean if a field has been set.

### GetServicePackages

`func (o *ApplianceImageBundleRelationship) GetServicePackages() []OnpremImagePackage`

GetServicePackages returns the ServicePackages field if non-nil, zero value otherwise.

### GetServicePackagesOk

`func (o *ApplianceImageBundleRelationship) GetServicePackagesOk() (*[]OnpremImagePackage, bool)`

GetServicePackagesOk returns a tuple with the ServicePackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePackages

`func (o *ApplianceImageBundleRelationship) SetServicePackages(v []OnpremImagePackage)`

SetServicePackages sets ServicePackages field to given value.

### HasServicePackages

`func (o *ApplianceImageBundleRelationship) HasServicePackages() bool`

HasServicePackages returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ApplianceImageBundleRelationship) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ApplianceImageBundleRelationship) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ApplianceImageBundleRelationship) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ApplianceImageBundleRelationship) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetSystemPackages

`func (o *ApplianceImageBundleRelationship) GetSystemPackages() []OnpremImagePackage`

GetSystemPackages returns the SystemPackages field if non-nil, zero value otherwise.

### GetSystemPackagesOk

`func (o *ApplianceImageBundleRelationship) GetSystemPackagesOk() (*[]OnpremImagePackage, bool)`

GetSystemPackagesOk returns a tuple with the SystemPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPackages

`func (o *ApplianceImageBundleRelationship) SetSystemPackages(v []OnpremImagePackage)`

SetSystemPackages sets SystemPackages field to given value.

### HasSystemPackages

`func (o *ApplianceImageBundleRelationship) HasSystemPackages() bool`

HasSystemPackages returns a boolean if a field has been set.

### GetUiPackages

`func (o *ApplianceImageBundleRelationship) GetUiPackages() []OnpremImagePackage`

GetUiPackages returns the UiPackages field if non-nil, zero value otherwise.

### GetUiPackagesOk

`func (o *ApplianceImageBundleRelationship) GetUiPackagesOk() (*[]OnpremImagePackage, bool)`

GetUiPackagesOk returns a tuple with the UiPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiPackages

`func (o *ApplianceImageBundleRelationship) SetUiPackages(v []OnpremImagePackage)`

SetUiPackages sets UiPackages field to given value.

### HasUiPackages

`func (o *ApplianceImageBundleRelationship) HasUiPackages() bool`

HasUiPackages returns a boolean if a field has been set.

### GetUpgradeEndTime

`func (o *ApplianceImageBundleRelationship) GetUpgradeEndTime() time.Time`

GetUpgradeEndTime returns the UpgradeEndTime field if non-nil, zero value otherwise.

### GetUpgradeEndTimeOk

`func (o *ApplianceImageBundleRelationship) GetUpgradeEndTimeOk() (*time.Time, bool)`

GetUpgradeEndTimeOk returns a tuple with the UpgradeEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeEndTime

`func (o *ApplianceImageBundleRelationship) SetUpgradeEndTime(v time.Time)`

SetUpgradeEndTime sets UpgradeEndTime field to given value.

### HasUpgradeEndTime

`func (o *ApplianceImageBundleRelationship) HasUpgradeEndTime() bool`

HasUpgradeEndTime returns a boolean if a field has been set.

### GetUpgradeGracePeriod

`func (o *ApplianceImageBundleRelationship) GetUpgradeGracePeriod() int64`

GetUpgradeGracePeriod returns the UpgradeGracePeriod field if non-nil, zero value otherwise.

### GetUpgradeGracePeriodOk

`func (o *ApplianceImageBundleRelationship) GetUpgradeGracePeriodOk() (*int64, bool)`

GetUpgradeGracePeriodOk returns a tuple with the UpgradeGracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeGracePeriod

`func (o *ApplianceImageBundleRelationship) SetUpgradeGracePeriod(v int64)`

SetUpgradeGracePeriod sets UpgradeGracePeriod field to given value.

### HasUpgradeGracePeriod

`func (o *ApplianceImageBundleRelationship) HasUpgradeGracePeriod() bool`

HasUpgradeGracePeriod returns a boolean if a field has been set.

### GetUpgradeImpactDuration

`func (o *ApplianceImageBundleRelationship) GetUpgradeImpactDuration() int64`

GetUpgradeImpactDuration returns the UpgradeImpactDuration field if non-nil, zero value otherwise.

### GetUpgradeImpactDurationOk

`func (o *ApplianceImageBundleRelationship) GetUpgradeImpactDurationOk() (*int64, bool)`

GetUpgradeImpactDurationOk returns a tuple with the UpgradeImpactDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeImpactDuration

`func (o *ApplianceImageBundleRelationship) SetUpgradeImpactDuration(v int64)`

SetUpgradeImpactDuration sets UpgradeImpactDuration field to given value.

### HasUpgradeImpactDuration

`func (o *ApplianceImageBundleRelationship) HasUpgradeImpactDuration() bool`

HasUpgradeImpactDuration returns a boolean if a field has been set.

### GetUpgradeImpactEnum

`func (o *ApplianceImageBundleRelationship) GetUpgradeImpactEnum() string`

GetUpgradeImpactEnum returns the UpgradeImpactEnum field if non-nil, zero value otherwise.

### GetUpgradeImpactEnumOk

`func (o *ApplianceImageBundleRelationship) GetUpgradeImpactEnumOk() (*string, bool)`

GetUpgradeImpactEnumOk returns a tuple with the UpgradeImpactEnum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeImpactEnum

`func (o *ApplianceImageBundleRelationship) SetUpgradeImpactEnum(v string)`

SetUpgradeImpactEnum sets UpgradeImpactEnum field to given value.

### HasUpgradeImpactEnum

`func (o *ApplianceImageBundleRelationship) HasUpgradeImpactEnum() bool`

HasUpgradeImpactEnum returns a boolean if a field has been set.

### GetUpgradeStartTime

`func (o *ApplianceImageBundleRelationship) GetUpgradeStartTime() time.Time`

GetUpgradeStartTime returns the UpgradeStartTime field if non-nil, zero value otherwise.

### GetUpgradeStartTimeOk

`func (o *ApplianceImageBundleRelationship) GetUpgradeStartTimeOk() (*time.Time, bool)`

GetUpgradeStartTimeOk returns a tuple with the UpgradeStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeStartTime

`func (o *ApplianceImageBundleRelationship) SetUpgradeStartTime(v time.Time)`

SetUpgradeStartTime sets UpgradeStartTime field to given value.

### HasUpgradeStartTime

`func (o *ApplianceImageBundleRelationship) HasUpgradeStartTime() bool`

HasUpgradeStartTime returns a boolean if a field has been set.

### GetVersion

`func (o *ApplianceImageBundleRelationship) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApplianceImageBundleRelationship) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApplianceImageBundleRelationship) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ApplianceImageBundleRelationship) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


