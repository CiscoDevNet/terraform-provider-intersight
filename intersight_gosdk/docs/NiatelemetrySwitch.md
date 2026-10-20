# NiatelemetrySwitch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.Switch"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.Switch"]
**AnomalyStatus** | Pointer to **string** | The anomaly status of the managed switch. | [optional] [readonly] 
**ComplianceStatus** | Pointer to **string** | The compliance status of the managed switch. | [optional] [readonly] 
**FabricType** | Pointer to **string** | The type of fabric the managed switch belongs to. | [optional] [readonly] 
**HardwareConformance** | Pointer to **string** | The hardware conformance status of the managed switch. | [optional] [readonly] 
**HardwareEndOfVulnerabilitySupportDate** | Pointer to **string** | The hardware end of vulnerability support date of the managed switch. | [optional] [readonly] 
**HardwareLastDateOfSupport** | Pointer to **string** | The hardware last date of support of the managed switch. | [optional] [readonly] 
**Hostname** | Pointer to **string** | The hostname of the managed switch. | [optional] [readonly] 
**InBandIp** | Pointer to **string** | The inband IPv4 address of the managed switch. | [optional] [readonly] 
**IpAddress** | Pointer to **string** | The IP address of the managed switch. | [optional] [readonly] 
**Model** | Pointer to **string** | The model of the managed switch. | [optional] [readonly] 
**OperatingSystemVersion** | Pointer to **string** | The OS version of the managed switch. | [optional] [readonly] 
**OutOfBandIp** | Pointer to **string** | The out-of-band IPv4 address of the managed switch. | [optional] [readonly] 
**Role** | Pointer to **string** | The role of the managed switch. | [optional] [readonly] 
**Serial** | Pointer to **string** | The serial number of the managed switch. | [optional] [readonly] 
**Site** | Pointer to **string** | The name of the managed fabric. | [optional] [readonly] 
**SoftwareConformance** | Pointer to **string** | The software conformance status of the managed switch. | [optional] [readonly] 
**SoftwareEndOfVulnerabilitySupportDate** | Pointer to **string** | The software end of vulnerability support date of the managed switch. | [optional] [readonly] 
**SoftwareLastDateOfSupport** | Pointer to **string** | The software last date of support of the managed switch. | [optional] [readonly] 
**SoftwareVersion** | Pointer to **string** | The software version of the managed switch. | [optional] [readonly] 
**StrongAssetId** | Pointer to **string** | The DBID of the managed switch. | [optional] [readonly] 
**SystemUpTime** | Pointer to **string** | The system uptime of the managed switch. | [optional] [readonly] 
**TelemetryEnabled** | Pointer to **bool** | Indicates whether telemetry is enabled for the fabric. | [optional] [readonly] 
**Fabric** | Pointer to [**NullableNiatelemetryFabricRelationship**](NiatelemetryFabricRelationship.md) |  | [optional] 
**SwitchInterfaces** | Pointer to [**[]NiatelemetrySwitchInterfaceRelationship**](NiatelemetrySwitchInterfaceRelationship.md) | An array of relationships to niatelemetrySwitchInterface resources. | [optional] [readonly] 

## Methods

### NewNiatelemetrySwitch

`func NewNiatelemetrySwitch(classId string, objectType string, ) *NiatelemetrySwitch`

NewNiatelemetrySwitch instantiates a new NiatelemetrySwitch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetrySwitchWithDefaults

`func NewNiatelemetrySwitchWithDefaults() *NiatelemetrySwitch`

NewNiatelemetrySwitchWithDefaults instantiates a new NiatelemetrySwitch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetrySwitch) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetrySwitch) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetrySwitch) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetrySwitch) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetrySwitch) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetrySwitch) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAnomalyStatus

`func (o *NiatelemetrySwitch) GetAnomalyStatus() string`

GetAnomalyStatus returns the AnomalyStatus field if non-nil, zero value otherwise.

### GetAnomalyStatusOk

