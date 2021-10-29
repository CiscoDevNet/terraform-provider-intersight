# NiatelemetryFabricPodSsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.FabricPodSs"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.FabricPodSs"]
**Dn** | Pointer to **string** | Dn of the Fabric podS for APIC. | [optional] 
**FabricPodProf** | Pointer to **string** | Parent PodP of the Fabric podS for APIC. | [optional] 
**PodBlk** | Pointer to **string** | Pod Block for the above Fabric PodS. | [optional] 
**PodPolGrp** | Pointer to **string** | Policy Group for the above Fabric PodS. | [optional] 
**PolList** | Pointer to **string** | List of Dn of CommPols, SnmpPols and TimePols. | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**SiteName** | Pointer to **string** | Name of the APIC site from which this data is being collected. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryFabricPodSsAllOf

`func NewNiatelemetryFabricPodSsAllOf(classId string, objectType string, ) *NiatelemetryFabricPodSsAllOf`

NewNiatelemetryFabricPodSsAllOf instantiates a new NiatelemetryFabricPodSsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryFabricPodSsAllOfWithDefaults

`func NewNiatelemetryFabricPodSsAllOfWithDefaults() *NiatelemetryFabricPodSsAllOf`

NewNiatelemetryFabricPodSsAllOfWithDefaults instantiates a new NiatelemetryFabricPodSsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryFabricPodSsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryFabricPodSsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryFabricPodSsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryFabricPodSsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryFabricPodSsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryFabricPodSsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDn

`func (o *NiatelemetryFabricPodSsAllOf) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetryFabricPodSsAllOf) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetryFabricPodSsAllOf) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetryFabricPodSsAllOf) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetFabricPodProf

`func (o *NiatelemetryFabricPodSsAllOf) GetFabricPodProf() string`

GetFabricPodProf returns the FabricPodProf field if non-nil, zero value otherwise.

### GetFabricPodProfOk

`func (o *NiatelemetryFabricPodSsAllOf) GetFabricPodProfOk() (*string, bool)`

GetFabricPodProfOk returns a tuple with the FabricPodProf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricPodProf

`func (o *NiatelemetryFabricPodSsAllOf) SetFabricPodProf(v string)`

SetFabricPodProf sets FabricPodProf field to given value.

### HasFabricPodProf

`func (o *NiatelemetryFabricPodSsAllOf) HasFabricPodProf() bool`

HasFabricPodProf returns a boolean if a field has been set.

### GetPodBlk

`func (o *NiatelemetryFabricPodSsAllOf) GetPodBlk() string`

GetPodBlk returns the PodBlk field if non-nil, zero value otherwise.

### GetPodBlkOk

`func (o *NiatelemetryFabricPodSsAllOf) GetPodBlkOk() (*string, bool)`

GetPodBlkOk returns a tuple with the PodBlk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodBlk

`func (o *NiatelemetryFabricPodSsAllOf) SetPodBlk(v string)`

SetPodBlk sets PodBlk field to given value.

### HasPodBlk

`func (o *NiatelemetryFabricPodSsAllOf) HasPodBlk() bool`

HasPodBlk returns a boolean if a field has been set.

### GetPodPolGrp

`func (o *NiatelemetryFabricPodSsAllOf) GetPodPolGrp() string`

GetPodPolGrp returns the PodPolGrp field if non-nil, zero value otherwise.

### GetPodPolGrpOk

`func (o *NiatelemetryFabricPodSsAllOf) GetPodPolGrpOk() (*string, bool)`

GetPodPolGrpOk returns a tuple with the PodPolGrp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodPolGrp

`func (o *NiatelemetryFabricPodSsAllOf) SetPodPolGrp(v string)`

SetPodPolGrp sets PodPolGrp field to given value.

### HasPodPolGrp

`func (o *NiatelemetryFabricPodSsAllOf) HasPodPolGrp() bool`

HasPodPolGrp returns a boolean if a field has been set.

### GetPolList

`func (o *NiatelemetryFabricPodSsAllOf) GetPolList() string`

GetPolList returns the PolList field if non-nil, zero value otherwise.

### GetPolListOk

`func (o *NiatelemetryFabricPodSsAllOf) GetPolListOk() (*string, bool)`

GetPolListOk returns a tuple with the PolList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolList

`func (o *NiatelemetryFabricPodSsAllOf) SetPolList(v string)`

SetPolList sets PolList field to given value.

### HasPolList

`func (o *NiatelemetryFabricPodSsAllOf) HasPolList() bool`

HasPolList returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryFabricPodSsAllOf) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryFabricPodSsAllOf) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryFabricPodSsAllOf) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryFabricPodSsAllOf) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryFabricPodSsAllOf) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryFabricPodSsAllOf) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryFabricPodSsAllOf) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryFabricPodSsAllOf) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryFabricPodSsAllOf) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryFabricPodSsAllOf) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryFabricPodSsAllOf) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryFabricPodSsAllOf) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryFabricPodSsAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryFabricPodSsAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryFabricPodSsAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryFabricPodSsAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


