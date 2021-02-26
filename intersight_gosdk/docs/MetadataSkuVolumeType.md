# MetadataSkuVolumeType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "metadata.SkuVolumeType"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "metadata.SkuVolumeType"]
**IopsUnit** | Pointer to **string** | The units to measure IOPS. | [optional] 
**IsBootable** | Pointer to **bool** | Indicates if this volume can be used as a boot volume. | [optional] 
**IsDefault** | Pointer to **bool** | Flag to indicate if this is a default volume. Default volumes will be used when an instance type is launched unless another volume type is specified. | [optional] 
**IsProvisionedIops** | Pointer to **bool** | Indicates if this volume type supports provisioned IOPS. | [optional] 
**MaxIops** | Pointer to **float64** | The max I/O operations per second that this volume type supports. Read or write numbers does not go beyond this value. | [optional] 
**MaxReadIops** | Pointer to **float64** | The maximum read IOPS that this volume type supports. | [optional] 
**MaxReadThroughput** | Pointer to **float64** | The maximum read throughput limit for this volume type. | [optional] 
**MaxThroughput** | Pointer to **float64** | The maximum throughput limit for this volume type. | [optional] 
**MaxVolumeSize** | Pointer to **float64** | The maximum storage size that this volume type supports. | [optional] 
**MaxWriteIops** | Pointer to **float64** | The maximum write IOPS that this volume type supports. | [optional] 
**MaxWriteThroughput** | Pointer to **float64** | The maximum write throughput limit for this volume type. | [optional] 
**MinVolumeSize** | Pointer to **float64** | The minimum storage size that this volume type supports. | [optional] 
**ThroughputUnit** | Pointer to **string** | The units for measuring throughput. | [optional] 
**Type** | Pointer to **string** | The volume type like gp2 or st1. | [optional] 
**VolumeSizeUnit** | Pointer to **string** | The units for measuring volume size. | [optional] 

## Methods

### NewMetadataSkuVolumeType

`func NewMetadataSkuVolumeType(classId string, objectType string, ) *MetadataSkuVolumeType`

NewMetadataSkuVolumeType instantiates a new MetadataSkuVolumeType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataSkuVolumeTypeWithDefaults

`func NewMetadataSkuVolumeTypeWithDefaults() *MetadataSkuVolumeType`

NewMetadataSkuVolumeTypeWithDefaults instantiates a new MetadataSkuVolumeType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MetadataSkuVolumeType) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MetadataSkuVolumeType) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MetadataSkuVolumeType) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MetadataSkuVolumeType) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MetadataSkuVolumeType) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MetadataSkuVolumeType) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIopsUnit

`func (o *MetadataSkuVolumeType) GetIopsUnit() string`

GetIopsUnit returns the IopsUnit field if non-nil, zero value otherwise.

### GetIopsUnitOk

`func (o *MetadataSkuVolumeType) GetIopsUnitOk() (*string, bool)`

GetIopsUnitOk returns a tuple with the IopsUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsUnit

`func (o *MetadataSkuVolumeType) SetIopsUnit(v string)`

SetIopsUnit sets IopsUnit field to given value.

### HasIopsUnit

`func (o *MetadataSkuVolumeType) HasIopsUnit() bool`

HasIopsUnit returns a boolean if a field has been set.

### GetIsBootable

`func (o *MetadataSkuVolumeType) GetIsBootable() bool`

GetIsBootable returns the IsBootable field if non-nil, zero value otherwise.

### GetIsBootableOk

`func (o *MetadataSkuVolumeType) GetIsBootableOk() (*bool, bool)`

GetIsBootableOk returns a tuple with the IsBootable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBootable

`func (o *MetadataSkuVolumeType) SetIsBootable(v bool)`

SetIsBootable sets IsBootable field to given value.

### HasIsBootable

`func (o *MetadataSkuVolumeType) HasIsBootable() bool`

HasIsBootable returns a boolean if a field has been set.

### GetIsDefault