`func (o *NiatelemetrySwitch) GetAnomalyStatusOk() (*string, bool)`

GetAnomalyStatusOk returns a tuple with the AnomalyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnomalyStatus

`func (o *NiatelemetrySwitch) SetAnomalyStatus(v string)`

SetAnomalyStatus sets AnomalyStatus field to given value.

### HasAnomalyStatus

`func (o *NiatelemetrySwitch) HasAnomalyStatus() bool`

HasAnomalyStatus returns a boolean if a field has been set.

### GetComplianceStatus

`func (o *NiatelemetrySwitch) GetComplianceStatus() string`

GetComplianceStatus returns the ComplianceStatus field if non-nil, zero value otherwise.

### GetComplianceStatusOk

`func (o *NiatelemetrySwitch) GetComplianceStatusOk() (*string, bool)`

GetComplianceStatusOk returns a tuple with the ComplianceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceStatus

`func (o *NiatelemetrySwitch) SetComplianceStatus(v string)`

SetComplianceStatus sets ComplianceStatus field to given value.

### HasComplianceStatus

`func (o *NiatelemetrySwitch) HasComplianceStatus() bool`

HasComplianceStatus returns a boolean if a field has been set.

### GetFabricType

`func (o *NiatelemetrySwitch) GetFabricType() string`

GetFabricType returns the FabricType field if non-nil, zero value otherwise.

### GetFabricTypeOk

`func (o *NiatelemetrySwitch) GetFabricTypeOk() (*string, bool)`

GetFabricTypeOk returns a tuple with the FabricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricType

`func (o *NiatelemetrySwitch) SetFabricType(v string)`

SetFabricType sets FabricType field to given value.

### HasFabricType

`func (o *NiatelemetrySwitch) HasFabricType() bool`

HasFabricType returns a boolean if a field has been set.

### GetHardwareConformance

`func (o *NiatelemetrySwitch) GetHardwareConformance() string`

GetHardwareConformance returns the HardwareConformance field if non-nil, zero value otherwise.

### GetHardwareConformanceOk

`func (o *NiatelemetrySwitch) GetHardwareConformanceOk() (*string, bool)`

GetHardwareConformanceOk returns a tuple with the HardwareConformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareConformance

`func (o *NiatelemetrySwitch) SetHardwareConformance(v string)`

SetHardwareConformance sets HardwareConformance field to given value.

### HasHardwareConformance

`func (o *NiatelemetrySwitch) HasHardwareConformance() bool`

HasHardwareConformance returns a boolean if a field has been set.

### GetHardwareEndOfVulnerabilitySupportDate

`func (o *NiatelemetrySwitch) GetHardwareEndOfVulnerabilitySupportDate() string`

GetHardwareEndOfVulnerabilitySupportDate returns the HardwareEndOfVulnerabilitySupportDate field if non-nil, zero value otherwise.

### GetHardwareEndOfVulnerabilitySupportDateOk

`func (o *NiatelemetrySwitch) GetHardwareEndOfVulnerabilitySupportDateOk() (*string, bool)`

GetHardwareEndOfVulnerabilitySupportDateOk returns a tuple with the HardwareEndOfVulnerabilitySupportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareEndOfVulnerabilitySupportDate

`func (o *NiatelemetrySwitch) SetHardwareEndOfVulnerabilitySupportDate(v string)`

SetHardwareEndOfVulnerabilitySupportDate sets HardwareEndOfVulnerabilitySupportDate field to given value.

### HasHardwareEndOfVulnerabilitySupportDate

`func (o *NiatelemetrySwitch) HasHardwareEndOfVulnerabilitySupportDate() bool`

HasHardwareEndOfVulnerabilitySupportDate returns a boolean if a field has been set.

### GetHardwareLastDateOfSupport

`func (o *NiatelemetrySwitch) GetHardwareLastDateOfSupport() string`

GetHardwareLastDateOfSupport returns the HardwareLastDateOfSupport field if non-nil, zero value otherwise.

