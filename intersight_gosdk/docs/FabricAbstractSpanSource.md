# FabricAbstractSpanSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | 
**Direction** | Pointer to **string** | Direction of the source SPAN. * &#x60;Receive&#x60; - SPAN incoming traffic on the SPAN source interface. * &#x60;Transmit&#x60; - SPAN outgoing traffic on the SPAN source interface. * &#x60;Both&#x60; - SPAN incoming and outgoing traffic on the SPAN source interface. | [optional] [default to "Receive"]

## Methods

### NewFabricAbstractSpanSource

`func NewFabricAbstractSpanSource(classId string, objectType string, ) *FabricAbstractSpanSource`

NewFabricAbstractSpanSource instantiates a new FabricAbstractSpanSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricAbstractSpanSourceWithDefaults

`func NewFabricAbstractSpanSourceWithDefaults() *FabricAbstractSpanSource`

NewFabricAbstractSpanSourceWithDefaults instantiates a new FabricAbstractSpanSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricAbstractSpanSource) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricAbstractSpanSource) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricAbstractSpanSource) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricAbstractSpanSource) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricAbstractSpanSource) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricAbstractSpanSource) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDirection

`func (o *FabricAbstractSpanSource) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *FabricAbstractSpanSource) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *FabricAbstractSpanSource) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *FabricAbstractSpanSource) HasDirection() bool`

HasDirection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


