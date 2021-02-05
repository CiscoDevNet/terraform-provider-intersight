# NiatelemetryNiaInventoryFabric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.NiaInventoryFabric"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.NiaInventoryFabric"]
**AnycastGwMac** | Pointer to **string** | Returns the aycast gateway mac. | [optional] 
**DcnmtrackerEnabled** | Pointer to **bool** | Returns the value of the dcnmtrackerEnabled field. | [optional] 
**FabricId** | Pointer to **string** | Uniquely identifies a fabric. | [optional] 
**FabricName** | Pointer to **string** | Returns the value of the Name of a fabric. | [optional] 
**IsNgoamEnabled** | Pointer to **string** | Returns if ngoam is enabled. | [optional] 
**IsScheduledBackUpEnabled** | Pointer to **string** | Returns if the scheduled backup is enabled. | [optional] 
**LeafCount** | Pointer to **int64** | Returns total number of leafs in the fabric. | [optional] 
**LogicalLinks** | Pointer to [**[]NiatelemetryLogicalLink**](NiatelemetryLogicalLink.md) |  | [optional] 
**NxosVrfCount** | Pointer to **int64** | Returns the value of the nxosVrfCount field. | [optional] 
**SpineCount** | Pointer to **int64** | Returns total number of spines in the fabric. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewNiatelemetryNiaInventoryFabric

`func NewNiatelemetryNiaInventoryFabric(classId string, objectType string, ) *NiatelemetryNiaInventoryFabric`

NewNiatelemetryNiaInventoryFabric instantiates a new NiatelemetryNiaInventoryFabric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryNiaInventoryFabricWithDefaults

`func NewNiatelemetryNiaInventoryFabricWithDefaults() *NiatelemetryNiaInventoryFabric`

NewNiatelemetryNiaInventoryFabricWithDefaults instantiates a new NiatelemetryNiaInventoryFabric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryNiaInventoryFabric) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryNiaInventoryFabric) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryNiaInventoryFabric) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryNiaInventoryFabric) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryNiaInventoryFabric) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryNiaInventoryFabric) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAnycastGwMac

`func (o *NiatelemetryNiaInventoryFabric) GetAnycastGwMac() string`

GetAnycastGwMac returns the AnycastGwMac field if non-nil, zero value otherwise.

### GetAnycastGwMacOk

`func (o *NiatelemetryNiaInventoryFabric) GetAnycastGwMacOk() (*string, bool)`

GetAnycastGwMacOk returns a tuple with the AnycastGwMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnycastGwMac

`func (o *NiatelemetryNiaInventoryFabric) SetAnycastGwMac(v string)`

SetAnycastGwMac sets AnycastGwMac field to given value.

### HasAnycastGwMac

`func (o *NiatelemetryNiaInventoryFabric) HasAnycastGwMac() bool`

HasAnycastGwMac returns a boolean if a field has been set.

### GetDcnmtrackerEnabled

`func (o *NiatelemetryNiaInventoryFabric) GetDcnmtrackerEnabled() bool`

GetDcnmtrackerEnabled returns the DcnmtrackerEnabled field if non-nil, zero value otherwise.

### GetDcnmtrackerEnabledOk

`func (o *NiatelemetryNiaInventoryFabric) GetDcnmtrackerEnabledOk() (*bool, bool)`

GetDcnmtrackerEnabledOk returns a tuple with the DcnmtrackerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcnmtrackerEnabled

`func (o *NiatelemetryNiaInventoryFabric) SetDcnmtrackerEnabled(v bool)`

SetDcnmtrackerEnabled sets DcnmtrackerEnabled field to given value.

### HasDcnmtrackerEnabled

`func (o *NiatelemetryNiaInventoryFabric) HasDcnmtrackerEnabled() bool`

HasDcnmtrackerEnabled returns a boolean if a field has been set.

### GetFabricId

`func (o *NiatelemetryNiaInventoryFabric) GetFabricId() string`

GetFabricId returns the FabricId field if non-nil, zero value otherwise.

### GetFabricIdOk

`func (o *NiatelemetryNiaInventoryFabric) GetFabricIdOk() (*string, bool)`

GetFabricIdOk returns a tuple with the FabricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricId

`func (o *NiatelemetryNiaInventoryFabric) SetFabricId(v string)`

SetFabricId sets FabricId field to given value.

### HasFabricId

`func (o *NiatelemetryNiaInventoryFabric) HasFabricId() bool`

HasFabricId returns a boolean if a field has been set.

### GetFabricName

`func (o *NiatelemetryNiaInventoryFabric) GetFabricName() string`

GetFabricName returns the FabricName field if non-nil, zero value otherwise.

### GetFabricNameOk

`func (o *NiatelemetryNiaInventoryFabric) GetFabricNameOk() (*string, bool)`

GetFabricNameOk returns a tuple with the FabricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricName

`func (o *NiatelemetryNiaInventoryFabric) SetFabricName(v string)`

SetFabricName sets FabricName field to given value.

### HasFabricName

`func (o *NiatelemetryNiaInventoryFabric) HasFabricName() bool`

HasFabricName returns a boolean if a field has been set.

### GetIsNgoamEnabled

`func (o *NiatelemetryNiaInventoryFabric) GetIsNgoamEnabled() string`

