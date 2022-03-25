# NiatelemetrySpinePolGrpDetailsAllOf

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

### NewNiatelemetrySpinePolGrpDetailsAllOf

`func NewNiatelemetrySpinePolGrpDetailsAllOf(classId string, objectType string, ) *NiatelemetrySpinePolGrpDetailsAllOf`

NewNiatelemetrySpinePolGrpDetailsAllOf instantiates a new NiatelemetrySpinePolGrpDetailsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetrySpinePolGrpDetailsAllOfWithDefaults

`func NewNiatelemetrySpinePolGrpDetailsAllOfWithDefaults() *NiatelemetrySpinePolGrpDetailsAllOf`

NewNiatelemetrySpinePolGrpDetailsAllOfWithDefaults instantiates a new NiatelemetrySpinePolGrpDetailsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDn

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetFabricNodeControlDn

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) GetFabricNodeControlDn() string`

GetFabricNodeControlDn returns the FabricNodeControlDn field if non-nil, zero value otherwise.

### GetFabricNodeControlDnOk

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) GetFabricNodeControlDnOk() (*string, bool)`

GetFabricNodeControlDnOk returns a tuple with the FabricNodeControlDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricNodeControlDn

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) SetFabricNodeControlDn(v string)`

SetFabricNodeControlDn sets FabricNodeControlDn field to given value.

### HasFabricNodeControlDn

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) HasFabricNodeControlDn() bool`

HasFabricNodeControlDn returns a boolean if a field has been set.

### GetFabricNodeControlPolName

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) GetFabricNodeControlPolName() string`

GetFabricNodeControlPolName returns the FabricNodeControlPolName field if non-nil, zero value otherwise.

### GetFabricNodeControlPolNameOk

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) GetFabricNodeControlPolNameOk() (*string, bool)`

GetFabricNodeControlPolNameOk returns a tuple with the FabricNodeControlPolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricNodeControlPolName

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) SetFabricNodeControlPolName(v string)`

SetFabricNodeControlPolName sets FabricNodeControlPolName field to given value.

### HasFabricNodeControlPolName

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) HasFabricNodeControlPolName() bool`

HasFabricNodeControlPolName returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSpinePolGroupName

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) GetSpinePolGroupName() string`

GetSpinePolGroupName returns the SpinePolGroupName field if non-nil, zero value otherwise.

### GetSpinePolGroupNameOk

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) GetSpinePolGroupNameOk() (*string, bool)`

GetSpinePolGroupNameOk returns a tuple with the SpinePolGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpinePolGroupName

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) SetSpinePolGroupName(v string)`

SetSpinePolGroupName sets SpinePolGroupName field to given value.

### HasSpinePolGroupName

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) HasSpinePolGroupName() bool`

HasSpinePolGroupName returns a boolean if a field has been set.

### GetSpineProfileName

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) GetSpineProfileName() string`

GetSpineProfileName returns the SpineProfileName field if non-nil, zero value otherwise.

### GetSpineProfileNameOk

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) GetSpineProfileNameOk() (*string, bool)`

GetSpineProfileNameOk returns a tuple with the SpineProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpineProfileName

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) SetSpineProfileName(v string)`

SetSpineProfileName sets SpineProfileName field to given value.

### HasSpineProfileName

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) HasSpineProfileName() bool`

HasSpineProfileName returns a boolean if a field has been set.

### GetState

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) HasState() bool`

HasState returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetrySpinePolGrpDetailsAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


