import Vue from 'vue'

import VueRouter from 'vue-router'
Vue.use(VueRouter)

import Vuex from 'vuex'
Vue.use(Vuex)
//从本地存储读取数据
var car = JSON.parse(localStorage.getItem("car") || '[]')

var store = new Vuex.Store({
  state: {
    car: car //购物车数据

  },
  mutations: {
    addToCar(state, goodsinfo) {
      var flag = false

      state.car.some(item => {
        if (item.id == goodsinfo.id) {
          item.count +=  parseInt(goodsinfo.count)
          flag = true
          return true
        }
      })
      //没有找到就进行添加
      if(!flag){
        state.car.push(goodsinfo)
      }

      //保存到本地
      localStorage.setItem('car',JSON.stringify(state.car))

      console.info("stateCar",state.car)
    },
    updateCar(state,goodsinfo){
        //修改购物车数量值
        state.car.some(item =>{
            if(item.id == goodsinfo.id){
              item.count = parseInt(goodsinfo.count)
              return true
            }
        })

        localStorage.setItem('car',JSON.stringify(state.car))
    },
    removeCar(state,id){
      state.car.some((item,i)=>{
        if(item.id == id){
            state.car.splice(i,1)
            return true
        }
      })
      localStorage.setItem('car',JSON.stringify(state.car))
    },
    updataGoodsSelected(state,info){
        state.car.some(item=>{
          if(item.id == info.id){
              item.selected = info.selected
          }
        })
        localStorage.setItem('car',JSON.stringify(state.car))
    }
  },
  getters: {
      getAllCount(state){
          var c = 0;
          state.car.forEach(item=>{
              c += item.count
          })
          return c
      },
      getGoodsCount(state){
        var o = {}
        state.car.forEach(item =>{
          o[item.id] = item.count
        })
        return o
      },
      getGoodsSelected(state){
        var o = {}
        state.car.forEach(item =>{
            o[item.id] = item.selected
        })
        return o
      },
      getGoodsCountAndAmount(state){
          var o ={
            count:0,
            amount:0
          }
          state.car.forEach(item=>{
            console.info("it```",item)
              if(item.selected){
                o.count += item.count
                o.amount += item.price * item.count
              }
          })

          return o
      }
  }
})



import moment from 'moment'
//定义全局过滤器
Vue.filter("dateFormat", function (dataStr, pattern = "YYYY-MM-DD HH:mm:ss") {
  return moment(dataStr).format(pattern)
})


import router from './router.js'

import VueResource from 'vue-resource'
Vue.use(VueResource)
Vue.http.options.root = 'http://127.0.0.1:8080'
Vue.http.options.emulateJSON = true


// import { Header,Swipe, SwipeItem,Button,Lazyload } from 'mint-ui';
// Vue.component(Header.name, Header); 
// Vue.component(Swipe.name, Swipe);
// Vue.component(SwipeItem.name, SwipeItem);
// Vue.component(Button.name, Button);

// Vue.use(Lazyload);

import MintUI from 'mint-ui'
Vue.use(MintUI)
import 'mint-ui/lib/style.css'


// 安装 图片预览插件
import VuePreview from 'vue-preview'
Vue.use(VuePreview)


import './lib/mui/css/mui.min.css'
import './lib/mui/css/icons-extra.css'


import app from './App.vue'
var vm = new Vue({
  el: '#app',
  render: c => c(app),
  router,
  store
}) 