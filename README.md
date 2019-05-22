# GoproAHRS
Convert imu metadata from Gopro Hero 5 to AHRS using the .csv file from DashWare

## How to Use
- You must have Golang installed<br/>
- You must have my [goahrs](https://github.com/brunocannavina/goahrs) package installed<br/>
- You must have [DashWare](http://www.dashware.net/) installed<br/>
- And of course a video of Gopro Hero 5 or later<br/>

Inside any folder execute the command:<br/>
`$ git clone https://github.com/brunocannavina/goproahrs.git`<br/>
_(or you can simply download it from github)_<br/>
then, open the folder goproahrs:<br/>
`$ cd goproahrs`<br/>
and execute the following command:<br/>
`$ go run main.go -i `<br/>
