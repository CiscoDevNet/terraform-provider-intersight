# MetaDefinitionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "meta.Definition"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "meta.Definition"]
**AccessPrivileges** | Pointer to [**[]MetaAccessPrivilege**](MetaAccessPrivilege.md) |  | [optional] 
**AncestorClasses** | Pointer to **[]string** |  | [optional] 
**DisplayNameMetas** | Pointer to [**[]MetaDisplayNameDefinition**](MetaDisplayNameDefinition.md) |  | [optional] 
**IdentityConstraints** | Pointer to [**[]MetaIdentityDefinition**](MetaIdentityDefinition.md) |  | [optional] 
**IsConcrete** | Pointer to **bool** | Boolean flag to specify whether the meta class is a concrete class or not. | [optional] [readonly] 
**MetaType** | Pointer to **string** | Indicates whether the meta class is a complex type or managed object. * &#x60;ManagedObject&#x60; - The meta.Definition object describes a managed object. * &#x60;ComplexType&#x60; - The meta.Definition object describes a nested complex type within a managed object. | [optional] [readonly] [default to "ManagedObject"]
**Name** | Pointer to **string** | The fully-qualified class name of the Managed Object or complex type. For example, \&quot;compute:Blade\&quot; where the Managed Object is \&quot;Blade\&quot; and the package is &#39;compute&#39;. | [optional] [readonly] 
**Namespace** | Pointer to **string** | The namespace of the meta. | [optional] [readonly] 
**ParentClass** | Pointer to **string** | The fully-qualified name of the parent metaclass in the class inheritance hierarchy. | [optional] [readonly] 
**PermissionSupported** | Pointer to **bool** | Boolean flag to specify whether instances of this class type can be specified in permissions for instance based access control. Permissions can be created for entire Intersight account or to a subset of resources (instance based access control). In the first release, permissions are supported for entire account or for a subset of organizations. | [optional] [readonly] 
**Properties** | Pointer to [**[]MetaPropDefinition**](MetaPropDefinition.md) |  | [optional] 
**RbacResource** | Pointer to **bool** | Boolean flag to specify whether instances of this class type can be assigned to resource groups that are part of an organization for access control. Inventoried physical/logical objects which needs access control should have rbacResource&#x3D;true. These objects are not part of any organization by default like device registrations and should be assigned to organizations for access control. Profiles, policies, workflow definitions which are created by specifying organization need not have this flag set. | [optional] [readonly] 
**Relationships** | Pointer to [**[]MetaRelationshipDefinition**](MetaRelationshipDefinition.md) |  | [optional] 
**ResourcePoolTypes** | Pointer to **[]string** |  | [optional] 
**RestPath** | Pointer to **string** | Restful URL path for the meta. | [optional] [readonly] 
**Version** | Pointer to **string** | The version of the service that defines the meta-data. | [optional] [readonly] 

## Methods

### NewMetaDefinitionAllOf

`func NewMetaDefinitionAllOf(classId string, objectType string, ) *MetaDefinitionAllOf`

NewMetaDefinitionAllOf instantiates a new MetaDefinitionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaDefinitionAllOfWithDefaults

`func NewMetaDefinitionAllOfWithDefaults() *MetaDefinitionAllOf`

NewMetaDefinitionAllOfWithDefaults instantiates a new MetaDefinitionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *MetaDefinitionAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *MetaDefinitionAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *MetaDefinitionAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *MetaDefinitionAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MetaDefinitionAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MetaDefinitionAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccessPrivileges

`func (o *MetaDefinitionAllOf) GetAccessPrivileges() []MetaAccessPrivilege`

GetAccessPrivileges returns the AccessPrivileges field if non-nil, zero value otherwise.

### GetAccessPrivilegesOk

`func (o *MetaDefinitionAllOf) GetAccessPrivilegesOk() (*[]MetaAccessPrivilege, bool)`

GetAccessPrivilegesOk returns a tuple with the AccessPrivileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPrivileges

`func (o *MetaDefinitionAllOf) SetAccessPrivileges(v []MetaAccessPrivilege)`

SetAccessPrivileges sets AccessPrivileges field to given value.

### HasAccessPrivileges

`func (o *MetaDefinitionAllOf) HasAccessPrivileges() bool`

HasAccessPrivileges returns a boolean if a field has been set.

### SetAccessPrivilegesNil

`func (o *MetaDefinitionAllOf) SetAccessPrivilegesNil(b bool)`

 SetAccessPrivilegesNil sets the value for AccessPrivileges to be an explicit nil

### UnsetAccessPrivileges
`func (o *MetaDefinitionAllOf) UnsetAccessPrivileges()`

UnsetAccessPrivileges ensures that no value is present for AccessPrivileges, not even an explicit nil
### GetAncestorClasses

`func (o *MetaDefinitionAllOf) GetAncestorClasses() []string`

GetAncestorClasses returns the AncestorClasses field if non-nil, zero value otherwise.

### GetAncestorClassesOk

