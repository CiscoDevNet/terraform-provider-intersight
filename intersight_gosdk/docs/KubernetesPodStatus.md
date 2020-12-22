# KubernetesPodStatus

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

### NewKubernetesPodStatus

`func NewKubernetesPodStatus(classId string, objectType string, ) *KubernetesPodStatus`

NewKubernetesPodStatus instantiates a new KubernetesPodStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesPodStatusWithDefaults

`func NewKubernetesPodStatusWithDefaults() *KubernetesPodStatus`

NewKubernetesPodStatusWithDefaults instantiates a new KubernetesPodStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesPodStatus) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesPodStatus) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesPodStatus) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesPodStatus) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesPodStatus) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesPodStatus) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetHostIp

`func (o *KubernetesPodStatus) GetHostIp() string`

GetHostIp returns the HostIp field if non-nil, zero value otherwise.

### GetHostIpOk

`func (o *KubernetesPodStatus) GetHostIpOk() (*string, bool)`

GetHostIpOk returns a tuple with the HostIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostIp

`func (o *KubernetesPodStatus) SetHostIp(v string)`

SetHostIp sets HostIp field to given value.

### HasHostIp

`func (o *KubernetesPodStatus) HasHostIp() bool`

HasHostIp returns a boolean if a field has been set.

### GetPhase

`func (o *KubernetesPodStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *KubernetesPodStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *KubernetesPodStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *KubernetesPodStatus) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetPodIp

`func (o *KubernetesPodStatus) GetPodIp() string`

GetPodIp returns the PodIp field if non-nil, zero value otherwise.

### GetPodIpOk

`func (o *KubernetesPodStatus) GetPodIpOk() (*string, bool)`

GetPodIpOk returns a tuple with the PodIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodIp

`func (o *KubernetesPodStatus) SetPodIp(v string)`

SetPodIp sets PodIp field to given value.

### HasPodIp

`func (o *KubernetesPodStatus) HasPodIp() bool`

HasPodIp returns a boolean if a field has been set.

### GetQosClass

`func (o *KubernetesPodStatus) GetQosClass() string`

GetQosClass returns the QosClass field if non-nil, zero value otherwise.

### GetQosClassOk

`func (o *KubernetesPodStatus) GetQosClassOk() (*string, bool)`

GetQosClassOk returns a tuple with the QosClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosClass

`func (o *KubernetesPodStatus) SetQosClass(v string)`

SetQosClass sets QosClass field to given value.

### HasQosClass

`func (o *KubernetesPodStatus) HasQosClass() bool`

HasQosClass returns a boolean if a field has been set.

### GetStartTime

`func (o *KubernetesPodStatus) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *KubernetesPodStatus) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *KubernetesPodStatus) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *KubernetesPodStatus) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


