# How to build
Navigate to the directory where you cloned and then run the following command in the terminal:

`go build`

# How to run
Once that the build process has completed, execute the program by running the command:

`./pokemon-api`

This will start the application in the portal 10000. To call the API, use Postman or Curl:

`curl --location --request GET 'http://localhost:10000/pokemons'`

 Response should be:
 
 
 ```
[
     {
         "Id": "1",
         "Name": "Pikachu",
         "Type": "Electric"
     },
     {
         "Id": "2",
         "Name": "Charmeleon",
         "Type": "Fire"
     }
 ]
```
