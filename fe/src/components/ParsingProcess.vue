<template>
  <div class="parsing-process">
    <div class="input-section">
      <textarea v-model="sourceCode" placeholder="请输入源代码"></textarea>
      <button @click="importDefaultCode">导入源代码</button>
    </div>
    <div class="analysis-section">
      <img src="../assets/dfa.png" alt="词法分析图" />
      <button @click="performLexicalAnalysis">词法分析</button>
    </div>
    <div class="result-section">
      <div class="table-container">
        <table>
          <thead class="table-header">
            <tr>
              <th>行号</th>
              <th>单词内容</th>
              <th>单词种类</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="result in analysisResults" :key="result.lineNum + result.content">
              <td>{{ result.lineNum }}</td>
              <td>{{ result.content }}</td>
              <td>{{ result.type }}</td>
            </tr>
          </tbody>
        </table>
      </div>
      <button class="export-button" @click="exportWords">导出单词</button>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import axios from 'axios'

// 初始化数据
const sourceCode = ref('')
const analysisResults = ref([])

// 默认代码
const defaultCode = `
program p
var i : integer ;
  x : array [1..16] of integer ;
begin
  i := 5 + 10 ;
  x[i] := 0 ;
end .
`

// 导入默认源代码
const importDefaultCode = () => {
  sourceCode.value = defaultCode
}

// 词法分析
const performLexicalAnalysis = async () => {
  try {
    const response = await axios.post('/analysis/lexical', {
      code: sourceCode.value
    })
    if (response.data.code === 1) {
      analysisResults.value = response.data.data
      console.log('Lexical analysis successful:', analysisResults.value)
    } else {
      alert('词法分析失败！' + response.data.msg)
    }
  } catch (error) {
    console.error('Error during lexical analysis:', error)
    alert('词法分析失败！' + error.response.data.msg)
  }
}

// 导出单词
const exportWords = () => {
  const csvContent = "data:text/csv;charset=utf-8,"
    + "行号,单词内容,单词种类\n"
    + analysisResults.value.map(e => `${e.lineNum},${e.content},${e.type}`).join("\n")
  const encodedUri = encodeURI(csvContent)
  const link = document.createElement("a")
  link.setAttribute("href", encodedUri)
  link.setAttribute("download", "lexical_analysis_results.csv")
  document.body.appendChild(link)
  link.click()
  document.body.removeChild(link)
}
</script>

<style scoped>
.parsing-process {
  display: flex;
  justify-content: space-between;
  padding: 30px;
  border: 1px solid #ccc;
  border-radius: 8px;
  background-color: #f9f9f9;
  font-family: Arial, sans-serif;
  width: 100%;
  max-width: 1200px;
  margin: 0 auto;
}

.input-section,
.analysis-section,
.result-section {
  width: 30%;
  display: flex;
  flex-direction: column;
  align-items: center;
}

textarea {
  width: 100%;
  height: 520px;
  padding: 10px;
  box-sizing: border-box;
  border: 1px solid #ccc;
  border-radius: 4px;
}

button {
  width: 100%;
  padding: 10px 0;
  margin-top: 10px;
  border: none;
  border-radius: 4px;
  background-color: #007bff;
  color: white;
  cursor: pointer;
  font-size: 16px;
  transition: background-color 0.3s ease;
}

button:hover {
  background-color: #0056b3;
}

.export-button {
  margin-top: 25px;
}

.analysis-section img {
  width: 100%;
  height: auto;
  border: 1px solid #ccc;
  border-radius: 4px;
  margin-bottom: 10px;
}

.table-container {
  width: 100%;
  height: 510px;
  /* 限制表格的高度 */
  overflow-y: auto;
  /* 超出高度时显示滚动条 */
}

.table-header {
  position: sticky;
  top: 0;
  background-color: white;
  /* 防止滚动时内容透过表头可见 */
  z-index: 10;
  /* 确保表头在内容之上 */
}

table {
  width: 100%;
  border-collapse: collapse;
}

th,
td {
  border: 1px solid #ddd;
  padding: 12px;
  text-align: center;
  background-color: #fff;
}

th {
  background-color: #f1f1f1;
}
</style>
