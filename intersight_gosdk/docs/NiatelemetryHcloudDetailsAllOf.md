# NiatelemetryHcloudDetailsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.HcloudDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.HcloudDetails"]
**Dn** | Pointer to **string** | Dn for the inventories present. | [optional] 
**EpgCount** | Pointer to **int64** | Returns the total number of EPGs deployed. | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**Router** | Pointer to [**[]NiatelemetryCloudRoutersElement**](NiatelemetryCloudRoutersElement.md) |  | [optional] 
**RouterCount** | Pointer to **int64** | Returns the total number of Cisco Cloud Routers deployed. | [optional] 
**SiteName** | Pointer to **string** | Name of the APIC site from which this data is being collected. | [optional] 
**SubnetsAddress** | Pointer to **string** | Returns the IP addresses of the subnets. | [optional] 
**SubnetsCount** | Pointer to **int64** | Returns the total number of subnets deployed. | [optional] 
**TransitGatewaysCount** | Pointer to **int64** | Returns the total number of Transit Gateways deployed. | [optional] 
**VpcCount** | Pointer to **int64** | Returns the total number of VPCs deployed in Azure/AWS platforms. | [optional] 
**VpcCountGcp** | Pointer to **int64** | Returns the total number of VPCs deployed in GCP. | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](AssetDeviceRegistrationRelationship.md) |  | [optional] 

## Methods

### NewNiatelemetryHcloudDetailsAllOf

`func NewNiatelemetryHcloudDetailsAllOf(classId string, objectType string, ) *NiatelemetryHcloudDetailsAllOf`

NewNiatelemetryHcloudDetailsAllOf instantiates a new NiatelemetryHcloudDetailsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryHcloudDetailsAllOfWithDefaults

`func NewNiatelemetryHcloudDetailsAllOfWithDefaults() *NiatelemetryHcloudDetailsAllOf`

NewNiatelemetryHcloudDetailsAllOfWithDefaults instantiates a new NiatelemetryHcloudDetailsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetryHcloudDetailsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryHcloudDetailsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryHcloudDetailsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetryHcloudDetailsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryHcloudDetailsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryHcloudDetailsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDn

`func (o *NiatelemetryHcloudDetailsAllOf) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *NiatelemetryHcloudDetailsAllOf) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *NiatelemetryHcloudDetailsAllOf) SetDn(v string)`

SetDn sets Dn field to given value.

### HasDn

`func (o *NiatelemetryHcloudDetailsAllOf) HasDn() bool`

HasDn returns a boolean if a field has been set.

### GetEpgCount

`func (o *NiatelemetryHcloudDetailsAllOf) GetEpgCount() int64`

GetEpgCount returns the EpgCount field if non-nil, zero value otherwise.

### GetEpgCountOk

`func (o *NiatelemetryHcloudDetailsAllOf) GetEpgCountOk() (*int64, bool)`

GetEpgCountOk returns a tuple with the EpgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpgCount

`func (o *NiatelemetryHcloudDetailsAllOf) SetEpgCount(v int64)`

SetEpgCount sets EpgCount field to given value.

### HasEpgCount

`func (o *NiatelemetryHcloudDetailsAllOf) HasEpgCount() bool`

HasEpgCount returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryHcloudDetailsAllOf) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryHcloudDetailsAllOf) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryHcloudDetailsAllOf) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryHcloudDetailsAllOf) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryHcloudDetailsAllOf) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryHcloudDetailsAllOf) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryHcloudDetailsAllOf) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryHcloudDetailsAllOf) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetRouter

`func (o *NiatelemetryHcloudDetailsAllOf) GetRouter() []NiatelemetryCloudRoutersElement`

GetRouter returns the Router field if non-nil, zero value otherwise.

### GetRouterOk

`func (o *NiatelemetryHcloudDetailsAllOf) GetRouterOk() (*[]NiatelemetryCloudRoutersElement, bool)`

GetRouterOk returns a tuple with the Router field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouter

`func (o *NiatelemetryHcloudDetailsAllOf) SetRouter(v []NiatelemetryCloudRoutersElement)`

SetRouter sets Router field to given value.

### HasRouter

`func (o *NiatelemetryHcloudDetailsAllOf) HasRouter() bool`

HasRouter returns a boolean if a field has been set.

### SetRouterNil

`func (o *NiatelemetryHcloudDetailsAllOf) SetRouterNil(b bool)`

 SetRouterNil sets the value for Router to be an explicit nil

### UnsetRouter
`func (o *NiatelemetryHcloudDetailsAllOf) UnsetRouter()`

UnsetRouter ensures that no value is present for Router, not even an explicit nil
### GetRouterCount

`func (o *NiatelemetryHcloudDetailsAllOf) GetRouterCount() int64`

GetRouterCount returns the RouterCount field if non-nil, zero value otherwise.

### GetRouterCountOk

