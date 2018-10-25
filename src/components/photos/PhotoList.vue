<template>
    <div>

      <div id="slider" class="mui-slider">
				<div id="sliderSegmentedControl" class="mui-scroll-wrapper mui-slider-indicator mui-segmented-control mui-segmented-control-inverted">
					<div class="mui-scroll">
						<a v-for="cate in categorylist" :key="cate.id" @click="getPicList"
						:class="['mui-control-item',cate.id === 1? 'mui-active':'']">
							{{cate.title}}
						</a>
					</div>
				</div>
			</div>

			<ul class="photolist">
				<li v-for="item in piclist" :key="item">
					<img v-lazy="item">
				</li>
		  </ul>


    </div>
</template>

<script>
import mui from "../../lib/mui/js/mui.min.js";

export default {
  data() {
    return {
      categorylist: [],
      piclist: []
    };
  },
  created() {
    this.getAllCategory();
    this.getPicList();
  },
  mounted() {
    mui(".mui-scroll-wrapper").scroll({
      deceleration: 0.0005 //flick 减速系数，系数越大，滚动速度越慢，滚动距离越小，默认值0.0006
    });
  },
  methods: {
    getPicList() {
      this.$http.get("pingg").then(result => {
        console.info(result);
        if (result.status === 200) {
          this.piclist = result.body.lunbo;
        } else {
          Toast("图片列表获取失败");
        }
      });
    },
    getAllCategory() {
      this.$http.get("getimgcategory").then(result => {
        console.info(result);
        if (result.status === 200) {
          this.categorylist = result.body;
        } else {
          Toast("轮播图片请求失败");
        }
      });
    }
  }
};
</script>

<style lang="scss" scoped>
.photolist {
  list-style: none;
  margin: 0;
  padding: 10px;
  padding-bottom: 0;
  li {
    background-color: #cccccc;
    text-align: center;
    margin-bottom: 10px;
    box-shadow: 0 0 6px #999;
    img[lazy="loading"] {
      width: 40px;
      height: 300px;
      margin: auto;
    }
    img{
      width:100%;
      vertical-align: middle;
    }
  }
}

* {
  touch-action: pan-y;
}
</style>
