# HyperflexSiteDetailsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.SiteDetails"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.SiteDetails"]
**Name** | Pointer to **string** | The name of the site for stretch cluster. | [optional] [readonly] 
**NumNodes** | Pointer to **int64** | Number of nodes in the Zone. | [optional] [readonly] 
**ZoneUuid** | Pointer to **string** | The unique identifier of the zone for stretch cluster. | [optional] [readonly] 

## Methods

### NewHyperflexSiteDetailsAllOf

`func NewHyperflexSiteDetailsAllOf(classId string, objectType string, ) *HyperflexSiteDetailsAllOf`

NewHyperflexSiteDetailsAllOf instantiates a new HyperflexSiteDetailsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexSiteDetailsAllOfWithDefaults

`func NewHyperflexSiteDetailsAllOfWithDefaults() *HyperflexSiteDetailsAllOf`

NewHyperflexSiteDetailsAllOfWithDefaults instantiates a new HyperflexSiteDetailsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexSiteDetailsAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexSiteDetailsAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexSiteDetailsAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexSiteDetailsAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexSiteDetailsAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexSiteDetailsAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetName

`func (o *HyperflexSiteDetailsAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HyperflexSiteDetailsAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HyperflexSiteDetailsAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HyperflexSiteDetailsAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumNodes

`func (o *HyperflexSiteDetailsAllOf) GetNumNodes() int64`

GetNumNodes returns the NumNodes field if non-nil, zero value otherwise.

### GetNumNodesOk

`func (o *HyperflexSiteDetailsAllOf) GetNumNodesOk() (*int64, bool)`

GetNumNodesOk returns a tuple with the NumNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumNodes

`func (o *HyperflexSiteDetailsAllOf) SetNumNodes(v int64)`

SetNumNodes sets NumNodes field to given value.

### HasNumNodes

`func (o *HyperflexSiteDetailsAllOf) HasNumNodes() bool`

HasNumNodes returns a boolean if a field has been set.

### GetZoneUuid

`func (o *HyperflexSiteDetailsAllOf) GetZoneUuid() string`

GetZoneUuid returns the ZoneUuid field if non-nil, zero value otherwise.

### GetZoneUuidOk

`func (o *HyperflexSiteDetailsAllOf) GetZoneUuidOk() (*string, bool)`

GetZoneUuidOk returns a tuple with the ZoneUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneUuid

`func (o *HyperflexSiteDetailsAllOf) SetZoneUuid(v string)`

SetZoneUuid sets ZoneUuid field to given value.

### HasZoneUuid

`func (o *HyperflexSiteDetailsAllOf) HasZoneUuid() bool`

HasZoneUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


