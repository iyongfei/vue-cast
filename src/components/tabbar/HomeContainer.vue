<template>

<div>

<swiper :lunbotuList="lunbotuList" :isfull="true"></swiper>

<ul class="mui-table-view mui-grid-view mui-grid-9">
    <li class="mui-table-view-cell mui-media mui-col-xs-4 mui-col-sm-3">
        <router-link to="/home/newslist">
            <img src="../../images/menu1.png"/>
            <div class="mui-media-body">新闻资讯</div>
         </router-link>
    </li>
   <li class="mui-table-view-cell mui-media mui-col-xs-4 mui-col-sm-3">
        <router-link to="/home/photolist">
              <img src="../../images/menu2.png" alt="">
              <div class="mui-media-body">图片分享</div>
        </router-link>
    </li>

      <li class="mui-table-view-cell mui-media mui-col-xs-4 mui-col-sm-3">
        <router-link to="/home/goods">
            <img src="../../images/menu3.png" alt="">
            <div class="mui-media-body">商品购买
            </div>
        </router-link>
      </li>

      <li class="mui-table-view-cell mui-media mui-col-xs-4 mui-col-sm-3"><a href="#">
              <img src="../../images/menu4.png" alt="">
              <div class="mui-media-body">留言反馈</div></a></li>
      <li class="mui-table-view-cell mui-media mui-col-xs-4 mui-col-sm-3"><a href="#">
              <img src="../../images/menu5.png" alt="">
              <div class="mui-media-body">视频专区</div></a></li>
      <li class="mui-table-view-cell mui-media mui-col-xs-4 mui-col-sm-3"><a href="#">
              <img src="../../images/menu6.png" alt="">
              <div class="mui-media-body">联系我们</div></a></li>
</ul> 

</div>

</template>

<script>
import { Toast } from "mint-ui";
import swiper from '../subcomponets/Swiper.vue'

export default {
  data() {
    return {
      lunbotuList: []
    };
  },
  created() {
    this.getLunbo();
  },

  methods: {
    getLunbo() {
      //写成/pingg是加载不出来图片的
      this.$http.get("pingg").then(result => {
        console.info(result.body.lunbo);
        console.info(result.status);
        if (result.status === 200) {
          this.lunbotuList = result.body.lunbo;
        } else {
          Toast("轮播图片请求失败");
        }
      });
    }
  },
  components:{
    swiper
  }
};
</script>

<style lang="scss" scoped>
.mui-grid-view.mui-grid-9 .mui-table-view-cell {
  border: 0;
}
.mui-grid-view.mui-grid-9 {
  background-color: #fff;
  border: 0;
  img {
    width: 60px;
    height: 60px;
  }
}



</style>