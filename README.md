HACHI
===================
Open Catalog & API for Mexican 10 Digit Public Numbers

<p align="center">
  <img width="400px" src="https://raw.githubusercontent.com/mafairnet/hachi/master/resources/hachi_logo.png">
</p>

<!---
![Hachi](https://raw.githubusercontent.com/mafairnet/hachi/master/resources/hachi_logo.png)
-->

About
-----------
#### Objective: Provide a Catalog and API to obtain data of a 10 digit mexican number.

Some companies need to know if a certain number is a mobile number or a landline number. This catalog is ontained from the IFT (Telecomunications Federal Institute) containing the basic data of all the registered numbers at the current Telecomunications Companies. 

System Architecture
-----------
The project includes an importer tool to download the data to your computer/server the store into a relational DB and a HTTP Server that runs the API that provides the data of the 10 digit mexican numbers.

<p align="center">
  <img width="700px" src="https://raw.githubusercontent.com/mafairnet/hachi/master/resources/system_architecture.png">
</p>

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
go get- u "github.com/go-sql-driver/mysql"
```
3. Build binaries
```
go build
```
4. Edit the "config/config.development.json". Set the mysql/mariadb server IP, username and password.
```
{
    "IftCatalogURL" : "https://github.com/mafairnet/hachi/blob/master/ift/",
    "DbServer" : "YOUR_SERVER_IP",
    "DbUsername" : "YOUR_DB_USERNAME",
    "DbPassword" : "YOUR_DB_PASSWORD",
    "DbPort" : "3306",
    "DbSchema" : "hachi"
}
```
5. Create the database from the file "database/model.sql"
```
>mysql -u root -p 
>create database hachi;
>exit
>mysql -u username -p hachi < database/model.sql
```
6. Add this Store Procedure
```
DELIMITER //
 CREATE PROCEDURE CleanDb()
   BEGIN
   SET FOREIGN_KEY_CHECKS = 0;
   truncate table state;
   truncate table township;
   truncate table town;
   truncate table provider;
   truncate table number_type;
   truncate table number;
   SET FOREIGN_KEY_CHECKS = 1;
   END //
DELIMITER ;
```
5. Run your binary
```
//Windows
.\importer.exe
//Linux/MacOS
./importer
```

API - HTTP Server
-----------
<!---
![Importer](http://www.maf.mx/astricon/2017/images/spectrogram_peaks.png)
-->

The API HTTP Server is a Golang based program. It retreives the data of the 10 digit Mexican number from the relational mysql/mariadb database that was previously populated by your importer.

It returns a json object containing the prefix, series, initial numeration, final numeration, type, provider, town, township and state where the phone number belongs.

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
go get- u "github.com/go-sql-driver/mysql"
go get- u "github.com/gorilla/mux"
```
3. Build binaries
```
go build
```
4. Edit the "config/config.development.json". Set the mysql/mariadb server IP, username and password.
```
{
    "DbServer" : "YOUR_SERVER_IP",
    "DbUsername" : "YOUR_DB_USERNAME",
    "DbPassword" : "YOUR_DB_PASSWORD",
    "DbPort" : "3306",
    "DbSchema" : "hachi"
}
```
5. Run your binary
```
//Windows
.\http_server.exe
//Linux/MacOS
./http_server
```
6. Request data through the API endpoint
```
http://SERVER_IP:8080/number/10_DIGIT_NUMBER
``` 
The endpoint will return a JSON data
```
{
  "id_number": 102930,
  "prefix": 998,
  "series": 123,
  "initial_numeration": 0,
  "final_numeration": 9999,
  "provider": {
    "id_provider": 58,
    "description": "RADIOMOVIL DIPSA, S.A. DE C.V."
  },
  "number_type": {
    "id_number_type": 2,
    "description": "MOVIL"
  },
  "town": {
    "id_town": 824,
    "description": "CANCUN",
    "township": {
      "id_township": 1456,
      "description": "BENITO JUAREZ",
      "state": {
        "id_state": 23,
        "description": "QROO"
      }
    }
  }
}
```