module load golang/1.16.2

cd /editor

echo "ws"
echo "----"

for i in {1..3}
do
    go run editor.go -m ws -n 20000 -i 10 -t 64
done
echo "----"

echo "wb"
echo "----"
for i in {1..3}
do
    go run editor.go -m wb -n 20000 -i 10 -t 64
done
echo "----"

echo "s"
echo "----"
for i in {1..3}
do
    go run editor.go -m s -n 20000 -i 10
done
echo "----"

echo "------------------"

echo "ws"
echo "----"
go run editor.go -m ws -n 20000 -i 10 -t 2
go run editor.go -m ws -n 20000 -i 10 -t 4
go run editor.go -m ws -n 20000 -i 10 -t 8
go run editor.go -m ws -n 20000 -i 10 -t 16
go run editor.go -m ws -n 20000 -i 10 -t 32
go run editor.go -m ws -n 20000 -i 10 -t 64
go run editor.go -m ws -n 20000 -i 10 -t 128
echo "----"

echo "wb"
echo "----"
go run editor.go -m wb -n 20000 -i 10 -t 2
go run editor.go -m wb -n 20000 -i 10 -t 4
go run editor.go -m wb -n 20000 -i 10 -t 8
go run editor.go -m wb -n 20000 -i 10 -t 16
go run editor.go -m wb -n 20000 -i 10 -t 32
go run editor.go -m wb -n 20000 -i 10 -t 64
go run editor.go -m wb -n 20000 -i 10 -t 128
echo "----"