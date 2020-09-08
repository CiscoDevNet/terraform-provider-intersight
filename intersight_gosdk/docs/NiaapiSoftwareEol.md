# NiaapiSoftwareEol

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedVersions** | Pointer to **string** | String contains the Release versions affected by this notice, seperated by comma. | [optional] 
**AnnouncementDate** | Pointer to [**time.Time**](time.Time.md) | Date time of this notice Announced. | [optional] 
**AnnouncementDateEpoch** | Pointer to **int64** | Epoch time of this notice Announced. | [optional] 
**BulletinNo** | Pointer to **string** | The bulletinno of this software release end of life notice. | [optional] 
**Description** | Pointer to **string** | The description of this software release end of life notice. | [optional] 
**EndofNewServiceAttachmentDate** | Pointer to [**time.Time**](time.Time.md) | Date time of End of New service attachment . | [optional] 
**EndofNewServiceAttachmentDateEpoch** | Pointer to **int64** | Epoch time of End of New service attachment . | [optional] 
**EndofServiceContractRenewalDate** | Pointer to [**time.Time**](time.Time.md) | Date time of End of Renewal of service contract . | [optional] 
**EndofServiceContractRenewalDateEpoch** | Pointer to **int64** | Epoch time of End of Renewal of service contract . | [optional] 
**EndofSwMaintenanceDate** | Pointer to [**time.Time**](time.Time.md) | Date time of End of Maintenance. | [optional] 
**EndofSwMaintenanceDateEpoch** | Pointer to **int64** | Epoch time of End of Maintenance. | [optional] 
**Headline** | Pointer to **string** | The title of this software release end of life notice. | [optional] 
**LastDateofSupport** | Pointer to [**time.Time**](time.Time.md) | Date time of Last day of Support . | [optional] 
**LastDateofSupportEpoch** | Pointer to **int64** | Epoch time of Last day of Support . | [optional] 
**LastShipDate** | Pointer to [**time.Time**](time.Time.md) | Date time of Lastship Date. | [optional] 
**LastShipDateEpoch** | Pointer to **int64** | Epoch time of Lastship Date. | [optional] 
**MigrationUrl** | Pointer to **string** | The URL of this migration notice. | [optional] 
**SoftwareEolUrl** | Pointer to **string** | Software end of life notice URL link to the notice webpage. | [optional] 

## Methods

### NewNiaapiSoftwareEol

`func NewNiaapiSoftwareEol() *NiaapiSoftwareEol`

NewNiaapiSoftwareEol instantiates a new NiaapiSoftwareEol object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiaapiSoftwareEolWithDefaults

`func NewNiaapiSoftwareEolWithDefaults() *NiaapiSoftwareEol`

NewNiaapiSoftwareEolWithDefaults instantiates a new NiaapiSoftwareEol object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedVersions

`func (o *NiaapiSoftwareEol) GetAffectedVersions() string`

GetAffectedVersions returns the AffectedVersions field if non-nil, zero value otherwise.

### GetAffectedVersionsOk

`func (o *NiaapiSoftwareEol) GetAffectedVersionsOk() (*string, bool)`

GetAffectedVersionsOk returns a tuple with the AffectedVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedVersions

`func (o *NiaapiSoftwareEol) SetAffectedVersions(v string)`

SetAffectedVersions sets AffectedVersions field to given value.

### HasAffectedVersions

`func (o *NiaapiSoftwareEol) HasAffectedVersions() bool`

HasAffectedVersions returns a boolean if a field has been set.

### GetAnnouncementDate

`func (o *NiaapiSoftwareEol) GetAnnouncementDate() time.Time`

GetAnnouncementDate returns the AnnouncementDate field if non-nil, zero value otherwise.

### GetAnnouncementDateOk

`func (o *NiaapiSoftwareEol) GetAnnouncementDateOk() (*time.Time, bool)`

GetAnnouncementDateOk returns a tuple with the AnnouncementDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementDate

`func (o *NiaapiSoftwareEol) SetAnnouncementDate(v time.Time)`

SetAnnouncementDate sets AnnouncementDate field to given value.

### HasAnnouncementDate

`func (o *NiaapiSoftwareEol) HasAnnouncementDate() bool`

HasAnnouncementDate returns a boolean if a field has been set.

### GetAnnouncementDateEpoch

`func (o *NiaapiSoftwareEol) GetAnnouncementDateEpoch() int64`

