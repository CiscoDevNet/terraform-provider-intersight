package intersight

import (
	"context"
	"log"
	"reflect"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceHyperflexDatastoreStatistic() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHyperflexDatastoreStatisticRead,
		Schema: map[string]*schema.Schema{
			"accessibility_summary": {
				Description: "HyperFlex datastore accessibility summary.\n* `ACCESSIBLE` - The HyperFlex Accessibility Summary is Accessible.\n* `NOT_ACCESSIBLE` - The HyperFlex Accessibility Summary is Not Accessible.\n* `PARTIALLY_ACCESSIBLE` - The HyperFlex Accessibility Summary is Partially Accessible.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"class_id": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"creation_time": {
				Description: "Timestamp the datastore object was created.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"datastore_status": {
				Description: "HyperFlex datastore status.\n* `NORMAL` - The HyperFlex datastore status is normal.\n* `ALERT` - The HyperFlex datastore status is alert.\n* `FAILED` - The HyperFlex datastore status is failed.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"free_capacity_in_bytes": {
				Description: "Free capacity of the datastore in bytes.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"last_modified_time": {
				Description: "Timestamp the datastore object was last modified.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"mount_summary": {
				Description: "HyperFlex datastore mount summary.\n* `MOUNTED` - The HyperFlex mount summary is mounted.\n* `UNMOUNTED` - The HyperFlex mount summary is unmounted.\n* `MOUNT_FAILURE` - The HyperFlex mount summary is mount failure.\n* `UNMOUNT_FAILURE` - The HyperFlex mount summary is unmount failure.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"object_type": {
				Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"parent_uuid": {
				Description: "UUID of the parent datastore object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"total_capacity_in_bytes": {
				Description: "Total capacity of the datastore object.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"un_compressed_used_bytes": {
				Description: "Number of uncompressed used bytes in the datastore.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"unshared_used_bytes": {
				Description: "Unshared used capacity of the datastore in bytes.",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"uuid": {
				Description: "UUID for the datastore object.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"results": {
				Type: schema.TypeList,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{"accessibility_summary": {
					Description: "HyperFlex datastore accessibility summary.\n* `ACCESSIBLE` - The HyperFlex Accessibility Summary is Accessible.\n* `NOT_ACCESSIBLE` - The HyperFlex Accessibility Summary is Not Accessible.\n* `PARTIALLY_ACCESSIBLE` - The HyperFlex Accessibility Summary is Partially Accessible.",
					Type:        schema.TypeString,
					Optional:    true,
					Computed:    true,
				},
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"creation_time": {
						Description: "Timestamp the datastore object was created.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"data_protection_peer": {
						Description: "A reference to a hyperflexDataProtectionPeer resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Computed:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
					},
					"datastore_status": {
						Description: "HyperFlex datastore status.\n* `NORMAL` - The HyperFlex datastore status is normal.\n* `ALERT` - The HyperFlex datastore status is alert.\n* `FAILED` - The HyperFlex datastore status is failed.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"dsconfig": {
						Description: "HyperFlex platform datastore configuration.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Computed:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"data_block_size": {
									Description: "Size of datablock for this HyperFlex datastore.",
									Type:        schema.TypeInt,
									Optional:    true,
									Computed:    true,
								},
								"name": {
									Description: "Unique name for the datastore.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"num_mirrors": {
									Description: "Number of copies of data maintained for data redundancy.",
									Type:        schema.TypeInt,
									Optional:    true,
									Computed:    true,
								},
								"num_stripes_for_large_files": {
									Description: "Number of stripes to be used for large files in datastore.",
									Type:        schema.TypeInt,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"provisioned_capacity": {
									Description: "Provisioned capacity of datastore in bytes.",
									Type:        schema.TypeInt,
									Optional:    true,
									Computed:    true,
								},
								"system_datastore": {
									Description: "Specifies if this datastore is a system datastore or not.",
									Type:        schema.TypeBool,
									Optional:    true,
									Computed:    true,
								},
							},
						},
					},
					"free_capacity_in_bytes": {
						Description: "Free capacity of the datastore in bytes.",
						Type:        schema.TypeInt,
						Optional:    true,
						Computed:    true,
					},
					"host_mount_status": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"accessibility": {
									Description: "Accessibility of HyperFlex datastore.\n* `ACCESSIBLE` - The HyperFlex datastore accessibility state is accessible.\n* `NOT_ACCESSIBLE` - The HyperFlex datastore accessibility state is not accessible.\n* `PARTIALLY_ACCESSIBLE` - The HyperFlex datastore accessibility state is partially accessible.\n* `READONLY` - The HyperFlex datastore accessibility state is read-only.\n* `HOST_NOT_REACHABLE` - The HyperFlex datastore accessibility state is host not reachable.\n* `UNKNOWN` - The HyperFlex datastore accessibility state is unknown.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"host_name": {
									Description: "HyperFlex name of host for this datastore.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"mounted": {
									Description: "Is the HyperFlex datastore mounted or not.",
									Type:        schema.TypeBool,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"reason": {
									Description: "Reason for inaccessibility for this datastore.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
							},
						},
						Computed: true,
					},
					"last_modified_time": {
						Description: "Timestamp the datastore object was last modified.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"moid": {
						Description: "The unique identifier of this Managed Object instance.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"mount_summary": {
						Description: "HyperFlex datastore mount summary.\n* `MOUNTED` - The HyperFlex mount summary is mounted.\n* `UNMOUNTED` - The HyperFlex mount summary is unmounted.\n* `MOUNT_FAILURE` - The HyperFlex mount summary is mount failure.\n* `UNMOUNT_FAILURE` - The HyperFlex mount summary is unmount failure.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"parent_uuid": {
						Description: "UUID of the parent datastore object.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
					"site": {
						Description: "HyperFlex site for the datastore.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Computed:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"name": {
									Description: "Name of the site for this HyperFlex cluster.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
									Computed:    true,
								},
								"zone": {
									Description: "Zone information for this HyperFlex cluster site.",
									Type:        schema.TypeList,
									MaxItems:    1,
									Optional:    true,
									Computed:    true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"additional_properties": {
												Type:             schema.TypeString,
												Optional:         true,
												DiffSuppressFunc: SuppressDiffAdditionProps,
											},
											"class_id": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"num_nodes": {
												Description: "Number of nodes in the Zone.",
												Type:        schema.TypeInt,
												Optional:    true,
												Computed:    true,
											},
											"object_type": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
											"uuid": {
												Description: "Zone UUID for this HyperFlex node.",
												Type:        schema.TypeString,
												Optional:    true,
												Computed:    true,
											},
										},
									},
								},
							},
						},
					},
					"tags": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"key": {
									Description: "The string representation of a tag key.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"value": {
									Description: "The string representation of a tag value.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
					"total_capacity_in_bytes": {
						Description: "Total capacity of the datastore object.",
						Type:        schema.TypeInt,
						Optional:    true,
						Computed:    true,
					},
					"un_compressed_used_bytes": {
						Description: "Number of uncompressed used bytes in the datastore.",
						Type:        schema.TypeInt,
						Optional:    true,
						Computed:    true,
					},
					"unshared_used_bytes": {
						Description: "Unshared used capacity of the datastore in bytes.",
						Type:        schema.TypeInt,
						Optional:    true,
						Computed:    true,
					},
					"uuid": {
						Description: "UUID for the datastore object.",
						Type:        schema.TypeString,
						Optional:    true,
						Computed:    true,
					},
				}},
				Computed: true,
			}},
	}
}

func dataSourceHyperflexDatastoreStatisticRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.HyperflexDatastoreStatistic{}
	if v, ok := d.GetOk("accessibility_summary"); ok {
		x := (v.(string))
		o.SetAccessibilitySummary(x)
	}
	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}
	if v, ok := d.GetOk("creation_time"); ok {
		x := (v.(string))
		o.SetCreationTime(x)
	}
	if v, ok := d.GetOk("datastore_status"); ok {
		x := (v.(string))
		o.SetDatastoreStatus(x)
	}
	if v, ok := d.GetOk("free_capacity_in_bytes"); ok {
		x := int64(v.(int))
		o.SetFreeCapacityInBytes(x)
	}
	if v, ok := d.GetOk("last_modified_time"); ok {
		x := (v.(string))
		o.SetLastModifiedTime(x)
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}
	if v, ok := d.GetOk("mount_summary"); ok {
		x := (v.(string))
		o.SetMountSummary(x)
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}
	if v, ok := d.GetOk("parent_uuid"); ok {
		x := (v.(string))
		o.SetParentUuid(x)
	}
	if v, ok := d.GetOk("total_capacity_in_bytes"); ok {
		x := int64(v.(int))
		o.SetTotalCapacityInBytes(x)
	}
	if v, ok := d.GetOk("un_compressed_used_bytes"); ok {
		x := int64(v.(int))
		o.SetUnCompressedUsedBytes(x)
	}
	if v, ok := d.GetOk("unshared_used_bytes"); ok {
		x := int64(v.(int))
		o.SetUnsharedUsedBytes(x)
	}
	if v, ok := d.GetOk("uuid"); ok {
		x := (v.(string))
		o.SetUuid(x)
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of HyperflexDatastoreStatistic object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexDatastoreStatisticList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		responseErr := responseErr.(models.GenericOpenAPIError)
		return diag.Errorf("error occurred while fetching count of HyperflexDatastoreStatistic: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
	}
	count := countResponse.HyperflexDatastoreStatisticList.GetCount()
	var i int32
	var hyperflexDatastoreStatisticResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexDatastoreStatisticList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching HyperflexDatastoreStatistic: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		results := resMo.HyperflexDatastoreStatisticList.GetResults()
		length := len(results)
		if length == 0 {
			return diag.Errorf("your query for HyperflexDatastoreStatistic data source did not return results. Please change your search criteria and try again")
		}
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["accessibility_summary"] = (s.GetAccessibilitySummary())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)
				temp["class_id"] = (s.GetClassId())
				temp["creation_time"] = (s.GetCreationTime())

				temp["data_protection_peer"] = flattenMapHyperflexDataProtectionPeerRelationship(s.GetDataProtectionPeer(), d)
				temp["datastore_status"] = (s.GetDatastoreStatus())

				temp["dsconfig"] = flattenMapHyperflexHxPlatformDatastoreConfigDt(s.GetDsconfig(), d)
				temp["free_capacity_in_bytes"] = (s.GetFreeCapacityInBytes())

				temp["host_mount_status"] = flattenListHyperflexHxHostMountStatusDt(s.GetHostMountStatus(), d)
				temp["last_modified_time"] = (s.GetLastModifiedTime())
				temp["moid"] = (s.GetMoid())
				temp["mount_summary"] = (s.GetMountSummary())
				temp["object_type"] = (s.GetObjectType())
				temp["parent_uuid"] = (s.GetParentUuid())

				temp["site"] = flattenMapHyperflexHxSiteDt(s.GetSite(), d)

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["total_capacity_in_bytes"] = (s.GetTotalCapacityInBytes())
				temp["un_compressed_used_bytes"] = (s.GetUnCompressedUsedBytes())
				temp["unshared_used_bytes"] = (s.GetUnsharedUsedBytes())
				temp["uuid"] = (s.GetUuid())
				hyperflexDatastoreStatisticResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(hyperflexDatastoreStatisticResults))
	if err := d.Set("results", hyperflexDatastoreStatisticResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(hyperflexDatastoreStatisticResults[0]["moid"].(string))
	return de
}
