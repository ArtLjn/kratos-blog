<template>
  <el-card v-loading="loading"
  element-loading-text="Loading..."
  :element-loading-spinner="svg"
  element-loading-svg-view-box="-10, -10, 50, 50"
  element-loading-background="rgba(122, 122, 122, 0.3)"
   class="card2">
    <div class="title">
      <font-awesome-icon class="icon" :icon="['fas', 'comments']"></font-awesome-icon>
      <span class="title-text">评论</span>
    </div>
    <form  v-if="replyIndex === null" @submit.prevent="saveComment">
      <textarea class="textarea-field" placeholder="好言一句三冬暖 恶语伤人六月寒" v-model="addComment.comment"  @keydown.enter.prevent></textarea>
      <div class="expression" @click="clickShowEmjio">(っ◔◡◔)っ</div>
      <div class="button-group">
        <el-button class="submit-button" @click="saveComment" :disabled="isCountingDown">
          {{ isCountingDown ? '发表评论 (' + countdown + ')' : '发表评论' }}
        </el-button>
      </div>
    </form>
    <Picker :data="emojiIndex" set="twitter" @select="showEmoji" v-if="hideEmoji" class="a"/>
    <div class="comment">
      <span v-if="commentList.length === 0" class="com">来评论吧~</span>
      <span class="comment-title" v-if="commentList.length > 0">评论</span>
      <div class="comment-list">
        <div v-for="(comment, index) in commentList" :key="index" class="comment-item">
          <div class="comment-header">
            <div style="display:flex;">
              <img :src="`http://q2.qlogo.cn/headimg_dl?dst_uin=${comment.email}&spec=1`" class="photo"/>
              <span class="comment-name">{{ comment.name }}</span>
            </div>
            <span class="comment-time">{{ comment.commentTime }}</span>
            <span class="comment-addr">
              <font-awesome-icon :icon="['fas', 'location-dot']" style="color: #495a79;" /> {{comment.commentAddr}}</span>
            <span v-if="replyIndex !== index" @click="replyIndex = index" class="reply-button">回复</span>
          </div>   
          <div class="comment-body">
            <div>{{ comment.comment }}</div>
          </div>

            <div v-if="comment.child" >
              <div v-for="(reward,index) in comment.child" :key="index" style="margin-left:35px;">
                <div style="display:flex;">
                <img :src="`http://q2.qlogo.cn/headimg_dl?dst_uin=${reward.email}&spec=1`" class="photo"/>
                <span class="comment-name">{{reward.name}}</span></div>
                <span class="comment-time">{{reward.commentTime}}</span>
                <span class="comment-addr">
                  <font-awesome-icon :icon="['fas', 'location-dot']" style="color: #495a79;" /> {{reward.commentAddr}}</span>
                <div class="comment-body">
                  <span>{{reward.comment}}</span>
                </div>
            </div>
          </div>
          <div v-if="replyIndex === index">
            <form class="reply-form" @submit.prevent="saveReward(comment.id)">
              <textarea class="textarea-field" placeholder="好言一句三冬暖 恶语伤人六月寒" v-model="RewardData.comment"  @keydown.enter.prevent></textarea>
              <div style="float:left;" @click="clickShowEmjio2" class="expression">(っ◔◡◔)っ</div>
              <div class="button-group">
                <el-button class="submit-button" @click="saveReward(comment.id)" :disabled="isCountingDown">
                  {{ isCountingDown ? '发表评论 (' + countdown + ')' : '发表评论' }}
                </el-button>
                <el-button class="cancel-button" @click="cancelReply">取消</el-button>
              </div>
            </form>
            <Picker :data="emojiIndex" set="twitter" @select="showEmoji2" v-if="hideEmoji2" class="b"/>
          </div>
        </div>
      </div>
    </div>
  </el-card>
</template>

<script>
import { onMounted, reactive, ref } from 'vue';
import { ElMessage} from 'element-plus';
import { useRoute } from 'vue-router';
import data from "emoji-mart-vue-fast/data/all.json";
import "emoji-mart-vue-fast/css/emoji-mart.css";
import { Picker, EmojiIndex } from "emoji-mart-vue-fast/src";
import { getIp } from '@/assets/js/ip';
import {AddComment, AddReward, checkPath, getComment, validComment} from "@/api/commentFunc";
import {SUCCESS_REQUEST} from "@/api/status";

