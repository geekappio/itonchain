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
 6. 启动Golang, 打开项目，路径xxx/itonchain/github.com/geekappio/itonchain为当前项目的路径  
 7. 设置项目的GOPATH为xxx/itonchain  
      
## 代码结构 
 >app             代码    
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