GetAnnouncementDateEpoch returns the AnnouncementDateEpoch field if non-nil, zero value otherwise.

### GetAnnouncementDateEpochOk

`func (o *NiaapiSoftwareEol) GetAnnouncementDateEpochOk() (*int64, bool)`

GetAnnouncementDateEpochOk returns a tuple with the AnnouncementDateEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementDateEpoch

`func (o *NiaapiSoftwareEol) SetAnnouncementDateEpoch(v int64)`

SetAnnouncementDateEpoch sets AnnouncementDateEpoch field to given value.

### HasAnnouncementDateEpoch

`func (o *NiaapiSoftwareEol) HasAnnouncementDateEpoch() bool`

HasAnnouncementDateEpoch returns a boolean if a field has been set.

### GetBulletinNo

`func (o *NiaapiSoftwareEol) GetBulletinNo() string`

GetBulletinNo returns the BulletinNo field if non-nil, zero value otherwise.

### GetBulletinNoOk

`func (o *NiaapiSoftwareEol) GetBulletinNoOk() (*string, bool)`

GetBulletinNoOk returns a tuple with the BulletinNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulletinNo

`func (o *NiaapiSoftwareEol) SetBulletinNo(v string)`

SetBulletinNo sets BulletinNo field to given value.

### HasBulletinNo

`func (o *NiaapiSoftwareEol) HasBulletinNo() bool`

HasBulletinNo returns a boolean if a field has been set.

### GetDescription

