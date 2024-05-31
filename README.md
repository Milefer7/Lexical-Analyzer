# 词法分析器

## Preview预览

* 主窗口

  ![image-20240531090344115](https://my-note-drawing-bed-1322822796.cos.ap-shanghai.myqcloud.com/picture/202405310903196.png)

* 词法分析过程展示

  ![image-20240531090303721](https://my-note-drawing-bed-1322822796.cos.ap-shanghai.myqcloud.com/picture/202405310903908.png)

## 项目结构

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

## 准备环境

* Go
* npm
* MySQL
* Linux

## 运行

* 在根目录下执行 `start.sh`即可执行

## 贡献者

* [Milefer7](https://github.com/Milefer7)
* [tabzzz1](https://github.com/tabzzz1)

## 协议

* [MIT License](LICENSE).
