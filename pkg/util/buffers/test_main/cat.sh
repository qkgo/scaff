echo '---------------------- '
echo 'you can execute next code: '
echo '---------------------- '
echo 'ps -ef |grep app_cancel | xargs kill -9 '
echo '---------------------- '
echo 'starting ... '
echo 'top -n 1 |grep app_cancel'
echo 'ifconfig |grep eth0 -A 6  |grep packets'
echo '---------------------- '
top -n 1 |grep app_cancel
ifconfig |grep eth0 -A 6  |grep packets