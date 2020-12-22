# NiaapiHardwareEolAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**AffectedPids** | Pointer to **string** | String contains the PID of hardwares affected by this notice, seperated by comma. | [optional] 
**AnnouncementDate** | Pointer to **time.Time** | When this notice is announced. | [optional] 
**AnnouncementDateEpoch** | Pointer to **int64** | Epoch time of Announcement Date. | [optional] 
**BulletinNo** | Pointer to **string** | The bulletinno of this hardware end of life notice. | [optional] 
**Description** | Pointer to **string** | The description of this hardware end of life notice. | [optional] 
**EndofNewServiceAttachmentDate** | Pointer to **time.Time** | Date time of end of new services attachment  . | [optional] 
**EndofNewServiceAttachmentDateEpoch** | Pointer to **int64** | Epoch time of New service attachment Date . | [optional] 
**EndofRoutineFailureAnalysisDate** | Pointer to **time.Time** | Date time of end of routinefailure analysis. | [optional] 
**EndofRoutineFailureAnalysisDateEpoch** | Pointer to **int64** | Epoch time of End of Routine Failure Analysis Date. | [optional] 
**EndofSaleDate** | Pointer to **time.Time** | When this hardware will end sale. | [optional] 
**EndofSaleDateEpoch** | Pointer to **int64** | Epoch time of End of Sale Date. | [optional] 
**EndofSecuritySupport** | Pointer to **time.Time** | Date time of end of security support . | [optional] 
**EndofSecuritySupportEpoch** | Pointer to **int64** | Epoch time of End of Security Support Date . | [optional] 
**EndofServiceContractRenewalDate** | Pointer to **time.Time** | Date time of end of service contract renew . | [optional] 
**EndofServiceContractRenewalDateEpoch** | Pointer to **int64** | Epoch time of End of Renewal service contract. | [optional] 
**EndofSwMaintenanceDate** | Pointer to **time.Time** | The date of end of maintainance. | [optional] 
**EndofSwMaintenanceDateEpoch** | Pointer to **int64** | Epoch time of End of maintenance Date. | [optional] 
**HardwareEolUrl** | Pointer to **string** | Hardware end of sale URL link to the notice webpage. | [optional] 
**Headline** | Pointer to **string** | The title of this hardware end of life notice. | [optional] 
**LastDateofSupport** | Pointer to **time.Time** | Date time of end of support . | [optional] 
**LastDateofSupportEpoch** | Pointer to **int64** | Epoch time of last date of support . | [optional] 
**LastShipDate** | Pointer to **time.Time** | Date time of Lastship Date. | [optional] 
**LastShipDateEpoch** | Pointer to **int64** | Epoch time of last ship Date. | [optional] 
**MigrationUrl** | Pointer to **string** | The URL of this migration notice. | [optional] 

## Methods

### NewNiaapiHardwareEolAllOf

`func NewNiaapiHardwareEolAllOf(classId string, objectType string, ) *NiaapiHardwareEolAllOf`

NewNiaapiHardwareEolAllOf instantiates a new NiaapiHardwareEolAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiaapiHardwareEolAllOfWithDefaults

`func NewNiaapiHardwareEolAllOfWithDefaults() *NiaapiHardwareEolAllOf`

NewNiaapiHardwareEolAllOfWithDefaults instantiates a new NiaapiHardwareEolAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiaapiHardwareEolAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiaapiHardwareEolAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiaapiHardwareEolAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiaapiHardwareEolAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiaapiHardwareEolAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiaapiHardwareEolAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAffectedPids

`func (o *NiaapiHardwareEolAllOf) GetAffectedPids() string`

GetAffectedPids returns the AffectedPids field if non-nil, zero value otherwise.

### GetAffectedPidsOk

`func (o *NiaapiHardwareEolAllOf) GetAffectedPidsOk() (*string, bool)`

GetAffectedPidsOk returns a tuple with the AffectedPids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedPids

`func (o *NiaapiHardwareEolAllOf) SetAffectedPids(v string)`

SetAffectedPids sets AffectedPids field to given value.

### HasAffectedPids

`func (o *NiaapiHardwareEolAllOf) HasAffectedPids() bool`

