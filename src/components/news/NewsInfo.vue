<template>
<div class="newsinfo-container">
  <h3 class="title">新闻详情 {{id}} </h3>

  <p class="subtitle">
      <span>发表时间</span>
      <span>点击：0次</span>
  </p>

  <hr>

  <div class="content"></div>

  <!-- 评论区域 -->
  <comment-box :id="this.id"></comment-box>


</div>
  
   
</template>

<script>
    import comment from '../subcomponets/component.vue'

    export default{
        data(){
            return{
                id:this.$route.params.id
            }
        },
        created() {
            
        },
        methods:{

        },
       components:{
           'comment-box':comment
       }
    }
</script>

<style lang="scss" scoped>
    .newsinfo-container{
        padding: 0 4px;
        .title{
            font-size: 16px;
            text-align: center;
            margin: 15px 0;
            color: red($color: #000000)
        } 
        .subtitle{
            font-size: 13px;
            color: #2262ff;
            display: flex;
            justify-content: space-between;

        } 
        .content{
            img{
                width: 100%;
            }
        }

    }


</style>
