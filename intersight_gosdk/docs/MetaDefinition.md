# MetaDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessPrivileges** | Pointer to [**[]MetaAccessPrivilege**](meta.AccessPrivilege.md) |  | [optional] 
**AncestorClasses** | Pointer to **[]string** |  | [optional] 
**DisplayNameMetas** | Pointer to [**[]MetaDisplayNameDefinition**](meta.DisplayNameDefinition.md) |  | [optional] 
**IsConcrete** | Pointer to **bool** | Boolean flag to specify whether the meta class is a concrete class or not. | [optional] [readonly] 
**MetaType** | Pointer to **string** | Indicates whether the meta class is a complex type or managed object. * &#x60;ManagedObject&#x60; - The meta.Definition object describes a managed object. * &#x60;ComplexType&#x60; - The meta.Definition object describes a nested complex type within a managed object. | [optional] [readonly] [default to "ManagedObject"]
**Name** | Pointer to **string** | The fully-qualified class name of the Managed Object or complex type. For example, \&quot;compute:Blade\&quot; where the Managed Object is \&quot;Blade\&quot; and the package is &#39;compute&#39;. | [optional] [readonly] 
**Namespace** | Pointer to **string** | The namespace of the meta. | [optional] [readonly] 
**ParentClass** | Pointer to **string** | The fully-qualified name of the parent metaclass in the class inheritance hierarchy. | [optional] [readonly] 
**PermissionSupported** | Pointer to **bool** | Boolean flag to specify whether instances of this class type can be specified in permissions for instance based access control. Permissions can be created for entire Intersight account or to a subset of resources (instance based access control). In the first release, permissions are supported for entire account or for a subset of organizations. | [optional] [readonly] 
**Properties** | Pointer to [**[]MetaPropDefinition**](meta.PropDefinition.md) |  | [optional] 
**RbacResource** | Pointer to **bool** | Boolean flag to specify whether instances of this class type can be assigned to resource groups that are part of an organization for access control. Inventoried physical/logical objects which needs access control should have rbacResource&#x3D;yes. These objects are not part of any organization by default like device registrations and should be assigned to organizations for access control. Profiles, policies, workflow definitions which are created by specifying organization need not have this flag set. | [optional] [readonly] 
**Relationships** | Pointer to [**[]MetaRelationshipDefinition**](meta.RelationshipDefinition.md) |  | [optional] 
**RestPath** | Pointer to **string** | Restful URL path for the meta. | [optional] [readonly] 
**Version** | Pointer to **string** | The version of the service that defines the meta-data. | [optional] [readonly] 

## Methods

### NewMetaDefinition

`func NewMetaDefinition() *MetaDefinition`

NewMetaDefinition instantiates a new MetaDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaDefinitionWithDefaults

`func NewMetaDefinitionWithDefaults() *MetaDefinition`

NewMetaDefinitionWithDefaults instantiates a new MetaDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessPrivileges

`func (o *MetaDefinition) GetAccessPrivileges() []MetaAccessPrivilege`

GetAccessPrivileges returns the AccessPrivileges field if non-nil, zero value otherwise.

### GetAccessPrivilegesOk

`func (o *MetaDefinition) GetAccessPrivilegesOk() (*[]MetaAccessPrivilege, bool)`

GetAccessPrivilegesOk returns a tuple with the AccessPrivileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPrivileges

`func (o *MetaDefinition) SetAccessPrivileges(v []MetaAccessPrivilege)`

SetAccessPrivileges sets AccessPrivileges field to given value.

### HasAccessPrivileges

`func (o *MetaDefinition) HasAccessPrivileges() bool`

HasAccessPrivileges returns a boolean if a field has been set.

### GetAncestorClasses

`func (o *MetaDefinition) GetAncestorClasses() []string`

GetAncestorClasses returns the AncestorClasses field if non-nil, zero value otherwise.

### GetAncestorClassesOk

`func (o *MetaDefinition) GetAncestorClassesOk() (*[]string, bool)`

GetAncestorClassesOk returns a tuple with the AncestorClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestorClasses

`func (o *MetaDefinition) SetAncestorClasses(v []string)`

SetAncestorClasses sets AncestorClasses field to given value.

### HasAncestorClasses

`func (o *MetaDefinition) HasAncestorClasses() bool`

HasAncestorClasses returns a boolean if a field has been set.

### GetDisplayNameMetas

`func (o *MetaDefinition) GetDisplayNameMetas() []MetaDisplayNameDefinition`

GetDisplayNameMetas returns the DisplayNameMetas field if non-nil, zero value otherwise.

### GetDisplayNameMetasOk

`func (o *MetaDefinition) GetDisplayNameMetasOk() (*[]MetaDisplayNameDefinition, bool)`

GetDisplayNameMetasOk returns a tuple with the DisplayNameMetas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNameMetas

`func (o *MetaDefinition) SetDisplayNameMetas(v []MetaDisplayNameDefinition)`

SetDisplayNameMetas sets DisplayNameMetas field to given value.

### HasDisplayNameMetas

`func (o *MetaDefinition) HasDisplayNameMetas() bool`

HasDisplayNameMetas returns a boolean if a field has been set.

### GetIsConcrete

`func (o *MetaDefinition) GetIsConcrete() bool`

GetIsConcrete returns the IsConcrete field if non-nil, zero value otherwise.

### GetIsConcreteOk

`func (o *MetaDefinition) GetIsConcreteOk() (*bool, bool)`

GetIsConcreteOk returns a tuple with the IsConcrete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConcrete

`func (o *MetaDefinition) SetIsConcrete(v bool)`

SetIsConcrete sets IsConcrete field to given value.

### HasIsConcrete

`func (o *MetaDefinition) HasIsConcrete() bool`

HasIsConcrete returns a boolean if a field has been set.

### GetMetaType

`func (o *MetaDefinition) GetMetaType() string`

GetMetaType returns the MetaType field if non-nil, zero value otherwise.

### GetMetaTypeOk

`func (o *MetaDefinition) GetMetaTypeOk() (*string, bool)`

GetMetaTypeOk returns a tuple with the MetaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaType

`func (o *MetaDefinition) SetMetaType(v string)`

SetMetaType sets MetaType field to given value.

### HasMetaType

`func (o *MetaDefinition) HasMetaType() bool`

HasMetaType returns a boolean if a field has been set.

### GetName

`func (o *MetaDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetaDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetaDefinition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetaDefinition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *MetaDefinition) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *MetaDefinition) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *MetaDefinition) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *MetaDefinition) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetParentClass

`func (o *MetaDefinition) GetParentClass() string`

GetParentClass returns the ParentClass field if non-nil, zero value otherwise.

### GetParentClassOk

`func (o *MetaDefinition) GetParentClassOk() (*string, bool)`

GetParentClassOk returns a tuple with the ParentClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentClass

`func (o *MetaDefinition) SetParentClass(v string)`

SetParentClass sets ParentClass field to given value.

### HasParentClass

`func (o *MetaDefinition) HasParentClass() bool`

HasParentClass returns a boolean if a field has been set.

### GetPermissionSupported

`func (o *MetaDefinition) GetPermissionSupported() bool`

GetPermissionSupported returns the PermissionSupported field if non-nil, zero value otherwise.

### GetPermissionSupportedOk

`func (o *MetaDefinition) GetPermissionSupportedOk() (*bool, bool)`

GetPermissionSupportedOk returns a tuple with the PermissionSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionSupported

`func (o *MetaDefinition) SetPermissionSupported(v bool)`

SetPermissionSupported sets PermissionSupported field to given value.

### HasPermissionSupported

`func (o *MetaDefinition) HasPermissionSupported() bool`

HasPermissionSupported returns a boolean if a field has been set.

### GetProperties

`func (o *MetaDefinition) GetProperties() []MetaPropDefinition`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *MetaDefinition) GetPropertiesOk() (*[]MetaPropDefinition, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *MetaDefinition) SetProperties(v []MetaPropDefinition)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *MetaDefinition) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetRbacResource

`func (o *MetaDefinition) GetRbacResource() bool`

GetRbacResource returns the RbacResource field if non-nil, zero value otherwise.

### GetRbacResourceOk

`func (o *MetaDefinition) GetRbacResourceOk() (*bool, bool)`

GetRbacResourceOk returns a tuple with the RbacResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRbacResource

`func (o *MetaDefinition) SetRbacResource(v bool)`

SetRbacResource sets RbacResource field to given value.

### HasRbacResource

`func (o *MetaDefinition) HasRbacResource() bool`

HasRbacResource returns a boolean if a field has been set.

### GetRelationships

`func (o *MetaDefinition) GetRelationships() []MetaRelationshipDefinition`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *MetaDefinition) GetRelationshipsOk() (*[]MetaRelationshipDefinition, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *MetaDefinition) SetRelationships(v []MetaRelationshipDefinition)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *MetaDefinition) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetRestPath

`func (o *MetaDefinition) GetRestPath() string`

GetRestPath returns the RestPath field if non-nil, zero value otherwise.

### GetRestPathOk

`func (o *MetaDefinition) GetRestPathOk() (*string, bool)`

GetRestPathOk returns a tuple with the RestPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestPath

`func (o *MetaDefinition) SetRestPath(v string)`

SetRestPath sets RestPath field to given value.

### HasRestPath

`func (o *MetaDefinition) HasRestPath() bool`

HasRestPath returns a boolean if a field has been set.

### GetVersion

`func (o *MetaDefinition) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MetaDefinition) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MetaDefinition) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MetaDefinition) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