### GetHardwareLastDateOfSupportOk

`func (o *NiatelemetrySwitch) GetHardwareLastDateOfSupportOk() (*string, bool)`

GetHardwareLastDateOfSupportOk returns a tuple with the HardwareLastDateOfSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareLastDateOfSupport

`func (o *NiatelemetrySwitch) SetHardwareLastDateOfSupport(v string)`

SetHardwareLastDateOfSupport sets HardwareLastDateOfSupport field to given value.

### HasHardwareLastDateOfSupport

`func (o *NiatelemetrySwitch) HasHardwareLastDateOfSupport() bool`

HasHardwareLastDateOfSupport returns a boolean if a field has been set.

### GetHostname

`func (o *NiatelemetrySwitch) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *NiatelemetrySwitch) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *NiatelemetrySwitch) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *NiatelemetrySwitch) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetInBandIp

`func (o *NiatelemetrySwitch) GetInBandIp() string`

GetInBandIp returns the InBandIp field if non-nil, zero value otherwise.

### GetInBandIpOk

`func (o *NiatelemetrySwitch) GetInBandIpOk() (*string, bool)`

GetInBandIpOk returns a tuple with the InBandIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInBandIp

`func (o *NiatelemetrySwitch) SetInBandIp(v string)`

SetInBandIp sets InBandIp field to given value.

### HasInBandIp

`func (o *NiatelemetrySwitch) HasInBandIp() bool`

HasInBandIp returns a boolean if a field has been set.

### GetIpAddress

`func (o *NiatelemetrySwitch) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *NiatelemetrySwitch) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *NiatelemetrySwitch) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *NiatelemetrySwitch) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetModel

`func (o *NiatelemetrySwitch) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *NiatelemetrySwitch) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *NiatelemetrySwitch) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *NiatelemetrySwitch) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetOperatingSystemVersion

`func (o *NiatelemetrySwitch) GetOperatingSystemVersion() string`

GetOperatingSystemVersion returns the OperatingSystemVersion field if non-nil, zero value otherwise.

### GetOperatingSystemVersionOk

`func (o *NiatelemetrySwitch) GetOperatingSystemVersionOk() (*string, bool)`

GetOperatingSystemVersionOk returns a tuple with the OperatingSystemVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemVersion

`func (o *NiatelemetrySwitch) SetOperatingSystemVersion(v string)`

SetOperatingSystemVersion sets OperatingSystemVersion field to given value.

### HasOperatingSystemVersion

`func (o *NiatelemetrySwitch) HasOperatingSystemVersion() bool`

HasOperatingSystemVersion returns a boolean if a field has been set.

### GetOutOfBandIp

`func (o *NiatelemetrySwitch) GetOutOfBandIp() string`

GetOutOfBandIp returns the OutOfBandIp field if non-nil, zero value otherwise.

### GetOutOfBandIpOk

`func (o *NiatelemetrySwitch) GetOutOfBandIpOk() (*string, bool)`

GetOutOfBandIpOk returns a tuple with the OutOfBandIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandIp

`func (o *NiatelemetrySwitch) SetOutOfBandIp(v string)`

SetOutOfBandIp sets OutOfBandIp field to given value.

### HasOutOfBandIp

`func (o *NiatelemetrySwitch) HasOutOfBandIp() bool`

HasOutOfBandIp returns a boolean if a field has been set.

### GetRole

`func (o *NiatelemetrySwitch) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *NiatelemetrySwitch) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *NiatelemetrySwitch) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *NiatelemetrySwitch) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSerial

`func (o *NiatelemetrySwitch) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *NiatelemetrySwitch) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *NiatelemetrySwitch) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *NiatelemetrySwitch) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetSite

`func (o *NiatelemetrySwitch) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *NiatelemetrySwitch) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *NiatelemetrySwitch) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *NiatelemetrySwitch) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetSoftwareConformance

`func (o *NiatelemetrySwitch) GetSoftwareConformance() string`

