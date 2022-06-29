# bluebell

bluebell是一个基于gin框架和Vue框架搭建的前后端分离的web项目。

## 技能清单

1. gin框架：

   1. 使用gin框架作为底层框架启动

2. zap日志库：

   1. 使用zap日志库来替换gin框架原本log的记录，使用dev mode来去区分打印级别，在log配置中使用 mode模式区分

   2. ```go
      	if mode == "dev" {
      
      		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
      		core = zapcore.NewTee(
      			zapcore.NewCore(encoder, writeSyncer, l),
      			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
      		)
      	} else {
      
      		core = zapcore.NewCore(encoder, writeSyncer, l)
      
      	}
      ```

      

3. Viper配置管理

   1. viper库用yaml格式文件管理配置参数：

   2. ```yaml
      name: "web_app"
      mode: "dev"
      port: 8080
      version: "v0.1.4"
      start_time: "2020-07-01"
      machine_id: 1
      
      auth:
        jwt_expire: 8760
      
      log:
        level: "debug"
        filename: "web_app.log"
        max_size: 200
        max_age: 30
        max_backups: 7
      mysql:
        host: "127.0.0.1"
        port: 3306
        user: "root"
        password: "123456"
        dbname: "bluebell"
        max_open_conns: 200
        max_idle_conns: 50
      redis:
        host: "127.0.0.1"
        port: 6379
        password: ""
        db: 0
        pool_size: 100
      
      ```

      对启动中的参数可以通过配置文件的方式修改

4. swagger生成文档

5. JWT认证

   1. 使用jwt中间件来处理auth鉴权的处理

6. Go语言操作MySQL

7. Go语言操作Redis

