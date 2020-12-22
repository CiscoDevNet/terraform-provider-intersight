# KubernetesDaemonSetStatusAllOf

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

### NewKubernetesDaemonSetStatusAllOf

`func NewKubernetesDaemonSetStatusAllOf(classId string, objectType string, ) *KubernetesDaemonSetStatusAllOf`

NewKubernetesDaemonSetStatusAllOf instantiates a new KubernetesDaemonSetStatusAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesDaemonSetStatusAllOfWithDefaults

`func NewKubernetesDaemonSetStatusAllOfWithDefaults() *KubernetesDaemonSetStatusAllOf`

NewKubernetesDaemonSetStatusAllOfWithDefaults instantiates a new KubernetesDaemonSetStatusAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *KubernetesDaemonSetStatusAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *KubernetesDaemonSetStatusAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *KubernetesDaemonSetStatusAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *KubernetesDaemonSetStatusAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KubernetesDaemonSetStatusAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KubernetesDaemonSetStatusAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCurrentNumberScheduled

`func (o *KubernetesDaemonSetStatusAllOf) GetCurrentNumberScheduled() int64`

GetCurrentNumberScheduled returns the CurrentNumberScheduled field if non-nil, zero value otherwise.

### GetCurrentNumberScheduledOk

`func (o *KubernetesDaemonSetStatusAllOf) GetCurrentNumberScheduledOk() (*int64, bool)`

GetCurrentNumberScheduledOk returns a tuple with the CurrentNumberScheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentNumberScheduled

`func (o *KubernetesDaemonSetStatusAllOf) SetCurrentNumberScheduled(v int64)`

SetCurrentNumberScheduled sets CurrentNumberScheduled field to given value.

### HasCurrentNumberScheduled

`func (o *KubernetesDaemonSetStatusAllOf) HasCurrentNumberScheduled() bool`

HasCurrentNumberScheduled returns a boolean if a field has been set.

### GetDesiredNumberScheduled

`func (o *KubernetesDaemonSetStatusAllOf) GetDesiredNumberScheduled() int64`

GetDesiredNumberScheduled returns the DesiredNumberScheduled field if non-nil, zero value otherwise.

### GetDesiredNumberScheduledOk

`func (o *KubernetesDaemonSetStatusAllOf) GetDesiredNumberScheduledOk() (*int64, bool)`

GetDesiredNumberScheduledOk returns a tuple with the DesiredNumberScheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredNumberScheduled

`func (o *KubernetesDaemonSetStatusAllOf) SetDesiredNumberScheduled(v int64)`

SetDesiredNumberScheduled sets DesiredNumberScheduled field to given value.

### HasDesiredNumberScheduled

`func (o *KubernetesDaemonSetStatusAllOf) HasDesiredNumberScheduled() bool`

HasDesiredNumberScheduled returns a boolean if a field has been set.

### GetNumberAvailable

`func (o *KubernetesDaemonSetStatusAllOf) GetNumberAvailable() string`

GetNumberAvailable returns the NumberAvailable field if non-nil, zero value otherwise.

### GetNumberAvailableOk

`func (o *KubernetesDaemonSetStatusAllOf) GetNumberAvailableOk() (*string, bool)`

GetNumberAvailableOk returns a tuple with the NumberAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberAvailable

`func (o *KubernetesDaemonSetStatusAllOf) SetNumberAvailable(v string)`

SetNumberAvailable sets NumberAvailable field to given value.

### HasNumberAvailable

`func (o *KubernetesDaemonSetStatusAllOf) HasNumberAvailable() bool`

HasNumberAvailable returns a boolean if a field has been set.

### GetNumberMisscheduled

`func (o *KubernetesDaemonSetStatusAllOf) GetNumberMisscheduled() int64`

GetNumberMisscheduled returns the NumberMisscheduled field if non-nil, zero value otherwise.

### GetNumberMisscheduledOk

`func (o *KubernetesDaemonSetStatusAllOf) GetNumberMisscheduledOk() (*int64, bool)`

GetNumberMisscheduledOk returns a tuple with the NumberMisscheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberMisscheduled

`func (o *KubernetesDaemonSetStatusAllOf) SetNumberMisscheduled(v int64)`

SetNumberMisscheduled sets NumberMisscheduled field to given value.

### HasNumberMisscheduled

`func (o *KubernetesDaemonSetStatusAllOf) HasNumberMisscheduled() bool`

HasNumberMisscheduled returns a boolean if a field has been set.

### GetNumberReady

`func (o *KubernetesDaemonSetStatusAllOf) GetNumberReady() int64`

GetNumberReady returns the NumberReady field if non-nil, zero value otherwise.

### GetNumberReadyOk

`func (o *KubernetesDaemonSetStatusAllOf) GetNumberReadyOk() (*int64, bool)`

GetNumberReadyOk returns a tuple with the NumberReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberReady

`func (o *KubernetesDaemonSetStatusAllOf) SetNumberReady(v int64)`

SetNumberReady sets NumberReady field to given value.

### HasNumberReady

`func (o *KubernetesDaemonSetStatusAllOf) HasNumberReady() bool`

HasNumberReady returns a boolean if a field has been set.

### GetObservedGeneration

`func (o *KubernetesDaemonSetStatusAllOf) GetObservedGeneration() int64`

GetObservedGeneration returns the ObservedGeneration field if non-nil, zero value otherwise.

### GetObservedGenerationOk

`func (o *KubernetesDaemonSetStatusAllOf) GetObservedGenerationOk() (*int64, bool)`

GetObservedGenerationOk returns a tuple with the ObservedGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedGeneration

`func (o *KubernetesDaemonSetStatusAllOf) SetObservedGeneration(v int64)`

SetObservedGeneration sets ObservedGeneration field to given value.

### HasObservedGeneration

`func (o *KubernetesDaemonSetStatusAllOf) HasObservedGeneration() bool`

HasObservedGeneration returns a boolean if a field has been set.

### GetUpdatedNumberScheduled

`func (o *KubernetesDaemonSetStatusAllOf) GetUpdatedNumberScheduled() string`

GetUpdatedNumberScheduled returns the UpdatedNumberScheduled field if non-nil, zero value otherwise.

### GetUpdatedNumberScheduledOk

`func (o *KubernetesDaemonSetStatusAllOf) GetUpdatedNumberScheduledOk() (*string, bool)`

GetUpdatedNumberScheduledOk returns a tuple with the UpdatedNumberScheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedNumberScheduled

`func (o *KubernetesDaemonSetStatusAllOf) SetUpdatedNumberScheduled(v string)`

SetUpdatedNumberScheduled sets UpdatedNumberScheduled field to given value.

### HasUpdatedNumberScheduled

`func (o *KubernetesDaemonSetStatusAllOf) HasUpdatedNumberScheduled() bool`

HasUpdatedNumberScheduled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


