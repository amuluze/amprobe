/**
 * @Author     : Amu
 * @Date       : 2024/10/17 16:35:32
 * @Description:
 */

export interface EmailSetting {
    id: number
    server: string
    port: number
    sender: string
    password: string
    receiver: string
}

export interface AlarmThreshold {
    id: number
    type: string
    duration: number
    threshold: number
}

export interface AlarmThresholdQueryResult {
    data: AlarmThreshold[]
}

export interface AlarmThresholdUpdateArgs {
    id: number
    type: string
    duration: number
    threshold: number
}
