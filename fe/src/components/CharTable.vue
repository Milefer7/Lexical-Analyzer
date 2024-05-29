<template>
    <div class="char-table">
        <h2>字符表</h2>
        <p class="note">注：英文字母区分大小写</p>
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
                        <input v-model="char.name" />
                    </td>
                </tr>
                <tr v-if="hiddenCharacters.length">
                    <td colspan="2">
                        <select v-model="selectedCharId" @change="handleSelectChange">
                            <option disabled value="">选择更多字符</option>
                            <option v-for="char in hiddenCharacters" :key="char.id" :value="char.id">
                                {{ char.key }}: {{ char.value }}
                            </option>
                        </select>
                    </td>
                </tr>
            </tbody>
        </table>
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
const visibleCharacters = computed(() => state.characters.slice(0, 10))
const hiddenCharacters = computed(() => state.characters.slice(79))

// Define methods
const fetchCharacters = async () => {
    try {
        const response = await axios.get('/alphabets')
        if (response.data.code === 1) {
            state.characters = response.data.data
            // 写入到框中展示
            console.log('Characters fetched successfully:', state.characters)
        } else {
            console.error('Error fetching characters:', response.data.msg)
        }
    } catch (error) {
        console.error('Error fetching characters:', error)
    }
}

const updateCharacters = async () => {
    const updates = state.characters.map(char => ({
        key: char.id,
        value: char.name
    }))
    try {
        const response = await axios.put('/alphabets', updates)
        console.log('Characters updated successfully:', response.data)
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
    padding: 20px;
    border: 1px solid #ccc;
    border-radius: 8px;
    width: 300px;
    margin: 0 auto;
    font-family: Arial, sans-serif;
}

h2 {
    margin-top: 0;
}

.note {
    font-size: 12px;
    color: #555;
    margin-bottom: 10px;
}

table {
    width: 100%;
    border-collapse: collapse;
    margin-bottom: 10px;
}

th,
td {
    border: 1px solid #ddd;
    padding: 8px;
    text-align: center;
}

input {
    width: 90%;
    padding: 4px;
    text-align: center;
    box-sizing: border-box;
}

select {
    width: 100%;
    padding: 4px;
    box-sizing: border-box;
}

.buttons {
    display: flex;
    justify-content: space-between;
}

button {
    padding: 6px 12px;
    border: none;
    border-radius: 4px;
    background-color: #007BFF;
    color: white;
    cursor: pointer;
}

button:hover {
    background-color: #0056b3;
}
</style>