# IppoolUniverseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "ippool.Universe"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "ippool.Universe"]
**Vrf** | Pointer to [**VrfVrfRelationship**](VrfVrfRelationship.md) |  | [optional] 

## Methods

### NewIppoolUniverseAllOf

`func NewIppoolUniverseAllOf(classId string, objectType string, ) *IppoolUniverseAllOf`

NewIppoolUniverseAllOf instantiates a new IppoolUniverseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIppoolUniverseAllOfWithDefaults

`func NewIppoolUniverseAllOfWithDefaults() *IppoolUniverseAllOf`

NewIppoolUniverseAllOfWithDefaults instantiates a new IppoolUniverseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *IppoolUniverseAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IppoolUniverseAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IppoolUniverseAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *IppoolUniverseAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IppoolUniverseAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IppoolUniverseAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetVrf

`func (o *IppoolUniverseAllOf) GetVrf() VrfVrfRelationship`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *IppoolUniverseAllOf) GetVrfOk() (*VrfVrfRelationship, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *IppoolUniverseAllOf) SetVrf(v VrfVrfRelationship)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *IppoolUniverseAllOf) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


