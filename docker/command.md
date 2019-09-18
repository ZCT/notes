check

 curl https://raw.githubusercontent.com/docker/docker/master/contrib/check-config.sh > check-config.sh



sudo systemctl status docker.service



sudo journalctl -u docker.service --no-pager





dockerd --iptables=False



sudo dockerd --iptables=False  2>&1 &



sudo apt-get remove ––purge docker-ce (彻底删除)