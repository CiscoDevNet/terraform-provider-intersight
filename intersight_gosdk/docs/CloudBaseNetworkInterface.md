# CloudBaseNetworkInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "cloud.AwsNetworkInterface"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "cloud.AwsNetworkInterface"]
**BillingUnit** | Pointer to [**NullableCloudBillingUnit**](cloud.BillingUnit.md) |  | [optional] 
**Cidr** | Pointer to **string** | CIDR scheme for defining an IP block. | [optional] [readonly] 
**Description** | Pointer to **string** | User friendly description of network interface. | [optional] [readonly] 
**Identity** | Pointer to **string** | Internally generated identity of this network interface. | [optional] [readonly] 
**RegionInfo** | Pointer to [**NullableCloudCloudRegion**](cloud.CloudRegion.md) |  | [optional] 
**ZoneInfo** | Pointer to [**NullableCloudAvailabilityZone**](cloud.AvailabilityZone.md) |  | [optional] 

## Methods

### NewCloudBaseNetworkInterface

`func NewCloudBaseNetworkInterface(classId string, objectType string, ) *CloudBaseNetworkInterface`

NewCloudBaseNetworkInterface instantiates a new CloudBaseNetworkInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBaseNetworkInterfaceWithDefaults

`func NewCloudBaseNetworkInterfaceWithDefaults() *CloudBaseNetworkInterface`

NewCloudBaseNetworkInterfaceWithDefaults instantiates a new CloudBaseNetworkInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudBaseNetworkInterface) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudBaseNetworkInterface) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudBaseNetworkInterface) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudBaseNetworkInterface) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudBaseNetworkInterface) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudBaseNetworkInterface) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBillingUnit

`func (o *CloudBaseNetworkInterface) GetBillingUnit() CloudBillingUnit`

GetBillingUnit returns the BillingUnit field if non-nil, zero value otherwise.

### GetBillingUnitOk

`func (o *CloudBaseNetworkInterface) GetBillingUnitOk() (*CloudBillingUnit, bool)`

GetBillingUnitOk returns a tuple with the BillingUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingUnit

`func (o *CloudBaseNetworkInterface) SetBillingUnit(v CloudBillingUnit)`

SetBillingUnit sets BillingUnit field to given value.

### HasBillingUnit

`func (o *CloudBaseNetworkInterface) HasBillingUnit() bool`

HasBillingUnit returns a boolean if a field has been set.

### SetBillingUnitNil

`func (o *CloudBaseNetworkInterface) SetBillingUnitNil(b bool)`

 SetBillingUnitNil sets the value for BillingUnit to be an explicit nil

### UnsetBillingUnit
`func (o *CloudBaseNetworkInterface) UnsetBillingUnit()`

UnsetBillingUnit ensures that no value is present for BillingUnit, not even an explicit nil
### GetCidr

`func (o *CloudBaseNetworkInterface) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *CloudBaseNetworkInterface) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *CloudBaseNetworkInterface) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *CloudBaseNetworkInterface) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetDescription

`func (o *CloudBaseNetworkInterface) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudBaseNetworkInterface) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudBaseNetworkInterface) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudBaseNetworkInterface) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIdentity

`func (o *CloudBaseNetworkInterface) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *CloudBaseNetworkInterface) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *CloudBaseNetworkInterface) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *CloudBaseNetworkInterface) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetRegionInfo

`func (o *CloudBaseNetworkInterface) GetRegionInfo() CloudCloudRegion`

GetRegionInfo returns the RegionInfo field if non-nil, zero value otherwise.

### GetRegionInfoOk

`func (o *CloudBaseNetworkInterface) GetRegionInfoOk() (*CloudCloudRegion, bool)`

GetRegionInfoOk returns a tuple with the RegionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionInfo

`func (o *CloudBaseNetworkInterface) SetRegionInfo(v CloudCloudRegion)`

SetRegionInfo sets RegionInfo field to given value.

### HasRegionInfo

`func (o *CloudBaseNetworkInterface) HasRegionInfo() bool`

HasRegionInfo returns a boolean if a field has been set.

### SetRegionInfoNil

`func (o *CloudBaseNetworkInterface) SetRegionInfoNil(b bool)`

 SetRegionInfoNil sets the value for RegionInfo to be an explicit nil

### UnsetRegionInfo
`func (o *CloudBaseNetworkInterface) UnsetRegionInfo()`

UnsetRegionInfo ensures that no value is present for RegionInfo, not even an explicit nil
### GetZoneInfo

`func (o *CloudBaseNetworkInterface) GetZoneInfo() CloudAvailabilityZone`

GetZoneInfo returns the ZoneInfo field if non-nil, zero value otherwise.

### GetZoneInfoOk

`func (o *CloudBaseNetworkInterface) GetZoneInfoOk() (*CloudAvailabilityZone, bool)`

GetZoneInfoOk returns a tuple with the ZoneInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneInfo

`func (o *CloudBaseNetworkInterface) SetZoneInfo(v CloudAvailabilityZone)`

SetZoneInfo sets ZoneInfo field to given value.

### HasZoneInfo

`func (o *CloudBaseNetworkInterface) HasZoneInfo() bool`

HasZoneInfo returns a boolean if a field has been set.

### SetZoneInfoNil

`func (o *CloudBaseNetworkInterface) SetZoneInfoNil(b bool)`

 SetZoneInfoNil sets the value for ZoneInfo to be an explicit nil

### UnsetZoneInfo
`func (o *CloudBaseNetworkInterface) UnsetZoneInfo()`

UnsetZoneInfo ensures that no value is present for ZoneInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


