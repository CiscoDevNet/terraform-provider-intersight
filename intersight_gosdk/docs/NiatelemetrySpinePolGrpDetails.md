# NiatelemetrySpinePolGrpDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.SpinePolGrpDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.SpinePolGrpDetails"]
**Dn** | Pointer to **string** | Dn of the Pol group and leaf profile relational object. | [optional] 
**FabricNodeControlDn** | Pointer to **string** | Fabric node control dn associated with the pol group. | [optional] 
**FabricNodeControlPolName** | Pointer to **string** | Fabric node control policy name associated with the pol group. | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**SiteName** | Pointer to **string** | Name of the APIC site from which this data is being collected. | [optional] 
**SpinePolGroupName** | Pointer to **string** | Spine policy group name in APIC. | [optional] 
**SpineProfileName** | Pointer to **string** | Spine profile group name in APIC. | [optional] 
**State** | Pointer to **string** | State of fabric node control pol. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetrySpinePolGrpDetails

`func NewNiatelemetrySpinePolGrpDetails(classId string, objectType string, ) *NiatelemetrySpinePolGrpDetails`

NewNiatelemetrySpinePolGrpDetails instantiates a new NiatelemetrySpinePolGrpDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetrySpinePolGrpDetailsWithDefaults

`func NewNiatelemetrySpinePolGrpDetailsWithDefaults() *NiatelemetrySpinePolGrpDetails`

NewNiatelemetrySpinePolGrpDetailsWithDefaults instantiates a new NiatelemetrySpinePolGrpDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetrySpinePolGrpDetails) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetrySpinePolGrpDetails) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetrySpinePolGrpDetails) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetrySpinePolGrpDetails) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetrySpinePolGrpDetails) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetrySpinePolGrpDetails) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDn

`func (o *NiatelemetrySpinePolGrpDetails) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetrySpinePolGrpDetails) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetrySpinePolGrpDetails) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetrySpinePolGrpDetails) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetFabricNodeControlDn

`func (o *NiatelemetrySpinePolGrpDetails) GetFabricNodeControlDn() string`

GetFabricNodeControlDn returns the FabricNodeControlDn field if non-nil, zero value otherwise.

### GetFabricNodeControlDnOk

`func (o *NiatelemetrySpinePolGrpDetails) GetFabricNodeControlDnOk() (*string, bool)`

GetFabricNodeControlDnOk returns a tuple with the FabricNodeControlDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricNodeControlDn

`func (o *NiatelemetrySpinePolGrpDetails) SetFabricNodeControlDn(v string)`

SetFabricNodeControlDn sets FabricNodeControlDn field to given value.

### HasFabricNodeControlDn

`func (o *NiatelemetrySpinePolGrpDetails) HasFabricNodeControlDn() bool`

HasFabricNodeControlDn returns a boolean if a field has been set.

### GetFabricNodeControlPolName

`func (o *NiatelemetrySpinePolGrpDetails) GetFabricNodeControlPolName() string`

GetFabricNodeControlPolName returns the FabricNodeControlPolName field if non-nil, zero value otherwise.

### GetFabricNodeControlPolNameOk

`func (o *NiatelemetrySpinePolGrpDetails) GetFabricNodeControlPolNameOk() (*string, bool)`

GetFabricNodeControlPolNameOk returns a tuple with the FabricNodeControlPolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricNodeControlPolName

`func (o *NiatelemetrySpinePolGrpDetails) SetFabricNodeControlPolName(v string)`

SetFabricNodeControlPolName sets FabricNodeControlPolName field to given value.

### HasFabricNodeControlPolName

`func (o *NiatelemetrySpinePolGrpDetails) HasFabricNodeControlPolName() bool`

HasFabricNodeControlPolName returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetrySpinePolGrpDetails) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetrySpinePolGrpDetails) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetrySpinePolGrpDetails) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetrySpinePolGrpDetails) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetrySpinePolGrpDetails) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetrySpinePolGrpDetails) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetrySpinePolGrpDetails) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetrySpinePolGrpDetails) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetrySpinePolGrpDetails) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetrySpinePolGrpDetails) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetrySpinePolGrpDetails) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetrySpinePolGrpDetails) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSpinePolGroupName

`func (o *NiatelemetrySpinePolGrpDetails) GetSpinePolGroupName() string`

GetSpinePolGroupName returns the SpinePolGroupName field if non-nil, zero value otherwise.

### GetSpinePolGroupNameOk

`func (o *NiatelemetrySpinePolGrpDetails) GetSpinePolGroupNameOk() (*string, bool)`

GetSpinePolGroupNameOk returns a tuple with the SpinePolGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpinePolGroupName

`func (o *NiatelemetrySpinePolGrpDetails) SetSpinePolGroupName(v string)`

SetSpinePolGroupName sets SpinePolGroupName field to given value.

### HasSpinePolGroupName

`func (o *NiatelemetrySpinePolGrpDetails) HasSpinePolGroupName() bool`

HasSpinePolGroupName returns a boolean if a field has been set.

### GetSpineProfileName

`func (o *NiatelemetrySpinePolGrpDetails) GetSpineProfileName() string`

GetSpineProfileName returns the SpineProfileName field if non-nil, zero value otherwise.

### GetSpineProfileNameOk

`func (o *NiatelemetrySpinePolGrpDetails) GetSpineProfileNameOk() (*string, bool)`

GetSpineProfileNameOk returns a tuple with the SpineProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpineProfileName

`func (o *NiatelemetrySpinePolGrpDetails) SetSpineProfileName(v string)`

SetSpineProfileName sets SpineProfileName field to given value.

### HasSpineProfileName

`func (o *NiatelemetrySpinePolGrpDetails) HasSpineProfileName() bool`

HasSpineProfileName returns a boolean if a field has been set.

### GetState

`func (o *NiatelemetrySpinePolGrpDetails) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NiatelemetrySpinePolGrpDetails) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NiatelemetrySpinePolGrpDetails) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *NiatelemetrySpinePolGrpDetails) HasState() bool`

HasState returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetrySpinePolGrpDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetrySpinePolGrpDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetrySpinePolGrpDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetrySpinePolGrpDetails) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


