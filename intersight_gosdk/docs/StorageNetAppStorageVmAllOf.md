# StorageNetAppStorageVmAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NetAppStorageVm"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NetAppStorageVm"]
**CifsEnabled** | Pointer to **bool** | Status for Common Internet File System protocol ( CIFS ) allowed to run on Vservers. | [optional] [readonly] 
**FcpEnabled** | Pointer to **bool** | Status for Fibre Channel Protocol ( FCP ) allowed to run on Vservers. | [optional] [readonly] 
**IscsiEnabled** | Pointer to **bool** | Status for iSCSI protocol allowed to run on Vservers. | [optional] [readonly] 
**NfsEnabled** | Pointer to **bool** | Status for Network File System Protocol ( NFS ) allowed to run on  Vservers. | [optional] [readonly] 
**NvmeEnabled** | Pointer to **bool** | Status for NVME protocol allowed to run on Vservers. | [optional] [readonly] 
**Array** | Pointer to [**StorageNetAppClusterRelationship**](StorageNetAppClusterRelationship.md) |  | [optional] 
**DiskPool** | Pointer to [**[]StorageNetAppAggregateRelationship**](StorageNetAppAggregateRelationship.md) | An array of relationships to storageNetAppAggregate resources. | [optional] [readonly] 

## Methods

### NewStorageNetAppStorageVmAllOf

`func NewStorageNetAppStorageVmAllOf(classId string, objectType string, ) *StorageNetAppStorageVmAllOf`

NewStorageNetAppStorageVmAllOf instantiates a new StorageNetAppStorageVmAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNetAppStorageVmAllOfWithDefaults

`func NewStorageNetAppStorageVmAllOfWithDefaults() *StorageNetAppStorageVmAllOf`

NewStorageNetAppStorageVmAllOfWithDefaults instantiates a new StorageNetAppStorageVmAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNetAppStorageVmAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNetAppStorageVmAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNetAppStorageVmAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNetAppStorageVmAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNetAppStorageVmAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNetAppStorageVmAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCifsEnabled

`func (o *StorageNetAppStorageVmAllOf) GetCifsEnabled() bool`

GetCifsEnabled returns the CifsEnabled field if non-nil, zero value otherwise.

### GetCifsEnabledOk

`func (o *StorageNetAppStorageVmAllOf) GetCifsEnabledOk() (*bool, bool)`

GetCifsEnabledOk returns a tuple with the CifsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCifsEnabled

`func (o *StorageNetAppStorageVmAllOf) SetCifsEnabled(v bool)`

SetCifsEnabled sets CifsEnabled field to given value.

### HasCifsEnabled

`func (o *StorageNetAppStorageVmAllOf) HasCifsEnabled() bool`

HasCifsEnabled returns a boolean if a field has been set.

### GetFcpEnabled

`func (o *StorageNetAppStorageVmAllOf) GetFcpEnabled() bool`

GetFcpEnabled returns the FcpEnabled field if non-nil, zero value otherwise.

### GetFcpEnabledOk

`func (o *StorageNetAppStorageVmAllOf) GetFcpEnabledOk() (*bool, bool)`

GetFcpEnabledOk returns a tuple with the FcpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcpEnabled

`func (o *StorageNetAppStorageVmAllOf) SetFcpEnabled(v bool)`

SetFcpEnabled sets FcpEnabled field to given value.

### HasFcpEnabled

`func (o *StorageNetAppStorageVmAllOf) HasFcpEnabled() bool`

HasFcpEnabled returns a boolean if a field has been set.

### GetIscsiEnabled

`func (o *StorageNetAppStorageVmAllOf) GetIscsiEnabled() bool`

GetIscsiEnabled returns the IscsiEnabled field if non-nil, zero value otherwise.

### GetIscsiEnabledOk

`func (o *StorageNetAppStorageVmAllOf) GetIscsiEnabledOk() (*bool, bool)`

