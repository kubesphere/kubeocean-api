package qingcloud

type InstanceType string

const (
	InstanceType_S1_LARGE_R2   InstanceType = "s1.large.r2"
	InstanceType_S1_XLARGE_R2  InstanceType = "s1.xlarge.r2"
	InstanceType_S1_2XLARGE_R2 InstanceType = "s1.2xlarge.r2"
)

var InstanceCPUMap = map[InstanceType]int32{
	InstanceType_S1_LARGE_R2:   2,
	InstanceType_S1_XLARGE_R2:  4,
	InstanceType_S1_2XLARGE_R2: 8,
}

var InstanceMemMap = map[InstanceType]int32{
	InstanceType_S1_LARGE_R2:   4096,
	InstanceType_S1_XLARGE_R2:  8192,
	InstanceType_S1_2XLARGE_R2: 16384,
}
