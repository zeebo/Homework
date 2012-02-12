echo "Exact..."
cd exact
go run exact.go kernel1.go f1.go -n $N -top $TOP > ../results/exact.0.0.$N.$TOP.log
go run exact.go kernel1.go f2.go -n $N -top $TOP > ../results/exact.0.1.$N.$TOP.log
go run exact.go kernel1.go f3.go -n $N -top $TOP > ../results/exact.0.2.$N.$TOP.log
go run exact.go kernel2.go f1.go -n $N -top $TOP > ../results/exact.1.0.$N.$TOP.log
go run exact.go kernel2.go f2.go -n $N -top $TOP > ../results/exact.1.1.$N.$TOP.log
go run exact.go kernel2.go f3.go -n $N -top $TOP > ../results/exact.1.2.$N.$TOP.log
cd ..

echo "Lubich..."
cd lubich
go run lubich.go kernel1.go f1.go -n $N -top $TOP > ../results/lubich.0.0.$N.$TOP.log
go run lubich.go kernel1.go f2.go -n $N -top $TOP > ../results/lubich.0.1.$N.$TOP.log
go run lubich.go kernel1.go f3.go -n $N -top $TOP > ../results/lubich.0.2.$N.$TOP.log
go run lubich.go kernel2.go f1.go -n $N -top $TOP > ../results/lubich.1.0.$N.$TOP.log
go run lubich.go kernel2.go f2.go -n $N -top $TOP > ../results/lubich.1.1.$N.$TOP.log
go run lubich.go kernel2.go f3.go -n $N -top $TOP > ../results/lubich.1.2.$N.$TOP.log
cd ..

echo "Rect..."
cd rect
go run rect.go kernel1.go f1.go -n $N -top $TOP > ../results/rect.0.0.$N.$TOP.log
go run rect.go kernel1.go f2.go -n $N -top $TOP > ../results/rect.0.1.$N.$TOP.log
go run rect.go kernel1.go f3.go -n $N -top $TOP > ../results/rect.0.2.$N.$TOP.log
go run rect.go kernel2.go f1.go -n $N -top $TOP > ../results/rect.1.0.$N.$TOP.log
go run rect.go kernel2.go f2.go -n $N -top $TOP > ../results/rect.1.1.$N.$TOP.log
go run rect.go kernel2.go f3.go -n $N -top $TOP > ../results/rect.1.2.$N.$TOP.log
cd ..

echo "Mid..."
cd mid
go run mid.go kernel1.go f1.go -n $N -top $TOP > ../results/mid.0.0.$N.$TOP.log
go run mid.go kernel1.go f2.go -n $N -top $TOP > ../results/mid.0.1.$N.$TOP.log
go run mid.go kernel1.go f3.go -n $N -top $TOP > ../results/mid.0.2.$N.$TOP.log
go run mid.go kernel2.go f1.go -n $N -top $TOP > ../results/mid.1.0.$N.$TOP.log
go run mid.go kernel2.go f2.go -n $N -top $TOP > ../results/mid.1.1.$N.$TOP.log
go run mid.go kernel2.go f3.go -n $N -top $TOP > ../results/mid.1.2.$N.$TOP.log
cd ..

