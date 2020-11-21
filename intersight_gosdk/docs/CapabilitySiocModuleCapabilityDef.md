# CapabilitySiocModuleCapabilityDef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.SiocModuleCapabilityDef"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.SiocModuleCapabilityDef"]
**DcSupported** | Pointer to **bool** | Device connector support on SIOC. | [optional] 

## Methods

### NewCapabilitySiocModuleCapabilityDef

`func NewCapabilitySiocModuleCapabilityDef(classId string, objectType string, ) *CapabilitySiocModuleCapabilityDef`

NewCapabilitySiocModuleCapabilityDef instantiates a new CapabilitySiocModuleCapabilityDef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilitySiocModuleCapabilityDefWithDefaults

`func NewCapabilitySiocModuleCapabilityDefWithDefaults() *CapabilitySiocModuleCapabilityDef`

NewCapabilitySiocModuleCapabilityDefWithDefaults instantiates a new CapabilitySiocModuleCapabilityDef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilitySiocModuleCapabilityDef) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilitySiocModuleCapabilityDef) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilitySiocModuleCapabilityDef) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilitySiocModuleCapabilityDef) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilitySiocModuleCapabilityDef) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilitySiocModuleCapabilityDef) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDcSupported

`func (o *CapabilitySiocModuleCapabilityDef) GetDcSupported() bool`

GetDcSupported returns the DcSupported field if non-nil, zero value otherwise.

### GetDcSupportedOk

`func (o *CapabilitySiocModuleCapabilityDef) GetDcSupportedOk() (*bool, bool)`

GetDcSupportedOk returns a tuple with the DcSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcSupported

`func (o *CapabilitySiocModuleCapabilityDef) SetDcSupported(v bool)`

SetDcSupported sets DcSupported field to given value.

### HasDcSupported

`func (o *CapabilitySiocModuleCapabilityDef) HasDcSupported() bool`

HasDcSupported returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


