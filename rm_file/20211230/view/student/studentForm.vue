<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="学生编号:">
          <el-input v-model="formData.sn" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="姓名:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="性别:">
          <el-select v-model="formData.sex" placeholder="请选择" clearable>
            <el-option v-for="(item,key) in genderOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
        </el-form-item>
        <el-form-item label="邮箱:">
          <el-input v-model="formData.email" clearable placeholder="请输入" />
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
  createStudent,
  updateStudent,
  findStudent
} from '@/api/student' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'Student',
  mixins: [infoList],
  data() {
    return {
      type: '',
      genderOptions: [],
      formData: {
        sn: '',
        name: '',
        sex: undefined,
        email: '',
      }
    }
  },
  async created() {
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findStudent({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.restudent
        this.type = 'update'
      }
    } else {
      this.type = 'create'
    }
    await this.getDict('gender')
  },
  methods: {
    async save() {
      let res
      switch (this.type) {
        case 'create':
          res = await createStudent(this.formData)
          break
        case 'update':
          res = await updateStudent(this.formData)
          break
        default:
          res = await createStudent(this.formData)
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
