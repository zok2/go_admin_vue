<template>
  <el-select
    :value="value"
    clearable
    filterable
    :filter-method="storesFilter"
    @change="handleChange"
    @clear="handleClear"
  >
    <el-option v-for="store in storeList" :key="store.id" :label="store.name" :value="store.id" />
  </el-select>
</template>

<script>
export default {
  model: {
    prop: 'value',
    event: 'change'
  },
  props: {
    value: [String, Number],
  },
  data() {
    return {
      List: [],
      copiedList: []
    }
  },
  mounted() {
    this.fetchList()
  },
  methods: {
    fetchList() {
      this.List = this.value
      this.copiedList = this.value
    },
    storesFilter(value) {
      if (!value) {
        this.List = this.copiedList
      } else {
        this.List = this.copiedList.filter(item => {
          if (item.id.indexOf(value) !== -1) {
            return true
          }
          if (item.name.indexOf(value) !== -1) {
            return true
          }
        })
      }
    },
    handleChange(storeId) {
      this.$emit('change', storeId)
      this.$emit('select', storeId)
      this.List = this.copiedList
    },
    handleClear() {
      this.$emit('change', '')
      this.List = this.copiedList
    }
  }
}
</script>
