# HyperflexDatastoreInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hyperflex.DatastoreInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hyperflex.DatastoreInfo"]
**DsBackendId** | Pointer to **string** | This datastore&#39;s backend unique id. | [optional] [readonly] 
**DsFrontendId** | Pointer to **string** | This datastore&#39;s frontend id. | [optional] [readonly] 

## Methods

### NewHyperflexDatastoreInfo

`func NewHyperflexDatastoreInfo(classId string, objectType string, ) *HyperflexDatastoreInfo`

NewHyperflexDatastoreInfo instantiates a new HyperflexDatastoreInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexDatastoreInfoWithDefaults

`func NewHyperflexDatastoreInfoWithDefaults() *HyperflexDatastoreInfo`

NewHyperflexDatastoreInfoWithDefaults instantiates a new HyperflexDatastoreInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HyperflexDatastoreInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexDatastoreInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexDatastoreInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HyperflexDatastoreInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexDatastoreInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexDatastoreInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDsBackendId

`func (o *HyperflexDatastoreInfo) GetDsBackendId() string`

GetDsBackendId returns the DsBackendId field if non-nil, zero value otherwise.

### GetDsBackendIdOk

`func (o *HyperflexDatastoreInfo) GetDsBackendIdOk() (*string, bool)`

GetDsBackendIdOk returns a tuple with the DsBackendId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsBackendId

`func (o *HyperflexDatastoreInfo) SetDsBackendId(v string)`

SetDsBackendId sets DsBackendId field to given value.

### HasDsBackendId

`func (o *HyperflexDatastoreInfo) HasDsBackendId() bool`

HasDsBackendId returns a boolean if a field has been set.

### GetDsFrontendId

`func (o *HyperflexDatastoreInfo) GetDsFrontendId() string`

GetDsFrontendId returns the DsFrontendId field if non-nil, zero value otherwise.

### GetDsFrontendIdOk

`func (o *HyperflexDatastoreInfo) GetDsFrontendIdOk() (*string, bool)`

GetDsFrontendIdOk returns a tuple with the DsFrontendId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsFrontendId

`func (o *HyperflexDatastoreInfo) SetDsFrontendId(v string)`

SetDsFrontendId sets DsFrontendId field to given value.

### HasDsFrontendId

`func (o *HyperflexDatastoreInfo) HasDsFrontendId() bool`

HasDsFrontendId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


