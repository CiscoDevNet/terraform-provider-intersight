# LicenseAccountLicenseDataRelationship

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
**AccountId** | Pointer to **string** | Root user&#39;s ID of the account. | [optional] [readonly] 
**AgentData** | Pointer to **string** | Agent trusted store data. | [optional] [readonly] 
**AuthExpireTime** | Pointer to **string** | Authorization expiration time. | [optional] [readonly] 
**AuthInitialTime** | Pointer to **string** | Intial authorization time. | [optional] [readonly] 
**AuthNextTime** | Pointer to **string** | Next time for the authorization. | [optional] [readonly] 
**Category** | Pointer to **string** | Account license data category name. | [optional] [readonly] 
**DefaultLicenseType** | Pointer to **string** | Default license tier set by user. * &#x60;Base&#x60; - Base as a License type. It is default license type. * &#x60;Essential&#x60; - Essential as a License type. * &#x60;Standard&#x60; - Standard as a License type. * &#x60;Advantage&#x60; - Advantage as a License type. * &#x60;Premier&#x60; - Premier as a License type. | [optional] [default to "Base"]
**ErrorDesc** | Pointer to **string** | The detailed error message when there is any error related to license sync of this account. | [optional] [readonly] 
**Group** | Pointer to **string** | Account license data group name. | [optional] [readonly] 
**HighestCompliantLicenseTier** | Pointer to **string** | The highest license tier which is in compliant of this account. * &#x60;Base&#x60; - Base as a License type. It is default license type. * &#x60;Essential&#x60; - Essential as a License type. * &#x60;Standard&#x60; - Standard as a License type. * &#x60;Advantage&#x60; - Advantage as a License type. * &#x60;Premier&#x60; - Premier as a License type. | [optional] [readonly] [default to "Base"]
**LastSync** | Pointer to [**time.Time**](time.Time.md) | Specifies last sync time with SA. | [optional] [readonly] 
**LastUpdatedTime** | Pointer to [**time.Time**](time.Time.md) | Record&#39;s last update datetime. | [optional] [readonly] 
**LicenseState** | Pointer to **string** | Aggregrated mode for the agent. | [optional] [readonly] 
**LicenseTechSupportInfo** | Pointer to **string** | Tech-support info of a smart-agent. | [optional] [readonly] 
**RegisterExpireTime** | Pointer to **string** | Registration exipiration time. | [optional] [readonly] 
**RegisterInitialTime** | Pointer to **string** | Initial time of registration. | [optional] [readonly] 
**RegisterNextTime** | Pointer to **string** | Next time for the license registration. | [optional] [readonly] 
**RegistrationStatus** | Pointer to **string** | Registration status of a smart-agent. | [optional] [readonly] 
**RenewFailureString** | Pointer to **string** | License renewal failure message. | [optional] [readonly] 
**SmartAccount** | Pointer to **string** | Name of the smart account. | [optional] [readonly] 
**SyncStatus** | Pointer to **string** | Current sync status for the account. | [optional] [readonly] 
**VirtualAccount** | Pointer to **string** | Name of the virtual account. | [optional] [readonly] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 
**CustomerOp** | Pointer to [**LicenseCustomerOpRelationship**](license.CustomerOp.Relationship.md) |  | [optional] 
**Licenseinfos** | Pointer to [**[]LicenseLicenseInfoRelationship**](license.LicenseInfo.Relationship.md) | An array of relationships to licenseLicenseInfo resources. | [optional] 
**SmartlicenseToken** | Pointer to [**LicenseSmartlicenseTokenRelationship**](license.SmartlicenseToken.Relationship.md) |  | [optional] 

## Methods

### NewLicenseAccountLicenseDataRelationship

`func NewLicenseAccountLicenseDataRelationship(classId string, objectType string, ) *LicenseAccountLicenseDataRelationship`

NewLicenseAccountLicenseDataRelationship instantiates a new LicenseAccountLicenseDataRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseAccountLicenseDataRelationshipWithDefaults

`func NewLicenseAccountLicenseDataRelationshipWithDefaults() *LicenseAccountLicenseDataRelationship`

