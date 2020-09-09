# IamCertificateRequestRelationship

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
**EmailAddress** | Pointer to **string** | User input email address, an optional part of the subject of the certificate request. | [optional] 
**Name** | Pointer to **string** | Name of the certificate request. | [optional] 
**Request** | Pointer to **string** | Generated certificate signing request (CSR) in PEM format. | [optional] [readonly] 
**SelfSigned** | Pointer to **bool** | Whether the user wants the generated CSR to be self-signed by the appliance. | [optional] 
**Subject** | Pointer to [**PkixDistinguishedName**](pkix.DistinguishedName.md) |  | [optional] 
**SubjectAlternateName** | Pointer to [**PkixSubjectAlternateName**](pkix.SubjectAlternateName.md) |  | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 
**Certificate** | Pointer to [**IamCertificateRelationship**](iam.Certificate.Relationship.md) |  | [optional] 
**PrivateKeySpec** | Pointer to [**IamPrivateKeySpecRelationship**](iam.PrivateKeySpec.Relationship.md) |  | [optional] 

## Methods

### NewIamCertificateRequestRelationship

`func NewIamCertificateRequestRelationship(classId string, objectType string, ) *IamCertificateRequestRelationship`

NewIamCertificateRequestRelationship instantiates a new IamCertificateRequestRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamCertificateRequestRelationshipWithDefaults

`func NewIamCertificateRequestRelationshipWithDefaults() *IamCertificateRequestRelationship`

NewIamCertificateRequestRelationshipWithDefaults instantiates a new IamCertificateRequestRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *IamCertificateRequestRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *IamCertificateRequestRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *IamCertificateRequestRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *IamCertificateRequestRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *IamCertificateRequestRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamCertificateRequestRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamCertificateRequestRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *IamCertificateRequestRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *IamCertificateRequestRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *IamCertificateRequestRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *IamCertificateRequestRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *IamCertificateRequestRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *IamCertificateRequestRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *IamCertificateRequestRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *IamCertificateRequestRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *IamCertificateRequestRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *IamCertificateRequestRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *IamCertificateRequestRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *IamCertificateRequestRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *IamCertificateRequestRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *IamCertificateRequestRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *IamCertificateRequestRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *IamCertificateRequestRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *IamCertificateRequestRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamCertificateRequestRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamCertificateRequestRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *IamCertificateRequestRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *IamCertificateRequestRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *IamCertificateRequestRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *IamCertificateRequestRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *IamCertificateRequestRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *IamCertificateRequestRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *IamCertificateRequestRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *IamCertificateRequestRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *IamCertificateRequestRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IamCertificateRequestRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IamCertificateRequestRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IamCertificateRequestRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *IamCertificateRequestRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *IamCertificateRequestRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *IamCertificateRequestRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *IamCertificateRequestRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *IamCertificateRequestRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *IamCertificateRequestRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *IamCertificateRequestRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *IamCertificateRequestRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *IamCertificateRequestRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *IamCertificateRequestRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *IamCertificateRequestRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *IamCertificateRequestRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *IamCertificateRequestRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *IamCertificateRequestRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *IamCertificateRequestRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *IamCertificateRequestRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *IamCertificateRequestRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *IamCertificateRequestRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *IamCertificateRequestRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *IamCertificateRequestRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *IamCertificateRequestRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *IamCertificateRequestRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *IamCertificateRequestRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *IamCertificateRequestRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *IamCertificateRequestRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *IamCertificateRequestRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetEmailAddress

`func (o *IamCertificateRequestRelationship) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *IamCertificateRequestRelationship) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *IamCertificateRequestRelationship) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *IamCertificateRequestRelationship) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetName

`func (o *IamCertificateRequestRelationship) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamCertificateRequestRelationship) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamCertificateRequestRelationship) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamCertificateRequestRelationship) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRequest

`func (o *IamCertificateRequestRelationship) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *IamCertificateRequestRelationship) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *IamCertificateRequestRelationship) SetRequest(v string)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *IamCertificateRequestRelationship) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetSelfSigned

`func (o *IamCertificateRequestRelationship) GetSelfSigned() bool`

GetSelfSigned returns the SelfSigned field if non-nil, zero value otherwise.

### GetSelfSignedOk

`func (o *IamCertificateRequestRelationship) GetSelfSignedOk() (*bool, bool)`

GetSelfSignedOk returns a tuple with the SelfSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfSigned

`func (o *IamCertificateRequestRelationship) SetSelfSigned(v bool)`

SetSelfSigned sets SelfSigned field to given value.

### HasSelfSigned

`func (o *IamCertificateRequestRelationship) HasSelfSigned() bool`

HasSelfSigned returns a boolean if a field has been set.

### GetSubject

`func (o *IamCertificateRequestRelationship) GetSubject() PkixDistinguishedName`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *IamCertificateRequestRelationship) GetSubjectOk() (*PkixDistinguishedName, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *IamCertificateRequestRelationship) SetSubject(v PkixDistinguishedName)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *IamCertificateRequestRelationship) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetSubjectAlternateName

`func (o *IamCertificateRequestRelationship) GetSubjectAlternateName() PkixSubjectAlternateName`

GetSubjectAlternateName returns the SubjectAlternateName field if non-nil, zero value otherwise.

### GetSubjectAlternateNameOk

`func (o *IamCertificateRequestRelationship) GetSubjectAlternateNameOk() (*PkixSubjectAlternateName, bool)`

GetSubjectAlternateNameOk returns a tuple with the SubjectAlternateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAlternateName

`func (o *IamCertificateRequestRelationship) SetSubjectAlternateName(v PkixSubjectAlternateName)`

SetSubjectAlternateName sets SubjectAlternateName field to given value.

### HasSubjectAlternateName

`func (o *IamCertificateRequestRelationship) HasSubjectAlternateName() bool`

HasSubjectAlternateName returns a boolean if a field has been set.

### GetAccount

`func (o *IamCertificateRequestRelationship) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamCertificateRequestRelationship) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamCertificateRequestRelationship) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamCertificateRequestRelationship) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCertificate

`func (o *IamCertificateRequestRelationship) GetCertificate() IamCertificateRelationship`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *IamCertificateRequestRelationship) GetCertificateOk() (*IamCertificateRelationship, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *IamCertificateRequestRelationship) SetCertificate(v IamCertificateRelationship)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *IamCertificateRequestRelationship) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetPrivateKeySpec

`func (o *IamCertificateRequestRelationship) GetPrivateKeySpec() IamPrivateKeySpecRelationship`

GetPrivateKeySpec returns the PrivateKeySpec field if non-nil, zero value otherwise.

### GetPrivateKeySpecOk

`func (o *IamCertificateRequestRelationship) GetPrivateKeySpecOk() (*IamPrivateKeySpecRelationship, bool)`

GetPrivateKeySpecOk returns a tuple with the PrivateKeySpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeySpec

`func (o *IamCertificateRequestRelationship) SetPrivateKeySpec(v IamPrivateKeySpecRelationship)`

SetPrivateKeySpec sets PrivateKeySpec field to given value.

### HasPrivateKeySpec

`func (o *IamCertificateRequestRelationship) HasPrivateKeySpec() bool`

HasPrivateKeySpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