GetSoftwareConformance returns the SoftwareConformance field if non-nil, zero value otherwise.

### GetSoftwareConformanceOk

`func (o *NiatelemetrySwitch) GetSoftwareConformanceOk() (*string, bool)`

GetSoftwareConformanceOk returns a tuple with the SoftwareConformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareConformance

`func (o *NiatelemetrySwitch) SetSoftwareConformance(v string)`

SetSoftwareConformance sets SoftwareConformance field to given value.

### HasSoftwareConformance

`func (o *NiatelemetrySwitch) HasSoftwareConformance() bool`

HasSoftwareConformance returns a boolean if a field has been set.

### GetSoftwareEndOfVulnerabilitySupportDate

`func (o *NiatelemetrySwitch) GetSoftwareEndOfVulnerabilitySupportDate() string`

GetSoftwareEndOfVulnerabilitySupportDate returns the SoftwareEndOfVulnerabilitySupportDate field if non-nil, zero value otherwise.

### GetSoftwareEndOfVulnerabilitySupportDateOk

`func (o *NiatelemetrySwitch) GetSoftwareEndOfVulnerabilitySupportDateOk() (*string, bool)`

GetSoftwareEndOfVulnerabilitySupportDateOk returns a tuple with the SoftwareEndOfVulnerabilitySupportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareEndOfVulnerabilitySupportDate

`func (o *NiatelemetrySwitch) SetSoftwareEndOfVulnerabilitySupportDate(v string)`

SetSoftwareEndOfVulnerabilitySupportDate sets SoftwareEndOfVulnerabilitySupportDate field to given value.

### HasSoftwareEndOfVulnerabilitySupportDate

`func (o *NiatelemetrySwitch) HasSoftwareEndOfVulnerabilitySupportDate() bool`

HasSoftwareEndOfVulnerabilitySupportDate returns a boolean if a field has been set.

### GetSoftwareLastDateOfSupport

`func (o *NiatelemetrySwitch) GetSoftwareLastDateOfSupport() string`

GetSoftwareLastDateOfSupport returns the SoftwareLastDateOfSupport field if non-nil, zero value otherwise.

### GetSoftwareLastDateOfSupportOk

`func (o *NiatelemetrySwitch) GetSoftwareLastDateOfSupportOk() (*string, bool)`

GetSoftwareLastDateOfSupportOk returns a tuple with the SoftwareLastDateOfSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareLastDateOfSupport

`func (o *NiatelemetrySwitch) SetSoftwareLastDateOfSupport(v string)`

SetSoftwareLastDateOfSupport sets SoftwareLastDateOfSupport field to given value.

### HasSoftwareLastDateOfSupport

`func (o *NiatelemetrySwitch) HasSoftwareLastDateOfSupport() bool`

HasSoftwareLastDateOfSupport returns a boolean if a field has been set.

### GetSoftwareVersion