NewLicenseAccountLicenseDataRelationshipWithDefaults instantiates a new LicenseAccountLicenseDataRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *LicenseAccountLicenseDataRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *LicenseAccountLicenseDataRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *LicenseAccountLicenseDataRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *LicenseAccountLicenseDataRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *LicenseAccountLicenseDataRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *LicenseAccountLicenseDataRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *LicenseAccountLicenseDataRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *LicenseAccountLicenseDataRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *LicenseAccountLicenseDataRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *LicenseAccountLicenseDataRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *LicenseAccountLicenseDataRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *LicenseAccountLicenseDataRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *LicenseAccountLicenseDataRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *LicenseAccountLicenseDataRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *LicenseAccountLicenseDataRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *LicenseAccountLicenseDataRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *LicenseAccountLicenseDataRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *LicenseAccountLicenseDataRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *LicenseAccountLicenseDataRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *LicenseAccountLicenseDataRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *LicenseAccountLicenseDataRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *LicenseAccountLicenseDataRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *LicenseAccountLicenseDataRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *LicenseAccountLicenseDataRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *LicenseAccountLicenseDataRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *LicenseAccountLicenseDataRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *LicenseAccountLicenseDataRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *LicenseAccountLicenseDataRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *LicenseAccountLicenseDataRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *LicenseAccountLicenseDataRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *LicenseAccountLicenseDataRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *LicenseAccountLicenseDataRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *LicenseAccountLicenseDataRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *LicenseAccountLicenseDataRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *LicenseAccountLicenseDataRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LicenseAccountLicenseDataRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LicenseAccountLicenseDataRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *LicenseAccountLicenseDataRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *LicenseAccountLicenseDataRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *LicenseAccountLicenseDataRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *LicenseAccountLicenseDataRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *LicenseAccountLicenseDataRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *LicenseAccountLicenseDataRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *LicenseAccountLicenseDataRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *LicenseAccountLicenseDataRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *LicenseAccountLicenseDataRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *LicenseAccountLicenseDataRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *LicenseAccountLicenseDataRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *LicenseAccountLicenseDataRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *LicenseAccountLicenseDataRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *LicenseAccountLicenseDataRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *LicenseAccountLicenseDataRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *LicenseAccountLicenseDataRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *LicenseAccountLicenseDataRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *LicenseAccountLicenseDataRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *LicenseAccountLicenseDataRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *LicenseAccountLicenseDataRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *LicenseAccountLicenseDataRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *LicenseAccountLicenseDataRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *LicenseAccountLicenseDataRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *LicenseAccountLicenseDataRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *LicenseAccountLicenseDataRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *LicenseAccountLicenseDataRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *LicenseAccountLicenseDataRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetAccountId

