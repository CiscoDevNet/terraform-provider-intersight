# NiatelemetryApicPerformanceDataAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.ApicPerformanceData"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.ApicPerformanceData"]
**DigitalOpticalMonitoring** | Pointer to [**[]NiatelemetryDigitalOpticalMonitoring**](NiatelemetryDigitalOpticalMonitoring.md) |  | [optional] 
**Dn** | Pointer to **string** | Dn of the fabric nodes in the apic. | [optional] 
**EqptStorageFirmware** | Pointer to [**NullableNiatelemetryEqptStorageFirmware**](NiatelemetryEqptStorageFirmware.md) |  | [optional] 
**EqptcapacityPolUsage5min** | Pointer to [**NullableNiatelemetryEqptcapacityPolUsage5min**](NiatelemetryEqptcapacityPolUsage5min.md) |  | [optional] 
**EqptcapacityPrefixEntries15min** | Pointer to [**NullableNiatelemetryEqptcapacityPrefixEntries15min**](NiatelemetryEqptcapacityPrefixEntries15min.md) |  | [optional] 
**EqptcapacityPrefixEntries5min** | Pointer to [**NullableNiatelemetryEqptcapacityPrefixEntries5min**](NiatelemetryEqptcapacityPrefixEntries5min.md) |  | [optional] 
**NodeHealth** | Pointer to **int64** | Health of the fabric nodes in the apic. | [optional] 
**ProcSysCpu15min** | Pointer to [**NullableNiatelemetryProcSysCpu15min**](NiatelemetryProcSysCpu15min.md) |  | [optional] 
**ProcSysCpu5min** | Pointer to [**NullableNiatelemetryProcSysCpu5min**](NiatelemetryProcSysCpu5min.md) |  | [optional] 
**ProcSysMem15min** | Pointer to [**NullableNiatelemetryProcSysMem15min**](NiatelemetryProcSysMem15min.md) |  | [optional] 
**ProcSysMem5min** | Pointer to [**NullableNiatelemetryProcSysMem5min**](NiatelemetryProcSysMem5min.md) |  | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**SiteName** | Pointer to **string** | Name of the APIC site from which this data is being collected.. | [optional] 
**SwitchDiskUtilization** | Pointer to [**[]NiatelemetrySwitchDiskUtilization**](NiatelemetrySwitchDiskUtilization.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryApicPerformanceDataAllOf

`func NewNiatelemetryApicPerformanceDataAllOf(classId string, objectType string, ) *NiatelemetryApicPerformanceDataAllOf`

NewNiatelemetryApicPerformanceDataAllOf instantiates a new NiatelemetryApicPerformanceDataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryApicPerformanceDataAllOfWithDefaults

`func NewNiatelemetryApicPerformanceDataAllOfWithDefaults() *NiatelemetryApicPerformanceDataAllOf`

NewNiatelemetryApicPerformanceDataAllOfWithDefaults instantiates a new NiatelemetryApicPerformanceDataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryApicPerformanceDataAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryApicPerformanceDataAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryApicPerformanceDataAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryApicPerformanceDataAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryApicPerformanceDataAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryApicPerformanceDataAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDigitalOpticalMonitoring

`func (o *NiatelemetryApicPerformanceDataAllOf) GetDigitalOpticalMonitoring() []NiatelemetryDigitalOpticalMonitoring`

GetDigitalOpticalMonitoring returns the DigitalOpticalMonitoring field if non-nil, zero value otherwise.

### GetDigitalOpticalMonitoringOk

`func (o *NiatelemetryApicPerformanceDataAllOf) GetDigitalOpticalMonitoringOk() (*[]NiatelemetryDigitalOpticalMonitoring, bool)`

GetDigitalOpticalMonitoringOk returns a tuple with the DigitalOpticalMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalOpticalMonitoring

`func (o *NiatelemetryApicPerformanceDataAllOf) SetDigitalOpticalMonitoring(v []NiatelemetryDigitalOpticalMonitoring)`

SetDigitalOpticalMonitoring sets DigitalOpticalMonitoring field to given value.

### HasDigitalOpticalMonitoring

`func (o *NiatelemetryApicPerformanceDataAllOf) HasDigitalOpticalMonitoring() bool`

HasDigitalOpticalMonitoring returns a boolean if a field has been set.

### SetDigitalOpticalMonitoringNil

`func (o *NiatelemetryApicPerformanceDataAllOf) SetDigitalOpticalMonitoringNil(b bool)`

 SetDigitalOpticalMonitoringNil sets the value for DigitalOpticalMonitoring to be an explicit nil

### UnsetDigitalOpticalMonitoring
`func (o *NiatelemetryApicPerformanceDataAllOf) UnsetDigitalOpticalMonitoring()`

UnsetDigitalOpticalMonitoring ensures that no value is present for DigitalOpticalMonitoring, not even an explicit nil
### GetDn

`func (o *NiatelemetryApicPerformanceDataAllOf) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetryApicPerformanceDataAllOf) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetryApicPerformanceDataAllOf) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetryApicPerformanceDataAllOf) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetEqptStorageFirmware

`func (o *NiatelemetryApicPerformanceDataAllOf) GetEqptStorageFirmware() NiatelemetryEqptStorageFirmware`

GetEqptStorageFirmware returns the EqptStorageFirmware field if non-nil, zero value otherwise.

### GetEqptStorageFirmwareOk

`func (o *NiatelemetryApicPerformanceDataAllOf) GetEqptStorageFirmwareOk() (*NiatelemetryEqptStorageFirmware, bool)`

GetEqptStorageFirmwareOk returns a tuple with the EqptStorageFirmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEqptStorageFirmware

`func (o *NiatelemetryApicPerformanceDataAllOf) SetEqptStorageFirmware(v NiatelemetryEqptStorageFirmware)`

SetEqptStorageFirmware sets EqptStorageFirmware field to given value.

### HasEqptStorageFirmware

`func (o *NiatelemetryApicPerformanceDataAllOf) HasEqptStorageFirmware() bool`

HasEqptStorageFirmware returns a boolean if a field has been set.

### SetEqptStorageFirmwareNil

`func (o *NiatelemetryApicPerformanceDataAllOf) SetEqptStorageFirmwareNil(b bool)`

 SetEqptStorageFirmwareNil sets the value for EqptStorageFirmware to be an explicit nil

### UnsetEqptStorageFirmware
`func (o *NiatelemetryApicPerformanceDataAllOf) UnsetEqptStorageFirmware()`

UnsetEqptStorageFirmware ensures that no value is present for EqptStorageFirmware, not even an explicit nil
### GetEqptcapacityPolUsage5min

`func (o *NiatelemetryApicPerformanceDataAllOf) GetEqptcapacityPolUsage5min() NiatelemetryEqptcapacityPolUsage5min`

GetEqptcapacityPolUsage5min returns the EqptcapacityPolUsage5min field if non-nil, zero value otherwise.

### GetEqptcapacityPolUsage5minOk

`func (o *NiatelemetryApicPerformanceDataAllOf) GetEqptcapacityPolUsage5minOk() (*NiatelemetryEqptcapacityPolUsage5min, bool)`

GetEqptcapacityPolUsage5minOk returns a tuple with the EqptcapacityPolUsage5min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEqptcapacityPolUsage5min

`func (o *NiatelemetryApicPerformanceDataAllOf) SetEqptcapacityPolUsage5min(v NiatelemetryEqptcapacityPolUsage5min)`

SetEqptcapacityPolUsage5min sets EqptcapacityPolUsage5min field to given value.

### HasEqptcapacityPolUsage5min

`func (o *NiatelemetryApicPerformanceDataAllOf) HasEqptcapacityPolUsage5min() bool`

HasEqptcapacityPolUsage5min returns a boolean if a field has been set.

### SetEqptcapacityPolUsage5minNil

`func (o *NiatelemetryApicPerformanceDataAllOf) SetEqptcapacityPolUsage5minNil(b bool)`

 SetEqptcapacityPolUsage5minNil sets the value for EqptcapacityPolUsage5min to be an explicit nil

### UnsetEqptcapacityPolUsage5min
`func (o *NiatelemetryApicPerformanceDataAllOf) UnsetEqptcapacityPolUsage5min()`

UnsetEqptcapacityPolUsage5min ensures that no value is present for EqptcapacityPolUsage5min, not even an explicit nil
### GetEqptcapacityPrefixEntries15min

`func (o *NiatelemetryApicPerformanceDataAllOf) GetEqptcapacityPrefixEntries15min() NiatelemetryEqptcapacityPrefixEntries15min`

GetEqptcapacityPrefixEntries15min returns the EqptcapacityPrefixEntries15min field if non-nil, zero value otherwise.

### GetEqptcapacityPrefixEntries15minOk

`func (o *NiatelemetryApicPerformanceDataAllOf) GetEqptcapacityPrefixEntries15minOk() (*NiatelemetryEqptcapacityPrefixEntries15min, bool)`

GetEqptcapacityPrefixEntries15minOk returns a tuple with the EqptcapacityPrefixEntries15min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEqptcapacityPrefixEntries15min

`func (o *NiatelemetryApicPerformanceDataAllOf) SetEqptcapacityPrefixEntries15min(v NiatelemetryEqptcapacityPrefixEntries15min)`

SetEqptcapacityPrefixEntries15min sets EqptcapacityPrefixEntries15min field to given value.

### HasEqptcapacityPrefixEntries15min

`func (o *NiatelemetryApicPerformanceDataAllOf) HasEqptcapacityPrefixEntries15min() bool`

HasEqptcapacityPrefixEntries15min returns a boolean if a field has been set.

### SetEqptcapacityPrefixEntries15minNil

`func (o *NiatelemetryApicPerformanceDataAllOf) SetEqptcapacityPrefixEntries15minNil(b bool)`

 SetEqptcapacityPrefixEntries15minNil sets the value for EqptcapacityPrefixEntries15min to be an explicit nil

### UnsetEqptcapacityPrefixEntries15min
`func (o *NiatelemetryApicPerformanceDataAllOf) UnsetEqptcapacityPrefixEntries15min()`

UnsetEqptcapacityPrefixEntries15min ensures that no value is present for EqptcapacityPrefixEntries15min, not even an explicit nil
### GetEqptcapacityPrefixEntries5min

`func (o *NiatelemetryApicPerformanceDataAllOf) GetEqptcapacityPrefixEntries5min() NiatelemetryEqptcapacityPrefixEntries5min`

GetEqptcapacityPrefixEntries5min returns the EqptcapacityPrefixEntries5min field if non-nil, zero value otherwise.

### GetEqptcapacityPrefixEntries5minOk

`func (o *NiatelemetryApicPerformanceDataAllOf) GetEqptcapacityPrefixEntries5minOk() (*NiatelemetryEqptcapacityPrefixEntries5min, bool)`

GetEqptcapacityPrefixEntries5minOk returns a tuple with the EqptcapacityPrefixEntries5min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEqptcapacityPrefixEntries5min

`func (o *NiatelemetryApicPerformanceDataAllOf) SetEqptcapacityPrefixEntries5min(v NiatelemetryEqptcapacityPrefixEntries5min)`

SetEqptcapacityPrefixEntries5min sets EqptcapacityPrefixEntries5min field to given value.

### HasEqptcapacityPrefixEntries5min

`func (o *NiatelemetryApicPerformanceDataAllOf) HasEqptcapacityPrefixEntries5min() bool`

HasEqptcapacityPrefixEntries5min returns a boolean if a field has been set.

### SetEqptcapacityPrefixEntries5minNil

`func (o *NiatelemetryApicPerformanceDataAllOf) SetEqptcapacityPrefixEntries5minNil(b bool)`

 SetEqptcapacityPrefixEntries5minNil sets the value for EqptcapacityPrefixEntries5min to be an explicit nil

### UnsetEqptcapacityPrefixEntries5min
`func (o *NiatelemetryApicPerformanceDataAllOf) UnsetEqptcapacityPrefixEntries5min()`

UnsetEqptcapacityPrefixEntries5min ensures that no value is present for EqptcapacityPrefixEntries5min, not even an explicit nil
### GetNodeHealth

`func (o *NiatelemetryApicPerformanceDataAllOf) GetNodeHealth() int64`

GetNodeHealth returns the NodeHealth field if non-nil, zero value otherwise.

### GetNodeHealthOk

`func (o *NiatelemetryApicPerformanceDataAllOf) GetNodeHealthOk() (*int64, bool)`

GetNodeHealthOk returns a tuple with the NodeHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeHealth

`func (o *NiatelemetryApicPerformanceDataAllOf) SetNodeHealth(v int64)`

SetNodeHealth sets NodeHealth field to given value.

### HasNodeHealth

`func (o *NiatelemetryApicPerformanceDataAllOf) HasNodeHealth() bool`

HasNodeHealth returns a boolean if a field has been set.

### GetProcSysCpu15min

`func (o *NiatelemetryApicPerformanceDataAllOf) GetProcSysCpu15min() NiatelemetryProcSysCpu15min`

GetProcSysCpu15min returns the ProcSysCpu15min field if non-nil, zero value otherwise.

### GetProcSysCpu15minOk

`func (o *NiatelemetryApicPerformanceDataAllOf) GetProcSysCpu15minOk() (*NiatelemetryProcSysCpu15min, bool)`

GetProcSysCpu15minOk returns a tuple with the ProcSysCpu15min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcSysCpu15min

`func (o *NiatelemetryApicPerformanceDataAllOf) SetProcSysCpu15min(v NiatelemetryProcSysCpu15min)`

SetProcSysCpu15min sets ProcSysCpu15min field to given value.

### HasProcSysCpu15min

`func (o *NiatelemetryApicPerformanceDataAllOf) HasProcSysCpu15min() bool`

HasProcSysCpu15min returns a boolean if a field has been set.

### SetProcSysCpu15minNil

`func (o *NiatelemetryApicPerformanceDataAllOf) SetProcSysCpu15minNil(b bool)`

 SetProcSysCpu15minNil sets the value for ProcSysCpu15min to be an explicit nil

### UnsetProcSysCpu15min
`func (o *NiatelemetryApicPerformanceDataAllOf) UnsetProcSysCpu15min()`

UnsetProcSysCpu15min ensures that no value is present for ProcSysCpu15min, not even an explicit nil
### GetProcSysCpu5min

`func (o *NiatelemetryApicPerformanceDataAllOf) GetProcSysCpu5min() NiatelemetryProcSysCpu5min`

GetProcSysCpu5min returns the ProcSysCpu5min field if non-nil, zero value otherwise.

### GetProcSysCpu5minOk

`func (o *NiatelemetryApicPerformanceDataAllOf) GetProcSysCpu5minOk() (*NiatelemetryProcSysCpu5min, bool)`

GetProcSysCpu5minOk returns a tuple with the ProcSysCpu5min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcSysCpu5min

`func (o *NiatelemetryApicPerformanceDataAllOf) SetProcSysCpu5min(v NiatelemetryProcSysCpu5min)`

SetProcSysCpu5min sets ProcSysCpu5min field to given value.

### HasProcSysCpu5min

`func (o *NiatelemetryApicPerformanceDataAllOf) HasProcSysCpu5min() bool`

HasProcSysCpu5min returns a boolean if a field has been set.

### SetProcSysCpu5minNil

`func (o *NiatelemetryApicPerformanceDataAllOf) SetProcSysCpu5minNil(b bool)`

 SetProcSysCpu5minNil sets the value for ProcSysCpu5min to be an explicit nil

### UnsetProcSysCpu5min
`func (o *NiatelemetryApicPerformanceDataAllOf) UnsetProcSysCpu5min()`

UnsetProcSysCpu5min ensures that no value is present for ProcSysCpu5min, not even an explicit nil
### GetProcSysMem15min

`func (o *NiatelemetryApicPerformanceDataAllOf) GetProcSysMem15min() NiatelemetryProcSysMem15min`

GetProcSysMem15min returns the ProcSysMem15min field if non-nil, zero value otherwise.

### GetProcSysMem15minOk

`func (o *NiatelemetryApicPerformanceDataAllOf) GetProcSysMem15minOk() (*NiatelemetryProcSysMem15min, bool)`

GetProcSysMem15minOk returns a tuple with the ProcSysMem15min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcSysMem15min

`func (o *NiatelemetryApicPerformanceDataAllOf) SetProcSysMem15min(v NiatelemetryProcSysMem15min)`

SetProcSysMem15min sets ProcSysMem15min field to given value.

### HasProcSysMem15min

`func (o *NiatelemetryApicPerformanceDataAllOf) HasProcSysMem15min() bool`

HasProcSysMem15min returns a boolean if a field has been set.

### SetProcSysMem15minNil

`func (o *NiatelemetryApicPerformanceDataAllOf) SetProcSysMem15minNil(b bool)`

 SetProcSysMem15minNil sets the value for ProcSysMem15min to be an explicit nil

### UnsetProcSysMem15min
`func (o *NiatelemetryApicPerformanceDataAllOf) UnsetProcSysMem15min()`

UnsetProcSysMem15min ensures that no value is present for ProcSysMem15min, not even an explicit nil
### GetProcSysMem5min

`func (o *NiatelemetryApicPerformanceDataAllOf) GetProcSysMem5min() NiatelemetryProcSysMem5min`

GetProcSysMem5min returns the ProcSysMem5min field if non-nil, zero value otherwise.

### GetProcSysMem5minOk

`func (o *NiatelemetryApicPerformanceDataAllOf) GetProcSysMem5minOk() (*NiatelemetryProcSysMem5min, bool)`

GetProcSysMem5minOk returns a tuple with the ProcSysMem5min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcSysMem5min

`func (o *NiatelemetryApicPerformanceDataAllOf) SetProcSysMem5min(v NiatelemetryProcSysMem5min)`

SetProcSysMem5min sets ProcSysMem5min field to given value.

### HasProcSysMem5min

`func (o *NiatelemetryApicPerformanceDataAllOf) HasProcSysMem5min() bool`

HasProcSysMem5min returns a boolean if a field has been set.

### SetProcSysMem5minNil

`func (o *NiatelemetryApicPerformanceDataAllOf) SetProcSysMem5minNil(b bool)`

 SetProcSysMem5minNil sets the value for ProcSysMem5min to be an explicit nil

### UnsetProcSysMem5min
`func (o *NiatelemetryApicPerformanceDataAllOf) UnsetProcSysMem5min()`

UnsetProcSysMem5min ensures that no value is present for ProcSysMem5min, not even an explicit nil
### GetRecordType

`func (o *NiatelemetryApicPerformanceDataAllOf) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryApicPerformanceDataAllOf) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryApicPerformanceDataAllOf) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryApicPerformanceDataAllOf) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryApicPerformanceDataAllOf) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryApicPerformanceDataAllOf) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryApicPerformanceDataAllOf) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryApicPerformanceDataAllOf) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryApicPerformanceDataAllOf) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryApicPerformanceDataAllOf) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryApicPerformanceDataAllOf) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryApicPerformanceDataAllOf) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSwitchDiskUtilization

