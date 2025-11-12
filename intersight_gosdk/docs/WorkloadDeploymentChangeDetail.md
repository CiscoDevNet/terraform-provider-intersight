# WorkloadDeploymentChangeDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workload.DeploymentChangeDetail"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workload.DeploymentChangeDetail"]
**ChangeType** | Pointer to **string** | The type of change that was applied. * &#x60;None&#x60; - No changes have been made. * &#x60;InputChange&#x60; - A change has been made to an input parameter. For example, an update to the NTP server address. * &#x60;WorkloadDefinitionChange&#x60; - The associated workload definition has changed, such as updating to a new version. * &#x60;WorkloadPreferredVersionChange&#x60; - The deployment was created or updated with the default workload definition version, but the default version was later changed. | [optional] [readonly] [default to "None"]

## Methods

### NewWorkloadDeploymentChangeDetail

`func NewWorkloadDeploymentChangeDetail(classId string, objectType string, ) *WorkloadDeploymentChangeDetail`

NewWorkloadDeploymentChangeDetail instantiates a new WorkloadDeploymentChangeDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadDeploymentChangeDetailWithDefaults

`func NewWorkloadDeploymentChangeDetailWithDefaults() *WorkloadDeploymentChangeDetail`

NewWorkloadDeploymentChangeDetailWithDefaults instantiates a new WorkloadDeploymentChangeDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkloadDeploymentChangeDetail) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkloadDeploymentChangeDetail) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkloadDeploymentChangeDetail) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkloadDeploymentChangeDetail) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkloadDeploymentChangeDetail) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkloadDeploymentChangeDetail) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetChangeType

`func (o *WorkloadDeploymentChangeDetail) GetChangeType() string`

GetChangeType returns the ChangeType field if non-nil, zero value otherwise.

### GetChangeTypeOk

`func (o *WorkloadDeploymentChangeDetail) GetChangeTypeOk() (*string, bool)`

GetChangeTypeOk returns a tuple with the ChangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeType

`func (o *WorkloadDeploymentChangeDetail) SetChangeType(v string)`

SetChangeType sets ChangeType field to given value.

### HasChangeType

`func (o *WorkloadDeploymentChangeDetail) HasChangeType() bool`

HasChangeType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


