# NiatelemetryInterfaceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.Interface"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.Interface"]
**InterfaceDownCount** | Pointer to **int64** | Return value of number of interafces down. | [optional] 
**InterfaceUpCount** | Pointer to **int64** | Return value of number of interafces up. | [optional] 

## Methods

### NewNiatelemetryInterfaceAllOf

`func NewNiatelemetryInterfaceAllOf(classId string, objectType string, ) *NiatelemetryInterfaceAllOf`

NewNiatelemetryInterfaceAllOf instantiates a new NiatelemetryInterfaceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryInterfaceAllOfWithDefaults

`func NewNiatelemetryInterfaceAllOfWithDefaults() *NiatelemetryInterfaceAllOf`

NewNiatelemetryInterfaceAllOfWithDefaults instantiates a new NiatelemetryInterfaceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryInterfaceAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryInterfaceAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryInterfaceAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryInterfaceAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryInterfaceAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryInterfaceAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetInterfaceDownCount

`func (o *NiatelemetryInterfaceAllOf) GetInterfaceDownCount() int64`

GetInterfaceDownCount returns the InterfaceDownCount field if non-nil, zero value otherwise.

### GetInterfaceDownCountOk

`func (o *NiatelemetryInterfaceAllOf) GetInterfaceDownCountOk() (*int64, bool)`

GetInterfaceDownCountOk returns a tuple with the InterfaceDownCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceDownCount

`func (o *NiatelemetryInterfaceAllOf) SetInterfaceDownCount(v int64)`

SetInterfaceDownCount sets InterfaceDownCount field to given value.

### HasInterfaceDownCount

`func (o *NiatelemetryInterfaceAllOf) HasInterfaceDownCount() bool`

HasInterfaceDownCount returns a boolean if a field has been set.

### GetInterfaceUpCount

`func (o *NiatelemetryInterfaceAllOf) GetInterfaceUpCount() int64`

GetInterfaceUpCount returns the InterfaceUpCount field if non-nil, zero value otherwise.

### GetInterfaceUpCountOk

`func (o *NiatelemetryInterfaceAllOf) GetInterfaceUpCountOk() (*int64, bool)`

GetInterfaceUpCountOk returns a tuple with the InterfaceUpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceUpCount

`func (o *NiatelemetryInterfaceAllOf) SetInterfaceUpCount(v int64)`

SetInterfaceUpCount sets InterfaceUpCount field to given value.

### HasInterfaceUpCount

`func (o *NiatelemetryInterfaceAllOf) HasInterfaceUpCount() bool`

HasInterfaceUpCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


