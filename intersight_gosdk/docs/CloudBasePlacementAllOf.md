# CloudBasePlacementAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "cloud.AwsVpc"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "cloud.AwsVpc"]
**BillingUnit** | Pointer to [**NullableCloudBillingUnit**](cloud.BillingUnit.md) |  | [optional] 
**Identity** | Pointer to **string** | The internally generated identity of this placement. This entity is not manipulated by users. It aids in uniquely identifying the placement object. | [optional] [readonly] 
**RegionInfo** | Pointer to [**NullableCloudCloudRegion**](cloud.CloudRegion.md) |  | [optional] 
**ZoneInfo** | Pointer to [**NullableCloudAvailabilityZone**](cloud.AvailabilityZone.md) |  | [optional] 

## Methods

### NewCloudBasePlacementAllOf

`func NewCloudBasePlacementAllOf(classId string, objectType string, ) *CloudBasePlacementAllOf`

NewCloudBasePlacementAllOf instantiates a new CloudBasePlacementAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBasePlacementAllOfWithDefaults

`func NewCloudBasePlacementAllOfWithDefaults() *CloudBasePlacementAllOf`

NewCloudBasePlacementAllOfWithDefaults instantiates a new CloudBasePlacementAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudBasePlacementAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudBasePlacementAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudBasePlacementAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudBasePlacementAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudBasePlacementAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudBasePlacementAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBillingUnit

`func (o *CloudBasePlacementAllOf) GetBillingUnit() CloudBillingUnit`

GetBillingUnit returns the BillingUnit field if non-nil, zero value otherwise.

### GetBillingUnitOk

`func (o *CloudBasePlacementAllOf) GetBillingUnitOk() (*CloudBillingUnit, bool)`

GetBillingUnitOk returns a tuple with the BillingUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingUnit

`func (o *CloudBasePlacementAllOf) SetBillingUnit(v CloudBillingUnit)`

SetBillingUnit sets BillingUnit field to given value.

### HasBillingUnit

`func (o *CloudBasePlacementAllOf) HasBillingUnit() bool`

HasBillingUnit returns a boolean if a field has been set.

### SetBillingUnitNil

`func (o *CloudBasePlacementAllOf) SetBillingUnitNil(b bool)`

 SetBillingUnitNil sets the value for BillingUnit to be an explicit nil

### UnsetBillingUnit
`func (o *CloudBasePlacementAllOf) UnsetBillingUnit()`

UnsetBillingUnit ensures that no value is present for BillingUnit, not even an explicit nil
### GetIdentity

`func (o *CloudBasePlacementAllOf) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *CloudBasePlacementAllOf) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *CloudBasePlacementAllOf) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *CloudBasePlacementAllOf) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetRegionInfo

`func (o *CloudBasePlacementAllOf) GetRegionInfo() CloudCloudRegion`

GetRegionInfo returns the RegionInfo field if non-nil, zero value otherwise.

### GetRegionInfoOk

`func (o *CloudBasePlacementAllOf) GetRegionInfoOk() (*CloudCloudRegion, bool)`

GetRegionInfoOk returns a tuple with the RegionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionInfo

`func (o *CloudBasePlacementAllOf) SetRegionInfo(v CloudCloudRegion)`

SetRegionInfo sets RegionInfo field to given value.

### HasRegionInfo

`func (o *CloudBasePlacementAllOf) HasRegionInfo() bool`

HasRegionInfo returns a boolean if a field has been set.

### SetRegionInfoNil

`func (o *CloudBasePlacementAllOf) SetRegionInfoNil(b bool)`

 SetRegionInfoNil sets the value for RegionInfo to be an explicit nil

### UnsetRegionInfo
`func (o *CloudBasePlacementAllOf) UnsetRegionInfo()`

UnsetRegionInfo ensures that no value is present for RegionInfo, not even an explicit nil
### GetZoneInfo

`func (o *CloudBasePlacementAllOf) GetZoneInfo() CloudAvailabilityZone`

GetZoneInfo returns the ZoneInfo field if non-nil, zero value otherwise.

### GetZoneInfoOk

`func (o *CloudBasePlacementAllOf) GetZoneInfoOk() (*CloudAvailabilityZone, bool)`

GetZoneInfoOk returns a tuple with the ZoneInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneInfo

`func (o *CloudBasePlacementAllOf) SetZoneInfo(v CloudAvailabilityZone)`

SetZoneInfo sets ZoneInfo field to given value.

### HasZoneInfo

`func (o *CloudBasePlacementAllOf) HasZoneInfo() bool`

HasZoneInfo returns a boolean if a field has been set.

### SetZoneInfoNil

`func (o *CloudBasePlacementAllOf) SetZoneInfoNil(b bool)`

 SetZoneInfoNil sets the value for ZoneInfo to be an explicit nil

### UnsetZoneInfo
`func (o *CloudBasePlacementAllOf) UnsetZoneInfo()`

UnsetZoneInfo ensures that no value is present for ZoneInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


