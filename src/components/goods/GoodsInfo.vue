<template>
    <div class="goodsinfo-container">
        <transition
            @before-enter="beforeEnter"
            @enter="enter"
            @after-enter="afterEnter">
            <div class="ball" v-show="ballFlag" ref="ball"></div>
        </transition>
      
        
        <div class="mui-card">
				<div class="mui-card-content">
					<div class="mui-card-content-inner">
                        <swiper :lunbotuList="lunbotuList" :isfull="false"></swiper>
					</div>
				</div>
		</div>
        
        <div class="mui-card">
				<div class="mui-card-header">页眉</div>
				<div class="mui-card-content">
					<div class="mui-card-content-inner">
                        <p class="price">
                            市场价：<del>￥232</del>&nbsp;&nbsp;
                            销售价：<span class="now_price">￥43534</span>
                        </p>
                        购买数量：<numbox @getCount="getSelectedCount" :max="goodsMax"></numbox>

                        <p>
                            <mt-button type="primary" size="small">立即购买</mt-button>
                            <mt-button type="danger" size="small" @click="addToShopCar">加入购物车</mt-button>
                            <!-- 分析： 如何实现加入购物车时候，拿到 选择的数量 -->
                            <!-- 1. 经过分析发现： 按钮属于 goodsinfo 页面， 数字 属于 numberbox 组件 -->
                            <!-- 2. 由于涉及到了父子组件的嵌套了，所以，无法直接在 goodsinfo 页面zhong 中获取到 选中的商品数量值-->
                            <!-- 3. 怎么解决这个问题：涉及到了 子组件向父组件传值了（事件调用机制） -->
                            <!-- 4. 事件调用的本质： 父向子传递方法，子调用这个方法， 同时把 数据当作参数 传递给这个方法 -->
                        </p>
					</div>
				</div>
		</div>

        <div class="mui-card">
				<div class="mui-card-header">页眉</div>
				<div class="mui-card-content">
					<div class="mui-card-content-inner">
                        <p>商品货号：</p>
                        <p>库存情况：</p>
                        <p>上架时间：</p>
					</div>
				</div>
				<div class="mui-card-footer">页脚</div>
		</div>
    </div>
</template>

<script>
import swiper from "../subcomponets/Swiper.vue";
import numbox from "../subcomponets/NumberBox.vue";
export default {
  data() {
    return {
      //路由参数挂在
      id: this.$route.params.id,
      lunbotuList: [],
      ballFlag: false,
      selectedCount: 1, //默认一个
      goodsMax: 10
    };
  },
  created() {
    this.getGoodsInfoLunbo();
  },
  methods: {
    getSelectedCount(count) {
      this.selectedCount = count;
      console.info("父组件拿到的数量", this.selectedCount);
    },
    addToShopCar() {
      //添加到购物车
      this.ballFlag = !this.ballFlag;

      var goodsinfo = { id: this.id, count: this.selectedCount,selected:true };
      this.$store.commit("addToCar",goodsinfo)
    },
    getGoodsInfoLunbo() {
      this.$http.get("pingg").then(result => {
        console.info(result.body.lunbo);
        console.info(result.status);
        if (result.status === 200) {
          this.goodsMax = 20;
          this.lunbotuList = result.body.lunbo;
        } else {
          Toast("轮播图片请求失败");
        }
      });
    },
    beforeEnter(el) {
      el.style.transform = "translate(0, 0)";
    },
    enter(el, done) {
      el.offsetWidth;

      const ballPosition = this.$refs.ball.getBoundingClientRect();
      const badgePosition = document
        .getElementById("badge")
        .getBoundingClientRect();

      const xDist = badgePosition.left - ballPosition.left;
      const yDist = badgePosition.top - ballPosition.top;

      el.style.transform = `translate(${xDist}px, ${yDist}px)`;
      el.style.transition = "all 0.5s cubic-bezier(.4,-0.3,1,.68)";
      done();
    },
    afterEnter(el) {
      this.ballFlag = !this.ballFlag;
    }
  },
  components: {
    swiper,
    numbox
  }
};
</script>

<style lang="scss" scoped>
.ball {
  width: 15px;
  height: 15px;
  border-radius: 50%;
  background-color: red;
  position: absolute;
  z-index: 99;
  top: 340px;
  left: 148px;
}
.now_price {
  color: red;
  font-size: 16px;
  font-weight: bold;
}
.goodsinfo-container {
  background-color: #eeeeee;
  overflow: hidden;
}
</style>
