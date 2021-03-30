# VirtualizationVmwareSharesInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "virtualization.VmwareSharesInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "virtualization.VmwareSharesInfo"]
**Level** | Pointer to **string** | The allocation level. The level is a simplified view of shares. Levels map to a pre-determined set of numeric values for shares. If the shares value does not map to a predefined size, then the level is set as custom. | [optional] 
**Shares** | Pointer to **int64** | The number of shares allocated. It is used to determine resource allocation in case of resource contention. Set if level is custom. If level is not set to custom, this value is ignored. Therefore, only shares with custom values can be compared. | [optional] 

## Methods

### NewVirtualizationVmwareSharesInfoAllOf

`func NewVirtualizationVmwareSharesInfoAllOf(classId string, objectType string, ) *VirtualizationVmwareSharesInfoAllOf`

NewVirtualizationVmwareSharesInfoAllOf instantiates a new VirtualizationVmwareSharesInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVmwareSharesInfoAllOfWithDefaults

`func NewVirtualizationVmwareSharesInfoAllOfWithDefaults() *VirtualizationVmwareSharesInfoAllOf`

NewVirtualizationVmwareSharesInfoAllOfWithDefaults instantiates a new VirtualizationVmwareSharesInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *VirtualizationVmwareSharesInfoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *VirtualizationVmwareSharesInfoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *VirtualizationVmwareSharesInfoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *VirtualizationVmwareSharesInfoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VirtualizationVmwareSharesInfoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VirtualizationVmwareSharesInfoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetLevel

`func (o *VirtualizationVmwareSharesInfoAllOf) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *VirtualizationVmwareSharesInfoAllOf) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *VirtualizationVmwareSharesInfoAllOf) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *VirtualizationVmwareSharesInfoAllOf) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetShares

`func (o *VirtualizationVmwareSharesInfoAllOf) GetShares() int64`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *VirtualizationVmwareSharesInfoAllOf) GetSharesOk() (*int64, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *VirtualizationVmwareSharesInfoAllOf) SetShares(v int64)`

SetShares sets Shares field to given value.

### HasShares

`func (o *VirtualizationVmwareSharesInfoAllOf) HasShares() bool`

HasShares returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


