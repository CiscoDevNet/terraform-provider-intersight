# NiatelemetryEpg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzurePackCount** | Pointer to **int64** | Azure Pack NAT with ASA feature usage. | [optional] 
**Dn** | Pointer to **string** | Dn value for the End Point Groups present. | [optional] 
**EpgDelimiterCount** | Pointer to **int64** | EPG Delimiter scale where the delimiter value is present. | [optional] 
**FcNpvCount** | Pointer to **int64** | Number of ports with FC path attribute of type FC. | [optional] 
**FcoeCount** | Pointer to **string** | Number of FCoE per End Point Group. | [optional] 
**FvRsDomAttCount** | Pointer to **int64** | FvRsDomAtt scale per End Point Group with VMware configured. | [optional] 
**IntraEpgDvsBmCount** | Pointer to **int64** | Intra End Point Group Contract for Distributed Virtual Switch and BM feature usage. | [optional] 
**IntraEpgHyperv** | Pointer to **string** | Intra EPG Isolation for Hyper-V, enabled if pcEnfPref attribute is set to enforced. | [optional] 
**IsAttrBased** | Pointer to **string** | Gets the state of End Point Groups with isAttrBasedEPg value as configured. | [optional] 
**Microsegmentation** | Pointer to **string** | Gets the state of End Point Groups where microsegmentation is present. | [optional] 
**MicrosoftUsegCount** | Pointer to **int64** | FvRsDomAtt scale per End Point Group with Microsoft configured. | [optional] 
**Name** | Pointer to **string** | Name value for the End Point Groups present. | [optional] 
**OrchslDevVipCfgCount** | Pointer to **int64** | Simplified Service Graph Integration with Windows Azure Pack count scale. | [optional] 
**SiteName** | Pointer to **string** | The Site name represents an APIC cluster. Service Engine can onboard multiple APIC clusters / sites. | [optional] 
**UsegHypervCount** | Pointer to **int64** | Logical Operators for attribute based microsegmentation in a hypervisor. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewNiatelemetryEpg

`func NewNiatelemetryEpg() *NiatelemetryEpg`

NewNiatelemetryEpg instantiates a new NiatelemetryEpg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryEpgWithDefaults

`func NewNiatelemetryEpgWithDefaults() *NiatelemetryEpg`

NewNiatelemetryEpgWithDefaults instantiates a new NiatelemetryEpg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzurePackCount

`func (o *NiatelemetryEpg) GetAzurePackCount() int64`

GetAzurePackCount returns the AzurePackCount field if non-nil, zero value otherwise.

### GetAzurePackCountOk

`func (o *NiatelemetryEpg) GetAzurePackCountOk() (*int64, bool)`

GetAzurePackCountOk returns a tuple with the AzurePackCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzurePackCount

`func (o *NiatelemetryEpg) SetAzurePackCount(v int64)`

SetAzurePackCount sets AzurePackCount field to given value.

### HasAzurePackCount

`func (o *NiatelemetryEpg) HasAzurePackCount() bool`

HasAzurePackCount returns a boolean if a field has been set.

### GetDn

