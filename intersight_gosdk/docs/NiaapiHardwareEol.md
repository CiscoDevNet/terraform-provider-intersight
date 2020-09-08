# NiaapiHardwareEol

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedPids** | Pointer to **string** | String contains the PID of hardwares affected by this notice, seperated by comma. | [optional] 
**AnnouncementDate** | Pointer to [**time.Time**](time.Time.md) | When this notice is announced. | [optional] 
**AnnouncementDateEpoch** | Pointer to **int64** | Epoch time of Announcement Date. | [optional] 
**BulletinNo** | Pointer to **string** | The bulletinno of this hardware end of life notice. | [optional] 
**Description** | Pointer to **string** | The description of this hardware end of life notice. | [optional] 
**EndofNewServiceAttachmentDate** | Pointer to [**time.Time**](time.Time.md) | Date time of end of new services attachment  . | [optional] 
**EndofNewServiceAttachmentDateEpoch** | Pointer to **int64** | Epoch time of New service attachment Date . | [optional] 
**EndofRoutineFailureAnalysisDate** | Pointer to [**time.Time**](time.Time.md) | Date time of end of routinefailure analysis. | [optional] 
**EndofRoutineFailureAnalysisDateEpoch** | Pointer to **int64** | Epoch time of End of Routine Failure Analysis Date. | [optional] 
**EndofSaleDate** | Pointer to [**time.Time**](time.Time.md) | When this hardware will end sale. | [optional] 
**EndofSaleDateEpoch** | Pointer to **int64** | Epoch time of End of Sale Date. | [optional] 
**EndofSecuritySupport** | Pointer to [**time.Time**](time.Time.md) | Date time of end of security support . | [optional] 
**EndofSecuritySupportEpoch** | Pointer to **int64** | Epoch time of End of Security Support Date . | [optional] 
**EndofServiceContractRenewalDate** | Pointer to [**time.Time**](time.Time.md) | Date time of end of service contract renew . | [optional] 
**EndofServiceContractRenewalDateEpoch** | Pointer to **int64** | Epoch time of End of Renewal service contract. | [optional] 
**EndofSwMaintenanceDate** | Pointer to [**time.Time**](time.Time.md) | The date of end of maintainance. | [optional] 
**EndofSwMaintenanceDateEpoch** | Pointer to **int64** | Epoch time of End of maintenance Date. | [optional] 
**HardwareEolUrl** | Pointer to **string** | Hardware end of sale URL link to the notice webpage. | [optional] 
**Headline** | Pointer to **string** | The title of this hardware end of life notice. | [optional] 
**LastDateofSupport** | Pointer to [**time.Time**](time.Time.md) | Date time of end of support . | [optional] 
**LastDateofSupportEpoch** | Pointer to **int64** | Epoch time of last date of support . | [optional] 
**LastShipDate** | Pointer to [**time.Time**](time.Time.md) | Date time of Lastship Date. | [optional] 
**LastShipDateEpoch** | Pointer to **int64** | Epoch time of last ship Date. | [optional] 
**MigrationUrl** | Pointer to **string** | The URL of this migration notice. | [optional] 

## Methods

### NewNiaapiHardwareEol

`func NewNiaapiHardwareEol() *NiaapiHardwareEol`

NewNiaapiHardwareEol instantiates a new NiaapiHardwareEol object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiaapiHardwareEolWithDefaults

`func NewNiaapiHardwareEolWithDefaults() *NiaapiHardwareEol`

NewNiaapiHardwareEolWithDefaults instantiates a new NiaapiHardwareEol object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedPids

`func (o *NiaapiHardwareEol) GetAffectedPids() string`

GetAffectedPids returns the AffectedPids field if non-nil, zero value otherwise.

### GetAffectedPidsOk

`func (o *NiaapiHardwareEol) GetAffectedPidsOk() (*string, bool)`

GetAffectedPidsOk returns a tuple with the AffectedPids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedPids

`func (o *NiaapiHardwareEol) SetAffectedPids(v string)`

