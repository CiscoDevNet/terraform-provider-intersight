# NiatelemetryLeafPolGrpDetailsAllOf

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
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryLeafPolGrpDetailsAllOf

`func NewNiatelemetryLeafPolGrpDetailsAllOf(classId string, objectType string, ) *NiatelemetryLeafPolGrpDetailsAllOf`

NewNiatelemetryLeafPolGrpDetailsAllOf instantiates a new NiatelemetryLeafPolGrpDetailsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryLeafPolGrpDetailsAllOfWithDefaults

`func NewNiatelemetryLeafPolGrpDetailsAllOfWithDefaults() *NiatelemetryLeafPolGrpDetailsAllOf`

NewNiatelemetryLeafPolGrpDetailsAllOfWithDefaults instantiates a new NiatelemetryLeafPolGrpDetailsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDn

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetFabricNodeControlDn

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetFabricNodeControlDn() string`

GetFabricNodeControlDn returns the FabricNodeControlDn field if non-nil, zero value otherwise.

### GetFabricNodeControlDnOk

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetFabricNodeControlDnOk() (*string, bool)`

GetFabricNodeControlDnOk returns a tuple with the FabricNodeControlDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricNodeControlDn

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) SetFabricNodeControlDn(v string)`

SetFabricNodeControlDn sets FabricNodeControlDn field to given value.

### HasFabricNodeControlDn

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) HasFabricNodeControlDn() bool`

HasFabricNodeControlDn returns a boolean if a field has been set.

### GetFabricNodeControlPolName

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetFabricNodeControlPolName() string`

GetFabricNodeControlPolName returns the FabricNodeControlPolName field if non-nil, zero value otherwise.

### GetFabricNodeControlPolNameOk

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetFabricNodeControlPolNameOk() (*string, bool)`

GetFabricNodeControlPolNameOk returns a tuple with the FabricNodeControlPolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricNodeControlPolName

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) SetFabricNodeControlPolName(v string)`

SetFabricNodeControlPolName sets FabricNodeControlPolName field to given value.

### HasFabricNodeControlPolName

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) HasFabricNodeControlPolName() bool`

HasFabricNodeControlPolName returns a boolean if a field has been set.

### GetLeafPolGroupName

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetLeafPolGroupName() string`

GetLeafPolGroupName returns the LeafPolGroupName field if non-nil, zero value otherwise.

### GetLeafPolGroupNameOk

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetLeafPolGroupNameOk() (*string, bool)`

GetLeafPolGroupNameOk returns a tuple with the LeafPolGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeafPolGroupName

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) SetLeafPolGroupName(v string)`

SetLeafPolGroupName sets LeafPolGroupName field to given value.

### HasLeafPolGroupName

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) HasLeafPolGroupName() bool`

HasLeafPolGroupName returns a boolean if a field has been set.

### GetLeafProfileName

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetLeafProfileName() string`

GetLeafProfileName returns the LeafProfileName field if non-nil, zero value otherwise.

### GetLeafProfileNameOk

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetLeafProfileNameOk() (*string, bool)`

GetLeafProfileNameOk returns a tuple with the LeafProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeafProfileName

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) SetLeafProfileName(v string)`

SetLeafProfileName sets LeafProfileName field to given value.

### HasLeafProfileName

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) HasLeafProfileName() bool`

HasLeafProfileName returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetState

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) HasState() bool`

HasState returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryLeafPolGrpDetailsAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


