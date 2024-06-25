/**
 * @Author     : Amu
 * @Date       : 2024/3/7 13:38
 * @Description:
 */
export interface Container {
    id: string
    name: string
    image: string
    ip: string
    state: string
    uptime: string
    cpu_percent: number
    memory_percent: number
    memory_usage: number
    memory_limit: number
}

export interface StartContainerArgs {
    container_id: string
}
export interface StopContainerArgs {
    container_id: string
}
export interface RemoveContainerArgs {
    container_id: string
}
export interface RestartContainerArgs {
    container_id: string
}

export interface ContainerQueryResult {
    data: Container[]
    total: number
    page: number
    size: number
}

export interface Image {
    id: string
    name: string
    tag: string
    created: string
    size: string
    number: number
}

export interface RemoveImageArgs {
    image_id: string
}

export interface ImageQueryResult {
    data: Image[]
    total: number
    page: number
    size: number
}

export interface DockerInfo {
    docker_version: string
    api_version: string
    min_api_version: string
    git_commit: string
    go_version: string
    os: string
    arch: string
}

export interface GetDockerRegistryMirrorsResult {
    registry_mirrors: string[]
}

export interface SetDockerRegistryMirrorsArgs {
    registry_mirrors: string[]
}
