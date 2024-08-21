# ApplianceDeviceState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "appliance.DeviceState"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "appliance.DeviceState"]
**BlockReasons** | Pointer to **[]string** |  | [optional] 
**BlockedVersion** | Pointer to **string** | Version string of the current software bundle that is blocked for upgrade in the Intersight Appliance. It is used by UI to show banner message for blocked upgrade. | [optional] [readonly] 
**Capabilities** | Pointer to [**[]ApplianceKeyValuePair**](ApplianceKeyValuePair.md) |  | [optional] 
**Certificate** | Pointer to **string** | Certificate to be used for verifying software upgrade bundles. Intersight&#39;s upgrade service sets the certificate dynamically when the Intersight Appliance queries DeviceState. | [optional] [readonly] 
**CertificateNotAfter** | Pointer to **time.Time** | Expiration date of the software bundle verification certificate. | [optional] [readonly] 
**ClusterInfo** | Pointer to [**NullableOnpremClusterInfo**](OnpremClusterInfo.md) |  | [optional] 
**ConnectionStatus** | Pointer to **string** | Intersight Appliance&#39;s connectivity status. ConnectionStatus field is updated infrequently, and value may not be up to date. However, upgrade service will populate this field with actual value when queried. * &#x60;&#x60; - The target details have been persisted but Intersight has not yet attempted to connect to the target. * &#x60;Connected&#x60; - Intersight is able to establish a connection to the target and initiate management activities. * &#x60;NotConnected&#x60; - Intersight is unable to establish a connection to the target. * &#x60;ClaimInProgress&#x60; - Claim of the target is in progress. A connection to the target has not been fully established. * &#x60;UnclaimInProgress&#x60; - Unclaim of the target is in progress. Intersight is able to connect to the target and all management operations are supported. * &#x60;Unclaimed&#x60; - The device was un-claimed from the users account by an Administrator of the device. Also indicates the failure to claim Targets of type HTTP Endpoint in Intersight. * &#x60;Claimed&#x60; - Target of type HTTP Endpoint is successfully claimed in Intersight. Currently no validation is performed to verify the Target connectivity from Intersight at the time of creation. However invoking API from Intersight Orchestrator fails if this Target is not reachable from Intersight or if Target API credentials are incorrect. | [optional] [readonly] [default to ""]
**CurrentFingerprint** | Pointer to **string** | Fingerprint of the software bundle that is currently installed in the Intersight Appliance. | [optional] [readonly] 
**CurrentVersion** | Pointer to **string** | Version string of the current software bundle that is installed in the Intersight Appliance. | [optional] [readonly] 
**DcVersion** | Pointer to **string** | Version string of the Intersight Appliance&#39;s device connector. Device connector reports version number during the initial handshake. | [optional] [readonly] 
**DesiredVersion** | Pointer to **string** | The desired software bundle version of the Intersight Appliance. | [optional] [readonly] 
**DownloadedFingerprint** | Pointer to **string** | Fingerprint of the downloaded software bundle. | [optional] [readonly] 
**DownloadedMetadataVersion** | Pointer to [**[]ApplianceMetadataManifestVersion**](ApplianceMetadataManifestVersion.md) |  | [optional] 
**DownloadedVersion** | Pointer to **string** | Intersight appliance software bundle version downloaded on the endpoint. Once Intersight Appliance upgrade service starts processing the version, it is updated as pending version. | [optional] [readonly] 
**Hostname** | Pointer to **string** | Hostname of the Intersight Appliance. | [optional] [readonly] 
**PendingFingerprint** | Pointer to **string** | Fingerprint of the pending software bundle. | [optional] [readonly] 
**PendingVersion** | Pointer to **string** | Version string of the pending software bundle that the Intersight Appliance will install. | [optional] [readonly] 
**SerialId** | Pointer to **string** | SerialId of the Intersight Appliance. SerialId is generated when the Intersight Appliance is setup. SerialId is a unique UUID string, and it will not change for the life time of the Intersight Appliance. | [optional] [readonly] 
**UpgradeBlocked** | Pointer to **bool** | Flag to indicate whether upgrade on this Intersight Appliance is blocked. | [optional] [readonly] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 
**UpgradePolicy** | Pointer to [**NullableApplianceDeviceUpgradePolicyRelationship**](ApplianceDeviceUpgradePolicyRelationship.md) |  | [optional] 

## Methods

### NewApplianceDeviceState

`func NewApplianceDeviceState(classId string, objectType string, ) *ApplianceDeviceState`

NewApplianceDeviceState instantiates a new ApplianceDeviceState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplianceDeviceStateWithDefaults

`func NewApplianceDeviceStateWithDefaults() *ApplianceDeviceState`

