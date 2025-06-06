# LicenseAccountLicenseData

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
**DefaultLicenseType** | Pointer to **string** | Default license tier set by the user. * &#x60;Base&#x60; - Base as a License type. It is default license type. * &#x60;Essential&#x60; - Essential as a License type. * &#x60;Standard&#x60; - Standard as a License type. * &#x60;Advantage&#x60; - Advantage as a License type. * &#x60;Premier&#x60; - Premier as a License type. * &#x60;IWO-Essential&#x60; - IWO-Essential as a License type. * &#x60;IWO-Advantage&#x60; - IWO-Advantage as a License type. * &#x60;IWO-Premier&#x60; - IWO-Premier as a License type. * &#x60;IKS-Advantage&#x60; - IKS-Advantage as a License type. * &#x60;INC-Premier-1GFixed&#x60; - Premier 1G Fixed license tier for Intersight Nexus Cloud. * &#x60;INC-Premier-10GFixed&#x60; - Premier 10G Fixed license tier for Intersight Nexus Cloud. * &#x60;INC-Premier-100GFixed&#x60; - Premier 100G Fixed license tier for Intersight Nexus Cloud. * &#x60;INC-Premier-Mod4Slot&#x60; - Premier Modular 4 slot license tier for Intersight Nexus Cloud. * &#x60;INC-Premier-Mod8Slot&#x60; - Premier Modular 8 slot license tier for Intersight Nexus Cloud. * &#x60;INC-Premier-D2OpsFixed&#x60; - Premier D2Ops fixed license tier for Intersight Nexus Cloud. * &#x60;INC-Premier-D2OpsMod&#x60; - Premier D2Ops modular license tier for Intersight Nexus Cloud. * &#x60;INC-Premier-CentralizedMod8Slot&#x60; - Premier modular license tier of switch type CentralizedMod8Slot for Intersight Nexus Cloud. * &#x60;INC-Premier-DistributedMod8Slot&#x60; - Premier modular license tier of switch type DistributedMod8Slot for Intersight Nexus Cloud. * &#x60;ERP-Advantage&#x60; - Advantage license tier for ERP workflows. * &#x60;IntersightTrial&#x60; - Virtual dummy license type to indicate trial. Used for UI display of trial mode Intersight tiers. * &#x60;IWOTrial&#x60; - Virtual dummy license type to indicate trial. Used for UI display of trial mode IKS tiers. * &#x60;IKSTrial&#x60; - Virtual dummy license type to indicate trial. Used for UI display of trial mode IWO tiers. * &#x60;INCTrial&#x60; - Virtual dummy license type to indicate trial. Used for UI display of trial mode Nexus tiers. | [optional] [default to "Base"]
**DefaultLicenseTypeNewerModels** | Pointer to **string** | Default license tier for newer model M7+ servers set by the user. * &#x60;Base&#x60; - Base as a License type. It is default license type. * &#x60;Essential&#x60; - Essential as a License type. * &#x60;Standard&#x60; - Standard as a License type. * &#x60;Advantage&#x60; - Advantage as a License type. * &#x60;Premier&#x60; - Premier as a License type. * &#x60;IWO-Essential&#x60; - IWO-Essential as a License type. * &#x60;IWO-Advantage&#x60; - IWO-Advantage as a License type. * &#x60;IWO-Premier&#x60; - IWO-Premier as a License type. * &#x60;IKS-Advantage&#x60; - IKS-Advantage as a License type. * &#x60;INC-Premier-1GFixed&#x60; - Premier 1G Fixed license tier for Intersight Nexus Cloud. * &#x60;INC-Premier-10GFixed&#x60; - Premier 10G Fixed license tier for Intersight Nexus Cloud. * &#x60;INC-Premier-100GFixed&#x60; - Premier 100G Fixed license tier for Intersight Nexus Cloud. * &#x60;INC-Premier-Mod4Slot&#x60; - Premier Modular 4 slot license tier for Intersight Nexus Cloud. * &#x60;INC-Premier-Mod8Slot&#x60; - Premier Modular 8 slot license tier for Intersight Nexus Cloud. * &#x60;INC-Premier-D2OpsFixed&#x60; - Premier D2Ops fixed license tier for Intersight Nexus Cloud. * &#x60;INC-Premier-D2OpsMod&#x60; - Premier D2Ops modular license tier for Intersight Nexus Cloud. * &#x60;INC-Premier-CentralizedMod8Slot&#x60; - Premier modular license tier of switch type CentralizedMod8Slot for Intersight Nexus Cloud. * &#x60;INC-Premier-DistributedMod8Slot&#x60; - Premier modular license tier of switch type DistributedMod8Slot for Intersight Nexus Cloud. * &#x60;ERP-Advantage&#x60; - Advantage license tier for ERP workflows. * &#x60;IntersightTrial&#x60; - Virtual dummy license type to indicate trial. Used for UI display of trial mode Intersight tiers. * &#x60;IWOTrial&#x60; - Virtual dummy license type to indicate trial. Used for UI display of trial mode IKS tiers. * &#x60;IKSTrial&#x60; - Virtual dummy license type to indicate trial. Used for UI display of trial mode IWO tiers. * &#x60;INCTrial&#x60; - Virtual dummy license type to indicate trial. Used for UI display of trial mode Nexus tiers. | [optional] [default to "Base"]
**ErrorDesc** | Pointer to **string** | The detailed error message when there is any error related to license sync of this account. | [optional] [readonly] 
**Group** | Pointer to **string** | Account license data group name. | [optional] [readonly] 
**HighestCompliantLicenseTier** | Pointer to **string** | The highest license tier which is in compliant of this account. * &#x60;Base&#x60; - Base as a License type. It is default license type. * &#x60;Essential&#x60; - Essential as a License type. * &#x60;Standard&#x60; - Standard as a License type. * &#x60;Advantage&#x60; - Advantage as a License type. * &#x60;Premier&#x60; - Premier as a License type. * &#x60;IWO-Essential&#x60; - IWO-Essential as a License type. * &#x60;IWO-Advantage&#x60; - IWO-Advantage as a License type. * &#x60;IWO-Premier&#x60; - IWO-Premier as a License type. * &#x60;IKS-Advantage&#x60; - IKS-Advantage as a License type. * &#x60;INC-Premier-1GFixed&#x60; - Premier 1G Fixed license tier for Intersight Nexus Cloud. * &#x60;INC-Premier-10GFixed&#x60; - Premier 10G Fixed license tier for Intersight Nexus Cloud. * &#x60;INC-Premier-100GFixed&#x60; - Premier 100G Fixed license tier for Intersight Nexus Cloud. * &#x60;INC-Premier-Mod4Slot&#x60; - Premier Modular 4 slot license tier for Intersight Nexus Cloud. * &#x60;INC-Premier-Mod8Slot&#x60; - Premier Modular 8 slot license tier for Intersight Nexus Cloud. * &#x60;INC-Premier-D2OpsFixed&#x60; - Premier D2Ops fixed license tier for Intersight Nexus Cloud. * &#x60;INC-Premier-D2OpsMod&#x60; - Premier D2Ops modular license tier for Intersight Nexus Cloud. * &#x60;INC-Premier-CentralizedMod8Slot&#x60; - Premier modular license tier of switch type CentralizedMod8Slot for Intersight Nexus Cloud. * &#x60;INC-Premier-DistributedMod8Slot&#x60; - Premier modular license tier of switch type DistributedMod8Slot for Intersight Nexus Cloud. * &#x60;ERP-Advantage&#x60; - Advantage license tier for ERP workflows. * &#x60;IntersightTrial&#x60; - Virtual dummy license type to indicate trial. Used for UI display of trial mode Intersight tiers. * &#x60;IWOTrial&#x60; - Virtual dummy license type to indicate trial. Used for UI display of trial mode IKS tiers. * &#x60;IKSTrial&#x60; - Virtual dummy license type to indicate trial. Used for UI display of trial mode IWO tiers. * &#x60;INCTrial&#x60; - Virtual dummy license type to indicate trial. Used for UI display of trial mode Nexus tiers. | [optional] [readonly] [default to "Base"]
**LastCssmSync** | Pointer to **time.Time** | Specifies last sync time with CSSM. | [optional] [readonly] 
**LastRenew** | Pointer to **time.Time** | Specifies last certificate renew time with SA. | [optional] [readonly] 
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
**SmartAccountDomain** | Pointer to **string** | Domain Name of the smart account. | [optional] [readonly] 
**SmartApiEnabled** | Pointer to **bool** | Indicate whether API integration is enabled. | [optional] [readonly] 
**SmartApiSyncStatus** | Pointer to **string** | The detailed error message when there is any smart API sync error related to this account. | [optional] [readonly] 
**SyncStatus** | Pointer to **string** | Current sync status for the account. | [optional] [readonly] 
**VirtualAccount** | Pointer to **string** | Name of the virtual account. | [optional] [readonly] 
**Account** | Pointer to [**NullableIamAccountRelationship**](IamAccountRelationship.md) |  | [optional] 
**CustomerOp** | Pointer to [**NullableLicenseCustomerOpRelationship**](LicenseCustomerOpRelationship.md) |  | [optional] 
**ErpCustomerOp** | Pointer to [**NullableLicenseErpCustomerOpRelationship**](LicenseErpCustomerOpRelationship.md) |  | [optional] 
**ErpLicenseCount** | Pointer to [**NullableLicenseErpLicenseCountRelationship**](LicenseErpLicenseCountRelationship.md) |  | [optional] 
**IksCustomerOp** | Pointer to [**NullableLicenseIksCustomerOpRelationship**](LicenseIksCustomerOpRelationship.md) |  | [optional] 
**IksLicenseCount** | Pointer to [**NullableLicenseIksLicenseCountRelationship**](LicenseIksLicenseCountRelationship.md) |  | [optional] 
**IncCustomerOp** | Pointer to [**NullableLicenseIncCustomerOpRelationship**](LicenseIncCustomerOpRelationship.md) |  | [optional] 
**IncLicenseCount** | Pointer to [**NullableLicenseIncLicenseCountRelationship**](LicenseIncLicenseCountRelationship.md) |  | [optional] 
**IwoCustomerOp** | Pointer to [**NullableLicenseIwoCustomerOpRelationship**](LicenseIwoCustomerOpRelationship.md) |  | [optional] 
**IwoLicenseCount** | Pointer to [**NullableLicenseIwoLicenseCountRelationship**](LicenseIwoLicenseCountRelationship.md) |  | [optional] 
**LicenseInfoView** | Pointer to [**NullableLicenseLicenseInfoViewRelationship**](LicenseLicenseInfoViewRelationship.md) |  | [optional] 
**LicenseRegistrationStatus** | Pointer to [**NullableLicenseLicenseRegistrationStatusRelationship**](LicenseLicenseRegistrationStatusRelationship.md) |  | [optional] 
**Licenseinfos** | Pointer to [**[]LicenseLicenseInfoRelationship**](LicenseLicenseInfoRelationship.md) | An array of relationships to licenseLicenseInfo resources. | [optional] 
**SmartlicenseToken** | Pointer to [**NullableLicenseSmartlicenseTokenRelationship**](LicenseSmartlicenseTokenRelationship.md) |  | [optional] 

