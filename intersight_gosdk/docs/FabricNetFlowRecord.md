# FabricNetFlowRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.NetFlowRecord"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.NetFlowRecord"]
**FlowNonKey** | Pointer to [**NullableFabricFlowNonKey**](FabricFlowNonKey.md) |  | [optional] 
**Ipv4FlowKey** | Pointer to [**NullableFabricIpv4FlowKey**](FabricIpv4FlowKey.md) |  | [optional] 
**Ipv6FlowKey** | Pointer to [**NullableFabricIpv6FlowKey**](FabricIpv6FlowKey.md) |  | [optional] 
**L2FlowKey** | Pointer to [**NullableFabricL2FlowKey**](FabricL2FlowKey.md) |  | [optional] 
**Name** | Pointer to **string** | Netflow record name. Must be a maximum of 63 characters, without spacing. | [optional] 
**RecordType** | Pointer to **string** | Netflow Record Type IPv4, IPv6 and L2. * &#x60;Invalid&#x60; - Netflow record invlaid type. * &#x60;IPv4&#x60; - Netflow record type for IPv4 packet. * &#x60;IPv6&#x60; - Netflow record type for IPv6 packet. * &#x60;L2&#x60; - Netflow record type for L2 packet. | [optional] [default to "Invalid"]
**NetFlowPolicy** | Pointer to [**NullableFabricNetFlowPolicyRelationship**](FabricNetFlowPolicyRelationship.md) |  | [optional] 
**NetflowMonitors** | Pointer to [**[]FabricNetFlowMonitorRelationship**](FabricNetFlowMonitorRelationship.md) | An array of relationships to fabricNetFlowMonitor resources. | [optional] [readonly] 

## Methods

### NewFabricNetFlowRecord

`func NewFabricNetFlowRecord(classId string, objectType string, ) *FabricNetFlowRecord`

NewFabricNetFlowRecord instantiates a new FabricNetFlowRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricNetFlowRecordWithDefaults

`func NewFabricNetFlowRecordWithDefaults() *FabricNetFlowRecord`

NewFabricNetFlowRecordWithDefaults instantiates a new FabricNetFlowRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricNetFlowRecord) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricNetFlowRecord) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricNetFlowRecord) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricNetFlowRecord) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricNetFlowRecord) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricNetFlowRecord) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFlowNonKey

`func (o *FabricNetFlowRecord) GetFlowNonKey() FabricFlowNonKey`

GetFlowNonKey returns the FlowNonKey field if non-nil, zero value otherwise.

### GetFlowNonKeyOk

`func (o *FabricNetFlowRecord) GetFlowNonKeyOk() (*FabricFlowNonKey, bool)`

GetFlowNonKeyOk returns a tuple with the FlowNonKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowNonKey

`func (o *FabricNetFlowRecord) SetFlowNonKey(v FabricFlowNonKey)`

SetFlowNonKey sets FlowNonKey field to given value.

### HasFlowNonKey

`func (o *FabricNetFlowRecord) HasFlowNonKey() bool`

HasFlowNonKey returns a boolean if a field has been set.

### SetFlowNonKeyNil

`func (o *FabricNetFlowRecord) SetFlowNonKeyNil(b bool)`

 SetFlowNonKeyNil sets the value for FlowNonKey to be an explicit nil

### UnsetFlowNonKey
`func (o *FabricNetFlowRecord) UnsetFlowNonKey()`

UnsetFlowNonKey ensures that no value is present for FlowNonKey, not even an explicit nil
### GetIpv4FlowKey

`func (o *FabricNetFlowRecord) GetIpv4FlowKey() FabricIpv4FlowKey`

GetIpv4FlowKey returns the Ipv4FlowKey field if non-nil, zero value otherwise.

### GetIpv4FlowKeyOk

`func (o *FabricNetFlowRecord) GetIpv4FlowKeyOk() (*FabricIpv4FlowKey, bool)`

GetIpv4FlowKeyOk returns a tuple with the Ipv4FlowKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4FlowKey

`func (o *FabricNetFlowRecord) SetIpv4FlowKey(v FabricIpv4FlowKey)`

SetIpv4FlowKey sets Ipv4FlowKey field to given value.

### HasIpv4FlowKey

`func (o *FabricNetFlowRecord) HasIpv4FlowKey() bool`

HasIpv4FlowKey returns a boolean if a field has been set.

### SetIpv4FlowKeyNil

`func (o *FabricNetFlowRecord) SetIpv4FlowKeyNil(b bool)`

 SetIpv4FlowKeyNil sets the value for Ipv4FlowKey to be an explicit nil

### UnsetIpv4FlowKey
`func (o *FabricNetFlowRecord) UnsetIpv4FlowKey()`

UnsetIpv4FlowKey ensures that no value is present for Ipv4FlowKey, not even an explicit nil
### GetIpv6FlowKey

`func (o *FabricNetFlowRecord) GetIpv6FlowKey() FabricIpv6FlowKey`

GetIpv6FlowKey returns the Ipv6FlowKey field if non-nil, zero value otherwise.

### GetIpv6FlowKeyOk