GetIscsiEnabledOk returns a tuple with the IscsiEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiEnabled

`func (o *StorageNetAppStorageVmAllOf) SetIscsiEnabled(v bool)`

SetIscsiEnabled sets IscsiEnabled field to given value.

### HasIscsiEnabled

`func (o *StorageNetAppStorageVmAllOf) HasIscsiEnabled() bool`

HasIscsiEnabled returns a boolean if a field has been set.

### GetNfsEnabled

`func (o *StorageNetAppStorageVmAllOf) GetNfsEnabled() bool`

GetNfsEnabled returns the NfsEnabled field if non-nil, zero value otherwise.

### GetNfsEnabledOk

`func (o *StorageNetAppStorageVmAllOf) GetNfsEnabledOk() (*bool, bool)`

GetNfsEnabledOk returns a tuple with the NfsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsEnabled

`func (o *StorageNetAppStorageVmAllOf) SetNfsEnabled(v bool)`

SetNfsEnabled sets NfsEnabled field to given value.

### HasNfsEnabled

`func (o *StorageNetAppStorageVmAllOf) HasNfsEnabled() bool`

HasNfsEnabled returns a boolean if a field has been set.

### GetNvmeEnabled

`func (o *StorageNetAppStorageVmAllOf) GetNvmeEnabled() bool`

GetNvmeEnabled returns the NvmeEnabled field if non-nil, zero value otherwise.

### GetNvmeEnabledOk

`func (o *StorageNetAppStorageVmAllOf) GetNvmeEnabledOk() (*bool, bool)`

GetNvmeEnabledOk returns a tuple with the NvmeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNvmeEnabled

`func (o *StorageNetAppStorageVmAllOf) SetNvmeEnabled(v bool)`

SetNvmeEnabled sets NvmeEnabled field to given value.

### HasNvmeEnabled

`func (o *StorageNetAppStorageVmAllOf) HasNvmeEnabled() bool`

HasNvmeEnabled returns a boolean if a field has been set.

### GetArray

`func (o *StorageNetAppStorageVmAllOf) GetArray() StorageNetAppClusterRelationship`

GetArray returns the Array field if non-nil, zero value otherwise.

### GetArrayOk

`func (o *StorageNetAppStorageVmAllOf) GetArrayOk() (*StorageNetAppClusterRelationship, bool)`

GetArrayOk returns a tuple with the Array field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArray

`func (o *StorageNetAppStorageVmAllOf) SetArray(v StorageNetAppClusterRelationship)`

SetArray sets Array field to given value.

### HasArray

`func (o *StorageNetAppStorageVmAllOf) HasArray() bool`

HasArray returns a boolean if a field has been set.

### GetDiskPool

`func (o *StorageNetAppStorageVmAllOf) GetDiskPool() []StorageNetAppAggregateRelationship`

GetDiskPool returns the DiskPool field if non-nil, zero value otherwise.

### GetDiskPoolOk

`func (o *StorageNetAppStorageVmAllOf) GetDiskPoolOk() (*[]StorageNetAppAggregateRelationship, bool)`

GetDiskPoolOk returns a tuple with the DiskPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskPool

`func (o *StorageNetAppStorageVmAllOf) SetDiskPool(v []StorageNetAppAggregateRelationship)`

SetDiskPool sets DiskPool field to given value.

### HasDiskPool

`func (o *StorageNetAppStorageVmAllOf) HasDiskPool() bool`

HasDiskPool returns a boolean if a field has been set.

### SetDiskPoolNil

`func (o *StorageNetAppStorageVmAllOf) SetDiskPoolNil(b bool)`

 SetDiskPoolNil sets the value for DiskPool to be an explicit nil

### UnsetDiskPool
`func (o *StorageNetAppStorageVmAllOf) UnsetDiskPool()`

UnsetDiskPool ensures that no value is present for DiskPool, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


