<template>
  <div class="background">
    <div class="register_container">
      <div class="line">
        <h1>
          RABUMS
        </h1>
      </div>
      <div class="line">
        <h3>토큰</h3>
      </div>
      <el-input
        v-model="token"
        :autosize="{ minRows: 2, maxRows: 6 }"
        placeholder="이메일로 받은 토큰을 입력해주세요"
        type="textarea"
        style="width: 80%; max-width: 480px; margin: 0 auto;"
      ></el-input>
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
          :disabled="token.length == 0"
          @click="submit"
        >
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