## Methods

### NewLicenseAccountLicenseData

`func NewLicenseAccountLicenseData(classId string, objectType string, ) *LicenseAccountLicenseData`

NewLicenseAccountLicenseData instantiates a new LicenseAccountLicenseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseAccountLicenseDataWithDefaults

`func NewLicenseAccountLicenseDataWithDefaults() *LicenseAccountLicenseData`

NewLicenseAccountLicenseDataWithDefaults instantiates a new LicenseAccountLicenseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *LicenseAccountLicenseData) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *LicenseAccountLicenseData) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *LicenseAccountLicenseData) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *LicenseAccountLicenseData) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *LicenseAccountLicenseData) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *LicenseAccountLicenseData) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccountId

`func (o *LicenseAccountLicenseData) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *LicenseAccountLicenseData) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *LicenseAccountLicenseData) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *LicenseAccountLicenseData) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAgentData

`func (o *LicenseAccountLicenseData) GetAgentData() string`

GetAgentData returns the AgentData field if non-nil, zero value otherwise.

### GetAgentDataOk

`func (o *LicenseAccountLicenseData) GetAgentDataOk() (*string, bool)`

GetAgentDataOk returns a tuple with the AgentData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentData

`func (o *LicenseAccountLicenseData) SetAgentData(v string)`

SetAgentData sets AgentData field to given value.

### HasAgentData

`func (o *LicenseAccountLicenseData) HasAgentData() bool`

HasAgentData returns a boolean if a field has been set.

### GetAuthExpireTime

`func (o *LicenseAccountLicenseData) GetAuthExpireTime() string`

GetAuthExpireTime returns the AuthExpireTime field if non-nil, zero value otherwise.

### GetAuthExpireTimeOk

`func (o *LicenseAccountLicenseData) GetAuthExpireTimeOk() (*string, bool)`

GetAuthExpireTimeOk returns a tuple with the AuthExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthExpireTime

`func (o *LicenseAccountLicenseData) SetAuthExpireTime(v string)`

SetAuthExpireTime sets AuthExpireTime field to given value.

### HasAuthExpireTime

`func (o *LicenseAccountLicenseData) HasAuthExpireTime() bool`

HasAuthExpireTime returns a boolean if a field has been set.

### GetAuthInitialTime

`func (o *LicenseAccountLicenseData) GetAuthInitialTime() string`

GetAuthInitialTime returns the AuthInitialTime field if non-nil, zero value otherwise.

### GetAuthInitialTimeOk

`func (o *LicenseAccountLicenseData) GetAuthInitialTimeOk() (*string, bool)`

GetAuthInitialTimeOk returns a tuple with the AuthInitialTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthInitialTime

`func (o *LicenseAccountLicenseData) SetAuthInitialTime(v string)`

SetAuthInitialTime sets AuthInitialTime field to given value.

### HasAuthInitialTime

`func (o *LicenseAccountLicenseData) HasAuthInitialTime() bool`

HasAuthInitialTime returns a boolean if a field has been set.

### GetAuthNextTime

`func (o *LicenseAccountLicenseData) GetAuthNextTime() string`

GetAuthNextTime returns the AuthNextTime field if non-nil, zero value otherwise.

### GetAuthNextTimeOk

`func (o *LicenseAccountLicenseData) GetAuthNextTimeOk() (*string, bool)`

GetAuthNextTimeOk returns a tuple with the AuthNextTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthNextTime

`func (o *LicenseAccountLicenseData) SetAuthNextTime(v string)`

SetAuthNextTime sets AuthNextTime field to given value.

### HasAuthNextTime

`func (o *LicenseAccountLicenseData) HasAuthNextTime() bool`

HasAuthNextTime returns a boolean if a field has been set.

### GetCategory

`func (o *LicenseAccountLicenseData) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *LicenseAccountLicenseData) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *LicenseAccountLicenseData) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *LicenseAccountLicenseData) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDefaultLicenseType

`func (o *LicenseAccountLicenseData) GetDefaultLicenseType() string`

GetDefaultLicenseType returns the DefaultLicenseType field if non-nil, zero value otherwise.

### GetDefaultLicenseTypeOk

`func (o *LicenseAccountLicenseData) GetDefaultLicenseTypeOk() (*string, bool)`

GetDefaultLicenseTypeOk returns a tuple with the DefaultLicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLicenseType

`func (o *LicenseAccountLicenseData) SetDefaultLicenseType(v string)`

SetDefaultLicenseType sets DefaultLicenseType field to given value.

### HasDefaultLicenseType

`func (o *LicenseAccountLicenseData) HasDefaultLicenseType() bool`

HasDefaultLicenseType returns a boolean if a field has been set.

### GetDefaultLicenseTypeNewerModels

`func (o *LicenseAccountLicenseData) GetDefaultLicenseTypeNewerModels() string`

GetDefaultLicenseTypeNewerModels returns the DefaultLicenseTypeNewerModels field if non-nil, zero value otherwise.

### GetDefaultLicenseTypeNewerModelsOk

`func (o *LicenseAccountLicenseData) GetDefaultLicenseTypeNewerModelsOk() (*string, bool)`

GetDefaultLicenseTypeNewerModelsOk returns a tuple with the DefaultLicenseTypeNewerModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLicenseTypeNewerModels

`func (o *LicenseAccountLicenseData) SetDefaultLicenseTypeNewerModels(v string)`

SetDefaultLicenseTypeNewerModels sets DefaultLicenseTypeNewerModels field to given value.

### HasDefaultLicenseTypeNewerModels

`func (o *LicenseAccountLicenseData) HasDefaultLicenseTypeNewerModels() bool`

HasDefaultLicenseTypeNewerModels returns a boolean if a field has been set.

### GetErrorDesc

`func (o *LicenseAccountLicenseData) GetErrorDesc() string`

GetErrorDesc returns the ErrorDesc field if non-nil, zero value otherwise.

### GetErrorDescOk

`func (o *LicenseAccountLicenseData) GetErrorDescOk() (*string, bool)`

GetErrorDescOk returns a tuple with the ErrorDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDesc

`func (o *LicenseAccountLicenseData) SetErrorDesc(v string)`

SetErrorDesc sets ErrorDesc field to given value.

### HasErrorDesc

`func (o *LicenseAccountLicenseData) HasErrorDesc() bool`

HasErrorDesc returns a boolean if a field has been set.

### GetGroup