`func (o *NiatelemetryEpg) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetryEpg) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetryEpg) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetryEpg) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetEpgDelimiterCount

`func (o *NiatelemetryEpg) GetEpgDelimiterCount() int64`

GetEpgDelimiterCount returns the EpgDelimiterCount field if non-nil, zero value otherwise.

### GetEpgDelimiterCountOk

`func (o *NiatelemetryEpg) GetEpgDelimiterCountOk() (*int64, bool)`

GetEpgDelimiterCountOk returns a tuple with the EpgDelimiterCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpgDelimiterCount

`func (o *NiatelemetryEpg) SetEpgDelimiterCount(v int64)`

SetEpgDelimiterCount sets EpgDelimiterCount field to given value.

### HasEpgDelimiterCount

`func (o *NiatelemetryEpg) HasEpgDelimiterCount() bool`

HasEpgDelimiterCount returns a boolean if a field has been set.

### GetFcNpvCount

`func (o *NiatelemetryEpg) GetFcNpvCount() int64`

GetFcNpvCount returns the FcNpvCount field if non-nil, zero value otherwise.

### GetFcNpvCountOk

`func (o *NiatelemetryEpg) GetFcNpvCountOk() (*int64, bool)`

GetFcNpvCountOk returns a tuple with the FcNpvCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcNpvCount

`func (o *NiatelemetryEpg) SetFcNpvCount(v int64)`

SetFcNpvCount sets FcNpvCount field to given value.

### HasFcNpvCount

`func (o *NiatelemetryEpg) HasFcNpvCount() bool`

HasFcNpvCount returns a boolean if a field has been set.

### GetFcoeCount

`func (o *NiatelemetryEpg) GetFcoeCount() string`

GetFcoeCount returns the FcoeCount field if non-nil, zero value otherwise.

### GetFcoeCountOk

`func (o *NiatelemetryEpg) GetFcoeCountOk() (*string, bool)`

GetFcoeCountOk returns a tuple with the FcoeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcoeCount

`func (o *NiatelemetryEpg) SetFcoeCount(v string)`

SetFcoeCount sets FcoeCount field to given value.

### HasFcoeCount

`func (o *NiatelemetryEpg) HasFcoeCount() bool`

HasFcoeCount returns a boolean if a field has been set.

### GetFvRsDomAttCount

`func (o *NiatelemetryEpg) GetFvRsDomAttCount() int64`

GetFvRsDomAttCount returns the FvRsDomAttCount field if non-nil, zero value otherwise.

### GetFvRsDomAttCountOk

`func (o *NiatelemetryEpg) GetFvRsDomAttCountOk() (*int64, bool)`

GetFvRsDomAttCountOk returns a tuple with the FvRsDomAttCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFvRsDomAttCount

`func (o *NiatelemetryEpg) SetFvRsDomAttCount(v int64)`

SetFvRsDomAttCount sets FvRsDomAttCount field to given value.

### HasFvRsDomAttCount

`func (o *NiatelemetryEpg) HasFvRsDomAttCount() bool`

HasFvRsDomAttCount returns a boolean if a field has been set.

### GetIntraEpgDvsBmCount

`func (o *NiatelemetryEpg) GetIntraEpgDvsBmCount() int64`

GetIntraEpgDvsBmCount returns the IntraEpgDvsBmCount field if non-nil, zero value otherwise.

### GetIntraEpgDvsBmCountOk

`func (o *NiatelemetryEpg) GetIntraEpgDvsBmCountOk() (*int64, bool)`

GetIntraEpgDvsBmCountOk returns a tuple with the IntraEpgDvsBmCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntraEpgDvsBmCount

`func (o *NiatelemetryEpg) SetIntraEpgDvsBmCount(v int64)`

SetIntraEpgDvsBmCount sets IntraEpgDvsBmCount field to given value.

### HasIntraEpgDvsBmCount

`func (o *NiatelemetryEpg) HasIntraEpgDvsBmCount() bool`

HasIntraEpgDvsBmCount returns a boolean if a field has been set.

### GetIntraEpgHyperv

`func (o *NiatelemetryEpg) GetIntraEpgHyperv() string`

GetIntraEpgHyperv returns the IntraEpgHyperv field if non-nil, zero value otherwise.

### GetIntraEpgHypervOk

`func (o *NiatelemetryEpg) GetIntraEpgHypervOk() (*string, bool)`

GetIntraEpgHypervOk returns a tuple with the IntraEpgHyperv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntraEpgHyperv

`func (o *NiatelemetryEpg) SetIntraEpgHyperv(v string)`

SetIntraEpgHyperv sets IntraEpgHyperv field to given value.

### HasIntraEpgHyperv

`func (o *NiatelemetryEpg) HasIntraEpgHyperv() bool`

HasIntraEpgHyperv returns a boolean if a field has been set.

### GetIsAttrBased

`func (o *NiatelemetryEpg) GetIsAttrBased() string`

GetIsAttrBased returns the IsAttrBased field if non-nil, zero value otherwise.

### GetIsAttrBasedOk

`func (o *NiatelemetryEpg) GetIsAttrBasedOk() (*string, bool)`

GetIsAttrBasedOk returns a tuple with the IsAttrBased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAttrBased

`func (o *NiatelemetryEpg) SetIsAttrBased(v string)`

SetIsAttrBased sets IsAttrBased field to given value.

### HasIsAttrBased

`func (o *NiatelemetryEpg) HasIsAttrBased() bool`

HasIsAttrBased returns a boolean if a field has been set.

### GetMicrosegmentation

`func (o *NiatelemetryEpg) GetMicrosegmentation() string`

GetMicrosegmentation returns the Microsegmentation field if non-nil, zero value otherwise.

### GetMicrosegmentationOk

`func (o *NiatelemetryEpg) GetMicrosegmentationOk() (*string, bool)`

GetMicrosegmentationOk returns a tuple with the Microsegmentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosegmentation

`func (o *NiatelemetryEpg) SetMicrosegmentation(v string)`

SetMicrosegmentation sets Microsegmentation field to given value.

### HasMicrosegmentation

`func (o *NiatelemetryEpg) HasMicrosegmentation() bool`

HasMicrosegmentation returns a boolean if a field has been set.

### GetMicrosoftUsegCount

`func (o *NiatelemetryEpg) GetMicrosoftUsegCount() int64`

GetMicrosoftUsegCount returns the MicrosoftUsegCount field if non-nil, zero value otherwise.

### GetMicrosoftUsegCountOk

`func (o *NiatelemetryEpg) GetMicrosoftUsegCountOk() (*int64, bool)`

GetMicrosoftUsegCountOk returns a tuple with the MicrosoftUsegCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosoftUsegCount

`func (o *NiatelemetryEpg) SetMicrosoftUsegCount(v int64)`

SetMicrosoftUsegCount sets MicrosoftUsegCount field to given value.

### HasMicrosoftUsegCount

`func (o *NiatelemetryEpg) HasMicrosoftUsegCount() bool`

HasMicrosoftUsegCount returns a boolean if a field has been set.

### GetName

`func (o *NiatelemetryEpg) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NiatelemetryEpg) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NiatelemetryEpg) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NiatelemetryEpg) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrchslDevVipCfgCount

`func (o *NiatelemetryEpg) GetOrchslDevVipCfgCount() int64`

GetOrchslDevVipCfgCount returns the OrchslDevVipCfgCount field if non-nil, zero value otherwise.

### GetOrchslDevVipCfgCountOk

`func (o *NiatelemetryEpg) GetOrchslDevVipCfgCountOk() (*int64, bool)`

GetOrchslDevVipCfgCountOk returns a tuple with the OrchslDevVipCfgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrchslDevVipCfgCount

`func (o *NiatelemetryEpg) SetOrchslDevVipCfgCount(v int64)`

SetOrchslDevVipCfgCount sets OrchslDevVipCfgCount field to given value.

### HasOrchslDevVipCfgCount

`func (o *NiatelemetryEpg) HasOrchslDevVipCfgCount() bool`

HasOrchslDevVipCfgCount returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryEpg) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryEpg) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryEpg) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryEpg) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetUsegHypervCount

`func (o *NiatelemetryEpg) GetUsegHypervCount() int64`

GetUsegHypervCount returns the UsegHypervCount field if non-nil, zero value otherwise.

### GetUsegHypervCountOk

`func (o *NiatelemetryEpg) GetUsegHypervCountOk() (*int64, bool)`

GetUsegHypervCountOk returns a tuple with the UsegHypervCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsegHypervCount

`func (o *NiatelemetryEpg) SetUsegHypervCount(v int64)`

SetUsegHypervCount sets UsegHypervCount field to given value.

### HasUsegHypervCount

`func (o *NiatelemetryEpg) HasUsegHypervCount() bool`

HasUsegHypervCount returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryEpg) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryEpg) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryEpg) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryEpg) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


