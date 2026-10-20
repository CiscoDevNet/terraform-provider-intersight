# NiatelemetryLeafPolGrpDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.LeafPolGrpDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.LeafPolGrpDetails"]
**Dn** | Pointer to **string** | Dn of the Pol group and leaf profile relational object. | [optional] 
**FabricNodeControlDn** | Pointer to **string** | Fabric node control dn associated with the pol group. | [optional] 
**FabricNodeControlPolName** | Pointer to **string** | Fabric node control policy name associated with the pol group. | [optional] 
**LeafPolGroupName** | Pointer to **string** | Leaf policy group name in APIC. | [optional] 
**LeafProfileName** | Pointer to **string** | Leaf profile group name in APIC. | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**SiteName** | Pointer to **string** | Name of the APIC site from which this data is being collected. | [optional] 
**State** | Pointer to **string** | State of fabric node control pol. | [optional] 
**RegisteredDevice** | Pointer to [**NullableAssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryLeafPolGrpDetails

`func NewNiatelemetryLeafPolGrpDetails(classId string, objectType string, ) *NiatelemetryLeafPolGrpDetails`

NewNiatelemetryLeafPolGrpDetails instantiates a new NiatelemetryLeafPolGrpDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryLeafPolGrpDetailsWithDefaults

`func NewNiatelemetryLeafPolGrpDetailsWithDefaults() *NiatelemetryLeafPolGrpDetails`

NewNiatelemetryLeafPolGrpDetailsWithDefaults instantiates a new NiatelemetryLeafPolGrpDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryLeafPolGrpDetails) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryLeafPolGrpDetails) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryLeafPolGrpDetails) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryLeafPolGrpDetails) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryLeafPolGrpDetails) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryLeafPolGrpDetails) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDn

`func (o *NiatelemetryLeafPolGrpDetails) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetryLeafPolGrpDetails) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetryLeafPolGrpDetails) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetryLeafPolGrpDetails) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetFabricNodeControlDn

`func (o *NiatelemetryLeafPolGrpDetails) GetFabricNodeControlDn() string`

GetFabricNodeControlDn returns the FabricNodeControlDn field if non-nil, zero value otherwise.

### GetFabricNodeControlDnOk

`func (o *NiatelemetryLeafPolGrpDetails) GetFabricNodeControlDnOk() (*string, bool)`

GetFabricNodeControlDnOk returns a tuple with the FabricNodeControlDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricNodeControlDn

`func (o *NiatelemetryLeafPolGrpDetails) SetFabricNodeControlDn(v string)`

SetFabricNodeControlDn sets FabricNodeControlDn field to given value.

### HasFabricNodeControlDn

`func (o *NiatelemetryLeafPolGrpDetails) HasFabricNodeControlDn() bool`

HasFabricNodeControlDn returns a boolean if a field has been set.

### GetFabricNodeControlPolName

`func (o *NiatelemetryLeafPolGrpDetails) GetFabricNodeControlPolName() string`

GetFabricNodeControlPolName returns the FabricNodeControlPolName field if non-nil, zero value otherwise.

### GetFabricNodeControlPolNameOk

`func (o *NiatelemetryLeafPolGrpDetails) GetFabricNodeControlPolNameOk() (*string, bool)`

GetFabricNodeControlPolNameOk returns a tuple with the FabricNodeControlPolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricNodeControlPolName

`func (o *NiatelemetryLeafPolGrpDetails) SetFabricNodeControlPolName(v string)`

SetFabricNodeControlPolName sets FabricNodeControlPolName field to given value.

### HasFabricNodeControlPolName

`func (o *NiatelemetryLeafPolGrpDetails) HasFabricNodeControlPolName() bool`

HasFabricNodeControlPolName returns a boolean if a field has been set.

### GetLeafPolGroupName

`func (o *NiatelemetryLeafPolGrpDetails) GetLeafPolGroupName() string`

GetLeafPolGroupName returns the LeafPolGroupName field if non-nil, zero value otherwise.

### GetLeafPolGroupNameOk

`func (o *NiatelemetryLeafPolGrpDetails) GetLeafPolGroupNameOk() (*string, bool)`

GetLeafPolGroupNameOk returns a tuple with the LeafPolGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeafPolGroupName

`func (o *NiatelemetryLeafPolGrpDetails) SetLeafPolGroupName(v string)`

SetLeafPolGroupName sets LeafPolGroupName field to given value.

### HasLeafPolGroupName

`func (o *NiatelemetryLeafPolGrpDetails) HasLeafPolGroupName() bool`

HasLeafPolGroupName returns a boolean if a field has been set.

### GetLeafProfileName

`func (o *NiatelemetryLeafPolGrpDetails) GetLeafProfileName() string`

GetLeafProfileName returns the LeafProfileName field if non-nil, zero value otherwise.

### GetLeafProfileNameOk

`func (o *NiatelemetryLeafPolGrpDetails) GetLeafProfileNameOk() (*string, bool)`

GetLeafProfileNameOk returns a tuple with the LeafProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeafProfileName

`func (o *NiatelemetryLeafPolGrpDetails) SetLeafProfileName(v string)`

SetLeafProfileName sets LeafProfileName field to given value.

### HasLeafProfileName

`func (o *NiatelemetryLeafPolGrpDetails) HasLeafProfileName() bool`

HasLeafProfileName returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryLeafPolGrpDetails) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryLeafPolGrpDetails) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryLeafPolGrpDetails) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryLeafPolGrpDetails) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryLeafPolGrpDetails) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryLeafPolGrpDetails) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryLeafPolGrpDetails) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryLeafPolGrpDetails) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryLeafPolGrpDetails) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryLeafPolGrpDetails) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryLeafPolGrpDetails) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryLeafPolGrpDetails) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetState

`func (o *NiatelemetryLeafPolGrpDetails) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NiatelemetryLeafPolGrpDetails) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NiatelemetryLeafPolGrpDetails) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *NiatelemetryLeafPolGrpDetails) HasState() bool`

HasState returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryLeafPolGrpDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryLeafPolGrpDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryLeafPolGrpDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryLeafPolGrpDetails) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### SetRegisteredDeviceNil

`func (o *NiatelemetryLeafPolGrpDetails) SetRegisteredDeviceNil(b bool)`

 SetRegisteredDeviceNil sets the value for RegisteredDevice to be an explicit nil

### UnsetRegisteredDevice
`func (o *NiatelemetryLeafPolGrpDetails) UnsetRegisteredDevice()`

UnsetRegisteredDevice ensures that no value is present for RegisteredDevice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


