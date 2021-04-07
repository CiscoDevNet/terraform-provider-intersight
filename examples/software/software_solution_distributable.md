### Resource Creation

```hcl
resource "intersight_software_solution_distributable" "software_solution_distributable1" {
    name = "software_solution_distributable1"
    sha_512_sum = "23ccf875003a8459559a77816d0fedd395b0fb6a489091778c169496ddc600
    99297f64b4887dae8e384d97f927c438f638b165bccf59f6fdb6e359145702273f"
    size = 7471044747
    component_meta = [object_type = "firmware.ComponentMeta"
        component_label = "BIOS"
    ]
    model = ""
    mdfid = ""
    platform_type = ""
    release_date {
        object_type = "softwarerepository.Release"
        moid = var.release
    }
    catalog {
        object_type = "softwarerepository.Catalog"
        moid = var.catalog
    }

    file_path = "IKS/shared/ccp-tenant-image-1.18.12-ubuntu18-94a8577-2.ova/ccp-tenant-image-1.18.12-ubuntu18-94a8577-2.ova"
    vendor = "Cisco"
    distributable_metas = [{
        object_type = "firmware.DistributableMeta"
    }]
    release_url = < url
    for release notes of this image >
        suppotred_models = [
            "HyperFlex Data Platform"
        ]
    sub_type = "osimage"
    solution_name = "IKS"
    version = "ccp-tenant-image-1.18.12-ubuntu18-94a8577-2.ova"
}
```