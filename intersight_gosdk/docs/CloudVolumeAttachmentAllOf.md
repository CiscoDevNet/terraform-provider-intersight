# CloudVolumeAttachmentAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cloud.VolumeAttachment"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cloud.VolumeAttachment"]
**AttachedTime** | Pointer to **time.Time** | The time stamp when the attachment of volume to virtual machine is initiated. | [optional] [readonly] 
**AutoDelete** | Pointer to **bool** | If true, volume is deleted on virtual machine termination. | [optional] [readonly] 
**DetachedTime** | Pointer to **time.Time** | The time stamp when the detached of volume to virtual machine is initiated. | [optional] [readonly] 
**DeviceName** | Pointer to **string** | The device name.For example, /dev/sdh or xvdh in case of AWS cloud. | [optional] [readonly] 
**Identity** | Pointer to **string** | The internally generated identity of this volume. | [optional] 
**Index** | Pointer to **int64** | The position of the volume attachment in the virtual machine. | [optional] [readonly] 
**IsRoot** | Pointer to **bool** | If set to true, then it is the root volume. | [optional] [readonly] 

## Methods

### NewCloudVolumeAttachmentAllOf

`func NewCloudVolumeAttachmentAllOf(classId string, objectType string, ) *CloudVolumeAttachmentAllOf`

NewCloudVolumeAttachmentAllOf instantiates a new CloudVolumeAttachmentAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudVolumeAttachmentAllOfWithDefaults

`func NewCloudVolumeAttachmentAllOfWithDefaults() *CloudVolumeAttachmentAllOf`

NewCloudVolumeAttachmentAllOfWithDefaults instantiates a new CloudVolumeAttachmentAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudVolumeAttachmentAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudVolumeAttachmentAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudVolumeAttachmentAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudVolumeAttachmentAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudVolumeAttachmentAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudVolumeAttachmentAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAttachedTime

`func (o *CloudVolumeAttachmentAllOf) GetAttachedTime() time.Time`

GetAttachedTime returns the AttachedTime field if non-nil, zero value otherwise.

### GetAttachedTimeOk

`func (o *CloudVolumeAttachmentAllOf) GetAttachedTimeOk() (*time.Time, bool)`

GetAttachedTimeOk returns a tuple with the AttachedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedTime

`func (o *CloudVolumeAttachmentAllOf) SetAttachedTime(v time.Time)`

SetAttachedTime sets AttachedTime field to given value.

### HasAttachedTime

`func (o *CloudVolumeAttachmentAllOf) HasAttachedTime() bool`

HasAttachedTime returns a boolean if a field has been set.

### GetAutoDelete

`func (o *CloudVolumeAttachmentAllOf) GetAutoDelete() bool`

GetAutoDelete returns the AutoDelete field if non-nil, zero value otherwise.

### GetAutoDeleteOk

`func (o *CloudVolumeAttachmentAllOf) GetAutoDeleteOk() (*bool, bool)`

GetAutoDeleteOk returns a tuple with the AutoDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDelete

`func (o *CloudVolumeAttachmentAllOf) SetAutoDelete(v bool)`

SetAutoDelete sets AutoDelete field to given value.

### HasAutoDelete

`func (o *CloudVolumeAttachmentAllOf) HasAutoDelete() bool`

HasAutoDelete returns a boolean if a field has been set.

### GetDetachedTime

`func (o *CloudVolumeAttachmentAllOf) GetDetachedTime() time.Time`

GetDetachedTime returns the DetachedTime field if non-nil, zero value otherwise.

### GetDetachedTimeOk

`func (o *CloudVolumeAttachmentAllOf) GetDetachedTimeOk() (*time.Time, bool)`

GetDetachedTimeOk returns a tuple with the DetachedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetachedTime

`func (o *CloudVolumeAttachmentAllOf) SetDetachedTime(v time.Time)`

SetDetachedTime sets DetachedTime field to given value.

### HasDetachedTime

`func (o *CloudVolumeAttachmentAllOf) HasDetachedTime() bool`

HasDetachedTime returns a boolean if a field has been set.

### GetDeviceName

`func (o *CloudVolumeAttachmentAllOf) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *CloudVolumeAttachmentAllOf) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *CloudVolumeAttachmentAllOf) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *CloudVolumeAttachmentAllOf) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetIdentity

`func (o *CloudVolumeAttachmentAllOf) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *CloudVolumeAttachmentAllOf) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *CloudVolumeAttachmentAllOf) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *CloudVolumeAttachmentAllOf) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetIndex

`func (o *CloudVolumeAttachmentAllOf) GetIndex() int64`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *CloudVolumeAttachmentAllOf) GetIndexOk() (*int64, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *CloudVolumeAttachmentAllOf) SetIndex(v int64)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *CloudVolumeAttachmentAllOf) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetIsRoot

`func (o *CloudVolumeAttachmentAllOf) GetIsRoot() bool`

GetIsRoot returns the IsRoot field if non-nil, zero value otherwise.

### GetIsRootOk

`func (o *CloudVolumeAttachmentAllOf) GetIsRootOk() (*bool, bool)`

GetIsRootOk returns a tuple with the IsRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRoot

`func (o *CloudVolumeAttachmentAllOf) SetIsRoot(v bool)`

SetIsRoot sets IsRoot field to given value.

### HasIsRoot

`func (o *CloudVolumeAttachmentAllOf) HasIsRoot() bool`

HasIsRoot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


