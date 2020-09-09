# LicenseLicenseReservationOpRelationship

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
**AuthCode** | Pointer to **string** | Revervation code used to install the license. | [optional] 
**AuthCodeInstalled** | Pointer to **bool** | Flag to indicate whether authorization code is installed. | [optional] [readonly] 
**ConfirmCode** | Pointer to **string** | Confirm code used to complete the license update on smart license account. | [optional] [readonly] 
**GenerateRequestCode** | Pointer to **bool** | Trigger the generation of request code for specific license reservation. | [optional] 
**GenerateReturnCode** | Pointer to **bool** | Trigger the generation of return code for specific license reservation. | [optional] 
**RequestCode** | Pointer to **string** | Revervation code used to generate authorization code from CSSM. | [optional] [readonly] 
**ReturnCode** | Pointer to **string** | Return code used to return the reserved license to smart license account. | [optional] [readonly] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 

## Methods

### NewLicenseLicenseReservationOpRelationship

`func NewLicenseLicenseReservationOpRelationship(classId string, objectType string, ) *LicenseLicenseReservationOpRelationship`

NewLicenseLicenseReservationOpRelationship instantiates a new LicenseLicenseReservationOpRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseLicenseReservationOpRelationshipWithDefaults

`func NewLicenseLicenseReservationOpRelationshipWithDefaults() *LicenseLicenseReservationOpRelationship`

NewLicenseLicenseReservationOpRelationshipWithDefaults instantiates a new LicenseLicenseReservationOpRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *LicenseLicenseReservationOpRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *LicenseLicenseReservationOpRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *LicenseLicenseReservationOpRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *LicenseLicenseReservationOpRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *LicenseLicenseReservationOpRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *LicenseLicenseReservationOpRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *LicenseLicenseReservationOpRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *LicenseLicenseReservationOpRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *LicenseLicenseReservationOpRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *LicenseLicenseReservationOpRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *LicenseLicenseReservationOpRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *LicenseLicenseReservationOpRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *LicenseLicenseReservationOpRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *LicenseLicenseReservationOpRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *LicenseLicenseReservationOpRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *LicenseLicenseReservationOpRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *LicenseLicenseReservationOpRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *LicenseLicenseReservationOpRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *LicenseLicenseReservationOpRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *LicenseLicenseReservationOpRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *LicenseLicenseReservationOpRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *LicenseLicenseReservationOpRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *LicenseLicenseReservationOpRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *LicenseLicenseReservationOpRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *LicenseLicenseReservationOpRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *LicenseLicenseReservationOpRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *LicenseLicenseReservationOpRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *LicenseLicenseReservationOpRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *LicenseLicenseReservationOpRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *LicenseLicenseReservationOpRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *LicenseLicenseReservationOpRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *LicenseLicenseReservationOpRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *LicenseLicenseReservationOpRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *LicenseLicenseReservationOpRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *LicenseLicenseReservationOpRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LicenseLicenseReservationOpRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LicenseLicenseReservationOpRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *LicenseLicenseReservationOpRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *LicenseLicenseReservationOpRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *LicenseLicenseReservationOpRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *LicenseLicenseReservationOpRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *LicenseLicenseReservationOpRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *LicenseLicenseReservationOpRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *LicenseLicenseReservationOpRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *LicenseLicenseReservationOpRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *LicenseLicenseReservationOpRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *LicenseLicenseReservationOpRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *LicenseLicenseReservationOpRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *LicenseLicenseReservationOpRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *LicenseLicenseReservationOpRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *LicenseLicenseReservationOpRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *LicenseLicenseReservationOpRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *LicenseLicenseReservationOpRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *LicenseLicenseReservationOpRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *LicenseLicenseReservationOpRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *LicenseLicenseReservationOpRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *LicenseLicenseReservationOpRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *LicenseLicenseReservationOpRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *LicenseLicenseReservationOpRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *LicenseLicenseReservationOpRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *LicenseLicenseReservationOpRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *LicenseLicenseReservationOpRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *LicenseLicenseReservationOpRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *LicenseLicenseReservationOpRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetAuthCode

`func (o *LicenseLicenseReservationOpRelationship) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *LicenseLicenseReservationOpRelationship) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *LicenseLicenseReservationOpRelationship) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.

### HasAuthCode

`func (o *LicenseLicenseReservationOpRelationship) HasAuthCode() bool`

HasAuthCode returns a boolean if a field has been set.

### GetAuthCodeInstalled

`func (o *LicenseLicenseReservationOpRelationship) GetAuthCodeInstalled() bool`

GetAuthCodeInstalled returns the AuthCodeInstalled field if non-nil, zero value otherwise.

### GetAuthCodeInstalledOk

`func (o *LicenseLicenseReservationOpRelationship) GetAuthCodeInstalledOk() (*bool, bool)`

GetAuthCodeInstalledOk returns a tuple with the AuthCodeInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCodeInstalled

`func (o *LicenseLicenseReservationOpRelationship) SetAuthCodeInstalled(v bool)`

SetAuthCodeInstalled sets AuthCodeInstalled field to given value.

### HasAuthCodeInstalled

`func (o *LicenseLicenseReservationOpRelationship) HasAuthCodeInstalled() bool`

HasAuthCodeInstalled returns a boolean if a field has been set.

### GetConfirmCode

`func (o *LicenseLicenseReservationOpRelationship) GetConfirmCode() string`

GetConfirmCode returns the ConfirmCode field if non-nil, zero value otherwise.

### GetConfirmCodeOk

`func (o *LicenseLicenseReservationOpRelationship) GetConfirmCodeOk() (*string, bool)`

GetConfirmCodeOk returns a tuple with the ConfirmCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmCode

`func (o *LicenseLicenseReservationOpRelationship) SetConfirmCode(v string)`

SetConfirmCode sets ConfirmCode field to given value.

### HasConfirmCode

`func (o *LicenseLicenseReservationOpRelationship) HasConfirmCode() bool`

HasConfirmCode returns a boolean if a field has been set.

### GetGenerateRequestCode

`func (o *LicenseLicenseReservationOpRelationship) GetGenerateRequestCode() bool`

GetGenerateRequestCode returns the GenerateRequestCode field if non-nil, zero value otherwise.

### GetGenerateRequestCodeOk

`func (o *LicenseLicenseReservationOpRelationship) GetGenerateRequestCodeOk() (*bool, bool)`

GetGenerateRequestCodeOk returns a tuple with the GenerateRequestCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateRequestCode

`func (o *LicenseLicenseReservationOpRelationship) SetGenerateRequestCode(v bool)`

SetGenerateRequestCode sets GenerateRequestCode field to given value.

### HasGenerateRequestCode

`func (o *LicenseLicenseReservationOpRelationship) HasGenerateRequestCode() bool`

HasGenerateRequestCode returns a boolean if a field has been set.

### GetGenerateReturnCode

`func (o *LicenseLicenseReservationOpRelationship) GetGenerateReturnCode() bool`

GetGenerateReturnCode returns the GenerateReturnCode field if non-nil, zero value otherwise.

### GetGenerateReturnCodeOk

`func (o *LicenseLicenseReservationOpRelationship) GetGenerateReturnCodeOk() (*bool, bool)`

GetGenerateReturnCodeOk returns a tuple with the GenerateReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateReturnCode

`func (o *LicenseLicenseReservationOpRelationship) SetGenerateReturnCode(v bool)`

SetGenerateReturnCode sets GenerateReturnCode field to given value.

### HasGenerateReturnCode

`func (o *LicenseLicenseReservationOpRelationship) HasGenerateReturnCode() bool`

HasGenerateReturnCode returns a boolean if a field has been set.

### GetRequestCode

`func (o *LicenseLicenseReservationOpRelationship) GetRequestCode() string`

GetRequestCode returns the RequestCode field if non-nil, zero value otherwise.

### GetRequestCodeOk

`func (o *LicenseLicenseReservationOpRelationship) GetRequestCodeOk() (*string, bool)`

GetRequestCodeOk returns a tuple with the RequestCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCode

`func (o *LicenseLicenseReservationOpRelationship) SetRequestCode(v string)`

SetRequestCode sets RequestCode field to given value.

### HasRequestCode

`func (o *LicenseLicenseReservationOpRelationship) HasRequestCode() bool`

HasRequestCode returns a boolean if a field has been set.

### GetReturnCode

`func (o *LicenseLicenseReservationOpRelationship) GetReturnCode() string`

GetReturnCode returns the ReturnCode field if non-nil, zero value otherwise.

### GetReturnCodeOk

`func (o *LicenseLicenseReservationOpRelationship) GetReturnCodeOk() (*string, bool)`

GetReturnCodeOk returns a tuple with the ReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnCode

`func (o *LicenseLicenseReservationOpRelationship) SetReturnCode(v string)`

SetReturnCode sets ReturnCode field to given value.

### HasReturnCode

`func (o *LicenseLicenseReservationOpRelationship) HasReturnCode() bool`

HasReturnCode returns a boolean if a field has been set.

### GetAccount

`func (o *LicenseLicenseReservationOpRelationship) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *LicenseLicenseReservationOpRelationship) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *LicenseLicenseReservationOpRelationship) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *LicenseLicenseReservationOpRelationship) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


