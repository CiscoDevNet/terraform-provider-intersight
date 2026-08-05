# IamCuiIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "iam.CuiIntegration"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "iam.CuiIntegration"]
**CiscoCloudControlInventorySync** | Pointer to **string** | CCC inventory synchronization setting. * &#x60;enable&#x60; - Cisco Cloud Control inventory synchronization is enabled. * &#x60;disable&#x60; - Cisco Cloud Control inventory synchronization is disabled. | [optional] [readonly] [default to "enable"]
**LastFullSyncTime** | Pointer to **time.Time** | The last-modified timestamp of the tenant as returned by the remote API during the most recent workflow-based synchronization. | [optional] [readonly] 
**LastNotificationModifiedTime** | Pointer to **time.Time** | The modified-time carried by the most recent notification applied to this CUI integration MO, used to detect and skip stale or duplicate notifications. | [optional] [readonly] 
**LinkedTenants** | Pointer to [**[]IamTenantLinking**](IamTenantLinking.md) |  | [optional] 
**OptIn** | Pointer to **bool** | Indicates whether the account has opted in to CUI identity synchronization. | [optional] 
**OptInTime** | Pointer to **time.Time** | Timestamp representing when the customer opted in to CUI identity synchronization. | [optional] [readonly] 
**OptInUserIdentity** | Pointer to **string** | Identity or name of the engineer who performed the opt-in action. | [optional] [readonly] 
**SyncStatus** | Pointer to **string** | Current synchronization status of CUI integration data. * &#x60;Needs-Sync&#x60; - CUI integration data needs to be synchronized. * &#x60;InProgress&#x60; - CUI integration synchronization is in progress. * &#x60;Done&#x60; - CUI integration synchronization is completed. | [optional] [readonly] [default to "Needs-Sync"]
**TenantId** | Pointer to **string** | Primary tenant ID for CUI integration. | [optional] [readonly] 
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**User** | Pointer to [**NullableIamUserRelationship**](IamUserRelationship.md) |  | [optional] 

## Methods

### NewIamCuiIntegration

`func NewIamCuiIntegration(classId string, objectType string, ) *IamCuiIntegration`

NewIamCuiIntegration instantiates a new IamCuiIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamCuiIntegrationWithDefaults

`func NewIamCuiIntegrationWithDefaults() *IamCuiIntegration`

NewIamCuiIntegrationWithDefaults instantiates a new IamCuiIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IamCuiIntegration) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IamCuiIntegration) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IamCuiIntegration) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IamCuiIntegration) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IamCuiIntegration) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IamCuiIntegration) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCiscoCloudControlInventorySync

`func (o *IamCuiIntegration) GetCiscoCloudControlInventorySync() string`

GetCiscoCloudControlInventorySync returns the CiscoCloudControlInventorySync field if non-nil, zero value otherwise.

### GetCiscoCloudControlInventorySyncOk

`func (o *IamCuiIntegration) GetCiscoCloudControlInventorySyncOk() (*string, bool)`

GetCiscoCloudControlInventorySyncOk returns a tuple with the CiscoCloudControlInventorySync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiscoCloudControlInventorySync

`func (o *IamCuiIntegration) SetCiscoCloudControlInventorySync(v string)`

SetCiscoCloudControlInventorySync sets CiscoCloudControlInventorySync field to given value.

### HasCiscoCloudControlInventorySync

`func (o *IamCuiIntegration) HasCiscoCloudControlInventorySync() bool`

HasCiscoCloudControlInventorySync returns a boolean if a field has been set.

### GetLastFullSyncTime

`func (o *IamCuiIntegration) GetLastFullSyncTime() time.Time`

GetLastFullSyncTime returns the LastFullSyncTime field if non-nil, zero value otherwise.

### GetLastFullSyncTimeOk

`func (o *IamCuiIntegration) GetLastFullSyncTimeOk() (*time.Time, bool)`

GetLastFullSyncTimeOk returns a tuple with the LastFullSyncTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFullSyncTime

`func (o *IamCuiIntegration) SetLastFullSyncTime(v time.Time)`

SetLastFullSyncTime sets LastFullSyncTime field to given value.

### HasLastFullSyncTime

`func (o *IamCuiIntegration) HasLastFullSyncTime() bool`

HasLastFullSyncTime returns a boolean if a field has been set.

### GetLastNotificationModifiedTime

`func (o *IamCuiIntegration) GetLastNotificationModifiedTime() time.Time`

GetLastNotificationModifiedTime returns the LastNotificationModifiedTime field if non-nil, zero value otherwise.

### GetLastNotificationModifiedTimeOk

`func (o *IamCuiIntegration) GetLastNotificationModifiedTimeOk() (*time.Time, bool)`

GetLastNotificationModifiedTimeOk returns a tuple with the LastNotificationModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNotificationModifiedTime

