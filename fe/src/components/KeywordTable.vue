<template>
    <div class="keywords-table">
        <h2>关键字</h2>
        <div class="table-container">
            <table>
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>关键字</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="keyword in keywords" :key="keyword.id">
                        <td>{{ keyword.id }}</td>
                        <td>
                            <input v-model="keyword.keyword" />
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
        <div class="buttons">
            <button @click="fetchKeywords">查询</button>
            <button @click="addKeyword">添加</button>
            <button @click="updateKeywords">修改</button>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

// Define reactive data
const keywords = ref([])

// Define methods
const fetchKeywords = async () => {
    try {
        const response = await axios.get('/keywords')
        keywords.value = response.data.data
        // console.log('Keywords fetched successfully:', keywords.value)
        alert('获取关键字成功！')
    } catch (error) {
        console.error('Error fetching keywords:', error)
        alert('获取关键字失败！' + error.response.data.msg)
    }
}

const addKeyword = () => {
    const id = keywords.value.length > 0 ? keywords.value[keywords.value.length - 1].id + 1 : 1
    keywords.value.push({ id, name: '' })
}

const updateKeywords = async () => {
    try {
        await axios.put('/keywords', keywords.value)
        console.log('Keywords updated successfully')
        alert('修改成功！')
    } catch (error) {
        console.error('Error updating keywords:', error)
        alert('修改关键字失败！' + error.response.data.msg)
    }
}

// Fetch keywords on component mount
onMounted(fetchKeywords)
</script>

<style scoped>
.keywords-table {
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
