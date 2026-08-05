# HciStorageContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.StorageContainer"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.StorageContainer"]
**CacheDeduplication** | Pointer to **string** | Current status of Cache Deduplication for the Storage Container. Possible values: ON, OFF. Does not apply to external storage container case. | [optional] [readonly] 
**ClusterExtId** | Pointer to **string** | The external identifier of the cluster owning the Storage Container. | [optional] [readonly] 
**ClusterName** | Pointer to **string** | The corresponding name of the cluster owning the Storage Container instance. | [optional] [readonly] 
**CompressionDelaySecs** | Pointer to **int32** | The compression delay in seconds. Does not apply to external storage container case. | [optional] [readonly] 
**ErasureCode** | Pointer to **string** | Current status value of Erasure Coding for the Storage Container. Possible values: ON, OFF. Note: Erasure coding is available with certain licensing levels (e.g., Pro licensing or above). Does not apply to external storage container case. | [optional] [readonly] 
**ErasureCodeDelaySecs** | Pointer to **int32** | Delay in performing Erasure Code for the current Storage Container instance. Does not apply to external storage container case. | [optional] [readonly] 
**ExternalStorageExtId** | Pointer to **string** | External Storage extId. Presence of externalStorageExtId indicates this storage-container is an external storage container instead of a HCI storage container. | [optional] [readonly] 
**HasHigherEcFaultDomainPreference** | Pointer to **bool** | Indicates whether to prefer a higher Erasure Code fault domain. Does not apply to external storage container case. | [optional] [readonly] 
**IsCompressionEnabled** | Pointer to **bool** | Indicates whether the compression is enabled for the Storage Container. Does not apply to external storage container case. | [optional] [readonly] 
**IsEncrypted** | Pointer to **bool** | Indicates whether the Storage Container is encrypted or not. Does not apply to external storage container case. | [optional] [readonly] 
**IsExternalStorage** | Pointer to **bool** | Derived value to indicate if the container is backed by an external storage or not. The value is true when the externalStorageExtId is not empty. | [optional] [readonly] 
**IsInlineEcEnabled** | Pointer to **bool** | Indicates whether data written to this Storage Container should be inline erasure-coded or not. This field is only considered if ErasureCoding is enabled. Inline Erasure-Coded refers to a data protection method where erasure coding is applied directly and immediately to data as it is written (in-line) to the storage system. Does not apply to external storage container case. | [optional] [readonly] 
**IsInternal** | Pointer to **bool** | Indicates whether the Storage Container is internal and is managed by Nutanix. | [optional] [readonly] 
**IsMarkedForRemoval** | Pointer to **bool** | Indicates whether the Storage Container is marked for removal. This field is set when the Storage Container is about to be destroyed. | [optional] [readonly] 
**IsNfsAllowlistInherited** | Pointer to **bool** | Indicates whether the NFS allowlist is inherited from the global configuration. Does not apply to external storage container case. | [optional] [readonly] 
**IsShared** | Pointer to **bool** | Indicates whether the Storage Container is shared. When shared, all PEs registered under the same PC can access and utilize its storage. Once the field is set during the creation, it is immutable except in the case of SelfServiceContainer. Even for SelfServiceContainer only FALSE to TRUE is allowed. Does not apply to external storage container case. | [optional] [readonly] 
**IsSoftwareEncryptionEnabled** | Pointer to **bool** | Indicates whether the Storage Container instance has software encryption enabled. Does not apply to external storage container case. | [optional] [readonly] 
**LogicalAdvertisedCapacityBytes** | Pointer to **int64** | Maximum capacity of the Storage Container as defined by the user. Does not apply to external storage container case. | [optional] [readonly] 
**LogicalExplicitReservedCapacityBytes** | Pointer to **int64** | Total reserved size (in bytes) of the Storage Container (set by Admin). This also includes the replication factor of the Storage Container. The actual reserved capacity of the Storage Container will be the maximum of explicitReservedCapacity and implicitReservedCapacity. Does not apply to external storage container case. | [optional] [readonly] 
**LogicalImplicitReservedCapacityBytes** | Pointer to **int64** | Sum of the reservations provisioned on all vDisks in the Storage Container. The actual reserved capacity of the Storage Container will be the maximum of explicitReservedCapacity and implicitReservedCapacity. Does not apply to external storage container case. | [optional] [readonly] 
**MaxCapacityBytes** | Pointer to **int64** | Maximum physical capacity of the Storage Container in bytes. Does not apply to external storage container case. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the Storage Container. Prism Central requires that the name of Storage Container should be unique in every registered cluster. | [optional] [readonly] 
**OnDiskDedup** | Pointer to **string** | Current status of Disk deduplication for the Storage Container. Possible values: - POST_PROCESS: Deduplication is enabled for the Storage Container instance. - OFF: Deduplication is disabled for the Storage Container instance. Does not apply to external storage container case. | [optional] [readonly] 
**ReplicationFactor** | Pointer to **int32** | Replication factor of the Storage Container. Does not apply to external storage container case. | [optional] [readonly] 
**StorageContainerExtId** | Pointer to **string** | The external identifier of the Storage Container. | [optional] [readonly] 
**StoragePoolExtId** | Pointer to **string** | The external identifier of the Storage Pool owning the Storage Container instance. Does not apply to external storage container case. | [optional] [readonly] 
**Cluster** | Pointer to [**NullableHciClusterRelationship**](HciClusterRelationship.md) |  | [optional] 
**ExternalStorage** | Pointer to [**NullableHciExternalStorageRelationship**](HciExternalStorageRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewHciStorageContainer

`func NewHciStorageContainer(classId string, objectType string, ) *HciStorageContainer`

NewHciStorageContainer instantiates a new HciStorageContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciStorageContainerWithDefaults

`func NewHciStorageContainerWithDefaults() *HciStorageContainer`

NewHciStorageContainerWithDefaults instantiates a new HciStorageContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciStorageContainer) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciStorageContainer) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciStorageContainer) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciStorageContainer) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciStorageContainer) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciStorageContainer) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCacheDeduplication

`func (o *HciStorageContainer) GetCacheDeduplication() string`

GetCacheDeduplication returns the CacheDeduplication field if non-nil, zero value otherwise.

### GetCacheDeduplicationOk

`func (o *HciStorageContainer) GetCacheDeduplicationOk() (*string, bool)`

GetCacheDeduplicationOk returns a tuple with the CacheDeduplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheDeduplication

`func (o *HciStorageContainer) SetCacheDeduplication(v string)`

SetCacheDeduplication sets CacheDeduplication field to given value.

### HasCacheDeduplication

`func (o *HciStorageContainer) HasCacheDeduplication() bool`

HasCacheDeduplication returns a boolean if a field has been set.

### GetClusterExtId

`func (o *HciStorageContainer) GetClusterExtId() string`

GetClusterExtId returns the ClusterExtId field if non-nil, zero value otherwise.

### GetClusterExtIdOk

`func (o *HciStorageContainer) GetClusterExtIdOk() (*string, bool)`

GetClusterExtIdOk returns a tuple with the ClusterExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterExtId

`func (o *HciStorageContainer) SetClusterExtId(v string)`

SetClusterExtId sets ClusterExtId field to given value.

### HasClusterExtId

`func (o *HciStorageContainer) HasClusterExtId() bool`

HasClusterExtId returns a boolean if a field has been set.

### GetClusterName

