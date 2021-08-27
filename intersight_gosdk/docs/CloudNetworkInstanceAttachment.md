# CloudNetworkInstanceAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cloud.NetworkInstanceAttachment"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cloud.NetworkInstanceAttachment"]
**AttachTime** | Pointer to **time.Time** | Time stamp when the attachment was initiated. | [optional] [readonly] 
**AutoDelete** | Pointer to **bool** | Indicates whether the attachment is deleted when an instance is terminated. | [optional] [readonly] 
**DeviceIndex** | Pointer to **int64** | The device index to which the network interface is attached. | [optional] [readonly] 
**InstanceId** | Pointer to **string** | The ID of the instance to which the network interface is attached. | [optional] [readonly] 
**State** | Pointer to **string** | The status of the attachment. It is one of attaching, attached, detaching, or detached. * &#x60;UnAttached&#x60; - Network interface is not attached to a virtual machine. * &#x60;Attached&#x60; - Network interface is attached to a virtual machine. * &#x60;Attaching&#x60; - Network interface is being attached to a virtual machine. * &#x60;Detaching&#x60; - Network interface is being attached to a virtual machine. * &#x60;Detached&#x60; - Network interface is detached from a virtual machine. | [optional] [readonly] [default to "UnAttached"]

## Methods

### NewCloudNetworkInstanceAttachment

`func NewCloudNetworkInstanceAttachment(classId string, objectType string, ) *CloudNetworkInstanceAttachment`

NewCloudNetworkInstanceAttachment instantiates a new CloudNetworkInstanceAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudNetworkInstanceAttachmentWithDefaults

`func NewCloudNetworkInstanceAttachmentWithDefaults() *CloudNetworkInstanceAttachment`

NewCloudNetworkInstanceAttachmentWithDefaults instantiates a new CloudNetworkInstanceAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudNetworkInstanceAttachment) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudNetworkInstanceAttachment) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudNetworkInstanceAttachment) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudNetworkInstanceAttachment) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudNetworkInstanceAttachment) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudNetworkInstanceAttachment) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAttachTime

`func (o *CloudNetworkInstanceAttachment) GetAttachTime() time.Time`

GetAttachTime returns the AttachTime field if non-nil, zero value otherwise.

### GetAttachTimeOk

`func (o *CloudNetworkInstanceAttachment) GetAttachTimeOk() (*time.Time, bool)`

GetAttachTimeOk returns a tuple with the AttachTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachTime

`func (o *CloudNetworkInstanceAttachment) SetAttachTime(v time.Time)`

SetAttachTime sets AttachTime field to given value.

### HasAttachTime

`func (o *CloudNetworkInstanceAttachment) HasAttachTime() bool`

HasAttachTime returns a boolean if a field has been set.

### GetAutoDelete

`func (o *CloudNetworkInstanceAttachment) GetAutoDelete() bool`

GetAutoDelete returns the AutoDelete field if non-nil, zero value otherwise.

### GetAutoDeleteOk

`func (o *CloudNetworkInstanceAttachment) GetAutoDeleteOk() (*bool, bool)`

GetAutoDeleteOk returns a tuple with the AutoDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDelete

`func (o *CloudNetworkInstanceAttachment) SetAutoDelete(v bool)`

SetAutoDelete sets AutoDelete field to given value.

### HasAutoDelete

`func (o *CloudNetworkInstanceAttachment) HasAutoDelete() bool`

HasAutoDelete returns a boolean if a field has been set.

### GetDeviceIndex

`func (o *CloudNetworkInstanceAttachment) GetDeviceIndex() int64`

GetDeviceIndex returns the DeviceIndex field if non-nil, zero value otherwise.

### GetDeviceIndexOk

`func (o *CloudNetworkInstanceAttachment) GetDeviceIndexOk() (*int64, bool)`

GetDeviceIndexOk returns a tuple with the DeviceIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIndex

`func (o *CloudNetworkInstanceAttachment) SetDeviceIndex(v int64)`

SetDeviceIndex sets DeviceIndex field to given value.

### HasDeviceIndex

`func (o *CloudNetworkInstanceAttachment) HasDeviceIndex() bool`

HasDeviceIndex returns a boolean if a field has been set.

### GetInstanceId

`func (o *CloudNetworkInstanceAttachment) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *CloudNetworkInstanceAttachment) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *CloudNetworkInstanceAttachment) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *CloudNetworkInstanceAttachment) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetState

`func (o *CloudNetworkInstanceAttachment) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CloudNetworkInstanceAttachment) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CloudNetworkInstanceAttachment) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CloudNetworkInstanceAttachment) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


