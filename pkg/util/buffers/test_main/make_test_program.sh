mkdir -p ./out/
GOOS=linux go build -o ./out/app_cancel_copy_test.linuxexe ./ignore_cancel_copy_main.go

sh ./download_pprof.sh

ssh root@etest1 'mkdir -p /root/tests'
ssh root@etest1 'ps -ef |grep app_cancel | xargs kill -9 ' || true
ssh root@etest1 'ls -lath /root/tests'
ssh root@etest1 'du -sh /root/tests'
ssh root@etest1 'rm -rf /root/tests/file_create_by_cancelable_copy_*'
ssh root@etest1 'ls -lath /root/tests'

chmod +x ./test_scripts.sh
chmod +x ./cat.sh
scp ./out/app_cancel_copy_test.linuxexe root@etest1:~/tests/
scp ./test_scripts.sh root@etest1:~/tests/
scp ./cat.sh root@etest1:~/tests/

ssh root@etest1
