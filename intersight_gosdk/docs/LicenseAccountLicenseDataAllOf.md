# LicenseAccountLicenseDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "license.AccountLicenseData"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "license.AccountLicenseData"]
**AccountId** | Pointer to **string** | Root user&#39;s ID of the account. | [optional] [readonly] 
**AgentData** | Pointer to **string** | Agent trusted store data. | [optional] [readonly] 
**AuthExpireTime** | Pointer to **string** | Authorization expiration time. | [optional] [readonly] 
**AuthInitialTime** | Pointer to **string** | Intial authorization time. | [optional] [readonly] 
**AuthNextTime** | Pointer to **string** | Next time for the authorization. | [optional] [readonly] 
**Category** | Pointer to **string** | Account license data category name. | [optional] [readonly] 
**DefaultLicenseType** | Pointer to **string** | Default license tier set by user. * &#x60;Base&#x60; - Base as a License type. It is default license type. * &#x60;Essential&#x60; - Essential as a License type. * &#x60;Standard&#x60; - Standard as a License type. * &#x60;Advantage&#x60; - Advantage as a License type. * &#x60;Premier&#x60; - Premier as a License type. * &#x60;IWO-Essential&#x60; - IWO-Essential as a License type. * &#x60;IWO-Advantage&#x60; - IWO-Advantage as a License type. * &#x60;IWO-Premier&#x60; - IWO-Premier as a License type. | [optional] [default to "Base"]
**ErrorDesc** | Pointer to **string** | The detailed error message when there is any error related to license sync of this account. | [optional] [readonly] 
**Group** | Pointer to **string** | Account license data group name. | [optional] [readonly] 
**HighestCompliantLicenseTier** | Pointer to **string** | The highest license tier which is in compliant of this account. * &#x60;Base&#x60; - Base as a License type. It is default license type. * &#x60;Essential&#x60; - Essential as a License type. * &#x60;Standard&#x60; - Standard as a License type. * &#x60;Advantage&#x60; - Advantage as a License type. * &#x60;Premier&#x60; - Premier as a License type. * &#x60;IWO-Essential&#x60; - IWO-Essential as a License type. * &#x60;IWO-Advantage&#x60; - IWO-Advantage as a License type. * &#x60;IWO-Premier&#x60; - IWO-Premier as a License type. | [optional] [readonly] [default to "Base"]
**LastSync** | Pointer to **time.Time** | Specifies last sync time with SA. | [optional] [readonly] 
**LastUpdatedTime** | Pointer to **time.Time** | Record&#39;s last update datetime. | [optional] [readonly] 
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
**Account** | Pointer to [**IamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**CustomerOp** | Pointer to [**LicenseCustomerOpRelationship**](LicenseCustomerOpRelationship.md) |  | [optional] 
**IwoCustomerOp** | Pointer to [**LicenseIwoCustomerOpRelationship**](LicenseIwoCustomerOpRelationship.md) |  | [optional] 
**IwoLicenseCount** | Pointer to [**LicenseIwoLicenseCountRelationship**](LicenseIwoLicenseCountRelationship.md) |  | [optional] 
**Licenseinfos** | Pointer to [**[]LicenseLicenseInfoRelationship**](LicenseLicenseInfoRelationship.md) | An array of relationships to licenseLicenseInfo resources. | [optional] 
**SmartlicenseToken** | Pointer to [**LicenseSmartlicenseTokenRelationship**](LicenseSmartlicenseTokenRelationship.md) |  | [optional] 

## Methods

### NewLicenseAccountLicenseDataAllOf

`func NewLicenseAccountLicenseDataAllOf(classId string, objectType string, ) *LicenseAccountLicenseDataAllOf`

NewLicenseAccountLicenseDataAllOf instantiates a new LicenseAccountLicenseDataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseAccountLicenseDataAllOfWithDefaults

`func NewLicenseAccountLicenseDataAllOfWithDefaults() *LicenseAccountLicenseDataAllOf`

NewLicenseAccountLicenseDataAllOfWithDefaults instantiates a new LicenseAccountLicenseDataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *LicenseAccountLicenseDataAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *LicenseAccountLicenseDataAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *LicenseAccountLicenseDataAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *LicenseAccountLicenseDataAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *LicenseAccountLicenseDataAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *LicenseAccountLicenseDataAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccountId

`func (o *LicenseAccountLicenseDataAllOf) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *LicenseAccountLicenseDataAllOf) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *LicenseAccountLicenseDataAllOf) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *LicenseAccountLicenseDataAllOf) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAgentData

`func (o *LicenseAccountLicenseDataAllOf) GetAgentData() string`

GetAgentData returns the AgentData field if non-nil, zero value otherwise.

### GetAgentDataOk

`func (o *LicenseAccountLicenseDataAllOf) GetAgentDataOk() (*string, bool)`

GetAgentDataOk returns a tuple with the AgentData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentData

`func (o *LicenseAccountLicenseDataAllOf) SetAgentData(v string)`

SetAgentData sets AgentData field to given value.

### HasAgentData

`func (o *LicenseAccountLicenseDataAllOf) HasAgentData() bool`

HasAgentData returns a boolean if a field has been set.

### GetAuthExpireTime

`func (o *LicenseAccountLicenseDataAllOf) GetAuthExpireTime() string`

GetAuthExpireTime returns the AuthExpireTime field if non-nil, zero value otherwise.

### GetAuthExpireTimeOk

`func (o *LicenseAccountLicenseDataAllOf) GetAuthExpireTimeOk() (*string, bool)`

GetAuthExpireTimeOk returns a tuple with the AuthExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthExpireTime

`func (o *LicenseAccountLicenseDataAllOf) SetAuthExpireTime(v string)`

SetAuthExpireTime sets AuthExpireTime field to given value.

### HasAuthExpireTime

`func (o *LicenseAccountLicenseDataAllOf) HasAuthExpireTime() bool`

HasAuthExpireTime returns a boolean if a field has been set.

### GetAuthInitialTime

`func (o *LicenseAccountLicenseDataAllOf) GetAuthInitialTime() string`

GetAuthInitialTime returns the AuthInitialTime field if non-nil, zero value otherwise.

### GetAuthInitialTimeOk

`func (o *LicenseAccountLicenseDataAllOf) GetAuthInitialTimeOk() (*string, bool)`

GetAuthInitialTimeOk returns a tuple with the AuthInitialTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthInitialTime

`func (o *LicenseAccountLicenseDataAllOf) SetAuthInitialTime(v string)`

SetAuthInitialTime sets AuthInitialTime field to given value.

### HasAuthInitialTime

`func (o *LicenseAccountLicenseDataAllOf) HasAuthInitialTime() bool`

HasAuthInitialTime returns a boolean if a field has been set.

### GetAuthNextTime

`func (o *LicenseAccountLicenseDataAllOf) GetAuthNextTime() string`

GetAuthNextTime returns the AuthNextTime field if non-nil, zero value otherwise.

### GetAuthNextTimeOk

`func (o *LicenseAccountLicenseDataAllOf) GetAuthNextTimeOk() (*string, bool)`

GetAuthNextTimeOk returns a tuple with the AuthNextTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthNextTime

`func (o *LicenseAccountLicenseDataAllOf) SetAuthNextTime(v string)`

SetAuthNextTime sets AuthNextTime field to given value.

### HasAuthNextTime

`func (o *LicenseAccountLicenseDataAllOf) HasAuthNextTime() bool`

HasAuthNextTime returns a boolean if a field has been set.

### GetCategory

`func (o *LicenseAccountLicenseDataAllOf) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *LicenseAccountLicenseDataAllOf) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *LicenseAccountLicenseDataAllOf) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *LicenseAccountLicenseDataAllOf) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDefaultLicenseType

