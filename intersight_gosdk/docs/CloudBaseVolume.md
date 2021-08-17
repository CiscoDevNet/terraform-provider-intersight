# CloudBaseVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "cloud.AwsVolume"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "cloud.AwsVolume"]
**BillingUnit** | Pointer to [**NullableCloudBillingUnit**](cloud.BillingUnit.md) |  | [optional] 
**EncryptionState** | Pointer to **string** | Encryption state of this volume.For example ENCRYPTED, NOT ENCRYPTED, etc. * &#x60;Automatic&#x60; - Volume encryption state is automatic.Cloud provider takes the decision of when to encrypt the data. * &#x60;Encrypted&#x60; - Volume data is encrypted. Can be decrypted only by authorized users. * &#x60;Not_Encrypted&#x60; - Volume data is not encrypted. | [optional] [readonly] [default to "Automatic"]
**Identity** | Pointer to **string** | The internally generated identity of this VM. It aids in uniquely identifying the volume object. | [optional] [readonly] 
**InstanceAttachments** | Pointer to [**[]CloudVolumeInstanceAttachment**](CloudVolumeInstanceAttachment.md) |  | [optional] 
**IopsInfo** | Pointer to [**NullableCloudVolumeIopsInfo**](cloud.VolumeIopsInfo.md) |  | [optional] 
**RegionInfo** | Pointer to [**NullableCloudCloudRegion**](cloud.CloudRegion.md) |  | [optional] 
**SourceImageId** | Pointer to **string** | Source Image Id used for the volume. | [optional] [readonly] 
**State** | Pointer to **string** | The volume state.For example AVAILABLE, IN_USE, DELETED, etc. * &#x60;UnRecognized&#x60; - Volume is in unrecognized state. * &#x60;Pending&#x60; - Volume is  in pending state, due to errors encountered during volume creation. * &#x60;Bound&#x60; - Volume is in bound state. * &#x60;Available&#x60; - Volume is in available state. * &#x60;Error&#x60; - Volume is in error state. * &#x60;Released&#x60; - Volume is in released state. * &#x60;in-use&#x60; - Volume is in in-use state. * &#x60;Creating&#x60; - Volume is in creating state. * &#x60;Deleting&#x60; - Volume is in deleting state. * &#x60;Deleted&#x60; - Volume is in deleted state. | [optional] [readonly] [default to "UnRecognized"]
**Uuid** | Pointer to **string** | UUID (Universally Unique IDentifier) is a 128-bit value used for unique identification of Volume. | [optional] [readonly] 
**VolumeCreateTime** | Pointer to **time.Time** | Time when this volume is created. | [optional] [readonly] 
**VolumeTags** | Pointer to [**[]CloudCloudTag**](CloudCloudTag.md) |  | [optional] 
**VolumeType** | Pointer to [**NullableCloudVolumeType**](cloud.VolumeType.md) |  | [optional] 
**ZoneInfo** | Pointer to [**NullableCloudAvailabilityZone**](cloud.AvailabilityZone.md) |  | [optional] 

## Methods

### NewCloudBaseVolume

`func NewCloudBaseVolume(classId string, objectType string, ) *CloudBaseVolume`

NewCloudBaseVolume instantiates a new CloudBaseVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBaseVolumeWithDefaults

`func NewCloudBaseVolumeWithDefaults() *CloudBaseVolume`

NewCloudBaseVolumeWithDefaults instantiates a new CloudBaseVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudBaseVolume) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudBaseVolume) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudBaseVolume) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudBaseVolume) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudBaseVolume) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudBaseVolume) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBillingUnit

`func (o *CloudBaseVolume) GetBillingUnit() CloudBillingUnit`

GetBillingUnit returns the BillingUnit field if non-nil, zero value otherwise.

### GetBillingUnitOk

`func (o *CloudBaseVolume) GetBillingUnitOk() (*CloudBillingUnit, bool)`

GetBillingUnitOk returns a tuple with the BillingUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingUnit

`func (o *CloudBaseVolume) SetBillingUnit(v CloudBillingUnit)`

SetBillingUnit sets BillingUnit field to given value.

### HasBillingUnit

`func (o *CloudBaseVolume) HasBillingUnit() bool`

HasBillingUnit returns a boolean if a field has been set.

### SetBillingUnitNil

`func (o *CloudBaseVolume) SetBillingUnitNil(b bool)`

 SetBillingUnitNil sets the value for BillingUnit to be an explicit nil

