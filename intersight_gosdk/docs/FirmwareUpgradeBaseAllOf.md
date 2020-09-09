# FirmwareUpgradeBaseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectDownload** | Pointer to [**FirmwareDirectDownload**](firmware.DirectDownload.md) |  | [optional] 
**FileServer** | Pointer to [**SoftwarerepositoryFileServer**](softwarerepository.FileServer.md) |  | [optional] 
**NetworkShare** | Pointer to [**FirmwareNetworkShare**](firmware.NetworkShare.md) |  | [optional] 
**SkipEstimateImpact** | Pointer to **bool** | User has the option to skip the estimate impact calculation. | [optional] 
**Status** | Pointer to **string** | Status of the upgrade operation. * &#x60;NONE&#x60; - Upgrade status is not populated. * &#x60;IN_PROGRESS&#x60; - The upgrade is in progress. * &#x60;SUCCESSFUL&#x60; - The upgrade successfully completed. * &#x60;FAILED&#x60; - The upgrade shows failed status. * &#x60;TERMINATED&#x60; - The upgrade has been terminated. | [optional] [default to "NONE"]
**UpgradeType** | Pointer to **string** | Desired upgrade mode to choose either direct download based upgrade or network share upgrade. * &#x60;direct_upgrade&#x60; - Upgrade mode is direct download. * &#x60;network_upgrade&#x60; - Upgrade mode is network upgrade. | [optional] [default to "direct_upgrade"]
**Distributable** | Pointer to [**FirmwareDistributableRelationship**](firmware.Distributable.Relationship.md) |  | [optional] 
**Release** | Pointer to [**SoftwarerepositoryReleaseRelationship**](softwarerepository.Release.Relationship.md) |  | [optional] 
**UpgradeImpact** | Pointer to [**FirmwareUpgradeImpactStatusRelationship**](firmware.UpgradeImpactStatus.Relationship.md) |  | [optional] 
**UpgradeStatus** | Pointer to [**FirmwareUpgradeStatusRelationship**](firmware.UpgradeStatus.Relationship.md) |  | [optional] 

## Methods

### NewFirmwareUpgradeBaseAllOf

`func NewFirmwareUpgradeBaseAllOf() *FirmwareUpgradeBaseAllOf`

NewFirmwareUpgradeBaseAllOf instantiates a new FirmwareUpgradeBaseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareUpgradeBaseAllOfWithDefaults

`func NewFirmwareUpgradeBaseAllOfWithDefaults() *FirmwareUpgradeBaseAllOf`

NewFirmwareUpgradeBaseAllOfWithDefaults instantiates a new FirmwareUpgradeBaseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectDownload

`func (o *FirmwareUpgradeBaseAllOf) GetDirectDownload() FirmwareDirectDownload`

GetDirectDownload returns the DirectDownload field if non-nil, zero value otherwise.

### GetDirectDownloadOk

`func (o *FirmwareUpgradeBaseAllOf) GetDirectDownloadOk() (*FirmwareDirectDownload, bool)`

GetDirectDownloadOk returns a tuple with the DirectDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDownload

`func (o *FirmwareUpgradeBaseAllOf) SetDirectDownload(v FirmwareDirectDownload)`

SetDirectDownload sets DirectDownload field to given value.

### HasDirectDownload

`func (o *FirmwareUpgradeBaseAllOf) HasDirectDownload() bool`

HasDirectDownload returns a boolean if a field has been set.

### GetFileServer

`func (o *FirmwareUpgradeBaseAllOf) GetFileServer() SoftwarerepositoryFileServer`

GetFileServer returns the FileServer field if non-nil, zero value otherwise.

### GetFileServerOk

`func (o *FirmwareUpgradeBaseAllOf) GetFileServerOk() (*SoftwarerepositoryFileServer, bool)`

GetFileServerOk returns a tuple with the FileServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileServer

`func (o *FirmwareUpgradeBaseAllOf) SetFileServer(v SoftwarerepositoryFileServer)`

SetFileServer sets FileServer field to given value.

### HasFileServer

`func (o *FirmwareUpgradeBaseAllOf) HasFileServer() bool`

HasFileServer returns a boolean if a field has been set.

### GetNetworkShare

`func (o *FirmwareUpgradeBaseAllOf) GetNetworkShare() FirmwareNetworkShare`

GetNetworkShare returns the NetworkShare field if non-nil, zero value otherwise.

### GetNetworkShareOk

`func (o *FirmwareUpgradeBaseAllOf) GetNetworkShareOk() (*FirmwareNetworkShare, bool)`

GetNetworkShareOk returns a tuple with the NetworkShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkShare

`func (o *FirmwareUpgradeBaseAllOf) SetNetworkShare(v FirmwareNetworkShare)`

SetNetworkShare sets NetworkShare field to given value.

### HasNetworkShare

`func (o *FirmwareUpgradeBaseAllOf) HasNetworkShare() bool`

