export function formatBytes(bytes: number, decimals = 1) {
    if (!+bytes) return '0 B'

    const k = 1024
    const dm = decimals < 0 ? 0 : decimals
    const sizes = ['B', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB']

    const i = Math.floor(Math.log(bytes) / Math.log(k))

    return `${parseFloat((bytes / Math.pow(k, i)).toFixed(dm))} ${sizes[i]}`
}

export function invokeInterval(fn: ()=>void, ts: number) {
    fn()
    setInterval(fn, ts)
}

export function cutStr(s: string, maxLen: number) {
    if (s.length <= maxLen) {
        return s
    }
    return s.substring(0, maxLen - 3) + "..."
}