`func (o *MetadataSkuVolumeType) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *MetadataSkuVolumeType) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *MetadataSkuVolumeType) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *MetadataSkuVolumeType) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetIsProvisionedIops

`func (o *MetadataSkuVolumeType) GetIsProvisionedIops() bool`

GetIsProvisionedIops returns the IsProvisionedIops field if non-nil, zero value otherwise.

### GetIsProvisionedIopsOk

`func (o *MetadataSkuVolumeType) GetIsProvisionedIopsOk() (*bool, bool)`

GetIsProvisionedIopsOk returns a tuple with the IsProvisionedIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProvisionedIops

`func (o *MetadataSkuVolumeType) SetIsProvisionedIops(v bool)`

SetIsProvisionedIops sets IsProvisionedIops field to given value.

### HasIsProvisionedIops

`func (o *MetadataSkuVolumeType) HasIsProvisionedIops() bool`

HasIsProvisionedIops returns a boolean if a field has been set.

### GetMaxIops

`func (o *MetadataSkuVolumeType) GetMaxIops() float64`

GetMaxIops returns the MaxIops field if non-nil, zero value otherwise.

### GetMaxIopsOk

`func (o *MetadataSkuVolumeType) GetMaxIopsOk() (*float64, bool)`

GetMaxIopsOk returns a tuple with the MaxIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIops

`func (o *MetadataSkuVolumeType) SetMaxIops(v float64)`

SetMaxIops sets MaxIops field to given value.

### HasMaxIops

`func (o *MetadataSkuVolumeType) HasMaxIops() bool`

HasMaxIops returns a boolean if a field has been set.

### GetMaxReadIops

`func (o *MetadataSkuVolumeType) GetMaxReadIops() float64`

GetMaxReadIops returns the MaxReadIops field if non-nil, zero value otherwise.

### GetMaxReadIopsOk

`func (o *MetadataSkuVolumeType) GetMaxReadIopsOk() (*float64, bool)`

GetMaxReadIopsOk returns a tuple with the MaxReadIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReadIops

`func (o *MetadataSkuVolumeType) SetMaxReadIops(v float64)`

SetMaxReadIops sets MaxReadIops field to given value.

### HasMaxReadIops

`func (o *MetadataSkuVolumeType) HasMaxReadIops() bool`

HasMaxReadIops returns a boolean if a field has been set.

### GetMaxReadThroughput

`func (o *MetadataSkuVolumeType) GetMaxReadThroughput() float64`

GetMaxReadThroughput returns the MaxReadThroughput field if non-nil, zero value otherwise.

### GetMaxReadThroughputOk

`func (o *MetadataSkuVolumeType) GetMaxReadThroughputOk() (*float64, bool)`

GetMaxReadThroughputOk returns a tuple with the MaxReadThroughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReadThroughput

`func (o *MetadataSkuVolumeType) SetMaxReadThroughput(v float64)`

SetMaxReadThroughput sets MaxReadThroughput field to given value.

### HasMaxReadThroughput

`func (o *MetadataSkuVolumeType) HasMaxReadThroughput() bool`

HasMaxReadThroughput returns a boolean if a field has been set.

### GetMaxThroughput

`func (o *MetadataSkuVolumeType) GetMaxThroughput() float64`

GetMaxThroughput returns the MaxThroughput field if non-nil, zero value otherwise.

### GetMaxThroughputOk

`func (o *MetadataSkuVolumeType) GetMaxThroughputOk() (*float64, bool)`

GetMaxThroughputOk returns a tuple with the MaxThroughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxThroughput

`func (o *MetadataSkuVolumeType) SetMaxThroughput(v float64)`

SetMaxThroughput sets MaxThroughput field to given value.

### HasMaxThroughput

`func (o *MetadataSkuVolumeType) HasMaxThroughput() bool`

HasMaxThroughput returns a boolean if a field has been set.

### GetMaxVolumeSize

`func (o *MetadataSkuVolumeType) GetMaxVolumeSize() float64`

GetMaxVolumeSize returns the MaxVolumeSize field if non-nil, zero value otherwise.

### GetMaxVolumeSizeOk

`func (o *MetadataSkuVolumeType) GetMaxVolumeSizeOk() (*float64, bool)`

GetMaxVolumeSizeOk returns a tuple with the MaxVolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVolumeSize

`func (o *MetadataSkuVolumeType) SetMaxVolumeSize(v float64)`

SetMaxVolumeSize sets MaxVolumeSize field to given value.

### HasMaxVolumeSize

`func (o *MetadataSkuVolumeType) HasMaxVolumeSize() bool`

HasMaxVolumeSize returns a boolean if a field has been set.

### GetMaxWriteIops

`func (o *MetadataSkuVolumeType) GetMaxWriteIops() float64`

GetMaxWriteIops returns the MaxWriteIops field if non-nil, zero value otherwise.

### GetMaxWriteIopsOk

`func (o *MetadataSkuVolumeType) GetMaxWriteIopsOk() (*float64, bool)`

GetMaxWriteIopsOk returns a tuple with the MaxWriteIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWriteIops

`func (o *MetadataSkuVolumeType) SetMaxWriteIops(v float64)`

SetMaxWriteIops sets MaxWriteIops field to given value.

### HasMaxWriteIops

`func (o *MetadataSkuVolumeType) HasMaxWriteIops() bool`

HasMaxWriteIops returns a boolean if a field has been set.

### GetMaxWriteThroughput

`func (o *MetadataSkuVolumeType) GetMaxWriteThroughput() float64`

GetMaxWriteThroughput returns the MaxWriteThroughput field if non-nil, zero value otherwise.

### GetMaxWriteThroughputOk

`func (o *MetadataSkuVolumeType) GetMaxWriteThroughputOk() (*float64, bool)`

GetMaxWriteThroughputOk returns a tuple with the MaxWriteThroughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWriteThroughput

`func (o *MetadataSkuVolumeType) SetMaxWriteThroughput(v float64)`

SetMaxWriteThroughput sets MaxWriteThroughput field to given value.

### HasMaxWriteThroughput

`func (o *MetadataSkuVolumeType) HasMaxWriteThroughput() bool`

HasMaxWriteThroughput returns a boolean if a field has been set.

### GetMinVolumeSize

`func (o *MetadataSkuVolumeType) GetMinVolumeSize() float64`

GetMinVolumeSize returns the MinVolumeSize field if non-nil, zero value otherwise.

### GetMinVolumeSizeOk

`func (o *MetadataSkuVolumeType) GetMinVolumeSizeOk() (*float64, bool)`

GetMinVolumeSizeOk returns a tuple with the MinVolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinVolumeSize

`func (o *MetadataSkuVolumeType) SetMinVolumeSize(v float64)`

SetMinVolumeSize sets MinVolumeSize field to given value.

### HasMinVolumeSize

`func (o *MetadataSkuVolumeType) HasMinVolumeSize() bool`

HasMinVolumeSize returns a boolean if a field has been set.

### GetThroughputUnit

`func (o *MetadataSkuVolumeType) GetThroughputUnit() string`

GetThroughputUnit returns the ThroughputUnit field if non-nil, zero value otherwise.

### GetThroughputUnitOk

`func (o *MetadataSkuVolumeType) GetThroughputUnitOk() (*string, bool)`

GetThroughputUnitOk returns a tuple with the ThroughputUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughputUnit

`func (o *MetadataSkuVolumeType) SetThroughputUnit(v string)`

SetThroughputUnit sets ThroughputUnit field to given value.

### HasThroughputUnit

`func (o *MetadataSkuVolumeType) HasThroughputUnit() bool`

HasThroughputUnit returns a boolean if a field has been set.

### GetType

`func (o *MetadataSkuVolumeType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetadataSkuVolumeType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetadataSkuVolumeType) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MetadataSkuVolumeType) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVolumeSizeUnit

`func (o *MetadataSkuVolumeType) GetVolumeSizeUnit() string`

GetVolumeSizeUnit returns the VolumeSizeUnit field if non-nil, zero value otherwise.

### GetVolumeSizeUnitOk

`func (o *MetadataSkuVolumeType) GetVolumeSizeUnitOk() (*string, bool)`

GetVolumeSizeUnitOk returns a tuple with the VolumeSizeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSizeUnit

`func (o *MetadataSkuVolumeType) SetVolumeSizeUnit(v string)`

SetVolumeSizeUnit sets VolumeSizeUnit field to given value.

### HasVolumeSizeUnit

`func (o *MetadataSkuVolumeType) HasVolumeSizeUnit() bool`

HasVolumeSizeUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


