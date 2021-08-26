# HyperflexDatastoreStatistic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.DatastoreStatistic"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.DatastoreStatistic"]
**AccessibilitySummary** | Pointer to **string** | HyperFlex datastore accessibility summary. * &#x60;ACCESSIBLE&#x60; - The HyperFlex Accessibility Summary is Accessible. * &#x60;NOT_ACCESSIBLE&#x60; - The HyperFlex Accessibility Summary is Not Accessible. * &#x60;PARTIALLY_ACCESSIBLE&#x60; - The HyperFlex Accessibility Summary is Partially Accessible. | [optional] [readonly] [default to "ACCESSIBLE"]
**CreationTime** | Pointer to **string** | Timestamp the datastore object was created. | [optional] [readonly] 
**DatastoreStatus** | Pointer to **string** | HyperFlex datastore status. * &#x60;NORMAL&#x60; - The HyperFlex datastore status is normal. * &#x60;ALERT&#x60; - The HyperFlex datastore status is alert. * &#x60;FAILED&#x60; - The HyperFlex datastore status is failed. | [optional] [readonly] [default to "NORMAL"]
**Dsconfig** | Pointer to [**NullableHyperflexHxPlatformDatastoreConfigDt**](HyperflexHxPlatformDatastoreConfigDt.md) |  | [optional] 
**FreeCapacityInBytes** | Pointer to **int64** | Free capacity of the datastore in bytes. | [optional] [readonly] 
**HostMountStatus** | Pointer to [**[]HyperflexHxHostMountStatusDt**](HyperflexHxHostMountStatusDt.md) |  | [optional] 
**LastAccessTime** | Pointer to **string** | Timestamp the datastore object was last accessed. | [optional] [readonly] 
**LastModifiedTime** | Pointer to **string** | Timestamp the datastore object was last modified. | [optional] [readonly] 
**MountSummary** | Pointer to **string** | HyperFlex datastore mount summary. * &#x60;MOUNTED&#x60; - The HyperFlex mount summary is mounted. * &#x60;UNMOUNTED&#x60; - The HyperFlex mount summary is unmounted. * &#x60;MOUNT_FAILURE&#x60; - The HyperFlex mount summary is mount failure. * &#x60;UNMOUNT_FAILURE&#x60; - The HyperFlex mount summary is unmount failure. | [optional] [readonly] [default to "MOUNTED"]
**ParentUuid** | Pointer to **string** | UUID of the parent datastore object. | [optional] [readonly] 
**Site** | Pointer to [**NullableHyperflexHxSiteDt**](HyperflexHxSiteDt.md) |  | [optional] 
**TotalCapacityInBytes** | Pointer to **int64** | Total capacity of the datastore object. | [optional] [readonly] 
**UnCompressedUsedBytes** | Pointer to **int64** | Number of uncompressed used bytes in the datastore. | [optional] [readonly] 
**UnsharedUsedBytes** | Pointer to **int64** | Unshared used capacity of the datastore in bytes. | [optional] [readonly] 
**Uuid** | Pointer to **string** | UUID for the datastore object. | [optional] [readonly] 
**DataProtectionPeer** | Pointer to [**HyperflexDataProtectionPeerRelationship**](HyperflexDataProtectionPeerRelationship.md) |  | [optional] 

## Methods

### NewHyperflexDatastoreStatistic

`func NewHyperflexDatastoreStatistic(classId string, objectType string, ) *HyperflexDatastoreStatistic`

NewHyperflexDatastoreStatistic instantiates a new HyperflexDatastoreStatistic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexDatastoreStatisticWithDefaults

`func NewHyperflexDatastoreStatisticWithDefaults() *HyperflexDatastoreStatistic`

NewHyperflexDatastoreStatisticWithDefaults instantiates a new HyperflexDatastoreStatistic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexDatastoreStatistic) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexDatastoreStatistic) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexDatastoreStatistic) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexDatastoreStatistic) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexDatastoreStatistic) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexDatastoreStatistic) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccessibilitySummary

`func (o *HyperflexDatastoreStatistic) GetAccessibilitySummary() string`

