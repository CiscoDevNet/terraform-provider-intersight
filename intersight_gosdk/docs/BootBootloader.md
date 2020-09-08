# BootBootloader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Carries more information about the bootloader. | [optional] 
**Name** | Pointer to **string** | Name of the bootloader image. | [optional] 
**Path** | Pointer to **string** | Path to the bootloader image. | [optional] 

## Methods

### NewBootBootloader

`func NewBootBootloader() *BootBootloader`

NewBootBootloader instantiates a new BootBootloader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootBootloaderWithDefaults

`func NewBootBootloaderWithDefaults() *BootBootloader`

NewBootBootloaderWithDefaults instantiates a new BootBootloader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *BootBootloader) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BootBootloader) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BootBootloader) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BootBootloader) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *BootBootloader) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BootBootloader) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BootBootloader) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BootBootloader) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *BootBootloader) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *BootBootloader) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *BootBootloader) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *BootBootloader) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


