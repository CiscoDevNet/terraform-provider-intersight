# CapabilityChassisDescriptorAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.ChassisDescriptor"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.ChassisDescriptor"]
**Revision** | Pointer to **string** | Revision for the chassis enclosure. | [optional] 

## Methods

### NewCapabilityChassisDescriptorAllOf

`func NewCapabilityChassisDescriptorAllOf(classId string, objectType string, ) *CapabilityChassisDescriptorAllOf`

NewCapabilityChassisDescriptorAllOf instantiates a new CapabilityChassisDescriptorAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityChassisDescriptorAllOfWithDefaults

`func NewCapabilityChassisDescriptorAllOfWithDefaults() *CapabilityChassisDescriptorAllOf`

NewCapabilityChassisDescriptorAllOfWithDefaults instantiates a new CapabilityChassisDescriptorAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityChassisDescriptorAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityChassisDescriptorAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityChassisDescriptorAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityChassisDescriptorAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityChassisDescriptorAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityChassisDescriptorAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetRevision

`func (o *CapabilityChassisDescriptorAllOf) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *CapabilityChassisDescriptorAllOf) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *CapabilityChassisDescriptorAllOf) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *CapabilityChassisDescriptorAllOf) HasRevision() bool`

HasRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


