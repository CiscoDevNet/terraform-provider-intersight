# FabricNetFlowExporter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.NetFlowExporter"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.NetFlowExporter"]
**Destination** | Pointer to **string** | Netflow collector IP address, The Netflow collector receives flow records from one or more exporters.  It processes the received export packets and stores the flow record information. | [optional] 
**Dscp** | Pointer to **int64** | DSCP value for export packets to ensure they receive proper QoS treatment. By default, NetFlow export packets may use the default DSCP value (usually 0, equivalent to Best Effort). | [optional] [default to 0]
**GatewayIp** | Pointer to **string** | Gateway IP address for the export interface network. | [optional] 
**Name** | Pointer to **string** | Name of netflow exporter. Must be a maximum of 31 characters, without spacing. | [optional] 
**OptionExporterStatsTimeout** | Pointer to **int64** | The time interval, in seconds, during which a NetFlow collector maintains an option template after it has been received from an exporter. An Option Template Record is a special type of template in NetFlow used to export metadata or control information, rather than flow data such as sampling parameters or exporter statistics. | [optional] [default to 600]
**OptionInterfaceTableTimeout** | Pointer to **int64** | The time interval, in seconds, during which a NetFlow collector maintains an interface option template after it has been received from an exporter. The optionInterfaceTable refers to an option data record exported by NetFlow exporters that provides metadata about network interfaces such as interface names, types, and speeds. | [optional] [default to 600]
**SourceInterface** | Pointer to [**NullableFabricVlanExportInterface**](FabricVlanExportInterface.md) |  | [optional] 
**TemplateDataTimeout** | Pointer to **int64** | The time interval, in seconds, during which a NetFlow collector maintains a template after it has been received from an exporter. templateData refers to the actual flow record data that is exported from a exporter to a collector, using a previously defined template. The template specifies the structure and the templateData provides the corresponding values for those fields for each flow. | [optional] [default to 600]
**UdpPort** | Pointer to **int64** | NetFlow export packets are encapsulated within UDP datagrams for transmission to the NetFlow collector. | [optional] [default to 1]
**Version** | Pointer to **int64** | Version of flow record format exported in exporter packet.  The value of this field is 9 for the current version. | [optional] [readonly] [default to 9]
**NetFlowPolicy** | Pointer to [**NullableFabricNetFlowPolicyRelationship**](FabricNetFlowPolicyRelationship.md) |  | [optional] 
**NetflowMonitors** | Pointer to [**[]FabricNetFlowMonitorRelationship**](FabricNetFlowMonitorRelationship.md) | An array of relationships to fabricNetFlowMonitor resources. | [optional] [readonly] 

## Methods

### NewFabricNetFlowExporter

`func NewFabricNetFlowExporter(classId string, objectType string, ) *FabricNetFlowExporter`

NewFabricNetFlowExporter instantiates a new FabricNetFlowExporter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricNetFlowExporterWithDefaults

`func NewFabricNetFlowExporterWithDefaults() *FabricNetFlowExporter`

NewFabricNetFlowExporterWithDefaults instantiates a new FabricNetFlowExporter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricNetFlowExporter) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricNetFlowExporter) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricNetFlowExporter) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricNetFlowExporter) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricNetFlowExporter) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricNetFlowExporter) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDestination

