# StorageNetAppNamespace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppNamespace"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppNamespace"]
**ContainerState** | Pointer to **string** | The state of the volume and aggregate that contain the NVMe namespace. Namespaces are only available when their containers are available. | [optional] [readonly] 
**IsMapped** | Pointer to **string** | Reports if the NVMe namespace is mapped to an NVMe subsystem. | [optional] [readonly] 
**Name** | Pointer to **string** | The base name component of the NVMe namespace. | [optional] [readonly] 
**NamespaceId** | Pointer to **string** | The NVMe namespace identifier. An identifier used by an NVMe controller to provide access to the NVMe namespace. The format for an NVMe namespace identifier is 8 hexadecimal digits (zero-filled) followed by a lower case \&quot;h\&quot;. | [optional] [readonly] 
**Path** | Pointer to **string** | The fully qualified path name of the NVMe namespace composed of a \&quot;/vol\&quot; prefix, the volume name, the (optional) qtree name and base name of the namespace. | [optional] [readonly] 
**State** | Pointer to **string** | The state of the NVMe namespace. Normal states for a namespace are online and offline. Other states indicate errors (nvfail, space_error). | [optional] [readonly] 
**StorageUtilization** | Pointer to [**NullableStorageBaseCapacity**](StorageBaseCapacity.md) |  | [optional] 
**SubsystemName** | Pointer to **string** | The NVMe subsystem name to which the NVMe namespace is mapped. | [optional] [readonly] 
**SvmName** | Pointer to **string** | The storage virtual machine name for the NVMe namespace. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Universally unique identifier of the NVMe namespace. | [optional] [readonly] 
**VolumeName** | Pointer to **string** | The volume name in which the NVMe namespace is located. | [optional] [readonly] 
**Array** | Pointer to [**StorageNetAppClusterRelationship**](StorageNetAppClusterRelationship.md) |  | [optional] 
**StorageContainer** | Pointer to [**StorageNetAppVolumeRelationship**](StorageNetAppVolumeRelationship.md) |  | [optional] 

## Methods

### NewStorageNetAppNamespace

`func NewStorageNetAppNamespace(classId string, objectType string, ) *StorageNetAppNamespace`

NewStorageNetAppNamespace instantiates a new StorageNetAppNamespace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppNamespaceWithDefaults

`func NewStorageNetAppNamespaceWithDefaults() *StorageNetAppNamespace`

NewStorageNetAppNamespaceWithDefaults instantiates a new StorageNetAppNamespace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppNamespace) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppNamespace) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppNamespace) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppNamespace) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppNamespace) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppNamespace) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetContainerState

`func (o *StorageNetAppNamespace) GetContainerState() string`

GetContainerState returns the ContainerState field if non-nil, zero value otherwise.

### GetContainerStateOk

`func (o *StorageNetAppNamespace) GetContainerStateOk() (*string, bool)`

GetContainerStateOk returns a tuple with the ContainerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerState

`func (o *StorageNetAppNamespace) SetContainerState(v string)`

SetContainerState sets ContainerState field to given value.

### HasContainerState

`func (o *StorageNetAppNamespace) HasContainerState() bool`

HasContainerState returns a boolean if a field has been set.

### GetIsMapped

`func (o *StorageNetAppNamespace) GetIsMapped() string`

GetIsMapped returns the IsMapped field if non-nil, zero value otherwise.

### GetIsMappedOk

`func (o *StorageNetAppNamespace) GetIsMappedOk() (*string, bool)`

GetIsMappedOk returns a tuple with the IsMapped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMapped

`func (o *StorageNetAppNamespace) SetIsMapped(v string)`

SetIsMapped sets IsMapped field to given value.

### HasIsMapped

`func (o *StorageNetAppNamespace) HasIsMapped() bool`

HasIsMapped returns a boolean if a field has been set.

### GetName

`func (o *StorageNetAppNamespace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageNetAppNamespace) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageNetAppNamespace) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageNetAppNamespace) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespaceId

`func (o *StorageNetAppNamespace) GetNamespaceId() string`

GetNamespaceId returns the NamespaceId field if non-nil, zero value otherwise.

### GetNamespaceIdOk

`func (o *StorageNetAppNamespace) GetNamespaceIdOk() (*string, bool)`

GetNamespaceIdOk returns a tuple with the NamespaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceId

`func (o *StorageNetAppNamespace) SetNamespaceId(v string)`

SetNamespaceId sets NamespaceId field to given value.

### HasNamespaceId

`func (o *StorageNetAppNamespace) HasNamespaceId() bool`

HasNamespaceId returns a boolean if a field has been set.

### GetPath

`func (o *StorageNetAppNamespace) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *StorageNetAppNamespace) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *StorageNetAppNamespace) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *StorageNetAppNamespace) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetState

`func (o *StorageNetAppNamespace) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StorageNetAppNamespace) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StorageNetAppNamespace) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StorageNetAppNamespace) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStorageUtilization

`func (o *StorageNetAppNamespace) GetStorageUtilization() StorageBaseCapacity`

GetStorageUtilization returns the StorageUtilization field if non-nil, zero value otherwise.

### GetStorageUtilizationOk

`func (o *StorageNetAppNamespace) GetStorageUtilizationOk() (*StorageBaseCapacity, bool)`

GetStorageUtilizationOk returns a tuple with the StorageUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUtilization

`func (o *StorageNetAppNamespace) SetStorageUtilization(v StorageBaseCapacity)`

SetStorageUtilization sets StorageUtilization field to given value.

### HasStorageUtilization

`func (o *StorageNetAppNamespace) HasStorageUtilization() bool`

HasStorageUtilization returns a boolean if a field has been set.

### SetStorageUtilizationNil

`func (o *StorageNetAppNamespace) SetStorageUtilizationNil(b bool)`

 SetStorageUtilizationNil sets the value for StorageUtilization to be an explicit nil

### UnsetStorageUtilization
`func (o *StorageNetAppNamespace) UnsetStorageUtilization()`

UnsetStorageUtilization ensures that no value is present for StorageUtilization, not even an explicit nil
### GetSubsystemName

`func (o *StorageNetAppNamespace) GetSubsystemName() string`

GetSubsystemName returns the SubsystemName field if non-nil, zero value otherwise.

### GetSubsystemNameOk

`func (o *StorageNetAppNamespace) GetSubsystemNameOk() (*string, bool)`

GetSubsystemNameOk returns a tuple with the SubsystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsystemName

`func (o *StorageNetAppNamespace) SetSubsystemName(v string)`

SetSubsystemName sets SubsystemName field to given value.

### HasSubsystemName

`func (o *StorageNetAppNamespace) HasSubsystemName() bool`

HasSubsystemName returns a boolean if a field has been set.

### GetSvmName

`func (o *StorageNetAppNamespace) GetSvmName() string`

GetSvmName returns the SvmName field if non-nil, zero value otherwise.

### GetSvmNameOk

`func (o *StorageNetAppNamespace) GetSvmNameOk() (*string, bool)`

GetSvmNameOk returns a tuple with the SvmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvmName

`func (o *StorageNetAppNamespace) SetSvmName(v string)`

SetSvmName sets SvmName field to given value.

### HasSvmName

`func (o *StorageNetAppNamespace) HasSvmName() bool`

HasSvmName returns a boolean if a field has been set.

### GetUuid

`func (o *StorageNetAppNamespace) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StorageNetAppNamespace) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StorageNetAppNamespace) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StorageNetAppNamespace) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVolumeName

`func (o *StorageNetAppNamespace) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *StorageNetAppNamespace) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *StorageNetAppNamespace) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *StorageNetAppNamespace) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.

### GetArray

`func (o *StorageNetAppNamespace) GetArray() StorageNetAppClusterRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageNetAppNamespace) GetArrayOk() (*StorageNetAppClusterRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageNetAppNamespace) SetArray(v StorageNetAppClusterRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageNetAppNamespace) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetStorageContainer

`func (o *StorageNetAppNamespace) GetStorageContainer() StorageNetAppVolumeRelationship`

GetStorageContainer returns the StorageContainer field if non-nil, zero value otherwise.

### GetStorageContainerOk

`func (o *StorageNetAppNamespace) GetStorageContainerOk() (*StorageNetAppVolumeRelationship, bool)`

GetStorageContainerOk returns a tuple with the StorageContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageContainer

`func (o *StorageNetAppNamespace) SetStorageContainer(v StorageNetAppVolumeRelationship)`

SetStorageContainer sets StorageContainer field to given value.

### HasStorageContainer

`func (o *StorageNetAppNamespace) HasStorageContainer() bool`

HasStorageContainer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


