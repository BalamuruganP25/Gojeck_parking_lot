# Gojeck_parking_lot

Before run the program please follow below steps

Language - GO (GOLANG)
OS - Centos 
version - 7.1
go version - go version go1.11.2 linux/amd64

1.Install GO

	https://golang.org/doc/install

	please use above link for install go in centos

2.set working directory path 
	open - below file
	vim/etc/bashrc

	Copy and past below commends in bashrc file
	----------------------------------------------
	 export GOROOT=/usr/local/go
	 export GOPATH=$HOME/Projects
	 export PATH=$GOPATH/bin:$GOROOT/bin:$PATH

3. after set the path please check go version
	go version
	output - example
	go version go1.11.2 linux/amd64


4.create working directory folder 

	open terminal 
	cd /root
	mkdir Projects
	cd Projects
	mkdir src
	mkdir pkg
	mkdir bin

5.Copy the parking_lot folder and past below location 

   /root/Projects/src

7.user upload file should be store in below path 

      /root/Projects/src/Gojeck_Parking_lot_problem_result/src/Useruploadfile

6.After copy the folder  compile and run the program

 using below commends

 cd /root/Projects/src/Gojeck_Parking_lot_problem_result/bin

 ./setup
 ./parking_lot

 7.Now Program hasbeen started

 	check log file below path

 	  cd /home/Parking_Log_Details
