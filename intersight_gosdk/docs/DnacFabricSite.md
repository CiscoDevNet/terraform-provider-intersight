# DnacFabricSite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "dnac.FabricSite"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "dnac.FabricSite"]
**AuthenticationProfileName** | Pointer to **string** | Authentication profile name. | [optional] 
**FabricSiteId** | Pointer to **string** | UUID for the Fabric Site. | [optional] 
**FabricSiteNameHierarchy** | Pointer to **string** | Fabric site name hierarchy. | [optional] 
**IsPubSubEnabled** | Pointer to **string** | Pub sub for the fabric site. | [optional] 
**SiteId** | Pointer to **string** | Site id for the fabric site. | [optional] 

## Methods

### NewDnacFabricSite

`func NewDnacFabricSite(classId string, objectType string, ) *DnacFabricSite`

NewDnacFabricSite instantiates a new DnacFabricSite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnacFabricSiteWithDefaults

`func NewDnacFabricSiteWithDefaults() *DnacFabricSite`

NewDnacFabricSiteWithDefaults instantiates a new DnacFabricSite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *DnacFabricSite) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *DnacFabricSite) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *DnacFabricSite) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *DnacFabricSite) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *DnacFabricSite) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *DnacFabricSite) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAuthenticationProfileName

`func (o *DnacFabricSite) GetAuthenticationProfileName() string`

GetAuthenticationProfileName returns the AuthenticationProfileName field if non-nil, zero value otherwise.

### GetAuthenticationProfileNameOk

`func (o *DnacFabricSite) GetAuthenticationProfileNameOk() (*string, bool)`

GetAuthenticationProfileNameOk returns a tuple with the AuthenticationProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationProfileName

`func (o *DnacFabricSite) SetAuthenticationProfileName(v string)`

SetAuthenticationProfileName sets AuthenticationProfileName field to given value.

### HasAuthenticationProfileName

`func (o *DnacFabricSite) HasAuthenticationProfileName() bool`

HasAuthenticationProfileName returns a boolean if a field has been set.

### GetFabricSiteId

`func (o *DnacFabricSite) GetFabricSiteId() string`

GetFabricSiteId returns the FabricSiteId field if non-nil, zero value otherwise.

### GetFabricSiteIdOk

`func (o *DnacFabricSite) GetFabricSiteIdOk() (*string, bool)`

GetFabricSiteIdOk returns a tuple with the FabricSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricSiteId

`func (o *DnacFabricSite) SetFabricSiteId(v string)`

SetFabricSiteId sets FabricSiteId field to given value.

### HasFabricSiteId

`func (o *DnacFabricSite) HasFabricSiteId() bool`

HasFabricSiteId returns a boolean if a field has been set.

### GetFabricSiteNameHierarchy

`func (o *DnacFabricSite) GetFabricSiteNameHierarchy() string`

GetFabricSiteNameHierarchy returns the FabricSiteNameHierarchy field if non-nil, zero value otherwise.

### GetFabricSiteNameHierarchyOk

`func (o *DnacFabricSite) GetFabricSiteNameHierarchyOk() (*string, bool)`

GetFabricSiteNameHierarchyOk returns a tuple with the FabricSiteNameHierarchy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricSiteNameHierarchy

`func (o *DnacFabricSite) SetFabricSiteNameHierarchy(v string)`

SetFabricSiteNameHierarchy sets FabricSiteNameHierarchy field to given value.

### HasFabricSiteNameHierarchy

`func (o *DnacFabricSite) HasFabricSiteNameHierarchy() bool`

HasFabricSiteNameHierarchy returns a boolean if a field has been set.

### GetIsPubSubEnabled

`func (o *DnacFabricSite) GetIsPubSubEnabled() string`

GetIsPubSubEnabled returns the IsPubSubEnabled field if non-nil, zero value otherwise.

### GetIsPubSubEnabledOk

`func (o *DnacFabricSite) GetIsPubSubEnabledOk() (*string, bool)`

GetIsPubSubEnabledOk returns a tuple with the IsPubSubEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPubSubEnabled

`func (o *DnacFabricSite) SetIsPubSubEnabled(v string)`

SetIsPubSubEnabled sets IsPubSubEnabled field to given value.

### HasIsPubSubEnabled

`func (o *DnacFabricSite) HasIsPubSubEnabled() bool`

HasIsPubSubEnabled returns a boolean if a field has been set.

### GetSiteId

`func (o *DnacFabricSite) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *DnacFabricSite) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *DnacFabricSite) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *DnacFabricSite) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


