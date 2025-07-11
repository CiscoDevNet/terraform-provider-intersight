# VirtualizationVmwareAttachedResourceTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VmwareAttachedResourceTag"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VmwareAttachedResourceTag"]
**Category** | Pointer to **string** | The category of the tag assigned to the resource. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the tag assigned to the resource. | [optional] [readonly] 

## Methods

### NewVirtualizationVmwareAttachedResourceTag

`func NewVirtualizationVmwareAttachedResourceTag(classId string, objectType string, ) *VirtualizationVmwareAttachedResourceTag`

NewVirtualizationVmwareAttachedResourceTag instantiates a new VirtualizationVmwareAttachedResourceTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareAttachedResourceTagWithDefaults

`func NewVirtualizationVmwareAttachedResourceTagWithDefaults() *VirtualizationVmwareAttachedResourceTag`

NewVirtualizationVmwareAttachedResourceTagWithDefaults instantiates a new VirtualizationVmwareAttachedResourceTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmwareAttachedResourceTag) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwareAttachedResourceTag) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwareAttachedResourceTag) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmwareAttachedResourceTag) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwareAttachedResourceTag) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwareAttachedResourceTag) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCategory

`func (o *VirtualizationVmwareAttachedResourceTag) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *VirtualizationVmwareAttachedResourceTag) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *VirtualizationVmwareAttachedResourceTag) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *VirtualizationVmwareAttachedResourceTag) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetName

`func (o *VirtualizationVmwareAttachedResourceTag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationVmwareAttachedResourceTag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationVmwareAttachedResourceTag) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationVmwareAttachedResourceTag) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


