package collector

import (
    "context"
    "runtime"
    "time"

    "github.com/shirou/gopsutil/v3/cpu"
    "github.com/shirou/gopsutil/v3/disk"
    "github.com/shirou/gopsutil/v3/mem"
    "github.com/shirou/gopsutil/v3/net"
    "github.com/shirou/gopsutil/v3/process"

    "sysmc/internal/models"
)

type Collector interface {
    Collect(ctx context.Context) (*models.Metrics, error)
}

type SystemCollector struct{}

func NewSystemCollector() *SystemCollector {
    return &SystemCollector{}
}

func (c *SystemCollector) Collect(ctx context.Context) (*models.Metrics, error) {
    metrics := &models.Metrics{
        Timestamp: time.Now(),
    }

    // CPU
    cpuPercents, err := cpu.Percent(0, false)
    if err == nil && len(cpuPercents) > 0 {
        metrics.CPU.Percent = cpuPercents[0]
    }
    metrics.CPU.Cores = runtime.NumCPU()

    // Memory
    vmStat, err := mem.VirtualMemory()
    if err == nil {
        metrics.Memory = models.MemoryStats{
            Total:       vmStat.Total,
            Available:   vmStat.Available,
            Used:        vmStat.Used,
            UsedPercent: vmStat.UsedPercent,
        }
    }

    // Disk
    partitions, err := disk.Partitions(false)
    if err == nil {
        for _, p := range partitions {
            usage, err := disk.Usage(p.Mountpoint)
            if err != nil {
                continue
            }
            metrics.Disk = append(metrics.Disk, models.DiskStats{
                Mountpoint:  p.Mountpoint,
                Total:       usage.Total,
                Free:        usage.Free,
                Used:        usage.Used,
                UsedPercent: usage.UsedPercent,
            })
        }
    }

    // Network
    netStats, err := net.IOCounters(true)
    if err == nil {
        for _, n := range netStats {
            metrics.Network = append(metrics.Network, models.NetworkStats{
                InterfaceName: n.Name,
                BytesSent:     n.BytesSent,
                BytesRecv:     n.BytesRecv,
            })
        }
    }

    // Processes count
    procs, err := process.Processes()
    if err == nil {
        metrics.Processes = len(procs)
    }

    return metrics, nil
}