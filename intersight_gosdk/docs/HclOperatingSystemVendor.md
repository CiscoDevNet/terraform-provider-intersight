# HclOperatingSystemVendor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hcl.OperatingSystemVendor"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hcl.OperatingSystemVendor"]
**Name** | Pointer to **string** | Name of the vendor of the operating system. | [optional] 

## Methods

### NewHclOperatingSystemVendor

`func NewHclOperatingSystemVendor(classId string, objectType string, ) *HclOperatingSystemVendor`

NewHclOperatingSystemVendor instantiates a new HclOperatingSystemVendor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHclOperatingSystemVendorWithDefaults

`func NewHclOperatingSystemVendorWithDefaults() *HclOperatingSystemVendor`

NewHclOperatingSystemVendorWithDefaults instantiates a new HclOperatingSystemVendor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HclOperatingSystemVendor) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HclOperatingSystemVendor) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HclOperatingSystemVendor) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HclOperatingSystemVendor) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HclOperatingSystemVendor) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HclOperatingSystemVendor) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *HclOperatingSystemVendor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HclOperatingSystemVendor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HclOperatingSystemVendor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HclOperatingSystemVendor) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


