# HyperflexHxPlatformDatastoreConfigDt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.HxPlatformDatastoreConfigDt"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.HxPlatformDatastoreConfigDt"]
**DataBlockSize** | Pointer to **int64** | Size of datablock for this HyperFlex datastore. | [optional] [readonly] 
**Name** | Pointer to **string** | Unique name for the datastore. | [optional] [readonly] 
**NumMirrors** | Pointer to **int64** | Number of copies of data maintained for data redundancy. | [optional] [readonly] 
**NumStripesForLargeFiles** | Pointer to **int64** | Number of stripes to be used for large files in datastore. | [optional] [readonly] 
**ProvisionedCapacity** | Pointer to **int64** | Provisioned capacity of datastore in bytes. | [optional] [readonly] 
**SystemDatastore** | Pointer to **bool** | Specifies if this datastore is a system datastore or not. | [optional] [readonly] 

## Methods

### NewHyperflexHxPlatformDatastoreConfigDt

`func NewHyperflexHxPlatformDatastoreConfigDt(classId string, objectType string, ) *HyperflexHxPlatformDatastoreConfigDt`

NewHyperflexHxPlatformDatastoreConfigDt instantiates a new HyperflexHxPlatformDatastoreConfigDt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHxPlatformDatastoreConfigDtWithDefaults

`func NewHyperflexHxPlatformDatastoreConfigDtWithDefaults() *HyperflexHxPlatformDatastoreConfigDt`

NewHyperflexHxPlatformDatastoreConfigDtWithDefaults instantiates a new HyperflexHxPlatformDatastoreConfigDt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHxPlatformDatastoreConfigDt) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHxPlatformDatastoreConfigDt) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHxPlatformDatastoreConfigDt) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHxPlatformDatastoreConfigDt) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHxPlatformDatastoreConfigDt) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHxPlatformDatastoreConfigDt) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDataBlockSize

`func (o *HyperflexHxPlatformDatastoreConfigDt) GetDataBlockSize() int64`

GetDataBlockSize returns the DataBlockSize field if non-nil, zero value otherwise.

### GetDataBlockSizeOk

`func (o *HyperflexHxPlatformDatastoreConfigDt) GetDataBlockSizeOk() (*int64, bool)`

GetDataBlockSizeOk returns a tuple with the DataBlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataBlockSize

`func (o *HyperflexHxPlatformDatastoreConfigDt) SetDataBlockSize(v int64)`

SetDataBlockSize sets DataBlockSize field to given value.

### HasDataBlockSize

`func (o *HyperflexHxPlatformDatastoreConfigDt) HasDataBlockSize() bool`

HasDataBlockSize returns a boolean if a field has been set.

### GetName

`func (o *HyperflexHxPlatformDatastoreConfigDt) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HyperflexHxPlatformDatastoreConfigDt) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HyperflexHxPlatformDatastoreConfigDt) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HyperflexHxPlatformDatastoreConfigDt) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumMirrors

`func (o *HyperflexHxPlatformDatastoreConfigDt) GetNumMirrors() int64`

GetNumMirrors returns the NumMirrors field if non-nil, zero value otherwise.

### GetNumMirrorsOk

`func (o *HyperflexHxPlatformDatastoreConfigDt) GetNumMirrorsOk() (*int64, bool)`

GetNumMirrorsOk returns a tuple with the NumMirrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumMirrors

`func (o *HyperflexHxPlatformDatastoreConfigDt) SetNumMirrors(v int64)`

SetNumMirrors sets NumMirrors field to given value.

### HasNumMirrors

`func (o *HyperflexHxPlatformDatastoreConfigDt) HasNumMirrors() bool`

HasNumMirrors returns a boolean if a field has been set.

### GetNumStripesForLargeFiles

`func (o *HyperflexHxPlatformDatastoreConfigDt) GetNumStripesForLargeFiles() int64`

GetNumStripesForLargeFiles returns the NumStripesForLargeFiles field if non-nil, zero value otherwise.

### GetNumStripesForLargeFilesOk

`func (o *HyperflexHxPlatformDatastoreConfigDt) GetNumStripesForLargeFilesOk() (*int64, bool)`

GetNumStripesForLargeFilesOk returns a tuple with the NumStripesForLargeFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumStripesForLargeFiles

`func (o *HyperflexHxPlatformDatastoreConfigDt) SetNumStripesForLargeFiles(v int64)`

SetNumStripesForLargeFiles sets NumStripesForLargeFiles field to given value.

### HasNumStripesForLargeFiles

`func (o *HyperflexHxPlatformDatastoreConfigDt) HasNumStripesForLargeFiles() bool`

HasNumStripesForLargeFiles returns a boolean if a field has been set.

### GetProvisionedCapacity

`func (o *HyperflexHxPlatformDatastoreConfigDt) GetProvisionedCapacity() int64`

GetProvisionedCapacity returns the ProvisionedCapacity field if non-nil, zero value otherwise.

### GetProvisionedCapacityOk

`func (o *HyperflexHxPlatformDatastoreConfigDt) GetProvisionedCapacityOk() (*int64, bool)`

GetProvisionedCapacityOk returns a tuple with the ProvisionedCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionedCapacity

`func (o *HyperflexHxPlatformDatastoreConfigDt) SetProvisionedCapacity(v int64)`

SetProvisionedCapacity sets ProvisionedCapacity field to given value.

### HasProvisionedCapacity

`func (o *HyperflexHxPlatformDatastoreConfigDt) HasProvisionedCapacity() bool`

HasProvisionedCapacity returns a boolean if a field has been set.

### GetSystemDatastore

`func (o *HyperflexHxPlatformDatastoreConfigDt) GetSystemDatastore() bool`

GetSystemDatastore returns the SystemDatastore field if non-nil, zero value otherwise.

### GetSystemDatastoreOk

`func (o *HyperflexHxPlatformDatastoreConfigDt) GetSystemDatastoreOk() (*bool, bool)`

GetSystemDatastoreOk returns a tuple with the SystemDatastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDatastore

`func (o *HyperflexHxPlatformDatastoreConfigDt) SetSystemDatastore(v bool)`

SetSystemDatastore sets SystemDatastore field to given value.

### HasSystemDatastore

`func (o *HyperflexHxPlatformDatastoreConfigDt) HasSystemDatastore() bool`

HasSystemDatastore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


