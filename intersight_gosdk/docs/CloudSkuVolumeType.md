# CloudSkuVolumeType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "cloud.SkuVolumeType"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "cloud.SkuVolumeType"]
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

### NewCloudSkuVolumeType

`func NewCloudSkuVolumeType(classId string, objectType string, ) *CloudSkuVolumeType`

NewCloudSkuVolumeType instantiates a new CloudSkuVolumeType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSkuVolumeTypeWithDefaults

`func NewCloudSkuVolumeTypeWithDefaults() *CloudSkuVolumeType`

NewCloudSkuVolumeTypeWithDefaults instantiates a new CloudSkuVolumeType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudSkuVolumeType) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudSkuVolumeType) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudSkuVolumeType) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudSkuVolumeType) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudSkuVolumeType) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudSkuVolumeType) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIopsUnit

`func (o *CloudSkuVolumeType) GetIopsUnit() string`

GetIopsUnit returns the IopsUnit field if non-nil, zero value otherwise.

### GetIopsUnitOk

`func (o *CloudSkuVolumeType) GetIopsUnitOk() (*string, bool)`

GetIopsUnitOk returns a tuple with the IopsUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIopsUnit

`func (o *CloudSkuVolumeType) SetIopsUnit(v string)`

SetIopsUnit sets IopsUnit field to given value.

### HasIopsUnit

`func (o *CloudSkuVolumeType) HasIopsUnit() bool`

HasIopsUnit returns a boolean if a field has been set.

### GetIsBootable

`func (o *CloudSkuVolumeType) GetIsBootable() bool`

GetIsBootable returns the IsBootable field if non-nil, zero value otherwise.

### GetIsBootableOk

`func (o *CloudSkuVolumeType) GetIsBootableOk() (*bool, bool)`

GetIsBootableOk returns a tuple with the IsBootable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBootable

`func (o *CloudSkuVolumeType) SetIsBootable(v bool)`

SetIsBootable sets IsBootable field to given value.

### HasIsBootable

`func (o *CloudSkuVolumeType) HasIsBootable() bool`

HasIsBootable returns a boolean if a field has been set.

### GetIsDefault

