# VirtualizationBaseVirtualDisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "virtualization.VmwareVirtualDisk"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "virtualization.VmwareVirtualDisk"]
**Name** | Pointer to **string** | Name of the storage disk. Name must be unique within a Datastore. | [optional] 
**Size** | Pointer to **int64** | Disk size represented in bytes. | [optional] 

## Methods

### NewVirtualizationBaseVirtualDisk

`func NewVirtualizationBaseVirtualDisk(classId string, objectType string, ) *VirtualizationBaseVirtualDisk`

NewVirtualizationBaseVirtualDisk instantiates a new VirtualizationBaseVirtualDisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationBaseVirtualDiskWithDefaults

`func NewVirtualizationBaseVirtualDiskWithDefaults() *VirtualizationBaseVirtualDisk`

NewVirtualizationBaseVirtualDiskWithDefaults instantiates a new VirtualizationBaseVirtualDisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationBaseVirtualDisk) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationBaseVirtualDisk) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationBaseVirtualDisk) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationBaseVirtualDisk) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationBaseVirtualDisk) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationBaseVirtualDisk) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *VirtualizationBaseVirtualDisk) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationBaseVirtualDisk) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationBaseVirtualDisk) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationBaseVirtualDisk) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *VirtualizationBaseVirtualDisk) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VirtualizationBaseVirtualDisk) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VirtualizationBaseVirtualDisk) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *VirtualizationBaseVirtualDisk) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