`func (o *FabricNetFlowExporter) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *FabricNetFlowExporter) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *FabricNetFlowExporter) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *FabricNetFlowExporter) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetDscp

`func (o *FabricNetFlowExporter) GetDscp() int64`

GetDscp returns the Dscp field if non-nil, zero value otherwise.

### GetDscpOk

`func (o *FabricNetFlowExporter) GetDscpOk() (*int64, bool)`

GetDscpOk returns a tuple with the Dscp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscp

`func (o *FabricNetFlowExporter) SetDscp(v int64)`

SetDscp sets Dscp field to given value.

### HasDscp

`func (o *FabricNetFlowExporter) HasDscp() bool`

HasDscp returns a boolean if a field has been set.

### GetGatewayIp

`func (o *FabricNetFlowExporter) GetGatewayIp() string`

GetGatewayIp returns the GatewayIp field if non-nil, zero value otherwise.

### GetGatewayIpOk

`func (o *FabricNetFlowExporter) GetGatewayIpOk() (*string, bool)`

GetGatewayIpOk returns a tuple with the GatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIp

`func (o *FabricNetFlowExporter) SetGatewayIp(v string)`

SetGatewayIp sets GatewayIp field to given value.

### HasGatewayIp

`func (o *FabricNetFlowExporter) HasGatewayIp() bool`

HasGatewayIp returns a boolean if a field has been set.

### GetName

`func (o *FabricNetFlowExporter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FabricNetFlowExporter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FabricNetFlowExporter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FabricNetFlowExporter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptionExporterStatsTimeout

`func (o *FabricNetFlowExporter) GetOptionExporterStatsTimeout() int64`

GetOptionExporterStatsTimeout returns the OptionExporterStatsTimeout field if non-nil, zero value otherwise.

### GetOptionExporterStatsTimeoutOk

`func (o *FabricNetFlowExporter) GetOptionExporterStatsTimeoutOk() (*int64, bool)`

GetOptionExporterStatsTimeoutOk returns a tuple with the OptionExporterStatsTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionExporterStatsTimeout

`func (o *FabricNetFlowExporter) SetOptionExporterStatsTimeout(v int64)`

SetOptionExporterStatsTimeout sets OptionExporterStatsTimeout field to given value.

### HasOptionExporterStatsTimeout

`func (o *FabricNetFlowExporter) HasOptionExporterStatsTimeout() bool`

HasOptionExporterStatsTimeout returns a boolean if a field has been set.

### GetOptionInterfaceTableTimeout

`func (o *FabricNetFlowExporter) GetOptionInterfaceTableTimeout() int64`

GetOptionInterfaceTableTimeout returns the OptionInterfaceTableTimeout field if non-nil, zero value otherwise.

### GetOptionInterfaceTableTimeoutOk

`func (o *FabricNetFlowExporter) GetOptionInterfaceTableTimeoutOk() (*int64, bool)`

GetOptionInterfaceTableTimeoutOk returns a tuple with the OptionInterfaceTableTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionInterfaceTableTimeout

`func (o *FabricNetFlowExporter) SetOptionInterfaceTableTimeout(v int64)`

SetOptionInterfaceTableTimeout sets OptionInterfaceTableTimeout field to given value.

### HasOptionInterfaceTableTimeout

`func (o *FabricNetFlowExporter) HasOptionInterfaceTableTimeout() bool`

HasOptionInterfaceTableTimeout returns a boolean if a field has been set.

### GetSourceInterface

`func (o *FabricNetFlowExporter) GetSourceInterface() FabricVlanExportInterface`

GetSourceInterface returns the SourceInterface field if non-nil, zero value otherwise.

### GetSourceInterfaceOk

`func (o *FabricNetFlowExporter) GetSourceInterfaceOk() (*FabricVlanExportInterface, bool)`

GetSourceInterfaceOk returns a tuple with the SourceInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInterface

`func (o *FabricNetFlowExporter) SetSourceInterface(v FabricVlanExportInterface)`

SetSourceInterface sets SourceInterface field to given value.

### HasSourceInterface

`func (o *FabricNetFlowExporter) HasSourceInterface() bool`

HasSourceInterface returns a boolean if a field has been set.

### SetSourceInterfaceNil

`func (o *FabricNetFlowExporter) SetSourceInterfaceNil(b bool)`

 SetSourceInterfaceNil sets the value for SourceInterface to be an explicit nil

### UnsetSourceInterface
`func (o *FabricNetFlowExporter) UnsetSourceInterface()`

UnsetSourceInterface ensures that no value is present for SourceInterface, not even an explicit nil
### GetTemplateDataTimeout

`func (o *FabricNetFlowExporter) GetTemplateDataTimeout() int64`

GetTemplateDataTimeout returns the TemplateDataTimeout field if non-nil, zero value otherwise.

### GetTemplateDataTimeoutOk

`func (o *FabricNetFlowExporter) GetTemplateDataTimeoutOk() (*int64, bool)`

GetTemplateDataTimeoutOk returns a tuple with the TemplateDataTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateDataTimeout

`func (o *FabricNetFlowExporter) SetTemplateDataTimeout(v int64)`

SetTemplateDataTimeout sets TemplateDataTimeout field to given value.

### HasTemplateDataTimeout

`func (o *FabricNetFlowExporter) HasTemplateDataTimeout() bool`

HasTemplateDataTimeout returns a boolean if a field has been set.

### GetUdpPort

`func (o *FabricNetFlowExporter) GetUdpPort() int64`

GetUdpPort returns the UdpPort field if non-nil, zero value otherwise.

### GetUdpPortOk

`func (o *FabricNetFlowExporter) GetUdpPortOk() (*int64, bool)`

GetUdpPortOk returns a tuple with the UdpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdpPort

`func (o *FabricNetFlowExporter) SetUdpPort(v int64)`

SetUdpPort sets UdpPort field to given value.

### HasUdpPort

`func (o *FabricNetFlowExporter) HasUdpPort() bool`

HasUdpPort returns a boolean if a field has been set.

### GetVersion

`func (o *FabricNetFlowExporter) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FabricNetFlowExporter) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FabricNetFlowExporter) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FabricNetFlowExporter) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetNetFlowPolicy

`func (o *FabricNetFlowExporter) GetNetFlowPolicy() FabricNetFlowPolicyRelationship`

GetNetFlowPolicy returns the NetFlowPolicy field if non-nil, zero value otherwise.

### GetNetFlowPolicyOk

`func (o *FabricNetFlowExporter) GetNetFlowPolicyOk() (*FabricNetFlowPolicyRelationship, bool)`

GetNetFlowPolicyOk returns a tuple with the NetFlowPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetFlowPolicy

`func (o *FabricNetFlowExporter) SetNetFlowPolicy(v FabricNetFlowPolicyRelationship)`

SetNetFlowPolicy sets NetFlowPolicy field to given value.

### HasNetFlowPolicy

`func (o *FabricNetFlowExporter) HasNetFlowPolicy() bool`

HasNetFlowPolicy returns a boolean if a field has been set.

### SetNetFlowPolicyNil

`func (o *FabricNetFlowExporter) SetNetFlowPolicyNil(b bool)`

 SetNetFlowPolicyNil sets the value for NetFlowPolicy to be an explicit nil

### UnsetNetFlowPolicy
`func (o *FabricNetFlowExporter) UnsetNetFlowPolicy()`

UnsetNetFlowPolicy ensures that no value is present for NetFlowPolicy, not even an explicit nil
### GetNetflowMonitors

`func (o *FabricNetFlowExporter) GetNetflowMonitors() []FabricNetFlowMonitorRelationship`

GetNetflowMonitors returns the NetflowMonitors field if non-nil, zero value otherwise.

### GetNetflowMonitorsOk

`func (o *FabricNetFlowExporter) GetNetflowMonitorsOk() (*[]FabricNetFlowMonitorRelationship, bool)`

GetNetflowMonitorsOk returns a tuple with the NetflowMonitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetflowMonitors

`func (o *FabricNetFlowExporter) SetNetflowMonitors(v []FabricNetFlowMonitorRelationship)`

SetNetflowMonitors sets NetflowMonitors field to given value.

### HasNetflowMonitors

`func (o *FabricNetFlowExporter) HasNetflowMonitors() bool`

HasNetflowMonitors returns a boolean if a field has been set.

### SetNetflowMonitorsNil

`func (o *FabricNetFlowExporter) SetNetflowMonitorsNil(b bool)`

 SetNetflowMonitorsNil sets the value for NetflowMonitors to be an explicit nil

### UnsetNetflowMonitors
`func (o *FabricNetFlowExporter) UnsetNetflowMonitors()`

UnsetNetflowMonitors ensures that no value is present for NetflowMonitors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


