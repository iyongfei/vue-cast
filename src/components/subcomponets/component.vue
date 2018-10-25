<template>
    <div class="cmt-container">
        <h3>评论子组件</h3>

        <textarea name="" id="" cols="30" rows="10"
         placeholder="最多吐槽120字" maxlength="120" v-model="msg">
        </textarea>

        <mt-button type="primary" size="large" @click="commit">
            发表评论
        </mt-button>
        <!-- 评论列表 -->

        <div class="cmt-list">
            <div class="cmt-item" v-for="(comment,index)  in comments" :key="comment.addtime">
                <div class="cmt-title">
                    第{{index+1}}楼&nbsp;
                    发表时间是{{comment.addtime | dateFormat}}
                </div>
                <div class="cmt-body">
                    {{comment.content ==='undefined' ? "用户很懒没有发表" : comment.content}}
                </div>
            </div>
        </div>


        <mt-button type="primary" size="large" plain @click="getMore">
        加载更多
        </mt-button>
    </div>
</template>

<script>
import { Toast } from "mint-ui";
export default {
  data() {
    return {
      pageIndex: 1,
      comments: [],
      msg: "" //评论内容
    };
  },
  created() {
    this.getComments();
  },
  methods: {
    getComments() {
      this.$http
        .get("commentlist/" + this.id + "?pageindex=" + this.pageIndex)
        .then(result => {
          if (result.status === 200) {
            console.info(result.body);
            this.comments = this.comments.concat(result.body);
          } else {
            Toast("轮播图片请求失败");
          }
        });
    },

    getMore() {
      this.pageIndex++;
      this.getComments();
    },
    commit() {
      if (this.msg.trim().length === 0) {
        return Toast("评论内容不能为空");
      }

      this.$http
        .post("commit/" + this.$route.params.id, {
          content: this.msg.trim(),
          username: "saflyer",
          addtime: Date.now()
        })
        .then(function(result) {
          console.info(result);
          var cmt = {
            username: "saflyer",
            addtime: Date.now(),
            content: this.msg.trim()
          };

          if (result.status === 200) {
            this.comments.unshift(cmt);
            this.msg = "";
          } else {
            Toast("评论失败");
          }
        });
    }
  },
  props: ["id"]
};
</script> 


<style lang="scss" scoped>
.cmt-container {
  h3 {
    font-size: 18px;
  }
  textarea {
    font-size: 14px;
    height: 80px;
    margin: 0;
  }
  .cmt-list {
    margin-top: 5px;
    .cmt-item {
      font-size: 13px;
      .cmt-title {
        background-color: #cccccc;
        line-height: 30px;
      }
      .cmt-body {
        line-height: 35px;
        text-indent: 2em;
      }
    }
  }
}
</style>