`func (o *FabricNetFlowRecord) GetIpv6FlowKeyOk() (*FabricIpv6FlowKey, bool)`

GetIpv6FlowKeyOk returns a tuple with the Ipv6FlowKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6FlowKey

`func (o *FabricNetFlowRecord) SetIpv6FlowKey(v FabricIpv6FlowKey)`

SetIpv6FlowKey sets Ipv6FlowKey field to given value.

### HasIpv6FlowKey

`func (o *FabricNetFlowRecord) HasIpv6FlowKey() bool`

HasIpv6FlowKey returns a boolean if a field has been set.

### SetIpv6FlowKeyNil

`func (o *FabricNetFlowRecord) SetIpv6FlowKeyNil(b bool)`

 SetIpv6FlowKeyNil sets the value for Ipv6FlowKey to be an explicit nil

### UnsetIpv6FlowKey
`func (o *FabricNetFlowRecord) UnsetIpv6FlowKey()`

UnsetIpv6FlowKey ensures that no value is present for Ipv6FlowKey, not even an explicit nil
### GetL2FlowKey

`func (o *FabricNetFlowRecord) GetL2FlowKey() FabricL2FlowKey`

GetL2FlowKey returns the L2FlowKey field if non-nil, zero value otherwise.

### GetL2FlowKeyOk

`func (o *FabricNetFlowRecord) GetL2FlowKeyOk() (*FabricL2FlowKey, bool)`

GetL2FlowKeyOk returns a tuple with the L2FlowKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL2FlowKey

`func (o *FabricNetFlowRecord) SetL2FlowKey(v FabricL2FlowKey)`

SetL2FlowKey sets L2FlowKey field to given value.

### HasL2FlowKey

`func (o *FabricNetFlowRecord) HasL2FlowKey() bool`

HasL2FlowKey returns a boolean if a field has been set.

### SetL2FlowKeyNil

`func (o *FabricNetFlowRecord) SetL2FlowKeyNil(b bool)`

 SetL2FlowKeyNil sets the value for L2FlowKey to be an explicit nil

### UnsetL2FlowKey
`func (o *FabricNetFlowRecord) UnsetL2FlowKey()`

UnsetL2FlowKey ensures that no value is present for L2FlowKey, not even an explicit nil
### GetName

`func (o *FabricNetFlowRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FabricNetFlowRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FabricNetFlowRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FabricNetFlowRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRecordType

`func (o *FabricNetFlowRecord) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *FabricNetFlowRecord) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *FabricNetFlowRecord) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *FabricNetFlowRecord) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetNetFlowPolicy

`func (o *FabricNetFlowRecord) GetNetFlowPolicy() FabricNetFlowPolicyRelationship`

GetNetFlowPolicy returns the NetFlowPolicy field if non-nil, zero value otherwise.

### GetNetFlowPolicyOk

`func (o *FabricNetFlowRecord) GetNetFlowPolicyOk() (*FabricNetFlowPolicyRelationship, bool)`

GetNetFlowPolicyOk returns a tuple with the NetFlowPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetFlowPolicy

`func (o *FabricNetFlowRecord) SetNetFlowPolicy(v FabricNetFlowPolicyRelationship)`

SetNetFlowPolicy sets NetFlowPolicy field to given value.

### HasNetFlowPolicy

`func (o *FabricNetFlowRecord) HasNetFlowPolicy() bool`

HasNetFlowPolicy returns a boolean if a field has been set.

### SetNetFlowPolicyNil

`func (o *FabricNetFlowRecord) SetNetFlowPolicyNil(b bool)`

 SetNetFlowPolicyNil sets the value for NetFlowPolicy to be an explicit nil

### UnsetNetFlowPolicy
`func (o *FabricNetFlowRecord) UnsetNetFlowPolicy()`

UnsetNetFlowPolicy ensures that no value is present for NetFlowPolicy, not even an explicit nil
### GetNetflowMonitors

`func (o *FabricNetFlowRecord) GetNetflowMonitors() []FabricNetFlowMonitorRelationship`

GetNetflowMonitors returns the NetflowMonitors field if non-nil, zero value otherwise.

### GetNetflowMonitorsOk

`func (o *FabricNetFlowRecord) GetNetflowMonitorsOk() (*[]FabricNetFlowMonitorRelationship, bool)`

GetNetflowMonitorsOk returns a tuple with the NetflowMonitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetflowMonitors

`func (o *FabricNetFlowRecord) SetNetflowMonitors(v []FabricNetFlowMonitorRelationship)`

SetNetflowMonitors sets NetflowMonitors field to given value.

### HasNetflowMonitors

`func (o *FabricNetFlowRecord) HasNetflowMonitors() bool`

HasNetflowMonitors returns a boolean if a field has been set.

### SetNetflowMonitorsNil

`func (o *FabricNetFlowRecord) SetNetflowMonitorsNil(b bool)`

 SetNetflowMonitorsNil sets the value for NetflowMonitors to be an explicit nil

### UnsetNetflowMonitors
`func (o *FabricNetFlowRecord) UnsetNetflowMonitors()`

UnsetNetflowMonitors ensures that no value is present for NetflowMonitors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


