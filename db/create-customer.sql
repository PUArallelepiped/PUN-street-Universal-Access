CREATE TABLE IF NOT EXISTS UserData(
  ID int unique not null,
  Name VARCHAR(32) not null,
  Password VARCHAR(32) not null,
  Email VARCHAR(64) unique not null,
  Address VARCHAR(64) not null,
  PhoneNumber VARCHAR(16) not null,
  Birthday timestamp,
  Authority int not null
);