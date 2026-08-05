# BulkExport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "bulk.Export"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "bulk.Export"]
**Action** | Pointer to **string** | Action to be performed on the export operation. * &#x60;Start&#x60; - Starts the export operation. * &#x60;Cancel&#x60; - Cancels the export operation that is in progress. | [optional] [default to "Start"]
**ExcludePeers** | Pointer to **[]string** |  | [optional] 
**ExcludeRelations** | Pointer to **bool** | Used to specify that none of the relationships should be exported. | [optional] 
**ExportTags** | Pointer to **bool** | Specifies whether tags must be exported and will be considered for all the items MOs. | [optional] [default to true]
**ExportedObjects** | Pointer to [**[]BulkSubRequest**](BulkSubRequest.md) |  | [optional] 
**ExportedObjectsAsGroups** | Pointer to [**[]BulkSubRequestsGroup**](BulkSubRequestsGroup.md) |  | [optional] 
**ImportOrder** | Pointer to **interface{}** | Contains the list of import order. This field is deprecated and will be removed in future versions. | [optional] [readonly] 
**IncludeOrgIdentity** | Pointer to **bool** | Indicates that exported references for objects which are organization owned should include the organization reference along with the other identity properties. | [optional] 
**IsAesKeySet** | Pointer to **bool** | Indicates whether the value of the &#39;aesKey&#39; property has been set. | [optional] [readonly] [default to false]
**ItemNames** | Pointer to **[]string** |  | [optional] 
**Items** | Pointer to [**[]MoMoRef**](MoMoRef.md) |  | [optional] 
**Name** | Pointer to **string** | An identifier for the export instance. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_). | [optional] 
**PermissionId** | Pointer to **string** | The permission identifier which indicates the permission that current user has that will allow to start this export operation. | [optional] [readonly] 
**PreserveIdentities** | Pointer to **bool** | The flag set by the user during a configuration backup to preserve static or dynamic IDs assigned to an export item. | [optional] 
**RelatedTypeOptions** | Pointer to [**[]BulkRelatedTypeExportOption**](BulkRelatedTypeExportOption.md) |  | [optional] 
**SplitRequests** | Pointer to **bool** | Intersight supports the import operation using bulk requests. A single bulk request is created using the bulk sub requests created in export operation. There is a restriction of a single API request size in Intersight. If the exported MOs are many in an export operation, this flag will store the bulk sub requests as two dimensional collection property exportedObjectsAsGroups instead of one dimensional exportedObjects value where the sub requests are split into multiple smaller groups. Each group can be sent in a single bulk request during import operation. | [optional] 
**Status** | Pointer to **string** | Status of the export operation. * &#x60;&#x60; - The operation has not started. * &#x60;InProgress&#x60; - The operation is in progress. * &#x60;OrderInProgress&#x60; - The archive operation is in progress. * &#x60;Success&#x60; - The operation has succeeded. * &#x60;Failed&#x60; - The operation has failed. * &#x60;OperationTimedOut&#x60; - The operation has timed out. * &#x60;OperationCancelled&#x60; - The operation has been cancelled. * &#x60;CancelInProgress&#x60; - The operation is being cancelled. | [optional] [readonly] [default to ""]
**StatusMessage** | Pointer to **string** | Status message associated with failures or progress indication. | [optional] [readonly] 
**UserId** | Pointer to **string** | The user identifier which indicates the user that started this export operation. | [optional] [readonly] 
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**ExportedItems** | Pointer to [**[]BulkExportedItemRelationship**](BulkExportedItemRelationship.md) | An array of relationships to bulkExportedItem resources. | [optional] [readonly] 
**Idp** | Pointer to [**NullableIamIdpRelationship**](IamIdpRelationship.md) |  | [optional] 
**IdpReference** | Pointer to [**NullableIamIdpReferenceRelationship**](IamIdpReferenceRelationship.md) |  | [optional] 
**Organization** | Pointer to [**NullableOrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**User** | Pointer to [**NullableIamUserRelationship**](IamUserRelationship.md) |  | [optional] 

## Methods

### NewBulkExport

`func NewBulkExport(classId string, objectType string, ) *BulkExport`

NewBulkExport instantiates a new BulkExport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkExportWithDefaults

`func NewBulkExportWithDefaults() *BulkExport`

NewBulkExportWithDefaults instantiates a new BulkExport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *BulkExport) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *BulkExport) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *BulkExport) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *BulkExport) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BulkExport) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BulkExport) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *BulkExport) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *BulkExport) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *BulkExport) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *BulkExport) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetExcludePeers

