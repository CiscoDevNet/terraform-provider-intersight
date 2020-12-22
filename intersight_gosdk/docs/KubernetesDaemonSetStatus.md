# KubernetesDaemonSetStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "kubernetes.DaemonSetStatus"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "kubernetes.DaemonSetStatus"]
**CurrentNumberScheduled** | Pointer to **int64** | The number of nodes that are running at least 1 daemon pod and are supposed to run the daemon pod. | [optional] 
**DesiredNumberScheduled** | Pointer to **int64** | The total number of nodes that should be running the daemon pod (including nodes correctly running the daemon pod). | [optional] 
**NumberAvailable** | Pointer to **string** | The number of nodes that should be running the daemon pod and have one or more of the daemon pod running and available (ready for at least spec.minReadySeconds). | [optional] 
**NumberMisscheduled** | Pointer to **int64** | The number of nodes that are running the daemon pod, but are not supposed to run the daemon pod. | [optional] 
**NumberReady** | Pointer to **int64** | The number of nodes that should be running the daemon pod and have one or more of the daemon pod running and ready. | [optional] 
**ObservedGeneration** | Pointer to **int64** | The most recent generation observed by the daemon set controller. | [optional] 
**UpdatedNumberScheduled** | Pointer to **string** | The total number of nodes that are running updated daemon pod. | [optional] 

## Methods

### NewKubernetesDaemonSetStatus

`func NewKubernetesDaemonSetStatus(classId string, objectType string, ) *KubernetesDaemonSetStatus`

NewKubernetesDaemonSetStatus instantiates a new KubernetesDaemonSetStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesDaemonSetStatusWithDefaults

`func NewKubernetesDaemonSetStatusWithDefaults() *KubernetesDaemonSetStatus`

NewKubernetesDaemonSetStatusWithDefaults instantiates a new KubernetesDaemonSetStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesDaemonSetStatus) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesDaemonSetStatus) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesDaemonSetStatus) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesDaemonSetStatus) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesDaemonSetStatus) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesDaemonSetStatus) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCurrentNumberScheduled

`func (o *KubernetesDaemonSetStatus) GetCurrentNumberScheduled() int64`

GetCurrentNumberScheduled returns the CurrentNumberScheduled field if non-nil, zero value otherwise.

### GetCurrentNumberScheduledOk

`func (o *KubernetesDaemonSetStatus) GetCurrentNumberScheduledOk() (*int64, bool)`

GetCurrentNumberScheduledOk returns a tuple with the CurrentNumberScheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentNumberScheduled

`func (o *KubernetesDaemonSetStatus) SetCurrentNumberScheduled(v int64)`

SetCurrentNumberScheduled sets CurrentNumberScheduled field to given value.

### HasCurrentNumberScheduled

`func (o *KubernetesDaemonSetStatus) HasCurrentNumberScheduled() bool`

HasCurrentNumberScheduled returns a boolean if a field has been set.

### GetDesiredNumberScheduled

`func (o *KubernetesDaemonSetStatus) GetDesiredNumberScheduled() int64`

GetDesiredNumberScheduled returns the DesiredNumberScheduled field if non-nil, zero value otherwise.

### GetDesiredNumberScheduledOk

`func (o *KubernetesDaemonSetStatus) GetDesiredNumberScheduledOk() (*int64, bool)`

GetDesiredNumberScheduledOk returns a tuple with the DesiredNumberScheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredNumberScheduled

`func (o *KubernetesDaemonSetStatus) SetDesiredNumberScheduled(v int64)`

SetDesiredNumberScheduled sets DesiredNumberScheduled field to given value.

### HasDesiredNumberScheduled

`func (o *KubernetesDaemonSetStatus) HasDesiredNumberScheduled() bool`

HasDesiredNumberScheduled returns a boolean if a field has been set.

### GetNumberAvailable

`func (o *KubernetesDaemonSetStatus) GetNumberAvailable() string`

GetNumberAvailable returns the NumberAvailable field if non-nil, zero value otherwise.

### GetNumberAvailableOk

`func (o *KubernetesDaemonSetStatus) GetNumberAvailableOk() (*string, bool)`

GetNumberAvailableOk returns a tuple with the NumberAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberAvailable

`func (o *KubernetesDaemonSetStatus) SetNumberAvailable(v string)`

SetNumberAvailable sets NumberAvailable field to given value.

### HasNumberAvailable

`func (o *KubernetesDaemonSetStatus) HasNumberAvailable() bool`

HasNumberAvailable returns a boolean if a field has been set.

### GetNumberMisscheduled

`func (o *KubernetesDaemonSetStatus) GetNumberMisscheduled() int64`

GetNumberMisscheduled returns the NumberMisscheduled field if non-nil, zero value otherwise.

### GetNumberMisscheduledOk

`func (o *KubernetesDaemonSetStatus) GetNumberMisscheduledOk() (*int64, bool)`

GetNumberMisscheduledOk returns a tuple with the NumberMisscheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberMisscheduled

`func (o *KubernetesDaemonSetStatus) SetNumberMisscheduled(v int64)`

SetNumberMisscheduled sets NumberMisscheduled field to given value.

### HasNumberMisscheduled

`func (o *KubernetesDaemonSetStatus) HasNumberMisscheduled() bool`

HasNumberMisscheduled returns a boolean if a field has been set.

### GetNumberReady

`func (o *KubernetesDaemonSetStatus) GetNumberReady() int64`

GetNumberReady returns the NumberReady field if non-nil, zero value otherwise.

### GetNumberReadyOk

`func (o *KubernetesDaemonSetStatus) GetNumberReadyOk() (*int64, bool)`

GetNumberReadyOk returns a tuple with the NumberReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberReady

`func (o *KubernetesDaemonSetStatus) SetNumberReady(v int64)`

SetNumberReady sets NumberReady field to given value.

### HasNumberReady

`func (o *KubernetesDaemonSetStatus) HasNumberReady() bool`

HasNumberReady returns a boolean if a field has been set.

### GetObservedGeneration

`func (o *KubernetesDaemonSetStatus) GetObservedGeneration() int64`

GetObservedGeneration returns the ObservedGeneration field if non-nil, zero value otherwise.

### GetObservedGenerationOk

`func (o *KubernetesDaemonSetStatus) GetObservedGenerationOk() (*int64, bool)`

GetObservedGenerationOk returns a tuple with the ObservedGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedGeneration

`func (o *KubernetesDaemonSetStatus) SetObservedGeneration(v int64)`

SetObservedGeneration sets ObservedGeneration field to given value.

### HasObservedGeneration

`func (o *KubernetesDaemonSetStatus) HasObservedGeneration() bool`

HasObservedGeneration returns a boolean if a field has been set.

### GetUpdatedNumberScheduled

`func (o *KubernetesDaemonSetStatus) GetUpdatedNumberScheduled() string`

GetUpdatedNumberScheduled returns the UpdatedNumberScheduled field if non-nil, zero value otherwise.

### GetUpdatedNumberScheduledOk

`func (o *KubernetesDaemonSetStatus) GetUpdatedNumberScheduledOk() (*string, bool)`

GetUpdatedNumberScheduledOk returns a tuple with the UpdatedNumberScheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedNumberScheduled

`func (o *KubernetesDaemonSetStatus) SetUpdatedNumberScheduled(v string)`

SetUpdatedNumberScheduled sets UpdatedNumberScheduled field to given value.

### HasUpdatedNumberScheduled

`func (o *KubernetesDaemonSetStatus) HasUpdatedNumberScheduled() bool`

HasUpdatedNumberScheduled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


