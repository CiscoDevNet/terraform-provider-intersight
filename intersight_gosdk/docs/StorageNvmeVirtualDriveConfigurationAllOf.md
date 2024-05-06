# StorageNvmeVirtualDriveConfigurationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.NvmeVirtualDriveConfiguration"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.NvmeVirtualDriveConfiguration"]
**AccessPolicy** | Pointer to **string** | This defines the characteristics of a specific virtual drive. | [optional] [readonly] 
**Bootable** | Pointer to **bool** | This defines the characteristics of a specific virtual drive. | [optional] [readonly] 
**ControllerDn** | Pointer to **string** | This defines the characteristics of a specific storage controller. | [optional] [readonly] 
**DedicatedHotSpare** | Pointer to **string** | This defines the characteristics of a specific virtual drive. | [optional] [readonly] 
**DiskCachePolicy** | Pointer to **string** | This defines the characteristics of a specific virtual drive. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the virtual drive. The name can be between 1 and 15 alphanumeric characters. Spaces or any special characters other than - (hyphen) and _ (underscore) are not allowed. | [optional] [readonly] 
**RaidLevel** | Pointer to **string** | This defines the characteristics of a specific virtual drive. | [optional] [readonly] 
**ReadPolicy** | Pointer to **string** | This defines the characteristics of a specific virtual drive. | [optional] [readonly] 
**SelfEncrypt** | Pointer to **string** | This defines the characteristics of a specific virtual drive. | [optional] [readonly] 
**Size** | Pointer to **string** | This defines the characteristics of a specific virtual drive. | [optional] [readonly] 
**SpanDisks** | Pointer to **[]string** |  | [optional] 
**StripSize** | Pointer to **string** | Virtual drive strip size. | [optional] [readonly] 
**WritePolicy** | Pointer to **string** | This defines the characteristics of a specific virtual drive. | [optional] [readonly] 

## Methods

### NewStorageNvmeVirtualDriveConfigurationAllOf

`func NewStorageNvmeVirtualDriveConfigurationAllOf(classId string, objectType string, ) *StorageNvmeVirtualDriveConfigurationAllOf`

NewStorageNvmeVirtualDriveConfigurationAllOf instantiates a new StorageNvmeVirtualDriveConfigurationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageNvmeVirtualDriveConfigurationAllOfWithDefaults

`func NewStorageNvmeVirtualDriveConfigurationAllOfWithDefaults() *StorageNvmeVirtualDriveConfigurationAllOf`

NewStorageNvmeVirtualDriveConfigurationAllOfWithDefaults instantiates a new StorageNvmeVirtualDriveConfigurationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccessPolicy

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) GetAccessPolicy() string`

GetAccessPolicy returns the AccessPolicy field if non-nil, zero value otherwise.

### GetAccessPolicyOk

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) GetAccessPolicyOk() (*string, bool)`

GetAccessPolicyOk returns a tuple with the AccessPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPolicy

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) SetAccessPolicy(v string)`

SetAccessPolicy sets AccessPolicy field to given value.

### HasAccessPolicy

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) HasAccessPolicy() bool`

HasAccessPolicy returns a boolean if a field has been set.

### GetBootable

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) GetBootable() bool`

GetBootable returns the Bootable field if non-nil, zero value otherwise.

### GetBootableOk

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) GetBootableOk() (*bool, bool)`

GetBootableOk returns a tuple with the Bootable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootable

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) SetBootable(v bool)`

SetBootable sets Bootable field to given value.

### HasBootable

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) HasBootable() bool`

HasBootable returns a boolean if a field has been set.

### GetControllerDn

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) GetControllerDn() string`

GetControllerDn returns the ControllerDn field if non-nil, zero value otherwise.

### GetControllerDnOk

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) GetControllerDnOk() (*string, bool)`

GetControllerDnOk returns a tuple with the ControllerDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerDn

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) SetControllerDn(v string)`

SetControllerDn sets ControllerDn field to given value.

### HasControllerDn

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) HasControllerDn() bool`

HasControllerDn returns a boolean if a field has been set.

### GetDedicatedHotSpare

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) GetDedicatedHotSpare() string`

GetDedicatedHotSpare returns the DedicatedHotSpare field if non-nil, zero value otherwise.

### GetDedicatedHotSpareOk

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) GetDedicatedHotSpareOk() (*string, bool)`

