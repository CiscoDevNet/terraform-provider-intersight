# CapabilityServerTopologyMap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "capability.ServerTopologyMap"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "capability.ServerTopologyMap"]
**Handler** | Pointer to **string** | Handler identifier for managing this topology configuration. | [optional] 
**PcieNodeDetail** | Pointer to [**NullableCapabilityDeviceInventory**](CapabilityDeviceInventory.md) |  | [optional] 
**ServerDetail** | Pointer to [**NullableCapabilityDeviceInventory**](CapabilityDeviceInventory.md) |  | [optional] 
**SupportedTopologyName** | Pointer to **string** | Server model information for which this topology configuration is defined. | [optional] 
**XfmDetail** | Pointer to [**NullableCapabilityDeviceInventory**](CapabilityDeviceInventory.md) |  | [optional] 

## Methods

### NewCapabilityServerTopologyMap

`func NewCapabilityServerTopologyMap(classId string, objectType string, ) *CapabilityServerTopologyMap`

NewCapabilityServerTopologyMap instantiates a new CapabilityServerTopologyMap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityServerTopologyMapWithDefaults

`func NewCapabilityServerTopologyMapWithDefaults() *CapabilityServerTopologyMap`

NewCapabilityServerTopologyMapWithDefaults instantiates a new CapabilityServerTopologyMap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CapabilityServerTopologyMap) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CapabilityServerTopologyMap) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CapabilityServerTopologyMap) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CapabilityServerTopologyMap) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CapabilityServerTopologyMap) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CapabilityServerTopologyMap) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetHandler

`func (o *CapabilityServerTopologyMap) GetHandler() string`

GetHandler returns the Handler field if non-nil, zero value otherwise.

### GetHandlerOk

`func (o *CapabilityServerTopologyMap) GetHandlerOk() (*string, bool)`

GetHandlerOk returns a tuple with the Handler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandler

`func (o *CapabilityServerTopologyMap) SetHandler(v string)`

SetHandler sets Handler field to given value.

### HasHandler

`func (o *CapabilityServerTopologyMap) HasHandler() bool`

HasHandler returns a boolean if a field has been set.

### GetPcieNodeDetail

`func (o *CapabilityServerTopologyMap) GetPcieNodeDetail() CapabilityDeviceInventory`

GetPcieNodeDetail returns the PcieNodeDetail field if non-nil, zero value otherwise.

### GetPcieNodeDetailOk

`func (o *CapabilityServerTopologyMap) GetPcieNodeDetailOk() (*CapabilityDeviceInventory, bool)`

GetPcieNodeDetailOk returns a tuple with the PcieNodeDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcieNodeDetail

`func (o *CapabilityServerTopologyMap) SetPcieNodeDetail(v CapabilityDeviceInventory)`

SetPcieNodeDetail sets PcieNodeDetail field to given value.

### HasPcieNodeDetail

`func (o *CapabilityServerTopologyMap) HasPcieNodeDetail() bool`

HasPcieNodeDetail returns a boolean if a field has been set.

### SetPcieNodeDetailNil

`func (o *CapabilityServerTopologyMap) SetPcieNodeDetailNil(b bool)`

 SetPcieNodeDetailNil sets the value for PcieNodeDetail to be an explicit nil

### UnsetPcieNodeDetail
`func (o *CapabilityServerTopologyMap) UnsetPcieNodeDetail()`

UnsetPcieNodeDetail ensures that no value is present for PcieNodeDetail, not even an explicit nil
### GetServerDetail

`func (o *CapabilityServerTopologyMap) GetServerDetail() CapabilityDeviceInventory`

GetServerDetail returns the ServerDetail field if non-nil, zero value otherwise.

### GetServerDetailOk

`func (o *CapabilityServerTopologyMap) GetServerDetailOk() (*CapabilityDeviceInventory, bool)`

GetServerDetailOk returns a tuple with the ServerDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerDetail

`func (o *CapabilityServerTopologyMap) SetServerDetail(v CapabilityDeviceInventory)`

SetServerDetail sets ServerDetail field to given value.

### HasServerDetail

`func (o *CapabilityServerTopologyMap) HasServerDetail() bool`

HasServerDetail returns a boolean if a field has been set.

### SetServerDetailNil

`func (o *CapabilityServerTopologyMap) SetServerDetailNil(b bool)`

 SetServerDetailNil sets the value for ServerDetail to be an explicit nil

### UnsetServerDetail
`func (o *CapabilityServerTopologyMap) UnsetServerDetail()`

UnsetServerDetail ensures that no value is present for ServerDetail, not even an explicit nil
### GetSupportedTopologyName

`func (o *CapabilityServerTopologyMap) GetSupportedTopologyName() string`

GetSupportedTopologyName returns the SupportedTopologyName field if non-nil, zero value otherwise.

### GetSupportedTopologyNameOk

`func (o *CapabilityServerTopologyMap) GetSupportedTopologyNameOk() (*string, bool)`

GetSupportedTopologyNameOk returns a tuple with the SupportedTopologyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedTopologyName

`func (o *CapabilityServerTopologyMap) SetSupportedTopologyName(v string)`

SetSupportedTopologyName sets SupportedTopologyName field to given value.

### HasSupportedTopologyName

`func (o *CapabilityServerTopologyMap) HasSupportedTopologyName() bool`

HasSupportedTopologyName returns a boolean if a field has been set.

### GetXfmDetail

`func (o *CapabilityServerTopologyMap) GetXfmDetail() CapabilityDeviceInventory`

GetXfmDetail returns the XfmDetail field if non-nil, zero value otherwise.

### GetXfmDetailOk

`func (o *CapabilityServerTopologyMap) GetXfmDetailOk() (*CapabilityDeviceInventory, bool)`

GetXfmDetailOk returns a tuple with the XfmDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXfmDetail

`func (o *CapabilityServerTopologyMap) SetXfmDetail(v CapabilityDeviceInventory)`

SetXfmDetail sets XfmDetail field to given value.

### HasXfmDetail

`func (o *CapabilityServerTopologyMap) HasXfmDetail() bool`

HasXfmDetail returns a boolean if a field has been set.

### SetXfmDetailNil

`func (o *CapabilityServerTopologyMap) SetXfmDetailNil(b bool)`

 SetXfmDetailNil sets the value for XfmDetail to be an explicit nil

### UnsetXfmDetail
`func (o *CapabilityServerTopologyMap) UnsetXfmDetail()`

UnsetXfmDetail ensures that no value is present for XfmDetail, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


