# CloudBaseVirtualMachine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "cloud.AwsVirtualMachine"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. The enum values provides the list of concrete types that can be instantiated from this abstract type. | [default to "cloud.AwsVirtualMachine"]
**BillingUnit** | Pointer to [**NullableCloudBillingUnit**](cloud.BillingUnit.md) |  | [optional] 
**ImageInfo** | Pointer to [**NullableCloudImageReference**](cloud.ImageReference.md) |  | [optional] 
**InstanceType** | Pointer to [**NullableCloudInstanceType**](cloud.InstanceType.md) |  | [optional] 
**NetworkInterfaceAttachments** | Pointer to [**[]CloudNetworkInterfaceAttachment**](CloudNetworkInterfaceAttachment.md) |  | [optional] 
**PrivateDns** | Pointer to **string** | The private DNS hostname name assigned to the instance. | [optional] [readonly] 
**PublicDns** | Pointer to **string** | The public DNS name assigned to the instance. | [optional] [readonly] 
**Region** | Pointer to [**NullableCloudCloudRegion**](cloud.CloudRegion.md) |  | [optional] 
**Tenancy** | Pointer to **string** | How virtual machines are distributed across physical hardware and affects pricing. | [optional] [readonly] 
**TerminationTime** | Pointer to **time.Time** | Time when this virtualmachine is terminated. | [optional] [readonly] 
**VirtualMachineTags** | Pointer to [**[]CloudCloudTag**](CloudCloudTag.md) |  | [optional] 
**VolumeAttachments** | Pointer to [**[]CloudVolumeAttachment**](CloudVolumeAttachment.md) |  | [optional] 
**Zone** | Pointer to [**NullableCloudAvailabilityZone**](cloud.AvailabilityZone.md) |  | [optional] 

## Methods

### NewCloudBaseVirtualMachine

`func NewCloudBaseVirtualMachine(classId string, objectType string, ) *CloudBaseVirtualMachine`

NewCloudBaseVirtualMachine instantiates a new CloudBaseVirtualMachine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBaseVirtualMachineWithDefaults

`func NewCloudBaseVirtualMachineWithDefaults() *CloudBaseVirtualMachine`

NewCloudBaseVirtualMachineWithDefaults instantiates a new CloudBaseVirtualMachine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *CloudBaseVirtualMachine) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *CloudBaseVirtualMachine) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *CloudBaseVirtualMachine) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *CloudBaseVirtualMachine) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudBaseVirtualMachine) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudBaseVirtualMachine) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetBillingUnit

`func (o *CloudBaseVirtualMachine) GetBillingUnit() CloudBillingUnit`

GetBillingUnit returns the BillingUnit field if non-nil, zero value otherwise.

### GetBillingUnitOk

`func (o *CloudBaseVirtualMachine) GetBillingUnitOk() (*CloudBillingUnit, bool)`

GetBillingUnitOk returns a tuple with the BillingUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingUnit

`func (o *CloudBaseVirtualMachine) SetBillingUnit(v CloudBillingUnit)`

SetBillingUnit sets BillingUnit field to given value.

### HasBillingUnit

`func (o *CloudBaseVirtualMachine) HasBillingUnit() bool`

HasBillingUnit returns a boolean if a field has been set.

### SetBillingUnitNil

`func (o *CloudBaseVirtualMachine) SetBillingUnitNil(b bool)`

 SetBillingUnitNil sets the value for BillingUnit to be an explicit nil

### UnsetBillingUnit
`func (o *CloudBaseVirtualMachine) UnsetBillingUnit()`

UnsetBillingUnit ensures that no value is present for BillingUnit, not even an explicit nil
### GetImageInfo

`func (o *CloudBaseVirtualMachine) GetImageInfo() CloudImageReference`

GetImageInfo returns the ImageInfo field if non-nil, zero value otherwise.

### GetImageInfoOk

`func (o *CloudBaseVirtualMachine) GetImageInfoOk() (*CloudImageReference, bool)`

GetImageInfoOk returns a tuple with the ImageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageInfo

`func (o *CloudBaseVirtualMachine) SetImageInfo(v CloudImageReference)`

SetImageInfo sets ImageInfo field to given value.

### HasImageInfo

`func (o *CloudBaseVirtualMachine) HasImageInfo() bool`

HasImageInfo returns a boolean if a field has been set.

### SetImageInfoNil

`func (o *CloudBaseVirtualMachine) SetImageInfoNil(b bool)`

 SetImageInfoNil sets the value for ImageInfo to be an explicit nil

