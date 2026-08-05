# NiatelemetryFabric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.Fabric"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.Fabric"]
**Advisories** | Pointer to [**[]NiatelemetryAdvisories**](NiatelemetryAdvisories.md) |  | [optional] 
**LicenseTier** | Pointer to **string** | The license tier of the fabric. | [optional] [readonly] 
**Name** | Pointer to **string** | Returns the name of the fabric. | [optional] [readonly] 
**NetworkCount** | Pointer to **int64** | The network count of the fabric. | [optional] [readonly] 
**OwnerClusterName** | Pointer to **string** | The name of the cluster this fabric belongs to. | [optional] [readonly] 
**PowerStatus** | Pointer to **string** | The power status of the fabric. | [optional] [readonly] 
**PowerUnit** | Pointer to **string** | The power unit of the fabric. | [optional] [readonly] 
**PowerValue** | Pointer to **string** | The power value of the fabric. | [optional] [readonly] 
**SecurityDomain** | Pointer to **string** | The security domain of the fabric. | [optional] [readonly] 
**TelemetryEnabled** | Pointer to **bool** | Indicates whether telemetry is enabled for the fabric. | [optional] [readonly] 
**VrfCount** | Pointer to **int64** | The VRF count of the fabric. | [optional] [readonly] 
**Cluster** | Pointer to [**NullableNiatelemetryClusterRelationship**](NiatelemetryClusterRelationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryFabric

`func NewNiatelemetryFabric(classId string, objectType string, ) *NiatelemetryFabric`

NewNiatelemetryFabric instantiates a new NiatelemetryFabric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryFabricWithDefaults

`func NewNiatelemetryFabricWithDefaults() *NiatelemetryFabric`

NewNiatelemetryFabricWithDefaults instantiates a new NiatelemetryFabric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryFabric) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryFabric) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryFabric) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryFabric) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryFabric) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryFabric) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAdvisories

`func (o *NiatelemetryFabric) GetAdvisories() []NiatelemetryAdvisories`

GetAdvisories returns the Advisories field if non-nil, zero value otherwise.

### GetAdvisoriesOk

`func (o *NiatelemetryFabric) GetAdvisoriesOk() (*[]NiatelemetryAdvisories, bool)`

GetAdvisoriesOk returns a tuple with the Advisories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvisories

`func (o *NiatelemetryFabric) SetAdvisories(v []NiatelemetryAdvisories)`

SetAdvisories sets Advisories field to given value.

### HasAdvisories

`func (o *NiatelemetryFabric) HasAdvisories() bool`

HasAdvisories returns a boolean if a field has been set.

### SetAdvisoriesNil

`func (o *NiatelemetryFabric) SetAdvisoriesNil(b bool)`

 SetAdvisoriesNil sets the value for Advisories to be an explicit nil

### UnsetAdvisories
`func (o *NiatelemetryFabric) UnsetAdvisories()`

UnsetAdvisories ensures that no value is present for Advisories, not even an explicit nil
### GetLicenseTier

`func (o *NiatelemetryFabric) GetLicenseTier() string`

GetLicenseTier returns the LicenseTier field if non-nil, zero value otherwise.

### GetLicenseTierOk

`func (o *NiatelemetryFabric) GetLicenseTierOk() (*string, bool)`

GetLicenseTierOk returns a tuple with the LicenseTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseTier

`func (o *NiatelemetryFabric) SetLicenseTier(v string)`

SetLicenseTier sets LicenseTier field to given value.

### HasLicenseTier

`func (o *NiatelemetryFabric) HasLicenseTier() bool`

HasLicenseTier returns a boolean if a field has been set.

### GetName