export default {
  props: {
    message: String
  },
  components:{
    Picker
  },
  setup(props) {
    const emojiIndex = new EmojiIndex(data);
    const showEmoji = (emoji) => {
      addComment.comment += emoji.native;
    }
    const showEmoji2 = (emoji) => {
      RewardData.comment += emoji.native;
    }
    const hideEmoji = ref(false);
    const clickShowEmjio = () => {
      hideEmoji.value = !hideEmoji.value;
    }
    const hideEmoji2 = ref(false);
    const clickShowEmjio2 = () => {
      hideEmoji2.value = !hideEmoji2.value;
    }
    const route = useRoute();
    const addComment = reactive({
      comment: '',
      article_id: 'message',
      comment_addr:""
    });
    const RewardData = reactive({
        comment_addr:"",
        comment: '',
        article_id: 'message',
        parent_id:0,
    })
    const replyIndex = ref(null);


    /*
    评论接口
    */
    /*请求禁用*/
    const initialCountdown = 5; // 倒计时初始值
    const countdown = ref(initialCountdown); // 倒计时时间，单位为秒
    const isCountingDown = ref(false); // 按钮状态
    const loading = ref(false); // 加载状态，默认为false
    const startCountdown = () => {
      let timer = setInterval(() => {
        countdown.value--;
        if (countdown.value === 0) {
          clearInterval(timer);
          isCountingDown.value = false;
        }
      }, 1000);
    };

    const saveComment = () => {
      if (!validComment(addComment)) {
        return;
      } else if (props.message) {
        addComment.article_id = props.message;
      } else {
        checkPath(route.path, addComment);
      }
      if (!isCountingDown.value) {
        countdown.value = initialCountdown; // 重置倒计时为初始值
        isCountingDown.value = true;
        loading.value = true; // 显示加载状态

        startCountdown(); // 开始倒计时

        addComment.comment_addr = Ip.value;
        AddComment(addComment).then((res) => {
          if (res.code === SUCCESS_REQUEST) {
            ElMessage.success('评论成功');
            Object.keys(addComment).forEach(key => {
              addComment[key] = null;
            });
            const userAgent = window.navigator.userAgent;
            console.log(userAgent);
            getAllComment();
            loading.value = false; // 取消加载状态
          } else {
            ElMessage.error(res.result);
            loading.value = false; // 取消加载状态
          }
        });
      }
    };
    const saveReward = (id) => {
      RewardData.parent_id = id;
      if (!validComment(RewardData)) {
        return;
      } else if (props.message) {
          RewardData.article_id = props.message;
      } else {
        checkPath(route.path,RewardData)
      }
      if (!isCountingDown.value) {
        countdown.value = initialCountdown; // 重置倒计时为初始值
        isCountingDown.value = true;
        loading.value = true; // 显示加载状态

        startCountdown(); // 开始倒计时

        RewardData.comment_addr = Ip.value;
        AddReward(RewardData).then((res) => {
          if (res.code == SUCCESS_REQUEST) {
            ElMessage.success("回复成功")
            Object.keys(RewardData).forEach(key => {
              RewardData[key] = null;
            })
            getAllComment();
            loading.value = false; // 取消加载状态
          } else {
            ElMessage.error(res.result);
            loading.value = false; // 取消加载状态
          }
        })
      }
    }


    let commentList = ref([]);

    const GetComment = (path) => {
      getComment(path).then((res) => {
        if (res.code === SUCCESS_REQUEST) {
          commentList.value = res.list
        }
      })
    }

    const getAllComment = () => {
      switch (true) {
        case !!props.message:
          GetComment(props.message)
          break;
        case route.path === '/message':
          GetComment("message")
          break;
        case route.path === '/friendlink':
          GetComment("friendlink")
          break;
        case route.path === '/talk':
          GetComment("talk")
          break;
        default:
          // 如果没有匹配到任何情况，可以在这里添加默认操作。
      }
    };
    const cancelReply = () => {
      replyIndex.value = null;
    };
    
    const Ip = ref('');
    getIp().then(data => {
      Ip.value = data.ip;
    })
    onMounted(() => {
      getAllComment();
      getIp();
    });
    return {
      commentList,
      addComment,
      saveComment,
      replyIndex,
      cancelReply,
      RewardData,
      saveReward,
      emojiIndex,
      showEmoji,
      showEmoji2,
      hideEmoji,
      clickShowEmjio,
      clickShowEmjio2,
      hideEmoji2,
      countdown,
      isCountingDown,
      loading,
    };
  }
}
</script>