`func (o *LicenseAccountLicenseData) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *LicenseAccountLicenseData) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *LicenseAccountLicenseData) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *LicenseAccountLicenseData) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetHighestCompliantLicenseTier

`func (o *LicenseAccountLicenseData) GetHighestCompliantLicenseTier() string`

GetHighestCompliantLicenseTier returns the HighestCompliantLicenseTier field if non-nil, zero value otherwise.

### GetHighestCompliantLicenseTierOk

`func (o *LicenseAccountLicenseData) GetHighestCompliantLicenseTierOk() (*string, bool)`

GetHighestCompliantLicenseTierOk returns a tuple with the HighestCompliantLicenseTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighestCompliantLicenseTier

`func (o *LicenseAccountLicenseData) SetHighestCompliantLicenseTier(v string)`

SetHighestCompliantLicenseTier sets HighestCompliantLicenseTier field to given value.

### HasHighestCompliantLicenseTier

`func (o *LicenseAccountLicenseData) HasHighestCompliantLicenseTier() bool`

HasHighestCompliantLicenseTier returns a boolean if a field has been set.

### GetLastCssmSync

`func (o *LicenseAccountLicenseData) GetLastCssmSync() time.Time`

GetLastCssmSync returns the LastCssmSync field if non-nil, zero value otherwise.

### GetLastCssmSyncOk

`func (o *LicenseAccountLicenseData) GetLastCssmSyncOk() (*time.Time, bool)`

GetLastCssmSyncOk returns a tuple with the LastCssmSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCssmSync

`func (o *LicenseAccountLicenseData) SetLastCssmSync(v time.Time)`

SetLastCssmSync sets LastCssmSync field to given value.

### HasLastCssmSync

`func (o *LicenseAccountLicenseData) HasLastCssmSync() bool`

HasLastCssmSync returns a boolean if a field has been set.

### GetLastRenew

`func (o *LicenseAccountLicenseData) GetLastRenew() time.Time`

GetLastRenew returns the LastRenew field if non-nil, zero value otherwise.

### GetLastRenewOk

`func (o *LicenseAccountLicenseData) GetLastRenewOk() (*time.Time, bool)`

GetLastRenewOk returns a tuple with the LastRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRenew

`func (o *LicenseAccountLicenseData) SetLastRenew(v time.Time)`

SetLastRenew sets LastRenew field to given value.

### HasLastRenew

`func (o *LicenseAccountLicenseData) HasLastRenew() bool`

HasLastRenew returns a boolean if a field has been set.

### GetLastSync

`func (o *LicenseAccountLicenseData) GetLastSync() time.Time`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *LicenseAccountLicenseData) GetLastSyncOk() (*time.Time, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *LicenseAccountLicenseData) SetLastSync(v time.Time)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *LicenseAccountLicenseData) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### GetLastUpdatedTime

`func (o *LicenseAccountLicenseData) GetLastUpdatedTime() time.Time`

GetLastUpdatedTime returns the LastUpdatedTime field if non-nil, zero value otherwise.

### GetLastUpdatedTimeOk

`func (o *LicenseAccountLicenseData) GetLastUpdatedTimeOk() (*time.Time, bool)`

GetLastUpdatedTimeOk returns a tuple with the LastUpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTime

`func (o *LicenseAccountLicenseData) SetLastUpdatedTime(v time.Time)`

SetLastUpdatedTime sets LastUpdatedTime field to given value.

### HasLastUpdatedTime

`func (o *LicenseAccountLicenseData) HasLastUpdatedTime() bool`

HasLastUpdatedTime returns a boolean if a field has been set.

### GetLicenseState

`func (o *LicenseAccountLicenseData) GetLicenseState() string`

GetLicenseState returns the LicenseState field if non-nil, zero value otherwise.

### GetLicenseStateOk

`func (o *LicenseAccountLicenseData) GetLicenseStateOk() (*string, bool)`

GetLicenseStateOk returns a tuple with the LicenseState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseState

`func (o *LicenseAccountLicenseData) SetLicenseState(v string)`

SetLicenseState sets LicenseState field to given value.

### HasLicenseState

`func (o *LicenseAccountLicenseData) HasLicenseState() bool`

HasLicenseState returns a boolean if a field has been set.

### GetLicenseTechSupportInfo

`func (o *LicenseAccountLicenseData) GetLicenseTechSupportInfo() string`

GetLicenseTechSupportInfo returns the LicenseTechSupportInfo field if non-nil, zero value otherwise.

### GetLicenseTechSupportInfoOk

`func (o *LicenseAccountLicenseData) GetLicenseTechSupportInfoOk() (*string, bool)`

GetLicenseTechSupportInfoOk returns a tuple with the LicenseTechSupportInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseTechSupportInfo

`func (o *LicenseAccountLicenseData) SetLicenseTechSupportInfo(v string)`

SetLicenseTechSupportInfo sets LicenseTechSupportInfo field to given value.

### HasLicenseTechSupportInfo

`func (o *LicenseAccountLicenseData) HasLicenseTechSupportInfo() bool`

HasLicenseTechSupportInfo returns a boolean if a field has been set.

### GetRegisterExpireTime

`func (o *LicenseAccountLicenseData) GetRegisterExpireTime() string`

GetRegisterExpireTime returns the RegisterExpireTime field if non-nil, zero value otherwise.

### GetRegisterExpireTimeOk

`func (o *LicenseAccountLicenseData) GetRegisterExpireTimeOk() (*string, bool)`

GetRegisterExpireTimeOk returns a tuple with the RegisterExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterExpireTime

`func (o *LicenseAccountLicenseData) SetRegisterExpireTime(v string)`

SetRegisterExpireTime sets RegisterExpireTime field to given value.

### HasRegisterExpireTime

`func (o *LicenseAccountLicenseData) HasRegisterExpireTime() bool`

HasRegisterExpireTime returns a boolean if a field has been set.

### GetRegisterInitialTime

`func (o *LicenseAccountLicenseData) GetRegisterInitialTime() string`

GetRegisterInitialTime returns the RegisterInitialTime field if non-nil, zero value otherwise.

### GetRegisterInitialTimeOk

`func (o *LicenseAccountLicenseData) GetRegisterInitialTimeOk() (*string, bool)`

GetRegisterInitialTimeOk returns a tuple with the RegisterInitialTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterInitialTime

`func (o *LicenseAccountLicenseData) SetRegisterInitialTime(v string)`

SetRegisterInitialTime sets RegisterInitialTime field to given value.

### HasRegisterInitialTime

`func (o *LicenseAccountLicenseData) HasRegisterInitialTime() bool`

HasRegisterInitialTime returns a boolean if a field has been set.

### GetRegisterNextTime

`func (o *LicenseAccountLicenseData) GetRegisterNextTime() string`

GetRegisterNextTime returns the RegisterNextTime field if non-nil, zero value otherwise.

### GetRegisterNextTimeOk

`func (o *LicenseAccountLicenseData) GetRegisterNextTimeOk() (*string, bool)`

GetRegisterNextTimeOk returns a tuple with the RegisterNextTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterNextTime

`func (o *LicenseAccountLicenseData) SetRegisterNextTime(v string)`

SetRegisterNextTime sets RegisterNextTime field to given value.

### HasRegisterNextTime

`func (o *LicenseAccountLicenseData) HasRegisterNextTime() bool`

HasRegisterNextTime returns a boolean if a field has been set.

### GetRegistrationStatus

`func (o *LicenseAccountLicenseData) GetRegistrationStatus() string`

GetRegistrationStatus returns the RegistrationStatus field if non-nil, zero value otherwise.

### GetRegistrationStatusOk

`func (o *LicenseAccountLicenseData) GetRegistrationStatusOk() (*string, bool)`

GetRegistrationStatusOk returns a tuple with the RegistrationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationStatus

`func (o *LicenseAccountLicenseData) SetRegistrationStatus(v string)`

SetRegistrationStatus sets RegistrationStatus field to given value.

### HasRegistrationStatus

`func (o *LicenseAccountLicenseData) HasRegistrationStatus() bool`

HasRegistrationStatus returns a boolean if a field has been set.

### GetRenewFailureString

`func (o *LicenseAccountLicenseData) GetRenewFailureString() string`

GetRenewFailureString returns the RenewFailureString field if non-nil, zero value otherwise.

### GetRenewFailureStringOk

`func (o *LicenseAccountLicenseData) GetRenewFailureStringOk() (*string, bool)`

GetRenewFailureStringOk returns a tuple with the RenewFailureString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewFailureString

`func (o *LicenseAccountLicenseData) SetRenewFailureString(v string)`

SetRenewFailureString sets RenewFailureString field to given value.

### HasRenewFailureString

`func (o *LicenseAccountLicenseData) HasRenewFailureString() bool`

HasRenewFailureString returns a boolean if a field has been set.

### GetSmartAccount

`func (o *LicenseAccountLicenseData) GetSmartAccount() string`

GetSmartAccount returns the SmartAccount field if non-nil, zero value otherwise.

### GetSmartAccountOk

`func (o *LicenseAccountLicenseData) GetSmartAccountOk() (*string, bool)`

GetSmartAccountOk returns a tuple with the SmartAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartAccount

`func (o *LicenseAccountLicenseData) SetSmartAccount(v string)`

SetSmartAccount sets SmartAccount field to given value.

### HasSmartAccount

`func (o *LicenseAccountLicenseData) HasSmartAccount() bool`

HasSmartAccount returns a boolean if a field has been set.

### GetSmartAccountDomain

`func (o *LicenseAccountLicenseData) GetSmartAccountDomain() string`

GetSmartAccountDomain returns the SmartAccountDomain field if non-nil, zero value otherwise.

### GetSmartAccountDomainOk

`func (o *LicenseAccountLicenseData) GetSmartAccountDomainOk() (*string, bool)`

GetSmartAccountDomainOk returns a tuple with the SmartAccountDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartAccountDomain

`func (o *LicenseAccountLicenseData) SetSmartAccountDomain(v string)`

SetSmartAccountDomain sets SmartAccountDomain field to given value.

### HasSmartAccountDomain

`func (o *LicenseAccountLicenseData) HasSmartAccountDomain() bool`

HasSmartAccountDomain returns a boolean if a field has been set.

### GetSmartApiEnabled

`func (o *LicenseAccountLicenseData) GetSmartApiEnabled() bool`

GetSmartApiEnabled returns the SmartApiEnabled field if non-nil, zero value otherwise.

### GetSmartApiEnabledOk

`func (o *LicenseAccountLicenseData) GetSmartApiEnabledOk() (*bool, bool)`

GetSmartApiEnabledOk returns a tuple with the SmartApiEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartApiEnabled

`func (o *LicenseAccountLicenseData) SetSmartApiEnabled(v bool)`

SetSmartApiEnabled sets SmartApiEnabled field to given value.

### HasSmartApiEnabled

`func (o *LicenseAccountLicenseData) HasSmartApiEnabled() bool`

HasSmartApiEnabled returns a boolean if a field has been set.

### GetSmartApiSyncStatus

`func (o *LicenseAccountLicenseData) GetSmartApiSyncStatus() string`

GetSmartApiSyncStatus returns the SmartApiSyncStatus field if non-nil, zero value otherwise.

### GetSmartApiSyncStatusOk

`func (o *LicenseAccountLicenseData) GetSmartApiSyncStatusOk() (*string, bool)`

GetSmartApiSyncStatusOk returns a tuple with the SmartApiSyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartApiSyncStatus

`func (o *LicenseAccountLicenseData) SetSmartApiSyncStatus(v string)`

SetSmartApiSyncStatus sets SmartApiSyncStatus field to given value.

### HasSmartApiSyncStatus

`func (o *LicenseAccountLicenseData) HasSmartApiSyncStatus() bool`

HasSmartApiSyncStatus returns a boolean if a field has been set.

### GetSyncStatus

`func (o *LicenseAccountLicenseData) GetSyncStatus() string`

GetSyncStatus returns the SyncStatus field if non-nil, zero value otherwise.

### GetSyncStatusOk

`func (o *LicenseAccountLicenseData) GetSyncStatusOk() (*string, bool)`

GetSyncStatusOk returns a tuple with the SyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncStatus

`func (o *LicenseAccountLicenseData) SetSyncStatus(v string)`

SetSyncStatus sets SyncStatus field to given value.

### HasSyncStatus

`func (o *LicenseAccountLicenseData) HasSyncStatus() bool`

HasSyncStatus returns a boolean if a field has been set.

### GetVirtualAccount

`func (o *LicenseAccountLicenseData) GetVirtualAccount() string`

GetVirtualAccount returns the VirtualAccount field if non-nil, zero value otherwise.

### GetVirtualAccountOk

`func (o *LicenseAccountLicenseData) GetVirtualAccountOk() (*string, bool)`

GetVirtualAccountOk returns a tuple with the VirtualAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualAccount

`func (o *LicenseAccountLicenseData) SetVirtualAccount(v string)`

SetVirtualAccount sets VirtualAccount field to given value.

### HasVirtualAccount

`func (o *LicenseAccountLicenseData) HasVirtualAccount() bool`

HasVirtualAccount returns a boolean if a field has been set.

### GetAccount

`func (o *LicenseAccountLicenseData) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *LicenseAccountLicenseData) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *LicenseAccountLicenseData) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *LicenseAccountLicenseData) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *LicenseAccountLicenseData) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *LicenseAccountLicenseData) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetCustomerOp

`func (o *LicenseAccountLicenseData) GetCustomerOp() LicenseCustomerOpRelationship`

GetCustomerOp returns the CustomerOp field if non-nil, zero value otherwise.

### GetCustomerOpOk

`func (o *LicenseAccountLicenseData) GetCustomerOpOk() (*LicenseCustomerOpRelationship, bool)`

GetCustomerOpOk returns a tuple with the CustomerOp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerOp

`func (o *LicenseAccountLicenseData) SetCustomerOp(v LicenseCustomerOpRelationship)`

SetCustomerOp sets CustomerOp field to given value.

### HasCustomerOp

`func (o *LicenseAccountLicenseData) HasCustomerOp() bool`

HasCustomerOp returns a boolean if a field has been set.

### SetCustomerOpNil

`func (o *LicenseAccountLicenseData) SetCustomerOpNil(b bool)`

 SetCustomerOpNil sets the value for CustomerOp to be an explicit nil

### UnsetCustomerOp
`func (o *LicenseAccountLicenseData) UnsetCustomerOp()`

UnsetCustomerOp ensures that no value is present for CustomerOp, not even an explicit nil
### GetErpCustomerOp

`func (o *LicenseAccountLicenseData) GetErpCustomerOp() LicenseErpCustomerOpRelationship`