`func (o *NiatelemetryApicPerformanceDataAllOf) GetSwitchDiskUtilization() []NiatelemetrySwitchDiskUtilization`

GetSwitchDiskUtilization returns the SwitchDiskUtilization field if non-nil, zero value otherwise.

### GetSwitchDiskUtilizationOk

`func (o *NiatelemetryApicPerformanceDataAllOf) GetSwitchDiskUtilizationOk() (*[]NiatelemetrySwitchDiskUtilization, bool)`

GetSwitchDiskUtilizationOk returns a tuple with the SwitchDiskUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchDiskUtilization

`func (o *NiatelemetryApicPerformanceDataAllOf) SetSwitchDiskUtilization(v []NiatelemetrySwitchDiskUtilization)`

SetSwitchDiskUtilization sets SwitchDiskUtilization field to given value.

### HasSwitchDiskUtilization

`func (o *NiatelemetryApicPerformanceDataAllOf) HasSwitchDiskUtilization() bool`

HasSwitchDiskUtilization returns a boolean if a field has been set.

### SetSwitchDiskUtilizationNil

`func (o *NiatelemetryApicPerformanceDataAllOf) SetSwitchDiskUtilizationNil(b bool)`

 SetSwitchDiskUtilizationNil sets the value for SwitchDiskUtilization to be an explicit nil

### UnsetSwitchDiskUtilization
`func (o *NiatelemetryApicPerformanceDataAllOf) UnsetSwitchDiskUtilization()`

UnsetSwitchDiskUtilization ensures that no value is present for SwitchDiskUtilization, not even an explicit nil
### GetRegisteredDevice

`func (o *NiatelemetryApicPerformanceDataAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryApicPerformanceDataAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryApicPerformanceDataAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryApicPerformanceDataAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