GetAccessibilitySummary returns the AccessibilitySummary field if non-nil, zero value otherwise.

### GetAccessibilitySummaryOk

`func (o *HyperflexDatastoreStatistic) GetAccessibilitySummaryOk() (*string, bool)`

GetAccessibilitySummaryOk returns a tuple with the AccessibilitySummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessibilitySummary

`func (o *HyperflexDatastoreStatistic) SetAccessibilitySummary(v string)`

SetAccessibilitySummary sets AccessibilitySummary field to given value.

### HasAccessibilitySummary

`func (o *HyperflexDatastoreStatistic) HasAccessibilitySummary() bool`

HasAccessibilitySummary returns a boolean if a field has been set.

### GetCreationTime

`func (o *HyperflexDatastoreStatistic) GetCreationTime() string`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *HyperflexDatastoreStatistic) GetCreationTimeOk() (*string, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *HyperflexDatastoreStatistic) SetCreationTime(v string)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *HyperflexDatastoreStatistic) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetDatastoreStatus

`func (o *HyperflexDatastoreStatistic) GetDatastoreStatus() string`

GetDatastoreStatus returns the DatastoreStatus field if non-nil, zero value otherwise.

### GetDatastoreStatusOk

`func (o *HyperflexDatastoreStatistic) GetDatastoreStatusOk() (*string, bool)`

GetDatastoreStatusOk returns a tuple with the DatastoreStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreStatus

`func (o *HyperflexDatastoreStatistic) SetDatastoreStatus(v string)`

SetDatastoreStatus sets DatastoreStatus field to given value.

### HasDatastoreStatus

`func (o *HyperflexDatastoreStatistic) HasDatastoreStatus() bool`

HasDatastoreStatus returns a boolean if a field has been set.

### GetDsconfig

`func (o *HyperflexDatastoreStatistic) GetDsconfig() HyperflexHxPlatformDatastoreConfigDt`

GetDsconfig returns the Dsconfig field if non-nil, zero value otherwise.

### GetDsconfigOk

`func (o *HyperflexDatastoreStatistic) GetDsconfigOk() (*HyperflexHxPlatformDatastoreConfigDt, bool)`

GetDsconfigOk returns a tuple with the Dsconfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsconfig

`func (o *HyperflexDatastoreStatistic) SetDsconfig(v HyperflexHxPlatformDatastoreConfigDt)`

SetDsconfig sets Dsconfig field to given value.

### HasDsconfig

`func (o *HyperflexDatastoreStatistic) HasDsconfig() bool`

HasDsconfig returns a boolean if a field has been set.

### SetDsconfigNil

`func (o *HyperflexDatastoreStatistic) SetDsconfigNil(b bool)`

 SetDsconfigNil sets the value for Dsconfig to be an explicit nil

### UnsetDsconfig
`func (o *HyperflexDatastoreStatistic) UnsetDsconfig()`

UnsetDsconfig ensures that no value is present for Dsconfig, not even an explicit nil
### GetFreeCapacityInBytes

`func (o *HyperflexDatastoreStatistic) GetFreeCapacityInBytes() int64`

GetFreeCapacityInBytes returns the FreeCapacityInBytes field if non-nil, zero value otherwise.

### GetFreeCapacityInBytesOk

`func (o *HyperflexDatastoreStatistic) GetFreeCapacityInBytesOk() (*int64, bool)`

GetFreeCapacityInBytesOk returns a tuple with the FreeCapacityInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeCapacityInBytes

`func (o *HyperflexDatastoreStatistic) SetFreeCapacityInBytes(v int64)`

SetFreeCapacityInBytes sets FreeCapacityInBytes field to given value.

### HasFreeCapacityInBytes

`func (o *HyperflexDatastoreStatistic) HasFreeCapacityInBytes() bool`

HasFreeCapacityInBytes returns a boolean if a field has been set.

### GetHostMountStatus

`func (o *HyperflexDatastoreStatistic) GetHostMountStatus() []HyperflexHxHostMountStatusDt`

GetHostMountStatus returns the HostMountStatus field if non-nil, zero value otherwise.