GetErpCustomerOp returns the ErpCustomerOp field if non-nil, zero value otherwise.

### GetErpCustomerOpOk

`func (o *LicenseAccountLicenseData) GetErpCustomerOpOk() (*LicenseErpCustomerOpRelationship, bool)`

GetErpCustomerOpOk returns a tuple with the ErpCustomerOp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErpCustomerOp

`func (o *LicenseAccountLicenseData) SetErpCustomerOp(v LicenseErpCustomerOpRelationship)`

SetErpCustomerOp sets ErpCustomerOp field to given value.

### HasErpCustomerOp

`func (o *LicenseAccountLicenseData) HasErpCustomerOp() bool`

HasErpCustomerOp returns a boolean if a field has been set.

### SetErpCustomerOpNil

`func (o *LicenseAccountLicenseData) SetErpCustomerOpNil(b bool)`

 SetErpCustomerOpNil sets the value for ErpCustomerOp to be an explicit nil

### UnsetErpCustomerOp
`func (o *LicenseAccountLicenseData) UnsetErpCustomerOp()`

UnsetErpCustomerOp ensures that no value is present for ErpCustomerOp, not even an explicit nil
### GetErpLicenseCount

`func (o *LicenseAccountLicenseData) GetErpLicenseCount() LicenseErpLicenseCountRelationship`

GetErpLicenseCount returns the ErpLicenseCount field if non-nil, zero value otherwise.

### GetErpLicenseCountOk

`func (o *LicenseAccountLicenseData) GetErpLicenseCountOk() (*LicenseErpLicenseCountRelationship, bool)`

GetErpLicenseCountOk returns a tuple with the ErpLicenseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErpLicenseCount

`func (o *LicenseAccountLicenseData) SetErpLicenseCount(v LicenseErpLicenseCountRelationship)`

SetErpLicenseCount sets ErpLicenseCount field to given value.

### HasErpLicenseCount

`func (o *LicenseAccountLicenseData) HasErpLicenseCount() bool`

HasErpLicenseCount returns a boolean if a field has been set.

### SetErpLicenseCountNil

`func (o *LicenseAccountLicenseData) SetErpLicenseCountNil(b bool)`

 SetErpLicenseCountNil sets the value for ErpLicenseCount to be an explicit nil

### UnsetErpLicenseCount
`func (o *LicenseAccountLicenseData) UnsetErpLicenseCount()`

UnsetErpLicenseCount ensures that no value is present for ErpLicenseCount, not even an explicit nil
### GetIksCustomerOp

`func (o *LicenseAccountLicenseData) GetIksCustomerOp() LicenseIksCustomerOpRelationship`

GetIksCustomerOp returns the IksCustomerOp field if non-nil, zero value otherwise.

### GetIksCustomerOpOk

`func (o *LicenseAccountLicenseData) GetIksCustomerOpOk() (*LicenseIksCustomerOpRelationship, bool)`

GetIksCustomerOpOk returns a tuple with the IksCustomerOp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIksCustomerOp

`func (o *LicenseAccountLicenseData) SetIksCustomerOp(v LicenseIksCustomerOpRelationship)`

SetIksCustomerOp sets IksCustomerOp field to given value.

### HasIksCustomerOp

`func (o *LicenseAccountLicenseData) HasIksCustomerOp() bool`

HasIksCustomerOp returns a boolean if a field has been set.

### SetIksCustomerOpNil

`func (o *LicenseAccountLicenseData) SetIksCustomerOpNil(b bool)`

 SetIksCustomerOpNil sets the value for IksCustomerOp to be an explicit nil

### UnsetIksCustomerOp
`func (o *LicenseAccountLicenseData) UnsetIksCustomerOp()`

