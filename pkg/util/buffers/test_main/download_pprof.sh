

mkdir -p ./out
scp -rp root@etest1:/root/tests/heap.pprof.out ./out ||true
scp -rp root@etest1:/root/tests/profile.pprof.out ./out ||true
scp -rp root@etest1:/root/tests/block.pprof.out ./out ||true
scp -rp root@etest1:/root/tests/mutex.pprof.out ./out ||true


cd ./out/
mkdir -p ./oldtestlog
#go tool pprof --png /bin/ls  ./out/profile.out
go tool pprof --pdf /bin/ls  ./heap.pprof.out
go tool pprof --pdf /bin/ls  ./profile.pprof.out
go tool pprof --pdf /bin/ls  ./block.pprof.out
go tool pprof --pdf /bin/ls  ./mutex.pprof.out
cd -