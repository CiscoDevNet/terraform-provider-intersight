# DnacSiteIpPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "dnac.SiteIpPool"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "dnac.SiteIpPool"]
**IpPoolCidr** | Pointer to **string** | Ip Pool Cidr in format e.g. 10.0.0.0/24. | [optional] 
**IpPoolId** | Pointer to **string** | Site ip pool identification. | [optional] 
**IpPoolName** | Pointer to **string** | Name for the site ip pool. | [optional] 
**SiteId** | Pointer to **string** | Site to which ip pool is associated with. | [optional] 
**SiteNameHierarchy** | Pointer to **string** | Hierarchy of the site names. | [optional] 

## Methods

### NewDnacSiteIpPool

`func NewDnacSiteIpPool(classId string, objectType string, ) *DnacSiteIpPool`

NewDnacSiteIpPool instantiates a new DnacSiteIpPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnacSiteIpPoolWithDefaults

`func NewDnacSiteIpPoolWithDefaults() *DnacSiteIpPool`

NewDnacSiteIpPoolWithDefaults instantiates a new DnacSiteIpPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *DnacSiteIpPool) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *DnacSiteIpPool) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *DnacSiteIpPool) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *DnacSiteIpPool) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *DnacSiteIpPool) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *DnacSiteIpPool) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIpPoolCidr

`func (o *DnacSiteIpPool) GetIpPoolCidr() string`

GetIpPoolCidr returns the IpPoolCidr field if non-nil, zero value otherwise.

### GetIpPoolCidrOk

`func (o *DnacSiteIpPool) GetIpPoolCidrOk() (*string, bool)`

GetIpPoolCidrOk returns a tuple with the IpPoolCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPoolCidr

`func (o *DnacSiteIpPool) SetIpPoolCidr(v string)`

SetIpPoolCidr sets IpPoolCidr field to given value.

### HasIpPoolCidr

`func (o *DnacSiteIpPool) HasIpPoolCidr() bool`

HasIpPoolCidr returns a boolean if a field has been set.

### GetIpPoolId

`func (o *DnacSiteIpPool) GetIpPoolId() string`

GetIpPoolId returns the IpPoolId field if non-nil, zero value otherwise.

### GetIpPoolIdOk

`func (o *DnacSiteIpPool) GetIpPoolIdOk() (*string, bool)`

GetIpPoolIdOk returns a tuple with the IpPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPoolId

`func (o *DnacSiteIpPool) SetIpPoolId(v string)`

SetIpPoolId sets IpPoolId field to given value.

### HasIpPoolId

`func (o *DnacSiteIpPool) HasIpPoolId() bool`

HasIpPoolId returns a boolean if a field has been set.

### GetIpPoolName

`func (o *DnacSiteIpPool) GetIpPoolName() string`

GetIpPoolName returns the IpPoolName field if non-nil, zero value otherwise.

### GetIpPoolNameOk

`func (o *DnacSiteIpPool) GetIpPoolNameOk() (*string, bool)`

GetIpPoolNameOk returns a tuple with the IpPoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPoolName

`func (o *DnacSiteIpPool) SetIpPoolName(v string)`

SetIpPoolName sets IpPoolName field to given value.

### HasIpPoolName

`func (o *DnacSiteIpPool) HasIpPoolName() bool`

HasIpPoolName returns a boolean if a field has been set.

### GetSiteId

`func (o *DnacSiteIpPool) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *DnacSiteIpPool) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *DnacSiteIpPool) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *DnacSiteIpPool) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSiteNameHierarchy

`func (o *DnacSiteIpPool) GetSiteNameHierarchy() string`

GetSiteNameHierarchy returns the SiteNameHierarchy field if non-nil, zero value otherwise.

### GetSiteNameHierarchyOk

`func (o *DnacSiteIpPool) GetSiteNameHierarchyOk() (*string, bool)`

GetSiteNameHierarchyOk returns a tuple with the SiteNameHierarchy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteNameHierarchy

`func (o *DnacSiteIpPool) SetSiteNameHierarchy(v string)`

SetSiteNameHierarchy sets SiteNameHierarchy field to given value.

### HasSiteNameHierarchy

`func (o *DnacSiteIpPool) HasSiteNameHierarchy() bool`

HasSiteNameHierarchy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


