# RABUMS
Rest API Based User Management System

`/v1/user`
* GET
  * Healthcheck 용
* POST
  * 유저 정보 가져오는 요청
  * 구현됨, 테스트 필요
* PUT
  * 유저 정보 입력하는 요청
* DELETE
  * 유저 정보 삭제하는 요청

`/v1/client`
* GET
  * 현재 사용 가능한 클라이언트 리스트 및 사용량 정보
* POST
  * 클라이언트 세부 개인 정보 가져오는 요청
* PUT
  * 클라이언트 추가 혹은 갱신하는 요청
  * 구현중
* DELETE
  * 클라이언트 삭제하는 요청