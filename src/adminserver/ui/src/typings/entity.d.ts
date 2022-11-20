declare interface Version {
    hash: string
    size: number
    sequence: number
    ts: number
    ecAlgo: number
    dataShards: number
    parityShards: number
    shardSize: number
    locate: string[]
}

declare interface Metadata {
    name: string
    createTime: number
    updateTime: number
}

declare interface DiskInfo {
    total: number
    free: number
    used: number
}

declare interface MemStat {
    all: number
    used: number
    free: number
    self: number
}

declare interface CpuStat {
    usedPercent: number
    logicalCount: number
    physicalCount: number
}

declare interface SystemInfo {
    diskInfo: DiskInfo
    memStatus: MemStat
    cpuStatus: CpuStat
}

declare interface TimeStat {
    time: string
    percent: number
}

declare interface ServerInfo {
    serverId: string
    httpAddr: string
    rpcAddr: string
    isMaster?: boolean
    sysInfo: SystemInfo
}
