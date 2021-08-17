# CloudBaseVolumeAllOf

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

### NewCloudBaseVolumeAllOf

`func NewCloudBaseVolumeAllOf(classId string, objectType string, ) *CloudBaseVolumeAllOf`

NewCloudBaseVolumeAllOf instantiates a new CloudBaseVolumeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBaseVolumeAllOfWithDefaults

`func NewCloudBaseVolumeAllOfWithDefaults() *CloudBaseVolumeAllOf`

NewCloudBaseVolumeAllOfWithDefaults instantiates a new CloudBaseVolumeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudBaseVolumeAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudBaseVolumeAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudBaseVolumeAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudBaseVolumeAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudBaseVolumeAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudBaseVolumeAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBillingUnit

`func (o *CloudBaseVolumeAllOf) GetBillingUnit() CloudBillingUnit`

GetBillingUnit returns the BillingUnit field if non-nil, zero value otherwise.

### GetBillingUnitOk

`func (o *CloudBaseVolumeAllOf) GetBillingUnitOk() (*CloudBillingUnit, bool)`

GetBillingUnitOk returns a tuple with the BillingUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingUnit

`func (o *CloudBaseVolumeAllOf) SetBillingUnit(v CloudBillingUnit)`

SetBillingUnit sets BillingUnit field to given value.

### HasBillingUnit

`func (o *CloudBaseVolumeAllOf) HasBillingUnit() bool`

HasBillingUnit returns a boolean if a field has been set.

### SetBillingUnitNil

`func (o *CloudBaseVolumeAllOf) SetBillingUnitNil(b bool)`

 SetBillingUnitNil sets the value for BillingUnit to be an explicit nil

### UnsetBillingUnit
`func (o *CloudBaseVolumeAllOf) UnsetBillingUnit()`

UnsetBillingUnit ensures that no value is present for BillingUnit, not even an explicit nil
### GetEncryptionState

`func (o *CloudBaseVolumeAllOf) GetEncryptionState() string`

GetEncryptionState returns the EncryptionState field if non-nil, zero value otherwise.

### GetEncryptionStateOk

`func (o *CloudBaseVolumeAllOf) GetEncryptionStateOk() (*string, bool)`

GetEncryptionStateOk returns a tuple with the EncryptionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionState

`func (o *CloudBaseVolumeAllOf) SetEncryptionState(v string)`

SetEncryptionState sets EncryptionState field to given value.

### HasEncryptionState

`func (o *CloudBaseVolumeAllOf) HasEncryptionState() bool`

HasEncryptionState returns a boolean if a field has been set.

### GetIdentity

`func (o *CloudBaseVolumeAllOf) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *CloudBaseVolumeAllOf) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *CloudBaseVolumeAllOf) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *CloudBaseVolumeAllOf) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetInstanceAttachments

`func (o *CloudBaseVolumeAllOf) GetInstanceAttachments() []CloudVolumeInstanceAttachment`

GetInstanceAttachments returns the InstanceAttachments field if non-nil, zero value otherwise.

### GetInstanceAttachmentsOk

`func (o *CloudBaseVolumeAllOf) GetInstanceAttachmentsOk() (*[]CloudVolumeInstanceAttachment, bool)`

GetInstanceAttachmentsOk returns a tuple with the InstanceAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceAttachments

`func (o *CloudBaseVolumeAllOf) SetInstanceAttachments(v []CloudVolumeInstanceAttachment)`

SetInstanceAttachments sets InstanceAttachments field to given value.

### HasInstanceAttachments

`func (o *CloudBaseVolumeAllOf) HasInstanceAttachments() bool`

HasInstanceAttachments returns a boolean if a field has been set.

### SetInstanceAttachmentsNil

`func (o *CloudBaseVolumeAllOf) SetInstanceAttachmentsNil(b bool)`

 SetInstanceAttachmentsNil sets the value for InstanceAttachments to be an explicit nil

### UnsetInstanceAttachments
`func (o *CloudBaseVolumeAllOf) UnsetInstanceAttachments()`

UnsetInstanceAttachments ensures that no value is present for InstanceAttachments, not even an explicit nil
### GetIopsInfo

`func (o *CloudBaseVolumeAllOf) GetIopsInfo() CloudVolumeIopsInfo`

GetIopsInfo returns the IopsInfo field if non-nil, zero value otherwise.

### GetIopsInfoOk

`func (o *CloudBaseVolumeAllOf) GetIopsInfoOk() (*CloudVolumeIopsInfo, bool)`

GetIopsInfoOk returns a tuple with the IopsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsInfo

`func (o *CloudBaseVolumeAllOf) SetIopsInfo(v CloudVolumeIopsInfo)`

SetIopsInfo sets IopsInfo field to given value.

### HasIopsInfo

`func (o *CloudBaseVolumeAllOf) HasIopsInfo() bool`

HasIopsInfo returns a boolean if a field has been set.

### SetIopsInfoNil

`func (o *CloudBaseVolumeAllOf) SetIopsInfoNil(b bool)`

 SetIopsInfoNil sets the value for IopsInfo to be an explicit nil

### UnsetIopsInfo
`func (o *CloudBaseVolumeAllOf) UnsetIopsInfo()`

UnsetIopsInfo ensures that no value is present for IopsInfo, not even an explicit nil
### GetRegionInfo

`func (o *CloudBaseVolumeAllOf) GetRegionInfo() CloudCloudRegion`

GetRegionInfo returns the RegionInfo field if non-nil, zero value otherwise.

### GetRegionInfoOk

`func (o *CloudBaseVolumeAllOf) GetRegionInfoOk() (*CloudCloudRegion, bool)`

GetRegionInfoOk returns a tuple with the RegionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionInfo

`func (o *CloudBaseVolumeAllOf) SetRegionInfo(v CloudCloudRegion)`

SetRegionInfo sets RegionInfo field to given value.

### HasRegionInfo

`func (o *CloudBaseVolumeAllOf) HasRegionInfo() bool`

HasRegionInfo returns a boolean if a field has been set.

### SetRegionInfoNil

`func (o *CloudBaseVolumeAllOf) SetRegionInfoNil(b bool)`

 SetRegionInfoNil sets the value for RegionInfo to be an explicit nil

### UnsetRegionInfo
`func (o *CloudBaseVolumeAllOf) UnsetRegionInfo()`

UnsetRegionInfo ensures that no value is present for RegionInfo, not even an explicit nil
### GetSourceImageId

`func (o *CloudBaseVolumeAllOf) GetSourceImageId() string`

GetSourceImageId returns the SourceImageId field if non-nil, zero value otherwise.

### GetSourceImageIdOk

`func (o *CloudBaseVolumeAllOf) GetSourceImageIdOk() (*string, bool)`

GetSourceImageIdOk returns a tuple with the SourceImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceImageId

`func (o *CloudBaseVolumeAllOf) SetSourceImageId(v string)`

SetSourceImageId sets SourceImageId field to given value.

### HasSourceImageId

`func (o *CloudBaseVolumeAllOf) HasSourceImageId() bool`

HasSourceImageId returns a boolean if a field has been set.

### GetState

`func (o *CloudBaseVolumeAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CloudBaseVolumeAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CloudBaseVolumeAllOf) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CloudBaseVolumeAllOf) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUuid

`func (o *CloudBaseVolumeAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CloudBaseVolumeAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CloudBaseVolumeAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *CloudBaseVolumeAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVolumeCreateTime

`func (o *CloudBaseVolumeAllOf) GetVolumeCreateTime() time.Time`

GetVolumeCreateTime returns the VolumeCreateTime field if non-nil, zero value otherwise.

### GetVolumeCreateTimeOk

`func (o *CloudBaseVolumeAllOf) GetVolumeCreateTimeOk() (*time.Time, bool)`

GetVolumeCreateTimeOk returns a tuple with the VolumeCreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeCreateTime

`func (o *CloudBaseVolumeAllOf) SetVolumeCreateTime(v time.Time)`

SetVolumeCreateTime sets VolumeCreateTime field to given value.

### HasVolumeCreateTime

`func (o *CloudBaseVolumeAllOf) HasVolumeCreateTime() bool`

HasVolumeCreateTime returns a boolean if a field has been set.

### GetVolumeTags

`func (o *CloudBaseVolumeAllOf) GetVolumeTags() []CloudCloudTag`

GetVolumeTags returns the VolumeTags field if non-nil, zero value otherwise.

### GetVolumeTagsOk

`func (o *CloudBaseVolumeAllOf) GetVolumeTagsOk() (*[]CloudCloudTag, bool)`

GetVolumeTagsOk returns a tuple with the VolumeTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeTags

`func (o *CloudBaseVolumeAllOf) SetVolumeTags(v []CloudCloudTag)`

SetVolumeTags sets VolumeTags field to given value.

### HasVolumeTags

`func (o *CloudBaseVolumeAllOf) HasVolumeTags() bool`

HasVolumeTags returns a boolean if a field has been set.

### SetVolumeTagsNil

`func (o *CloudBaseVolumeAllOf) SetVolumeTagsNil(b bool)`

 SetVolumeTagsNil sets the value for VolumeTags to be an explicit nil

### UnsetVolumeTags
`func (o *CloudBaseVolumeAllOf) UnsetVolumeTags()`

UnsetVolumeTags ensures that no value is present for VolumeTags, not even an explicit nil
### GetVolumeType

`func (o *CloudBaseVolumeAllOf) GetVolumeType() CloudVolumeType`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *CloudBaseVolumeAllOf) GetVolumeTypeOk() (*CloudVolumeType, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *CloudBaseVolumeAllOf) SetVolumeType(v CloudVolumeType)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *CloudBaseVolumeAllOf) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### SetVolumeTypeNil

`func (o *CloudBaseVolumeAllOf) SetVolumeTypeNil(b bool)`

 SetVolumeTypeNil sets the value for VolumeType to be an explicit nil

### UnsetVolumeType
`func (o *CloudBaseVolumeAllOf) UnsetVolumeType()`

UnsetVolumeType ensures that no value is present for VolumeType, not even an explicit nil
### GetZoneInfo

`func (o *CloudBaseVolumeAllOf) GetZoneInfo() CloudAvailabilityZone`

GetZoneInfo returns the ZoneInfo field if non-nil, zero value otherwise.

### GetZoneInfoOk

`func (o *CloudBaseVolumeAllOf) GetZoneInfoOk() (*CloudAvailabilityZone, bool)`

GetZoneInfoOk returns a tuple with the ZoneInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneInfo

`func (o *CloudBaseVolumeAllOf) SetZoneInfo(v CloudAvailabilityZone)`

SetZoneInfo sets ZoneInfo field to given value.

### HasZoneInfo

`func (o *CloudBaseVolumeAllOf) HasZoneInfo() bool`

HasZoneInfo returns a boolean if a field has been set.

### SetZoneInfoNil

`func (o *CloudBaseVolumeAllOf) SetZoneInfoNil(b bool)`

 SetZoneInfoNil sets the value for ZoneInfo to be an explicit nil

### UnsetZoneInfo
`func (o *CloudBaseVolumeAllOf) UnsetZoneInfo()`

UnsetZoneInfo ensures that no value is present for ZoneInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


