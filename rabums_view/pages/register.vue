<template>
  <div class="background">
    <div class="register_container">
      <div class="line">
        <h2>
          RABUMS
        </h2>
      </div>
      <div class="line">
        <el-radio v-model="type" :label="1" border>학생</el-radio>
        <el-radio v-model="type" :label="2" border>교수/강사</el-radio>
      </div>
      <div class="line">
        <div class="title">
          <p style="margin: 10px auto;">이름 :</p>
        </div>
        <el-input
          v-model="user.userName"
          maxlength="15"
          show-word-limit
        ></el-input>
      </div>
      <div class="line">
        <div class="title">
          <p style="margin: 10px auto;">학번 :</p>
        </div>
        <el-input
          v-model="user.userNumber"
          maxlength="10"
          show-word-limit
          :disabled="type === 2"
        >
        </el-input>
      </div>
      <div class="line">
        <el-tag v-if="!validation.validUserNumber" type="danger"
          >적절한 학번이 아님</el-tag
        >
        <el-tag v-else type="success">:-)</el-tag>
      </div>
      <div class="line">
        <div class="title">
          <p style="margin: 10px auto;">아이디 :</p>
        </div>
        <el-input
          v-model="user.userId"
          maxlength="15"
          show-word-limit
        ></el-input>
      </div>
      <div class="line">
        <div class="title">
          <p style="margin: 10px auto;">비밀번호 :</p>
        </div>
        <el-input v-model="user.userPw" show-password></el-input>
      </div>
      <div class="line">
        <div class="title">
          <p style="margin: 10px auto;">비밀번호 확인 :</p>
        </div>
        <el-input v-model="user.userPwCheck" show-password></el-input>
      </div>
      <div class="line">
        <el-tag v-if="!validation.equalPassword" type="danger"
          >비밀번호 불일치</el-tag
        >
        <el-tag v-else type="success">비밀번호 일치</el-tag>
        <el-tag v-if="!validation.validPasswordLength" type="warning">
          너무 짧음
        </el-tag>
        <el-tag v-else type="success">적당함</el-tag>
      </div>
      <div class="line">
        <div class="title">
          <p style="margin: 10px auto;">이메일주소 :</p>
        </div>
        <el-input v-model="user.userEmail" maxlength="30">
          <template slot="append">@uos.ac.kr</template>
        </el-input>
      </div>
      <div class="line">
        <nuxt-link to="/" style="margin: 0 10px;">
          <el-button type="info" plain>
            취소
          </el-button>
        </nuxt-link>
        <el-button
          v-loading.fullscreen.lock="onLoading"
          style="margin: 0 10px;"
          type="primary"
          plain
          :disabled="!validation.consulusion"
          @click="submit"
        >
          제출
        </el-button>
      </div>
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
      type: 1,
      user: {
        userName: '',
        userNumber: '',
        userId: '',
        userPw: '',
        userPwCheck: '',
        userEmail: '',
      },
      onLoading: false,
    }
  },
  computed: {
    validation() {
      const output = {
        noneEmpty:
          this.user.userName.length !== 0 &&
          this.user.userNumber.length !== 0 &&
          this.user.userId.length !== 0 &&
          this.user.userPw.length !== 0 &&
          this.user.userEmail.length !== 0,
        equalPassword: this.user.userPw === this.user.userPwCheck,
        validPasswordLength: this.user.userPw.length > 6,
        validUserNumber:
          this.type === 1 ? this.user.userNumber * 1 > 2000000000 : true,
      }
      let consulusion = true
      Object.keys(output).forEach((k) => {
        if (output[k] !== true) consulusion = false
      })
      output.consulusion = consulusion
      return output
    },
  },
  watch: {
    type(e) {
      if (e === 2) {
        this.user.userNumber = '-1'
      } else {
        this.user.userNumber = ''
      }
    },
  },
  mounted() {},
  methods: {
    async submit() {
      if (this.onLoading === true) return
      this.onLoading = true
      try {
        await this.$axios.put('/api/v1/user/request/token', {
          userName: this.user.userName,
          userNumber: this.user.userNumber * 1,
          userEmail: this.user.userEmail + '@uos.ac.kr',
          userId: this.user.userId,
          userPw: RabumsHASH(RabumsHASH(this.user.userPw)),
        })
        this.onLoading = false
        this.$message({
          message: `${this.user.userEmail}@uos.ac.kr 을 확인해주세요! :-)`,
          type: 'success',
        })
        this.$router.push('/token')
      } catch (error) {
        console.log(error)
        console.log(error.response.data)
        this.$message.error(JSON.stringify(error.response.data))
      }
      this.onLoading = false
    },
  },
}
</script>

<style>
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
.register_container {
  margin: auto;
  min-height: 640px;
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
