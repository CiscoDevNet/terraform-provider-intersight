# CapabilityFexCapabilityDef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.FexCapabilityDef"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.FexCapabilityDef"]
**FecConfigOnHifPortSupported** | Pointer to **bool** | FEC config on HIF port for Fabric Extender. | [optional] 

## Methods

### NewCapabilityFexCapabilityDef

`func NewCapabilityFexCapabilityDef(classId string, objectType string, ) *CapabilityFexCapabilityDef`

NewCapabilityFexCapabilityDef instantiates a new CapabilityFexCapabilityDef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityFexCapabilityDefWithDefaults

`func NewCapabilityFexCapabilityDefWithDefaults() *CapabilityFexCapabilityDef`

NewCapabilityFexCapabilityDefWithDefaults instantiates a new CapabilityFexCapabilityDef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityFexCapabilityDef) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityFexCapabilityDef) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityFexCapabilityDef) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityFexCapabilityDef) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityFexCapabilityDef) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityFexCapabilityDef) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFecConfigOnHifPortSupported

`func (o *CapabilityFexCapabilityDef) GetFecConfigOnHifPortSupported() bool`

GetFecConfigOnHifPortSupported returns the FecConfigOnHifPortSupported field if non-nil, zero value otherwise.

### GetFecConfigOnHifPortSupportedOk

`func (o *CapabilityFexCapabilityDef) GetFecConfigOnHifPortSupportedOk() (*bool, bool)`

GetFecConfigOnHifPortSupportedOk returns a tuple with the FecConfigOnHifPortSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFecConfigOnHifPortSupported

`func (o *CapabilityFexCapabilityDef) SetFecConfigOnHifPortSupported(v bool)`

SetFecConfigOnHifPortSupported sets FecConfigOnHifPortSupported field to given value.

### HasFecConfigOnHifPortSupported

`func (o *CapabilityFexCapabilityDef) HasFecConfigOnHifPortSupported() bool`

HasFecConfigOnHifPortSupported returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


