# KubernetesDaemonSetAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.DaemonSet"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.DaemonSet"]
**Status** | Pointer to [**NullableKubernetesDaemonSetStatus**](KubernetesDaemonSetStatus.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewKubernetesDaemonSetAllOf

`func NewKubernetesDaemonSetAllOf(classId string, objectType string, ) *KubernetesDaemonSetAllOf`

NewKubernetesDaemonSetAllOf instantiates a new KubernetesDaemonSetAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesDaemonSetAllOfWithDefaults

`func NewKubernetesDaemonSetAllOfWithDefaults() *KubernetesDaemonSetAllOf`

NewKubernetesDaemonSetAllOfWithDefaults instantiates a new KubernetesDaemonSetAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesDaemonSetAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesDaemonSetAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesDaemonSetAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesDaemonSetAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesDaemonSetAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesDaemonSetAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetStatus

`func (o *KubernetesDaemonSetAllOf) GetStatus() KubernetesDaemonSetStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KubernetesDaemonSetAllOf) GetStatusOk() (*KubernetesDaemonSetStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KubernetesDaemonSetAllOf) SetStatus(v KubernetesDaemonSetStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KubernetesDaemonSetAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *KubernetesDaemonSetAllOf) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *KubernetesDaemonSetAllOf) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetRegisteredDevice

`func (o *KubernetesDaemonSetAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *KubernetesDaemonSetAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *KubernetesDaemonSetAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *KubernetesDaemonSetAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


