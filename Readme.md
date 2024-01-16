#### Go language folder structure for web app/server ####

1. Place all read me here or any other info needed here. for set up and development.

2. create a ".env" file where we will keep all environment related constants. like database username,password, secret key  and mail info.

3. Create a "main.go" file in root, this will serve as the project/web-server entry point.

4. Create a "Makefile" this will contain some basic commands to run the project.

5. Create these folders:
   1. Controllers: contains a controller that accepts a request and calls a particular service to process it.

   2. Services: contains service that has the logic all the process like manipulating the DB and giving back the desired result.

   3. Models: contains the struct for storing data in DB OR fetching data from DB.

   4. DB: keep you DB related operations here. these can be used in services to fetch/update/insert/delete data from DB.
         *** Place DB connection file as well inside this folder.
   
   5. Routes: Keep all your routes here and pass the name of controller.

   6. Config: keep all your configuration things here like. fetching variables from .env file or even db config can be place here.

   7. Constants: keep all constants here, can also categorise them in separate files.

   8. Utils: Keeps all commonly used functions here like capitalising first Latter, etc.

We are now ready to start coding with this folder structure.