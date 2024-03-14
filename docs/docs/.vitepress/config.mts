import { defineConfig } from "vitepress";

// https://vitepress.dev/reference/site-config
export default defineConfig({
    title: "Amprobe",
    description: "amprobe",

    themeConfig: {
        // https://vitepress.dev/reference/default-theme-config
        nav: [
            { text: "联系我们", link: "/views/tutorials" },
        ],

        sidebar: [
            {
                text: "使用手册",
                items: [
                    { text: "docker 安装", link: "/views/docker" },
                    { text: "docker-compose 安装(推荐)", link: "/views/docker-compose" },
                    { text: "后续计划", link: "/views/next" },
                ],
            },
            {
                text: "配置文件",
                items: [
                    { text: "config", link: "/views/config" },
                    { text: "compose", link: "/views/compose" },
                ],
            },
        ],

        socialLinks: [{ icon: "github", link: "https://github.com/amuluze/amprobe" }],
    },
});
