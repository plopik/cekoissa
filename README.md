# cekoissa

ssh -t ubuntu@music.douady.paris "cd cekoissa ; bash --login"
git pull
kill $(lsof -t -i:4277)
nohup go run . >> log.tmp & tail -f log.tmp