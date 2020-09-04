<template>
  <div class="container">
    <div>
      <div>
        <img src="/icon.png" style="width: 80%; max-width: 1024px" />
        <h1 style="font-size: xxx-large">CSUOS | Rabums</h1>
        <h3>Rest API Based User Management System</h3>
      </div>
      <div class="links">
        <nuxt-link to="/applist" style="margin: 0 10px">
          <el-button class="main_button" type="info" plain>
            APP 관리
          </el-button>
        </nuxt-link>
        <nuxt-link to="/privacy" style="margin: 0 10px">
          <el-button class="main_button" type="primary" plain>
            보안 관리
          </el-button>
        </nuxt-link>
        <nuxt-link to="/register" style="margin: 0 10px">
          <el-button class="main_button" type="success" plain>
            회원가입
          </el-button>
        </nuxt-link>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {}
  },
}
</script>

<style>
@font-face {
  font-family: 'IBMPlexSansKR-Regular';
  src: url('https://cdn.jsdelivr.net/gh/projectnoonnu/noonfonts_20-07@1.0/IBMPlexSansKR-Regular.woff')
    format('woff');
  font-weight: normal;
  font-style: normal;
}
body {
  font-family: IBMPlexSansKR-Regular;
}
a {
  color: black;
  text-decoration: none;
}
.main_button {
  font-size: large;
  padding: 20px;
}
.container {
  margin: 0 auto;
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}

.title {
  font-family: 'Quicksand', 'Source Sans Pro', -apple-system, BlinkMacSystemFont,
    'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
  display: block;
  font-weight: 300;
  font-size: 100px;
  color: #35495e;
  letter-spacing: 1px;
}

.subtitle {
  font-weight: 300;
  font-size: 42px;
  color: #526488;
  word-spacing: 5px;
  padding-bottom: 15px;
}

.links {
  padding-top: 15px;
}
</style>
