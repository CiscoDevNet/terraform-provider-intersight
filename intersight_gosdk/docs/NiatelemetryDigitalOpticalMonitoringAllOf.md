# NiatelemetryDigitalOpticalMonitoringAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.DigitalOpticalMonitoring"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.DigitalOpticalMonitoring"]
**Alerts** | Pointer to **string** | Alerts count for the interface in the node. | [optional] 
**Dn** | Pointer to **string** | Dn with interface name for the aci nodes. | [optional] 
**RxLos** | Pointer to **string** | RxLos count for the interface in the node. | [optional] 
**TxFaultCount** | Pointer to **string** | TxfaultCount for the interface in the node. | [optional] 

## Methods

### NewNiatelemetryDigitalOpticalMonitoringAllOf

`func NewNiatelemetryDigitalOpticalMonitoringAllOf(classId string, objectType string, ) *NiatelemetryDigitalOpticalMonitoringAllOf`

NewNiatelemetryDigitalOpticalMonitoringAllOf instantiates a new NiatelemetryDigitalOpticalMonitoringAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryDigitalOpticalMonitoringAllOfWithDefaults

`func NewNiatelemetryDigitalOpticalMonitoringAllOfWithDefaults() *NiatelemetryDigitalOpticalMonitoringAllOf`

NewNiatelemetryDigitalOpticalMonitoringAllOfWithDefaults instantiates a new NiatelemetryDigitalOpticalMonitoringAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryDigitalOpticalMonitoringAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryDigitalOpticalMonitoringAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryDigitalOpticalMonitoringAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryDigitalOpticalMonitoringAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryDigitalOpticalMonitoringAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryDigitalOpticalMonitoringAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAlerts

`func (o *NiatelemetryDigitalOpticalMonitoringAllOf) GetAlerts() string`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *NiatelemetryDigitalOpticalMonitoringAllOf) GetAlertsOk() (*string, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *NiatelemetryDigitalOpticalMonitoringAllOf) SetAlerts(v string)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *NiatelemetryDigitalOpticalMonitoringAllOf) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetDn

`func (o *NiatelemetryDigitalOpticalMonitoringAllOf) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetryDigitalOpticalMonitoringAllOf) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetryDigitalOpticalMonitoringAllOf) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetryDigitalOpticalMonitoringAllOf) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetRxLos

`func (o *NiatelemetryDigitalOpticalMonitoringAllOf) GetRxLos() string`

GetRxLos returns the RxLos field if non-nil, zero value otherwise.

### GetRxLosOk

`func (o *NiatelemetryDigitalOpticalMonitoringAllOf) GetRxLosOk() (*string, bool)`

GetRxLosOk returns a tuple with the RxLos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxLos

`func (o *NiatelemetryDigitalOpticalMonitoringAllOf) SetRxLos(v string)`

SetRxLos sets RxLos field to given value.

### HasRxLos

`func (o *NiatelemetryDigitalOpticalMonitoringAllOf) HasRxLos() bool`

HasRxLos returns a boolean if a field has been set.

### GetTxFaultCount

`func (o *NiatelemetryDigitalOpticalMonitoringAllOf) GetTxFaultCount() string`

GetTxFaultCount returns the TxFaultCount field if non-nil, zero value otherwise.

### GetTxFaultCountOk

`func (o *NiatelemetryDigitalOpticalMonitoringAllOf) GetTxFaultCountOk() (*string, bool)`

GetTxFaultCountOk returns a tuple with the TxFaultCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxFaultCount

`func (o *NiatelemetryDigitalOpticalMonitoringAllOf) SetTxFaultCount(v string)`

SetTxFaultCount sets TxFaultCount field to given value.

### HasTxFaultCount

`func (o *NiatelemetryDigitalOpticalMonitoringAllOf) HasTxFaultCount() bool`

HasTxFaultCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


