CREATE TABLE UserData(
  ID serial PRIMARY KEY,
  Name varchar(32) not null,
  Password varchar(32) not null,
  Email varchar(64) unique not null,
  Address char(64) not null,
  PhoneNumber varchar(16) not null,
  Birthday DATE,
  Authority int not null
);