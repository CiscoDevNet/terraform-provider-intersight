# HciEsxiGuestTools

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hci.EsxiGuestTools"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hci.EsxiGuestTools"]
**AvailableVersion** | Pointer to **string** | The available version of the Nutanix Guest Tools. | [optional] [readonly] 
**Capabilities** | Pointer to **[]string** |  | [optional] 
**GuestOsVersion** | Pointer to **string** | The guest OS version of the Nutanix Guest Tools. | [optional] [readonly] 
**IsEnabled** | Pointer to **bool** | Indicates if the Nutanix Guest Tools is enabled. | [optional] [readonly] 
**IsInstalled** | Pointer to **bool** | Indicates if the Nutanix Guest Tools is installed. | [optional] [readonly] 
**IsIsoInserted** | Pointer to **bool** | Indicates if the Nutanix Guest Tools ISO is inserted. | [optional] [readonly] 
**IsReachable** | Pointer to **bool** | Indicates if the Nutanix Guest Tools is reachable. | [optional] [readonly] 
**IsVmMobilityDriverInstalled** | Pointer to **bool** | Indicates if the Nutanix Guest Tools VM mobility driver is installed. | [optional] [readonly] 
**IsVssSnapshotCapable** | Pointer to **bool** | Indicates if the Nutanix Guest Tools is VSS snapshot capable. | [optional] [readonly] 
**Version** | Pointer to **string** | The version of the Nutanix Guest Tools. | [optional] [readonly] 

## Methods

### NewHciEsxiGuestTools

`func NewHciEsxiGuestTools(classId string, objectType string, ) *HciEsxiGuestTools`

NewHciEsxiGuestTools instantiates a new HciEsxiGuestTools object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHciEsxiGuestToolsWithDefaults

`func NewHciEsxiGuestToolsWithDefaults() *HciEsxiGuestTools`

NewHciEsxiGuestToolsWithDefaults instantiates a new HciEsxiGuestTools object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HciEsxiGuestTools) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HciEsxiGuestTools) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HciEsxiGuestTools) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HciEsxiGuestTools) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HciEsxiGuestTools) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HciEsxiGuestTools) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAvailableVersion

`func (o *HciEsxiGuestTools) GetAvailableVersion() string`

GetAvailableVersion returns the AvailableVersion field if non-nil, zero value otherwise.

### GetAvailableVersionOk

`func (o *HciEsxiGuestTools) GetAvailableVersionOk() (*string, bool)`

GetAvailableVersionOk returns a tuple with the AvailableVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableVersion

`func (o *HciEsxiGuestTools) SetAvailableVersion(v string)`

SetAvailableVersion sets AvailableVersion field to given value.

### HasAvailableVersion

`func (o *HciEsxiGuestTools) HasAvailableVersion() bool`

HasAvailableVersion returns a boolean if a field has been set.

### GetCapabilities

`func (o *HciEsxiGuestTools) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *HciEsxiGuestTools) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *HciEsxiGuestTools) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *HciEsxiGuestTools) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### SetCapabilitiesNil

`func (o *HciEsxiGuestTools) SetCapabilitiesNil(b bool)`

 SetCapabilitiesNil sets the value for Capabilities to be an explicit nil

### UnsetCapabilities
`func (o *HciEsxiGuestTools) UnsetCapabilities()`

UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil
### GetGuestOsVersion

`func (o *HciEsxiGuestTools) GetGuestOsVersion() string`

GetGuestOsVersion returns the GuestOsVersion field if non-nil, zero value otherwise.

### GetGuestOsVersionOk

`func (o *HciEsxiGuestTools) GetGuestOsVersionOk() (*string, bool)`

GetGuestOsVersionOk returns a tuple with the GuestOsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestOsVersion

`func (o *HciEsxiGuestTools) SetGuestOsVersion(v string)`

SetGuestOsVersion sets GuestOsVersion field to given value.

### HasGuestOsVersion

`func (o *HciEsxiGuestTools) HasGuestOsVersion() bool`

HasGuestOsVersion returns a boolean if a field has been set.

### GetIsEnabled

`func (o *HciEsxiGuestTools) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *HciEsxiGuestTools) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *HciEsxiGuestTools) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *HciEsxiGuestTools) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetIsInstalled

`func (o *HciEsxiGuestTools) GetIsInstalled() bool`

GetIsInstalled returns the IsInstalled field if non-nil, zero value otherwise.

### GetIsInstalledOk

`func (o *HciEsxiGuestTools) GetIsInstalledOk() (*bool, bool)`

GetIsInstalledOk returns a tuple with the IsInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInstalled

`func (o *HciEsxiGuestTools) SetIsInstalled(v bool)`

SetIsInstalled sets IsInstalled field to given value.

### HasIsInstalled

`func (o *HciEsxiGuestTools) HasIsInstalled() bool`

HasIsInstalled returns a boolean if a field has been set.

### GetIsIsoInserted

`func (o *HciEsxiGuestTools) GetIsIsoInserted() bool`

GetIsIsoInserted returns the IsIsoInserted field if non-nil, zero value otherwise.

### GetIsIsoInsertedOk

`func (o *HciEsxiGuestTools) GetIsIsoInsertedOk() (*bool, bool)`

GetIsIsoInsertedOk returns a tuple with the IsIsoInserted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIsoInserted

`func (o *HciEsxiGuestTools) SetIsIsoInserted(v bool)`

SetIsIsoInserted sets IsIsoInserted field to given value.

### HasIsIsoInserted

`func (o *HciEsxiGuestTools) HasIsIsoInserted() bool`

HasIsIsoInserted returns a boolean if a field has been set.

### GetIsReachable

`func (o *HciEsxiGuestTools) GetIsReachable() bool`

GetIsReachable returns the IsReachable field if non-nil, zero value otherwise.

### GetIsReachableOk

`func (o *HciEsxiGuestTools) GetIsReachableOk() (*bool, bool)`

GetIsReachableOk returns a tuple with the IsReachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReachable

`func (o *HciEsxiGuestTools) SetIsReachable(v bool)`

SetIsReachable sets IsReachable field to given value.

### HasIsReachable

`func (o *HciEsxiGuestTools) HasIsReachable() bool`

HasIsReachable returns a boolean if a field has been set.

### GetIsVmMobilityDriverInstalled

`func (o *HciEsxiGuestTools) GetIsVmMobilityDriverInstalled() bool`

GetIsVmMobilityDriverInstalled returns the IsVmMobilityDriverInstalled field if non-nil, zero value otherwise.

### GetIsVmMobilityDriverInstalledOk

`func (o *HciEsxiGuestTools) GetIsVmMobilityDriverInstalledOk() (*bool, bool)`

GetIsVmMobilityDriverInstalledOk returns a tuple with the IsVmMobilityDriverInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVmMobilityDriverInstalled

`func (o *HciEsxiGuestTools) SetIsVmMobilityDriverInstalled(v bool)`

SetIsVmMobilityDriverInstalled sets IsVmMobilityDriverInstalled field to given value.

### HasIsVmMobilityDriverInstalled

`func (o *HciEsxiGuestTools) HasIsVmMobilityDriverInstalled() bool`

HasIsVmMobilityDriverInstalled returns a boolean if a field has been set.

### GetIsVssSnapshotCapable

`func (o *HciEsxiGuestTools) GetIsVssSnapshotCapable() bool`

GetIsVssSnapshotCapable returns the IsVssSnapshotCapable field if non-nil, zero value otherwise.

### GetIsVssSnapshotCapableOk

`func (o *HciEsxiGuestTools) GetIsVssSnapshotCapableOk() (*bool, bool)`

GetIsVssSnapshotCapableOk returns a tuple with the IsVssSnapshotCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVssSnapshotCapable

`func (o *HciEsxiGuestTools) SetIsVssSnapshotCapable(v bool)`

SetIsVssSnapshotCapable sets IsVssSnapshotCapable field to given value.

### HasIsVssSnapshotCapable

`func (o *HciEsxiGuestTools) HasIsVssSnapshotCapable() bool`

HasIsVssSnapshotCapable returns a boolean if a field has been set.

### GetVersion

`func (o *HciEsxiGuestTools) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HciEsxiGuestTools) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HciEsxiGuestTools) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HciEsxiGuestTools) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