NewApplianceDeviceStateWithDefaults instantiates a new ApplianceDeviceState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *ApplianceDeviceState) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ApplianceDeviceState) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ApplianceDeviceState) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *ApplianceDeviceState) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApplianceDeviceState) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApplianceDeviceState) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBlockReasons

`func (o *ApplianceDeviceState) GetBlockReasons() []string`

GetBlockReasons returns the BlockReasons field if non-nil, zero value otherwise.

### GetBlockReasonsOk

`func (o *ApplianceDeviceState) GetBlockReasonsOk() (*[]string, bool)`

GetBlockReasonsOk returns a tuple with the BlockReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockReasons

`func (o *ApplianceDeviceState) SetBlockReasons(v []string)`

SetBlockReasons sets BlockReasons field to given value.

### HasBlockReasons

`func (o *ApplianceDeviceState) HasBlockReasons() bool`

HasBlockReasons returns a boolean if a field has been set.

### SetBlockReasonsNil

`func (o *ApplianceDeviceState) SetBlockReasonsNil(b bool)`

 SetBlockReasonsNil sets the value for BlockReasons to be an explicit nil

### UnsetBlockReasons
`func (o *ApplianceDeviceState) UnsetBlockReasons()`

UnsetBlockReasons ensures that no value is present for BlockReasons, not even an explicit nil
### GetBlockedVersion

`func (o *ApplianceDeviceState) GetBlockedVersion() string`

GetBlockedVersion returns the BlockedVersion field if non-nil, zero value otherwise.

### GetBlockedVersionOk

`func (o *ApplianceDeviceState) GetBlockedVersionOk() (*string, bool)`

GetBlockedVersionOk returns a tuple with the BlockedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedVersion

`func (o *ApplianceDeviceState) SetBlockedVersion(v string)`

SetBlockedVersion sets BlockedVersion field to given value.

### HasBlockedVersion

`func (o *ApplianceDeviceState) HasBlockedVersion() bool`

HasBlockedVersion returns a boolean if a field has been set.

### GetCapabilities

`func (o *ApplianceDeviceState) GetCapabilities() []ApplianceKeyValuePair`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *ApplianceDeviceState) GetCapabilitiesOk() (*[]ApplianceKeyValuePair, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *ApplianceDeviceState) SetCapabilities(v []ApplianceKeyValuePair)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *ApplianceDeviceState) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### SetCapabilitiesNil

`func (o *ApplianceDeviceState) SetCapabilitiesNil(b bool)`

 SetCapabilitiesNil sets the value for Capabilities to be an explicit nil

### UnsetCapabilities
`func (o *ApplianceDeviceState) UnsetCapabilities()`

UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil
### GetCertificate

`func (o *ApplianceDeviceState) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *ApplianceDeviceState) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *ApplianceDeviceState) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *ApplianceDeviceState) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetCertificateNotAfter

`func (o *ApplianceDeviceState) GetCertificateNotAfter() time.Time`

GetCertificateNotAfter returns the CertificateNotAfter field if non-nil, zero value otherwise.

### GetCertificateNotAfterOk

`func (o *ApplianceDeviceState) GetCertificateNotAfterOk() (*time.Time, bool)`

GetCertificateNotAfterOk returns a tuple with the CertificateNotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateNotAfter

`func (o *ApplianceDeviceState) SetCertificateNotAfter(v time.Time)`

SetCertificateNotAfter sets CertificateNotAfter field to given value.

### HasCertificateNotAfter

`func (o *ApplianceDeviceState) HasCertificateNotAfter() bool`

HasCertificateNotAfter returns a boolean if a field has been set.

### GetClusterInfo

`func (o *ApplianceDeviceState) GetClusterInfo() OnpremClusterInfo`

GetClusterInfo returns the ClusterInfo field if non-nil, zero value otherwise.

### GetClusterInfoOk

`func (o *ApplianceDeviceState) GetClusterInfoOk() (*OnpremClusterInfo, bool)`

GetClusterInfoOk returns a tuple with the ClusterInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterInfo

`func (o *ApplianceDeviceState) SetClusterInfo(v OnpremClusterInfo)`

SetClusterInfo sets ClusterInfo field to given value.

### HasClusterInfo

`func (o *ApplianceDeviceState) HasClusterInfo() bool`

HasClusterInfo returns a boolean if a field has been set.

### SetClusterInfoNil

`func (o *ApplianceDeviceState) SetClusterInfoNil(b bool)`

 SetClusterInfoNil sets the value for ClusterInfo to be an explicit nil

### UnsetClusterInfo
`func (o *ApplianceDeviceState) UnsetClusterInfo()`

UnsetClusterInfo ensures that no value is present for ClusterInfo, not even an explicit nil
### GetConnectionStatus

`func (o *ApplianceDeviceState) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *ApplianceDeviceState) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *ApplianceDeviceState) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *ApplianceDeviceState) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetCurrentFingerprint

`func (o *ApplianceDeviceState) GetCurrentFingerprint() string`

GetCurrentFingerprint returns the CurrentFingerprint field if non-nil, zero value otherwise.

### GetCurrentFingerprintOk

`func (o *ApplianceDeviceState) GetCurrentFingerprintOk() (*string, bool)`

GetCurrentFingerprintOk returns a tuple with the CurrentFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentFingerprint

`func (o *ApplianceDeviceState) SetCurrentFingerprint(v string)`

SetCurrentFingerprint sets CurrentFingerprint field to given value.

### HasCurrentFingerprint

`func (o *ApplianceDeviceState) HasCurrentFingerprint() bool`

HasCurrentFingerprint returns a boolean if a field has been set.

### GetCurrentVersion

`func (o *ApplianceDeviceState) GetCurrentVersion() string`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *ApplianceDeviceState) GetCurrentVersionOk() (*string, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *ApplianceDeviceState) SetCurrentVersion(v string)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *ApplianceDeviceState) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetDcVersion

`func (o *ApplianceDeviceState) GetDcVersion() string`

GetDcVersion returns the DcVersion field if non-nil, zero value otherwise.

### GetDcVersionOk

`func (o *ApplianceDeviceState) GetDcVersionOk() (*string, bool)`

GetDcVersionOk returns a tuple with the DcVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcVersion

`func (o *ApplianceDeviceState) SetDcVersion(v string)`

SetDcVersion sets DcVersion field to given value.

### HasDcVersion

`func (o *ApplianceDeviceState) HasDcVersion() bool`

HasDcVersion returns a boolean if a field has been set.

### GetDesiredVersion

`func (o *ApplianceDeviceState) GetDesiredVersion() string`

GetDesiredVersion returns the DesiredVersion field if non-nil, zero value otherwise.

### GetDesiredVersionOk

`func (o *ApplianceDeviceState) GetDesiredVersionOk() (*string, bool)`

GetDesiredVersionOk returns a tuple with the DesiredVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredVersion

`func (o *ApplianceDeviceState) SetDesiredVersion(v string)`

SetDesiredVersion sets DesiredVersion field to given value.

### HasDesiredVersion

`func (o *ApplianceDeviceState) HasDesiredVersion() bool`

HasDesiredVersion returns a boolean if a field has been set.

### GetDownloadedFingerprint

`func (o *ApplianceDeviceState) GetDownloadedFingerprint() string`

GetDownloadedFingerprint returns the DownloadedFingerprint field if non-nil, zero value otherwise.

### GetDownloadedFingerprintOk

`func (o *ApplianceDeviceState) GetDownloadedFingerprintOk() (*string, bool)`

GetDownloadedFingerprintOk returns a tuple with the DownloadedFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadedFingerprint

`func (o *ApplianceDeviceState) SetDownloadedFingerprint(v string)`

SetDownloadedFingerprint sets DownloadedFingerprint field to given value.

### HasDownloadedFingerprint

`func (o *ApplianceDeviceState) HasDownloadedFingerprint() bool`

HasDownloadedFingerprint returns a boolean if a field has been set.

### GetDownloadedMetadataVersion

`func (o *ApplianceDeviceState) GetDownloadedMetadataVersion() []ApplianceMetadataManifestVersion`

GetDownloadedMetadataVersion returns the DownloadedMetadataVersion field if non-nil, zero value otherwise.

### GetDownloadedMetadataVersionOk

`func (o *ApplianceDeviceState) GetDownloadedMetadataVersionOk() (*[]ApplianceMetadataManifestVersion, bool)`

GetDownloadedMetadataVersionOk returns a tuple with the DownloadedMetadataVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadedMetadataVersion

`func (o *ApplianceDeviceState) SetDownloadedMetadataVersion(v []ApplianceMetadataManifestVersion)`

SetDownloadedMetadataVersion sets DownloadedMetadataVersion field to given value.

### HasDownloadedMetadataVersion

`func (o *ApplianceDeviceState) HasDownloadedMetadataVersion() bool`

HasDownloadedMetadataVersion returns a boolean if a field has been set.

### SetDownloadedMetadataVersionNil

`func (o *ApplianceDeviceState) SetDownloadedMetadataVersionNil(b bool)`

 SetDownloadedMetadataVersionNil sets the value for DownloadedMetadataVersion to be an explicit nil

### UnsetDownloadedMetadataVersion
`func (o *ApplianceDeviceState) UnsetDownloadedMetadataVersion()`

UnsetDownloadedMetadataVersion ensures that no value is present for DownloadedMetadataVersion, not even an explicit nil
### GetDownloadedVersion

`func (o *ApplianceDeviceState) GetDownloadedVersion() string`

GetDownloadedVersion returns the DownloadedVersion field if non-nil, zero value otherwise.

### GetDownloadedVersionOk