### UnsetImageInfo
`func (o *CloudBaseVirtualMachine) UnsetImageInfo()`

UnsetImageInfo ensures that no value is present for ImageInfo, not even an explicit nil
### GetInstanceType

`func (o *CloudBaseVirtualMachine) GetInstanceType() CloudInstanceType`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *CloudBaseVirtualMachine) GetInstanceTypeOk() (*CloudInstanceType, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *CloudBaseVirtualMachine) SetInstanceType(v CloudInstanceType)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *CloudBaseVirtualMachine) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### SetInstanceTypeNil

`func (o *CloudBaseVirtualMachine) SetInstanceTypeNil(b bool)`

 SetInstanceTypeNil sets the value for InstanceType to be an explicit nil

### UnsetInstanceType
`func (o *CloudBaseVirtualMachine) UnsetInstanceType()`

UnsetInstanceType ensures that no value is present for InstanceType, not even an explicit nil
### GetNetworkInterfaceAttachments

`func (o *CloudBaseVirtualMachine) GetNetworkInterfaceAttachments() []CloudNetworkInterfaceAttachment`

GetNetworkInterfaceAttachments returns the NetworkInterfaceAttachments field if non-nil, zero value otherwise.

### GetNetworkInterfaceAttachmentsOk

`func (o *CloudBaseVirtualMachine) GetNetworkInterfaceAttachmentsOk() (*[]CloudNetworkInterfaceAttachment, bool)`

GetNetworkInterfaceAttachmentsOk returns a tuple with the NetworkInterfaceAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceAttachments

`func (o *CloudBaseVirtualMachine) SetNetworkInterfaceAttachments(v []CloudNetworkInterfaceAttachment)`

SetNetworkInterfaceAttachments sets NetworkInterfaceAttachments field to given value.

### HasNetworkInterfaceAttachments

`func (o *CloudBaseVirtualMachine) HasNetworkInterfaceAttachments() bool`

HasNetworkInterfaceAttachments returns a boolean if a field has been set.

### SetNetworkInterfaceAttachmentsNil

`func (o *CloudBaseVirtualMachine) SetNetworkInterfaceAttachmentsNil(b bool)`

 SetNetworkInterfaceAttachmentsNil sets the value for NetworkInterfaceAttachments to be an explicit nil

### UnsetNetworkInterfaceAttachments
`func (o *CloudBaseVirtualMachine) UnsetNetworkInterfaceAttachments()`

UnsetNetworkInterfaceAttachments ensures that no value is present for NetworkInterfaceAttachments, not even an explicit nil
### GetPrivateDns

`func (o *CloudBaseVirtualMachine) GetPrivateDns() string`

GetPrivateDns returns the PrivateDns field if non-nil, zero value otherwise.

### GetPrivateDnsOk

`func (o *CloudBaseVirtualMachine) GetPrivateDnsOk() (*string, bool)`

GetPrivateDnsOk returns a tuple with the PrivateDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateDns

`func (o *CloudBaseVirtualMachine) SetPrivateDns(v string)`

SetPrivateDns sets PrivateDns field to given value.

### HasPrivateDns

`func (o *CloudBaseVirtualMachine) HasPrivateDns() bool`

HasPrivateDns returns a boolean if a field has been set.

### GetPublicDns

`func (o *CloudBaseVirtualMachine) GetPublicDns() string`

GetPublicDns returns the PublicDns field if non-nil, zero value otherwise.

### GetPublicDnsOk

`func (o *CloudBaseVirtualMachine) GetPublicDnsOk() (*string, bool)`

GetPublicDnsOk returns a tuple with the PublicDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicDns

`func (o *CloudBaseVirtualMachine) SetPublicDns(v string)`

SetPublicDns sets PublicDns field to given value.

### HasPublicDns

`func (o *CloudBaseVirtualMachine) HasPublicDns() bool`

HasPublicDns returns a boolean if a field has been set.

### GetRegion

