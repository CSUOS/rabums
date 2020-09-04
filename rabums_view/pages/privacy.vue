<template>
  <div class="background">
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
        <el-input placeholder="이완해" disabled></el-input>
      </div>
      <div class="line">
        <div class="title">
          <p style="margin: 10px auto">학번(사용자번호) :</p>
        </div>
        <el-input placeholder="2016920036" disabled></el-input>
      </div>
      <div class="line">
        <div class="title">
          <p style="margin: 10px auto">이메일주소 :</p>
        </div>
        <el-input placeholder="train96@uos.ac.kr" disabled></el-input>
      </div>
      <div class="line">
        <div class="title">
          <p style="margin: 10px auto">아이디 :</p>
        </div>
        <el-input placeholder=""></el-input>
      </div>
      <div class="line">
        <div class="title">
          <p style="margin: 10px auto">비밀번호 :</p>
        </div>
        <el-input></el-input>
      </div>
      <div class="line">
        <div class="title">
          <p style="margin: 10px auto">비밀번호 확인 :</p>
        </div>
        <el-input></el-input>
      </div>
      <div class="line">
        <nuxt-link to="/" style="margin: 0 10px">
          <el-button>돌아가기</el-button>
        </nuxt-link>
        <el-button type="primary">적용</el-button>
      </div>
      <el-divider />
      <div class="line">
        <h3>활동기록</h3>
      </div>
      <div class="line" style="height: 200px; overflow-y: scroll; width: 100%">
        <el-timeline>
          <el-timeline-item> 계정생성 </el-timeline-item>
          <el-timeline-item> 계정생성 </el-timeline-item>
          <el-timeline-item> 계정생성 </el-timeline-item>
          <el-timeline-item> 계정생성 </el-timeline-item>
          <el-timeline-item> 계정생성 </el-timeline-item>
          <el-timeline-item> 계정생성 </el-timeline-item>
          <el-timeline-item> 계정생성 </el-timeline-item>
          <el-timeline-item> 계정생성 </el-timeline-item>
          <el-timeline-item> 계정생성 </el-timeline-item>
          <el-timeline-item> 계정생성 </el-timeline-item>
          <el-timeline-item> 계정생성 </el-timeline-item>
          <el-timeline-item> 계정생성 </el-timeline-item>
        </el-timeline>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      token: '',
      onLoading: false,
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
