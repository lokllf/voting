# client
docker-compose -f ~/go/src/voting/sawtooth-default.yaml up -d
nohup go run ~/go/src/voting/processor/organizer/main.go -vvv &> organizer.out&
nohup go run ~/go/src/voting/processor/voter/main.go -vvv &> voter.out&
nohup go run ~/go/src/voting/client/main.go &> client.out&

# curl
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"private_key":"793b849da13c4829693fa555c54686e44951f227637929e0997bd1b67705ecec","vote":{"name":"vote 1","description":"test vote 1","candidates":[{"code":"1","name":"candidate 1","description":"test candidate 1","status":0},{"code":"2","name":"candidate 2","description":"test candidate 2","status":0}],"start_at":1,"end_at":9999999999}}' \
  http://localhost:9009/vote/create

curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"private_key":"793b849da13c4829693fa555c54686e44951f227637929e0997bd1b67705ecec","vote":{"id":"4f19267a-128c-11e9-bebf-12062d656d28","name":"new name","description":"test vote 1","candidates":[{"code":"1","name":"candidate 1","description":"test candidate 1","status":0},{"code":"2","name":"candidate 2","description":"test candidate 2","status":0}],"start_at":1,"end_at":9999999999}}' \
  http://localhost:9009/vote/update

curl --header "Content-Type: application/json" \
  --request DELETE \
  --data '{"private_key":"793b849da13c4829693fa555c54686e44951f227637929e0997bd1b67705ecec","vote_id":"4f19267a-128c-11e9-bebf-12062d656d28"}' \
  http://localhost:9009/vote

curl http://localhost:9009/vote/4f19267a-128c-11e9-bebf-12062d656d28

curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"private_key":"793b849da13c4829693fa555c54686e44951f227637929e0997bd1b67705ecec","vote_id":"4f19267a-128c-11e9-bebf-12062d656d28","quantity":1}' \
  http://localhost:9009/ballot/add

curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"ballot":{"vote_id":"4f19267a-128c-11e9-bebf-12062d656d28","choice":"1"},"code":"4bdf6976-1289-11e9-ba43-12062d656d28"}' \
  http://localhost:9009/ballot/cast

curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"private_key":"793b849da13c4829693fa555c54686e44951f227637929e0997bd1b67705ecec","vote_id":"4f19267a-128c-11e9-bebf-12062d656d28"}' \
  http://localhost:9009/ballot/count

curl http://localhost:9009/vote/4f19267a-128c-11e9-bebf-12062d656d28/result