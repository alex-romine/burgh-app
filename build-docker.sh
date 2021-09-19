echo 'building image'
docker build . -t golang:1 && echo 'running image' && docker run -it -p 8888:8888 -v $PWD:/app-dev golang:1

echo 'chowning'
sudo chown $USER:$USER -R $PWD
