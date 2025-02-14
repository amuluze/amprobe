/**
 * @Author     : Amu
 * @Date       : 2024/11/14 00:06
 * @Description:
 */

// * 主题配置
export const themeConfig: Record<Theme.ThemeType, { [ key: string ]: string }> = {
    light: {
        // 菜单样式#f4f0f0
        '--el-bg-color': '#f0f1f5',
        '--el-menu-bg-color': '#eff1f5',
        '--el-menu-hover-bg-color': '#71c4ef',
        '--el-menu-active-bg-color': '#409eff',
        '--el-menu-text-color': '#3b3c3d',
        '--el-menu-active-color': '#ffffff',
        '--el-menu-hover-text-color': '#ffffff',
        '--el-menu-horizontal-sub-item-height': '50px',
        // 侧边栏样式
        '--el-aside-logo-text-color': '#303133',
        '--el-aside-border-color': '#e4e7ed',
        // 头部样式
        '--el-header-logo-text-color': '#303133',
        '--el-header-bg-color': '#fffefb',
        '--el-header-text-color': '#303133',
        '--el-header-text-color-regular': '#606266',
        '--el-header-border-color': '#e4e7ed',
    },
    dark: {
        // 菜单样式
        '--el-bg-color': '#141414',
        '--el-menu-bg-color': '#1d1e1f',
        '--el-menu-hover-bg-color': '#000000',
        '--el-menu-active-bg-color': '#000000',
        '--el-menu-text-color': '#dadada',
        '--el-menu-active-color': '#ffffff',
        '--el-menu-hover-text-color': '#ffffff',
        '--el-menu-horizontal-sub-item-height': '50px',
        // 侧边栏样式
        '--el-aside-logo-text-color': '#dadada',
        '--el-aside-border-color': '#414243',
        // 头部样式
        '--el-header-logo-text-color': '#dadada',
        '--el-header-bg-color': '#141414',
        '--el-header-text-color': '#e5eaf3',
        '--el-header-text-color-regular': '#cfd3dc',
        '--el-header-border-color': '#414243',
    },
}

export const echartsThemeData = [
    {
        name: 'vintage',
        background: '#fef8ef',
        theme: [
            '#d87c7c',
            '#919e8b',
            '#d7ab82',
            '#6e7074',
            '#61a0a8',
            '#efa18d',
            '#787464',
            '#cc7e63',
            '#724e58',
            '#4b565b',
        ],
    },
    {
        name: 'dark',
        background: '#333',
        theme: [
            '#dd6b66',
            '#759aa0',
            '#e69d87',
            '#8dc1a9',
            '#ea7e53',
            '#eedd78',
            '#73a373',
            '#73b9bc',
            '#7289ab',
            '#91ca8c',
            '#f49f42',
        ],
    },
    {
        name: 'westeros',
        background: 'transparent',
        theme: ['#516b91', '#59c4e6', '#edafda', '#93b7e3', '#a5e7f0', '#cbb0e3'],
    },
    {
        name: 'essos',
        background: 'rgba(242,234,191,0.15)',
        theme: ['#893448', '#d95850', '#eb8146', '#ffb248', '#f2d643', '#ebdba4'],
    },
    {
        name: 'wonderland',
        background: 'transparent',
        theme: ['#4ea397', '#22c3aa', '#7bd9a5', '#d0648a', '#f58db2', '#f2b3c9'],
    },
    {
        name: 'walden',
        background: 'rgba(252,252,252,0)',
        theme: ['#3fb1e3', '#6be6c1', '#626c91', '#a0a7e6', '#c4ebad', '#96dee8'],
    },
    {
        name: 'chalk',
        background: '#293441',
        theme: ['#fc97af', '#87f7cf', '#f7f494', '#72ccff', '#f7c5a0', '#d4a4eb', '#d2f5a6', '#76f2f2'],
    },
    {
        name: 'infographic',
        background: 'transparent',
        theme: [
            '#C1232B',
            '#27727B',
            '#FCCE10',
            '#E87C25',
            '#B5C334',
            '#FE8463',
            '#9BCA63',
            '#FAD860',
            '#F3A43B',
            '#60C0DD',
            '#D7504B',
            '#C6E579',
            '#F4E001',
            '#F0805A',
            '#26C0C0',
        ],
    },
    {
        name: 'macarons',
        background: 'transparent',
        theme: [
            '#2ec7c9',
            '#b6a2de',
            '#5ab1ef',
            '#ffb980',
            '#d87a80',
            '#8d98b3',
            '#e5cf0d',
            '#97b552',
            '#95706d',
            '#dc69aa',
            '#07a2a4',
            '#9a7fd1',
            '#588dd5',
            '#f5994e',
            '#c05050',
            '#59678c',
            '#c9ab00',
            '#7eb00a',
            '#6f5553',
            '#c14089',
        ],
    },
    {
        name: 'roma',
        background: 'transparent',
        theme: [
            '#E01F54',
            '#001852',
            '#f5e8c8',
            '#b8d2c7',
            '#c6b38e',
            '#a4d8c2',
            '#f3d999',
            '#d3758f',
            '#dcc392',
            '#2e4783',
            '#82b6e9',
            '#ff6347',
            '#a092f1',
            '#0a915d',
            '#eaf889',
            '#6699FF',
            '#ff6666',
            '#3cb371',
            '#d5b158',
            '#38b6b6',
        ],
    },
    {
        name: 'shine',
        background: 'transparent',
        theme: ['#c12e34', '#e6b600', '#0098d9', '#2b821d', '#005eaa', '#339ca8', '#cda819', '#32a487'],
    },
    {
        name: 'purple-passion',
        background: 'rgba(91,92,110,1)',
        theme: ['#8a7ca8', '#e098c7', '#8fd3e8', '#71669e', '#cc70af', '#7cb4cc'],
    },
]
