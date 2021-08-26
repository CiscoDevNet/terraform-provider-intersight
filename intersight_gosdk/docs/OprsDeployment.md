# OprsDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "oprs.Deployment"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "oprs.Deployment"]
**AvailableReplicas** | Pointer to **int64** | Available number of replicas. | [optional] 
**DesiredReplicas** | Pointer to **int64** | The expected number of replicas. | [optional] 
**Event** | Pointer to **string** | The type of event which was triggered. | [optional] 
**Labels** | Pointer to [**[]OprsKvpair**](OprsKvpair.md) |  | [optional] 
**Name** | Pointer to **string** | Agent name for which the event is triggered. | [optional] 
**Namespace** | Pointer to **string** | Name space in which the agents are running. | [optional] 
**Status** | Pointer to **string** | Status which shows if the resource is healthy or not. * &#x60;&#x60; - An Unknown status indicates that the resource status is not known. * &#x60;Healthy&#x60; - A healthy status indicates that the resource is healthy and running as per spec. * &#x60;Unhealthy&#x60; - An unhealthy status indicates that the resource is down. | [optional] [default to ""]
**TimeStamp** | Pointer to **time.Time** | The time at which the event was generated. Date is accurate to Intersights clock. This time will be used to identify order of events. | [optional] 
**UnavailableReplicas** | Pointer to **int64** | Number of replicas Unavailable. | [optional] 
**Assist** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewOprsDeployment

`func NewOprsDeployment(classId string, objectType string, ) *OprsDeployment`

NewOprsDeployment instantiates a new OprsDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOprsDeploymentWithDefaults

`func NewOprsDeploymentWithDefaults() *OprsDeployment`

NewOprsDeploymentWithDefaults instantiates a new OprsDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *OprsDeployment) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OprsDeployment) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OprsDeployment) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *OprsDeployment) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OprsDeployment) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OprsDeployment) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAvailableReplicas

`func (o *OprsDeployment) GetAvailableReplicas() int64`

GetAvailableReplicas returns the AvailableReplicas field if non-nil, zero value otherwise.

### GetAvailableReplicasOk

`func (o *OprsDeployment) GetAvailableReplicasOk() (*int64, bool)`

GetAvailableReplicasOk returns a tuple with the AvailableReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableReplicas

`func (o *OprsDeployment) SetAvailableReplicas(v int64)`

SetAvailableReplicas sets AvailableReplicas field to given value.

### HasAvailableReplicas

`func (o *OprsDeployment) HasAvailableReplicas() bool`

HasAvailableReplicas returns a boolean if a field has been set.

### GetDesiredReplicas

`func (o *OprsDeployment) GetDesiredReplicas() int64`

GetDesiredReplicas returns the DesiredReplicas field if non-nil, zero value otherwise.

### GetDesiredReplicasOk

`func (o *OprsDeployment) GetDesiredReplicasOk() (*int64, bool)`

GetDesiredReplicasOk returns a tuple with the DesiredReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredReplicas

`func (o *OprsDeployment) SetDesiredReplicas(v int64)`

SetDesiredReplicas sets DesiredReplicas field to given value.

### HasDesiredReplicas

`func (o *OprsDeployment) HasDesiredReplicas() bool`

HasDesiredReplicas returns a boolean if a field has been set.

### GetEvent

`func (o *OprsDeployment) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *OprsDeployment) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *OprsDeployment) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *OprsDeployment) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetLabels

`func (o *OprsDeployment) GetLabels() []OprsKvpair`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *OprsDeployment) GetLabelsOk() (*[]OprsKvpair, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *OprsDeployment) SetLabels(v []OprsKvpair)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *OprsDeployment) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *OprsDeployment) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *OprsDeployment) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetName

`func (o *OprsDeployment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OprsDeployment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OprsDeployment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OprsDeployment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *OprsDeployment) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *OprsDeployment) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *OprsDeployment) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *OprsDeployment) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetStatus

`func (o *OprsDeployment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OprsDeployment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OprsDeployment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OprsDeployment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeStamp

`func (o *OprsDeployment) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *OprsDeployment) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *OprsDeployment) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.

### HasTimeStamp

`func (o *OprsDeployment) HasTimeStamp() bool`

HasTimeStamp returns a boolean if a field has been set.

### GetUnavailableReplicas

`func (o *OprsDeployment) GetUnavailableReplicas() int64`

GetUnavailableReplicas returns the UnavailableReplicas field if non-nil, zero value otherwise.

### GetUnavailableReplicasOk

`func (o *OprsDeployment) GetUnavailableReplicasOk() (*int64, bool)`

GetUnavailableReplicasOk returns a tuple with the UnavailableReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailableReplicas

`func (o *OprsDeployment) SetUnavailableReplicas(v int64)`

SetUnavailableReplicas sets UnavailableReplicas field to given value.

### HasUnavailableReplicas

`func (o *OprsDeployment) HasUnavailableReplicas() bool`

HasUnavailableReplicas returns a boolean if a field has been set.

### GetAssist

`func (o *OprsDeployment) GetAssist() AssetDeviceRegistrationRelationship`

GetAssist returns the Assist field if non-nil, zero value otherwise.

### GetAssistOk

`func (o *OprsDeployment) GetAssistOk() (*AssetDeviceRegistrationRelationship, bool)`

GetAssistOk returns a tuple with the Assist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssist

`func (o *OprsDeployment) SetAssist(v AssetDeviceRegistrationRelationship)`

SetAssist sets Assist field to given value.

### HasAssist

`func (o *OprsDeployment) HasAssist() bool`

HasAssist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