`func (o *IamCuiIntegration) SetLastNotificationModifiedTime(v time.Time)`

SetLastNotificationModifiedTime sets LastNotificationModifiedTime field to given value.

### HasLastNotificationModifiedTime

`func (o *IamCuiIntegration) HasLastNotificationModifiedTime() bool`

HasLastNotificationModifiedTime returns a boolean if a field has been set.

### GetLinkedTenants

`func (o *IamCuiIntegration) GetLinkedTenants() []IamTenantLinking`

GetLinkedTenants returns the LinkedTenants field if non-nil, zero value otherwise.

### GetLinkedTenantsOk

`func (o *IamCuiIntegration) GetLinkedTenantsOk() (*[]IamTenantLinking, bool)`

GetLinkedTenantsOk returns a tuple with the LinkedTenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedTenants

`func (o *IamCuiIntegration) SetLinkedTenants(v []IamTenantLinking)`

SetLinkedTenants sets LinkedTenants field to given value.

### HasLinkedTenants

`func (o *IamCuiIntegration) HasLinkedTenants() bool`

HasLinkedTenants returns a boolean if a field has been set.

### SetLinkedTenantsNil

`func (o *IamCuiIntegration) SetLinkedTenantsNil(b bool)`

 SetLinkedTenantsNil sets the value for LinkedTenants to be an explicit nil

### UnsetLinkedTenants
`func (o *IamCuiIntegration) UnsetLinkedTenants()`

UnsetLinkedTenants ensures that no value is present for LinkedTenants, not even an explicit nil
### GetOptIn

`func (o *IamCuiIntegration) GetOptIn() bool`

GetOptIn returns the OptIn field if non-nil, zero value otherwise.

### GetOptInOk

`func (o *IamCuiIntegration) GetOptInOk() (*bool, bool)`

GetOptInOk returns a tuple with the OptIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptIn

`func (o *IamCuiIntegration) SetOptIn(v bool)`

SetOptIn sets OptIn field to given value.

### HasOptIn

`func (o *IamCuiIntegration) HasOptIn() bool`

HasOptIn returns a boolean if a field has been set.

### GetOptInTime

`func (o *IamCuiIntegration) GetOptInTime() time.Time`

GetOptInTime returns the OptInTime field if non-nil, zero value otherwise.

### GetOptInTimeOk

`func (o *IamCuiIntegration) GetOptInTimeOk() (*time.Time, bool)`

GetOptInTimeOk returns a tuple with the OptInTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptInTime

`func (o *IamCuiIntegration) SetOptInTime(v time.Time)`

SetOptInTime sets OptInTime field to given value.

### HasOptInTime

`func (o *IamCuiIntegration) HasOptInTime() bool`

HasOptInTime returns a boolean if a field has been set.

### GetOptInUserIdentity

`func (o *IamCuiIntegration) GetOptInUserIdentity() string`

GetOptInUserIdentity returns the OptInUserIdentity field if non-nil, zero value otherwise.

### GetOptInUserIdentityOk

`func (o *IamCuiIntegration) GetOptInUserIdentityOk() (*string, bool)`

GetOptInUserIdentityOk returns a tuple with the OptInUserIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptInUserIdentity

`func (o *IamCuiIntegration) SetOptInUserIdentity(v string)`

SetOptInUserIdentity sets OptInUserIdentity field to given value.

### HasOptInUserIdentity

`func (o *IamCuiIntegration) HasOptInUserIdentity() bool`

HasOptInUserIdentity returns a boolean if a field has been set.

### GetSyncStatus

`func (o *IamCuiIntegration) GetSyncStatus() string`

GetSyncStatus returns the SyncStatus field if non-nil, zero value otherwise.

### GetSyncStatusOk

`func (o *IamCuiIntegration) GetSyncStatusOk() (*string, bool)`

GetSyncStatusOk returns a tuple with the SyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncStatus

`func (o *IamCuiIntegration) SetSyncStatus(v string)`

SetSyncStatus sets SyncStatus field to given value.

### HasSyncStatus

`func (o *IamCuiIntegration) HasSyncStatus() bool`

HasSyncStatus returns a boolean if a field has been set.

### GetTenantId

`func (o *IamCuiIntegration) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *IamCuiIntegration) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *IamCuiIntegration) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *IamCuiIntegration) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetAccount

`func (o *IamCuiIntegration) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IamCuiIntegration) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IamCuiIntegration) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IamCuiIntegration) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *IamCuiIntegration) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *IamCuiIntegration) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetUser

`func (o *IamCuiIntegration) GetUser() IamUserRelationship`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *IamCuiIntegration) GetUserOk() (*IamUserRelationship, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *IamCuiIntegration) SetUser(v IamUserRelationship)`

SetUser sets User field to given value.

### HasUser

`func (o *IamCuiIntegration) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *IamCuiIntegration) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *IamCuiIntegration) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


