<template>
  <div class="background">
    <div v-loading.fullscreen.lock="onLoading" class="container">
      <div class="line">
        <h1>RABUMS</h1>
      </div>
      <div class="line">
        <h3>앱리스트</h3>
      </div>
      <div class="line">현재 RABUMS를 사용중인 앱 리스트입니다.</div>
      <div class="line" style="overflow-y: scroll; height: 400px">
        <el-table
          :data="clientList"
          :border="true"
          style="min-width: 360px; width: 100%"
        >
          <el-table-column prop="createdAt" label="Date" width="100">
          </el-table-column>
          <el-table-column prop="clientId" label="Name"> </el-table-column>
          <el-table-column prop="link" label="Link" width="180">
          </el-table-column>
          <el-table-column prop="description" label="Description" width="300">
          </el-table-column>
          <el-table-column prop="validString" label="Valid" width="60">
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
      <div class="line">
        <nuxt-link to="/" style="margin: 0 10px">
          <el-button type="info"> 돌아가기 </el-button>
        </nuxt-link>
        <el-button @click="register()"> 추가 </el-button>
      </div>
      <el-dialog title="" :visible.sync="dialogFormVisible">
        <el-form :model="form">
          <el-form-item label="ID" :label-width="formLabelWidth">
            <el-input v-model="form.clientId" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="PW" :label-width="formLabelWidth">
            <el-input
              v-model="form.clientPw"
              autocomplete="off"
              :disabled="creatingNewClient"
              show-password
            ></el-input>
          </el-form-item>
          <el-form-item label="NewPW" :label-width="formLabelWidth">
            <el-input
              v-model="form.changedPw"
              autocomplete="off"
              placeholder="Leave blank, if not chaged."
              show-password
            ></el-input>
          </el-form-item>
          <el-form-item label="NewPWAgain" :label-width="formLabelWidth">
            <el-input
              v-model="form.changedPwChk"
              autocomplete="off"
              placeholder="Leave blank, if not chaged."
              show-password
            ></el-input>
          </el-form-item>
          <el-form-item label="Link" :label-width="formLabelWidth">
            <el-input v-model="form.link" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="Description" :label-width="formLabelWidth">
            <el-input v-model="form.description" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="Valid" :label-width="formLabelWidth">
            <el-select
              v-model="form.valid"
              placeholder="Only admin can change this"
            >
              <el-option label="true" value="true"></el-option>
              <el-option label="false" value="false"></el-option>
            </el-select>
          </el-form-item>
        </el-form>

        <span slot="footer" class="dialog-footer">
          <el-button @click="dialogFormVisible = false">Cancel</el-button>
          <el-button
            type="primary"
            :disabled="creatingNewClient || !isPasswdConfirmable"
            @click="submit(true)"
          >
            Renew Token
          </el-button>
          <el-button
            type="primary"
            :disabled="!isPasswdConfirmable"
            @click="submit(false)"
          >
            Confirm
          </el-button>
        </span>
      </el-dialog>
    </div>
  </div>
</template>

<script>
const shajs = require('sha.js')

const RabumsHASH = (text) => {
  text = Buffer.from(text).toString('base64')
  return shajs('sha256').update(text).digest('hex')
}

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
        clientId: '',
        clientPw: '',
        changedPwChk: '',
        changedPw: '',
        createdAt: new Date(),
        description: '',
        link: '',
        token: '',
        valid: '',
      },
      creatingNewClient: false,
      formLabelWidth: '120px',
      onLoading: false,
      dialogFormVisible: false,
    }
  },
  computed: {
    isPasswdConfirmable() {
      if (this.form.clientPw === '') return false
      if (this.creatingNewClient) {
        return this.form.changedPw === ''
          ? false
          : this.form.changedPw === this.form.changedPwChk
      } else {
        return this.form.changedPw === ''
          ? true
          : this.form.changedPw === this.form.changedPwChk
      }
    },
  },
  watch: {},
  async mounted() {
    this.onLoading = true
    await this.update()
    this.onLoading = false
  },
  methods: {
    async update() {
      try {
        const data = await this.$axios.get('/api/v1/client')
        this.clientList = data.data
        this.clientList.forEach((e) => {
          e.createdAt = new Date(Date.parse(e.createdAt)).toLocaleDateString()
          e.validString = e.valid ? '○' : '×'
        })
      } catch (err) {
        this.$message.error(JSON.stringify(err))
      }
    },
    async submit(renewToken) {
      try {
        await this.$confirm('정말로 변경하시겠습니까?', 'Warning', {
          confirmButtonText: 'OK',
          cancelButtonText: 'Cancel',
          type: 'warning',
        })
      } catch {
        return
      }
      this.onLoading = true
      try {
        const data = await this.$axios.put('/api/v1/client', {
          clientId: this.form.clientId,
          clientPw:
            this.form.clientPw === 'empty'
              ? 'empty'
              : RabumsHASH(this.form.clientPw),
          changedPw:
            this.form.changedPw === '' ? '' : RabumsHASH(this.form.changedPw),
          link: this.form.link,
          description: this.form.description,
          valid: this.form.valid === 'true',
          renewToken: renewToken === true,
        })
        this.dialogFormVisible = false
        this.$alert(
          `정보가 업데이트 되었습니다. 
          토큰: ${data.data.token}`,
          '성공',
          {
            confirmButtonText: 'OK',
          }
        )
        await this.update()
      } catch (err) {
        if (err.response.data)
          this.$message.error(JSON.stringify(err.response.data))
        else this.$message.error(JSON.stringify(err))
      }
      this.onLoading = false
    },
    edit(index, clientList) {
      this.dialogFormVisible = true
      this.creatingNewClient = false
      const data = clientList[index]
      this.form = {
        clientId: data.clientId,
        clientPw: '',
        changedPwChk: '',
        changedPw: '',
        createdAt: data.createdAt,
        description: data.description,
        link: data.link,
        token: data.token,
        valid: data.valid,
      }
    },
    register() {
      this.dialogFormVisible = true
      this.creatingNewClient = true
      this.form = {
        clientId: '',
        clientPw: 'empty',
        changedPwChk: '',
        changedPw: '',
        createdAt: new Date(),
        description: '',
        link: '',
        token: '',
        valid: 'false',
      }
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
