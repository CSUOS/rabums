<template>
  <div class="background">
    <div class="container">
      <div class="line">
        <h1>
          RABUMS
        </h1>
      </div>
      <div class="line">
        <h3>앱리스트</h3>
      </div>
      <div class="line">
        현재 RABUMS를 사용중인 앱 리스트입니다.
      </div>
      <div class="line" style="overflow-y: scroll; height: 400px;">
        <el-table
          :data="clientList"
          :border="true"
          style="min-width: 360px; width: 100%;"
        >
          <el-table-column prop="createdAt" label="Date" width="100">
          </el-table-column>
          <el-table-column prop="clientId" label="Name"> </el-table-column>
          <el-table-column prop="link" label="Link" width="180">
          </el-table-column>
          <el-table-column prop="description" label="Description" width="300">
          </el-table-column>
          <el-table-column prop="valid" label="Valid" width="60">
          </el-table-column>
          <el-table-column fixed="right" label="" width="50">
            <template slot-scope="scope">
              <el-button
                type="text"
                size="small"
                @click.native.prevent="edit(scope.$index, clientList)"
              >
                Edit
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
      <el-dialog title="" :visible.sync="dialogFormVisible">
        <el-form :model="form">
          <el-form-item label="ID" :label-width="formLabelWidth">
            <el-input v-model="form.name" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="OriginalPW" :label-width="formLabelWidth">
            <el-input v-model="form.name" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="NewPW" :label-width="formLabelWidth">
            <el-input
              v-model="form.name"
              autocomplete="off"
              placeholder="Leave blank, if not chaged."
            ></el-input>
          </el-form-item>
          <el-form-item label="NewPWAgain" :label-width="formLabelWidth">
            <el-input v-model="form.name" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="Link" :label-width="formLabelWidth">
            <el-input v-model="form.name" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="Description" :label-width="formLabelWidth">
            <el-input v-model="form.name" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="Valid" :label-width="formLabelWidth">
            <el-select
              v-model="form.region"
              placeholder="Only admin can change this"
            >
              <el-option label="true" value="true"></el-option>
              <el-option label="false" value="false"></el-option>
            </el-select>
          </el-form-item>
        </el-form>

        <span slot="footer" class="dialog-footer">
          <el-button @click="dialogFormVisible = false">Cancel</el-button>
          <el-button type="primary" @click="dialogFormVisible = false">
            Confirm
          </el-button>
        </span>
      </el-dialog>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      token: '',
      clientList: [
        {
          clientId: '',
          clientPw: '',
          createdAt: new Date(),
          description: '',
          link: '',
          token: '',
          valid: '',
        },
      ],
      form: {
        name: '',
        region: '',
        date1: '',
        date2: '',
        delivery: false,
        type: [],
        resource: '',
        desc: '',
      },
      formLabelWidth: '120px',
      onLoading: false,
      dialogFormVisible: false,
    }
  },
  computed: {},
  watch: {},
  async mounted() {
    this.onLoading = true
    try {
      const data = await this.$axios.get('/api/v1/client')
      this.clientList = data.data
      this.clientList.forEach((e) => {
        e.createdAt = new Date(Date.parse(e.createdAt)).toLocaleDateString()
        e.valid = e.valid ? '○' : '×'
      })
    } catch (err) {
      this.$message.error(JSON.stringify(err))
    }
    this.onLoading = false
  },
  methods: {
    async submit() {},
    edit(index, clientList) {
      this.dialogFormVisible = true
      console.log(index, clientList)
    },
  },
}
</script>

<style scoped>
.title {
  text-align: initial;
  font-size: 16px !important;
  min-width: 160px;
}
@media screen and (max-width: 500px) {
  .line {
    display: block !important;
  }
}

.line {
  margin: auto;
  display: flex;
  margin: 10px;
}
.background {
  display: flex;
  height: 100vh;
  background: whitesmoke;
}
.container {
  margin: auto;

  min-height: 80%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  text-align: center;
  background: white;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  border-radius: 4px;
  width: 80%;
  max-width: 1024px;
  min-width: 360px;
}
.el-input {
  width: 250px;
}
</style>
