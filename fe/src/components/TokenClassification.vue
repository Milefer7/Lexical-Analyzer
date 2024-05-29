<template>
    <div class="token-classification">
        <h2>单词分类</h2>
        <form @submit.prevent="updateTokens">
            <div v-for="(token) in tokens" :key="token.id" class="form-group">
                <label>{{ token.label }}</label>
                <input v-model="token.value" />
            </div>
            <div class="buttons">
                <button type="button" @click="fetchTokens">查询</button>
                <button type="submit">修改</button>
            </div>
        </form>
    </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import axios from 'axios'

export default {
    setup() {
        const tokens = ref([
            { id: 1, label: '标识符', value: '' },
            { id: 2, label: '保留字/关键字', value: '' },
            { id: 3, label: '无符号整数', value: '' },
            { id: 4, label: '单字符分界符', value: '' },
            { id: 5, label: '双字符分界符', value: '' },
            { id: 6, label: '注释头符', value: '' },
            { id: 7, label: '注释尾符', value: '' },
            { id: 8, label: '字符串起始、结束符', value: '' },
            { id: 9, label: '数组下标界限符', value: '' }
        ])

        const fetchTokens = async () => {
            try {
                const response = await axios.get('/words')
                const data = response.data.data

                console.log('Fetched data:', data) // 添加调试信息

                // 创建一个对象，键是 name 字段的值，值是 word 字段的值
                const wordMap = data.reduce((acc, item) => {
                    acc[item.name] = item.word
                    return acc
                }, {})

                console.log('Mapped word data:', wordMap) // 添加调试信息

                tokens.value = tokens.value.map(token => ({
                    ...token,
                    value: wordMap[token.label] || ''
                }))

                console.log('Updated tokens:', tokens.value) // 添加调试信息
            } catch (error) {
                console.error('Error fetching tokens:', error)
            }
        }

        const updateTokens = async () => {
            try {
                const data = tokens.value.map(token => ({
                    id: token.id,
                    name: token.label,
                    word: token.value
                }))
                await axios.put('/words', data)
                alert('单词成功更新！')
            } catch (error) {
                console.error('Error updating tokens:', error)
                alert('单词更新失败！ ' + error.response.data.msg)
            }
        }

        onMounted(fetchTokens)

        return {
            tokens,
            fetchTokens,
            updateTokens
        }
    }
}
</script>

<style scoped>
.token-classification {
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

.form-group {
    margin-bottom: 20px;
}

label {
    display: block;
    margin-bottom: 5px;
    font-weight: normal;
}

input {
    width: 100%;
    padding: 8px;
    text-align: left;
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

button[type="button"] {
    background-color: #6c757d;
}

button[type="button"]:hover {
    background-color: #5a6268;
}
</style>
