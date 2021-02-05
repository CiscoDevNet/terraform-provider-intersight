# NiatelemetryLogicalLinkAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.LogicalLink"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.LogicalLink"]
**DbId** | Pointer to **int64** | Return value of dbId attribute. | [optional] 
**IsPresent** | Pointer to **bool** | Return value of isPresent attribute. | [optional] 
**LinkAddr1** | Pointer to **string** | Return value of linkAddr1 attribute. | [optional] 
**LinkAddr2** | Pointer to **string** | Return value of linkAddr2 attribute. | [optional] 
**LinkState** | Pointer to **string** | Return value of linkState attribute. | [optional] 
**LinkType** | Pointer to **string** | Return value of linkType attribute. | [optional] 
**Uptime** | Pointer to **string** | Return value of uptime attribute. | [optional] 

## Methods

### NewNiatelemetryLogicalLinkAllOf

`func NewNiatelemetryLogicalLinkAllOf(classId string, objectType string, ) *NiatelemetryLogicalLinkAllOf`

NewNiatelemetryLogicalLinkAllOf instantiates a new NiatelemetryLogicalLinkAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryLogicalLinkAllOfWithDefaults

`func NewNiatelemetryLogicalLinkAllOfWithDefaults() *NiatelemetryLogicalLinkAllOf`

NewNiatelemetryLogicalLinkAllOfWithDefaults instantiates a new NiatelemetryLogicalLinkAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryLogicalLinkAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryLogicalLinkAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryLogicalLinkAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryLogicalLinkAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryLogicalLinkAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryLogicalLinkAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDbId

`func (o *NiatelemetryLogicalLinkAllOf) GetDbId() int64`

GetDbId returns the DbId field if non-nil, zero value otherwise.

### GetDbIdOk

`func (o *NiatelemetryLogicalLinkAllOf) GetDbIdOk() (*int64, bool)`

GetDbIdOk returns a tuple with the DbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbId

`func (o *NiatelemetryLogicalLinkAllOf) SetDbId(v int64)`

SetDbId sets DbId field to given value.

### HasDbId

`func (o *NiatelemetryLogicalLinkAllOf) HasDbId() bool`

HasDbId returns a boolean if a field has been set.

### GetIsPresent

`func (o *NiatelemetryLogicalLinkAllOf) GetIsPresent() bool`

GetIsPresent returns the IsPresent field if non-nil, zero value otherwise.

### GetIsPresentOk

`func (o *NiatelemetryLogicalLinkAllOf) GetIsPresentOk() (*bool, bool)`

GetIsPresentOk returns a tuple with the IsPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPresent

`func (o *NiatelemetryLogicalLinkAllOf) SetIsPresent(v bool)`

SetIsPresent sets IsPresent field to given value.

### HasIsPresent

`func (o *NiatelemetryLogicalLinkAllOf) HasIsPresent() bool`

HasIsPresent returns a boolean if a field has been set.

### GetLinkAddr1

`func (o *NiatelemetryLogicalLinkAllOf) GetLinkAddr1() string`

GetLinkAddr1 returns the LinkAddr1 field if non-nil, zero value otherwise.

### GetLinkAddr1Ok

`func (o *NiatelemetryLogicalLinkAllOf) GetLinkAddr1Ok() (*string, bool)`

GetLinkAddr1Ok returns a tuple with the LinkAddr1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkAddr1

`func (o *NiatelemetryLogicalLinkAllOf) SetLinkAddr1(v string)`

SetLinkAddr1 sets LinkAddr1 field to given value.

### HasLinkAddr1

`func (o *NiatelemetryLogicalLinkAllOf) HasLinkAddr1() bool`

HasLinkAddr1 returns a boolean if a field has been set.

### GetLinkAddr2

`func (o *NiatelemetryLogicalLinkAllOf) GetLinkAddr2() string`

GetLinkAddr2 returns the LinkAddr2 field if non-nil, zero value otherwise.

### GetLinkAddr2Ok

`func (o *NiatelemetryLogicalLinkAllOf) GetLinkAddr2Ok() (*string, bool)`

GetLinkAddr2Ok returns a tuple with the LinkAddr2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkAddr2

`func (o *NiatelemetryLogicalLinkAllOf) SetLinkAddr2(v string)`

SetLinkAddr2 sets LinkAddr2 field to given value.

### HasLinkAddr2

`func (o *NiatelemetryLogicalLinkAllOf) HasLinkAddr2() bool`

HasLinkAddr2 returns a boolean if a field has been set.

### GetLinkState

`func (o *NiatelemetryLogicalLinkAllOf) GetLinkState() string`

GetLinkState returns the LinkState field if non-nil, zero value otherwise.

### GetLinkStateOk

`func (o *NiatelemetryLogicalLinkAllOf) GetLinkStateOk() (*string, bool)`

GetLinkStateOk returns a tuple with the LinkState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkState

`func (o *NiatelemetryLogicalLinkAllOf) SetLinkState(v string)`

SetLinkState sets LinkState field to given value.

### HasLinkState

`func (o *NiatelemetryLogicalLinkAllOf) HasLinkState() bool`

HasLinkState returns a boolean if a field has been set.

### GetLinkType

`func (o *NiatelemetryLogicalLinkAllOf) GetLinkType() string`

GetLinkType returns the LinkType field if non-nil, zero value otherwise.

### GetLinkTypeOk

`func (o *NiatelemetryLogicalLinkAllOf) GetLinkTypeOk() (*string, bool)`

GetLinkTypeOk returns a tuple with the LinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkType

`func (o *NiatelemetryLogicalLinkAllOf) SetLinkType(v string)`

SetLinkType sets LinkType field to given value.

### HasLinkType

`func (o *NiatelemetryLogicalLinkAllOf) HasLinkType() bool`

HasLinkType returns a boolean if a field has been set.

### GetUptime

`func (o *NiatelemetryLogicalLinkAllOf) GetUptime() string`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *NiatelemetryLogicalLinkAllOf) GetUptimeOk() (*string, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *NiatelemetryLogicalLinkAllOf) SetUptime(v string)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *NiatelemetryLogicalLinkAllOf) HasUptime() bool`

HasUptime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


