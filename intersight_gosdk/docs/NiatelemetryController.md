# NiatelemetryController

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.Controller"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.Controller"]
**AnomalyStatus** | Pointer to **string** | The anomaly status of the managed Controller. | [optional] [readonly] 
**ComplianceStatus** | Pointer to **string** | The compliance status of the managed Controller. | [optional] [readonly] 
**FabricType** | Pointer to **string** | The type of fabric the managed switch belongs to. | [optional] [readonly] 
**HardwareConformance** | Pointer to **string** | The hardware conformance status of the managed Controller. | [optional] [readonly] 
**HardwareEndOfVulnerabilitySupportDate** | Pointer to **string** | The hardware end of vulnerability support date of the managed Controller. | [optional] [readonly] 
**HardwareLastDateOfSupport** | Pointer to **string** | The hardware last date of support of the managed Controller. | [optional] [readonly] 
**Hostname** | Pointer to **string** | The hostname of the managed Controller. | [optional] [readonly] 
**InBandIpV4Address** | Pointer to **string** | The inband IPv4 address of the managed Controller. | [optional] [readonly] 
**InBandIpV6Address** | Pointer to **string** | The inband IPv6 address of the managed Controller. | [optional] [readonly] 
**Model** | Pointer to **string** | The model of the managed Controller. | [optional] [readonly] 
**OperatingSystem** | Pointer to **string** | The operating system of the managed Controller. | [optional] [readonly] 
**OperatingSystemVersion** | Pointer to **string** | The OS version of the managed Controller. | [optional] [readonly] 
**OutOfBandIpV4Address** | Pointer to **string** | The out-of-band IPv4 address of the managed Controller. | [optional] [readonly] 
**OutOfBandIpV6Address** | Pointer to **string** | The out-of-band IPv6 address of the managed Controller. | [optional] [readonly] 
**Role** | Pointer to **string** | The role of the managed Controller. | [optional] [readonly] 
**Serial** | Pointer to **string** | The serial number of the managed Controller. | [optional] [readonly] 
**Site** | Pointer to **string** | The name of the managed fabric. | [optional] [readonly] 
**SoftwareConformance** | Pointer to **string** | The software conformance status of the managed Controller. | [optional] [readonly] 
**SoftwareEndOfVulnerabilitySupportDate** | Pointer to **string** | The software end of vulnerability support date of the managed Controller. | [optional] [readonly] 
**SoftwareLastDateOfSupport** | Pointer to **string** | The software last date of support of the managed Controller. | [optional] [readonly] 
**SoftwareVersion** | Pointer to **string** | The software version of the managed Controller. | [optional] [readonly] 
**StrongAssetId** | Pointer to **string** | The DBID of the managed Controller. | [optional] [readonly] 
**TelemetryEnabled** | Pointer to **bool** | Indicates whether telemetry is enabled for the fabric. | [optional] [readonly] 
**Fabric** | Pointer to [**NullableNiatelemetryFabricRelationship**](NiatelemetryFabricRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryController

`func NewNiatelemetryController(classId string, objectType string, ) *NiatelemetryController`

NewNiatelemetryController instantiates a new NiatelemetryController object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryControllerWithDefaults

`func NewNiatelemetryControllerWithDefaults() *NiatelemetryController`

NewNiatelemetryControllerWithDefaults instantiates a new NiatelemetryController object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryController) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryController) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryController) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryController) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryController) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryController) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAnomalyStatus

`func (o *NiatelemetryController) GetAnomalyStatus() string`

GetAnomalyStatus returns the AnomalyStatus field if non-nil, zero value otherwise.

### GetAnomalyStatusOk

`func (o *NiatelemetryController) GetAnomalyStatusOk() (*string, bool)`

GetAnomalyStatusOk returns a tuple with the AnomalyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnomalyStatus

`func (o *NiatelemetryController) SetAnomalyStatus(v string)`

SetAnomalyStatus sets AnomalyStatus field to given value.

### HasAnomalyStatus

`func (o *NiatelemetryController) HasAnomalyStatus() bool`

HasAnomalyStatus returns a boolean if a field has been set.

### GetComplianceStatus

`func (o *NiatelemetryController) GetComplianceStatus() string`

GetComplianceStatus returns the ComplianceStatus field if non-nil, zero value otherwise.

### GetComplianceStatusOk

`func (o *NiatelemetryController) GetComplianceStatusOk() (*string, bool)`

GetComplianceStatusOk returns a tuple with the ComplianceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceStatus

`func (o *NiatelemetryController) SetComplianceStatus(v string)`

SetComplianceStatus sets ComplianceStatus field to given value.

