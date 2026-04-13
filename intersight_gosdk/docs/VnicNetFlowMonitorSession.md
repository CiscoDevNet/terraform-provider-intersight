# VnicNetFlowMonitorSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "vnic.NetFlowMonitorSession"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "vnic.NetFlowMonitorSession"]
**FlowDirection** | Pointer to **string** | Direction of the flow to monitor (ingress, egress, or both). * &#x60;both&#x60; - Both direction for NetFlow monitor. * &#x60;rx&#x60; - Input direction for NetFlow monitor. * &#x60;tx&#x60; - Output direction for NetFlow monitor. | [optional] [default to "both"]
**MonitorName** | Pointer to **string** | Name of the NetFlow monitor to use for this vNIC. The NetFlow monitor must be pre-configured on the Fabric Interconnects. It is configured as part of the NetFlow policy attached to the Fabric Switch Cluster profile. | [optional] 

## Methods

### NewVnicNetFlowMonitorSession

`func NewVnicNetFlowMonitorSession(classId string, objectType string, ) *VnicNetFlowMonitorSession`

NewVnicNetFlowMonitorSession instantiates a new VnicNetFlowMonitorSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVnicNetFlowMonitorSessionWithDefaults

`func NewVnicNetFlowMonitorSessionWithDefaults() *VnicNetFlowMonitorSession`

NewVnicNetFlowMonitorSessionWithDefaults instantiates a new VnicNetFlowMonitorSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VnicNetFlowMonitorSession) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VnicNetFlowMonitorSession) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VnicNetFlowMonitorSession) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VnicNetFlowMonitorSession) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VnicNetFlowMonitorSession) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VnicNetFlowMonitorSession) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFlowDirection

`func (o *VnicNetFlowMonitorSession) GetFlowDirection() string`

GetFlowDirection returns the FlowDirection field if non-nil, zero value otherwise.

### GetFlowDirectionOk

`func (o *VnicNetFlowMonitorSession) GetFlowDirectionOk() (*string, bool)`

GetFlowDirectionOk returns a tuple with the FlowDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowDirection

`func (o *VnicNetFlowMonitorSession) SetFlowDirection(v string)`

SetFlowDirection sets FlowDirection field to given value.

### HasFlowDirection

`func (o *VnicNetFlowMonitorSession) HasFlowDirection() bool`

HasFlowDirection returns a boolean if a field has been set.

### GetMonitorName

`func (o *VnicNetFlowMonitorSession) GetMonitorName() string`

GetMonitorName returns the MonitorName field if non-nil, zero value otherwise.

### GetMonitorNameOk

`func (o *VnicNetFlowMonitorSession) GetMonitorNameOk() (*string, bool)`

GetMonitorNameOk returns a tuple with the MonitorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorName

`func (o *VnicNetFlowMonitorSession) SetMonitorName(v string)`

SetMonitorName sets MonitorName field to given value.

### HasMonitorName

`func (o *VnicNetFlowMonitorSession) HasMonitorName() bool`

HasMonitorName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


