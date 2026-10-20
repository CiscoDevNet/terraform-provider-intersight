# NiatelemetryRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.Route"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.Route"]
**FabricName** | Pointer to **string** | Fabric name associated with the route. | [optional] [readonly] 
**NextHops** | Pointer to [**[]NiatelemetryRouteNextHop**](NiatelemetryRouteNextHop.md) |  | [optional] 
**Route** | Pointer to **string** | Route prefix or destination for the static route. | [optional] [readonly] 
**RouteType** | Pointer to **string** | Route type reported by Nexus Dashboard. | [optional] [readonly] 
**SwitchName** | Pointer to **string** | Switch name associated with the static route. | [optional] [readonly] 
**Vrf** | Pointer to **string** | VRF name associated with the static route. | [optional] [readonly] 
**Fabric** | Pointer to [**NullableNiatelemetryFabricRelationship**](NiatelemetryFabricRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryRoute

`func NewNiatelemetryRoute(classId string, objectType string, ) *NiatelemetryRoute`

NewNiatelemetryRoute instantiates a new NiatelemetryRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryRouteWithDefaults

`func NewNiatelemetryRouteWithDefaults() *NiatelemetryRoute`

NewNiatelemetryRouteWithDefaults instantiates a new NiatelemetryRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryRoute) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryRoute) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryRoute) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryRoute) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryRoute) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryRoute) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFabricName

`func (o *NiatelemetryRoute) GetFabricName() string`

GetFabricName returns the FabricName field if non-nil, zero value otherwise.

### GetFabricNameOk

`func (o *NiatelemetryRoute) GetFabricNameOk() (*string, bool)`

GetFabricNameOk returns a tuple with the FabricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricName

`func (o *NiatelemetryRoute) SetFabricName(v string)`

SetFabricName sets FabricName field to given value.

### HasFabricName

`func (o *NiatelemetryRoute) HasFabricName() bool`

HasFabricName returns a boolean if a field has been set.

### GetNextHops

`func (o *NiatelemetryRoute) GetNextHops() []NiatelemetryRouteNextHop`

GetNextHops returns the NextHops field if non-nil, zero value otherwise.

### GetNextHopsOk

`func (o *NiatelemetryRoute) GetNextHopsOk() (*[]NiatelemetryRouteNextHop, bool)`

GetNextHopsOk returns a tuple with the NextHops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHops

`func (o *NiatelemetryRoute) SetNextHops(v []NiatelemetryRouteNextHop)`

SetNextHops sets NextHops field to given value.

### HasNextHops

`func (o *NiatelemetryRoute) HasNextHops() bool`

HasNextHops returns a boolean if a field has been set.

### SetNextHopsNil

`func (o *NiatelemetryRoute) SetNextHopsNil(b bool)`

 SetNextHopsNil sets the value for NextHops to be an explicit nil

### UnsetNextHops
`func (o *NiatelemetryRoute) UnsetNextHops()`

UnsetNextHops ensures that no value is present for NextHops, not even an explicit nil
### GetRoute

`func (o *NiatelemetryRoute) GetRoute() string`

GetRoute returns the Route field if non-nil, zero value otherwise.

### GetRouteOk

`func (o *NiatelemetryRoute) GetRouteOk() (*string, bool)`

GetRouteOk returns a tuple with the Route field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoute

`func (o *NiatelemetryRoute) SetRoute(v string)`

SetRoute sets Route field to given value.

### HasRoute

`func (o *NiatelemetryRoute) HasRoute() bool`

HasRoute returns a boolean if a field has been set.

### GetRouteType

`func (o *NiatelemetryRoute) GetRouteType() string`

GetRouteType returns the RouteType field if non-nil, zero value otherwise.

### GetRouteTypeOk

`func (o *NiatelemetryRoute) GetRouteTypeOk() (*string, bool)`

GetRouteTypeOk returns a tuple with the RouteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteType

`func (o *NiatelemetryRoute) SetRouteType(v string)`

SetRouteType sets RouteType field to given value.

### HasRouteType

`func (o *NiatelemetryRoute) HasRouteType() bool`

HasRouteType returns a boolean if a field has been set.

### GetSwitchName

`func (o *NiatelemetryRoute) GetSwitchName() string`

GetSwitchName returns the SwitchName field if non-nil, zero value otherwise.

### GetSwitchNameOk

`func (o *NiatelemetryRoute) GetSwitchNameOk() (*string, bool)`

GetSwitchNameOk returns a tuple with the SwitchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchName

`func (o *NiatelemetryRoute) SetSwitchName(v string)`

SetSwitchName sets SwitchName field to given value.

### HasSwitchName

`func (o *NiatelemetryRoute) HasSwitchName() bool`

HasSwitchName returns a boolean if a field has been set.

### GetVrf

`func (o *NiatelemetryRoute) GetVrf() string`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *NiatelemetryRoute) GetVrfOk() (*string, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *NiatelemetryRoute) SetVrf(v string)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *NiatelemetryRoute) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### GetFabric

`func (o *NiatelemetryRoute) GetFabric() NiatelemetryFabricRelationship`

GetFabric returns the Fabric field if non-nil, zero value otherwise.

### GetFabricOk

`func (o *NiatelemetryRoute) GetFabricOk() (*NiatelemetryFabricRelationship, bool)`

GetFabricOk returns a tuple with the Fabric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabric

`func (o *NiatelemetryRoute) SetFabric(v NiatelemetryFabricRelationship)`

SetFabric sets Fabric field to given value.

### HasFabric

`func (o *NiatelemetryRoute) HasFabric() bool`

HasFabric returns a boolean if a field has been set.

### SetFabricNil

`func (o *NiatelemetryRoute) SetFabricNil(b bool)`

 SetFabricNil sets the value for Fabric to be an explicit nil

### UnsetFabric
`func (o *NiatelemetryRoute) UnsetFabric()`

UnsetFabric ensures that no value is present for Fabric, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


