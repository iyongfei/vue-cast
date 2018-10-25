<template>
    <div>
        <mt-swipe :auto="2000">
            <mt-swipe-item v-for="item in lunbotuList" :key="item">
                <img :src='item' alt="" :class="{'full':isfull}" />
            </mt-swipe-item>
        </mt-swipe>
    </div>
</template>

<script>
export default {
  props: ["lunbotuList","isfull"]
};
</script>

<style lang="scss" scoped>
.mint-swipe {
  height: 150px;

  .mint-swipe-item {
      text-align: center;
  }

  img {
    // width: 100%;
    height: 100%;
  }
  .full{
      width: 100%;
  }
}
</style>
