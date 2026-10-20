# NiatelemetryRouteNextHop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.RouteNextHop"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.RouteNextHop"]
**Interface** | Pointer to **string** | Interface associated with the static route next hop. | [optional] [readonly] 
**NextHop** | Pointer to **string** | Next-hop address configured for the static route. | [optional] [readonly] 
**Vrf** | Pointer to **string** | VRF associated with the static route next hop. | [optional] [readonly] 

## Methods

### NewNiatelemetryRouteNextHop

`func NewNiatelemetryRouteNextHop(classId string, objectType string, ) *NiatelemetryRouteNextHop`

NewNiatelemetryRouteNextHop instantiates a new NiatelemetryRouteNextHop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryRouteNextHopWithDefaults

`func NewNiatelemetryRouteNextHopWithDefaults() *NiatelemetryRouteNextHop`

NewNiatelemetryRouteNextHopWithDefaults instantiates a new NiatelemetryRouteNextHop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryRouteNextHop) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryRouteNextHop) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryRouteNextHop) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryRouteNextHop) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryRouteNextHop) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryRouteNextHop) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetInterface

`func (o *NiatelemetryRouteNextHop) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *NiatelemetryRouteNextHop) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *NiatelemetryRouteNextHop) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *NiatelemetryRouteNextHop) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetNextHop

`func (o *NiatelemetryRouteNextHop) GetNextHop() string`

GetNextHop returns the NextHop field if non-nil, zero value otherwise.

### GetNextHopOk

`func (o *NiatelemetryRouteNextHop) GetNextHopOk() (*string, bool)`

GetNextHopOk returns a tuple with the NextHop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHop

`func (o *NiatelemetryRouteNextHop) SetNextHop(v string)`

SetNextHop sets NextHop field to given value.

### HasNextHop

`func (o *NiatelemetryRouteNextHop) HasNextHop() bool`

HasNextHop returns a boolean if a field has been set.

### GetVrf

`func (o *NiatelemetryRouteNextHop) GetVrf() string`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *NiatelemetryRouteNextHop) GetVrfOk() (*string, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *NiatelemetryRouteNextHop) SetVrf(v string)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *NiatelemetryRouteNextHop) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


