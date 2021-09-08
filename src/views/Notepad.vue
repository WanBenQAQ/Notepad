<template>
  <div class="notepad-card">
    <el-card class="box-card">
      <div slot="header">
        <el-input v-model="title" placeholder="请输入待办事项">
          <el-button slot="append" class="el-icon-circle-plus-outline" @click="addTodo"></el-button>
          <!--<el-button slot="append" icon="el-icon-search"></el-button>-->
        </el-input>
      </div>
      <el-table :data="todoList" height="250">
        <el-table-column type="index" label="#" width="140"></el-table-column>
        <el-table-column prop="title" label="待办事项" width="400"></el-table-column>
        <el-table-column label="操作">
          <template slot-scope="scope">
            <el-button v-if="scope.row.status" size="mini" type="success" icon="el-icon-check" circle @click="updateFalseTodo(scope.row.id)"></el-button>
            <el-button size="mini" v-if="!scope.row.status" type="info" class="el-icon-refresh-right" circle @click="updateTrueTodo(scope.row.id)"></el-button>
            <el-button size="mini" type="danger" icon="el-icon-delete" circle @click="deleteTodo(scope.row.id)"></el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script>
  export default {
    name: "Notepad",
    data() {
      return {
        todoList: [],
        title: ""
      }
    },
    created() {
      this.getTodoList()
    },
    methods: {
      getTodoList() {
        this.$http.get('api/v1/todo').then(res => {
          if(res.status !== 200) {
            return this.$message.error('获取用户列表失败')
          }
          this.todoList = res.data
        })
      },
      addTodo() {
        this.$http.post('api/v1/todo', {"title": this.title}).then(res => {
          if(res.status !== 200) {
            return this.$message.error("添加失败")
          }
          this.$message.success("添加成功")
          this.title = ''
          this.getTodoList()
        })
      },
      updateTrueTodo(id) {
        this.$http.put('api/v1/todo/'+id, {"status": true}).then(res => {
          if(res.status !== 200) {
            return this.$message.error("操作失败")
          }
          this.$message.success("操作成功")
          this.getTodoList()
        })
      },
      updateFalseTodo(id) {
        this.$http.put('api/v1/todo/'+id, {"status": false}).then(res => {
          if(res.status !== 200) {
            return this.$message.error("操作失败")
          }
          this.$message.success("操作成功")
          this.getTodoList()
        })
      },
      deleteTodo(id) {
        this.$http.delete('api/v1/todo/'+id).then(res => {
          if(res.status !== 200) {
            return this.$message.error("删除失败")
          }
          this.$message.success("删除成功")
          this.getTodoList()
        })
      }
    }
  }
</script>

<style scoped>
  .notepad-card {
    display:flex;
    justify-content: center;
    align-items: center;
  }
  .box-card {
    margin-top: 25px;
    width: 700px;
    text-align: center;
  }
</style>
