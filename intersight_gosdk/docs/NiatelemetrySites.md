# NiatelemetrySites

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.Sites"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.Sites"]
**Name** | Pointer to **string** | Returns the name of the site. | [optional] 
**SiteType** | Pointer to **string** | Returns the type of the site. | [optional] 
**Uuid** | Pointer to **string** | Returns the uuid of the site. | [optional] 

## Methods

### NewNiatelemetrySites

`func NewNiatelemetrySites(classId string, objectType string, ) *NiatelemetrySites`

NewNiatelemetrySites instantiates a new NiatelemetrySites object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetrySitesWithDefaults

`func NewNiatelemetrySitesWithDefaults() *NiatelemetrySites`

NewNiatelemetrySitesWithDefaults instantiates a new NiatelemetrySites object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetrySites) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetrySites) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetrySites) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetrySites) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetrySites) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetrySites) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *NiatelemetrySites) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NiatelemetrySites) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NiatelemetrySites) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NiatelemetrySites) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSiteType

`func (o *NiatelemetrySites) GetSiteType() string`

GetSiteType returns the SiteType field if non-nil, zero value otherwise.

### GetSiteTypeOk

`func (o *NiatelemetrySites) GetSiteTypeOk() (*string, bool)`

GetSiteTypeOk returns a tuple with the SiteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteType

`func (o *NiatelemetrySites) SetSiteType(v string)`

SetSiteType sets SiteType field to given value.

### HasSiteType

`func (o *NiatelemetrySites) HasSiteType() bool`

HasSiteType returns a boolean if a field has been set.

### GetUuid

`func (o *NiatelemetrySites) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *NiatelemetrySites) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *NiatelemetrySites) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *NiatelemetrySites) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