SetAffectedPids sets AffectedPids field to given value.

### HasAffectedPids

`func (o *NiaapiHardwareEol) HasAffectedPids() bool`

HasAffectedPids returns a boolean if a field has been set.

### GetAnnouncementDate

`func (o *NiaapiHardwareEol) GetAnnouncementDate() time.Time`

GetAnnouncementDate returns the AnnouncementDate field if non-nil, zero value otherwise.

### GetAnnouncementDateOk

`func (o *NiaapiHardwareEol) GetAnnouncementDateOk() (*time.Time, bool)`

GetAnnouncementDateOk returns a tuple with the AnnouncementDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementDate

`func (o *NiaapiHardwareEol) SetAnnouncementDate(v time.Time)`

SetAnnouncementDate sets AnnouncementDate field to given value.

### HasAnnouncementDate

`func (o *NiaapiHardwareEol) HasAnnouncementDate() bool`

HasAnnouncementDate returns a boolean if a field has been set.

### GetAnnouncementDateEpoch

`func (o *NiaapiHardwareEol) GetAnnouncementDateEpoch() int64`

GetAnnouncementDateEpoch returns the AnnouncementDateEpoch field if non-nil, zero value otherwise.

### GetAnnouncementDateEpochOk

`func (o *NiaapiHardwareEol) GetAnnouncementDateEpochOk() (*int64, bool)`

GetAnnouncementDateEpochOk returns a tuple with the AnnouncementDateEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementDateEpoch

`func (o *NiaapiHardwareEol) SetAnnouncementDateEpoch(v int64)`

SetAnnouncementDateEpoch sets AnnouncementDateEpoch field to given value.

### HasAnnouncementDateEpoch

`func (o *NiaapiHardwareEol) HasAnnouncementDateEpoch() bool`

HasAnnouncementDateEpoch returns a boolean if a field has been set.

### GetBulletinNo

`func (o *NiaapiHardwareEol) GetBulletinNo() string`

GetBulletinNo returns the BulletinNo field if non-nil, zero value otherwise.

### GetBulletinNoOk

`func (o *NiaapiHardwareEol) GetBulletinNoOk() (*string, bool)`

GetBulletinNoOk returns a tuple with the BulletinNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulletinNo

`func (o *NiaapiHardwareEol) SetBulletinNo(v string)`

SetBulletinNo sets BulletinNo field to given value.

### HasBulletinNo

`func (o *NiaapiHardwareEol) HasBulletinNo() bool`

HasBulletinNo returns a boolean if a field has been set.

### GetDescription

