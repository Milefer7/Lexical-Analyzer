<!-- src/components/KeywordsTable.vue -->
<template>
    <div class="keywords-table">
        <h2>关键字</h2>
        <div class="form-group">
            <input v-model="newKeyword" placeholder="请输入关键字" />
            <button @click="addKeyword">添加</button>
        </div>
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
                    <td>{{ keyword.name }}</td>
                </tr>
            </tbody>
        </table>
        <div class="actions">
            <button @click="fetchKeywords">查询</button>
            <button @click="updateKeywords">修改</button>
        </div>
    </div>
</template>

<script>
import axios from 'axios'

export default {
    data() {
        return {
            keywords: [],
            newKeyword: ''
        }
    },
    methods: {
        async fetchKeywords() {
            try {
                const response = await axios.get('http://127.0.0.1:8080/keywords')
                this.keywords = response.data
            } catch (error) {
                console.error('Error fetching keywords:', error)
            }
        },
        async addKeyword() {
            try {
                const response = await axios.post('http://127.0.0.1:8080/keywords', {
                    name: this.newKeyword
                })
                this.keywords.push(response.data)
                this.newKeyword = ''
            } catch (error) {
                console.error('Error adding keyword:', error)
            }
        },
        async updateKeywords() {
            try {
                await axios.put('http://127.0.0.1:8080/keywords', this.keywords)
                alert('Keywords updated successfully')
            } catch (error) {
                console.error('Error updating keywords:', error)
            }
        }
    },
    created() {
        this.fetchKeywords()
    }
}
</script>

<style scoped>
.keywords-table {
    width: 400px;
    margin: auto;
}

.form-group {
    display: flex;
    margin-bottom: 10px;
}

input {
    flex: 1;
    padding: 5px;
}

button {
    margin-left: 10px;
}

table {
    width: 100%;
    border-collapse: collapse;
}

th,
td {
    border: 1px solid #ddd;
    padding: 8px;
}

thead {
    background-color: #f2f2f2;
}

.actions {
    margin-top: 10px;
}
</style>