`func (o *MetaDefinitionAllOf) GetAncestorClassesOk() (*[]string, bool)`

GetAncestorClassesOk returns a tuple with the AncestorClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestorClasses

`func (o *MetaDefinitionAllOf) SetAncestorClasses(v []string)`

SetAncestorClasses sets AncestorClasses field to given value.

### HasAncestorClasses

`func (o *MetaDefinitionAllOf) HasAncestorClasses() bool`

HasAncestorClasses returns a boolean if a field has been set.

### SetAncestorClassesNil

`func (o *MetaDefinitionAllOf) SetAncestorClassesNil(b bool)`

 SetAncestorClassesNil sets the value for AncestorClasses to be an explicit nil

### UnsetAncestorClasses
`func (o *MetaDefinitionAllOf) UnsetAncestorClasses()`

UnsetAncestorClasses ensures that no value is present for AncestorClasses, not even an explicit nil
### GetDisplayNameMetas

`func (o *MetaDefinitionAllOf) GetDisplayNameMetas() []MetaDisplayNameDefinition`

GetDisplayNameMetas returns the DisplayNameMetas field if non-nil, zero value otherwise.

### GetDisplayNameMetasOk

`func (o *MetaDefinitionAllOf) GetDisplayNameMetasOk() (*[]MetaDisplayNameDefinition, bool)`

GetDisplayNameMetasOk returns a tuple with the DisplayNameMetas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNameMetas

`func (o *MetaDefinitionAllOf) SetDisplayNameMetas(v []MetaDisplayNameDefinition)`

SetDisplayNameMetas sets DisplayNameMetas field to given value.

### HasDisplayNameMetas

`func (o *MetaDefinitionAllOf) HasDisplayNameMetas() bool`

HasDisplayNameMetas returns a boolean if a field has been set.

### SetDisplayNameMetasNil

`func (o *MetaDefinitionAllOf) SetDisplayNameMetasNil(b bool)`

 SetDisplayNameMetasNil sets the value for DisplayNameMetas to be an explicit nil

### UnsetDisplayNameMetas
`func (o *MetaDefinitionAllOf) UnsetDisplayNameMetas()`

UnsetDisplayNameMetas ensures that no value is present for DisplayNameMetas, not even an explicit nil
### GetIdentityConstraints

`func (o *MetaDefinitionAllOf) GetIdentityConstraints() []MetaIdentityDefinition`

GetIdentityConstraints returns the IdentityConstraints field if non-nil, zero value otherwise.

### GetIdentityConstraintsOk

`func (o *MetaDefinitionAllOf) GetIdentityConstraintsOk() (*[]MetaIdentityDefinition, bool)`

GetIdentityConstraintsOk returns a tuple with the IdentityConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityConstraints

`func (o *MetaDefinitionAllOf) SetIdentityConstraints(v []MetaIdentityDefinition)`

SetIdentityConstraints sets IdentityConstraints field to given value.

### HasIdentityConstraints

`func (o *MetaDefinitionAllOf) HasIdentityConstraints() bool`

HasIdentityConstraints returns a boolean if a field has been set.

### SetIdentityConstraintsNil

`func (o *MetaDefinitionAllOf) SetIdentityConstraintsNil(b bool)`

 SetIdentityConstraintsNil sets the value for IdentityConstraints to be an explicit nil

### UnsetIdentityConstraints
`func (o *MetaDefinitionAllOf) UnsetIdentityConstraints()`

UnsetIdentityConstraints ensures that no value is present for IdentityConstraints, not even an explicit nil
### GetIsConcrete

`func (o *MetaDefinitionAllOf) GetIsConcrete() bool`

GetIsConcrete returns the IsConcrete field if non-nil, zero value otherwise.

### GetIsConcreteOk

`func (o *MetaDefinitionAllOf) GetIsConcreteOk() (*bool, bool)`

GetIsConcreteOk returns a tuple with the IsConcrete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConcrete

`func (o *MetaDefinitionAllOf) SetIsConcrete(v bool)`

SetIsConcrete sets IsConcrete field to given value.

### HasIsConcrete

`func (o *MetaDefinitionAllOf) HasIsConcrete() bool`

HasIsConcrete returns a boolean if a field has been set.

### GetMetaType

`func (o *MetaDefinitionAllOf) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *MetaDefinitionAllOf) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *MetaDefinitionAllOf) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.

### HasMetaType

`func (o *MetaDefinitionAllOf) HasMetaType() bool`

HasMetaType returns a boolean if a field has been set.

### GetName

`func (o *MetaDefinitionAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetaDefinitionAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetaDefinitionAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetaDefinitionAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *MetaDefinitionAllOf) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *MetaDefinitionAllOf) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *MetaDefinitionAllOf) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *MetaDefinitionAllOf) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetParentClass

`func (o *MetaDefinitionAllOf) GetParentClass() string`

GetParentClass returns the ParentClass field if non-nil, zero value otherwise.

