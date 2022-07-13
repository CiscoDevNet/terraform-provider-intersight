# NiatelemetryHealthInsightsDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.HealthInsightsData"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.HealthInsightsData"]
**FltEqptFlashMinorAlarmFaultCount** | Pointer to **int64** | Flt eqpt flash minor alarm fault count for the APIC. | [optional] 
**FltEqptFlashWornOutFaultCount** | Pointer to **int64** | Flt eqpt flash worn out fault count for the APIC. | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**SiteName** | Pointer to **string** | Name of the APIC site from which this data is being collected.. | [optional] 
**SsdMonitoringFaultCount** | Pointer to **int64** | Ssd fault monitoring count for the APIC. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryHealthInsightsDataAllOf

`func NewNiatelemetryHealthInsightsDataAllOf(classId string, objectType string, ) *NiatelemetryHealthInsightsDataAllOf`

NewNiatelemetryHealthInsightsDataAllOf instantiates a new NiatelemetryHealthInsightsDataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryHealthInsightsDataAllOfWithDefaults

`func NewNiatelemetryHealthInsightsDataAllOfWithDefaults() *NiatelemetryHealthInsightsDataAllOf`

NewNiatelemetryHealthInsightsDataAllOfWithDefaults instantiates a new NiatelemetryHealthInsightsDataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryHealthInsightsDataAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryHealthInsightsDataAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryHealthInsightsDataAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryHealthInsightsDataAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryHealthInsightsDataAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryHealthInsightsDataAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFltEqptFlashMinorAlarmFaultCount

`func (o *NiatelemetryHealthInsightsDataAllOf) GetFltEqptFlashMinorAlarmFaultCount() int64`

GetFltEqptFlashMinorAlarmFaultCount returns the FltEqptFlashMinorAlarmFaultCount field if non-nil, zero value otherwise.

### GetFltEqptFlashMinorAlarmFaultCountOk

`func (o *NiatelemetryHealthInsightsDataAllOf) GetFltEqptFlashMinorAlarmFaultCountOk() (*int64, bool)`

GetFltEqptFlashMinorAlarmFaultCountOk returns a tuple with the FltEqptFlashMinorAlarmFaultCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFltEqptFlashMinorAlarmFaultCount

`func (o *NiatelemetryHealthInsightsDataAllOf) SetFltEqptFlashMinorAlarmFaultCount(v int64)`

SetFltEqptFlashMinorAlarmFaultCount sets FltEqptFlashMinorAlarmFaultCount field to given value.

### HasFltEqptFlashMinorAlarmFaultCount

`func (o *NiatelemetryHealthInsightsDataAllOf) HasFltEqptFlashMinorAlarmFaultCount() bool`

HasFltEqptFlashMinorAlarmFaultCount returns a boolean if a field has been set.

### GetFltEqptFlashWornOutFaultCount

`func (o *NiatelemetryHealthInsightsDataAllOf) GetFltEqptFlashWornOutFaultCount() int64`

GetFltEqptFlashWornOutFaultCount returns the FltEqptFlashWornOutFaultCount field if non-nil, zero value otherwise.

### GetFltEqptFlashWornOutFaultCountOk

`func (o *NiatelemetryHealthInsightsDataAllOf) GetFltEqptFlashWornOutFaultCountOk() (*int64, bool)`

GetFltEqptFlashWornOutFaultCountOk returns a tuple with the FltEqptFlashWornOutFaultCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFltEqptFlashWornOutFaultCount

`func (o *NiatelemetryHealthInsightsDataAllOf) SetFltEqptFlashWornOutFaultCount(v int64)`

SetFltEqptFlashWornOutFaultCount sets FltEqptFlashWornOutFaultCount field to given value.

### HasFltEqptFlashWornOutFaultCount

`func (o *NiatelemetryHealthInsightsDataAllOf) HasFltEqptFlashWornOutFaultCount() bool`

HasFltEqptFlashWornOutFaultCount returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryHealthInsightsDataAllOf) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryHealthInsightsDataAllOf) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryHealthInsightsDataAllOf) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryHealthInsightsDataAllOf) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryHealthInsightsDataAllOf) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryHealthInsightsDataAllOf) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryHealthInsightsDataAllOf) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryHealthInsightsDataAllOf) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryHealthInsightsDataAllOf) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryHealthInsightsDataAllOf) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryHealthInsightsDataAllOf) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryHealthInsightsDataAllOf) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSsdMonitoringFaultCount

`func (o *NiatelemetryHealthInsightsDataAllOf) GetSsdMonitoringFaultCount() int64`

GetSsdMonitoringFaultCount returns the SsdMonitoringFaultCount field if non-nil, zero value otherwise.

### GetSsdMonitoringFaultCountOk

`func (o *NiatelemetryHealthInsightsDataAllOf) GetSsdMonitoringFaultCountOk() (*int64, bool)`

GetSsdMonitoringFaultCountOk returns a tuple with the SsdMonitoringFaultCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsdMonitoringFaultCount

`func (o *NiatelemetryHealthInsightsDataAllOf) SetSsdMonitoringFaultCount(v int64)`

SetSsdMonitoringFaultCount sets SsdMonitoringFaultCount field to given value.

### HasSsdMonitoringFaultCount

`func (o *NiatelemetryHealthInsightsDataAllOf) HasSsdMonitoringFaultCount() bool`

HasSsdMonitoringFaultCount returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryHealthInsightsDataAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryHealthInsightsDataAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryHealthInsightsDataAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryHealthInsightsDataAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