GetIsNgoamEnabled returns the IsNgoamEnabled field if non-nil, zero value otherwise.

### GetIsNgoamEnabledOk

`func (o *NiatelemetryNiaInventoryFabric) GetIsNgoamEnabledOk() (*string, bool)`

GetIsNgoamEnabledOk returns a tuple with the IsNgoamEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNgoamEnabled

`func (o *NiatelemetryNiaInventoryFabric) SetIsNgoamEnabled(v string)`

SetIsNgoamEnabled sets IsNgoamEnabled field to given value.

### HasIsNgoamEnabled

`func (o *NiatelemetryNiaInventoryFabric) HasIsNgoamEnabled() bool`

HasIsNgoamEnabled returns a boolean if a field has been set.

### GetIsScheduledBackUpEnabled

`func (o *NiatelemetryNiaInventoryFabric) GetIsScheduledBackUpEnabled() string`

GetIsScheduledBackUpEnabled returns the IsScheduledBackUpEnabled field if non-nil, zero value otherwise.

### GetIsScheduledBackUpEnabledOk

`func (o *NiatelemetryNiaInventoryFabric) GetIsScheduledBackUpEnabledOk() (*string, bool)`

GetIsScheduledBackUpEnabledOk returns a tuple with the IsScheduledBackUpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsScheduledBackUpEnabled

`func (o *NiatelemetryNiaInventoryFabric) SetIsScheduledBackUpEnabled(v string)`

SetIsScheduledBackUpEnabled sets IsScheduledBackUpEnabled field to given value.

### HasIsScheduledBackUpEnabled

`func (o *NiatelemetryNiaInventoryFabric) HasIsScheduledBackUpEnabled() bool`

HasIsScheduledBackUpEnabled returns a boolean if a field has been set.

### GetLeafCount

`func (o *NiatelemetryNiaInventoryFabric) GetLeafCount() int64`

GetLeafCount returns the LeafCount field if non-nil, zero value otherwise.

### GetLeafCountOk

`func (o *NiatelemetryNiaInventoryFabric) GetLeafCountOk() (*int64, bool)`

GetLeafCountOk returns a tuple with the LeafCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeafCount

`func (o *NiatelemetryNiaInventoryFabric) SetLeafCount(v int64)`

SetLeafCount sets LeafCount field to given value.

### HasLeafCount

`func (o *NiatelemetryNiaInventoryFabric) HasLeafCount() bool`

HasLeafCount returns a boolean if a field has been set.

### GetLogicalLinks

`func (o *NiatelemetryNiaInventoryFabric) GetLogicalLinks() []NiatelemetryLogicalLink`

GetLogicalLinks returns the LogicalLinks field if non-nil, zero value otherwise.

### GetLogicalLinksOk

`func (o *NiatelemetryNiaInventoryFabric) GetLogicalLinksOk() (*[]NiatelemetryLogicalLink, bool)`

GetLogicalLinksOk returns a tuple with the LogicalLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalLinks

`func (o *NiatelemetryNiaInventoryFabric) SetLogicalLinks(v []NiatelemetryLogicalLink)`

SetLogicalLinks sets LogicalLinks field to given value.

### HasLogicalLinks

`func (o *NiatelemetryNiaInventoryFabric) HasLogicalLinks() bool`

HasLogicalLinks returns a boolean if a field has been set.

### SetLogicalLinksNil

`func (o *NiatelemetryNiaInventoryFabric) SetLogicalLinksNil(b bool)`

 SetLogicalLinksNil sets the value for LogicalLinks to be an explicit nil

### UnsetLogicalLinks
`func (o *NiatelemetryNiaInventoryFabric) UnsetLogicalLinks()`

UnsetLogicalLinks ensures that no value is present for LogicalLinks, not even an explicit nil
### GetNxosVrfCount

`func (o *NiatelemetryNiaInventoryFabric) GetNxosVrfCount() int64`

GetNxosVrfCount returns the NxosVrfCount field if non-nil, zero value otherwise.

### GetNxosVrfCountOk

`func (o *NiatelemetryNiaInventoryFabric) GetNxosVrfCountOk() (*int64, bool)`

GetNxosVrfCountOk returns a tuple with the NxosVrfCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxosVrfCount

`func (o *NiatelemetryNiaInventoryFabric) SetNxosVrfCount(v int64)`

SetNxosVrfCount sets NxosVrfCount field to given value.

### HasNxosVrfCount

`func (o *NiatelemetryNiaInventoryFabric) HasNxosVrfCount() bool`

HasNxosVrfCount returns a boolean if a field has been set.

### GetSpineCount

`func (o *NiatelemetryNiaInventoryFabric) GetSpineCount() int64`

GetSpineCount returns the SpineCount field if non-nil, zero value otherwise.

### GetSpineCountOk

`func (o *NiatelemetryNiaInventoryFabric) GetSpineCountOk() (*int64, bool)`

GetSpineCountOk returns a tuple with the SpineCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpineCount

`func (o *NiatelemetryNiaInventoryFabric) SetSpineCount(v int64)`

SetSpineCount sets SpineCount field to given value.

### HasSpineCount

`func (o *NiatelemetryNiaInventoryFabric) HasSpineCount() bool`

HasSpineCount returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryNiaInventoryFabric) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryNiaInventoryFabric) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryNiaInventoryFabric) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryNiaInventoryFabric) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


