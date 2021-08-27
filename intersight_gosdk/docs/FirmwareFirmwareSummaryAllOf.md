# FirmwareFirmwareSummaryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "firmware.FirmwareSummary"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "firmware.FirmwareSummary"]
**BundleVersion** | Pointer to **string** | Version details at the bundle level for the each of server. | [optional] 
**ComponentsFwInventory** | Pointer to [**[]FirmwareFirmwareInventory**](FirmwareFirmwareInventory.md) |  | [optional] 
**Server** | Pointer to [**ComputePhysicalRelationship**](ComputePhysicalRelationship.md) |  | [optional] 

## Methods

### NewFirmwareFirmwareSummaryAllOf

`func NewFirmwareFirmwareSummaryAllOf(classId string, objectType string, ) *FirmwareFirmwareSummaryAllOf`

NewFirmwareFirmwareSummaryAllOf instantiates a new FirmwareFirmwareSummaryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareFirmwareSummaryAllOfWithDefaults

`func NewFirmwareFirmwareSummaryAllOfWithDefaults() *FirmwareFirmwareSummaryAllOf`

NewFirmwareFirmwareSummaryAllOfWithDefaults instantiates a new FirmwareFirmwareSummaryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *FirmwareFirmwareSummaryAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FirmwareFirmwareSummaryAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FirmwareFirmwareSummaryAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *FirmwareFirmwareSummaryAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FirmwareFirmwareSummaryAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FirmwareFirmwareSummaryAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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

### SetComponentsFwInventoryNil

`func (o *FirmwareFirmwareSummaryAllOf) SetComponentsFwInventoryNil(b bool)`

 SetComponentsFwInventoryNil sets the value for ComponentsFwInventory to be an explicit nil

### UnsetComponentsFwInventory
`func (o *FirmwareFirmwareSummaryAllOf) UnsetComponentsFwInventory()`

UnsetComponentsFwInventory ensures that no value is present for ComponentsFwInventory, not even an explicit nil
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


