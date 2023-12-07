# StorageNetAppNamespaceAllOf

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

### NewStorageNetAppNamespaceAllOf

`func NewStorageNetAppNamespaceAllOf(classId string, objectType string, ) *StorageNetAppNamespaceAllOf`

NewStorageNetAppNamespaceAllOf instantiates a new StorageNetAppNamespaceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppNamespaceAllOfWithDefaults

`func NewStorageNetAppNamespaceAllOfWithDefaults() *StorageNetAppNamespaceAllOf`

NewStorageNetAppNamespaceAllOfWithDefaults instantiates a new StorageNetAppNamespaceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppNamespaceAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppNamespaceAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppNamespaceAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppNamespaceAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppNamespaceAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppNamespaceAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetContainerState

`func (o *StorageNetAppNamespaceAllOf) GetContainerState() string`

GetContainerState returns the ContainerState field if non-nil, zero value otherwise.

### GetContainerStateOk

`func (o *StorageNetAppNamespaceAllOf) GetContainerStateOk() (*string, bool)`

GetContainerStateOk returns a tuple with the ContainerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerState

`func (o *StorageNetAppNamespaceAllOf) SetContainerState(v string)`

SetContainerState sets ContainerState field to given value.

### HasContainerState

`func (o *StorageNetAppNamespaceAllOf) HasContainerState() bool`

HasContainerState returns a boolean if a field has been set.

### GetIsMapped

`func (o *StorageNetAppNamespaceAllOf) GetIsMapped() string`

GetIsMapped returns the IsMapped field if non-nil, zero value otherwise.

### GetIsMappedOk

`func (o *StorageNetAppNamespaceAllOf) GetIsMappedOk() (*string, bool)`

GetIsMappedOk returns a tuple with the IsMapped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMapped

`func (o *StorageNetAppNamespaceAllOf) SetIsMapped(v string)`

SetIsMapped sets IsMapped field to given value.

### HasIsMapped

`func (o *StorageNetAppNamespaceAllOf) HasIsMapped() bool`

HasIsMapped returns a boolean if a field has been set.

### GetName

`func (o *StorageNetAppNamespaceAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageNetAppNamespaceAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageNetAppNamespaceAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageNetAppNamespaceAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespaceId

`func (o *StorageNetAppNamespaceAllOf) GetNamespaceId() string`

GetNamespaceId returns the NamespaceId field if non-nil, zero value otherwise.

### GetNamespaceIdOk

`func (o *StorageNetAppNamespaceAllOf) GetNamespaceIdOk() (*string, bool)`

GetNamespaceIdOk returns a tuple with the NamespaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceId

`func (o *StorageNetAppNamespaceAllOf) SetNamespaceId(v string)`

SetNamespaceId sets NamespaceId field to given value.

### HasNamespaceId

`func (o *StorageNetAppNamespaceAllOf) HasNamespaceId() bool`

HasNamespaceId returns a boolean if a field has been set.

### GetPath

`func (o *StorageNetAppNamespaceAllOf) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *StorageNetAppNamespaceAllOf) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *StorageNetAppNamespaceAllOf) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *StorageNetAppNamespaceAllOf) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetState

`func (o *StorageNetAppNamespaceAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StorageNetAppNamespaceAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StorageNetAppNamespaceAllOf) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StorageNetAppNamespaceAllOf) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStorageUtilization

`func (o *StorageNetAppNamespaceAllOf) GetStorageUtilization() StorageBaseCapacity`

GetStorageUtilization returns the StorageUtilization field if non-nil, zero value otherwise.

### GetStorageUtilizationOk

`func (o *StorageNetAppNamespaceAllOf) GetStorageUtilizationOk() (*StorageBaseCapacity, bool)`

GetStorageUtilizationOk returns a tuple with the StorageUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUtilization

`func (o *StorageNetAppNamespaceAllOf) SetStorageUtilization(v StorageBaseCapacity)`

SetStorageUtilization sets StorageUtilization field to given value.

### HasStorageUtilization

`func (o *StorageNetAppNamespaceAllOf) HasStorageUtilization() bool`

HasStorageUtilization returns a boolean if a field has been set.

### SetStorageUtilizationNil

`func (o *StorageNetAppNamespaceAllOf) SetStorageUtilizationNil(b bool)`

 SetStorageUtilizationNil sets the value for StorageUtilization to be an explicit nil

### UnsetStorageUtilization
`func (o *StorageNetAppNamespaceAllOf) UnsetStorageUtilization()`

UnsetStorageUtilization ensures that no value is present for StorageUtilization, not even an explicit nil
### GetSubsystemName

`func (o *StorageNetAppNamespaceAllOf) GetSubsystemName() string`

GetSubsystemName returns the SubsystemName field if non-nil, zero value otherwise.

### GetSubsystemNameOk

`func (o *StorageNetAppNamespaceAllOf) GetSubsystemNameOk() (*string, bool)`

GetSubsystemNameOk returns a tuple with the SubsystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsystemName

`func (o *StorageNetAppNamespaceAllOf) SetSubsystemName(v string)`

SetSubsystemName sets SubsystemName field to given value.

### HasSubsystemName

`func (o *StorageNetAppNamespaceAllOf) HasSubsystemName() bool`

HasSubsystemName returns a boolean if a field has been set.

### GetSvmName

`func (o *StorageNetAppNamespaceAllOf) GetSvmName() string`

GetSvmName returns the SvmName field if non-nil, zero value otherwise.

### GetSvmNameOk

`func (o *StorageNetAppNamespaceAllOf) GetSvmNameOk() (*string, bool)`

GetSvmNameOk returns a tuple with the SvmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvmName

`func (o *StorageNetAppNamespaceAllOf) SetSvmName(v string)`

SetSvmName sets SvmName field to given value.

### HasSvmName

`func (o *StorageNetAppNamespaceAllOf) HasSvmName() bool`

HasSvmName returns a boolean if a field has been set.

### GetUuid

`func (o *StorageNetAppNamespaceAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StorageNetAppNamespaceAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StorageNetAppNamespaceAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StorageNetAppNamespaceAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVolumeName

`func (o *StorageNetAppNamespaceAllOf) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *StorageNetAppNamespaceAllOf) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *StorageNetAppNamespaceAllOf) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *StorageNetAppNamespaceAllOf) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.

### GetArray

`func (o *StorageNetAppNamespaceAllOf) GetArray() StorageNetAppClusterRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageNetAppNamespaceAllOf) GetArrayOk() (*StorageNetAppClusterRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageNetAppNamespaceAllOf) SetArray(v StorageNetAppClusterRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageNetAppNamespaceAllOf) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetStorageContainer

`func (o *StorageNetAppNamespaceAllOf) GetStorageContainer() StorageNetAppVolumeRelationship`

GetStorageContainer returns the StorageContainer field if non-nil, zero value otherwise.

### GetStorageContainerOk

`func (o *StorageNetAppNamespaceAllOf) GetStorageContainerOk() (*StorageNetAppVolumeRelationship, bool)`

GetStorageContainerOk returns a tuple with the StorageContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageContainer

`func (o *StorageNetAppNamespaceAllOf) SetStorageContainer(v StorageNetAppVolumeRelationship)`

SetStorageContainer sets StorageContainer field to given value.

### HasStorageContainer

`func (o *StorageNetAppNamespaceAllOf) HasStorageContainer() bool`

HasStorageContainer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