`func (o *BulkExport) GetExcludePeers() []string`

GetExcludePeers returns the ExcludePeers field if non-nil, zero value otherwise.

### GetExcludePeersOk

`func (o *BulkExport) GetExcludePeersOk() (*[]string, bool)`

GetExcludePeersOk returns a tuple with the ExcludePeers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludePeers

`func (o *BulkExport) SetExcludePeers(v []string)`

SetExcludePeers sets ExcludePeers field to given value.

### HasExcludePeers

`func (o *BulkExport) HasExcludePeers() bool`

HasExcludePeers returns a boolean if a field has been set.

### SetExcludePeersNil

`func (o *BulkExport) SetExcludePeersNil(b bool)`

 SetExcludePeersNil sets the value for ExcludePeers to be an explicit nil

### UnsetExcludePeers
`func (o *BulkExport) UnsetExcludePeers()`

UnsetExcludePeers ensures that no value is present for ExcludePeers, not even an explicit nil
### GetExcludeRelations

`func (o *BulkExport) GetExcludeRelations() bool`

GetExcludeRelations returns the ExcludeRelations field if non-nil, zero value otherwise.

### GetExcludeRelationsOk

`func (o *BulkExport) GetExcludeRelationsOk() (*bool, bool)`

GetExcludeRelationsOk returns a tuple with the ExcludeRelations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeRelations

`func (o *BulkExport) SetExcludeRelations(v bool)`

SetExcludeRelations sets ExcludeRelations field to given value.

### HasExcludeRelations

`func (o *BulkExport) HasExcludeRelations() bool`

HasExcludeRelations returns a boolean if a field has been set.

### GetExportTags

`func (o *BulkExport) GetExportTags() bool`

GetExportTags returns the ExportTags field if non-nil, zero value otherwise.

### GetExportTagsOk

`func (o *BulkExport) GetExportTagsOk() (*bool, bool)`

GetExportTagsOk returns a tuple with the ExportTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportTags

`func (o *BulkExport) SetExportTags(v bool)`

SetExportTags sets ExportTags field to given value.

### HasExportTags

`func (o *BulkExport) HasExportTags() bool`

HasExportTags returns a boolean if a field has been set.

### GetExportedObjects

`func (o *BulkExport) GetExportedObjects() []BulkSubRequest`

GetExportedObjects returns the ExportedObjects field if non-nil, zero value otherwise.

### GetExportedObjectsOk

`func (o *BulkExport) GetExportedObjectsOk() (*[]BulkSubRequest, bool)`

GetExportedObjectsOk returns a tuple with the ExportedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportedObjects

`func (o *BulkExport) SetExportedObjects(v []BulkSubRequest)`

SetExportedObjects sets ExportedObjects field to given value.

### HasExportedObjects

`func (o *BulkExport) HasExportedObjects() bool`

HasExportedObjects returns a boolean if a field has been set.

### SetExportedObjectsNil

`func (o *BulkExport) SetExportedObjectsNil(b bool)`

 SetExportedObjectsNil sets the value for ExportedObjects to be an explicit nil

### UnsetExportedObjects
`func (o *BulkExport) UnsetExportedObjects()`

UnsetExportedObjects ensures that no value is present for ExportedObjects, not even an explicit nil
### GetExportedObjectsAsGroups

`func (o *BulkExport) GetExportedObjectsAsGroups() []BulkSubRequestsGroup`

