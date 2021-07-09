# NiatelemetryApicCoreFileDetailsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.ApicCoreFileDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.ApicCoreFileDetails"]
**Annotation** | Pointer to **string** | Annotation of the Core file in APIC. | [optional] 
**ChildAction** | Pointer to **string** | Child action of the Core file in APIC. | [optional] 
**CollectionTime** | Pointer to **string** | Collection Time of the Core file in APIC. | [optional] 
**DataType** | Pointer to **string** | Data type of the Core file in APIC. | [optional] 
**Dn** | Pointer to **string** | Dn for the Core file in the inventory. | [optional] 
**ExportFileUri** | Pointer to **string** | Export file URI of the Core file in APIC. | [optional] 
**ExportStatus** | Pointer to **string** | Export status of the Core file in APIC. | [optional] 
**ExportStatusStr** | Pointer to **int64** | Export status str of the Core file in APIC. | [optional] 
**ExportTechSupFileUri** | Pointer to **string** | Export tech Sup file URI of the Core file in APIC. | [optional] 
**ExportedToController** | Pointer to **string** | Return if file is exported To Controller or not in APIC. | [optional] 
**ExtMngdBy** | Pointer to **string** | Ext Mngd By of the Core file in APIC. | [optional] 
**FileSize** | Pointer to **int64** | File size of the Core file in APIC. | [optional] 
**HostName** | Pointer to **string** | Host Name of the Core file in APIC. | [optional] 
**LcOwn** | Pointer to **string** | Lc owner of the Core file in APIC. | [optional] 
**ModTs** | Pointer to **int64** | Mod Ts of the Core file in APIC. | [optional] 
**NodeId** | Pointer to **string** | Node Id of the Core file in APIC. | [optional] 
**PolName** | Pointer to **string** | Pol Name of the Core file in APIC. | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**SiteName** | Pointer to **string** | Name of the APIC site from which this data is being collected. | [optional] 
**Status** | Pointer to **string** | Status of the Core file in APIC. | [optional] 
**Uid** | Pointer to **string** | UId of the Core file in the APIC. | [optional] 
**Userdom** | Pointer to **string** | User dom of the Core file in APIC. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewNiatelemetryApicCoreFileDetailsAllOf

`func NewNiatelemetryApicCoreFileDetailsAllOf(classId string, objectType string, ) *NiatelemetryApicCoreFileDetailsAllOf`

NewNiatelemetryApicCoreFileDetailsAllOf instantiates a new NiatelemetryApicCoreFileDetailsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryApicCoreFileDetailsAllOfWithDefaults

`func NewNiatelemetryApicCoreFileDetailsAllOfWithDefaults() *NiatelemetryApicCoreFileDetailsAllOf`

NewNiatelemetryApicCoreFileDetailsAllOfWithDefaults instantiates a new NiatelemetryApicCoreFileDetailsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryApicCoreFileDetailsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryApicCoreFileDetailsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAnnotation

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetAnnotation() string`

GetAnnotation returns the Annotation field if non-nil, zero value otherwise.

### GetAnnotationOk

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetAnnotationOk() (*string, bool)`

GetAnnotationOk returns a tuple with the Annotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotation

`func (o *NiatelemetryApicCoreFileDetailsAllOf) SetAnnotation(v string)`

SetAnnotation sets Annotation field to given value.

### HasAnnotation

`func (o *NiatelemetryApicCoreFileDetailsAllOf) HasAnnotation() bool`

HasAnnotation returns a boolean if a field has been set.

### GetChildAction

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetChildAction() string`

GetChildAction returns the ChildAction field if non-nil, zero value otherwise.

### GetChildActionOk

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetChildActionOk() (*string, bool)`

GetChildActionOk returns a tuple with the ChildAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildAction

`func (o *NiatelemetryApicCoreFileDetailsAllOf) SetChildAction(v string)`

SetChildAction sets ChildAction field to given value.

### HasChildAction

`func (o *NiatelemetryApicCoreFileDetailsAllOf) HasChildAction() bool`

HasChildAction returns a boolean if a field has been set.

