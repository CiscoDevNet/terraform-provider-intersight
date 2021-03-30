# HyperflexHxPlatformDatastoreConfigDtAllOf

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

### NewHyperflexHxPlatformDatastoreConfigDtAllOf

`func NewHyperflexHxPlatformDatastoreConfigDtAllOf(classId string, objectType string, ) *HyperflexHxPlatformDatastoreConfigDtAllOf`

NewHyperflexHxPlatformDatastoreConfigDtAllOf instantiates a new HyperflexHxPlatformDatastoreConfigDtAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHxPlatformDatastoreConfigDtAllOfWithDefaults

`func NewHyperflexHxPlatformDatastoreConfigDtAllOfWithDefaults() *HyperflexHxPlatformDatastoreConfigDtAllOf`

NewHyperflexHxPlatformDatastoreConfigDtAllOfWithDefaults instantiates a new HyperflexHxPlatformDatastoreConfigDtAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHxPlatformDatastoreConfigDtAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHxPlatformDatastoreConfigDtAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHxPlatformDatastoreConfigDtAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHxPlatformDatastoreConfigDtAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHxPlatformDatastoreConfigDtAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHxPlatformDatastoreConfigDtAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDataBlockSize

`func (o *HyperflexHxPlatformDatastoreConfigDtAllOf) GetDataBlockSize() int64`

GetDataBlockSize returns the DataBlockSize field if non-nil, zero value otherwise.

### GetDataBlockSizeOk

`func (o *HyperflexHxPlatformDatastoreConfigDtAllOf) GetDataBlockSizeOk() (*int64, bool)`

GetDataBlockSizeOk returns a tuple with the DataBlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataBlockSize

`func (o *HyperflexHxPlatformDatastoreConfigDtAllOf) SetDataBlockSize(v int64)`

SetDataBlockSize sets DataBlockSize field to given value.

### HasDataBlockSize

`func (o *HyperflexHxPlatformDatastoreConfigDtAllOf) HasDataBlockSize() bool`

HasDataBlockSize returns a boolean if a field has been set.

### GetName

`func (o *HyperflexHxPlatformDatastoreConfigDtAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HyperflexHxPlatformDatastoreConfigDtAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HyperflexHxPlatformDatastoreConfigDtAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HyperflexHxPlatformDatastoreConfigDtAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumMirrors

`func (o *HyperflexHxPlatformDatastoreConfigDtAllOf) GetNumMirrors() int64`

GetNumMirrors returns the NumMirrors field if non-nil, zero value otherwise.

### GetNumMirrorsOk

`func (o *HyperflexHxPlatformDatastoreConfigDtAllOf) GetNumMirrorsOk() (*int64, bool)`

GetNumMirrorsOk returns a tuple with the NumMirrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumMirrors

`func (o *HyperflexHxPlatformDatastoreConfigDtAllOf) SetNumMirrors(v int64)`

SetNumMirrors sets NumMirrors field to given value.

### HasNumMirrors

`func (o *HyperflexHxPlatformDatastoreConfigDtAllOf) HasNumMirrors() bool`

HasNumMirrors returns a boolean if a field has been set.

### GetNumStripesForLargeFiles

`func (o *HyperflexHxPlatformDatastoreConfigDtAllOf) GetNumStripesForLargeFiles() int64`

GetNumStripesForLargeFiles returns the NumStripesForLargeFiles field if non-nil, zero value otherwise.

### GetNumStripesForLargeFilesOk

`func (o *HyperflexHxPlatformDatastoreConfigDtAllOf) GetNumStripesForLargeFilesOk() (*int64, bool)`

GetNumStripesForLargeFilesOk returns a tuple with the NumStripesForLargeFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumStripesForLargeFiles

`func (o *HyperflexHxPlatformDatastoreConfigDtAllOf) SetNumStripesForLargeFiles(v int64)`

SetNumStripesForLargeFiles sets NumStripesForLargeFiles field to given value.

### HasNumStripesForLargeFiles

`func (o *HyperflexHxPlatformDatastoreConfigDtAllOf) HasNumStripesForLargeFiles() bool`

HasNumStripesForLargeFiles returns a boolean if a field has been set.

### GetProvisionedCapacity

`func (o *HyperflexHxPlatformDatastoreConfigDtAllOf) GetProvisionedCapacity() int64`

GetProvisionedCapacity returns the ProvisionedCapacity field if non-nil, zero value otherwise.

### GetProvisionedCapacityOk

`func (o *HyperflexHxPlatformDatastoreConfigDtAllOf) GetProvisionedCapacityOk() (*int64, bool)`

GetProvisionedCapacityOk returns a tuple with the ProvisionedCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionedCapacity

`func (o *HyperflexHxPlatformDatastoreConfigDtAllOf) SetProvisionedCapacity(v int64)`

SetProvisionedCapacity sets ProvisionedCapacity field to given value.

### HasProvisionedCapacity

`func (o *HyperflexHxPlatformDatastoreConfigDtAllOf) HasProvisionedCapacity() bool`

HasProvisionedCapacity returns a boolean if a field has been set.

### GetSystemDatastore

`func (o *HyperflexHxPlatformDatastoreConfigDtAllOf) GetSystemDatastore() bool`

GetSystemDatastore returns the SystemDatastore field if non-nil, zero value otherwise.

### GetSystemDatastoreOk

`func (o *HyperflexHxPlatformDatastoreConfigDtAllOf) GetSystemDatastoreOk() (*bool, bool)`

GetSystemDatastoreOk returns a tuple with the SystemDatastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDatastore

`func (o *HyperflexHxPlatformDatastoreConfigDtAllOf) SetSystemDatastore(v bool)`

SetSystemDatastore sets SystemDatastore field to given value.

### HasSystemDatastore

`func (o *HyperflexHxPlatformDatastoreConfigDtAllOf) HasSystemDatastore() bool`

HasSystemDatastore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


