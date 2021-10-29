# NiatelemetryFabricPodSs

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

### NewNiatelemetryFabricPodSs

`func NewNiatelemetryFabricPodSs(classId string, objectType string, ) *NiatelemetryFabricPodSs`

NewNiatelemetryFabricPodSs instantiates a new NiatelemetryFabricPodSs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryFabricPodSsWithDefaults

`func NewNiatelemetryFabricPodSsWithDefaults() *NiatelemetryFabricPodSs`

NewNiatelemetryFabricPodSsWithDefaults instantiates a new NiatelemetryFabricPodSs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryFabricPodSs) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryFabricPodSs) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryFabricPodSs) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryFabricPodSs) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryFabricPodSs) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryFabricPodSs) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDn

`func (o *NiatelemetryFabricPodSs) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetryFabricPodSs) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetryFabricPodSs) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetryFabricPodSs) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetFabricPodProf

`func (o *NiatelemetryFabricPodSs) GetFabricPodProf() string`

GetFabricPodProf returns the FabricPodProf field if non-nil, zero value otherwise.

### GetFabricPodProfOk

`func (o *NiatelemetryFabricPodSs) GetFabricPodProfOk() (*string, bool)`

GetFabricPodProfOk returns a tuple with the FabricPodProf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricPodProf

`func (o *NiatelemetryFabricPodSs) SetFabricPodProf(v string)`

SetFabricPodProf sets FabricPodProf field to given value.

### HasFabricPodProf

`func (o *NiatelemetryFabricPodSs) HasFabricPodProf() bool`

HasFabricPodProf returns a boolean if a field has been set.

### GetPodBlk

`func (o *NiatelemetryFabricPodSs) GetPodBlk() string`

GetPodBlk returns the PodBlk field if non-nil, zero value otherwise.

### GetPodBlkOk

`func (o *NiatelemetryFabricPodSs) GetPodBlkOk() (*string, bool)`

GetPodBlkOk returns a tuple with the PodBlk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodBlk

`func (o *NiatelemetryFabricPodSs) SetPodBlk(v string)`

SetPodBlk sets PodBlk field to given value.

### HasPodBlk

`func (o *NiatelemetryFabricPodSs) HasPodBlk() bool`

HasPodBlk returns a boolean if a field has been set.

### GetPodPolGrp

`func (o *NiatelemetryFabricPodSs) GetPodPolGrp() string`

GetPodPolGrp returns the PodPolGrp field if non-nil, zero value otherwise.

### GetPodPolGrpOk

`func (o *NiatelemetryFabricPodSs) GetPodPolGrpOk() (*string, bool)`

GetPodPolGrpOk returns a tuple with the PodPolGrp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodPolGrp

`func (o *NiatelemetryFabricPodSs) SetPodPolGrp(v string)`

SetPodPolGrp sets PodPolGrp field to given value.

### HasPodPolGrp

`func (o *NiatelemetryFabricPodSs) HasPodPolGrp() bool`

HasPodPolGrp returns a boolean if a field has been set.

### GetPolList

`func (o *NiatelemetryFabricPodSs) GetPolList() string`

GetPolList returns the PolList field if non-nil, zero value otherwise.

### GetPolListOk

`func (o *NiatelemetryFabricPodSs) GetPolListOk() (*string, bool)`

GetPolListOk returns a tuple with the PolList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolList

`func (o *NiatelemetryFabricPodSs) SetPolList(v string)`

SetPolList sets PolList field to given value.

### HasPolList

`func (o *NiatelemetryFabricPodSs) HasPolList() bool`

HasPolList returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryFabricPodSs) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryFabricPodSs) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryFabricPodSs) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryFabricPodSs) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryFabricPodSs) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryFabricPodSs) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryFabricPodSs) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryFabricPodSs) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryFabricPodSs) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryFabricPodSs) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryFabricPodSs) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryFabricPodSs) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryFabricPodSs) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryFabricPodSs) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryFabricPodSs) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryFabricPodSs) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


