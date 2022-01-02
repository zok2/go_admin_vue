<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="名称">
          <el-input v-model="searchInfo.name" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="mini" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="mini" type="primary" icon="plus" @click="openDialog(0)">新增</el-button>
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button size="mini" type="text" @click="deleteVisible = false">取消</el-button>
            <el-button size="mini" type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button icon="delete" size="mini" style="margin-left: 10px;" :disabled="!multipleSelection.length">删除</el-button>
          </template>
        </el-popover>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="编号" prop="ID" />
        <el-table-column align="left" label="名称" prop="name"  />
        <el-table-column align="left" label="创建日期" >
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="按钮组">
          <template #default="scope">
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updateSysBookType(scope.row)">变更</el-button>
            <el-button
                icon="plus"
                size="mini"
                type="text"
                @click="openDialog(scope.row.ID)"
            >添加子分类</el-button>
            <el-button type="text" icon="delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="添加分类">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="父级分类" prop="pid">
          <el-cascader
            v-model="formData.pid"
            style="width:100%"
            :disabled="type=='create'"
            :options="TypeOption"
            :props="{ checkStrictly: true,label:'name',value:'ID',disabled:'disabled',emitPath:false}"
            :show-all-levels="false"
            filterable
          />
        </el-form-item>
        <el-form-item label="名称:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
import {
  createSysBookType,
  deleteSysBookType,
  deleteSysBookTypeByIds,
  updateSysBookType,
  findSysBookType,
  getSysBookTypeList
} from '@/api/sysBookType' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'SysBookType',
  mixins: [infoList],
  data() {
    return {
      TypeOption: [
        {
          ID: 0,
          name: '根分类'
        }
      ],
      listApi: getSysBookTypeList,
      dialogFormVisible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        name: '',
        ID: '',
        pid: 0,
      }
    }
  },
  async created() {
    await this.getTableData()
  },
  methods: {
    onReset() {
      this.searchInfo = {}
    },
    // 条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10
      this.getTableData()
    },
    handleSelectionChange(val) {
      this.multipleSelection = val
    },
    deleteRow(row) {
      this.$confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.deleteSysBookType(row)
      })
    },
    setOptions() {
      this.TypeOption = [
        {
          ID: 0,
          name: '根分类'
        }
      ]
      this.setTypeOptions(this.tableData, this.TypeOption, false)
    },
    setTypeOptions(TypeData, optionsData, disabled) {
      this.formData.ID = String(this.formData.ID)
      TypeData &&
      TypeData.forEach(item => {
        if (item.children && item.children.length) {
          const option = {
            Id: item.ID,
            name: item.name,
            disabled: disabled || item.ID === this.formData.ID,
            children: []
          }
          this.setTypeOptions(
            item.children,
            option.children,
            disabled || item.ID === this.formData.ID,
          )
          optionsData.push(option)
        } else {
          const option = {
            ID: item.ID,
            name: item.name,
            disabled: disabled || item.ID === this.formData.ID
          }
          optionsData.push(option)
        }
      })
    },
    async onDelete() {
      const ids = []
      if (this.multipleSelection.length === 0) {
        this.$message({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      this.multipleSelection &&
        this.multipleSelection.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteSysBookTypeByIds({ ids })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length === ids.length && this.page > 1) {
          this.page--
        }
        this.deleteVisible = false
        this.getTableData()
      }
    },
    async updateSysBookType(row) {
      const res = await findSysBookType({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        this.formData = res.data.resysBookType
        this.dialogFormVisible = true
      }
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.formData = {
        name: '',
        pid: 0,
      }
    },
    async deleteSysBookType(row) {
      const res = await deleteSysBookType({ ID: row.ID })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length === 1 && this.page > 1) {
          this.page--
        }
        this.getTableData()
      }
    },
    // 初始化表单
    initForm() {
      if (this.$refs.authorityForm) {
        this.$refs.authorityForm.resetFields()
      }
      this.formData = {
        name: '',
        pid: '0'
      }
    },
    async enterDialog() {
      let res
      switch (this.type) {
        case 'create':
          res = await createSysBookType(this.formData)
          break
        case 'update':
          res = await updateSysBookType(this.formData)
          break
        default:
          res = await createSysBookType(this.formData)
          break
      }
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '创建/更改成功'
        })
        this.closeDialog()
        this.getTableData()
      }
    },
    openDialog(pid) {
      this.type = 'create'
      this.initForm()
      this.formData.pid = pid
      this.setOptions()
      this.dialogFormVisible = true
    }
  },
}
</script>

<style>
</style>
