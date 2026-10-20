# NiatelemetryLogicalLink

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

### NewNiatelemetryLogicalLink

`func NewNiatelemetryLogicalLink(classId string, objectType string, ) *NiatelemetryLogicalLink`

NewNiatelemetryLogicalLink instantiates a new NiatelemetryLogicalLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryLogicalLinkWithDefaults

`func NewNiatelemetryLogicalLinkWithDefaults() *NiatelemetryLogicalLink`

NewNiatelemetryLogicalLinkWithDefaults instantiates a new NiatelemetryLogicalLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryLogicalLink) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryLogicalLink) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryLogicalLink) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryLogicalLink) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryLogicalLink) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryLogicalLink) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDbId

`func (o *NiatelemetryLogicalLink) GetDbId() int64`

GetDbId returns the DbId field if non-nil, zero value otherwise.

### GetDbIdOk

`func (o *NiatelemetryLogicalLink) GetDbIdOk() (*int64, bool)`

GetDbIdOk returns a tuple with the DbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbId

`func (o *NiatelemetryLogicalLink) SetDbId(v int64)`

SetDbId sets DbId field to given value.

### HasDbId

`func (o *NiatelemetryLogicalLink) HasDbId() bool`

HasDbId returns a boolean if a field has been set.

### GetIsPresent

`func (o *NiatelemetryLogicalLink) GetIsPresent() bool`

GetIsPresent returns the IsPresent field if non-nil, zero value otherwise.

### GetIsPresentOk

`func (o *NiatelemetryLogicalLink) GetIsPresentOk() (*bool, bool)`

GetIsPresentOk returns a tuple with the IsPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPresent

`func (o *NiatelemetryLogicalLink) SetIsPresent(v bool)`

SetIsPresent sets IsPresent field to given value.

### HasIsPresent

`func (o *NiatelemetryLogicalLink) HasIsPresent() bool`

HasIsPresent returns a boolean if a field has been set.

### GetLinkAddr1

`func (o *NiatelemetryLogicalLink) GetLinkAddr1() string`

GetLinkAddr1 returns the LinkAddr1 field if non-nil, zero value otherwise.

### GetLinkAddr1Ok

`func (o *NiatelemetryLogicalLink) GetLinkAddr1Ok() (*string, bool)`

GetLinkAddr1Ok returns a tuple with the LinkAddr1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkAddr1

`func (o *NiatelemetryLogicalLink) SetLinkAddr1(v string)`

SetLinkAddr1 sets LinkAddr1 field to given value.

### HasLinkAddr1

`func (o *NiatelemetryLogicalLink) HasLinkAddr1() bool`

HasLinkAddr1 returns a boolean if a field has been set.

### GetLinkAddr2

`func (o *NiatelemetryLogicalLink) GetLinkAddr2() string`

GetLinkAddr2 returns the LinkAddr2 field if non-nil, zero value otherwise.

### GetLinkAddr2Ok

`func (o *NiatelemetryLogicalLink) GetLinkAddr2Ok() (*string, bool)`

GetLinkAddr2Ok returns a tuple with the LinkAddr2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkAddr2

`func (o *NiatelemetryLogicalLink) SetLinkAddr2(v string)`

SetLinkAddr2 sets LinkAddr2 field to given value.

### HasLinkAddr2

`func (o *NiatelemetryLogicalLink) HasLinkAddr2() bool`

HasLinkAddr2 returns a boolean if a field has been set.

### GetLinkState

`func (o *NiatelemetryLogicalLink) GetLinkState() string`

GetLinkState returns the LinkState field if non-nil, zero value otherwise.

### GetLinkStateOk

`func (o *NiatelemetryLogicalLink) GetLinkStateOk() (*string, bool)`

GetLinkStateOk returns a tuple with the LinkState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkState

`func (o *NiatelemetryLogicalLink) SetLinkState(v string)`

SetLinkState sets LinkState field to given value.

### HasLinkState

`func (o *NiatelemetryLogicalLink) HasLinkState() bool`

HasLinkState returns a boolean if a field has been set.

### GetLinkType

`func (o *NiatelemetryLogicalLink) GetLinkType() string`

GetLinkType returns the LinkType field if non-nil, zero value otherwise.

### GetLinkTypeOk

`func (o *NiatelemetryLogicalLink) GetLinkTypeOk() (*string, bool)`

GetLinkTypeOk returns a tuple with the LinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkType

`func (o *NiatelemetryLogicalLink) SetLinkType(v string)`

SetLinkType sets LinkType field to given value.

### HasLinkType

`func (o *NiatelemetryLogicalLink) HasLinkType() bool`

HasLinkType returns a boolean if a field has been set.

### GetUptime

`func (o *NiatelemetryLogicalLink) GetUptime() string`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *NiatelemetryLogicalLink) GetUptimeOk() (*string, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *NiatelemetryLogicalLink) SetUptime(v string)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *NiatelemetryLogicalLink) HasUptime() bool`

HasUptime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


