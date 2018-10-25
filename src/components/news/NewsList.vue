<template>
    <div>
        <h3>新闻列表页面</h3>
        <ul class="mui-table-view">
            <li class="mui-table-view-cell mui-media" v-for="item in newsList" :key="item.newslogo">
                <router-link :to="'/home/newsinfo/' + item.id">
                <img class="mui-media-object mui-pull-left" :src="item.newslogo">
                <div class="mui-media-body">
                    <h1>{{ item.newscontent }}</h1>
                    <p class='mui-ellipsis'>
                    <span>发表时间：{{ item.newstime | dateFormat}}</span>
                    <span>点击：{{item.newsclick}}次</span>
                    </p>
                </div>
                </router-link>
            </li>
        </ul>

        
    </div>
</template>
<script>
import { Toast } from "mint-ui";

export default {
  data() {
    return {
      newsList: []
    };
  },
  created() {
    this.getNewsList();
  },

  methods: {
    getNewsList() {
      //写成/pingg是加载不出来图片的
      this.$http.get("newslist").then(result => {
          console.info(result)
        if (result.status === 200) {
            this.newsList = result.body 
        } else {
          Toast("轮播图片请求失败");
        }
      });
    }
  }
};
</script>

<style lang="scss" scoped>
.mui-table-view {
  li {
    h1 {
      font-size: 14px;
    }
    .mui-ellipsis {
      font-size: 12px;
      color: #226aff;
      display: flex;
      justify-content: space-between;
    }
  }
}
</style>
