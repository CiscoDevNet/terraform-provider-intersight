# HclHwCatalogInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "hcl.HwCatalogInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "hcl.HwCatalogInfo"]
**ServerModel** | Pointer to **string** | Server model information for HyperFlex servers. | [optional] 
**ServerType** | Pointer to **string** | Server type of the server hardware. For example, server type AF is for an all-flash server. | [optional] 
**ServerHwCatalogInfo** | Pointer to [**[]HclServerHwCatalogInfoRelationship**](HclServerHwCatalogInfoRelationship.md) | An array of relationships to hclServerHwCatalogInfo resources. | [optional] 

## Methods

### NewHclHwCatalogInfo

`func NewHclHwCatalogInfo(classId string, objectType string, ) *HclHwCatalogInfo`

NewHclHwCatalogInfo instantiates a new HclHwCatalogInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHclHwCatalogInfoWithDefaults

`func NewHclHwCatalogInfoWithDefaults() *HclHwCatalogInfo`

NewHclHwCatalogInfoWithDefaults instantiates a new HclHwCatalogInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *HclHwCatalogInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HclHwCatalogInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HclHwCatalogInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *HclHwCatalogInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HclHwCatalogInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HclHwCatalogInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetServerModel

`func (o *HclHwCatalogInfo) GetServerModel() string`

GetServerModel returns the ServerModel field if non-nil, zero value otherwise.

### GetServerModelOk

`func (o *HclHwCatalogInfo) GetServerModelOk() (*string, bool)`

GetServerModelOk returns a tuple with the ServerModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerModel

`func (o *HclHwCatalogInfo) SetServerModel(v string)`

SetServerModel sets ServerModel field to given value.

### HasServerModel

`func (o *HclHwCatalogInfo) HasServerModel() bool`

HasServerModel returns a boolean if a field has been set.

### GetServerType

`func (o *HclHwCatalogInfo) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *HclHwCatalogInfo) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *HclHwCatalogInfo) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *HclHwCatalogInfo) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetServerHwCatalogInfo

`func (o *HclHwCatalogInfo) GetServerHwCatalogInfo() []HclServerHwCatalogInfoRelationship`

GetServerHwCatalogInfo returns the ServerHwCatalogInfo field if non-nil, zero value otherwise.

### GetServerHwCatalogInfoOk

`func (o *HclHwCatalogInfo) GetServerHwCatalogInfoOk() (*[]HclServerHwCatalogInfoRelationship, bool)`

GetServerHwCatalogInfoOk returns a tuple with the ServerHwCatalogInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHwCatalogInfo

`func (o *HclHwCatalogInfo) SetServerHwCatalogInfo(v []HclServerHwCatalogInfoRelationship)`

SetServerHwCatalogInfo sets ServerHwCatalogInfo field to given value.

### HasServerHwCatalogInfo

`func (o *HclHwCatalogInfo) HasServerHwCatalogInfo() bool`

HasServerHwCatalogInfo returns a boolean if a field has been set.

### SetServerHwCatalogInfoNil

`func (o *HclHwCatalogInfo) SetServerHwCatalogInfoNil(b bool)`

 SetServerHwCatalogInfoNil sets the value for ServerHwCatalogInfo to be an explicit nil

### UnsetServerHwCatalogInfo
`func (o *HclHwCatalogInfo) UnsetServerHwCatalogInfo()`

UnsetServerHwCatalogInfo ensures that no value is present for ServerHwCatalogInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


