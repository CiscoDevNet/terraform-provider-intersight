# VirtualizationVmwareTeamingAndFailover

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VmwareTeamingAndFailover"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VmwareTeamingAndFailover"]
**ActiveAdapters** | Pointer to **[]string** |  | [optional] 
**Failback** | Pointer to **bool** | By default, a failback policy is enabled on a NIC team. If a failed physical NIC returns online, the network component sets the NIC back to active by replacing the standby NIC that took over its slot. | [optional] 
**LoadBalancing** | Pointer to **string** | Determines how network traffic is distributed between the network adapters in a NIC team. * &#x60;loadbalanceIP&#x60; - Load balance based on IP hash. * &#x60;loadbalanceSrcmac&#x60; - Route based on source MAC hash. * &#x60;loadbalanceSrcid&#x60; - Route based on originating virtual port. * &#x60;failoverExplicit&#x60; - Use explicit failover order. * &#x60;loadbalanceSrcnic&#x60; - Route based on physical NIC load. | [optional] [default to "loadbalanceIP"]
**Name** | Pointer to **string** | Name of the network component, example dvswitch, dvnetwork, vswitch or standard network. | [optional] 
**NetworkFailureDetection** | Pointer to **string** | Methods used by network component for failover detection. * &#x60;linkStatus&#x60; - This option detects failures such as removed cables and physical switch power failures. * &#x60;beaconProbing&#x60; - Sends out and listens for beacon probes on all NICs in the team, and uses this information, in addition to link status, to determine link failure. ESXi sends beacon packets every second. | [optional] [default to "linkStatus"]
**NotifySwitches** | Pointer to **bool** | Determines how network traffic is distributed between the network adapters in a NIC team. | [optional] 
**StandbyAdapters** | Pointer to **[]string** |  | [optional] 

## Methods

### NewVirtualizationVmwareTeamingAndFailover

`func NewVirtualizationVmwareTeamingAndFailover(classId string, objectType string, ) *VirtualizationVmwareTeamingAndFailover`

NewVirtualizationVmwareTeamingAndFailover instantiates a new VirtualizationVmwareTeamingAndFailover object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareTeamingAndFailoverWithDefaults

`func NewVirtualizationVmwareTeamingAndFailoverWithDefaults() *VirtualizationVmwareTeamingAndFailover`

NewVirtualizationVmwareTeamingAndFailoverWithDefaults instantiates a new VirtualizationVmwareTeamingAndFailover object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmwareTeamingAndFailover) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwareTeamingAndFailover) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwareTeamingAndFailover) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmwareTeamingAndFailover) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwareTeamingAndFailover) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwareTeamingAndFailover) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetActiveAdapters

`func (o *VirtualizationVmwareTeamingAndFailover) GetActiveAdapters() []string`

GetActiveAdapters returns the ActiveAdapters field if non-nil, zero value otherwise.

### GetActiveAdaptersOk

`func (o *VirtualizationVmwareTeamingAndFailover) GetActiveAdaptersOk() (*[]string, bool)`

GetActiveAdaptersOk returns a tuple with the ActiveAdapters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveAdapters

`func (o *VirtualizationVmwareTeamingAndFailover) SetActiveAdapters(v []string)`

SetActiveAdapters sets ActiveAdapters field to given value.

### HasActiveAdapters

`func (o *VirtualizationVmwareTeamingAndFailover) HasActiveAdapters() bool`

HasActiveAdapters returns a boolean if a field has been set.

### SetActiveAdaptersNil

`func (o *VirtualizationVmwareTeamingAndFailover) SetActiveAdaptersNil(b bool)`

 SetActiveAdaptersNil sets the value for ActiveAdapters to be an explicit nil

### UnsetActiveAdapters
`func (o *VirtualizationVmwareTeamingAndFailover) UnsetActiveAdapters()`

UnsetActiveAdapters ensures that no value is present for ActiveAdapters, not even an explicit nil
### GetFailback

`func (o *VirtualizationVmwareTeamingAndFailover) GetFailback() bool`

GetFailback returns the Failback field if non-nil, zero value otherwise.

### GetFailbackOk

`func (o *VirtualizationVmwareTeamingAndFailover) GetFailbackOk() (*bool, bool)`

GetFailbackOk returns a tuple with the Failback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailback

`func (o *VirtualizationVmwareTeamingAndFailover) SetFailback(v bool)`

SetFailback sets Failback field to given value.

### HasFailback

`func (o *VirtualizationVmwareTeamingAndFailover) HasFailback() bool`

HasFailback returns a boolean if a field has been set.

### GetLoadBalancing

`func (o *VirtualizationVmwareTeamingAndFailover) GetLoadBalancing() string`

GetLoadBalancing returns the LoadBalancing field if non-nil, zero value otherwise.

### GetLoadBalancingOk

`func (o *VirtualizationVmwareTeamingAndFailover) GetLoadBalancingOk() (*string, bool)`

GetLoadBalancingOk returns a tuple with the LoadBalancing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancing

`func (o *VirtualizationVmwareTeamingAndFailover) SetLoadBalancing(v string)`

SetLoadBalancing sets LoadBalancing field to given value.

### HasLoadBalancing

`func (o *VirtualizationVmwareTeamingAndFailover) HasLoadBalancing() bool`

HasLoadBalancing returns a boolean if a field has been set.

### GetName

`func (o *VirtualizationVmwareTeamingAndFailover) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationVmwareTeamingAndFailover) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationVmwareTeamingAndFailover) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationVmwareTeamingAndFailover) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkFailureDetection

`func (o *VirtualizationVmwareTeamingAndFailover) GetNetworkFailureDetection() string`

GetNetworkFailureDetection returns the NetworkFailureDetection field if non-nil, zero value otherwise.

### GetNetworkFailureDetectionOk

`func (o *VirtualizationVmwareTeamingAndFailover) GetNetworkFailureDetectionOk() (*string, bool)`

GetNetworkFailureDetectionOk returns a tuple with the NetworkFailureDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFailureDetection

`func (o *VirtualizationVmwareTeamingAndFailover) SetNetworkFailureDetection(v string)`

SetNetworkFailureDetection sets NetworkFailureDetection field to given value.

### HasNetworkFailureDetection

`func (o *VirtualizationVmwareTeamingAndFailover) HasNetworkFailureDetection() bool`

HasNetworkFailureDetection returns a boolean if a field has been set.

### GetNotifySwitches

`func (o *VirtualizationVmwareTeamingAndFailover) GetNotifySwitches() bool`

GetNotifySwitches returns the NotifySwitches field if non-nil, zero value otherwise.

### GetNotifySwitchesOk

`func (o *VirtualizationVmwareTeamingAndFailover) GetNotifySwitchesOk() (*bool, bool)`

GetNotifySwitchesOk returns a tuple with the NotifySwitches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifySwitches

`func (o *VirtualizationVmwareTeamingAndFailover) SetNotifySwitches(v bool)`

SetNotifySwitches sets NotifySwitches field to given value.

### HasNotifySwitches

`func (o *VirtualizationVmwareTeamingAndFailover) HasNotifySwitches() bool`

HasNotifySwitches returns a boolean if a field has been set.

### GetStandbyAdapters

`func (o *VirtualizationVmwareTeamingAndFailover) GetStandbyAdapters() []string`

GetStandbyAdapters returns the StandbyAdapters field if non-nil, zero value otherwise.

### GetStandbyAdaptersOk

`func (o *VirtualizationVmwareTeamingAndFailover) GetStandbyAdaptersOk() (*[]string, bool)`

GetStandbyAdaptersOk returns a tuple with the StandbyAdapters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandbyAdapters

`func (o *VirtualizationVmwareTeamingAndFailover) SetStandbyAdapters(v []string)`

SetStandbyAdapters sets StandbyAdapters field to given value.

### HasStandbyAdapters

`func (o *VirtualizationVmwareTeamingAndFailover) HasStandbyAdapters() bool`

HasStandbyAdapters returns a boolean if a field has been set.

### SetStandbyAdaptersNil

`func (o *VirtualizationVmwareTeamingAndFailover) SetStandbyAdaptersNil(b bool)`

 SetStandbyAdaptersNil sets the value for StandbyAdapters to be an explicit nil

### UnsetStandbyAdapters
`func (o *VirtualizationVmwareTeamingAndFailover) UnsetStandbyAdapters()`

UnsetStandbyAdapters ensures that no value is present for StandbyAdapters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


