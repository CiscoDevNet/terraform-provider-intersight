# NiatelemetrySitesAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "niatelemetry.Sites"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "niatelemetry.Sites"]
**Name** | Pointer to **string** | Returns the name of the site. | [optional] 
**SiteType** | Pointer to **string** | Returns the type of the site. | [optional] 
**Uuid** | Pointer to **string** | Returns the uuid of the site. | [optional] 

## Methods

### NewNiatelemetrySitesAllOf

`func NewNiatelemetrySitesAllOf(classId string, objectType string, ) *NiatelemetrySitesAllOf`

NewNiatelemetrySitesAllOf instantiates a new NiatelemetrySitesAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetrySitesAllOfWithDefaults

`func NewNiatelemetrySitesAllOfWithDefaults() *NiatelemetrySitesAllOf`

NewNiatelemetrySitesAllOfWithDefaults instantiates a new NiatelemetrySitesAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *NiatelemetrySitesAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetrySitesAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetrySitesAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *NiatelemetrySitesAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetrySitesAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetrySitesAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *NiatelemetrySitesAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NiatelemetrySitesAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NiatelemetrySitesAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NiatelemetrySitesAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSiteType

`func (o *NiatelemetrySitesAllOf) GetSiteType() string`

GetSiteType returns the SiteType field if non-nil, zero value otherwise.

### GetSiteTypeOk

`func (o *NiatelemetrySitesAllOf) GetSiteTypeOk() (*string, bool)`

GetSiteTypeOk returns a tuple with the SiteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteType

`func (o *NiatelemetrySitesAllOf) SetSiteType(v string)`

SetSiteType sets SiteType field to given value.

### HasSiteType

`func (o *NiatelemetrySitesAllOf) HasSiteType() bool`

HasSiteType returns a boolean if a field has been set.

### GetUuid

`func (o *NiatelemetrySitesAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *NiatelemetrySitesAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *NiatelemetrySitesAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *NiatelemetrySitesAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


