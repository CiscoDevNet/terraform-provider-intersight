# ComputeStoragePhysicalDriveAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "compute.StoragePhysicalDrive"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "compute.StoragePhysicalDrive"]
**SlotNumber** | Pointer to **int64** | Physical Drive ID of the storage on the server. | [optional] 

## Methods

### NewComputeStoragePhysicalDriveAllOf

`func NewComputeStoragePhysicalDriveAllOf(classId string, objectType string, ) *ComputeStoragePhysicalDriveAllOf`

NewComputeStoragePhysicalDriveAllOf instantiates a new ComputeStoragePhysicalDriveAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputeStoragePhysicalDriveAllOfWithDefaults

`func NewComputeStoragePhysicalDriveAllOfWithDefaults() *ComputeStoragePhysicalDriveAllOf`

NewComputeStoragePhysicalDriveAllOfWithDefaults instantiates a new ComputeStoragePhysicalDriveAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ComputeStoragePhysicalDriveAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ComputeStoragePhysicalDriveAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ComputeStoragePhysicalDriveAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ComputeStoragePhysicalDriveAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ComputeStoragePhysicalDriveAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ComputeStoragePhysicalDriveAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetSlotNumber

`func (o *ComputeStoragePhysicalDriveAllOf) GetSlotNumber() int64`

GetSlotNumber returns the SlotNumber field if non-nil, zero value otherwise.

### GetSlotNumberOk

`func (o *ComputeStoragePhysicalDriveAllOf) GetSlotNumberOk() (*int64, bool)`

GetSlotNumberOk returns a tuple with the SlotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotNumber

`func (o *ComputeStoragePhysicalDriveAllOf) SetSlotNumber(v int64)`

SetSlotNumber sets SlotNumber field to given value.

### HasSlotNumber

`func (o *ComputeStoragePhysicalDriveAllOf) HasSlotNumber() bool`

HasSlotNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


