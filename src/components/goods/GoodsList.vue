<template>
    <div class="goods-list">
        <div class="goods-item">
            <img src="http://ofv795nmp.bkt.clouddn.com/vuelogobanner1.jpg" alt=""
            @click="goDetail(1)">
            <h1 class="title">小米手机开售了</h1>
            <div class="info">
                <p class="price">
                    <span class="now">￥999</span>
                    <span class="old">￥1099</span>
                </p>
                <p class="sell">
                    <span>热卖中</span>
                    <span>剩60件</span>
                </p>
            </div>
        </div>

        <div class="goods-item">
            <img src="http://ofv795nmp.bkt.clouddn.com/vuelogobanner1.jpg" alt=""
            @click="goDetail(2)">
            <h1 class="title">小米手机开售了小米手机开售了小米手机开开售了小米手机开开售了小米手机开开售了小米手机开售了小米手机开售了小米手机开售了</h1>
            <div class="info">
                <p class="price">
                    <span class="now">￥999</span>
                    <span class="old">￥1099</span>
                </p>
                <p class="sell">
                    <span>热卖中</span>
                    <span>剩60件</span>
                </p>
            </div>
        </div>


        
    </div>
</template>

<script>
    export default{
        data(){
            return {}
        },
        created() {
            
        },
        methods:{
            goDetail(id){
                console.info(this)
                this.$router.push('/home/goodsinfo/'+id)
            }
        }
    }
    
</script>

<style lang="scss" scoped>
.goods-list{
    display: flex;
    flex-wrap: wrap;
    padding: 7px;
    justify-content: space-between;

    .goods-item{
        width: 49%;
        border: 1px solid #ccc;
        box-shadow: 0 0 15px #ccc;
        margin: 3px 0;
        padding: 2px;
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        img{
            width: 100%;
        }
        .title{
            font-size: 14px;
        }
        .info{
            background-color: #eeeeee;
            p{
                margin: 0;
                padding: 5px;
            }
            .price{
                .now{
                    color: red;
                    font-weight: bold;
                    font-size: 16px;
                }
                .old{
                    text-decoration: line-through;
                }
            }
            .sell{
                display: flex;
                justify-content: space-between;
                font-size: 12px;
            }
        }
    }
}
</style>
