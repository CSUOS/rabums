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
        <el-tag v-if="user.userNumber * 1 < 2000000000" type="danger"
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
        <el-tag v-if="!equalPassword" type="danger">비밀번호 불일치</el-tag>
        <el-tag v-else type="success">비밀번호 일치</el-tag>
        <el-tag v-if="user.userPw.length < 6" type="warning">
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
        <el-button style="margin: 0 10px;" type="primary" plain>
          제출
        </el-button>
      </div>
    </div>
  </div>
</template>

<script>
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
    }
  },
  computed: {
    equalPassword() {
      console.log(this.user)
      return this.user.userPw === this.user.userPwCheck
    },
  },
  watch: {
    type(e) {
      if (e === 2) {
        this.user.userNumber = '0'
      } else {
        this.user.userNumber = ''
      }
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
