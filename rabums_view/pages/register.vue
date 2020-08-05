<template>
  <div class="background">
    <div class="register_container">
      <div class="line">
        <h2>
          RABUMS
        </h2>
      </div>
      <div class="line">
        <el-radio v-model="type" label="1" border>학생</el-radio>
        <el-radio v-model="type" label="2" border>교수/강사</el-radio>
      </div>
      <div class="line">
        <div class="title">
          <p style="margin: 10px auto;">이름 :</p>
        </div>
        <el-input placeholder="Please input" v-model="input"></el-input>
      </div>
      <div class="line">
        <div class="title">
          <p style="margin: 10px auto;">학번 :</p>
        </div>
        <el-input placeholder="Please input" v-model="input"></el-input>
      </div>
      <div class="line">
        <div class="title">
          <p style="margin: 10px auto;">아이디 :</p>
        </div>
        <el-input placeholder="Please input" v-model="input"></el-input>
      </div>
      <div class="line">
        <div class="title">
          <p style="margin: 10px auto;">비밀번호 :</p>
        </div>
        <el-input placeholder="Please input" v-model="input"></el-input>
      </div>
      <div class="line">
        <div class="title">
          <p style="margin: 10px auto;">비밀번호 확인 :</p>
        </div>
        <el-input placeholder="Please input" v-model="input"></el-input>
      </div>
      <div class="line">
        <div class="title">
          <p style="margin: 10px auto;">이메일주소 :</p>
        </div>
        <el-input placeholder="Please input" v-model="input"></el-input>
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
      type: 0,
    }
  },
}
</script>

<style>
.title {
  text-align: initial;
  font-size: 16px !important;
  min-width: 160px;
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
  height: 80%;
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
</style>
