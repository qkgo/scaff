# if [ $(uname) = "Darwin" ]; then echo 1; fi
if [ $(uname) = "Darwin" ]; then
  echo "please don't run in macos"
  exit 0
fi


echo '---------------------- '
echo 'you can execute next code: '
echo '---------------------- '
echo 'top -n 1 |grep app_cancel'
echo 'ifconfig |grep eth0 -A 6  |grep packets |grep  RX '
echo '---------------------- '
top -n 1 | grep app_cancel
ifconfig |grep eth0 -A 6  |grep packets |grep  RX