<style scoped>
.card2 {
  margin-bottom: 150px;
  min-height: 220px;
  border-radius: 10px;
  padding: 20px;
  width: calc(66%);
  background-color: #f5f5f5;
  overflow: visible !important;
}

.title {
  display: flex;
  align-items: center;
  color: rgb(33, 14, 67);
  font-size: 20px;
  font-weight: bold;
  margin-bottom: 10px;
}

.icon {
  margin-right: 5px;
}

.title-text {
  line-height: 1;
}


.form-group {
  display: flex;
  margin-bottom: 10px;
  margin-left: 5px;
}

.input-field {
  width: 50%;
  height: 40px;
  outline: none;
  border: none;
  border-radius: 10px;
  text-indent: 15px;
  background-color: #ffffff;
  box-shadow: 2px 2px 2px 1px rgba(0, 0, 0, 0.2);
}

.textarea-field {
  width: 99%;
  border-radius: 10px;
  height: auto;
  min-height: 90px;
  outline: none;
  resize: none;
  border: none;
  font-size: 16px;
  text-indent: 10px;
  padding-top: 10px;
  direction: ltr;
  color: black;
  font-weight: 100;
  margin: 0px 5px auto 5px;
  box-shadow: 2px 2px 2px 1px rgba(0, 0, 0, 0.2);
  font-family: Verdana, Geneva, Tahoma, sans-serif;
}

.button-group {
  display: flex;
  justify-content: flex-end;
  position: relative;
  z-index: 1;
}

.submit-button {
  background-color: red;
  border-radius: 10px;
  color: white;
  margin: 10px;
  padding: 18px;
  box-shadow: 2px 2px 2px 1px rgba(0, 0, 0, 0.2);
}

.cancel-button {
  background-color: rgb(153, 188, 197);
  border-radius: 10px;
  margin: 10px;
  padding: 18px;
  box-shadow: 2px 2px 2px 1px rgba(0, 0, 0, 0.2);
}

.comment-list {
  margin-top: 10px;
  position: relative;
  z-index: 1;
}

.comment-title {
  color: rgb(33, 14, 67);
  font-size: 20px;
  font-weight: bold;
}

.comment-item {
  margin-bottom: 20px;
}

.comment-name {
  font-size: 12px;
  color: rgb(5, 0, 0);
  font-weight: 800;
  margin:3px auto auto 5px;
}

.comment-time {
  font-size: 12px;
  color: #4b4a4a;
  margin-top: 4px;
}
.comment-addr {
  font-size: 12px;
  color: #4b4a4a;
  margin-top: 4px;
  margin-left: 10px;
}

.reply-button {
  margin-left: auto;
  font-size: 12px;
  color: red;
  margin-bottom: 1px;
  cursor: pointer;
  float: right;
}

.comment-body {
  margin-top: 10px;
  text-indent: 10px;
  margin-bottom: 15px;
}

.com {
  color: rgb(33, 14, 67);
  font-size: 20px;
  font-weight: bold;
  justify-content: center;
}
@media screen and (max-width: 600px) {
  .card2{
    width: calc(88%);
  }
}
.photo{
  border-radius: 50%;
}
.b{
  position: absolute;
  z-index: 3;
}
.a{
  position: absolute;
  z-index: 2;
}
.expression{
  float:left;margin:15px;
}
.comment{
  position: relative;
}
</style>
