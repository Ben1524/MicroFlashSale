## 一、项目概述
MicroFlashSale 是一个包含多个子服务的综合性项目，旨在实现一个完整的微服务架构下的秒杀系统。该项目涵盖前端展示、后端服务以及数据生成等多个方面，具备产品管理、秒杀服务、用户服务等核心功能。

## 二、项目结构
项目主要由以下几个部分组成：
1. **vue_gin**：基于 Vue 和 Gin 的前端项目，负责用户界面的展示和交互。
2. **web**：Web 服务，处理前端请求并与各个后端服务进行交互。
3. **product_srv**：产品服务，负责产品数据的管理，包括产品信息的增删改查等操作。
4. **seckill_srv**：秒杀服务，处理秒杀业务逻辑，确保秒杀活动的公平性和高并发处理能力。
5. **user_srv**：用户服务，负责用户信息的管理和认证。
6. **data_generation.py**：数据生成脚本，用于生成测试数据。

## 三、前端项目（vue_gin）
### 1. 项目设置
在 `vue_gin` 目录下执行以下命令进行项目依赖安装：
```bash
npm install
```

### 2. 开发环境
在开发环境下进行代码的编译和热更新：
```bash
npm run serve
```

### 3. 生产环境
在生产环境下进行代码的编译和压缩：
```bash
npm run build
```

### 4. 代码检查和修复
对代码进行检查并修复一些常见问题：
```bash
npm run lint
```

### 5. 自定义配置
更多自定义配置请参考 [Configuration Reference](https://cli.vuejs.org/config/)。

## 四、Web 服务（web）
### 1. 生成 proto 代码
在 `web` 目录下，使用以下命令生成 proto 代码：
```bash
make proto
```

### 2. 运行服务
在 `web` 目录下，使用以下命令运行服务：
```bash
micro run .
```

## 五、产品服务（product_srv）
产品服务提供了编辑产品数据的功能。相关代码位于 `controller/product.go` 文件中，通过调用 `ProductDoEdit` 方法可以编辑产品数据。示例代码如下：
## 六、数据生成脚本（data_generation.py）
该脚本用于生成测试数据，具体使用方法可根据脚本内的注释进行操作。

## 七、依赖管理
项目中使用了多个第三方依赖，具体依赖版本信息可以参考以下文件：
- `MicroFlashSale/go.work.sum`
- `MicroFlashSale/web/go.sum`
- `MicroFlashSale/user_srv/go.sum`

## 八、注意事项
1. 请确保你已经安装了必要的开发环境和工具，如 Node.js、Go、Micro 等。
2. 在运行项目之前，需要配置好各个服务的环境变量和数据库连接信息。
3. 对于秒杀服务，需要考虑高并发情况下的性能和稳定性问题，可以进行相应的压力测试和优化。

## 九、贡献指南
如果你想为该项目做出贡献，请遵循以下步骤：
1. Fork 该项目到你的 GitHub 账户。
2. 创建一个新的分支进行开发：`git checkout -b feature/your-feature-name`
3. 提交你的代码并推送至你的分支：`git push origin feature/your-feature-name`
4. 创建一个 Pull Request，详细描述你的更改和贡献。

## 十、联系方式
如果你在使用过程中遇到任何问题或有任何建议，请通过以下方式联系我们：
- 邮箱：[wy17674094293@163.com]