# NiaapiSoftwareEolAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**AffectedVersions** | Pointer to **string** | String contains the Release versions affected by this notice, seperated by comma. | [optional] 
**AnnouncementDate** | Pointer to **time.Time** | Date time of this notice Announced. | [optional] 
**AnnouncementDateEpoch** | Pointer to **int64** | Epoch time of this notice Announced. | [optional] 
**BulletinNo** | Pointer to **string** | The bulletinno of this software release end of life notice. | [optional] 
**Description** | Pointer to **string** | The description of this software release end of life notice. | [optional] 
**EndofNewServiceAttachmentDate** | Pointer to **time.Time** | Date time of End of New service attachment . | [optional] 
**EndofNewServiceAttachmentDateEpoch** | Pointer to **int64** | Epoch time of End of New service attachment . | [optional] 
**EndofServiceContractRenewalDate** | Pointer to **time.Time** | Date time of End of Renewal of service contract . | [optional] 
**EndofServiceContractRenewalDateEpoch** | Pointer to **int64** | Epoch time of End of Renewal of service contract . | [optional] 
**EndofSwMaintenanceDate** | Pointer to **time.Time** | Date time of End of Maintenance. | [optional] 
**EndofSwMaintenanceDateEpoch** | Pointer to **int64** | Epoch time of End of Maintenance. | [optional] 
**Headline** | Pointer to **string** | The title of this software release end of life notice. | [optional] 
**LastDateofSupport** | Pointer to **time.Time** | Date time of Last day of Support . | [optional] 
**LastDateofSupportEpoch** | Pointer to **int64** | Epoch time of Last day of Support . | [optional] 
**LastShipDate** | Pointer to **time.Time** | Date time of Lastship Date. | [optional] 
**LastShipDateEpoch** | Pointer to **int64** | Epoch time of Lastship Date. | [optional] 
**MigrationUrl** | Pointer to **string** | The URL of this migration notice. | [optional] 
**SoftwareEolUrl** | Pointer to **string** | Software end of life notice URL link to the notice webpage. | [optional] 

## Methods

### NewNiaapiSoftwareEolAllOf

`func NewNiaapiSoftwareEolAllOf(classId string, objectType string, ) *NiaapiSoftwareEolAllOf`

NewNiaapiSoftwareEolAllOf instantiates a new NiaapiSoftwareEolAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiaapiSoftwareEolAllOfWithDefaults

`func NewNiaapiSoftwareEolAllOfWithDefaults() *NiaapiSoftwareEolAllOf`

NewNiaapiSoftwareEolAllOfWithDefaults instantiates a new NiaapiSoftwareEolAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiaapiSoftwareEolAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiaapiSoftwareEolAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiaapiSoftwareEolAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiaapiSoftwareEolAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiaapiSoftwareEolAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiaapiSoftwareEolAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAffectedVersions

`func (o *NiaapiSoftwareEolAllOf) GetAffectedVersions() string`

GetAffectedVersions returns the AffectedVersions field if non-nil, zero value otherwise.

### GetAffectedVersionsOk

`func (o *NiaapiSoftwareEolAllOf) GetAffectedVersionsOk() (*string, bool)`

GetAffectedVersionsOk returns a tuple with the AffectedVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedVersions

`func (o *NiaapiSoftwareEolAllOf) SetAffectedVersions(v string)`

SetAffectedVersions sets AffectedVersions field to given value.

### HasAffectedVersions

`func (o *NiaapiSoftwareEolAllOf) HasAffectedVersions() bool`

HasAffectedVersions returns a boolean if a field has been set.

### GetAnnouncementDate

`func (o *NiaapiSoftwareEolAllOf) GetAnnouncementDate() time.Time`

GetAnnouncementDate returns the AnnouncementDate field if non-nil, zero value otherwise.

### GetAnnouncementDateOk

`func (o *NiaapiSoftwareEolAllOf) GetAnnouncementDateOk() (*time.Time, bool)`

GetAnnouncementDateOk returns a tuple with the AnnouncementDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementDate

`func (o *NiaapiSoftwareEolAllOf) SetAnnouncementDate(v time.Time)`

SetAnnouncementDate sets AnnouncementDate field to given value.

### HasAnnouncementDate

`func (o *NiaapiSoftwareEolAllOf) HasAnnouncementDate() bool`

HasAnnouncementDate returns a boolean if a field has been set.

### GetAnnouncementDateEpoch

`func (o *NiaapiSoftwareEolAllOf) GetAnnouncementDateEpoch() int64`

GetAnnouncementDateEpoch returns the AnnouncementDateEpoch field if non-nil, zero value otherwise.

