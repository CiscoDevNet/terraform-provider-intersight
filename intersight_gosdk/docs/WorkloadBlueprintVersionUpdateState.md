# WorkloadBlueprintVersionUpdateState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workload.BlueprintVersionUpdateState"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workload.BlueprintVersionUpdateState"]
**Status** | Pointer to **string** | The status of the blueprint currently referred to in the workload definition. * &#x60;Updated&#x60; - The blueprint reference is at the latest available version. * &#x60;AutoUpdate&#x60; - The blueprint reference can be updated to the latest version automatically. * &#x60;NeedsNewVersion&#x60; - The blueprint reference cannot be updated without a new version of the definition. * &#x60;Deprecated&#x60; - The referred blueprint is deprecated. * &#x60;UpdateInProgress&#x60; - The blueprint reference update is in progress. The update is in an intermediate state between AutoUpdate and Updated. | [optional] [readonly] [default to "Updated"]
**TargetVersion** | Pointer to [**BlueprintVersionReference**](BlueprintVersionReference.md) |  | [optional] 

## Methods

### NewWorkloadBlueprintVersionUpdateState

`func NewWorkloadBlueprintVersionUpdateState(classId string, objectType string, ) *WorkloadBlueprintVersionUpdateState`

NewWorkloadBlueprintVersionUpdateState instantiates a new WorkloadBlueprintVersionUpdateState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadBlueprintVersionUpdateStateWithDefaults

`func NewWorkloadBlueprintVersionUpdateStateWithDefaults() *WorkloadBlueprintVersionUpdateState`

NewWorkloadBlueprintVersionUpdateStateWithDefaults instantiates a new WorkloadBlueprintVersionUpdateState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkloadBlueprintVersionUpdateState) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkloadBlueprintVersionUpdateState) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkloadBlueprintVersionUpdateState) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkloadBlueprintVersionUpdateState) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkloadBlueprintVersionUpdateState) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkloadBlueprintVersionUpdateState) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetStatus

`func (o *WorkloadBlueprintVersionUpdateState) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkloadBlueprintVersionUpdateState) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkloadBlueprintVersionUpdateState) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkloadBlueprintVersionUpdateState) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTargetVersion

`func (o *WorkloadBlueprintVersionUpdateState) GetTargetVersion() BlueprintVersionReference`

GetTargetVersion returns the TargetVersion field if non-nil, zero value otherwise.

### GetTargetVersionOk

`func (o *WorkloadBlueprintVersionUpdateState) GetTargetVersionOk() (*BlueprintVersionReference, bool)`

GetTargetVersionOk returns a tuple with the TargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVersion

`func (o *WorkloadBlueprintVersionUpdateState) SetTargetVersion(v BlueprintVersionReference)`

SetTargetVersion sets TargetVersion field to given value.

### HasTargetVersion

`func (o *WorkloadBlueprintVersionUpdateState) HasTargetVersion() bool`

HasTargetVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


