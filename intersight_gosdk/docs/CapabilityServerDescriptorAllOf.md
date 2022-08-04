# CapabilityServerDescriptorAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.ServerDescriptor"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.ServerDescriptor"]
**IsNcsiEnabled** | Pointer to **bool** | Indicates whether the CIMC to VIC side-band interface is enabled on the server. | [optional] 
**ServerFormFactor** | Pointer to **string** | The form factor (blade/rack/etc) of the server. * &#x60;unknown&#x60; - The form factor of the server is unknown. * &#x60;blade&#x60; - Blade server form factor. * &#x60;rack&#x60; - Rack unit server form factor. | [optional] [readonly] [default to "unknown"]

## Methods

### NewCapabilityServerDescriptorAllOf

`func NewCapabilityServerDescriptorAllOf(classId string, objectType string, ) *CapabilityServerDescriptorAllOf`

NewCapabilityServerDescriptorAllOf instantiates a new CapabilityServerDescriptorAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityServerDescriptorAllOfWithDefaults

`func NewCapabilityServerDescriptorAllOfWithDefaults() *CapabilityServerDescriptorAllOf`

NewCapabilityServerDescriptorAllOfWithDefaults instantiates a new CapabilityServerDescriptorAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityServerDescriptorAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityServerDescriptorAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityServerDescriptorAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityServerDescriptorAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityServerDescriptorAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityServerDescriptorAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsNcsiEnabled

`func (o *CapabilityServerDescriptorAllOf) GetIsNcsiEnabled() bool`

GetIsNcsiEnabled returns the IsNcsiEnabled field if non-nil, zero value otherwise.

### GetIsNcsiEnabledOk

`func (o *CapabilityServerDescriptorAllOf) GetIsNcsiEnabledOk() (*bool, bool)`

GetIsNcsiEnabledOk returns a tuple with the IsNcsiEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNcsiEnabled

`func (o *CapabilityServerDescriptorAllOf) SetIsNcsiEnabled(v bool)`

SetIsNcsiEnabled sets IsNcsiEnabled field to given value.

### HasIsNcsiEnabled

`func (o *CapabilityServerDescriptorAllOf) HasIsNcsiEnabled() bool`

HasIsNcsiEnabled returns a boolean if a field has been set.

### GetServerFormFactor

`func (o *CapabilityServerDescriptorAllOf) GetServerFormFactor() string`

GetServerFormFactor returns the ServerFormFactor field if non-nil, zero value otherwise.

### GetServerFormFactorOk

`func (o *CapabilityServerDescriptorAllOf) GetServerFormFactorOk() (*string, bool)`

GetServerFormFactorOk returns a tuple with the ServerFormFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFormFactor

`func (o *CapabilityServerDescriptorAllOf) SetServerFormFactor(v string)`

SetServerFormFactor sets ServerFormFactor field to given value.

### HasServerFormFactor

`func (o *CapabilityServerDescriptorAllOf) HasServerFormFactor() bool`

HasServerFormFactor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


