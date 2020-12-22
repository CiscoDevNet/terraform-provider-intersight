# KubernetesPodStatusAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.PodStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.PodStatus"]
**HostIp** | Pointer to **string** | The IP of the host that the pod is running on. | [optional] 
**Phase** | Pointer to **string** | The phase of a Pod is a simple, high-level summary of where the Pod is in its lifecycle. | [optional] 
**PodIp** | Pointer to **string** | The IP of the pod. The IP is allocated by the Pod CIDR from the kubernetes cluster itself. | [optional] 
**QosClass** | Pointer to **string** | The Quality of Service (QOS) classification assigned to the pod based on resource requirements. | [optional] 
**StartTime** | Pointer to **string** | The time that the pod was started. | [optional] 

## Methods

### NewKubernetesPodStatusAllOf

`func NewKubernetesPodStatusAllOf(classId string, objectType string, ) *KubernetesPodStatusAllOf`

NewKubernetesPodStatusAllOf instantiates a new KubernetesPodStatusAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesPodStatusAllOfWithDefaults

`func NewKubernetesPodStatusAllOfWithDefaults() *KubernetesPodStatusAllOf`

NewKubernetesPodStatusAllOfWithDefaults instantiates a new KubernetesPodStatusAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesPodStatusAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesPodStatusAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesPodStatusAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesPodStatusAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesPodStatusAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesPodStatusAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetHostIp

`func (o *KubernetesPodStatusAllOf) GetHostIp() string`

GetHostIp returns the HostIp field if non-nil, zero value otherwise.

### GetHostIpOk

`func (o *KubernetesPodStatusAllOf) GetHostIpOk() (*string, bool)`

GetHostIpOk returns a tuple with the HostIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostIp

`func (o *KubernetesPodStatusAllOf) SetHostIp(v string)`

SetHostIp sets HostIp field to given value.

### HasHostIp

`func (o *KubernetesPodStatusAllOf) HasHostIp() bool`

HasHostIp returns a boolean if a field has been set.

### GetPhase

`func (o *KubernetesPodStatusAllOf) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *KubernetesPodStatusAllOf) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *KubernetesPodStatusAllOf) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *KubernetesPodStatusAllOf) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetPodIp

`func (o *KubernetesPodStatusAllOf) GetPodIp() string`

GetPodIp returns the PodIp field if non-nil, zero value otherwise.

### GetPodIpOk

`func (o *KubernetesPodStatusAllOf) GetPodIpOk() (*string, bool)`

GetPodIpOk returns a tuple with the PodIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodIp

`func (o *KubernetesPodStatusAllOf) SetPodIp(v string)`

SetPodIp sets PodIp field to given value.

### HasPodIp

`func (o *KubernetesPodStatusAllOf) HasPodIp() bool`

HasPodIp returns a boolean if a field has been set.

### GetQosClass

`func (o *KubernetesPodStatusAllOf) GetQosClass() string`

GetQosClass returns the QosClass field if non-nil, zero value otherwise.

### GetQosClassOk

`func (o *KubernetesPodStatusAllOf) GetQosClassOk() (*string, bool)`

GetQosClassOk returns a tuple with the QosClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosClass

`func (o *KubernetesPodStatusAllOf) SetQosClass(v string)`

SetQosClass sets QosClass field to given value.

### HasQosClass

`func (o *KubernetesPodStatusAllOf) HasQosClass() bool`

HasQosClass returns a boolean if a field has been set.

### GetStartTime

`func (o *KubernetesPodStatusAllOf) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *KubernetesPodStatusAllOf) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *KubernetesPodStatusAllOf) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *KubernetesPodStatusAllOf) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


