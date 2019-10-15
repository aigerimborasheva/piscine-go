curl https://api.github.com/users | jq '.[]| select(.id==75) | .login'
