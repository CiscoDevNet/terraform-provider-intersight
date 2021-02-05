# CapabilityIoCardDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.IoCardDescriptor"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.IoCardDescriptor"]
**NumHifPorts** | Pointer to **int64** | Number of hif ports per blade for the iocard module. | [optional] 
**Revision** | Pointer to **string** | Revision for the iocard module. | [optional] 
**UifConnectivity** | Pointer to **string** | Connectivity information between UIF Uplink ports and IOM ports. * &#x60;inline&#x60; - UIF uplink ports and IOM ports are connected inline. * &#x60;cross-connected&#x60; - UIF uplink ports and IOM ports are cross-connected, a case in washington chassis. | [optional] [default to "inline"]

## Methods

### NewCapabilityIoCardDescriptor

`func NewCapabilityIoCardDescriptor(classId string, objectType string, ) *CapabilityIoCardDescriptor`

NewCapabilityIoCardDescriptor instantiates a new CapabilityIoCardDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityIoCardDescriptorWithDefaults

`func NewCapabilityIoCardDescriptorWithDefaults() *CapabilityIoCardDescriptor`

NewCapabilityIoCardDescriptorWithDefaults instantiates a new CapabilityIoCardDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityIoCardDescriptor) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityIoCardDescriptor) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityIoCardDescriptor) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityIoCardDescriptor) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityIoCardDescriptor) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityIoCardDescriptor) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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

### GetUifConnectivity

`func (o *CapabilityIoCardDescriptor) GetUifConnectivity() string`

GetUifConnectivity returns the UifConnectivity field if non-nil, zero value otherwise.

### GetUifConnectivityOk

`func (o *CapabilityIoCardDescriptor) GetUifConnectivityOk() (*string, bool)`

GetUifConnectivityOk returns a tuple with the UifConnectivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUifConnectivity

`func (o *CapabilityIoCardDescriptor) SetUifConnectivity(v string)`

SetUifConnectivity sets UifConnectivity field to given value.

### HasUifConnectivity

`func (o *CapabilityIoCardDescriptor) HasUifConnectivity() bool`

HasUifConnectivity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


