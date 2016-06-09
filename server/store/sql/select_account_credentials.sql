select id, role, password
  from account
 where email = ?
 limit 1
