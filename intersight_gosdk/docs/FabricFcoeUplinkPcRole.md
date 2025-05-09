# FabricFcoeUplinkPcRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.FcoeUplinkPcRole"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.FcoeUplinkPcRole"]
**AdminSpeed** | Pointer to **string** | Admin configured speed for the port. * &#x60;Auto&#x60; - Admin configurable speed AUTO ( default ). * &#x60;1Gbps&#x60; - Admin configurable speed 1Gbps. * &#x60;10Gbps&#x60; - Admin configurable speed 10Gbps. * &#x60;25Gbps&#x60; - Admin configurable speed 25Gbps. * &#x60;40Gbps&#x60; - Admin configurable speed 40Gbps. * &#x60;100Gbps&#x60; - Admin configurable speed 100Gbps. * &#x60;NegAuto25Gbps&#x60; - Admin configurable 25Gbps auto negotiation for ports and port-channels.Speed is applicable on Ethernet Uplink, Ethernet Appliance and FCoE Uplink port and port-channel roles.This speed config is only applicable to non-breakout ports on UCS-FI-6454 and UCS-FI-64108. | [optional] [default to "Auto"]
**Fec** | Pointer to **string** | Forward error correction configuration for Fcoe Uplink Port Channel member ports. * &#x60;Auto&#x60; - Forward error correction option &#39;Auto&#39;. Supported speeds are Auto, 1Gbps, 10Gbps, 25Gbps, 40Gbps and 100 Gbps. * &#x60;Cl91&#x60; - Forward error correction option &#39;cl91&#39;. Supported speeds are 25Gbps and 100 Gbps. * &#x60;Cl74&#x60; - Forward error correction option &#39;cl74&#39;. Supported speeds are 25Gbps. * &#x60;rs-cons16&#x60; - Forward error correction option \&quot;rs-cons16\&quot;. Supported speeds are 25Gbps. * &#x60;rs-ieee&#x60; - Forward error correction option \&quot;rs-ieee\&quot;. Supported speeds are 25Gbps. * &#x60;Off&#x60; - Turn off forward error correction. Supported speeds are 25Gbps and 100 Gbps. | [optional] [default to "Auto"]
**LinkAggregationPolicy** | Pointer to [**NullableFabricLinkAggregationPolicyRelationship**](FabricLinkAggregationPolicyRelationship.md) |  | [optional] 
**LinkControlPolicy** | Pointer to [**NullableFabricLinkControlPolicyRelationship**](FabricLinkControlPolicyRelationship.md) |  | [optional] 

## Methods

### NewFabricFcoeUplinkPcRole

`func NewFabricFcoeUplinkPcRole(classId string, objectType string, ) *FabricFcoeUplinkPcRole`

NewFabricFcoeUplinkPcRole instantiates a new FabricFcoeUplinkPcRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricFcoeUplinkPcRoleWithDefaults

`func NewFabricFcoeUplinkPcRoleWithDefaults() *FabricFcoeUplinkPcRole`

NewFabricFcoeUplinkPcRoleWithDefaults instantiates a new FabricFcoeUplinkPcRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricFcoeUplinkPcRole) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricFcoeUplinkPcRole) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricFcoeUplinkPcRole) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricFcoeUplinkPcRole) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricFcoeUplinkPcRole) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricFcoeUplinkPcRole) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminSpeed

`func (o *FabricFcoeUplinkPcRole) GetAdminSpeed() string`

GetAdminSpeed returns the AdminSpeed field if non-nil, zero value otherwise.

### GetAdminSpeedOk

`func (o *FabricFcoeUplinkPcRole) GetAdminSpeedOk() (*string, bool)`

GetAdminSpeedOk returns a tuple with the AdminSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminSpeed

`func (o *FabricFcoeUplinkPcRole) SetAdminSpeed(v string)`

SetAdminSpeed sets AdminSpeed field to given value.

### HasAdminSpeed

`func (o *FabricFcoeUplinkPcRole) HasAdminSpeed() bool`

HasAdminSpeed returns a boolean if a field has been set.

### GetFec

`func (o *FabricFcoeUplinkPcRole) GetFec() string`

GetFec returns the Fec field if non-nil, zero value otherwise.

### GetFecOk

`func (o *FabricFcoeUplinkPcRole) GetFecOk() (*string, bool)`

GetFecOk returns a tuple with the Fec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFec

`func (o *FabricFcoeUplinkPcRole) SetFec(v string)`

SetFec sets Fec field to given value.

### HasFec

`func (o *FabricFcoeUplinkPcRole) HasFec() bool`

HasFec returns a boolean if a field has been set.

### GetLinkAggregationPolicy

`func (o *FabricFcoeUplinkPcRole) GetLinkAggregationPolicy() FabricLinkAggregationPolicyRelationship`

GetLinkAggregationPolicy returns the LinkAggregationPolicy field if non-nil, zero value otherwise.

### GetLinkAggregationPolicyOk

`func (o *FabricFcoeUplinkPcRole) GetLinkAggregationPolicyOk() (*FabricLinkAggregationPolicyRelationship, bool)`

GetLinkAggregationPolicyOk returns a tuple with the LinkAggregationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkAggregationPolicy

`func (o *FabricFcoeUplinkPcRole) SetLinkAggregationPolicy(v FabricLinkAggregationPolicyRelationship)`

SetLinkAggregationPolicy sets LinkAggregationPolicy field to given value.

### HasLinkAggregationPolicy

`func (o *FabricFcoeUplinkPcRole) HasLinkAggregationPolicy() bool`

HasLinkAggregationPolicy returns a boolean if a field has been set.

### SetLinkAggregationPolicyNil

`func (o *FabricFcoeUplinkPcRole) SetLinkAggregationPolicyNil(b bool)`

 SetLinkAggregationPolicyNil sets the value for LinkAggregationPolicy to be an explicit nil

### UnsetLinkAggregationPolicy
`func (o *FabricFcoeUplinkPcRole) UnsetLinkAggregationPolicy()`

UnsetLinkAggregationPolicy ensures that no value is present for LinkAggregationPolicy, not even an explicit nil
### GetLinkControlPolicy

`func (o *FabricFcoeUplinkPcRole) GetLinkControlPolicy() FabricLinkControlPolicyRelationship`

GetLinkControlPolicy returns the LinkControlPolicy field if non-nil, zero value otherwise.

### GetLinkControlPolicyOk

`func (o *FabricFcoeUplinkPcRole) GetLinkControlPolicyOk() (*FabricLinkControlPolicyRelationship, bool)`

GetLinkControlPolicyOk returns a tuple with the LinkControlPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkControlPolicy

`func (o *FabricFcoeUplinkPcRole) SetLinkControlPolicy(v FabricLinkControlPolicyRelationship)`

SetLinkControlPolicy sets LinkControlPolicy field to given value.

### HasLinkControlPolicy

`func (o *FabricFcoeUplinkPcRole) HasLinkControlPolicy() bool`

HasLinkControlPolicy returns a boolean if a field has been set.

### SetLinkControlPolicyNil

`func (o *FabricFcoeUplinkPcRole) SetLinkControlPolicyNil(b bool)`

 SetLinkControlPolicyNil sets the value for LinkControlPolicy to be an explicit nil

### UnsetLinkControlPolicy
`func (o *FabricFcoeUplinkPcRole) UnsetLinkControlPolicy()`

UnsetLinkControlPolicy ensures that no value is present for LinkControlPolicy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