### GetAnnouncementDateEpochOk

`func (o *NiaapiSoftwareEolAllOf) GetAnnouncementDateEpochOk() (*int64, bool)`

GetAnnouncementDateEpochOk returns a tuple with the AnnouncementDateEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementDateEpoch

`func (o *NiaapiSoftwareEolAllOf) SetAnnouncementDateEpoch(v int64)`

SetAnnouncementDateEpoch sets AnnouncementDateEpoch field to given value.

### HasAnnouncementDateEpoch

`func (o *NiaapiSoftwareEolAllOf) HasAnnouncementDateEpoch() bool`

HasAnnouncementDateEpoch returns a boolean if a field has been set.

### GetBulletinNo

`func (o *NiaapiSoftwareEolAllOf) GetBulletinNo() string`

GetBulletinNo returns the BulletinNo field if non-nil, zero value otherwise.

### GetBulletinNoOk

`func (o *NiaapiSoftwareEolAllOf) GetBulletinNoOk() (*string, bool)`

GetBulletinNoOk returns a tuple with the BulletinNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulletinNo

`func (o *NiaapiSoftwareEolAllOf) SetBulletinNo(v string)`

SetBulletinNo sets BulletinNo field to given value.

### HasBulletinNo

`func (o *NiaapiSoftwareEolAllOf) HasBulletinNo() bool`

HasBulletinNo returns a boolean if a field has been set.

### GetDescription

`func (o *NiaapiSoftwareEolAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NiaapiSoftwareEolAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NiaapiSoftwareEolAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NiaapiSoftwareEolAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndofNewServiceAttachmentDate

`func (o *NiaapiSoftwareEolAllOf) GetEndofNewServiceAttachmentDate() time.Time`

GetEndofNewServiceAttachmentDate returns the EndofNewServiceAttachmentDate field if non-nil, zero value otherwise.

### GetEndofNewServiceAttachmentDateOk

`func (o *NiaapiSoftwareEolAllOf) GetEndofNewServiceAttachmentDateOk() (*time.Time, bool)`

GetEndofNewServiceAttachmentDateOk returns a tuple with the EndofNewServiceAttachmentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndofNewServiceAttachmentDate

`func (o *NiaapiSoftwareEolAllOf) SetEndofNewServiceAttachmentDate(v time.Time)`

SetEndofNewServiceAttachmentDate sets EndofNewServiceAttachmentDate field to given value.

### HasEndofNewServiceAttachmentDate

`func (o *NiaapiSoftwareEolAllOf) HasEndofNewServiceAttachmentDate() bool`

HasEndofNewServiceAttachmentDate returns a boolean if a field has been set.

### GetEndofNewServiceAttachmentDateEpoch

`func (o *NiaapiSoftwareEolAllOf) GetEndofNewServiceAttachmentDateEpoch() int64`

GetEndofNewServiceAttachmentDateEpoch returns the EndofNewServiceAttachmentDateEpoch field if non-nil, zero value otherwise.

### GetEndofNewServiceAttachmentDateEpochOk

`func (o *NiaapiSoftwareEolAllOf) GetEndofNewServiceAttachmentDateEpochOk() (*int64, bool)`

GetEndofNewServiceAttachmentDateEpochOk returns a tuple with the EndofNewServiceAttachmentDateEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndofNewServiceAttachmentDateEpoch

`func (o *NiaapiSoftwareEolAllOf) SetEndofNewServiceAttachmentDateEpoch(v int64)`

SetEndofNewServiceAttachmentDateEpoch sets EndofNewServiceAttachmentDateEpoch field to given value.

### HasEndofNewServiceAttachmentDateEpoch

`func (o *NiaapiSoftwareEolAllOf) HasEndofNewServiceAttachmentDateEpoch() bool`

HasEndofNewServiceAttachmentDateEpoch returns a boolean if a field has been set.

### GetEndofServiceContractRenewalDate

`func (o *NiaapiSoftwareEolAllOf) GetEndofServiceContractRenewalDate() time.Time`

GetEndofServiceContractRenewalDate returns the EndofServiceContractRenewalDate field if non-nil, zero value otherwise.

### GetEndofServiceContractRenewalDateOk

`func (o *NiaapiSoftwareEolAllOf) GetEndofServiceContractRenewalDateOk() (*time.Time, bool)`

GetEndofServiceContractRenewalDateOk returns a tuple with the EndofServiceContractRenewalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndofServiceContractRenewalDate

`func (o *NiaapiSoftwareEolAllOf) SetEndofServiceContractRenewalDate(v time.Time)`

SetEndofServiceContractRenewalDate sets EndofServiceContractRenewalDate field to given value.

### HasEndofServiceContractRenewalDate

`func (o *NiaapiSoftwareEolAllOf) HasEndofServiceContractRenewalDate() bool`

HasEndofServiceContractRenewalDate returns a boolean if a field has been set.

### GetEndofServiceContractRenewalDateEpoch

`func (o *NiaapiSoftwareEolAllOf) GetEndofServiceContractRenewalDateEpoch() int64`

GetEndofServiceContractRenewalDateEpoch returns the EndofServiceContractRenewalDateEpoch field if non-nil, zero value otherwise.

### GetEndofServiceContractRenewalDateEpochOk

`func (o *NiaapiSoftwareEolAllOf) GetEndofServiceContractRenewalDateEpochOk() (*int64, bool)`

GetEndofServiceContractRenewalDateEpochOk returns a tuple with the EndofServiceContractRenewalDateEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndofServiceContractRenewalDateEpoch

`func (o *NiaapiSoftwareEolAllOf) SetEndofServiceContractRenewalDateEpoch(v int64)`

SetEndofServiceContractRenewalDateEpoch sets EndofServiceContractRenewalDateEpoch field to given value.

### HasEndofServiceContractRenewalDateEpoch

`func (o *NiaapiSoftwareEolAllOf) HasEndofServiceContractRenewalDateEpoch() bool`

HasEndofServiceContractRenewalDateEpoch returns a boolean if a field has been set.

### GetEndofSwMaintenanceDate

`func (o *NiaapiSoftwareEolAllOf) GetEndofSwMaintenanceDate() time.Time`

GetEndofSwMaintenanceDate returns the EndofSwMaintenanceDate field if non-nil, zero value otherwise.

### GetEndofSwMaintenanceDateOk

`func (o *NiaapiSoftwareEolAllOf) GetEndofSwMaintenanceDateOk() (*time.Time, bool)`

GetEndofSwMaintenanceDateOk returns a tuple with the EndofSwMaintenanceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndofSwMaintenanceDate

`func (o *NiaapiSoftwareEolAllOf) SetEndofSwMaintenanceDate(v time.Time)`

SetEndofSwMaintenanceDate sets EndofSwMaintenanceDate field to given value.

### HasEndofSwMaintenanceDate

`func (o *NiaapiSoftwareEolAllOf) HasEndofSwMaintenanceDate() bool`

HasEndofSwMaintenanceDate returns a boolean if a field has been set.

### GetEndofSwMaintenanceDateEpoch

`func (o *NiaapiSoftwareEolAllOf) GetEndofSwMaintenanceDateEpoch() int64`

GetEndofSwMaintenanceDateEpoch returns the EndofSwMaintenanceDateEpoch field if non-nil, zero value otherwise.

### GetEndofSwMaintenanceDateEpochOk

`func (o *NiaapiSoftwareEolAllOf) GetEndofSwMaintenanceDateEpochOk() (*int64, bool)`

GetEndofSwMaintenanceDateEpochOk returns a tuple with the EndofSwMaintenanceDateEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndofSwMaintenanceDateEpoch

`func (o *NiaapiSoftwareEolAllOf) SetEndofSwMaintenanceDateEpoch(v int64)`

SetEndofSwMaintenanceDateEpoch sets EndofSwMaintenanceDateEpoch field to given value.

### HasEndofSwMaintenanceDateEpoch

`func (o *NiaapiSoftwareEolAllOf) HasEndofSwMaintenanceDateEpoch() bool`

HasEndofSwMaintenanceDateEpoch returns a boolean if a field has been set.

### GetHeadline

`func (o *NiaapiSoftwareEolAllOf) GetHeadline() string`

GetHeadline returns the Headline field if non-nil, zero value otherwise.

### GetHeadlineOk

`func (o *NiaapiSoftwareEolAllOf) GetHeadlineOk() (*string, bool)`

GetHeadlineOk returns a tuple with the Headline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadline

`func (o *NiaapiSoftwareEolAllOf) SetHeadline(v string)`

SetHeadline sets Headline field to given value.

### HasHeadline

`func (o *NiaapiSoftwareEolAllOf) HasHeadline() bool`

HasHeadline returns a boolean if a field has been set.

### GetLastDateofSupport

`func (o *NiaapiSoftwareEolAllOf) GetLastDateofSupport() time.Time`

GetLastDateofSupport returns the LastDateofSupport field if non-nil, zero value otherwise.

### GetLastDateofSupportOk

`func (o *NiaapiSoftwareEolAllOf) GetLastDateofSupportOk() (*time.Time, bool)`

GetLastDateofSupportOk returns a tuple with the LastDateofSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDateofSupport

`func (o *NiaapiSoftwareEolAllOf) SetLastDateofSupport(v time.Time)`

SetLastDateofSupport sets LastDateofSupport field to given value.

### HasLastDateofSupport

`func (o *NiaapiSoftwareEolAllOf) HasLastDateofSupport() bool`

HasLastDateofSupport returns a boolean if a field has been set.

### GetLastDateofSupportEpoch

`func (o *NiaapiSoftwareEolAllOf) GetLastDateofSupportEpoch() int64`

GetLastDateofSupportEpoch returns the LastDateofSupportEpoch field if non-nil, zero value otherwise.

### GetLastDateofSupportEpochOk

`func (o *NiaapiSoftwareEolAllOf) GetLastDateofSupportEpochOk() (*int64, bool)`

GetLastDateofSupportEpochOk returns a tuple with the LastDateofSupportEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDateofSupportEpoch

`func (o *NiaapiSoftwareEolAllOf) SetLastDateofSupportEpoch(v int64)`

SetLastDateofSupportEpoch sets LastDateofSupportEpoch field to given value.

### HasLastDateofSupportEpoch

`func (o *NiaapiSoftwareEolAllOf) HasLastDateofSupportEpoch() bool`

HasLastDateofSupportEpoch returns a boolean if a field has been set.

### GetLastShipDate

`func (o *NiaapiSoftwareEolAllOf) GetLastShipDate() time.Time`

GetLastShipDate returns the LastShipDate field if non-nil, zero value otherwise.

### GetLastShipDateOk

`func (o *NiaapiSoftwareEolAllOf) GetLastShipDateOk() (*time.Time, bool)`

GetLastShipDateOk returns a tuple with the LastShipDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastShipDate

`func (o *NiaapiSoftwareEolAllOf) SetLastShipDate(v time.Time)`

SetLastShipDate sets LastShipDate field to given value.

### HasLastShipDate

`func (o *NiaapiSoftwareEolAllOf) HasLastShipDate() bool`

HasLastShipDate returns a boolean if a field has been set.

### GetLastShipDateEpoch

`func (o *NiaapiSoftwareEolAllOf) GetLastShipDateEpoch() int64`

GetLastShipDateEpoch returns the LastShipDateEpoch field if non-nil, zero value otherwise.

### GetLastShipDateEpochOk

`func (o *NiaapiSoftwareEolAllOf) GetLastShipDateEpochOk() (*int64, bool)`

GetLastShipDateEpochOk returns a tuple with the LastShipDateEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastShipDateEpoch

`func (o *NiaapiSoftwareEolAllOf) SetLastShipDateEpoch(v int64)`

SetLastShipDateEpoch sets LastShipDateEpoch field to given value.

### HasLastShipDateEpoch

`func (o *NiaapiSoftwareEolAllOf) HasLastShipDateEpoch() bool`

HasLastShipDateEpoch returns a boolean if a field has been set.

### GetMigrationUrl

`func (o *NiaapiSoftwareEolAllOf) GetMigrationUrl() string`

GetMigrationUrl returns the MigrationUrl field if non-nil, zero value otherwise.

### GetMigrationUrlOk

`func (o *NiaapiSoftwareEolAllOf) GetMigrationUrlOk() (*string, bool)`

GetMigrationUrlOk returns a tuple with the MigrationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationUrl

`func (o *NiaapiSoftwareEolAllOf) SetMigrationUrl(v string)`

SetMigrationUrl sets MigrationUrl field to given value.

### HasMigrationUrl

`func (o *NiaapiSoftwareEolAllOf) HasMigrationUrl() bool`

HasMigrationUrl returns a boolean if a field has been set.

### GetSoftwareEolUrl

`func (o *NiaapiSoftwareEolAllOf) GetSoftwareEolUrl() string`

GetSoftwareEolUrl returns the SoftwareEolUrl field if non-nil, zero value otherwise.

### GetSoftwareEolUrlOk

`func (o *NiaapiSoftwareEolAllOf) GetSoftwareEolUrlOk() (*string, bool)`

GetSoftwareEolUrlOk returns a tuple with the SoftwareEolUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareEolUrl

`func (o *NiaapiSoftwareEolAllOf) SetSoftwareEolUrl(v string)`

SetSoftwareEolUrl sets SoftwareEolUrl field to given value.

### HasSoftwareEolUrl

`func (o *NiaapiSoftwareEolAllOf) HasSoftwareEolUrl() bool`

HasSoftwareEolUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


