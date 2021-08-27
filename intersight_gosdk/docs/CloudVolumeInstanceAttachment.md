# CloudVolumeInstanceAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cloud.VolumeInstanceAttachment"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cloud.VolumeInstanceAttachment"]
**AttachTime** | Pointer to **time.Time** | Time stamp when the attachment initiated. | [optional] [readonly] 
**AutoDelete** | Pointer to **bool** | If true, volume is deleted on virtual machine termination. | [optional] [readonly] 
**DeviceName** | Pointer to **string** | Device name assigned to the volume. | [optional] [readonly] 
**InstanceId** | Pointer to **string** | ID of the virtual machine, the volume is attached to. | [optional] [readonly] 
**State** | Pointer to **string** | The attachment state of the volume. * &#x60;UnRecognized&#x60; - Volume is in unrecognized state. * &#x60;Attached&#x60; - Volume is attached to the virtual machine. * &#x60;Attaching&#x60; - Volume is being attached to the virtual machine. * &#x60;Detaching&#x60; - Volume is being detached from the virtual machine. | [optional] [readonly] [default to "UnRecognized"]

## Methods

### NewCloudVolumeInstanceAttachment

`func NewCloudVolumeInstanceAttachment(classId string, objectType string, ) *CloudVolumeInstanceAttachment`

NewCloudVolumeInstanceAttachment instantiates a new CloudVolumeInstanceAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudVolumeInstanceAttachmentWithDefaults

`func NewCloudVolumeInstanceAttachmentWithDefaults() *CloudVolumeInstanceAttachment`

NewCloudVolumeInstanceAttachmentWithDefaults instantiates a new CloudVolumeInstanceAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudVolumeInstanceAttachment) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudVolumeInstanceAttachment) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudVolumeInstanceAttachment) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudVolumeInstanceAttachment) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudVolumeInstanceAttachment) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudVolumeInstanceAttachment) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAttachTime

`func (o *CloudVolumeInstanceAttachment) GetAttachTime() time.Time`

GetAttachTime returns the AttachTime field if non-nil, zero value otherwise.

### GetAttachTimeOk

`func (o *CloudVolumeInstanceAttachment) GetAttachTimeOk() (*time.Time, bool)`

GetAttachTimeOk returns a tuple with the AttachTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachTime

`func (o *CloudVolumeInstanceAttachment) SetAttachTime(v time.Time)`

SetAttachTime sets AttachTime field to given value.

### HasAttachTime

`func (o *CloudVolumeInstanceAttachment) HasAttachTime() bool`

HasAttachTime returns a boolean if a field has been set.

### GetAutoDelete

`func (o *CloudVolumeInstanceAttachment) GetAutoDelete() bool`

GetAutoDelete returns the AutoDelete field if non-nil, zero value otherwise.

### GetAutoDeleteOk

`func (o *CloudVolumeInstanceAttachment) GetAutoDeleteOk() (*bool, bool)`

GetAutoDeleteOk returns a tuple with the AutoDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDelete

`func (o *CloudVolumeInstanceAttachment) SetAutoDelete(v bool)`

SetAutoDelete sets AutoDelete field to given value.

### HasAutoDelete

`func (o *CloudVolumeInstanceAttachment) HasAutoDelete() bool`

HasAutoDelete returns a boolean if a field has been set.

### GetDeviceName

`func (o *CloudVolumeInstanceAttachment) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *CloudVolumeInstanceAttachment) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *CloudVolumeInstanceAttachment) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *CloudVolumeInstanceAttachment) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetInstanceId

`func (o *CloudVolumeInstanceAttachment) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *CloudVolumeInstanceAttachment) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *CloudVolumeInstanceAttachment) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *CloudVolumeInstanceAttachment) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetState

`func (o *CloudVolumeInstanceAttachment) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CloudVolumeInstanceAttachment) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CloudVolumeInstanceAttachment) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CloudVolumeInstanceAttachment) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


