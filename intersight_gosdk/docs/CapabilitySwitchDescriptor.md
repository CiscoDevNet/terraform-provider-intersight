# CapabilitySwitchDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.SwitchDescriptor"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.SwitchDescriptor"]
**ExpectedMemory** | Pointer to **int64** | The total expected memory for this hardware. | [optional] 
**Revision** | Pointer to **string** | Revision for the fabric interconnect. | [optional] 

## Methods

### NewCapabilitySwitchDescriptor

`func NewCapabilitySwitchDescriptor(classId string, objectType string, ) *CapabilitySwitchDescriptor`

NewCapabilitySwitchDescriptor instantiates a new CapabilitySwitchDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilitySwitchDescriptorWithDefaults

`func NewCapabilitySwitchDescriptorWithDefaults() *CapabilitySwitchDescriptor`

NewCapabilitySwitchDescriptorWithDefaults instantiates a new CapabilitySwitchDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilitySwitchDescriptor) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilitySwitchDescriptor) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilitySwitchDescriptor) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilitySwitchDescriptor) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilitySwitchDescriptor) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilitySwitchDescriptor) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExpectedMemory

`func (o *CapabilitySwitchDescriptor) GetExpectedMemory() int64`

GetExpectedMemory returns the ExpectedMemory field if non-nil, zero value otherwise.

### GetExpectedMemoryOk

`func (o *CapabilitySwitchDescriptor) GetExpectedMemoryOk() (*int64, bool)`

GetExpectedMemoryOk returns a tuple with the ExpectedMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedMemory

`func (o *CapabilitySwitchDescriptor) SetExpectedMemory(v int64)`

SetExpectedMemory sets ExpectedMemory field to given value.

### HasExpectedMemory

`func (o *CapabilitySwitchDescriptor) HasExpectedMemory() bool`

HasExpectedMemory returns a boolean if a field has been set.

### GetRevision

`func (o *CapabilitySwitchDescriptor) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *CapabilitySwitchDescriptor) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *CapabilitySwitchDescriptor) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *CapabilitySwitchDescriptor) HasRevision() bool`

HasRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


