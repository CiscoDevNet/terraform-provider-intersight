# HyperflexVmRestoreOperationAllOf

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

### NewHyperflexVmRestoreOperationAllOf

`func NewHyperflexVmRestoreOperationAllOf(classId string, objectType string, ) *HyperflexVmRestoreOperationAllOf`

NewHyperflexVmRestoreOperationAllOf instantiates a new HyperflexVmRestoreOperationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexVmRestoreOperationAllOfWithDefaults

`func NewHyperflexVmRestoreOperationAllOfWithDefaults() *HyperflexVmRestoreOperationAllOf`

NewHyperflexVmRestoreOperationAllOfWithDefaults instantiates a new HyperflexVmRestoreOperationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexVmRestoreOperationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexVmRestoreOperationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexVmRestoreOperationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexVmRestoreOperationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexVmRestoreOperationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexVmRestoreOperationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetNewName

`func (o *HyperflexVmRestoreOperationAllOf) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *HyperflexVmRestoreOperationAllOf) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *HyperflexVmRestoreOperationAllOf) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *HyperflexVmRestoreOperationAllOf) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetPowerOn

`func (o *HyperflexVmRestoreOperationAllOf) GetPowerOn() bool`

GetPowerOn returns the PowerOn field if non-nil, zero value otherwise.

### GetPowerOnOk

`func (o *HyperflexVmRestoreOperationAllOf) GetPowerOnOk() (*bool, bool)`

GetPowerOnOk returns a tuple with the PowerOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOn

`func (o *HyperflexVmRestoreOperationAllOf) SetPowerOn(v bool)`

SetPowerOn sets PowerOn field to given value.

### HasPowerOn

`func (o *HyperflexVmRestoreOperationAllOf) HasPowerOn() bool`

HasPowerOn returns a boolean if a field has been set.

### GetOrganization

`func (o *HyperflexVmRestoreOperationAllOf) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *HyperflexVmRestoreOperationAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *HyperflexVmRestoreOperationAllOf) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *HyperflexVmRestoreOperationAllOf) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetRestoreEdgeClusterMoid

`func (o *HyperflexVmRestoreOperationAllOf) GetRestoreEdgeClusterMoid() HyperflexClusterRelationship`

GetRestoreEdgeClusterMoid returns the RestoreEdgeClusterMoid field if non-nil, zero value otherwise.

### GetRestoreEdgeClusterMoidOk

`func (o *HyperflexVmRestoreOperationAllOf) GetRestoreEdgeClusterMoidOk() (*HyperflexClusterRelationship, bool)`

GetRestoreEdgeClusterMoidOk returns a tuple with the RestoreEdgeClusterMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreEdgeClusterMoid

`func (o *HyperflexVmRestoreOperationAllOf) SetRestoreEdgeClusterMoid(v HyperflexClusterRelationship)`

SetRestoreEdgeClusterMoid sets RestoreEdgeClusterMoid field to given value.

### HasRestoreEdgeClusterMoid

`func (o *HyperflexVmRestoreOperationAllOf) HasRestoreEdgeClusterMoid() bool`

HasRestoreEdgeClusterMoid returns a boolean if a field has been set.

### GetVmBackupInfoMoid

`func (o *HyperflexVmRestoreOperationAllOf) GetVmBackupInfoMoid() HyperflexVmBackupInfoRelationship`

GetVmBackupInfoMoid returns the VmBackupInfoMoid field if non-nil, zero value otherwise.

### GetVmBackupInfoMoidOk

`func (o *HyperflexVmRestoreOperationAllOf) GetVmBackupInfoMoidOk() (*HyperflexVmBackupInfoRelationship, bool)`

GetVmBackupInfoMoidOk returns a tuple with the VmBackupInfoMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmBackupInfoMoid

`func (o *HyperflexVmRestoreOperationAllOf) SetVmBackupInfoMoid(v HyperflexVmBackupInfoRelationship)`

SetVmBackupInfoMoid sets VmBackupInfoMoid field to given value.

### HasVmBackupInfoMoid

`func (o *HyperflexVmRestoreOperationAllOf) HasVmBackupInfoMoid() bool`

HasVmBackupInfoMoid returns a boolean if a field has been set.

### GetVmSnapshotInfoMoid

`func (o *HyperflexVmRestoreOperationAllOf) GetVmSnapshotInfoMoid() HyperflexVmSnapshotInfoRelationship`

GetVmSnapshotInfoMoid returns the VmSnapshotInfoMoid field if non-nil, zero value otherwise.

### GetVmSnapshotInfoMoidOk

`func (o *HyperflexVmRestoreOperationAllOf) GetVmSnapshotInfoMoidOk() (*HyperflexVmSnapshotInfoRelationship, bool)`

GetVmSnapshotInfoMoidOk returns a tuple with the VmSnapshotInfoMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmSnapshotInfoMoid

`func (o *HyperflexVmRestoreOperationAllOf) SetVmSnapshotInfoMoid(v HyperflexVmSnapshotInfoRelationship)`

SetVmSnapshotInfoMoid sets VmSnapshotInfoMoid field to given value.

### HasVmSnapshotInfoMoid

`func (o *HyperflexVmRestoreOperationAllOf) HasVmSnapshotInfoMoid() bool`

HasVmSnapshotInfoMoid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


