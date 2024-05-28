<template>
    <div class="container">
        <div class="title-bar">编译程序</div>
        <div class="main-content">
            <div class="panel">
                <h3>系统维护</h3>
                <button @click="openPopup('CharTable')">字符表</button>
                <button @click="openPopup('TokenClassification')">单词分类</button>
                <button @click="openPopup('SyntaxRules')">语法规则</button>
            </div>
            <div class="panel">
                <h3>词法分析</h3>
                <button @click="openPopup('KeywordTable')">关键字表</button>
                <button @click="openPopup('DelimeterTable')">单字符分界符</button>
                <button @click="openPopup('DFA')">词法分析DFA</button>
                <button @click="openPopup('ParsingProcess')">词法分析过程展示</button>
            </div>
        </div>
        <div class="status-bar">系统状态</div>
    </div>
</template>

<script>
export default {
    methods: {
        openPopup(component) {
            this.$router.push({ name: 'Popup', query: { component } });
        }
    }
}
</script>

<style>
body {
    font-family: Arial, sans-serif;
    margin: 0;
    padding: 0;
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    background-color: #f0f0f0;
}

.container {
    display: flex;
    flex-direction: column;
    width: 600px;
    background-color: white;
    border: 1px solid #ccc;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}

.title-bar {
    background-color: #2c3e50;
    color: white;
    padding: 10px;
    text-align: center;
    font-size: 18px;
    font-weight: bold;
}

.main-content {
    display: flex;
    justify-content: space-around;
    padding: 20px;
}

.panel {
    display: flex;
    flex-direction: column;
    align-items: center;
    width: 45%;
    border: 1px solid #ccc;
    padding: 10px;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}

.panel h3 {
    margin-bottom: 10px;
    font-size: 16px;
}

.panel button {
    margin: 5px 0;
    padding: 10px 20px;
    width: 100%;
    border: 1px solid #ccc;
    background-color: #f9f9f9;
    cursor: pointer;
    transition: background-color 0.3s;
}

.panel button:hover {
    background-color: #e9e9e9;
}

.status-bar {
    background-color: #ecf0f1;
    padding: 10px;
    text-align: center;
    font-size: 14px;
}
</style>
