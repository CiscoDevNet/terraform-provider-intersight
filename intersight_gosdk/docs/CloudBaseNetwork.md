# CloudBaseNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "cloud.AwsSubnet"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "cloud.AwsSubnet"]
**BillingUnit** | Pointer to [**NullableCloudBillingUnit**](cloud.BillingUnit.md) |  | [optional] 
**Cidr** | Pointer to **string** | CIDR scheme for defining an IP block. | [optional] [readonly] 
**RegionInfo** | Pointer to [**NullableCloudCloudRegion**](cloud.CloudRegion.md) |  | [optional] 
**ZoneInfo** | Pointer to [**NullableCloudAvailabilityZone**](cloud.AvailabilityZone.md) |  | [optional] 

## Methods

### NewCloudBaseNetwork

`func NewCloudBaseNetwork(classId string, objectType string, ) *CloudBaseNetwork`

NewCloudBaseNetwork instantiates a new CloudBaseNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBaseNetworkWithDefaults

`func NewCloudBaseNetworkWithDefaults() *CloudBaseNetwork`

NewCloudBaseNetworkWithDefaults instantiates a new CloudBaseNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudBaseNetwork) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudBaseNetwork) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudBaseNetwork) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudBaseNetwork) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudBaseNetwork) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudBaseNetwork) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBillingUnit

`func (o *CloudBaseNetwork) GetBillingUnit() CloudBillingUnit`

GetBillingUnit returns the BillingUnit field if non-nil, zero value otherwise.

### GetBillingUnitOk

`func (o *CloudBaseNetwork) GetBillingUnitOk() (*CloudBillingUnit, bool)`

GetBillingUnitOk returns a tuple with the BillingUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingUnit

`func (o *CloudBaseNetwork) SetBillingUnit(v CloudBillingUnit)`

SetBillingUnit sets BillingUnit field to given value.

### HasBillingUnit

`func (o *CloudBaseNetwork) HasBillingUnit() bool`

HasBillingUnit returns a boolean if a field has been set.

### SetBillingUnitNil

`func (o *CloudBaseNetwork) SetBillingUnitNil(b bool)`

 SetBillingUnitNil sets the value for BillingUnit to be an explicit nil

### UnsetBillingUnit
`func (o *CloudBaseNetwork) UnsetBillingUnit()`

UnsetBillingUnit ensures that no value is present for BillingUnit, not even an explicit nil
### GetCidr

`func (o *CloudBaseNetwork) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *CloudBaseNetwork) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *CloudBaseNetwork) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *CloudBaseNetwork) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetRegionInfo

`func (o *CloudBaseNetwork) GetRegionInfo() CloudCloudRegion`

GetRegionInfo returns the RegionInfo field if non-nil, zero value otherwise.

### GetRegionInfoOk

`func (o *CloudBaseNetwork) GetRegionInfoOk() (*CloudCloudRegion, bool)`

GetRegionInfoOk returns a tuple with the RegionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionInfo

`func (o *CloudBaseNetwork) SetRegionInfo(v CloudCloudRegion)`

SetRegionInfo sets RegionInfo field to given value.

### HasRegionInfo

`func (o *CloudBaseNetwork) HasRegionInfo() bool`

HasRegionInfo returns a boolean if a field has been set.

### SetRegionInfoNil

`func (o *CloudBaseNetwork) SetRegionInfoNil(b bool)`

 SetRegionInfoNil sets the value for RegionInfo to be an explicit nil

### UnsetRegionInfo
`func (o *CloudBaseNetwork) UnsetRegionInfo()`

UnsetRegionInfo ensures that no value is present for RegionInfo, not even an explicit nil
### GetZoneInfo

`func (o *CloudBaseNetwork) GetZoneInfo() CloudAvailabilityZone`

GetZoneInfo returns the ZoneInfo field if non-nil, zero value otherwise.

### GetZoneInfoOk

`func (o *CloudBaseNetwork) GetZoneInfoOk() (*CloudAvailabilityZone, bool)`

GetZoneInfoOk returns a tuple with the ZoneInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneInfo

`func (o *CloudBaseNetwork) SetZoneInfo(v CloudAvailabilityZone)`

SetZoneInfo sets ZoneInfo field to given value.

### HasZoneInfo

`func (o *CloudBaseNetwork) HasZoneInfo() bool`

HasZoneInfo returns a boolean if a field has been set.

### SetZoneInfoNil

`func (o *CloudBaseNetwork) SetZoneInfoNil(b bool)`

 SetZoneInfoNil sets the value for ZoneInfo to be an explicit nil

### UnsetZoneInfo
`func (o *CloudBaseNetwork) UnsetZoneInfo()`

UnsetZoneInfo ensures that no value is present for ZoneInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