### GetHostMountStatusOk

`func (o *HyperflexDatastoreStatistic) GetHostMountStatusOk() (*[]HyperflexHxHostMountStatusDt, bool)`

GetHostMountStatusOk returns a tuple with the HostMountStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostMountStatus

`func (o *HyperflexDatastoreStatistic) SetHostMountStatus(v []HyperflexHxHostMountStatusDt)`

SetHostMountStatus sets HostMountStatus field to given value.

### HasHostMountStatus

`func (o *HyperflexDatastoreStatistic) HasHostMountStatus() bool`

HasHostMountStatus returns a boolean if a field has been set.

### SetHostMountStatusNil

`func (o *HyperflexDatastoreStatistic) SetHostMountStatusNil(b bool)`

 SetHostMountStatusNil sets the value for HostMountStatus to be an explicit nil

### UnsetHostMountStatus
`func (o *HyperflexDatastoreStatistic) UnsetHostMountStatus()`

UnsetHostMountStatus ensures that no value is present for HostMountStatus, not even an explicit nil
### GetLastAccessTime

`func (o *HyperflexDatastoreStatistic) GetLastAccessTime() string`

GetLastAccessTime returns the LastAccessTime field if non-nil, zero value otherwise.

### GetLastAccessTimeOk

`func (o *HyperflexDatastoreStatistic) GetLastAccessTimeOk() (*string, bool)`

GetLastAccessTimeOk returns a tuple with the LastAccessTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessTime

`func (o *HyperflexDatastoreStatistic) SetLastAccessTime(v string)`

SetLastAccessTime sets LastAccessTime field to given value.

### HasLastAccessTime

`func (o *HyperflexDatastoreStatistic) HasLastAccessTime() bool`

HasLastAccessTime returns a boolean if a field has been set.

### GetLastModifiedTime

`func (o *HyperflexDatastoreStatistic) GetLastModifiedTime() string`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *HyperflexDatastoreStatistic) GetLastModifiedTimeOk() (*string, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *HyperflexDatastoreStatistic) SetLastModifiedTime(v string)`

SetLastModifiedTime sets LastModifiedTime field to given value.

### HasLastModifiedTime

`func (o *HyperflexDatastoreStatistic) HasLastModifiedTime() bool`

HasLastModifiedTime returns a boolean if a field has been set.

### GetMountSummary

`func (o *HyperflexDatastoreStatistic) GetMountSummary() string`

GetMountSummary returns the MountSummary field if non-nil, zero value otherwise.

### GetMountSummaryOk

`func (o *HyperflexDatastoreStatistic) GetMountSummaryOk() (*string, bool)`

GetMountSummaryOk returns a tuple with the MountSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountSummary

`func (o *HyperflexDatastoreStatistic) SetMountSummary(v string)`

SetMountSummary sets MountSummary field to given value.

### HasMountSummary

`func (o *HyperflexDatastoreStatistic) HasMountSummary() bool`

HasMountSummary returns a boolean if a field has been set.

### GetParentUuid

`func (o *HyperflexDatastoreStatistic) GetParentUuid() string`

GetParentUuid returns the ParentUuid field if non-nil, zero value otherwise.

### GetParentUuidOk

`func (o *HyperflexDatastoreStatistic) GetParentUuidOk() (*string, bool)`

GetParentUuidOk returns a tuple with the ParentUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentUuid

`func (o *HyperflexDatastoreStatistic) SetParentUuid(v string)`

SetParentUuid sets ParentUuid field to given value.

### HasParentUuid

`func (o *HyperflexDatastoreStatistic) HasParentUuid() bool`

HasParentUuid returns a boolean if a field has been set.

### GetSite

`func (o *HyperflexDatastoreStatistic) GetSite() HyperflexHxSiteDt`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *HyperflexDatastoreStatistic) GetSiteOk() (*HyperflexHxSiteDt, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *HyperflexDatastoreStatistic) SetSite(v HyperflexHxSiteDt)`

SetSite sets Site field to given value.

### HasSite

