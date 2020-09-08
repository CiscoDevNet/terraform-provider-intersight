# \IppoolApi

All URIs are relative to *https://intersight.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIppoolIpLease**](IppoolApi.md#CreateIppoolIpLease) | **Post** /api/v1/ippool/IpLeases | Create a &#39;ippool.IpLease&#39; resource.
[**CreateIppoolPool**](IppoolApi.md#CreateIppoolPool) | **Post** /api/v1/ippool/Pools | Create a &#39;ippool.Pool&#39; resource.
[**DeleteIppoolIpLease**](IppoolApi.md#DeleteIppoolIpLease) | **Delete** /api/v1/ippool/IpLeases/{Moid} | Delete a &#39;ippool.IpLease&#39; resource.
[**DeleteIppoolPool**](IppoolApi.md#DeleteIppoolPool) | **Delete** /api/v1/ippool/Pools/{Moid} | Delete a &#39;ippool.Pool&#39; resource.
[**GetIppoolIpLeaseByMoid**](IppoolApi.md#GetIppoolIpLeaseByMoid) | **Get** /api/v1/ippool/IpLeases/{Moid} | Read a &#39;ippool.IpLease&#39; resource.
[**GetIppoolIpLeaseList**](IppoolApi.md#GetIppoolIpLeaseList) | **Get** /api/v1/ippool/IpLeases | Read a &#39;ippool.IpLease&#39; resource.
[**GetIppoolPoolByMoid**](IppoolApi.md#GetIppoolPoolByMoid) | **Get** /api/v1/ippool/Pools/{Moid} | Read a &#39;ippool.Pool&#39; resource.
[**GetIppoolPoolList**](IppoolApi.md#GetIppoolPoolList) | **Get** /api/v1/ippool/Pools | Read a &#39;ippool.Pool&#39; resource.
[**GetIppoolPoolMemberByMoid**](IppoolApi.md#GetIppoolPoolMemberByMoid) | **Get** /api/v1/ippool/PoolMembers/{Moid} | Read a &#39;ippool.PoolMember&#39; resource.
[**GetIppoolPoolMemberList**](IppoolApi.md#GetIppoolPoolMemberList) | **Get** /api/v1/ippool/PoolMembers | Read a &#39;ippool.PoolMember&#39; resource.
[**GetIppoolShadowBlockByMoid**](IppoolApi.md#GetIppoolShadowBlockByMoid) | **Get** /api/v1/ippool/ShadowBlocks/{Moid} | Read a &#39;ippool.ShadowBlock&#39; resource.
[**GetIppoolShadowBlockList**](IppoolApi.md#GetIppoolShadowBlockList) | **Get** /api/v1/ippool/ShadowBlocks | Read a &#39;ippool.ShadowBlock&#39; resource.
[**GetIppoolShadowPoolByMoid**](IppoolApi.md#GetIppoolShadowPoolByMoid) | **Get** /api/v1/ippool/ShadowPools/{Moid} | Read a &#39;ippool.ShadowPool&#39; resource.
[**GetIppoolShadowPoolList**](IppoolApi.md#GetIppoolShadowPoolList) | **Get** /api/v1/ippool/ShadowPools | Read a &#39;ippool.ShadowPool&#39; resource.
[**GetIppoolUniverseByMoid**](IppoolApi.md#GetIppoolUniverseByMoid) | **Get** /api/v1/ippool/Universes/{Moid} | Read a &#39;ippool.Universe&#39; resource.
[**GetIppoolUniverseList**](IppoolApi.md#GetIppoolUniverseList) | **Get** /api/v1/ippool/Universes | Read a &#39;ippool.Universe&#39; resource.
[**PatchIppoolPool**](IppoolApi.md#PatchIppoolPool) | **Patch** /api/v1/ippool/Pools/{Moid} | Update a &#39;ippool.Pool&#39; resource.
[**UpdateIppoolPool**](IppoolApi.md#UpdateIppoolPool) | **Post** /api/v1/ippool/Pools/{Moid} | Update a &#39;ippool.Pool&#39; resource.



## CreateIppoolIpLease

> IppoolIpLease CreateIppoolIpLease(ctx).IppoolIpLease(ippoolIpLease).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()

Create a 'ippool.IpLease' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ippoolIpLease := openapiclient.ippool.IpLease{AccountMoid: "AccountMoid_example", ClassId: "ClassId_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", ObjectType: "ObjectType_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{openapiclient.mo.Tag{Key: "Key_example", Value: "Value_example"}), VersionContext: openapiclient.mo.VersionContext{ClassId: "ClassId_example", ObjectType: "ObjectType_example", InterestedMos: []MoMoRef{openapiclient.mo.MoRef{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example"}), RefMo: openapiclient.mo.MoRef{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example"}, Timestamp: "TODO", Version: "Version_example", VersionType: "VersionType_example"}, Ancestors: []MoBaseMoRelationship{openapiclient.mo.BaseMo.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{openapiclient.mo.Tag{Key: "Key_example", Value: "Value_example"}), VersionContext: openapiclient.mo.VersionContext{ClassId: "ClassId_example", ObjectType: "ObjectType_example", InterestedMos: []MoMoRef{), RefMo: , Timestamp: "TODO", Version: "Version_example", VersionType: "VersionType_example"}, Ancestors: []MoBaseMoRelationship{openapiclient.mo.BaseMo.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }}), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }}), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, IpV4Address: "IpV4Address_example", IpV4Config: openapiclient.ippool.IpV4Config{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Gateway: "Gateway_example", Netmask: "Netmask_example", PrimaryDns: "PrimaryDns_example", SecondaryDns: "SecondaryDns_example"}, AssignedToEntity: , Pool: openapiclient.ippool.Pool.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Assigned: int64(123), AssignmentOrder: "AssignmentOrder_example", Size: int64(123), IpV4Blocks: []IppoolIpBlock{openapiclient.ippool.IpBlock{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Size: int64(123), From: "From_example", To: "To_example"}), IpV4Config: openapiclient.ippool.IpV4Config{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Gateway: "Gateway_example", Netmask: "Netmask_example", PrimaryDns: "PrimaryDns_example", SecondaryDns: "SecondaryDns_example"}, V4Assigned: int64(123), V4Size: int64(123), V6Assigned: int64(123), V6Size: int64(123), Organization: openapiclient.organization.Organization.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Account: openapiclient.iam.Account.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", Status: "Status_example", Var2LicenseReservationOp: openapiclient.license.LicenseReservationOp.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, AuthCode: "AuthCode_example", AuthCodeInstalled: false, ConfirmCode: "ConfirmCode_example", GenerateRequestCode: false, GenerateReturnCode: false, RequestCode: "RequestCode_example", ReturnCode: "ReturnCode_example", Account: openapiclient.iam.Account.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", Status: "Status_example", Var2LicenseReservationOp: openapiclient.license.LicenseReservationOp.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, AuthCode: "AuthCode_example", AuthCodeInstalled: false, ConfirmCode: "ConfirmCode_example", GenerateRequestCode: false, GenerateReturnCode: false, RequestCode: "RequestCode_example", ReturnCode: "ReturnCode_example", Account: }, AppRegistrations: []IamAppRegistrationRelationship{openapiclient.iam.AppRegistration.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, ClientId: "ClientId_example", ClientName: "ClientName_example", ClientSecret: "ClientSecret_example", ClientType: "ClientType_example", Description: "Description_example", GrantTypes: []string{"GrantTypes_example"), RedirectUris: []string{"RedirectUris_example"), RenewClientSecret: false, ResponseTypes: []string{"ResponseTypes_example"), RevocationTimestamp: "TODO", Revoke: false, Account: , OauthTokens: []IamOAuthTokenRelationship{openapiclient.iam.OAuthToken.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, AccessExpirationTime: "TODO", ClientId: "ClientId_example", ClientIpAddress: "ClientIpAddress_example", ClientName: "ClientName_example", ExpirationTime: "TODO", LastLoginClient: "LastLoginClient_example", LastLoginTime: "TODO", TokenId: "TokenId_example", UserMeta: openapiclient.iam.ClientMeta{ClassId: "ClassId_example", ObjectType: "ObjectType_example", DeviceModel: "DeviceModel_example", UserAgent: "UserAgent_example"}, AppRegistration: openapiclient.iam.AppRegistration.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, ClientId: "ClientId_example", ClientName: "ClientName_example", ClientSecret: "ClientSecret_example", ClientType: "ClientType_example", Description: "Description_example", GrantTypes: []string{"GrantTypes_example"), RedirectUris: []string{"RedirectUris_example"), RenewClientSecret: false, ResponseTypes: []string{"ResponseTypes_example"), RevocationTimestamp: "TODO", Revoke: false, Account: , OauthTokens: []IamOAuthTokenRelationship{openapiclient.iam.OAuthToken.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, AccessExpirationTime: "TODO", ClientId: "ClientId_example", ClientIpAddress: "ClientIpAddress_example", ClientName: "ClientName_example", ExpirationTime: "TODO", LastLoginClient: "LastLoginClient_example", LastLoginTime: "TODO", TokenId: "TokenId_example", UserMeta: openapiclient.iam.ClientMeta{ClassId: "ClassId_example", ObjectType: "ObjectType_example", DeviceModel: "DeviceModel_example", UserAgent: "UserAgent_example"}, AppRegistration: , Permission: openapiclient.iam.Permission.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Account: , EndPointRoles: []IamEndPointRoleRelationship{openapiclient.iam.EndPointRole.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", RoleType: "RoleType_example", Type: "Type_example", Account: , EndPointPrivileges: []IamEndPointPrivilegeRelationship{openapiclient.iam.EndPointPrivilege.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Type: "Type_example", System: openapiclient.iam.System.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, EndPointPrivileges: []IamEndPointPrivilegeRelationship{openapiclient.iam.EndPointPrivilege.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Type: "Type_example", System: openapiclient.iam.System.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, EndPointPrivileges: []IamEndPointPrivilegeRelationship{), EndPointRoles: []IamEndPointRoleRelationship{openapiclient.iam.EndPointRole.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", RoleType: "RoleType_example", Type: "Type_example", Account: , EndPointPrivileges: []IamEndPointPrivilegeRelationship{), System: }), Idp: openapiclient.iam.Idp.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DomainName: "DomainName_example", IdpEntityId: "IdpEntityId_example", Metadata: "Metadata_example", Name: "Name_example", Type: "Type_example", Account: , LdapPolicy: openapiclient.iam.LdapPolicy.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", BaseProperties: openapiclient.iam.LdapBaseProperties{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Attribute: "Attribute_example", BaseDn: "BaseDn_example", BindDn: "BindDn_example", BindMethod: "BindMethod_example", Domain: "Domain_example", EnableEncryption: false, EnableGroupAuthorization: false, Filter: "Filter_example", GroupAttribute: "GroupAttribute_example", IsPasswordSet: false, NestedGroupSearchDepth: int64(123), Password: "Password_example", Timeout: int64(123)}, DnsParameters: openapiclient.iam.LdapDnsParameters{ClassId: "ClassId_example", ObjectType: "ObjectType_example", SearchDomain: "SearchDomain_example", SearchForest: "SearchForest_example", Source: "Source_example"}, EnableDns: false, Enabled: false, UserSearchPrecedence: "UserSearchPrecedence_example", Var0Idp: openapiclient.iam.Idp.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DomainName: "DomainName_example", IdpEntityId: "IdpEntityId_example", Metadata: "Metadata_example", Name: "Name_example", Type: "Type_example", Account: , LdapPolicy: openapiclient.iam.LdapPolicy.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", BaseProperties: openapiclient.iam.LdapBaseProperties{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Attribute: "Attribute_example", BaseDn: "BaseDn_example", BindDn: "BindDn_example", BindMethod: "BindMethod_example", Domain: "Domain_example", EnableEncryption: false, EnableGroupAuthorization: false, Filter: "Filter_example", GroupAttribute: "GroupAttribute_example", IsPasswordSet: false, NestedGroupSearchDepth: int64(123), Password: "Password_example", Timeout: int64(123)}, DnsParameters: openapiclient.iam.LdapDnsParameters{ClassId: "ClassId_example", ObjectType: "ObjectType_example", SearchDomain: "SearchDomain_example", SearchForest: "SearchForest_example", Source: "Source_example"}, EnableDns: false, Enabled: false, UserSearchPrecedence: "UserSearchPrecedence_example", Var0Idp: , ApplianceAccount: , Groups: []IamLdapGroupRelationship{openapiclient.iam.LdapGroup.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Domain: "Domain_example", Name: "Name_example", EndPointRole: []IamEndPointRoleRelationship{), LdapPolicy: }), Organization: openapiclient.organization.Organization.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Account: , ResourceGroups: []ResourceGroupRelationship{openapiclient.resource.Group.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", PerTypeCombinedSelector: []ResourcePerTypeCombinedSelector{openapiclient.resource.PerTypeCombinedSelector{ClassId: "ClassId_example", ObjectType: "ObjectType_example", CombinedSelector: "CombinedSelector_example", EmptyFilter: false, SelectorObjectType: "SelectorObjectType_example"}), Qualifier: "Qualifier_example", Selectors: []ResourceSelector{openapiclient.resource.Selector{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Selector: "Selector_example"}), Account: , Organizations: []OrganizationOrganizationRelationship{)})}, Profiles: []PolicyAbstractConfigProfileRelationship{openapiclient.policy.AbstractConfigProfile.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Type: "Type_example", SrcTemplate: openapiclient.policy.AbstractProfile.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Type: "Type_example", SrcTemplate: openapiclient.policy.AbstractProfile.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Type: "Type_example", SrcTemplate: }}, Action: "Action_example", ConfigContext: openapiclient.policy.ConfigContext{ClassId: "ClassId_example", ObjectType: "ObjectType_example", ConfigState: "ConfigState_example", ControlAction: "ControlAction_example", ErrorState: "ErrorState_example", OperState: "OperState_example"}}), Providers: []IamLdapProviderRelationship{openapiclient.iam.LdapProvider.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Port: int64(123), Server: "Server_example", LdapPolicy: })}, System: , UserPreferences: []IamUserPreferenceRelationship{openapiclient.iam.UserPreference.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Preference: 123, UserUniqueIdentifier: "UserUniqueIdentifier_example", Idp: , IdpReference: openapiclient.iam.IdpReference.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DomainName: "DomainName_example", IdpEntityId: "IdpEntityId_example", MultiFactorAuthentication: false, Name: "Name_example", Account: , Idp: , UserPreferences: []IamUserPreferenceRelationship{openapiclient.iam.UserPreference.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Preference: 123, UserUniqueIdentifier: "UserUniqueIdentifier_example", Idp: , IdpReference: openapiclient.iam.IdpReference.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, DomainName: "DomainName_example", IdpEntityId: "IdpEntityId_example", MultiFactorAuthentication: false, Name: "Name_example", Account: , Idp: , UserPreferences: []IamUserPreferenceRelationship{), Usergroups: []IamUserGroupRelationship{openapiclient.iam.UserGroup.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", Idp: , Idpreference: , Permissions: []IamPermissionRelationship{openapiclient.iam.Permission.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Account: , EndPointRoles: []IamEndPointRoleRelationship{), ResourceRoles: []IamResourceRolesRelationship{openapiclient.iam.ResourceRoles.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, EndPointRoles: []IamEndPointRoleRelationship{), Permission: , Resource: , Roles: []IamRoleRelationship{openapiclient.iam.Role.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", PrivilegeNames: []string{"PrivilegeNames_example"), Account: , PrivilegeSets: []IamPrivilegeSetRelationship{openapiclient.iam.PrivilegeSet.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", PrivilegeNames: []string{"PrivilegeNames_example"), Account: , AssociatedPrivilegeSets: []IamPrivilegeSetRelationship{openapiclient.iam.PrivilegeSet.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", PrivilegeNames: []string{"PrivilegeNames_example"), Account: , AssociatedPrivilegeSets: []IamPrivilegeSetRelationship{), Privileges: []IamPrivilegeRelationship{openapiclient.iam.Privilege.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, HostnamePrefix: "HostnamePrefix_example", Method: "Method_example", Name: "Name_example", RestPath: "RestPath_example", UrlPrefix: "UrlPrefix_example", Account: , System: }), System: }), Privileges: []IamPrivilegeRelationship{openapiclient.iam.Privilege.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, HostnamePrefix: "HostnamePrefix_example", Method: "Method_example", Name: "Name_example", RestPath: "RestPath_example", UrlPrefix: "UrlPrefix_example", Account: , System: }), System: }), System: })}), Roles: []IamRoleRelationship{openapiclient.iam.Role.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", PrivilegeNames: []string{"PrivilegeNames_example"), Account: , PrivilegeSets: []IamPrivilegeSetRelationship{), System: }), SessionLimits: openapiclient.iam.SessionLimits.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, IdleTimeOut: int64(123), MaximumLimit: int64(123), PerUserLimit: int64(123), SessionTimeOut: int64(123), Account: , Permission: }, UserGroups: []IamUserGroupRelationship{openapiclient.iam.UserGroup.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", Idp: , Idpreference: , Permissions: []IamPermissionRelationship{), Qualifier: openapiclient.iam.Qualifier.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", Value: []string{"Value_example"), Usergroup: }, Users: []IamUserRelationship{openapiclient.iam.User.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, ClientIpAddress: "ClientIpAddress_example", Email: "Email_example", FirstName: "FirstName_example", LastLoginTime: "TODO", LastName: "LastName_example", Name: "Name_example", UserIdOrEmail: "UserIdOrEmail_example", UserType: "UserType_example", UserUniqueIdentifier: "UserUniqueIdentifier_example", ApiKeys: []IamApiKeyRelationship{openapiclient.iam.ApiKey.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, HashAlgorithm: "HashAlgorithm_example", KeySpec: openapiclient.pkix.KeyGenerationSpec{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Name: "Name_example"}, PrivateKey: "PrivateKey_example", Purpose: "Purpose_example", SigningAlgorithm: "SigningAlgorithm_example", Permission: , User: openapiclient.iam.User.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, ClientIpAddress: "ClientIpAddress_example", Email: "Email_example", FirstName: "FirstName_example", LastLoginTime: "TODO", LastName: "LastName_example", Name: "Name_example", UserIdOrEmail: "UserIdOrEmail_example", UserType: "UserType_example", UserUniqueIdentifier: "UserUniqueIdentifier_example", ApiKeys: []IamApiKeyRelationship{openapiclient.iam.ApiKey.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, HashAlgorithm: "HashAlgorithm_example", KeySpec: openapiclient.pkix.KeyGenerationSpec{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Name: "Name_example"}, PrivateKey: "PrivateKey_example", Purpose: "Purpose_example", SigningAlgorithm: "SigningAlgorithm_example", Permission: , User: }), AppRegistrations: []IamAppRegistrationRelationship{), Idp: , Idpreference: , LocalUserPassword: openapiclient.iam.LocalUserPassword.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, CurrentPassword: "CurrentPassword_example", NewPassword: "NewPassword_example", Password: 123, User: }, OauthTokens: []IamOAuthTokenRelationship{), Permissions: []IamPermissionRelationship{), Sessions: []IamSessionRelationship{openapiclient.iam.Session.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, AccountPermissions: []IamAccountPermissions{openapiclient.iam.AccountPermissions{ClassId: "ClassId_example", ObjectType: "ObjectType_example", AccountIdentifier: "AccountIdentifier_example", AccountName: "AccountName_example", AccountStatus: "AccountStatus_example", Permissions: []IamPermissionReference{openapiclient.iam.PermissionReference{ClassId: "ClassId_example", ObjectType: "ObjectType_example", PermissionIdentifier: "PermissionIdentifier_example", PermissionName: "PermissionName_example"})}), ClientIpAddress: "ClientIpAddress_example", Expiration: "TODO", IdleTimeExpiration: "TODO", LastLoginClient: "LastLoginClient_example", LastLoginTime: "TODO", Permission: , User: })}}), AppRegistrations: []IamAppRegistrationRelationship{), Idp: , Idpreference: , LocalUserPassword: openapiclient.iam.LocalUserPassword.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, CurrentPassword: "CurrentPassword_example", NewPassword: "NewPassword_example", Password: 123, User: }, OauthTokens: []IamOAuthTokenRelationship{), Permissions: []IamPermissionRelationship{), Sessions: []IamSessionRelationship{openapiclient.iam.Session.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, AccountPermissions: []IamAccountPermissions{openapiclient.iam.AccountPermissions{ClassId: "ClassId_example", ObjectType: "ObjectType_example", AccountIdentifier: "AccountIdentifier_example", AccountName: "AccountName_example", AccountStatus: "AccountStatus_example", Permissions: []IamPermissionReference{openapiclient.iam.PermissionReference{ClassId: "ClassId_example", ObjectType: "ObjectType_example", PermissionIdentifier: "PermissionIdentifier_example", PermissionName: "PermissionName_example"})}), ClientIpAddress: "ClientIpAddress_example", Expiration: "TODO", IdleTimeExpiration: "TODO", LastLoginClient: "LastLoginClient_example", LastLoginTime: "TODO", Permission: , User: })})}), Users: []IamUserRelationship{)}), Qualifier: openapiclient.iam.Qualifier.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", Value: []string{"Value_example"), Usergroup: }, Users: []IamUserRelationship{)}), Users: []IamUserRelationship{)}}), Usergroups: []IamUserGroupRelationship{), Users: []IamUserRelationship{)}}), Usergroups: []IamUserGroupRelationship{), Users: []IamUserRelationship{)}, ApplianceAccount: , Groups: []IamLdapGroupRelationship{openapiclient.iam.LdapGroup.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Domain: "Domain_example", Name: "Name_example", EndPointRole: []IamEndPointRoleRelationship{), LdapPolicy: }), Organization: , Profiles: []PolicyAbstractConfigProfileRelationship{openapiclient.policy.AbstractConfigProfile.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Type: "Type_example", SrcTemplate: , Action: "Action_example", ConfigContext: openapiclient.policy.ConfigContext{ClassId: "ClassId_example", ObjectType: "ObjectType_example", ConfigState: "ConfigState_example", ControlAction: "ControlAction_example", ErrorState: "ErrorState_example", OperState: "OperState_example"}}), Providers: []IamLdapProviderRelationship{openapiclient.iam.LdapProvider.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Port: int64(123), Server: "Server_example", LdapPolicy: })}, System: , UserPreferences: []IamUserPreferenceRelationship{), Usergroups: []IamUserGroupRelationship{), Users: []IamUserRelationship{)}, PrivilegeSets: []IamPrivilegeSetRelationship{), Privileges: []IamPrivilegeRelationship{), Roles: []IamRoleRelationship{), ServiceProvider: openapiclient.iam.ServiceProvider.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, EntityId: "EntityId_example", Metadata: "Metadata_example", Name: "Name_example", System: }}}), EndPointRoles: []IamEndPointRoleRelationship{), Idp: , PrivilegeSets: []IamPrivilegeSetRelationship{), Privileges: []IamPrivilegeRelationship{), Roles: []IamRoleRelationship{), ServiceProvider: openapiclient.iam.ServiceProvider.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, EntityId: "EntityId_example", Metadata: "Metadata_example", Name: "Name_example", System: }}}), System: }), ResourceRoles: []IamResourceRolesRelationship{openapiclient.iam.ResourceRoles.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, EndPointRoles: []IamEndPointRoleRelationship{), Permission: , Resource: , Roles: []IamRoleRelationship{)}), Roles: []IamRoleRelationship{), SessionLimits: openapiclient.iam.SessionLimits.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, IdleTimeOut: int64(123), MaximumLimit: int64(123), PerUserLimit: int64(123), SessionTimeOut: int64(123), Account: , Permission: }, UserGroups: []IamUserGroupRelationship{), Users: []IamUserRelationship{)}, User: }), Permission: , Roles: []IamRoleRelationship{), User: }, Permission: , User: }), Permission: , Roles: []IamRoleRelationship{), User: }), DomainGroups: []IamDomainGroupRelationship{openapiclient.iam.DomainGroup.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", Partition1: int64(123), Partition2: int64(123), Partition3: int64(123), PartitionKey: "PartitionKey_example", Usage: int64(123), Account: }), EndPointRoles: []IamEndPointRoleRelationship{), Idpreferences: []IamIdpReferenceRelationship{), Idps: []IamIdpRelationship{), Permissions: []IamPermissionRelationship{), PrivilegeSets: []IamPrivilegeSetRelationship{), Privileges: []IamPrivilegeRelationship{), ResourceLimits: openapiclient.iam.ResourceLimits.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, PerAccountUserLimit: int64(123), Account: }, Roles: []IamRoleRelationship{), SecurityHolder: openapiclient.iam.SecurityHolder.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Account: , IpRulesConfiguration: openapiclient.iam.IpAccessManagement.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Enable: false, LastRecoveryTime: "TODO", Holder: openapiclient.iam.SecurityHolder.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Account: , IpRulesConfiguration: openapiclient.iam.IpAccessManagement.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Enable: false, LastRecoveryTime: "TODO", Holder: , IpAddresses: []IamIpAddressRelationship{openapiclient.iam.IpAddress.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Address: "Address_example", Description: "Description_example", IpAccessManagement: })}, ResourcePermissions: []IamResourcePermissionRelationship{openapiclient.iam.ResourcePermission.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, PermissionRoles: []IamPermissionToRoles{openapiclient.iam.PermissionToRoles{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Permission: openapiclient.cmrf.CmRf{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example"}, Roles: []CmrfCmRf{openapiclient.cmrf.CmRf{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example"})}), TargetApp: "TargetApp_example", Holder: , Resource: })}, IpAddresses: []IamIpAddressRelationship{openapiclient.iam.IpAddress.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Address: "Address_example", Description: "Description_example", IpAccessManagement: })}, ResourcePermissions: []IamResourcePermissionRelationship{openapiclient.iam.ResourcePermission.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, PermissionRoles: []IamPermissionToRoles{openapiclient.iam.PermissionToRoles{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Permission: , Roles: []CmrfCmRf{)}), TargetApp: "TargetApp_example", Holder: , Resource: })}, SessionLimits: }}, AppRegistrations: []IamAppRegistrationRelationship{), DomainGroups: []IamDomainGroupRelationship{openapiclient.iam.DomainGroup.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", Partition1: int64(123), Partition2: int64(123), Partition3: int64(123), PartitionKey: "PartitionKey_example", Usage: int64(123), Account: }), EndPointRoles: []IamEndPointRoleRelationship{), Idpreferences: []IamIdpReferenceRelationship{), Idps: []IamIdpRelationship{), Permissions: []IamPermissionRelationship{), PrivilegeSets: []IamPrivilegeSetRelationship{), Privileges: []IamPrivilegeRelationship{), ResourceLimits: openapiclient.iam.ResourceLimits.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, PerAccountUserLimit: int64(123), Account: }, Roles: []IamRoleRelationship{), SecurityHolder: , SessionLimits: }, ResourceGroups: []ResourceGroupRelationship{openapiclient.resource.Group.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Name: "Name_example", PerTypeCombinedSelector: []ResourcePerTypeCombinedSelector{openapiclient.resource.PerTypeCombinedSelector{ClassId: "ClassId_example", ObjectType: "ObjectType_example", CombinedSelector: "CombinedSelector_example", EmptyFilter: false, SelectorObjectType: "SelectorObjectType_example"}), Qualifier: "Qualifier_example", Selectors: []ResourceSelector{openapiclient.resource.Selector{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Selector: "Selector_example"}), Account: , Organizations: []OrganizationOrganizationRelationship{)})}, ShadowPools: []IppoolShadowPoolRelationship{openapiclient.ippool.ShadowPool.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Assigned: int64(123), AssignmentOrder: "AssignmentOrder_example", Size: int64(123), IpV4Blocks: []IppoolIpBlock{openapiclient.ippool.IpBlock{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Size: int64(123), From: "From_example", To: "To_example"}), IpV4Config: , V4Assigned: int64(123), V4Size: int64(123), V6Assigned: int64(123), V6Size: int64(123), Pool: openapiclient.ippool.Pool.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Assigned: int64(123), AssignmentOrder: "AssignmentOrder_example", Size: int64(123), IpV4Blocks: []IppoolIpBlock{), IpV4Config: , V4Assigned: int64(123), V4Size: int64(123), V6Assigned: int64(123), V6Size: int64(123), Organization: , ShadowPools: []IppoolShadowPoolRelationship{openapiclient.ippool.ShadowPool.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Assigned: int64(123), AssignmentOrder: "AssignmentOrder_example", Size: int64(123), IpV4Blocks: []IppoolIpBlock{), IpV4Config: , V4Assigned: int64(123), V4Size: int64(123), V6Assigned: int64(123), V6Size: int64(123), Pool: , V4BlockHeads: []IppoolShadowBlockRelationship{openapiclient.ippool.ShadowBlock.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, FreeBlockCount: int64(123), NextIdAllocator: int64(123), IpV4Block: , Pool: }), Vrf: openapiclient.vrf.Vrf.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Account: }})}, V4BlockHeads: []IppoolShadowBlockRelationship{openapiclient.ippool.ShadowBlock.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, FreeBlockCount: int64(123), NextIdAllocator: int64(123), IpV4Block: , Pool: }), Vrf: openapiclient.vrf.Vrf.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Account: }})}, PoolMember: openapiclient.ippool.PoolMember.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Assigned: false, IpV4Address: "IpV4Address_example", AssignedToEntity: , IpV4Block: , Peer: openapiclient.ippool.IpLease.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, IpV4Address: "IpV4Address_example", IpV4Config: , AssignedToEntity: , Pool: , PoolMember: openapiclient.ippool.PoolMember.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Assigned: false, IpV4Address: "IpV4Address_example", AssignedToEntity: , IpV4Block: , Peer: openapiclient.ippool.IpLease.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, IpV4Address: "IpV4Address_example", IpV4Config: , AssignedToEntity: , Pool: , PoolMember: , Universe: openapiclient.ippool.Universe.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Vrf: }, Vrf: }, Pool: }, Universe: openapiclient.ippool.Universe.Relationship{ClassId: "ClassId_example", ObjectType: "ObjectType_example", Moid: "Moid_example", Selector: "Selector_example", Link: "Link_example", AccountMoid: "AccountMoid_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Vrf: }, Vrf: }, Pool: }, Universe: , Vrf: } // IppoolIpLease | The 'ippool.IpLease' resource to create.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn't happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource's ETag doesn't match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don't have to be identical byte for byte. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IppoolApi.CreateIppoolIpLease(context.Background(), ippoolIpLease).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IppoolApi.CreateIppoolIpLease``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIppoolIpLease`: IppoolIpLease
    fmt.Fprintf(os.Stdout, "Response from `IppoolApi.CreateIppoolIpLease`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIppoolIpLeaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ippoolIpLease** | [**IppoolIpLease**](IppoolIpLease.md) | The &#39;ippool.IpLease&#39; resource to create. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 
 **ifNoneMatch** | **string** | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn&#39;t happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource&#39;s ETag doesn&#39;t match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don&#39;t have to be identical byte for byte. | 

### Return type

[**IppoolIpLease**](ippool.IpLease.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIppoolPool

> IppoolPool CreateIppoolPool(ctx).IppoolPool(ippoolPool).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()

Create a 'ippool.Pool' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ippoolPool := openapiclient.ippool.Pool{AccountMoid: "AccountMoid_example", ClassId: "ClassId_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", ObjectType: "ObjectType_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Assigned: int64(123), AssignmentOrder: "AssignmentOrder_example", Size: int64(123), IpV4Blocks: []IppoolIpBlock{), IpV4Config: , V4Assigned: int64(123), V4Size: int64(123), V6Assigned: int64(123), V6Size: int64(123), Organization: , ShadowPools: []IppoolShadowPoolRelationship{)} // IppoolPool | The 'ippool.Pool' resource to create.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn't happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource's ETag doesn't match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don't have to be identical byte for byte. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IppoolApi.CreateIppoolPool(context.Background(), ippoolPool).IfMatch(ifMatch).IfNoneMatch(ifNoneMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IppoolApi.CreateIppoolPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIppoolPool`: IppoolPool
    fmt.Fprintf(os.Stdout, "Response from `IppoolApi.CreateIppoolPool`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIppoolPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ippoolPool** | [**IppoolPool**](IppoolPool.md) | The &#39;ippool.Pool&#39; resource to create. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 
 **ifNoneMatch** | **string** | For methods that apply server-side changes, If-None-Match used with the * value can be used to create a resource not known to exist, guaranteeing that another resource creation didn&#39;t happen before, losing the data of the previous put. The request will be processed only if the eventually existing resource&#39;s ETag doesn&#39;t match any of the values listed. Otherwise, the status code 412 (Precondition Failed) is used. The asterisk is a special value representing any resource. It is only useful when creating a resource, usually with PUT, to check if another resource with the identity has already been created before. The comparison with the stored ETag uses the weak comparison algorithm, meaning two resources are considered identical if the content is equivalent - they don&#39;t have to be identical byte for byte. | 

### Return type

[**IppoolPool**](ippool.Pool.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIppoolIpLease

> DeleteIppoolIpLease(ctx, moid).Execute()

Delete a 'ippool.IpLease' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IppoolApi.DeleteIppoolIpLease(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IppoolApi.DeleteIppoolIpLease``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIppoolIpLeaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIppoolPool

> DeleteIppoolPool(ctx, moid).Execute()

Delete a 'ippool.Pool' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IppoolApi.DeleteIppoolPool(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IppoolApi.DeleteIppoolPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIppoolPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIppoolIpLeaseByMoid

> IppoolIpLease GetIppoolIpLeaseByMoid(ctx, moid).Execute()

Read a 'ippool.IpLease' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IppoolApi.GetIppoolIpLeaseByMoid(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IppoolApi.GetIppoolIpLeaseByMoid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIppoolIpLeaseByMoid`: IppoolIpLease
    fmt.Fprintf(os.Stdout, "Response from `IppoolApi.GetIppoolIpLeaseByMoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIppoolIpLeaseByMoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IppoolIpLease**](ippool.IpLease.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIppoolIpLeaseList

> IppoolIpLeaseResponse GetIppoolIpLeaseList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()

Read a 'ippool.IpLease' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    filter := "filter_example" // string | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). (optional) (default to "")
    orderby := "orderby_example" // string | Determines what properties are used to sort the collection of resources. (optional)
    top := 987 // int32 | Specifies the maximum number of resources to return in the response. (optional) (default to 100)
    skip := 987 // int32 | Specifies the number of resources to skip in the response. (optional) (default to 0)
    select_ := "select__example" // string | Specifies a subset of properties to return. (optional) (default to "")
    expand := "expand_example" // string | Specify additional attributes or related resources to return in addition to the primary resources. (optional)
    apply := "apply_example" // string | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \"$apply\" query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \"aggregate\" and \"groupby\". The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. (optional)
    count := false // bool | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. (optional)
    inlinecount := "inlinecount_example" // string | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. (optional) (default to "allpages")
    at := "at_example" // string | Similar to \"$filter\", but \"at\" is specifically used to filter versioning information properties for resources to return. A URI with an \"at\" Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. (optional)
    tags := "tags_example" // string | The 'tags' parameter is used to request a summary of the Tag utilization for this resource. When the 'tags' parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IppoolApi.GetIppoolIpLeaseList(context.Background(), ).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IppoolApi.GetIppoolIpLeaseList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIppoolIpLeaseList`: IppoolIpLeaseResponse
    fmt.Fprintf(os.Stdout, "Response from `IppoolApi.GetIppoolIpLeaseList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIppoolIpLeaseListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). | [default to &quot;&quot;]
 **orderby** | **string** | Determines what properties are used to sort the collection of resources. | 
 **top** | **int32** | Specifies the maximum number of resources to return in the response. | [default to 100]
 **skip** | **int32** | Specifies the number of resources to skip in the response. | [default to 0]
 **select_** | **string** | Specifies a subset of properties to return. | [default to &quot;&quot;]
 **expand** | **string** | Specify additional attributes or related resources to return in addition to the primary resources. | 
 **apply** | **string** | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \&quot;$apply\&quot; query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \&quot;aggregate\&quot; and \&quot;groupby\&quot;. The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. | 
 **count** | **bool** | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. | 
 **inlinecount** | **string** | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. | [default to &quot;allpages&quot;]
 **at** | **string** | Similar to \&quot;$filter\&quot;, but \&quot;at\&quot; is specifically used to filter versioning information properties for resources to return. A URI with an \&quot;at\&quot; Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. | 
 **tags** | **string** | The &#39;tags&#39; parameter is used to request a summary of the Tag utilization for this resource. When the &#39;tags&#39; parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. | 

### Return type

[**IppoolIpLeaseResponse**](ippool.IpLease.Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIppoolPoolByMoid

> IppoolPool GetIppoolPoolByMoid(ctx, moid).Execute()

Read a 'ippool.Pool' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IppoolApi.GetIppoolPoolByMoid(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IppoolApi.GetIppoolPoolByMoid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIppoolPoolByMoid`: IppoolPool
    fmt.Fprintf(os.Stdout, "Response from `IppoolApi.GetIppoolPoolByMoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIppoolPoolByMoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IppoolPool**](ippool.Pool.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIppoolPoolList

> IppoolPoolResponse GetIppoolPoolList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()

Read a 'ippool.Pool' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    filter := "filter_example" // string | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). (optional) (default to "")
    orderby := "orderby_example" // string | Determines what properties are used to sort the collection of resources. (optional)
    top := 987 // int32 | Specifies the maximum number of resources to return in the response. (optional) (default to 100)
    skip := 987 // int32 | Specifies the number of resources to skip in the response. (optional) (default to 0)
    select_ := "select__example" // string | Specifies a subset of properties to return. (optional) (default to "")
    expand := "expand_example" // string | Specify additional attributes or related resources to return in addition to the primary resources. (optional)
    apply := "apply_example" // string | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \"$apply\" query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \"aggregate\" and \"groupby\". The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. (optional)
    count := false // bool | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. (optional)
    inlinecount := "inlinecount_example" // string | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. (optional) (default to "allpages")
    at := "at_example" // string | Similar to \"$filter\", but \"at\" is specifically used to filter versioning information properties for resources to return. A URI with an \"at\" Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. (optional)
    tags := "tags_example" // string | The 'tags' parameter is used to request a summary of the Tag utilization for this resource. When the 'tags' parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IppoolApi.GetIppoolPoolList(context.Background(), ).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IppoolApi.GetIppoolPoolList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIppoolPoolList`: IppoolPoolResponse
    fmt.Fprintf(os.Stdout, "Response from `IppoolApi.GetIppoolPoolList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIppoolPoolListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). | [default to &quot;&quot;]
 **orderby** | **string** | Determines what properties are used to sort the collection of resources. | 
 **top** | **int32** | Specifies the maximum number of resources to return in the response. | [default to 100]
 **skip** | **int32** | Specifies the number of resources to skip in the response. | [default to 0]
 **select_** | **string** | Specifies a subset of properties to return. | [default to &quot;&quot;]
 **expand** | **string** | Specify additional attributes or related resources to return in addition to the primary resources. | 
 **apply** | **string** | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \&quot;$apply\&quot; query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \&quot;aggregate\&quot; and \&quot;groupby\&quot;. The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. | 
 **count** | **bool** | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. | 
 **inlinecount** | **string** | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. | [default to &quot;allpages&quot;]
 **at** | **string** | Similar to \&quot;$filter\&quot;, but \&quot;at\&quot; is specifically used to filter versioning information properties for resources to return. A URI with an \&quot;at\&quot; Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. | 
 **tags** | **string** | The &#39;tags&#39; parameter is used to request a summary of the Tag utilization for this resource. When the &#39;tags&#39; parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. | 

### Return type

[**IppoolPoolResponse**](ippool.Pool.Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIppoolPoolMemberByMoid

> IppoolPoolMember GetIppoolPoolMemberByMoid(ctx, moid).Execute()

Read a 'ippool.PoolMember' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IppoolApi.GetIppoolPoolMemberByMoid(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IppoolApi.GetIppoolPoolMemberByMoid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIppoolPoolMemberByMoid`: IppoolPoolMember
    fmt.Fprintf(os.Stdout, "Response from `IppoolApi.GetIppoolPoolMemberByMoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIppoolPoolMemberByMoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IppoolPoolMember**](ippool.PoolMember.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIppoolPoolMemberList

> IppoolPoolMemberResponse GetIppoolPoolMemberList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()

Read a 'ippool.PoolMember' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    filter := "filter_example" // string | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). (optional) (default to "")
    orderby := "orderby_example" // string | Determines what properties are used to sort the collection of resources. (optional)
    top := 987 // int32 | Specifies the maximum number of resources to return in the response. (optional) (default to 100)
    skip := 987 // int32 | Specifies the number of resources to skip in the response. (optional) (default to 0)
    select_ := "select__example" // string | Specifies a subset of properties to return. (optional) (default to "")
    expand := "expand_example" // string | Specify additional attributes or related resources to return in addition to the primary resources. (optional)
    apply := "apply_example" // string | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \"$apply\" query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \"aggregate\" and \"groupby\". The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. (optional)
    count := false // bool | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. (optional)
    inlinecount := "inlinecount_example" // string | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. (optional) (default to "allpages")
    at := "at_example" // string | Similar to \"$filter\", but \"at\" is specifically used to filter versioning information properties for resources to return. A URI with an \"at\" Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. (optional)
    tags := "tags_example" // string | The 'tags' parameter is used to request a summary of the Tag utilization for this resource. When the 'tags' parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IppoolApi.GetIppoolPoolMemberList(context.Background(), ).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IppoolApi.GetIppoolPoolMemberList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIppoolPoolMemberList`: IppoolPoolMemberResponse
    fmt.Fprintf(os.Stdout, "Response from `IppoolApi.GetIppoolPoolMemberList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIppoolPoolMemberListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). | [default to &quot;&quot;]
 **orderby** | **string** | Determines what properties are used to sort the collection of resources. | 
 **top** | **int32** | Specifies the maximum number of resources to return in the response. | [default to 100]
 **skip** | **int32** | Specifies the number of resources to skip in the response. | [default to 0]
 **select_** | **string** | Specifies a subset of properties to return. | [default to &quot;&quot;]
 **expand** | **string** | Specify additional attributes or related resources to return in addition to the primary resources. | 
 **apply** | **string** | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \&quot;$apply\&quot; query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \&quot;aggregate\&quot; and \&quot;groupby\&quot;. The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. | 
 **count** | **bool** | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. | 
 **inlinecount** | **string** | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. | [default to &quot;allpages&quot;]
 **at** | **string** | Similar to \&quot;$filter\&quot;, but \&quot;at\&quot; is specifically used to filter versioning information properties for resources to return. A URI with an \&quot;at\&quot; Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. | 
 **tags** | **string** | The &#39;tags&#39; parameter is used to request a summary of the Tag utilization for this resource. When the &#39;tags&#39; parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. | 

### Return type

[**IppoolPoolMemberResponse**](ippool.PoolMember.Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIppoolShadowBlockByMoid

> IppoolShadowBlock GetIppoolShadowBlockByMoid(ctx, moid).Execute()

Read a 'ippool.ShadowBlock' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IppoolApi.GetIppoolShadowBlockByMoid(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IppoolApi.GetIppoolShadowBlockByMoid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIppoolShadowBlockByMoid`: IppoolShadowBlock
    fmt.Fprintf(os.Stdout, "Response from `IppoolApi.GetIppoolShadowBlockByMoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIppoolShadowBlockByMoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IppoolShadowBlock**](ippool.ShadowBlock.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIppoolShadowBlockList

> IppoolShadowBlockResponse GetIppoolShadowBlockList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()

Read a 'ippool.ShadowBlock' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    filter := "filter_example" // string | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). (optional) (default to "")
    orderby := "orderby_example" // string | Determines what properties are used to sort the collection of resources. (optional)
    top := 987 // int32 | Specifies the maximum number of resources to return in the response. (optional) (default to 100)
    skip := 987 // int32 | Specifies the number of resources to skip in the response. (optional) (default to 0)
    select_ := "select__example" // string | Specifies a subset of properties to return. (optional) (default to "")
    expand := "expand_example" // string | Specify additional attributes or related resources to return in addition to the primary resources. (optional)
    apply := "apply_example" // string | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \"$apply\" query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \"aggregate\" and \"groupby\". The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. (optional)
    count := false // bool | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. (optional)
    inlinecount := "inlinecount_example" // string | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. (optional) (default to "allpages")
    at := "at_example" // string | Similar to \"$filter\", but \"at\" is specifically used to filter versioning information properties for resources to return. A URI with an \"at\" Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. (optional)
    tags := "tags_example" // string | The 'tags' parameter is used to request a summary of the Tag utilization for this resource. When the 'tags' parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IppoolApi.GetIppoolShadowBlockList(context.Background(), ).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IppoolApi.GetIppoolShadowBlockList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIppoolShadowBlockList`: IppoolShadowBlockResponse
    fmt.Fprintf(os.Stdout, "Response from `IppoolApi.GetIppoolShadowBlockList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIppoolShadowBlockListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). | [default to &quot;&quot;]
 **orderby** | **string** | Determines what properties are used to sort the collection of resources. | 
 **top** | **int32** | Specifies the maximum number of resources to return in the response. | [default to 100]
 **skip** | **int32** | Specifies the number of resources to skip in the response. | [default to 0]
 **select_** | **string** | Specifies a subset of properties to return. | [default to &quot;&quot;]
 **expand** | **string** | Specify additional attributes or related resources to return in addition to the primary resources. | 
 **apply** | **string** | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \&quot;$apply\&quot; query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \&quot;aggregate\&quot; and \&quot;groupby\&quot;. The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. | 
 **count** | **bool** | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. | 
 **inlinecount** | **string** | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. | [default to &quot;allpages&quot;]
 **at** | **string** | Similar to \&quot;$filter\&quot;, but \&quot;at\&quot; is specifically used to filter versioning information properties for resources to return. A URI with an \&quot;at\&quot; Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. | 
 **tags** | **string** | The &#39;tags&#39; parameter is used to request a summary of the Tag utilization for this resource. When the &#39;tags&#39; parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. | 

### Return type

[**IppoolShadowBlockResponse**](ippool.ShadowBlock.Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIppoolShadowPoolByMoid

> IppoolShadowPool GetIppoolShadowPoolByMoid(ctx, moid).Execute()

Read a 'ippool.ShadowPool' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IppoolApi.GetIppoolShadowPoolByMoid(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IppoolApi.GetIppoolShadowPoolByMoid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIppoolShadowPoolByMoid`: IppoolShadowPool
    fmt.Fprintf(os.Stdout, "Response from `IppoolApi.GetIppoolShadowPoolByMoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIppoolShadowPoolByMoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IppoolShadowPool**](ippool.ShadowPool.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIppoolShadowPoolList

> IppoolShadowPoolResponse GetIppoolShadowPoolList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()

Read a 'ippool.ShadowPool' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    filter := "filter_example" // string | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). (optional) (default to "")
    orderby := "orderby_example" // string | Determines what properties are used to sort the collection of resources. (optional)
    top := 987 // int32 | Specifies the maximum number of resources to return in the response. (optional) (default to 100)
    skip := 987 // int32 | Specifies the number of resources to skip in the response. (optional) (default to 0)
    select_ := "select__example" // string | Specifies a subset of properties to return. (optional) (default to "")
    expand := "expand_example" // string | Specify additional attributes or related resources to return in addition to the primary resources. (optional)
    apply := "apply_example" // string | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \"$apply\" query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \"aggregate\" and \"groupby\". The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. (optional)
    count := false // bool | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. (optional)
    inlinecount := "inlinecount_example" // string | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. (optional) (default to "allpages")
    at := "at_example" // string | Similar to \"$filter\", but \"at\" is specifically used to filter versioning information properties for resources to return. A URI with an \"at\" Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. (optional)
    tags := "tags_example" // string | The 'tags' parameter is used to request a summary of the Tag utilization for this resource. When the 'tags' parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IppoolApi.GetIppoolShadowPoolList(context.Background(), ).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IppoolApi.GetIppoolShadowPoolList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIppoolShadowPoolList`: IppoolShadowPoolResponse
    fmt.Fprintf(os.Stdout, "Response from `IppoolApi.GetIppoolShadowPoolList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIppoolShadowPoolListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). | [default to &quot;&quot;]
 **orderby** | **string** | Determines what properties are used to sort the collection of resources. | 
 **top** | **int32** | Specifies the maximum number of resources to return in the response. | [default to 100]
 **skip** | **int32** | Specifies the number of resources to skip in the response. | [default to 0]
 **select_** | **string** | Specifies a subset of properties to return. | [default to &quot;&quot;]
 **expand** | **string** | Specify additional attributes or related resources to return in addition to the primary resources. | 
 **apply** | **string** | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \&quot;$apply\&quot; query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \&quot;aggregate\&quot; and \&quot;groupby\&quot;. The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. | 
 **count** | **bool** | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. | 
 **inlinecount** | **string** | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. | [default to &quot;allpages&quot;]
 **at** | **string** | Similar to \&quot;$filter\&quot;, but \&quot;at\&quot; is specifically used to filter versioning information properties for resources to return. A URI with an \&quot;at\&quot; Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. | 
 **tags** | **string** | The &#39;tags&#39; parameter is used to request a summary of the Tag utilization for this resource. When the &#39;tags&#39; parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. | 

### Return type

[**IppoolShadowPoolResponse**](ippool.ShadowPool.Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIppoolUniverseByMoid

> IppoolUniverse GetIppoolUniverseByMoid(ctx, moid).Execute()

Read a 'ippool.Universe' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IppoolApi.GetIppoolUniverseByMoid(context.Background(), moid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IppoolApi.GetIppoolUniverseByMoid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIppoolUniverseByMoid`: IppoolUniverse
    fmt.Fprintf(os.Stdout, "Response from `IppoolApi.GetIppoolUniverseByMoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIppoolUniverseByMoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IppoolUniverse**](ippool.Universe.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIppoolUniverseList

> IppoolUniverseResponse GetIppoolUniverseList(ctx).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()

Read a 'ippool.Universe' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    filter := "filter_example" // string | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). (optional) (default to "")
    orderby := "orderby_example" // string | Determines what properties are used to sort the collection of resources. (optional)
    top := 987 // int32 | Specifies the maximum number of resources to return in the response. (optional) (default to 100)
    skip := 987 // int32 | Specifies the number of resources to skip in the response. (optional) (default to 0)
    select_ := "select__example" // string | Specifies a subset of properties to return. (optional) (default to "")
    expand := "expand_example" // string | Specify additional attributes or related resources to return in addition to the primary resources. (optional)
    apply := "apply_example" // string | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \"$apply\" query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \"aggregate\" and \"groupby\". The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. (optional)
    count := false // bool | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. (optional)
    inlinecount := "inlinecount_example" // string | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. (optional) (default to "allpages")
    at := "at_example" // string | Similar to \"$filter\", but \"at\" is specifically used to filter versioning information properties for resources to return. A URI with an \"at\" Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. (optional)
    tags := "tags_example" // string | The 'tags' parameter is used to request a summary of the Tag utilization for this resource. When the 'tags' parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IppoolApi.GetIppoolUniverseList(context.Background(), ).Filter(filter).Orderby(orderby).Top(top).Skip(skip).Select_(select_).Expand(expand).Apply(apply).Count(count).Inlinecount(inlinecount).At(at).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IppoolApi.GetIppoolUniverseList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIppoolUniverseList`: IppoolUniverseResponse
    fmt.Fprintf(os.Stdout, "Response from `IppoolApi.GetIppoolUniverseList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIppoolUniverseListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filter criteria for the resources to return. A URI with a $filter query option identifies a subset of the entries from the Collection of Entries. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the $filter option. The expression language that is used in $filter queries supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false). | [default to &quot;&quot;]
 **orderby** | **string** | Determines what properties are used to sort the collection of resources. | 
 **top** | **int32** | Specifies the maximum number of resources to return in the response. | [default to 100]
 **skip** | **int32** | Specifies the number of resources to skip in the response. | [default to 0]
 **select_** | **string** | Specifies a subset of properties to return. | [default to &quot;&quot;]
 **expand** | **string** | Specify additional attributes or related resources to return in addition to the primary resources. | 
 **apply** | **string** | Specify one or more transformation operations to perform aggregation on the resources. The transformations are processed in order with the output from a transformation being used as input for the subsequent transformation. The \&quot;$apply\&quot; query takes a sequence of set transformations, separated by forward slashes to express that they are consecutively applied, i.e. the result of each transformation is the input to the next transformation. Supported aggregation methods are \&quot;aggregate\&quot; and \&quot;groupby\&quot;. The **aggregate** transformation takes a comma-separated list of one or more aggregate expressions as parameters and returns a result set with a single instance, representing the aggregated value for all instances in the input set. The **groupby** transformation takes one or two parameters and 1. Splits the initial set into subsets where all instances in a subset have the same values for the grouping properties specified in the first parameter, 2. Applies set transformations to each subset according to the second parameter, resulting in a new set of potentially different structure and cardinality, 3. Ensures that the instances in the result set contain all grouping properties with the correct values for the group, 4. Concatenates the intermediate result sets into one result set. A groupby transformation affects the structure of the result set. | 
 **count** | **bool** | The $count query specifies the service should return the count of the matching resources, instead of returning the resources. | 
 **inlinecount** | **string** | The $inlinecount query option allows clients to request an inline count of the matching resources included with the resources in the response. | [default to &quot;allpages&quot;]
 **at** | **string** | Similar to \&quot;$filter\&quot;, but \&quot;at\&quot; is specifically used to filter versioning information properties for resources to return. A URI with an \&quot;at\&quot; Query Option identifies a subset of the Entries from the Collection of Entries identified by the Resource Path section of the URI. The subset is determined by selecting only the Entries that satisfy the predicate expression specified by the query option. The expression language that is used in at operators supports references to properties and literals. The literal values can be strings enclosed in single quotes, numbers and boolean values (true or false) or any of the additional literal representations shown in the Abstract Type System section. | 
 **tags** | **string** | The &#39;tags&#39; parameter is used to request a summary of the Tag utilization for this resource. When the &#39;tags&#39; parameter is specified, the response provides a list of tag keys, the number of times the key has been used across all documents, and the tag values that have been assigned to the tag key. | 

### Return type

[**IppoolUniverseResponse**](ippool.Universe.Response.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchIppoolPool

> IppoolPool PatchIppoolPool(ctx, moid).IppoolPool(ippoolPool).IfMatch(ifMatch).Execute()

Update a 'ippool.Pool' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.
    ippoolPool := openapiclient.ippool.Pool{AccountMoid: "AccountMoid_example", ClassId: "ClassId_example", CreateTime: "TODO", DomainGroupMoid: "DomainGroupMoid_example", ModTime: "TODO", Moid: "Moid_example", ObjectType: "ObjectType_example", Owners: []string{"Owners_example"), SharedScope: "SharedScope_example", Tags: []MoTag{), VersionContext: , Ancestors: []MoBaseMoRelationship{), Parent: , PermissionResources: []MoBaseMoRelationship{), DisplayNames: map[string]string{ "Key" = "Value" }, Description: "Description_example", Name: "Name_example", Assigned: int64(123), AssignmentOrder: "AssignmentOrder_example", Size: int64(123), IpV4Blocks: []IppoolIpBlock{), IpV4Config: , V4Assigned: int64(123), V4Size: int64(123), V6Assigned: int64(123), V6Size: int64(123), Organization: , ShadowPools: []IppoolShadowPoolRelationship{)} // IppoolPool | The 'ippool.Pool' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IppoolApi.PatchIppoolPool(context.Background(), moid, ippoolPool).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IppoolApi.PatchIppoolPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchIppoolPool`: IppoolPool
    fmt.Fprintf(os.Stdout, "Response from `IppoolApi.PatchIppoolPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchIppoolPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ippoolPool** | [**IppoolPool**](IppoolPool.md) | The &#39;ippool.Pool&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**IppoolPool**](ippool.Pool.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIppoolPool

> IppoolPool UpdateIppoolPool(ctx, moid).IppoolPool(ippoolPool).IfMatch(ifMatch).Execute()

Update a 'ippool.Pool' resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moid := "moid_example" // string | The unique Moid identifier of a resource instance.
    ippoolPool :=  // IppoolPool | The 'ippool.Pool' resource to update.
    ifMatch := "ifMatch_example" // string | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IppoolApi.UpdateIppoolPool(context.Background(), moid, ippoolPool).IfMatch(ifMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IppoolApi.UpdateIppoolPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIppoolPool`: IppoolPool
    fmt.Fprintf(os.Stdout, "Response from `IppoolApi.UpdateIppoolPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moid** | **string** | The unique Moid identifier of a resource instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIppoolPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ippoolPool** | [**IppoolPool**](IppoolPool.md) | The &#39;ippool.Pool&#39; resource to update. | 
 **ifMatch** | **string** | For methods that apply server-side changes, and in particular for PUT, If-Match can be used to prevent the lost update problem. It can check if the modification of a resource that the user wants to upload will not override another change that has been done since the original resource was fetched. If the request cannot be fulfilled, the 412 (Precondition Failed) response is returned. When modifying a resource using POST or PUT, the If-Match header must be set to the value of the resource ModTime property after which no lost update problem should occur. For example, a client send a GET request to obtain a resource, which includes the ModTime property. The ModTime indicates the last time the resource was created or modified. The client then sends a POST or PUT request with the If-Match header set to the ModTime property of the resource as obtained in the GET request. | 

### Return type

[**IppoolPool**](ippool.Pool.md)

### Authorization

[cookieAuth](../README.md#cookieAuth), [http_signature](../README.md#http_signature), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json, application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

