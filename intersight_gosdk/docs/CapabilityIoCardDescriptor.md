# CapabilityIoCardDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumHifPorts** | Pointer to **int64** | Number of hif ports per blade for the iocard module. | [optional] 
**Revision** | Pointer to **string** | Revision for the iocard module. | [optional] 

## Methods

### NewCapabilityIoCardDescriptor

`func NewCapabilityIoCardDescriptor() *CapabilityIoCardDescriptor`

NewCapabilityIoCardDescriptor instantiates a new CapabilityIoCardDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityIoCardDescriptorWithDefaults

`func NewCapabilityIoCardDescriptorWithDefaults() *CapabilityIoCardDescriptor`

NewCapabilityIoCardDescriptorWithDefaults instantiates a new CapabilityIoCardDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumHifPorts

`func (o *CapabilityIoCardDescriptor) GetNumHifPorts() int64`

GetNumHifPorts returns the NumHifPorts field if non-nil, zero value otherwise.

### GetNumHifPortsOk

`func (o *CapabilityIoCardDescriptor) GetNumHifPortsOk() (*int64, bool)`

GetNumHifPortsOk returns a tuple with the NumHifPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumHifPorts

`func (o *CapabilityIoCardDescriptor) SetNumHifPorts(v int64)`

SetNumHifPorts sets NumHifPorts field to given value.

### HasNumHifPorts

`func (o *CapabilityIoCardDescriptor) HasNumHifPorts() bool`

HasNumHifPorts returns a boolean if a field has been set.

### GetRevision

`func (o *CapabilityIoCardDescriptor) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *CapabilityIoCardDescriptor) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *CapabilityIoCardDescriptor) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *CapabilityIoCardDescriptor) HasRevision() bool`

HasRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


