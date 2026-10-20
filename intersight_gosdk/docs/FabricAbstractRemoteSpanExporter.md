# FabricAbstractRemoteSpanExporter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "fabric.RemoteSpanExporter"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "fabric.RemoteSpanExporter"]
**ConfigState** | Pointer to **string** | The configured state of Origin Interface configuration. If the configuration fails to deploy to the Fabric Interconnect, it can be redeployed by toggling the admin state. * &#x60;None&#x60; - Nil value when no action has been triggered by the user. * &#x60;Applied&#x60; - User configured settings are in applied state. * &#x60;Applying&#x60; - User settings are being applied on the target server. * &#x60;Failed&#x60; - User configured settings could not be applied. * &#x60;Scheduled&#x60; - User configured settings are scheduled to be applied. | [optional] [readonly] [default to "None"]
**Gateway** | Pointer to **string** | Source gateway that acts as the origin interface for ERSPAN tunnel frames on the Fabric Interconnect. | [optional] 
**IpAddressSource** | Pointer to **string** | Indicates whether the effective ERSPAN exporter IP address is sourced from this configuration or shared from a NetFlow exporter on the same VLAN. | [optional] [readonly] 
**OperState** | Pointer to **string** | Operational state of the Origin Interface. * &#x60;Unknown&#x60; - SPAN session operational state is Unknown. * &#x60;Up&#x60; - SPAN session operational state is Up. * &#x60;Down&#x60; - SPAN session operational state is Down. * &#x60;Error&#x60; - SPAN session operational state is Error. | [optional] [readonly] [default to "Unknown"]
**OperStateReason** | Pointer to **string** | Operational state reason of the SPAN session. * &#x60;Unknown&#x60; - Unknown operational state. * &#x60;Active&#x60; - Active and operational SPAN session. * &#x60;NoHardwareResource&#x60; - No hardware resources available. * &#x60;NoOperationalSrcDst&#x60; - No operational SPAN source or destination. * &#x60;GenericError&#x60; - Generic operational error. * &#x60;NoSourcesConfigured&#x60; - No source interfaces configured. * &#x60;NoDestinationConfigured&#x60; - No destination port configured. * &#x60;NoSourceDestinationConfigured&#x60; - No source or destination interface configured. * &#x60;SessionAdminShut&#x60; - Session is administratively disabled. * &#x60;WrongDestinationMode&#x60; - Wrong Destination mode configured. * &#x60;WrongSourceMode&#x60; - Wrong Source mode configured. * &#x60;TunnelMisconfDown&#x60; - Tunnel Misconfigured or Down. * &#x60;NoFlowIdSpecified&#x60; - No Flow ID specified for ERSPAN. * &#x60;NoErspanIdSpecified&#x60; - No ERSPAN ID is configured. * &#x60;MultiDestinationDenied&#x60; - Multiple destination is not allowed. * &#x60;VrfInvalid&#x60; - No valid VRF is configured. * &#x60;IpInvalid&#x60; - No valid destination IP address configured. * &#x60;GlobalIpInvalid&#x60; - No valid global origin IP address configured. * &#x60;NoRoute&#x60; - No route to destination IP address. * &#x60;ArpNotResolved&#x60; - ARP entry for destination is not resolved. * &#x60;NoEgressInterface&#x60; - Egress interface is not resolved. * &#x60;NoSviMember&#x60; - SVI member was not retrieved. | [optional] [readonly] [default to "Unknown"]
**OperationalIpAddress** | Pointer to **string** | Effective IP address programmed on the Fabric Interconnect for the ERSPAN exporter origin interface. When the exporter VLAN is shared with NetFlow, this reflects the NetFlow-owned SVI address; otherwise it matches sourceIp. | [optional] [readonly] 
**SourceIp** | Pointer to **string** | Source Interface IP Address that acts as the origin interface for ERSPAN tunnel frames on the Fabric Interconnect. | [optional] 
**SourceNetmask** | Pointer to **string** | Source Interface Netmask that acts as the origin interface for ERSPAN tunnel frames on the Fabric Interconnect. | [optional] 
**SourceVlanId** | Pointer to **int64** | Source Interface VLAN ID that acts as the origin interface for ERSPAN tunnel frames on the Fabric Interconnect. | [optional] 
**SwitchId** | Pointer to **string** | Switch ID (A or B) corresponding to network element. | [optional] 
**NetworkElement** | Pointer to [**NullableNetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 

## Methods

### NewFabricAbstractRemoteSpanExporter

`func NewFabricAbstractRemoteSpanExporter(classId string, objectType string, ) *FabricAbstractRemoteSpanExporter`

NewFabricAbstractRemoteSpanExporter instantiates a new FabricAbstractRemoteSpanExporter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricAbstractRemoteSpanExporterWithDefaults

`func NewFabricAbstractRemoteSpanExporterWithDefaults() *FabricAbstractRemoteSpanExporter`

NewFabricAbstractRemoteSpanExporterWithDefaults instantiates a new FabricAbstractRemoteSpanExporter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricAbstractRemoteSpanExporter) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricAbstractRemoteSpanExporter) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricAbstractRemoteSpanExporter) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricAbstractRemoteSpanExporter) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricAbstractRemoteSpanExporter) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricAbstractRemoteSpanExporter) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetConfigState

`func (o *FabricAbstractRemoteSpanExporter) GetConfigState() string`

GetConfigState returns the ConfigState field if non-nil, zero value otherwise.

### GetConfigStateOk

`func (o *FabricAbstractRemoteSpanExporter) GetConfigStateOk() (*string, bool)`

GetConfigStateOk returns a tuple with the ConfigState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigState

`func (o *FabricAbstractRemoteSpanExporter) SetConfigState(v string)`

SetConfigState sets ConfigState field to given value.

### HasConfigState

`func (o *FabricAbstractRemoteSpanExporter) HasConfigState() bool`

HasConfigState returns a boolean if a field has been set.

### GetGateway

`func (o *FabricAbstractRemoteSpanExporter) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *FabricAbstractRemoteSpanExporter) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *FabricAbstractRemoteSpanExporter) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *FabricAbstractRemoteSpanExporter) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetIpAddressSource

`func (o *FabricAbstractRemoteSpanExporter) GetIpAddressSource() string`

GetIpAddressSource returns the IpAddressSource field if non-nil, zero value otherwise.

### GetIpAddressSourceOk

`func (o *FabricAbstractRemoteSpanExporter) GetIpAddressSourceOk() (*string, bool)`

GetIpAddressSourceOk returns a tuple with the IpAddressSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddressSource

`func (o *FabricAbstractRemoteSpanExporter) SetIpAddressSource(v string)`

SetIpAddressSource sets IpAddressSource field to given value.

### HasIpAddressSource

`func (o *FabricAbstractRemoteSpanExporter) HasIpAddressSource() bool`

HasIpAddressSource returns a boolean if a field has been set.

### GetOperState

`func (o *FabricAbstractRemoteSpanExporter) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *FabricAbstractRemoteSpanExporter) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *FabricAbstractRemoteSpanExporter) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *FabricAbstractRemoteSpanExporter) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetOperStateReason

`func (o *FabricAbstractRemoteSpanExporter) GetOperStateReason() string`

GetOperStateReason returns the OperStateReason field if non-nil, zero value otherwise.

### GetOperStateReasonOk

`func (o *FabricAbstractRemoteSpanExporter) GetOperStateReasonOk() (*string, bool)`

GetOperStateReasonOk returns a tuple with the OperStateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperStateReason

`func (o *FabricAbstractRemoteSpanExporter) SetOperStateReason(v string)`

SetOperStateReason sets OperStateReason field to given value.

### HasOperStateReason

`func (o *FabricAbstractRemoteSpanExporter) HasOperStateReason() bool`

HasOperStateReason returns a boolean if a field has been set.

