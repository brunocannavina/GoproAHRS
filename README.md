# GoproAHRS
Convert IMU metadata from Gopro Hero 5 to AHRS using the CSV file from DashWare

## How to Use
- You must have Golang installed<br/>
- You must have my [goahrs](https://github.com/brunocannavina/goahrs) package installed<br/>
- You must have [DashWare](http://www.dashware.net/) installed<br/>
- And of course a video of Gopro Hero 5 or later<br/>

When you import the video into DashWare,<br/>
the software will automatically generate a CSV file inside the project folder.<br/>
Copy this file.<br/>
_(Contains all the metadata of the video)_<br/>

## First method
Inside any folder execute the command:<br/>
`$ git clone https://github.com/brunocannavina/goproahrs.git`<br/>
_(or you can simply download it from github)_<br/>
then, open the folder goproahrs:<br/>
`$ cd goproahrs`<br/>
paste the CSV file copied above<br/>
and execute the following command:<br/>
`$ go run main.go -i GOPR2630.csv`<br/>
finally a new file called _GOPR2630-ahrs.csv_ will be created<br/>
_(replace GOPR2630.csv with the name of your file)_<br/>

## Second method 
**_(this method does not require golang)_**
Inside any folder execute the command:<br/>
`$ git clone https://github.com/brunocannavina/goproahrs.git`<br/>
_(or you can simply download it from github)_<br/>
then, open the folder goproahrs:<br/>
`$ cd goproahrs`<br/>
paste the CSV file copied above<br/>
and execute the following command:<br/>
`$ ahrs.exe -i GOPR2630.csv`<br/>
finally a new file called _GOPR2630-ahrs.csv_ will be created<br/>
_(replace GOPR2630.csv with the name of your file)_<br/>
