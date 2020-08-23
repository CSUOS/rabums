export const state = () => ({
  userId: '',
  userPw: '',
  token: '',
})
const shajs = require('sha.js')

const RabumsHASH = (text) => {
  text = Buffer.from(text).toString('base64')
  return shajs('sha256').update(text).digest('hex')
}

export const mutations = {
  login(state, userId, userPw) {
    state.userPw = RabumsHASH(RabumsHASH(userPw))
    state.userId = userId
  },
}