`func (o *LicenseAccountLicenseDataRelationship) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *LicenseAccountLicenseDataRelationship) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *LicenseAccountLicenseDataRelationship) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *LicenseAccountLicenseDataRelationship) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAgentData

`func (o *LicenseAccountLicenseDataRelationship) GetAgentData() string`

GetAgentData returns the AgentData field if non-nil, zero value otherwise.

### GetAgentDataOk

`func (o *LicenseAccountLicenseDataRelationship) GetAgentDataOk() (*string, bool)`

GetAgentDataOk returns a tuple with the AgentData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentData

`func (o *LicenseAccountLicenseDataRelationship) SetAgentData(v string)`

SetAgentData sets AgentData field to given value.

### HasAgentData

`func (o *LicenseAccountLicenseDataRelationship) HasAgentData() bool`

HasAgentData returns a boolean if a field has been set.

### GetAuthExpireTime

`func (o *LicenseAccountLicenseDataRelationship) GetAuthExpireTime() string`

GetAuthExpireTime returns the AuthExpireTime field if non-nil, zero value otherwise.

### GetAuthExpireTimeOk

`func (o *LicenseAccountLicenseDataRelationship) GetAuthExpireTimeOk() (*string, bool)`

GetAuthExpireTimeOk returns a tuple with the AuthExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthExpireTime

`func (o *LicenseAccountLicenseDataRelationship) SetAuthExpireTime(v string)`

SetAuthExpireTime sets AuthExpireTime field to given value.

### HasAuthExpireTime

`func (o *LicenseAccountLicenseDataRelationship) HasAuthExpireTime() bool`

HasAuthExpireTime returns a boolean if a field has been set.

### GetAuthInitialTime

`func (o *LicenseAccountLicenseDataRelationship) GetAuthInitialTime() string`

GetAuthInitialTime returns the AuthInitialTime field if non-nil, zero value otherwise.

### GetAuthInitialTimeOk

`func (o *LicenseAccountLicenseDataRelationship) GetAuthInitialTimeOk() (*string, bool)`

GetAuthInitialTimeOk returns a tuple with the AuthInitialTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthInitialTime

`func (o *LicenseAccountLicenseDataRelationship) SetAuthInitialTime(v string)`

SetAuthInitialTime sets AuthInitialTime field to given value.

### HasAuthInitialTime

`func (o *LicenseAccountLicenseDataRelationship) HasAuthInitialTime() bool`

HasAuthInitialTime returns a boolean if a field has been set.

### GetAuthNextTime

`func (o *LicenseAccountLicenseDataRelationship) GetAuthNextTime() string`

GetAuthNextTime returns the AuthNextTime field if non-nil, zero value otherwise.

### GetAuthNextTimeOk

`func (o *LicenseAccountLicenseDataRelationship) GetAuthNextTimeOk() (*string, bool)`

GetAuthNextTimeOk returns a tuple with the AuthNextTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthNextTime

`func (o *LicenseAccountLicenseDataRelationship) SetAuthNextTime(v string)`

SetAuthNextTime sets AuthNextTime field to given value.

### HasAuthNextTime

`func (o *LicenseAccountLicenseDataRelationship) HasAuthNextTime() bool`

HasAuthNextTime returns a boolean if a field has been set.

### GetCategory

`func (o *LicenseAccountLicenseDataRelationship) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *LicenseAccountLicenseDataRelationship) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *LicenseAccountLicenseDataRelationship) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *LicenseAccountLicenseDataRelationship) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDefaultLicenseType

`func (o *LicenseAccountLicenseDataRelationship) GetDefaultLicenseType() string`

GetDefaultLicenseType returns the DefaultLicenseType field if non-nil, zero value otherwise.

### GetDefaultLicenseTypeOk

`func (o *LicenseAccountLicenseDataRelationship) GetDefaultLicenseTypeOk() (*string, bool)`

GetDefaultLicenseTypeOk returns a tuple with the DefaultLicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLicenseType

`func (o *LicenseAccountLicenseDataRelationship) SetDefaultLicenseType(v string)`

SetDefaultLicenseType sets DefaultLicenseType field to given value.

### HasDefaultLicenseType

`func (o *LicenseAccountLicenseDataRelationship) HasDefaultLicenseType() bool`

HasDefaultLicenseType returns a boolean if a field has been set.

### GetErrorDesc

`func (o *LicenseAccountLicenseDataRelationship) GetErrorDesc() string`

GetErrorDesc returns the ErrorDesc field if non-nil, zero value otherwise.

### GetErrorDescOk

`func (o *LicenseAccountLicenseDataRelationship) GetErrorDescOk() (*string, bool)`

GetErrorDescOk returns a tuple with the ErrorDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDesc

`func (o *LicenseAccountLicenseDataRelationship) SetErrorDesc(v string)`

SetErrorDesc sets ErrorDesc field to given value.

### HasErrorDesc

`func (o *LicenseAccountLicenseDataRelationship) HasErrorDesc() bool`

HasErrorDesc returns a boolean if a field has been set.

### GetGroup

