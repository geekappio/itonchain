# itonchain 
## 项目说明
基于微信小程序和公众平台的知识分享解答平台

功能： 采用币乎的模式，发行基于区块链的key，作为文章奖励。

通过微信注册用户，用户注册时选择自己感兴趣的文章类型。

用户可分类收藏各类微信上的文章。

用户可搜索知乎和csdn上的文章，后台转换格式后显示给用户，并提供微信内分享的功能。

用户可提问，其他用户可回答。

PC端功能类似知乎，可以编辑文章与微信端打通。

提供插件功能，可自定义开发自己的插件，比如备忘录插件、日程管理插件、课程表插件等等。

提高分类技术论坛、公众号的类目查询

提供各种会议、活动的信息发布

提供问卷调查的功能

创业企业黄页

区块链企业名录和介绍

一期：（3月17日-4月15日）

用户注册
文章分类管理、收藏
各个渠道的文章获取
通过公众号打通用户与小程序
计划：

设计： 17-19
评审： 20
开发： 3月21日 - 4月15日
技术选型： 微信内：小程序语言 服务端：Go 数据库：MySql 资源存储：https://github.com/chrislusf/seaweedfs

1. 文章列表项布局、详情式样调整、详情页面内功能、目录管理
2. 后台文章管理页面（文章列表、查询、编辑、发布、标签、上架、下架）
3. RSS订阅源自动处理，需要增加一个表

二期：

编辑、分享
搜集用户习惯，推送文章
三期：

增加实用小功能

## 如何创建开发环境 
1. 在本机安装Go环境    
2. 设置好本机GOROOT和默认的GOPATH路径    
3. 像下面一样在本地创建itonchain项目环境    
    >mkdir itonchain    
    cd itonchain  
    mkdir bin pkg src  
    cd src  
    mkdir -p github.com/geekappio/itonchain  
 4. 从github下拉代码    
    > cd github.com/geekappio/itonchain    
    git clone https://github.com/geekappio/itonchain 
    
    **注意：这里一定要在src/github.com/geekappio/itonchain目录下拉代码**  
 5. 设置GOBIN=xxx/itonchain/bin环境变量, 指定go insall命令的默认执行文件的位置（可选）  
 6. 启动Goland, 打开项目，路径xxx/itonchain/github.com/geekappio/itonchain为当前项目的路径  
 7. 设置项目的GOPATH为xxx/itonchain  
      
## 代码结构 
app             代码    
  common          公用代码    
    config          系统配置    
    dao             数据源管理，数据库操作    
    resource        配置和资源文件    
    service         业务层，处理业务    
    sql             数据库脚本    
    util            工具类    
    vender          第三方依赖项目    
    web             web层    
   reousrce        资源配置    
sql             数据库模板    
tmp             临时目录

## 开发约定
### 数据操作ORM 
 1. 使用xormplus作为数据库orm类库，使用stpl模板
 2. 默认模板都放置在项目的/resource/xorm目录下，每个Table创建一个子目录，每个查询语句创建一个模板文件


## 开发
### Goland安装远程发布插件
参考：https://blog.csdn.net/weixin_41571449/article/details/78957144

### 如何Debug
 1. 使用geekappio用户ssh登录 geekapp.itonchina.com
 2. 进入目录/home/geekappio/go/src/github.com/geekappio/itonchain
 3. 运行命令：go build -o itonchain -gcflags='-N -l' github.com/geekappio/itonchain/app && dlv --listen=:9000 --headless=true --api-version=2 exec ./itonchain

### 访问SeaWeedFS存储的文件
 1. 通过Nginx代理对外暴露内部的SeaWeedFS服务
 2. 对外暴露的URL路径为
    文件：https://geekapp.itonchain.com/itonchain/resource/article/[fid]
    图片：https://geekapp.itonchain.com/itonchain/resource/image/[fid]