`func (o *CloudBaseVirtualMachine) GetRegion() CloudCloudRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CloudBaseVirtualMachine) GetRegionOk() (*CloudCloudRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CloudBaseVirtualMachine) SetRegion(v CloudCloudRegion)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CloudBaseVirtualMachine) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *CloudBaseVirtualMachine) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *CloudBaseVirtualMachine) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetTenancy

`func (o *CloudBaseVirtualMachine) GetTenancy() string`

GetTenancy returns the Tenancy field if non-nil, zero value otherwise.

### GetTenancyOk

`func (o *CloudBaseVirtualMachine) GetTenancyOk() (*string, bool)`

GetTenancyOk returns a tuple with the Tenancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenancy

`func (o *CloudBaseVirtualMachine) SetTenancy(v string)`

SetTenancy sets Tenancy field to given value.

### HasTenancy

`func (o *CloudBaseVirtualMachine) HasTenancy() bool`

HasTenancy returns a boolean if a field has been set.

### GetTerminationTime

`func (o *CloudBaseVirtualMachine) GetTerminationTime() time.Time`

GetTerminationTime returns the TerminationTime field if non-nil, zero value otherwise.

### GetTerminationTimeOk

`func (o *CloudBaseVirtualMachine) GetTerminationTimeOk() (*time.Time, bool)`

GetTerminationTimeOk returns a tuple with the TerminationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationTime

`func (o *CloudBaseVirtualMachine) SetTerminationTime(v time.Time)`

SetTerminationTime sets TerminationTime field to given value.

### HasTerminationTime

`func (o *CloudBaseVirtualMachine) HasTerminationTime() bool`

HasTerminationTime returns a boolean if a field has been set.

### GetVirtualMachineTags

`func (o *CloudBaseVirtualMachine) GetVirtualMachineTags() []CloudCloudTag`

GetVirtualMachineTags returns the VirtualMachineTags field if non-nil, zero value otherwise.

### GetVirtualMachineTagsOk

`func (o *CloudBaseVirtualMachine) GetVirtualMachineTagsOk() (*[]CloudCloudTag, bool)`

GetVirtualMachineTagsOk returns a tuple with the VirtualMachineTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachineTags

`func (o *CloudBaseVirtualMachine) SetVirtualMachineTags(v []CloudCloudTag)`

SetVirtualMachineTags sets VirtualMachineTags field to given value.

### HasVirtualMachineTags

`func (o *CloudBaseVirtualMachine) HasVirtualMachineTags() bool`

HasVirtualMachineTags returns a boolean if a field has been set.

### SetVirtualMachineTagsNil

`func (o *CloudBaseVirtualMachine) SetVirtualMachineTagsNil(b bool)`

 SetVirtualMachineTagsNil sets the value for VirtualMachineTags to be an explicit nil

### UnsetVirtualMachineTags
`func (o *CloudBaseVirtualMachine) UnsetVirtualMachineTags()`

UnsetVirtualMachineTags ensures that no value is present for VirtualMachineTags, not even an explicit nil
### GetVolumeAttachments

`func (o *CloudBaseVirtualMachine) GetVolumeAttachments() []CloudVolumeAttachment`

GetVolumeAttachments returns the VolumeAttachments field if non-nil, zero value otherwise.

### GetVolumeAttachmentsOk

`func (o *CloudBaseVirtualMachine) GetVolumeAttachmentsOk() (*[]CloudVolumeAttachment, bool)`

GetVolumeAttachmentsOk returns a tuple with the VolumeAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeAttachments

`func (o *CloudBaseVirtualMachine) SetVolumeAttachments(v []CloudVolumeAttachment)`

SetVolumeAttachments sets VolumeAttachments field to given value.

### HasVolumeAttachments

`func (o *CloudBaseVirtualMachine) HasVolumeAttachments() bool`

HasVolumeAttachments returns a boolean if a field has been set.

### SetVolumeAttachmentsNil

`func (o *CloudBaseVirtualMachine) SetVolumeAttachmentsNil(b bool)`

 SetVolumeAttachmentsNil sets the value for VolumeAttachments to be an explicit nil

### UnsetVolumeAttachments
`func (o *CloudBaseVirtualMachine) UnsetVolumeAttachments()`

UnsetVolumeAttachments ensures that no value is present for VolumeAttachments, not even an explicit nil
### GetZone

`func (o *CloudBaseVirtualMachine) GetZone() CloudAvailabilityZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *CloudBaseVirtualMachine) GetZoneOk() (*CloudAvailabilityZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *CloudBaseVirtualMachine) SetZone(v CloudAvailabilityZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *CloudBaseVirtualMachine) HasZone() bool`

HasZone returns a boolean if a field has been set.

### SetZoneNil

`func (o *CloudBaseVirtualMachine) SetZoneNil(b bool)`

 SetZoneNil sets the value for Zone to be an explicit nil

### UnsetZone
`func (o *CloudBaseVirtualMachine) UnsetZone()`

UnsetZone ensures that no value is present for Zone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


