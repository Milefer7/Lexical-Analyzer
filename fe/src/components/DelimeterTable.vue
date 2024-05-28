<!-- src/components/DelimitersTable.vue -->
<template>
    <div class="delimiters-table">
      <h2>单字符分隔符</h2>
      <div class="form-group">
        <input v-model="newDelimiter" placeholder="请输入单字符分隔符" />
        <button @click="addDelimiter">添加</button>
      </div>
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
            <td>{{ delimiter.symbol }}</td>
          </tr>
        </tbody>
      </table>
      <div class="actions">
        <button @click="fetchDelimiters">查询</button>
        <button @click="updateDelimiters">修改</button>
      </div>
    </div>
  </template>
  
  <script>
  import axios from 'axios'
  
  export default {
    data() {
      return {
        delimiters: [],
        newDelimiter: ''
      }
    },
    methods: {
      async fetchDelimiters() {
        try {
          const response = await axios.get('http://127.0.0.1:8080/delimiters')
          this.delimiters = response.data
        } catch (error) {
          console.error('Error fetching delimiters:', error)
        }
      },
      async addDelimiter() {
        try {
          const response = await axios.post('http://127.0.0.1:8080/delimiters', {
            symbol: this.newDelimiter
          })
          this.delimiters.push(response.data)
          this.newDelimiter = ''
        } catch (error) {
          console.error('Error adding delimiter:', error)
        }
      },
      async updateDelimiters() {
        try {
          await axios.put('http://127.0.0.1:8080/delimiters', this.delimiters)
          alert('Delimiters updated successfully')
        } catch (error) {
          console.error('Error updating delimiters:', error)
        }
      }
    },
    created() {
      this.fetchDelimiters()
    }
  }
  </script>
  
  <style scoped>
  .delimiters-table {
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
  th, td {
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
  