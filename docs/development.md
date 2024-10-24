# XCoder
XCoder 是一款旨在提高开发效率的工具，它能够根据用户的需求和输入，自动生成相应的代码片段或完整的代码文件。该插件旨在集成到主流的代码编辑器中，为开发者提供智能的代码辅助功能，帮助他们提高开发效率、减少错误，并提升代码质量。

## 项目配置
项目配置主要是项目的一些组件配置，mysql，mongodb 等，可以直接使用项目默认的配置，也可以换成自己的数据库。

### 使用项目默认配置
* 项目使用 docker-compose 进行部署，部署文件参考：[`manifest/deploy/docker-compose/docker-compose.yml`](../manifest/deploy/docker-compose/docker-compose.yml)，包括：mysql，mongodb，xcoder 等服务。    
* mysql，mongodb 在 `docker-compose.yml` 中设置了初始用户，密码及数据库名称，用于初始创建库表的配置。     
* 也支持修改这些数据库的默认配置，如果修改了，也需要在文件 [`manifest/config/config.yaml`](../manifest/config/config.yaml) 中同步修改，这个文件用于配置数据库的连接信息。  

### 使用自己的数据库
1. 直接在 [`manifest/config/config.yaml`](../manifest/config/config.yaml) 中对应的 mysql，mongodb 部分修改成自己的数据库连接即可。    
2. 然后需要执行 [`manifest/deploy/initdb.d/init_mysql.sql`](../manifest/deploy/initdb.d/init_mysql.sql) 脚本，对 mysql 数据库表进行初始化创建操作。    

## 项目部署
### 部署依赖
* docker
* docker-compose
* golang 1.18+

### 镜像构建
git clone 下载项目之后，执行以下命令行，即可构建项目镜像：
```shell
$ cd XCoder-Server

$ docker build -t xcoderai/xcoder:v0.1.0 .
```

当然也可以直接使用 [`manifest/deploy/docker-compose/docker-compose.yml`](../manifest/deploy/docker-compose/docker-compose.yml) 中，我们构建好的 xcoder 官方镜像。

### 项目参数配置
项目参数配置文件，位于 `manifest/config/` 目录，包括以下两个文件：
* [xcoder_code_generate_llm_params_conf.json](../manifest/config/xcoder_code_generate_llm_params_conf.json): 代码生成大模型参数相关配置文件
* [xcoder_chat_llm_params_conf.json](../manifest/config/xcoder_chat_llm_params_conf.json): 智能聊天（chat，单元测试，代码编辑等）模块大模型参数相关配置文件

**配置文件具体参数：**
#### xcoder_code_generate_llm_params_conf.json

```json
{
  "config": {
    "codellama": {
      "modelName": "codellama",
      "modelVersion": "codellama-13b-hf",
      "totalMaxTokens": 2500,
      "singleLineMaxTokens": 80,
      "multiLineMaxTokens": 160,
      "singeLineTemperature": 0.2,
      "multiLineTemperature": 0,
      "singleLineStopWords": ["\n", "\r\n"],
      "multiLineStopWords": ["\n\n", "\r\n\r\n"],
      "singleLineTopP": 1,
      "multiLineTopP": 1,
      "connUrls": []
    },
    "deepseeker": {
      "modelName": "deepseeker",
      "modelVersion": "deepseek-coder-7b-base",
      "totalMaxTokens": 2500,
      "singleLineMaxTokens": 80,
      "multiLineMaxTokens": 160,
      "singeLineTemperature": 0.2,
      "multiLineTemperature": 0,
      "singleLineStopWords": ["\n", "\r\n"],
      "multiLineStopWords": ["\n\n", "\r\n\r\n"],
      "singleLineTopP": 1,
      "multiLineTopP": 1,
      "connUrls": []
    }
  },
  "selectedModel": "codellama"
}
```
**参数解释：**
- **modelName**：string, llm 大模型的名称。
- **modelVersion**：string, llm 大模型的版本。此项需要跟后面本地大模型部署的 llm 名称保持一致。
- **singleLineMaxTokens**：int, 单行代码生成的最大 token 数。
- **multiLineMaxTokens**：int, 多行代码生成的最大 token 数。
- **singleLineStopWords**：[]string{}, 单行代码生成的停止符。
- **multiLineStopWords**：[]string{}, 多行代码生成的停止符。
- **connUrls**：[]string, 大模型的连接地址，支持配置多个。
- **selectedModel**：string, 设置使用的模型名称，例如：设置了 codellama，则使用 codellama 模型生成代码。

#### xcoder_chat_llm_params_conf.json

```json
{
  "config": {
    "llmType": "",
    "llmParams": {
      "apiBase": "",
      "apiKey": "",
      "apiVersion": "",
      "model": "",
      "temperature": 0.1,
      "maxTokens": 2500,
      "topP": 1,
      "frequencyPenalty": 0,
      "presencePenalty": 0
    }
  }
}
```
**参数解释：**
1. **llmType**: string，用于指定语言模型的类型。目前支持 openai，azure.
2. **llmParams**: map，包含用于配置语言模型的具体参数。
    - **apiBase**: string，llm 服务的访问地址。如果是 openai 类型，支持官方模型，末尾需要带上版本号。也支持自己本地部署的模型，这时需要自行确定末尾是否需要带上版本号。
    - **apiKey**: string，访问 llm 服务的 api key。
    - **apiVersion**: string，API的版本号，用于指定使用的API版本。如果是 openai 类型，此项可留空。
    - **model**: string，llm 模型名称。
    - **temperature**: float，控制生成文本的随机性。值越高，生成的文本越随机；值越低，生成的文本越确定。默认值是 `0.1`。
    - **maxTokens**: int，指定生成文本的最大长度，以令牌（tokens）为单位。默认值是 `2500`。
    - **topP**: float，用于控制核采样（nucleus sampling）的参数，影响生成文本的多样性。默认值是 `1`，意味着使用全部可能的令牌。
    - **frequencyPenalty**: int，用于减少重复内容的生成。值越高，重复内容出现的可能性越低。默认值是 `0`。
    - **presencePenalty**:int，用于增加生成文本中新内容的比例。值越高，新内容出现的可能性越高。默认值是 `0`。


### 项目启动
如果自己在上一步中，构建了镜像，则在 `docker-compose.yml` 中，替换成自己的 xcoder 镜像。然后执行以下命令进行项目启动：
```shell
$ cd manifest/deploy/docker-compose
$ docker-compose -f docker-compose.yml up
```


## 大模型部署
### 镜像构建
这里可以构建自己的 vllm 镜像，也可以直接使用官方的 vllm 镜像。如果要使用自己的，则使用以下命令进行构建：
```shell
$ cd llm_deploy/

$ docker build -t vllm/vllm:v1.0.0 .
```

### codellama
部署 codellama 代码大模型，相关配置信息都在 docker-compose.codellama.yaml 中配置好了，直接使用以下命令部署即可：
```shell
$ cd llm_deploy/
$ docker-compose -f docker-compose.codellama.yaml up -d
```

### deepseeker
部署 deepseeker 代码大模型，相关配置信息都在 docker-compose.deepseeker.yaml 中配置好了，直接使用以下命令部署即可：
```shell
$ cd llm_deploy/
$ docker-compose -f docker-composer.deepseeker.yaml up -d
```