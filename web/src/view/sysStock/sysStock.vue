<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="编号">
          <el-input v-model="searchInfo.stockNo" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="图书">
          <el-input v-model="searchInfo.bookId" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchInfo.status" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in StockStatusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="借阅者">
          <el-input v-model="searchInfo.userId" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="mini" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button size="mini" type="primary" icon="plus" @click="openDialog">新增</el-button>
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

        <el-table-column align="left" label="编号" prop="stockNo" />
        <el-table-column align="left" label="图书" prop="" >
          <template #default="scope">{{ scope.row.book.name }}</template>
          </el-table-column>
        <el-table-column align="left" label="状态" prop="status" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.status,"StockStatus") }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="借阅者" prop="userId" width="120">
          <template #default="scope">{{ scope.row.borrower.nickName }}</template>
          </el-table-column>
          <el-table-column align="left" label="创建人" prop="creatorId" width="120">
            <template #default="scope">{{ scope.row.creator.nickName }}</template>
          </el-table-column>
        <el-table-column align="left" label="备注" prop="remark" />
          <el-table-column align="left" label="借阅日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.return_at) }}</template>
          </el-table-column>
          <el-table-column align="left" label="创建日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
          </el-table-column>
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
              <el-button type="text" icon="edit" size="small" class="table-button" @click="updateSysStock(scope.row)">修改图书状态</el-button>
              <el-button type="text" icon="edit" size="small" class="table-button" @click="getBookLog(scope.row)">借书记录</el-button>
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="编号:">
          <el-input    :disabled="type!='create'" v-model="formData.stockNo" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="图书:" v-if="type =='create'" >
          <el-input v-model.number="formData.bookId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="状态:">
          <el-select v-model="formData.status" placeholder="请选择" style="width:100%" clearable>
            <el-option v-for="(item,key) in StockStatusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="借阅者:"  v-if="type =='create'" >
          <el-input v-model.number="formData.userId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="创建人:"  v-if="type =='create'" >
          <el-input v-model.number="formData.creatorId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="备注:" >
          <el-input v-model="formData.remark" clearable placeholder="请输入" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>


    <el-dialog v-model="dialogList" :before-close="closeDialogList" title="借书记录">
      <el-table
          ref="multipleTable"
          style="width: 100%"
          tooltip-effect="dark"
          :data="bookLog"
          row-key="ID"
      >
        <el-table-column align="left" label="图书" prop=""  width="200" >
          <template #default="scope">{{ scope.row.book.name }}</template>
        </el-table-column>
        <el-table-column align="left" label="状态" prop="status" width="80">
          <template #default="scope">
            {{ filterDict(scope.row.type,"type") }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="借阅者" prop="userId" width="120">
          <template #default="scope">{{ scope.row.borrower.nickName }}</template>
        </el-table-column>
        <el-table-column align="left" label="创建人" prop="creatorId" width="120">
          <template #default="scope">{{ scope.row.creator.nickName }}</template>
        </el-table-column>
        <el-table-column align="left" label="备注" prop="remark" />
        <el-table-column align="left" label="借阅日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
      </el-table>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" type="primary" @click="closeDialogList">确 定</el-button>
        </div>
      </template>
    </el-dialog>


  </div>
</template>

<script>
import {
  createSysStock,
  deleteSysStock,
  deleteSysStockByIds,
  updateSysStock,
  findSysStock,
  getSysStockList
} from '@/api/sysStock' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
import { findSysBookRentLog } from '../../api/sysBookRentLog'
export default {
  name: 'SysStock',
  mixins: [infoList],
   data() {
    return {
      listApi: getSysStockList,
      dialogFormVisible: false,
      bookLog:[],
      dialogList :false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      StockStatusOptions: [],
      formData: {
        stockNo: '',
        bookId: 0,
        status: undefined,
        userId: 0,
        creatorId: 0,
        remark: '',
      }
    }
  },
  async created() {
    await this.getDict('StockStatus')
    await this.getDict('type')
    await this.getTableData()
  },
  methods: {
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
        this.deleteSysStock(row)
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
      const res = await deleteSysStockByIds({ ids })
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
    async updateSysStock(row) {
      const res = await findSysStock({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        this.formData = res.data.resysStock
        this.dialogFormVisible = true
      }
    },
    async getBookLog(row) {
      const res = await findSysBookRentLog({ stockId: row.ID })
      if (res.code === 0) {
        this.bookLog = res.data.resysBookRentLog
        this.dialogList = true
      }
    },
    closeDialogList() {
      this.dialogList = false
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.formData = {
        stockNo: '',
        bookId: 0,
        status: undefined,
        userId: 0,
        creatorId: 0,
        remark: '',
      }
    },
    async deleteSysStock(row) {
      const res = await deleteSysStock({ ID: row.ID })
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
    async enterDialog() {
      let res
      switch (this.type) {
        case 'create':
          res = await createSysStock(this.formData)
          break
        case 'update':
          res = await updateSysStock(this.formData)
          break
        default:
          res = await createSysStock(this.formData)
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
    openDialog() {
      this.type = 'create'
      this.dialogFormVisible = true
    }
  },
}
</script>

<style>
</style>
