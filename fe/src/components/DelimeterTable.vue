<template>
  <div class="delimiters-table">
    <h2>单字符分隔符</h2>
    <div class="table-container">
      <table>
        <thead>
          <tr>
            <th>ID</th>
            <th>单字符分隔符</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="delimiter in delimiters" :key="delimiter.id">
            <td>{{ delimiter.id }}</td>
            <td>
              <input v-model="delimiter.name" />
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    <div class="buttons">
      <button @click="fetchDelimiters">查询</button>
      <button @click="addDelimiter">添加</button>
      <button @click="updateDelimiters">修改</button>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

// Define reactive data
const delimiters = ref([])

// Define methods
const fetchDelimiters = async () => {
  try {
    const response = await axios.get('/delimiters')
    delimiters.value = response.data.data
    console.log('Delimiters fetched successfully:', delimiters.value)
  } catch (error) {
    console.error('Error fetching delimiters:', error)
    alert('获取单字符分隔符失败！' + error.response.data.msg)
  }
}

const addDelimiter = async () => {
  const ID = delimiters.value.length > 0 ? delimiters.value[delimiters.value.length - 1].id + 1 : 1
  delimiters.value.push({ id: ID, name: '' })
}

const updateDelimiters = async () => {
  try {
    await axios.put('/delimiters', delimiters.value)
    console.log('Delimiters updated successfully')
    alert('修改成功！')
  } catch (error) {
    console.error('Error updating delimiters:', error)
    alert('修改单字符分隔符失败！' + error.response.data.msg)
  }
}

// Fetch delimiters on component mount
onMounted(fetchDelimiters)
</script>

<style scoped>
.delimiters-table {
  padding: 30px;
  border: 1px solid #ccc;
  border-radius: 8px;
  width: 100%;
  max-width: 600px;
  margin: 0 auto;
  font-family: Arial, sans-serif;
  background-color: #f9f9f9;
}

h2 {
  margin-top: 0;
  font-size: 24px;
  color: #333;
}

.table-container {
  height: 400px;
  overflow-y: auto;
}

table {
  width: 100%;
  border-collapse: collapse;
  margin-bottom: 20px;
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

input {
  width: 100%;
  padding: 8px;
  text-align: center;
  box-sizing: border-box;
  border: 1px solid #ccc;
  border-radius: 4px;
}

.buttons {
  display: flex;
  justify-content: space-between;
  margin-top: 20px;
}

button {
  padding: 10px 16px;
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
</style>
