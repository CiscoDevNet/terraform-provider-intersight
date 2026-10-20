# FabricAbstractSpanDestRemote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "fabric.SpanDestRemote"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "fabric.SpanDestRemote"]
**DestinationIp** | Pointer to **string** | IP address of the remote ERSPAN destination. | [optional] 
**ErspanId** | Pointer to **int64** | ERSPAN ID identifies the ERSPAN session on the Fabric Interconnect. | [optional] [default to 1]
**IpDscp** | Pointer to **int64** | Differentiated Services Code Point (DSCP) value for the ERSPAN tunnel IP packets. | [optional] [default to 0]
**IpTtl** | Pointer to **int64** | Time To Live (TTL) value for the ERSPAN tunnel IP packets. | [optional] [default to 1]
**Mtu** | Pointer to **int64** | Enables or disables packet truncation for the ERSPAN session. | [optional] [default to 1518]
**PacketTruncation** | Pointer to **bool** | Enables or disables packet truncation for the ERSPAN session. | [optional] 
**SessionId** | Pointer to **int64** | Session ID identifies the SPAN session on the Fabric Interconnect. | [optional] [readonly] 
**SwitchId** | Pointer to **string** | Switch ID (A or B) corresponding to network element. | [optional] [readonly] 
**SpanSession** | Pointer to [**NullableFabricSpanSessionRelationship**](FabricSpanSessionRelationship.md) |  | [optional] 

## Methods

### NewFabricAbstractSpanDestRemote

`func NewFabricAbstractSpanDestRemote(classId string, objectType string, ) *FabricAbstractSpanDestRemote`

NewFabricAbstractSpanDestRemote instantiates a new FabricAbstractSpanDestRemote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricAbstractSpanDestRemoteWithDefaults

`func NewFabricAbstractSpanDestRemoteWithDefaults() *FabricAbstractSpanDestRemote`

NewFabricAbstractSpanDestRemoteWithDefaults instantiates a new FabricAbstractSpanDestRemote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricAbstractSpanDestRemote) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricAbstractSpanDestRemote) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricAbstractSpanDestRemote) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricAbstractSpanDestRemote) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricAbstractSpanDestRemote) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricAbstractSpanDestRemote) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDestinationIp

`func (o *FabricAbstractSpanDestRemote) GetDestinationIp() string`

GetDestinationIp returns the DestinationIp field if non-nil, zero value otherwise.

### GetDestinationIpOk

`func (o *FabricAbstractSpanDestRemote) GetDestinationIpOk() (*string, bool)`

GetDestinationIpOk returns a tuple with the DestinationIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIp

`func (o *FabricAbstractSpanDestRemote) SetDestinationIp(v string)`

SetDestinationIp sets DestinationIp field to given value.

### HasDestinationIp

`func (o *FabricAbstractSpanDestRemote) HasDestinationIp() bool`

HasDestinationIp returns a boolean if a field has been set.

### GetErspanId

`func (o *FabricAbstractSpanDestRemote) GetErspanId() int64`

GetErspanId returns the ErspanId field if non-nil, zero value otherwise.

### GetErspanIdOk

`func (o *FabricAbstractSpanDestRemote) GetErspanIdOk() (*int64, bool)`

GetErspanIdOk returns a tuple with the ErspanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErspanId

`func (o *FabricAbstractSpanDestRemote) SetErspanId(v int64)`

SetErspanId sets ErspanId field to given value.

### HasErspanId

`func (o *FabricAbstractSpanDestRemote) HasErspanId() bool`

HasErspanId returns a boolean if a field has been set.

### GetIpDscp

`func (o *FabricAbstractSpanDestRemote) GetIpDscp() int64`

GetIpDscp returns the IpDscp field if non-nil, zero value otherwise.

### GetIpDscpOk

`func (o *FabricAbstractSpanDestRemote) GetIpDscpOk() (*int64, bool)`

GetIpDscpOk returns a tuple with the IpDscp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpDscp

`func (o *FabricAbstractSpanDestRemote) SetIpDscp(v int64)`

SetIpDscp sets IpDscp field to given value.

### HasIpDscp

`func (o *FabricAbstractSpanDestRemote) HasIpDscp() bool`

HasIpDscp returns a boolean if a field has been set.

### GetIpTtl

`func (o *FabricAbstractSpanDestRemote) GetIpTtl() int64`

GetIpTtl returns the IpTtl field if non-nil, zero value otherwise.

### GetIpTtlOk

`func (o *FabricAbstractSpanDestRemote) GetIpTtlOk() (*int64, bool)`

GetIpTtlOk returns a tuple with the IpTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpTtl

`func (o *FabricAbstractSpanDestRemote) SetIpTtl(v int64)`

SetIpTtl sets IpTtl field to given value.

### HasIpTtl

`func (o *FabricAbstractSpanDestRemote) HasIpTtl() bool`

HasIpTtl returns a boolean if a field has been set.

### GetMtu

`func (o *FabricAbstractSpanDestRemote) GetMtu() int64`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *FabricAbstractSpanDestRemote) GetMtuOk() (*int64, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *FabricAbstractSpanDestRemote) SetMtu(v int64)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *FabricAbstractSpanDestRemote) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetPacketTruncation

`func (o *FabricAbstractSpanDestRemote) GetPacketTruncation() bool`

GetPacketTruncation returns the PacketTruncation field if non-nil, zero value otherwise.

### GetPacketTruncationOk

`func (o *FabricAbstractSpanDestRemote) GetPacketTruncationOk() (*bool, bool)`

GetPacketTruncationOk returns a tuple with the PacketTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketTruncation

`func (o *FabricAbstractSpanDestRemote) SetPacketTruncation(v bool)`

SetPacketTruncation sets PacketTruncation field to given value.

### HasPacketTruncation

`func (o *FabricAbstractSpanDestRemote) HasPacketTruncation() bool`

HasPacketTruncation returns a boolean if a field has been set.

### GetSessionId

`func (o *FabricAbstractSpanDestRemote) GetSessionId() int64`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *FabricAbstractSpanDestRemote) GetSessionIdOk() (*int64, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *FabricAbstractSpanDestRemote) SetSessionId(v int64)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *FabricAbstractSpanDestRemote) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetSwitchId

`func (o *FabricAbstractSpanDestRemote) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *FabricAbstractSpanDestRemote) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *FabricAbstractSpanDestRemote) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *FabricAbstractSpanDestRemote) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetSpanSession

`func (o *FabricAbstractSpanDestRemote) GetSpanSession() FabricSpanSessionRelationship`

GetSpanSession returns the SpanSession field if non-nil, zero value otherwise.

### GetSpanSessionOk

`func (o *FabricAbstractSpanDestRemote) GetSpanSessionOk() (*FabricSpanSessionRelationship, bool)`

GetSpanSessionOk returns a tuple with the SpanSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanSession

`func (o *FabricAbstractSpanDestRemote) SetSpanSession(v FabricSpanSessionRelationship)`

SetSpanSession sets SpanSession field to given value.

### HasSpanSession

`func (o *FabricAbstractSpanDestRemote) HasSpanSession() bool`

HasSpanSession returns a boolean if a field has been set.

### SetSpanSessionNil

`func (o *FabricAbstractSpanDestRemote) SetSpanSessionNil(b bool)`

 SetSpanSessionNil sets the value for SpanSession to be an explicit nil

### UnsetSpanSession
`func (o *FabricAbstractSpanDestRemote) UnsetSpanSession()`

UnsetSpanSession ensures that no value is present for SpanSession, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


