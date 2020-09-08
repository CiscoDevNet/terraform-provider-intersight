# NiatelemetryEpgAllOf

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

### NewNiatelemetryEpgAllOf

`func NewNiatelemetryEpgAllOf() *NiatelemetryEpgAllOf`

NewNiatelemetryEpgAllOf instantiates a new NiatelemetryEpgAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryEpgAllOfWithDefaults

`func NewNiatelemetryEpgAllOfWithDefaults() *NiatelemetryEpgAllOf`

NewNiatelemetryEpgAllOfWithDefaults instantiates a new NiatelemetryEpgAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzurePackCount

`func (o *NiatelemetryEpgAllOf) GetAzurePackCount() int64`

GetAzurePackCount returns the AzurePackCount field if non-nil, zero value otherwise.

### GetAzurePackCountOk

`func (o *NiatelemetryEpgAllOf) GetAzurePackCountOk() (*int64, bool)`

GetAzurePackCountOk returns a tuple with the AzurePackCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzurePackCount

`func (o *NiatelemetryEpgAllOf) SetAzurePackCount(v int64)`

SetAzurePackCount sets AzurePackCount field to given value.

### HasAzurePackCount

`func (o *NiatelemetryEpgAllOf) HasAzurePackCount() bool`

HasAzurePackCount returns a boolean if a field has been set.

### GetDn

