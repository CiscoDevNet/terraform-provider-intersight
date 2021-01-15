# HyperflexDatastoreInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.DatastoreInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.DatastoreInfo"]
**DsBackendId** | Pointer to **string** | This datastore&#39;s backend unique id. | [optional] [readonly] 
**DsFrontendId** | Pointer to **string** | This datastore&#39;s frontend id. | [optional] [readonly] 

## Methods

### NewHyperflexDatastoreInfoAllOf

`func NewHyperflexDatastoreInfoAllOf(classId string, objectType string, ) *HyperflexDatastoreInfoAllOf`

NewHyperflexDatastoreInfoAllOf instantiates a new HyperflexDatastoreInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexDatastoreInfoAllOfWithDefaults

`func NewHyperflexDatastoreInfoAllOfWithDefaults() *HyperflexDatastoreInfoAllOf`

NewHyperflexDatastoreInfoAllOfWithDefaults instantiates a new HyperflexDatastoreInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexDatastoreInfoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexDatastoreInfoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexDatastoreInfoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexDatastoreInfoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexDatastoreInfoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexDatastoreInfoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDsBackendId

`func (o *HyperflexDatastoreInfoAllOf) GetDsBackendId() string`

GetDsBackendId returns the DsBackendId field if non-nil, zero value otherwise.

### GetDsBackendIdOk

`func (o *HyperflexDatastoreInfoAllOf) GetDsBackendIdOk() (*string, bool)`

GetDsBackendIdOk returns a tuple with the DsBackendId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsBackendId

`func (o *HyperflexDatastoreInfoAllOf) SetDsBackendId(v string)`

SetDsBackendId sets DsBackendId field to given value.

### HasDsBackendId

`func (o *HyperflexDatastoreInfoAllOf) HasDsBackendId() bool`

HasDsBackendId returns a boolean if a field has been set.

### GetDsFrontendId

`func (o *HyperflexDatastoreInfoAllOf) GetDsFrontendId() string`

GetDsFrontendId returns the DsFrontendId field if non-nil, zero value otherwise.

### GetDsFrontendIdOk

`func (o *HyperflexDatastoreInfoAllOf) GetDsFrontendIdOk() (*string, bool)`

GetDsFrontendIdOk returns a tuple with the DsFrontendId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsFrontendId

`func (o *HyperflexDatastoreInfoAllOf) SetDsFrontendId(v string)`

SetDsFrontendId sets DsFrontendId field to given value.

### HasDsFrontendId

`func (o *HyperflexDatastoreInfoAllOf) HasDsFrontendId() bool`

HasDsFrontendId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


