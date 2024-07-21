# 词法分析器

## :book: Preview预览

* 主窗口

  ![image-20240531090344115](https://my-note-drawing-bed-1322822796.cos.ap-shanghai.myqcloud.com/picture/202405310903196.png)

* 词法分析过程展示

  ![image-20240531090303721](https://my-note-drawing-bed-1322822796.cos.ap-shanghai.myqcloud.com/picture/202405310903908.png)

## 项目结构

### 结构

```
├── README.md
├── be
│   ├── code
│   │   └── resp.go
│   ├── controller
│   │   └── controller.go
│   ├── dao
│   │   └── mysql
│   │       ├── alphabets.go
│   │       ├── delimiters.go
│   │       ├── init.go
│   │       ├── keywords.go
│   │       └── words.go
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   ├── model
│   │   ├── alphabets.go
│   │   ├── analysis.go
│   │   ├── delimiters.go
│   │   ├── keywords.go
│   │   └── words.go
│   ├── router
│   │   └── router.go
│   ├── service
│   │   ├── alphabets.go
│   │   ├── analysis.go
│   │   ├── delimiters.go
│   │   ├── keywords.go
│   │   ├── words.go
│   │   └── ws.go
│   └── utils
├── fe
│   ├── babel.config.js
│   ├── jsconfig.json
│   ├── package-lock.json
│   ├── package.json
│   ├── public
│   │   ├── favicon.ico
│   │   └── index.html
│   ├── src
│   │   ├── App.vue
│   │   ├── assets
│   │   │   └── dfa.png
│   │   ├── components
│   │   │   ├── ButtonPanel.vue
│   │   │   ├── CharTable.vue
│   │   │   ├── DFA.vue
│   │   │   ├── DelimeterTable.vue
│   │   │   ├── KeywordTable.vue
│   │   │   ├── ParsingProcess.vue
│   │   │   ├── SyntaxRules.vue
│   │   │   └── TokenClassification.vue
│   │   ├── main.js
│   │   ├── router.js
│   │   └── views
│   │       ├── HomeView.vue
│   │       └── PopupView.vue
│   └── vue.config.js
```

### 解释

- [README.md](./README.md) - 项目的README文件，包含项目的基本信息和使用说明
- be - 后端代码目录
  - [code](./be/code) - 
    - [resp.go](./be/code/resp.go) - 存放一些公用的状态码，以及响应处理
  - [controller](./be/controller) - 控制器目录，处理请求和响应
    - [controller.go](./be/controller/controller.go) - 控制器的主要代码
  - [dao](./be/dao) - 数据访问对象目录，处理数据库操作
    - mysql - MySQL相关的数据访问对象
      - [alphabets.go](./be/dao/mysql/alphabets.go) - 处理字母表相关的数据库操作
      - [delimiters.go](./be/dao/mysql/delimiters.go) - 处理分隔符相关的数据库操作
      - [init.go](./be/dao/mysql/init.go) - 初始化数据库连接
      - [keywords.go](./be/dao/mysql/keywords.go) - 处理关键字相关的数据库操作
      - [words.go](./be/dao/mysql/words.go) - 处理单词相关的数据库操作
  - [go.mod](./be/go.mod) - Go模块的依赖管理文件
  - [go.sum](./be/go.sum) - Go模块的依赖校验文件
  - [main.go](./be/main.go) - 后端服务的入口文件
  - [model](be/dal/model) - 数据模型目录，定义数据结构
    - [alphabets.go](be/dal/model/alphabets.go) - 字母表的数据模型
    - [analysis.go](be/dal/model/analysis.go) - 分析的数据模型
    - [delimiters.go](be/dal/model/delimiters.go) - 分隔符的数据模型
    - [keywords.go](be/dal/model/keywords.go) - 关键字的数据模型
    - [words.go](be/dal/model/words.go) - 单词的数据模型
  - [router](./be/router) - 路由目录，定义路由规则
    - [router.go](./be/router/router.go) - 路由的主要代码
  - [service](./be/service) - 服务目录，处理业务逻辑
    - [alphabets.go](./be/service/alphabets.go) - 字母表相关的业务逻辑
    - [analysis.go](./be/service/analysis.go) - 分析相关的业务逻辑
    - [delimiters.go](./be/service/delimiters.go) - 分隔符相关的业务逻辑
    - [keywords.go](./be/service/keywords.go) - 关键字相关的业务逻辑
    - [words.go](./be/service/words.go) - 单词相关的业务逻辑
    - [ws.go](./be/service/ws.go) - WebSocket相关的业务逻辑
  - [utils](./be/utils) - 工具函数目录
- fe - 前端代码目录
  - [babel.config.js](./fe/babel.config.js) - Babel的配置文件
  - [jsconfig.json](./fe/jsconfig.json) - JavaScript的配置文件
  - [package-lock.json](./fe/package-lock.json) - npm的依赖锁定文件
  - [package.json](./fe/package.json) - npm的依赖管理文件
  - public - 公共资源目录
    - [favicon.ico](./fe/public/favicon.ico) - 网站的图标
    - [index.html](./fe/public/index.html) - 网站的入口HTML文件
  - src - 源代码目录
    - [App.vue](./fe/src/App.vue) - Vue应用的主组件
    - assets - 静态资源目录
      - [dfa.png](./fe/src/assets/dfa.png) - DFA的图片资源
    - components - Vue组件目录
      - [ButtonPanel.vue](./fe/src/components/ButtonPanel.vue) - 按钮面板组件
      - [CharTable.vue](./fe/src/components/CharTable.vue) - 字符表组件
      - [DFA.vue](./fe/src/components/DFA.vue) - DFA组件
      - [DelimeterTable.vue](./fe/src/components/DelimeterTable.vue) - 分隔符表组件
      - [KeywordTable.vue](./fe/src/components/KeywordTable.vue) - 关键字表组件
      - [ParsingProcess.vue](./fe/src/components/ParsingProcess.vue) - 解析过程组件
      - [SyntaxRules.vue](./fe/src/components/SyntaxRules.vue) - 语法规则组件
      - [TokenClassification.vue](./fe/src/components/TokenClassification.vue) - 令牌分类组件
    - [main.js](./fe/src/main.js) - Vue应用的入口文件
    - [router.js](./fe/src/router.js) - Vue路由的配置文件
    - views - Vue视图目录
      - [HomeView.vue](./fe/src/views/HomeView.vue) - 主页视图
      - [PopupView.vue](./fe/src/views/PopupView.vue) - 弹出视图
  - [vue.config.js](./fe/vue.config.js) - Vue的配置文件

* ![image-20240531174922111](https://my-note-drawing-bed-1322822796.cos.ap-shanghai.myqcloud.com/picture/202405311749254.png)

## 准备

### 环境准备

* Go
* npm
* MySQL
* Linux

### 配置数据库

* 需在自行在根目录配置`.env`文件

  `DSN_LOCAL="加入你的数据库相关信息DSN"`

* 写入数据库示例

  > 相关功能未完善，需要用api测试工具手动写入数据库，统一用post方法，在body里用json格式传送

  * 关键词写入：http://127.0.0.1:8080/keywords

    ```json
    [
        { "id": 1, "keyword": "program" },
        { "id": 2, "keyword": "function" },
        { "id": 3, "keyword": "procedure" },
        { "id": 4, "keyword": "array" },
        { "id": 5, "keyword": "const" },
        { "id": 6, "keyword": "file" },
        { "id": 7, "keyword": "record" },
        { "id": 8, "keyword": "set" },
        { "id": 9, "keyword": "type" },
        { "id": 10, "keyword": "var" },
        { "id": 11, "keyword": "case" },
        { "id": 12, "keyword": "of" },
        { "id": 13, "keyword": "begin" },
        { "id": 14, "keyword": "end" },
        { "id": 15, "keyword": "do" },
        { "id": 16, "keyword": "if" },
        { "id": 17, "keyword": "else" },
        { "id": 18, "keyword": "for" },
        { "id": 19, "keyword": "repeat" },
        { "id": 20, "keyword": "then" },
        { "id": 21, "keyword": "while" },
        { "id": 22, "keyword": "with" },
        { "id": 23, "keyword": "string" },
        { "id": 24, "keyword": "integer" },
        { "id": 25, "keyword": "class" },
        { "id": 26, "keyword": "not" },
        { "id": 27, "keyword": "read" },
        { "id": 28, "keyword": "write" }
    ]
    ```

  * 分界符写入 http://127.0.0.1:8080/delimiters

    * ```json
      [
          { "id": 1, "name": "+" },
          { "id": 2, "name": "-" },
          { "id": 3, "name": "*" },
          { "id": 4, "name": "/" },
          { "id": 5, "name": "(" },
          { "id": 6, "name": ")" },
          { "id": 7, "name": "[" },
          { "id": 8, "name": "]" },
          { "id": 9, "name": ";" },
          { "id": 10, "name": "=" },
          { "id": 11, "name": "<" },
          { "id": 12, "name": ":" }
      ]
      ```

  * 字符表写入 http://127.0.0.1:8080

    * ```json
      [
          { "key": 1, "value": "a" },
          { "key": 2, "value": "b" },
          { "key": 3, "value": "c" },
          { "key": 4, "value": "d" },
          { "key": 5, "value": "e" },
          { "key": 6, "value": "f" },
          { "key": 7, "value": "g" },
          { "key": 8, "value": "h" },
          { "key": 9, "value": "i" },
          { "key": 10, "value": "j" },
          { "key": 11, "value": "k" },
          { "key": 12, "value": "l" },
          { "key": 13, "value": "m" },
          { "key": 14, "value": "n" },
          { "key": 15, "value": "o" },
          { "key": 16, "value": "p" },
          { "key": 17, "value": "q" },
          { "key": 18, "value": "r" },
          { "key": 19, "value": "s" },
          { "key": 20, "value": "t" },
          { "key": 21, "value": "u" },
          { "key": 22, "value": "v" },
          { "key": 23, "value": "w" },
          { "key": 24, "value": "x" },
          { "key": 25, "value": "y" },
          { "key": 26, "value": "z" },
          { "key": 27, "value": "A" },
          { "key": 28, "value": "B" },
          { "key": 29, "value": "C" },
          { "key": 30, "value": "D" },
          { "key": 31, "value": "E" },
          { "key": 32, "value": "F" },
          { "key": 33, "value": "G" },
          { "key": 34, "value": "H" },
          { "key": 35, "value": "I" },
          { "key": 36, "value": "J" },
          { "key": 37, "value": "K" },
          { "key": 38, "value": "L" },
          { "key": 39, "value": "M" },
          { "key": 40, "value": "N" },
          { "key": 41, "value": "O" },
          { "key": 42, "value": "P" },
          { "key": 43, "value": "Q" },
          { "key": 44, "value": "R" },
          { "key": 45, "value": "S" },
          { "key": 46, "value": "T" },
          { "key": 47, "value": "U" },
          { "key": 48, "value": "V" },
          { "key": 49, "value": "W" },
          { "key": 50, "value": "X" },
          { "key": 51, "value": "Y" },
          { "key": 52, "value": "Z" },
          { "key": 53, "value": "0" },
          { "key": 54, "value": "1" },
          { "key": 55, "value": "2" },
          { "key": 56, "value": "3" },
          { "key": 57, "value": "4" },
          { "key": 58, "value": "5" },
          { "key": 59, "value": "6" },
          { "key": 60, "value": "7" },
          { "key": 61, "value": "8" },
          { "key": 62, "value": "9" },
          { "key": 63, "value": "+" },
          { "key": 64, "value": "-" },
          { "key": 65, "value": "*" },
          { "key": 66, "value": "/" },
          { "key": 67, "value": "<" },
          { "key": 68, "value": "=" },
          { "key": 69, "value": "(" },
          { "key": 70, "value": ")" },
          { "key": 71, "value": "[" },
          { "key": 72, "value": "]" },
          { "key": 73, "value": "{" },
          { "key": 74, "value": "}" },
          { "key": 75, "value": ":" },
          { "key": 76, "value": "." },
          { "key": 77, "value": ";" },
          { "key": 78, "value": "'" },
          { "key": 79, "value": " " }
      ]
      ```

  * 单词写入 http://127.0.0.1:8080

    * ```json
      [
          { "id": 1, "name": "标识符", "word": "ID" },
          { "id": 2, "name": "保留字/关键字", "word": "标识符子集，内部表示" },
          { "id": 3, "name": "无符号整数", "word": "INTC" },
          { "id": 4, "name": "单字符分界符", "word": "一位长度界限符" },
          { "id": 5, "name": "双字符分界符", "word": ":=" },
          { "id": 6, "name": "注释头符", "word": "{" },
          { "id": 7, "name": "注释尾符", "word": "}" },
          { "id": 8, "name": "字符串起始、结束符", "word": "‘" },
          { "id": 9, "name": "数组下标界限符", "word": ".." }
      ]
      ```

## 运行

* 在根目录下执行 `start.sh`即可执行

## 贡献者

* [Milefer7](https://github.com/Milefer7)
* [tabzzz1](https://github.com/tabzzz1)

## 协议

* [MIT License](LICENSE).