GetExportedObjectsAsGroups returns the ExportedObjectsAsGroups field if non-nil, zero value otherwise.

### GetExportedObjectsAsGroupsOk

`func (o *BulkExport) GetExportedObjectsAsGroupsOk() (*[]BulkSubRequestsGroup, bool)`

GetExportedObjectsAsGroupsOk returns a tuple with the ExportedObjectsAsGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportedObjectsAsGroups

`func (o *BulkExport) SetExportedObjectsAsGroups(v []BulkSubRequestsGroup)`

SetExportedObjectsAsGroups sets ExportedObjectsAsGroups field to given value.

### HasExportedObjectsAsGroups

`func (o *BulkExport) HasExportedObjectsAsGroups() bool`

HasExportedObjectsAsGroups returns a boolean if a field has been set.

### SetExportedObjectsAsGroupsNil

`func (o *BulkExport) SetExportedObjectsAsGroupsNil(b bool)`

 SetExportedObjectsAsGroupsNil sets the value for ExportedObjectsAsGroups to be an explicit nil

### UnsetExportedObjectsAsGroups
`func (o *BulkExport) UnsetExportedObjectsAsGroups()`

UnsetExportedObjectsAsGroups ensures that no value is present for ExportedObjectsAsGroups, not even an explicit nil
### GetImportOrder

`func (o *BulkExport) GetImportOrder() interface{}`

GetImportOrder returns the ImportOrder field if non-nil, zero value otherwise.

### GetImportOrderOk

`func (o *BulkExport) GetImportOrderOk() (*interface{}, bool)`

GetImportOrderOk returns a tuple with the ImportOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportOrder

`func (o *BulkExport) SetImportOrder(v interface{})`

SetImportOrder sets ImportOrder field to given value.

### HasImportOrder

`func (o *BulkExport) HasImportOrder() bool`

HasImportOrder returns a boolean if a field has been set.

### SetImportOrderNil

`func (o *BulkExport) SetImportOrderNil(b bool)`

 SetImportOrderNil sets the value for ImportOrder to be an explicit nil

### UnsetImportOrder
`func (o *BulkExport) UnsetImportOrder()`

UnsetImportOrder ensures that no value is present for ImportOrder, not even an explicit nil
### GetIncludeOrgIdentity

`func (o *BulkExport) GetIncludeOrgIdentity() bool`

GetIncludeOrgIdentity returns the IncludeOrgIdentity field if non-nil, zero value otherwise.

### GetIncludeOrgIdentityOk

`func (o *BulkExport) GetIncludeOrgIdentityOk() (*bool, bool)`

GetIncludeOrgIdentityOk returns a tuple with the IncludeOrgIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeOrgIdentity

`func (o *BulkExport) SetIncludeOrgIdentity(v bool)`

SetIncludeOrgIdentity sets IncludeOrgIdentity field to given value.

### HasIncludeOrgIdentity

`func (o *BulkExport) HasIncludeOrgIdentity() bool`

HasIncludeOrgIdentity returns a boolean if a field has been set.

### GetIsAesKeySet

`func (o *BulkExport) GetIsAesKeySet() bool`

GetIsAesKeySet returns the IsAesKeySet field if non-nil, zero value otherwise.

### GetIsAesKeySetOk

`func (o *BulkExport) GetIsAesKeySetOk() (*bool, bool)`

GetIsAesKeySetOk returns a tuple with the IsAesKeySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAesKeySet

`func (o *BulkExport) SetIsAesKeySet(v bool)`

SetIsAesKeySet sets IsAesKeySet field to given value.

### HasIsAesKeySet

`func (o *BulkExport) HasIsAesKeySet() bool`

HasIsAesKeySet returns a boolean if a field has been set.

### GetItemNames

`func (o *BulkExport) GetItemNames() []string`

GetItemNames returns the ItemNames field if non-nil, zero value otherwise.

### GetItemNamesOk

`func (o *BulkExport) GetItemNamesOk() (*[]string, bool)`

GetItemNamesOk returns a tuple with the ItemNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemNames

`func (o *BulkExport) SetItemNames(v []string)`

SetItemNames sets ItemNames field to given value.

### HasItemNames

`func (o *BulkExport) HasItemNames() bool`

HasItemNames returns a boolean if a field has been set.

### SetItemNamesNil

`func (o *BulkExport) SetItemNamesNil(b bool)`

 SetItemNamesNil sets the value for ItemNames to be an explicit nil

### UnsetItemNames
`func (o *BulkExport) UnsetItemNames()`

UnsetItemNames ensures that no value is present for ItemNames, not even an explicit nil
### GetItems

`func (o *BulkExport) GetItems() []MoMoRef`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BulkExport) GetItemsOk() (*[]MoMoRef, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BulkExport) SetItems(v []MoMoRef)`

SetItems sets Items field to given value.

### HasItems

`func (o *BulkExport) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *BulkExport) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *BulkExport) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetName

`func (o *BulkExport) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkExport) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkExport) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BulkExport) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPermissionId

`func (o *BulkExport) GetPermissionId() string`

GetPermissionId returns the PermissionId field if non-nil, zero value otherwise.

### GetPermissionIdOk

`func (o *BulkExport) GetPermissionIdOk() (*string, bool)`

GetPermissionIdOk returns a tuple with the PermissionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionId

`func (o *BulkExport) SetPermissionId(v string)`

SetPermissionId sets PermissionId field to given value.

### HasPermissionId

`func (o *BulkExport) HasPermissionId() bool`

HasPermissionId returns a boolean if a field has been set.

### GetPreserveIdentities

`func (o *BulkExport) GetPreserveIdentities() bool`

GetPreserveIdentities returns the PreserveIdentities field if non-nil, zero value otherwise.

### GetPreserveIdentitiesOk

`func (o *BulkExport) GetPreserveIdentitiesOk() (*bool, bool)`

GetPreserveIdentitiesOk returns a tuple with the PreserveIdentities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveIdentities

`func (o *BulkExport) SetPreserveIdentities(v bool)`

SetPreserveIdentities sets PreserveIdentities field to given value.

### HasPreserveIdentities

`func (o *BulkExport) HasPreserveIdentities() bool`

HasPreserveIdentities returns a boolean if a field has been set.

### GetRelatedTypeOptions

`func (o *BulkExport) GetRelatedTypeOptions() []BulkRelatedTypeExportOption`

GetRelatedTypeOptions returns the RelatedTypeOptions field if non-nil, zero value otherwise.

### GetRelatedTypeOptionsOk

`func (o *BulkExport) GetRelatedTypeOptionsOk() (*[]BulkRelatedTypeExportOption, bool)`

GetRelatedTypeOptionsOk returns a tuple with the RelatedTypeOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedTypeOptions

`func (o *BulkExport) SetRelatedTypeOptions(v []BulkRelatedTypeExportOption)`

SetRelatedTypeOptions sets RelatedTypeOptions field to given value.

### HasRelatedTypeOptions

`func (o *BulkExport) HasRelatedTypeOptions() bool`

HasRelatedTypeOptions returns a boolean if a field has been set.

### SetRelatedTypeOptionsNil

`func (o *BulkExport) SetRelatedTypeOptionsNil(b bool)`

 SetRelatedTypeOptionsNil sets the value for RelatedTypeOptions to be an explicit nil

### UnsetRelatedTypeOptions
`func (o *BulkExport) UnsetRelatedTypeOptions()`

UnsetRelatedTypeOptions ensures that no value is present for RelatedTypeOptions, not even an explicit nil
### GetSplitRequests

`func (o *BulkExport) GetSplitRequests() bool`

GetSplitRequests returns the SplitRequests field if non-nil, zero value otherwise.

### GetSplitRequestsOk

`func (o *BulkExport) GetSplitRequestsOk() (*bool, bool)`

GetSplitRequestsOk returns a tuple with the SplitRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitRequests

`func (o *BulkExport) SetSplitRequests(v bool)`

SetSplitRequests sets SplitRequests field to given value.

### HasSplitRequests

`func (o *BulkExport) HasSplitRequests() bool`

HasSplitRequests returns a boolean if a field has been set.

### GetStatus

`func (o *BulkExport) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkExport) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkExport) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BulkExport) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *BulkExport) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *BulkExport) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *BulkExport) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *BulkExport) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetUserId

`func (o *BulkExport) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BulkExport) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BulkExport) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *BulkExport) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAccount

`func (o *BulkExport) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *BulkExport) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *BulkExport) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *BulkExport) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *BulkExport) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *BulkExport) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetExportedItems

`func (o *BulkExport) GetExportedItems() []BulkExportedItemRelationship`

GetExportedItems returns the ExportedItems field if non-nil, zero value otherwise.

### GetExportedItemsOk

`func (o *BulkExport) GetExportedItemsOk() (*[]BulkExportedItemRelationship, bool)`

GetExportedItemsOk returns a tuple with the ExportedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportedItems

`func (o *BulkExport) SetExportedItems(v []BulkExportedItemRelationship)`

SetExportedItems sets ExportedItems field to given value.

### HasExportedItems

`func (o *BulkExport) HasExportedItems() bool`

HasExportedItems returns a boolean if a field has been set.

### SetExportedItemsNil

`func (o *BulkExport) SetExportedItemsNil(b bool)`

 SetExportedItemsNil sets the value for ExportedItems to be an explicit nil

### UnsetExportedItems
`func (o *BulkExport) UnsetExportedItems()`

UnsetExportedItems ensures that no value is present for ExportedItems, not even an explicit nil
### GetIdp

`func (o *BulkExport) GetIdp() IamIdpRelationship`

GetIdp returns the Idp field if non-nil, zero value otherwise.

### GetIdpOk

`func (o *BulkExport) GetIdpOk() (*IamIdpRelationship, bool)`

GetIdpOk returns a tuple with the Idp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdp

`func (o *BulkExport) SetIdp(v IamIdpRelationship)`

SetIdp sets Idp field to given value.

### HasIdp

`func (o *BulkExport) HasIdp() bool`

HasIdp returns a boolean if a field has been set.

### SetIdpNil

`func (o *BulkExport) SetIdpNil(b bool)`

 SetIdpNil sets the value for Idp to be an explicit nil

### UnsetIdp
`func (o *BulkExport) UnsetIdp()`

UnsetIdp ensures that no value is present for Idp, not even an explicit nil
### GetIdpReference

`func (o *BulkExport) GetIdpReference() IamIdpReferenceRelationship`

GetIdpReference returns the IdpReference field if non-nil, zero value otherwise.

### GetIdpReferenceOk

`func (o *BulkExport) GetIdpReferenceOk() (*IamIdpReferenceRelationship, bool)`

GetIdpReferenceOk returns a tuple with the IdpReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpReference

`func (o *BulkExport) SetIdpReference(v IamIdpReferenceRelationship)`

SetIdpReference sets IdpReference field to given value.

### HasIdpReference

`func (o *BulkExport) HasIdpReference() bool`

HasIdpReference returns a boolean if a field has been set.

### SetIdpReferenceNil

`func (o *BulkExport) SetIdpReferenceNil(b bool)`

 SetIdpReferenceNil sets the value for IdpReference to be an explicit nil

### UnsetIdpReference
`func (o *BulkExport) UnsetIdpReference()`

UnsetIdpReference ensures that no value is present for IdpReference, not even an explicit nil
### GetOrganization

`func (o *BulkExport) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *BulkExport) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *BulkExport) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *BulkExport) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *BulkExport) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *BulkExport) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetUser

`func (o *BulkExport) GetUser() IamUserRelationship`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *BulkExport) GetUserOk() (*IamUserRelationship, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *BulkExport) SetUser(v IamUserRelationship)`

SetUser sets User field to given value.

### HasUser

`func (o *BulkExport) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *BulkExport) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *BulkExport) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