`func (o *NiatelemetrySwitch) GetSoftwareVersion() string`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *NiatelemetrySwitch) GetSoftwareVersionOk() (*string, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *NiatelemetrySwitch) SetSoftwareVersion(v string)`

SetSoftwareVersion sets SoftwareVersion field to given value.

### HasSoftwareVersion

`func (o *NiatelemetrySwitch) HasSoftwareVersion() bool`

HasSoftwareVersion returns a boolean if a field has been set.

### GetStrongAssetId

`func (o *NiatelemetrySwitch) GetStrongAssetId() string`

GetStrongAssetId returns the StrongAssetId field if non-nil, zero value otherwise.

### GetStrongAssetIdOk

`func (o *NiatelemetrySwitch) GetStrongAssetIdOk() (*string, bool)`

GetStrongAssetIdOk returns a tuple with the StrongAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrongAssetId

`func (o *NiatelemetrySwitch) SetStrongAssetId(v string)`

SetStrongAssetId sets StrongAssetId field to given value.

### HasStrongAssetId

`func (o *NiatelemetrySwitch) HasStrongAssetId() bool`

HasStrongAssetId returns a boolean if a field has been set.

### GetSystemUpTime

`func (o *NiatelemetrySwitch) GetSystemUpTime() string`

GetSystemUpTime returns the SystemUpTime field if non-nil, zero value otherwise.

### GetSystemUpTimeOk

`func (o *NiatelemetrySwitch) GetSystemUpTimeOk() (*string, bool)`

GetSystemUpTimeOk returns a tuple with the SystemUpTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemUpTime

`func (o *NiatelemetrySwitch) SetSystemUpTime(v string)`

SetSystemUpTime sets SystemUpTime field to given value.

### HasSystemUpTime

`func (o *NiatelemetrySwitch) HasSystemUpTime() bool`

HasSystemUpTime returns a boolean if a field has been set.

### GetTelemetryEnabled

`func (o *NiatelemetrySwitch) GetTelemetryEnabled() bool`

GetTelemetryEnabled returns the TelemetryEnabled field if non-nil, zero value otherwise.

### GetTelemetryEnabledOk

`func (o *NiatelemetrySwitch) GetTelemetryEnabledOk() (*bool, bool)`

GetTelemetryEnabledOk returns a tuple with the TelemetryEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelemetryEnabled

`func (o *NiatelemetrySwitch) SetTelemetryEnabled(v bool)`

SetTelemetryEnabled sets TelemetryEnabled field to given value.

### HasTelemetryEnabled

`func (o *NiatelemetrySwitch) HasTelemetryEnabled() bool`

HasTelemetryEnabled returns a boolean if a field has been set.

### GetFabric

`func (o *NiatelemetrySwitch) GetFabric() NiatelemetryFabricRelationship`

GetFabric returns the Fabric field if non-nil, zero value otherwise.

### GetFabricOk

`func (o *NiatelemetrySwitch) GetFabricOk() (*NiatelemetryFabricRelationship, bool)`

GetFabricOk returns a tuple with the Fabric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabric

`func (o *NiatelemetrySwitch) SetFabric(v NiatelemetryFabricRelationship)`

SetFabric sets Fabric field to given value.

### HasFabric

`func (o *NiatelemetrySwitch) HasFabric() bool`

HasFabric returns a boolean if a field has been set.

### SetFabricNil

`func (o *NiatelemetrySwitch) SetFabricNil(b bool)`

 SetFabricNil sets the value for Fabric to be an explicit nil

### UnsetFabric
`func (o *NiatelemetrySwitch) UnsetFabric()`

UnsetFabric ensures that no value is present for Fabric, not even an explicit nil
### GetSwitchInterfaces

`func (o *NiatelemetrySwitch) GetSwitchInterfaces() []NiatelemetrySwitchInterfaceRelationship`

GetSwitchInterfaces returns the SwitchInterfaces field if non-nil, zero value otherwise.

### GetSwitchInterfacesOk

`func (o *NiatelemetrySwitch) GetSwitchInterfacesOk() (*[]NiatelemetrySwitchInterfaceRelationship, bool)`

GetSwitchInterfacesOk returns a tuple with the SwitchInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchInterfaces

`func (o *NiatelemetrySwitch) SetSwitchInterfaces(v []NiatelemetrySwitchInterfaceRelationship)`

SetSwitchInterfaces sets SwitchInterfaces field to given value.

### HasSwitchInterfaces

`func (o *NiatelemetrySwitch) HasSwitchInterfaces() bool`

HasSwitchInterfaces returns a boolean if a field has been set.

### SetSwitchInterfacesNil

`func (o *NiatelemetrySwitch) SetSwitchInterfacesNil(b bool)`

 SetSwitchInterfacesNil sets the value for SwitchInterfaces to be an explicit nil

### UnsetSwitchInterfaces
`func (o *NiatelemetrySwitch) UnsetSwitchInterfaces()`

UnsetSwitchInterfaces ensures that no value is present for SwitchInterfaces, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


