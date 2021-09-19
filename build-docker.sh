docker build . -t golang:1 && docker run -it -v $PWD:/app-dev golang:1
sudo chown $USER:$USER -R $PWD
