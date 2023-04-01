

nohup ./app_cancel_copy_test.linuxexe 100 10 . &

rm ./heap.pprof.out
rm ./profile.pprof.out
rm ./block.pprof.out
rm ./mutex.pprof.out

sleep 10
nohup  curl -o heap.pprof.out    'http://localhost:7777/debug/pprof/heap' . &
nohup  curl -o profile.pprof.out 'http://localhost:7777/debug/pprof/profile?seconds=10' . &
nohup  curl -o block.pprof.out   'http://localhost:7777/debug/pprof/block' . &
nohup  curl -o mutex.pprof.out   'http://localhost:7777/debug/pprof/mutex' . &
sleep 2
sh ./cat.sh
sleep 8
echo 'test finished, you can stop the job by pid in before output'