`func (o *NiaapiHardwareEol) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NiaapiHardwareEol) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NiaapiHardwareEol) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NiaapiHardwareEol) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndofNewServiceAttachmentDate

`func (o *NiaapiHardwareEol) GetEndofNewServiceAttachmentDate() time.Time`

GetEndofNewServiceAttachmentDate returns the EndofNewServiceAttachmentDate field if non-nil, zero value otherwise.

### GetEndofNewServiceAttachmentDateOk

`func (o *NiaapiHardwareEol) GetEndofNewServiceAttachmentDateOk() (*time.Time, bool)`

GetEndofNewServiceAttachmentDateOk returns a tuple with the EndofNewServiceAttachmentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndofNewServiceAttachmentDate

`func (o *NiaapiHardwareEol) SetEndofNewServiceAttachmentDate(v time.Time)`

SetEndofNewServiceAttachmentDate sets EndofNewServiceAttachmentDate field to given value.

### HasEndofNewServiceAttachmentDate

`func (o *NiaapiHardwareEol) HasEndofNewServiceAttachmentDate() bool`

HasEndofNewServiceAttachmentDate returns a boolean if a field has been set.

### GetEndofNewServiceAttachmentDateEpoch

`func (o *NiaapiHardwareEol) GetEndofNewServiceAttachmentDateEpoch() int64`

GetEndofNewServiceAttachmentDateEpoch returns the EndofNewServiceAttachmentDateEpoch field if non-nil, zero value otherwise.

### GetEndofNewServiceAttachmentDateEpochOk

`func (o *NiaapiHardwareEol) GetEndofNewServiceAttachmentDateEpochOk() (*int64, bool)`

GetEndofNewServiceAttachmentDateEpochOk returns a tuple with the EndofNewServiceAttachmentDateEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndofNewServiceAttachmentDateEpoch

`func (o *NiaapiHardwareEol) SetEndofNewServiceAttachmentDateEpoch(v int64)`

SetEndofNewServiceAttachmentDateEpoch sets EndofNewServiceAttachmentDateEpoch field to given value.

### HasEndofNewServiceAttachmentDateEpoch

`func (o *NiaapiHardwareEol) HasEndofNewServiceAttachmentDateEpoch() bool`

HasEndofNewServiceAttachmentDateEpoch returns a boolean if a field has been set.

### GetEndofRoutineFailureAnalysisDate

`func (o *NiaapiHardwareEol) GetEndofRoutineFailureAnalysisDate() time.Time`

GetEndofRoutineFailureAnalysisDate returns the EndofRoutineFailureAnalysisDate field if non-nil, zero value otherwise.

### GetEndofRoutineFailureAnalysisDateOk

`func (o *NiaapiHardwareEol) GetEndofRoutineFailureAnalysisDateOk() (*time.Time, bool)`

GetEndofRoutineFailureAnalysisDateOk returns a tuple with the EndofRoutineFailureAnalysisDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndofRoutineFailureAnalysisDate

`func (o *NiaapiHardwareEol) SetEndofRoutineFailureAnalysisDate(v time.Time)`

SetEndofRoutineFailureAnalysisDate sets EndofRoutineFailureAnalysisDate field to given value.

### HasEndofRoutineFailureAnalysisDate

`func (o *NiaapiHardwareEol) HasEndofRoutineFailureAnalysisDate() bool`

HasEndofRoutineFailureAnalysisDate returns a boolean if a field has been set.

### GetEndofRoutineFailureAnalysisDateEpoch

`func (o *NiaapiHardwareEol) GetEndofRoutineFailureAnalysisDateEpoch() int64`

GetEndofRoutineFailureAnalysisDateEpoch returns the EndofRoutineFailureAnalysisDateEpoch field if non-nil, zero value otherwise.

### GetEndofRoutineFailureAnalysisDateEpochOk

`func (o *NiaapiHardwareEol) GetEndofRoutineFailureAnalysisDateEpochOk() (*int64, bool)`

GetEndofRoutineFailureAnalysisDateEpochOk returns a tuple with the EndofRoutineFailureAnalysisDateEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndofRoutineFailureAnalysisDateEpoch

`func (o *NiaapiHardwareEol) SetEndofRoutineFailureAnalysisDateEpoch(v int64)`

SetEndofRoutineFailureAnalysisDateEpoch sets EndofRoutineFailureAnalysisDateEpoch field to given value.

### HasEndofRoutineFailureAnalysisDateEpoch

`func (o *NiaapiHardwareEol) HasEndofRoutineFailureAnalysisDateEpoch() bool`

HasEndofRoutineFailureAnalysisDateEpoch returns a boolean if a field has been set.

### GetEndofSaleDate

`func (o *NiaapiHardwareEol) GetEndofSaleDate() time.Time`

GetEndofSaleDate returns the EndofSaleDate field if non-nil, zero value otherwise.

### GetEndofSaleDateOk

`func (o *NiaapiHardwareEol) GetEndofSaleDateOk() (*time.Time, bool)`

GetEndofSaleDateOk returns a tuple with the EndofSaleDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndofSaleDate

`func (o *NiaapiHardwareEol) SetEndofSaleDate(v time.Time)`

SetEndofSaleDate sets EndofSaleDate field to given value.

### HasEndofSaleDate

`func (o *NiaapiHardwareEol) HasEndofSaleDate() bool`

HasEndofSaleDate returns a boolean if a field has been set.

### GetEndofSaleDateEpoch

`func (o *NiaapiHardwareEol) GetEndofSaleDateEpoch() int64`

GetEndofSaleDateEpoch returns the EndofSaleDateEpoch field if non-nil, zero value otherwise.

### GetEndofSaleDateEpochOk

`func (o *NiaapiHardwareEol) GetEndofSaleDateEpochOk() (*int64, bool)`

GetEndofSaleDateEpochOk returns a tuple with the EndofSaleDateEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndofSaleDateEpoch

`func (o *NiaapiHardwareEol) SetEndofSaleDateEpoch(v int64)`

SetEndofSaleDateEpoch sets EndofSaleDateEpoch field to given value.

### HasEndofSaleDateEpoch

`func (o *NiaapiHardwareEol) HasEndofSaleDateEpoch() bool`

HasEndofSaleDateEpoch returns a boolean if a field has been set.

### GetEndofSecuritySupport

`func (o *NiaapiHardwareEol) GetEndofSecuritySupport() time.Time`

GetEndofSecuritySupport returns the EndofSecuritySupport field if non-nil, zero value otherwise.

### GetEndofSecuritySupportOk

`func (o *NiaapiHardwareEol) GetEndofSecuritySupportOk() (*time.Time, bool)`

GetEndofSecuritySupportOk returns a tuple with the EndofSecuritySupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndofSecuritySupport

`func (o *NiaapiHardwareEol) SetEndofSecuritySupport(v time.Time)`

SetEndofSecuritySupport sets EndofSecuritySupport field to given value.

### HasEndofSecuritySupport

`func (o *NiaapiHardwareEol) HasEndofSecuritySupport() bool`

HasEndofSecuritySupport returns a boolean if a field has been set.

### GetEndofSecuritySupportEpoch

`func (o *NiaapiHardwareEol) GetEndofSecuritySupportEpoch() int64`

GetEndofSecuritySupportEpoch returns the EndofSecuritySupportEpoch field if non-nil, zero value otherwise.

### GetEndofSecuritySupportEpochOk

`func (o *NiaapiHardwareEol) GetEndofSecuritySupportEpochOk() (*int64, bool)`

GetEndofSecuritySupportEpochOk returns a tuple with the EndofSecuritySupportEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndofSecuritySupportEpoch

`func (o *NiaapiHardwareEol) SetEndofSecuritySupportEpoch(v int64)`

SetEndofSecuritySupportEpoch sets EndofSecuritySupportEpoch field to given value.

### HasEndofSecuritySupportEpoch

`func (o *NiaapiHardwareEol) HasEndofSecuritySupportEpoch() bool`

HasEndofSecuritySupportEpoch returns a boolean if a field has been set.

### GetEndofServiceContractRenewalDate

`func (o *NiaapiHardwareEol) GetEndofServiceContractRenewalDate() time.Time`

GetEndofServiceContractRenewalDate returns the EndofServiceContractRenewalDate field if non-nil, zero value otherwise.

### GetEndofServiceContractRenewalDateOk

`func (o *NiaapiHardwareEol) GetEndofServiceContractRenewalDateOk() (*time.Time, bool)`

GetEndofServiceContractRenewalDateOk returns a tuple with the EndofServiceContractRenewalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndofServiceContractRenewalDate

`func (o *NiaapiHardwareEol) SetEndofServiceContractRenewalDate(v time.Time)`

SetEndofServiceContractRenewalDate sets EndofServiceContractRenewalDate field to given value.

### HasEndofServiceContractRenewalDate

`func (o *NiaapiHardwareEol) HasEndofServiceContractRenewalDate() bool`

HasEndofServiceContractRenewalDate returns a boolean if a field has been set.

### GetEndofServiceContractRenewalDateEpoch

`func (o *NiaapiHardwareEol) GetEndofServiceContractRenewalDateEpoch() int64`

GetEndofServiceContractRenewalDateEpoch returns the EndofServiceContractRenewalDateEpoch field if non-nil, zero value otherwise.

### GetEndofServiceContractRenewalDateEpochOk

`func (o *NiaapiHardwareEol) GetEndofServiceContractRenewalDateEpochOk() (*int64, bool)`

GetEndofServiceContractRenewalDateEpochOk returns a tuple with the EndofServiceContractRenewalDateEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndofServiceContractRenewalDateEpoch

`func (o *NiaapiHardwareEol) SetEndofServiceContractRenewalDateEpoch(v int64)`

SetEndofServiceContractRenewalDateEpoch sets EndofServiceContractRenewalDateEpoch field to given value.

### HasEndofServiceContractRenewalDateEpoch

`func (o *NiaapiHardwareEol) HasEndofServiceContractRenewalDateEpoch() bool`

HasEndofServiceContractRenewalDateEpoch returns a boolean if a field has been set.

### GetEndofSwMaintenanceDate

`func (o *NiaapiHardwareEol) GetEndofSwMaintenanceDate() time.Time`

GetEndofSwMaintenanceDate returns the EndofSwMaintenanceDate field if non-nil, zero value otherwise.

### GetEndofSwMaintenanceDateOk

`func (o *NiaapiHardwareEol) GetEndofSwMaintenanceDateOk() (*time.Time, bool)`

GetEndofSwMaintenanceDateOk returns a tuple with the EndofSwMaintenanceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndofSwMaintenanceDate

`func (o *NiaapiHardwareEol) SetEndofSwMaintenanceDate(v time.Time)`

SetEndofSwMaintenanceDate sets EndofSwMaintenanceDate field to given value.

### HasEndofSwMaintenanceDate

`func (o *NiaapiHardwareEol) HasEndofSwMaintenanceDate() bool`

HasEndofSwMaintenanceDate returns a boolean if a field has been set.

### GetEndofSwMaintenanceDateEpoch

`func (o *NiaapiHardwareEol) GetEndofSwMaintenanceDateEpoch() int64`

GetEndofSwMaintenanceDateEpoch returns the EndofSwMaintenanceDateEpoch field if non-nil, zero value otherwise.

### GetEndofSwMaintenanceDateEpochOk

`func (o *NiaapiHardwareEol) GetEndofSwMaintenanceDateEpochOk() (*int64, bool)`

GetEndofSwMaintenanceDateEpochOk returns a tuple with the EndofSwMaintenanceDateEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndofSwMaintenanceDateEpoch

`func (o *NiaapiHardwareEol) SetEndofSwMaintenanceDateEpoch(v int64)`

SetEndofSwMaintenanceDateEpoch sets EndofSwMaintenanceDateEpoch field to given value.

### HasEndofSwMaintenanceDateEpoch

`func (o *NiaapiHardwareEol) HasEndofSwMaintenanceDateEpoch() bool`

HasEndofSwMaintenanceDateEpoch returns a boolean if a field has been set.

### GetHardwareEolUrl

`func (o *NiaapiHardwareEol) GetHardwareEolUrl() string`

GetHardwareEolUrl returns the HardwareEolUrl field if non-nil, zero value otherwise.

### GetHardwareEolUrlOk

`func (o *NiaapiHardwareEol) GetHardwareEolUrlOk() (*string, bool)`

GetHardwareEolUrlOk returns a tuple with the HardwareEolUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareEolUrl

`func (o *NiaapiHardwareEol) SetHardwareEolUrl(v string)`

SetHardwareEolUrl sets HardwareEolUrl field to given value.

### HasHardwareEolUrl

`func (o *NiaapiHardwareEol) HasHardwareEolUrl() bool`

HasHardwareEolUrl returns a boolean if a field has been set.

### GetHeadline

`func (o *NiaapiHardwareEol) GetHeadline() string`

GetHeadline returns the Headline field if non-nil, zero value otherwise.

### GetHeadlineOk

`func (o *NiaapiHardwareEol) GetHeadlineOk() (*string, bool)`

GetHeadlineOk returns a tuple with the Headline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadline

`func (o *NiaapiHardwareEol) SetHeadline(v string)`

SetHeadline sets Headline field to given value.

### HasHeadline

`func (o *NiaapiHardwareEol) HasHeadline() bool`

HasHeadline returns a boolean if a field has been set.

### GetLastDateofSupport

`func (o *NiaapiHardwareEol) GetLastDateofSupport() time.Time`

GetLastDateofSupport returns the LastDateofSupport field if non-nil, zero value otherwise.

### GetLastDateofSupportOk

`func (o *NiaapiHardwareEol) GetLastDateofSupportOk() (*time.Time, bool)`

GetLastDateofSupportOk returns a tuple with the LastDateofSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDateofSupport

`func (o *NiaapiHardwareEol) SetLastDateofSupport(v time.Time)`

SetLastDateofSupport sets LastDateofSupport field to given value.

### HasLastDateofSupport

`func (o *NiaapiHardwareEol) HasLastDateofSupport() bool`

HasLastDateofSupport returns a boolean if a field has been set.

### GetLastDateofSupportEpoch

`func (o *NiaapiHardwareEol) GetLastDateofSupportEpoch() int64`

GetLastDateofSupportEpoch returns the LastDateofSupportEpoch field if non-nil, zero value otherwise.

### GetLastDateofSupportEpochOk

`func (o *NiaapiHardwareEol) GetLastDateofSupportEpochOk() (*int64, bool)`

GetLastDateofSupportEpochOk returns a tuple with the LastDateofSupportEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDateofSupportEpoch

`func (o *NiaapiHardwareEol) SetLastDateofSupportEpoch(v int64)`

SetLastDateofSupportEpoch sets LastDateofSupportEpoch field to given value.

### HasLastDateofSupportEpoch

`func (o *NiaapiHardwareEol) HasLastDateofSupportEpoch() bool`

HasLastDateofSupportEpoch returns a boolean if a field has been set.

### GetLastShipDate

`func (o *NiaapiHardwareEol) GetLastShipDate() time.Time`

GetLastShipDate returns the LastShipDate field if non-nil, zero value otherwise.

### GetLastShipDateOk

`func (o *NiaapiHardwareEol) GetLastShipDateOk() (*time.Time, bool)`

GetLastShipDateOk returns a tuple with the LastShipDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastShipDate

`func (o *NiaapiHardwareEol) SetLastShipDate(v time.Time)`

SetLastShipDate sets LastShipDate field to given value.

### HasLastShipDate

`func (o *NiaapiHardwareEol) HasLastShipDate() bool`

HasLastShipDate returns a boolean if a field has been set.

### GetLastShipDateEpoch

`func (o *NiaapiHardwareEol) GetLastShipDateEpoch() int64`

GetLastShipDateEpoch returns the LastShipDateEpoch field if non-nil, zero value otherwise.

### GetLastShipDateEpochOk

`func (o *NiaapiHardwareEol) GetLastShipDateEpochOk() (*int64, bool)`

GetLastShipDateEpochOk returns a tuple with the LastShipDateEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastShipDateEpoch

`func (o *NiaapiHardwareEol) SetLastShipDateEpoch(v int64)`

SetLastShipDateEpoch sets LastShipDateEpoch field to given value.

### HasLastShipDateEpoch

`func (o *NiaapiHardwareEol) HasLastShipDateEpoch() bool`

HasLastShipDateEpoch returns a boolean if a field has been set.

### GetMigrationUrl

`func (o *NiaapiHardwareEol) GetMigrationUrl() string`

GetMigrationUrl returns the MigrationUrl field if non-nil, zero value otherwise.

### GetMigrationUrlOk

`func (o *NiaapiHardwareEol) GetMigrationUrlOk() (*string, bool)`

GetMigrationUrlOk returns a tuple with the MigrationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationUrl

`func (o *NiaapiHardwareEol) SetMigrationUrl(v string)`

SetMigrationUrl sets MigrationUrl field to given value.

### HasMigrationUrl

`func (o *NiaapiHardwareEol) HasMigrationUrl() bool`

HasMigrationUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


