# CapabilityFanModuleDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.FanModuleDescriptor"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.FanModuleDescriptor"]
**Revision** | Pointer to **string** | Revision for the chassis enclosure. | [optional] 

## Methods

### NewCapabilityFanModuleDescriptor

`func NewCapabilityFanModuleDescriptor(classId string, objectType string, ) *CapabilityFanModuleDescriptor`

NewCapabilityFanModuleDescriptor instantiates a new CapabilityFanModuleDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityFanModuleDescriptorWithDefaults

`func NewCapabilityFanModuleDescriptorWithDefaults() *CapabilityFanModuleDescriptor`

NewCapabilityFanModuleDescriptorWithDefaults instantiates a new CapabilityFanModuleDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityFanModuleDescriptor) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityFanModuleDescriptor) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityFanModuleDescriptor) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityFanModuleDescriptor) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityFanModuleDescriptor) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityFanModuleDescriptor) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetRevision

`func (o *CapabilityFanModuleDescriptor) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *CapabilityFanModuleDescriptor) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *CapabilityFanModuleDescriptor) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *CapabilityFanModuleDescriptor) HasRevision() bool`

HasRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