`func (o *NiatelemetryHcloudDetailsAllOf) GetRouterCountOk() (*int64, bool)`

GetRouterCountOk returns a tuple with the RouterCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterCount

`func (o *NiatelemetryHcloudDetailsAllOf) SetRouterCount(v int64)`

SetRouterCount sets RouterCount field to given value.

### HasRouterCount

`func (o *NiatelemetryHcloudDetailsAllOf) HasRouterCount() bool`

HasRouterCount returns a boolean if a field has been set.

### GetSiteName

`func (o *NiatelemetryHcloudDetailsAllOf) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *NiatelemetryHcloudDetailsAllOf) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *NiatelemetryHcloudDetailsAllOf) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *NiatelemetryHcloudDetailsAllOf) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSubnetsAddress

`func (o *NiatelemetryHcloudDetailsAllOf) GetSubnetsAddress() string`

GetSubnetsAddress returns the SubnetsAddress field if non-nil, zero value otherwise.

### GetSubnetsAddressOk

`func (o *NiatelemetryHcloudDetailsAllOf) GetSubnetsAddressOk() (*string, bool)`

GetSubnetsAddressOk returns a tuple with the SubnetsAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetsAddress

`func (o *NiatelemetryHcloudDetailsAllOf) SetSubnetsAddress(v string)`

SetSubnetsAddress sets SubnetsAddress field to given value.

### HasSubnetsAddress

`func (o *NiatelemetryHcloudDetailsAllOf) HasSubnetsAddress() bool`

HasSubnetsAddress returns a boolean if a field has been set.

### GetSubnetsCount

`func (o *NiatelemetryHcloudDetailsAllOf) GetSubnetsCount() int64`

GetSubnetsCount returns the SubnetsCount field if non-nil, zero value otherwise.

### GetSubnetsCountOk

`func (o *NiatelemetryHcloudDetailsAllOf) GetSubnetsCountOk() (*int64, bool)`

GetSubnetsCountOk returns a tuple with the SubnetsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetsCount

`func (o *NiatelemetryHcloudDetailsAllOf) SetSubnetsCount(v int64)`

SetSubnetsCount sets SubnetsCount field to given value.

### HasSubnetsCount

`func (o *NiatelemetryHcloudDetailsAllOf) HasSubnetsCount() bool`

HasSubnetsCount returns a boolean if a field has been set.

### GetTransitGatewaysCount

`func (o *NiatelemetryHcloudDetailsAllOf) GetTransitGatewaysCount() int64`

GetTransitGatewaysCount returns the TransitGatewaysCount field if non-nil, zero value otherwise.

### GetTransitGatewaysCountOk

`func (o *NiatelemetryHcloudDetailsAllOf) GetTransitGatewaysCountOk() (*int64, bool)`

GetTransitGatewaysCountOk returns a tuple with the TransitGatewaysCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitGatewaysCount

`func (o *NiatelemetryHcloudDetailsAllOf) SetTransitGatewaysCount(v int64)`

SetTransitGatewaysCount sets TransitGatewaysCount field to given value.

### HasTransitGatewaysCount

`func (o *NiatelemetryHcloudDetailsAllOf) HasTransitGatewaysCount() bool`

HasTransitGatewaysCount returns a boolean if a field has been set.

### GetVpcCount

`func (o *NiatelemetryHcloudDetailsAllOf) GetVpcCount() int64`

GetVpcCount returns the VpcCount field if non-nil, zero value otherwise.

### GetVpcCountOk

`func (o *NiatelemetryHcloudDetailsAllOf) GetVpcCountOk() (*int64, bool)`

GetVpcCountOk returns a tuple with the VpcCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcCount

`func (o *NiatelemetryHcloudDetailsAllOf) SetVpcCount(v int64)`

SetVpcCount sets VpcCount field to given value.

### HasVpcCount

`func (o *NiatelemetryHcloudDetailsAllOf) HasVpcCount() bool`

HasVpcCount returns a boolean if a field has been set.

### GetVpcCountGcp

`func (o *NiatelemetryHcloudDetailsAllOf) GetVpcCountGcp() int64`

GetVpcCountGcp returns the VpcCountGcp field if non-nil, zero value otherwise.

### GetVpcCountGcpOk

`func (o *NiatelemetryHcloudDetailsAllOf) GetVpcCountGcpOk() (*int64, bool)`

GetVpcCountGcpOk returns a tuple with the VpcCountGcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcCountGcp

`func (o *NiatelemetryHcloudDetailsAllOf) SetVpcCountGcp(v int64)`

SetVpcCountGcp sets VpcCountGcp field to given value.

### HasVpcCountGcp

`func (o *NiatelemetryHcloudDetailsAllOf) HasVpcCountGcp() bool`

HasVpcCountGcp returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryHcloudDetailsAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryHcloudDetailsAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryHcloudDetailsAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryHcloudDetailsAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


