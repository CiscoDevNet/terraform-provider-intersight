# VirtualizationVirtualDiskConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VirtualDiskConfig"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VirtualDiskConfig"]
**Capacity** | Pointer to **string** | Disk capacity to be provided with units example - 10Gi. | [optional] 
**Mode** | Pointer to **string** | File mode of the disk, example - Filesystem, Block. * &#x60;Block&#x60; - It is a Block virtual disk. * &#x60;Filesystem&#x60; - It is a File system virtual disk. * &#x60;&#x60; - Disk mode is either unknown or not supported. | [optional] [default to "Block"]
**SourceCerts** | Pointer to **string** | Base64 encoded CA certificates of the https source to check against. | [optional] 
**SourceDiskToClone** | Pointer to **string** | Source disk name from where the clone is done. | [optional] 
**SourceFilePath** | Pointer to **string** | Disk image source for the virtual machine. | [optional] 

## Methods

### NewVirtualizationVirtualDiskConfig

`func NewVirtualizationVirtualDiskConfig(classId string, objectType string, ) *VirtualizationVirtualDiskConfig`

NewVirtualizationVirtualDiskConfig instantiates a new VirtualizationVirtualDiskConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVirtualDiskConfigWithDefaults

`func NewVirtualizationVirtualDiskConfigWithDefaults() *VirtualizationVirtualDiskConfig`

NewVirtualizationVirtualDiskConfigWithDefaults instantiates a new VirtualizationVirtualDiskConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVirtualDiskConfig) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVirtualDiskConfig) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVirtualDiskConfig) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVirtualDiskConfig) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVirtualDiskConfig) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVirtualDiskConfig) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCapacity

`func (o *VirtualizationVirtualDiskConfig) GetCapacity() string`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *VirtualizationVirtualDiskConfig) GetCapacityOk() (*string, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *VirtualizationVirtualDiskConfig) SetCapacity(v string)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *VirtualizationVirtualDiskConfig) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetMode

`func (o *VirtualizationVirtualDiskConfig) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *VirtualizationVirtualDiskConfig) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *VirtualizationVirtualDiskConfig) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *VirtualizationVirtualDiskConfig) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetSourceCerts

`func (o *VirtualizationVirtualDiskConfig) GetSourceCerts() string`

GetSourceCerts returns the SourceCerts field if non-nil, zero value otherwise.

### GetSourceCertsOk

`func (o *VirtualizationVirtualDiskConfig) GetSourceCertsOk() (*string, bool)`

GetSourceCertsOk returns a tuple with the SourceCerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCerts

`func (o *VirtualizationVirtualDiskConfig) SetSourceCerts(v string)`

SetSourceCerts sets SourceCerts field to given value.

### HasSourceCerts

`func (o *VirtualizationVirtualDiskConfig) HasSourceCerts() bool`

HasSourceCerts returns a boolean if a field has been set.

### GetSourceDiskToClone

`func (o *VirtualizationVirtualDiskConfig) GetSourceDiskToClone() string`

GetSourceDiskToClone returns the SourceDiskToClone field if non-nil, zero value otherwise.

### GetSourceDiskToCloneOk

`func (o *VirtualizationVirtualDiskConfig) GetSourceDiskToCloneOk() (*string, bool)`

GetSourceDiskToCloneOk returns a tuple with the SourceDiskToClone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDiskToClone

`func (o *VirtualizationVirtualDiskConfig) SetSourceDiskToClone(v string)`

SetSourceDiskToClone sets SourceDiskToClone field to given value.

### HasSourceDiskToClone

`func (o *VirtualizationVirtualDiskConfig) HasSourceDiskToClone() bool`

HasSourceDiskToClone returns a boolean if a field has been set.

### GetSourceFilePath

`func (o *VirtualizationVirtualDiskConfig) GetSourceFilePath() string`

GetSourceFilePath returns the SourceFilePath field if non-nil, zero value otherwise.

### GetSourceFilePathOk

`func (o *VirtualizationVirtualDiskConfig) GetSourceFilePathOk() (*string, bool)`

GetSourceFilePathOk returns a tuple with the SourceFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFilePath

`func (o *VirtualizationVirtualDiskConfig) SetSourceFilePath(v string)`

SetSourceFilePath sets SourceFilePath field to given value.

### HasSourceFilePath

`func (o *VirtualizationVirtualDiskConfig) HasSourceFilePath() bool`

HasSourceFilePath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


