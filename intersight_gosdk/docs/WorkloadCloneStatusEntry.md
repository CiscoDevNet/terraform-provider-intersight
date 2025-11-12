# WorkloadCloneStatusEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workload.CloneStatusEntry"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workload.CloneStatusEntry"]
**ClonedObject** | Pointer to [**MoMoRef**](MoMoRef.md) |  | [optional] 
**SourceInput** | Pointer to **string** | The source input data for the clone operation. It is used to track the original input data that is being cloned. | [optional] [readonly] 
**Status** | Pointer to **string** | The status of the clone operation for a workload blueprint input. * &#x60;ToBeCloned&#x60; - The clone operation for the workload blueprint input is pending and has not yet started. * &#x60;Cloned&#x60; - The clone operation for the workload blueprint input has been successfully completed. * &#x60;Failed&#x60; - The clone operation for the workload blueprint input has failed due to an error. * &#x60;InProgress&#x60; - The clone operation for the workload blueprint input is currently in progress. | [optional] [readonly] [default to "ToBeCloned"]

## Methods

### NewWorkloadCloneStatusEntry

`func NewWorkloadCloneStatusEntry(classId string, objectType string, ) *WorkloadCloneStatusEntry`

NewWorkloadCloneStatusEntry instantiates a new WorkloadCloneStatusEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadCloneStatusEntryWithDefaults

`func NewWorkloadCloneStatusEntryWithDefaults() *WorkloadCloneStatusEntry`

NewWorkloadCloneStatusEntryWithDefaults instantiates a new WorkloadCloneStatusEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkloadCloneStatusEntry) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkloadCloneStatusEntry) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkloadCloneStatusEntry) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkloadCloneStatusEntry) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkloadCloneStatusEntry) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkloadCloneStatusEntry) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetClonedObject

`func (o *WorkloadCloneStatusEntry) GetClonedObject() MoMoRef`

GetClonedObject returns the ClonedObject field if non-nil, zero value otherwise.

### GetClonedObjectOk

`func (o *WorkloadCloneStatusEntry) GetClonedObjectOk() (*MoMoRef, bool)`

GetClonedObjectOk returns a tuple with the ClonedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClonedObject

`func (o *WorkloadCloneStatusEntry) SetClonedObject(v MoMoRef)`

SetClonedObject sets ClonedObject field to given value.

### HasClonedObject

`func (o *WorkloadCloneStatusEntry) HasClonedObject() bool`

HasClonedObject returns a boolean if a field has been set.

### GetSourceInput

`func (o *WorkloadCloneStatusEntry) GetSourceInput() string`

GetSourceInput returns the SourceInput field if non-nil, zero value otherwise.

### GetSourceInputOk

`func (o *WorkloadCloneStatusEntry) GetSourceInputOk() (*string, bool)`

GetSourceInputOk returns a tuple with the SourceInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInput

`func (o *WorkloadCloneStatusEntry) SetSourceInput(v string)`

SetSourceInput sets SourceInput field to given value.

### HasSourceInput

`func (o *WorkloadCloneStatusEntry) HasSourceInput() bool`

HasSourceInput returns a boolean if a field has been set.

### GetStatus

`func (o *WorkloadCloneStatusEntry) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkloadCloneStatusEntry) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkloadCloneStatusEntry) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkloadCloneStatusEntry) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