### GetCollectionTime

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetCollectionTime() string`

GetCollectionTime returns the CollectionTime field if non-nil, zero value otherwise.

### GetCollectionTimeOk

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetCollectionTimeOk() (*string, bool)`

GetCollectionTimeOk returns a tuple with the CollectionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionTime

`func (o *NiatelemetryApicCoreFileDetailsAllOf) SetCollectionTime(v string)`

SetCollectionTime sets CollectionTime field to given value.

### HasCollectionTime

`func (o *NiatelemetryApicCoreFileDetailsAllOf) HasCollectionTime() bool`

HasCollectionTime returns a boolean if a field has been set.

### GetDataType

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *NiatelemetryApicCoreFileDetailsAllOf) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *NiatelemetryApicCoreFileDetailsAllOf) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetDn

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetryApicCoreFileDetailsAllOf) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetryApicCoreFileDetailsAllOf) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetExportFileUri

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetExportFileUri() string`

GetExportFileUri returns the ExportFileUri field if non-nil, zero value otherwise.

### GetExportFileUriOk

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetExportFileUriOk() (*string, bool)`

GetExportFileUriOk returns a tuple with the ExportFileUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportFileUri

`func (o *NiatelemetryApicCoreFileDetailsAllOf) SetExportFileUri(v string)`

SetExportFileUri sets ExportFileUri field to given value.

### HasExportFileUri

`func (o *NiatelemetryApicCoreFileDetailsAllOf) HasExportFileUri() bool`

HasExportFileUri returns a boolean if a field has been set.

### GetExportStatus

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetExportStatus() string`

GetExportStatus returns the ExportStatus field if non-nil, zero value otherwise.

### GetExportStatusOk

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetExportStatusOk() (*string, bool)`

GetExportStatusOk returns a tuple with the ExportStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportStatus

`func (o *NiatelemetryApicCoreFileDetailsAllOf) SetExportStatus(v string)`

SetExportStatus sets ExportStatus field to given value.

### HasExportStatus

`func (o *NiatelemetryApicCoreFileDetailsAllOf) HasExportStatus() bool`

HasExportStatus returns a boolean if a field has been set.

### GetExportStatusStr

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetExportStatusStr() int64`

GetExportStatusStr returns the ExportStatusStr field if non-nil, zero value otherwise.

### GetExportStatusStrOk

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetExportStatusStrOk() (*int64, bool)`

GetExportStatusStrOk returns a tuple with the ExportStatusStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportStatusStr

`func (o *NiatelemetryApicCoreFileDetailsAllOf) SetExportStatusStr(v int64)`

SetExportStatusStr sets ExportStatusStr field to given value.

### HasExportStatusStr

`func (o *NiatelemetryApicCoreFileDetailsAllOf) HasExportStatusStr() bool`

HasExportStatusStr returns a boolean if a field has been set.

### GetExportTechSupFileUri

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetExportTechSupFileUri() string`

GetExportTechSupFileUri returns the ExportTechSupFileUri field if non-nil, zero value otherwise.

### GetExportTechSupFileUriOk

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetExportTechSupFileUriOk() (*string, bool)`

GetExportTechSupFileUriOk returns a tuple with the ExportTechSupFileUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportTechSupFileUri

`func (o *NiatelemetryApicCoreFileDetailsAllOf) SetExportTechSupFileUri(v string)`

SetExportTechSupFileUri sets ExportTechSupFileUri field to given value.

### HasExportTechSupFileUri

`func (o *NiatelemetryApicCoreFileDetailsAllOf) HasExportTechSupFileUri() bool`

HasExportTechSupFileUri returns a boolean if a field has been set.

### GetExportedToController

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetExportedToController() string`

GetExportedToController returns the ExportedToController field if non-nil, zero value otherwise.

### GetExportedToControllerOk

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetExportedToControllerOk() (*string, bool)`

GetExportedToControllerOk returns a tuple with the ExportedToController field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportedToController

`func (o *NiatelemetryApicCoreFileDetailsAllOf) SetExportedToController(v string)`

SetExportedToController sets ExportedToController field to given value.

### HasExportedToController

`func (o *NiatelemetryApicCoreFileDetailsAllOf) HasExportedToController() bool`

