package system

import (
	"common/system/cpu"
	"common/system/disk"
	"common/system/mem"
)

//go:generate msgp -tests=false #system

type Info struct {
	DiskInfo  *disk.Info    `json:"diskInfo" msg:",inline"`
	MemStatus *mem.Status   `json:"memStatus" msg:",inline"`
	CpuStatus *cpu.Stat     `json:"cpuStatus" msg:",inline"`
	IoStatus  *disk.IOStats `json:"ioStatus" msg:",inline"`
}

func ZeroInfo() *Info {
	return &Info{
		DiskInfo:  &disk.Info{},
		MemStatus: &mem.Status{},
		CpuStatus: &cpu.Stat{},
		IoStatus:  &disk.IOStats{},
	}
}

func NewInfo(diskPath string) (*Info, error) {
	diskInfo, err := disk.GetInfo(diskPath)
	if err != nil {
		return nil, err
	}
	cpuStat, err := cpu.StatInfo()
	if err != nil {
		return nil, err
	}
	memStat, err := mem.Stat()
	if err != nil {
		return nil, err
	}
	ioStat, err := disk.GetAverageIOStats()
	if err != nil {
		return nil, err
	}
	return &Info{
		DiskInfo:  &diskInfo,
		MemStatus: &memStat,
		CpuStatus: &cpuStat,
		IoStatus:  ioStat,
	}, nil
}