`func (o *LicenseAccountLicenseDataRelationship) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *LicenseAccountLicenseDataRelationship) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *LicenseAccountLicenseDataRelationship) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *LicenseAccountLicenseDataRelationship) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetHighestCompliantLicenseTier

`func (o *LicenseAccountLicenseDataRelationship) GetHighestCompliantLicenseTier() string`

GetHighestCompliantLicenseTier returns the HighestCompliantLicenseTier field if non-nil, zero value otherwise.

### GetHighestCompliantLicenseTierOk

`func (o *LicenseAccountLicenseDataRelationship) GetHighestCompliantLicenseTierOk() (*string, bool)`

GetHighestCompliantLicenseTierOk returns a tuple with the HighestCompliantLicenseTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighestCompliantLicenseTier

`func (o *LicenseAccountLicenseDataRelationship) SetHighestCompliantLicenseTier(v string)`

SetHighestCompliantLicenseTier sets HighestCompliantLicenseTier field to given value.

### HasHighestCompliantLicenseTier

`func (o *LicenseAccountLicenseDataRelationship) HasHighestCompliantLicenseTier() bool`

HasHighestCompliantLicenseTier returns a boolean if a field has been set.

### GetLastSync

`func (o *LicenseAccountLicenseDataRelationship) GetLastSync() time.Time`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *LicenseAccountLicenseDataRelationship) GetLastSyncOk() (*time.Time, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *LicenseAccountLicenseDataRelationship) SetLastSync(v time.Time)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *LicenseAccountLicenseDataRelationship) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *LicenseAccountLicenseDataRelationship) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *LicenseAccountLicenseDataRelationship) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *LicenseAccountLicenseDataRelationship) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *LicenseAccountLicenseDataRelationship) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetLicenseState

`func (o *LicenseAccountLicenseDataRelationship) GetLicenseState() string`

GetLicenseState returns the LicenseState field if non-nil, zero value otherwise.

### GetLicenseStateOk

`func (o *LicenseAccountLicenseDataRelationship) GetLicenseStateOk() (*string, bool)`

GetLicenseStateOk returns a tuple with the LicenseState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseState

`func (o *LicenseAccountLicenseDataRelationship) SetLicenseState(v string)`

SetLicenseState sets LicenseState field to given value.

### HasLicenseState

`func (o *LicenseAccountLicenseDataRelationship) HasLicenseState() bool`

HasLicenseState returns a boolean if a field has been set.

### GetLicenseTechSupportInfo

`func (o *LicenseAccountLicenseDataRelationship) GetLicenseTechSupportInfo() string`

GetLicenseTechSupportInfo returns the LicenseTechSupportInfo field if non-nil, zero value otherwise.

### GetLicenseTechSupportInfoOk

`func (o *LicenseAccountLicenseDataRelationship) GetLicenseTechSupportInfoOk() (*string, bool)`

GetLicenseTechSupportInfoOk returns a tuple with the LicenseTechSupportInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseTechSupportInfo

`func (o *LicenseAccountLicenseDataRelationship) SetLicenseTechSupportInfo(v string)`

SetLicenseTechSupportInfo sets LicenseTechSupportInfo field to given value.

### HasLicenseTechSupportInfo

`func (o *LicenseAccountLicenseDataRelationship) HasLicenseTechSupportInfo() bool`

HasLicenseTechSupportInfo returns a boolean if a field has been set.

### GetRegisterExpireTime

`func (o *LicenseAccountLicenseDataRelationship) GetRegisterExpireTime() string`

GetRegisterExpireTime returns the RegisterExpireTime field if non-nil, zero value otherwise.

### GetRegisterExpireTimeOk

`func (o *LicenseAccountLicenseDataRelationship) GetRegisterExpireTimeOk() (*string, bool)`

GetRegisterExpireTimeOk returns a tuple with the RegisterExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterExpireTime

`func (o *LicenseAccountLicenseDataRelationship) SetRegisterExpireTime(v string)`

SetRegisterExpireTime sets RegisterExpireTime field to given value.

### HasRegisterExpireTime

`func (o *LicenseAccountLicenseDataRelationship) HasRegisterExpireTime() bool`

HasRegisterExpireTime returns a boolean if a field has been set.

### GetRegisterInitialTime

`func (o *LicenseAccountLicenseDataRelationship) GetRegisterInitialTime() string`

GetRegisterInitialTime returns the RegisterInitialTime field if non-nil, zero value otherwise.

### GetRegisterInitialTimeOk

`func (o *LicenseAccountLicenseDataRelationship) GetRegisterInitialTimeOk() (*string, bool)`

GetRegisterInitialTimeOk returns a tuple with the RegisterInitialTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterInitialTime

`func (o *LicenseAccountLicenseDataRelationship) SetRegisterInitialTime(v string)`

SetRegisterInitialTime sets RegisterInitialTime field to given value.

### HasRegisterInitialTime

`func (o *LicenseAccountLicenseDataRelationship) HasRegisterInitialTime() bool`

HasRegisterInitialTime returns a boolean if a field has been set.

### GetRegisterNextTime

`func (o *LicenseAccountLicenseDataRelationship) GetRegisterNextTime() string`

GetRegisterNextTime returns the RegisterNextTime field if non-nil, zero value otherwise.

### GetRegisterNextTimeOk

`func (o *LicenseAccountLicenseDataRelationship) GetRegisterNextTimeOk() (*string, bool)`

GetRegisterNextTimeOk returns a tuple with the RegisterNextTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterNextTime

`func (o *LicenseAccountLicenseDataRelationship) SetRegisterNextTime(v string)`

SetRegisterNextTime sets RegisterNextTime field to given value.

### HasRegisterNextTime

`func (o *LicenseAccountLicenseDataRelationship) HasRegisterNextTime() bool`

HasRegisterNextTime returns a boolean if a field has been set.

### GetRegistrationStatus

`func (o *LicenseAccountLicenseDataRelationship) GetRegistrationStatus() string`

GetRegistrationStatus returns the RegistrationStatus field if non-nil, zero value otherwise.

### GetRegistrationStatusOk

`func (o *LicenseAccountLicenseDataRelationship) GetRegistrationStatusOk() (*string, bool)`

GetRegistrationStatusOk returns a tuple with the RegistrationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationStatus

`func (o *LicenseAccountLicenseDataRelationship) SetRegistrationStatus(v string)`

SetRegistrationStatus sets RegistrationStatus field to given value.

### HasRegistrationStatus

`func (o *LicenseAccountLicenseDataRelationship) HasRegistrationStatus() bool`

HasRegistrationStatus returns a boolean if a field has been set.

### GetRenewFailureString

`func (o *LicenseAccountLicenseDataRelationship) GetRenewFailureString() string`

GetRenewFailureString returns the RenewFailureString field if non-nil, zero value otherwise.

### GetRenewFailureStringOk

`func (o *LicenseAccountLicenseDataRelationship) GetRenewFailureStringOk() (*string, bool)`

GetRenewFailureStringOk returns a tuple with the RenewFailureString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewFailureString

`func (o *LicenseAccountLicenseDataRelationship) SetRenewFailureString(v string)`

SetRenewFailureString sets RenewFailureString field to given value.

### HasRenewFailureString

`func (o *LicenseAccountLicenseDataRelationship) HasRenewFailureString() bool`

HasRenewFailureString returns a boolean if a field has been set.

### GetSmartAccount

`func (o *LicenseAccountLicenseDataRelationship) GetSmartAccount() string`

GetSmartAccount returns the SmartAccount field if non-nil, zero value otherwise.

### GetSmartAccountOk

`func (o *LicenseAccountLicenseDataRelationship) GetSmartAccountOk() (*string, bool)`

GetSmartAccountOk returns a tuple with the SmartAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartAccount

`func (o *LicenseAccountLicenseDataRelationship) SetSmartAccount(v string)`

SetSmartAccount sets SmartAccount field to given value.

### HasSmartAccount

`func (o *LicenseAccountLicenseDataRelationship) HasSmartAccount() bool`

HasSmartAccount returns a boolean if a field has been set.

### GetSyncStatus

`func (o *LicenseAccountLicenseDataRelationship) GetSyncStatus() string`

GetSyncStatus returns the SyncStatus field if non-nil, zero value otherwise.

### GetSyncStatusOk

`func (o *LicenseAccountLicenseDataRelationship) GetSyncStatusOk() (*string, bool)`

GetSyncStatusOk returns a tuple with the SyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncStatus

`func (o *LicenseAccountLicenseDataRelationship) SetSyncStatus(v string)`

SetSyncStatus sets SyncStatus field to given value.

### HasSyncStatus

`func (o *LicenseAccountLicenseDataRelationship) HasSyncStatus() bool`

HasSyncStatus returns a boolean if a field has been set.

### GetVirtualAccount

`func (o *LicenseAccountLicenseDataRelationship) GetVirtualAccount() string`

GetVirtualAccount returns the VirtualAccount field if non-nil, zero value otherwise.

### GetVirtualAccountOk

`func (o *LicenseAccountLicenseDataRelationship) GetVirtualAccountOk() (*string, bool)`

GetVirtualAccountOk returns a tuple with the VirtualAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualAccount

`func (o *LicenseAccountLicenseDataRelationship) SetVirtualAccount(v string)`

SetVirtualAccount sets VirtualAccount field to given value.

### HasVirtualAccount

`func (o *LicenseAccountLicenseDataRelationship) HasVirtualAccount() bool`

HasVirtualAccount returns a boolean if a field has been set.

### GetAccount

`func (o *LicenseAccountLicenseDataRelationship) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *LicenseAccountLicenseDataRelationship) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *LicenseAccountLicenseDataRelationship) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *LicenseAccountLicenseDataRelationship) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCustomerOp

`func (o *LicenseAccountLicenseDataRelationship) GetCustomerOp() LicenseCustomerOpRelationship`

GetCustomerOp returns the CustomerOp field if non-nil, zero value otherwise.

### GetCustomerOpOk

`func (o *LicenseAccountLicenseDataRelationship) GetCustomerOpOk() (*LicenseCustomerOpRelationship, bool)`

GetCustomerOpOk returns a tuple with the CustomerOp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerOp

`func (o *LicenseAccountLicenseDataRelationship) SetCustomerOp(v LicenseCustomerOpRelationship)`

SetCustomerOp sets CustomerOp field to given value.

### HasCustomerOp

`func (o *LicenseAccountLicenseDataRelationship) HasCustomerOp() bool`

HasCustomerOp returns a boolean if a field has been set.

### GetLicenseinfos

`func (o *LicenseAccountLicenseDataRelationship) GetLicenseinfos() []LicenseLicenseInfoRelationship`

GetLicenseinfos returns the Licenseinfos field if non-nil, zero value otherwise.

### GetLicenseinfosOk

`func (o *LicenseAccountLicenseDataRelationship) GetLicenseinfosOk() (*[]LicenseLicenseInfoRelationship, bool)`

GetLicenseinfosOk returns a tuple with the Licenseinfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseinfos

