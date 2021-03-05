# ComputeStorageVirtualDrive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "compute.StorageVirtualDrive"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "compute.StorageVirtualDrive"]
**Id** | Pointer to **string** | Virtual Drive ID of the storage on the server. | [optional] 

## Methods

### NewComputeStorageVirtualDrive

`func NewComputeStorageVirtualDrive(classId string, objectType string, ) *ComputeStorageVirtualDrive`

NewComputeStorageVirtualDrive instantiates a new ComputeStorageVirtualDrive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeStorageVirtualDriveWithDefaults

`func NewComputeStorageVirtualDriveWithDefaults() *ComputeStorageVirtualDrive`

NewComputeStorageVirtualDriveWithDefaults instantiates a new ComputeStorageVirtualDrive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ComputeStorageVirtualDrive) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputeStorageVirtualDrive) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputeStorageVirtualDrive) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ComputeStorageVirtualDrive) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputeStorageVirtualDrive) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputeStorageVirtualDrive) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetId

`func (o *ComputeStorageVirtualDrive) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComputeStorageVirtualDrive) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComputeStorageVirtualDrive) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ComputeStorageVirtualDrive) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


