module.exports = {
  EventType: {
    '-1': '알수없는 이벤트',
    100: '데이터가 조회됨',
    101: '🎉🎉🎉🍾회원가입 🍾🎉🎉🎉',
    102: '회원정보 갱신됨',
    103: '삭제됨',
    200: '로그인 성공',
    201: '로그아웃',
    401: 'Client 비밀번호 틀림',
    402: '인증되지 않은 클라이언트에서 로그인이 시도됨',
    403: '없는 유저로 로그인을 시도함',
    404: '비밀번호 틀림',
  },
}
/*
	UNKOWN   EventType = -1
	ACCESSED EventType = 100
	CREATED  EventType = 101
	UPDATED  EventType = 102
	DELETED  EventType = 103
	LOGIN    EventType = 200
	LOGOUT   EventType = 201

	INVALIDCLIENTPW   EventType = 401
	INVALIDTOKEN      EventType = 402
	USERNOTFOUND      EventType = 403
	INCORRECTPASSWORD EventType = 404
*/
