<template>
    <el-tooltip content="感谢打赏" placement="top">
        <button class="button" @click="tipping">
            <font-awesome-icon :icon="['fas', 'piggy-bank']" style="font-size:20px;"/>
        </button>
    </el-tooltip>
    <el-dialog v-model="giveShang" style="border-radius:10px;width:400px;height:450px;"
     :show-close="false" :close-delay="100" :lock-scroll="false" :append-to-body="true"
     >
        <div class="tipping">
            <h1>感谢打赏</h1>
            <img src="../../../assets/image/IMG_WeChat.jpg" alt="微信支付" class="img" v-if="payment === 'wechat'">
            <img src="../../../assets/image/IMG_20230705_145116.jpg" alt="支付宝支付" class="img" v-if="payment === 'zhifubao'">
            <div style="margin-top:20px;">
                <el-radio v-model="payment" class="radio" label="wechat"><img src="../../../assets/image/wx.png" ></el-radio>
                <el-radio v-model="payment" class="radio" label="zhifubao"><img src="../../../assets/image/ali.png" alt=""></el-radio>
            </div>
        </div>
    </el-dialog>
</template>

<script>
import { ref } from 'vue';

export default {
    setup() {
        const giveShang = ref(false);
        const payment = ref('wechat');

        const tipping = () => {
            giveShang.value = true;
        }

        return {
            giveShang,
            tipping,
            payment
        }
    }
}
</script>
<style scoped>
.button{
    border: rgb(95, 4, 4) 4px solid;
    border-radius: 50%;
    padding: 10px;
    margin: 15px;
}
.img{
    width: 200px;
    height: 200px;
}
.tipping{
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
}
.radio{
    font-size:large;
}
</style>