UnsetIksCustomerOp ensures that no value is present for IksCustomerOp, not even an explicit nil
### GetIksLicenseCount

`func (o *LicenseAccountLicenseData) GetIksLicenseCount() LicenseIksLicenseCountRelationship`

GetIksLicenseCount returns the IksLicenseCount field if non-nil, zero value otherwise.

### GetIksLicenseCountOk

`func (o *LicenseAccountLicenseData) GetIksLicenseCountOk() (*LicenseIksLicenseCountRelationship, bool)`

GetIksLicenseCountOk returns a tuple with the IksLicenseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIksLicenseCount

`func (o *LicenseAccountLicenseData) SetIksLicenseCount(v LicenseIksLicenseCountRelationship)`

SetIksLicenseCount sets IksLicenseCount field to given value.

### HasIksLicenseCount

`func (o *LicenseAccountLicenseData) HasIksLicenseCount() bool`

HasIksLicenseCount returns a boolean if a field has been set.

### SetIksLicenseCountNil

`func (o *LicenseAccountLicenseData) SetIksLicenseCountNil(b bool)`

 SetIksLicenseCountNil sets the value for IksLicenseCount to be an explicit nil

### UnsetIksLicenseCount
`func (o *LicenseAccountLicenseData) UnsetIksLicenseCount()`

UnsetIksLicenseCount ensures that no value is present for IksLicenseCount, not even an explicit nil
### GetIncCustomerOp

`func (o *LicenseAccountLicenseData) GetIncCustomerOp() LicenseIncCustomerOpRelationship`

GetIncCustomerOp returns the IncCustomerOp field if non-nil, zero value otherwise.

### GetIncCustomerOpOk

`func (o *LicenseAccountLicenseData) GetIncCustomerOpOk() (*LicenseIncCustomerOpRelationship, bool)`

GetIncCustomerOpOk returns a tuple with the IncCustomerOp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncCustomerOp

`func (o *LicenseAccountLicenseData) SetIncCustomerOp(v LicenseIncCustomerOpRelationship)`

SetIncCustomerOp sets IncCustomerOp field to given value.

### HasIncCustomerOp

`func (o *LicenseAccountLicenseData) HasIncCustomerOp() bool`

HasIncCustomerOp returns a boolean if a field has been set.

### SetIncCustomerOpNil

`func (o *LicenseAccountLicenseData) SetIncCustomerOpNil(b bool)`

 SetIncCustomerOpNil sets the value for IncCustomerOp to be an explicit nil

### UnsetIncCustomerOp
`func (o *LicenseAccountLicenseData) UnsetIncCustomerOp()`

UnsetIncCustomerOp ensures that no value is present for IncCustomerOp, not even an explicit nil
### GetIncLicenseCount

`func (o *LicenseAccountLicenseData) GetIncLicenseCount() LicenseIncLicenseCountRelationship`

GetIncLicenseCount returns the IncLicenseCount field if non-nil, zero value otherwise.

### GetIncLicenseCountOk

`func (o *LicenseAccountLicenseData) GetIncLicenseCountOk() (*LicenseIncLicenseCountRelationship, bool)`

GetIncLicenseCountOk returns a tuple with the IncLicenseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncLicenseCount

`func (o *LicenseAccountLicenseData) SetIncLicenseCount(v LicenseIncLicenseCountRelationship)`

SetIncLicenseCount sets IncLicenseCount field to given value.

### HasIncLicenseCount

`func (o *LicenseAccountLicenseData) HasIncLicenseCount() bool`

HasIncLicenseCount returns a boolean if a field has been set.

### SetIncLicenseCountNil

`func (o *LicenseAccountLicenseData) SetIncLicenseCountNil(b bool)`

 SetIncLicenseCountNil sets the value for IncLicenseCount to be an explicit nil

### UnsetIncLicenseCount
`func (o *LicenseAccountLicenseData) UnsetIncLicenseCount()`

UnsetIncLicenseCount ensures that no value is present for IncLicenseCount, not even an explicit nil
### GetIwoCustomerOp

`func (o *LicenseAccountLicenseData) GetIwoCustomerOp() LicenseIwoCustomerOpRelationship`

GetIwoCustomerOp returns the IwoCustomerOp field if non-nil, zero value otherwise.

### GetIwoCustomerOpOk

`func (o *LicenseAccountLicenseData) GetIwoCustomerOpOk() (*LicenseIwoCustomerOpRelationship, bool)`

GetIwoCustomerOpOk returns a tuple with the IwoCustomerOp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIwoCustomerOp

`func (o *LicenseAccountLicenseData) SetIwoCustomerOp(v LicenseIwoCustomerOpRelationship)`

SetIwoCustomerOp sets IwoCustomerOp field to given value.

### HasIwoCustomerOp

`func (o *LicenseAccountLicenseData) HasIwoCustomerOp() bool`

HasIwoCustomerOp returns a boolean if a field has been set.

### SetIwoCustomerOpNil

`func (o *LicenseAccountLicenseData) SetIwoCustomerOpNil(b bool)`

 SetIwoCustomerOpNil sets the value for IwoCustomerOp to be an explicit nil

### UnsetIwoCustomerOp
`func (o *LicenseAccountLicenseData) UnsetIwoCustomerOp()`

UnsetIwoCustomerOp ensures that no value is present for IwoCustomerOp, not even an explicit nil
### GetIwoLicenseCount

`func (o *LicenseAccountLicenseData) GetIwoLicenseCount() LicenseIwoLicenseCountRelationship`

GetIwoLicenseCount returns the IwoLicenseCount field if non-nil, zero value otherwise.

### GetIwoLicenseCountOk

`func (o *LicenseAccountLicenseData) GetIwoLicenseCountOk() (*LicenseIwoLicenseCountRelationship, bool)`

GetIwoLicenseCountOk returns a tuple with the IwoLicenseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIwoLicenseCount

`func (o *LicenseAccountLicenseData) SetIwoLicenseCount(v LicenseIwoLicenseCountRelationship)`

SetIwoLicenseCount sets IwoLicenseCount field to given value.

### HasIwoLicenseCount

