# VirtualizationIweHostVswitch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.IweHostVswitch"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.IweHostVswitch"]
**HostName** | Pointer to **string** | The name of the host to which this vSwitch belongs to. | [optional] 
**Ports** | Pointer to [**[]VirtualizationNetworkPort**](VirtualizationNetworkPort.md) |  | [optional] 
**Cluster** | Pointer to [**VirtualizationIweClusterRelationship**](VirtualizationIweClusterRelationship.md) |  | [optional] 
**DvSwitch** | Pointer to [**VirtualizationIweDvswitchRelationship**](VirtualizationIweDvswitchRelationship.md) |  | [optional] 
**Host** | Pointer to [**VirtualizationIweHostRelationship**](VirtualizationIweHostRelationship.md) |  | [optional] 

## Methods

### NewVirtualizationIweHostVswitch

`func NewVirtualizationIweHostVswitch(classId string, objectType string, ) *VirtualizationIweHostVswitch`

NewVirtualizationIweHostVswitch instantiates a new VirtualizationIweHostVswitch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationIweHostVswitchWithDefaults

`func NewVirtualizationIweHostVswitchWithDefaults() *VirtualizationIweHostVswitch`

NewVirtualizationIweHostVswitchWithDefaults instantiates a new VirtualizationIweHostVswitch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationIweHostVswitch) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationIweHostVswitch) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationIweHostVswitch) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationIweHostVswitch) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationIweHostVswitch) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationIweHostVswitch) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetHostName

`func (o *VirtualizationIweHostVswitch) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *VirtualizationIweHostVswitch) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *VirtualizationIweHostVswitch) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *VirtualizationIweHostVswitch) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetPorts

`func (o *VirtualizationIweHostVswitch) GetPorts() []VirtualizationNetworkPort`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *VirtualizationIweHostVswitch) GetPortsOk() (*[]VirtualizationNetworkPort, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *VirtualizationIweHostVswitch) SetPorts(v []VirtualizationNetworkPort)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *VirtualizationIweHostVswitch) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### SetPortsNil

`func (o *VirtualizationIweHostVswitch) SetPortsNil(b bool)`

 SetPortsNil sets the value for Ports to be an explicit nil

### UnsetPorts
`func (o *VirtualizationIweHostVswitch) UnsetPorts()`

UnsetPorts ensures that no value is present for Ports, not even an explicit nil
### GetCluster

`func (o *VirtualizationIweHostVswitch) GetCluster() VirtualizationIweClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VirtualizationIweHostVswitch) GetClusterOk() (*VirtualizationIweClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VirtualizationIweHostVswitch) SetCluster(v VirtualizationIweClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VirtualizationIweHostVswitch) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetDvSwitch

`func (o *VirtualizationIweHostVswitch) GetDvSwitch() VirtualizationIweDvswitchRelationship`

GetDvSwitch returns the DvSwitch field if non-nil, zero value otherwise.

### GetDvSwitchOk

`func (o *VirtualizationIweHostVswitch) GetDvSwitchOk() (*VirtualizationIweDvswitchRelationship, bool)`

GetDvSwitchOk returns a tuple with the DvSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDvSwitch

`func (o *VirtualizationIweHostVswitch) SetDvSwitch(v VirtualizationIweDvswitchRelationship)`

SetDvSwitch sets DvSwitch field to given value.

### HasDvSwitch

`func (o *VirtualizationIweHostVswitch) HasDvSwitch() bool`

HasDvSwitch returns a boolean if a field has been set.

### GetHost

`func (o *VirtualizationIweHostVswitch) GetHost() VirtualizationIweHostRelationship`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *VirtualizationIweHostVswitch) GetHostOk() (*VirtualizationIweHostRelationship, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *VirtualizationIweHostVswitch) SetHost(v VirtualizationIweHostRelationship)`

SetHost sets Host field to given value.

### HasHost

`func (o *VirtualizationIweHostVswitch) HasHost() bool`

HasHost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


