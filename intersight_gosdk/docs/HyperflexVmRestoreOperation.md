# HyperflexVmRestoreOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.VmRestoreOperation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.VmRestoreOperation"]
**NewName** | Pointer to **string** | New name for the Virtual Machine after recovery. | [optional] 
**PowerOn** | Pointer to **bool** | Power on the Virtual Machine after recovery. | [optional] [default to true]
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](OrganizationOrganizationRelationship.md) |  | [optional] 
**RestoreEdgeClusterMoid** | Pointer to [**HyperflexClusterRelationship**](HyperflexClusterRelationship.md) |  | [optional] 
**VmBackupInfoMoid** | Pointer to [**HyperflexVmBackupInfoRelationship**](HyperflexVmBackupInfoRelationship.md) |  | [optional] 
**VmSnapshotInfoMoid** | Pointer to [**HyperflexVmSnapshotInfoRelationship**](HyperflexVmSnapshotInfoRelationship.md) |  | [optional] 

## Methods

### NewHyperflexVmRestoreOperation

`func NewHyperflexVmRestoreOperation(classId string, objectType string, ) *HyperflexVmRestoreOperation`

NewHyperflexVmRestoreOperation instantiates a new HyperflexVmRestoreOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexVmRestoreOperationWithDefaults

`func NewHyperflexVmRestoreOperationWithDefaults() *HyperflexVmRestoreOperation`

NewHyperflexVmRestoreOperationWithDefaults instantiates a new HyperflexVmRestoreOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexVmRestoreOperation) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexVmRestoreOperation) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexVmRestoreOperation) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexVmRestoreOperation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexVmRestoreOperation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexVmRestoreOperation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetNewName

`func (o *HyperflexVmRestoreOperation) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *HyperflexVmRestoreOperation) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *HyperflexVmRestoreOperation) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *HyperflexVmRestoreOperation) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetPowerOn

`func (o *HyperflexVmRestoreOperation) GetPowerOn() bool`

GetPowerOn returns the PowerOn field if non-nil, zero value otherwise.

### GetPowerOnOk

`func (o *HyperflexVmRestoreOperation) GetPowerOnOk() (*bool, bool)`

GetPowerOnOk returns a tuple with the PowerOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOn

`func (o *HyperflexVmRestoreOperation) SetPowerOn(v bool)`

SetPowerOn sets PowerOn field to given value.

### HasPowerOn

`func (o *HyperflexVmRestoreOperation) HasPowerOn() bool`

HasPowerOn returns a boolean if a field has been set.

### GetOrganization

`func (o *HyperflexVmRestoreOperation) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *HyperflexVmRestoreOperation) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *HyperflexVmRestoreOperation) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *HyperflexVmRestoreOperation) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetRestoreEdgeClusterMoid

`func (o *HyperflexVmRestoreOperation) GetRestoreEdgeClusterMoid() HyperflexClusterRelationship`

GetRestoreEdgeClusterMoid returns the RestoreEdgeClusterMoid field if non-nil, zero value otherwise.

### GetRestoreEdgeClusterMoidOk

`func (o *HyperflexVmRestoreOperation) GetRestoreEdgeClusterMoidOk() (*HyperflexClusterRelationship, bool)`

GetRestoreEdgeClusterMoidOk returns a tuple with the RestoreEdgeClusterMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreEdgeClusterMoid

`func (o *HyperflexVmRestoreOperation) SetRestoreEdgeClusterMoid(v HyperflexClusterRelationship)`

SetRestoreEdgeClusterMoid sets RestoreEdgeClusterMoid field to given value.

### HasRestoreEdgeClusterMoid

`func (o *HyperflexVmRestoreOperation) HasRestoreEdgeClusterMoid() bool`

HasRestoreEdgeClusterMoid returns a boolean if a field has been set.

### GetVmBackupInfoMoid

`func (o *HyperflexVmRestoreOperation) GetVmBackupInfoMoid() HyperflexVmBackupInfoRelationship`

GetVmBackupInfoMoid returns the VmBackupInfoMoid field if non-nil, zero value otherwise.

### GetVmBackupInfoMoidOk

`func (o *HyperflexVmRestoreOperation) GetVmBackupInfoMoidOk() (*HyperflexVmBackupInfoRelationship, bool)`

GetVmBackupInfoMoidOk returns a tuple with the VmBackupInfoMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmBackupInfoMoid

`func (o *HyperflexVmRestoreOperation) SetVmBackupInfoMoid(v HyperflexVmBackupInfoRelationship)`

SetVmBackupInfoMoid sets VmBackupInfoMoid field to given value.

### HasVmBackupInfoMoid

`func (o *HyperflexVmRestoreOperation) HasVmBackupInfoMoid() bool`

HasVmBackupInfoMoid returns a boolean if a field has been set.

### GetVmSnapshotInfoMoid

`func (o *HyperflexVmRestoreOperation) GetVmSnapshotInfoMoid() HyperflexVmSnapshotInfoRelationship`

GetVmSnapshotInfoMoid returns the VmSnapshotInfoMoid field if non-nil, zero value otherwise.

### GetVmSnapshotInfoMoidOk

`func (o *HyperflexVmRestoreOperation) GetVmSnapshotInfoMoidOk() (*HyperflexVmSnapshotInfoRelationship, bool)`

GetVmSnapshotInfoMoidOk returns a tuple with the VmSnapshotInfoMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmSnapshotInfoMoid

`func (o *HyperflexVmRestoreOperation) SetVmSnapshotInfoMoid(v HyperflexVmSnapshotInfoRelationship)`

SetVmSnapshotInfoMoid sets VmSnapshotInfoMoid field to given value.

### HasVmSnapshotInfoMoid

`func (o *HyperflexVmRestoreOperation) HasVmSnapshotInfoMoid() bool`

HasVmSnapshotInfoMoid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


