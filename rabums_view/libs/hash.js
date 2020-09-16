const shajs = require('sha.js')

const RabumsHASH = (text) => {
  text = Buffer.from(text).toString('base64')
  return shajs('sha256').update(text).digest('hex')
}

const RabumsTokenHash = (hashedToken, password) => {
  const hashedpw = RabumsHASH(RabumsHASH(password))
  return RabumsHASH(hashedToken + hashedpw)
}

module.exports = { RabumsHASH, RabumsTokenHash }
