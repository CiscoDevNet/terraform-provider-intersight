# StorageAutomaticDriveGroupAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "storage.AutomaticDriveGroup"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "storage.AutomaticDriveGroup"]
**DriveType** | Pointer to **string** | Type of drive that should be used for this RAID group. * &#x60;Any&#x60; - Any type of drive can be used for virtual drive creation. * &#x60;HDD&#x60; - Hard disk drives should be used for virtual drive creation. * &#x60;SSD&#x60; - Solid state drives should be used for virtual drive creation. | [optional] [default to "Any"]
**DrivesPerSpan** | Pointer to **int64** | Number of drives within this span group. The minimum number of disks needed in a span group varies based on RAID level. RAID0 requires at least one disk. RAID1 and RAID10 requires at least 2 and in multiples of . RAID5 and RAID50 require at least 3 disks in a span group. RAID6 and RAID60 require atleast 4 disks in a span. | [optional] 
**MinimumDriveSize** | Pointer to **int64** | Minimum size of the drive to be used for creating this RAID group. | [optional] 
**NumDedicatedHotSpares** | Pointer to **string** | Number of dedicated hot spare disks for this RAID group. Allowed value is a comma or hyphen separated number range. | [optional] 
**NumberOfSpans** | Pointer to **int64** | Number of span groups to be created for this RAID group. Non-nested RAID levels have a single span. | [optional] [default to 0]
**UseRemainingDrives** | Pointer to **bool** | This flag enables the drive group to use all the remaining drives on the server. | [optional] 

## Methods

### NewStorageAutomaticDriveGroupAllOf

`func NewStorageAutomaticDriveGroupAllOf(classId string, objectType string, ) *StorageAutomaticDriveGroupAllOf`

NewStorageAutomaticDriveGroupAllOf instantiates a new StorageAutomaticDriveGroupAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageAutomaticDriveGroupAllOfWithDefaults

`func NewStorageAutomaticDriveGroupAllOfWithDefaults() *StorageAutomaticDriveGroupAllOf`

NewStorageAutomaticDriveGroupAllOfWithDefaults instantiates a new StorageAutomaticDriveGroupAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *StorageAutomaticDriveGroupAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *StorageAutomaticDriveGroupAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *StorageAutomaticDriveGroupAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *StorageAutomaticDriveGroupAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *StorageAutomaticDriveGroupAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *StorageAutomaticDriveGroupAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDriveType

`func (o *StorageAutomaticDriveGroupAllOf) GetDriveType() string`

GetDriveType returns the DriveType field if non-nil, zero value otherwise.

### GetDriveTypeOk

`func (o *StorageAutomaticDriveGroupAllOf) GetDriveTypeOk() (*string, bool)`

GetDriveTypeOk returns a tuple with the DriveType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveType

`func (o *StorageAutomaticDriveGroupAllOf) SetDriveType(v string)`

SetDriveType sets DriveType field to given value.

### HasDriveType

`func (o *StorageAutomaticDriveGroupAllOf) HasDriveType() bool`

HasDriveType returns a boolean if a field has been set.

### GetDrivesPerSpan

`func (o *StorageAutomaticDriveGroupAllOf) GetDrivesPerSpan() int64`

GetDrivesPerSpan returns the DrivesPerSpan field if non-nil, zero value otherwise.

### GetDrivesPerSpanOk

`func (o *StorageAutomaticDriveGroupAllOf) GetDrivesPerSpanOk() (*int64, bool)`

GetDrivesPerSpanOk returns a tuple with the DrivesPerSpan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivesPerSpan

`func (o *StorageAutomaticDriveGroupAllOf) SetDrivesPerSpan(v int64)`

SetDrivesPerSpan sets DrivesPerSpan field to given value.

### HasDrivesPerSpan

`func (o *StorageAutomaticDriveGroupAllOf) HasDrivesPerSpan() bool`

HasDrivesPerSpan returns a boolean if a field has been set.

### GetMinimumDriveSize

`func (o *StorageAutomaticDriveGroupAllOf) GetMinimumDriveSize() int64`

GetMinimumDriveSize returns the MinimumDriveSize field if non-nil, zero value otherwise.

### GetMinimumDriveSizeOk

`func (o *StorageAutomaticDriveGroupAllOf) GetMinimumDriveSizeOk() (*int64, bool)`

GetMinimumDriveSizeOk returns a tuple with the MinimumDriveSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumDriveSize

`func (o *StorageAutomaticDriveGroupAllOf) SetMinimumDriveSize(v int64)`

SetMinimumDriveSize sets MinimumDriveSize field to given value.

### HasMinimumDriveSize

`func (o *StorageAutomaticDriveGroupAllOf) HasMinimumDriveSize() bool`

HasMinimumDriveSize returns a boolean if a field has been set.

### GetNumDedicatedHotSpares

`func (o *StorageAutomaticDriveGroupAllOf) GetNumDedicatedHotSpares() string`

GetNumDedicatedHotSpares returns the NumDedicatedHotSpares field if non-nil, zero value otherwise.

### GetNumDedicatedHotSparesOk

`func (o *StorageAutomaticDriveGroupAllOf) GetNumDedicatedHotSparesOk() (*string, bool)`

GetNumDedicatedHotSparesOk returns a tuple with the NumDedicatedHotSpares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDedicatedHotSpares

`func (o *StorageAutomaticDriveGroupAllOf) SetNumDedicatedHotSpares(v string)`

SetNumDedicatedHotSpares sets NumDedicatedHotSpares field to given value.

### HasNumDedicatedHotSpares

`func (o *StorageAutomaticDriveGroupAllOf) HasNumDedicatedHotSpares() bool`

HasNumDedicatedHotSpares returns a boolean if a field has been set.

### GetNumberOfSpans

`func (o *StorageAutomaticDriveGroupAllOf) GetNumberOfSpans() int64`

GetNumberOfSpans returns the NumberOfSpans field if non-nil, zero value otherwise.

### GetNumberOfSpansOk

`func (o *StorageAutomaticDriveGroupAllOf) GetNumberOfSpansOk() (*int64, bool)`

GetNumberOfSpansOk returns a tuple with the NumberOfSpans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSpans

`func (o *StorageAutomaticDriveGroupAllOf) SetNumberOfSpans(v int64)`

SetNumberOfSpans sets NumberOfSpans field to given value.

### HasNumberOfSpans

`func (o *StorageAutomaticDriveGroupAllOf) HasNumberOfSpans() bool`

HasNumberOfSpans returns a boolean if a field has been set.

### GetUseRemainingDrives

`func (o *StorageAutomaticDriveGroupAllOf) GetUseRemainingDrives() bool`

GetUseRemainingDrives returns the UseRemainingDrives field if non-nil, zero value otherwise.

### GetUseRemainingDrivesOk

`func (o *StorageAutomaticDriveGroupAllOf) GetUseRemainingDrivesOk() (*bool, bool)`

GetUseRemainingDrivesOk returns a tuple with the UseRemainingDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRemainingDrives

`func (o *StorageAutomaticDriveGroupAllOf) SetUseRemainingDrives(v bool)`

SetUseRemainingDrives sets UseRemainingDrives field to given value.

### HasUseRemainingDrives

`func (o *StorageAutomaticDriveGroupAllOf) HasUseRemainingDrives() bool`

HasUseRemainingDrives returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


