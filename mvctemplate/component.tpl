<template>
  <el-select
    v-model.trim="input"
    filterable
    remote
    placeholder="{{.Name}}"
    @focus="getlist"
    :loading="loading"
    @change="handleSelect"
    clearable
    :disabled="disabled"
  >
    <el-option
      v-for="item in list"
      :key="item.{{.PrimaryKeyColumn.ColumnName}}"
      :label="item.{{.PrimaryKeyColumn.ColumnName}}"
      :value="item.{{.PrimaryKeyColumn.ColumnName}}"
    ></el-option>
  </el-select>
</template>

<script>
  import * as api from "@/api/{{.Name}}";
  export default {
    data() {
      return {
        loading: false,
        input: "",
        list: []
      };
    },
    props: { value: Number, disabled: Boolean },
    methods: {
      handleSelect(v) {
        if (typeof v == "number") {
          this.$emit("input", v);
          this.$emit("change", v);
        } else {
          this.$emit("input", 0);
          this.$emit("change", 0);
        }
      },
      getlist() {
        this.loading = true;
        api
          .List()
          .then(rsp => {
            this.list = rsp.data.data;
          })
          .catch(err => {})
          .then(() => {
            this.loading = false;
          });
      },
      init() {
        this.list = [];
        this.input = this.value;
        if (this.value > 0) {
          this.getlist();
        }
      }
    },
    watch: {
      value: function() {
        this.init();
      }
    },
    mounted() {
      this.init();
    }
  };
</script>

