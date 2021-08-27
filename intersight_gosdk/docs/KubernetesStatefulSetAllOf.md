# KubernetesStatefulSetAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.StatefulSet"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.StatefulSet"]
**Status** | Pointer to [**NullableKubernetesStatefulSetStatus**](KubernetesStatefulSetStatus.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewKubernetesStatefulSetAllOf

`func NewKubernetesStatefulSetAllOf(classId string, objectType string, ) *KubernetesStatefulSetAllOf`

NewKubernetesStatefulSetAllOf instantiates a new KubernetesStatefulSetAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesStatefulSetAllOfWithDefaults

`func NewKubernetesStatefulSetAllOfWithDefaults() *KubernetesStatefulSetAllOf`

NewKubernetesStatefulSetAllOfWithDefaults instantiates a new KubernetesStatefulSetAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesStatefulSetAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesStatefulSetAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesStatefulSetAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesStatefulSetAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesStatefulSetAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesStatefulSetAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetStatus

`func (o *KubernetesStatefulSetAllOf) GetStatus() KubernetesStatefulSetStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KubernetesStatefulSetAllOf) GetStatusOk() (*KubernetesStatefulSetStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KubernetesStatefulSetAllOf) SetStatus(v KubernetesStatefulSetStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KubernetesStatefulSetAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *KubernetesStatefulSetAllOf) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *KubernetesStatefulSetAllOf) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetRegisteredDevice

`func (o *KubernetesStatefulSetAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *KubernetesStatefulSetAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *KubernetesStatefulSetAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *KubernetesStatefulSetAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


