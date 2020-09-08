# FirmwareFirmwareSummaryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BundleVersion** | Pointer to **string** | Version details at the bundle level for the each of server. | [optional] 
**ComponentsFwInventory** | Pointer to [**[]FirmwareFirmwareInventory**](firmware.FirmwareInventory.md) |  | [optional] 
**Server** | Pointer to [**ComputePhysicalRelationship**](compute.Physical.Relationship.md) |  | [optional] 

## Methods

### NewFirmwareFirmwareSummaryAllOf

`func NewFirmwareFirmwareSummaryAllOf() *FirmwareFirmwareSummaryAllOf`

NewFirmwareFirmwareSummaryAllOf instantiates a new FirmwareFirmwareSummaryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareFirmwareSummaryAllOfWithDefaults

`func NewFirmwareFirmwareSummaryAllOfWithDefaults() *FirmwareFirmwareSummaryAllOf`

NewFirmwareFirmwareSummaryAllOfWithDefaults instantiates a new FirmwareFirmwareSummaryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundleVersion

`func (o *FirmwareFirmwareSummaryAllOf) GetBundleVersion() string`

GetBundleVersion returns the BundleVersion field if non-nil, zero value otherwise.

### GetBundleVersionOk

`func (o *FirmwareFirmwareSummaryAllOf) GetBundleVersionOk() (*string, bool)`

GetBundleVersionOk returns a tuple with the BundleVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleVersion

`func (o *FirmwareFirmwareSummaryAllOf) SetBundleVersion(v string)`

SetBundleVersion sets BundleVersion field to given value.

### HasBundleVersion

`func (o *FirmwareFirmwareSummaryAllOf) HasBundleVersion() bool`

HasBundleVersion returns a boolean if a field has been set.

### GetComponentsFwInventory

`func (o *FirmwareFirmwareSummaryAllOf) GetComponentsFwInventory() []FirmwareFirmwareInventory`

GetComponentsFwInventory returns the ComponentsFwInventory field if non-nil, zero value otherwise.

### GetComponentsFwInventoryOk

`func (o *FirmwareFirmwareSummaryAllOf) GetComponentsFwInventoryOk() (*[]FirmwareFirmwareInventory, bool)`

GetComponentsFwInventoryOk returns a tuple with the ComponentsFwInventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentsFwInventory

`func (o *FirmwareFirmwareSummaryAllOf) SetComponentsFwInventory(v []FirmwareFirmwareInventory)`

SetComponentsFwInventory sets ComponentsFwInventory field to given value.

### HasComponentsFwInventory

`func (o *FirmwareFirmwareSummaryAllOf) HasComponentsFwInventory() bool`

HasComponentsFwInventory returns a boolean if a field has been set.

### GetServer

`func (o *FirmwareFirmwareSummaryAllOf) GetServer() ComputePhysicalRelationship`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *FirmwareFirmwareSummaryAllOf) GetServerOk() (*ComputePhysicalRelationship, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *FirmwareFirmwareSummaryAllOf) SetServer(v ComputePhysicalRelationship)`

SetServer sets Server field to given value.

### HasServer

`func (o *FirmwareFirmwareSummaryAllOf) HasServer() bool`

HasServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


