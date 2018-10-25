package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type User struct {
	NewsLogo string     `json:"newslogo"`
	NewsContent string  `json:"newscontent"`
	NewsTime      int64   `json:"newstime"`
	NewsClick int    `json:"newsclick"`
	Id int `json:"id"`
}

type Comment struct {
	UserName string     `json:"username"`
	AddTime int64  `json:"addtime"`
	Content      string   `json:"content"`
}

type Category struct {
	Title string `json:"title"`
	Id int `json:"id"`
}

type Good struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Price int `json:"price"`
	Path string `json:"path"`
}

func main() {
	r := gin.Default()
	r.Use(Cors())
	//vue项目轮播图
	r.GET("/pingg", func(c *gin.Context) {
		var lunbo []string = []string{"http://ofv795nmp.bkt.clouddn.com/vuelogobanner1.jpg",
			"http://ofv795nmp.bkt.clouddn.com/vuelogobanner2-1.jpg"}
		c.JSON(200, gin.H{
			"lunbo":&lunbo,
		})
	})

	//新闻列表
	r.GET("/newslist", func(c *gin.Context) {
		user1 := User{
			NewsLogo: "http://ofv795nmp.bkt.clouddn.com/vuelogobanner1.jpg",
			NewsContent: "vue 与 vue-resource 跨域问题解决",
			NewsTime:      time.Now().UnixNano() / 1e6,
			NewsClick: 23,
			Id:1,

		}

		user2 := User{
			NewsLogo: "http://ofv795nmp.bkt.clouddn.com/vuelogobanner2-1.jpg",
			NewsContent: "在vue项目下的 config/index.js 文件里面配置代理proxyTable",
			NewsTime:      time.Now().UnixNano() / 1e6,
			NewsClick: 343,
			Id:2,
		}

		var newlist []User = []User{user1,user2}

		c.JSON(http.StatusOK, newlist)
	})

	//获取评论
	r.GET("/commentlist/:name", func(c *gin.Context) {
		name:=c.Param("name")
		pageindex := c.Query("pageindex")
		fmt.Println("name:",name,",pageindex:",pageindex)

		var comments []Comment
		for i:=0 ;i<5 ;i++  {
			comment:=Comment{
				UserName:"safly",
				AddTime:time.Now().UnixNano() / 1e6,
				Content:"我是新添加的第"+strconv.Itoa(i)+"条评论",
			}
			comments = append(comments,comment)
		}

		c.JSON(http.StatusOK, comments)
	})

	//获取购物车数据

	r.GET("/shoplist/:ids", func(c *gin.Context) {
		ids := c.Param("ids")
		fmt.Println("ids:",ids)

		var idArr = strings.Split(ids,",")

		var shopcar []Good
		//添加第一个
		good:=Good{
			Id:idArr[0],
			Title: "产品一",
			Price:1000,
		}
		shopcar = append(shopcar,good)

		//添加第二个产品
		good=Good{
			Id:idArr[1],
			Title: "产品二",
			Price:2000,
		}
		shopcar = append(shopcar,good)

		c.JSON(http.StatusOK, shopcar)
	})

	//发表评论
	r.POST("/commit/:id", func(c *gin.Context) {
		id:=c.Param("id")
		fmt.Println("id",id)

		content := c.PostForm("content")
		username := c.PostForm("username")
		addtime := c.PostForm("addtime")
		fmt.Println("content>>>",content,"username>>>",username,"addtime",addtime)

		c.JSON(200, gin.H{
			"message": "ok",
		})
	})


	//获取图片分类数据
	r.GET("/getimgcategory", func(c *gin.Context) {
		var categorylist []Category
		var cats = [7]string{"推荐","热点","北京","社会","娱乐","NBA","新闻"}

		for i:=0;i<len(cats) ;i++ {
			cat := Category{
				Title: cats[i],
				Id:    i + 1,
			}
			categorylist = append(categorylist,cat)
		}

		c.JSON(http.StatusOK, categorylist)
	})


	r.GET("/getpiclist", func(c *gin.Context) {
		var categorylist []Category
		var cats = [7]string{"推荐","热点","北京","社会","娱乐","NBA","新闻"}

		for i:=0;i<len(cats) ;i++ {
			cat := Category{
				Title: cats[i],
				Id:    i + 1,
			}
			categorylist = append(categorylist,cat)
		}

		c.JSON(http.StatusOK, categorylist)
	})






	r.POST("/pingp", func(c *gin.Context) {
		name := c.PostForm("name")
		sex := c.PostForm("sex")
		fmt.Println("name>>>",name,"sex>>>",sex)

		c.JSON(200, gin.H{
			"name": name,
			"sex":sex,
		})
	})


	r.Run()
}


func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method      //请求方法
		origin := c.Request.Header.Get("Origin")        //请求头部
		var headerKeys []string                             // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")        // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")      //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//              允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")      // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")        // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")       //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")       // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next()        //  处理请求
	}
}