### UnsetBillingUnit
`func (o *CloudBaseVolume) UnsetBillingUnit()`

UnsetBillingUnit ensures that no value is present for BillingUnit, not even an explicit nil
### GetEncryptionState

`func (o *CloudBaseVolume) GetEncryptionState() string`

GetEncryptionState returns the EncryptionState field if non-nil, zero value otherwise.

### GetEncryptionStateOk

`func (o *CloudBaseVolume) GetEncryptionStateOk() (*string, bool)`

GetEncryptionStateOk returns a tuple with the EncryptionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionState

`func (o *CloudBaseVolume) SetEncryptionState(v string)`

SetEncryptionState sets EncryptionState field to given value.

### HasEncryptionState

`func (o *CloudBaseVolume) HasEncryptionState() bool`

HasEncryptionState returns a boolean if a field has been set.

### GetIdentity

`func (o *CloudBaseVolume) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *CloudBaseVolume) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *CloudBaseVolume) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *CloudBaseVolume) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetInstanceAttachments

`func (o *CloudBaseVolume) GetInstanceAttachments() []CloudVolumeInstanceAttachment`

GetInstanceAttachments returns the InstanceAttachments field if non-nil, zero value otherwise.

### GetInstanceAttachmentsOk

`func (o *CloudBaseVolume) GetInstanceAttachmentsOk() (*[]CloudVolumeInstanceAttachment, bool)`

GetInstanceAttachmentsOk returns a tuple with the InstanceAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceAttachments

`func (o *CloudBaseVolume) SetInstanceAttachments(v []CloudVolumeInstanceAttachment)`

SetInstanceAttachments sets InstanceAttachments field to given value.

### HasInstanceAttachments

`func (o *CloudBaseVolume) HasInstanceAttachments() bool`

HasInstanceAttachments returns a boolean if a field has been set.

### SetInstanceAttachmentsNil

`func (o *CloudBaseVolume) SetInstanceAttachmentsNil(b bool)`

 SetInstanceAttachmentsNil sets the value for InstanceAttachments to be an explicit nil

### UnsetInstanceAttachments
`func (o *CloudBaseVolume) UnsetInstanceAttachments()`

UnsetInstanceAttachments ensures that no value is present for InstanceAttachments, not even an explicit nil
### GetIopsInfo

`func (o *CloudBaseVolume) GetIopsInfo() CloudVolumeIopsInfo`

GetIopsInfo returns the IopsInfo field if non-nil, zero value otherwise.

### GetIopsInfoOk

`func (o *CloudBaseVolume) GetIopsInfoOk() (*CloudVolumeIopsInfo, bool)`

GetIopsInfoOk returns a tuple with the IopsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsInfo

`func (o *CloudBaseVolume) SetIopsInfo(v CloudVolumeIopsInfo)`

SetIopsInfo sets IopsInfo field to given value.

### HasIopsInfo

`func (o *CloudBaseVolume) HasIopsInfo() bool`

HasIopsInfo returns a boolean if a field has been set.

### SetIopsInfoNil

`func (o *CloudBaseVolume) SetIopsInfoNil(b bool)`

 SetIopsInfoNil sets the value for IopsInfo to be an explicit nil

### UnsetIopsInfo
`func (o *CloudBaseVolume) UnsetIopsInfo()`

UnsetIopsInfo ensures that no value is present for IopsInfo, not even an explicit nil
### GetRegionInfo

`func (o *CloudBaseVolume) GetRegionInfo() CloudCloudRegion`

GetRegionInfo returns the RegionInfo field if non-nil, zero value otherwise.

### GetRegionInfoOk

`func (o *CloudBaseVolume) GetRegionInfoOk() (*CloudCloudRegion, bool)`

GetRegionInfoOk returns a tuple with the RegionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionInfo

`func (o *CloudBaseVolume) SetRegionInfo(v CloudCloudRegion)`

SetRegionInfo sets RegionInfo field to given value.

### HasRegionInfo

`func (o *CloudBaseVolume) HasRegionInfo() bool`

HasRegionInfo returns a boolean if a field has been set.

### SetRegionInfoNil

`func (o *CloudBaseVolume) SetRegionInfoNil(b bool)`

 SetRegionInfoNil sets the value for RegionInfo to be an explicit nil

### UnsetRegionInfo
`func (o *CloudBaseVolume) UnsetRegionInfo()`

UnsetRegionInfo ensures that no value is present for RegionInfo, not even an explicit nil
### GetSourceImageId

`func (o *CloudBaseVolume) GetSourceImageId() string`

GetSourceImageId returns the SourceImageId field if non-nil, zero value otherwise.

### GetSourceImageIdOk

`func (o *CloudBaseVolume) GetSourceImageIdOk() (*string, bool)`

GetSourceImageIdOk returns a tuple with the SourceImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceImageId

`func (o *CloudBaseVolume) SetSourceImageId(v string)`

SetSourceImageId sets SourceImageId field to given value.

### HasSourceImageId

`func (o *CloudBaseVolume) HasSourceImageId() bool`

HasSourceImageId returns a boolean if a field has been set.

### GetState

`func (o *CloudBaseVolume) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CloudBaseVolume) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CloudBaseVolume) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CloudBaseVolume) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUuid

`func (o *CloudBaseVolume) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CloudBaseVolume) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CloudBaseVolume) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *CloudBaseVolume) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVolumeCreateTime

`func (o *CloudBaseVolume) GetVolumeCreateTime() time.Time`

GetVolumeCreateTime returns the VolumeCreateTime field if non-nil, zero value otherwise.

### GetVolumeCreateTimeOk

`func (o *CloudBaseVolume) GetVolumeCreateTimeOk() (*time.Time, bool)`

GetVolumeCreateTimeOk returns a tuple with the VolumeCreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeCreateTime

`func (o *CloudBaseVolume) SetVolumeCreateTime(v time.Time)`

SetVolumeCreateTime sets VolumeCreateTime field to given value.

### HasVolumeCreateTime

`func (o *CloudBaseVolume) HasVolumeCreateTime() bool`

HasVolumeCreateTime returns a boolean if a field has been set.

### GetVolumeTags

`func (o *CloudBaseVolume) GetVolumeTags() []CloudCloudTag`

GetVolumeTags returns the VolumeTags field if non-nil, zero value otherwise.

### GetVolumeTagsOk

`func (o *CloudBaseVolume) GetVolumeTagsOk() (*[]CloudCloudTag, bool)`

GetVolumeTagsOk returns a tuple with the VolumeTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeTags

`func (o *CloudBaseVolume) SetVolumeTags(v []CloudCloudTag)`

SetVolumeTags sets VolumeTags field to given value.

### HasVolumeTags

`func (o *CloudBaseVolume) HasVolumeTags() bool`

HasVolumeTags returns a boolean if a field has been set.

### SetVolumeTagsNil

`func (o *CloudBaseVolume) SetVolumeTagsNil(b bool)`

 SetVolumeTagsNil sets the value for VolumeTags to be an explicit nil

### UnsetVolumeTags
`func (o *CloudBaseVolume) UnsetVolumeTags()`

UnsetVolumeTags ensures that no value is present for VolumeTags, not even an explicit nil
### GetVolumeType

`func (o *CloudBaseVolume) GetVolumeType() CloudVolumeType`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *CloudBaseVolume) GetVolumeTypeOk() (*CloudVolumeType, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *CloudBaseVolume) SetVolumeType(v CloudVolumeType)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *CloudBaseVolume) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### SetVolumeTypeNil

`func (o *CloudBaseVolume) SetVolumeTypeNil(b bool)`

 SetVolumeTypeNil sets the value for VolumeType to be an explicit nil

### UnsetVolumeType
`func (o *CloudBaseVolume) UnsetVolumeType()`

UnsetVolumeType ensures that no value is present for VolumeType, not even an explicit nil
### GetZoneInfo

`func (o *CloudBaseVolume) GetZoneInfo() CloudAvailabilityZone`

GetZoneInfo returns the ZoneInfo field if non-nil, zero value otherwise.

### GetZoneInfoOk

`func (o *CloudBaseVolume) GetZoneInfoOk() (*CloudAvailabilityZone, bool)`

GetZoneInfoOk returns a tuple with the ZoneInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneInfo

`func (o *CloudBaseVolume) SetZoneInfo(v CloudAvailabilityZone)`

SetZoneInfo sets ZoneInfo field to given value.

### HasZoneInfo

`func (o *CloudBaseVolume) HasZoneInfo() bool`

HasZoneInfo returns a boolean if a field has been set.

### SetZoneInfoNil

`func (o *CloudBaseVolume) SetZoneInfoNil(b bool)`

 SetZoneInfoNil sets the value for ZoneInfo to be an explicit nil

### UnsetZoneInfo
`func (o *CloudBaseVolume) UnsetZoneInfo()`

UnsetZoneInfo ensures that no value is present for ZoneInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


