# DnacVirtualNetworkFabricSite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "dnac.VirtualNetworkFabricSite"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "dnac.VirtualNetworkFabricSite"]
**FabricSiteNameHierarchy** | Pointer to **string** | Fabric site name hierarchy. | [optional] 
**SiteId** | Pointer to **string** | Site id for the virtual network fabric site. | [optional] 
**VirtualNetworkId** | Pointer to **string** | Virtual network id fro the virtual network fabric site. | [optional] 
**VirtualNetworkName** | Pointer to **string** | Virtual network name for the virtual network fabric site. | [optional] 

## Methods

### NewDnacVirtualNetworkFabricSite

`func NewDnacVirtualNetworkFabricSite(classId string, objectType string, ) *DnacVirtualNetworkFabricSite`

NewDnacVirtualNetworkFabricSite instantiates a new DnacVirtualNetworkFabricSite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnacVirtualNetworkFabricSiteWithDefaults

`func NewDnacVirtualNetworkFabricSiteWithDefaults() *DnacVirtualNetworkFabricSite`

NewDnacVirtualNetworkFabricSiteWithDefaults instantiates a new DnacVirtualNetworkFabricSite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *DnacVirtualNetworkFabricSite) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *DnacVirtualNetworkFabricSite) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *DnacVirtualNetworkFabricSite) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *DnacVirtualNetworkFabricSite) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *DnacVirtualNetworkFabricSite) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *DnacVirtualNetworkFabricSite) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFabricSiteNameHierarchy

`func (o *DnacVirtualNetworkFabricSite) GetFabricSiteNameHierarchy() string`

GetFabricSiteNameHierarchy returns the FabricSiteNameHierarchy field if non-nil, zero value otherwise.

### GetFabricSiteNameHierarchyOk

`func (o *DnacVirtualNetworkFabricSite) GetFabricSiteNameHierarchyOk() (*string, bool)`

GetFabricSiteNameHierarchyOk returns a tuple with the FabricSiteNameHierarchy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricSiteNameHierarchy

`func (o *DnacVirtualNetworkFabricSite) SetFabricSiteNameHierarchy(v string)`

SetFabricSiteNameHierarchy sets FabricSiteNameHierarchy field to given value.

### HasFabricSiteNameHierarchy

`func (o *DnacVirtualNetworkFabricSite) HasFabricSiteNameHierarchy() bool`

HasFabricSiteNameHierarchy returns a boolean if a field has been set.

### GetSiteId

`func (o *DnacVirtualNetworkFabricSite) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *DnacVirtualNetworkFabricSite) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *DnacVirtualNetworkFabricSite) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *DnacVirtualNetworkFabricSite) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetVirtualNetworkId

`func (o *DnacVirtualNetworkFabricSite) GetVirtualNetworkId() string`

GetVirtualNetworkId returns the VirtualNetworkId field if non-nil, zero value otherwise.

### GetVirtualNetworkIdOk

`func (o *DnacVirtualNetworkFabricSite) GetVirtualNetworkIdOk() (*string, bool)`

GetVirtualNetworkIdOk returns a tuple with the VirtualNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetworkId

`func (o *DnacVirtualNetworkFabricSite) SetVirtualNetworkId(v string)`

SetVirtualNetworkId sets VirtualNetworkId field to given value.

### HasVirtualNetworkId

`func (o *DnacVirtualNetworkFabricSite) HasVirtualNetworkId() bool`

HasVirtualNetworkId returns a boolean if a field has been set.

### GetVirtualNetworkName

`func (o *DnacVirtualNetworkFabricSite) GetVirtualNetworkName() string`

GetVirtualNetworkName returns the VirtualNetworkName field if non-nil, zero value otherwise.

### GetVirtualNetworkNameOk

`func (o *DnacVirtualNetworkFabricSite) GetVirtualNetworkNameOk() (*string, bool)`

GetVirtualNetworkNameOk returns a tuple with the VirtualNetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetworkName

`func (o *DnacVirtualNetworkFabricSite) SetVirtualNetworkName(v string)`

SetVirtualNetworkName sets VirtualNetworkName field to given value.

### HasVirtualNetworkName

`func (o *DnacVirtualNetworkFabricSite) HasVirtualNetworkName() bool`

HasVirtualNetworkName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