HasNetworkShare returns a boolean if a field has been set.

### GetSkipEstimateImpact

`func (o *FirmwareUpgradeBaseAllOf) GetSkipEstimateImpact() bool`

GetSkipEstimateImpact returns the SkipEstimateImpact field if non-nil, zero value otherwise.

### GetSkipEstimateImpactOk

`func (o *FirmwareUpgradeBaseAllOf) GetSkipEstimateImpactOk() (*bool, bool)`

GetSkipEstimateImpactOk returns a tuple with the SkipEstimateImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipEstimateImpact

`func (o *FirmwareUpgradeBaseAllOf) SetSkipEstimateImpact(v bool)`

SetSkipEstimateImpact sets SkipEstimateImpact field to given value.

### HasSkipEstimateImpact

`func (o *FirmwareUpgradeBaseAllOf) HasSkipEstimateImpact() bool`

HasSkipEstimateImpact returns a boolean if a field has been set.

### GetStatus

`func (o *FirmwareUpgradeBaseAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FirmwareUpgradeBaseAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FirmwareUpgradeBaseAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FirmwareUpgradeBaseAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpgradeType

`func (o *FirmwareUpgradeBaseAllOf) GetUpgradeType() string`

GetUpgradeType returns the UpgradeType field if non-nil, zero value otherwise.

### GetUpgradeTypeOk

`func (o *FirmwareUpgradeBaseAllOf) GetUpgradeTypeOk() (*string, bool)`

GetUpgradeTypeOk returns a tuple with the UpgradeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeType

`func (o *FirmwareUpgradeBaseAllOf) SetUpgradeType(v string)`

SetUpgradeType sets UpgradeType field to given value.

### HasUpgradeType

`func (o *FirmwareUpgradeBaseAllOf) HasUpgradeType() bool`

HasUpgradeType returns a boolean if a field has been set.

### GetDistributable

`func (o *FirmwareUpgradeBaseAllOf) GetDistributable() FirmwareDistributableRelationship`

GetDistributable returns the Distributable field if non-nil, zero value otherwise.

### GetDistributableOk

`func (o *FirmwareUpgradeBaseAllOf) GetDistributableOk() (*FirmwareDistributableRelationship, bool)`

GetDistributableOk returns a tuple with the Distributable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributable

`func (o *FirmwareUpgradeBaseAllOf) SetDistributable(v FirmwareDistributableRelationship)`

SetDistributable sets Distributable field to given value.

### HasDistributable

`func (o *FirmwareUpgradeBaseAllOf) HasDistributable() bool`

HasDistributable returns a boolean if a field has been set.

### GetRelease

`func (o *FirmwareUpgradeBaseAllOf) GetRelease() SoftwarerepositoryReleaseRelationship`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *FirmwareUpgradeBaseAllOf) GetReleaseOk() (*SoftwarerepositoryReleaseRelationship, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *FirmwareUpgradeBaseAllOf) SetRelease(v SoftwarerepositoryReleaseRelationship)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *FirmwareUpgradeBaseAllOf) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetUpgradeImpact

`func (o *FirmwareUpgradeBaseAllOf) GetUpgradeImpact() FirmwareUpgradeImpactStatusRelationship`

GetUpgradeImpact returns the UpgradeImpact field if non-nil, zero value otherwise.

### GetUpgradeImpactOk

`func (o *FirmwareUpgradeBaseAllOf) GetUpgradeImpactOk() (*FirmwareUpgradeImpactStatusRelationship, bool)`

GetUpgradeImpactOk returns a tuple with the UpgradeImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeImpact

`func (o *FirmwareUpgradeBaseAllOf) SetUpgradeImpact(v FirmwareUpgradeImpactStatusRelationship)`

SetUpgradeImpact sets UpgradeImpact field to given value.

### HasUpgradeImpact

`func (o *FirmwareUpgradeBaseAllOf) HasUpgradeImpact() bool`

HasUpgradeImpact returns a boolean if a field has been set.

### GetUpgradeStatus

`func (o *FirmwareUpgradeBaseAllOf) GetUpgradeStatus() FirmwareUpgradeStatusRelationship`

GetUpgradeStatus returns the UpgradeStatus field if non-nil, zero value otherwise.

### GetUpgradeStatusOk

`func (o *FirmwareUpgradeBaseAllOf) GetUpgradeStatusOk() (*FirmwareUpgradeStatusRelationship, bool)`

GetUpgradeStatusOk returns a tuple with the UpgradeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeStatus

`func (o *FirmwareUpgradeBaseAllOf) SetUpgradeStatus(v FirmwareUpgradeStatusRelationship)`

SetUpgradeStatus sets UpgradeStatus field to given value.

### HasUpgradeStatus

`func (o *FirmwareUpgradeBaseAllOf) HasUpgradeStatus() bool`

HasUpgradeStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


