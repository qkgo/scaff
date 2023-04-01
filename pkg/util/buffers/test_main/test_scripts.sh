# if [ $(uname) = "Darwin" ]; then echo 1; fi
if [ $(uname) = "Darwin" ]; then
  echo "please don't run in macos"
  exit 0
fi


# wget 'https://consumer-tkb.huawei.com/weknow/servlet/download/public?contextNo=W00036307'
rm -rf ./nohup.out
rm -rf ./file_create_by_cancelable_copy_*
ps -ef |grep app_cancel | awk '{print $2}'  | xargs kill -9  || true
nohup ./app_cancel_copy_test.linuxexe 100 2.5 'https://consumer-tkb.huawei.com/weknow/servlet/download/public?contextNo=W00036307' . &


rm ./heap.pprof.out || true
rm ./profile.pprof.out || true
rm ./block.pprof.out || true
rm ./mutex.pprof.out || true

sleep 2.5
nohup $(curl -o profile.pprof.out 'http://localhost:7777/debug/pprof/profile?seconds=10') . &
nohup $(curl -o heap.pprof.out 'http://localhost:7777/debug/pprof/heap') . &
nohup $(curl -o block.pprof.out 'http://localhost:7777/debug/pprof/block') . &
nohup $(curl -o mutex.pprof.out 'http://localhost:7777/debug/pprof/mutex') . &
sleep 1

echo 'starting ... '
sh ./cat.sh
echo 'test finished, you can stop the job by pid in before output'
