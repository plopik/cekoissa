# cekoissa

ssh ubuntu@music.douady.paris
cd cekoissa
git pull
kill $(lsof -t -i:4277)
nohup go run . > log.tmp & cat log.tmp