`func (o *HciStorageContainer) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *HciStorageContainer) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *HciStorageContainer) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *HciStorageContainer) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetCompressionDelaySecs

`func (o *HciStorageContainer) GetCompressionDelaySecs() int32`

GetCompressionDelaySecs returns the CompressionDelaySecs field if non-nil, zero value otherwise.

### GetCompressionDelaySecsOk

`func (o *HciStorageContainer) GetCompressionDelaySecsOk() (*int32, bool)`

GetCompressionDelaySecsOk returns a tuple with the CompressionDelaySecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionDelaySecs

`func (o *HciStorageContainer) SetCompressionDelaySecs(v int32)`

SetCompressionDelaySecs sets CompressionDelaySecs field to given value.

### HasCompressionDelaySecs

`func (o *HciStorageContainer) HasCompressionDelaySecs() bool`

HasCompressionDelaySecs returns a boolean if a field has been set.

### GetErasureCode

`func (o *HciStorageContainer) GetErasureCode() string`

GetErasureCode returns the ErasureCode field if non-nil, zero value otherwise.

### GetErasureCodeOk

`func (o *HciStorageContainer) GetErasureCodeOk() (*string, bool)`

GetErasureCodeOk returns a tuple with the ErasureCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErasureCode

`func (o *HciStorageContainer) SetErasureCode(v string)`

SetErasureCode sets ErasureCode field to given value.

### HasErasureCode

`func (o *HciStorageContainer) HasErasureCode() bool`

HasErasureCode returns a boolean if a field has been set.

### GetErasureCodeDelaySecs

`func (o *HciStorageContainer) GetErasureCodeDelaySecs() int32`

GetErasureCodeDelaySecs returns the ErasureCodeDelaySecs field if non-nil, zero value otherwise.

### GetErasureCodeDelaySecsOk

`func (o *HciStorageContainer) GetErasureCodeDelaySecsOk() (*int32, bool)`

GetErasureCodeDelaySecsOk returns a tuple with the ErasureCodeDelaySecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErasureCodeDelaySecs

`func (o *HciStorageContainer) SetErasureCodeDelaySecs(v int32)`

SetErasureCodeDelaySecs sets ErasureCodeDelaySecs field to given value.

### HasErasureCodeDelaySecs

`func (o *HciStorageContainer) HasErasureCodeDelaySecs() bool`

HasErasureCodeDelaySecs returns a boolean if a field has been set.

### GetExternalStorageExtId

`func (o *HciStorageContainer) GetExternalStorageExtId() string`

GetExternalStorageExtId returns the ExternalStorageExtId field if non-nil, zero value otherwise.

### GetExternalStorageExtIdOk

`func (o *HciStorageContainer) GetExternalStorageExtIdOk() (*string, bool)`

GetExternalStorageExtIdOk returns a tuple with the ExternalStorageExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStorageExtId

`func (o *HciStorageContainer) SetExternalStorageExtId(v string)`

SetExternalStorageExtId sets ExternalStorageExtId field to given value.

### HasExternalStorageExtId

`func (o *HciStorageContainer) HasExternalStorageExtId() bool`

HasExternalStorageExtId returns a boolean if a field has been set.

### GetHasHigherEcFaultDomainPreference

`func (o *HciStorageContainer) GetHasHigherEcFaultDomainPreference() bool`

GetHasHigherEcFaultDomainPreference returns the HasHigherEcFaultDomainPreference field if non-nil, zero value otherwise.

### GetHasHigherEcFaultDomainPreferenceOk

`func (o *HciStorageContainer) GetHasHigherEcFaultDomainPreferenceOk() (*bool, bool)`

GetHasHigherEcFaultDomainPreferenceOk returns a tuple with the HasHigherEcFaultDomainPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasHigherEcFaultDomainPreference

`func (o *HciStorageContainer) SetHasHigherEcFaultDomainPreference(v bool)`

SetHasHigherEcFaultDomainPreference sets HasHigherEcFaultDomainPreference field to given value.

### HasHasHigherEcFaultDomainPreference

`func (o *HciStorageContainer) HasHasHigherEcFaultDomainPreference() bool`

HasHasHigherEcFaultDomainPreference returns a boolean if a field has been set.

### GetIsCompressionEnabled

`func (o *HciStorageContainer) GetIsCompressionEnabled() bool`

GetIsCompressionEnabled returns the IsCompressionEnabled field if non-nil, zero value otherwise.

### GetIsCompressionEnabledOk

`func (o *HciStorageContainer) GetIsCompressionEnabledOk() (*bool, bool)`

GetIsCompressionEnabledOk returns a tuple with the IsCompressionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompressionEnabled

`func (o *HciStorageContainer) SetIsCompressionEnabled(v bool)`

SetIsCompressionEnabled sets IsCompressionEnabled field to given value.

### HasIsCompressionEnabled

`func (o *HciStorageContainer) HasIsCompressionEnabled() bool`

HasIsCompressionEnabled returns a boolean if a field has been set.

### GetIsEncrypted

`func (o *HciStorageContainer) GetIsEncrypted() bool`

GetIsEncrypted returns the IsEncrypted field if non-nil, zero value otherwise.

### GetIsEncryptedOk

`func (o *HciStorageContainer) GetIsEncryptedOk() (*bool, bool)`

GetIsEncryptedOk returns a tuple with the IsEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEncrypted

`func (o *HciStorageContainer) SetIsEncrypted(v bool)`

SetIsEncrypted sets IsEncrypted field to given value.

### HasIsEncrypted

`func (o *HciStorageContainer) HasIsEncrypted() bool`

HasIsEncrypted returns a boolean if a field has been set.

### GetIsExternalStorage

`func (o *HciStorageContainer) GetIsExternalStorage() bool`

GetIsExternalStorage returns the IsExternalStorage field if non-nil, zero value otherwise.

### GetIsExternalStorageOk

`func (o *HciStorageContainer) GetIsExternalStorageOk() (*bool, bool)`

GetIsExternalStorageOk returns a tuple with the IsExternalStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternalStorage

`func (o *HciStorageContainer) SetIsExternalStorage(v bool)`

SetIsExternalStorage sets IsExternalStorage field to given value.

### HasIsExternalStorage

`func (o *HciStorageContainer) HasIsExternalStorage() bool`

HasIsExternalStorage returns a boolean if a field has been set.

### GetIsInlineEcEnabled

`func (o *HciStorageContainer) GetIsInlineEcEnabled() bool`

GetIsInlineEcEnabled returns the IsInlineEcEnabled field if non-nil, zero value otherwise.

### GetIsInlineEcEnabledOk

`func (o *HciStorageContainer) GetIsInlineEcEnabledOk() (*bool, bool)`

GetIsInlineEcEnabledOk returns a tuple with the IsInlineEcEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInlineEcEnabled

`func (o *HciStorageContainer) SetIsInlineEcEnabled(v bool)`

SetIsInlineEcEnabled sets IsInlineEcEnabled field to given value.

### HasIsInlineEcEnabled

`func (o *HciStorageContainer) HasIsInlineEcEnabled() bool`

HasIsInlineEcEnabled returns a boolean if a field has been set.

### GetIsInternal

`func (o *HciStorageContainer) GetIsInternal() bool`

GetIsInternal returns the IsInternal field if non-nil, zero value otherwise.

### GetIsInternalOk

`func (o *HciStorageContainer) GetIsInternalOk() (*bool, bool)`

GetIsInternalOk returns a tuple with the IsInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInternal

`func (o *HciStorageContainer) SetIsInternal(v bool)`

SetIsInternal sets IsInternal field to given value.

### HasIsInternal

`func (o *HciStorageContainer) HasIsInternal() bool`

HasIsInternal returns a boolean if a field has been set.

### GetIsMarkedForRemoval

`func (o *HciStorageContainer) GetIsMarkedForRemoval() bool`

GetIsMarkedForRemoval returns the IsMarkedForRemoval field if non-nil, zero value otherwise.

### GetIsMarkedForRemovalOk

`func (o *HciStorageContainer) GetIsMarkedForRemovalOk() (*bool, bool)`

GetIsMarkedForRemovalOk returns a tuple with the IsMarkedForRemoval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMarkedForRemoval

`func (o *HciStorageContainer) SetIsMarkedForRemoval(v bool)`

SetIsMarkedForRemoval sets IsMarkedForRemoval field to given value.

### HasIsMarkedForRemoval

`func (o *HciStorageContainer) HasIsMarkedForRemoval() bool`

HasIsMarkedForRemoval returns a boolean if a field has been set.

### GetIsNfsAllowlistInherited

`func (o *HciStorageContainer) GetIsNfsAllowlistInherited() bool`

GetIsNfsAllowlistInherited returns the IsNfsAllowlistInherited field if non-nil, zero value otherwise.

### GetIsNfsAllowlistInheritedOk

`func (o *HciStorageContainer) GetIsNfsAllowlistInheritedOk() (*bool, bool)`

GetIsNfsAllowlistInheritedOk returns a tuple with the IsNfsAllowlistInherited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNfsAllowlistInherited

`func (o *HciStorageContainer) SetIsNfsAllowlistInherited(v bool)`

SetIsNfsAllowlistInherited sets IsNfsAllowlistInherited field to given value.

### HasIsNfsAllowlistInherited

`func (o *HciStorageContainer) HasIsNfsAllowlistInherited() bool`

HasIsNfsAllowlistInherited returns a boolean if a field has been set.

### GetIsShared

`func (o *HciStorageContainer) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *HciStorageContainer) GetIsSharedOk() (*bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShared

`func (o *HciStorageContainer) SetIsShared(v bool)`

SetIsShared sets IsShared field to given value.

### HasIsShared

`func (o *HciStorageContainer) HasIsShared() bool`

HasIsShared returns a boolean if a field has been set.

### GetIsSoftwareEncryptionEnabled

`func (o *HciStorageContainer) GetIsSoftwareEncryptionEnabled() bool`

GetIsSoftwareEncryptionEnabled returns the IsSoftwareEncryptionEnabled field if non-nil, zero value otherwise.

### GetIsSoftwareEncryptionEnabledOk

`func (o *HciStorageContainer) GetIsSoftwareEncryptionEnabledOk() (*bool, bool)`

GetIsSoftwareEncryptionEnabledOk returns a tuple with the IsSoftwareEncryptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSoftwareEncryptionEnabled

`func (o *HciStorageContainer) SetIsSoftwareEncryptionEnabled(v bool)`

SetIsSoftwareEncryptionEnabled sets IsSoftwareEncryptionEnabled field to given value.

### HasIsSoftwareEncryptionEnabled

`func (o *HciStorageContainer) HasIsSoftwareEncryptionEnabled() bool`

HasIsSoftwareEncryptionEnabled returns a boolean if a field has been set.

### GetLogicalAdvertisedCapacityBytes

`func (o *HciStorageContainer) GetLogicalAdvertisedCapacityBytes() int64`

GetLogicalAdvertisedCapacityBytes returns the LogicalAdvertisedCapacityBytes field if non-nil, zero value otherwise.

### GetLogicalAdvertisedCapacityBytesOk

`func (o *HciStorageContainer) GetLogicalAdvertisedCapacityBytesOk() (*int64, bool)`

GetLogicalAdvertisedCapacityBytesOk returns a tuple with the LogicalAdvertisedCapacityBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalAdvertisedCapacityBytes

`func (o *HciStorageContainer) SetLogicalAdvertisedCapacityBytes(v int64)`

SetLogicalAdvertisedCapacityBytes sets LogicalAdvertisedCapacityBytes field to given value.

### HasLogicalAdvertisedCapacityBytes

`func (o *HciStorageContainer) HasLogicalAdvertisedCapacityBytes() bool`

HasLogicalAdvertisedCapacityBytes returns a boolean if a field has been set.

### GetLogicalExplicitReservedCapacityBytes

`func (o *HciStorageContainer) GetLogicalExplicitReservedCapacityBytes() int64`

GetLogicalExplicitReservedCapacityBytes returns the LogicalExplicitReservedCapacityBytes field if non-nil, zero value otherwise.

### GetLogicalExplicitReservedCapacityBytesOk

`func (o *HciStorageContainer) GetLogicalExplicitReservedCapacityBytesOk() (*int64, bool)`

GetLogicalExplicitReservedCapacityBytesOk returns a tuple with the LogicalExplicitReservedCapacityBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalExplicitReservedCapacityBytes

`func (o *HciStorageContainer) SetLogicalExplicitReservedCapacityBytes(v int64)`

SetLogicalExplicitReservedCapacityBytes sets LogicalExplicitReservedCapacityBytes field to given value.

### HasLogicalExplicitReservedCapacityBytes

`func (o *HciStorageContainer) HasLogicalExplicitReservedCapacityBytes() bool`

HasLogicalExplicitReservedCapacityBytes returns a boolean if a field has been set.

### GetLogicalImplicitReservedCapacityBytes

`func (o *HciStorageContainer) GetLogicalImplicitReservedCapacityBytes() int64`

GetLogicalImplicitReservedCapacityBytes returns the LogicalImplicitReservedCapacityBytes field if non-nil, zero value otherwise.

### GetLogicalImplicitReservedCapacityBytesOk

`func (o *HciStorageContainer) GetLogicalImplicitReservedCapacityBytesOk() (*int64, bool)`

GetLogicalImplicitReservedCapacityBytesOk returns a tuple with the LogicalImplicitReservedCapacityBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalImplicitReservedCapacityBytes

`func (o *HciStorageContainer) SetLogicalImplicitReservedCapacityBytes(v int64)`

SetLogicalImplicitReservedCapacityBytes sets LogicalImplicitReservedCapacityBytes field to given value.

### HasLogicalImplicitReservedCapacityBytes

`func (o *HciStorageContainer) HasLogicalImplicitReservedCapacityBytes() bool`

HasLogicalImplicitReservedCapacityBytes returns a boolean if a field has been set.

### GetMaxCapacityBytes

`func (o *HciStorageContainer) GetMaxCapacityBytes() int64`

GetMaxCapacityBytes returns the MaxCapacityBytes field if non-nil, zero value otherwise.

### GetMaxCapacityBytesOk

`func (o *HciStorageContainer) GetMaxCapacityBytesOk() (*int64, bool)`

GetMaxCapacityBytesOk returns a tuple with the MaxCapacityBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCapacityBytes

`func (o *HciStorageContainer) SetMaxCapacityBytes(v int64)`

SetMaxCapacityBytes sets MaxCapacityBytes field to given value.

### HasMaxCapacityBytes

`func (o *HciStorageContainer) HasMaxCapacityBytes() bool`

HasMaxCapacityBytes returns a boolean if a field has been set.

### GetName

`func (o *HciStorageContainer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HciStorageContainer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HciStorageContainer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HciStorageContainer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOnDiskDedup

`func (o *HciStorageContainer) GetOnDiskDedup() string`

GetOnDiskDedup returns the OnDiskDedup field if non-nil, zero value otherwise.

### GetOnDiskDedupOk

`func (o *HciStorageContainer) GetOnDiskDedupOk() (*string, bool)`

GetOnDiskDedupOk returns a tuple with the OnDiskDedup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDiskDedup

`func (o *HciStorageContainer) SetOnDiskDedup(v string)`

SetOnDiskDedup sets OnDiskDedup field to given value.

### HasOnDiskDedup

`func (o *HciStorageContainer) HasOnDiskDedup() bool`

HasOnDiskDedup returns a boolean if a field has been set.

### GetReplicationFactor

`func (o *HciStorageContainer) GetReplicationFactor() int32`

GetReplicationFactor returns the ReplicationFactor field if non-nil, zero value otherwise.

### GetReplicationFactorOk

`func (o *HciStorageContainer) GetReplicationFactorOk() (*int32, bool)`

GetReplicationFactorOk returns a tuple with the ReplicationFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationFactor

`func (o *HciStorageContainer) SetReplicationFactor(v int32)`

SetReplicationFactor sets ReplicationFactor field to given value.

### HasReplicationFactor

`func (o *HciStorageContainer) HasReplicationFactor() bool`

HasReplicationFactor returns a boolean if a field has been set.

### GetStorageContainerExtId

`func (o *HciStorageContainer) GetStorageContainerExtId() string`

GetStorageContainerExtId returns the StorageContainerExtId field if non-nil, zero value otherwise.

### GetStorageContainerExtIdOk

`func (o *HciStorageContainer) GetStorageContainerExtIdOk() (*string, bool)`

GetStorageContainerExtIdOk returns a tuple with the StorageContainerExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageContainerExtId

`func (o *HciStorageContainer) SetStorageContainerExtId(v string)`

SetStorageContainerExtId sets StorageContainerExtId field to given value.

### HasStorageContainerExtId

`func (o *HciStorageContainer) HasStorageContainerExtId() bool`

HasStorageContainerExtId returns a boolean if a field has been set.

### GetStoragePoolExtId

`func (o *HciStorageContainer) GetStoragePoolExtId() string`

GetStoragePoolExtId returns the StoragePoolExtId field if non-nil, zero value otherwise.

### GetStoragePoolExtIdOk

`func (o *HciStorageContainer) GetStoragePoolExtIdOk() (*string, bool)`

GetStoragePoolExtIdOk returns a tuple with the StoragePoolExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePoolExtId

`func (o *HciStorageContainer) SetStoragePoolExtId(v string)`

SetStoragePoolExtId sets StoragePoolExtId field to given value.

### HasStoragePoolExtId

`func (o *HciStorageContainer) HasStoragePoolExtId() bool`

HasStoragePoolExtId returns a boolean if a field has been set.

### GetCluster

`func (o *HciStorageContainer) GetCluster() HciClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HciStorageContainer) GetClusterOk() (*HciClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HciStorageContainer) SetCluster(v HciClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HciStorageContainer) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *HciStorageContainer) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *HciStorageContainer) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetExternalStorage

`func (o *HciStorageContainer) GetExternalStorage() HciExternalStorageRelationship`

GetExternalStorage returns the ExternalStorage field if non-nil, zero value otherwise.

### GetExternalStorageOk

`func (o *HciStorageContainer) GetExternalStorageOk() (*HciExternalStorageRelationship, bool)`

GetExternalStorageOk returns a tuple with the ExternalStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStorage

`func (o *HciStorageContainer) SetExternalStorage(v HciExternalStorageRelationship)`

SetExternalStorage sets ExternalStorage field to given value.

### HasExternalStorage

`func (o *HciStorageContainer) HasExternalStorage() bool`

HasExternalStorage returns a boolean if a field has been set.

### SetExternalStorageNil

`func (o *HciStorageContainer) SetExternalStorageNil(b bool)`

 SetExternalStorageNil sets the value for ExternalStorage to be an explicit nil

### UnsetExternalStorage
`func (o *HciStorageContainer) UnsetExternalStorage()`

UnsetExternalStorage ensures that no value is present for ExternalStorage, not even an explicit nil
### GetRegisteredDevice

`func (o *HciStorageContainer) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *HciStorageContainer) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *HciStorageContainer) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *HciStorageContainer) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *HciStorageContainer) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *HciStorageContainer) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