GetDedicatedHotSpareOk returns a tuple with the DedicatedHotSpare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedHotSpare

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) SetDedicatedHotSpare(v string)`

SetDedicatedHotSpare sets DedicatedHotSpare field to given value.

### HasDedicatedHotSpare

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) HasDedicatedHotSpare() bool`

HasDedicatedHotSpare returns a boolean if a field has been set.

### GetDiskCachePolicy

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) GetDiskCachePolicy() string`

GetDiskCachePolicy returns the DiskCachePolicy field if non-nil, zero value otherwise.

### GetDiskCachePolicyOk

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) GetDiskCachePolicyOk() (*string, bool)`

GetDiskCachePolicyOk returns a tuple with the DiskCachePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskCachePolicy

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) SetDiskCachePolicy(v string)`

SetDiskCachePolicy sets DiskCachePolicy field to given value.

### HasDiskCachePolicy

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) HasDiskCachePolicy() bool`

HasDiskCachePolicy returns a boolean if a field has been set.

### GetName

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRaidLevel

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) GetRaidLevel() string`

GetRaidLevel returns the RaidLevel field if non-nil, zero value otherwise.

### GetRaidLevelOk

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) GetRaidLevelOk() (*string, bool)`

GetRaidLevelOk returns a tuple with the RaidLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaidLevel

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) SetRaidLevel(v string)`

SetRaidLevel sets RaidLevel field to given value.

### HasRaidLevel

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) HasRaidLevel() bool`

HasRaidLevel returns a boolean if a field has been set.

### GetReadPolicy

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) GetReadPolicy() string`

GetReadPolicy returns the ReadPolicy field if non-nil, zero value otherwise.

### GetReadPolicyOk

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) GetReadPolicyOk() (*string, bool)`

GetReadPolicyOk returns a tuple with the ReadPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadPolicy

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) SetReadPolicy(v string)`

SetReadPolicy sets ReadPolicy field to given value.

### HasReadPolicy

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) HasReadPolicy() bool`

HasReadPolicy returns a boolean if a field has been set.

### GetSelfEncrypt

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) GetSelfEncrypt() string`

GetSelfEncrypt returns the SelfEncrypt field if non-nil, zero value otherwise.

### GetSelfEncryptOk

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) GetSelfEncryptOk() (*string, bool)`

GetSelfEncryptOk returns a tuple with the SelfEncrypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfEncrypt

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) SetSelfEncrypt(v string)`

SetSelfEncrypt sets SelfEncrypt field to given value.

### HasSelfEncrypt

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) HasSelfEncrypt() bool`

HasSelfEncrypt returns a boolean if a field has been set.

### GetSize

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSpanDisks

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) GetSpanDisks() []string`

GetSpanDisks returns the SpanDisks field if non-nil, zero value otherwise.

### GetSpanDisksOk

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) GetSpanDisksOk() (*[]string, bool)`

GetSpanDisksOk returns a tuple with the SpanDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanDisks

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) SetSpanDisks(v []string)`

SetSpanDisks sets SpanDisks field to given value.

### HasSpanDisks

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) HasSpanDisks() bool`

HasSpanDisks returns a boolean if a field has been set.

### SetSpanDisksNil

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) SetSpanDisksNil(b bool)`

 SetSpanDisksNil sets the value for SpanDisks to be an explicit nil

### UnsetSpanDisks
`func (o *StorageNvmeVirtualDriveConfigurationAllOf) UnsetSpanDisks()`

UnsetSpanDisks ensures that no value is present for SpanDisks, not even an explicit nil
### GetStripSize

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) GetStripSize() string`

GetStripSize returns the StripSize field if non-nil, zero value otherwise.

### GetStripSizeOk

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) GetStripSizeOk() (*string, bool)`

GetStripSizeOk returns a tuple with the StripSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripSize

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) SetStripSize(v string)`

SetStripSize sets StripSize field to given value.

### HasStripSize

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) HasStripSize() bool`

HasStripSize returns a boolean if a field has been set.

### GetWritePolicy

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) GetWritePolicy() string`

GetWritePolicy returns the WritePolicy field if non-nil, zero value otherwise.

### GetWritePolicyOk

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) GetWritePolicyOk() (*string, bool)`

GetWritePolicyOk returns a tuple with the WritePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritePolicy

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) SetWritePolicy(v string)`

SetWritePolicy sets WritePolicy field to given value.

### HasWritePolicy

`func (o *StorageNvmeVirtualDriveConfigurationAllOf) HasWritePolicy() bool`

HasWritePolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


