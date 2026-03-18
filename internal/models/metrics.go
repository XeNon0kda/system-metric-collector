package models

import (
    "time"
)

type Metrics struct {
    Timestamp time.Time           `json:"timestamp"`
    CPU       CPUStats             `json:"cpu"`
    Memory    MemoryStats          `json:"memory"`
    Disk      []DiskStats          `json:"disk"`
    Network   []NetworkStats       `json:"network"`
    Processes int                  `json:"processes_count"`
}

type CPUStats struct {
    Percent float64 `json:"percent"`
    Cores   int     `json:"cores"`
}

type MemoryStats struct {
    Total       uint64  `json:"total"`
    Available   uint64  `json:"available"`
    Used        uint64  `json:"used"`
    UsedPercent float64 `json:"used_percent"`
}

type DiskStats struct {
    Mountpoint  string  `json:"mountpoint"`
    Total       uint64  `json:"total"`
    Free        uint64  `json:"free"`
    Used        uint64  `json:"used"`
    UsedPercent float64 `json:"used_percent"`
}

type NetworkStats struct {
    InterfaceName string `json:"interface_name"`
    BytesSent     uint64 `json:"bytes_sent"`
    BytesRecv     uint64 `json:"bytes_recv"`
}