`func (o *HyperflexDatastoreStatistic) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSiteNil

`func (o *HyperflexDatastoreStatistic) SetSiteNil(b bool)`

 SetSiteNil sets the value for Site to be an explicit nil

### UnsetSite
`func (o *HyperflexDatastoreStatistic) UnsetSite()`

UnsetSite ensures that no value is present for Site, not even an explicit nil
### GetTotalCapacityInBytes

`func (o *HyperflexDatastoreStatistic) GetTotalCapacityInBytes() int64`

GetTotalCapacityInBytes returns the TotalCapacityInBytes field if non-nil, zero value otherwise.

### GetTotalCapacityInBytesOk

`func (o *HyperflexDatastoreStatistic) GetTotalCapacityInBytesOk() (*int64, bool)`

GetTotalCapacityInBytesOk returns a tuple with the TotalCapacityInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCapacityInBytes

`func (o *HyperflexDatastoreStatistic) SetTotalCapacityInBytes(v int64)`

SetTotalCapacityInBytes sets TotalCapacityInBytes field to given value.

### HasTotalCapacityInBytes

`func (o *HyperflexDatastoreStatistic) HasTotalCapacityInBytes() bool`

HasTotalCapacityInBytes returns a boolean if a field has been set.

### GetUnCompressedUsedBytes

`func (o *HyperflexDatastoreStatistic) GetUnCompressedUsedBytes() int64`

GetUnCompressedUsedBytes returns the UnCompressedUsedBytes field if non-nil, zero value otherwise.

### GetUnCompressedUsedBytesOk

`func (o *HyperflexDatastoreStatistic) GetUnCompressedUsedBytesOk() (*int64, bool)`

GetUnCompressedUsedBytesOk returns a tuple with the UnCompressedUsedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnCompressedUsedBytes

`func (o *HyperflexDatastoreStatistic) SetUnCompressedUsedBytes(v int64)`

SetUnCompressedUsedBytes sets UnCompressedUsedBytes field to given value.

### HasUnCompressedUsedBytes

`func (o *HyperflexDatastoreStatistic) HasUnCompressedUsedBytes() bool`

HasUnCompressedUsedBytes returns a boolean if a field has been set.

### GetUnsharedUsedBytes

`func (o *HyperflexDatastoreStatistic) GetUnsharedUsedBytes() int64`

GetUnsharedUsedBytes returns the UnsharedUsedBytes field if non-nil, zero value otherwise.

### GetUnsharedUsedBytesOk

`func (o *HyperflexDatastoreStatistic) GetUnsharedUsedBytesOk() (*int64, bool)`

GetUnsharedUsedBytesOk returns a tuple with the UnsharedUsedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsharedUsedBytes

`func (o *HyperflexDatastoreStatistic) SetUnsharedUsedBytes(v int64)`

SetUnsharedUsedBytes sets UnsharedUsedBytes field to given value.

### HasUnsharedUsedBytes

`func (o *HyperflexDatastoreStatistic) HasUnsharedUsedBytes() bool`

HasUnsharedUsedBytes returns a boolean if a field has been set.

### GetUuid

`func (o *HyperflexDatastoreStatistic) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *HyperflexDatastoreStatistic) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *HyperflexDatastoreStatistic) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *HyperflexDatastoreStatistic) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetDataProtectionPeer

`func (o *HyperflexDatastoreStatistic) GetDataProtectionPeer() HyperflexDataProtectionPeerRelationship`

GetDataProtectionPeer returns the DataProtectionPeer field if non-nil, zero value otherwise.

### GetDataProtectionPeerOk

`func (o *HyperflexDatastoreStatistic) GetDataProtectionPeerOk() (*HyperflexDataProtectionPeerRelationship, bool)`

GetDataProtectionPeerOk returns a tuple with the DataProtectionPeer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataProtectionPeer

`func (o *HyperflexDatastoreStatistic) SetDataProtectionPeer(v HyperflexDataProtectionPeerRelationship)`

SetDataProtectionPeer sets DataProtectionPeer field to given value.

### HasDataProtectionPeer

`func (o *HyperflexDatastoreStatistic) HasDataProtectionPeer() bool`

HasDataProtectionPeer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


