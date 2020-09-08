# CapabilitySwitchDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpectedMemory** | Pointer to **int64** | The total expected memory for this hardware. | [optional] 
**Revision** | Pointer to **string** | Revision for the fabric interconnect. | [optional] 

## Methods

### NewCapabilitySwitchDescriptor

`func NewCapabilitySwitchDescriptor() *CapabilitySwitchDescriptor`

NewCapabilitySwitchDescriptor instantiates a new CapabilitySwitchDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilitySwitchDescriptorWithDefaults

`func NewCapabilitySwitchDescriptorWithDefaults() *CapabilitySwitchDescriptor`

NewCapabilitySwitchDescriptorWithDefaults instantiates a new CapabilitySwitchDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


