
import service from '@/utils/request'

export function Commits(page) {
  return service({
    url: '/sysBookRentLog/getSysBookRentLogList?page=1&pageSize=10&type=1',
    method: 'get',
  })
}

export function Members() {
  return service({
    url: 'https://api.github.com/orgs/FLIPPED-AURORA/members',
    method: 'get'
  })
}
