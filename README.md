HACHI
===================
Open Catalog & API for Mexican 10 Digit Public Numbers

![Hachi](https://raw.githubusercontent.com/mafairnet/hachi/master/resources/hachi_logo.png)

About
-----------
#### Objective: Provide a Catalog and API to obtain data of a 10 digit mexican number.

Some companies need to know if a certain number is a mobile number or a landline number. This catalog is ontained from the IFT (Telecomunications Federal Institute) containing the basic data of all the registered numbers at the current Telecomunications Companies. 

Importer
-----------
<!---
![Importer](http://www.maf.mx/astricon/2017/images/spectrogram_peaks.png)
-->

The importer is a Golang based program. It downloads the catalog to your server/computer and stores it into a relational mysql/mariadb database.

The number data is stored by town, township, state, number type and provider. With a simple SQL query you can get all the data you need for a certain 10 digit number.

Prerequisites
-----------
- Golang Enviroment and dependencies
- MySQL Server (5.1 or later)
- Windows, MacOS or Linux SO

Compilation
-----------
1. Install golang (further info about that at https://golang.org/doc/install)
2. Install golang dependencies
```
go get- u"github.com/go-sql-driver/mysql"
```
3. Build binaries
```
go build
```
4. Run your binary
```
//Windows
.\importer.exe
//Linux/MacOS
./importer
```