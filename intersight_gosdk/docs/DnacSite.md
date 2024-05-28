# DnacSite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "dnac.Site"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "dnac.Site"]
**FabricSiteNameHierarchy** | Pointer to **string** | Fabric site name hierarchy. | [optional] 
**SiteId** | Pointer to **string** | Identification for the Site. | [optional] 

## Methods

### NewDnacSite

`func NewDnacSite(classId string, objectType string, ) *DnacSite`

NewDnacSite instantiates a new DnacSite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnacSiteWithDefaults

`func NewDnacSiteWithDefaults() *DnacSite`

NewDnacSiteWithDefaults instantiates a new DnacSite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *DnacSite) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *DnacSite) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *DnacSite) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *DnacSite) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *DnacSite) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *DnacSite) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetFabricSiteNameHierarchy

`func (o *DnacSite) GetFabricSiteNameHierarchy() string`

GetFabricSiteNameHierarchy returns the FabricSiteNameHierarchy field if non-nil, zero value otherwise.

### GetFabricSiteNameHierarchyOk

`func (o *DnacSite) GetFabricSiteNameHierarchyOk() (*string, bool)`

GetFabricSiteNameHierarchyOk returns a tuple with the FabricSiteNameHierarchy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricSiteNameHierarchy

`func (o *DnacSite) SetFabricSiteNameHierarchy(v string)`

SetFabricSiteNameHierarchy sets FabricSiteNameHierarchy field to given value.

### HasFabricSiteNameHierarchy

`func (o *DnacSite) HasFabricSiteNameHierarchy() bool`

HasFabricSiteNameHierarchy returns a boolean if a field has been set.

### GetSiteId

`func (o *DnacSite) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *DnacSite) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *DnacSite) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *DnacSite) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


