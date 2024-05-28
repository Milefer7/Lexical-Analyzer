<!-- src/components/TokenClassification.vue -->
<template>
    <div class="token-classification">
        <h2>单词分类</h2>
        <form @submit.prevent="updateTokens">
            <div v-for="(token, index) in tokens" :key="index" class="form-group">
                <label>{{ token.label }}</label>
                <input v-model="token.value" :placeholder="token.placeholder" />
            </div>
            <button type="submit">修改</button>
            <button type="button" @click="fetchTokens">取消</button>
        </form>
    </div>
</template>

<script>
import axios from 'axios'

export default {
    data() {
        return {
            tokens: [
                { label: '标识符', value: '', placeholder: 'ID' },
                { label: '保留字', value: '', placeholder: '标识符子集，内部表示' },
                { label: '无符号整数', value: '', placeholder: 'INTC' },
                { label: '单字符分界符', value: '', placeholder: '单字符界限符' },
                { label: '双字符分界符', value: '', placeholder: ':=' },
                { label: '注释头符', value: '', placeholder: '{' },
                { label: '注释尾符', value: '', placeholder: '}' },
                { label: '字符起始、结束符', value: '', placeholder: '\'' },
                { label: '数组下标界限符', value: '', placeholder: '..' }
            ]
        }
    },
    methods: {
        async fetchTokens() {
            try {
                const response = await axios.get('http://127.0.0.1:8080/words')
                const data = response.data
                this.tokens = this.tokens.map(token => ({
                    ...token,
                    value: data[token.placeholder.toLowerCase()]
                }))
            } catch (error) {
                console.error('Error fetching tokens:', error)
            }
        },
        async updateTokens() {
            try {
                const data = this.tokens.reduce((acc, token) => {
                    acc[token.placeholder.toLowerCase()] = token.value
                    return acc
                }, {})
                await axios.put('http://127.0.0.1:8080/words', data)
                alert('Tokens updated successfully')
            } catch (error) {
                console.error('Error updating tokens:', error)
            }
        }
    },
    created() {
        this.fetchTokens()
    }
}
</script>

<style scoped>
.token-classification {
    width: 400px;
    margin: auto;
}

.form-group {
    margin-bottom: 10px;
}

label {
    display: block;
    margin-bottom: 5px;
}

input {
    width: 100%;
    padding: 5px;
    box-sizing: border-box;
}

button {
    margin-top: 10px;
    margin-right: 10px;
}
</style>