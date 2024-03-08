<template>
    <div class="footer-content">
        <div class="lh">
            Copyright © 2023 红豆南墙 | 个人主页 | Welcome to HongDouNanQiang !
        </div>
        <div class="lh">
            本站已安全运行 {{ uptime }}
        </div>
        <div class="lh"><a href="https://beian.miit.gov.cn/#/Integrated/recordQuery" class="link">苏ICP备2023021145号</a></div>
    </div>
</template>

<script>
import { onMounted, ref } from 'vue';

export default {
    name: "newFooter",
    setup() {
        const uptime = ref('');

        const updateUptime = () => {
            const launchDate = new Date(2023, 6, 1); // 设置网站启动日期
            const now = new Date();
            const uptimeValue = now - launchDate;
            const seconds = Math.floor(uptimeValue / 1000) % 60;
            const minutes = Math.floor(uptimeValue / (1000 * 60)) % 60;
            const hours = Math.floor(uptimeValue / (1000 * 60 * 60)) % 24;
            const days = Math.floor(uptimeValue / (1000 * 60 * 60 * 24)) % 365;
            const years = Math.floor(uptimeValue / (1000 * 60 * 60 * 24 * 365));

            uptime.value = `${years}年 ${days}天 ${hours}小时 ${minutes}分钟 ${seconds}秒`;
        };

        onMounted(() => {
            updateUptime();
            setInterval(updateUptime, 1000);
        });

        return {
            uptime,
        };
    },
};
</script>

<style scoped>
.footer-content{
    text-align: center;
}
.lh{
    margin-bottom:7px;
}
</style>