HasAffectedPids returns a boolean if a field has been set.

### GetAnnouncementDate

`func (o *NiaapiHardwareEolAllOf) GetAnnouncementDate() time.Time`

GetAnnouncementDate returns the AnnouncementDate field if non-nil, zero value otherwise.

### GetAnnouncementDateOk

`func (o *NiaapiHardwareEolAllOf) GetAnnouncementDateOk() (*time.Time, bool)`

GetAnnouncementDateOk returns a tuple with the AnnouncementDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementDate

`func (o *NiaapiHardwareEolAllOf) SetAnnouncementDate(v time.Time)`

SetAnnouncementDate sets AnnouncementDate field to given value.

### HasAnnouncementDate

`func (o *NiaapiHardwareEolAllOf) HasAnnouncementDate() bool`

HasAnnouncementDate returns a boolean if a field has been set.

### GetAnnouncementDateEpoch

`func (o *NiaapiHardwareEolAllOf) GetAnnouncementDateEpoch() int64`

GetAnnouncementDateEpoch returns the AnnouncementDateEpoch field if non-nil, zero value otherwise.

### GetAnnouncementDateEpochOk

`func (o *NiaapiHardwareEolAllOf) GetAnnouncementDateEpochOk() (*int64, bool)`

GetAnnouncementDateEpochOk returns a tuple with the AnnouncementDateEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementDateEpoch

`func (o *NiaapiHardwareEolAllOf) SetAnnouncementDateEpoch(v int64)`

SetAnnouncementDateEpoch sets AnnouncementDateEpoch field to given value.

### HasAnnouncementDateEpoch

`func (o *NiaapiHardwareEolAllOf) HasAnnouncementDateEpoch() bool`

HasAnnouncementDateEpoch returns a boolean if a field has been set.

### GetBulletinNo

`func (o *NiaapiHardwareEolAllOf) GetBulletinNo() string`

GetBulletinNo returns the BulletinNo field if non-nil, zero value otherwise.

### GetBulletinNoOk

`func (o *NiaapiHardwareEolAllOf) GetBulletinNoOk() (*string, bool)`

GetBulletinNoOk returns a tuple with the BulletinNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulletinNo

`func (o *NiaapiHardwareEolAllOf) SetBulletinNo(v string)`

SetBulletinNo sets BulletinNo field to given value.

### HasBulletinNo

`func (o *NiaapiHardwareEolAllOf) HasBulletinNo() bool`

HasBulletinNo returns a boolean if a field has been set.

### GetDescription