`func (o *LicenseAccountLicenseDataRelationship) SetLicenseinfos(v []LicenseLicenseInfoRelationship)`

SetLicenseinfos sets Licenseinfos field to given value.

### HasLicenseinfos

`func (o *LicenseAccountLicenseDataRelationship) HasLicenseinfos() bool`

HasLicenseinfos returns a boolean if a field has been set.

### SetLicenseinfosNil

`func (o *LicenseAccountLicenseDataRelationship) SetLicenseinfosNil(b bool)`

 SetLicenseinfosNil sets the value for Licenseinfos to be an explicit nil

### UnsetLicenseinfos
`func (o *LicenseAccountLicenseDataRelationship) UnsetLicenseinfos()`

UnsetLicenseinfos ensures that no value is present for Licenseinfos, not even an explicit nil
### GetSmartlicenseToken

`func (o *LicenseAccountLicenseDataRelationship) GetSmartlicenseToken() LicenseSmartlicenseTokenRelationship`

GetSmartlicenseToken returns the SmartlicenseToken field if non-nil, zero value otherwise.

### GetSmartlicenseTokenOk

`func (o *LicenseAccountLicenseDataRelationship) GetSmartlicenseTokenOk() (*LicenseSmartlicenseTokenRelationship, bool)`

GetSmartlicenseTokenOk returns a tuple with the SmartlicenseToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartlicenseToken

`func (o *LicenseAccountLicenseDataRelationship) SetSmartlicenseToken(v LicenseSmartlicenseTokenRelationship)`

SetSmartlicenseToken sets SmartlicenseToken field to given value.

### HasSmartlicenseToken

`func (o *LicenseAccountLicenseDataRelationship) HasSmartlicenseToken() bool`

HasSmartlicenseToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