### GetOperationalIpAddress

`func (o *FabricAbstractRemoteSpanExporter) GetOperationalIpAddress() string`

GetOperationalIpAddress returns the OperationalIpAddress field if non-nil, zero value otherwise.

### GetOperationalIpAddressOk

`func (o *FabricAbstractRemoteSpanExporter) GetOperationalIpAddressOk() (*string, bool)`

GetOperationalIpAddressOk returns a tuple with the OperationalIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalIpAddress

`func (o *FabricAbstractRemoteSpanExporter) SetOperationalIpAddress(v string)`

SetOperationalIpAddress sets OperationalIpAddress field to given value.

### HasOperationalIpAddress

`func (o *FabricAbstractRemoteSpanExporter) HasOperationalIpAddress() bool`

HasOperationalIpAddress returns a boolean if a field has been set.

### GetSourceIp

`func (o *FabricAbstractRemoteSpanExporter) GetSourceIp() string`

GetSourceIp returns the SourceIp field if non-nil, zero value otherwise.

### GetSourceIpOk

`func (o *FabricAbstractRemoteSpanExporter) GetSourceIpOk() (*string, bool)`

GetSourceIpOk returns a tuple with the SourceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIp

`func (o *FabricAbstractRemoteSpanExporter) SetSourceIp(v string)`

SetSourceIp sets SourceIp field to given value.

### HasSourceIp

`func (o *FabricAbstractRemoteSpanExporter) HasSourceIp() bool`

HasSourceIp returns a boolean if a field has been set.

### GetSourceNetmask

`func (o *FabricAbstractRemoteSpanExporter) GetSourceNetmask() string`

GetSourceNetmask returns the SourceNetmask field if non-nil, zero value otherwise.

### GetSourceNetmaskOk

`func (o *FabricAbstractRemoteSpanExporter) GetSourceNetmaskOk() (*string, bool)`

GetSourceNetmaskOk returns a tuple with the SourceNetmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceNetmask

`func (o *FabricAbstractRemoteSpanExporter) SetSourceNetmask(v string)`

SetSourceNetmask sets SourceNetmask field to given value.

### HasSourceNetmask

`func (o *FabricAbstractRemoteSpanExporter) HasSourceNetmask() bool`

HasSourceNetmask returns a boolean if a field has been set.

### GetSourceVlanId

`func (o *FabricAbstractRemoteSpanExporter) GetSourceVlanId() int64`

GetSourceVlanId returns the SourceVlanId field if non-nil, zero value otherwise.

### GetSourceVlanIdOk

`func (o *FabricAbstractRemoteSpanExporter) GetSourceVlanIdOk() (*int64, bool)`

GetSourceVlanIdOk returns a tuple with the SourceVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVlanId

`func (o *FabricAbstractRemoteSpanExporter) SetSourceVlanId(v int64)`

SetSourceVlanId sets SourceVlanId field to given value.

### HasSourceVlanId

`func (o *FabricAbstractRemoteSpanExporter) HasSourceVlanId() bool`

HasSourceVlanId returns a boolean if a field has been set.

### GetSwitchId

`func (o *FabricAbstractRemoteSpanExporter) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *FabricAbstractRemoteSpanExporter) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *FabricAbstractRemoteSpanExporter) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *FabricAbstractRemoteSpanExporter) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetNetworkElement

`func (o *FabricAbstractRemoteSpanExporter) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *FabricAbstractRemoteSpanExporter) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *FabricAbstractRemoteSpanExporter) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *FabricAbstractRemoteSpanExporter) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### SetNetworkElementNil

`func (o *FabricAbstractRemoteSpanExporter) SetNetworkElementNil(b bool)`

 SetNetworkElementNil sets the value for NetworkElement to be an explicit nil

### UnsetNetworkElement
`func (o *FabricAbstractRemoteSpanExporter) UnsetNetworkElement()`

UnsetNetworkElement ensures that no value is present for NetworkElement, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


