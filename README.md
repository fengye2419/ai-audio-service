# ai-audio-service

## 简介
AI Audio Service 是一个提供 AI 音频处理功能的 Web 服务。

## 构建和运行

### 构建
使用以下命令构建项目：
```sh
make build
```

### 运行
构建完成后，可以使用以下命令运行服务：
```sh
./build/ai-audio-service --config ./configs/app.ini
```

### 配置
配置文件位于 configs/app.ini。你可以在此文件中修改服务的配置。

### 生成 Swagger 文档
使用以下命令生成 Swagger 文档：
```sh
make swagger
```

### 运行测试
使用以下命令运行测试：
```sh
make test
```

### 代码风格检查
使用以下命令运行 golangci-lint 进行代码风格检查：
```sh
make golangci-lint
```

### 帮助
使用以下命令查看可用的 Makefile 目标：
```sh
make help
```