<template>
    <div class="char-table">
        <h2>字符表</h2>
        <p class="note">注：英文字母区分大小写</p>
        <div class="table-container">
            <table>
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>字符</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="char in visibleCharacters" :key="char.key">
                        <td>{{ char.key }}</td>
                        <td>
                            <input v-model="char.value" />
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
        <div v-if="hiddenCharacters.length" class="more-characters">
            <select v-model="selectedCharId" @change="handleSelectChange">
                <option disabled value="">选择更多字符</option>
                <option v-for="char in hiddenCharacters" :key="char.id" :value="char.id">
                    {{ char.key }}: {{ char.value }}
                </option>
            </select>
        </div>
        <div class="buttons">
            <button @click="fetchCharacters">查询</button>
            <button @click="updateCharacters">修改</button>
        </div>
    </div>
</template>

<script setup>
import { reactive, computed } from 'vue'
import axios from 'axios'

// Define reactive data
const state = reactive({
    characters: [],
    selectedCharId: ''
})

// Define computed properties
const visibleCharacters = computed(() => state.characters.slice(0, 79))
const hiddenCharacters = computed(() => state.characters.slice(79))

// Define methods
const fetchCharacters = async () => {
    try {
        const response = await axios.get('/alphabets')
        if (response.data.code === 1) {
            state.characters = response.data.data
            // console.log('Characters fetched successfully:', state.characters)
        } else {
            console.error('Error fetching characters:', response.data.msg)
        }
    } catch (error) {
        console.error('Error fetching characters:', error)
        alert('获取数据失败！', + error.response.data.msg)
    }
}

const updateCharacters = async () => {
    const updates = state.characters.map(char => ({
        key: char.key,
        value: char.value
    }))
    try {
        const response = await axios.put('/alphabets', updates)
        console.log('Characters updated successfully:', response.data)
        // 增加一个弹窗提示
        alert('修改成功！')
    } catch (error) {
        console.error('Error updating characters:', error)
    }
}

const handleSelectChange = () => {
    const selectedChar = hiddenCharacters.value.find(char => char.key === state.selectedCharId)
    if (selectedChar) {
        state.characters = state.characters.filter(char => char.key !== state.selectedCharId)
        state.characters.splice(10, 0, selectedChar)
        state.selectedCharId = ''
    }
}
</script>

<style scoped>
.char-table {
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

.note {
    font-size: 14px;
    color: #555;
    margin-bottom: 20px;
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

.more-characters {
    margin-bottom: 20px;
}

select {
    width: 100%;
    padding: 8px;
    box-sizing: border-box;
    border: 1px solid #ccc;
    border-radius: 4px;
    background-color: #fff;
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