### HasComplianceStatus

`func (o *NiatelemetryController) HasComplianceStatus() bool`

HasComplianceStatus returns a boolean if a field has been set.

### GetFabricType

`func (o *NiatelemetryController) GetFabricType() string`

GetFabricType returns the FabricType field if non-nil, zero value otherwise.

### GetFabricTypeOk

`func (o *NiatelemetryController) GetFabricTypeOk() (*string, bool)`

GetFabricTypeOk returns a tuple with the FabricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricType

`func (o *NiatelemetryController) SetFabricType(v string)`

SetFabricType sets FabricType field to given value.

### HasFabricType

`func (o *NiatelemetryController) HasFabricType() bool`

HasFabricType returns a boolean if a field has been set.

### GetHardwareConformance

`func (o *NiatelemetryController) GetHardwareConformance() string`

GetHardwareConformance returns the HardwareConformance field if non-nil, zero value otherwise.

### GetHardwareConformanceOk

`func (o *NiatelemetryController) GetHardwareConformanceOk() (*string, bool)`

GetHardwareConformanceOk returns a tuple with the HardwareConformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareConformance

`func (o *NiatelemetryController) SetHardwareConformance(v string)`

SetHardwareConformance sets HardwareConformance field to given value.

### HasHardwareConformance

`func (o *NiatelemetryController) HasHardwareConformance() bool`

HasHardwareConformance returns a boolean if a field has been set.

### GetHardwareEndOfVulnerabilitySupportDate

`func (o *NiatelemetryController) GetHardwareEndOfVulnerabilitySupportDate() string`

GetHardwareEndOfVulnerabilitySupportDate returns the HardwareEndOfVulnerabilitySupportDate field if non-nil, zero value otherwise.

### GetHardwareEndOfVulnerabilitySupportDateOk

`func (o *NiatelemetryController) GetHardwareEndOfVulnerabilitySupportDateOk() (*string, bool)`

GetHardwareEndOfVulnerabilitySupportDateOk returns a tuple with the HardwareEndOfVulnerabilitySupportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareEndOfVulnerabilitySupportDate

`func (o *NiatelemetryController) SetHardwareEndOfVulnerabilitySupportDate(v string)`

SetHardwareEndOfVulnerabilitySupportDate sets HardwareEndOfVulnerabilitySupportDate field to given value.

### HasHardwareEndOfVulnerabilitySupportDate

`func (o *NiatelemetryController) HasHardwareEndOfVulnerabilitySupportDate() bool`

HasHardwareEndOfVulnerabilitySupportDate returns a boolean if a field has been set.

### GetHardwareLastDateOfSupport

`func (o *NiatelemetryController) GetHardwareLastDateOfSupport() string`

GetHardwareLastDateOfSupport returns the HardwareLastDateOfSupport field if non-nil, zero value otherwise.

### GetHardwareLastDateOfSupportOk

`func (o *NiatelemetryController) GetHardwareLastDateOfSupportOk() (*string, bool)`

GetHardwareLastDateOfSupportOk returns a tuple with the HardwareLastDateOfSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareLastDateOfSupport

`func (o *NiatelemetryController) SetHardwareLastDateOfSupport(v string)`

SetHardwareLastDateOfSupport sets HardwareLastDateOfSupport field to given value.

### HasHardwareLastDateOfSupport

`func (o *NiatelemetryController) HasHardwareLastDateOfSupport() bool`

HasHardwareLastDateOfSupport returns a boolean if a field has been set.

### GetHostname

`func (o *NiatelemetryController) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *NiatelemetryController) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *NiatelemetryController) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *NiatelemetryController) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetInBandIpV4Address

`func (o *NiatelemetryController) GetInBandIpV4Address() string`

GetInBandIpV4Address returns the InBandIpV4Address field if non-nil, zero value otherwise.

### GetInBandIpV4AddressOk

`func (o *NiatelemetryController) GetInBandIpV4AddressOk() (*string, bool)`

GetInBandIpV4AddressOk returns a tuple with the InBandIpV4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInBandIpV4Address

`func (o *NiatelemetryController) SetInBandIpV4Address(v string)`

SetInBandIpV4Address sets InBandIpV4Address field to given value.

### HasInBandIpV4Address

`func (o *NiatelemetryController) HasInBandIpV4Address() bool`

HasInBandIpV4Address returns a boolean if a field has been set.

### GetInBandIpV6Address

`func (o *NiatelemetryController) GetInBandIpV6Address() string`

GetInBandIpV6Address returns the InBandIpV6Address field if non-nil, zero value otherwise.

### GetInBandIpV6AddressOk

`func (o *NiatelemetryController) GetInBandIpV6AddressOk() (*string, bool)`

GetInBandIpV6AddressOk returns a tuple with the InBandIpV6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInBandIpV6Address

`func (o *NiatelemetryController) SetInBandIpV6Address(v string)`

SetInBandIpV6Address sets InBandIpV6Address field to given value.

### HasInBandIpV6Address

`func (o *NiatelemetryController) HasInBandIpV6Address() bool`

HasInBandIpV6Address returns a boolean if a field has been set.

### GetModel

`func (o *NiatelemetryController) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *NiatelemetryController) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *NiatelemetryController) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *NiatelemetryController) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *NiatelemetryController) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *NiatelemetryController) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *NiatelemetryController) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *NiatelemetryController) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetOperatingSystemVersion

