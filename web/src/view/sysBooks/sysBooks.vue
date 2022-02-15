<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="名称">
          <el-input v-model="searchInfo.name" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="图书编码">
          <el-input v-model="searchInfo.upc" placeholder="搜索条件" />
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
        <el-table-column align="left" label="名称" prop="name" />
        <el-table-column align="left" label="作者" prop="author" />
        <el-table-column align="left" label="出版社" prop="press" />
        <el-table-column align="left" label="出版时间" prop="pubDate" />
        <el-table-column align="left" label="图书编码" prop="upc" />
        <el-table-column align="left" label="书架编号" prop="bookcaseId" />
        <el-table-column align="left" label="价格" prop="price" />
        <el-table-column align="left" label="类型" prop="type.name" />
        <el-table-column align="left" label="封面" prop="photo" />
        <el-table-column align="left" label="状态" prop="status"><template #default="scope">{{ filterDict(scope.row.status,"status") }}</template></el-table-column>
        <el-table-column align="left" label="可借/总数量" prop="amount"><template #default="scope">{{ filterNum (scope.row.amount, scope.row.stock) }}</template></el-table-column><el-table-column align="left" label="日期" width="180"><template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template></el-table-column>
        <el-table-column align="left" label="按钮组">
          <template #default="scope">
            <el-button type="text" icon="edit" size="small" class="table-button" @click="borrowSysBooks(scope.row)">借书</el-button>
            <el-button type="text" icon="edit" size="small" class="table-button" @click="alsoSysBooks(scope.row)">还书</el-button>
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updateSysBooks(scope.row)">变更</el-button>
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="名称:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="作者:">
          <el-input v-model="formData.author" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="出版社:">
          <el-input v-model="formData.press" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="出版时间:">
          <el-date-picker v-model="formData.pubDate" type="date" style="width:100%" placeholder="选择日期" clearable />
        </el-form-item>
        <el-form-item label="图书编码:">
          <el-input v-model="formData.upc" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="书架编号:">
          <el-input v-model="formData.bookcaseId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="价格:">
          <el-input-number v-model="formData.price" style="width:100%" :precision="2" clearable />
        </el-form-item>
        <el-form-item label="类型" prop="typeId">
          <el-cascader
            ref="cascader"
            v-model="formData.typeId"
            style="width:100%"
            :options="typeOptions"
            :show-all-levels="false"
            :props="{ multiple:false,checkStrictly: false,label:'name',value:'ID',disabled:'disabled',emitPath:false}"
            :clearable="false"
          />
        </el-form-item>
        <el-form-item label="封面:">
          <el-input v-model="formData.photo" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="状态:">
          <el-select v-model="formData.status" placeholder="请选择" style="width:100%" clearable>
            <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="数量:">
          <el-input v-model.number="formData.amount" :disabled="type === `update`" clearable placeholder="请输入" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
    <el-dialog v-model="borrowFormVisible" :before-close="closeDialog" :title=" `${(type=='borrow')?'借-':'还-'}`+`${borrowData.name}`">
      <el-form :model="borrowData" label-position="right" label-width="80px">
        <el-form-item label="名称:">
          <el-input v-model="borrowData.name" :disabled="true" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="作者:">
          <el-input v-model="borrowData.author" :disabled="true" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="出版社:">
          <el-input v-model="borrowData.press" :disabled="true" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="出版时间:">
          <el-date-picker v-model="borrowData.pubDate" :disabled="true" type="date" style="width:100%" placeholder="选择日期" clearable />
        </el-form-item>
        <el-form-item label="图书编码:">
          <el-input v-model="borrowData.upc" :disabled="true" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="库存编码:">
          <el-select
            v-model="borrowData.stockId"
            clearable
            filterable
            :filter-method="stockFilter"
            @change="handleChange"
            @clear="handleClear"
          >
            <el-option v-for="stock in stockList" :key="stock.ID" :label="stock.stockNo" :value="stock.ID" />
          </el-select>
        </el-form-item>
        <el-form-item label="借书人:">
          <el-select
            v-model="borrowData.userId"
            :disabled="type!='borrow'"
            clearable
            filterable
            :filter-method="userFilter"
            @change="handleUserChange"
            @clear="handleClear"
          >
            <el-option v-for="user in userList" :key="user.ID" :label="user.nickName" :value="user.ID" />
          </el-select>
        </el-form-item>
        <el-form-item label="备注:">
          <el-input v-model="borrowData.remark" type="textarea" :autosize="{ minRows: 4, maxRows: 20}" placeholder="请输入内容" class="filter-item" clearable />
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
  createSysBooks,
  deleteSysBooks,
  deleteSysBooksByIds,
  updateSysBooks,
  borrowSysBooks,
  alsoSysBooks,
  findSysBooks,
  getSysBooksList
} from '@/api/sysBooks' //  此处请自行替换地址
import { findUsers } from '../../api/user'
import infoList from '@/mixins/infoList'
import { getSysBookTypeList } from '../../api/sysBookType'
export default {
  name: 'SysBooks',
  mixins: [infoList],
  model: {
    prop: 'stockValue',
    event: 'change'
  },
  data() {
    return {
      listApi: getSysBooksList,
      dialogFormVisible: false,
      borrowFormVisible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      statusOptions: [],
      typeOptions: [],
      stockList: [],
      copiedStockList: [],
      userList: [],
      copiedUserList: [],
      formData: {
        name: '',
        author: '',
        press: '',
        pubDate: new Date(),
        upc: '',
        bookcaseId: '',
        price: 0,
        typeId: 0,
        photo: '',
        status: undefined,
        amount: 0,
      },
      borrowData: {
        name: '',
        author: '',
        press: '',
        pubDate: new Date(),
        upc: '',
        bookcaseId: '',
        price: 0,
        typeId: 0,
        photo: '',
        status: undefined,
        amount: 0,
        stockId: '',
        userId: '',
      }
    }
  },
  async created() {
    await this.getTableData()
    await this.getDict('status')
    const res = await getSysBookTypeList({ page: 1, pageSize: 999 })
    this.typeOptions = []
    this.setTypeOptions(res.data.list, this.typeOptions)
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
        this.deleteSysBooks(row)
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
      const res = await deleteSysBooksByIds({ ids })
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
    async updateSysBooks(row) {
      const res = await findSysBooks({ ID: row.ID })
      this.type = 'update'
      if (res.code === 0) {
        this.formData = res.data.resysBooks
        this.dialogFormVisible = true
      }
    },
    async alsoSysBooks(row) {
      const res = await findSysBooks({ ID: row.ID, StockStatus: 2 })
      this.type = 'also'
      if (res.code === 0) {
        this.borrowData = res.data.resysBooks
        this.stockList = this.borrowData.stock
        this.copiedStockList = this.borrowData.stock
        this.borrowFormVisible = true
        const userRes = await findUsers({})
        if (userRes.code === 0) {
          this.userList = userRes.data
          this.copiedUserList = userRes.data
          this.borrowFormVisible = true
        }
      }
    },
    async borrowSysBooks(row) {
      const res = await findSysBooks({ ID: row.ID, StockStatus: 3})
      this.type = 'borrow'
      if (res.code === 0) {
        if (res.data.resysBooks.stock.length && res.data.resysBooks.status) {
          this.borrowData = res.data.resysBooks
          this.stockList = this.borrowData.stock
          this.copiedStockList = this.borrowData.stock
          this.borrowFormVisible = true
          const userRes = await findUsers({})
          if (userRes.code === 0) {
            this.userList = userRes.data
            this.copiedUserList = userRes.data
            this.borrowFormVisible = true
          }
        } else {
          this.$message({
            type: 'error',
            message: '商品下架或库存不足'
          })
        }
      }
    },
    setTypeOptions(TypeData, optionsData) {
      console.log(TypeData)
      TypeData &&
      TypeData.forEach(item => {
        if (item.children && item.children.length) {
          const option = {
            Id: item.ID,
            name: item.name,
            children: []
          }
          this.setTypeOptions(
            item.children,
            option.children,
          )
          optionsData.push(option)
        } else {
          const option = {
            ID: item.ID,
            name: item.name,
          }
          optionsData.push(option)
        }
      })
    },
    handleChange(stockId) {
      this.stockList = this.copiedStockList
      this.$emit('change', stockId)
      this.stockList.map(value => {
        // eslint-disable-next-line eqeqeq
        if (value.ID == stockId && value.creatorId != 0) {
          this.borrowData.creatorId = value.creatorId
        }
      })
    },
    handleUserChange(userId) {
      this.userList = this.copiedUserList
      this.$emit('change', userId)
    },
    handleClear() {
      this.stockList = this.copiedStockList
      this.$emit('change', '')
    },
    stockFilter(stockValue) {
      if (!stockValue) {
        this.stockList = this.copiedStockList
      } else {
        this.stockList = this.copiedStockList.filter(item => {
          console.log(item)
          if ((item.ID).toString().indexOf(stockValue) !== -1) {
            return true
          }
          if (item.stockNo.indexOf(stockValue) !== -1) {
            this.$emit('change', '')
            return true
          }
        })
      }
    },
    userFilter(userValue) {
      if (!userValue) {
        this.userList = this.copiedUserList
      } else {
        this.userList = this.copiedUserList.filter(item => {
          if ((item.ID).toString().indexOf(userValue) !== -1) {
            return true
          }
          if (item.nickName.indexOf(userValue) !== -1) {
            this.$emit('change', '')
            return true
          }
          if (item.sn.indexOf(userValue) !== -1) {
            this.$emit('change', '')
            return true
          }
        })
      }
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.borrowFormVisible = false
      this.formData = {
        name: '',
        author: '',
        press: '',
        pubDate: new Date(),
        upc: '',
        bookcaseId: '',
        price: 0,
        typeId: 0,
        photo: '',
        status: undefined,
        amount: 0,
      }
      this.borrowData = {
        name: '',
        author: '',
        press: '',
        pubDate: new Date(),
        upc: '',
        bookcaseId: '',
        price: 0,
        typeId: 0,
        photo: '',
        status: undefined,
        amount: 0,
      }
    },
    async deleteSysBooks(row) {
      const res = await deleteSysBooks({ ID: row.ID })
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
          res = await createSysBooks(this.formData)
          break
        case 'update':
          res = await updateSysBooks(this.formData)
          break
        case 'borrow':
          res = await borrowSysBooks({
            bookId: this.borrowData.ID,
            stockId: this.borrowData.stockId,
            userId: this.borrowData.userId,
            remark: this.borrowData.remark,
            type: 1
          })
          break
        case 'also':
          // eslint-disable-next-line no-undef
          res = await alsoSysBooks({
            bookId: this.borrowData.ID,
            stockId: this.borrowData.stockId,
            userId: this.borrowData.userId,
            remark: this.borrowData.remark,
            type: 2
          })
          break
        default:
          res = await createSysBooks(this.formData)
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