`func (o *NiaapiHardwareEolAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NiaapiHardwareEolAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NiaapiHardwareEolAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NiaapiHardwareEolAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndofNewServiceAttachmentDate

`func (o *NiaapiHardwareEolAllOf) GetEndofNewServiceAttachmentDate() time.Time`

GetEndofNewServiceAttachmentDate returns the EndofNewServiceAttachmentDate field if non-nil, zero value otherwise.

### GetEndofNewServiceAttachmentDateOk

`func (o *NiaapiHardwareEolAllOf) GetEndofNewServiceAttachmentDateOk() (*time.Time, bool)`

GetEndofNewServiceAttachmentDateOk returns a tuple with the EndofNewServiceAttachmentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndofNewServiceAttachmentDate

`func (o *NiaapiHardwareEolAllOf) SetEndofNewServiceAttachmentDate(v time.Time)`

SetEndofNewServiceAttachmentDate sets EndofNewServiceAttachmentDate field to given value.

### HasEndofNewServiceAttachmentDate

`func (o *NiaapiHardwareEolAllOf) HasEndofNewServiceAttachmentDate() bool`

HasEndofNewServiceAttachmentDate returns a boolean if a field has been set.

### GetEndofNewServiceAttachmentDateEpoch

`func (o *NiaapiHardwareEolAllOf) GetEndofNewServiceAttachmentDateEpoch() int64`

GetEndofNewServiceAttachmentDateEpoch returns the EndofNewServiceAttachmentDateEpoch field if non-nil, zero value otherwise.

### GetEndofNewServiceAttachmentDateEpochOk

`func (o *NiaapiHardwareEolAllOf) GetEndofNewServiceAttachmentDateEpochOk() (*int64, bool)`

GetEndofNewServiceAttachmentDateEpochOk returns a tuple with the EndofNewServiceAttachmentDateEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndofNewServiceAttachmentDateEpoch

`func (o *NiaapiHardwareEolAllOf) SetEndofNewServiceAttachmentDateEpoch(v int64)`

SetEndofNewServiceAttachmentDateEpoch sets EndofNewServiceAttachmentDateEpoch field to given value.

### HasEndofNewServiceAttachmentDateEpoch

`func (o *NiaapiHardwareEolAllOf) HasEndofNewServiceAttachmentDateEpoch() bool`

HasEndofNewServiceAttachmentDateEpoch returns a boolean if a field has been set.

### GetEndofRoutineFailureAnalysisDate

`func (o *NiaapiHardwareEolAllOf) GetEndofRoutineFailureAnalysisDate() time.Time`

GetEndofRoutineFailureAnalysisDate returns the EndofRoutineFailureAnalysisDate field if non-nil, zero value otherwise.

### GetEndofRoutineFailureAnalysisDateOk

`func (o *NiaapiHardwareEolAllOf) GetEndofRoutineFailureAnalysisDateOk() (*time.Time, bool)`

GetEndofRoutineFailureAnalysisDateOk returns a tuple with the EndofRoutineFailureAnalysisDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndofRoutineFailureAnalysisDate

`func (o *NiaapiHardwareEolAllOf) SetEndofRoutineFailureAnalysisDate(v time.Time)`

SetEndofRoutineFailureAnalysisDate sets EndofRoutineFailureAnalysisDate field to given value.

### HasEndofRoutineFailureAnalysisDate

`func (o *NiaapiHardwareEolAllOf) HasEndofRoutineFailureAnalysisDate() bool`

HasEndofRoutineFailureAnalysisDate returns a boolean if a field has been set.

### GetEndofRoutineFailureAnalysisDateEpoch

`func (o *NiaapiHardwareEolAllOf) GetEndofRoutineFailureAnalysisDateEpoch() int64`

GetEndofRoutineFailureAnalysisDateEpoch returns the EndofRoutineFailureAnalysisDateEpoch field if non-nil, zero value otherwise.

### GetEndofRoutineFailureAnalysisDateEpochOk

`func (o *NiaapiHardwareEolAllOf) GetEndofRoutineFailureAnalysisDateEpochOk() (*int64, bool)`

GetEndofRoutineFailureAnalysisDateEpochOk returns a tuple with the EndofRoutineFailureAnalysisDateEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndofRoutineFailureAnalysisDateEpoch

`func (o *NiaapiHardwareEolAllOf) SetEndofRoutineFailureAnalysisDateEpoch(v int64)`

SetEndofRoutineFailureAnalysisDateEpoch sets EndofRoutineFailureAnalysisDateEpoch field to given value.

### HasEndofRoutineFailureAnalysisDateEpoch

`func (o *NiaapiHardwareEolAllOf) HasEndofRoutineFailureAnalysisDateEpoch() bool`

HasEndofRoutineFailureAnalysisDateEpoch returns a boolean if a field has been set.

### GetEndofSaleDate

`func (o *NiaapiHardwareEolAllOf) GetEndofSaleDate() time.Time`

GetEndofSaleDate returns the EndofSaleDate field if non-nil, zero value otherwise.

### GetEndofSaleDateOk

`func (o *NiaapiHardwareEolAllOf) GetEndofSaleDateOk() (*time.Time, bool)`

GetEndofSaleDateOk returns a tuple with the EndofSaleDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndofSaleDate

`func (o *NiaapiHardwareEolAllOf) SetEndofSaleDate(v time.Time)`

SetEndofSaleDate sets EndofSaleDate field to given value.

### HasEndofSaleDate

`func (o *NiaapiHardwareEolAllOf) HasEndofSaleDate() bool`

HasEndofSaleDate returns a boolean if a field has been set.

### GetEndofSaleDateEpoch

`func (o *NiaapiHardwareEolAllOf) GetEndofSaleDateEpoch() int64`

GetEndofSaleDateEpoch returns the EndofSaleDateEpoch field if non-nil, zero value otherwise.

### GetEndofSaleDateEpochOk

`func (o *NiaapiHardwareEolAllOf) GetEndofSaleDateEpochOk() (*int64, bool)`

GetEndofSaleDateEpochOk returns a tuple with the EndofSaleDateEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndofSaleDateEpoch

`func (o *NiaapiHardwareEolAllOf) SetEndofSaleDateEpoch(v int64)`

SetEndofSaleDateEpoch sets EndofSaleDateEpoch field to given value.

### HasEndofSaleDateEpoch

`func (o *NiaapiHardwareEolAllOf) HasEndofSaleDateEpoch() bool`

HasEndofSaleDateEpoch returns a boolean if a field has been set.

### GetEndofSecuritySupport

`func (o *NiaapiHardwareEolAllOf) GetEndofSecuritySupport() time.Time`

GetEndofSecuritySupport returns the EndofSecuritySupport field if non-nil, zero value otherwise.

### GetEndofSecuritySupportOk

`func (o *NiaapiHardwareEolAllOf) GetEndofSecuritySupportOk() (*time.Time, bool)`

GetEndofSecuritySupportOk returns a tuple with the EndofSecuritySupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndofSecuritySupport

`func (o *NiaapiHardwareEolAllOf) SetEndofSecuritySupport(v time.Time)`

SetEndofSecuritySupport sets EndofSecuritySupport field to given value.

### HasEndofSecuritySupport

`func (o *NiaapiHardwareEolAllOf) HasEndofSecuritySupport() bool`

HasEndofSecuritySupport returns a boolean if a field has been set.

### GetEndofSecuritySupportEpoch

`func (o *NiaapiHardwareEolAllOf) GetEndofSecuritySupportEpoch() int64`

GetEndofSecuritySupportEpoch returns the EndofSecuritySupportEpoch field if non-nil, zero value otherwise.

### GetEndofSecuritySupportEpochOk

`func (o *NiaapiHardwareEolAllOf) GetEndofSecuritySupportEpochOk() (*int64, bool)`

GetEndofSecuritySupportEpochOk returns a tuple with the EndofSecuritySupportEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndofSecuritySupportEpoch

`func (o *NiaapiHardwareEolAllOf) SetEndofSecuritySupportEpoch(v int64)`

SetEndofSecuritySupportEpoch sets EndofSecuritySupportEpoch field to given value.

### HasEndofSecuritySupportEpoch

`func (o *NiaapiHardwareEolAllOf) HasEndofSecuritySupportEpoch() bool`

HasEndofSecuritySupportEpoch returns a boolean if a field has been set.

### GetEndofServiceContractRenewalDate

`func (o *NiaapiHardwareEolAllOf) GetEndofServiceContractRenewalDate() time.Time`

GetEndofServiceContractRenewalDate returns the EndofServiceContractRenewalDate field if non-nil, zero value otherwise.

### GetEndofServiceContractRenewalDateOk

`func (o *NiaapiHardwareEolAllOf) GetEndofServiceContractRenewalDateOk() (*time.Time, bool)`

GetEndofServiceContractRenewalDateOk returns a tuple with the EndofServiceContractRenewalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndofServiceContractRenewalDate

`func (o *NiaapiHardwareEolAllOf) SetEndofServiceContractRenewalDate(v time.Time)`

SetEndofServiceContractRenewalDate sets EndofServiceContractRenewalDate field to given value.

### HasEndofServiceContractRenewalDate

`func (o *NiaapiHardwareEolAllOf) HasEndofServiceContractRenewalDate() bool`

HasEndofServiceContractRenewalDate returns a boolean if a field has been set.

### GetEndofServiceContractRenewalDateEpoch

`func (o *NiaapiHardwareEolAllOf) GetEndofServiceContractRenewalDateEpoch() int64`

GetEndofServiceContractRenewalDateEpoch returns the EndofServiceContractRenewalDateEpoch field if non-nil, zero value otherwise.

### GetEndofServiceContractRenewalDateEpochOk

`func (o *NiaapiHardwareEolAllOf) GetEndofServiceContractRenewalDateEpochOk() (*int64, bool)`

GetEndofServiceContractRenewalDateEpochOk returns a tuple with the EndofServiceContractRenewalDateEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndofServiceContractRenewalDateEpoch

`func (o *NiaapiHardwareEolAllOf) SetEndofServiceContractRenewalDateEpoch(v int64)`

SetEndofServiceContractRenewalDateEpoch sets EndofServiceContractRenewalDateEpoch field to given value.

### HasEndofServiceContractRenewalDateEpoch

`func (o *NiaapiHardwareEolAllOf) HasEndofServiceContractRenewalDateEpoch() bool`

HasEndofServiceContractRenewalDateEpoch returns a boolean if a field has been set.

### GetEndofSwMaintenanceDate

`func (o *NiaapiHardwareEolAllOf) GetEndofSwMaintenanceDate() time.Time`

GetEndofSwMaintenanceDate returns the EndofSwMaintenanceDate field if non-nil, zero value otherwise.

### GetEndofSwMaintenanceDateOk

`func (o *NiaapiHardwareEolAllOf) GetEndofSwMaintenanceDateOk() (*time.Time, bool)`

GetEndofSwMaintenanceDateOk returns a tuple with the EndofSwMaintenanceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndofSwMaintenanceDate

`func (o *NiaapiHardwareEolAllOf) SetEndofSwMaintenanceDate(v time.Time)`

SetEndofSwMaintenanceDate sets EndofSwMaintenanceDate field to given value.

### HasEndofSwMaintenanceDate

`func (o *NiaapiHardwareEolAllOf) HasEndofSwMaintenanceDate() bool`

HasEndofSwMaintenanceDate returns a boolean if a field has been set.

### GetEndofSwMaintenanceDateEpoch

`func (o *NiaapiHardwareEolAllOf) GetEndofSwMaintenanceDateEpoch() int64`

GetEndofSwMaintenanceDateEpoch returns the EndofSwMaintenanceDateEpoch field if non-nil, zero value otherwise.

### GetEndofSwMaintenanceDateEpochOk

`func (o *NiaapiHardwareEolAllOf) GetEndofSwMaintenanceDateEpochOk() (*int64, bool)`

GetEndofSwMaintenanceDateEpochOk returns a tuple with the EndofSwMaintenanceDateEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndofSwMaintenanceDateEpoch

`func (o *NiaapiHardwareEolAllOf) SetEndofSwMaintenanceDateEpoch(v int64)`

SetEndofSwMaintenanceDateEpoch sets EndofSwMaintenanceDateEpoch field to given value.

### HasEndofSwMaintenanceDateEpoch

`func (o *NiaapiHardwareEolAllOf) HasEndofSwMaintenanceDateEpoch() bool`

HasEndofSwMaintenanceDateEpoch returns a boolean if a field has been set.

### GetHardwareEolUrl

`func (o *NiaapiHardwareEolAllOf) GetHardwareEolUrl() string`

GetHardwareEolUrl returns the HardwareEolUrl field if non-nil, zero value otherwise.

### GetHardwareEolUrlOk

`func (o *NiaapiHardwareEolAllOf) GetHardwareEolUrlOk() (*string, bool)`

GetHardwareEolUrlOk returns a tuple with the HardwareEolUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareEolUrl

`func (o *NiaapiHardwareEolAllOf) SetHardwareEolUrl(v string)`

SetHardwareEolUrl sets HardwareEolUrl field to given value.

### HasHardwareEolUrl

`func (o *NiaapiHardwareEolAllOf) HasHardwareEolUrl() bool`

HasHardwareEolUrl returns a boolean if a field has been set.

### GetHeadline

`func (o *NiaapiHardwareEolAllOf) GetHeadline() string`

GetHeadline returns the Headline field if non-nil, zero value otherwise.

### GetHeadlineOk

`func (o *NiaapiHardwareEolAllOf) GetHeadlineOk() (*string, bool)`

GetHeadlineOk returns a tuple with the Headline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadline

`func (o *NiaapiHardwareEolAllOf) SetHeadline(v string)`

SetHeadline sets Headline field to given value.

### HasHeadline

`func (o *NiaapiHardwareEolAllOf) HasHeadline() bool`

HasHeadline returns a boolean if a field has been set.

### GetLastDateofSupport

`func (o *NiaapiHardwareEolAllOf) GetLastDateofSupport() time.Time`

GetLastDateofSupport returns the LastDateofSupport field if non-nil, zero value otherwise.

### GetLastDateofSupportOk

`func (o *NiaapiHardwareEolAllOf) GetLastDateofSupportOk() (*time.Time, bool)`

GetLastDateofSupportOk returns a tuple with the LastDateofSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDateofSupport

`func (o *NiaapiHardwareEolAllOf) SetLastDateofSupport(v time.Time)`

SetLastDateofSupport sets LastDateofSupport field to given value.

### HasLastDateofSupport

`func (o *NiaapiHardwareEolAllOf) HasLastDateofSupport() bool`

HasLastDateofSupport returns a boolean if a field has been set.

### GetLastDateofSupportEpoch

`func (o *NiaapiHardwareEolAllOf) GetLastDateofSupportEpoch() int64`

GetLastDateofSupportEpoch returns the LastDateofSupportEpoch field if non-nil, zero value otherwise.

### GetLastDateofSupportEpochOk

`func (o *NiaapiHardwareEolAllOf) GetLastDateofSupportEpochOk() (*int64, bool)`

GetLastDateofSupportEpochOk returns a tuple with the LastDateofSupportEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDateofSupportEpoch

`func (o *NiaapiHardwareEolAllOf) SetLastDateofSupportEpoch(v int64)`

SetLastDateofSupportEpoch sets LastDateofSupportEpoch field to given value.

### HasLastDateofSupportEpoch

`func (o *NiaapiHardwareEolAllOf) HasLastDateofSupportEpoch() bool`

HasLastDateofSupportEpoch returns a boolean if a field has been set.

### GetLastShipDate

`func (o *NiaapiHardwareEolAllOf) GetLastShipDate() time.Time`

GetLastShipDate returns the LastShipDate field if non-nil, zero value otherwise.

### GetLastShipDateOk

`func (o *NiaapiHardwareEolAllOf) GetLastShipDateOk() (*time.Time, bool)`

GetLastShipDateOk returns a tuple with the LastShipDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastShipDate

`func (o *NiaapiHardwareEolAllOf) SetLastShipDate(v time.Time)`

SetLastShipDate sets LastShipDate field to given value.

### HasLastShipDate

`func (o *NiaapiHardwareEolAllOf) HasLastShipDate() bool`

HasLastShipDate returns a boolean if a field has been set.

### GetLastShipDateEpoch

`func (o *NiaapiHardwareEolAllOf) GetLastShipDateEpoch() int64`

GetLastShipDateEpoch returns the LastShipDateEpoch field if non-nil, zero value otherwise.

### GetLastShipDateEpochOk

`func (o *NiaapiHardwareEolAllOf) GetLastShipDateEpochOk() (*int64, bool)`

GetLastShipDateEpochOk returns a tuple with the LastShipDateEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastShipDateEpoch

`func (o *NiaapiHardwareEolAllOf) SetLastShipDateEpoch(v int64)`

SetLastShipDateEpoch sets LastShipDateEpoch field to given value.

### HasLastShipDateEpoch

`func (o *NiaapiHardwareEolAllOf) HasLastShipDateEpoch() bool`

HasLastShipDateEpoch returns a boolean if a field has been set.

### GetMigrationUrl

`func (o *NiaapiHardwareEolAllOf) GetMigrationUrl() string`

GetMigrationUrl returns the MigrationUrl field if non-nil, zero value otherwise.

### GetMigrationUrlOk

`func (o *NiaapiHardwareEolAllOf) GetMigrationUrlOk() (*string, bool)`

GetMigrationUrlOk returns a tuple with the MigrationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationUrl

`func (o *NiaapiHardwareEolAllOf) SetMigrationUrl(v string)`

SetMigrationUrl sets MigrationUrl field to given value.

### HasMigrationUrl

`func (o *NiaapiHardwareEolAllOf) HasMigrationUrl() bool`

HasMigrationUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


