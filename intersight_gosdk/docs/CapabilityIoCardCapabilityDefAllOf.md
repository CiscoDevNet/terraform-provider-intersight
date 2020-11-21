# CapabilityIoCardCapabilityDefAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.IoCardCapabilityDef"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.IoCardCapabilityDef"]
**DcSupported** | Pointer to **bool** | Device connector support on Iocard. | [optional] 

## Methods

### NewCapabilityIoCardCapabilityDefAllOf

`func NewCapabilityIoCardCapabilityDefAllOf(classId string, objectType string, ) *CapabilityIoCardCapabilityDefAllOf`

NewCapabilityIoCardCapabilityDefAllOf instantiates a new CapabilityIoCardCapabilityDefAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityIoCardCapabilityDefAllOfWithDefaults

`func NewCapabilityIoCardCapabilityDefAllOfWithDefaults() *CapabilityIoCardCapabilityDefAllOf`

NewCapabilityIoCardCapabilityDefAllOfWithDefaults instantiates a new CapabilityIoCardCapabilityDefAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityIoCardCapabilityDefAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityIoCardCapabilityDefAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityIoCardCapabilityDefAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityIoCardCapabilityDefAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityIoCardCapabilityDefAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityIoCardCapabilityDefAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDcSupported

`func (o *CapabilityIoCardCapabilityDefAllOf) GetDcSupported() bool`

GetDcSupported returns the DcSupported field if non-nil, zero value otherwise.

### GetDcSupportedOk

`func (o *CapabilityIoCardCapabilityDefAllOf) GetDcSupportedOk() (*bool, bool)`

GetDcSupportedOk returns a tuple with the DcSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcSupported

`func (o *CapabilityIoCardCapabilityDefAllOf) SetDcSupported(v bool)`

SetDcSupported sets DcSupported field to given value.

### HasDcSupported

`func (o *CapabilityIoCardCapabilityDefAllOf) HasDcSupported() bool`

HasDcSupported returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


