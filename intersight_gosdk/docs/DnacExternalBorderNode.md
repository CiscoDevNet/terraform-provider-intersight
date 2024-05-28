# DnacExternalBorderNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "dnac.ExternalBorderNode"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "dnac.ExternalBorderNode"]
**ExternalBorderNodeId** | Pointer to **string** | External border node&#39;s id. | [optional] 
**ExternalBorderNodeName** | Pointer to **string** | External border node&#39;s name. | [optional] 
**FabricSiteId** | Pointer to **string** | Fabric Site id in UUID format. | [optional] 
**ImportExternalRoutes** | Pointer to **bool** | Flag to determine if Border Node is External or Internal. | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 

## Methods

### NewDnacExternalBorderNode

`func NewDnacExternalBorderNode(classId string, objectType string, ) *DnacExternalBorderNode`

NewDnacExternalBorderNode instantiates a new DnacExternalBorderNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnacExternalBorderNodeWithDefaults

`func NewDnacExternalBorderNodeWithDefaults() *DnacExternalBorderNode`

NewDnacExternalBorderNodeWithDefaults instantiates a new DnacExternalBorderNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *DnacExternalBorderNode) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *DnacExternalBorderNode) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *DnacExternalBorderNode) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *DnacExternalBorderNode) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *DnacExternalBorderNode) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *DnacExternalBorderNode) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExternalBorderNodeId

`func (o *DnacExternalBorderNode) GetExternalBorderNodeId() string`

GetExternalBorderNodeId returns the ExternalBorderNodeId field if non-nil, zero value otherwise.

### GetExternalBorderNodeIdOk

`func (o *DnacExternalBorderNode) GetExternalBorderNodeIdOk() (*string, bool)`

GetExternalBorderNodeIdOk returns a tuple with the ExternalBorderNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalBorderNodeId

`func (o *DnacExternalBorderNode) SetExternalBorderNodeId(v string)`

SetExternalBorderNodeId sets ExternalBorderNodeId field to given value.

### HasExternalBorderNodeId

`func (o *DnacExternalBorderNode) HasExternalBorderNodeId() bool`

HasExternalBorderNodeId returns a boolean if a field has been set.

### GetExternalBorderNodeName

`func (o *DnacExternalBorderNode) GetExternalBorderNodeName() string`

GetExternalBorderNodeName returns the ExternalBorderNodeName field if non-nil, zero value otherwise.

### GetExternalBorderNodeNameOk

`func (o *DnacExternalBorderNode) GetExternalBorderNodeNameOk() (*string, bool)`

GetExternalBorderNodeNameOk returns a tuple with the ExternalBorderNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalBorderNodeName

`func (o *DnacExternalBorderNode) SetExternalBorderNodeName(v string)`

SetExternalBorderNodeName sets ExternalBorderNodeName field to given value.

### HasExternalBorderNodeName

`func (o *DnacExternalBorderNode) HasExternalBorderNodeName() bool`

HasExternalBorderNodeName returns a boolean if a field has been set.

### GetFabricSiteId

`func (o *DnacExternalBorderNode) GetFabricSiteId() string`

GetFabricSiteId returns the FabricSiteId field if non-nil, zero value otherwise.

### GetFabricSiteIdOk

`func (o *DnacExternalBorderNode) GetFabricSiteIdOk() (*string, bool)`

GetFabricSiteIdOk returns a tuple with the FabricSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFabricSiteId

`func (o *DnacExternalBorderNode) SetFabricSiteId(v string)`

SetFabricSiteId sets FabricSiteId field to given value.

### HasFabricSiteId

`func (o *DnacExternalBorderNode) HasFabricSiteId() bool`

HasFabricSiteId returns a boolean if a field has been set.

### GetImportExternalRoutes

`func (o *DnacExternalBorderNode) GetImportExternalRoutes() bool`

GetImportExternalRoutes returns the ImportExternalRoutes field if non-nil, zero value otherwise.

### GetImportExternalRoutesOk

`func (o *DnacExternalBorderNode) GetImportExternalRoutesOk() (*bool, bool)`

GetImportExternalRoutesOk returns a tuple with the ImportExternalRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportExternalRoutes

`func (o *DnacExternalBorderNode) SetImportExternalRoutes(v bool)`

SetImportExternalRoutes sets ImportExternalRoutes field to given value.

### HasImportExternalRoutes

`func (o *DnacExternalBorderNode) HasImportExternalRoutes() bool`

HasImportExternalRoutes returns a boolean if a field has been set.

### GetRoles

`func (o *DnacExternalBorderNode) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *DnacExternalBorderNode) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *DnacExternalBorderNode) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *DnacExternalBorderNode) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *DnacExternalBorderNode) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *DnacExternalBorderNode) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


