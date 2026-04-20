# WorkloadRenameRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workload.RenameRequest"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workload.RenameRequest"]
**Action** | Pointer to **string** | The action to be taken for the rename operation on an object. * &#x60;None&#x60; - No action is to be taken for the rename request. * &#x60;Rename&#x60; - The object is to be renamed with the new name in the request. | [optional] [default to "None"]
**NewName** | Pointer to **string** | The new name for the object. This name will be used to rename all objects associated with it. | [optional] 
**Status** | Pointer to **string** | The status of the rename operation for an object. * &#x60;None&#x60; - No rename operation is in progress, the last rename operation succeeded, or rename has not been performed on the object. * &#x60;Updating&#x60; - The object is currently being renamed. * &#x60;UpdateScheduled&#x60; - The rename request for the object has been accepted and will be processed. | [optional] [readonly] [default to "None"]

## Methods

### NewWorkloadRenameRequest

`func NewWorkloadRenameRequest(classId string, objectType string, ) *WorkloadRenameRequest`

NewWorkloadRenameRequest instantiates a new WorkloadRenameRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadRenameRequestWithDefaults

`func NewWorkloadRenameRequestWithDefaults() *WorkloadRenameRequest`

NewWorkloadRenameRequestWithDefaults instantiates a new WorkloadRenameRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkloadRenameRequest) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkloadRenameRequest) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkloadRenameRequest) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkloadRenameRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkloadRenameRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkloadRenameRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAction

`func (o *WorkloadRenameRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WorkloadRenameRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WorkloadRenameRequest) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *WorkloadRenameRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetNewName

`func (o *WorkloadRenameRequest) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *WorkloadRenameRequest) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *WorkloadRenameRequest) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *WorkloadRenameRequest) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetStatus

`func (o *WorkloadRenameRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkloadRenameRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkloadRenameRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkloadRenameRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