`func (o *NiaapiSoftwareEol) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NiaapiSoftwareEol) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NiaapiSoftwareEol) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NiaapiSoftwareEol) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndofNewServiceAttachmentDate

`func (o *NiaapiSoftwareEol) GetEndofNewServiceAttachmentDate() time.Time`

GetEndofNewServiceAttachmentDate returns the EndofNewServiceAttachmentDate field if non-nil, zero value otherwise.

### GetEndofNewServiceAttachmentDateOk

`func (o *NiaapiSoftwareEol) GetEndofNewServiceAttachmentDateOk() (*time.Time, bool)`

GetEndofNewServiceAttachmentDateOk returns a tuple with the EndofNewServiceAttachmentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndofNewServiceAttachmentDate

`func (o *NiaapiSoftwareEol) SetEndofNewServiceAttachmentDate(v time.Time)`

SetEndofNewServiceAttachmentDate sets EndofNewServiceAttachmentDate field to given value.

### HasEndofNewServiceAttachmentDate

`func (o *NiaapiSoftwareEol) HasEndofNewServiceAttachmentDate() bool`

HasEndofNewServiceAttachmentDate returns a boolean if a field has been set.

### GetEndofNewServiceAttachmentDateEpoch

`func (o *NiaapiSoftwareEol) GetEndofNewServiceAttachmentDateEpoch() int64`

GetEndofNewServiceAttachmentDateEpoch returns the EndofNewServiceAttachmentDateEpoch field if non-nil, zero value otherwise.

### GetEndofNewServiceAttachmentDateEpochOk

`func (o *NiaapiSoftwareEol) GetEndofNewServiceAttachmentDateEpochOk() (*int64, bool)`

GetEndofNewServiceAttachmentDateEpochOk returns a tuple with the EndofNewServiceAttachmentDateEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndofNewServiceAttachmentDateEpoch

`func (o *NiaapiSoftwareEol) SetEndofNewServiceAttachmentDateEpoch(v int64)`

SetEndofNewServiceAttachmentDateEpoch sets EndofNewServiceAttachmentDateEpoch field to given value.

### HasEndofNewServiceAttachmentDateEpoch

`func (o *NiaapiSoftwareEol) HasEndofNewServiceAttachmentDateEpoch() bool`

HasEndofNewServiceAttachmentDateEpoch returns a boolean if a field has been set.

### GetEndofServiceContractRenewalDate

`func (o *NiaapiSoftwareEol) GetEndofServiceContractRenewalDate() time.Time`

GetEndofServiceContractRenewalDate returns the EndofServiceContractRenewalDate field if non-nil, zero value otherwise.

### GetEndofServiceContractRenewalDateOk

`func (o *NiaapiSoftwareEol) GetEndofServiceContractRenewalDateOk() (*time.Time, bool)`

GetEndofServiceContractRenewalDateOk returns a tuple with the EndofServiceContractRenewalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndofServiceContractRenewalDate

`func (o *NiaapiSoftwareEol) SetEndofServiceContractRenewalDate(v time.Time)`

SetEndofServiceContractRenewalDate sets EndofServiceContractRenewalDate field to given value.

### HasEndofServiceContractRenewalDate

`func (o *NiaapiSoftwareEol) HasEndofServiceContractRenewalDate() bool`

HasEndofServiceContractRenewalDate returns a boolean if a field has been set.

### GetEndofServiceContractRenewalDateEpoch

`func (o *NiaapiSoftwareEol) GetEndofServiceContractRenewalDateEpoch() int64`

GetEndofServiceContractRenewalDateEpoch returns the EndofServiceContractRenewalDateEpoch field if non-nil, zero value otherwise.

### GetEndofServiceContractRenewalDateEpochOk

`func (o *NiaapiSoftwareEol) GetEndofServiceContractRenewalDateEpochOk() (*int64, bool)`

GetEndofServiceContractRenewalDateEpochOk returns a tuple with the EndofServiceContractRenewalDateEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndofServiceContractRenewalDateEpoch

`func (o *NiaapiSoftwareEol) SetEndofServiceContractRenewalDateEpoch(v int64)`

SetEndofServiceContractRenewalDateEpoch sets EndofServiceContractRenewalDateEpoch field to given value.

### HasEndofServiceContractRenewalDateEpoch

`func (o *NiaapiSoftwareEol) HasEndofServiceContractRenewalDateEpoch() bool`

HasEndofServiceContractRenewalDateEpoch returns a boolean if a field has been set.

### GetEndofSwMaintenanceDate

`func (o *NiaapiSoftwareEol) GetEndofSwMaintenanceDate() time.Time`

GetEndofSwMaintenanceDate returns the EndofSwMaintenanceDate field if non-nil, zero value otherwise.

### GetEndofSwMaintenanceDateOk

`func (o *NiaapiSoftwareEol) GetEndofSwMaintenanceDateOk() (*time.Time, bool)`

GetEndofSwMaintenanceDateOk returns a tuple with the EndofSwMaintenanceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndofSwMaintenanceDate

`func (o *NiaapiSoftwareEol) SetEndofSwMaintenanceDate(v time.Time)`

SetEndofSwMaintenanceDate sets EndofSwMaintenanceDate field to given value.

### HasEndofSwMaintenanceDate

`func (o *NiaapiSoftwareEol) HasEndofSwMaintenanceDate() bool`

HasEndofSwMaintenanceDate returns a boolean if a field has been set.

### GetEndofSwMaintenanceDateEpoch

`func (o *NiaapiSoftwareEol) GetEndofSwMaintenanceDateEpoch() int64`

GetEndofSwMaintenanceDateEpoch returns the EndofSwMaintenanceDateEpoch field if non-nil, zero value otherwise.

### GetEndofSwMaintenanceDateEpochOk

`func (o *NiaapiSoftwareEol) GetEndofSwMaintenanceDateEpochOk() (*int64, bool)`

GetEndofSwMaintenanceDateEpochOk returns a tuple with the EndofSwMaintenanceDateEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndofSwMaintenanceDateEpoch

`func (o *NiaapiSoftwareEol) SetEndofSwMaintenanceDateEpoch(v int64)`

SetEndofSwMaintenanceDateEpoch sets EndofSwMaintenanceDateEpoch field to given value.

### HasEndofSwMaintenanceDateEpoch

`func (o *NiaapiSoftwareEol) HasEndofSwMaintenanceDateEpoch() bool`

HasEndofSwMaintenanceDateEpoch returns a boolean if a field has been set.

### GetHeadline

`func (o *NiaapiSoftwareEol) GetHeadline() string`

GetHeadline returns the Headline field if non-nil, zero value otherwise.

### GetHeadlineOk

`func (o *NiaapiSoftwareEol) GetHeadlineOk() (*string, bool)`

GetHeadlineOk returns a tuple with the Headline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadline

`func (o *NiaapiSoftwareEol) SetHeadline(v string)`

SetHeadline sets Headline field to given value.

### HasHeadline

`func (o *NiaapiSoftwareEol) HasHeadline() bool`

HasHeadline returns a boolean if a field has been set.

### GetLastDateofSupport

`func (o *NiaapiSoftwareEol) GetLastDateofSupport() time.Time`

GetLastDateofSupport returns the LastDateofSupport field if non-nil, zero value otherwise.

### GetLastDateofSupportOk

`func (o *NiaapiSoftwareEol) GetLastDateofSupportOk() (*time.Time, bool)`

GetLastDateofSupportOk returns a tuple with the LastDateofSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDateofSupport

`func (o *NiaapiSoftwareEol) SetLastDateofSupport(v time.Time)`

SetLastDateofSupport sets LastDateofSupport field to given value.

### HasLastDateofSupport

`func (o *NiaapiSoftwareEol) HasLastDateofSupport() bool`

HasLastDateofSupport returns a boolean if a field has been set.

### GetLastDateofSupportEpoch

`func (o *NiaapiSoftwareEol) GetLastDateofSupportEpoch() int64`

GetLastDateofSupportEpoch returns the LastDateofSupportEpoch field if non-nil, zero value otherwise.

### GetLastDateofSupportEpochOk

`func (o *NiaapiSoftwareEol) GetLastDateofSupportEpochOk() (*int64, bool)`

GetLastDateofSupportEpochOk returns a tuple with the LastDateofSupportEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDateofSupportEpoch

`func (o *NiaapiSoftwareEol) SetLastDateofSupportEpoch(v int64)`

SetLastDateofSupportEpoch sets LastDateofSupportEpoch field to given value.

### HasLastDateofSupportEpoch

`func (o *NiaapiSoftwareEol) HasLastDateofSupportEpoch() bool`

HasLastDateofSupportEpoch returns a boolean if a field has been set.

### GetLastShipDate

`func (o *NiaapiSoftwareEol) GetLastShipDate() time.Time`

GetLastShipDate returns the LastShipDate field if non-nil, zero value otherwise.

### GetLastShipDateOk

`func (o *NiaapiSoftwareEol) GetLastShipDateOk() (*time.Time, bool)`

GetLastShipDateOk returns a tuple with the LastShipDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastShipDate

`func (o *NiaapiSoftwareEol) SetLastShipDate(v time.Time)`

SetLastShipDate sets LastShipDate field to given value.

### HasLastShipDate

`func (o *NiaapiSoftwareEol) HasLastShipDate() bool`

HasLastShipDate returns a boolean if a field has been set.

### GetLastShipDateEpoch

`func (o *NiaapiSoftwareEol) GetLastShipDateEpoch() int64`

GetLastShipDateEpoch returns the LastShipDateEpoch field if non-nil, zero value otherwise.

### GetLastShipDateEpochOk

`func (o *NiaapiSoftwareEol) GetLastShipDateEpochOk() (*int64, bool)`

GetLastShipDateEpochOk returns a tuple with the LastShipDateEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastShipDateEpoch

`func (o *NiaapiSoftwareEol) SetLastShipDateEpoch(v int64)`

SetLastShipDateEpoch sets LastShipDateEpoch field to given value.

### HasLastShipDateEpoch

`func (o *NiaapiSoftwareEol) HasLastShipDateEpoch() bool`

HasLastShipDateEpoch returns a boolean if a field has been set.

### GetMigrationUrl

`func (o *NiaapiSoftwareEol) GetMigrationUrl() string`

GetMigrationUrl returns the MigrationUrl field if non-nil, zero value otherwise.

### GetMigrationUrlOk

`func (o *NiaapiSoftwareEol) GetMigrationUrlOk() (*string, bool)`

GetMigrationUrlOk returns a tuple with the MigrationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationUrl

`func (o *NiaapiSoftwareEol) SetMigrationUrl(v string)`

SetMigrationUrl sets MigrationUrl field to given value.

### HasMigrationUrl

`func (o *NiaapiSoftwareEol) HasMigrationUrl() bool`

HasMigrationUrl returns a boolean if a field has been set.

### GetSoftwareEolUrl

`func (o *NiaapiSoftwareEol) GetSoftwareEolUrl() string`

GetSoftwareEolUrl returns the SoftwareEolUrl field if non-nil, zero value otherwise.

### GetSoftwareEolUrlOk

`func (o *NiaapiSoftwareEol) GetSoftwareEolUrlOk() (*string, bool)`

GetSoftwareEolUrlOk returns a tuple with the SoftwareEolUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareEolUrl

`func (o *NiaapiSoftwareEol) SetSoftwareEolUrl(v string)`

SetSoftwareEolUrl sets SoftwareEolUrl field to given value.

### HasSoftwareEolUrl

`func (o *NiaapiSoftwareEol) HasSoftwareEolUrl() bool`

HasSoftwareEolUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


