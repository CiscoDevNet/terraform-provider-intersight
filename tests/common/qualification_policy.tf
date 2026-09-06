resource "intersight_resourcepool_qualification_policy" "test" {
  name        = "quaificationPolicy1"
  description = "demo quaification policy"
  object_type = "resourcepool.QualificationPolicy"
  organization {
    object_type = "organization.Organization"
    moid        = data.intersight_organization_organization.default.results.0.moid
  }
  qualifiers = [
    {
      additional_properties = jsonencode(
        {
          AdaptorsRange = {
            MaxValue      = 2
            MinValue      = 1
            ObjectType    = "resource.AdaptorsRangeFilter"
          }
        }
      )
      class_id    = "resource.NetworkAdaptorQualifier"
      object_type = "resource.NetworkAdaptorQualifier"
    },
    {
      additional_properties = jsonencode(
        {
          ChassisAndSlotIdRange = [
            {
              ChassisIdRange = {
                ClassId       = "resource.ChassisIdRangeFilter"
                ConditionType = "RANGE"
                MaxValue      = 5
                MinValue      = 1
                ObjectType    = "resource.ChassisIdRangeFilter"
              }
              ObjectType = "resource.ChassisAndSlotQualification"
              SlotIdRanges = [
                {
                  ClassId       = "resource.SlotIdRangeFilter"
                  ConditionType = "RANGE"
                  MaxValue      = 3
                  MinValue      = 1
                  ObjectType    = "resource.SlotIdRangeFilter"
                },
                {
                  ClassId       = "resource.SlotIdRangeFilter"
                  ConditionType = "RANGE"
                  MaxValue      = 7
                  MinValue      = 4
                  ObjectType    = "resource.SlotIdRangeFilter"
                }
              ]
            }
          ]
        }
      )
      class_id    = "resource.BladeQualifier"
      object_type = "resource.BladeQualifier"
    },
    {
      additional_properties = jsonencode(
        {
          CpuCoresRange = {
            ConditionType = "RANGE"
            MaxValue      = 48
            MinValue      = 24
            ObjectType    = "resource.CpuCoreRangeFilter"
          }
          Pids = []
          SpeedRange = {
            ConditionType = "RANGE"
            MaxValue      = 0
            MinValue      = 0
            ObjectType    = "resource.CpuSpeedRangeFilter"
          }
          Vendor = ""
        }
      )
      class_id    = "resource.ProcessorQualifier"
      object_type = "resource.ProcessorQualifier"
    },
    {
      additional_properties = jsonencode(
        {
          FabricInterConnectPids = ["UCS-FI-6454"]
        }
      )
      class_id    = "resource.DomainQualifier"
      object_type = "resource.DomainQualifier"
    },
    {
      additional_properties = jsonencode(
        {
          GpuControllersRange = {
            ConditionType = "RANGE"
            MaxValue      = 0
            MinValue      = 0
            ObjectType    = "resource.GpuControllersRangeFilter"
          }
          GpuEvaluationType = "ServerWithoutGpu"
          Pids              = []
          Vendor            = ""
        }
      )
      class_id    = "resource.GpuQualifier"
      object_type = "resource.GpuQualifier"
    },
    {
      additional_properties = jsonencode(
        {
          MemoryCapacityRange = {
            ConditionType = "RANGE"
            MaxValue      = 512
            MinValue      = 256
            ObjectType    = "resource.MemoryCapacityRangeFilter"
          }
          MemoryUnitsRange = {
            ConditionType = "RANGE"
            MaxValue      = 0
            MinValue      = 0
            ObjectType    = "resource.MemoryUnitsRangeFilter"
          }
        }
      )
      class_id    = "resource.MemoryQualifier"
      object_type = "resource.MemoryQualifier"
    },
    {
      additional_properties = jsonencode(
        {
          RackIdRange = [
            {
              ConditionType = "RANGE"
              MaxValue      = 5
              MinValue      = 1
              ObjectType    = "resource.RackIdRangeFilter"
            }
          ]
        }
      )
      class_id    = "resource.RackServerQualifier"
      object_type = "resource.RackServerQualifier"
    },
    {
      additional_properties = jsonencode(
        {
          ChassisTags = [
            {
              Key        = "dummy1"
              ObjectType = "resource.Tag"
              Value      = "dummy2"
            }
          ]
          ServerTags = [
            {
              Key        = "dummy8"
              ObjectType = "resource.Tag"
              Value      = "Dummy9"
            }
          ]
        }
      )
      class_id    = "resource.TagQualifier"
      object_type = "resource.TagQualifier"
    }
  ]
  tags = []
}