### GetParentClassOk

`func (o *MetaDefinitionAllOf) GetParentClassOk() (*string, bool)`

GetParentClassOk returns a tuple with the ParentClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentClass

`func (o *MetaDefinitionAllOf) SetParentClass(v string)`

SetParentClass sets ParentClass field to given value.

### HasParentClass

`func (o *MetaDefinitionAllOf) HasParentClass() bool`

HasParentClass returns a boolean if a field has been set.

### GetPermissionSupported

`func (o *MetaDefinitionAllOf) GetPermissionSupported() bool`

GetPermissionSupported returns the PermissionSupported field if non-nil, zero value otherwise.

### GetPermissionSupportedOk

`func (o *MetaDefinitionAllOf) GetPermissionSupportedOk() (*bool, bool)`

GetPermissionSupportedOk returns a tuple with the PermissionSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionSupported

`func (o *MetaDefinitionAllOf) SetPermissionSupported(v bool)`

SetPermissionSupported sets PermissionSupported field to given value.

### HasPermissionSupported

`func (o *MetaDefinitionAllOf) HasPermissionSupported() bool`

HasPermissionSupported returns a boolean if a field has been set.

### GetProperties

`func (o *MetaDefinitionAllOf) GetProperties() []MetaPropDefinition`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *MetaDefinitionAllOf) GetPropertiesOk() (*[]MetaPropDefinition, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *MetaDefinitionAllOf) SetProperties(v []MetaPropDefinition)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *MetaDefinitionAllOf) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *MetaDefinitionAllOf) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *MetaDefinitionAllOf) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetRbacResource

`func (o *MetaDefinitionAllOf) GetRbacResource() bool`

GetRbacResource returns the RbacResource field if non-nil, zero value otherwise.

### GetRbacResourceOk

`func (o *MetaDefinitionAllOf) GetRbacResourceOk() (*bool, bool)`

GetRbacResourceOk returns a tuple with the RbacResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRbacResource

`func (o *MetaDefinitionAllOf) SetRbacResource(v bool)`

SetRbacResource sets RbacResource field to given value.

### HasRbacResource

`func (o *MetaDefinitionAllOf) HasRbacResource() bool`

HasRbacResource returns a boolean if a field has been set.

### GetRelationships

`func (o *MetaDefinitionAllOf) GetRelationships() []MetaRelationshipDefinition`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *MetaDefinitionAllOf) GetRelationshipsOk() (*[]MetaRelationshipDefinition, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *MetaDefinitionAllOf) SetRelationships(v []MetaRelationshipDefinition)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *MetaDefinitionAllOf) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### SetRelationshipsNil

`func (o *MetaDefinitionAllOf) SetRelationshipsNil(b bool)`

 SetRelationshipsNil sets the value for Relationships to be an explicit nil

### UnsetRelationships
`func (o *MetaDefinitionAllOf) UnsetRelationships()`

UnsetRelationships ensures that no value is present for Relationships, not even an explicit nil
### GetResourcePoolTypes

`func (o *MetaDefinitionAllOf) GetResourcePoolTypes() []string`

GetResourcePoolTypes returns the ResourcePoolTypes field if non-nil, zero value otherwise.

### GetResourcePoolTypesOk

`func (o *MetaDefinitionAllOf) GetResourcePoolTypesOk() (*[]string, bool)`

GetResourcePoolTypesOk returns a tuple with the ResourcePoolTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolTypes

`func (o *MetaDefinitionAllOf) SetResourcePoolTypes(v []string)`

SetResourcePoolTypes sets ResourcePoolTypes field to given value.

### HasResourcePoolTypes

`func (o *MetaDefinitionAllOf) HasResourcePoolTypes() bool`

HasResourcePoolTypes returns a boolean if a field has been set.

### SetResourcePoolTypesNil

`func (o *MetaDefinitionAllOf) SetResourcePoolTypesNil(b bool)`

 SetResourcePoolTypesNil sets the value for ResourcePoolTypes to be an explicit nil

### UnsetResourcePoolTypes
`func (o *MetaDefinitionAllOf) UnsetResourcePoolTypes()`

UnsetResourcePoolTypes ensures that no value is present for ResourcePoolTypes, not even an explicit nil
### GetRestPath

`func (o *MetaDefinitionAllOf) GetRestPath() string`

GetRestPath returns the RestPath field if non-nil, zero value otherwise.

### GetRestPathOk

`func (o *MetaDefinitionAllOf) GetRestPathOk() (*string, bool)`

GetRestPathOk returns a tuple with the RestPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestPath

`func (o *MetaDefinitionAllOf) SetRestPath(v string)`

SetRestPath sets RestPath field to given value.

### HasRestPath

`func (o *MetaDefinitionAllOf) HasRestPath() bool`

HasRestPath returns a boolean if a field has been set.

### GetVersion

`func (o *MetaDefinitionAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MetaDefinitionAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MetaDefinitionAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MetaDefinitionAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


