# NiatelemetryInterfaceElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.InterfaceElement"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.InterfaceElement"]
**Name** | Pointer to **string** | Return value of name of the port. | [optional] 
**OperState** | Pointer to **string** | Return value of operState attribute. | [optional] 
**XcvrPresent** | Pointer to **string** | Return whether sfp is present or not. | [optional] 

## Methods

### NewNiatelemetryInterfaceElement

`func NewNiatelemetryInterfaceElement(classId string, objectType string, ) *NiatelemetryInterfaceElement`

NewNiatelemetryInterfaceElement instantiates a new NiatelemetryInterfaceElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryInterfaceElementWithDefaults

`func NewNiatelemetryInterfaceElementWithDefaults() *NiatelemetryInterfaceElement`

NewNiatelemetryInterfaceElementWithDefaults instantiates a new NiatelemetryInterfaceElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryInterfaceElement) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryInterfaceElement) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryInterfaceElement) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryInterfaceElement) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryInterfaceElement) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryInterfaceElement) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *NiatelemetryInterfaceElement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NiatelemetryInterfaceElement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NiatelemetryInterfaceElement) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NiatelemetryInterfaceElement) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperState

`func (o *NiatelemetryInterfaceElement) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *NiatelemetryInterfaceElement) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *NiatelemetryInterfaceElement) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *NiatelemetryInterfaceElement) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetXcvrPresent

`func (o *NiatelemetryInterfaceElement) GetXcvrPresent() string`

GetXcvrPresent returns the XcvrPresent field if non-nil, zero value otherwise.

### GetXcvrPresentOk

`func (o *NiatelemetryInterfaceElement) GetXcvrPresentOk() (*string, bool)`

GetXcvrPresentOk returns a tuple with the XcvrPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXcvrPresent

`func (o *NiatelemetryInterfaceElement) SetXcvrPresent(v string)`

SetXcvrPresent sets XcvrPresent field to given value.

### HasXcvrPresent

`func (o *NiatelemetryInterfaceElement) HasXcvrPresent() bool`

HasXcvrPresent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