`func (o *NiatelemetryController) GetOperatingSystemVersion() string`

GetOperatingSystemVersion returns the OperatingSystemVersion field if non-nil, zero value otherwise.

### GetOperatingSystemVersionOk

`func (o *NiatelemetryController) GetOperatingSystemVersionOk() (*string, bool)`

GetOperatingSystemVersionOk returns a tuple with the OperatingSystemVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemVersion

`func (o *NiatelemetryController) SetOperatingSystemVersion(v string)`

SetOperatingSystemVersion sets OperatingSystemVersion field to given value.

### HasOperatingSystemVersion

`func (o *NiatelemetryController) HasOperatingSystemVersion() bool`

HasOperatingSystemVersion returns a boolean if a field has been set.

### GetOutOfBandIpV4Address

`func (o *NiatelemetryController) GetOutOfBandIpV4Address() string`

GetOutOfBandIpV4Address returns the OutOfBandIpV4Address field if non-nil, zero value otherwise.

### GetOutOfBandIpV4AddressOk

`func (o *NiatelemetryController) GetOutOfBandIpV4AddressOk() (*string, bool)`

GetOutOfBandIpV4AddressOk returns a tuple with the OutOfBandIpV4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandIpV4Address

`func (o *NiatelemetryController) SetOutOfBandIpV4Address(v string)`

SetOutOfBandIpV4Address sets OutOfBandIpV4Address field to given value.

### HasOutOfBandIpV4Address

`func (o *NiatelemetryController) HasOutOfBandIpV4Address() bool`

HasOutOfBandIpV4Address returns a boolean if a field has been set.

### GetOutOfBandIpV6Address

`func (o *NiatelemetryController) GetOutOfBandIpV6Address() string`

GetOutOfBandIpV6Address returns the OutOfBandIpV6Address field if non-nil, zero value otherwise.

### GetOutOfBandIpV6AddressOk

`func (o *NiatelemetryController) GetOutOfBandIpV6AddressOk() (*string, bool)`

GetOutOfBandIpV6AddressOk returns a tuple with the OutOfBandIpV6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfBandIpV6Address

`func (o *NiatelemetryController) SetOutOfBandIpV6Address(v string)`

SetOutOfBandIpV6Address sets OutOfBandIpV6Address field to given value.

### HasOutOfBandIpV6Address

`func (o *NiatelemetryController) HasOutOfBandIpV6Address() bool`

HasOutOfBandIpV6Address returns a boolean if a field has been set.

### GetRole

`func (o *NiatelemetryController) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *NiatelemetryController) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *NiatelemetryController) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *NiatelemetryController) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSerial

`func (o *NiatelemetryController) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *NiatelemetryController) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *NiatelemetryController) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *NiatelemetryController) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetSite

`func (o *NiatelemetryController) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *NiatelemetryController) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *NiatelemetryController) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *NiatelemetryController) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetSoftwareConformance

`func (o *NiatelemetryController) GetSoftwareConformance() string`

GetSoftwareConformance returns the SoftwareConformance field if non-nil, zero value otherwise.

### GetSoftwareConformanceOk

`func (o *NiatelemetryController) GetSoftwareConformanceOk() (*string, bool)`

GetSoftwareConformanceOk returns a tuple with the SoftwareConformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareConformance

`func (o *NiatelemetryController) SetSoftwareConformance(v string)`

SetSoftwareConformance sets SoftwareConformance field to given value.

### HasSoftwareConformance

`func (o *NiatelemetryController) HasSoftwareConformance() bool`

HasSoftwareConformance returns a boolean if a field has been set.

### GetSoftwareEndOfVulnerabilitySupportDate

`func (o *NiatelemetryController) GetSoftwareEndOfVulnerabilitySupportDate() string`

GetSoftwareEndOfVulnerabilitySupportDate returns the SoftwareEndOfVulnerabilitySupportDate field if non-nil, zero value otherwise.

### GetSoftwareEndOfVulnerabilitySupportDateOk

`func (o *NiatelemetryController) GetSoftwareEndOfVulnerabilitySupportDateOk() (*string, bool)`

GetSoftwareEndOfVulnerabilitySupportDateOk returns a tuple with the SoftwareEndOfVulnerabilitySupportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareEndOfVulnerabilitySupportDate

`func (o *NiatelemetryController) SetSoftwareEndOfVulnerabilitySupportDate(v string)`

SetSoftwareEndOfVulnerabilitySupportDate sets SoftwareEndOfVulnerabilitySupportDate field to given value.

### HasSoftwareEndOfVulnerabilitySupportDate

`func (o *NiatelemetryController) HasSoftwareEndOfVulnerabilitySupportDate() bool`

HasSoftwareEndOfVulnerabilitySupportDate returns a boolean if a field has been set.

### GetSoftwareLastDateOfSupport

`func (o *NiatelemetryController) GetSoftwareLastDateOfSupport() string`

GetSoftwareLastDateOfSupport returns the SoftwareLastDateOfSupport field if non-nil, zero value otherwise.

### GetSoftwareLastDateOfSupportOk

`func (o *NiatelemetryController) GetSoftwareLastDateOfSupportOk() (*string, bool)`

GetSoftwareLastDateOfSupportOk returns a tuple with the SoftwareLastDateOfSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareLastDateOfSupport

`func (o *NiatelemetryController) SetSoftwareLastDateOfSupport(v string)`

SetSoftwareLastDateOfSupport sets SoftwareLastDateOfSupport field to given value.

### HasSoftwareLastDateOfSupport

`func (o *NiatelemetryController) HasSoftwareLastDateOfSupport() bool`

HasSoftwareLastDateOfSupport returns a boolean if a field has been set.

### GetSoftwareVersion

`func (o *NiatelemetryController) GetSoftwareVersion() string`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *NiatelemetryController) GetSoftwareVersionOk() (*string, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *NiatelemetryController) SetSoftwareVersion(v string)`

SetSoftwareVersion sets SoftwareVersion field to given value.

### HasSoftwareVersion

`func (o *NiatelemetryController) HasSoftwareVersion() bool`

HasSoftwareVersion returns a boolean if a field has been set.

### GetStrongAssetId

`func (o *NiatelemetryController) GetStrongAssetId() string`

GetStrongAssetId returns the StrongAssetId field if non-nil, zero value otherwise.

### GetStrongAssetIdOk

`func (o *NiatelemetryController) GetStrongAssetIdOk() (*string, bool)`

GetStrongAssetIdOk returns a tuple with the StrongAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrongAssetId

`func (o *NiatelemetryController) SetStrongAssetId(v string)`

SetStrongAssetId sets StrongAssetId field to given value.

### HasStrongAssetId

`func (o *NiatelemetryController) HasStrongAssetId() bool`

HasStrongAssetId returns a boolean if a field has been set.

### GetTelemetryEnabled

`func (o *NiatelemetryController) GetTelemetryEnabled() bool`

GetTelemetryEnabled returns the TelemetryEnabled field if non-nil, zero value otherwise.

### GetTelemetryEnabledOk

`func (o *NiatelemetryController) GetTelemetryEnabledOk() (*bool, bool)`

GetTelemetryEnabledOk returns a tuple with the TelemetryEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelemetryEnabled

`func (o *NiatelemetryController) SetTelemetryEnabled(v bool)`

SetTelemetryEnabled sets TelemetryEnabled field to given value.

### HasTelemetryEnabled

`func (o *NiatelemetryController) HasTelemetryEnabled() bool`

HasTelemetryEnabled returns a boolean if a field has been set.

### GetFabric

`func (o *NiatelemetryController) GetFabric() NiatelemetryFabricRelationship`

GetFabric returns the Fabric field if non-nil, zero value otherwise.

### GetFabricOk

`func (o *NiatelemetryController) GetFabricOk() (*NiatelemetryFabricRelationship, bool)`

GetFabricOk returns a tuple with the Fabric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabric

`func (o *NiatelemetryController) SetFabric(v NiatelemetryFabricRelationship)`

SetFabric sets Fabric field to given value.

### HasFabric

`func (o *NiatelemetryController) HasFabric() bool`

HasFabric returns a boolean if a field has been set.

### SetFabricNil

`func (o *NiatelemetryController) SetFabricNil(b bool)`

 SetFabricNil sets the value for Fabric to be an explicit nil

### UnsetFabric
`func (o *NiatelemetryController) UnsetFabric()`

UnsetFabric ensures that no value is present for Fabric, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