`func (o *ApplianceDeviceState) GetDownloadedVersionOk() (*string, bool)`

GetDownloadedVersionOk returns a tuple with the DownloadedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadedVersion

`func (o *ApplianceDeviceState) SetDownloadedVersion(v string)`

SetDownloadedVersion sets DownloadedVersion field to given value.

### HasDownloadedVersion

`func (o *ApplianceDeviceState) HasDownloadedVersion() bool`

HasDownloadedVersion returns a boolean if a field has been set.

### GetHostname

`func (o *ApplianceDeviceState) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ApplianceDeviceState) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ApplianceDeviceState) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ApplianceDeviceState) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetPendingFingerprint

`func (o *ApplianceDeviceState) GetPendingFingerprint() string`

GetPendingFingerprint returns the PendingFingerprint field if non-nil, zero value otherwise.

### GetPendingFingerprintOk

`func (o *ApplianceDeviceState) GetPendingFingerprintOk() (*string, bool)`

GetPendingFingerprintOk returns a tuple with the PendingFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingFingerprint

`func (o *ApplianceDeviceState) SetPendingFingerprint(v string)`

SetPendingFingerprint sets PendingFingerprint field to given value.

### HasPendingFingerprint

`func (o *ApplianceDeviceState) HasPendingFingerprint() bool`

HasPendingFingerprint returns a boolean if a field has been set.

### GetPendingVersion

`func (o *ApplianceDeviceState) GetPendingVersion() string`

GetPendingVersion returns the PendingVersion field if non-nil, zero value otherwise.

### GetPendingVersionOk

`func (o *ApplianceDeviceState) GetPendingVersionOk() (*string, bool)`

GetPendingVersionOk returns a tuple with the PendingVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingVersion

`func (o *ApplianceDeviceState) SetPendingVersion(v string)`

SetPendingVersion sets PendingVersion field to given value.

### HasPendingVersion

`func (o *ApplianceDeviceState) HasPendingVersion() bool`

HasPendingVersion returns a boolean if a field has been set.

### GetSerialId

`func (o *ApplianceDeviceState) GetSerialId() string`

GetSerialId returns the SerialId field if non-nil, zero value otherwise.

### GetSerialIdOk

`func (o *ApplianceDeviceState) GetSerialIdOk() (*string, bool)`

GetSerialIdOk returns a tuple with the SerialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialId

`func (o *ApplianceDeviceState) SetSerialId(v string)`

SetSerialId sets SerialId field to given value.

### HasSerialId

`func (o *ApplianceDeviceState) HasSerialId() bool`

HasSerialId returns a boolean if a field has been set.

### GetUpgradeBlocked

`func (o *ApplianceDeviceState) GetUpgradeBlocked() bool`

GetUpgradeBlocked returns the UpgradeBlocked field if non-nil, zero value otherwise.

### GetUpgradeBlockedOk

`func (o *ApplianceDeviceState) GetUpgradeBlockedOk() (*bool, bool)`

GetUpgradeBlockedOk returns a tuple with the UpgradeBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeBlocked

`func (o *ApplianceDeviceState) SetUpgradeBlocked(v bool)`

SetUpgradeBlocked sets UpgradeBlocked field to given value.

### HasUpgradeBlocked

`func (o *ApplianceDeviceState) HasUpgradeBlocked() bool`

HasUpgradeBlocked returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *ApplianceDeviceState) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *ApplianceDeviceState) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *ApplianceDeviceState) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *ApplianceDeviceState) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *ApplianceDeviceState) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *ApplianceDeviceState) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil
### GetUpgradePolicy

`func (o *ApplianceDeviceState) GetUpgradePolicy() ApplianceDeviceUpgradePolicyRelationship`

GetUpgradePolicy returns the UpgradePolicy field if non-nil, zero value otherwise.

### GetUpgradePolicyOk

`func (o *ApplianceDeviceState) GetUpgradePolicyOk() (*ApplianceDeviceUpgradePolicyRelationship, bool)`

GetUpgradePolicyOk returns a tuple with the UpgradePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradePolicy

`func (o *ApplianceDeviceState) SetUpgradePolicy(v ApplianceDeviceUpgradePolicyRelationship)`

SetUpgradePolicy sets UpgradePolicy field to given value.

### HasUpgradePolicy

`func (o *ApplianceDeviceState) HasUpgradePolicy() bool`

HasUpgradePolicy returns a boolean if a field has been set.

### SetUpgradePolicyNil

`func (o *ApplianceDeviceState) SetUpgradePolicyNil(b bool)`

 SetUpgradePolicyNil sets the value for UpgradePolicy to be an explicit nil

### UnsetUpgradePolicy
`func (o *ApplianceDeviceState) UnsetUpgradePolicy()`

UnsetUpgradePolicy ensures that no value is present for UpgradePolicy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


