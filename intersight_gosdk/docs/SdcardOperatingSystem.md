# SdcardOperatingSystem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "sdcard.OperatingSystem"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "sdcard.OperatingSystem"]
**Name** | Pointer to **string** | Name of virtual drive for operating system partition. | [optional] [default to "Hypervisor"]

## Methods

### NewSdcardOperatingSystem

`func NewSdcardOperatingSystem(classId string, objectType string, ) *SdcardOperatingSystem`

NewSdcardOperatingSystem instantiates a new SdcardOperatingSystem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdcardOperatingSystemWithDefaults

`func NewSdcardOperatingSystemWithDefaults() *SdcardOperatingSystem`

NewSdcardOperatingSystemWithDefaults instantiates a new SdcardOperatingSystem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *SdcardOperatingSystem) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SdcardOperatingSystem) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SdcardOperatingSystem) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *SdcardOperatingSystem) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SdcardOperatingSystem) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SdcardOperatingSystem) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *SdcardOperatingSystem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SdcardOperatingSystem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SdcardOperatingSystem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SdcardOperatingSystem) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


