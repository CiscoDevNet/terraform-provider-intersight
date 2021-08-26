# KubernetesDeploymentAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.Deployment"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.Deployment"]
**Status** | Pointer to [**NullableKubernetesDeploymentStatus**](KubernetesDeploymentStatus.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewKubernetesDeploymentAllOf

`func NewKubernetesDeploymentAllOf(classId string, objectType string, ) *KubernetesDeploymentAllOf`

NewKubernetesDeploymentAllOf instantiates a new KubernetesDeploymentAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesDeploymentAllOfWithDefaults

`func NewKubernetesDeploymentAllOfWithDefaults() *KubernetesDeploymentAllOf`

NewKubernetesDeploymentAllOfWithDefaults instantiates a new KubernetesDeploymentAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesDeploymentAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesDeploymentAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesDeploymentAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesDeploymentAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesDeploymentAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesDeploymentAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetStatus

`func (o *KubernetesDeploymentAllOf) GetStatus() KubernetesDeploymentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KubernetesDeploymentAllOf) GetStatusOk() (*KubernetesDeploymentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KubernetesDeploymentAllOf) SetStatus(v KubernetesDeploymentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KubernetesDeploymentAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *KubernetesDeploymentAllOf) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *KubernetesDeploymentAllOf) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetRegisteredDevice

`func (o *KubernetesDeploymentAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *KubernetesDeploymentAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *KubernetesDeploymentAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *KubernetesDeploymentAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