`func (o *NiatelemetryEpgAllOf) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetryEpgAllOf) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetryEpgAllOf) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetryEpgAllOf) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetEpgDelimiterCount

`func (o *NiatelemetryEpgAllOf) GetEpgDelimiterCount() int64`

GetEpgDelimiterCount returns the EpgDelimiterCount field if non-nil, zero value otherwise.

### GetEpgDelimiterCountOk

`func (o *NiatelemetryEpgAllOf) GetEpgDelimiterCountOk() (*int64, bool)`

GetEpgDelimiterCountOk returns a tuple with the EpgDelimiterCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpgDelimiterCount

`func (o *NiatelemetryEpgAllOf) SetEpgDelimiterCount(v int64)`

SetEpgDelimiterCount sets EpgDelimiterCount field to given value.

### HasEpgDelimiterCount

`func (o *NiatelemetryEpgAllOf) HasEpgDelimiterCount() bool`

HasEpgDelimiterCount returns a boolean if a field has been set.

### GetFcNpvCount

`func (o *NiatelemetryEpgAllOf) GetFcNpvCount() int64`

GetFcNpvCount returns the FcNpvCount field if non-nil, zero value otherwise.

### GetFcNpvCountOk

`func (o *NiatelemetryEpgAllOf) GetFcNpvCountOk() (*int64, bool)`

GetFcNpvCountOk returns a tuple with the FcNpvCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcNpvCount

`func (o *NiatelemetryEpgAllOf) SetFcNpvCount(v int64)`

SetFcNpvCount sets FcNpvCount field to given value.

### HasFcNpvCount

`func (o *NiatelemetryEpgAllOf) HasFcNpvCount() bool`

HasFcNpvCount returns a boolean if a field has been set.

### GetFcoeCount

`func (o *NiatelemetryEpgAllOf) GetFcoeCount() string`

GetFcoeCount returns the FcoeCount field if non-nil, zero value otherwise.

### GetFcoeCountOk

`func (o *NiatelemetryEpgAllOf) GetFcoeCountOk() (*string, bool)`

GetFcoeCountOk returns a tuple with the FcoeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFcoeCount

`func (o *NiatelemetryEpgAllOf) SetFcoeCount(v string)`

SetFcoeCount sets FcoeCount field to given value.

### HasFcoeCount

`func (o *NiatelemetryEpgAllOf) HasFcoeCount() bool`

HasFcoeCount returns a boolean if a field has been set.

### GetFvRsDomAttCount

`func (o *NiatelemetryEpgAllOf) GetFvRsDomAttCount() int64`

GetFvRsDomAttCount returns the FvRsDomAttCount field if non-nil, zero value otherwise.

### GetFvRsDomAttCountOk

`func (o *NiatelemetryEpgAllOf) GetFvRsDomAttCountOk() (*int64, bool)`

GetFvRsDomAttCountOk returns a tuple with the FvRsDomAttCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFvRsDomAttCount

`func (o *NiatelemetryEpgAllOf) SetFvRsDomAttCount(v int64)`

SetFvRsDomAttCount sets FvRsDomAttCount field to given value.

### HasFvRsDomAttCount

`func (o *NiatelemetryEpgAllOf) HasFvRsDomAttCount() bool`

HasFvRsDomAttCount returns a boolean if a field has been set.

### GetIntraEpgDvsBmCount

`func (o *NiatelemetryEpgAllOf) GetIntraEpgDvsBmCount() int64`

GetIntraEpgDvsBmCount returns the IntraEpgDvsBmCount field if non-nil, zero value otherwise.

### GetIntraEpgDvsBmCountOk

`func (o *NiatelemetryEpgAllOf) GetIntraEpgDvsBmCountOk() (*int64, bool)`

GetIntraEpgDvsBmCountOk returns a tuple with the IntraEpgDvsBmCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntraEpgDvsBmCount

`func (o *NiatelemetryEpgAllOf) SetIntraEpgDvsBmCount(v int64)`

SetIntraEpgDvsBmCount sets IntraEpgDvsBmCount field to given value.

### HasIntraEpgDvsBmCount

`func (o *NiatelemetryEpgAllOf) HasIntraEpgDvsBmCount() bool`

HasIntraEpgDvsBmCount returns a boolean if a field has been set.

### GetIntraEpgHyperv

`func (o *NiatelemetryEpgAllOf) GetIntraEpgHyperv() string`

GetIntraEpgHyperv returns the IntraEpgHyperv field if non-nil, zero value otherwise.

### GetIntraEpgHypervOk

`func (o *NiatelemetryEpgAllOf) GetIntraEpgHypervOk() (*string, bool)`

GetIntraEpgHypervOk returns a tuple with the IntraEpgHyperv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntraEpgHyperv

`func (o *NiatelemetryEpgAllOf) SetIntraEpgHyperv(v string)`

SetIntraEpgHyperv sets IntraEpgHyperv field to given value.

### HasIntraEpgHyperv

`func (o *NiatelemetryEpgAllOf) HasIntraEpgHyperv() bool`

HasIntraEpgHyperv returns a boolean if a field has been set.

### GetIsAttrBased

`func (o *NiatelemetryEpgAllOf) GetIsAttrBased() string`

GetIsAttrBased returns the IsAttrBased field if non-nil, zero value otherwise.

### GetIsAttrBasedOk

`func (o *NiatelemetryEpgAllOf) GetIsAttrBasedOk() (*string, bool)`

GetIsAttrBasedOk returns a tuple with the IsAttrBased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAttrBased

`func (o *NiatelemetryEpgAllOf) SetIsAttrBased(v string)`

SetIsAttrBased sets IsAttrBased field to given value.

### HasIsAttrBased

`func (o *NiatelemetryEpgAllOf) HasIsAttrBased() bool`

HasIsAttrBased returns a boolean if a field has been set.

### GetMicrosegmentation

`func (o *NiatelemetryEpgAllOf) GetMicrosegmentation() string`

GetMicrosegmentation returns the Microsegmentation field if non-nil, zero value otherwise.

### GetMicrosegmentationOk

`func (o *NiatelemetryEpgAllOf) GetMicrosegmentationOk() (*string, bool)`

GetMicrosegmentationOk returns a tuple with the Microsegmentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosegmentation

`func (o *NiatelemetryEpgAllOf) SetMicrosegmentation(v string)`

SetMicrosegmentation sets Microsegmentation field to given value.

### HasMicrosegmentation

`func (o *NiatelemetryEpgAllOf) HasMicrosegmentation() bool`

HasMicrosegmentation returns a boolean if a field has been set.

### GetMicrosoftUsegCount

`func (o *NiatelemetryEpgAllOf) GetMicrosoftUsegCount() int64`

GetMicrosoftUsegCount returns the MicrosoftUsegCount field if non-nil, zero value otherwise.

### GetMicrosoftUsegCountOk

`func (o *NiatelemetryEpgAllOf) GetMicrosoftUsegCountOk() (*int64, bool)`

GetMicrosoftUsegCountOk returns a tuple with the MicrosoftUsegCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosoftUsegCount

`func (o *NiatelemetryEpgAllOf) SetMicrosoftUsegCount(v int64)`

SetMicrosoftUsegCount sets MicrosoftUsegCount field to given value.

### HasMicrosoftUsegCount

`func (o *NiatelemetryEpgAllOf) HasMicrosoftUsegCount() bool`

HasMicrosoftUsegCount returns a boolean if a field has been set.

### GetName

`func (o *NiatelemetryEpgAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NiatelemetryEpgAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NiatelemetryEpgAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NiatelemetryEpgAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrchslDevVipCfgCount

`func (o *NiatelemetryEpgAllOf) GetOrchslDevVipCfgCount() int64`

GetOrchslDevVipCfgCount returns the OrchslDevVipCfgCount field if non-nil, zero value otherwise.

### GetOrchslDevVipCfgCountOk

`func (o *NiatelemetryEpgAllOf) GetOrchslDevVipCfgCountOk() (*int64, bool)`

GetOrchslDevVipCfgCountOk returns a tuple with the OrchslDevVipCfgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrchslDevVipCfgCount

`func (o *NiatelemetryEpgAllOf) SetOrchslDevVipCfgCount(v int64)`

SetOrchslDevVipCfgCount sets OrchslDevVipCfgCount field to given value.

### HasOrchslDevVipCfgCount

`func (o *NiatelemetryEpgAllOf) HasOrchslDevVipCfgCount() bool`

HasOrchslDevVipCfgCount returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryEpgAllOf) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryEpgAllOf) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryEpgAllOf) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryEpgAllOf) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetUsegHypervCount

`func (o *NiatelemetryEpgAllOf) GetUsegHypervCount() int64`

GetUsegHypervCount returns the UsegHypervCount field if non-nil, zero value otherwise.

### GetUsegHypervCountOk

`func (o *NiatelemetryEpgAllOf) GetUsegHypervCountOk() (*int64, bool)`

GetUsegHypervCountOk returns a tuple with the UsegHypervCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsegHypervCount

`func (o *NiatelemetryEpgAllOf) SetUsegHypervCount(v int64)`

SetUsegHypervCount sets UsegHypervCount field to given value.

### HasUsegHypervCount

`func (o *NiatelemetryEpgAllOf) HasUsegHypervCount() bool`

HasUsegHypervCount returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryEpgAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryEpgAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryEpgAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryEpgAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


