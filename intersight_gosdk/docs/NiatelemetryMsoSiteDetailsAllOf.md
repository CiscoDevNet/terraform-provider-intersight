# NiatelemetryMsoSiteDetailsAllOf

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

### NewNiatelemetryMsoSiteDetailsAllOf

`func NewNiatelemetryMsoSiteDetailsAllOf(classId string, objectType string, ) *NiatelemetryMsoSiteDetailsAllOf`

NewNiatelemetryMsoSiteDetailsAllOf instantiates a new NiatelemetryMsoSiteDetailsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryMsoSiteDetailsAllOfWithDefaults

`func NewNiatelemetryMsoSiteDetailsAllOfWithDefaults() *NiatelemetryMsoSiteDetailsAllOf`

NewNiatelemetryMsoSiteDetailsAllOfWithDefaults instantiates a new NiatelemetryMsoSiteDetailsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryMsoSiteDetailsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryMsoSiteDetailsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryMsoSiteDetailsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryMsoSiteDetailsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryMsoSiteDetailsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryMsoSiteDetailsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIsCloudSecEnabled

`func (o *NiatelemetryMsoSiteDetailsAllOf) GetIsCloudSecEnabled() string`

GetIsCloudSecEnabled returns the IsCloudSecEnabled field if non-nil, zero value otherwise.

### GetIsCloudSecEnabledOk

`func (o *NiatelemetryMsoSiteDetailsAllOf) GetIsCloudSecEnabledOk() (*string, bool)`

GetIsCloudSecEnabledOk returns a tuple with the IsCloudSecEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloudSecEnabled

`func (o *NiatelemetryMsoSiteDetailsAllOf) SetIsCloudSecEnabled(v string)`

SetIsCloudSecEnabled sets IsCloudSecEnabled field to given value.

### HasIsCloudSecEnabled

`func (o *NiatelemetryMsoSiteDetailsAllOf) HasIsCloudSecEnabled() bool`

HasIsCloudSecEnabled returns a boolean if a field has been set.

### GetNumberOfLeafsPerSiteInMso

`func (o *NiatelemetryMsoSiteDetailsAllOf) GetNumberOfLeafsPerSiteInMso() int64`

GetNumberOfLeafsPerSiteInMso returns the NumberOfLeafsPerSiteInMso field if non-nil, zero value otherwise.

### GetNumberOfLeafsPerSiteInMsoOk

`func (o *NiatelemetryMsoSiteDetailsAllOf) GetNumberOfLeafsPerSiteInMsoOk() (*int64, bool)`

GetNumberOfLeafsPerSiteInMsoOk returns a tuple with the NumberOfLeafsPerSiteInMso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfLeafsPerSiteInMso

`func (o *NiatelemetryMsoSiteDetailsAllOf) SetNumberOfLeafsPerSiteInMso(v int64)`

SetNumberOfLeafsPerSiteInMso sets NumberOfLeafsPerSiteInMso field to given value.

### HasNumberOfLeafsPerSiteInMso

`func (o *NiatelemetryMsoSiteDetailsAllOf) HasNumberOfLeafsPerSiteInMso() bool`

HasNumberOfLeafsPerSiteInMso returns a boolean if a field has been set.

### GetNumberOfPodsPerSiteInMso

`func (o *NiatelemetryMsoSiteDetailsAllOf) GetNumberOfPodsPerSiteInMso() int64`

GetNumberOfPodsPerSiteInMso returns the NumberOfPodsPerSiteInMso field if non-nil, zero value otherwise.

### GetNumberOfPodsPerSiteInMsoOk

`func (o *NiatelemetryMsoSiteDetailsAllOf) GetNumberOfPodsPerSiteInMsoOk() (*int64, bool)`

GetNumberOfPodsPerSiteInMsoOk returns a tuple with the NumberOfPodsPerSiteInMso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfPodsPerSiteInMso

`func (o *NiatelemetryMsoSiteDetailsAllOf) SetNumberOfPodsPerSiteInMso(v int64)`

SetNumberOfPodsPerSiteInMso sets NumberOfPodsPerSiteInMso field to given value.

### HasNumberOfPodsPerSiteInMso

`func (o *NiatelemetryMsoSiteDetailsAllOf) HasNumberOfPodsPerSiteInMso() bool`

HasNumberOfPodsPerSiteInMso returns a boolean if a field has been set.

### GetNumberOfSpinesPerSiteInMso

`func (o *NiatelemetryMsoSiteDetailsAllOf) GetNumberOfSpinesPerSiteInMso() int64`

GetNumberOfSpinesPerSiteInMso returns the NumberOfSpinesPerSiteInMso field if non-nil, zero value otherwise.

### GetNumberOfSpinesPerSiteInMsoOk

`func (o *NiatelemetryMsoSiteDetailsAllOf) GetNumberOfSpinesPerSiteInMsoOk() (*int64, bool)`

GetNumberOfSpinesPerSiteInMsoOk returns a tuple with the NumberOfSpinesPerSiteInMso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSpinesPerSiteInMso

`func (o *NiatelemetryMsoSiteDetailsAllOf) SetNumberOfSpinesPerSiteInMso(v int64)`

SetNumberOfSpinesPerSiteInMso sets NumberOfSpinesPerSiteInMso field to given value.

### HasNumberOfSpinesPerSiteInMso

`func (o *NiatelemetryMsoSiteDetailsAllOf) HasNumberOfSpinesPerSiteInMso() bool`

HasNumberOfSpinesPerSiteInMso returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryMsoSiteDetailsAllOf) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryMsoSiteDetailsAllOf) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryMsoSiteDetailsAllOf) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryMsoSiteDetailsAllOf) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetSiteId

`func (o *NiatelemetryMsoSiteDetailsAllOf) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *NiatelemetryMsoSiteDetailsAllOf) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *NiatelemetryMsoSiteDetailsAllOf) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *NiatelemetryMsoSiteDetailsAllOf) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryMsoSiteDetailsAllOf) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryMsoSiteDetailsAllOf) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryMsoSiteDetailsAllOf) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryMsoSiteDetailsAllOf) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSiteVersion

`func (o *NiatelemetryMsoSiteDetailsAllOf) GetSiteVersion() string`

GetSiteVersion returns the SiteVersion field if non-nil, zero value otherwise.

### GetSiteVersionOk

`func (o *NiatelemetryMsoSiteDetailsAllOf) GetSiteVersionOk() (*string, bool)`

GetSiteVersionOk returns a tuple with the SiteVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteVersion

`func (o *NiatelemetryMsoSiteDetailsAllOf) SetSiteVersion(v string)`

SetSiteVersion sets SiteVersion field to given value.

### HasSiteVersion

`func (o *NiatelemetryMsoSiteDetailsAllOf) HasSiteVersion() bool`

HasSiteVersion returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryMsoSiteDetailsAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryMsoSiteDetailsAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryMsoSiteDetailsAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryMsoSiteDetailsAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