`func (o *CloudSkuVolumeType) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *CloudSkuVolumeType) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *CloudSkuVolumeType) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *CloudSkuVolumeType) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetIsProvisionedIops

`func (o *CloudSkuVolumeType) GetIsProvisionedIops() bool`

GetIsProvisionedIops returns the IsProvisionedIops field if non-nil, zero value otherwise.

### GetIsProvisionedIopsOk

`func (o *CloudSkuVolumeType) GetIsProvisionedIopsOk() (*bool, bool)`

GetIsProvisionedIopsOk returns a tuple with the IsProvisionedIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProvisionedIops

`func (o *CloudSkuVolumeType) SetIsProvisionedIops(v bool)`

SetIsProvisionedIops sets IsProvisionedIops field to given value.

### HasIsProvisionedIops

`func (o *CloudSkuVolumeType) HasIsProvisionedIops() bool`

HasIsProvisionedIops returns a boolean if a field has been set.

### GetMaxIops

`func (o *CloudSkuVolumeType) GetMaxIops() float64`

GetMaxIops returns the MaxIops field if non-nil, zero value otherwise.

### GetMaxIopsOk

`func (o *CloudSkuVolumeType) GetMaxIopsOk() (*float64, bool)`

GetMaxIopsOk returns a tuple with the MaxIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIops

`func (o *CloudSkuVolumeType) SetMaxIops(v float64)`

SetMaxIops sets MaxIops field to given value.

### HasMaxIops

`func (o *CloudSkuVolumeType) HasMaxIops() bool`

HasMaxIops returns a boolean if a field has been set.

### GetMaxReadIops

`func (o *CloudSkuVolumeType) GetMaxReadIops() float64`

GetMaxReadIops returns the MaxReadIops field if non-nil, zero value otherwise.

### GetMaxReadIopsOk

`func (o *CloudSkuVolumeType) GetMaxReadIopsOk() (*float64, bool)`

GetMaxReadIopsOk returns a tuple with the MaxReadIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReadIops

`func (o *CloudSkuVolumeType) SetMaxReadIops(v float64)`

SetMaxReadIops sets MaxReadIops field to given value.

### HasMaxReadIops

`func (o *CloudSkuVolumeType) HasMaxReadIops() bool`

HasMaxReadIops returns a boolean if a field has been set.

### GetMaxReadThroughput

`func (o *CloudSkuVolumeType) GetMaxReadThroughput() float64`

GetMaxReadThroughput returns the MaxReadThroughput field if non-nil, zero value otherwise.

### GetMaxReadThroughputOk

`func (o *CloudSkuVolumeType) GetMaxReadThroughputOk() (*float64, bool)`

GetMaxReadThroughputOk returns a tuple with the MaxReadThroughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReadThroughput

`func (o *CloudSkuVolumeType) SetMaxReadThroughput(v float64)`

SetMaxReadThroughput sets MaxReadThroughput field to given value.

### HasMaxReadThroughput

`func (o *CloudSkuVolumeType) HasMaxReadThroughput() bool`

HasMaxReadThroughput returns a boolean if a field has been set.

### GetMaxThroughput

`func (o *CloudSkuVolumeType) GetMaxThroughput() float64`

GetMaxThroughput returns the MaxThroughput field if non-nil, zero value otherwise.

### GetMaxThroughputOk

`func (o *CloudSkuVolumeType) GetMaxThroughputOk() (*float64, bool)`

GetMaxThroughputOk returns a tuple with the MaxThroughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxThroughput

`func (o *CloudSkuVolumeType) SetMaxThroughput(v float64)`

SetMaxThroughput sets MaxThroughput field to given value.

### HasMaxThroughput

`func (o *CloudSkuVolumeType) HasMaxThroughput() bool`

HasMaxThroughput returns a boolean if a field has been set.

### GetMaxVolumeSize

`func (o *CloudSkuVolumeType) GetMaxVolumeSize() float64`

GetMaxVolumeSize returns the MaxVolumeSize field if non-nil, zero value otherwise.

### GetMaxVolumeSizeOk

`func (o *CloudSkuVolumeType) GetMaxVolumeSizeOk() (*float64, bool)`

GetMaxVolumeSizeOk returns a tuple with the MaxVolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVolumeSize

`func (o *CloudSkuVolumeType) SetMaxVolumeSize(v float64)`

SetMaxVolumeSize sets MaxVolumeSize field to given value.

### HasMaxVolumeSize

`func (o *CloudSkuVolumeType) HasMaxVolumeSize() bool`

HasMaxVolumeSize returns a boolean if a field has been set.

### GetMaxWriteIops

`func (o *CloudSkuVolumeType) GetMaxWriteIops() float64`

GetMaxWriteIops returns the MaxWriteIops field if non-nil, zero value otherwise.

### GetMaxWriteIopsOk

`func (o *CloudSkuVolumeType) GetMaxWriteIopsOk() (*float64, bool)`

GetMaxWriteIopsOk returns a tuple with the MaxWriteIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWriteIops

`func (o *CloudSkuVolumeType) SetMaxWriteIops(v float64)`

SetMaxWriteIops sets MaxWriteIops field to given value.

### HasMaxWriteIops

`func (o *CloudSkuVolumeType) HasMaxWriteIops() bool`

HasMaxWriteIops returns a boolean if a field has been set.

### GetMaxWriteThroughput

`func (o *CloudSkuVolumeType) GetMaxWriteThroughput() float64`

GetMaxWriteThroughput returns the MaxWriteThroughput field if non-nil, zero value otherwise.

### GetMaxWriteThroughputOk

`func (o *CloudSkuVolumeType) GetMaxWriteThroughputOk() (*float64, bool)`

GetMaxWriteThroughputOk returns a tuple with the MaxWriteThroughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWriteThroughput

`func (o *CloudSkuVolumeType) SetMaxWriteThroughput(v float64)`

SetMaxWriteThroughput sets MaxWriteThroughput field to given value.

### HasMaxWriteThroughput

`func (o *CloudSkuVolumeType) HasMaxWriteThroughput() bool`

HasMaxWriteThroughput returns a boolean if a field has been set.

### GetMinVolumeSize

`func (o *CloudSkuVolumeType) GetMinVolumeSize() float64`

GetMinVolumeSize returns the MinVolumeSize field if non-nil, zero value otherwise.

### GetMinVolumeSizeOk

`func (o *CloudSkuVolumeType) GetMinVolumeSizeOk() (*float64, bool)`

GetMinVolumeSizeOk returns a tuple with the MinVolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinVolumeSize

`func (o *CloudSkuVolumeType) SetMinVolumeSize(v float64)`

SetMinVolumeSize sets MinVolumeSize field to given value.

### HasMinVolumeSize

`func (o *CloudSkuVolumeType) HasMinVolumeSize() bool`

HasMinVolumeSize returns a boolean if a field has been set.

### GetThroughputUnit

`func (o *CloudSkuVolumeType) GetThroughputUnit() string`

GetThroughputUnit returns the ThroughputUnit field if non-nil, zero value otherwise.

### GetThroughputUnitOk

`func (o *CloudSkuVolumeType) GetThroughputUnitOk() (*string, bool)`

GetThroughputUnitOk returns a tuple with the ThroughputUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughputUnit

`func (o *CloudSkuVolumeType) SetThroughputUnit(v string)`

SetThroughputUnit sets ThroughputUnit field to given value.

### HasThroughputUnit

`func (o *CloudSkuVolumeType) HasThroughputUnit() bool`

HasThroughputUnit returns a boolean if a field has been set.

### GetType

`func (o *CloudSkuVolumeType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudSkuVolumeType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudSkuVolumeType) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CloudSkuVolumeType) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVolumeSizeUnit

`func (o *CloudSkuVolumeType) GetVolumeSizeUnit() string`

GetVolumeSizeUnit returns the VolumeSizeUnit field if non-nil, zero value otherwise.

### GetVolumeSizeUnitOk

`func (o *CloudSkuVolumeType) GetVolumeSizeUnitOk() (*string, bool)`

GetVolumeSizeUnitOk returns a tuple with the VolumeSizeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSizeUnit

`func (o *CloudSkuVolumeType) SetVolumeSizeUnit(v string)`

SetVolumeSizeUnit sets VolumeSizeUnit field to given value.

### HasVolumeSizeUnit

`func (o *CloudSkuVolumeType) HasVolumeSizeUnit() bool`

HasVolumeSizeUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


