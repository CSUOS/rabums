<template>
  <div style="margin: auto">
    <el-timeline :v-loading="loading">
      <el-timeline-item
        v-for="l in logs"
        :key="l.date"
        :color="l.color"
        :timestamp="l.date"
      >
        <a :href="l.clientLink">{{ l.clientId }}</a> 에서
        {{ l.event }}
      </el-timeline-item>
    </el-timeline>
  </div>
</template>

<script>
const { EventType } = require('../libs/types')
export default {
  data() {
    return {
      loading: false,
      logs: [
        {
          clientId: '',
          clientLink: '',
          event: 100,
          date: '',
          color: 'red',
        },
      ],
    }
  },
  async mounted() {
    this.loading = true
    try {
      const client = {}
      const clientList = (await this.$axios.get('/api/v1/client')).data
      clientList.forEach((e) => {
        client[e.id] = e
      })

      const logs = (await this.$axios.get('/api/v1frontend/logs')).data
      this.logs = []
      logs.forEach((e) => {
        let color = '#409eff'
        if (e.event >= 400) color = '#f56c6c'
        this.logs.push({
          clientId: client[e.clientId].clientId,
          clientLink: client[e.clientId].link,
          event: EventType[e.event],
          date: new Date(Date.parse(e.date)).toLocaleString(),
          color,
        })
      })
    } catch (error) {}
    this.loading = false
  },
}
</script>
<style></style>
