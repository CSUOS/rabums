<template>
  <div class="background">
    <el-dialog
      title="회원정보관리"
      :visible.sync="showDialog"
      :fullscreen="true"
      :show-close="false"
      :center="true"
    >
      <div style="max-width: 720px; margin: auto">
        <h3>회원정보를 수정하실려면 먼저 로그인을 해주세요</h3>
        <div style="display: flex; margin: 10px">
          <h3 style="margin: auto; width: 70px">아이디 :</h3>
          <el-input v-model="user.userId" style="margin: auto"></el-input>
        </div>
        <div style="display: flex; margin: 10px">
          <h3 style="margin: auto; width: 70px">비밀번호 :</h3>
          <el-input
            v-model="user.userPw"
            style="margin: auto"
            show-password
          ></el-input>
        </div>
      </div>
      <span slot="footer" class="dialog-footer">
        <nuxt-link to="/" style="margin: 0 10px">
          <el-button>돌아가기</el-button>
        </nuxt-link>
        <el-button type="primary" @click="login"> 로그인 </el-button>
      </span>
    </el-dialog>
    <div class="container">
      <div class="line">
        <h1>RABUMS</h1>
      </div>
      <div class="line">
        <h3>회원정보수정</h3>
      </div>
      <div class="line">
        <div class="title">
          <p style="margin: 10px auto">이름 :</p>
        </div>
        <el-input :placeholder="user.userName" disabled></el-input>
      </div>
      <div class="line">
        <div class="title">
          <p style="margin: 10px auto">학번(사용자번호) :</p>
        </div>
        <el-input :placeholder="user.userNumber" disabled></el-input>
      </div>
      <div class="line">
        <div class="title">
          <p style="margin: 10px auto">이메일주소 :</p>
        </div>
        <el-input :placeholder="user.userEmail" disabled></el-input>
      </div>
      <div class="line">
        <div class="title">
          <p style="margin: 10px auto">아이디 :</p>
        </div>
        <el-input v-model="user.userId"></el-input>
      </div>
      <div class="line">
        <div class="title">
          <p style="margin: 10px auto">비밀번호 :</p>
        </div>
        <el-input v-model="user.userPw" show-password></el-input>
      </div>
      <div class="line">
        <div class="title">
          <p style="margin: 10px auto">비밀번호 확인 :</p>
        </div>
        <el-input v-model="user.userPwChk" show-password></el-input>
      </div>
      <div class="line">
        <el-button @click="logout">로그아웃</el-button>
        <el-button type="primary" disabled>적용</el-button>
      </div>
      <el-divider />
      <div class="line">
        <h3>활동기록</h3>
      </div>
      <div class="line" style="height: 200px; overflow-y: scroll; width: 100%">
        <Logs v-if="!showDialog" />
      </div>
    </div>
  </div>
</template>

<script>
import Logs from '../components/Logs'
const { RabumsTokenHash } = require('../libs/hash')

export default {
  components: { Logs },
  data() {
    return {
      token: '',
      onLoading: false,
      showDialog: true,
      user: {
        userName: '',
        userId: '',
        userNumber: '',
        userEmail: '',
        userPw: '',
        userPwChk: '',
      },
    }
  },
  computed: {},
  watch: {},
  methods: {
    async submit() {
      if (this.onLoading === true) return
      this.onLoading = true
      try {
        const data = await this.$axios.put('/api/v1/user/request/register', {
          token: this.token,
        })
        this.$message({
          message: `환영합니다! '${data.data.userId}'님!`,
          type: 'success',
        })
        this.$router.push('/')
      } catch (error) {
        this.$message.error(JSON.stringify(error.response.data))
      }
      this.onLoading = false
    },
    async login() {
      if (this.onLoading === true) return
      this.onLoading = true
      try {
        const { token } = (await this.$axios.get('/api/v1frontend/token')).data
        this.user = (
          await this.$axios.post('/api/v1frontend/login', {
            userId: this.user.userId,
            userPw: RabumsTokenHash(token, this.user.userPw),
          })
        ).data
        this.showDialog = false
      } catch (error) {
        if (error.response)
          this.$message.error(JSON.stringify(error.response.data))
        else this.$message.error('알수없는 에러가 발생했습니다. :-(')
      }
      this.onLoading = false
    },
    async logout() {
      if (this.onLoading === true) return
      this.onLoading = true
      try {
        await this.$axios.get('/api/v1frontend/logout')
      } catch (error) {
        if (error.response)
          this.$message.error(JSON.stringify(error.response.data))
        else this.$message.error('알수없는 에러가 발생했습니다. :-(')
      }
      this.user = {
        userName: '',
        userId: '',
        userNumber: '',
        userEmail: '',
        userPw: '',
        userPwChk: '',
      }
      this.showDialog = true
      this.onLoading = false
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