`func (o *LicenseAccountLicenseData) HasIwoLicenseCount() bool`

HasIwoLicenseCount returns a boolean if a field has been set.

### SetIwoLicenseCountNil

`func (o *LicenseAccountLicenseData) SetIwoLicenseCountNil(b bool)`

 SetIwoLicenseCountNil sets the value for IwoLicenseCount to be an explicit nil

### UnsetIwoLicenseCount
`func (o *LicenseAccountLicenseData) UnsetIwoLicenseCount()`

UnsetIwoLicenseCount ensures that no value is present for IwoLicenseCount, not even an explicit nil
### GetLicenseInfoView

`func (o *LicenseAccountLicenseData) GetLicenseInfoView() LicenseLicenseInfoViewRelationship`

GetLicenseInfoView returns the LicenseInfoView field if non-nil, zero value otherwise.

### GetLicenseInfoViewOk

`func (o *LicenseAccountLicenseData) GetLicenseInfoViewOk() (*LicenseLicenseInfoViewRelationship, bool)`

GetLicenseInfoViewOk returns a tuple with the LicenseInfoView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseInfoView

`func (o *LicenseAccountLicenseData) SetLicenseInfoView(v LicenseLicenseInfoViewRelationship)`

SetLicenseInfoView sets LicenseInfoView field to given value.

### HasLicenseInfoView

`func (o *LicenseAccountLicenseData) HasLicenseInfoView() bool`

HasLicenseInfoView returns a boolean if a field has been set.

### SetLicenseInfoViewNil

`func (o *LicenseAccountLicenseData) SetLicenseInfoViewNil(b bool)`

 SetLicenseInfoViewNil sets the value for LicenseInfoView to be an explicit nil

### UnsetLicenseInfoView
`func (o *LicenseAccountLicenseData) UnsetLicenseInfoView()`

UnsetLicenseInfoView ensures that no value is present for LicenseInfoView, not even an explicit nil
### GetLicenseRegistrationStatus

`func (o *LicenseAccountLicenseData) GetLicenseRegistrationStatus() LicenseLicenseRegistrationStatusRelationship`

GetLicenseRegistrationStatus returns the LicenseRegistrationStatus field if non-nil, zero value otherwise.

### GetLicenseRegistrationStatusOk

`func (o *LicenseAccountLicenseData) GetLicenseRegistrationStatusOk() (*LicenseLicenseRegistrationStatusRelationship, bool)`

GetLicenseRegistrationStatusOk returns a tuple with the LicenseRegistrationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseRegistrationStatus

`func (o *LicenseAccountLicenseData) SetLicenseRegistrationStatus(v LicenseLicenseRegistrationStatusRelationship)`

SetLicenseRegistrationStatus sets LicenseRegistrationStatus field to given value.

### HasLicenseRegistrationStatus

`func (o *LicenseAccountLicenseData) HasLicenseRegistrationStatus() bool`

HasLicenseRegistrationStatus returns a boolean if a field has been set.

### SetLicenseRegistrationStatusNil

`func (o *LicenseAccountLicenseData) SetLicenseRegistrationStatusNil(b bool)`

 SetLicenseRegistrationStatusNil sets the value for LicenseRegistrationStatus to be an explicit nil

### UnsetLicenseRegistrationStatus
`func (o *LicenseAccountLicenseData) UnsetLicenseRegistrationStatus()`

UnsetLicenseRegistrationStatus ensures that no value is present for LicenseRegistrationStatus, not even an explicit nil
### GetLicenseinfos

`func (o *LicenseAccountLicenseData) GetLicenseinfos() []LicenseLicenseInfoRelationship`

GetLicenseinfos returns the Licenseinfos field if non-nil, zero value otherwise.

### GetLicenseinfosOk

`func (o *LicenseAccountLicenseData) GetLicenseinfosOk() (*[]LicenseLicenseInfoRelationship, bool)`

GetLicenseinfosOk returns a tuple with the Licenseinfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseinfos

`func (o *LicenseAccountLicenseData) SetLicenseinfos(v []LicenseLicenseInfoRelationship)`

SetLicenseinfos sets Licenseinfos field to given value.

### HasLicenseinfos

`func (o *LicenseAccountLicenseData) HasLicenseinfos() bool`

HasLicenseinfos returns a boolean if a field has been set.

### SetLicenseinfosNil

`func (o *LicenseAccountLicenseData) SetLicenseinfosNil(b bool)`

 SetLicenseinfosNil sets the value for Licenseinfos to be an explicit nil

### UnsetLicenseinfos
`func (o *LicenseAccountLicenseData) UnsetLicenseinfos()`

UnsetLicenseinfos ensures that no value is present for Licenseinfos, not even an explicit nil
### GetSmartlicenseToken

`func (o *LicenseAccountLicenseData) GetSmartlicenseToken() LicenseSmartlicenseTokenRelationship`

GetSmartlicenseToken returns the SmartlicenseToken field if non-nil, zero value otherwise.

### GetSmartlicenseTokenOk

`func (o *LicenseAccountLicenseData) GetSmartlicenseTokenOk() (*LicenseSmartlicenseTokenRelationship, bool)`

GetSmartlicenseTokenOk returns a tuple with the SmartlicenseToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartlicenseToken

`func (o *LicenseAccountLicenseData) SetSmartlicenseToken(v LicenseSmartlicenseTokenRelationship)`

SetSmartlicenseToken sets SmartlicenseToken field to given value.

### HasSmartlicenseToken

`func (o *LicenseAccountLicenseData) HasSmartlicenseToken() bool`

HasSmartlicenseToken returns a boolean if a field has been set.

### SetSmartlicenseTokenNil

`func (o *LicenseAccountLicenseData) SetSmartlicenseTokenNil(b bool)`

 SetSmartlicenseTokenNil sets the value for SmartlicenseToken to be an explicit nil

### UnsetSmartlicenseToken
`func (o *LicenseAccountLicenseData) UnsetSmartlicenseToken()`

UnsetSmartlicenseToken ensures that no value is present for SmartlicenseToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


