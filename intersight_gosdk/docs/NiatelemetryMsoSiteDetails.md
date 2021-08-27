# NiatelemetryMsoSiteDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.MsoSiteDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.MsoSiteDetails"]
**IsCloudSecEnabled** | Pointer to **string** | Status of cloudSec on Multi-Site Orchestrator site. | [optional] 
**NumberOfLeafsPerSiteInMso** | Pointer to **int64** | Number of leafs per site in Multi-Site Orchestrator. | [optional] 
**NumberOfPodsPerSiteInMso** | Pointer to **int64** | Number of pods per site in Multi-Site Orchestrator. | [optional] 
**NumberOfSpinesPerSiteInMso** | Pointer to **int64** | Number of spines per site in Multi-Site Orchestrator. | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**SiteId** | Pointer to **string** | ID of site in Multi-Site Orchestrator. | [optional] 
**SiteName** | Pointer to **string** | Name of the site in Multi-Site Orchestrator. | [optional] 
**SiteVersion** | Pointer to **string** | Version of the controller in the site. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryMsoSiteDetails

`func NewNiatelemetryMsoSiteDetails(classId string, objectType string, ) *NiatelemetryMsoSiteDetails`

NewNiatelemetryMsoSiteDetails instantiates a new NiatelemetryMsoSiteDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryMsoSiteDetailsWithDefaults

`func NewNiatelemetryMsoSiteDetailsWithDefaults() *NiatelemetryMsoSiteDetails`

NewNiatelemetryMsoSiteDetailsWithDefaults instantiates a new NiatelemetryMsoSiteDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryMsoSiteDetails) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryMsoSiteDetails) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryMsoSiteDetails) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryMsoSiteDetails) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryMsoSiteDetails) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryMsoSiteDetails) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsCloudSecEnabled

`func (o *NiatelemetryMsoSiteDetails) GetIsCloudSecEnabled() string`

GetIsCloudSecEnabled returns the IsCloudSecEnabled field if non-nil, zero value otherwise.

### GetIsCloudSecEnabledOk

`func (o *NiatelemetryMsoSiteDetails) GetIsCloudSecEnabledOk() (*string, bool)`

GetIsCloudSecEnabledOk returns a tuple with the IsCloudSecEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloudSecEnabled

`func (o *NiatelemetryMsoSiteDetails) SetIsCloudSecEnabled(v string)`

SetIsCloudSecEnabled sets IsCloudSecEnabled field to given value.

### HasIsCloudSecEnabled

`func (o *NiatelemetryMsoSiteDetails) HasIsCloudSecEnabled() bool`

HasIsCloudSecEnabled returns a boolean if a field has been set.

### GetNumberOfLeafsPerSiteInMso

`func (o *NiatelemetryMsoSiteDetails) GetNumberOfLeafsPerSiteInMso() int64`

GetNumberOfLeafsPerSiteInMso returns the NumberOfLeafsPerSiteInMso field if non-nil, zero value otherwise.

### GetNumberOfLeafsPerSiteInMsoOk

`func (o *NiatelemetryMsoSiteDetails) GetNumberOfLeafsPerSiteInMsoOk() (*int64, bool)`

GetNumberOfLeafsPerSiteInMsoOk returns a tuple with the NumberOfLeafsPerSiteInMso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfLeafsPerSiteInMso

`func (o *NiatelemetryMsoSiteDetails) SetNumberOfLeafsPerSiteInMso(v int64)`

SetNumberOfLeafsPerSiteInMso sets NumberOfLeafsPerSiteInMso field to given value.

### HasNumberOfLeafsPerSiteInMso

`func (o *NiatelemetryMsoSiteDetails) HasNumberOfLeafsPerSiteInMso() bool`

HasNumberOfLeafsPerSiteInMso returns a boolean if a field has been set.

### GetNumberOfPodsPerSiteInMso

`func (o *NiatelemetryMsoSiteDetails) GetNumberOfPodsPerSiteInMso() int64`

GetNumberOfPodsPerSiteInMso returns the NumberOfPodsPerSiteInMso field if non-nil, zero value otherwise.

### GetNumberOfPodsPerSiteInMsoOk

`func (o *NiatelemetryMsoSiteDetails) GetNumberOfPodsPerSiteInMsoOk() (*int64, bool)`

GetNumberOfPodsPerSiteInMsoOk returns a tuple with the NumberOfPodsPerSiteInMso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfPodsPerSiteInMso

`func (o *NiatelemetryMsoSiteDetails) SetNumberOfPodsPerSiteInMso(v int64)`

SetNumberOfPodsPerSiteInMso sets NumberOfPodsPerSiteInMso field to given value.

### HasNumberOfPodsPerSiteInMso

`func (o *NiatelemetryMsoSiteDetails) HasNumberOfPodsPerSiteInMso() bool`

HasNumberOfPodsPerSiteInMso returns a boolean if a field has been set.

### GetNumberOfSpinesPerSiteInMso

`func (o *NiatelemetryMsoSiteDetails) GetNumberOfSpinesPerSiteInMso() int64`

GetNumberOfSpinesPerSiteInMso returns the NumberOfSpinesPerSiteInMso field if non-nil, zero value otherwise.

### GetNumberOfSpinesPerSiteInMsoOk

`func (o *NiatelemetryMsoSiteDetails) GetNumberOfSpinesPerSiteInMsoOk() (*int64, bool)`

GetNumberOfSpinesPerSiteInMsoOk returns a tuple with the NumberOfSpinesPerSiteInMso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSpinesPerSiteInMso

`func (o *NiatelemetryMsoSiteDetails) SetNumberOfSpinesPerSiteInMso(v int64)`

SetNumberOfSpinesPerSiteInMso sets NumberOfSpinesPerSiteInMso field to given value.

### HasNumberOfSpinesPerSiteInMso

`func (o *NiatelemetryMsoSiteDetails) HasNumberOfSpinesPerSiteInMso() bool`

HasNumberOfSpinesPerSiteInMso returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryMsoSiteDetails) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryMsoSiteDetails) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryMsoSiteDetails) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryMsoSiteDetails) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetSiteId

`func (o *NiatelemetryMsoSiteDetails) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *NiatelemetryMsoSiteDetails) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *NiatelemetryMsoSiteDetails) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *NiatelemetryMsoSiteDetails) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryMsoSiteDetails) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryMsoSiteDetails) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryMsoSiteDetails) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryMsoSiteDetails) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSiteVersion

`func (o *NiatelemetryMsoSiteDetails) GetSiteVersion() string`

GetSiteVersion returns the SiteVersion field if non-nil, zero value otherwise.

### GetSiteVersionOk

`func (o *NiatelemetryMsoSiteDetails) GetSiteVersionOk() (*string, bool)`

GetSiteVersionOk returns a tuple with the SiteVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteVersion

`func (o *NiatelemetryMsoSiteDetails) SetSiteVersion(v string)`

SetSiteVersion sets SiteVersion field to given value.

### HasSiteVersion

`func (o *NiatelemetryMsoSiteDetails) HasSiteVersion() bool`

HasSiteVersion returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryMsoSiteDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryMsoSiteDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryMsoSiteDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryMsoSiteDetails) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


