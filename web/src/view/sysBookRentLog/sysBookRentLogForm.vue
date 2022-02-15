<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="图书id:">
          <el-input v-model.number="formData.bookId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="创建人id:">
          <el-input v-model.number="formData.creatorId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="借阅者id:">
          <el-input v-model.number="formData.userId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="类型:">
          <el-select v-model="formData.type" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in typeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="备注:">
          <el-input v-model="formData.remark" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="库存id:">
          <el-input v-model.number="formData.stockId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
import {
  createSysBookRentLog,
  updateSysBookRentLog,
  findSysBookRentLog
} from '@/api/sysBookRentLog' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'SysBookRentLog',
  mixins: [infoList],
  data() {
    return {
      type: '',
      typeOptions: [],
      formData: {
        bookId: 0,
        creatorId: 0,
        userId: 0,
        type: undefined,
        remark: '',
        stockId: 0,
      }
    }
  },
  async created() {
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findSysBookRentLog({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.resysBookRentLog
        this.type = 'update'
      }
    } else {
      this.type = 'create'
    }
    await this.getDict('type')
  },
  methods: {
    async save() {
      let res
      switch (this.type) {
        case 'create':
          res = await createSysBookRentLog(this.formData)
          break
        case 'update':
          res = await updateSysBookRentLog(this.formData)
          break
        default:
          res = await createSysBookRentLog(this.formData)
          break
      }
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '创建/更改成功'
        })
      }
    },
    back() {
      this.$router.go(-1)
    }
  }
}
</script>

<style>
</style>
