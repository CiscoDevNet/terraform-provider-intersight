# CapabilityCapability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Name** | Pointer to **string** | An unique identifer for a capability descriptor. | [optional] 

## Methods

### NewCapabilityCapability

`func NewCapabilityCapability(classId string, objectType string, ) *CapabilityCapability`

NewCapabilityCapability instantiates a new CapabilityCapability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityCapabilityWithDefaults

`func NewCapabilityCapabilityWithDefaults() *CapabilityCapability`

NewCapabilityCapabilityWithDefaults instantiates a new CapabilityCapability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityCapability) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityCapability) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityCapability) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityCapability) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityCapability) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityCapability) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *CapabilityCapability) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CapabilityCapability) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CapabilityCapability) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CapabilityCapability) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


