<template>
    <div class="shopcar-container">
        <div class="goods-list">
            

                <div class="mui-card" v-for="(item,i) in goodsList" :key="item.id">
                    <div class="mui-card-content">
                        <div class="mui-card-content-inner">
                            <mt-switch 
                            @change="selectedChanged(item.id,$store.getters.getGoodsSelected[item.id])"
                            v-model="$store.getters.getGoodsSelected[item.id]"></mt-switch>

                            <div class="info">
                                <h1>{{item.title}}</h1>  
                                <p>
                                    <span class="price">{{item.price}}</span>
                                    <numbox :goodsid="item.id"
                                    :initCount="$store.getters.getGoodsCount[item.id]"></numbox>
                                    <a href="#" @click.prevent="remove(item.id,i)">删除</a>
                                </p>
                            </div>
                        </div>
                    </div>
                </div>
        </div>
        
        

    <!-- 结算区域 -->
         <div class="mui-card">
            <div class="mui-card-content">
                <div class="mui-card-content-inner jiesuan">
                    <div class="left">
                        <p>总计不含运费</p>
                        <p>已勾选商品<span class="red">{{$store.getters.getGoodsCountAndAmount.count}}</span>件，
                        总价<span class="red">{{$store.getters.getGoodsCountAndAmount.amount}}</span></p>
                    </div>

                    <mt-button type="danger">去结算</mt-button>
                </div>
            </div>
    </div>

    <p>{{ $store.getters.getGoodsSelected }}</p>

    </div>
</template>

<script>
import numbox from "../subcomponets/ShopCarNumBox.vue";

export default {
  data() {
    return {
      goodsList: []
    };
  },
  created() {
    this.getGoodsList();
  },
  methods: {
    remove(item_id, index) {
      this.goodsList.splice(index, 1);
      this.$store.commit("removeCar", item_id);
    },

    getGoodsList() {
      var idArr = [];
      this.$store.state.car.forEach(item => idArr.push(item.id));
      if (idArr.length <= 0) {
        return;
      }
      //获取购物车列表
      this.$http.get("shoplist/" + idArr.join(",")).then(result => {
        if (result.status === 200) {
          console.info(result.body);
          this.goodsList = result.body;
        } else {
          Toast("轮播图片请求失败");
        }
      });
    },
    selectedChanged(id,val){
        //每当点击开关，同步最新的开关状态
        this.$store.commit("updataGoodsSelected",{id,selected:val})

    }
  },
  components: {
    numbox
  }
};
</script>

<style lang="scss" scoped>
.shopcar-container {
  background-color: #eeeeee;
  overflow: hidden;
  .goods-list {
    .mui-card-content-inner {
      display: flex;
      align-items: center;
    }
    img {
      width: 60px;
      height: 60px;
    }
    h1 {
      font-size: 14px;
    }
    .info {
      display: flex;
      flex-direction: column;
      justify-content: space-between;
      .price {
        color: red;
        font-weight: bold;
      }
    }
  }
  .jiesuan {
    display: flex;
    justify-content: space-between;
    align-items: center;
    .red {
      color: red;
      font-weight: bold;
      font-size: 16px;
    }
  }
}
</style>