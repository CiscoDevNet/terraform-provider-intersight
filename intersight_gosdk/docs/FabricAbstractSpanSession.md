# FabricAbstractSpanSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "fabric.SpanSession"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "fabric.SpanSession"]
**AdminState** | Pointer to **string** | Admin state to enable or disable the SPAN session. * &#x60;Disabled&#x60; - Admin configured Disabled State. * &#x60;Enabled&#x60; - Admin configured Enabled State. | [optional] [default to "Disabled"]
**ConfigState** | Pointer to **string** | The configured state of SPAN configuration. If the configuration fails to deploy to the Fabric Interconnect, it can be redeployed by toggling the admin state. * &#x60;None&#x60; - Nil value when no action has been triggered by the user. * &#x60;Applied&#x60; - User configured settings are in applied state. * &#x60;Applying&#x60; - User settings are being applied on the target server. * &#x60;Failed&#x60; - User configured settings could not be applied. | [optional] [readonly] [default to "None"]
**Name** | Pointer to **string** | Name of the SPAN session. The name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.) or an underscore (_). | [optional] 
**OperState** | Pointer to **string** | Operational state of the SPAN session. * &#x60;Unknown&#x60; - SPAN session operational state is Unknown. * &#x60;Up&#x60; - SPAN session operational state is Up. * &#x60;Down&#x60; - SPAN session operational state is Down. * &#x60;Error&#x60; - SPAN session operational state is Error. | [optional] [readonly] [default to "Unknown"]
**OperStateReason** | Pointer to **string** | Operational state reason of the SPAN session. * &#x60;Unknown&#x60; - Unknown operational state. * &#x60;Active&#x60; - Active and operational SPAN session. * &#x60;NoHardwareResource&#x60; - No hardware resources available. * &#x60;NoOperationalSrcDst&#x60; - No operational SPAN source or destination. * &#x60;GenericError&#x60; - Generic operational error. * &#x60;NoSourcesConfigured&#x60; - No source interfaces configured. * &#x60;NoDestinationConfigured&#x60; - No destination port configured. * &#x60;NoSourceDestinationConfigured&#x60; - No source or destination interface configured. * &#x60;SessionAdminShut&#x60; - Session is administratively disabled. * &#x60;WrongDestinationMode&#x60; - Wrong Destination mode configured. * &#x60;WrongSourceMode&#x60; - Wrong Source mode configured. * &#x60;TunnelMisconfDown&#x60; - Tunnel Misconfigured or Down. * &#x60;NoFlowIdSpecified&#x60; - No Flow ID specified for ERSPAN. | [optional] [readonly] [default to "Unknown"]
**SessionId** | Pointer to **int64** | Session ID identifies the SPAN session on the Fabric Interconnect. | [optional] [readonly] 
**SourceCount** | Pointer to **int64** | Count of total number of sources in the SPAN session. | [optional] [readonly] 
**SpanControlPackets** | Pointer to **string** | Admin state to enable or disable spanning control packets. * &#x60;Disabled&#x60; - Admin configured Disabled State. * &#x60;Enabled&#x60; - Admin configured Enabled State. | [optional] [default to "Disabled"]
**SwitchId** | Pointer to **string** | Switch ID (A or B) corresponding to network element. * &#x60;A&#x60; - Switch Identifier of Fabric Interconnect A. * &#x60;B&#x60; - Switch Identifier of Fabric Interconnect B. | [optional] [readonly] [default to "A"]
**NetworkElement** | Pointer to [**NullableNetworkElementRelationship**](NetworkElementRelationship.md) |  | [optional] 

## Methods

### NewFabricAbstractSpanSession

`func NewFabricAbstractSpanSession(classId string, objectType string, ) *FabricAbstractSpanSession`

NewFabricAbstractSpanSession instantiates a new FabricAbstractSpanSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricAbstractSpanSessionWithDefaults

`func NewFabricAbstractSpanSessionWithDefaults() *FabricAbstractSpanSession`

NewFabricAbstractSpanSessionWithDefaults instantiates a new FabricAbstractSpanSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricAbstractSpanSession) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricAbstractSpanSession) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricAbstractSpanSession) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricAbstractSpanSession) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricAbstractSpanSession) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricAbstractSpanSession) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdminState

`func (o *FabricAbstractSpanSession) GetAdminState() string`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *FabricAbstractSpanSession) GetAdminStateOk() (*string, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *FabricAbstractSpanSession) SetAdminState(v string)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *FabricAbstractSpanSession) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetConfigState

`func (o *FabricAbstractSpanSession) GetConfigState() string`

GetConfigState returns the ConfigState field if non-nil, zero value otherwise.

### GetConfigStateOk

`func (o *FabricAbstractSpanSession) GetConfigStateOk() (*string, bool)`

GetConfigStateOk returns a tuple with the ConfigState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigState

`func (o *FabricAbstractSpanSession) SetConfigState(v string)`

SetConfigState sets ConfigState field to given value.

### HasConfigState

`func (o *FabricAbstractSpanSession) HasConfigState() bool`

HasConfigState returns a boolean if a field has been set.

### GetName

`func (o *FabricAbstractSpanSession) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FabricAbstractSpanSession) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FabricAbstractSpanSession) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FabricAbstractSpanSession) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperState

`func (o *FabricAbstractSpanSession) GetOperState() string`

GetOperState returns the OperState field if non-nil, zero value otherwise.

### GetOperStateOk

`func (o *FabricAbstractSpanSession) GetOperStateOk() (*string, bool)`

GetOperStateOk returns a tuple with the OperState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperState

`func (o *FabricAbstractSpanSession) SetOperState(v string)`

SetOperState sets OperState field to given value.

### HasOperState

`func (o *FabricAbstractSpanSession) HasOperState() bool`

HasOperState returns a boolean if a field has been set.

### GetOperStateReason

`func (o *FabricAbstractSpanSession) GetOperStateReason() string`

GetOperStateReason returns the OperStateReason field if non-nil, zero value otherwise.

### GetOperStateReasonOk

`func (o *FabricAbstractSpanSession) GetOperStateReasonOk() (*string, bool)`

GetOperStateReasonOk returns a tuple with the OperStateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperStateReason

`func (o *FabricAbstractSpanSession) SetOperStateReason(v string)`

SetOperStateReason sets OperStateReason field to given value.

### HasOperStateReason

`func (o *FabricAbstractSpanSession) HasOperStateReason() bool`

HasOperStateReason returns a boolean if a field has been set.

### GetSessionId

`func (o *FabricAbstractSpanSession) GetSessionId() int64`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *FabricAbstractSpanSession) GetSessionIdOk() (*int64, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *FabricAbstractSpanSession) SetSessionId(v int64)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *FabricAbstractSpanSession) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetSourceCount

`func (o *FabricAbstractSpanSession) GetSourceCount() int64`

GetSourceCount returns the SourceCount field if non-nil, zero value otherwise.

### GetSourceCountOk

`func (o *FabricAbstractSpanSession) GetSourceCountOk() (*int64, bool)`

GetSourceCountOk returns a tuple with the SourceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCount

`func (o *FabricAbstractSpanSession) SetSourceCount(v int64)`

SetSourceCount sets SourceCount field to given value.

### HasSourceCount

`func (o *FabricAbstractSpanSession) HasSourceCount() bool`

HasSourceCount returns a boolean if a field has been set.

### GetSpanControlPackets

`func (o *FabricAbstractSpanSession) GetSpanControlPackets() string`

GetSpanControlPackets returns the SpanControlPackets field if non-nil, zero value otherwise.

### GetSpanControlPacketsOk

`func (o *FabricAbstractSpanSession) GetSpanControlPacketsOk() (*string, bool)`

GetSpanControlPacketsOk returns a tuple with the SpanControlPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanControlPackets

`func (o *FabricAbstractSpanSession) SetSpanControlPackets(v string)`

SetSpanControlPackets sets SpanControlPackets field to given value.

### HasSpanControlPackets

`func (o *FabricAbstractSpanSession) HasSpanControlPackets() bool`

HasSpanControlPackets returns a boolean if a field has been set.

### GetSwitchId

`func (o *FabricAbstractSpanSession) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *FabricAbstractSpanSession) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *FabricAbstractSpanSession) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *FabricAbstractSpanSession) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetNetworkElement

`func (o *FabricAbstractSpanSession) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *FabricAbstractSpanSession) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *FabricAbstractSpanSession) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *FabricAbstractSpanSession) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.

### SetNetworkElementNil

`func (o *FabricAbstractSpanSession) SetNetworkElementNil(b bool)`

 SetNetworkElementNil sets the value for NetworkElement to be an explicit nil

### UnsetNetworkElement
`func (o *FabricAbstractSpanSession) UnsetNetworkElement()`

UnsetNetworkElement ensures that no value is present for NetworkElement, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


