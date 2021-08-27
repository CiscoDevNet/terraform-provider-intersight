# HyperflexHxapHostVswitchAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.HxapHostVswitch"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.HxapHostVswitch"]
**HostName** | Pointer to **string** | The name of the host to which this vSwitch belongs to. | [optional] 
**Ports** | Pointer to [**[]HyperflexNetworkPort**](HyperflexNetworkPort.md) |  | [optional] 
**Cluster** | Pointer to [**HyperflexHxapClusterRelationship**](HyperflexHxapClusterRelationship.md) |  | [optional] 
**DvSwitch** | Pointer to [**HyperflexHxapDvswitchRelationship**](HyperflexHxapDvswitchRelationship.md) |  | [optional] 
**Host** | Pointer to [**HyperflexHxapHostRelationship**](HyperflexHxapHostRelationship.md) |  | [optional] 

## Methods

### NewHyperflexHxapHostVswitchAllOf

`func NewHyperflexHxapHostVswitchAllOf(classId string, objectType string, ) *HyperflexHxapHostVswitchAllOf`

NewHyperflexHxapHostVswitchAllOf instantiates a new HyperflexHxapHostVswitchAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHxapHostVswitchAllOfWithDefaults

`func NewHyperflexHxapHostVswitchAllOfWithDefaults() *HyperflexHxapHostVswitchAllOf`

NewHyperflexHxapHostVswitchAllOfWithDefaults instantiates a new HyperflexHxapHostVswitchAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexHxapHostVswitchAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexHxapHostVswitchAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexHxapHostVswitchAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexHxapHostVswitchAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexHxapHostVswitchAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexHxapHostVswitchAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetHostName

`func (o *HyperflexHxapHostVswitchAllOf) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *HyperflexHxapHostVswitchAllOf) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *HyperflexHxapHostVswitchAllOf) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *HyperflexHxapHostVswitchAllOf) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetPorts

`func (o *HyperflexHxapHostVswitchAllOf) GetPorts() []HyperflexNetworkPort`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *HyperflexHxapHostVswitchAllOf) GetPortsOk() (*[]HyperflexNetworkPort, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *HyperflexHxapHostVswitchAllOf) SetPorts(v []HyperflexNetworkPort)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *HyperflexHxapHostVswitchAllOf) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### SetPortsNil

`func (o *HyperflexHxapHostVswitchAllOf) SetPortsNil(b bool)`

 SetPortsNil sets the value for Ports to be an explicit nil

### UnsetPorts
`func (o *HyperflexHxapHostVswitchAllOf) UnsetPorts()`

UnsetPorts ensures that no value is present for Ports, not even an explicit nil
### GetCluster

`func (o *HyperflexHxapHostVswitchAllOf) GetCluster() HyperflexHxapClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HyperflexHxapHostVswitchAllOf) GetClusterOk() (*HyperflexHxapClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HyperflexHxapHostVswitchAllOf) SetCluster(v HyperflexHxapClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HyperflexHxapHostVswitchAllOf) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetDvSwitch

`func (o *HyperflexHxapHostVswitchAllOf) GetDvSwitch() HyperflexHxapDvswitchRelationship`

GetDvSwitch returns the DvSwitch field if non-nil, zero value otherwise.

### GetDvSwitchOk

`func (o *HyperflexHxapHostVswitchAllOf) GetDvSwitchOk() (*HyperflexHxapDvswitchRelationship, bool)`

GetDvSwitchOk returns a tuple with the DvSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDvSwitch

`func (o *HyperflexHxapHostVswitchAllOf) SetDvSwitch(v HyperflexHxapDvswitchRelationship)`

SetDvSwitch sets DvSwitch field to given value.

### HasDvSwitch

`func (o *HyperflexHxapHostVswitchAllOf) HasDvSwitch() bool`

HasDvSwitch returns a boolean if a field has been set.

### GetHost

`func (o *HyperflexHxapHostVswitchAllOf) GetHost() HyperflexHxapHostRelationship`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HyperflexHxapHostVswitchAllOf) GetHostOk() (*HyperflexHxapHostRelationship, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HyperflexHxapHostVswitchAllOf) SetHost(v HyperflexHxapHostRelationship)`

SetHost sets Host field to given value.

### HasHost

`func (o *HyperflexHxapHostVswitchAllOf) HasHost() bool`

HasHost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