HasExportedToController returns a boolean if a field has been set.

### GetExtMngdBy

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetExtMngdBy() string`

GetExtMngdBy returns the ExtMngdBy field if non-nil, zero value otherwise.

### GetExtMngdByOk

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetExtMngdByOk() (*string, bool)`

GetExtMngdByOk returns a tuple with the ExtMngdBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtMngdBy

`func (o *NiatelemetryApicCoreFileDetailsAllOf) SetExtMngdBy(v string)`

SetExtMngdBy sets ExtMngdBy field to given value.

### HasExtMngdBy

`func (o *NiatelemetryApicCoreFileDetailsAllOf) HasExtMngdBy() bool`

HasExtMngdBy returns a boolean if a field has been set.

### GetFileSize

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetFileSize() int64`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetFileSizeOk() (*int64, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *NiatelemetryApicCoreFileDetailsAllOf) SetFileSize(v int64)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *NiatelemetryApicCoreFileDetailsAllOf) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetHostName

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *NiatelemetryApicCoreFileDetailsAllOf) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *NiatelemetryApicCoreFileDetailsAllOf) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetLcOwn

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetLcOwn() string`

GetLcOwn returns the LcOwn field if non-nil, zero value otherwise.

### GetLcOwnOk

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetLcOwnOk() (*string, bool)`

GetLcOwnOk returns a tuple with the LcOwn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcOwn

`func (o *NiatelemetryApicCoreFileDetailsAllOf) SetLcOwn(v string)`

SetLcOwn sets LcOwn field to given value.

### HasLcOwn

`func (o *NiatelemetryApicCoreFileDetailsAllOf) HasLcOwn() bool`

HasLcOwn returns a boolean if a field has been set.

### GetModTs

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetModTs() int64`

GetModTs returns the ModTs field if non-nil, zero value otherwise.

### GetModTsOk

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetModTsOk() (*int64, bool)`

GetModTsOk returns a tuple with the ModTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTs

`func (o *NiatelemetryApicCoreFileDetailsAllOf) SetModTs(v int64)`

SetModTs sets ModTs field to given value.

### HasModTs

`func (o *NiatelemetryApicCoreFileDetailsAllOf) HasModTs() bool`

HasModTs returns a boolean if a field has been set.

### GetNodeId

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *NiatelemetryApicCoreFileDetailsAllOf) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *NiatelemetryApicCoreFileDetailsAllOf) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetPolName

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetPolName() string`

GetPolName returns the PolName field if non-nil, zero value otherwise.

### GetPolNameOk

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetPolNameOk() (*string, bool)`

GetPolNameOk returns a tuple with the PolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolName

`func (o *NiatelemetryApicCoreFileDetailsAllOf) SetPolName(v string)`

SetPolName sets PolName field to given value.

### HasPolName

`func (o *NiatelemetryApicCoreFileDetailsAllOf) HasPolName() bool`

HasPolName returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryApicCoreFileDetailsAllOf) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryApicCoreFileDetailsAllOf) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryApicCoreFileDetailsAllOf) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryApicCoreFileDetailsAllOf) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryApicCoreFileDetailsAllOf) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryApicCoreFileDetailsAllOf) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetStatus

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NiatelemetryApicCoreFileDetailsAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NiatelemetryApicCoreFileDetailsAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUid

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *NiatelemetryApicCoreFileDetailsAllOf) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *NiatelemetryApicCoreFileDetailsAllOf) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetUserdom

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetUserdom() string`

GetUserdom returns the Userdom field if non-nil, zero value otherwise.

### GetUserdomOk

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetUserdomOk() (*string, bool)`

GetUserdomOk returns a tuple with the Userdom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserdom

`func (o *NiatelemetryApicCoreFileDetailsAllOf) SetUserdom(v string)`

SetUserdom sets Userdom field to given value.

### HasUserdom

`func (o *NiatelemetryApicCoreFileDetailsAllOf) HasUserdom() bool`

HasUserdom returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryApicCoreFileDetailsAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryApicCoreFileDetailsAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryApicCoreFileDetailsAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


