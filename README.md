# itonchain 
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