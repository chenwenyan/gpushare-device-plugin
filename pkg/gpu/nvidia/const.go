package nvidia

import (
	pluginapi "k8s.io/kubernetes/pkg/kubelet/apis/deviceplugin/v1beta1"
)

// MemoryUnit describes GPU Memory, now only supports Gi, Mi
type MemoryUnit string

const (
	resourceName    = "aliyun.com/gpu-mem"
	resourceCount   = "aliyun.com/gpu-count"
	gpu0Utilization = "aliyun.com/gpu0-util"
	mem0Utilization = "aliyun.com/mem0-util"
	gpu1Utilization = "aliyun.com/gpu1-util"
	mem1Utilization = "aliyun.com/mem1-util"
	gpu0Processes   = "aliyun.com/gpu0-processes"
	gpu1Processes   = "aliyun.com/gpu1-processes"

	serverSock = pluginapi.DevicePluginPath + "aliyungpushare.sock"

	OptimisticLockErrorMsg = "the object has been modified; please apply your changes to the latest version and try again"

	allHealthChecks             = "xids"
	containerTypeLabelKey       = "io.kubernetes.docker.type"
	containerTypeLabelSandbox   = "podsandbox"
	containerTypeLabelContainer = "container"
	containerLogPathLabelKey    = "io.kubernetes.container.logpath"
	sandboxIDLabelKey           = "io.kubernetes.sandbox.id"

	envNVGPU                   = "NVIDIA_VISIBLE_DEVICES"
	EnvResourceIndex           = "ALIYUN_COM_GPU_MEM_IDX"
	EnvResourceByPod           = "ALIYUN_COM_GPU_MEM_POD"
	EnvResourceByContainer     = "ALIYUN_COM_GPU_MEM_CONTAINER"
	EnvResourceByDev           = "ALIYUN_COM_GPU_MEM_DEV"
	EnvAssignedFlag            = "ALIYUN_COM_GPU_MEM_ASSIGNED"
	EnvResourceAssumeTime      = "ALIYUN_COM_GPU_MEM_ASSUME_TIME"
	EnvResourceAssignTime      = "ALIYUN_COM_GPU_MEM_ASSIGN_TIME"
	EnvNodeLabelForDisableCGPU = "cgpu.disable.isolation"

	GiBPrefix = MemoryUnit("GiB")
	MiBPrefix = MemoryUnit("MiB")
)
