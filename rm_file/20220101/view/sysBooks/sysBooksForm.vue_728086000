<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="名称:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="作者:">
          <el-input v-model="formData.author" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="国际标准图书编码:">
          <el-input v-model="formData.isbn" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="书架编号:">
          <el-input v-model.number="formData.bookcaseId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="价格:">
          <el-input v-model.number="formData.price" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="类型:">
          <el-input v-model.number="formData.typeId" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="封面:">
          <el-input v-model="formData.photo" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="状态:">
          <el-select v-model="formData.status" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
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
  createSysBooks,
  updateSysBooks,
  findSysBooks
} from '@/api/sysBooks' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'SysBooks',
  mixins: [infoList],
  data() {
    return {
      type: '',
      statusOptions: [],
      formData: {
        name: '',
        author: '',
        isbn: '',
        bookcaseId: 0,
        price: 0,
        typeId: 0,
        photo: '',
        status: undefined,
      }
    }
  },
  async created() {
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findSysBooks({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.resysBooks
        this.type = 'update'
      }
    } else {
      this.type = 'create'
    }
    await this.getDict('status')
  },
  methods: {
    async save() {
      let res
      switch (this.type) {
        case 'create':
          res = await createSysBooks(this.formData)
          break
        case 'update':
          res = await updateSysBooks(this.formData)
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