`func (o *LicenseAccountLicenseDataAllOf) GetDefaultLicenseType() string`

GetDefaultLicenseType returns the DefaultLicenseType field if non-nil, zero value otherwise.

### GetDefaultLicenseTypeOk

`func (o *LicenseAccountLicenseDataAllOf) GetDefaultLicenseTypeOk() (*string, bool)`

GetDefaultLicenseTypeOk returns a tuple with the DefaultLicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLicenseType

`func (o *LicenseAccountLicenseDataAllOf) SetDefaultLicenseType(v string)`

SetDefaultLicenseType sets DefaultLicenseType field to given value.

### HasDefaultLicenseType

`func (o *LicenseAccountLicenseDataAllOf) HasDefaultLicenseType() bool`

HasDefaultLicenseType returns a boolean if a field has been set.

### GetErrorDesc

`func (o *LicenseAccountLicenseDataAllOf) GetErrorDesc() string`

GetErrorDesc returns the ErrorDesc field if non-nil, zero value otherwise.

### GetErrorDescOk

`func (o *LicenseAccountLicenseDataAllOf) GetErrorDescOk() (*string, bool)`

GetErrorDescOk returns a tuple with the ErrorDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDesc

`func (o *LicenseAccountLicenseDataAllOf) SetErrorDesc(v string)`

