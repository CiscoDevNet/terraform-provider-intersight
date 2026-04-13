# FabricNetFlowMonitor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "fabric.NetFlowMonitor"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "fabric.NetFlowMonitor"]
**Name** | Pointer to **string** | Netflow Monitor name, must be a maximum of 63 characters, without spacing. | [optional] 
**VnicUsageCount** | Pointer to **int64** | The count of the NetFlow monitor usage on vNICs. | [optional] [readonly] 
**FlowExporters** | Pointer to [**[]FabricNetFlowExporterRelationship**](FabricNetFlowExporterRelationship.md) | An array of relationships to fabricNetFlowExporter resources. | [optional] 
**FlowRecord** | Pointer to [**NullableFabricNetFlowRecordRelationship**](FabricNetFlowRecordRelationship.md) |  | [optional] 
**NetFlowPolicy** | Pointer to [**NullableFabricNetFlowPolicyRelationship**](FabricNetFlowPolicyRelationship.md) |  | [optional] 

## Methods

### NewFabricNetFlowMonitor

`func NewFabricNetFlowMonitor(classId string, objectType string, ) *FabricNetFlowMonitor`

NewFabricNetFlowMonitor instantiates a new FabricNetFlowMonitor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricNetFlowMonitorWithDefaults

`func NewFabricNetFlowMonitorWithDefaults() *FabricNetFlowMonitor`

NewFabricNetFlowMonitorWithDefaults instantiates a new FabricNetFlowMonitor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FabricNetFlowMonitor) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricNetFlowMonitor) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricNetFlowMonitor) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FabricNetFlowMonitor) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricNetFlowMonitor) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricNetFlowMonitor) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *FabricNetFlowMonitor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FabricNetFlowMonitor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FabricNetFlowMonitor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FabricNetFlowMonitor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVnicUsageCount

`func (o *FabricNetFlowMonitor) GetVnicUsageCount() int64`

GetVnicUsageCount returns the VnicUsageCount field if non-nil, zero value otherwise.

### GetVnicUsageCountOk

`func (o *FabricNetFlowMonitor) GetVnicUsageCountOk() (*int64, bool)`

GetVnicUsageCountOk returns a tuple with the VnicUsageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnicUsageCount

`func (o *FabricNetFlowMonitor) SetVnicUsageCount(v int64)`

SetVnicUsageCount sets VnicUsageCount field to given value.

### HasVnicUsageCount

`func (o *FabricNetFlowMonitor) HasVnicUsageCount() bool`

HasVnicUsageCount returns a boolean if a field has been set.

### GetFlowExporters

`func (o *FabricNetFlowMonitor) GetFlowExporters() []FabricNetFlowExporterRelationship`

GetFlowExporters returns the FlowExporters field if non-nil, zero value otherwise.

### GetFlowExportersOk

`func (o *FabricNetFlowMonitor) GetFlowExportersOk() (*[]FabricNetFlowExporterRelationship, bool)`

GetFlowExportersOk returns a tuple with the FlowExporters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowExporters

`func (o *FabricNetFlowMonitor) SetFlowExporters(v []FabricNetFlowExporterRelationship)`

SetFlowExporters sets FlowExporters field to given value.

### HasFlowExporters

`func (o *FabricNetFlowMonitor) HasFlowExporters() bool`

HasFlowExporters returns a boolean if a field has been set.

### SetFlowExportersNil

`func (o *FabricNetFlowMonitor) SetFlowExportersNil(b bool)`

 SetFlowExportersNil sets the value for FlowExporters to be an explicit nil

### UnsetFlowExporters
`func (o *FabricNetFlowMonitor) UnsetFlowExporters()`

UnsetFlowExporters ensures that no value is present for FlowExporters, not even an explicit nil
### GetFlowRecord

`func (o *FabricNetFlowMonitor) GetFlowRecord() FabricNetFlowRecordRelationship`

GetFlowRecord returns the FlowRecord field if non-nil, zero value otherwise.

### GetFlowRecordOk

`func (o *FabricNetFlowMonitor) GetFlowRecordOk() (*FabricNetFlowRecordRelationship, bool)`

GetFlowRecordOk returns a tuple with the FlowRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowRecord

`func (o *FabricNetFlowMonitor) SetFlowRecord(v FabricNetFlowRecordRelationship)`

SetFlowRecord sets FlowRecord field to given value.

### HasFlowRecord

`func (o *FabricNetFlowMonitor) HasFlowRecord() bool`

HasFlowRecord returns a boolean if a field has been set.

### SetFlowRecordNil

`func (o *FabricNetFlowMonitor) SetFlowRecordNil(b bool)`

 SetFlowRecordNil sets the value for FlowRecord to be an explicit nil

### UnsetFlowRecord
`func (o *FabricNetFlowMonitor) UnsetFlowRecord()`

UnsetFlowRecord ensures that no value is present for FlowRecord, not even an explicit nil
### GetNetFlowPolicy

`func (o *FabricNetFlowMonitor) GetNetFlowPolicy() FabricNetFlowPolicyRelationship`

GetNetFlowPolicy returns the NetFlowPolicy field if non-nil, zero value otherwise.

### GetNetFlowPolicyOk

`func (o *FabricNetFlowMonitor) GetNetFlowPolicyOk() (*FabricNetFlowPolicyRelationship, bool)`

GetNetFlowPolicyOk returns a tuple with the NetFlowPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetFlowPolicy

`func (o *FabricNetFlowMonitor) SetNetFlowPolicy(v FabricNetFlowPolicyRelationship)`

SetNetFlowPolicy sets NetFlowPolicy field to given value.

### HasNetFlowPolicy

`func (o *FabricNetFlowMonitor) HasNetFlowPolicy() bool`

HasNetFlowPolicy returns a boolean if a field has been set.

### SetNetFlowPolicyNil

`func (o *FabricNetFlowMonitor) SetNetFlowPolicyNil(b bool)`

 SetNetFlowPolicyNil sets the value for NetFlowPolicy to be an explicit nil

### UnsetNetFlowPolicy
`func (o *FabricNetFlowMonitor) UnsetNetFlowPolicy()`

UnsetNetFlowPolicy ensures that no value is present for NetFlowPolicy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