`func (o *NiatelemetryFabric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NiatelemetryFabric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NiatelemetryFabric) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NiatelemetryFabric) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkCount

`func (o *NiatelemetryFabric) GetNetworkCount() int64`

GetNetworkCount returns the NetworkCount field if non-nil, zero value otherwise.

### GetNetworkCountOk

`func (o *NiatelemetryFabric) GetNetworkCountOk() (*int64, bool)`

GetNetworkCountOk returns a tuple with the NetworkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkCount

`func (o *NiatelemetryFabric) SetNetworkCount(v int64)`

SetNetworkCount sets NetworkCount field to given value.

### HasNetworkCount

`func (o *NiatelemetryFabric) HasNetworkCount() bool`

HasNetworkCount returns a boolean if a field has been set.

### GetOwnerClusterName

`func (o *NiatelemetryFabric) GetOwnerClusterName() string`

GetOwnerClusterName returns the OwnerClusterName field if non-nil, zero value otherwise.

### GetOwnerClusterNameOk

`func (o *NiatelemetryFabric) GetOwnerClusterNameOk() (*string, bool)`

GetOwnerClusterNameOk returns a tuple with the OwnerClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerClusterName

`func (o *NiatelemetryFabric) SetOwnerClusterName(v string)`

SetOwnerClusterName sets OwnerClusterName field to given value.

### HasOwnerClusterName

`func (o *NiatelemetryFabric) HasOwnerClusterName() bool`

HasOwnerClusterName returns a boolean if a field has been set.

### GetPowerStatus

`func (o *NiatelemetryFabric) GetPowerStatus() string`

GetPowerStatus returns the PowerStatus field if non-nil, zero value otherwise.

### GetPowerStatusOk

`func (o *NiatelemetryFabric) GetPowerStatusOk() (*string, bool)`

GetPowerStatusOk returns a tuple with the PowerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerStatus

`func (o *NiatelemetryFabric) SetPowerStatus(v string)`

SetPowerStatus sets PowerStatus field to given value.

### HasPowerStatus

`func (o *NiatelemetryFabric) HasPowerStatus() bool`

HasPowerStatus returns a boolean if a field has been set.

### GetPowerUnit

`func (o *NiatelemetryFabric) GetPowerUnit() string`

GetPowerUnit returns the PowerUnit field if non-nil, zero value otherwise.

### GetPowerUnitOk

`func (o *NiatelemetryFabric) GetPowerUnitOk() (*string, bool)`

GetPowerUnitOk returns a tuple with the PowerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerUnit

`func (o *NiatelemetryFabric) SetPowerUnit(v string)`

SetPowerUnit sets PowerUnit field to given value.

### HasPowerUnit

`func (o *NiatelemetryFabric) HasPowerUnit() bool`

HasPowerUnit returns a boolean if a field has been set.

### GetPowerValue

`func (o *NiatelemetryFabric) GetPowerValue() string`

GetPowerValue returns the PowerValue field if non-nil, zero value otherwise.

### GetPowerValueOk

`func (o *NiatelemetryFabric) GetPowerValueOk() (*string, bool)`

GetPowerValueOk returns a tuple with the PowerValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerValue

`func (o *NiatelemetryFabric) SetPowerValue(v string)`

SetPowerValue sets PowerValue field to given value.

### HasPowerValue

`func (o *NiatelemetryFabric) HasPowerValue() bool`

HasPowerValue returns a boolean if a field has been set.

### GetSecurityDomain

`func (o *NiatelemetryFabric) GetSecurityDomain() string`

GetSecurityDomain returns the SecurityDomain field if non-nil, zero value otherwise.

### GetSecurityDomainOk

`func (o *NiatelemetryFabric) GetSecurityDomainOk() (*string, bool)`

GetSecurityDomainOk returns a tuple with the SecurityDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityDomain

`func (o *NiatelemetryFabric) SetSecurityDomain(v string)`

SetSecurityDomain sets SecurityDomain field to given value.

### HasSecurityDomain

`func (o *NiatelemetryFabric) HasSecurityDomain() bool`

HasSecurityDomain returns a boolean if a field has been set.

### GetTelemetryEnabled

`func (o *NiatelemetryFabric) GetTelemetryEnabled() bool`

GetTelemetryEnabled returns the TelemetryEnabled field if non-nil, zero value otherwise.

### GetTelemetryEnabledOk

`func (o *NiatelemetryFabric) GetTelemetryEnabledOk() (*bool, bool)`

GetTelemetryEnabledOk returns a tuple with the TelemetryEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelemetryEnabled

`func (o *NiatelemetryFabric) SetTelemetryEnabled(v bool)`

SetTelemetryEnabled sets TelemetryEnabled field to given value.

### HasTelemetryEnabled

`func (o *NiatelemetryFabric) HasTelemetryEnabled() bool`

HasTelemetryEnabled returns a boolean if a field has been set.

### GetVrfCount

`func (o *NiatelemetryFabric) GetVrfCount() int64`

GetVrfCount returns the VrfCount field if non-nil, zero value otherwise.

### GetVrfCountOk

`func (o *NiatelemetryFabric) GetVrfCountOk() (*int64, bool)`

GetVrfCountOk returns a tuple with the VrfCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfCount

`func (o *NiatelemetryFabric) SetVrfCount(v int64)`

SetVrfCount sets VrfCount field to given value.

### HasVrfCount

`func (o *NiatelemetryFabric) HasVrfCount() bool`

HasVrfCount returns a boolean if a field has been set.

### GetCluster

`func (o *NiatelemetryFabric) GetCluster() NiatelemetryClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *NiatelemetryFabric) GetClusterOk() (*NiatelemetryClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *NiatelemetryFabric) SetCluster(v NiatelemetryClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *NiatelemetryFabric) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### SetClusterNil

`func (o *NiatelemetryFabric) SetClusterNil(b bool)`

 SetClusterNil sets the value for Cluster to be an explicit nil

### UnsetCluster
`func (o *NiatelemetryFabric) UnsetCluster()`

UnsetCluster ensures that no value is present for Cluster, not even an explicit nil
### GetRegisteredDevice

`func (o *NiatelemetryFabric) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryFabric) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryFabric) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryFabric) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *NiatelemetryFabric) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *NiatelemetryFabric) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


