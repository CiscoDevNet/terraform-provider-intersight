# HyperflexSiteDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.SiteDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.SiteDetails"]
**Name** | Pointer to **string** | The name of the site for stretch cluster. | [optional] [readonly] 
**NumNodes** | Pointer to **int64** | Number of nodes in the Zone. | [optional] [readonly] 
**ZoneUuid** | Pointer to **string** | The unique identifier of the zone for stretch cluster. | [optional] [readonly] 

## Methods

### NewHyperflexSiteDetails

`func NewHyperflexSiteDetails(classId string, objectType string, ) *HyperflexSiteDetails`

NewHyperflexSiteDetails instantiates a new HyperflexSiteDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexSiteDetailsWithDefaults

`func NewHyperflexSiteDetailsWithDefaults() *HyperflexSiteDetails`

NewHyperflexSiteDetailsWithDefaults instantiates a new HyperflexSiteDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexSiteDetails) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexSiteDetails) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexSiteDetails) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexSiteDetails) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexSiteDetails) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexSiteDetails) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *HyperflexSiteDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HyperflexSiteDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HyperflexSiteDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HyperflexSiteDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumNodes

`func (o *HyperflexSiteDetails) GetNumNodes() int64`

GetNumNodes returns the NumNodes field if non-nil, zero value otherwise.

### GetNumNodesOk

`func (o *HyperflexSiteDetails) GetNumNodesOk() (*int64, bool)`

GetNumNodesOk returns a tuple with the NumNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumNodes

`func (o *HyperflexSiteDetails) SetNumNodes(v int64)`

SetNumNodes sets NumNodes field to given value.

### HasNumNodes

`func (o *HyperflexSiteDetails) HasNumNodes() bool`

HasNumNodes returns a boolean if a field has been set.

### GetZoneUuid

`func (o *HyperflexSiteDetails) GetZoneUuid() string`

GetZoneUuid returns the ZoneUuid field if non-nil, zero value otherwise.

### GetZoneUuidOk

`func (o *HyperflexSiteDetails) GetZoneUuidOk() (*string, bool)`

GetZoneUuidOk returns a tuple with the ZoneUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneUuid

`func (o *HyperflexSiteDetails) SetZoneUuid(v string)`

SetZoneUuid sets ZoneUuid field to given value.

### HasZoneUuid

`func (o *HyperflexSiteDetails) HasZoneUuid() bool`

HasZoneUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