SetErrorDesc sets ErrorDesc field to given value.

### HasErrorDesc

`func (o *LicenseAccountLicenseDataAllOf) HasErrorDesc() bool`

HasErrorDesc returns a boolean if a field has been set.

### GetGroup

`func (o *LicenseAccountLicenseDataAllOf) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *LicenseAccountLicenseDataAllOf) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *LicenseAccountLicenseDataAllOf) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *LicenseAccountLicenseDataAllOf) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetHighestCompliantLicenseTier

`func (o *LicenseAccountLicenseDataAllOf) GetHighestCompliantLicenseTier() string`

GetHighestCompliantLicenseTier returns the HighestCompliantLicenseTier field if non-nil, zero value otherwise.

### GetHighestCompliantLicenseTierOk

`func (o *LicenseAccountLicenseDataAllOf) GetHighestCompliantLicenseTierOk() (*string, bool)`

GetHighestCompliantLicenseTierOk returns a tuple with the HighestCompliantLicenseTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighestCompliantLicenseTier

`func (o *LicenseAccountLicenseDataAllOf) SetHighestCompliantLicenseTier(v string)`

SetHighestCompliantLicenseTier sets HighestCompliantLicenseTier field to given value.

### HasHighestCompliantLicenseTier

`func (o *LicenseAccountLicenseDataAllOf) HasHighestCompliantLicenseTier() bool`

HasHighestCompliantLicenseTier returns a boolean if a field has been set.

### GetLastSync

`func (o *LicenseAccountLicenseDataAllOf) GetLastSync() time.Time`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *LicenseAccountLicenseDataAllOf) GetLastSyncOk() (*time.Time, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *LicenseAccountLicenseDataAllOf) SetLastSync(v time.Time)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *LicenseAccountLicenseDataAllOf) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *LicenseAccountLicenseDataAllOf) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *LicenseAccountLicenseDataAllOf) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *LicenseAccountLicenseDataAllOf) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *LicenseAccountLicenseDataAllOf) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetLicenseState

`func (o *LicenseAccountLicenseDataAllOf) GetLicenseState() string`

GetLicenseState returns the LicenseState field if non-nil, zero value otherwise.

### GetLicenseStateOk

`func (o *LicenseAccountLicenseDataAllOf) GetLicenseStateOk() (*string, bool)`

GetLicenseStateOk returns a tuple with the LicenseState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseState

`func (o *LicenseAccountLicenseDataAllOf) SetLicenseState(v string)`

SetLicenseState sets LicenseState field to given value.

### HasLicenseState

`func (o *LicenseAccountLicenseDataAllOf) HasLicenseState() bool`

HasLicenseState returns a boolean if a field has been set.

### GetLicenseTechSupportInfo

`func (o *LicenseAccountLicenseDataAllOf) GetLicenseTechSupportInfo() string`

GetLicenseTechSupportInfo returns the LicenseTechSupportInfo field if non-nil, zero value otherwise.

### GetLicenseTechSupportInfoOk

`func (o *LicenseAccountLicenseDataAllOf) GetLicenseTechSupportInfoOk() (*string, bool)`

GetLicenseTechSupportInfoOk returns a tuple with the LicenseTechSupportInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseTechSupportInfo

`func (o *LicenseAccountLicenseDataAllOf) SetLicenseTechSupportInfo(v string)`

SetLicenseTechSupportInfo sets LicenseTechSupportInfo field to given value.

### HasLicenseTechSupportInfo

`func (o *LicenseAccountLicenseDataAllOf) HasLicenseTechSupportInfo() bool`

HasLicenseTechSupportInfo returns a boolean if a field has been set.

### GetRegisterExpireTime

`func (o *LicenseAccountLicenseDataAllOf) GetRegisterExpireTime() string`

GetRegisterExpireTime returns the RegisterExpireTime field if non-nil, zero value otherwise.

### GetRegisterExpireTimeOk

`func (o *LicenseAccountLicenseDataAllOf) GetRegisterExpireTimeOk() (*string, bool)`

GetRegisterExpireTimeOk returns a tuple with the RegisterExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterExpireTime

`func (o *LicenseAccountLicenseDataAllOf) SetRegisterExpireTime(v string)`

SetRegisterExpireTime sets RegisterExpireTime field to given value.

### HasRegisterExpireTime

`func (o *LicenseAccountLicenseDataAllOf) HasRegisterExpireTime() bool`

HasRegisterExpireTime returns a boolean if a field has been set.

### GetRegisterInitialTime

`func (o *LicenseAccountLicenseDataAllOf) GetRegisterInitialTime() string`

GetRegisterInitialTime returns the RegisterInitialTime field if non-nil, zero value otherwise.

### GetRegisterInitialTimeOk

`func (o *LicenseAccountLicenseDataAllOf) GetRegisterInitialTimeOk() (*string, bool)`

GetRegisterInitialTimeOk returns a tuple with the RegisterInitialTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterInitialTime

`func (o *LicenseAccountLicenseDataAllOf) SetRegisterInitialTime(v string)`

SetRegisterInitialTime sets RegisterInitialTime field to given value.

### HasRegisterInitialTime

`func (o *LicenseAccountLicenseDataAllOf) HasRegisterInitialTime() bool`

HasRegisterInitialTime returns a boolean if a field has been set.

### GetRegisterNextTime

`func (o *LicenseAccountLicenseDataAllOf) GetRegisterNextTime() string`

GetRegisterNextTime returns the RegisterNextTime field if non-nil, zero value otherwise.

### GetRegisterNextTimeOk

`func (o *LicenseAccountLicenseDataAllOf) GetRegisterNextTimeOk() (*string, bool)`

GetRegisterNextTimeOk returns a tuple with the RegisterNextTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterNextTime

`func (o *LicenseAccountLicenseDataAllOf) SetRegisterNextTime(v string)`

SetRegisterNextTime sets RegisterNextTime field to given value.

### HasRegisterNextTime

`func (o *LicenseAccountLicenseDataAllOf) HasRegisterNextTime() bool`

HasRegisterNextTime returns a boolean if a field has been set.

### GetRegistrationStatus

`func (o *LicenseAccountLicenseDataAllOf) GetRegistrationStatus() string`

GetRegistrationStatus returns the RegistrationStatus field if non-nil, zero value otherwise.

### GetRegistrationStatusOk

`func (o *LicenseAccountLicenseDataAllOf) GetRegistrationStatusOk() (*string, bool)`

GetRegistrationStatusOk returns a tuple with the RegistrationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationStatus

`func (o *LicenseAccountLicenseDataAllOf) SetRegistrationStatus(v string)`

SetRegistrationStatus sets RegistrationStatus field to given value.

### HasRegistrationStatus

`func (o *LicenseAccountLicenseDataAllOf) HasRegistrationStatus() bool`

HasRegistrationStatus returns a boolean if a field has been set.

### GetRenewFailureString

`func (o *LicenseAccountLicenseDataAllOf) GetRenewFailureString() string`

GetRenewFailureString returns the RenewFailureString field if non-nil, zero value otherwise.

### GetRenewFailureStringOk

`func (o *LicenseAccountLicenseDataAllOf) GetRenewFailureStringOk() (*string, bool)`

GetRenewFailureStringOk returns a tuple with the RenewFailureString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewFailureString

`func (o *LicenseAccountLicenseDataAllOf) SetRenewFailureString(v string)`

SetRenewFailureString sets RenewFailureString field to given value.

### HasRenewFailureString

`func (o *LicenseAccountLicenseDataAllOf) HasRenewFailureString() bool`

HasRenewFailureString returns a boolean if a field has been set.

### GetSmartAccount

`func (o *LicenseAccountLicenseDataAllOf) GetSmartAccount() string`

GetSmartAccount returns the SmartAccount field if non-nil, zero value otherwise.

### GetSmartAccountOk

`func (o *LicenseAccountLicenseDataAllOf) GetSmartAccountOk() (*string, bool)`

GetSmartAccountOk returns a tuple with the SmartAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartAccount

`func (o *LicenseAccountLicenseDataAllOf) SetSmartAccount(v string)`

SetSmartAccount sets SmartAccount field to given value.

### HasSmartAccount

`func (o *LicenseAccountLicenseDataAllOf) HasSmartAccount() bool`

HasSmartAccount returns a boolean if a field has been set.

### GetSyncStatus

`func (o *LicenseAccountLicenseDataAllOf) GetSyncStatus() string`

GetSyncStatus returns the SyncStatus field if non-nil, zero value otherwise.

### GetSyncStatusOk

`func (o *LicenseAccountLicenseDataAllOf) GetSyncStatusOk() (*string, bool)`

GetSyncStatusOk returns a tuple with the SyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncStatus

`func (o *LicenseAccountLicenseDataAllOf) SetSyncStatus(v string)`

SetSyncStatus sets SyncStatus field to given value.

### HasSyncStatus

`func (o *LicenseAccountLicenseDataAllOf) HasSyncStatus() bool`

HasSyncStatus returns a boolean if a field has been set.

### GetVirtualAccount

`func (o *LicenseAccountLicenseDataAllOf) GetVirtualAccount() string`

GetVirtualAccount returns the VirtualAccount field if non-nil, zero value otherwise.

### GetVirtualAccountOk

`func (o *LicenseAccountLicenseDataAllOf) GetVirtualAccountOk() (*string, bool)`

GetVirtualAccountOk returns a tuple with the VirtualAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualAccount

`func (o *LicenseAccountLicenseDataAllOf) SetVirtualAccount(v string)`

SetVirtualAccount sets VirtualAccount field to given value.

### HasVirtualAccount

`func (o *LicenseAccountLicenseDataAllOf) HasVirtualAccount() bool`

HasVirtualAccount returns a boolean if a field has been set.

### GetAccount

`func (o *LicenseAccountLicenseDataAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *LicenseAccountLicenseDataAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *LicenseAccountLicenseDataAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *LicenseAccountLicenseDataAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCustomerOp

`func (o *LicenseAccountLicenseDataAllOf) GetCustomerOp() LicenseCustomerOpRelationship`

GetCustomerOp returns the CustomerOp field if non-nil, zero value otherwise.

### GetCustomerOpOk

`func (o *LicenseAccountLicenseDataAllOf) GetCustomerOpOk() (*LicenseCustomerOpRelationship, bool)`

GetCustomerOpOk returns a tuple with the CustomerOp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerOp

`func (o *LicenseAccountLicenseDataAllOf) SetCustomerOp(v LicenseCustomerOpRelationship)`

SetCustomerOp sets CustomerOp field to given value.

### HasCustomerOp

`func (o *LicenseAccountLicenseDataAllOf) HasCustomerOp() bool`

HasCustomerOp returns a boolean if a field has been set.

### GetIwoCustomerOp

`func (o *LicenseAccountLicenseDataAllOf) GetIwoCustomerOp() LicenseIwoCustomerOpRelationship`

GetIwoCustomerOp returns the IwoCustomerOp field if non-nil, zero value otherwise.

### GetIwoCustomerOpOk

`func (o *LicenseAccountLicenseDataAllOf) GetIwoCustomerOpOk() (*LicenseIwoCustomerOpRelationship, bool)`

GetIwoCustomerOpOk returns a tuple with the IwoCustomerOp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIwoCustomerOp

`func (o *LicenseAccountLicenseDataAllOf) SetIwoCustomerOp(v LicenseIwoCustomerOpRelationship)`

SetIwoCustomerOp sets IwoCustomerOp field to given value.

### HasIwoCustomerOp

`func (o *LicenseAccountLicenseDataAllOf) HasIwoCustomerOp() bool`

HasIwoCustomerOp returns a boolean if a field has been set.

### GetIwoLicenseCount

`func (o *LicenseAccountLicenseDataAllOf) GetIwoLicenseCount() LicenseIwoLicenseCountRelationship`

GetIwoLicenseCount returns the IwoLicenseCount field if non-nil, zero value otherwise.

### GetIwoLicenseCountOk

`func (o *LicenseAccountLicenseDataAllOf) GetIwoLicenseCountOk() (*LicenseIwoLicenseCountRelationship, bool)`

GetIwoLicenseCountOk returns a tuple with the IwoLicenseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIwoLicenseCount

`func (o *LicenseAccountLicenseDataAllOf) SetIwoLicenseCount(v LicenseIwoLicenseCountRelationship)`

SetIwoLicenseCount sets IwoLicenseCount field to given value.

### HasIwoLicenseCount

`func (o *LicenseAccountLicenseDataAllOf) HasIwoLicenseCount() bool`

HasIwoLicenseCount returns a boolean if a field has been set.

### GetLicenseinfos

`func (o *LicenseAccountLicenseDataAllOf) GetLicenseinfos() []LicenseLicenseInfoRelationship`

GetLicenseinfos returns the Licenseinfos field if non-nil, zero value otherwise.

### GetLicenseinfosOk

`func (o *LicenseAccountLicenseDataAllOf) GetLicenseinfosOk() (*[]LicenseLicenseInfoRelationship, bool)`

GetLicenseinfosOk returns a tuple with the Licenseinfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseinfos

`func (o *LicenseAccountLicenseDataAllOf) SetLicenseinfos(v []LicenseLicenseInfoRelationship)`

SetLicenseinfos sets Licenseinfos field to given value.

### HasLicenseinfos

`func (o *LicenseAccountLicenseDataAllOf) HasLicenseinfos() bool`

HasLicenseinfos returns a boolean if a field has been set.

### SetLicenseinfosNil

`func (o *LicenseAccountLicenseDataAllOf) SetLicenseinfosNil(b bool)`

 SetLicenseinfosNil sets the value for Licenseinfos to be an explicit nil

### UnsetLicenseinfos
`func (o *LicenseAccountLicenseDataAllOf) UnsetLicenseinfos()`

UnsetLicenseinfos ensures that no value is present for Licenseinfos, not even an explicit nil
### GetSmartlicenseToken

`func (o *LicenseAccountLicenseDataAllOf) GetSmartlicenseToken() LicenseSmartlicenseTokenRelationship`

GetSmartlicenseToken returns the SmartlicenseToken field if non-nil, zero value otherwise.

### GetSmartlicenseTokenOk

`func (o *LicenseAccountLicenseDataAllOf) GetSmartlicenseTokenOk() (*LicenseSmartlicenseTokenRelationship, bool)`

GetSmartlicenseTokenOk returns a tuple with the SmartlicenseToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartlicenseToken

`func (o *LicenseAccountLicenseDataAllOf) SetSmartlicenseToken(v LicenseSmartlicenseTokenRelationship)`

SetSmartlicenseToken sets SmartlicenseToken field to given value.

### HasSmartlicenseToken

`func (o *LicenseAccountLicenseDataAllOf) HasSmartlicenseToken() bool`

HasSmartlicenseToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


